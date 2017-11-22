package validinput

import (
	"errors"
	"reflect"
)

type InputData struct {
	DataType    reflect.Kind
	Data        interface{}
	Validations []Validation
}

//
// Input Validation
//==================================================================
func NewInput(data interface{}) (input InputData) {
	input.DataType = reflect.TypeOf(data).Kind()
	if input.DataType != reflect.Invalid {
		input.Data = data
	}
	return input
}

//
// Input Data Helpers
//==================================================================
func (input InputData) IsValid() bool {
	return (len(input.InputErrors()) == 0)
}

//
// Validations
//==================================================================
func (input InputData) ValidationDescriptions() (descriptions []string) {
	for _, v := range input.Validations {
		descriptions = append(descriptions, v.Text.Description)
	}
	return descriptions
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
		errorMessages = append(errorMessages, inputError.Error())
	}
	return errorMessages
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
			DataType: input.DataType.String(),
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
//func (input InputData) SetValidationText(key string, text ValidationText) map[string]*ValidationText {
//	if IsTextKeyValid(key) {
//		// Validate: Assign if supplied validationText content valid
//		if IsTextContentValid(text.Error) {
//			input.ValidationText[key].Error = text.Error
//		}
//		if IsTextContentValid(text.Description) {
//			input.ValidationText[key].Description = text.Description
//		}
//	}
//	return input.ValidationText
//}
func (input InputData) SetAllErrorText(textMap map[string]string) map[string]*ValidationText {
	for key, text := range textMap {
		input.SetValidationText(key, ValidationText{Error: text})
	}
	return input.ValidationText
}
func (input InputData) SetAllText(textMap map[string]string) map[string]*ValidationText {
	for key, text := range textMap {
		input.SetValidationText(key, ValidationText{Description: text})
	}
	return input.ValidationText
}
