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

type Storage interface {
	// Uploads a file to the specified path with the given content type.
	// Returns file identifier.
	Upload(ctx context.Context, path string, reader io.Reader, size int64, contentType string) (string, error)

	// Retrieves a file from storage.
	Download(ctx context.Context, path string) (io.ReadCloser, error)

	// Removes a file from storage.
	Delete(ctx context.Context, path string) error

	// Returns the public URL for a file.
	GetPublicURL(ctx context.Context, path string) (string, error)

	// Returns a presigned URL for a file with an expiration time.
	GetSignedURL(ctx context.Context, path string, expiry time.Duration) (string, error)

	// Generates a presigned URL for direct client uploads.
	GenerateUploadURL(ctx context.Context, path string, contentType string, expiry time.Duration) (string, error)
}
