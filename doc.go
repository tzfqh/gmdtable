// Package gmdtable converts ordered column headers and map-backed row data into Markdown tables.
//
// Convert uses the header slice as both the table schema and the column order. Cell values are
// formatted with fmt's default representation, then structural Markdown table characters are
// escaped. The package performs no I/O and has no runtime dependencies outside the Go standard
// library.
package gmdtable
