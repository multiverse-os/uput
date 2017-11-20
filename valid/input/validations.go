package validinput

import (
	"fmt"
	"strconv" // DEV
	"strings"
)

type Validation struct {
	DataType    string
	Key         string
	Description string
	Values      []string
}

//
// Return String
//==================================================================
func (v Validation) String() string {
	switch len(v.Values) {
	case 0:
		return v.DataType + ": " + v.Description
	case 1:
		return v.DataType + ": " + v.Description + ": " + v.Values[0]
	case 2:
		return v.DataType + ": " + v.Description + ": " + v.Values[0] + "-" + v.Values[1]
	default:
		return v.DataType + ": " + v.Description + ": [ " + strings.Join(v.Values, ", ") + " ]"
	}
}

//
// Validation Helpers
//==================================================================
func (input InputData) LastValidation() Validation {
	if len(input.Validations) > 0 {
		return input.Validations[len(input.Validations)-1]
	}
	return Validation{}
}

//
// Append Validation
//==================================================================
func (input InputData) AppendValidation(key string, values []string) InputData {
	validation := Validation{
		DataType:    input.DataTypeName,
		Key:         key,
		Description: input.ValidationText[key].Description,
		Values:      values,
	}
	input.Validations = append(input.Validations, validation)
	return input
}

//
// Development Printing (remove later, don't assume logging style)
//==================================================================
func (input InputData) PrintValidations() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(input.Validations) > 0 {
		fmt.Println("{")
		fmt.Println("  \"validation_count\": " + strconv.Itoa(len(input.Validations)) + ",")
		fmt.Println("  \"validations\": {")
		for _, v := range input.Validations {
			fmt.Println("    \"" + v.String() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}
