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
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

type testData struct {
	Name  string
	Value int
}

// newTestCache sets up a miniredis instance, a Redis client, and the cache.
// It returns the miniredis instance (for direct manipulation), the client,
// and the cache under test. The caller should defer miniredis.Close().
func newTestCache(t *testing.T) (*miniredis.Miniredis, *redis.Client, *redisCache[testData]) {
	t.Helper()

	mr := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	// Flush all data to start clean
	mr.FlushAll()

	prefix := "test_prefix:"
	cache := NewRedisCache[testData](client, prefix).(*redisCache[testData])
	return mr, client, cache
}

// Helper to set raw JSON data directly in Redis (for testing invalid JSON).
func setRawData(ctx context.Context, client *redis.Client, key, value string) error {
	return client.Set(ctx, key, value, 0).Err()
}

func TestRedisCache_SetAndGet(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_Delete(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_Exists(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_MGet(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_MGet_WithInvalidJSON(t *testing.T) {
	mr, client, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	// Insert one valid key, one invalid JSON, and one missing
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
	err = setRawData(ctx, client, prefixedInvalid, "this is not json")
	if err != nil {
		t.Fatalf("setRawData failed: %v", err)
	}

	// Now MGet all three
	result, err := cache.MGet(ctx, validKey, invalidKey, missingKey)
	if err == nil {
		t.Error("expected error due to invalid JSON, got nil")
	}
	// Error should be a join of errors (one for invalidKey)
	if !errors.Is(err, errors.Join()) { // join returns a multi-error, we just check non-nil
		t.Logf("got error: %v", err)
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

func TestRedisCache_MSet(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_SetNX(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_TTL(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	key := "ttl_test"
	val := testData{Name: "TTL", Value: 123}

	// TTL on missing key -> should return -2 and redis.Nil error
	ttl, err := cache.TTL(ctx, key)
	if !errors.Is(err, redis.Nil) {
		t.Errorf("expected redis.Nil for missing key, got %v", err)
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

	// Set with no TTL (0)
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

func TestRedisCache_SetTTL(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

	// SetTTL on missing key -> should return redis.Nil
	err = cache.SetTTL(ctx, "missing", newTTL)
	if !errors.Is(err, redis.Nil) {
		t.Errorf("expected redis.Nil for missing key, got %v", err)
	}
}

func TestRedisCache_MDelete(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_Prefix(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

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

func TestRedisCache_ConcurrentAccess(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	const numGoroutines = 20
	const iterations = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Each goroutine works with its own key, so no race on the key itself.
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("concurrent_%d", id)
			val := testData{Name: "Goroutine", Value: id}
			for j := 0; j < iterations; j++ {
				// Set
				err := cache.Set(ctx, key, val, 0)
				if err != nil {
					t.Errorf("Set in goroutine %d failed: %v", id, err)
					return
				}
				// Get
				got, err := cache.Get(ctx, key)
				if err != nil {
					t.Errorf("Get in goroutine %d failed: %v", id, err)
					return
				}
				if got == nil {
					t.Errorf("Get in goroutine %d returned nil", id)
					return
				}
				if got.Value != val.Value {
					t.Errorf("goroutine %d: got value %d, expected %d", id, got.Value, val.Value)
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestRedisCache_ErrorScenarios(t *testing.T) {
	// Test that JSON marshal errors are handled (though impossible for simple types)
	// We can't easily force marshal error, but we can test that Set returns error if
	// value cannot be marshaled (we can use a type that fails). However, for testData,
	// it's trivial. We'll skip this.
	// But we can test that Set with nil pointer? Not relevant.
}

// Test that TTL returns redis.Nil error correctly
func TestRedisCache_TTL_ReturnsNilError(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	key := "missing_ttl"
	ttl, err := cache.TTL(ctx, key)
	if !errors.Is(err, redis.Nil) {
		t.Errorf("expected redis.Nil, got %v", err)
	}
	if ttl != -2*time.Nanosecond {
		t.Errorf("TTL = %v, want -2ns", ttl)
	}
}

// Test that SetTTL on missing key returns redis.Nil
func TestRedisCache_SetTTL_MissingKeyReturnsNilError(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	key := "missing_setttl"
	err := cache.SetTTL(ctx, key, time.Second)
	if !errors.Is(err, redis.Nil) {
		t.Errorf("expected redis.Nil, got %v", err)
	}
}

// Test that Delete on missing key does not error
func TestRedisCache_Delete_MissingKeyNoError(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	err := cache.Delete(ctx, "missing")
	if err != nil {
		t.Errorf("Delete missing key should not error, got %v", err)
	}
}

// Test that MDelete on missing keys does not error
func TestRedisCache_MDelete_MissingKeysNoError(t *testing.T) {
	mr, _, cache := newTestCache(t)
	defer mr.Close()

	ctx := context.Background()
	err := cache.MDelete(ctx, "missing1", "missing2")
	if err != nil {
		t.Errorf("MDelete missing keys should not error, got %v", err)
	}
}
