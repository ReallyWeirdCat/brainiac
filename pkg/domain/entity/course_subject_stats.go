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

type StudentCourseStats struct {
	GUID        valueobject.GUID
	AppUserGUID valueobject.GUID
	CourseGUID  valueobject.GUID
	Experience  int64
	Level       int16
	Meta        *valueobject.Metadata
	CreatedAt   time.Time
	DeletedAt   *time.Time
}

var _ Entity = &StudentCourseStats{}

func (s *StudentCourseStats) IsValid() bool {
	if s.GUID == nil || !s.GUID.IsValid() {
		return false
	}
	if s.AppUserGUID == nil || !s.AppUserGUID.IsValid() {
		return false
	}
	if s.CourseGUID == nil || !s.CourseGUID.IsValid() {
		return false
	}
	if s.Experience < 0 {
		return false
	}
	if s.Level < 0 {
		return false
	}
	if s.Meta != nil && !s.Meta.IsValid() {
		return false
	}
	if s.CreatedAt.IsZero() {
		return false
	}
	return true
}
