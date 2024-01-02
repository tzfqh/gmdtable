# gmdtable

A Go library for converting slices of maps into Markdown tables.

## Overview

The `gmdtable` library provides an easy way to take structured data in the form of a slice of maps and convert it into a Markdown formatted table. This can be particularly useful when generating documentation or reports in Markdown format programmatically.

## Installation

To install `gmdtable`, use the following `go get` command:

```sh
go get -u github.com/tzfqh/gmdtable
```

Replace `github.com/tzfqh/gmdtable` with the actual path to your library.

## Usage

Here is a simple example of how to use `gmdtable`:

```go
package main

import (
 "fmt"
 "github.com/tzfqh/gmdtable"
)

func main() {
 headers := map[string]string{
  "key1": "Title 1",
  "key2": "Title 2",
 }

 data := []map[string]interface{}{
  {"key1": "value1", "key2": 123},
  {"key1": "value2", "key2": 456},
 }

 markdownTable, err := gmdtable.Convert(headers, data)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }

 fmt.Println(markdownTable)
}
```

This will output:

```markdown
| Title 1 | Title 2 |
| ------------ | ------------ |
| value1 | 123 |
| value2 | 456 |
```

## Features

- Convert slices of maps to Markdown tables.
- Customizable headers.
- Error handling for consistent table generation.

## Documentation

For detailed documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/tzfqh/gmdtable) or generate the documentation using `godoc`.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This library is distributed under the MIT license. See the `LICENSE` file.
