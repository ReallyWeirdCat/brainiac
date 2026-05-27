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

// ProfileDiscoveryEnum represents profile discovery level (0=admins_only, 1=teachers, 2=group_mates, 3=course_mates, 4=users, 5=guests)
type ProfileDiscoveryEnum int16

const (
	AdminsOnly  ProfileDiscoveryEnum = 0
	Teachers    ProfileDiscoveryEnum = 1
	GroupMates  ProfileDiscoveryEnum = 2
	CourseMates ProfileDiscoveryEnum = 3
	AllUsers    ProfileDiscoveryEnum = 4
	Guests      ProfileDiscoveryEnum = 5
)

func (e ProfileDiscoveryEnum) String() string {
	switch e {
	case AdminsOnly:
		return "admins_only"
	case Teachers:
		return "teachers"
	case GroupMates:
		return "group_mates"
	case CourseMates:
		return "course_mates"
	case AllUsers:
		return "users"
	case Guests:
		return "guests"
	default:
		return "unknown"
	}
}

func (e ProfileDiscoveryEnum) Valid() bool {
	return e >= 0 && e <= 5
}

// Value returns the integer value for database operations
func (e ProfileDiscoveryEnum) Value() int16 {
	return int16(e)
}
