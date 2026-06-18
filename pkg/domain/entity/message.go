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

type Message struct {
	GUID           valueobject.GUID
	ChatMemberGUID valueobject.GUID
	ChatGUID       valueobject.GUID
	ResourceURL    *valueobject.HttpUrl
	Content        *string
	Urgency        enum.UrgencyEnum
	Meta           *valueobject.Metadata
	SeenBy         *valueobject.Metadata
	EditedAt       *time.Time
	PinnedAt       *time.Time
	ExpireAt       *time.Time
	CreatedAt      time.Time
	DeletedAt      *time.Time
}

var _ Entity = Message{}

func (m Message) IsValid() bool {
	if m.ResourceURL != nil && !m.ResourceURL.IsValid() {
		return false
	}
	if m.Content != nil && len(*m.Content) > 1024 {
		return false
	}
	if m.Meta != nil && !m.Meta.IsValid() {
		return false
	}
	return m.GUID.IsValid() && m.ChatMemberGUID.IsValid() && m.ChatMemberGUID.IsValid() && m.Urgency.IsValid()
}
