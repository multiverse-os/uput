package valid

import (
	"errors"
	"strconv"
	"strings"

	"lib/uput/valid/str"
)

var stringErrMessages = map[string]string{
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
	input.validations++
	if !validstr.IsInSlice(input.stringData, listOptions) {
		return appendStringErrMessage("isin", "", ": ["+strings.Join(listOptions, ", ")+"]")
	}
}
func (input InputData) NotIn(listOptions []string) InputData {
	input.validations++
	if !validstr.NotInSlice(input.stringData, listOptions) {
		return appendStringErrMessage("notin", "", ": ["+strings.Join(listOptions, ", ")+"]")
	}
}

//
// String Length Validations
func (input InputData) Required() InputData {
	if !validstr.Required(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["empty"]))
	}
	return input
}
func (input InputData) IsEmpty() InputData {
	input.validations++
	if !validstr.IsEmpty(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["empty"]))
	}
	return input
}
func (input InputData) IsNotEmpty() InputData {
	input.validations++
	if !validstr.IsNotEmpty(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["notempty"]))
	}
	return input
}
func (input InputData) IsBetween(start, end int) InputData {
	input.validations++
	if !validstr.IsBetween(input.stringData, start, end) {
		input.errors = append(input.errors, errors.New(stringErrMessages["between"]+": "+strconv.Itoa(start)+"-"+strconv.Itoa(end)))
	}
	return input
}
func (input InputData) IsLessThan(lt int) InputData {
	input.validations++
	if !validstr.IsLessThan(input.stringData, lt) {
		input.errors = append(input.errors, errors.New(stringErrMessages["lessthan"]+": "+strconv.Itoa(lt)))
	}
	return input
}
func (input InputData) IsGreaterThan(gt int) InputData {
	input.validations++
	if !validstr.IsGreaterThan(input.stringData, gt) {
		input.errors = append(input.errors, errors.New(stringErrMessages["greaterthan"]+": "+strconv.Itoa(gt)))
	}
	return input
}

//
// Substring Validation
func (input InputData) IsContaining(ss string) InputData {
	input.validations++
	if !validstr.IsContaining(input.stringData, ss) {
		input.errors = append(input.errors, errors.New(stringErrMessages["containing"]+": '"+ss+"'"))
	}
	return input
}
func (input InputData) NotContaining(ss string) InputData {
	input.validations++
	if !validstr.NotContaining(input.stringData, ss) {
		input.errors = append(input.errors, errors.New(stringErrMessages["notcontaining"]+": '"+ss+"'"))
	}
	return input
}

//
// Regex Validation
func (input InputData) IsRegexMatch(pattern string) InputData {
	input.validations++
	if !validstr.IsRegexMatch(input.stringData, pattern) {
		input.errors = append(input.errors, errors.New(stringErrMessages["regexmatch"]))
	}
	return input
}
func (input InputData) NoRegexMatch(pattern string) InputData {
	input.validations++
	if !validstr.NoRegexMatch(input.stringData, pattern) {
		input.errors = append(input.errors, errors.New(stringErrMessages["noregexmatch"]))
	}
	return input
}

//
// UTF8 Rune Validation
func (input InputData) IsUTF8() InputData {
	input.validations++
	if !validstr.IsUTF8(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["utf8"]))
	}
	return input
}
func (input InputData) NoUTF8() InputData {
	input.validations++
	if !validstr.NoUTF8(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["noutf8"]))
	}
	return input
}
func (input InputData) IsUppercase() InputData {
	input.validations++
	if !validstr.IsUppercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["uppercase"]))
	}
	return input
}
func (input InputData) NoUppercase() InputData {
	input.validations++
	if !validstr.NoUppercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nouppercase"]))
	}
	return input
}
func (input InputData) IsLowercase() InputData {
	input.validations++
	if !validstr.IsLowercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["lowercase"]))
	}
	return input
}
func (input InputData) NoLowercase() InputData {
	input.validations++
	if !validstr.NoLowercase(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nolowercase"]))
	}
	return input
}
func (input InputData) IsPrintable() InputData {
	input.validations++
	if !validstr.IsPrintable(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["printable"]))
	}
	return input
}
func (input InputData) NoPrintable() InputData {
	input.validations++
	if !validstr.NoPrintable(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["noprintable"]))
	}
	return input
}
func (input InputData) IsAlphabetic() InputData {
	input.validations++
	if !validstr.IsAlphabetic(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["alphabetic"]))
	}
	return input
}
func (input InputData) NoAlphabetic() InputData {
	input.validations++
	if !validstr.NoAlphabetic(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["noalphabetic"]))
	}
	return input
}
func (input InputData) IsAlphaNumeric() InputData {
	input.validations++
	if !validstr.IsAlphaNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["alphanumeric"]))
	}
	return input
}
func (input InputData) NoAlphaNumeric() InputData {
	input.validations++
	if !validstr.NoAlphaNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["noalphanumeric"]))
	}
	return input
}
func (input InputData) IsNumeric() InputData {
	input.validations++
	if !validstr.IsNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["numeric"]))
	}
	return input
}
func (input InputData) NoNumeric() InputData {
	input.validations++
	if !validstr.NoNumeric(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nonumeric"]))
	}
	return input
}
func (input InputData) IsDigits() InputData {
	input.validations++
	if !validstr.IsDigits(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["digits"]))
	}
	return input
}
func (input InputData) NoDigits() InputData {
	input.validations++
	if !validstr.NoDigits(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nodigits"]))
	}
	return input
}
func (input InputData) IsPunctuation() InputData {
	input.validations++
	if !validstr.IsPunctuation(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["punctuation"]))
	}
	return input
}
func (input InputData) NoPunctuation() InputData {
	input.validations++
	if !validstr.NoPunctuation(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nopunctuation"]))
	}
	return input
}
func (input InputData) IsSymbols() InputData {
	input.validations++
	if !validstr.IsSymbols(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["symbols"]))
	}
	return input
}
func (input InputData) NoSymbols() InputData {
	input.validations++
	if !validstr.NoSymbols(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nosymbols"]))
	}
	return input
}
func (input InputData) IsMarkCharacters() InputData {
	input.validations++
	if !validstr.IsMarkCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["markchars"]))
	}
	return input
}
func (input InputData) NoMarkCharacters() InputData {
	input.validations++
	if !validstr.NoMarkCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nomarkchars"]))
	}
	return input
}
func (input InputData) IsWhitespaces() InputData {
	input.validations++
	if !validstr.IsWhitespaces(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["spaces"]))
	}
	return input
}
func (input InputData) NoWhitespaces() InputData {
	input.validations++
	if !validstr.NoWhitespaces(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nospaces"]))
	}
	return input
}
func (input InputData) IsControlCharacters() InputData {
	input.validations++
	if !validstr.IsControlCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["controlchars"]))
	}
	return input
}
func (input InputData) NoControlCharacters() InputData {
	input.validations++
	if !validstr.NoControlCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nocontrolchars"]))
	}
	return input
}
func (input InputData) IsGraphicCharacters() InputData {
	input.validations++
	if !validstr.IsGraphicCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["graphicchars"]))
	}
	return input
}
func (input InputData) NoGraphicCharacters() InputData {
	input.validations++
	if !validstr.NoGraphicCharacters(input.stringData) {
		input.errors = append(input.errors, errors.New(stringErrMessages["nographicchars"]))
	}
	return input
}
