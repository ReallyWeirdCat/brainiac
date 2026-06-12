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

type Course struct {
	GUID        valueobject.GUID
	Title       *valueobject.I18nText
	Description *valueobject.I18nText
	Style       *valueobject.Metadata
	Meta        *valueobject.Metadata
	PublishedAt *time.Time
	CreatedAt   time.Time
	DeletedAt   *time.Time
}

var _ Entity = Course{}

func (c Course) IsValid() bool {

	if c.Title != nil && !c.Title.IsValid() {
		return false
	}
	if c.Description != nil && !c.Description.IsValid() {
		return false
	}
	if c.Style != nil && !c.Style.IsValid() {
		return false
	}
	if c.Meta != nil && !c.Meta.IsValid() {
		return false
	}

	return c.GUID.IsValid()
}
