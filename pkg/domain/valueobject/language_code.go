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
	"regexp"
	"strings"
)

var languageCodePattern = regexp.MustCompile(`^[a-z]{2}$`)

// LanguageCode represents a validated language code.
type LanguageCode struct {
	value string
}

var _ ValueObject = LanguageCode{}

// NewLanguageCode creates a LanguageCode after validation.
func NewLanguageCode(languageCode string) (LanguageCode, error) {

	sanitized := strings.ToLower(languageCode)

	if !languageCodePattern.MatchString(sanitized) {
		return LanguageCode{}, fmt.Errorf("%q is not a valid language code", sanitized)
	}
	return LanguageCode{value: sanitized}, nil
}

// String returns the language code string.
func (l LanguageCode) String() string {
	return l.value
}

// Equals returns true if the other object is a LanguageCode with the same value.
func (l LanguageCode) Equals(other any) bool {
	otherCode, ok := other.(LanguageCode)
	if !ok {
		return false
	}
	return l.value == otherCode.value
}

// IsValid returns true because the constructor guarantees validity.
func (l LanguageCode) IsValid() bool {
	return true
}

// IsZero returns true if the LanguageCode is the zero value (empty string).
func (l LanguageCode) IsZero() bool {
	return l.value == ""
}
