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
	"net/mail"
	"strings"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

type Email struct {
	value string
}

var _ ValueObject = (*Email)(nil)

func NewEmail(email string) (*Email, error) {

	sanitized := strings.TrimSpace(email)

	parsed, err := mail.ParseAddress(sanitized)

	if err != nil {
		return nil, &errors.ErrInvalidEmail
	}

	// Extract domain from parsed email
	parts := strings.Split(parsed.Address, "@")
	if len(parts) != 2 {
		return nil, &errors.ErrInvalidEmail
	}

	domain := parts[1]

	// Check if domain contains a dot
	if !strings.Contains(domain, ".") {
		return nil, &errors.ErrInvalidEmail
	}

	return &Email{value: sanitized}, nil
}

func (e Email) String() string {
	return e.value
}

// Equals returns true if the other object is an *Email with the same value.
// A nil *Email is only equal to another nil *Email.
func (e *Email) Equals(other any) bool {
	if e == nil {
		return other == nil
	}
	otherEmail, ok := other.(*Email)
	if !ok {
		return false
	}
	if otherEmail == nil {
		return false
	}
	return e.value == otherEmail.value
}

// IsValid returns true because the constructor guarantees validity.
func (e *Email) IsValid() bool {
	return e != nil
}

// IsZero returns true if the Email is nil or contains an empty value.
func (e *Email) IsZero() bool {
	return e == nil || e.value == ""
}
