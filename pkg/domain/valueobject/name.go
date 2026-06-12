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
)

var namePattern = regexp.MustCompile(`^[\p{L}\p{M} -]{1,70}$`)

// Name represents a validated name.
type Name struct {
	value string
}

// NewName creates a Name after validation.
func NewName(name string) (Name, error) {
	if !namePattern.MatchString(name) {
		return Name{}, fmt.Errorf("Names, surnames and patronymics must be between 1 and 80 characters with no special symbols, got %q", name)
	}
	return Name{value: name}, nil
}

// String returns the name string.
func (n Name) String() string {
	return n.value
}
