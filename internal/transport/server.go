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

// Package transport wires the standard-library HTTP server: health, the SPA,
// and (in future) the JSON API. No third-party router or web framework is used.
package transport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
	"go.uber.org/fx"
)

// NewRouter assembles all HTTP routes.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /health", healthHandler())
	mux.Handle("/", spaHandler())
	return mux
}

// NewHTTPServer builds the *http.Server and registers start/stop lifecycle
// hooks so fx owns the listener and shuts the server down gracefully.
func NewHTTPServer(lc fx.Lifecycle, cfg config.AppConfigProvider, logger ports.Logger) *http.Server {
	srv := &http.Server{
		Addr:              cfg.Get().Server.Addr,
		Handler:           NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return fmt.Errorf("listen %s: %w", srv.Addr, err)
			}
			go func() {
				if err := srv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("http server stopped unexpectedly", "error", err)
				}
			}()
			logger.Info("http server listening", "addr", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return srv.Shutdown(shutdownCtx)
		},
	})

	return srv
}

// Module is the fx option that registers the HTTP transport.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewRouter, NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	)
}
