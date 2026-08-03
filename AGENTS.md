# Contributor guide

## Repository scope

`gmdtable` is a Go 1.19 library that converts ordered headers and map-backed row data into a Markdown table. It has no command-line executable, server, database, runtime configuration, or generated source. Keep changes focused on the library's public behavior and its tests.

## Repository map

```text
.
├── gmdtable.go             # Package implementation and public Convert API
├── gmdtable_test.go        # Unit tests for package behavior
├── go.mod                  # Module path and minimum Go language version
├── README.md               # Installation, usage, and public documentation
├── .github/workflows/
│   └── test.yml            # Push CI: dependency download, build, and tests
├── .vscode/settings.json   # Optional editor window-title setting
├── .gitignore              # Repository ignore rules; currently empty
└── LICENSE                 # MIT license
```

The repository root is the package root. There are no internal packages, vendored dependencies, generated directories, or committed build artifacts.

## Entry point and startup flow

The consumer entry point is:

```go
func Convert(headers []string, data []map[string]interface{}) (string, error)
```

Consumers import `github.com/tzfqh/gmdtable` and call `gmdtable.Convert`. There is no application startup sequence. A call follows this path:

1. `Convert` rejects an empty `data` slice with `the data slice is empty`.
2. A `strings.Builder` writes the header row in the order supplied by `headers`.
3. The function writes one separator cell for each header.
4. For each map in `data`, it looks up values by iterating over `headers`; map iteration order never controls column order.
5. `fmt.Sprintf("%v", value)` formats each cell value.
6. `strings.TrimSpace` removes leading and trailing whitespace from the completed table before the function returns it.

Missing map keys resolve to `nil` and are currently formatted as `<nil>`. The implementation does not escape Markdown-sensitive cell content such as pipes or newlines. Treat these details, column ordering, spacing, and the empty-data error as observable behavior when changing the package.

## Configuration

- `go.mod` is the authoritative module configuration. Preserve the module path unless a migration is explicitly planned, and keep code compatible with its declared `go 1.19` version.
- `.github/workflows/test.yml` is the authoritative continuous integration pipeline. It uses Go 1.19 on `ubuntu-latest` and runs for pushes to every branch.
- `.vscode/settings.json` affects only local editor presentation. Do not make builds or tests depend on VS Code settings.
- The package has no environment variables or runtime configuration files.

## Build and validation pipeline

Run the same core checks locally that CI runs:

```bash
go mod download
go build -v ./...
go test -v ./...
```

Before those checks, format changed Go files:

```bash
gofmt -w gmdtable.go gmdtable_test.go
```

`go build` validates the library but does not produce a committed application binary. Do not add generated binaries or test artifacts to the repository.

## Data flow and component boundaries

```text
Consumer
  │ headers []string + data []map[string]interface{}
  ▼
Convert
  ├── validates that data is nonempty
  ├── uses headers as the column schema and ordering
  ├── reads each row map by header key
  └── formats values into a strings.Builder
  ▼
Markdown table string or error
```

All transformation happens synchronously in `Convert`. The function does not mutate the input slices or maps, perform I/O, retain state, or call external services. Tests should exercise the public function directly rather than reproduce its formatting algorithm.

## Make changes safely

- Keep the public API small and backward compatible. Coordinate changes to the signature or output format across implementation, tests, and `README.md`.
- Add table-driven tests for new cases, especially ordering, value formatting, whitespace, missing keys, and error behavior.
- Compare exact strings when testing Markdown output so spacing and separators remain intentional.
- Add dependencies only when the standard library cannot meet a demonstrated requirement. Run `go mod tidy` after changing imports and commit any resulting module metadata.
- Update the CI workflow only when build requirements change, and keep its Go version aligned with `go.mod`.
- Run `gofmt`, `go build -v ./...`, and `go test -v ./...` before submitting a change.
