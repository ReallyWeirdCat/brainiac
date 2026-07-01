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

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type PgAppUserRepo struct {
	Queries *generated.Queries
}

func (p *PgAppUserRepo) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) Create(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) CreateBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) GetAll(ctx context.Context) ([]*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) Save(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) SaveBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) Update(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	panic("unimplemented")
}

func (p *PgAppUserRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	panic("unimplemented")
}

var _ repo.AppUserRepository = &PgAppUserRepo{}
var _ entity.AppUser = entity.AppUser(generated.AppUser{})
