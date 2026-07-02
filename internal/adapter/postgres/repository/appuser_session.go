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
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type PgAppUserSessionRepo struct {
	queries *generated.Queries
	mu      *sync.Mutex
}

func NewPgAppUserSessionRepo(queries *generated.Queries, mu *sync.Mutex) repository.AppUserSessionRepository {
	if queries == nil || mu == nil {
		panic("queries and mu must not be nil")
	}
	return &PgAppUserSessionRepo{
		mu:      mu,
		queries: queries,
	}
}

func (p *PgAppUserSessionRepo) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) Create(ctx context.Context, obj *entity.AppUserSession) (*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserSession) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) GetAll(ctx context.Context) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) GetAllActiveByUsername(ctx context.Context, username valueobject.Username) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) GetAllAppUserGUIDsByIP(ctx context.Context, ip string) ([]valueobject.GUID, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) GetAllByLastIP(ctx context.Context, ip string) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) GetAllInactiveByUsername(ctx context.Context, username valueobject.Username) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) Save(ctx context.Context, obj *entity.AppUserSession) (*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserSession) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) Update(ctx context.Context, obj *entity.AppUserSession) (*entity.AppUserSession, error) {
	panic("unimplemented")
}

func (p *PgAppUserSessionRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserSession) ([]*entity.AppUserSession, error) {
	panic("unimplemented")
}

var _ repository.AppUserSessionRepository = &PgAppUserSessionRepo{}
var _ entity.AppUserSession = entity.AppUserSession(generated.AppUserSession{})
