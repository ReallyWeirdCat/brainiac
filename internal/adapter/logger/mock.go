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

package logger

import (
	"context"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
)

type mockLogger struct{}

func (m *mockLogger) Debug(msg string, args ...any) {
}

func (m *mockLogger) Error(msg string, args ...any) {
}

func (m *mockLogger) Info(msg string, args ...any) {
}

func (m *mockLogger) Warn(msg string, args ...any) {
}

func (m *mockLogger) With(args ...any) ports.Logger {
	return m
}

func (m *mockLogger) WithContext(ctx context.Context) ports.Logger {
	return m
}

func NewMockLogger() ports.Logger {
	return &mockLogger{}
}
