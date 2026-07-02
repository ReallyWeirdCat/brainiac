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
	"sync"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type PgAppUserProfileRepo struct {
	queries *generated.Queries
	mu      *sync.Mutex
}

func NewPgAppUserProfileRepo(queries *generated.Queries, mu *sync.Mutex) repo.AppUserProfileRepository {
	if queries == nil || mu == nil {
		panic("queries and mu must not be nil")
	}
	return &PgAppUserProfileRepo{
		mu:      mu,
		queries: queries,
	}
}

func (p *PgAppUserProfileRepo) GetByUsername(ctx context.Context, username valueobject.Username) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Create(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) GetAll(ctx context.Context) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Save(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) Update(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

func (p *PgAppUserProfileRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

var _ repo.AppUserProfileRepository = &PgAppUserProfileRepo{}
var _ entity.AppUserCredential = entity.AppUserCredential(generated.AppUserCredential{})
