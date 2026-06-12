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

// Username represents a validated username.
type Username struct {
	value string
}

// NewUsername creates a Username after validation.
func NewUsername(name string) (Username, error) {
	if !usernamePattern.MatchString(name) {
		return Username{}, &domerr.ErrInvalidUsername
	}
	return Username{value: name}, nil
}

// String returns the username string.
func (u Username) String() string {
	return u.value
}
