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
	AppUserGUID       valueobject.GUID
	Name              *valueobject.Name
	Surname           *valueobject.Name
	Patronymic        *valueobject.Name
	Nickname          *valueobject.Nickname
	Bio               *valueobject.Bio
	PreferredLanguage valueobject.LanguageCode
	ProfileDiscovery   enum.ProfileDiscoveryEnum
	AvatarUrl         *valueobject.HttpUrl
	EditingLockedAt   *time.Time
	CreatedAt         time.Time
	DeletedAt         *time.Time
}

func (a *AppUserProfile) IsValid() bool {
	if !a.AppUserGUID.IsValid() {
		return false
	}
	// Required fields (non-nil and valid)
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
	if !a.ProfileDisovery.IsValid() {
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
