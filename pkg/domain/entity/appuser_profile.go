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

var _ Entity = &AppUserProfile{}

type AppUserProfile struct {
	AppUserGUID       valueobject.GUID          `json:"app_user_guid"`
	Name              *valueobject.Name         `json:"name,omitempty"`
	Surname           *valueobject.Name         `json:"surname,omitempty"`
	Patronymic        *valueobject.Name         `json:"patronymic,omitempty"`
	Nickname          *valueobject.Nickname     `json:"nickname,omitempty"`
	Bio               *valueobject.Bio          `json:"bio,omitempty"`
	PreferredLanguage valueobject.LanguageCode  `json:"preferred_language"`
	ProfileDiscovery  enum.ProfileDiscoveryEnum `json:"profile_discovery"`
	AvatarUrl         *valueobject.HttpUrl      `json:"avatar_url,omitempty"`
	EditingLockedAt   *time.Time                `json:"editing_locked_at,omitempty"`
	CreatedAt         time.Time                 `json:"created_at"`
	DeletedAt         *time.Time                `json:"deleted_at,omitempty"`
}

func (a *AppUserProfile) IsValid() bool {
	if a.AppUserGUID == nil || !a.AppUserGUID.IsValid() {
		return false
	}
	if a.Name != nil && !a.Name.IsValid() {
		return false
	}
	if a.Surname != nil && !a.Surname.IsValid() {
		return false
	}
	if a.Patronymic != nil && !a.Patronymic.IsValid() {
		return false
	}
	if a.Nickname != nil && !a.Nickname.IsValid() {
		return false
	}
	if a.Bio != nil && !a.Bio.IsValid() {
		return false
	}
	if !a.PreferredLanguage.IsValid() {
		return false
	}
	if !a.ProfileDiscovery.IsValid() {
		return false
	}
	if a.AvatarUrl != nil && !a.AvatarUrl.IsValid() {
		return false
	}
	if a.CreatedAt.IsZero() {
		return false
	}
	return true
}
