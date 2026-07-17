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

type Notification struct {
	GUID        valueobject.GUID      `json:"guid"`
	AppUserGUID valueobject.GUID      `json:"app_user_guid"`
	TitleI18n   *valueobject.I18nText `json:"title_i18n,omitempty"`
	ContentI18n *valueobject.I18nText `json:"content_i18n,omitempty"`
	ResourceURL *valueobject.HttpUrl  `json:"resource_url,omitempty"`
	Urgency     enum.UrgencyEnum      `json:"urgency"`
	Meta        *valueobject.Metadata `json:"meta,omitempty"`
	SeenAt      *time.Time            `json:"seen_at,omitempty"`
	ExpireAt    *time.Time            `json:"expire_at,omitempty"`
	CreatedAt   time.Time             `json:"created_at"`
	DeletedAt   *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = &Notification{}

func (n *Notification) IsValid() bool {
	if !n.GUID.IsValid() || !n.AppUserGUID.IsValid() || !n.Urgency.IsValid() {
		return false
	}
	if n.TitleI18n != nil && !n.TitleI18n.IsValid() {
		return false
	}
	if n.ContentI18n != nil && !n.ContentI18n.IsValid() {
		return false
	}
	if n.ResourceURL != nil && !n.ResourceURL.IsValid() {
		return false
	}
	if n.Meta != nil && !n.Meta.IsValid() {
		return false
	}
	if n.CreatedAt.IsZero() {
		return false
	}

	return true
}
