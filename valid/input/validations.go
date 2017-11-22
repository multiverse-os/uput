package validinput

import (
	"reflect"
	"strings"
	"unicode"
)

type ValidationKey int
type ValidationText struct {
	Error       string
	Description string
}

//
// Global Loaded Localized Validation Text
//==================================================================
var GlobalLocalizedText map[ValidationKey]ValidationText

//var TextKeys map[string]ValidationKey

func InitializeLocalizedText() {
	if GlobalLocalizedText == nil {
		GlobalLocalizedText = make(map[ValidationKey]ValidationText)
	}
}

// Load From map[ValidationKey]ValidationText form used in DefaultText maps
func LoadGlobalLocalizedText(textMap map[ValidationKey]ValidationText) (loadCount int) {
	InitializeLocalizedText()
	for key, text := range textMap {
		globalText, exists := GlobalLocalizedText[key]
		if !exists {
			globalText = ValidationText{}
		}
		if IsTextValid(text.Description) {
			globalText.Description = text.Description
		}
		if IsTextValid(text.Error) {
			globalText.Error = text.Error
		}
		if IsTextValidOrEmpty(globalText.Description) && IsTextValidOrEmpty(globalText.Error) {
			GlobalLocalizedText[key] = globalText
			loadCount++
		}
	}
	return loadCount
}

//
// Individual InputData Validation Management
//==================================================================
func (input InputData) GetValidation(key ValidationKey) (ValidationText, int, bool) {
	for index, validation := range input.Validations {
		if validation.Key == key {
			return validation.Text, index, true
		}
	}
	return ValidationText{}, 0, false
}

// TODO: Determine if []interface{} would be more fluid for Values field
type Validation struct {
	Kind    reflect.Kind
	Key     ValidationKey
	Values  []string
	Text    ValidationText
	IsValid bool
}

//
// Validation/Error Text (string) Validations
//==================================================================
func IsTextValid(text string) bool {
	// valid.IfValidationText.Content.IsBetween(2, 64)
	if !(2 <= len(text) && len(text) <= 64) {
		return false
	} else {
		// valid.IfValidationText.Content.IsPrintable
		// TODO: Should be iterating over runes? Think this may not be assuming UTF8 runes
		for _, r := range text {
			if !unicode.IsPrint(r) {
				return false
			}
		}
	}
	return true
}

func IsTextValidOrEmpty(text string) bool {
	return (IsTextValid(text) || len(text) == 0)
}

//
// Compile Output Message
//==================================================================
func (v Validation) output(text string) string {
	switch len(v.Values) {
	case 0:
		return v.Kind.String() + ": " + text
	case 1:
		return v.Kind.String() + ": " + text + ": " + v.Values[0]
	case 2:
		return v.Kind.String() + ": " + text + ": " + v.Values[0] + " - " + v.Values[1]
	default:
		return v.Kind.String() + ": " + text + ": [ " + strings.Join(v.Values, ", ") + " ]"
	}
}

//
// Output Strings for Validation and Error
//==================================================================
func (v Validation) Error() string {
	return v.output(v.Text.Error)
}
func (v Validation) String() string {
	return v.output(v.Text.Description)
}
