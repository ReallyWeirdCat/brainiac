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

package entity

import (
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/enum"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type StudentGroupMemberAttendance struct {
	GUID                   valueobject.GUID
	StudentGroupMemberGUID valueobject.GUID
	FirstSeenAt            *time.Time
	LastSeenAt             *time.Time
	PresentAt              *time.Time
	ObjectivesCompletion   *valueobject.Metadata
	Comment                *string
	Grade                  enum.GradeEnum
	CreatedAt              time.Time
	DeletedAt              *time.Time
}

var _ Entity = StudentGroupMemberAttendance{}

func (s StudentGroupMemberAttendance) IsValid() bool {
	if s.ObjectivesCompletion != nil && !s.ObjectivesCompletion.IsValid() {
		return false
	}
	if s.Comment != nil && len(*s.Comment) > 1024 {
		return false
	}
	return s.GUID.IsValid() && s.StudentGroupMemberGUID.IsValid() && s.Grade.IsValid()
}
