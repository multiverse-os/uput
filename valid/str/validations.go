package validstr

import (
//
)

//
// String Validation Descriptions
//===============================================================
func DefaultValidationDescriptions() map[string]string {
	return map[string]string{
		// String In Slice
		"isin":  "is included in",
		"notin": "not included in",
		// Size/Memory Space Validations
		"oversized":    "is over minimum size",
		"notoversized": "is not over size",
		// Length Validaitons
		"empty":       "is empty",
		"notempty":    "is not empty",
		"between":     "length is between",
		"lessthan":    "length less than",
		"greaterthan": "length greater than",
		// Substring Validations
		"iscontaining":  "containing substring",
		"notcontaining": "not ccontaining substring",
		// Regex Validations
		"regexmatch":   "has regex matches",
		"noregexmatch": "no regex matches",
		// UTF Rune Validations
		"utf8":           "only utf8 characters",
		"noutf8":         "no utf8 characters",
		"uppercase":      "no lowercase characters",
		"nouppercase":    "no uppercase characters",
		"lowercase":      "contains uppercase characters",
		"nolowercase":    "no lowercase characters",
		"printable":      "only printable characters",
		"noprintable":    "no printable characters",
		"alphabetic":     "only alphabetic characters",
		"noalphabetic":   "no alphabetic characters",
		"alphanumeric":   "only alphanumeric characters",
		"noalphanumeric": "no alphanumeric characters",
		"numeric":        "only numeric characters",
		"nonumeric":      "no numeric characters",
		"digits":         "only digits",
		"nodigits":       "no digits",
		"symbols":        "only symbol characters",
		"nosymbols":      "no symbols characters",
		"punctuation":    "only punctuation characters",
		"nopunctuation":  "no punctuation characters",
		"marks":          "only mark characters",
		"nomarks":        "no mark characters",
		"graphics":       "only graphic characters",
		"nographics":     "no graphic characters",
		"spaces":         "only whitespace characters",
		"nospaces":       "no whitespace characters",
	}
}

func ValidationDescription(key string) string {
	return (DefaultValidationDescriptions())[key]
}
