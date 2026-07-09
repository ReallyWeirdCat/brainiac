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

var languageCodePattern = regexp.MustCompile(`^[a-z]{2}$`)

var ErrInvalidLanguageCode = domerr.NewDomainError("invalid language code format", nil).WithType(domerr.Validation)

// LanguageCode represents a validated language code.
type LanguageCode string

var _ ValueObject = LanguageCode("")

func NewLanguageCode(languageCode string) (LanguageCode, error) {

	sanitized := strings.ToLower(languageCode)

	if !languageCodePattern.MatchString(sanitized) {
		return LanguageCode(""), ErrInvalidLanguageCode
	}
	return LanguageCode(sanitized), nil
}

func (l LanguageCode) String() string {
	return string(l)
}

func (l LanguageCode) Equals(other any) bool {
	otherCode, ok := other.(LanguageCode)
	if !ok {
		return false
	}
	return string(l) == string(otherCode)
}

func (l LanguageCode) IsValid() bool {
	_, err := NewLanguageCode(string(l))
	return err == nil
}

func (l LanguageCode) IsZero() bool {
	return string(l) == ""
}
