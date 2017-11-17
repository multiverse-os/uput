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

// InputData.InputDataFunc(): Validation Functions
//type InputDataFunc func() InputData

type InputDataFunc func(input InputData) ValidateInput

//type ValidateStringFunction func(input string) (output string, errors []error)
type InputData struct {
	DataType dataType
	//fieldName string

	StringData string
	IntData    int
	//UintData   uint
	//BoolType   bool
	TimeType *time.Time
	MapData  map[interface{}]interface{}

	Errors        []error
	ErrorMessages map[string]string

	Validations map[string]InputDataFunc
	IsValid     bool
}

// Extend this to work with: [struct, single-variable]
type Validate interface {
	//InputDataFunc() InputData

	// InputData
	Value() interface{}
	SetValue(value interface{}) bool
	IsType() dataType
	// Errors
	AddError(key, value string) InputData
	PrintErrors()
	// Validations
	Validations() map[string]InputData
	IsValid() (bool, interface{}, []error)
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
func If(input interface{}) InputData {
	switch reflect.ValueOf(input).Kind() {
	case reflect.Array, reflect.String:
		fmt.Println("reflect.ValueOf(input).Kind(): ", reflect.ValueOf(input).Kind())
		fmt.Println("reflect.ValueOf(input): ", reflect.ValueOf(input))

		//IfString(string(input))
	}
	return InputData{
		DataType: MapType,
		//mapData:  input,
	}
}

//
// Panic & Recover (similar to try/rescue in Ruby)
/////////////////////////////////////////////////////
//
// https://golang.org/src/encoding/json/decode.go
// Line 171 has a working example in the json lib,
// a library that is worth reading because it also
// uses the tag system and will show the preferred
// way to implement it.
//

//
// Type switch without reflect library
///////////////////////////////////////
// Pkg https://golang.org/pkg/reflect
//
// It appears most people only know how to do data type comparision using the
// reflect library. Somet validation libraries boast how they avoid this, and
// it appears that type validation can be done without any additional libraries by
// doing the following:
//
// 			var t interface{}
// 			t = functionOfRandomType()
// 			switch t := t.(type) {
//      // First time I've seen default used first
// 			default:
// 				fmt.Println("unexpected type %T", t)
// 			case bool:
// 				fmt.Println("boolean data type %T", t)
// 			case int:
// 				fmt.Println("int data type %T", t)
// 			case *bool:
// 				fmt.Println("pointer (address) boolean data type %T", t)
// 			case *int:
// 				fmt.Println("pointer (address) int data type %T", t)
//
// **Arguments to avoid reflect?**
// "The reflect package is great way to make descision at runtime. However, we
// should be aware that it gives us some performance penalties. I would try to
// avoid using reflection. It’s not idiomatic, but it’s very powerfull in particular
// cases. Do not forget to follow the laws of reflection."
// http://blog.ralch.com/tutorial/golang-reflection/
//
// A separate argument against reflect can be found here:
// http://www.jerf.org/iri/post/2945
//
// ---other-reflect-topics---
//
// For types that support the equality operation, you can just compare interface{}
// variables holding the zero value and field value. Something like this:
//
// 			v.Interface() == reflect.Zero(v.Type()).Interface()
//
// For functions, maps and slices though, this comparison will fail, so we still
// need to include some special casing. Further more, while arrays and structs are
// comparable, the comparison will fail if they contain non-comparable types. So
// you probably need something like:
//
// 		func isZero(v reflect.Value) bool {
// 		    switch v.Kind() {
// 		    case reflect.Func, reflect.Map, reflect.Slice:
// 		        return v.IsNil()
// 		    case reflect.Array:
// 		        z := true
// 		        for i := 0; i < v.Len(); i++ {
// 		            z = z && isZero(v.Index(i))
// 		        }
// 		        return z
// 		    case reflect.Struct:
// 		        z := true
// 		        for i := 0; i < v.NumField(); i++ {
// 		            z = z && isZero(v.Field(i))
// 		        }
// 		        return z
// 		    }
// 		    // Compare other types directly:
// 		    z := reflect.Zero(v.Type())
// 		    return v.Interface() == z.Interface()
// 		}
//
// ---other-reflect-topics---
//
// Laws Of Reflection: https://blog.golang.org/laws-of-reflection
// Here again are the laws of reflection:
//
//    Reflection goes from interface value to reflection object.
//
//    Reflection goes from reflection object to interface value.
//
//    To modify a reflection object, the value must be settable.
//
// Once you understand these laws reflection in Go becomes much easier to use,
// although it remains subtle. It's a powerful tool that should be used with
// care and avoided unless strictly necessary.
//
// ---other-reflect-topics---
//
// GoFix
//
// switch f := value; f.Kind() {
// case reflect.Bool:
//     p.fmtBool(f.Bool(), verb, field)
// case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//     p.fmtInt64(f.Int(), verb, field)
// // ...
// case reflect.Array, reflect.Slice:
//     // Byte slices are special.
//     if f.Type().Elem().Kind() == reflect.Uint8 {
//         // ...
//     }
// //
// }
//
// Nearly every line above changed in some small way. The changes involved in
// the rewrite are extensive but nearly entirely mechanical, just the kind of
// thing that computers are great at doing.

//
// Basic Validation Types
///////////////////////////////////////////////////////////////////////////////////
//
// Eventually valid should be able to handle any type of input and properly
// classify it, have a working validation sub-package.
//
// bool
// string
//
// Numeric types:
//
// uint        either 32 or 64 bits
// int         same size as uint
// uintptr     an unsigned integer large enough to store the uninterpreted bits of
//             a pointer value
// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
//
// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers
//             (-9223372036854775808 to 9223372036854775807)
//
// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers
//
// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts
//
// byte        alias for uint8
// rune        alias for int32 (represents a Unicode code point)
