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
		"notcontaining": "containing substring",
		// Regex Validations
		"regexmatch":   "no regex matches",
		"noregexmatch": "has regex matches",
		// UTF Rune Validations
		"utf8":           "has no utf8 characters",
		"noutf8":         "has utf8 characters",
		"uppercase":      "has lowercase characters",
		"nouppercase":    "has non-uppercase characters",
		"lowercase":      "has uppercase characters",
		"nolowercase":    "has non-lowercase characters",
		"printable":      "has non-printable characters",
		"noprintable":    "has printable characters",
		"alphabetic":     "has non-alphabetic characters",
		"noalphabetic":   "has alphabetic characters",
		"alphanumeric":   "has non-alphanumeric characters",
		"noalphanumeric": "has alphanumeric characters",
		"numeric":        "has non-numeric characters",
		"nonumeric":      "has numeric characters",
		"digits":         "has non-digit characters",
		"nodigits":       "has digits characters",
		"symbols":        "has non-symbol characters",
		"nosymbols":      "has symbols characters",
		"punctuation":    "has non-punctuation characters",
		"nopunctuation":  "has punctuation characters",
		"marks":          "has non-mark characters",
		"nomarks":        "has mark characters",
		"graphics":       "has non-graphic characters",
		"nographics":     "has graphic characters",
		"spaces":         "has non-whitespace characters",
		"nospaces":       "has whitespace characters",
	}
}

func ErrorMessage(key string) string {
	return (DefaultErrorMessages())[key]
}
