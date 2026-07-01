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
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

type AppUserProfileRepo struct {
	b            *generated.DBTX
	querier      *generated.Querier
	guidProvider *ports.GuidProvider
}

// Count implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// Create implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Create(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// CreateBatch implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// Delete implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	panic("unimplemented")
}

// DeleteBatch implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	panic("unimplemented")
}

// Exists implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

// ExistsBatch implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

// Get implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// GetAll implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) GetAll(ctx context.Context) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// GetBatch implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// Save implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Save(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// SaveBatch implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// Update implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) Update(ctx context.Context, obj *entity.AppUserProfile) (*entity.AppUserProfile, error) {
	panic("unimplemented")
}

// UpdateBatch implements repository.AppUserProfileRepository.
func (a AppUserProfileRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserProfile) ([]*entity.AppUserProfile, error) {
	panic("unimplemented")
}

var _ repo.AppUserProfileRepository = AppUserProfileRepo{}
var _ entity.AppUserCredential = entity.AppUserCredential(generated.AppUserCredential{})
