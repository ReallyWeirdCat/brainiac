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

	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

var ErrInvalidEmail = domerr.NewDomainError("invalid email format", nil).WithType(domerr.Validation)

type Email string

var _ ValueObject = Email("")

func NewEmail(email string) (Email, error) {

	sanitized := strings.TrimSpace(email)

	parsed, err := mail.ParseAddress(sanitized)

	if err != nil {
		return Email(""), ErrInvalidEmail
	}

	// Extract domain from parsed email
	parts := strings.Split(parsed.Address, "@")
	if len(parts) != 2 {
		return Email(""), ErrInvalidEmail
	}

	domain := parts[1]

	// Check if domain contains a dot
	if !strings.Contains(domain, ".") {
		return Email(""), ErrInvalidEmail
	}

	return Email(sanitized), nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Equals(other any) bool {
	otherEmail, ok := other.(Email)
	if !ok {
		return false
	}
	return string(e) == string(otherEmail)
}

func (e Email) IsValid() bool {
	_, err := NewEmail(string(e))
	return err == nil
}

func (e Email) IsZero() bool {
	return string(e) == ""
}
