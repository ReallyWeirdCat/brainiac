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
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/redis/go-redis/v9"
)

type redisCache[T any] struct {
	client *redis.Client
	prefix string
}

func NewRedisCache[T any](client *redis.Client, prefix string) ports.Cache[T] {
	return &redisCache[T]{client: client, prefix: prefix}
}

func (r *redisCache[T]) prefixKey(key string) string {
	return r.prefix + key
}

// Delete removes a single key. Missing key is a no-op.
func (r *redisCache[T]) Delete(ctx context.Context, key string) error {
	cmd := r.client.Del(ctx, r.prefixKey(key))
	if err := cmd.Err(); err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}

func (r *redisCache[T]) Exists(ctx context.Context, key string) (bool, error) {
	cmd := r.client.Exists(ctx, r.prefixKey(key))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Get retrieves a value by key. Returns (nil, nil) if key does not exist.
func (r *redisCache[T]) Get(ctx context.Context, key string) (*T, error) {
	cmd := r.client.Get(ctx, r.prefixKey(key))
	data, err := cmd.Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}
	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return nil, err
	}
	return &val, nil
}

// MDelete deletes multiple keys. Missing keys are ignored.
func (r *redisCache[T]) MDelete(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	prefixed := make([]string, len(keys))
	for i, k := range keys {
		prefixed[i] = r.prefixKey(k)
	}
	cmd := r.client.Del(ctx, prefixed...)
	return cmd.Err()
}

func (r *redisCache[T]) MGet(ctx context.Context, keys ...string) (map[string]*T, error) {
	if len(keys) == 0 {
		return map[string]*T{}, nil
	}
	prefixed := make([]string, len(keys))
	for i, k := range keys {
		prefixed[i] = r.prefixKey(k)
	}
	cmd := r.client.MGet(ctx, prefixed...)
	vals, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	res := make(map[string]*T, len(keys))
	var errs []error

	for i, val := range vals {
		key := keys[i]

		// Key does not exist
		if val == nil {
			res[key] = nil
			continue
		}

		b, ok := val.(string)
		if !ok {
			errs = append(errs, fmt.Errorf("key %q: unexpected Redis type %T", key, val))
			res[key] = nil
			continue
		}

		var t T
		if err := json.Unmarshal([]byte(b), &t); err != nil {
			// Collect error, but keep result as nil for this key
			errs = append(errs, fmt.Errorf("key %q: %w", key, err))
			res[key] = nil
			continue
		}
		res[key] = &t
	}

	if len(errs) == 0 {
		return res, nil
	}
	// Return partial results along with aggregated errors
	return res, errors.Join(errs...)
}

// MSet sets multiple key-value pairs with the same TTL using a pipeline.
func (r *redisCache[T]) MSet(ctx context.Context, items map[string]T, ttl time.Duration) error {
	if len(items) == 0 {
		return nil
	}
	pipe := r.client.Pipeline()
	for key, val := range items {
		data, err := json.Marshal(val)
		if err != nil {
			return err
		}
		pipe.Set(ctx, r.prefixKey(key), data, ttl)
	}
	_, err := pipe.Exec(ctx)
	return err
}

// Set stores a value with a TTL.
func (r *redisCache[T]) Set(ctx context.Context, key string, value T, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, r.prefixKey(key), data, ttl).Err()
}

// SetNX atomically sets a value only if the key does not exist.
// Returns true if the key was set, false if it already existed.
func (r *redisCache[T]) SetNX(ctx context.Context, key string, value T, ttl time.Duration) (bool, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return false, err
	}
	return r.client.SetNX(ctx, r.prefixKey(key), data, ttl).Result()
}

// SetTTL updates the TTL of an existing key. Returns error if key does not exist.
func (r *redisCache[T]) SetTTL(ctx context.Context, key string, ttl time.Duration) error {
	cmd := r.client.Expire(ctx, r.prefixKey(key), ttl)
	ok, err := cmd.Result()
	if err != nil {
		return err
	}
	if !ok {
		return redis.Nil // key does not exist
	}
	return nil
}

// TTL returns the remaining time-to-live for a key.
// Nil key causes redis.Nil error.
func (r *redisCache[T]) TTL(ctx context.Context, key string) (time.Duration, error) {
	cmd := r.client.TTL(ctx, r.prefixKey(key))
	result, err := cmd.Result()
	if err != nil {
		return result, err
	}

	if result == -2*time.Nanosecond {
		return result, redis.Nil

	}
	return result, nil
}

var _ ports.Cache[any] = &redisCache[any]{}
