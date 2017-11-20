package validstruct

import (
	"testing"
)

func Test_filterValidationTag(t *testing.T) {

	tests := []struct {
		tagType, value, out string
	}{
		{"tname", "presence", "presence"},
		{"tname", "max=0", "max"},
		{"tname", "min=200", "min"},
		{"tvalue", "presence", ""},
		{"tvalue", "max=0", "0"},
		{"tvalue", "min=200", "200"},
	}

}
