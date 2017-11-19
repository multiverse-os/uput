package validinput

import (
	"reflect"
)

type InputData struct {
	DataType               reflect.Kind
	DataTypeName           string
	Data                   interface{}
	Errors                 []error
	Validations            []string
	ErrorMessages          map[string]string
	ValidationDescriptions map[string]string
	CustomValidations      map[string]interface{}
	Valid                  bool
	// Struct Validation
	//fieldName string
}
