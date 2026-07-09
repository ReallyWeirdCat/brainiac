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

	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

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

	ErrPasswordTooLong       = domerr.NewDomainError("password exceeds maximum length", nil).WithType(domerr.Validation)
	ErrPasswordNotAcceptable = domerr.NewDomainError("password contains forbidden characters", nil).WithType(domerr.Validation)
	ErrWeakPassword          = domerr.NewDomainError("password is too weak", nil).WithType(domerr.Validation)
	ErrCompromisedPassword   = domerr.NewDomainError("password is known to be compromised", nil).WithType(domerr.Validation)
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 72
)

// Password represents a validated password.
type Password string

var _ ValueObject = Password("")

func NewPassword(password string) (Password, error) {

	if len(password) < MinPasswordLength {
		return Password(""), ErrWeakPassword
	}
	if len(password) > MaxPasswordLength {
		return Password(""), ErrPasswordTooLong
	}
	if !allowedChars.MatchString(password) {
		return Password(""), ErrPasswordNotAcceptable
	}
	if !hasDigit.MatchString(password) {
		return Password(""), ErrWeakPassword
	}
	if !hasLower.MatchString(password) {
		return Password(""), ErrWeakPassword
	}
	if !hasUpper.MatchString(password) {
		return Password(""), ErrWeakPassword
	}
	if !hasSymbol.MatchString(password) {
		return Password(""), ErrWeakPassword
	}
	return Password(password), nil
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Equals(other any) bool {
	otherPassword, ok := other.(Password)
	if !ok {
		return false
	}
	return string(p) == string(otherPassword)
}

func (p Password) IsValid() bool {
	_, err := NewPassword(string(p))
	return err == nil
}

func (p Password) IsZero() bool {
	return string(p) == ""
}
