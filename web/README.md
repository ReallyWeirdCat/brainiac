# Brainiac frontend

Svelte 5 + Vite single-page application. The production build (`web/dist`) is
embedded into the Go binary via `web/embed.go` and served by the backend — no
separate frontend container.

## Setup

```sh
pnpm install
```

## Develop (with Vite dev server)

```sh
pnpm dev
```

## Build (outputs to `web/dist`, consumed by the Go embed)

```sh
pnpm build
```
