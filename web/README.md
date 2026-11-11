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

## GitHub Pages

The landing page is deployed by `.github/workflows/pages.yml` to
`https://reallyweirdcat.github.io/brainiac/`. GitHub Pages serves project sites
from a subpath, so that build uses `pnpm build:pages`, which sets the Vite base
to `/brainiac/`.

Keep using `pnpm build` for the Go embed: the backend serves the SPA from `/`,
so the Pages base path must not be baked into the embedded build.
