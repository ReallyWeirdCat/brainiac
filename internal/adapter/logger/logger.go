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
	"fmt"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/rs/zerolog"
)

type ZerologLogger struct {
	log zerolog.Logger
}

func NewZerologLogger(log zerolog.Logger) ports.Logger {
	return &ZerologLogger{log: log}
}

// argsToFields converts a list of alternating key-value pairs into a map suitable for zerolog.
// Non-string keys are converted with fmt.Sprint. An odd trailing key is given value '?'.
func argsToFields(args ...any) map[string]any {
	if len(args)%2 != 0 {
		args = append(args, "?")
	}
	if len(args) == 0 {
		return nil
	}
	fields := make(map[string]any, len(args)/2)
	for i := 0; i < len(args)-1; i += 2 {
		key := fmt.Sprint(args[i])
		fields[key] = args[i+1]
	}
	return fields
}

// Debug logs a message at debug level with the given key-value pairs.
func (l *ZerologLogger) Debug(msg string, args ...any) {
	l.log.Debug().Fields(argsToFields(args...)).Msg(msg)
}

// Info logs a message at info level with the given key-value pairs.
func (l *ZerologLogger) Info(msg string, args ...any) {
	l.log.Info().Fields(argsToFields(args...)).Msg(msg)
}

// Warn logs a message at warn level with the given key-value pairs.
func (l *ZerologLogger) Warn(msg string, args ...any) {
	l.log.Warn().Fields(argsToFields(args...)).Msg(msg)
}

// Error logs a message at error level with the given key-value pairs.
func (l *ZerologLogger) Error(msg string, args ...any) {
	l.log.Error().Fields(argsToFields(args...)).Msg(msg)
}

// With returns a new Logger with the given key-value pairs added to the logging context.
func (l *ZerologLogger) With(args ...any) ports.Logger {
	return &ZerologLogger{
		log: l.log.With().Fields(argsToFields(args...)).Logger(),
	}
}

// WithContext returns a logger that picks up contextual fields from the context.
// If a zerolog logger is present in the context (set via (*Logger).WithContext(ctx)),
// that logger is used directly, preserving all its fields. Otherwise the current logger is returned.
func (l *ZerologLogger) WithContext(ctx context.Context) ports.Logger {
	if ctxLogger := zerolog.Ctx(ctx); ctxLogger.GetLevel() != zerolog.Disabled {
		return &ZerologLogger{log: *ctxLogger}
	}
	return l
}
