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

package ports

import (
	"context"
	"time"
)

const (
	CacheDurationInfinite = time.Duration(-1) * time.Nanosecond
)

type Cache[T any] interface {
	// nil if object does not exist in cache
	Get(ctx context.Context, key string) (*T, error)

	Set(ctx context.Context, key string, value T, ttl time.Duration) error

	// SetNX sets a value only if the key does not exist (atomic lock).
	// Returns true if the lock was acquired, false if the key already exists.
	SetNX(ctx context.Context, key string, value T, ttl time.Duration) (bool, error)

	// Discards cache by key. Must not error when key is unset.
	Delete(ctx context.Context, key string) error

	Exists(ctx context.Context, key string) (bool, error)

	// Get time-to-live for key.
	TTL(ctx context.Context, key string) (time.Duration, error)

	// Set time-to-live for key.
	SetTTL(ctx context.Context, key string, ttl time.Duration) error

	// Returns a map of key -> *value (nil if missing)
	MGet(ctx context.Context, keys ...string) (map[string]*T, error)

	// Sets multiple keys atomically (or via pipeline)
	MSet(ctx context.Context, items map[string]T, ttl time.Duration) error

	// Delete multiple keys. Unexistent keys must be skipped.
	MDelete(ctx context.Context, keys ...string) error
}
