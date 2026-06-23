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

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

type CompromisedPasswordChecker interface {
	IsCompromised(password string) bool
}

var (
	// allowedChars ensures only printable ASCII (space to ~).
	allowedChars = regexp.MustCompile(`^[\x20-\x7E]+$`)
	// hasDigit requires at least one ASCII digit.
	hasDigit = regexp.MustCompile(`[0-9]`)
	// hasLower requires at least one lowercase ASCII letter.
	hasLower = regexp.MustCompile(`[a-z]`)
	// hasUpper requires at least one uppercase ASCII letter.
	hasUpper = regexp.MustCompile(`[A-Z]`)
	// hasSymbol requires at least one non-alphanumeric printable ASCII character.
	hasSymbol = regexp.MustCompile(`[^a-zA-Z0-9]`)

	minPasswordLength = 8
	maxPasswordLength = 72
)

// Password represents a validated password.
type Password struct {
	value string
}

var _ ValueObject = Password{}

// NewPassword creates a Password after validation.
func NewPassword(password string, compromisedPasswordChecker CompromisedPasswordChecker) (Password, error) {

	if len(password) < minPasswordLength {
		return Password{}, errors.ErrWeakPassword
	}
	if len(password) > maxPasswordLength {
		return Password{}, errors.ErrPasswordTooLong
	}
	if !allowedChars.MatchString(password) {
		return Password{}, errors.ErrPasswordNotAcceptable
	}
	if !hasDigit.MatchString(password) {
		return Password{}, errors.ErrWeakPassword
	}
	if !hasLower.MatchString(password) {
		return Password{}, errors.ErrWeakPassword
	}
	if !hasUpper.MatchString(password) {
		return Password{}, errors.ErrWeakPassword
	}
	if !hasSymbol.MatchString(password) {
		return Password{}, errors.ErrWeakPassword
	}
	if compromisedPasswordChecker.IsCompromised(password) {
		return Password{}, errors.ErrCompromisedPassword
	}
	return Password{value: password}, nil
}

// String returns the password string.
func (p Password) String() string {
	return p.value
}

// Equals returns true if the other object is a Password with the same value.
func (p Password) Equals(other any) bool {
	otherPassword, ok := other.(Password)
	if !ok {
		return false
	}
	return p.value == otherPassword.value
}

// IsValid returns true because the constructor guarantees validity.
func (p Password) IsValid() bool {
	return true
}

// IsZero returns true if the Password is the zero value (empty string).
func (p Password) IsZero() bool {
	return p.value == ""
}
