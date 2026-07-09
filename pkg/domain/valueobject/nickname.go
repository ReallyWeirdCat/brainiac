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

var nicknamePattern = regexp.MustCompile(`^[\p{L}\p{N}_ ]{3,30}$`)
var ErrInvalidNickname = domerr.NewDomainError("invalid nickname format", nil).WithType(domerr.Validation)

// Nickname represents a validated nickname.
type Nickname string

var _ ValueObject = Nickname("")

// NewNickname creates a Nickname after validation.
func NewNickname(nickname string) (Nickname, error) {

	sanitized := strings.TrimSpace(nickname)

	if !nicknamePattern.MatchString(sanitized) {
		return Nickname(""), ErrInvalidNickname
	}
	return Nickname(sanitized), nil
}

func (n Nickname) String() string {
	return string(n)
}

func (n Nickname) Equals(other any) bool {
	otherNick, ok := other.(Nickname)
	if !ok {
		return false
	}
	return string(n) == string(otherNick)
}

func (n Nickname) Validate() error {
	_, err := NewNickname(string(n))
	return err
}

func (n Nickname) IsValid() bool {
	return n.Validate() == nil
}

func (n Nickname) IsZero() bool {
	return string(n) == ""
}
