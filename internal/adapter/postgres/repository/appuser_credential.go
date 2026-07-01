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

type AppUserCredentialRepo struct {
	b            *generated.DBTX
	querier      *generated.Querier
	guidProvider *ports.GuidProvider
}

// Count implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// Create implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Create(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// CreateBatch implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) CreateBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// Delete implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Delete(ctx context.Context, guid valueobject.GUID) error {
	panic("unimplemented")
}

// DeleteBatch implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) DeleteBatch(ctx context.Context, guids []valueobject.GUID) error {
	panic("unimplemented")
}

// Exists implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Exists(ctx context.Context, guid valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

// ExistsBatch implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) ExistsBatch(ctx context.Context, guids []valueobject.GUID) (bool, error) {
	panic("unimplemented")
}

// Get implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Get(ctx context.Context, guid valueobject.GUID) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// GetAll implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) GetAll(ctx context.Context) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// GetBatch implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) GetBatch(ctx context.Context, guids []valueobject.GUID) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// Save implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Save(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// SaveBatch implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) SaveBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// Update implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) Update(ctx context.Context, obj *entity.AppUserCredential) (*entity.AppUserCredential, error) {
	panic("unimplemented")
}

// UpdateBatch implements repository.AppUserCredentialRepository.
func (a AppUserCredentialRepo) UpdateBatch(ctx context.Context, objs []*entity.AppUserCredential) ([]*entity.AppUserCredential, error) {
	panic("unimplemented")
}

var _ repo.AppUserCredentialRepository = AppUserCredentialRepo{}
var _ entity.AppUserCredential = entity.AppUserCredential(generated.AppUserCredential{})
