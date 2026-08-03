# AGENTS.md

Guidance for automated coding agents working in `gmdtable`.

## Package purpose and contract

`gmdtable` is a Go 1.19 module that converts ordered headers and map-backed row data into a Markdown
table. `go.mod` owns the public module path `github.com/tzfqh/gmdtable`. The module-root `gmdtable`
package exposes `Convert(headers []string, data []map[string]interface{}) (string, error)` and uses
`go-runewidth` for Unicode-aware column sizing.

Preserve header-defined column order, default `fmt` value formatting, exact Markdown spacing, the
`<nil>` representation for missing keys, structural cell escaping, and empty-input errors unless a
deliberate public behavior change is approved. Website code must not be imported by the public package,
and website changes must not silently change runtime behavior.

Public API changes require aligned implementation, GoDoc, behavior tests, executable examples,
`examples/basic`, README usage, website examples, and generated API documentation.

## Project structure

- `doc.go`, `gmdtable.go`: public package documentation and implementation.
- `gmdtable_test.go`, `example_test.go`: external-package behavior and executable documentation tests.
- `examples/basic`: runnable consumer-style example.
- `cmd/sitegen`: Go CLI with `build`, `validate`, and `preview` commands.
- `site.config.json`: website identity, metadata, theme palette, PWA identity, and page definitions.
- `internal/site/config.go`: strict configuration loading, derived URLs and keys, and identity checks.
- `internal/site/apidoc.go`: exported module-root API and GoDoc extraction; it excludes `internal` code.
- `internal/site`: SEO/PWA generation, template rendering, artifact validation, preview, and tests.
- `site/templates`: shared layout and navigation plus crawlable home, examples, and API pages.
- `site/assets/css` and `site/assets/js`: small Bootstrap-variable styling and shared browser behavior.
- `site/assets/vendor`: website-only Bootstrap 5.3.8 assets with notices under `guides/`.
- `site/assets/icons`: source favicon, PWA, maskable, Apple-touch, and Open Graph images.
- `dist/pages`: generated Pages artifact; it is ignored and produced in CI.
- `.github/workflows/ci.yml`: package, example, site generation, and artifact quality gates.
- `.github/workflows/pages.yml`: validated `dist/pages` upload and GitHub Pages deployment.

## Stable Pages contract

Keep real static pages at `index.html`, `examples/index.html`, and `api/index.html`, deployed below
`/gmdtable/`. All internal links, assets, manifest entries, service-worker paths, and preview routes
must remain repository-subpath-aware.

The examples page reads checked-in Go source. The API page parses actual exported root-package
declarations and GoDoc; do not manually maintain API declarations or document `internal` packages as
consumer API. SEO metadata, JSON-LD, `robots.txt`, and `sitemap.xml` derive from `site.config.json`.

Each new public page requires crawlable initial HTML, semantic landmarks, exactly one descriptive
`h1`, keyboard-accessible controls, and accurate unique canonical and Open Graph metadata. Preserve
visible focus indicators and accessible mobile navigation.

## Theme, SEO, and PWA behavior

The shared theme control supports exactly `system`, `light`, and `dark` through Bootstrap color
modes. Keep theme and cache keys target-specific. Persist explicit theme preferences, respond to
system-theme changes, synchronize `theme-color`, and retain safe storage failure handling.

The generated service worker is optional website behavior. It handles only same-origin GET requests
within `/gmdtable/`, uses network-first HTML and cache-first local assets, bounds runtime entries, and
cleans only `gmdtable` caches. Ordinary non-production file viewing must not register it.

## Commands and generated files

Run from the repository root:

```bash
gofmt -w .
go mod tidy
go test ./...
go test -race ./...
go vet ./...
go build ./...
go list ./...
go list ./... | xargs -n 1 go doc
go run ./examples/basic
go run ./cmd/sitegen build
go run ./cmd/sitegen validate
go run ./cmd/sitegen preview
```

`preview` rebuilds and validates before serving `http://127.0.0.1:4173/gmdtable/`. Stop the server
after verification.

Generated Pages files must not be edited manually. Change package GoDoc, configuration, templates,
assets, or generator code and rebuild. `dist/pages` should be generated in CI and not committed.

## Change discipline

Keep `site.config.json` synchronized with `go.mod` where values intentionally mirror each other. Add
deterministic tests for new configuration, route, metadata, theme, cache, scope, extraction, or
validation behavior. Keep all checks independent of external network services.

Do not add Node.js for site generation. Do not import Bootstrap or site tooling into the public Go
package. Preserve third-party license notices. Do not publish releases or tags unless explicitly
requested.
