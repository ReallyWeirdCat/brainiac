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
	GUID                valueobject.GUID
	CourseGUID          valueobject.GUID
	ParentGUID          *valueobject.GUID
	TitleI18n           *valueobject.I18nText
	DescriptionI18n     *valueobject.I18nText
	RewardExperience    int64
	CompletionCondition enum.CompletionConditionEnum
	HideChildren        bool
	Style               *valueobject.Metadata
	Meta                *valueobject.Metadata
	PublishedAt         *time.Time
	CreatedAt           time.Time
	DeletedAt           *time.Time
}

var _ Entity = Course{}

func (c CourseSubject) IsValid() bool {

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
	if c.GUID == nil || c.CourseGUID == nil || c.ParentGUID != nil && (!(*c.ParentGUID).IsValid() || c.ParentGUID == &c.GUID) {
		return false
	}

	return c.CompletionCondition.IsValid()
}
