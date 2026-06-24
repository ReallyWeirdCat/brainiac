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
	AppUserGUID           valueobject.GUID `json:"app_user_guid"`
	Level                 int16            `json:"level"`
	OtherExperience       int64            `json:"other_experience"`
	Money                 int64            `json:"money"`
	MoneyEarned           int64            `json:"money_earned"`
	SubjectsCompleted     int32            `json:"subjects_completed"`
	AssessmentsCompleted  int32            `json:"assessments_completed"`
	AssessmentsFailed     int32            `json:"assessments_failed"`
	AssessmentsTerminated int32            `json:"assessments_terminated"`
	WorksSubmitted        int32            `json:"works_submitted"`
	ItemsCollected        int32            `json:"items_collected"`
	ItemsUsed             int32            `json:"items_used"`
	ItemsSold             int32            `json:"items_sold"`
	ItemsExchanged        int32            `json:"items_exchanged"`
	BoxesOpened           int32            `json:"boxes_opened"`
	MaxDailyStreak        int32            `json:"max_daily_streak"`
	MessagesSent          int32            `json:"messages_sent"`
	CreatedAt             time.Time        `json:"created_at"`
	DeletedAt             *time.Time       `json:"deleted_at,omitempty"`
}

var _ Entity = &StudentGlobalStats{}

func (s *StudentGlobalStats) IsValid() bool {
	if !s.AppUserGUID.IsValid() {
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
