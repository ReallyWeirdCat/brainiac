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

package auth

import (
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
)

type LoginUsecase struct {
	userRepo       repository.AppUserRepository
	credentialRepo repository.AppUserCredentialRepository
	profileRepo    repository.AppUserProfileRepository
	uow            ports.UnitOfWork
	hasher         ports.PasswordHasher
	tokenGenerator ports.TokenGenerator
}

func NewLoginUsecase(
	userRepo repository.AppUserRepository,
	credentialRepo repository.AppUserCredentialRepository,
	profileRepo repository.AppUserProfileRepository,
	uow ports.UnitOfWork,
	hasher ports.PasswordHasher,
	tokenGenerator ports.TokenGenerator,
) *LoginUsecase {
	return &LoginUsecase{
		userRepo:       userRepo,
		credentialRepo: credentialRepo,
		profileRepo:    profileRepo,
		uow:            uow,
		hasher:         hasher,
		tokenGenerator: tokenGenerator,
	}
}

func (u *LoginUsecase) Execute(req LoginRequest) (LoginResponse, error) {
	// TODO:
	panic("not implemented")
}
