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

type StudentGroupMemberClassAttendance struct {
	GUID                   valueobject.GUID      `json:"guid"`
	StudentGroupMemberGUID valueobject.GUID      `json:"student_group_member_guid"`
	FirstSeenAt            *time.Time            `json:"first_seen_at,omitempty"`
	LastSeenAt             *time.Time            `json:"last_seen_at,omitempty"`
	PresentAt              *time.Time            `json:"present_at,omitempty"`
	ObjectivesCompletion   *valueobject.Metadata `json:"objectives_completion,omitempty"`
	Comment                *string               `json:"comment,omitempty"`
	Grade                  enum.GradeEnum        `json:"grade"`
	CreatedAt              time.Time             `json:"created_at"`
	DeletedAt              *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = StudentGroupMemberClassAttendance{}

func (s StudentGroupMemberClassAttendance) IsValid() bool {
	if s.ObjectivesCompletion != nil && !s.ObjectivesCompletion.IsValid() {
		return false
	}
	if s.Comment != nil && len(*s.Comment) > 1024 {
		return false
	}
	if s.GUID == nil || s.StudentGroupMemberGUID == nil {
		return false
	}
	return s.GUID.IsValid() && s.StudentGroupMemberGUID.IsValid() && s.Grade.IsValid()
}
