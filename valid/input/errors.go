package validinput

import (
	"errors"
	"fmt"
	"strconv" // development package
	"unicode"
)

//
// Error Tips
////////////////////////////////////////////////////////////////////
// In other stdlibs errors look like:
// 			"strings.Reader.UnreadByte: at beginning of string"
//
//
// An error is anything that can describe itself as an error string
// The idea is captured by the predefined, built-in interface type,
// error, with its single method, Error, returning a string:
//
// type error interface {
//     Error() string
// }
//
// type MyError struct {
//     When time.Time
//     What string
// }
//
// func (e *MyError) Error() string {
//     return fmt.Sprintf("at %v, %s",
//         e.When, e.What)
// }
//

//
// Transformations / Normalization

// Error Message (string) Normalization
// https://blog.golang.org/normalization
func NormalizeErrorMessage(message string) string {
	for index, character := range message {
		// Replace alternative whitespace characters
		fmt.Println("i, c: ", index, character)
		if unicode.IsSpace(character) {
			//message[index] = " "
		}
	}
	return message
}

//
// Error Key/Message (string) Validation
func IsErrorKeyValid(key string) bool {
	// valid.IfErrorKey.IsBetween(2, 12)
	if !(2 <= len(key) && len(key) <= 12) {
		return false
	} else {
		// valid.IfErrorKey.IsAlphanumeric
		for _, c := range key {
			if !unicode.IsLetter(c) && !unicode.IsNumber(c) || unicode.IsSpace(c) {
				return false
			}
		}
	}
	return true
}
func IsErrorMessageValid(message string) bool {
	// valid.IfErrorMessage.IsBetween(2, 64)
	if !(2 <= len(message) && len(message) <= 64) {
		return false
	} else {
		// valid.IfErrorKey.IsPrintable
		for _, c := range message {
			if !unicode.IsPrint(c) {
				return false
			}
		}
	}
	return true
}

//
// InputData []errorMessages Validation
func (input InputData) ValidateErrorMessages() InputData {
	for key, value := range input.ErrorMessages {
		// valid.IfErrorMessages.IsLessThan(255)
		if len(key) <= 255 {
			if !IsErrorKeyValid(key) {
				// If key is invalid, delete the errorMessage from map
			}
			if !IsErrorMessageValid(value) {
				// If message is invalid, replace message with key
			}
		}
	}
	return input
}

// Development Printing (remove later, don't assume logging style)
func (in InputData) PrintErrors() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(in.Errors) > 0 {
		fmt.Println("{")
		fmt.Println("  \"error_count\": \"" + strconv.Itoa(len(in.Errors)) + ",")
		fmt.Println("  \"errors\": {")
		for _, err := range in.Errors {
			fmt.Println("    \"string\": \"" + err.Error() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}

func (input InputData) AppendError(key string, values []string) InputData {
	switch lenValues := len(values); lenValues {
	case 1:
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages[key]+": "+values[0]))
	case 2:
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages[key]+": "+values[0]+" - "+values[1]))
	default:
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages[key]))
	}
	return input
}
