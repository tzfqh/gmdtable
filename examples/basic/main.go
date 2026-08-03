// Command basic demonstrates the public Markdown-table conversion API.
package main

import (
	"fmt"
	"log"

	"github.com/tzfqh/gmdtable"
)

func main() {
	headers := []string{"Name", "Language"}
	data := []map[string]interface{}{{"Name": "gmdtable", "Language": "Go"}}

	table, err := gmdtable.Convert(headers, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(table)
}
