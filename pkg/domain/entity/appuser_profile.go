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

type AppUserProfile struct {
	AppUserGUID       valueobject.GUID
	Name              *valueobject.Name
	Surname           *valueobject.Name
	Patronymic        *valueobject.Name
	Nickname          *valueobject.Nickname
	Bio               *valueobject.Bio
	PreferredLanguage valueobject.LanguageCode
	ProfileDisovery   enum.ProfileDiscoveryEnum
	AvatarUrl         *valueobject.HttpUrl
	EditingLockedAt   *time.Time
	CreatedAt         time.Time
	DeletedAt         *time.Time
}

func (a *AppUserProfile) IsValid() bool {
	return a.AppUserGUID.IsValid() && a.Name.IsValid() && a.Surname.IsValid() && a.Patronymic.IsValid() && a.Nickname.IsValid() && a.Bio.IsValid() && a.PreferredLanguage.IsValid() &&
		a.ProfileDisovery.IsValid() && a.AvatarUrl.IsValid()
}
