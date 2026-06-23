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

type Chat struct {
	GUID               valueobject.GUID      `json:"guid"`
	Title              string                `json:"title"`
	Topic              *string               `json:"topic,omitempty"`
	AvatarURL          *valueobject.HttpUrl  `json:"avatar_url,omitempty"`
	StudentsCanText    bool                  `json:"students_can_text"`
	StudentsSeeMembers bool                  `json:"students_see_members"`
	PreserveMessages   bool                  `json:"preserve_messages"`
	Meta               *valueobject.Metadata `json:"meta,omitempty"`
	CreatedAt          time.Time             `json:"created_at"`
	DeletedAt          *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = Chat{}

func (c Chat) IsValid() bool {
	if len(c.Title) > 50 {
		return false
	}
	if c.Topic != nil && len(*c.Topic) > 50 {
		return false
	}
	if c.AvatarURL != nil && !c.AvatarURL.IsValid() {
		return false
	}
	if c.Meta != nil && !c.Meta.IsValid() {
		return false
	}
	return c.GUID != nil && c.GUID.IsValid()
}
