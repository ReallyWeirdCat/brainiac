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
	"context"
	"errors"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
)

// ErrNotImplemented is returned until the login flow is implemented.
var ErrNotImplemented = errors.New("login usecase is not implemented yet")

// LoginUsecase depends on the UnitOfWorkProvider (not on a concrete
// UnitOfWork value or standalone repositories) so it gets a fresh transaction
// per request, matching the RegistrationBeginUsecase pattern.
type LoginUsecase struct {
	uow            ports.UnitOfWorkProvider
	hasher         ports.PasswordHasher
	tokenGenerator ports.TokenGenerator
}

func NewLoginUsecase(
	uow ports.UnitOfWorkProvider,
	hasher ports.PasswordHasher,
	tokenGenerator ports.TokenGenerator,
) *LoginUsecase {
	return &LoginUsecase{
		uow:            uow,
		hasher:         hasher,
		tokenGenerator: tokenGenerator,
	}
}

func (u *LoginUsecase) Execute(ctx context.Context, _ LoginRequest) (LoginResponse, error) {
	// TODO: verify credentials, issue access/refresh tokens, persist a session.
	return LoginResponse{}, ErrNotImplemented
}
