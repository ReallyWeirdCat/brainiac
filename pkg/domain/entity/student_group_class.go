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

type StudentGroupClass struct {
	GUID                          valueobject.GUID
	StudentGroupGUID              valueobject.GUID
	Title                         string
	StartTime                     time.Time
	EndTime                       time.Time
	Objectives                    *valueobject.Metadata
	AttestationAssessmentsEnabled bool
	AttendanceClosedAt            *time.Time
	MarksApprovedAt               *time.Time
	TOTPSecret                    []byte
	Room                          *string
	ClassType                     enum.ClassTypeEnum
	CreatedAt                     time.Time
	DeletedAt                     *time.Time
}

var _ Entity = StudentGroupClass{}

func (s StudentGroupClass) IsValid() bool {

	if len(s.Title) > 50 {
		return false
	}
	if s.EndTime.Before(s.StartTime) || s.StartTime.Equal(s.EndTime) {
		return false
	}
	if s.Room != nil && len(*s.Room) > 15 {
		return false
	}

	return s.GUID.IsValid() && s.StudentGroupGUID.IsValid() && s.ClassType.IsValid()
}
