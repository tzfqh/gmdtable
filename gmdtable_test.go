package gmdtable

import (
	"strings"
	"testing"
)

func TestConvertToMDTable(t *testing.T) {
	headers := map[string]string{"key1": "title1", "key2": "title2"}
	data := []map[string]interface{}{
		{"key1": "value1", "key2": 123},
		{"key1": "value2", "key2": 456},
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
