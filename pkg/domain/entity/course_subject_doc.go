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

type CourseSubjectDoc struct {
	GUID              valueobject.GUID        `json:"guid"`
	CourseSubjectGUID valueobject.GUID        `json:"course_subject_guid"`
	TitleI18n         *valueobject.I18nText   `json:"title_i18n,omitempty"`
	DescriptionI18n   *valueobject.I18nText   `json:"description_i18n,omitempty"`
	SubjectDocType    enum.SubjectDocTypeEnum `json:"subject_doc_type"`
	ExampleMeta       *valueobject.Metadata   `json:"example_meta,omitempty"`
	SubmissionMeta    *valueobject.Metadata   `json:"submission_meta,omitempty"`
	LiteratureMeta    *valueobject.Metadata   `json:"literature_meta,omitempty"`
	DocIndex          int16                   `json:"doc_index"`
	PublishedAt       *time.Time              `json:"published_at,omitempty"`
	CreatedAt         time.Time               `json:"created_at"`
	DeletedAt         *time.Time              `json:"deleted_at,omitempty"`
}

var _ Entity = CourseSubjectDoc{}

func (c CourseSubjectDoc) IsValid() bool {

	if c.TitleI18n != nil && !c.TitleI18n.IsValid() {
		return false
	}
	if c.DescriptionI18n != nil && !c.DescriptionI18n.IsValid() {
		return false
	}
	if c.ExampleMeta != nil && !c.ExampleMeta.IsValid() {
		return false
	}
	if c.SubmissionMeta != nil && !c.SubmissionMeta.IsValid() {
		return false
	}
	if c.LiteratureMeta != nil && !c.LiteratureMeta.IsValid() {
		return false
	}
	if c.GUID == nil || c.CourseSubjectGUID == nil {
		return false
	}

	return c.GUID.IsValid() && c.CourseSubjectGUID.IsValid() && c.SubjectDocType.IsValid()
}
