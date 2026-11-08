# syntax=docker/dockerfile:1

# 1. Build the Svelte frontend (output: /app/web/dist)
FROM node:22-alpine AS frontend-builder
WORKDIR /app/web

# Enable pnpm via corepack, pinned to the version in package.json.
COPY web/package.json web/pnpm-lock.yaml web/pnpm-workspace.yaml ./
RUN corepack enable && corepack prepare pnpm@11.28.5 --activate
RUN pnpm install --frozen-lockfile

# Copy the rest of the frontend source and build.
COPY web/ ./
RUN pnpm build

# 2. Build the Go backend (embeds web/dist via go:embed)
FROM golang:1.27-alpine AS backend-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# The go:embed directive in web/embed.go requires web/dist to exist at build time.
COPY --from=frontend-builder /app/web/dist ./web/dist
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/brainiac ./cmd/brainiac

# 3. Minimal final image
FROM alpine:3.21

# CA certificates for outbound TLS (SMTP, S3, etc.).
RUN apk add --no-cache ca-certificates

COPY --from=backend-builder /app/brainiac /usr/local/bin/brainiac

# Non-root user + writable data dir for the local-disk storage backend.
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup && \
    mkdir -p /data/storage && chown -R appuser:appgroup /data
WORKDIR /data
USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/health || exit 1

ENTRYPOINT ["/usr/local/bin/brainiac"]
