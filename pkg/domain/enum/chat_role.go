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

// ChatRoleEnum represents chat role (0=member, 1=admin, 2=owner)
type ChatRoleEnum int16

const (
	Member ChatRoleEnum = 0
	Admin  ChatRoleEnum = 1
	Owner  ChatRoleEnum = 2
)

var _ Enum = (*ChatRoleEnum)(nil)

func (e ChatRoleEnum) String() string {
	switch e {
	case Member:
		return "member"
	case Admin:
		return "admin"
	case Owner:
		return "owner"
	default:
		return "unknown"
	}
}

func (e ChatRoleEnum) IsValid() bool {
	return e >= 0 && e <= 2
}

// Value returns the integer value for database operations
func (e ChatRoleEnum) Value() int16 {
	return int16(e)
}
