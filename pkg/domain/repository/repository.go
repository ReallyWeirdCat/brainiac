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

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type Repository[T any] interface {
	Create(ctx context.Context, obj *T) (*T, error)
	CreateBatch(ctx context.Context, objs []*T) ([]*T, error)
	Update(ctx context.Context, obj *T) (*T, error)
	UpdateBatch(ctx context.Context, objs []*T) ([]*T, error)
	Save(ctx context.Context, obj *T) (*T, error)
	SaveBatch(ctx context.Context, objs []*T) ([]*T, error)
	Delete(ctx context.Context, guid valueobject.GUID) error
	DeleteBatch(ctx context.Context, guids []valueobject.GUID) error
	Get(ctx context.Context, guid valueobject.GUID) (*T, error)
	GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*T, error)
	GetAll(ctx context.Context) ([]*T, error)
	Count(ctx context.Context) (int64, error)
	Exists(ctx context.Context, guid valueobject.GUID) (bool, error)
	ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error)
	IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error)
}
