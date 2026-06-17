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
	GUID                  valueobject.GUID
	CourseSubjectTestGUID valueobject.GUID
	TitleI18n             *valueobject.I18nText
	DescriptionI18n       *valueobject.I18nText
	QuestionType          *enum.QuestionTypeEnum
	AttachmentURL         *valueobject.HttpUrl
	Options               *valueobject.I18nText
	CorrectOptions        *valueobject.I18nText
	MaxCorrectOptions     int16
	IsMultipleChoice      bool
	MaxOptions            int16
	MaxAnswerTime         int16
	UseTextAnswer         bool
	CorrectTextAnswer     *valueobject.I18nText
	ExampleUrl            *valueobject.HttpUrl
	PublishedAt           *time.Time
	CreatedAt             time.Time
	DeletedAt             *time.Time
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
	if c.MaxAnswerTime < 10 && c.MaxAnswerTime != 0 {
		return false
	}

	return c.GUID.IsValid() && c.CourseSubjectTestGUID.IsValid() && c.QuestionType.IsValid()
}
