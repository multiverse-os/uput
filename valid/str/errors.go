package validstr

import (
//
)

func DefaultErrorMessages() map[string]string {
	return map[string]string{
		// String In Slice
		"isin":  "not included in",
		"notin": "included in",
		// Size/Memory Space Validations
		"oversized":    "is too large",
		"notoversized": "data not oversized",
		// Length Validaitons
		"empty":       "is not empty",
		"notempty":    "is empty",
		"between":     "length not between",
		"lessthan":    "length not less than",
		"greaterthan": "length not greater than",
		// Substring Validations
		"iscontaining":  "not containing substring",
		"notcontaining": "ccontaining substring",
		// Regex Validations
		"regexmatch":   "no regex matches",
		"noregexmatch": "has regex matches",
		// UTF Rune Validations
		"utf8":           "no utf8 characters",
		"noutf8":         "contains utf8 characters",
		"uppercase":      "contains lowercase characters",
		"nouppercase":    "non-uppercase characters",
		"lowercase":      "contains uppercase characters",
		"nolowercase":    "non-lowercase characters",
		"printable":      "non-printable characters",
		"noprintable":    "contains printable characters",
		"alphabetic":     "non-alphabetic characters",
		"noalphabetic":   "contains alphabetic characters",
		"alphanumeric":   "non-alphanumeric characters",
		"noalphanumeric": "contains alphanumeric characters",
		"numeric":        "non-numeric characters",
		"nonumeric":      "is numeric characters",
		"digits":         "contains non-digit characters",
		"nodigits":       "contains digits characters",
		"symbols":        "non-symbol characters",
		"nosymbols":      "contains symbols characters",
		"punctuation":    "contains non-punctuation characters",
		"nopunctuation":  "contains punctuation characters",
		"marks":          "non-mark characters",
		"nomarks":        "contains mark characters",
		"graphics":       "non-graphic characters",
		"nographics":     "contains graphic characters",
		"spaces":         "non-whitespace characters",
		"nospaces":       "contains whitespace characters",
	}
}

func ErrorMessage(key string) string {
	return (DefaultErrorMessages())[key]
}
