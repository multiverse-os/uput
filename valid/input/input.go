package validinput

import (
	"reflect"
)

type InputData struct {
	DataType               reflect.Kind
	DataTypeName           string
	Data                   interface{}
	Errors                 []error
	ErrorMessages          map[string]string
	Validations            []string
	ValidationDescriptions map[string]string
	Valid                  bool
	// Struct Validation
	//fieldName string
}

//
// Generic Output Function
//func (input InputData) IsValid() (bool, interface{}, []error) {
//	return (len(input.Errors) == 0), input.Data, input.Errors
//}
