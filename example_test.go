package gmdtable_test

import (
	"fmt"

	"github.com/tzfqh/gmdtable"
)

func ExampleConvert() {
	headers := []string{"Name", "Language"}
	data := []map[string]interface{}{{"Name": "gmdtable", "Language": "Go"}}

	table, err := gmdtable.Convert(headers, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(table)
	// Output:
	// | Name     | Language |
	// | :------- | :------- |
	// | gmdtable | Go       |
}
