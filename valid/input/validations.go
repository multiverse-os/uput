package validinput

import (
	"strings"
	"unicode"
)

type Validation struct {
	IsValid      bool
	DataType     string
	Text         *ValidationText
	Key          string
	Values       []string
	ValidateFunc interface{}
}

//
// Validation/Error Text (string) Validations
//==================================================================
func IsTextKeyValid(key string) bool {
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
func (v Validation) compileMessage(message string) string {
	switch len(v.Values) {
	case 0:
		return v.DataType + ": " + message
	case 1:
		return v.DataType + ": " + message + ": " + v.Values[0]
	case 2:
		return v.DataType + ": " + message + ": " + v.Values[0] + " - " + v.Values[1]
	default:
		return v.DataType + ": " + message + ": [ " + strings.Join(v.Values, ", ") + " ]"
	}
}

//
// Output Strings for Validation and Error
//==================================================================
func (v Validation) Error() string {
	return v.compileMessage(v.Text.Error)
}
func (v Validation) String() string {
	return v.compileMessage(v.Text.Description)
}
