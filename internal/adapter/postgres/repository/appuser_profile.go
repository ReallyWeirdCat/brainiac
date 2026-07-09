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
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserProfileByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserProfile(obj)
	return &model, nil
}

func (p *PgAppUserProfileRepo) IsDeleted(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.IsDeletedAppUserProfile(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, repo.ErrEntityNotFound.FromError(err)
		}
		return false, err
	}
	return obj, nil
}

func (p *PgAppUserProfileRepo) Count(ctx context.Context) (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	count, err := p.queries.CountAppUserProfiles(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PgAppUserProfileRepo) Create(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.CreateAppUserProfileParams{
		AppUserGUID:       obj.AppUserGUID,
		Name:              obj.Name,
		Surname:           obj.Surname,
		Patronymic:        obj.Patronymic,
		Nickname:          obj.Nickname,
		Bio:               obj.Bio,
		PreferredLanguage: obj.PreferredLanguage,
		ProfileDiscovery:  obj.ProfileDiscovery,
		AvatarUrl:         obj.AvatarUrl,
		EditingLockedAt:   obj.EditingLockedAt,
	}
	newObj, err := p.queries.CreateAppUserProfile(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUserProfile(newObj)
	return &model, nil
}

func (p *PgAppUserProfileRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.CreateAppUserProfileBatchParams
	for _, obj := range objs {
		params = append(params, generated.CreateAppUserProfileBatchParams{
			AppUserGUID:       obj.AppUserGUID,
			Name:              obj.Name,
			Surname:           obj.Surname,
			Patronymic:        obj.Patronymic,
			Nickname:          obj.Nickname,
			Bio:               obj.Bio,
			PreferredLanguage: obj.PreferredLanguage,
			ProfileDiscovery:  obj.ProfileDiscovery,
			AvatarUrl:         obj.AvatarUrl,
			EditingLockedAt:   obj.EditingLockedAt,
		})
	}
	batch := p.queries.CreateAppUserProfileBatch(ctx, params)

	var models []*entity.AppUserProfile
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserProfile, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, newObj := range items {
			model := entity.AppUserProfile(newObj)
			models = append(models, &model)
		}
	})

	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserProfileRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.queries.DeleteAppUserProfile(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.ErrEntityNotFound.FromError(err)
		}
		return err
	}
	return nil
}

func (p *PgAppUserProfileRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.DeleteAppUserProfileBatch(ctx, guids)
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

func (p *PgAppUserProfileRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.queries.ExistsAppUserProfile(ctx, guid)
}

func (p *PgAppUserProfileRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) ([]valueobject.GUID, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.ExistsAppUserProfileBatch(ctx, guids)
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

func (p *PgAppUserProfileRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	obj, err := p.queries.GetAppUserProfile(ctx, guid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserProfile(obj)
	return &model, nil
}

func (p *PgAppUserProfileRepo) GetAll(ctx context.Context) ([]*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var models []*entity.AppUserProfile
	objs, err := p.queries.GetAllAppUserProfiles(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models, nil
		}
		return nil, err
	}
	for _, obj := range objs {
		model := entity.AppUserProfile(obj)
		models = append(models, &model)
	}
	return models, nil
}

func (p *PgAppUserProfileRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	batch := p.queries.GetAppUserProfileBatch(ctx, guids)
	var models []*entity.AppUserProfile
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserProfile, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserProfile(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserProfileRepo) Save(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.SaveAppUserProfileParams{
		AppUserGUID:       obj.AppUserGUID,
		Name:              obj.Name,
		Surname:           obj.Surname,
		Patronymic:        obj.Patronymic,
		Nickname:          obj.Nickname,
		Bio:               obj.Bio,
		PreferredLanguage: obj.PreferredLanguage,
		ProfileDiscovery:  obj.ProfileDiscovery,
		AvatarUrl:         obj.AvatarUrl,
		EditingLockedAt:   obj.EditingLockedAt,
	}
	newObj, err := p.queries.SaveAppUserProfile(ctx, params)
	if err != nil {
		return nil, err
	}
	model := entity.AppUserProfile(newObj)
	return &model, nil
}

func (p *PgAppUserProfileRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.SaveAppUserProfileBatchParams
	for _, obj := range objs {
		params = append(params, generated.SaveAppUserProfileBatchParams{
			AppUserGUID:       obj.AppUserGUID,
			Name:              obj.Name,
			Surname:           obj.Surname,
			Patronymic:        obj.Patronymic,
			Nickname:          obj.Nickname,
			Bio:               obj.Bio,
			PreferredLanguage: obj.PreferredLanguage,
			ProfileDiscovery:  obj.ProfileDiscovery,
			AvatarUrl:         obj.AvatarUrl,
			EditingLockedAt:   obj.EditingLockedAt,
		})
	}
	batch := p.queries.SaveAppUserProfileBatch(ctx, params)
	var models []*entity.AppUserProfile
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserProfile, err error) {
		if err != nil {
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserProfile(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

func (p *PgAppUserProfileRepo) Update(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	params := generated.UpdateAppUserProfileParams{
		AppUserGUID:       obj.AppUserGUID,
		Name:              obj.Name,
		Surname:           obj.Surname,
		Patronymic:        obj.Patronymic,
		Nickname:          obj.Nickname,
		Bio:               obj.Bio,
		PreferredLanguage: obj.PreferredLanguage,
		ProfileDiscovery:  obj.ProfileDiscovery,
		AvatarUrl:         obj.AvatarUrl,
		EditingLockedAt:   obj.EditingLockedAt,
	}
	newObj, err := p.queries.UpdateAppUserProfile(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrEntityNotFound.FromError(err)
		}
		return nil, err
	}
	model := entity.AppUserProfile(newObj)
	return &model, nil
}

func (p *PgAppUserProfileRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var params []generated.UpdateAppUserProfileBatchParams
	for _, obj := range objs {
		params = append(params, generated.UpdateAppUserProfileBatchParams{
			AppUserGUID:       obj.AppUserGUID,
			Name:              obj.Name,
			Surname:           obj.Surname,
			Patronymic:        obj.Patronymic,
			Nickname:          obj.Nickname,
			Bio:               obj.Bio,
			PreferredLanguage: obj.PreferredLanguage,
			ProfileDiscovery:  obj.ProfileDiscovery,
			AvatarUrl:         obj.AvatarUrl,
			EditingLockedAt:   obj.EditingLockedAt,
		})
	}
	batch := p.queries.UpdateAppUserProfileBatch(ctx, params)
	var models []*entity.AppUserProfile
	var batchErrs []error
	batch.Query(func(idx int, items []generated.AppUserProfile, err error) {
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				batchErrs = append(batchErrs, repo.ErrEntityNotFound.FromError(err))
				return
			}
			batchErrs = append(batchErrs, err)
			return
		}
		for _, item := range items {
			model := entity.AppUserProfile(item)
			models = append(models, &model)
		}
	})
	if len(batchErrs) != 0 {
		return models, errors.Join(batchErrs...)
	}
	return models, nil
}

var _ repo.AppUserProfileRepository = &PgAppUserProfileRepo{}
var _ entity.AppUserProfile = entity.AppUserProfile(generated.AppUserProfile{})
