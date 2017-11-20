package validinput

import (
	"errors"
	"fmt"     // DEV
	"strconv" // DEV
	"strings"
)

type InputError struct {
	DataType              string
	Key                   string
	Message               string
	ValidationDescription string
	ValidationValues      []string
}

//
// Implement Error Interface
//==================================================================

func (err InputError) Error() string {
	switch len(err.ValidationValues) {
	case 0:
		return err.DataType + ": " + err.Message
	case 1:
		return err.DataType + ": " + err.Message + ": " + err.ValidationValues[0]
	case 2:
		return err.DataType + ": " + err.Message + ": " + err.ValidationValues[0] + " - " + err.ValidationValues[1]
	default:
		return err.DataType + ": " + err.Message + ": [ " + strings.Join(err.ValidationValues, ", ") + " ]"
	}
}

//
// Error Helpers
//==================================================================
func (input InputData) LastInputError() InputError {
	if len(input.InputErrors) > 0 {
		return input.InputErrors[len(input.InputErrors)-1]
	}
	return InputError{}
}
func (input InputData) Errors() (outputErrors []error) {
	for _, inputError := range input.InputErrors {
		outputErrors = append(outputErrors, errors.New((inputError.Error())))
	}
	return outputErrors
}

//
// Append Errors
//==================================================================
func (input InputData) AppendError(key string, values []string) InputData {
	err := InputError{
		DataType:              input.DataTypeName,
		Key:                   key,
		Message:               input.ValidationText[key].Error,
		ValidationDescription: input.ValidationText[key].Description,
		ValidationValues:      values,
	}
	input.InputErrors = append(input.InputErrors, err)
	return input
}

//
// Development Printing (remove later, don't assume logging style)
//==================================================================
func (input InputData) PrintErrors() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(input.InputErrors) > 0 {
		fmt.Println("{")
		fmt.Println("  \"error_count\": " + strconv.Itoa(len(input.InputErrors)) + ",")
		fmt.Println("  \"errors\": {")
		for _, err := range input.InputErrors {
			fmt.Println("    \"" + err.Error() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}
