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

type PgAppUserCredentialRepo struct {
	queries *generated.Queries
	mu      *sync.Mutex
}

func NewPgAppUserCredentialRepo(queries *generated.Queries, mu *sync.Mutex) repo.AppUserCredentialRepository {
	if queries == nil || mu == nil {
		panic("queries and mu must not be nil")
	}
	return &PgAppUserCredentialRepo{
		mu:      mu,
		queries: queries,
	}
}

func (p *PgAppUserCredentialRepo) GetByEmail(ctx context.Context, email valueobject.Email) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Create(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) GetAll(ctx context.Context) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Save(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) Update(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

func (p *PgAppUserCredentialRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

var _ repo.AppUserCredentialRepository = &PgAppUserCredentialRepo{}
var _ entity.AppUserCredential = entity.AppUserCredential(generated.AppUserCredential{})
