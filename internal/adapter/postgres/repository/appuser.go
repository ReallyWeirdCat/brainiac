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

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	domerr "github.com/ReallyWeirdCat/brainiac/pkg/domain/errors"
	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
	"github.com/jackc/pgx/v5"
)

type PgAppUserRepo struct {
	Queries *generated.Queries
}

func (p *PgAppUserRepo) GetByEmail(ctx context.Context, email valueobject.Email) (*entity.AppUser, error) {
	obj, err := p.Queries.GetAppUserByEmail(ctx, &email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(obj)
	return &model, err
}

func (p *PgAppUserRepo) GetByUsername(ctx context.Context, username valueobject.Username) (*entity.AppUser, error) {
	obj, err := p.Queries.GetAppUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(obj)
	return &model, err
}

func (p *PgAppUserRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	obj, err := p.Queries.IsDeletedAppUser(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, domerr.ErrEntityNotFound.FromError(err)
		}
		return false, err
	}
	return obj, nil
}

func (p *PgAppUserRepo) Count(ctx context.Context) (int64, error) {
	count, err := p.Queries.CountAppUsers(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PgAppUserRepo) Create(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	params := generated.CreateAppUserParams{
		GUID:        obj.GUID,
		Username:    obj.Username,
		ActivatedAt: obj.ActivatedAt,
		DeletedAt:   obj.DeletedAt,
	}
	newObj, err := p.Queries.CreateAppUser(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUser(newObj)
	return &model, nil
}

func (p *PgAppUserRepo) CreateBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	var params []generated.CreateAppUserBatchParams
	for _, obj := range objs {
		params = append(params, generated.CreateAppUserBatchParams{
			GUID:        obj.GUID,
			Username:    obj.Username,
			ActivatedAt: obj.ActivatedAt,
			DeletedAt:   obj.DeletedAt,
		})
	}
	batch := p.Queries.CreateAppUserBatch(ctx, params)

	var models []*entity.AppUser
	var batchErr error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			batchErr = err
			return
		}
		for _, newObj := range items {
			model := entity.AppUser(newObj)
			models = append(models, &model)
		}
	})

	if batchErr != nil {
		return nil, batchErr
	}
	return models, nil
}

func (p *PgAppUserRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	return p.Queries.DeleteAppUser(ctx, guid)
}

func (p *PgAppUserRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	batch := p.Queries.DeleteAppUserBatch(ctx, guids)
	var batchErr error
	batch.Exec(func(_ int, err error) {
		if err != nil {
			batchErr = err
			return
		}
	})
	if batchErr != nil {
		return batchErr
	}
	return nil
}

func (p *PgAppUserRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	return p.Queries.ExistsAppUser(ctx, guid)
}

func (p *PgAppUserRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	batch := p.Queries.ExistsAppUserBatch(ctx, guids)
	var existingGuids []valueobject.GUID
	var batchErr error
	batch.Query(func(idx int, items []valueobject.GUID, err error) {
		if err != nil {
			batchErr = err
			return
		}
		existingGuids = items
	})
	if batchErr != nil {
		return nil, batchErr
	}
	return existingGuids, nil
}

func (p *PgAppUserRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUser, error) {
	obj, err := p.Queries.GetAppUser(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domerr.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUser(obj)
	return &model, nil
}

func (p *PgAppUserRepo) GetAll(ctx context.Context) ([]*entity.AppUser, error) {
	var models []*entity.AppUser
	objs, err := p.Queries.GetAllAppUsers(ctx)
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
	batch := p.Queries.GetAppUserBatch(ctx, guids)
	var models []*entity.AppUser
	var batchErr error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			batchErr = err
			return
		}
		for _, item := range items {
			model := entity.AppUser(item)
			models = append(models, &model)
		}
	})
	if batchErr != nil {
		return nil, batchErr
	}
	return models, nil
}

func (p *PgAppUserRepo) Save(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	params := generated.SaveAppUserParams{
		GUID:        obj.GUID,
		Username:    obj.Username,
		ActivatedAt: obj.ActivatedAt,
		DeletedAt:   obj.DeletedAt,
	}
	newObj, err := p.Queries.SaveAppUser(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUser(newObj)
	return &model, nil
}

func (p *PgAppUserRepo) SaveBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	var params []generated.SaveAppUserBatchParams
	for _, obj := range objs {
		params = append(params, generated.SaveAppUserBatchParams{
			GUID:        obj.GUID,
			Username:    obj.Username,
			ActivatedAt: obj.ActivatedAt,
			DeletedAt:   obj.DeletedAt,
		})
	}
	batch := p.Queries.SaveAppUserBatch(ctx, params)
	var models []*entity.AppUser
	var batchErr error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			batchErr = err
			return
		}
		for _, item := range items {
			model := entity.AppUser(item)
			models = append(models, &model)
		}
	})
	if batchErr != nil {
		return nil, batchErr
	}
	return models, nil
}

func (p *PgAppUserRepo) Update(ctx context.Context, obj *entity.AppUser) (*entity.AppUser, error) {
	params := generated.UpdateAppUserParams{
		GUID:        obj.GUID,
		Username:    obj.Username,
		ActivatedAt: obj.ActivatedAt,
		DeletedAt:   obj.DeletedAt,
	}
	newObj, err := p.Queries.UpdateAppUser(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUser(newObj)
	return &model, nil
}

func (p *PgAppUserRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUser) ([]*entity.AppUser, error) {
	var params []generated.UpdateAppUserBatchParams
	for _, obj := range objs {
		params = append(params, generated.UpdateAppUserBatchParams{
			GUID:        obj.GUID,
			Username:    obj.Username,
			ActivatedAt: obj.ActivatedAt,
			DeletedAt:   obj.DeletedAt,
		})
	}
	batch := p.Queries.UpdateAppUserBatch(ctx, params)
	var models []*entity.AppUser
	var batchErr error
	batch.Query(func(idx int, items []generated.AppUser, err error) {
		if err != nil {
			batchErr = err
			return
		}
		for _, item := range items {
			model := entity.AppUser(item)
			models = append(models, &model)
		}
	})
	if batchErr != nil {
		return nil, batchErr
	}
	return models, nil
}

var _ repo.AppUserRepository = &PgAppUserRepo{}
var _ entity.AppUser = entity.AppUser(generated.AppUser{})
