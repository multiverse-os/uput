package validinput

import (
	"fmt"
	"strconv" // development package
	"strings"
	"unicode"
)

//
// Validation Descriptions
//==================================================================

//
// Transformations / Normalization

// Validation Description (string) Normalization
// https://blog.golang.org/normalization
func NormalizeValidationDescription(description string) string {
	for index, character := range description {
		// Replace alternative whitespace characters
		fmt.Println("i, c: ", index, character)
		if unicode.IsSpace(character) {
			//description[index] = " "
		}
	}
	return description
}

//
// Validation Key/Description (string) Validation
func IsValidationKeyValid(key string) bool {
	// valid.IfValidationKey.IsBetween(2, 12)
	if !(2 <= len(key) && len(key) <= 12) {
		return false
	} else {
		// valid.IfValidationKey.IsAlphanumeric
		for _, c := range key {
			if !unicode.IsLetter(c) && !unicode.IsNumber(c) || unicode.IsSpace(c) {
				return false
			}
		}
	}
	return true
}
func IsValidationDescriptionValid(description string) bool {
	// valid.IfValidationDescription.IsBetween(2, 64)
	if !(2 <= len(description) && len(description) <= 64) {
		return false
	} else {
		// valid.IfValidationKey.IsPrintable
		for _, c := range description {
			if !unicode.IsPrint(c) {
				return false
			}
		}
	}
	return true
}

//
// InputData []validationDescriptions Validation
func (input InputData) ValidateValidationDescription() InputData {
	for key, value := range input.ErrorMessages {
		// valid.IfErrorMessages.IsLessThan(255)
		if len(key) <= 255 {
			if !IsValidationKeyValid(key) {
				// If key is invalid, delete the errorMessage from map
			}
			if !IsValidationDescriptionValid(value) {
				// If message is invalid, replace message with key
			}
		}
	}
	return input
}

// Development Printing (remove later, don't assume logging style)
func (input InputData) PrintValidations() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(input.Validations) > 0 {
		fmt.Println("{")
		fmt.Println("  \"validation_count\": " + strconv.Itoa(len(input.Validations)) + ",")
		fmt.Println("  \"validations\": {")
		for _, v := range input.Validations {
			fmt.Println("    \"" + input.DataTypeName + "\": \"" + v + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}

func (input InputData) AppendValidation(key string, values []string) InputData {
	switch lenValues := len(values); lenValues {
	case 0:
		input.Validations = append(input.Validations, input.ValidationDescriptions[key])
	case 1:
		input.Validations = append(input.Validations, input.ValidationDescriptions[key]+": "+values[0])
	case 2:
		input.Validations = append(input.Validations, input.ValidationDescriptions[key]+": "+values[0]+" - "+values[1])
	default:
		input.Validations = append(input.Validations, input.ValidationDescriptions[key]+": [ "+strings.Join(values, ", ")+" ]")
	}
	return input
}
