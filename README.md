# gmdtable

`gmdtable` converts ordered Go map data into Markdown tables. The module-root package has one public
function, uses only the Go standard library at runtime, and preserves the header slice as the table's
column order.

## Install

Go 1.19 or later is supported.

```bash
go get github.com/tzfqh/gmdtable
```

```go
import "github.com/tzfqh/gmdtable"
```

## Usage

```go
headers := []string{"Name", "Language"}
data := []map[string]interface{}{
	{"Name": "gmdtable", "Language": "Go"},
}

table, err := gmdtable.Convert(headers, data)
if err != nil {
	log.Fatal(err)
}
fmt.Println(table)
```

The example prints:

```markdown
| Name | Language |
| ------------ | ------------ |
| gmdtable | Go |
```

`Convert` returns an error when `headers` or `data` is empty. Values use `fmt`'s default formatting,
missing keys render as `<nil>`, and structural pipes, backslashes, and line breaks are escaped so the
content remains in its cell.

Run the complete example with `go run ./examples/basic`.

## Website and documentation

- [Public website](https://tzfqh.github.io/gmdtable/)
- [Executable examples](https://tzfqh.github.io/gmdtable/examples/)
- [Generated API reference](https://tzfqh.github.io/gmdtable/api/)
- [pkg.go.dev API documentation](https://pkg.go.dev/github.com/tzfqh/gmdtable)
- [Basic example](examples/basic/main.go)

The examples page renders `examples/basic/main.go` and `example_test.go` directly from checked-in Go
source. The API page extracts exported declarations and GoDoc from the module-root package.

## Development

Run the package checks from the repository root:

```bash
gofmt -w .
go mod tidy
go test ./...
go test -race ./...
go vet ./...
go build ./...
go run ./examples/basic
```

Generate and validate the final GitHub Pages artifact:

```bash
go run ./cmd/sitegen build
go run ./cmd/sitegen validate
```

Preview the production-like artifact at `http://127.0.0.1:4173/gmdtable/`:

```bash
go run ./cmd/sitegen preview
```

`dist/pages` is generated and ignored by Git. Update `site.config.json`, Go source, templates, or site
assets, then rebuild instead of editing generated Pages files.

## License

The Go package and project-authored website code are available under the [MIT License](LICENSE).
Vendored Bootstrap notices are in [guides/THIRD_PARTY_NOTICES.md](guides/THIRD_PARTY_NOTICES.md).
