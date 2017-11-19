package validinput

import (
	"reflect"
)

type InputData struct {
	DataType      reflect.Kind
	Input         interface{}
	Errors        []error
	ErrorMessages map[string]string
	Validations   []string
	Valid         bool
	// Struct Validation
	//fieldName string
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
