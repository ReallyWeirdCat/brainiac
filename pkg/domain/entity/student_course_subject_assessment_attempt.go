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

type StudentCourseSubjectAssessmentAttempt struct {
	GUID          valueobject.GUID       `json:"guid"`
	AppUserGUID   valueobject.GUID       `json:"app_user_guid"`
	AttemptStatus enum.AttemptStatusEnum `json:"attempt_status"`
	Score         int16                  `json:"score"`
	IsSuccess     bool                   `json:"is_success"`
	Answers       valueobject.Metadata   `json:"answers"`
	CreatedAt     time.Time              `json:"created_at"`
	DeletedAt     *time.Time             `json:"deleted_at,omitempty"`
}

var _ Entity = StudentCourseSubjectAssessmentAttempt{}

func (s StudentCourseSubjectAssessmentAttempt) IsValid() bool {
	if s.Score < 0 || s.Score > 100 {
		return false
	}
	return s.GUID.IsValid() && s.AppUserGUID.IsValid() && s.AttemptStatus.IsValid() && s.Answers.IsValid()
}
