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
