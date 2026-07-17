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
	"testing"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/database/postgres/generated"
	"github.com/jackc/pgx/v5"
)

type mockTx struct {
	pgx.Tx
	committed  bool
	rolledBack bool
}

func (m *mockTx) Commit(ctx context.Context) error {
	if m.committed {
		return errors.New("already committed")
	}
	m.committed = true
	return nil
}

func (m *mockTx) Rollback(ctx context.Context) error {
	if m.committed {
		return nil
	}
	m.rolledBack = true
	return nil
}

func newTestUnitOfWork(tx pgx.Tx) *UnitOfWork {
	return NewUnitOfWork(&generated.Queries{}, tx).(*UnitOfWork)
}

func TestUnitOfWork_Commit(t *testing.T) {
	t.Run("transactional UoW commits successfully", func(t *testing.T) {
		tx := &mockTx{}
		uow := newTestUnitOfWork(tx)

		if err := uow.Commit(context.Background()); err != nil {
			t.Fatalf("Commit() unexpected error: %v", err)
		}
		if !tx.committed {
			t.Error("Commit() should have called tx.Commit()")
		}
	})

	t.Run("non-transactional UoW returns nil", func(t *testing.T) {
		uow := newTestUnitOfWork(nil)

		if err := uow.Commit(context.Background()); err != nil {
			t.Fatalf("Commit() on non‑transactional UoW should return nil, got: %v", err)
		}
	})

	t.Run("double commit returns error", func(t *testing.T) {
		tx := &mockTx{}
		uow := newTestUnitOfWork(tx)

		if err := uow.Commit(context.Background()); err != nil {
			t.Fatalf("first Commit() unexpected error: %v", err)
		}
		err := uow.Commit(context.Background())
		if err == nil {
			t.Fatal("second Commit() should return an error")
		}
		if err.Error() != "transaction already committed" {
			t.Errorf("unexpected error message: %v", err)
		}
	})
}

func TestUnitOfWork_Rollback(t *testing.T) {
	t.Run("rollback without commit calls tx.Rollback", func(t *testing.T) {
		tx := &mockTx{}
		uow := newTestUnitOfWork(tx)

		if err := uow.Rollback(context.Background()); err != nil {
			t.Fatalf("Rollback() unexpected error: %v", err)
		}
		if !tx.rolledBack {
			t.Error("Rollback() should have called tx.Rollback()")
		}
	})

	t.Run("rollback after commit is no‑op", func(t *testing.T) {
		tx := &mockTx{}
		uow := newTestUnitOfWork(tx)

		if err := uow.Commit(context.Background()); err != nil {
			t.Fatalf("Commit() failed: %v", err)
		}
		if err := uow.Rollback(context.Background()); err != nil {
			t.Fatalf("Rollback() after commit should return nil, got: %v", err)
		}
		if tx.rolledBack {
			t.Error("Rollback() after commit should not call tx.Rollback()")
		}
	})

	t.Run("non-transactional UoW rollback returns nil", func(t *testing.T) {
		uow := newTestUnitOfWork(nil)

		if err := uow.Rollback(context.Background()); err != nil {
			t.Fatalf("Rollback() on non‑transactional UoW should return nil, got: %v", err)
		}
	})
}

func TestUnitOfWork_RepositoryGetters(t *testing.T) {
	uow := NewUnitOfWork(&generated.Queries{}, nil).(*UnitOfWork)

	if uow.AppUsers() == nil {
		t.Error("AppUsers() should not return nil")
	}
	if uow.AppUserCredentials() == nil {
		t.Error("AppUserCredentials() should not return nil")
	}
	if uow.AppUserProfiles() == nil {
		t.Error("AppUserProfiles() should not return nil")
	}
	// TODO: check other repos
}

func TestNewUnitOfWork(t *testing.T) {
	t.Run("returns non‑nil and initialises repositories", func(t *testing.T) {
		uow := NewUnitOfWork(&generated.Queries{}, nil)
		if uow == nil {
			t.Fatal("NewUnitOfWork must not return nil")
		}
	})

	t.Run("accepts nil tx for non‑transactional mode", func(t *testing.T) {
		uow := NewUnitOfWork(&generated.Queries{}, nil).(*UnitOfWork)
		if uow.tx != nil {
			t.Error("non‑transactional UoW should have tx == nil")
		}
	})

	t.Run("accepts a real pgx.Tx", func(t *testing.T) {
		tx := &mockTx{}
		uow := NewUnitOfWork(&generated.Queries{}, tx).(*UnitOfWork)
		if uow.tx != tx {
			t.Error("transactional UoW should hold the provided tx")
		}
	})
}
