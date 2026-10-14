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
	"errors"

	repo "github.com/ReallyWeirdCat/brainiac/pkg/domain/repository"
	"github.com/jackc/pgx/v5"
)

// MapNotFound translates a pgx "no rows" error into the domain not-found error.
// Every other error passes through unchanged.
func MapNotFound(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return repo.ErrEntityNotFound.FromError(err)
	}
	return err
}

// ptr returns a pointer to v. Kept as a tiny helper so generated adapters read
// as "return ptr(entity.X(m)), nil" instead of a local variable + &x.
func ptr[T any](v T) *T { return &v }
