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
	GUID                 valueobject.GUID `json:"guid"`
	AppUserGUID          valueobject.GUID `json:"app_user_guid"`
	Day                  time.Time        `json:"day"`
	ExperienceEarned     int64            `json:"experience_earned"`
	LevelsEarned         int16            `json:"levels_earned"`
	SubjectsCompleted    int16            `json:"subjects_completed"`
	AssessmentsCompleted int16            `json:"assessments_completed"`
	PracticesCompleted   int16            `json:"practices_completed"`
	CreatedAt            time.Time        `json:"created_at"`
	DeletedAt            *time.Time       `json:"deleted_at,omitempty"`
}

var _ Entity = &DailyActivity{}

func (d *DailyActivity) IsValid() bool {
	if !d.GUID.IsValid() {
		return false
	}
	if !d.AppUserGUID.IsValid() {
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
