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

package repository

import (
	"context"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/enum"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type ItemRepository interface {
	Save(ctx context.Context, item entity.Item) error
	Delete(ctx context.Context, guid valueobject.GUID) error
	GetByGUID(ctx context.Context, guid valueobject.GUID) (*entity.Item, error)
	GetByRarity(ctx context.Context, rarity enum.RarityEnum) ([]*entity.Item, error)
	GetPublished(ctx context.Context) ([]*entity.Item, error)
	ExistsByGUID(ctx context.Context, guid valueobject.GUID) (bool, error)
}
