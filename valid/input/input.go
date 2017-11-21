package validinput

import (
	"errors"
	"reflect"
)

type ValidationText struct {
	Error       string
	Description string
}

type TextType int

const (
	ErrorText TextType = iota
	DescriptionText
)

type InputData struct {
	DataType       reflect.Kind
	DataTypeName   string
	Data           interface{}
	Validations    []Validation
	ValidationText map[string]*ValidationText
}

//
// Input Data Helpers
//==================================================================
func (input InputData) IsValid() bool {
	return (len(input.InputErrors()) == 0)
}

//
// InputErrors
//==================================================================
func (input InputData) InputErrors() (inputErrors []Validation) {
	for _, validation := range input.Validations {
		if !validation.IsValid {
			inputErrors = append(inputErrors, validation)
		}
	}
	return inputErrors
}
func (input InputData) Errors() (outputErrors []error) {
	for _, inputError := range input.InputErrors() {
		outputErrors = append(outputErrors, errors.New((inputError.Error())))
	}
	return outputErrors
}
func (input InputData) ErrorMessages() (errorMessages []string) {
	for _, inputError := range input.InputErrors() {
		outputErrors = append(outputErrors, inputError.Error())
	}
	return outputErrors
}
func (input InputData) ErrorCount() int {
	return len(input.InputErrors())
}

//
// Append Validations/Errors
//==================================================================
func (input InputData) AppendValidation(key string, values []string, isValid bool) InputData {
	if IsTextKeyValid(key) {
		input.Validations = append(input.Validations, Validation{
			DataType: input.DataTypeName,
			Key:      key,
			Text:     input.ValidationText[key],
			Values:   values,
			IsValid:  isValid,
		})
	}
	return input
}

//
// Localize Validation Descriptions
//==================================================================
// Update Last Added Validation/Error Text
func (input InputData) SetLastValidationText(text ValidationText) map[string]*ValidationText {
	if len(input.Validations) > 0 {
		textKey := input.Validations[len(input.Validations)-1].Key
		lastValidationText, exists := input.ValidationText[textKey]
		if exists {
			if IsTextContentValid(text.Error) {
				lastValidationText.Error = text.Error
			}
			if IsTextContentValid(text.Description) {
				lastValidationText.Description = text.Description
			}
			input.ValidationText[textKey] = lastValidationText
		}
	}
	return input.ValidationText
}

// Localize Descriptions and Error Messages
func (input InputData) SetValidationText(key string, text ValidationText) map[string]*ValidationText {
	if IsTextKeyValid(key) {
		// Validate: Assign if supplied validationText content valid
		if IsTextContentValid(text.Error) {
			input.ValidationText[key].Error = text.Error
		}
		if IsTextContentValid(text.Description) {
			input.ValidationText[key].Description = text.Description
		}
	}
	return input.ValidationText
}
func (input InputData) SetAllTextOfType(textType string, textMap map[string]string) map[string]*ValidationText {
	for key, text := range textMap {
		validationText := ValidationText{}
		if IsTextKeyValid(key) {
			if IsTextContentValid(text) {
				if textType == ErrorText || textType == DescriptionText {
					err := SetField(&validationText, textType, text)
					if err == nil {
						input.ValidationText = input.SetValidationText(key, validationText)
					}
				}
			}
		}
	}
	return input.ValidationText
}
