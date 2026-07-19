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

package di

import (
	"context"
	"fmt"

	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/cache"
	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/config"
	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/guid"
	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/mail"
	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/repository"
	"github.com/ReallyWeirdCat/brainiac/internal/infrastructure/security"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/app/ports"
	cfg "github.com/ReallyWeirdCat/brainiac/pkg/domain/config"
	"github.com/ReallyWeirdCat/brainiac/pkg/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func provideCache[T any](tag, prefix string) interface{} {
	return fx.Annotate(
		func(redisClient *redis.Client) ports.Cache[T] {
			if redisClient != nil {
				return cache.NewRedisCache[T](redisClient, prefix)
			}
			return cache.NewInMemoryCache[T]()
		},
		fx.ResultTags(fmt.Sprintf(`name:"%s"`, tag)),
	)
}

func newPgxPool(lc fx.Lifecycle, cfg cfg.AppConfigProvider) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			cancel()
			return nil
		},
	})

	pool, err := pgxpool.New(ctx, cfg.Get().Database.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}
	return pool, nil
}

func newRedisClient(cfg cfg.AppConfigProvider) (*redis.Client, error) {
	c := cfg.Get().Cache
	if c.InMemory {
		return nil, nil
	}
	options, err := redis.ParseURL(c.URI)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URL: %w", err)
	}
	return redis.NewClient(options), nil
}

var InfrastructureModule = fx.Module(
	"infrastructure",
	fx.Provide(
		config.NewViperConfig,
		fx.Annotate(
			newPgxPool,
			fx.OnStart(func(ctx context.Context, pool *pgxpool.Pool) error {
				return pool.Ping(ctx)
			}),
			fx.OnStop(func(ctx context.Context, pool *pgxpool.Pool) error {
				pool.Close()
				return nil
			}),
		),
		fx.Annotate(
			newRedisClient,
			fx.OnStop(func(_ context.Context, client *redis.Client) error {
				if client != nil {
					return client.Close()
				}
				return nil
			}),
		),
		guid.NewUuidGuidProvider,
		security.NewBcryptHasher,
		security.NewPwnedPasswordChecker,
		mail.NewSmtpMailer,
		repository.NewUnitOfWorkProvider,
		provideCache[entity.AppUser]("appUsersCache", "usr"),
		provideCache[entity.AppUserCredential]("appUserCredentialsCache", "crd"),
		provideCache[entity.RegistrationCode]("registrationCodesCache", "reg"),
	),
)
