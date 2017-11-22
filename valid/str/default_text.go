package validstr

import (
	validinput "lib/uput/valid/input"
)

func GetDefeaultValidationTextString(key validinput.ValidationKey) validinput.ValidationText {
	return (DefaultStringValidationText())[key]
}

func DefaultStringValidationText() map[validinput.ValidationKey]validinput.ValidationText {
	return map[validinput.ValidationKey]validinput.ValidationText{
		In: validinput.ValidationText{
			Error:       "not included in",
			Description: "is included in",
		},
		NotIn: validinput.ValidationText{
			Error:       "included in",
			Description: "not included in",
		},
		Required: validinput.ValidationText{
			Error:       "is not present",
			Description: "is requred",
		},
		Empty: validinput.ValidationText{
			Error:       "is not empty",
			Description: "is empty",
		},
		NotEmpty: validinput.ValidationText{
			Error:       "is empty",
			Description: "is not empty",
		},
		Blank: validinput.ValidationText{
			Error:       "is not blank",
			Description: "is blank",
		},
		NotBlank: validinput.ValidationText{
			Error:       "is blank",
			Description: "is not blank",
		},
		Between: validinput.ValidationText{
			Error:       "length not between",
			Description: "length is between",
		},
		LessThan: validinput.ValidationText{
			Error:       "length not less than",
			Description: "length less than",
		},
		GreaterThan: validinput.ValidationText{
			Error:       "length not greater than",
			Description: "length greater than",
		},
		Contains: validinput.ValidationText{
			Error:       "not containing substring",
			Description: "contains substring",
		},
		NotContaining: validinput.ValidationText{
			Error:       "contains substring",
			Description: "not containing substring",
		},
		RegexMatch: validinput.ValidationText{
			Error:       "has no regex matches",
			Description: "matches regex pattern",
		},
		NoRegexMatch: validinput.ValidationText{
			Error:       "matches regex pattern",
			Description: "has no regex matches",
		},
		UTF8: validinput.ValidationText{
			Error:       "only UTF8 characters",
			Description: "contains no UTF8 characters",
		},
		NoUTF8: validinput.ValidationText{
			Error:       "only UTF8 characters",
			Description: "contains no UTF8 characters",
		},
		Uppercase: validinput.ValidationText{
			Error:       "contains lowercase characters",
			Description: "only uppercase characters",
		},
		NoUppercase: validinput.ValidationText{
			Error:       "has lowercase characters",
			Description: "only lowercase characters",
		},
		MinUppercase: validinput.ValidationText{
			Error:       "below minimum uppercase character count",
			Description: "minimum uppercase character count",
		},
		Lowercase: validinput.ValidationText{
			Error:       "has uppercase characters",
			Description: "only lowercase characters",
		},
		NoLowercase: validinput.ValidationText{
			Error:       "has lowercase characters",
			Description: "only uppercase characters",
		},
		MinLowercase: validinput.ValidationText{
			Error:       "below minimum lowercase character count",
			Description: "minimum lowercase character count",
		},
		Printable: validinput.ValidationText{
			Error:       "has non-printable characters",
			Description: "only printable characters",
		},
		NoPrintable: validinput.ValidationText{
			Error:       "has printable characters",
			Description: "only non-printable characters",
		},
		Alphabetic: validinput.ValidationText{
			Error:       "has non-alphabetic characters",
			Description: "only alphabetic characters",
		},
		NoAlphabetic: validinput.ValidationText{
			Error:       "has alphabetic characters",
			Description: "only non-alphabetic characters",
		},
		MinAlphabetic: validinput.ValidationText{
			Error:       "below minimum alphabetic character count",
			Description: "minimum alphabetic character count",
		},
		Alphanumeric: validinput.ValidationText{
			Error:       "has non-alphanumeric characters",
			Description: "only alphanumeric characters",
		},
		NoAlphanumeric: validinput.ValidationText{
			Error:       "has alphanumeric characters",
			Description: "only non-alphanumeric characters",
		},
		MinAlphanumeric: validinput.ValidationText{
			Error:       "below minimum alphanumeric character count",
			Description: "minimum alphanumeric character count",
		},
		Numeric: validinput.ValidationText{
			Error:       "has non-numeric characters",
			Description: "only numeric characters",
		},
		NoNumeric: validinput.ValidationText{
			Error:       "has numeric characters",
			Description: "only non-numeric characters",
		},
		MinNumeric: validinput.ValidationText{
			Error:       "below minimum numeric character count",
			Description: "minimum numeric character count",
		},
		Digits: validinput.ValidationText{
			Error:       "has non-digits",
			Description: "only digit characters",
		},
		NoDigits: validinput.ValidationText{
			Error:       "has digit characters",
			Description: "only non-digits",
		},
		MinDigits: validinput.ValidationText{
			Error:       "below minimum digit character count",
			Description: "minimum digit character count",
		},
		Symbols: validinput.ValidationText{
			Error:       "has non-symbol characters",
			Description: "only symbols",
		},
		NoSymbols: validinput.ValidationText{
			Error:       "has symbols",
			Description: "only non-symbol characters",
		},
		MinSymbols: validinput.ValidationText{
			Error:       "below minimum symbol character count",
			Description: "minimum symbol character count",
		},
		Punctuation: validinput.ValidationText{
			Error:       "has non-punctuation characters",
			Description: "only punctuation",
		},
		NoPunctuation: validinput.ValidationText{
			Error:       "has punctuation characters",
			Description: "only non-punctuation characters",
		},
		MinPunctuation: validinput.ValidationText{
			Error:       "below minimum punctuation character count",
			Description: "minimum punctuation character count",
		},
		Marks: validinput.ValidationText{
			Error:       "has non-mark UTF8 characters",
			Description: "only UTF8 mark characters",
		},
		NoMarks: validinput.ValidationText{
			Error:       "has UTF8 mark characters",
			Description: "only non-mark UTF8 characters",
		},
		Graphics: validinput.ValidationText{
			Error:       "has non-graphic UTF8 characters",
			Description: "only UTF8 graphic characters",
		},
		NoGraphics: validinput.ValidationText{
			Error:       "has graphic characters",
			Description: "only non-graphic UTF8 characters",
		},
		Controls: validinput.ValidationText{
			Error:       "has non-control UTF8 characters",
			Description: "only UTF8 control characters",
		},
		NoControls: validinput.ValidationText{
			Error:       "has UTF8 control characters",
			Description: "only non-control UTF8 characters",
		},
		Spaces: validinput.ValidationText{
			Error:       "has non-whitespace characters",
			Description: "only whitespace characters",
		},
		NoSpaces: validinput.ValidationText{
			Error:       "has whitespace characters",
			Description: "only non-whitespace characters",
		},
	}
}
