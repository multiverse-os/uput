package validstr

import (
	"reflect"
	"strconv"
	"strings"

	validinput "lib/uput/valid/input"
)

// Nesting the generic inputData type
type StringInput struct {
	input validinput.InputData
}

//
// Validation Input Function
// TODO: Add way to customize errors
func If(s string) StringInput {
	return StringInput{
		input: validinput.InputData{
			DataType:      reflect.String,
			Data:          s,
			ErrorMessages: (DefaultErrorMessages()),
		},
	}
}

//
// Validation Output Function
func (s StringInput) IsValid() (bool, interface{}, []error, []string) {
	// TODO: Now that validations are tracked, a better ouput that includes errors / validations
	// can be presented

	// TODO: Potentially remove validations in each check in favor of a generic loop
	// that runs the corresponding function and adds corresponding error
	return (len(s.input.Errors) == 0), s.input.Data, s.input.Errors
}

//
// Chainable String Validations
// ==========================================================================

//
// String Slice Validations
func (s StringInput) IsIn(list []string) StringInput {
	s.input = s.input.AppendValidation("isin")
	if !IsInSlice(s.input.Data, list) {
		s.input = s.input.AppendError("isin", []string{strings.Join(list, ", ")})
	}
	return s
}
func (s StringInput) NotIn(list []string) StringInput {
	s.input = s.input.AppendValidation("notin")
	if NotInSlice(s.input.Data, list) {
		s.input = s.input.AppendError("notin", []string{strings.Join(list, ", ")})
	}
	return s
}

//
// String Length Validations
func (s StringInput) Required() StringInput {
	s.input = s.input.AppendValidation("required")
	if !Required(s.input.Data) {
		s.input = s.input.AppendError("required", nil)
	}
	return s
}
func (s StringInput) IsEmpty() StringInput {
	s.input = s.input.AppendValidation("empty")
	if !IsEmpty(s.input.Data) {
		s.input = s.input.AppendError("empty", nil)
	}
	return s
}
func (s StringInput) IsNotEmpty() StringInput {
	// Add validaiton to the data.validations map
	s.input = s.input.AppendValidation("notempty")
	if !NotEmpty(s.input.Data) {
		s.input = s.input.AppendError("notempty", nil)
	}
	return s
}
func (s StringInput) IsBetween(start, end int) StringInput {
	s.input = s.input.AppendValidation("between")
	if !IsBetween(s.input.Data, start, end) {
		s.input = s.input.AppendError("between", []string{strconv.Itoa(start), strconv.Itoa(end)})
	}
	return s
}
func (s StringInput) IsLessThan(lt int) StringInput {
	s.input = s.input.AppendValidation("lessthan")
	if !IsLessThan(s.input.Data, lt) {
		s.input = s.input.AppendError("lessthan", []string{strconv.Itoa(lt)})
	}
	return s
}
func (s StringInput) IsGreaterThan(gt int) StringInput {
	s.input = s.input.AppendValidation("greaterthan")
	if !IsGreaterThan(s.input.Data, gt) {
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
	if !IsContaining(s.input.Data, ss) {
		s.input = s.input.AppendError("iscontaining", []string{ss})
	}
	return s
}
func (s StringInput) NotContaining(ss string) StringInput {
	s.input = s.input.AppendValidation("notcontaiyyning")
	if !NotContaining(s.input.Data, ss) {
		s.input = s.input.AppendError("notcontaining", []string{ss})
	}
	return s
}

//
// Regex Validation
func (s StringInput) IsRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation("regexmatch")
	if !IsRegexMatch(s.input.Data, pattern) {
		s.input = s.input.AppendError("regexmatch", []string{pattern})
	}
	return s
}
func (s StringInput) NoRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation("noregexmatch")
	if !NoRegexMatch(s.input.Data, pattern) {
		s.input = s.input.AppendError("noregexmatch", []string{pattern})
	}
	return s
}

//
// UTF8 Rune Validation
func (s StringInput) IsUTF8() StringInput {
	s.input = s.input.AppendValidation("utf8")
	if !IsUTF8(s.input.Data) {
		s.input = s.input.AppendError("utf8", nil)
	}
	return s
}
func (s StringInput) NoUTF8() StringInput {
	s.input = s.input.AppendValidation("noutf8")
	if !NoUTF8(s.input.Data) {
		s.input = s.input.AppendError("noutf8", nil)
	}
	return s
}
func (s StringInput) IsUppercase() StringInput {
	s.input = s.input.AppendValidation("uppercase")
	if !IsUppercase(s.input.Data) {
		s.input = s.input.AppendError("uppercase", nil)
	}
	return s
}
func (s StringInput) NoUppercase() StringInput {
	s.input = s.input.AppendValidation("regexmatch")
	if !NoUppercase(s.input.Data) {
		s.input = s.input.AppendError("nouppercase", nil)
	}
	return s
}
func (s StringInput) IsLowercase() StringInput {
	s.input = s.input.AppendValidation("lowercase")
	if !IsLowercase(s.input.Data) {
		s.input = s.input.AppendError("lowercase", nil)
	}
	return s
}
func (s StringInput) NoLowercase() StringInput {
	s.input = s.input.AppendValidation("nolowercase")
	if !NoLowercase(s.input.Data) {
		s.input = s.input.AppendError("nolowercase", nil)
	}
	return s
}
func (s StringInput) IsPrintable() StringInput {
	s.input = s.input.AppendValidation("printable")
	if !IsPrintable(s.input.Data) {
		s.input = s.input.AppendError("printable", nil)
	}
	return s
}
func (s StringInput) NoPrintable() StringInput {
	s.input = s.input.AppendValidation("noprintable")
	if !NoPrintable(s.input.Data) {
		s.input = s.input.AppendError("noprintable", nil)
	}
	return s
}
func (s StringInput) IsAlphabetic() StringInput {
	s.input = s.input.AppendValidation("alphabetic")
	if !IsAlphabetic(s.input.Data) {
		s.input = s.input.AppendError("alphabetic", nil)
	}
	return s
}
func (s StringInput) NoAlphabetic() StringInput {
	s.input = s.input.AppendValidation("noalphabetic")
	if !NoAlphabetic(s.input.Data) {
		s.input = s.input.AppendError("noalphabetic", nil)
	}
	return s
}
func (s StringInput) IsAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("alphanumeric")
	if !IsAlphanumeric(s.input.Data) {
		s.input = s.input.AppendError("alphanumeric", nil)
	}
	return s
}
func (s StringInput) NoAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("noalphanumeric")
	if !NoAlphanumeric(s.input.Data) {
		s.input = s.input.AppendError("noalphanumeric", nil)
	}
	return s
}
func (s StringInput) IsNumeric() StringInput {
	s.input = s.input.AppendValidation("numeric")
	if !IsNumeric(s.input.Data) {
		s.input = s.input.AppendError("numeric", nil)
	}
	return s
}
func (s StringInput) NoNumeric() StringInput {
	s.input = s.input.AppendValidation("nonumeric")
	if !NoNumeric(s.input.Data) {
		s.input = s.input.AppendError("nonumeric", nil)
	}
	return s
}
func (s StringInput) IsDigits() StringInput {
	s.input = s.input.AppendValidation("digits")
	if !IsDigits(s.input.Data) {
		s.input = s.input.AppendError("digits", nil)
	}
	return s
}
func (s StringInput) NoDigits() StringInput {
	s.input = s.input.AppendValidation("nodigits")
	if !NoDigits(s.input.Data) {
		s.input = s.input.AppendError("nodigits", nil)
	}
	return s
}
func (s StringInput) IsPunctuation() StringInput {
	s.input = s.input.AppendValidation("punctuation")
	if !IsPunctuation(s.input.Data) {
		s.input = s.input.AppendError("punctuation", nil)
	}
	return s
}
func (s StringInput) NoPunctuation() StringInput {
	s.input = s.input.AppendValidation("nopunctuation")
	if !NoPunctuation(s.input.Data) {
		s.input = s.input.AppendError("nopunctuation", nil)
	}
	return s
}
func (s StringInput) IsSymbols() StringInput {
	s.input = s.input.AppendValidation("symbols")
	if !IsSymbols(s.input.Data) {
		s.input = s.input.AppendError("symbols", nil)
	}
	return s
}
func (s StringInput) NoSymbols() StringInput {
	s.input = s.input.AppendValidation("nosymbols")
	if !NoSymbols(s.input.Data) {
		s.input = s.input.AppendError("nosymbols", nil)
	}
	return s
}
func (s StringInput) IsMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("markchars")
	if !IsMarkCharacters(s.input.Data) {
		s.input = s.input.AppendError("markchars", nil)
	}
	return s
}
func (s StringInput) NoMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("nomarkchars")
	if !NoMarkCharacters(s.input.Data) {
		s.input = s.input.AppendError("nomarkchars", nil)
	}
	return s
}
func (s StringInput) IsWhitespaces() StringInput {
	s.input = s.input.AppendValidation("spaces")
	if !IsWhitespaces(s.input.Data) {
		s.input = s.input.AppendError("spaces", nil)
	}
	return s
}
func (s StringInput) NoWhitespaces() StringInput {
	s.input = s.input.AppendValidation("nospaces")
	if !NoWhitespaces(s.input.Data) {
		s.input = s.input.AppendError("nospaces", nil)
	}
	return s
}
func (s StringInput) IsControlCharacters() StringInput {
	s.input = s.input.AppendValidation("controlchars")
	if !IsControlCharacters(s.input.Data) {
		s.input = s.input.AppendError("controlchars", nil)
	}
	return s
}
func (s StringInput) NoControlCharacters() StringInput {
	s.input = s.input.AppendValidation("nocontrolchars")
	if !NoControlCharacters(s.input.Data) {
		s.input = s.input.AppendError("nocontrolchars", nil)
	}
	return s
}
func (s StringInput) IsGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("graphicchars")
	if !IsGraphicCharacters(s.input.Data) {
		s.input = s.input.AppendError("graphicchars", nil)
	}
	return s
}
func (s StringInput) NoGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("nographicchars")
	if !NoGraphicCharacters(s.input.Data) {
		s.input = s.input.AppendError("nographicchars", nil)
	}
	return s
}
