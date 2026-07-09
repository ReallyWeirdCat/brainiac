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
	"errors"
	"sync"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
	"github.com/jackc/pgx/v5"
)

type PgAppUserSessionRepo struct {
	queries *generated.Queries
	mu      *sync.Mutex
}

func NewPgAppUserSessionRepo(queries *generated.Queries, mu *sync.Mutex) repo.AppUserSessionRepository {
	if queries == nil || mu == nil {
		panic("queries and mu must not be nil")
	}
	return &PgAppUserSessionRepo{
		queries: queries,
		mu:      mu,
	}
}

func (p *PgAppUserSessionRepo) Create(ctx context.Context, obj *entity.AppUserSession) (*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.CreateAppUserSessionParams{
		GUID:        obj.GUID,
		AppUserGUID: obj.AppUserGUID,
		LastIPV4:    obj.LastIPV4,
		LastIPV6:    obj.LastIPV6,
		LastAgent:   obj.LastAgent,
		LastSeenAt:  obj.LastSeenAt,
		ExpireAt:    obj.ExpireAt,
	}
	newObj, err := p.queries.CreateAppUserSession(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUserSession(newObj)
	return &model, nil
}

func (p *PgAppUserSessionRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserSession) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.CreateAppUserSessionBatchParams
	for _, obj := range objs {
		params = append(params, generated.CreateAppUserSessionBatchParams{
			GUID:        obj.GUID,
			AppUserGUID: obj.AppUserGUID,
			LastIPV4:    obj.LastIPV4,
			LastIPV6:    obj.LastIPV6,
			LastAgent:   obj.LastAgent,
			LastSeenAt:  obj.LastSeenAt,
			ExpireAt:    obj.ExpireAt,
		})
	}
	batch := p.queries.CreateAppUserSessionBatch(ctx, params)

	var models []*entity.AppUserSession
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserSession, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserSession(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) Update(ctx context.Context, obj *entity.AppUserSession) (*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.UpdateAppUserSessionParams{
		GUID:        obj.GUID,
		AppUserGUID: obj.AppUserGUID,
		LastIPV4:    obj.LastIPV4,
		LastIPV6:    obj.LastIPV6,
		LastAgent:   obj.LastAgent,
		LastSeenAt:  obj.LastSeenAt,
		ExpireAt:    obj.ExpireAt,
	}
	newObj, err := p.queries.UpdateAppUserSession(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserSession(newObj)
	return &model, nil
}

func (p *PgAppUserSessionRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserSession) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.UpdateAppUserSessionBatchParams
	for _, obj := range objs {
		params = append(params, generated.UpdateAppUserSessionBatchParams{
			GUID:        obj.GUID,
			AppUserGUID: obj.AppUserGUID,
			LastIPV4:    obj.LastIPV4,
			LastIPV6:    obj.LastIPV6,
			LastAgent:   obj.LastAgent,
			LastSeenAt:  obj.LastSeenAt,
			ExpireAt:    obj.ExpireAt,
		})
	}
	batch := p.queries.UpdateAppUserSessionBatch(ctx, params)

	var models []*entity.AppUserSession
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserSession, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserSession(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) Save(ctx context.Context, obj *entity.AppUserSession) (*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.SaveAppUserSessionParams{
		GUID:        obj.GUID,
		AppUserGUID: obj.AppUserGUID,
		LastIPV4:    obj.LastIPV4,
		LastIPV6:    obj.LastIPV6,
		LastAgent:   obj.LastAgent,
		LastSeenAt:  obj.LastSeenAt,
		ExpireAt:    obj.ExpireAt,
	}
	newObj, err := p.queries.SaveAppUserSession(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUserSession(newObj)
	return &model, nil
}

func (p *PgAppUserSessionRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserSession) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.SaveAppUserSessionBatchParams
	for _, obj := range objs {
		params = append(params, generated.SaveAppUserSessionBatchParams{
			GUID:        obj.GUID,
			AppUserGUID: obj.AppUserGUID,
			LastIPV4:    obj.LastIPV4,
			LastIPV6:    obj.LastIPV6,
			LastAgent:   obj.LastAgent,
			LastSeenAt:  obj.LastSeenAt,
			ExpireAt:    obj.ExpireAt,
		})
	}
	batch := p.queries.SaveAppUserSessionBatch(ctx, params)

	var models []*entity.AppUserSession
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserSession, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserSession(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.queries.DeleteAppUserSession(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.ErrEntityNotFound.FromError(err)
		}
		return err
	}
	return nil
}

func (p *PgAppUserSessionRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.DeleteAppUserSessionBatch(ctx, guids)
	var batchErrs []error
	batch.Exec(func(_ int, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
		}
	})
	if len(batchErrs) != 0 {
		return errors.Join(batchErrs...)
	}
	return nil
}

func (p *PgAppUserSessionRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserSession(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserSession(obj)
	return &model, nil
}

func (p *PgAppUserSessionRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.GetAppUserSessionBatch(ctx, guids)
	var models []*entity.AppUserSession
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserSession, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserSession(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) GetAll(ctx context.Context) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var models []*entity.AppUserSession
	objs, err := p.queries.GetAllAppUserSessions(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models, nil
		}
		return nil, err
	}
	for _, obj := range objs {
		model := entity.AppUserSession(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) Count(ctx context.Context) (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	count, err := p.queries.CountAppUserSessions(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PgAppUserSessionRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.queries.ExistsAppUserSession(ctx, guid)
}

func (p *PgAppUserSessionRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.ExistsAppUserSessionBatch(ctx, guids)
	var existingGuids []valueobject.GUID
	var batchErrs []error
	batch.Query(func(idx int, items []valueobject.GUID, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		existingGuids = append(existingGuids, items...)
	})
	if len(batchErrs) != 0 {
		return existingGuids, errors.Join(batchErrs...)
	}
	return existingGuids, nil
}

func (p *PgAppUserSessionRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	deleted, err := p.queries.IsDeletedAppUserSession(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, repo.ErrEntityNotFound.FromError(err)
		}
		return false, err
	}
	return deleted, nil
}

func (p *PgAppUserSessionRepo) GetAllActiveByUsername(ctx context.Context, username valueobject.Username) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	objs, err := p.queries.GetAllActiveSessionsByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entity.AppUserSession{}, nil
		}
		return nil, err
	}
	var models []*entity.AppUserSession
	for _, obj := range objs {
		model := entity.AppUserSession(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) GetAllInactiveByUsername(ctx context.Context, username valueobject.Username) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	objs, err := p.queries.GetAllInactiveSessionsByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entity.AppUserSession{}, nil
		}
		return nil, err
	}
	var models []*entity.AppUserSession
	for _, obj := range objs {
		model := entity.AppUserSession(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) GetAllByLastIP(ctx context.Context, ip string) ([]*entity.AppUserSession, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	objs, err := p.queries.GetAllSessionsByLastIP(ctx, &ip)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entity.AppUserSession{}, nil
		}
		return nil, err
	}
	var models []*entity.AppUserSession
	for _, obj := range objs {
		model := entity.AppUserSession(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserSessionRepo) GetAllAppUserGUIDsByIP(ctx context.Context, ip string) ([]valueobject.GUID, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	guids, err := p.queries.GetAllAppUserGUIDsByIP(ctx, &ip)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []valueobject.GUID{}, nil
		}
		return nil, err
	}
	return guids, nil
}

var _ repo.AppUserSessionRepository = &PgAppUserSessionRepo{}
var _ entity.AppUserSession = entity.AppUserSession(generated.AppUserSession{})
