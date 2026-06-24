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

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type TeacherGroup struct {
	GUID               valueobject.GUID `json:"guid"`
	TeacherGUID        valueobject.GUID `json:"teacher_guid"`
	StudentGroupGUID   valueobject.GUID `json:"student_group_guid"`
	ActiveSince        *time.Time       `json:"active_since,omitempty"`
	ManageAttendace    bool             `json:"manage_attendance"`
	ManageResults      bool             `json:"manage_results"`
	ManageStudents     bool             `json:"manage_students"`
	ManageAchievements bool             `json:"manage_achievements"`
	TeacherUntil       *time.Time       `json:"teacher_until,omitempty"`
	CreatedAt          time.Time        `json:"created_at"`
	DeletedAt          *time.Time       `json:"deleted_at,omitempty"`
}

var _ Entity = TeacherGroup{}

func (t TeacherGroup) IsValid() bool {
	return t.GUID.IsValid() && t.TeacherGUID.IsValid() && t.StudentGroupGUID.IsValid()
}
