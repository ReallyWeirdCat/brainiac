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

// SubjectStatusEnum represents subject status (0=untouched, 1=in_progress, 2=forgotten, 3=done)
type SubjectStatusEnum int16

const (
	Untouched  SubjectStatusEnum = 0
	InProgress SubjectStatusEnum = 1
	Forgotten  SubjectStatusEnum = 2
	Done       SubjectStatusEnum = 3
)

var _ Enum = (*SubjectStatusEnum)(nil)

func (e SubjectStatusEnum) String() string {
	switch e {
	case Untouched:
		return "untouched"
	case InProgress:
		return "in_progress"
	case Forgotten:
		return "forgotten"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}

func (e SubjectStatusEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// Value returns the integer value for database operations
func (e SubjectStatusEnum) Value() int16 {
	return int16(e)
}
