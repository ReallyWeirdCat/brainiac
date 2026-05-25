# syntax=docker/dockerfile:1

# 1. Extract curl (healthchecks)
FROM alpine/curl AS curl

# 2. Build the Vue frontend
FROM node:22-alpine AS frontend-builder
WORKDIR /app/web

# Enable pnpm via corepack (bundled with Node)
COPY web/package.json web/pnpm-lock.yaml ./
RUN corepack enable && corepack prepare pnpm@latest --activate

# Disable the "minimum release age" policy
RUN pnpm config set minimum-release-age 0

# Install dependencies (cached unless package files change)
RUN pnpm install --frozen-lockfile

# Remove pnpm config that might affect reproducibility
RUN pnpm config delete minimum-release-age

# Copy the rest of the frontend source and build
COPY web/ ./
RUN pnpm format
RUN pnpm build
# Output is in /app/web/dist (default Vite output)

# 3. Build the Go backend
FROM golang:1.26-alpine AS backend-builder
WORKDIR /app

# Download module dependencies (cache this layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code and build a static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/brainiac ./cmd/brainiac

# 4. Create the final minimal image
FROM alpine:latest

# Copy curl and all required shared libraries + CA certificates
COPY --from=curl /usr/bin/curl /usr/bin/curl
COPY --from=curl /lib/ld-musl-x86_64.so.1 /lib/
COPY --from=curl /usr/lib/libcurl.so.* /usr/lib/
COPY --from=curl /usr/lib/libssl.so.* /usr/lib/
COPY --from=curl /usr/lib/libcrypto.so.* /usr/lib/
COPY --from=curl /usr/lib/libz.so.* /usr/lib/
COPY --from=curl /usr/lib/libnghttp2.so.* /usr/lib/
COPY --from=curl /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the backend binary
COPY --from=backend-builder /app/brainiac /usr/local/bin/brainiac

# Copy the built frontend (serve it from wherever your Go app expects)
COPY --from=frontend-builder /app/web/dist /var/www/html

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup
USER appuser

ENTRYPOINT ["/usr/local/bin/brainiac"]
