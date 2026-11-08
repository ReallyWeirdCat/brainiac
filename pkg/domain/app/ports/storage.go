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
	// Upload stores content from reader at the given path.
	// Returns the file identifier used to download the file.
	Upload(ctx context.Context, path string, reader io.Reader, size int64, contentType string) (string, error)

	// Download retrieves the content stored at path.
	Download(ctx context.Context, id string) (io.ReadCloser, error)

	// Delete removes the content stored by id.
	Delete(ctx context.Context, id string) error
}

// URLProvider combines all URL generation capabilities.
type URLProvider interface {
	PublicURLProvider
	SignedURLProvider
	UploadURLProvider
}

// PublicURLProvider can generate a permanent public URL for a file.
type PublicURLProvider interface {
	GetPublicURL(ctx context.Context, path string) (string, error)
}

// SignedURLProvider can generate a temporary signed URL for downloading a file.
type SignedURLProvider interface {
	GetSignedURL(ctx context.Context, path string, expiry time.Duration) (string, error)
}

// UploadURLProvider can generate a temporary signed URL for direct client uploads.
type UploadURLProvider interface {
	GenerateUploadURL(ctx context.Context, path string, contentType string, expiry time.Duration) (string, error)
}
