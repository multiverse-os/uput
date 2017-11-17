package valid

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type dataType int

const (
	stringType dataType = iota
	intType
	uintType
	mapType
	boolType
	timeType
)

//type ValidateStringFunction func(input string) (output string, errors []error)
type InputData struct {
	dataType
	input       interface{}
	fieldName   string
	validations int

	stringData string
	intData    int
	uintData   uint
	boolType   bool
	timeType   *time.Time
	mapData    map[interface{}]interface{}

	errors        []error
	errorMessages map[string]string
	//validations   map[string]ValidateStringFunction
}

//
// Error Functions
func (input InputData) AppendErrorMessages(errorMessages map[string]string) InputData {
	for key, value := range errorMessages {
		// valid.IfKey.IsBetween(2, 12)
		if len(key) >= 2 && len(key) <= 12 {
			// valid.IfValue.IsBetween(2, 32)
			if len(key) >= 2 && len(key) <= 32 {
				input.errorMessages[key] = value
			}
		}
	}
	return input
}

func (input InputData) AddErrorMessage(key string) InputData {
	return append(input.errors, errors.New(input.errorMessages[key]))
}

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

//
// Output Function
func (input InputData) IsValid() (bool, interface{}, []error) {
	if input.dataType == stringType {
		return (len(input.errors) == 0), input.stringData, input.errors
	} else {
		input.errors = append(input.errors, errors.New("unknown data type"))
		return false, nil, input.errors
	}
}

//
// Input Functions
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
	// If Map Length 0 (IsEmpty)
	if len(input) > 0 {
		// In A Map Each Key/Value Set Is An Entry
		firstKey := (reflect.ValueOf(input).MapKeys())[0]
		// Prepare InputData Based On Key Type/Kind
		switch firstKey.Kind() {
		case reflect.Array, reflect.String:
			// TODO: Add stringErrorMessages for [key validations]
		}

		// Prepare InputData Based On Value Type/Kind
		//switch input[firstKey].Kind() {
		//case reflect.Array, reflect.String:
		//	// TODO: Add stringErrorMessages for [value validations]
		//}
	} else {
		// Empty/Nil Map will likely fail most validaitons
	}
	return InputData{
		dataType: mapType,
		mapData:  input,
	}
}

// Generic/Dynamic Input Function
func Validate(input interface{}) InputData {
	//switch reflect.ValueOf(input).Kind() {
	//case reflect.Array, reflect.String:
	//	IfString(string(input))
	//}
	return InputData{
		dataType: mapType,
		//mapData:  input,
	}
}
