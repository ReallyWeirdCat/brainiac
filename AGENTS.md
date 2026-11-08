# Brainiac

Brainiac is an education platform: courses, student groups and teachers, an internal messenger, course subjects (Markdown documents, assessments, tasks with submissions), gamification (experience, levels, virtual items, achievements) and personal user stats.

Business rules live in [DOMAIN.md](./DOMAIN.md). Read it before changing any domain logic (attendance, assessments, permissions, inventory, achievements).

## Stack

- Architecture: monolith, Clean Architecture, partially DDD (anemic models, no aggregates)
- Backend: Go 1.26.3, built-in Go HTTP (no web framework), SQLC, zerolog
- Storage: PostgreSQL, Redis, RustFS (S3-compatible object storage)
- Edge: Nginx
- Frontend: SvelteKit in `web/`, built to static files and served by a Go endpoint

## Repository layout

```
.
├── cmd/brainiac                 # entry point
├── internal                     # infrastructure and adapters layer
│   ├── adapter
│   ├── controller               # grpc, http, scheduler
│   ├── di                       # dependency wiring
│   └── infrastructure
│       ├── cache, config, guid, logger, mail, security
│       ├── database/postgres
│       │   ├── generated        # SQLC output
│       │   ├── migrations
│       │   └── queries          # SQLC input (.sql)
│       └── repository           # repository implementations
├── pkg/domain                   # business layer
│   ├── app
│   │   ├── ports
│   │   └── usecase              # grouped by area (auth, mail, ...)
│   ├── config
│   ├── entity
│   ├── enum
│   ├── errors
│   ├── repository               # repository interfaces
│   └── valueobject
└── web                          # SvelteKit frontend
```

### Dependency rule

- `pkg/domain` must never import from `internal`. Dependencies point inward only.
- The domain defines interfaces (`pkg/domain/repository`, `pkg/domain/app/ports`). `internal` implements them.
- Controllers call use cases. They contain no business logic.
- Wiring happens in `internal/di`.

### Where new code goes

| Adding | Location |
|---|---|
| Entity, value object, enum | `pkg/domain/entity`, `valueobject`, `enum` |
| Domain error | `pkg/domain/errors` |
| Use case | `pkg/domain/app/usecase/<area>` |
| Repository interface | `pkg/domain/repository` |
| Repository implementation | `internal/infrastructure/repository` |
| SQL query | `internal/infrastructure/database/postgres/queries`, then regenerate |
| Migration | `internal/infrastructure/database/postgres/migrations` |
| HTTP / gRPC handler, scheduled job | `internal/controller/{http,grpc,scheduler}` |

## Commands

<!-- TODO: fill in the real commands. An agent cannot verify its work without them. -->

| Task | Command |
|---|---|
| Build | `go build ./...` |
| Test | `go test ./...` |
| Lint | `go fmt ./...` |
| Regenerate SQLC | `sqlc generate` |
| Run migrations | TODO |
| Run locally (DB, Redis, RustFS) | TODO |
| Build frontend | TODO |

## Conventions

- Avoid reflection-based code. Prefer SQLC and code generation for compile-time safety and performance.
- Document through readable code and good names. Add comments only where the *why* is not obvious.
- Every non-generated code file starts with the GPL license header.
- Wrap and propagate errors using the types in `pkg/domain/errors`. TODO: document the wrapping pattern and how errors map to HTTP status codes.

Commits are done strictly in English language in the following format:
```
feat: implement feature;
refactor: restructure object;
fix: page does not open correctly;
docs: update readme;
experiment: try a new approach.
```
One change per line with a marker feat/refactor/fix/docs/experiment, devided by semicolon, period at the last line.

## Testing

- Unit tests with table-driven cases and mocks.
- No integration tests. Do not add tests that need a real database, Redis or network.
- New use cases and repository-facing logic come with tests.

## Logging and privacy

- Log operations verbosely using existing logger with zerolog implementation.
- Never log plain personal data or secrets: passwords, tokens, emails, names. Obfuscate personal data before logging.

## Security

- Passwords are hashed with bcrypt. Never store or log them in plain text.
- Passwords are checked against a bad-password list (rockyou.txt) on the relevant flows.
- On login a session is created. Access and refresh tokens carry the session id. An invalid session invalidates all of its tokens, so always validate the session, not just the token signature.
- TODO: document token lifetimes, where validation happens, and the session invalidation check.

## Boundaries

**Never**
- Create a commit without the user's review.
- Hand-edit generated code in `internal/infrastructure/database/postgres/generated`. Change the query and regenerate.
- Edit a migration that has already been applied. Add a new one.

**Workflow**
- Split work into small, logical commits. Each one is reviewed by the user before the next begins.
- A change is done when it builds, tests pass, generated code is up to date, and new files carry the license header.

**Ask first**
- Schema changes, new dependencies, changes to auth or session handling.
