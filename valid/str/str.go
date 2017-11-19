package validstr

import (
	"reflect"
	"strconv"
	"strings"

	validinput "lib/uput/valid/input"
	validate "lib/uput/valid/str/is"
)

//
// DEV
func PrintErrors(errors []error) {
	validinput.PrintErrors(errors)
}
func PrintValidations(validations []string) {
	validinput.PrintValidations(validations)
}

// Nesting the generic inputData type
type StringInput struct {
	stringData string
	input      validinput.InputData
}

//
// Validation Input Function
// TODO: Add way to customize errors
func If(s string) StringInput {
	return StringInput{
		stringData: s,
		input: validinput.InputData{
			DataType:      reflect.String,
			Data:          s,
			ErrorMessages: (DefaultErrorMessages()),
		},
	}
}

//
// Validation Output Function
func (s StringInput) IsValid() (bool, string, []error, []string) {
	// TODO: Now that validations are tracked, a better ouput that includes errors / validations
	// can be presented

	// TODO: Potentially remove validations in each check in favor of a generic loop
	// that runs the corresponding function and adds corresponding error
	return (len(s.input.Errors) == 0), s.stringData, s.input.Errors, s.input.Validations
}

//
// Chainable String Validations
// ==========================================================================

//
// String Slice Validations
func (s StringInput) IsIn(list []string) StringInput {
	s.input = s.input.AppendValidation("isin")
	if !validate.IsInSlice(s.stringData, list) {
		s.input = s.input.AppendError("isin", []string{strings.Join(list, ", ")})
	}
	return s
}
func (s StringInput) NotIn(list []string) StringInput {
	s.input = s.input.AppendValidation("notin")
	if validate.NotInSlice(s.stringData, list) {
		s.input = s.input.AppendError("notin", []string{strings.Join(list, ", ")})
	}
	return s
}

//
// String Length Validations
func (s StringInput) Required() StringInput {
	s.input = s.input.AppendValidation("required")
	if !validate.Required(s.stringData) {
		s.input = s.input.AppendError("required", nil)
	}
	return s
}
func (s StringInput) IsEmpty() StringInput {
	s.input = s.input.AppendValidation("empty")
	if !validate.IsEmpty(s.stringData) {
		s.input = s.input.AppendError("empty", nil)
	}
	return s
}
func (s StringInput) IsNotEmpty() StringInput {
	// Add validaiton to the data.validations map
	s.input = s.input.AppendValidation("notempty")
	if !validate.NotEmpty(s.stringData) {
		s.input = s.input.AppendError("notempty", nil)
	}
	return s
}
func (s StringInput) IsBetween(start, end int) StringInput {
	s.input = s.input.AppendValidation("between")
	if !validate.IsBetween(s.stringData, start, end) {
		s.input = s.input.AppendError("between", []string{strconv.Itoa(start), strconv.Itoa(end)})
	}
	return s
}
func (s StringInput) IsLessThan(lt int) StringInput {
	s.input = s.input.AppendValidation("lessthan")
	if !validate.IsLessThan(s.stringData, lt) {
		s.input = s.input.AppendError("lessthan", []string{strconv.Itoa(lt)})
	}
	return s
}
func (s StringInput) IsGreaterThan(gt int) StringInput {
	s.input = s.input.AppendValidation("greaterthan")
	if !validate.IsGreaterThan(s.stringData, gt) {
		s.input = s.input.AppendError("greaterthan", []string{strconv.Itoa(gt)})
	}
	return s
}

//
// Substring Validation
// WARNING: DOES NOT WORK FOR UTF8 MATCHING
// This will let through look-alikes
func (s StringInput) IsContaining(ss string) StringInput {
	s.input = s.input.AppendValidation("iscontaining")
	if !validate.IsContaining(s.stringData, ss) {
		s.input = s.input.AppendError("iscontaining", []string{ss})
	}
	return s
}
func (s StringInput) NotContaining(ss string) StringInput {
	s.input = s.input.AppendValidation("notcontaiyyning")
	if !validate.NotContaining(s.stringData, ss) {
		s.input = s.input.AppendError("notcontaining", []string{ss})
	}
	return s
}

//
// Regex Validation
func (s StringInput) IsRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation("regexmatch")
	if !validate.IsRegexMatch(s.stringData, pattern) {
		s.input = s.input.AppendError("regexmatch", []string{pattern})
	}
	return s
}
func (s StringInput) NoRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation("noregexmatch")
	if !validate.NoRegexMatch(s.stringData, pattern) {
		s.input = s.input.AppendError("noregexmatch", []string{pattern})
	}
	return s
}

//
// UTF8 Rune Validation
func (s StringInput) IsUTF8() StringInput {
	s.input = s.input.AppendValidation("utf8")
	if !validate.IsUTF8(s.stringData) {
		s.input = s.input.AppendError("utf8", nil)
	}
	return s
}
func (s StringInput) NoUTF8() StringInput {
	s.input = s.input.AppendValidation("noutf8")
	if !validate.NoUTF8(s.stringData) {
		s.input = s.input.AppendError("noutf8", nil)
	}
	return s
}
func (s StringInput) IsUppercase() StringInput {
	s.input = s.input.AppendValidation("uppercase")
	if !validate.IsUppercase(s.stringData) {
		s.input = s.input.AppendError("uppercase", nil)
	}
	return s
}
func (s StringInput) NoUppercase() StringInput {
	s.input = s.input.AppendValidation("regexmatch")
	if !validate.NoUppercase(s.stringData) {
		s.input = s.input.AppendError("nouppercase", nil)
	}
	return s
}
func (s StringInput) IsLowercase() StringInput {
	s.input = s.input.AppendValidation("lowercase")
	if !validate.IsLowercase(s.stringData) {
		s.input = s.input.AppendError("lowercase", nil)
	}
	return s
}
func (s StringInput) NoLowercase() StringInput {
	s.input = s.input.AppendValidation("nolowercase")
	if !validate.NoLowercase(s.stringData) {
		s.input = s.input.AppendError("nolowercase", nil)
	}
	return s
}
func (s StringInput) IsPrintable() StringInput {
	s.input = s.input.AppendValidation("printable")
	if !validate.IsPrintable(s.stringData) {
		s.input = s.input.AppendError("printable", nil)
	}
	return s
}
func (s StringInput) NoPrintable() StringInput {
	s.input = s.input.AppendValidation("noprintable")
	if !validate.NoPrintable(s.stringData) {
		s.input = s.input.AppendError("noprintable", nil)
	}
	return s
}
func (s StringInput) IsAlphabetic() StringInput {
	s.input = s.input.AppendValidation("alphabetic")
	if !validate.IsAlphabetic(s.stringData) {
		s.input = s.input.AppendError("alphabetic", nil)
	}
	return s
}
func (s StringInput) NoAlphabetic() StringInput {
	s.input = s.input.AppendValidation("noalphabetic")
	if !validate.NoAlphabetic(s.stringData) {
		s.input = s.input.AppendError("noalphabetic", nil)
	}
	return s
}
func (s StringInput) IsAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("alphanumeric")
	if !validate.IsAlphanumeric(s.stringData) {
		s.input = s.input.AppendError("alphanumeric", nil)
	}
	return s
}
func (s StringInput) NoAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("noalphanumeric")
	if !validate.NoAlphanumeric(s.stringData) {
		s.input = s.input.AppendError("noalphanumeric", nil)
	}
	return s
}
func (s StringInput) IsNumeric() StringInput {
	s.input = s.input.AppendValidation("numeric")
	if !validate.IsNumeric(s.stringData) {
		s.input = s.input.AppendError("numeric", nil)
	}
	return s
}
func (s StringInput) NoNumeric() StringInput {
	s.input = s.input.AppendValidation("nonumeric")
	if !validate.NoNumeric(s.stringData) {
		s.input = s.input.AppendError("nonumeric", nil)
	}
	return s
}
func (s StringInput) IsDigits() StringInput {
	s.input = s.input.AppendValidation("digits")
	if !validate.IsDigits(s.stringData) {
		s.input = s.input.AppendError("digits", nil)
	}
	return s
}
func (s StringInput) NoDigits() StringInput {
	s.input = s.input.AppendValidation("nodigits")
	if !validate.NoDigits(s.stringData) {
		s.input = s.input.AppendError("nodigits", nil)
	}
	return s
}
func (s StringInput) IsPunctuation() StringInput {
	s.input = s.input.AppendValidation("punctuation")
	if !validate.IsPunctuation(s.stringData) {
		s.input = s.input.AppendError("punctuation", nil)
	}
	return s
}
func (s StringInput) NoPunctuation() StringInput {
	s.input = s.input.AppendValidation("nopunctuation")
	if !validate.NoPunctuation(s.stringData) {
		s.input = s.input.AppendError("nopunctuation", nil)
	}
	return s
}
func (s StringInput) IsSymbols() StringInput {
	s.input = s.input.AppendValidation("symbols")
	if !validate.IsSymbols(s.stringData) {
		s.input = s.input.AppendError("symbols", nil)
	}
	return s
}
func (s StringInput) NoSymbols() StringInput {
	s.input = s.input.AppendValidation("nosymbols")
	if !validate.NoSymbols(s.stringData) {
		s.input = s.input.AppendError("nosymbols", nil)
	}
	return s
}
func (s StringInput) IsMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("markchars")
	if !validate.IsMarkCharacters(s.stringData) {
		s.input = s.input.AppendError("markchars", nil)
	}
	return s
}
func (s StringInput) NoMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("nomarkchars")
	if !validate.NoMarkCharacters(s.stringData) {
		s.input = s.input.AppendError("nomarkchars", nil)
	}
	return s
}
func (s StringInput) IsWhitespaces() StringInput {
	s.input = s.input.AppendValidation("spaces")
	if !validate.IsWhitespaces(s.stringData) {
		s.input = s.input.AppendError("spaces", nil)
	}
	return s
}
func (s StringInput) NoWhitespaces() StringInput {
	s.input = s.input.AppendValidation("nospaces")
	if !validate.NoWhitespaces(s.stringData) {
		s.input = s.input.AppendError("nospaces", nil)
	}
	return s
}
func (s StringInput) IsControlCharacters() StringInput {
	s.input = s.input.AppendValidation("controlchars")
	if !validate.IsControlCharacters(s.stringData) {
		s.input = s.input.AppendError("controlchars", nil)
	}
	return s
}
func (s StringInput) NoControlCharacters() StringInput {
	s.input = s.input.AppendValidation("nocontrolchars")
	if !validate.NoControlCharacters(s.stringData) {
		s.input = s.input.AppendError("nocontrolchars", nil)
	}
	return s
}
func (s StringInput) IsGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("graphicchars")
	if !validate.IsGraphicCharacters(s.stringData) {
		s.input = s.input.AppendError("graphicchars", nil)
	}
	return s
}
func (s StringInput) NoGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("nographicchars")
	if !validate.NoGraphicCharacters(s.stringData) {
		s.input = s.input.AppendError("nographicchars", nil)
	}
	return s
}
