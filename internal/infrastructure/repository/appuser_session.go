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
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/valueobject"
	"github.com/jackc/pgx/v5"
)

// Bespoke AppUserSession queries. CRUD lives in zz_generated_appuser_session.go.

func (r *pgAppUserSessionRepo) GetAllActiveByUsername(ctx context.Context, username valueobject.Username) ([]*entity.AppUserSession, error) {
	objs, err := r.q.GetAllActiveSessionsByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entity.AppUserSession{}, nil
		}
		return nil, err
	}
	return mapSessionRows(objs), nil
}

func (r *pgAppUserSessionRepo) GetAllInactiveByUsername(ctx context.Context, username valueobject.Username) ([]*entity.AppUserSession, error) {
	objs, err := r.q.GetAllInactiveSessionsByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entity.AppUserSession{}, nil
		}
		return nil, err
	}
	return mapSessionRows(objs), nil
}

func (r *pgAppUserSessionRepo) GetAllByLastIP(ctx context.Context, ip string) ([]*entity.AppUserSession, error) {
	objs, err := r.q.GetAllSessionsByLastIP(ctx, &ip)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*entity.AppUserSession{}, nil
		}
		return nil, err
	}
	return mapSessionRows(objs), nil
}

func (r *pgAppUserSessionRepo) GetAllAppUserGUIDsByIP(ctx context.Context, ip string) ([]valueobject.GUID, error) {
	guids, err := r.q.GetAllAppUserGUIDsByIP(ctx, &ip)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []valueobject.GUID{}, nil
		}
		return nil, err
	}
	return guids, nil
}

func mapSessionRows(objs []generated.AppUserSession) []*entity.AppUserSession {
	out := make([]*entity.AppUserSession, 0, len(objs))
	for i := range objs {
		e := entity.AppUserSession(objs[i])
		out = append(out, &e)
	}
	return out
}
