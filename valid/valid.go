package valid

import (
	"reflect"

	//"lib/uput/valid/input"
	"lib/uput/valid/str"
)

//
// Output Function
// >> Output functions will be in the given datatype validation subpackage

//
// Generic/Dynamic Input Function
func If(input interface{}) {
	switch inputValue := reflect.ValueOf(input); inputValue.Kind() {
	case reflect.String:

		validstr.If(inputValue.String())
	}
}

//
// Input Functions
func IfString(input string) validstr.StringInput {
	return validstr.If(input)
}

//func IfInt(input int) InputData {
//	return InputData{
//		DataType: reflect.Int,
//		IntData:  input,
//	}
//}
//func IfUInt(input uint) InputData {
//	return InputData{
//		DataType: reflect.Uint,
//		UintData: input,
//	}
//}
//func IfMap(input map[interface{}]interface{}) InputData {
//	// If Map Length 0 (IsEmpty)
//	if len(input) > 0 {
//		// In A Map Each Key/Value Set Is An Entry
//		value := reflect.ValueOf(input)
//		fmt.Println("Reflected.ValueOf().Kind() type using basic type check: %s", value.Kind())
//		firstKey := (value.MapKeys())[0]
//		// Prepare InputData Based On Key Type/Kind
//		switch firstKey.Kind() {
//		case reflect.Array, reflect.String:
//			// TODO: Add stringErrorMessages for [key validations]
//		}
//
//		// Prepare InputData Based On Value Type/Kind
//		//switch input[firstKey].Kind() {
//		//case reflect.Array, reflect.String:
//		//	// TODO: Add stringErrorMessages for [value validations]
//		//}
//	} else {
//		// Empty/Nil Map will likely fail most validaitons
//	}
//	return InputData{
//		DataType: reflect.Map,
//		MapData:  input,
//	}
//}

//
// Making Errors?
///////////////////////////////////////
// Then follow the established convention of the message:
// 			"reflect: call of " + e.Method + " on " + e.Kind.String() + " Value"

//
// Reflect Notes:
//
// [deep equals]
// People saying don't use reflect pkg? Okay well you better implement deep equals
// that exists in this lib if you are not going to import it. Otherwise you
// can not do secure comparisons.
//
// 			func DeepEqual(x, y interface{}) bool {
// 				if x == nil || y == nil {
// 					return x == y
// 				}
// 				v1 := ValueOf(x)
// 				v2 := ValueOf(y)
// 				if v1.Type() != v2.Type() {
// 					return false
// 				}
// 				return deepValueEqual(v1, v2, make(map[visit]bool), 0)
// 			}
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
// ---or---
//
// Header reflect data used is not stable!
//
// [StringHeader] is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
//
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
//
// 			type StringHeader struct {
// 				Data uintptr
// 				Len  int
// 			}
//
// [SliceHeader] is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
//
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
//
// 		type SliceHeader struct {
// 			Data uintptr
//			Len  int
//			Cap  int
// 		}
//
// So people relying on Cap() for their size?
// [!][Slice is always bigger than a word; assume flagIndir.]
//
// 		return (*sliceHeader)(v.ptr).Cap
//
//  Element flag same as Elem of Ptr.
// 	Addressable, indirect, possibly read-only.
//		s := (*sliceHeader)(v.ptr)
//		if uint(i) >= uint(s.Len) {
//			panic("reflect: slice index out of range")
//		}
//
// 		val := arrayAt(s.Data, i, typ.size)
//
//  If CanAddr returns false, calling Addr will panic.
// 		func (v Value) CanAddr() bool {
// 			return v.flag&flagAddr != 0
// 		}
//
// [!][INFO][Check if its even possible to interface?]
// CanInterface reports whether Interface can be used without panicking.
// Keep in mind that v Value is interface{}, so it could be any tyope.
//
// 			func (v Value) Type() Type
// 			func (v Value) CanInterface() bool
//
//  Other IMPORTANT validations inside of reflect:
//
// 			func Zero(typ Type) Value
//
//      func (v Value) Len() int
//      func (v Value) Cap() int
// 			func (v Value) IsNil() bool
// 			func (v Value) IsValid() bool
// 			func (v Value) Kind() Kind
// 			func (v Value) Addr() Value
// 			func (v Value) CanSet() bool
//
// 			func (v Value) NumMethod() int
// 			func (v Value) Method(i int) Value
// 			func (v Value) MethodByName(name string) Value
//
// For general [] type
//
//  		func (v Value) Index(i int) Value
//
// For Struct Type
//
//  		func (v Value) Field(i int) Value
// 			func (v Value) NumField() int
// 			func (v Value) FieldByIndex(index []int) Value
// 			func (v Value) FieldByName(name string) Value
// 			func (v Value) FieldByNameFunc(match func(string) bool) Value
//
// For Map Type
//
// 			func (v Value) Elem() Value
// 			func (v Value) MapIndex(key Value) Value
// 			func (v Value) MapKeys() []Value
//
// For Channel Type
//
// 			func (v Value) Recv() (x Value, ok bool)
// 			func (v Value) TrySend(x Value) bool
//
// IMPORTANT TYPE validation
//
// 			func (v Value) OverflowComplex(x complex128) bool
// 			func (v Value) OverflowFloat(x float64) bool
// 			func (v Value) OverflowInt(x int64) bool
// 			func (v Value) OverflowUint(x uint64) bool
//
// CONVERTING FROM REFLECT
//
//      func (v Value) Interface() (i interface{})
//
// 			func (v Value) String() string // DOES NOT PANIC :)
// 			func (v Value) Complex() complex128
// 			func (v Value) Float() float64
//     	func (v Value) Int() int64
// 			func (v Value) Uint() uint64 // will panic if not uint
//
// 			func (v Value) Pointer() uintptr  // Does not require "unsafe" pkg
//
// Notice we use StringHeader and SliceHeader in Len() function!
//
// 			case Slice:
// 				// Slice is bigger than a word; assume flagIndir.
// 				return (*sliceHeader)(v.ptr).Len
// 			case String:
// 				// String is bigger than a word; assume flagIndir.
// 				return (*stringHeader)(v.ptr).Len
//
// [!][INFO][So how do we get our own len/count/size ?]
// [ Either use recover, to prevent these panics from crashing the program,
// or reimplement the code without the panics so we save even more time not
// panicing and recovering.]
//
// var uint8Type = TypeOf(uint8(0)).(*rtype)
// Index returns v's i'th element.
// It panics if v's Kind is not Array, Slice, or String or i is out of range.
// 			func (v Value) Index(i int) Value {
// 				switch v.kind() {
// 				case Array:
// 					tt := (*arrayType)(unsafe.Pointer(v.typ))
// 					if uint(i) >= uint(tt.len) {
// 						panic("reflect: array index out of range")
// 					}
// 					typ := tt.elem
// 					offset := uintptr(i) * typ.size
// 					// Either flagIndir is set and v.ptr points at array,
// 					// or flagIndir is not set and v.ptr is the actual array data.
// 					// In the former case, we want v.ptr + offset.
// 					// In the latter case, we must be doing Index(0), so offset = 0,
// 					// so v.ptr + offset is still okay.
// 					val := unsafe.Pointer(uintptr(v.ptr) + offset)
// 					fl := v.flag&(flagRO|flagIndir|flagAddr) | flag(typ.Kind()) // bits same as overall array
// 					return Value{typ, val, fl}
// 				case Slice:
// 					// Element flag same as Elem of Ptr.
// 					// Addressable, indirect, possibly read-only.
// 					s := (*sliceHeader)(v.ptr)
// 					if uint(i) >= uint(s.Len) {
// 						panic("reflect: slice index out of range")
// 					}
// 					tt := (*sliceType)(unsafe.Pointer(v.typ))
// 					typ := tt.elem
// 					val := arrayAt(s.Data, i, typ.size)
// 					fl := flagAddr | flagIndir | v.flag&flagRO | flag(typ.Kind())
// 					return Value{typ, val, fl}
// 				case String:
// 					s := (*stringHeader)(v.ptr)
// 					if uint(i) >= uint(s.Len) {
// 						panic("reflect: string index out of range")
// 					}
// 					p := arrayAt(s.Data, i, 1)
// 					fl := v.flag&flagRO | flag(Uint8) | flagIndir
// 					return Value{uint8Type, p, fl}
// 				}
// 				panic(&ValueError{"reflect.Value.Index", v.kind()})
// 			}
//
// /!\ PAY ATTENTION WHEN [StringHeader OR SliceHeader] is used, \
// keep in mind the above!
///////////////////////////////////////////////////////////////////////

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
