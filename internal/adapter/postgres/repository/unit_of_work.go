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
	"fmt"
	"sync"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWorkProvider struct {
	pool *pgxpool.Pool
}

func NewUnitOfWorkProvider(pool *pgxpool.Pool) (ports.UnitOfWorkProvider, error) {
	return &UnitOfWorkProvider{pool: pool}, nil
}

func (u *UnitOfWorkProvider) Begin(ctx context.Context) (ports.UnitOfWork, error) {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	queries := generated.New(tx)
	return NewUnitOfWork(queries, tx), nil
}

func (u *UnitOfWorkProvider) New(ctx context.Context) (ports.UnitOfWork, error) {
	queries := generated.New(u.pool)
	return NewUnitOfWork(queries, nil), nil
}

var _ ports.UnitOfWorkProvider = &UnitOfWorkProvider{}

type UnitOfWork struct {
	mu                 sync.Mutex
	queries            *generated.Queries
	tx                 pgx.Tx
	committed          bool
	appUsers           repository.AppUserRepository
	appUserCredentials repository.AppUserCredentialRepository
	appUserProfiles    repository.AppUserProfileRepository
	appUserSessions    repository.AppUserSessionRepository
}

func NewUnitOfWork(queries *generated.Queries, tx pgx.Tx) ports.UnitOfWork {
	uow := &UnitOfWork{
		queries: queries,
		tx:      tx,
	}
	uow.appUsers = NewPgAppUserRepo(queries, &uow.mu)
	uow.appUserCredentials = NewPgAppUserCredentialRepo(queries, &uow.mu)
	uow.appUserProfiles = NewPgAppUserProfileRepo(queries, &uow.mu)
	uow.appUserSessions = NewPgAppUserSessionRepo(queries, &uow.mu)
	return uow
}

func (u *UnitOfWork) AppUserCredentials() repository.AppUserCredentialRepository {
	return u.appUserCredentials
}

func (u *UnitOfWork) AppUserProfiles() repository.AppUserProfileRepository {
	return u.appUserProfiles
}

func (u *UnitOfWork) AppUsers() repository.AppUserRepository {
	return u.appUsers
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.tx == nil {
		return nil
	}
	if u.committed {
		return fmt.Errorf("transaction already committed")
	}
	err := u.tx.Commit(ctx)
	if err == nil {
		u.committed = true
	}
	return err
}

func (u *UnitOfWork) Rollback(ctx context.Context) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.tx == nil {
		return nil
	}
	if u.committed {
		return nil
	}
	return u.tx.Rollback(ctx)
}

var _ ports.UnitOfWork = &UnitOfWork{}
