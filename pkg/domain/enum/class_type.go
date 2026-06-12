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

package enum

// ClassTypeEnum represents class type (0=practice, 1=lecture, 2=attestation, 3=consult)
type ClassTypeEnum int16

const (
	PracticeClass    ClassTypeEnum = 0
	LectureClass     ClassTypeEnum = 1
	AttestationClass ClassTypeEnum = 2
	ConsultClass     ClassTypeEnum = 3
)

var _ Enum = (*ChatRoleEnum)(nil)

func (e ClassTypeEnum) String() string {
	switch e {
	case PracticeClass:
		return "practice"
	case LectureClass:
		return "lecture"
	case AttestationClass:
		return "attestation"
	case ConsultClass:
		return "consult"
	default:
		return "unknown"
	}
}

func (e ClassTypeEnum) IsValid() bool {
	return e >= 0 && e <= 3
}

// Value returns the integer value for database operations
func (e ClassTypeEnum) Value() int16 {
	return int16(e)
}
