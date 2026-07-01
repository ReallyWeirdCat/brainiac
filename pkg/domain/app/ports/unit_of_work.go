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

package ports

import (
	"context"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
)

type UnitOfWorkProvider interface {
	// Returns an instance of UnitOfWork with a transaction
	Begin(ctx context.Context) (UnitOfWork, error)
	// Returns an instance of UnitOfWork with no transaction
	New(ctx context.Context) (UnitOfWork, error)
}

type UnitOfWork interface {
	AppUsers() repository.AppUserRepository
	AppUserCredentials() repository.AppUserCredentialRepository
	AppUserProfiles() repository.AppUserProfileRepository
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
