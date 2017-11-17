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
	StringType dataType = iota
	IntType
	UintType
	MapType
	BoolType
	TimeType
)

//type ValidateStringFunction func(input string) (output string, errors []error)
type InputData struct {
	DataType  dataType
	input     interface{}
	fieldName string

	StringData string
	IntData    int
	UintData   uint
	BoolType   bool
	TimeType   *time.Time
	MapData    map[interface{}]interface{}

	Errors        []error
	ErrorMessages map[string]string
	//validations   map[string]ValidateStringFunction
}

type Input interface {
	IsValid()
}

//
// Error Functions
func (input InputData) AddError(key, value string) InputData {
	message := input.ErrorMessages[key]
	if len(value) > 0 {
		message += ": [ " + value + " ]"
	}
	input.Errors = append(input.Errors, errors.New(message))
	return input
}

// Development Printing (remove later, don't assume logging style)
func (input InputData) PrintErrors() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(input.Errors) > 0 {
		fmt.Println("{")
		fmt.Println("  \"error_count\": \"" + strconv.Itoa(len(input.Errors)) + ",")
		fmt.Println("  \"errors\": {")
		for _, err := range input.Errors {
			fmt.Println("    \"string\": \"" + err.Error() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}

//
// Output Function
func (input InputData) IsValid() (bool, interface{}, []error) {
	if input.DataType == StringType {
		return (len(input.Errors) == 0), input.StringData, input.Errors
	} else {
		input.Errors = append(input.Errors, errors.New("unknown data type"))
		return false, nil, input.Errors
	}
}

//
// Input Functions
func IfString(input string) InputData {
	return InputData{
		DataType:      StringType,
		StringData:    input,
		ErrorMessages: make(map[string]string),
	}
}
func IfInt(input int) InputData {
	return InputData{
		DataType: IntType,
		IntData:  input,
	}
}
func IfUInt(input uint) InputData {
	return InputData{
		DataType: UintType,
		UintData: input,
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
		DataType: MapType,
		MapData:  input,
	}
}

// Generic/Dynamic Input Function
func Validate(input interface{}) InputData {
	//switch reflect.ValueOf(input).Kind() {
	//case reflect.Array, reflect.String:
	//	IfString(string(input))
	//}
	return InputData{
		DataType: MapType,
		//mapData:  input,
	}
}
