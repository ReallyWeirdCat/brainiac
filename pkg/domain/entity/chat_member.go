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

type ChatMember struct {
	GUID           valueobject.GUID  `json:"guid"`
	AppUserGUID    valueobject.GUID  `json:"app_user_guid"`
	ChatGUID       valueobject.GUID  `json:"chat_guid"`
	ChatRole       enum.ChatRoleEnum `json:"chat_role"`
	Name           *string           `json:"name,omitempty"`
	CustomRoleName *string           `json:"custom_role_name,omitempty"`
	LeftAt         *time.Time        `json:"left_at,omitempty"`
	SeenAt         *time.Time        `json:"seen_at,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
	DeletedAt      time.Time         `json:"deleted_at"`
}

var _ Entity = ChatMember{}

func (c ChatMember) IsValid() bool {
	if c.Name != nil && len(*c.Name) > 150 {
		return false
	}
	if c.CustomRoleName != nil && len(*c.CustomRoleName) > 100 {
		return false
	}
	if c.GUID == nil || c.AppUserGUID == nil || c.ChatGUID == nil {
		return false
	}

	return c.GUID.IsValid() && c.AppUserGUID.IsValid() && c.ChatGUID.IsValid() && c.ChatRole.IsValid()
}
