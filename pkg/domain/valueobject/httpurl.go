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
	"net/url"
	"strings"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

// HttpUrl represents a validated HTTP/HTTPS URL.
type HttpUrl string

var _ ValueObject = HttpUrl("")

// NewHttpUrl creates an HttpUrl after validation and trimming.
// If no scheme is provided, "https://" will be added automatically.
func NewHttpUrl(rawUrl string) (HttpUrl, error) {
	sanitized := strings.TrimSpace(rawUrl)

	if sanitized == "" {
		return HttpUrl(""), errors.ErrInvalidHttpUrl
	}

	// Check if URL has a scheme; if not, add https://
	if !strings.Contains(sanitized, "://") {
		sanitized = "https://" + sanitized
	}

	parsed, err := url.Parse(sanitized)
	if err != nil {
		return HttpUrl(""), errors.ErrInvalidHttpUrl
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return HttpUrl(""), errors.ErrInvalidHttpUrl
	}

	if parsed.Host == "" {
		return HttpUrl(""), errors.ErrInvalidHttpUrl
	}

	// Reconstruct the URL to ensure consistent formatting
	normalized := parsed.String()

	return HttpUrl(normalized), nil
}

func (h HttpUrl) String() string {
	return string(h)
}

func (h HttpUrl) Equals(other any) bool {
	otherUrl, ok := other.(HttpUrl)
	if !ok {
		return false
	}
	return string(h) == string(otherUrl)
}

func (h HttpUrl) IsValid() bool {
	_, err := NewHttpUrl(string(h))
	return err == nil
}

func (h HttpUrl) IsZero() bool {
	return string(h) == ""
}
