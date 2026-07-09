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

	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

var urlPattern = regexp.MustCompile(`\b(?:https?|ftp|file|sftp|ws|wss)://`)
var ErrInvalidBio = domerr.NewDomainError("invalid bio format", nil).WithType(domerr.Validation)

type Bio string

var _ ValueObject = Bio("")

// NewBio creates a Bio after validation.
func NewBio(bio string) (Bio, error) {
	sanitized := strings.TrimSpace(bio)

	// Check length
	if len(sanitized) < 1 || len(sanitized) > 175 {
		return Bio(""), ErrInvalidBio
	}

	// Check for @ symbol
	if strings.Contains(sanitized, "@") {
		return Bio(""), ErrInvalidBio
	}

	if urlPattern.MatchString(sanitized) {
		return Bio(""), ErrInvalidBio
	}
	return Bio(sanitized), nil
}

func (b Bio) String() string {
	return string(b)
}

func (b Bio) Equals(other any) bool {
	otherBio, ok := other.(Bio)
	if !ok {
		return false
	}
	return string(b) == string(otherBio)
}

func (b Bio) Validate() error {
	_, err := NewBio(string(b))
	return err
}

func (b Bio) IsValid() bool {
	return b.Validate() == nil
}

func (b Bio) IsZero() bool {
	return string(b) == ""
}
