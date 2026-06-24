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

type InventorySlot struct {
	GUID             valueobject.GUID `json:"guid"`
	InventoryGUID    valueobject.GUID `json:"inventory_guid"`
	ItemGUID         valueobject.GUID `json:"item_guid"`
	Amount           int16            `json:"amount"`
	AmountInExchange int16            `json:"amount_in_exchange"`
	ExpireAt         *time.Time       `json:"expire_at,omitempty"`
	CreatedAt        time.Time        `json:"created_at"`
	DeletedAt        *time.Time       `json:"deleted_at,omitempty"`
}

var _ Entity = &InventorySlot{}

func (i *InventorySlot) IsValid() bool {
	if !i.GUID.IsValid() {
		return false
	}
	if !i.InventoryGUID.IsValid() {
		return false
	}
	if !i.ItemGUID.IsValid() {
		return false
	}
	if i.Amount < 0 || i.AmountInExchange < 0 {
		return false
	}
	if i.CreatedAt.IsZero() {
		return false
	}
	return true
}
