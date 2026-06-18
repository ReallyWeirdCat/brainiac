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

type DailyActivity struct {
	GUID                 valueobject.GUID
	AppUserGUID          valueobject.GUID
	Day                  time.Time
	ExperienceEarned     int64
	LevelsEarned         int16
	SubjectsCompleted    int16
	AssessmentsCompleted int16
	PracticesCompleted   int16
	CreatedAt            time.Time
	DeletedAt            *time.Time
}

var _ Entity = &DailyActivity{}

func (d *DailyActivity) IsValid() bool {
	if d.GUID == nil || !d.GUID.IsValid() {
		return false
	}
	if d.AppUserGUID == nil || !d.AppUserGUID.IsValid() {
		return false
	}
	if d.Day.IsZero() {
		return false
	}
	if d.ExperienceEarned < 0 || d.LevelsEarned < 0 ||
		d.SubjectsCompleted < 0 || d.AssessmentsCompleted < 0 || d.PracticesCompleted < 0 {
		return false
	}
	if d.CreatedAt.IsZero() {
		return false
	}
	return true
}
