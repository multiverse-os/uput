// Utility functions for working with text
package textual

import (
	"testing"
)

var Format string = "\ninput:    %q\nexpected: %q\noutput:   %q"

type Test struct {
	input    string
	expected string
}

var truncations = []Test{
	{"This is a very long string which should be truncated", "This is a very long string which shou…"},
	{"<p>Some text</p>", "<p>Some text</p>"},
	{"So", "So"},
	{"Something some more", "Something some more"},
	{"This is a very long string which can be truncated.This is a very long string which should be truncated.This is a very long string which should be truncated.This is a very long string which should be truncated.This is a very long string which should be truncated.This is a very long string which should be truncated", "This is a very long string which can …"},
}

func TestTruncate(t *testing.T) {
	for _, test := range truncations {
		output := Truncate(test.input, 40)
		if output != test.expected {
			t.Fatalf(Format, test.input, test.expected, output)
		}
	}
}

var plurals = []Test{
	{"model", "models"},
	{"sheep", "sheep"},
	{"page", "pages"},
	{"pedant", "pedants"},
	{"lase", "lases"},
	{"datum", "data"},
}

func TestPlurals(t *testing.T) {
	for _, test := range plurals {
		output := ToPlural(test.input)
		if output != test.expected {
			t.Fatalf(Format, test.input, test.expected, output)
		}
	}
}

// Default is to translate to public fields
var colNames = []Test{
	{"id", "Id"},
	{"updated_at", "UpdatedAt"},
	{"translate_me_please", "TranslateMePlease"},
	{"long_strange_col_name", "LongStrangeColName"},
}

var colNamesPrivate = []Test{
	{"id", "id"},
	{"updated_at", "updatedAt"},
	{"translate_me_please", "translateMePlease"},
	{"long_strange_col_name", "longStrangeColName"},
}

func TestSnakeToCamel(t *testing.T) {
	// Test public fields
	for _, test := range colNames {
		output := ToCamel(test.input)
		if output != test.expected {
			t.Fatalf(Format, test.input, test.expected, output)
		}
	}

	// Test private fields
	for _, test := range colNamesPrivate {
		output := ToCamel(test.input, true)
		if output != test.expected {
			t.Fatalf(Format, test.input, test.expected, output)
		}
	}
}

var fieldNames = []Test{
	{"Id", "id"},
	{"sheepField", "sheep_field"},
	{"LongStrangeColName", "long_strange_col_name"},
}

func TestCamelToSnake(t *testing.T) {
	for _, test := range fieldNames {
		output := ToSnake(test.input)
		if output != test.expected {
			t.Fatalf(Format, test.input, test.expected, output)
		}
	}
}
