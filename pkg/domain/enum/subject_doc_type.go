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

// SubjectDocTypeEnum represents subject document type (0=theory, 1=practice, 2=example, 3=literature)
type SubjectDocTypeEnum int16

const (
	TheoryDoc     SubjectDocTypeEnum = 0
	PracticeDoc   SubjectDocTypeEnum = 1
	ExampleDoc    SubjectDocTypeEnum = 2
	LiteratureDoc SubjectDocTypeEnum = 3
)

var _ Enum = (*SubjectDocTypeEnum)(nil)

func (e SubjectDocTypeEnum) String() string {
	switch e {
	case TheoryDoc:
		return "theory"
	case PracticeDoc:
		return "practice"
	case ExampleDoc:
		return "example"
	case LiteratureDoc:
		return "literature"
	default:
		return "unknown"
	}
}

func (e SubjectDocTypeEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// Value returns the integer value for database operations
func (e SubjectDocTypeEnum) Value() int16 {
	return int16(e)
}
