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

type CourseSubject struct {
	GUID                valueobject.GUID             `json:"guid"`
	CourseGUID          valueobject.GUID             `json:"course_guid"`
	ParentGUID          *valueobject.GUID            `json:"parent_guid,omitempty"`
	TitleI18n           *valueobject.I18nText        `json:"title_i18n,omitempty"`
	DescriptionI18n     *valueobject.I18nText        `json:"description_i18n,omitempty"`
	RewardExperience    int64                        `json:"reward_experience"`
	CompletionCondition enum.CompletionConditionEnum `json:"completion_condition"`
	HideChildren        bool                         `json:"hide_children"`
	Style               *valueobject.Metadata        `json:"style,omitempty"`
	Meta                *valueobject.Metadata        `json:"meta,omitempty"`
	PublishedAt         *time.Time                   `json:"published_at,omitempty"`
	CreatedAt           time.Time                    `json:"created_at"`
	DeletedAt           *time.Time                   `json:"deleted_at,omitempty"`
}

var _ Entity = &Course{}

func (c *CourseSubject) IsValid() bool {

	if !c.GUID.IsValid() || !c.CourseGUID.IsValid() {
		return false
	}
	if c.TitleI18n != nil && !c.TitleI18n.IsValid() {
		return false
	}
	if c.DescriptionI18n != nil && !c.DescriptionI18n.IsValid() {
		return false
	}
	if c.Style != nil && !c.Style.IsValid() {
		return false
	}
	if c.Meta != nil && !c.Meta.IsValid() {
		return false
	}
	if c.ParentGUID != nil && (!(*c.ParentGUID).IsValid() || *c.ParentGUID == c.GUID) {
		return false
	}

	return c.CompletionCondition.IsValid()
}
