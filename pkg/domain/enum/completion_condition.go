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

// CompletionConditionEnum represents completion condition (0=manual_completion, 1=attestation_test, 2=approved_practice, 3=all)
type CompletionConditionEnum int16

const (
	ManualCompletion CompletionConditionEnum = 0
	AttestationTest  CompletionConditionEnum = 1
	ApprovedPractice CompletionConditionEnum = 2
	AllCompletion    CompletionConditionEnum = 3
)

var _ Enum = (*CompletionConditionEnum)(nil)

func (e CompletionConditionEnum) String() string {
	switch e {
	case ManualCompletion:
		return "manual_completion"
	case AttestationTest:
		return "attestation_test"
	case ApprovedPractice:
		return "approved_practice"
	case AllCompletion:
		return "all"
	default:
		return "unknown"
	}
}

func (e CompletionConditionEnum) Valid() bool {
	return e >= 0 && e <= 3
}

// Value returns the integer value for database operations
func (e CompletionConditionEnum) Value() int16 {
	return int16(e)
}
