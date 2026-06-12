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
	"regexp"
	"strings"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

var urlPattern = regexp.MustCompile(`\b(?:https?|ftp|file|sftp|ws|wss)://`)

// Bio represents a validated bio.
type Bio struct {
	value string
}

var _ ValueObject = Bio{}

// NewBio creates a Bio after validation.
func NewBio(bio string) (Bio, error) {

	sanitized := strings.TrimSpace(bio)

	// Check length
	if len(sanitized) < 1 || len(sanitized) > 175 {
		return Bio{}, &errors.ErrInvalidBio
	}

	// Check for @ symbol
	if strings.Contains(sanitized, "@") {
		return Bio{}, &errors.ErrInvalidBio
	}

	if urlPattern.MatchString(sanitized) {
		return Bio{}, &errors.ErrInvalidBio
	}
	return Bio{value: sanitized}, nil
}

// String returns the bio string.
func (b Bio) String() string {
	return b.value
}

// Equals returns true if the other object is a Bio with the same value.
func (b Bio) Equals(other any) bool {
	otherBio, ok := other.(Bio)
	if !ok {
		return false
	}
	return b.value == otherBio.value
}

// IsValid returns true because the constructor guarantees validity.
func (b Bio) IsValid() bool {
	return true
}

// IsZero returns true if the Bio is the zero value (empty string).
func (b Bio) IsZero() bool {
	return b.value == ""
}
