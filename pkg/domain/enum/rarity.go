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

// RarityEnum represents rarity level (0=common, 1=uncommon, 2=rare, 3=epic, 4=legendary, 5=mythical)
type RarityEnum int16

const (
	CommonRarity    RarityEnum = 0
	UncommonRarity  RarityEnum = 1
	RareRarity      RarityEnum = 2
	EpicRarity      RarityEnum = 3
	LegendaryRarity RarityEnum = 4
	MythicalRarity  RarityEnum = 5
)

var _ Enum = (*RarityEnum)(nil)

func (e RarityEnum) String() string {
	switch e {
	case CommonRarity:
		return "common"
	case UncommonRarity:
		return "uncommon"
	case RareRarity:
		return "rare"
	case EpicRarity:
		return "epic"
	case LegendaryRarity:
		return "legendary"
	case MythicalRarity:
		return "mythical"
	default:
		return "unknown"
	}
}

func (e RarityEnum) Valid() bool {
	return e >= 0 && e <= 5
}

// Value returns the integer value for database operations
func (e RarityEnum) Value() int16 {
	return int16(e)
}
