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

package enums

import (
	"fmt"
)

// GradeEnum represents grade (0=unspecified, 1=very_poor, 2=poor, 3=satisfactory, 4=good, 5=very_good)
type GradeEnum int16

const (
	Unspecified  GradeEnum = 0
	VeryPoor     GradeEnum = 1
	Poor         GradeEnum = 2
	Satisfactory GradeEnum = 3
	Good         GradeEnum = 4
	VeryGood     GradeEnum = 5
)

func (e GradeEnum) String() string {
	switch e {
	case Unspecified:
		return "unspecified"
	case VeryPoor:
		return "very_poor"
	case Poor:
		return "poor"
	case Satisfactory:
		return "satisfactory"
	case Good:
		return "good"
	case VeryGood:
		return "very_good"
	default:
		return "unknown"
	}
}

func (e GradeEnum) Valid() bool {
	return e >= 0 && e <= 5
}

// ParseGradeEnum parses a string into a GradeEnum
func ParseGradeEnum(s string) (GradeEnum, error) {
	switch s {
	case "unspecified":
		return Unspecified, nil
	case "very_poor":
		return VeryPoor, nil
	case "poor":
		return Poor, nil
	case "satisfactory":
		return Satisfactory, nil
	case "good":
		return Good, nil
	case "very_good":
		return VeryGood, nil
	default:
		return -1, fmt.Errorf("invalid grade: %s", s)
	}
}

// Value returns the integer value for database operations
func (e GradeEnum) Value() int16 {
	return int16(e)
}
