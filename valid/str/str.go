package validstr

import (
	"reflect"
	"strconv"
	"strings"

	// TODO: Should avoid this and move any thing used into a core/common valid package
	// so that valid could be used for including all validation and this for a single 1
	//valid "lib/uput/valid"
	input "lib/uput/valid/input"
)

// Nesting the common inputData type
type StringInput struct {
	inputData input.InputData
}

//
// Validation Input Function
func If(si string) StringInput {
	// TODO: Right now we are using the entire reflect package jsut for
	// a const Int value, either manually enter the value or delegate this to valid
	return StringInput{
		inputData: input.InputData{
			DataType:      reflect.String,
			Data:          si,
			StringData:    si,
			ErrorMessages: (DefaultErrorMessages()),
		},
	}
}

//
// Validation Output Function
func (si StringInput) IsValid() (bool, interface{}, []error) {
	return (len(si.inputData.Errors) == 0), si.inputData.StringData, si.inputData.Errors
}

//
// Option in slice of strings
func (si StringInput) IsIn(listOptions []string) StringInput {
	if !IsInSlice(si.inputData.StringData, listOptions) {
		si.inputData = si.inputData.AppendError("isin", []string{strings.Join(listOptions, ", ")})
	}
	return si
}
func (si StringInput) NotIn(listOptions []string) StringInput {
	if NotInSlice(si.inputData.StringData, listOptions) {
		si.inputData = si.inputData.AppendError("isin", []string{strings.Join(listOptions, ", ")})
	}
	return si
}

//
// String Length Validations
func (si StringInput) Required() StringInput {
	if !Required(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("required", nil)
	}
	return si
}
func (si StringInput) IsEmpty() StringInput {
	if !IsEmpty(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("empty", nil)
	}
	return si
}
func (si StringInput) IsNotEmpty() StringInput {
	if !NotEmpty(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("notempty", nil)
	}
	return si
}
func (si StringInput) IsBetween(start, end int) StringInput {
	if !IsBetween(si.inputData.StringData, start, end) {
		si.inputData = si.inputData.AppendError("between", []string{strconv.Itoa(start), strconv.Itoa(end)})
	}
	return si
}
func (si StringInput) IsLessThan(lt int) StringInput {
	if !IsLessThan(si.inputData.StringData, lt) {
		si.inputData = si.inputData.AppendError("lessthan", []string{strconv.Itoa(lt)})
	}
	return si
}
func (si StringInput) IsGreaterThan(gt int) StringInput {
	if !IsGreaterThan(si.inputData.StringData, gt) {
		si.inputData = si.inputData.AppendError("greaterthan", []string{strconv.Itoa(gt)})
	}
	return si
}

//
// Substring Validation
// WARNING: DOES NOT WORK FOR UTF8 MATCHING
// This will let through look-alikes
func (si StringInput) IsContaining(ss string) StringInput {
	if !IsContaining(si.inputData.StringData, ss) {
		si.inputData = si.inputData.AppendError("containing", []string{ss})
	}
	return si
}
func (si StringInput) NotContaining(ss string) StringInput {
	if !NotContaining(si.inputData.StringData, ss) {
		si.inputData = si.inputData.AppendError("notcontaining", []string{ss})
	}
	return si
}

//
// Regex Validation
func (si StringInput) IsRegexMatch(pattern string) StringInput {
	if !IsRegexMatch(si.inputData.StringData, pattern) {
		si.inputData = si.inputData.AppendError("regexmatch", []string{pattern})
	}
	return si
}
func (si StringInput) NoRegexMatch(pattern string) StringInput {
	if !NoRegexMatch(si.inputData.StringData, pattern) {
		si.inputData = si.inputData.AppendError("noregexmatch", []string{pattern})
	}
	return si
}

//
// UTF8 Rune Validation
func (si StringInput) IsUTF8() StringInput {
	if !IsUTF8(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("utf8", nil)
	}
	return si
}
func (si StringInput) NoUTF8() StringInput {
	if !NoUTF8(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("noutf8", nil)
	}
	return si
}
func (si StringInput) IsUppercase() StringInput {
	if !IsUppercase(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("uppercase", nil)
	}
	return si
}
func (si StringInput) NoUppercase() StringInput {
	if !NoUppercase(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nouppercase", nil)
	}
	return si
}
func (si StringInput) IsLowercase() StringInput {
	if !IsLowercase(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("lowercase", nil)
	}
	return si
}
func (si StringInput) NoLowercase() StringInput {
	if !NoLowercase(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nolowercase", nil)
	}
	return si
}
func (si StringInput) IsPrintable() StringInput {
	if !IsPrintable(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("printable", nil)
	}
	return si
}
func (si StringInput) NoPrintable() StringInput {
	if !NoPrintable(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("noprintable", nil)
	}
	return si
}
func (si StringInput) IsAlphabetic() StringInput {
	if !IsAlphabetic(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("alphabetic", nil)
	}
	return si
}
func (si StringInput) NoAlphabetic() StringInput {
	if !NoAlphabetic(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("noalphabetic", nil)
	}
	return si
}
func (si StringInput) IsAlphanumeric() StringInput {
	if !IsAlphanumeric(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("alphanumeric", nil)
	}
	return si
}
func (si StringInput) NoAlphanumeric() StringInput {
	if !NoAlphanumeric(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("noalphanumeric", nil)
	}
	return si
}
func (si StringInput) IsNumeric() StringInput {
	if !IsNumeric(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("numeric", nil)
	}
	return si
}
func (si StringInput) NoNumeric() StringInput {
	if !NoNumeric(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nonumeric", nil)
	}
	return si
}
func (si StringInput) IsDigits() StringInput {
	if !IsDigits(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("digits", nil)
	}
	return si
}
func (si StringInput) NoDigits() StringInput {
	if !NoDigits(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nodigits", nil)
	}
	return si
}
func (si StringInput) IsPunctuation() StringInput {
	if !IsPunctuation(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("punctuation", nil)
	}
	return si
}
func (si StringInput) NoPunctuation() StringInput {
	if !NoPunctuation(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nopunctuation", nil)
	}
	return si
}
func (si StringInput) IsSymbols() StringInput {
	if !IsSymbols(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("symbols", nil)
	}
	return si
}
func (si StringInput) NoSymbols() StringInput {
	if !NoSymbols(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nosymbols", nil)
	}
	return si
}
func (si StringInput) IsMarkCharacters() StringInput {
	if !IsMarkCharacters(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("markchars", nil)
	}
	return si
}
func (si StringInput) NoMarkCharacters() StringInput {
	if !NoMarkCharacters(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nomarkchars", nil)
	}
	return si
}
func (si StringInput) IsWhitespaces() StringInput {
	if !IsWhitespaces(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("spaces", nil)
	}
	return si
}
func (si StringInput) NoWhitespaces() StringInput {
	if !NoWhitespaces(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nospaces", nil)
	}
	return si
}
func (si StringInput) IsControlCharacters() StringInput {
	if !IsControlCharacters(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("controlchars", nil)
	}
	return si
}
func (si StringInput) NoControlCharacters() StringInput {
	if !NoControlCharacters(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nocontrolchars", nil)
	}
	return si
}
func (si StringInput) IsGraphicCharacters() StringInput {
	if !IsGraphicCharacters(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("graphicchars", nil)
	}
	return si
}
func (si StringInput) NoGraphicCharacters() StringInput {
	if !NoGraphicCharacters(si.inputData.StringData) {
		si.inputData = si.inputData.AppendError("nographicchars", nil)
	}
	return si
}
