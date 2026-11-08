# Brainiac

Brainiac is an education platform: courses, student groups and teachers, an internal messenger, course subjects (Markdown documents, assessments, tasks with submissions), gamification (experience, levels, virtual items, achievements) and personal user stats.

Business rules live in [DOMAIN.md](./DOMAIN.md). Read it before changing any domain logic (attendance, assessments, permissions, inventory, achievements).

## Stack

- Architecture: monolith, Clean Architecture, partially DDD (anemic models, no aggregates)
- Backend: Go 1.26.3, standard-library `net/http` (no web framework), Uber `fx` for wiring, `pgx/v5`, SQLC, zerolog
- Storage: PostgreSQL; Redis is optional (in-memory cache is the default); local disk is the only object-store implementation (no S3 backend yet)
- Edge: Nginx (deployment concern; there is no Nginx configuration in this repository)
- Frontend: Svelte 5 + Vite SPA in `web/`, built to static files and embedded into the Go binary
- Tooling: SQLC v1.31.1, Node.js 22 in the Docker image, pnpm 11.28.5 (`web/package.json`), Docker/Compose optional

## Repository layout

```
.
├── cmd/
│   ├── brainiac/                 # application entry point (fx)
│   └── genrepo/                  # generator for zz_generated_* repository adapters
├── internal/
│   ├── di/                       # fx modules: app, infrastructure, adapters, use cases
│   ├── transport/                # net/http server, /health, embedded SPA handler
│   └── infrastructure/
│       ├── cache/                # in-memory and Redis cache implementations
│       ├── config/               # Viper + envconfig configuration loader
│       ├── database/postgres/
│       │   ├── generated/        # SQLC output (do not edit)
│       │   ├── migrations/       # SQL migrations
│       │   └── queries/          # SQLC input (.sql)
│       ├── guid/                 # UUID provider
│       ├── logger/               # zerolog adapter
│       ├── mail/                 # SMTP mailer
│       ├── repository/           # pgx repository adapters and UnitOfWork
│       ├── security/             # bcrypt, JWT, compromised-password checker
│       └── storage/              # local-disk object store
├── pkg/
│   └── domain/
│       ├── app/
│       │   ├── ports/            # infrastructure interfaces (cache, mail, storage, tokens, UnitOfWork)
│       │   └── usecase/          # application services grouped by area (auth, mail, ...)
│       ├── config/               # AppConfig and validation
│       ├── entity/               # domain entities (anemic)
│       ├── enum/                 # domain enums
│       ├── errors/               # DomainError and error types
│       ├── repository/           # persistence interfaces
│       └── valueobject/          # GUID, Email, Bio, I18nText, Metadata, ...
├── web/                          # Svelte 5 + Vite SPA; web/dist (gitignored) is embedded
├── sqlc.yaml                     # SQLC config and Go type overrides
├── docker-compose.yml            # app + PostgreSQL + Redis
├── Dockerfile                    # frontend build, Go build, runtime image
├── example.env                   # documented environment variables
└── DOMAIN.md                     # business rules
```

There is no `internal/controller` or `internal/adapter` package yet. HTTP handlers live in `internal/transport`; gRPC and scheduler adapters have not been introduced.

### Layer flow

`cmd/brainiac` → `internal/di` → `internal/transport` / use cases → domain interfaces → `internal/infrastructure/...`.

### Dependency rule

- `pkg/domain` must never import from `internal`. Dependencies point inward only.
- The domain defines interfaces (`pkg/domain/repository`, `pkg/domain/app/ports`). `internal` implements them.
- Controllers/transport call use cases. They contain no business logic.
- Wiring happens in `internal/di`.

### Where new code goes

| Adding | Location |
|---|---|
| Entity, value object, enum | `pkg/domain/entity`, `pkg/domain/valueobject`, `pkg/domain/enum` |
| Domain error | `pkg/domain/errors` |
| Use case | `pkg/domain/app/usecase/<area>` |
| Repository interface | `pkg/domain/repository` |
| Repository implementation | `internal/infrastructure/repository` |
| SQL query | `internal/infrastructure/database/postgres/queries`, then regenerate |
| Migration | `internal/infrastructure/database/postgres/migrations` |
| HTTP handler or route | `internal/transport` |
| gRPC handler or scheduled job | not implemented yet; introduce `internal/controller/...` when needed |
| Configuration field | `pkg/domain/config/app_config.go` plus an `envconfig` tag and `example.env` |
| Cache, mail, storage or security adapter | `internal/infrastructure/<area>` plus a port in `pkg/domain/app/ports` |
| DI registration | `internal/di` |

## Commands

| Task | Command |
|---|---|
| Build frontend | `cd web && pnpm install && pnpm build` |
| Build backend (requires `web/dist`) | `go build ./...` |
| Build binary | `go build -o brainiac ./cmd/brainiac` |
| Test | `go test ./...` |
| Vet | `go vet ./...` |
| Format | `go fmt ./...` |
| Regenerate SQLC | `sqlc generate` |
| Regenerate repositories | `go generate ./internal/infrastructure/repository` |
| Run full stack | `docker compose up --build` |
| Run app natively | `go run ./cmd/brainiac` |

`web/embed.go` uses `//go:embed all:dist`, so `web/dist` must exist before any Go build or `go run`. A missing build fails with `pattern all:dist: no matching files found`. The Dockerfile builds the frontend before the Go binary; for local work run `pnpm build` in `web/` first.

### Local run

Full stack in Docker (PostgreSQL and Redis run inside the Compose network):

```sh
docker compose up --build -d
docker compose exec -T postgres psql -U brainiac -d brainiac \
  -v ON_ERROR_STOP=1 < internal/infrastructure/database/postgres/migrations/0001_initial.sql
```

Native process (PostgreSQL and, if `CACHE_IN_MEMORY=false`, Redis must already be reachable):

```sh
export DATABASE_URI="postgres://brainiac:brainiac@localhost:5432/brainiac?sslmode=disable"
export SECURITY_JWT_SECRET="$(openssl rand -base64 48)"
psql "$DATABASE_URI" -v ON_ERROR_STOP=1 \
  -f internal/infrastructure/database/postgres/migrations/0001_initial.sql
cd web && pnpm install && pnpm build && cd ..
go run ./cmd/brainiac
```

## Configuration

The application reads `./config.yaml` and overlays environment variables using the `envconfig` tags in `pkg/domain/config/app_config.go`. `config.yaml` is gitignored and is created with defaults on first start if missing. `example.env` documents the variables; it is a template, not loaded automatically. Export the variables or edit `config.yaml`.

Required:

- `DATABASE_URI` — PostgreSQL connection string.
- `SECURITY_JWT_SECRET` — at least 32 characters. The generated default is empty, so set it via environment or edit `config.yaml` before the first start.

Common optional variables:

- `CACHE_IN_MEMORY` (default `true`); when `false`, `CACHE_URI` must point to Redis.
- `SERVER_ADDR` (default `:8080`).
- `STORAGE_ROOT` (default `./storage`) and `STORAGE_BASE_URL`.
- `SMTP_*` — required only when `SMTP_ENABLE=true`.
- `REGISTRATION_*` and `LOGIN_*`.
- `SECURITY_PASSWORDS_COMPROMISED_*` — compromised-password checking; a file path is required when enabled.

## Database and migrations

Migrations are plain SQL files in `internal/infrastructure/database/postgres/migrations` and are not applied automatically by the app or `docker-compose.yml`. Apply each new file once, in filename order:

```sh
psql "$DATABASE_URI" -v ON_ERROR_STOP=1 \
  -f internal/infrastructure/database/postgres/migrations/0001_initial.sql
```

Never edit an applied migration; add a new numbered file. After changing the schema, run `sqlc generate` so the generated models and query parameters match the migrations.

## Persistence and code generation

- `internal/infrastructure/database/postgres/queries/*.sql` is SQLC input.
- `internal/infrastructure/database/postgres/generated/` is SQLC output. Do not edit it by hand; change the query or `sqlc.yaml` and run `sqlc generate`. SQLC reads the migrations and query files directly and does not need a running database.
- `internal/infrastructure/repository/zz_generated_*.go` is produced by `cmd/genrepo`. Do not edit it by hand; change `cmd/genrepo/main.go` and run `go generate ./internal/infrastructure/repository`.
- `sqlc.yaml` maps PostgreSQL columns and enums to domain types (`GUID`, `Bio`, `I18nText`, `Metadata`, enums, ...). Add an override there when a column should use a domain value object.
- Generated repository adapters convert `generated.X` to `entity.X` with a direct struct conversion. Field names, types and order must match exactly. When adding a database column, update the query, the SQLC override if needed, the entity, and the entity's `Writable` list in `cmd/genrepo/main.go`, then regenerate both layers.
- `internal/infrastructure/repository/unit_of_work.go` is hand-written. Add an accessor there when adding a new entity repository.

## Current HTTP surface

Only `GET /health` and the embedded SPA are routed today. The JSON API, gRPC server and scheduler are not implemented yet. Use cases live in `pkg/domain/app/usecase`, but not every use case is wired in `internal/di/usecase.go`.

## Conventions

- Avoid reflection-based code. Prefer SQLC and code generation for compile-time safety and performance.
- Document through readable code and good names. Add comments only where the *why* is not obvious.
- Every non-generated code file starts with the GPL license header.
- Return typed domain errors from `pkg/domain/errors` (`Validation`, `NotFound`, `Conflict`, `Forbidden`, `Unauthorized`, ...). Repository adapters translate `pgx.ErrNoRows` with `repository.MapNotFound`; HTTP error mapping is not implemented yet.
- Keep `pkg/domain` free of `internal` imports.

Commits are done strictly in English language in the following format:
```
feat: implement feature;
refactor: restructure object;
fix: page does not open correctly;
docs: update readme;
experiment: try a new approach.
```
One change per line with a marker feat/refactor/fix/docs/experiment, divided by semicolon, period at the last line.

## Testing

- Unit tests with table-driven cases and mocks.
- No integration tests. Do not add tests that need a real database, Redis or network.
- Use in-memory fakes for dependencies where practical (for example `miniredis` and `go-smtp-mock`, both already in `go.mod`).
- New use cases and repository-facing logic come with tests.
- Run the whole suite with `go test ./...`.

## Logging and privacy

- Log operations verbosely using existing logger with zerolog implementation.
- Never log plain personal data or secrets: passwords, tokens, emails, names. Obfuscate personal data before logging.

## Security

- Passwords are hashed with bcrypt. Never store or log them in plain text.
- Passwords are checked against a bad-password list (rockyou.txt) on the relevant flows.
- On login a session is created. Access and refresh tokens carry the session id. An invalid session invalidates all of its tokens, so always validate the session, not just the token signature.
- JWT token lifetimes default to access `3m`, refresh `168h` and TOTP `10m`; override them with `SECURITY_JWT_ACCESS_TOKEN_TTL`, `SECURITY_JWT_REFRESH_TOKEN_TTL` and `SECURITY_JWT_TOTP_TOKEN_TTL`. The refresh token lifetime must be longer than the access token lifetime.
- `SECURITY_JWT_SECRET` must be at least 32 characters and should be supplied by the deployment environment, not committed.

## Boundaries

**Never**
- Create a commit without the user's review.
- Hand-edit generated code in `internal/infrastructure/database/postgres/generated`. Change the query or `sqlc.yaml` and regenerate.
- Hand-edit `internal/infrastructure/repository/zz_generated_*.go`. Change `cmd/genrepo` and run `go generate`.
- Edit a migration that has already been applied. Add a new one.

**Workflow**
- Split work into small, logical commits. Each one is reviewed by the user before the next begins.
- A change is done when it builds, tests pass, generated code is up to date, and new files carry the license header.

**Ask first**
- Schema changes, new dependencies, changes to auth or session handling.
