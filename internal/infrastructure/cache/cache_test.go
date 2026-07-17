// Copyright (c) 2026 Nikolai Papin
//
// This file is part of Brainiac gamification and education platform
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

type testData struct {
	Name  string
	Value int
}

// cacheFactory returns a new cache instance and a cleanup function.
type cacheFactory func(t testing.TB) (cache ports.Cache[testData], cleanup func())

// newInMemoryTestCache returns an in-memory cache for testing.
func newInMemoryTestCache(t testing.TB) (ports.Cache[testData], func()) {
	return NewInMemoryCache[testData](), func() {}
}

// newRedisTestCache returns a Redis-backed cache (miniredis) for testing.
func newRedisTestCache(t testing.TB) (ports.Cache[testData], func()) {
	t.Helper()
	mr := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	prefix := "test_prefix:"
	cache := NewRedisCache[testData](client, prefix)
	return cache, func() { mr.Close() }
}

// Helper to set raw JSON data directly in Redis (used only in Redis-specific tests).
func setRawData(ctx context.Context, client *redis.Client, key, value string) error {
	return client.Set(ctx, key, value, 0).Err()
}

// ---------------------------------------------------------------------------
// Common test functions that work with any ports.Cache[testData] implementation.
// Each function receives a factory to create the cache under test.
// ---------------------------------------------------------------------------

func testSetAndGet(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	key := "testkey"
	val := testData{Name: "Alice", Value: 42}
	ttl := 10 * time.Minute

	// Set
	err := cache.Set(ctx, key, val, ttl)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Get
	got, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if got == nil {
		t.Fatal("Get returned nil, expected value")
	}
	if got.Name != val.Name || got.Value != val.Value {
		t.Errorf("got %+v, want %+v", got, val)
	}

	// TTL should be set
	ttlRemaining, err := cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttlRemaining <= 0 {
		t.Errorf("TTL = %v, expected > 0", ttlRemaining)
	}

	// Get non-existent key
	missing, err := cache.Get(ctx, "missing")
	if err != nil {
		t.Fatalf("Get missing failed: %v", err)
	}
	if missing != nil {
		t.Errorf("expected nil for missing key, got %+v", missing)
	}
}

func testDelete(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	key := "delete_me"
	val := testData{Name: "Bob", Value: 100}

	// Set and delete
	err := cache.Set(ctx, key, val, 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}
	err = cache.Delete(ctx, key)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	exists, err := cache.Exists(ctx, key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists {
		t.Error("key still exists after Delete")
	}

	// Delete missing key (should be no-op)
	err = cache.Delete(ctx, "missing")
	if err != nil {
		t.Errorf("Delete on missing key should not error, got %v", err)
	}
}

func testExists(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	key := "exists_test"
	val := testData{Name: "Charlie", Value: 7}

	// Should be false initially
	exists, err := cache.Exists(ctx, key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists {
		t.Error("expected false for missing key")
	}

	// Set and check
	err = cache.Set(ctx, key, val, 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}
	exists, err = cache.Exists(ctx, key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !exists {
		t.Error("expected true for existing key")
	}
}

func testMGet(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	items := map[string]testData{
		"a": {Name: "A", Value: 1},
		"b": {Name: "B", Value: 2},
		"c": {Name: "C", Value: 3},
	}

	// Set multiple
	err := cache.MSet(ctx, items, 0)
	if err != nil {
		t.Fatalf("MSet failed: %v", err)
	}

	// MGet all
	keys := []string{"a", "b", "c", "missing"}
	result, err := cache.MGet(ctx, keys...)
	if err != nil {
		t.Fatalf("MGet failed: %v", err)
	}
	// Check all present keys
	for _, k := range []string{"a", "b", "c"} {
		got := result[k]
		if got == nil {
			t.Errorf("key %q is nil, expected value", k)
			continue
		}
		expected := items[k]
		if got.Name != expected.Name || got.Value != expected.Value {
			t.Errorf("key %q: got %+v, want %+v", k, got, expected)
		}
	}
	// Check missing key
	if result["missing"] != nil {
		t.Errorf("missing key should be nil, got %+v", result["missing"])
	}
}

func testMSet(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	items := map[string]testData{
		"x": {Name: "X", Value: 10},
		"y": {Name: "Y", Value: 20},
	}
	ttl := 5 * time.Second

	err := cache.MSet(ctx, items, ttl)
	if err != nil {
		t.Fatalf("MSet failed: %v", err)
	}

	// Verify each key exists and has correct TTL
	for key, expected := range items {
		got, err := cache.Get(ctx, key)
		if err != nil {
			t.Errorf("Get %q failed: %v", key, err)
			continue
		}
		if got == nil {
			t.Errorf("key %q missing", key)
			continue
		}
		if got.Name != expected.Name || got.Value != expected.Value {
			t.Errorf("key %q: got %+v, want %+v", key, got, expected)
		}
		ttlRemaining, err := cache.TTL(ctx, key)
		if err != nil {
			t.Errorf("TTL %q failed: %v", key, err)
		} else if ttlRemaining <= 0 {
			t.Errorf("TTL %q = %v, expected > 0", key, ttlRemaining)
		}
	}

	// MSet with empty map should succeed
	err = cache.MSet(ctx, map[string]testData{}, ttl)
	if err != nil {
		t.Errorf("MSet with empty map should not error, got %v", err)
	}
}

func testSetNX(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	key := "nx_test"
	val1 := testData{Name: "First", Value: 1}
	val2 := testData{Name: "Second", Value: 2}

	// SetNX when key does not exist -> should succeed
	ok, err := cache.SetNX(ctx, key, val1, 0)
	if err != nil {
		t.Fatalf("SetNX failed: %v", err)
	}
	if !ok {
		t.Error("SetNX returned false, expected true for new key")
	}
	got, _ := cache.Get(ctx, key)
	if got == nil || got.Name != val1.Name || got.Value != val1.Value {
		t.Errorf("after SetNX, got %+v, want %+v", got, val1)
	}

	// SetNX when key exists -> should fail
	ok, err = cache.SetNX(ctx, key, val2, 0)
	if err != nil {
		t.Fatalf("SetNX failed: %v", err)
	}
	if ok {
		t.Error("SetNX returned true, expected false for existing key")
	}
	// Value should remain unchanged
	got, _ = cache.Get(ctx, key)
	if got == nil || got.Name != val1.Name || got.Value != val1.Value {
		t.Errorf("after second SetNX, got %+v, want %+v", got, val1)
	}
}

func testTTL(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	key := "ttl_test"
	val := testData{Name: "TTL", Value: 123}

	// TTL on missing key -> should return -2 and an error
	ttl, err := cache.TTL(ctx, key)
	if err == nil {
		t.Error("expected error for missing key, got nil")
	}
	if ttl != -2*time.Nanosecond {
		t.Errorf("TTL for missing = %v, want -2ns", ttl)
	}

	// Set with TTL
	err = cache.Set(ctx, key, val, 1*time.Hour)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}
	ttl, err = cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttl <= 0 {
		t.Errorf("TTL = %v, expected > 0", ttl)
	}

	// Set with no TTL (0) -> key never expires
	err = cache.Set(ctx, "no_ttl", val, 0)
	if err != nil {
		t.Fatalf("Set no TTL failed: %v", err)
	}
	ttl, err = cache.TTL(ctx, "no_ttl")
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttl != -1*time.Nanosecond {
		t.Errorf("TTL for no expiry = %v, want -1ns", ttl)
	}
}

func testSetTTL(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	key := "setttl_test"
	val := testData{Name: "TTL", Value: 99}

	// Set without TTL
	err := cache.Set(ctx, key, val, 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}
	// Update TTL
	newTTL := 30 * time.Minute
	err = cache.SetTTL(ctx, key, newTTL)
	if err != nil {
		t.Fatalf("SetTTL failed: %v", err)
	}
	ttl, err := cache.TTL(ctx, key)
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttl <= 0 || ttl > newTTL+time.Second {
		t.Errorf("TTL = %v, expected around %v", ttl, newTTL)
	}

	// SetTTL on missing key -> should return error
	err = cache.SetTTL(ctx, "missing", newTTL)
	if err == nil {
		t.Error("expected error for missing key, got nil")
	}
}

func testMDelete(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	// Set a few keys
	items := map[string]testData{
		"a": {Name: "A", Value: 1},
		"b": {Name: "B", Value: 2},
		"c": {Name: "C", Value: 3},
	}
	err := cache.MSet(ctx, items, 0)
	if err != nil {
		t.Fatalf("MSet failed: %v", err)
	}

	// Delete two keys, one missing
	err = cache.MDelete(ctx, "a", "c", "missing")
	if err != nil {
		t.Fatalf("MDelete failed: %v", err)
	}

	// Verify only b remains
	for k := range items {
		exists, err := cache.Exists(ctx, k)
		if err != nil {
			t.Errorf("Exists %q failed: %v", k, err)
		}
		if k == "a" || k == "c" {
			if exists {
				t.Errorf("key %q still exists after MDelete", k)
			}
		} else if k == "b" {
			if !exists {
				t.Errorf("key %q should exist, but doesn't", k)
			}
		}
	}

	// MDelete with empty list -> no-op
	err = cache.MDelete(ctx)
	if err != nil {
		t.Errorf("MDelete with no keys should not error, got %v", err)
	}
}

func testConcurrentAccess(t *testing.T, factory cacheFactory) {
	cache, cleanup := factory(t)
	defer cleanup()

	ctx := context.Background()
	const numGoroutines = 20
	const iterations = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("concurrent_%d", id)
			val := testData{Name: "Goroutine", Value: id}
			for j := 0; j < iterations; j++ {
				err := cache.Set(ctx, key, val, 0)
				if err != nil {
					t.Errorf("Set in goroutine %d failed: %v", id, err)
					return
				}
				got, err := cache.Get(ctx, key)
				if err != nil {
					t.Errorf("Get in goroutine %d failed: %v", id, err)
					return
				}
				if got == nil || got.Value != val.Value {
					t.Errorf("goroutine %d: got value %v, expected %d", id, got, val.Value)
				}
			}
		}(i)
	}
	wg.Wait()
}

// ---------------------------------------------------------------------------
// Redis-specific tests that cannot be run against the in-memory implementation.
// ---------------------------------------------------------------------------

// testRedisPrefix checks that the Redis cache properly prefixes keys.
func testRedisPrefix(t *testing.T, factory cacheFactory) {
	// This test is only applicable for Redis.
	cacheIf, cleanup := factory(t)
	defer cleanup()
	// We know this factory returns a *redisCache, but we need the concrete type
	// to access prefixKey and client. The test will only be invoked for Redis.
	cache, ok := cacheIf.(*redisCache[testData])
	if !ok {
		t.Fatal("expected *redisCache for prefix test")
	}

	ctx := context.Background()
	key := "prefixed"
	val := testData{Name: "PrefixTest", Value: 77}

	err := cache.Set(ctx, key, val, 0)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Directly check Redis for the prefixed key
	prefixed := cache.prefixKey(key)
	cmd := cache.client.Get(ctx, prefixed)
	data, err := cmd.Bytes()
	if err != nil {
		t.Fatalf("raw Get failed: %v", err)
	}
	var raw testData
	err = json.Unmarshal(data, &raw)
	if err != nil {
		t.Fatalf("unmarshal raw data failed: %v", err)
	}
	if raw != val {
		t.Errorf("raw value = %+v, want %+v", raw, val)
	}
}

// testRedisMGetWithInvalidJSON verifies that MGet returns an error when a key
// contains invalid JSON, while still returning partial results for valid keys.
func testRedisMGetWithInvalidJSON(t *testing.T, factory cacheFactory) {
	cacheIf, cleanup := factory(t)
	defer cleanup()
	cache, ok := cacheIf.(*redisCache[testData])
	if !ok {
		t.Fatal("expected *redisCache for invalid JSON test")
	}

	ctx := context.Background()
	validKey := "valid"
	invalidKey := "invalid"
	missingKey := "missing"

	validData := testData{Name: "Valid", Value: 99}
	// Set valid via cache
	err := cache.Set(ctx, validKey, validData, 0)
	if err != nil {
		t.Fatalf("Set valid failed: %v", err)
	}
	// Insert invalid JSON directly (note the prefix)
	prefixedInvalid := cache.prefixKey(invalidKey)
	err = setRawData(ctx, cache.client, prefixedInvalid, "this is not json")
	if err != nil {
		t.Fatalf("setRawData failed: %v", err)
	}

	// Now MGet all three
	result, err := cache.MGet(ctx, validKey, invalidKey, missingKey)
	if err == nil {
		t.Error("expected error due to invalid JSON, got nil")
	}
	// Check partial results
	if result[validKey] == nil {
		t.Error("validKey result is nil, expected value")
	} else {
		got := result[validKey]
		if got.Name != validData.Name || got.Value != validData.Value {
			t.Errorf("validKey: got %+v, want %+v", got, validData)
		}
	}
	if result[invalidKey] != nil {
		t.Errorf("invalidKey should be nil, got %+v", result[invalidKey])
	}
	if result[missingKey] != nil {
		t.Errorf("missingKey should be nil, got %+v", result[missingKey])
	}
}

// ---------------------------------------------------------------------------
// Main test runner – runs all common tests for both implementations,
// plus Redis‑specific tests only when the Redis factory is used.
// ---------------------------------------------------------------------------

func TestCache(t *testing.T) {
	factories := map[string]cacheFactory{
		"InMemory": newInMemoryTestCache,
		"Redis":    newRedisTestCache,
	}

	for impl, factory := range factories {
		t.Run(impl, func(t *testing.T) {
			t.Run("SetAndGet", func(t *testing.T) { testSetAndGet(t, factory) })
			t.Run("Delete", func(t *testing.T) { testDelete(t, factory) })
			t.Run("Exists", func(t *testing.T) { testExists(t, factory) })
			t.Run("MGet", func(t *testing.T) { testMGet(t, factory) })
			t.Run("MSet", func(t *testing.T) { testMSet(t, factory) })
			t.Run("SetNX", func(t *testing.T) { testSetNX(t, factory) })
			t.Run("TTL", func(t *testing.T) { testTTL(t, factory) })
			t.Run("SetTTL", func(t *testing.T) { testSetTTL(t, factory) })
			t.Run("MDelete", func(t *testing.T) { testMDelete(t, factory) })
			t.Run("ConcurrentAccess", func(t *testing.T) { testConcurrentAccess(t, factory) })

			// Redis‑specific tests
			if impl == "Redis" {
				t.Run("Prefix", func(t *testing.T) { testRedisPrefix(t, factory) })
				t.Run("MGetWithInvalidJSON", func(t *testing.T) { testRedisMGetWithInvalidJSON(t, factory) })
			}
		})
	}
}
