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

type StudentGlobalStats struct {
	AppUserGUID           valueobject.GUID
	Level                 int16
	OtherExperience       int64
	Money                 int64
	MoneyEarned           int64
	SubjectsCompleted     int32
	AssessmentsCompleted  int32
	AssessmentsFailed     int32
	AssessmentsTerminated int32
	WorksSubmitted        int32
	ItemsCollected        int32
	ItemsUsed             int32
	ItemsSold             int32
	ItemsExchanged        int32
	BoxesOpened           int32
	MaxDailyStreak        int32
	MessagesSent          int32
	CreatedAt             time.Time
	DeletedAt             *time.Time
}

var _ Entity = &StudentGlobalStats{}

func (s *StudentGlobalStats) IsValid() bool {
	if s.AppUserGUID == nil || !s.AppUserGUID.IsValid() {
		return false
	}
	if s.Level < 0 {
		return false
	}
	if s.OtherExperience < 0 || s.Money < 0 || s.MoneyEarned < 0 {
		return false
	}
	if s.SubjectsCompleted < 0 || s.AssessmentsCompleted < 0 || s.AssessmentsFailed < 0 ||
		s.AssessmentsTerminated < 0 || s.WorksSubmitted < 0 || s.ItemsCollected < 0 ||
		s.ItemsUsed < 0 || s.ItemsSold < 0 || s.ItemsExchanged < 0 || s.BoxesOpened < 0 ||
		s.MaxDailyStreak < 0 || s.MessagesSent < 0 {
		return false
	}
	if s.CreatedAt.IsZero() {
		return false
	}
	return true
}
