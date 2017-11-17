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
	stringData    string
	intData       int
	uintData      uint
	mapData       map[interface{}]interface{}
	errors        []error
	errorMessages map[string]string
	//validations   map[string]ValidateStringFunction
}

func PrintErrors(errors []error) {
	// TODO: Obviously should just be marshalling to JSON and printing
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

// Output function - maybe a function with the ability specify messages
func (input InputData) IsValid() (bool, interface{}, []error) {
	if input.dataType == stringType {
		return (len(input.errors) == 0), input.stringData, input.errors
	} else {
		input.errors = append(input.errors, errors.New("unknown data type"))
		return false, nil, input.errors
	}
}

// Input functions
func IfString(input string) InputData { return InputData{dataType: stringType, stringData: input} }
func IfInt(input int) InputData       { return InputData{dataType: intType, intData: input} }
func IfUInt(input uint) InputData     { return InputData{dataType: uintType, uintData: input} }
func IfMap(input map[interface{}]interface{}) InputData {
	return InputData{dataType: mapType, mapData: input}
}
