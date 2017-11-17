package validstr

import (
	"errors"
	"strconv"

	valid "lib/uput/valid"
)

var stringErrorMessages = map[string]string{
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

//type StringInput valid.InputData
type StringInput valid.InputData

//
// Output Function
func (input StringInput) IsValid() (bool, interface{}, []error) {
	return (len(input.Errors) == 0), input.StringData, input.Errors
}

//type StringInput struct {
//	inputData valid.InputData
//}

func If(input string) StringInput {
	return StringInput{
		DataType:      valid.StringType,
		StringData:    input,
		ErrorMessages: stringErrorMessages,
	}
}

//
// Option in slice of strings
func (input StringInput) IsIn(listOptions []string) StringInput {
	if !IsInSlice(input.StringData, listOptions) {
		//return input.AppendError("isin", strings.Join(listOptions, ", "))
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["isin"]))
	}
	return input
}
func (input StringInput) NotIn(listOptions []string) StringInput {
	if NotInSlice(input.StringData, listOptions) {
		//return input.AppendError("notin", strings.Join(listOptions, ", "))
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["isin"]))
	}
	return input
}

//
// String Length Validations
func (input StringInput) Required() StringInput {
	if !Required(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["empty"]))
	}
	return input
}
func (input StringInput) IsEmpty() StringInput {
	if !IsEmpty(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["empty"]))
	}
	return input
}
func (input StringInput) IsNotEmpty() StringInput {
	if !IsNotEmpty(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["notempty"]))
	}
	return input
}
func (input StringInput) IsBetween(start, end int) StringInput {
	if !IsBetween(input.StringData, start, end) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["between"]+": "+strconv.Itoa(start)+"-"+strconv.Itoa(end)))
	}
	return input
}
func (input StringInput) IsLessThan(lt int) StringInput {
	if !IsLessThan(input.StringData, lt) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["lessthan"]+": "+strconv.Itoa(lt)))
	}
	return input
}
func (input StringInput) IsGreaterThan(gt int) StringInput {
	if !IsGreaterThan(input.StringData, gt) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["greaterthan"]+": "+strconv.Itoa(gt)))
	}
	return input
}

//
// Substring Validation
// WARNING: DOES NOT WORK FOR UTF8 MATCHING
func (input StringInput) Contains(ss string) StringInput {
	if !IsContaining(input.StringData, ss) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["containing"]+": '"+ss+"'"))
	}
	return input
}
func (input StringInput) NotContaining(ss string) StringInput {
	if !NotContaining(input.StringData, ss) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["notcontaining"]+": '"+ss+"'"))
	}
	return input
}

//
// Regex Validation
func (input StringInput) IsRegexMatch(pattern string) StringInput {
	if !IsRegexMatch(input.StringData, pattern) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["regexmatch"]))
	}
	return input
}
func (input StringInput) NoRegexMatch(pattern string) StringInput {
	if !NoRegexMatch(input.StringData, pattern) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["noregexmatch"]))
	}
	return input
}

//
// UTF8 Rune Validation
func (input StringInput) IsUTF8() StringInput {
	if !IsUTF8(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["utf8"]))
	}
	return input
}
func (input StringInput) NoUTF8() StringInput {
	if !NoUTF8(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["noutf8"]))
	}
	return input
}
func (input StringInput) IsUppercase() StringInput {
	if !IsUppercase(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["uppercase"]))
	}
	return input
}
func (input StringInput) NoUppercase() StringInput {
	if !NoUppercase(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nouppercase"]))
	}
	return input
}
func (input StringInput) IsLowercase() StringInput {
	if !IsLowercase(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["lowercase"]))
	}
	return input
}
func (input StringInput) NoLowercase() StringInput {
	if !NoLowercase(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nolowercase"]))
	}
	return input
}
func (input StringInput) IsPrintable() StringInput {
	if !IsPrintable(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["printable"]))
	}
	return input
}
func (input StringInput) NoPrintable() StringInput {
	if !NoPrintable(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["noprintable"]))
	}
	return input
}
func (input StringInput) IsAlphabetic() StringInput {
	if !IsAlphabetic(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["alphabetic"]))
	}
	return input
}
func (input StringInput) NoAlphabetic() StringInput {
	if !NoAlphabetic(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["noalphabetic"]))
	}
	return input
}
func (input StringInput) IsAlphaNumeric() StringInput {
	if !IsAlphaNumeric(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["alphanumeric"]))
	}
	return input
}
func (input StringInput) NoAlphaNumeric() StringInput {
	if !NoAlphaNumeric(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["noalphanumeric"]))
	}
	return input
}
func (input StringInput) IsNumeric() StringInput {
	if !IsNumeric(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["numeric"]))
	}
	return input
}
func (input StringInput) NoNumeric() StringInput {
	if !NoNumeric(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nonumeric"]))
	}
	return input
}
func (input StringInput) IsDigits() StringInput {
	if !IsDigits(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["digits"]))
	}
	return input
}
func (input StringInput) NoDigits() StringInput {
	if !NoDigits(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nodigits"]))
	}
	return input
}
func (input StringInput) IsPunctuation() StringInput {
	if !IsPunctuation(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["punctuation"]))
	}
	return input
}
func (input StringInput) NoPunctuation() StringInput {
	if !NoPunctuation(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nopunctuation"]))
	}
	return input
}
func (input StringInput) IsSymbols() StringInput {
	if !IsSymbols(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["symbols"]))
	}
	return input
}
func (input StringInput) NoSymbols() StringInput {
	if !NoSymbols(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nosymbols"]))
	}
	return input
}
func (input StringInput) IsMarkCharacters() StringInput {
	if !IsMarkCharacters(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["markchars"]))
	}
	return input
}
func (input StringInput) NoMarkCharacters() StringInput {
	if !NoMarkCharacters(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nomarkchars"]))
	}
	return input
}
func (input StringInput) IsWhitespaces() StringInput {
	if !IsWhitespaces(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["spaces"]))
	}
	return input
}
func (input StringInput) NoWhitespaces() StringInput {
	if !NoWhitespaces(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nospaces"]))
	}
	return input
}
func (input StringInput) IsControlCharacters() StringInput {
	if !IsControlCharacters(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["controlchars"]))
	}
	return input
}
func (input StringInput) NoControlCharacters() StringInput {
	if !NoControlCharacters(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nocontrolchars"]))
	}
	return input
}
func (input StringInput) IsGraphicCharacters() StringInput {
	if !IsGraphicCharacters(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["graphicchars"]))
	}
	return input
}
func (input StringInput) NoGraphicCharacters() StringInput {
	if !NoGraphicCharacters(input.StringData) {
		input.Errors = append(input.Errors, errors.New(input.ErrorMessages["nographicchars"]))
	}
	return input
}
