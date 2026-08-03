package gmdtable

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

// Convert returns a Markdown table containing headers and data.
//
// The order of headers determines the column order. Each row is read by header key, and values use
// fmt's default formatting. Pipes, backslashes, and line breaks are escaped so they remain within
// their cells. Columns are padded to their widest cell. Convert returns an error when headers or
// data is empty.
func Convert(headers []string, data []map[string]interface{}) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("the data slice is empty")
	}
	if len(headers) == 0 {
		return "", fmt.Errorf("the headers slice is empty")
	}

	rows := make([][]string, len(data)+1)
	rows[0] = make([]string, len(headers))
	widths := make([]int, len(headers))
	for column, header := range headers {
		rows[0][column] = escapeCell(header)
		widths[column] = separatorWidth
		if width := displayWidth(rows[0][column]); width > widths[column] {
			widths[column] = width
		}
	}
	for rowIndex, dataRow := range data {
		rows[rowIndex+1] = make([]string, len(headers))
		for column, key := range headers {
			value := escapeCell(fmt.Sprintf("%v", dataRow[key]))
			rows[rowIndex+1][column] = value
			if width := displayWidth(value); width > widths[column] {
				widths[column] = width
			}
		}
	}

	var markdownTable strings.Builder
	writeRow(&markdownTable, rows[0], widths)
	markdownTable.WriteByte('\n')
	markdownTable.WriteString("|")
	for _, width := range widths {
		markdownTable.WriteString(" :")
		markdownTable.WriteString(strings.Repeat("-", width-1))
		markdownTable.WriteString(" |")
	}
	for _, row := range rows[1:] {
		markdownTable.WriteByte('\n')
		writeRow(&markdownTable, row, widths)
	}

	return markdownTable.String(), nil
}

const separatorWidth = 4 // A leading colon followed by Markdown's minimum three hyphens.

var widthCondition = &runewidth.Condition{
	EastAsianWidth:     false,
	StrictEmojiNeutral: true,
}

func writeRow(table *strings.Builder, cells []string, widths []int) {
	table.WriteString("|")
	for column, cell := range cells {
		table.WriteByte(' ')
		table.WriteString(cell)
		table.WriteString(strings.Repeat(" ", widths[column]-displayWidth(cell)))
		table.WriteString(" |")
	}
}

func displayWidth(value string) int {
	return widthCondition.StringWidth(value)
}

func escapeCell(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, "\r\n", "<br>")
	value = strings.ReplaceAll(value, "\r", "<br>")
	value = strings.ReplaceAll(value, "\n", "<br>")
	return strings.ReplaceAll(value, "|", `\|`)
}
