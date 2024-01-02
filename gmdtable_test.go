package gmdtable

import (
	"strings"
	"testing"
)

func TestConvertToMDTable(t *testing.T) {
	headers := []string{"title1", "title2"}
	data := []map[string]interface{}{
		{"title1": "value1", "title2": 123},
		{"title1": "value2", "title2": 456},
	}

	want := strings.TrimSpace(`
| title1 | title2 |
| ------------ | ------------ |
| value1 | 123 |
| value2 | 456 |
`)

	got, err := ConvertToMDTable(headers, data)
	if err != nil {
		t.Errorf("ConvertToMarkdown returned an unexpected error: %v", err)
	}

	if strings.TrimSpace(got) != want {
		t.Errorf("ConvertToMarkdown output was incorrect, got:\n%s\nwant:\n%s", got, want)
	}
}
