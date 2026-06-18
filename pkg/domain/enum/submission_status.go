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

// SubmissionStatusEnum represents submission status (0=submitted, 1=approved, 2=rejected)
type SubmissionStatusEnum int16

const (
	Submitted SubmissionStatusEnum = 0
	Approved  SubmissionStatusEnum = 1
	Rejected  SubmissionStatusEnum = 2
)

var _ Enum = (*SubmissionStatusEnum)(nil)

func (e SubmissionStatusEnum) String() string {
	switch e {
	case Submitted:
		return "submitted"
	case Approved:
		return "approved"
	case Rejected:
		return "rejected"
	default:
		return "unknown"
	}
}

func (e SubmissionStatusEnum) IsValid() bool {
	return e >= 0 && e <= 2
}

// Value returns the integer value for database operations
func (e SubmissionStatusEnum) Value() int16 {
	return int16(e)
}
