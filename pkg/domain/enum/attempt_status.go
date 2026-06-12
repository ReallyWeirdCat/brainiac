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

// AttemptStatusEnum represents attempt status (0=in_progress, 1=finished, 2=flagged)
type AttemptStatusEnum int16

const (
	AttemptInProgress AttemptStatusEnum = 0
	AttemptFinished   AttemptStatusEnum = 1
	AttemptFlagged    AttemptStatusEnum = 2
)

var _ Enum = (*AttemptStatusEnum)(nil)

func (e AttemptStatusEnum) String() string {
	switch e {
	case AttemptInProgress:
		return "in_progress"
	case AttemptFinished:
		return "finished"
	case AttemptFlagged:
		return "flagged"
	default:
		return "unknown"
	}
}

func (e AttemptStatusEnum) IsValid() bool {
	return e >= 0 && e <= 2
}

// Value returns the integer value for database operations
func (e AttemptStatusEnum) Value() int16 {
	return int16(e)
}
