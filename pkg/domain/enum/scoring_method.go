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

import (
	"fmt"
)

// ScoringMethodEnum represents the scoring method (0=max_score, 1=average_score, 2=latest_score)
type ScoringMethodEnum int16

const (
	MaxScore     ScoringMethodEnum = 0
	AverageScore ScoringMethodEnum = 1
	LatestScore  ScoringMethodEnum = 2
)

var _ Enum = (*ScoringMethodEnum)(nil)

func (e ScoringMethodEnum) String() string {
	switch e {
	case MaxScore:
		return "max_score"
	case AverageScore:
		return "average_score"
	case LatestScore:
		return "latest_score"
	default:
		return "unknown"
	}
}

func (e ScoringMethodEnum) IsValid() bool {
	return e >= 0 && e <= 2
}

// ParseScoringMethodEnum parses a string into a ScoringMethodEnum
func ParseScoringMethodEnum(s string) (ScoringMethodEnum, error) {
	switch s {
	case "max_score":
		return MaxScore, nil
	case "average_score":
		return AverageScore, nil
	case "latest_score":
		return LatestScore, nil
	default:
		return -1, fmt.Errorf("invalid scoring method: %s", s)
	}
}

// Value returns the integer value for database operations
func (e ScoringMethodEnum) Value() int16 {
	return int16(e)
}
