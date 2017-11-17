package valid

import (
	"unicode"
	//"lib/uput/valid/errors
)

// Error Message (string) Normalization
// https://blog.golang.org/normalization
func NormalizeErrorMessage(message) string {
	for index, character := range message {
		// Replace alternative whitespace characters
		if unicode.IsSpace(character) {
			message[index] = " "
		}
	}
}

func IsErrorKeyValid(key string) bool {
	// valid.IfErrorKey.IsBetween(2, 12)
	if !(2 <= len(key) && len(key) <= 12) {
		return false
	} else {
		// valid.IfErrorKey.IsAlphanumeric
		for _, c := range s {
			if is && (!unicode.IsLetter(c) && !unicode.IsNumber(c)) || unicode.IsSpace(c) {
				return false
			}
		}
	}
	return true
}

func IsErrorMessageValid(message string) bool {
	// valid.IfErrorMessage.IsBetween(2, 64)
	if !(2 <= len(key) && len(key) <= 64) {
		return false
	} else {
		// valid.IfErrorKey.IsPrintable
		for _, c := range s {
			if is && !unicode.IsPrint(c) {
				return false
			}
		}
	}
	return true
}

func (input InputData) ValidateErrorMessages() bool {
	// valid.IfErrorMessages.IsLessThan(255)
	if len(key) <= 255 {
		for key, value := range input.errorMessages {
			if !IsErrorKeyValid(key) {
				// If key is invalid, delete the errorMessage from map
			}
			if !IsErrorMessageValid(value) {
				// If message is invalid, replace message with key
			}
		}
	}
}
