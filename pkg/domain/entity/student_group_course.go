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

type StudentGroupCourse struct {
	GUID             valueobject.GUID `json:"guid"`
	StudentGroupGUID valueobject.GUID `json:"student_group_guid"`
	CourseGUID       valueobject.GUID `json:"course_guid"`
	FullAccess       bool             `json:"full_access"`
	CloseSubmissions bool             `json:"close_submissions"`
	CreatedAt        time.Time        `json:"created_at"`
	DeletedAt        *time.Time       `json:"deleted_at,omitempty"`
}

var _ Entity = StudentGroupCourse{}

func (s StudentGroupCourse) IsValid() bool {
	if s.GUID == nil || s.StudentGroupGUID == nil || s.CourseGUID == nil {
		return false
	}
	return s.GUID.IsValid() && s.StudentGroupGUID.IsValid() && s.CourseGUID.IsValid()
}
