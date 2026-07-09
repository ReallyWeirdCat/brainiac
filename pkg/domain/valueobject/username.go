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

// usernamePattern: only latin letters (a-z, A-Z), digits, underscores; length 3-18.
var usernamePattern = regexp.MustCompile(`^[a-zA-Z0-9_]{3,18}$`)
var ErrInvalidUsername = domerr.NewDomainError("invalid username format", nil).WithType(domerr.Validation)

// Username represents a validated username.
type Username string

var _ ValueObject = Username("")

func NewUsername(name string) (Username, error) {
	if !usernamePattern.MatchString(name) {
		return Username(""), ErrInvalidUsername
	}
	return Username(name), nil
}

func (u Username) String() string {
	return string(u)
}

func (u Username) Equals(other any) bool {
	otherUser, ok := other.(Username)
	if !ok {
		return false
	}
	return string(u) == string(otherUser)
}

func (u Username) Validate() error {
	_, err := NewUsername(string(u))
	return err
}

func (u Username) IsValid() bool {
	return u.Validate() == nil
}

func (u Username) IsZero() bool {
	return string(u) == ""
}
