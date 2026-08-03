package gmdtable_test

import (
	"testing"

	"github.com/tzfqh/gmdtable"
)

func TestConvert(t *testing.T) {
	t.Parallel()

	headers := []string{"title1", "title2"}
	data := []map[string]interface{}{
		{"title1": "value1", "title2": 123},
		{"title1": "value2", "title2": 456},
	}
	want := "| title1 | title2 |\n| ------------ | ------------ |\n| value1 | 123 |\n| value2 | 456 |"

	got, err := gmdtable.Convert(headers, data)
	if err != nil {
		t.Fatalf("Convert returned an unexpected error: %v", err)
	}
	if got != want {
		t.Fatalf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertUsesHeaderOrderAndFormatsMissingValues(t *testing.T) {
	t.Parallel()

	headers := []string{"second", "missing", "first"}
	data := []map[string]interface{}{{"first": 1, "second": true}}
	want := "| second | missing | first |\n| ------------ | ------------ | ------------ |\n| true | <nil> | 1 |"

	got, err := gmdtable.Convert(headers, data)
	if err != nil {
		t.Fatalf("Convert returned an unexpected error: %v", err)
	}
	if got != want {
		t.Fatalf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertRejectsEmptyData(t *testing.T) {
	t.Parallel()

	got, err := gmdtable.Convert([]string{"name"}, nil)
	if err == nil || err.Error() != "the data slice is empty" {
		t.Fatalf("Convert error = %v, want %q", err, "the data slice is empty")
	}
	if got != "" {
		t.Fatalf("Convert result = %q, want empty string", got)
	}
}

func TestConvertRejectsEmptyHeaders(t *testing.T) {
	t.Parallel()

	got, err := gmdtable.Convert(nil, []map[string]interface{}{{"name": "gmdtable"}})
	if err == nil || err.Error() != "the headers slice is empty" {
		t.Fatalf("Convert error = %v, want %q", err, "the headers slice is empty")
	}
	if got != "" {
		t.Fatalf("Convert result = %q, want empty string", got)
	}
}

func TestConvertEscapesStructuralCellContent(t *testing.T) {
	t.Parallel()

	headers := []string{`Name|Path`}
	data := []map[string]interface{}{{"Name|Path": "first line\r\nsecond | C:\\table"}}
	want := "| Name\\|Path |\n| ------------ |\n| first line<br>second \\| C:\\\\table |"

	got, err := gmdtable.Convert(headers, data)
	if err != nil {
		t.Fatalf("Convert returned an unexpected error: %v", err)
	}
	if got != want {
		t.Fatalf("Convert() = %q, want %q", got, want)
	}
}
