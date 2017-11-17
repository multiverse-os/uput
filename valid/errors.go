package valid

import (
	"fmt"
	"unicode"
	//"lib/uput/valid/errors
)

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
