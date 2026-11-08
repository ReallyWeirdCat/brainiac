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
	"io"
	"time"
)

// ObjectStore is the minimal object-store capability: put, get, delete.
// A local filesystem backend implements this fully; an S3/Garage backend can
// implement the same contract later without changing consumers.
type ObjectStore interface {
	// Put writes a file and returns its storage key.
	Put(ctx context.Context, key string, r io.Reader, size int64, contentType string) (string, error)

	// Get opens a file for reading. The caller must close it.
	Get(ctx context.Context, key string) (io.ReadCloser, error)

	// Delete removes a file. Deleting a missing key is a no-op.
	Delete(ctx context.Context, key string) error
}

// UrlSigner is an optional capability for backends that can produce URLs
// (public or presigned). Backends that cannot (e.g. local disk) do not
// implement it; consumers should type-assert for it.
type UrlSigner interface {
	PublicURL(ctx context.Context, key string) (string, error)
	SignedURL(ctx context.Context, key string, expiry time.Duration) (string, error)
	UploadURL(ctx context.Context, key, contentType string, expiry time.Duration) (string, error)
}

// Storage is a backwards-compatible alias for ObjectStore.
type Storage = ObjectStore
