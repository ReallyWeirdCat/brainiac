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

// QuestionTypeEnum represents the type of question (0=short_question, 1=normal_question, 2=long_question)
type QuestionTypeEnum int16

const (
	ShortQuestion  QuestionTypeEnum = 0
	NormalQuestion QuestionTypeEnum = 1
	LongQuestion   QuestionTypeEnum = 2
)

func (e QuestionTypeEnum) String() string {
	switch e {
	case ShortQuestion:
		return "short_question"
	case NormalQuestion:
		return "normal_question"
	case LongQuestion:
		return "long_question"
	default:
		return "unknown"
	}
}

func (e QuestionTypeEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// ParseQuestionTypeEnum parses a string into a QuestionTypeEnum
func ParseQuestionTypeEnum(s string) (QuestionTypeEnum, error) {
	switch s {
	case "short_question":
		return ShortQuestion, nil
	case "normal_question":
		return NormalQuestion, nil
	case "long_question":
		return LongQuestion, nil
	default:
		return -1, fmt.Errorf("invalid question type: %s", s)
	}
}

// Value returns the integer value for database operations
func (e QuestionTypeEnum) Value() int16 {
	return int16(e)
}
