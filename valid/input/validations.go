package validinput

import (
	"reflect"
	"strings"
	"unicode"
)

type Key int

type Text struct {
	Error       string
	Description string
}

func (input InputData) GetText(key Key) (text Text) {
	text, exists := input.ValidationText[key]
	if exists {
		return text.Error, text.Description
	}
	return text
}

// TODO: Determine if []interface{} would be more fluid for Values field
type Validation struct {
	Key
	Kind    reflect.Kind
	Values  []string
	IsValid bool
	Text
}

//
// Validation/Error Text (string) Validations
//==================================================================
func IsKeyValid(key string) bool {
	// valid.IfValidationKey.IsBetween(2, 12)
	if !(2 <= len(key) && len(key) <= 12) {
		return false
	} else {
		for _, c := range key {
			// valid.IfValidationKey.IsAlphanumeric.NoSpaces
			if (!unicode.IsLetter(c) && !unicode.IsNumber(c)) || unicode.IsSpace(c) {
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
	return v.output(v.ErrorMessage)
}
func (v Validation) String() string {
	return v.ouput(v.Description)
}
