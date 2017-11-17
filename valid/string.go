package valid

import (
	"errors"
	"strconv"
	"strings"

	"lib/uput/valid/str"
)

var stringErrorMessages = map[string]string{
	// String In Slice
	"isin":  "not included in",
	"notin": "included in",
	// Length Validaitons
	"empty":       "is not empty",
	"notempty":    "is empty",
	"between":     "length not between",
	"lessthan":    "length not less than",
	"greaterthan": "length not greater than",
	// Substring Validations
	"containing":    "not containing substring",
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
	"digits":         "non-digits",
	"nodigits":       "contains digits",
	"marks":          "non-mark characters",
	"nomarks":        "contains mark characters",
	"graphics":       "non-graphic characters",
	"nographics":     "contains graphic characters",
	"symbols":        "non-symbol characters",
	"nosymbols":      "contains symbols characters",
	"spaces":         "non-whitespace characters",
	"nospaces":       "contains whitespace characters",
}

//
// Option in slice of strings
func (input InputData) IsIn(listOptions []string) InputData {
	if !validstr.IsInSlice(input.stringData, listOptions) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["isin"]+": ["+strings.Join(listOptions, ", ")+"]"))
	}
	return input
}

func (input InputData) NotIn(listOptions []string) InputData {
	if !validstr.NotInSlice(input.stringData, listOptions) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["notin"]+": ["+strings.Join(listOptions, ", ")+"]"))
	}
	return input
}

//
// String Length Validations
func (input InputData) IsEmpty() InputData {
	if !validstr.IsEmpty(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["empty"]))
	}
	return input
}

func (input InputData) IsNotEmpty() InputData {
	if !validstr.IsNotEmpty(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["notempty"]))
	}
	return input
}

func (input InputData) IsBetween(start, end int) InputData {
	if !validstr.IsBetween(input.stringData, start, end) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["between"]+": "+strconv.Itoa(start)+"-"+strconv.Itoa(end)))
	}
	return input
}

func (input InputData) IsLessThan(lt int) InputData {
	if !validstr.IsLessThan(input.stringData, lt) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["lessthan"]+": "+strconv.Itoa(lt)))
	}
	return input
}

func (input InputData) IsGreaterThan(gt int) InputData {
	if !validstr.IsGreaterThan(input.stringData, gt) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["greaterthan"]+": "+strconv.Itoa(gt)))
	}
	return input
}

//
// Substring Validation
func (input InputData) IsContaining(ss string) InputData {
	if !validstr.IsContaining(input.stringData, ss) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["containing"]+": '"+ss+"'"))
	}
	return input
}

func (input InputData) NotContaining(ss string) InputData {
	if !validstr.NotContaining(input.stringData, ss) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["notcontaining"]+": '"+ss+"'"))
	}
	return input
}

//
// Regex Validation
func (input InputData) IsRegexMatch(pattern string) InputData {
	if !validstr.IsRegexMatch(input.stringData, pattern) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["regexmatch"]))
	}
	return input
}

func (input InputData) NoRegexMatch(pattern string) InputData {
	if !validstr.NoRegexMatch(input.stringData, pattern) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["noregexmatch"]))
	}
	return input
}

//
// UTF8 Rune Validation
func (input InputData) IsUTF8() InputData {
	if !validstr.IsUTF8(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["utf8"]))
	}
	return input
}

func (input InputData) NoUTF8() InputData {
	if !validstr.NoUTF8(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["noutf8"]))
	}
	return input
}

func (input InputData) IsUppercase() InputData {
	if !validstr.IsUppercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["uppercase"]))
	}
	return input
}

func (input InputData) NoUppercase() InputData {
	if !validstr.NoUppercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nouppercase"]))
	}
	return input
}

func (input InputData) IsLowercase() InputData {
	if !validstr.IsLowercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["lowercase"]))
	}
	return input
}

func (input InputData) NoLowercase() InputData {
	if !validstr.NoLowercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nolowercase"]))
	}
	return input
}

func (input InputData) IsPrintable() InputData {
	if !validstr.IsPrintable(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["printable"]))
	}
	return input
}

func (input InputData) NoPrintable() InputData {
	if !validstr.NoPrintable(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["noprintable"]))
	}
	return input
}

func (input InputData) IsAlphabetic() InputData {
	if !validstr.IsAlphabetic(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["alphabetic"]))
	}
	return input
}

func (input InputData) NoAlphabetic() InputData {
	if !validstr.NoAlphabetic(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["noalphabetic"]))
	}
	return input
}

func (input InputData) IsAlphaNumeric() InputData {
	if !validstr.IsAlphaNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["alphanumeric"]))
	}
	return input
}

func (input InputData) NoAlphaNumeric() InputData {
	if !validstr.NoAlphaNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["noalphanumeric"]))
	}
	return input
}

func (input InputData) IsNumeric() InputData {
	if !validstr.IsNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["numeric"]))
	}
	return input
}

func (input InputData) NoNumeric() InputData {
	if !validstr.NoNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nonumeric"]))
	}
	return input
}

func (input InputData) IsDigits() InputData {
	if !validstr.IsDigits(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["digits"]))
	}
	return input
}

func (input InputData) NoDigits() InputData {
	if !validstr.NoDigits(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nodigits"]))
	}
	return input
}

func (input InputData) IsPunctuation() InputData {
	if !validstr.IsPunctuation(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["punctuation"]))
	}
	return input
}

func (input InputData) NoPunctuation() InputData {
	if !validstr.NoPunctuation(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nopunctuation"]))
	}
	return input
}

func (input InputData) IsSymbols() InputData {
	if !validstr.IsSymbols(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["symbols"]))
	}
	return input
}

func (input InputData) NoSymbols() InputData {
	if !validstr.NoSymbols(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nosymbols"]))
	}
	return input
}

func (input InputData) IsMarkCharacters() InputData {
	if !validstr.IsMarkCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["markchars"]))
	}
	return input
}

func (input InputData) NoMarkCharacters() InputData {
	if !validstr.NoMarkCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nomarkchars"]))
	}
	return input
}

func (input InputData) IsWhitespaces() InputData {
	if !validstr.IsWhitespaces(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["spaces"]))
	}
	return input
}

func (input InputData) NoWhitespaces() InputData {
	if !validstr.NoWhitespaces(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nospaces"]))
	}
	return input
}

func (input InputData) IsControlCharacters() InputData {
	if !validstr.IsControlCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["controlchars"]))
	}
	return input
}

func (input InputData) NoControlCharacters() InputData {
	if !validstr.NoControlCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nocontrolchars"]))
	}
	return input
}

func (input InputData) IsGraphicCharacters() InputData {
	if !validstr.IsGraphicCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["graphicchars"]))
	}
	return input
}

func (input InputData) NoGraphicCharacters() InputData {
	if !validstr.NoGraphicCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrorMessages["nographicchars"]))
	}
	return input
}
