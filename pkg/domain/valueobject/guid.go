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

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
)

const uuidPattern = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

type GUID struct {
	value string
}

var _ ValueObject = GUID{}

func NewGUID(guid string) (GUID, error) {
	matched, _ := regexp.MatchString(uuidPattern, guid)

	if !matched {
		return GUID{}, errors.ErrInvalidGUID
	}

	return GUID{value: guid}, nil
}

func (g GUID) IsValid() bool {
	_, err := NewGUID(g.value)
	return err == nil
}

func (g GUID) Equals(other any) bool {
	if other == nil {
		return false
	}
	switch v := other.(type) {
	case GUID:
		return g.value == v.value
	case string:
		return g.value == v
	default:
		return false
	}
}

func (g GUID) IsZero() bool {
	return g.value == ""
}

func (g GUID) String() string {
	return g.value
}
