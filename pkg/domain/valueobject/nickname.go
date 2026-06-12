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

var nicknamePattern = regexp.MustCompile(`^[\p{L}\p{N}_ ]{3,30}$`)

// Nickname represents a validated nickname.
type Nickname struct {
	value string
}

var _ ValueObject = Nickname{}

// NewNickname creates a Nickname after validation.
func NewNickname(nickname string) (Nickname, error) {

	sanitized := strings.TrimSpace(nickname)

	if !nicknamePattern.MatchString(sanitized) {
		return Nickname{}, fmt.Errorf("nickname must be between 3 and 30 characters and only contain language characters, underscores, spaces and numbers %q", nickname)
	}
	return Nickname{value: sanitized}, nil
}

// String returns the username string.
func (n Nickname) String() string {
	return n.value
}

// Equals returns true if the other object is a Nickname with the same value.
func (n Nickname) Equals(other any) bool {
	otherNick, ok := other.(Nickname)
	if !ok {
		return false
	}
	return n.value == otherNick.value
}

// IsValid returns true because the constructor guarantees validity.
func (n Nickname) IsValid() bool {
	return true
}

// IsZero returns true if the Nickname is the zero value (empty string).
func (n Nickname) IsZero() bool {
	return n.value == ""
}
