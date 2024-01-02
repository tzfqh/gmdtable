# gmdtable - Go Markdown Table Generator

The `gmdtable` library provides a simple way to convert slices of maps into well-formatted Markdown tables. This can be particularly useful when generating Markdown documentation or output from Go programs.

## Installation

To start using `gmdtable`, install Go and run `go get`:

```sh
go get github.com/tzfqh/gmdtable
```

This will retrieve the library.

## Usage

To use the `gmdtable` library, you need to import it in your Go application:

```go
import "github.com/tzfqh/gmdtable"
```

Here is a simple example of how to use `gmdtable`:

```go
package main

import (
	"fmt"
	"github.com/tzfqh/gmdtable"
)

func main() {
	headers := []string{"ID", "Name", "Age"}
	data := []map[string]interface{}{
		{"ID": 1, "Name": "John Doe", "Age": 30},
		{"ID": 2, "Name": "Jane Smith", "Age": 25},
	}

	table, err := gmdtable.Convert(headers, data)
	if err != nil {
		fmt.Println("Error generating table:", err)
		return
	}

	fmt.Println(table)
}
```

The output will be:

```markdown
| ID | Name      | Age |
| ------------ | ------------ | ------------ |
| 1  | John Doe  | 30  |
| 2  | Jane Smith | 25  |
```

Preview:

| ID | Name      | Age |
| ------------ | ------------ | ------------ |
| 1  | John Doe  | 30  |
| 2  | Jane Smith | 25  |

## Features

- Easy to use: Just pass a slice of headers and a slice of map data.
- Customizable: You can easily change the headers and the data will adjust accordingly.
- Error handling: The function will return an error if the data slice is empty.

## Documentation

For detailed documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/tzfqh/gmdtable) or generate the documentation using `godoc`.

## Contributing

Feel free to dive in! [Open an issue](https://github.com/tzfqh/gmdtable/issues/new) or submit PRs.

## License

This library is distributed under the MIT license. See the `LICENSE` file.
