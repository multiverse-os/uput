package valid

import (
	"errors"
	"fmt"
	"strconv"
)

type dataType int

const (
	stringType dataType = iota
	intType
	uintType
	mapType
)

//type ValidateStringFunction func(input string) (output string, errors []error)
type InputData struct {
	dataType
	fieldName     string
	stringData    string
	intData       int
	uintData      uint
	mapData       map[interface{}]interface{}
	errors        []error
	errorMessages map[string]string
	//validations   map[string]ValidateStringFunction
}

//
// Errors

// Development Printing (remove later, don't assume logging style)
func PrintErrors(errors []error) {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(errors) > 0 {
		fmt.Println("{")
		fmt.Println("  \"error_count\": \"" + strconv.Itoa(len(errors)) + ",")
		fmt.Println("  \"errors\": {")
		for _, err := range errors {
			fmt.Println("    \"string\": \"" + err.Error() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}

func (input InputData) ErrorMessages(errorMessages map[string]string) InputData {
	for key, value := range errorMessages {
	}
}

//
// Output function - maybe a function with the ability specify messages
func (input InputData) IsValid() (bool, interface{}, []error) {
	if input.dataType == stringType {
		return (len(input.errors) == 0), input.stringData, input.errors
	} else {
		input.errors = append(input.errors, errors.New("unknown data type"))
		return false, nil, input.errors
	}
}

//
// Input functions
func IfString(input string) InputData {
	return InputData{
		dataType:      stringType,
		stringData:    input,
		errorMessages: stringErrorMessages,
	}
}
func IfInt(input int) InputData {
	return InputData{
		dataType: intType,
		intData:  input,
	}
}
func IfUInt(input uint) InputData {
	return InputData{
		dataType: uintType,
		uintData: input,
	}
}
func IfMap(input map[interface{}]interface{}) InputData {
	return InputData{
		dataType: mapType,
		mapData:  input,
	}
}
