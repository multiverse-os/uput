package valid

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

//
// Reflect Notes:
//
// reflect.Value has a function Interface() that converts it to interface{}
// 		Keep in mind that a struct can not have more than 32 methods:
// 		func (v Value) NumMethod() int
// 			NumMethod returns the number of exported methods in the value's method set.
//
// 		func (v Value) OverflowInt(x int64) bool
// 			OverflowInt reports whether the int64 x cannot be represented by v's type.
// 			It panics if v's Kind is not Int, Int8, int16, Int32, or Int64.
//
// 		Use this for Uint, Float, Int
//
// 		func (v Value) Pointer() uinptr
//      Pointer returns v's value as a uintptr. It returns uintptr instead of
//      unsafe.Pointer so that code using reflect cannot obtain unsafe.Pointers
//      without importing the unsafe package explicitly. It panics if v's Kind is
// 			not Chan, Func, Map, Ptr, Slice, or UnsafePointer.
//
// 			The only guarantee is that the result is zero if and only if v is a nil
// 			func Value.
//
// 			If v's Kind is Slice, the returned pointer is to the first element of the
// 			slice. If the slice is nil the returned value is 0. If the slice is empty
// 			but non-nil the return value is non-zero.
//
//
// 	==/!\[ WARNING! ][ Bug Known In Reflect Package(go1.9.2) ]========================
//
// 	 FieldByName and related functions consider struct field names to be equal if
//   the names are equal, even if they are unexported names originating in different
//   packages. The practical effect of this is that the result of t.FieldByName("x")
//   is not well defined if the struct type t contains multiple fields named x
//   (embedded from different packages). FieldByName may return one of the fields
//   named x or may report that there are none. See golang.org/issue/4876 for more
//   details.
//
// 	========================/!\[ WARNING! ][ Bug Known In Reflect Package(go1.9.2) ]==
//
//
//  ==/!\[ WARNING! ][ Thoughtless use of pointers is not secure! ]===================
//  if size > 0 && lastzero == size {
///   // TODO: THINK ABOUT THIS!
//		// This is a non-zero sized struct that ends in a
//		// zero-sized field. We add an extra byte of padding,
//		// to ensure that taking the address of the final
//		// zero-sized field can't manufacture a pointer to the
//		// next object in the heap. See issue 9401.
//		size++
//	}
//
//	var typ *structType
//	var ut *uncommonType
//	switch {
//	case len(methods) == 0:
//		t := new(structTypeUncommon)
//		typ = &t.structType
//		ut = &t.u
//	case len(methods) <= 4:
//		t := new(structTypeFixed4)
//		typ = &t.structType
//		ut = &t.u
//		copy(t.m[:], methods)
//	case len(methods) <= 8:
//		t := new(structTypeFixed8)
//		typ = &t.structType
//		ut = &t.u
//		copy(t.m[:], methods)
//	case len(methods) <= 16:
//		t := new(structTypeFixed16)
//		typ = &t.structType
//		ut = &t.u
//		copy(t.m[:], methods)
//	case len(methods) <= 32:
//		t := new(structTypeFixed32)
//		typ = &t.structType
//		ut = &t.u
//		copy(t.m[:], methods)
//	default:
//		panic("reflect.StructOf: too many methods")
//	}

//
// Data Kind
//////////////////////////////////////////////////////////////////////////////
// Pkg ~@250 line on https://golang.org/src/reflect/type.go
//
// TODO: Don't need to initialize our own datatypes if we are already calling
// reflect since it has the following:
//
// These data structures are known to the compiler
// (../../cmd/internal/gc/reflect.go).
//
// A few are known to ../runtime/type.go to convey to debuggers.
// They are also known to ../runtime/type.go.
//
// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
//
//    tflag is used by an rtype to signal what extra type information
//    type tflag uint8
// 		type Kind uint
// 		const (
// 			Invalid Kind = iota
// 			Bool
// 			Int
// 			Int8
// 			Int16
// 			Int32
// 			Int64
// 			Uint
// 			Uint8
// 			Uint16
// 			Uint32
// 			Uint64
// 			Uintptr
// 			Float32
// 			Float64
// 			Complex64
// 			Complex128
// 			Array
// 			Chan
// 			Func
// 			Interface
// 			Map
// 			Ptr
// 			Slice
// 			String
// 			Struct
// 			UnsafePointer
// 		)
//

// TODO: I don't to count over compared value
//
// Possible ways to not count more than
// needed when checking.
//
// Size() uintptr, Bite() int

// May just be useful reference
var kindNames = []string{
	Invalid: "invalid",
	Bool:    "bool",
	Int:     "int",
	//Int8:       "int8",
	//Int16:      "int16",
	//Int32:      "int32",
	//Int64:      "int64",
	Uint: "uint",
	//Uint8:      "uint8",
	//Uint16:     "uint16",
	//Uint32:     "uint32",
	//Uint64:     "uint64",
	//Uintptr:    "uintptr",
	//Float32:    "float32",
	//Float64:    "float64",
	//Complex64:  "complex64",
	//Complex128: "complex128",
	//Array:      "array",
	//Chan:       "chan",
	//Func:       "func",
	//Interface:  "interface",
	//Map:        "map",
	//Ptr:        "ptr",
	//Slice:      "slice",
	String: "string",
	//Struct:        "struct",
	//UnsafePointer: "unsafe.Pointer",
}

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
		value := reflect.ValueOf(input)
		fmt.Println("Input type using basic type check: %T", input.(type))
		fmt.Println("Reflected.ValueOf() type using basic type check: %T", value.(type))
		fmt.Println("Reflected.ValueOf().Kind() type using basic type check: %s", value.Kind())
		firstKey := (value.MapKeys())[0]
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

// TODO: Should consider using reflect.Value().MustBe()
// or at least look at the strategy used:
// https://golang.org/src/reflect/value.go

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
// The most common way to handle type differentiation and type checking is using
// the *reflect* package. This package is so ubiquitiuos that it is found throughout
// the Go source, for example, it is in the commonly used json package.
//
// An Example of using it can also be found:
// https://golang.org/src/reflect/value.go @ 320
//
// Additionally @ line 1021 Len() function shows how Len is fo-und
// for Array, Chan, Map, Slice, or String
// func (v Value) Len() int {
// 	k := v.kind()
// 	switch k {
//		return (*sliceHeader)(v.ptr).Len
//	case String:
//		// String is bigger than a word; assume flagIndir.
//		return (*stringHeader)(v.ptr).Len
//	}
//
// MapIndex returns the value associated with key in the map v.
// It panics if v's Kind is not Map.
// It returns the zero Value if key is not found in the map or if v represents a nil map.
// As in Go, the key's value must be assignable to the map's key type.
// func (v Value) MapIndex(key Value) Value
// v.mustBe(Map)
// 	tt := (*mapType)(unsafe.Pointer(v.typ))
//
//
//
// Despite this, people will warn against using it, and it is worth understanding
// their point of view before making your decision on weather it does properly
// fit your use case.
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
