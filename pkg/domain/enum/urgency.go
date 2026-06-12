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

// UrgencyEnum represents urgency level (0=low, 1=normal, 2=important)
type UrgencyEnum int16

const (
	LowUrgency       UrgencyEnum = 0
	NormalUrgency    UrgencyEnum = 1
	ImportantUrgency UrgencyEnum = 2
)

var _ Enum = (*UrgencyEnum)(nil)

func (e UrgencyEnum) String() string {
	switch e {
	case LowUrgency:
		return "low"
	case NormalUrgency:
		return "normal"
	case ImportantUrgency:
		return "important"
	default:
		return "unknown"
	}
}

func (e UrgencyEnum) Valid() bool {
	return e >= 0 && e <= 2
}

// Value returns the integer value for database operations
func (e UrgencyEnum) Value() int16 {
	return int16(e)
}
