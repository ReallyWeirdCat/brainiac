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

type Item struct {
	GUID            valueobject.GUID      `json:"guid"`
	TitleI18n       valueobject.I18nText  `json:"title_i18n"`
	DescriptionI18n *valueobject.I18nText `json:"description_i18n,omitempty"`
	ResourceURL     *valueobject.HttpUrl  `json:"resource_url,omitempty"`
	Meta            valueobject.Metadata  `json:"meta"`
	StackSize       int16                 `json:"stack_size"`
	Rarity          enum.RarityEnum       `json:"rarity"`
	AllowExchange   bool                  `json:"allow_exchange"`
	ShopPrice       *int64                `json:"shop_price,omitempty"`
	OneTimePurchase bool                  `json:"one_time_purchase"`
	ShopQuantity    *int16                `json:"shop_quantity,omitempty"`
	LevelRequired   *int16                `json:"level_required,omitempty"`
	OnSaleSince     *time.Time            `json:"on_sale_since,omitempty"`
	OnSaleUntil     *time.Time            `json:"on_sale_until,omitempty"`
	PublishedAt     *time.Time            `json:"published_at,omitempty"`
	CreatedAt       time.Time             `json:"created_at"`
	DeletedAt       *time.Time            `json:"deleted_at,omitempty"`
}

var _ Entity = &Item{}

func (i *Item) IsValid() bool {
	if !i.GUID.IsValid() {
		return false
	}
	if !i.TitleI18n.IsValid() {
		return false
	}
	if i.DescriptionI18n != nil && !i.DescriptionI18n.IsValid() {
		return false
	}
	if i.ResourceURL != nil && !i.ResourceURL.IsValid() {
		return false
	}
	if !i.Meta.IsValid() {
		return false
	}
	if i.StackSize < 1 {
		return false
	}
	if !i.Rarity.IsValid() {
		return false
	}
	if i.ShopPrice != nil && *i.ShopPrice < 0 {
		return false
	}
	if i.ShopQuantity != nil && *i.ShopQuantity < 0 {
		return false
	}
	if i.LevelRequired != nil && *i.LevelRequired < 0 {
		return false
	}
	if i.CreatedAt.IsZero() {
		return false
	}
	return true
}
