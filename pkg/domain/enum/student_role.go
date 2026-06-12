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

// StudentRoleEnum represents student role (0=unspecified, 1=student, 2=basic, 3=intermediate)
type StudentRoleEnum int16

const (
	UnspecifiedRole  StudentRoleEnum = 0
	StudentRole      StudentRoleEnum = 1
	BasicRole        StudentRoleEnum = 2
	IntermediateRole StudentRoleEnum = 3
)

var _ Enum = (*StudentRoleEnum)(nil)

func (e StudentRoleEnum) String() string {
	switch e {
	case UnspecifiedRole:
		return "unspecified"
	case StudentRole:
		return "student"
	case BasicRole:
		return "basic"
	case IntermediateRole:
		return "intermediate"
	default:
		return "unknown"
	}
}

func (e StudentRoleEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// Value returns the integer value for database operations
func (e StudentRoleEnum) Value() int16 {
	return int16(e)
}
