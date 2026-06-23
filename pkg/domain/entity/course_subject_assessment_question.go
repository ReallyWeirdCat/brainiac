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

type CourseSubjectAssessmentQuestion struct {
	GUID                        valueobject.GUID      `json:"guid"`
	CourseSubjectAssessmentGUID valueobject.GUID      `json:"course_subject_assessment_guid"`
	TitleI18n                   *valueobject.I18nText `json:"title_i18n,omitempty"`
	DescriptionI18n             *valueobject.I18nText `json:"description_i18n,omitempty"`
	QuestionType                enum.QuestionTypeEnum `json:"question_type"`
	AttachmentURL               *valueobject.HttpUrl  `json:"attachment_url,omitempty"`
	Options                     *valueobject.I18nText `json:"options,omitempty"`
	CorrectOptions              *valueobject.I18nText `json:"correct_options,omitempty"`
	MaxCorrectOptions           int16                 `json:"max_correct_options"`
	IsMultipleChoice            bool                  `json:"is_multiple_choice"`
	MaxOptions                  int16                 `json:"max_options"`
	MaxAnswerTime               int16                 `json:"max_answer_time"`
	UseTextAnswer               bool                  `json:"use_text_answer"`
	CorrectTextAnswer           *valueobject.I18nText `json:"correct_text_answer,omitempty"`
	ExampleUrl                  *valueobject.HttpUrl  `json:"example_url,omitempty"`
	PublishedAt                 *time.Time            `json:"published_at,omitempty"`
	CreatedAt                   time.Time             `json:"created_at"`
	DeletedAt                   *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = CourseSubjectAssessment{}

func (c CourseSubjectAssessmentQuestion) IsValid() bool {
	if c.TitleI18n != nil && !c.TitleI18n.IsValid() {
		return false
	}
	if c.DescriptionI18n != nil && !c.DescriptionI18n.IsValid() {
		return false
	}
	if c.AttachmentURL != nil && !c.AttachmentURL.IsValid() {
		return false
	}
	if c.CorrectOptions != nil && !c.CorrectOptions.IsValid() {
		return false
	}
	if c.CorrectTextAnswer != nil && !c.CorrectTextAnswer.IsValid() {
		return false
	}
	if c.ExampleUrl != nil && !c.ExampleUrl.IsValid() {
		return false
	}
	if c.MaxCorrectOptions < 1 || c.MaxCorrectOptions > c.MaxOptions {
		return false
	}
	if c.MaxAnswerTime < 10 {
		return false
	}
	if c.GUID == nil {
		return false
	}

	return c.GUID.IsValid() && c.CourseSubjectAssessmentGUID.IsValid() && c.QuestionType.IsValid()
}
