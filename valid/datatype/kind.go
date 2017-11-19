package datatype

import (
//
)

//
// Data Types (Kind/Type in the Reflect package)
///////////////////////////////////////////////////////////////////////////////////
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
// All of this was pulled from reflect, the idea is to
// use only the aspects of reflect useful to this library
// to reduce the amount of code.
//

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

var kindNames = []string{
	Invalid:       "invalid",
	Bool:          "bool",
	Int:           "int",
	Int8:          "int8",
	Int16:         "int16",
	Int32:         "int32",
	Int64:         "int64",
	Uint:          "uint",
	Uint8:         "uint8",
	Uint16:        "uint16",
	Uint32:        "uint32",
	Uint64:        "uint64",
	Uintptr:       "uintptr",
	Float32:       "float32",
	Float64:       "float64",
	Complex64:     "complex64",
	Complex128:    "complex128",
	Array:         "array",
	Chan:          "chan",
	Func:          "func",
	Interface:     "interface",
	Map:           "map",
	Ptr:           "ptr",
	Slice:         "slice",
	String:        "string",
	Struct:        "struct",
	UnsafePointer: "unsafe.Pointer",
}
