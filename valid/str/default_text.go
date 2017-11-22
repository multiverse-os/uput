package validstr

import (
	validinput "lib/uput/valid/input"
)

// Should this be in str/text to avoid loading any unncessary data for developers with
// their own localized strings?
func DefaultStringText() map[validinput.ValidationKey]validinput.Text {
	return map[string]ValidationText{
		In: ValidationText{
			Error:       "not included in",
			Description: "is included in",
		},
		NotIn: ValidationText{
			Error:       "included in",
			Description: "not included in",
		},
		Required: ValidationText{
			Error:       "is not present",
			Description: "is requred",
		},
		Empty: ValidationText{
			Error:       "is not empty",
			Description: "is empty",
		},
		NotEmpty: ValidationText{
			Error:       "is empty",
			Description: "is not empty",
		},
		Blank: ValidationText{
			Error:       "is not blank",
			Description: "is blank",
		},
		NotBlank: ValidationText{
			Error:       "is blank",
			Description: "is not blank",
		},
		Between: ValidationText{
			Error:       "length not between",
			Description: "length is between",
		},
		LessThan: ValidationText{
			Error:       "length not less than",
			Description: "length less than",
		},
		GreaterThan: ValidationText{
			Error:       "length not greater than",
			Description: "length greater than",
		},
		Contains: ValidationText{
			Error:       "not containing substring",
			Description: "contains substring",
		},
		NotContaining: ValidationText{
			Error:       "contains substring",
			Description: "not containing substring",
		},
		RegexMatch: ValidationText{
			Error:       "has no regex matches",
			Description: "matches regex pattern",
		},
		NoRegexMatch: ValidationText{
			Error:       "matches regex pattern",
			Description: "has no regex matches",
		},
		UTF8: ValidationText{
			Error:       "only UTF8 characters",
			Description: "contains no UTF8 characters",
		},
		NoUTF8: ValidationText{
			Error:       "only UTF8 characters",
			Description: "contains no UTF8 characters",
		},
		Uppercase: ValidationText{
			Error:       "contains lowercase characters",
			Description: "only uppercase characters",
		},
		NoUppercase: ValidationText{
			Error:       "has lowercase characters",
			Description: "only lowercase characters",
		},
		Lowercase: ValidationText{
			Error:       "has uppercase characters",
			Description: "only lowercase characters",
		},
		NoLowercase: ValidationText{
			Error:       "has lowercase characters",
			Description: "only uppercase characters",
		},
		Printable: ValidationText{
			Error:       "has non-printable characters",
			Description: "only printable characters",
		},
		NoPrintable: ValidationText{
			Error:       "has printable characters",
			Description: "only non-printable characters",
		},
		Alphabetic: ValidationText{
			Error:       "has non-alphabetic characters",
			Description: "only alphabetic characters",
		},
		NoAlphabetic: ValidationText{
			Error:       "has alphabetic characters",
			Description: "only non-alphabetic characters",
		},
		Alphanumeric: ValidationText{
			Error:       "has non-alphanumeric characters",
			Description: "only alphanumeric characters",
		},
		NoAlphanumeric: ValidationText{
			Error:       "has alphanumeric characters",
			Description: "only non-alphanumeric characters",
		},
		Numeric: ValidationText{
			Error:       "has non-numeric characters",
			Description: "only numeric characters",
		},
		NoNumeric: ValidationText{
			Error:       "has numeric characters",
			Description: "only non-numeric characters",
		},
		Digits: ValidationText{
			Error:       "has non-digit characters",
			Description: "only digits",
		},
		NoDigits: ValidationText{
			Error:       "has digits characters",
			Description: "only non-digits",
		},
		Symbols: ValidationText{
			Error:       "has non-symbol characters",
			Description: "only symbols",
		},
		NoSymbols: ValidationText{
			Error:       "has symbols characters",
			Description: "only non-symbol characters",
		},
		Punctuation: ValidationText{
			Error:       "has non-punctuation characters",
			Description: "only punctuation",
		},
		NoPunctuation: ValidationText{
			Error:       "has punctuation characters",
			Description: "only non-punctuation characters",
		},
		Marks: ValidationText{
			Error:       "has non-mark characters",
			Description: "only UTF8 mark characters",
		},
		NoMarks: ValidationText{
			Error:       "has mark characters",
			Description: "no UTF8 mark characters",
		},
		Graphics: ValidationText{
			Error:       "has non-graphic characters",
			Description: "no UTF8 graphic characters",
		},
		NoGraphics: ValidationText{
			Error:       "has graphic characters",
			Description: "only UTF8 graphic characters",
		},
		Controls: ValidationText{
			Error:       "has control characters",
			Description: "only UTF8 control characters",
		},
		NoControls: ValidationText{
			Error:       "has non-control characters",
			Description: "only UTF8 non-control characters",
		},
		Spaces: ValidationText{
			Error:       "has non-whitespace characters",
			Description: "only whitespace characters",
		},
		NoSpaces: ValidationText{
			Error:       "has whitespace characters",
			Description: "no whitespace characters",
		},
	}
}
