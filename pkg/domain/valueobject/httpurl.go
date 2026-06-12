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

package valueobject

import (
	"fmt"
	"net/url"
	"strings"
)

// HttpUrl represents a validated HTTP/HTTPS URL.
type HttpUrl struct {
	value string
	url   url.URL
}

var _ ValueObject = HttpUrl{}

// NewHttpUrl creates an HttpUrl after validation and trimming.
// If no scheme is provided, "https://" will be added automatically.
func NewHttpUrl(rawUrl string) (HttpUrl, error) {
	sanitized := strings.TrimSpace(rawUrl)

	if sanitized == "" {
		return HttpUrl{}, fmt.Errorf("URL cannot be empty")
	}

	// Check if URL has a scheme; if not, add https://
	if !strings.Contains(sanitized, "://") {
		sanitized = "https://" + sanitized
	}

	parsed, err := url.Parse(sanitized)
	if err != nil {
		return HttpUrl{}, fmt.Errorf("invalid URL format: %q", sanitized)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return HttpUrl{}, fmt.Errorf("URL must have http or https scheme, got %q", parsed.Scheme)
	}

	if parsed.Host == "" {
		return HttpUrl{}, fmt.Errorf("URL must contain a host, got %q", sanitized)
	}

	// Reconstruct the URL to ensure consistent formatting
	normalized := parsed.String()

	return HttpUrl{value: normalized, url: *parsed}, nil
}

// String returns the URL string.
func (h HttpUrl) String() string {
	return h.value
}

// Equals returns true if the other object is an HttpUrl with the same value.
func (h HttpUrl) Equals(other any) bool {
	otherUrl, ok := other.(HttpUrl)
	if !ok {
		return false
	}
	return h.value == otherUrl.value
}

// IsValid returns true because the constructor guarantees validity.
func (h HttpUrl) IsValid() bool {
	return true
}

// IsZero returns true if the HttpUrl is the zero value (empty URL).
func (h HttpUrl) IsZero() bool {
	return h.value == ""
}

func (h HttpUrl) Url() url.URL {
	return h.url
}
