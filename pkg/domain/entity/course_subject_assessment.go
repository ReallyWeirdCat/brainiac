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

type CourseSubjectAssessment struct {
	GUID                 valueobject.GUID       `json:"guid"`
	TitleI18n            *valueobject.I18nText  `json:"title_i18n,omitempty"`
	DescriptionI18n      *valueobject.I18nText  `json:"description_i18n,omitempty"`
	IsAttestation        bool                   `json:"is_attestation"`
	HideAnswers          bool                   `json:"hide_answers"`
	HideAnswerValidity   bool                   `json:"hide_answer_validity"`
	HideScore            bool                   `json:"hide_score"`
	UseTimedQuestions    bool                   `json:"use_timed_questions"`
	MaxTime              int16                  `json:"max_time"`
	MaxAttempts          int16                  `json:"max_attempts"`
	ScoringMethod        enum.ScoringMethodEnum `json:"scoring_method"`
	RequiredScore        int8                   `json:"required_score"`
	ShortQuestionsCount  int16                  `json:"short_questions_count"`
	NormalQuestionsCount int16                  `json:"normal_questions_count"`
	LongQuestionsCount   int16                  `json:"long_questions_count"`
	PublishedAt          *time.Time             `json:"published_at,omitempty"`
	CreatedAt            time.Time              `json:"created_at"`
	DeletedAt            *time.Time             `json:"deleted_at,omitempty"`
}

var _ Entity = CourseSubjectAssessment{}

func (c CourseSubjectAssessment) IsValid() bool {
	if c.TitleI18n != nil && !c.TitleI18n.IsValid() {
		return false
	}
	if c.DescriptionI18n != nil && !c.DescriptionI18n.IsValid() {
		return false
	}
	if c.RequiredScore < 0 || c.RequiredScore > 100 {
		return false
	}
	return c.GUID.IsValid() && c.ScoringMethod.IsValid()
}
