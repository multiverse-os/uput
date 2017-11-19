package validinput

import (
	"fmt" // DEV
	"reflect"
	"strconv" // DEV
)

type InputData struct {
	DataType      reflect.Kind
	Data          interface{}
	Errors        []error
	ErrorMessages map[string]string
	Validations   []string
	Valid         bool
	// Struct Validation
	//fieldName string
}

// Development Printing (remove later, don't assume logging style)
func PrintValidations(validations []string) {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(validations) > 0 {
		fmt.Println("{")
		fmt.Println("  \"validation_count\": \"" + strconv.Itoa(len(validations)) + ",")
		fmt.Println("  \"validations\": {")
		for _, v := range validations {
			fmt.Println("    \"string\": \"" + v + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}

func (input InputData) AppendValidation(key string) InputData {
	// TODO: Validate if key is valid, perhaps have a availableValidations array
	// TODO: Validate if key does not already exist in Validations?
	input.Validations = append(input.Validations, key)
	return input
}

//
// Generic Output Function
//func (input InputData) IsValid() (bool, interface{}, []error) {
//	return (len(input.Errors) == 0), input.Data, input.Errors
//}
