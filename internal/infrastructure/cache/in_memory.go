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
	"errors"
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/patrickmn/go-cache"
)

// cacheItem wraps a value with its expiration timestamp so that remaining TTL
// can be queried later.
type cacheItem[T any] struct {
	Value     T
	ExpiresAt time.Time // zero value means no expiration
}

type inMemoryCache[T any] struct {
	c *cache.Cache
}

// NewInMemoryCache creates an in-memory cache backed by patrickmn/go-cache.
// The default expiration is NoExpiration (items live forever unless a TTL is
// given), and expired items are purged every 10 minutes.
func NewInMemoryCache[T any]() ports.Cache[T] {
	return &inMemoryCache[T]{
		c: cache.New(cache.NoExpiration, 10*time.Minute),
	}
}

// Delete removes a single key. Missing key is a no-op.
func (i *inMemoryCache[T]) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	i.c.Delete(key)
	return nil
}

// Exists reports whether a key is present (including non-expired items).
func (i *inMemoryCache[T]) Exists(ctx context.Context, key string) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	_, found := i.c.Get(key)
	return found, nil
}

// Get retrieves a value by key. Returns (nil, nil) if the key does not exist.
func (i *inMemoryCache[T]) Get(ctx context.Context, key string) (*T, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	raw, found := i.c.Get(key)
	if !found {
		return nil, nil
	}
	item, ok := raw.(cacheItem[T])
	if !ok {
		return nil, errors.New("cache: unexpected type stored")
	}
	return &item.Value, nil
}

// MDelete deletes multiple keys. Missing keys are ignored.
func (i *inMemoryCache[T]) MDelete(ctx context.Context, keys ...string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	for _, k := range keys {
		i.c.Delete(k)
	}
	return nil
}

// MGet retrieves multiple keys. Missing keys appear as nil in the map.
func (i *inMemoryCache[T]) MGet(ctx context.Context, keys ...string) (map[string]*T, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	res := make(map[string]*T, len(keys))
	for _, k := range keys {
		v, err := i.Get(ctx, k)
		if err != nil {
			return nil, err
		}
		res[k] = v
	}
	return res, nil
}

// MSet sets multiple key-value pairs with the same TTL.
func (i *inMemoryCache[T]) MSet(ctx context.Context, items map[string]T, ttl time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	for key, val := range items {
		if err := i.Set(ctx, key, val, ttl); err != nil {
			return err
		}
	}
	return nil
}

// Set stores a value with a TTL. A TTL of 0 means no expiration.
func (i *inMemoryCache[T]) Set(ctx context.Context, key string, value T, ttl time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	var expiresAt time.Time
	d := ttl
	if ttl == 0 {
		d = cache.NoExpiration
	} else {
		expiresAt = time.Now().Add(ttl)
	}
	i.c.Set(key, cacheItem[T]{Value: value, ExpiresAt: expiresAt}, d)
	return nil
}

// SetNX atomically sets a value only if the key does not exist.
// Returns true if the key was set, false if it already existed.
func (i *inMemoryCache[T]) SetNX(ctx context.Context, key string, value T, ttl time.Duration) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}
	var expiresAt time.Time
	d := ttl
	if ttl == 0 {
		d = cache.NoExpiration
	} else {
		expiresAt = time.Now().Add(ttl)
	}
	err := i.c.Add(key, cacheItem[T]{Value: value, ExpiresAt: expiresAt}, d)
	if err != nil {
		// go-cache.Add returns an error when the item already exists.
		return false, nil
	}
	return true, nil
}

// SetTTL updates the TTL of an existing key. Returns an error if the key does
// not exist.
func (i *inMemoryCache[T]) SetTTL(ctx context.Context, key string, ttl time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	raw, found := i.c.Get(key)
	if !found {
		return errors.New("cache: key not found")
	}
	item, ok := raw.(cacheItem[T])
	if !ok {
		return errors.New("cache: unexpected type stored")
	}
	var expiresAt time.Time
	d := ttl
	if ttl == 0 {
		d = cache.NoExpiration
	} else {
		expiresAt = time.Now().Add(ttl)
	}
	item.ExpiresAt = expiresAt
	i.c.Set(key, item, d)
	return nil
}

// TTL returns the remaining time-to-live for a key.
//   - If the key does not exist, it returns (-2 * time.Nanosecond) and an error.
//   - If the key exists but has no expiration, it returns (-1 * time.Nanosecond).
//   - Otherwise it returns the positive duration until expiration.
func (i *inMemoryCache[T]) TTL(ctx context.Context, key string) (time.Duration, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	raw, found := i.c.Get(key)
	if !found {
		return -2 * time.Nanosecond, errors.New("cache: key not found")
	}
	item, ok := raw.(cacheItem[T])
	if !ok {
		return 0, errors.New("cache: unexpected type stored")
	}
	if item.ExpiresAt.IsZero() {
		return -1 * time.Nanosecond, nil // no expiration
	}
	remaining := time.Until(item.ExpiresAt)
	if remaining <= 0 {
		return -2 * time.Nanosecond, errors.New("cache: key expired")
	}
	return remaining, nil
}

var _ ports.Cache[any] = &inMemoryCache[any]{}
