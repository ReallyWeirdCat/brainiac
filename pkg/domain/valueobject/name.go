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

var namePattern = regexp.MustCompile(`^[\p{L}\p{M} -]{1,70}$`)
var ErrInvalidName = domerr.NewDomainError("invalid name format", nil).WithType(domerr.Validation)

// Name represents a validated name.
type Name string

var _ ValueObject = Name("")

func NewName(name string) (Name, error) {
	if !namePattern.MatchString(name) {
		return Name(""), ErrInvalidName
	}
	return Name(name), nil
}

func (n Name) String() string {
	return string(n)
}

func (n Name) Equals(other any) bool {
	otherName, ok := other.(Name)
	if !ok {
		return false
	}
	return string(n) == string(otherName)
}

func (n Name) IsValid() bool {
	_, err := NewName(string(n))
	return err == nil
}

func (n Name) IsZero() bool {
	return string(n) == ""
}
