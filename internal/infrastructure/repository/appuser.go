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

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
)

// Bespoke AppUser queries that are not part of the generic CRUD contract.
// The CRUD methods live in zz_generated_appuser.go.

func (r *pgAppUserRepo) GetByUsername(ctx context.Context, username valueobject.Username) (*entity.AppUser, error) {
	m, err := r.q.GetAppUserByUsername(ctx, username)
	if err != nil {
		return nil, MapNotFound(err)
	}
	return ptr(entity.AppUser(m)), nil
}

func (r *pgAppUserRepo) GetByEmail(ctx context.Context, email valueobject.Email) (*entity.AppUser, error) {
	m, err := r.q.GetAppUserByEmail(ctx, &email)
	if err != nil {
		return nil, MapNotFound(err)
	}
	return ptr(entity.AppUser(m)), nil
}

func (r *pgAppUserRepo) ExistsByUsername(ctx context.Context, username valueobject.Username) (bool, error) {
	return r.q.ExistsAppUserByUsername(ctx, username)
}
