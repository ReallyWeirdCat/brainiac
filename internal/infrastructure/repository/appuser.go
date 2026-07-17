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

type PgAppUserRepo struct {
	queries *generated.Queries
	mu      *sync.Mutex
}

func NewPgAppUserRepo(queries *generated.Queries, mu *sync.Mutex) repo.AppUserRepository {
	if queries == nil || mu == nil {
		panic("queries and mu must not be nil")
	}
	return &PgAppUserRepo{
		mu:      mu,
		queries: queries,
	}
}

func (p *PgAppUserRepo) GetByEmail(ctx context.Context, email valueobject.Email) (*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserByEmail(ctx, &email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(obj)
	return &model, err
}

func (p *PgAppUserRepo) GetByUsername(ctx context.Context, username valueobject.Username) (*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(obj)
	return &model, err
}

func (p *PgAppUserRepo) ExistsByUsername(ctx context.Context, username valueobject.Username) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.queries.ExistsAppUserByUsername(ctx, username)
}

func (p *PgAppUserRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.IsDeletedAppUser(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, repo.ErrEntityNotFound.FromError(err)
		}
		return false, err
	}
	return obj, nil
}

func (p *PgAppUserRepo) Count(ctx context.Context) (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	count, err := p.queries.CountAppUsers(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PgAppUserRepo) Create(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.CreateAppUserParams{
		GUID:        obj.GUID,
		Username:    obj.Username,
		ActivatedAt: obj.ActivatedAt,
	}
	newObj, err := p.queries.CreateAppUser(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUser(newObj)
	return &model, nil
}

func (p *PgAppUserRepo) CreateBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.CreateAppUserBatchParams
	for _, obj := range objs {
		params = append(params, generated.CreateAppUserBatchParams{
			GUID:        obj.GUID,
			Username:    obj.Username,
			ActivatedAt: obj.ActivatedAt,
		})
	}
	batch := p.queries.CreateAppUserBatch(ctx, params)

	var models []*entity.AppUser
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, newObj := range items {
			model := entity.AppUser(newObj)
			models = append(models, &model)
		}
	})

	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.queries.DeleteAppUser(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.ErrEntityNotFound.FromError(err)
		}
		return err
	}
	return nil
}

func (p *PgAppUserRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.DeleteAppUserBatch(ctx, guids)
	var batchErrs []error
	batch.Exec(func(_ int, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
	})
	if len(batchErrs) != 0 {
		return errors.Join(batchErrs...)
	}
	return nil
}

func (p *PgAppUserRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.queries.ExistsAppUser(ctx, guid)
}

func (p *PgAppUserRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.ExistsAppUserBatch(ctx, guids)
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

func (p *PgAppUserRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUser(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(obj)
	return &model, nil
}

func (p *PgAppUserRepo) GetAll(ctx context.Context) ([]*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var models []*entity.AppUser
	objs, err := p.queries.GetAllAppUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models, nil
		}
		return nil, err
	}
	for _, obj := range objs {
		model := entity.AppUser(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.GetAppUserBatch(ctx, guids)
	var models []*entity.AppUser
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUser(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserRepo) Save(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.SaveAppUserParams{
		GUID:        obj.GUID,
		Username:    obj.Username,
		ActivatedAt: obj.ActivatedAt,
	}
	newObj, err := p.queries.SaveAppUser(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUser(newObj)
	return &model, nil
}

func (p *PgAppUserRepo) SaveBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.SaveAppUserBatchParams
	for _, obj := range objs {
		params = append(params, generated.SaveAppUserBatchParams{
			GUID:        obj.GUID,
			Username:    obj.Username,
			ActivatedAt: obj.ActivatedAt,
		})
	}
	batch := p.queries.SaveAppUserBatch(ctx, params)
	var models []*entity.AppUser
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUser(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserRepo) Update(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.UpdateAppUserParams{
		GUID:        obj.GUID,
		Username:    obj.Username,
		ActivatedAt: obj.ActivatedAt,
	}
	newObj, err := p.queries.UpdateAppUser(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(newObj)
	return &model, nil
}

func (p *PgAppUserRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	var params []generated.UpdateAppUserBatchParams
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, obj := range objs {
		params = append(params, generated.UpdateAppUserBatchParams{
			GUID:        obj.GUID,
			Username:    obj.Username,
			ActivatedAt: obj.ActivatedAt,
		})
	}
	batch := p.queries.UpdateAppUserBatch(ctx, params)
	var models []*entity.AppUser
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUser(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

var _ repo.AppUserRepository = &PgAppUserRepo{}
var _ entity.AppUser = entity.AppUser(generated.AppUser{})
