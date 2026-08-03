package gmdtable

import (
	"fmt"
	"strings"
)

// Convert returns a Markdown table containing headers and data.
//
// The order of headers determines the column order. Each row is read by header key, and values use
// fmt's default formatting. Pipes, backslashes, and line breaks are escaped so they remain within
// their cells. Convert returns an error when headers or data is empty.
func Convert(headers []string, data []map[string]interface{}) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("the data slice is empty")
	}
	if len(headers) == 0 {
		return "", fmt.Errorf("the headers slice is empty")
	}

	var markdownTable strings.Builder
	var tableStr string

	// Write table headers
	markdownTable.WriteString("| ")
	headersLen := len(headers)
	for headersIndex, title := range headers {
		title = escapeCell(title)
		if headersLen == (headersIndex + 1) {
			markdownTable.WriteString(title + " |")
		} else {
			markdownTable.WriteString(title + " | ")
		}
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
		for _, key := range headers {
			value := escapeCell(fmt.Sprintf("%v", row[key]))
			if headersLen == (rowIndex + 1) {
				markdownTable.WriteString(value + " |")
			} else {
				markdownTable.WriteString(value + " | ")
			}
			rowIndex++
		}
		markdownTable.WriteString("\n")
	}

	tableStr = strings.TrimSpace(markdownTable.String())

	return tableStr, nil
}

func escapeCell(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, "\r\n", "<br>")
	value = strings.ReplaceAll(value, "\r", "<br>")
	value = strings.ReplaceAll(value, "\n", "<br>")
	return strings.ReplaceAll(value, "|", `\|`)
}
