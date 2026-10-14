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

// Repository is the minimal persistence contract shared by every entity.
//
// It deliberately excludes batch, list-all, count and existence checks: those
// are rarely needed on a per-aggregate basis and, when they are, should be
// expressed as bespoke methods on the entity's own repository interface (or a
// separate admin/analytics interface) rather than forced onto every entity.
type Repository[T any] interface {
	Get(ctx context.Context, guid valueobject.GUID) (*T, error)
	Create(ctx context.Context, obj *T) (*T, error)
	Save(ctx context.Context, obj *T) (*T, error)
	Update(ctx context.Context, obj *T) (*T, error)
	Delete(ctx context.Context, guid valueobject.GUID) error
}
