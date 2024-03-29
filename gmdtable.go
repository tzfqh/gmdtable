package gmdtable

import (
	"fmt"
	"strings"
)

// ConvertToMDTable takes a slice of maps and returns a string representing a Markdown table.
func ConvertToMDTable(headers map[string]string, data []map[string]interface{}) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("the data slice is empty")
	}

	var markdownTable strings.Builder

	// Write table headers
	markdownTable.WriteString("| ")
	headersLen := len(headers)
	headersIndex := 0
	for _, title := range headers {
		// markdownTable.WriteString(title + " | ")
		if headersLen == (headersIndex + 1) {
			markdownTable.WriteString(title + " |")
		} else {
			markdownTable.WriteString(title + " | ")
		}
		headersIndex++
	}
	markdownTable.WriteString("\n")

	// Write separator
	markdownTable.WriteString("|")
	for range headers {
		markdownTable.WriteString(" ------------ |")
	}
	markdownTable.WriteString("\n")

	// Write table rows
	for _, row := range data {
		markdownTable.WriteString("| ")
		rowIndex := 0
		for key := range headers {
			value := row[key]
			if headersLen == (rowIndex + 1) {
				markdownTable.WriteString(fmt.Sprintf("%v |", value))
			} else {
				markdownTable.WriteString(fmt.Sprintf("%v | ", value))
			}
			rowIndex++
		}
		markdownTable.WriteString("\n")
	}

	return markdownTable.String(), nil
}
