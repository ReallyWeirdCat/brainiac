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
	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
	"github.com/jackc/pgx/v5"
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
		queries: queries,
		mu:      mu,
	}
}

func (p *PgAppUserCredentialRepo) GetByAppUserGUID(ctx context.Context, appUserGUID valueobject.GUID) (*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserCredential(ctx, appUserGUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserCredential(obj)
	return &model, nil
}

func (p *PgAppUserCredentialRepo) GetByEmail(ctx context.Context, email valueobject.Email) (*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserCredentialByEmail(ctx, &email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserCredential(obj)
	return &model, nil
}

func (p *PgAppUserCredentialRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.IsDeletedAppUserCredential(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, domerr.ErrEntityNotFound.FromError(err)
		}
		return false, err
	}
	return obj, nil
}

func (p *PgAppUserCredentialRepo) Count(ctx context.Context) (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	count, err := p.queries.CountAppUserCredentials(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PgAppUserCredentialRepo) Create(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.CreateAppUserCredentialParams{
		AppUserGUID:  obj.AppUserGUID,
		Email:        obj.Email,
		PasswordHash: obj.PasswordHash,
	}
	newObj, err := p.queries.CreateAppUserCredential(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUserCredential(newObj)
	return &model, nil
}

func (p *PgAppUserCredentialRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.CreateAppUserCredentialBatchParams
	for _, obj := range objs {
		params = append(params, generated.CreateAppUserCredentialBatchParams{
			AppUserGUID:  obj.AppUserGUID,
			Email:        obj.Email,
			PasswordHash: obj.PasswordHash,
		})
	}
	batch := p.queries.CreateAppUserCredentialBatch(ctx, params)

	var models []*entity.AppUserCredential
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserCredential, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, newObj := range items {
			model := entity.AppUserCredential(newObj)
			models = append(models, &model)
		}
	})

	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserCredentialRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.queries.DeleteAppUserCredential(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domerr.ErrEntityNotFound.FromError(err)
		}
		return err
	}
	return nil
}

func (p *PgAppUserCredentialRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.DeleteAppUserCredentialBatch(ctx, guids)
	var batchErrs []error
	batch.Exec(func(_ int, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, domerr.ErrEntityNotFound.FromError(err))
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

func (p *PgAppUserCredentialRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.queries.ExistsAppUserCredential(ctx, guid)
}

func (p *PgAppUserCredentialRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.ExistsAppUserCredentialBatch(ctx, guids)
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

func (p *PgAppUserCredentialRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserCredential(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserCredential(obj)
	return &model, nil
}

func (p *PgAppUserCredentialRepo) GetAll(ctx context.Context) ([]*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var models []*entity.AppUserCredential
	objs, err := p.queries.GetAllAppUserCredentials(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models, nil
		}
		return nil, err
	}
	for _, obj := range objs {
		model := entity.AppUserCredential(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserCredentialRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.GetAppUserCredentialBatch(ctx, guids)
	var models []*entity.AppUserCredential
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserCredential, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, domerr.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserCredential(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserCredentialRepo) Save(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.SaveAppUserCredentialParams{
		AppUserGUID:  obj.AppUserGUID,
		Email:        obj.Email,
		PasswordHash: obj.PasswordHash,
	}
	newObj, err := p.queries.SaveAppUserCredential(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUserCredential(newObj)
	return &model, nil
}

func (p *PgAppUserCredentialRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.SaveAppUserCredentialBatchParams
	for _, obj := range objs {
		params = append(params, generated.SaveAppUserCredentialBatchParams{
			AppUserGUID:  obj.AppUserGUID,
			Email:        obj.Email,
			PasswordHash: obj.PasswordHash,
		})
	}
	batch := p.queries.SaveAppUserCredentialBatch(ctx, params)
	var models []*entity.AppUserCredential
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserCredential, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserCredential(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserCredentialRepo) Update(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.UpdateAppUserCredentialParams{
		AppUserGUID:  obj.AppUserGUID,
		Email:        obj.Email,
		PasswordHash: obj.PasswordHash,
	}
	newObj, err := p.queries.UpdateAppUserCredential(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserCredential(newObj)
	return &model, nil
}

func (p *PgAppUserCredentialRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.UpdateAppUserCredentialBatchParams
	for _, obj := range objs {
		params = append(params, generated.UpdateAppUserCredentialBatchParams{
			AppUserGUID:  obj.AppUserGUID,
			Email:        obj.Email,
			PasswordHash: obj.PasswordHash,
		})
	}
	batch := p.queries.UpdateAppUserCredentialBatch(ctx, params)
	var models []*entity.AppUserCredential
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserCredential, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, domerr.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserCredential(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

var _ repo.AppUserCredentialRepository = &PgAppUserCredentialRepo{}
var _ entity.AppUserCredential = entity.AppUserCredential(generated.AppUserCredential{})
