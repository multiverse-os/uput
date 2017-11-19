package input

import (
	"reflect"
)

type InputData struct {
	DataType reflect.Kind
	Data     interface{}
	//fieldName string

	StringData string
	IntData    int
	UintData   uint
	MapData    map[interface{}]interface{}

	Errors        []error
	ErrorMessages map[string]string

	Validations map[string]interface{}
	Valid       bool
}

//
// Generic Output Function
//func (input InputData) IsValid() (bool, interface{}, []error) {
//	return (len(input.Errors) == 0), input.Data, input.Errors
//}
