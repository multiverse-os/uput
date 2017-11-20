package validinput

import (
	"fmt"
	"reflect"
	"unicode"
)

type InputData struct {
	DataType       reflect.Kind
	DataTypeName   string
	Data           interface{}
	InputErrors    []InputError
	Validations    []Validation
	ValidationText map[string]ValidationText
}

type ValidationText struct {
	Error       string
	Description string
}

//
// Validation/Error Text (string) Validations
//==================================================================
func IsTextKeyValid(key string) bool {
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
func IsTextContentValid(content string) bool {
	// valid.IfValidationText.Content.IsBetween(2, 64)
	if !(2 <= len(content) && len(content) <= 64) {
		return false
	} else {
		// valid.IfValidationText.Content.IsPrintable
		for _, c := range content {
			if !unicode.IsPrint(c) {
				return false
			}
		}
	}
	return true
}
func (input InputData) ValidateValidationText() InputData {
	for key, value := range input.ValidationText {
		// valid.IfErrorMessages.IsLessThan(255)
		if len(key) <= 255 {
			if !IsTextKeyValid(key) {
				delete(input.ValidationText, key)
			}
			if !IsTextContentValid(value.Description) {
				value.Description = key
				input.ValidationText[key] = value
			}
			if !IsTextContentValid(value.Error) {
				value.Error = key
				input.ValidationText[key] = value
			}
		}
	}
	return input
}

//
// Localize Validation Descriptions
//==================================================================
func (input InputData) UpdateText(key string, text ValidationText) map[string]ValidationText {
	if IsTextKeyValid(key) {
		currentText, exists := input.ValidationText[key]
		if !exists {
			currentText = ValidationText{}
		}
		if len(text.Error) > 0 && !IsTextContentValid(text.Error) {
			currentText.Error = text.Error
		}
		if len(text.Description) > 0 && !IsTextContentValid(text.Description) {
			currentText.Description = text.Description
		}
		input.ValidationText[key] = text
	}
	return input.ValidationText
}
func (input InputData) UpdateValidationText(textType string, textMap map[string]string) map[string]ValidationText {
	for key, text := range textMap {
		validationText := ValidationText{}
		_ = SetField(&validationText, textType, text)
		fmt.Println("validationText: Error:", validationText.Error, ", Validation:", validationText.Description)
		input.ValidationText = input.UpdateText(key, validationText)
	}
	return input.ValidationText
}
