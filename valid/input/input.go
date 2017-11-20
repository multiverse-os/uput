package validinput

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type ValidationText struct {
	Error       string
	Description string
}

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
// Update Last Added Validation/Error Text
//==================================================================
func (input InputData) UpdateLastValidationText(text ValidationText) map[string]*ValidationText {
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

//
// Localize Validation Descriptions
//==================================================================
func (input InputData) SetValidationText(key string, text ValidationText) map[string]*ValidationText {
	if IsTextKeyValid(key) {
		currentText, exists := input.ValidationText[key]
		if !exists {
			currentText = &ValidationText{}
		}
		// Assign if supplied validationText values are valid
		if IsTextContentValid(text.Error) {
			currentText.Error = text.Error
		}
		if IsTextContentValid(text.Description) {
			currentText.Description = text.Description
		}
		// Confirm both Error and Description are assigned
		if len(currentText.Error) > 0 && len(currentText.Description) > 0 {
			input.ValidationText[key] = currentText
		}
	}
	return input.ValidationText
}
func (input InputData) SetAllTextOfType(textType string, textMap map[string]string) map[string]*ValidationText {
	for key, text := range textMap {
		validationText := ValidationText{}
		if IsTextKeyValid(key) {
			if IsTextContentValid(text) {
				var err error
				if textType == "Error" || textType == "Description" {
					err = SetField(&validationText, textType, text)
				}
				if err != nil {
					fmt.Println("[DEV] Error using SetField() in struct.go:,", err)
				} else {
					fmt.Println("validationText map UPDATED: Error:", validationText.Error, ", Validation:", validationText.Description)
					input.ValidationText = input.SetValidationText(key, validationText)
				}
			}
		}
	}
	return input.ValidationText
}

//
// Development Printing (remove later, don't assume logging style)
//==================================================================
func (input InputData) PrintValidations() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(input.Validations) > 0 {
		fmt.Println("{")
		fmt.Println("  \"validation_count\": " + strconv.Itoa(len(input.Validations)) + ",")
		fmt.Println("  \"validations\": {")
		for _, v := range input.Validations {
			fmt.Println("    \"" + v.String() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}

//
// Development Printing (remove later, don't assume logging style)
//==================================================================
func (input InputData) PrintErrors() {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(input.InputErrors()) > 0 {
		fmt.Println("{")
		fmt.Println("  \"error_count\": " + strconv.Itoa(len(input.InputErrors())) + ",")
		fmt.Println("  \"errors\": {")
		for _, err := range input.InputErrors() {
			fmt.Println("    \"" + err.Error() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}
