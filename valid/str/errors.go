package validstr

import (
	validinput "lib/uput/valid/input"
)

func DefaultValidationText() map[string]*validinput.ValidationText {
	return map[string]*validinput.ValidationText{
		"isin": &validinput.ValidationText{
			Error:       "not included in",
			Description: "is included in",
		},
		"notin": &validinput.ValidationText{
			Error:       "included in",
			Description: "not included in",
		},
		"empty": &validinput.ValidationText{
			Error:       "is not empty",
			Description: "is empty",
		},
		"notempty": &validinput.ValidationText{
			Error:       "is empty",
			Description: "is not empty",
		},
		"between": &validinput.ValidationText{
			Error:       "length not between",
			Description: "length is between",
		},
		"lessthan": &validinput.ValidationText{
			Error:       "length not less than",
			Description: "length less than",
		},
		"greaterthan": &validinput.ValidationText{
			Error:       "length not greater than",
			Description: "length greater than",
		},
		"iscontaining": &validinput.ValidationText{
			Error:       "not containing substring",
			Description: "containing substring",
		},
		"notcontaining": &validinput.ValidationText{
			Error:       "containing substring",
			Description: "not containing substring",
		},
		"regexmatch": &validinput.ValidationText{
			Error:       "has no regex matches",
			Description: "matches regex pattern",
		},
		"noregexmatch": &validinput.ValidationText{
			Error:       "matches regex pattern",
			Description: "has no regex matches",
		},
		"utf8": &validinput.ValidationText{
			Error:       "only UTF8 characters",
			Description: "contains no UTF8 characters",
		},
		"noutf8": &validinput.ValidationText{
			Error:       "only UTF8 characters",
			Description: "contains no UTF8 characters",
		},
		"uppercase": &validinput.ValidationText{
			Error:       "contains lowercase characters",
			Description: "only uppercase characters",
		},
		"nouppercase": &validinput.ValidationText{
			Error:       "has lowercase characters",
			Description: "only lowercase characters",
		},
		"lowercase": &validinput.ValidationText{
			Error:       "has uppercase characters",
			Description: "only lowercase characters",
		},
		"nolowercase": &validinput.ValidationText{
			Error:       "has lowercase characters",
			Description: "only uppercase characters",
		},
		"printable": &validinput.ValidationText{
			Error:       "has non-printable characters",
			Description: "only printable characters",
		},
		"noprintable": &validinput.ValidationText{
			Error:       "has printable characters",
			Description: "only non-printable characters",
		},
		"alphabetic": &validinput.ValidationText{
			Error:       "has non-alphabetic characters",
			Description: "only alphabetic characters",
		},
		"noalphabetic": &validinput.ValidationText{
			Error:       "has alphabetic characters",
			Description: "only non-alphabetic characters",
		},
		"alphanumeric": &validinput.ValidationText{
			Error:       "has non-alphanumeric characters",
			Description: "only alphanumeric characters",
		},
		"noalphanumeric": &validinput.ValidationText{
			Error:       "has alphanumeric characters",
			Description: "only non-alphanumeric characters",
		},
		"numeric": &validinput.ValidationText{
			Error:       "has non-numeric characters",
			Description: "only numeric characters",
		},
		"nonumeric": &validinput.ValidationText{
			Error:       "has numeric characters",
			Description: "only non-numeric characters",
		},
		"digits": &validinput.ValidationText{
			Error:       "has non-digit characters",
			Description: "only digits",
		},
		"nodigits": &validinput.ValidationText{
			Error:       "has digits characters",
			Description: "only non-digits",
		},
		"symbols": &validinput.ValidationText{
			Error:       "has non-symbol characters",
			Description: "only symbols",
		},
		"nosymbols": &validinput.ValidationText{
			Error:       "has symbols characters",
			Description: "only non-symbol characters",
		},
		"punctuation": &validinput.ValidationText{
			Error:       "has non-punctuation characters",
			Description: "only punctuation",
		},
		"nopunctuation": &validinput.ValidationText{
			Error:       "has punctuation characters",
			Description: "only non-punctuation characters",
		},
		"marks": &validinput.ValidationText{
			Error:       "has non-mark characters",
			Description: "only UTF8 mark characters",
		},
		"nomarks": &validinput.ValidationText{
			Error:       "has mark characters",
			Description: "no UTF8 mark characters",
		},
		"graphics": &validinput.ValidationText{
			Error:       "has non-graphic characters",
			Description: "no UTF8 graphic characters",
		},
		"nographics": &validinput.ValidationText{
			Error:       "has graphic characters",
			Description: "only UTF8 graphic characters",
		},
		"spaces": &validinput.ValidationText{
			Error:       "has non-whitespace characters",
			Description: "only whitespace characters",
		},
		"nospaces": &validinput.ValidationText{
			Error:       "has whitespace characters",
			Description: "no whitespace characters",
		},
	}
}
