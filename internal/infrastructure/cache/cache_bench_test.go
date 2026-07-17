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
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkSet(b *testing.B) {
	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()
			key := "bench_key"
			val := testData{Name: "Bench", Value: 42}
			ttl := 10 * time.Minute

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = cache.Set(ctx, key, val, ttl)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()
			key := "bench_key"
			val := testData{Name: "Bench", Value: 42}
			_ = cache.Set(ctx, key, val, 0)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = cache.Get(ctx, key)
			}
		})
	}
}

func BenchmarkMSet(b *testing.B) {
	const numKeys = 1000
	keys := make([]string, numKeys)
	items := make(map[string]testData, numKeys)
	for i := 0; i < numKeys; i++ {
		k := fmt.Sprintf("key_%d", i)
		keys[i] = k
		items[k] = testData{Name: "Bench", Value: i}
	}

	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = cache.MSet(ctx, items, 0)
			}
		})
	}
}

func BenchmarkMGet(b *testing.B) {
	const numKeys = 1000
	keys := make([]string, numKeys)
	items := make(map[string]testData, numKeys)
	for i := 0; i < numKeys; i++ {
		k := fmt.Sprintf("key_%d", i)
		keys[i] = k
		items[k] = testData{Name: "Bench", Value: i}
	}

	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()
			_ = cache.MSet(ctx, items, 0)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = cache.MGet(ctx, keys...)
			}
		})
	}
}

func BenchmarkSetNX(b *testing.B) {
	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()
			key := "nx_bench"
			val := testData{Name: "Bench", Value: 42}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// For SetNX to succeed, we need to delete or use a unique key each iteration.
				// Using a unique key per iteration avoids conflicts and measures SetNX performance.
				uniqueKey := fmt.Sprintf("%s_%d", key, i)
				_, _ = cache.SetNX(ctx, uniqueKey, val, 0)
			}
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()
			key := "delete_bench"
			val := testData{Name: "Bench", Value: 42}
			// Pre‑set the key so each iteration deletes an existing key.
			_ = cache.Set(ctx, key, val, 0)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = cache.Delete(ctx, key)
				// Re‑set for next iteration
				_ = cache.Set(ctx, key, val, 0)
			}
		})
	}
}

func BenchmarkConcurrent(b *testing.B) {
	benchmarks := []struct {
		name    string
		factory cacheFactory
	}{
		{"InMemory", newInMemoryTestCache},
		{"Redis", newRedisTestCache},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			cache, cleanup := bm.factory(b)
			defer cleanup()
			ctx := context.Background()

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				// Each goroutine works with its own key to avoid contention.
				// Use the goroutine ID or a random suffix.
				// The standard approach: use a counter from the test's state, but we can use b.N to generate unique keys.
				// However, b.RunParallel doesn't give a per‑goroutine ID. We can use atomic.AddUint64, but for simplicity
				// we can let each operation use a random key from a pool pre‑generated.
				// To keep it simple, we use a per‑iteration key based on the loop counter.
				// Since b.N is the total number of operations across all goroutines, we can use a global counter.
				// Instead, we'll generate a key from a global counter using atomic.
				var counter uint64
				for pb.Next() {
					idx := atomic.AddUint64(&counter, 1)
					key := fmt.Sprintf("concurrent_%d", idx)
					val := testData{Name: "Bench", Value: int(idx)}
					_ = cache.Set(ctx, key, val, 0)
					_, _ = cache.Get(ctx, key)
					_ = cache.Delete(ctx, key)
				}
			})
		})
	}
}
