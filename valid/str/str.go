package validstr

import (
	"fmt" // DEV
	"reflect"
	"strconv"

	validinput "lib/uput/valid/input"
	validate "lib/uput/valid/str/is"
)

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
			DataType:       reflect.String,
			DataTypeName:   "string",
			ValidationText: (DefaultValidationText()),
		},
	}
}

func (s StringInput) isValid() bool {
	return (len(s.input.InputErrors) == 0)
}

//
// Validation Output Function
// ==========================================================================
func (s StringInput) IsValid() (bool, string, []error) {
	s.input.PrintValidations()
	s.input.PrintErrors()
	return s.isValid(), s.stringData, s.input.Errors()
}

//
// Custom Validations
// ==========================================================================

// TODO: Add niche validations from /is/ like /is/email

//
// Localize Error Message & Validation Descriptions
// ==========================================================================

func (s StringInput) ErrorMessage(message string) StringInput {
	lastError := s.input.LastInputError()
	fmt.Println("test message: ", message)
	fmt.Println("last validation: ", lastError)
	return s
}

func (s StringInput) UpdateValidationText(key string, text validinput.ValidationText) map[string]validinput.ValidationText {
	s.input.ValidationText = s.input.UpdateText(key, text)
	return s.input.ValidationText
}

func (s StringInput) UpdateErrorMessages(errorMessages map[string]string) map[string]validinput.ValidationText {
	s.input.ValidationText = s.input.UpdateValidationText("Error", errorMessages)
	return s.input.ValidationText
}

func (s StringInput) UpdateValidationDescriptions(descriptions map[string]string) map[string]validinput.ValidationText {
	s.input.ValidationText = s.input.UpdateValidationText("Description", descriptions)
	return s.input.ValidationText
}

//
// Chainable String Validations
// ==========================================================================

//
// String Slice Validations
func (s StringInput) IsIn(list []string) StringInput {
	s.input = s.input.AppendValidation("isin", list)
	if !validate.IsInSlice(s.stringData, list) {
		s.input = s.input.AppendError("isin", list)
	}
	return s
}
func (s StringInput) NotIn(list []string) StringInput {
	s.input = s.input.AppendValidation("notin", list)
	if validate.NotInSlice(s.stringData, list) {
		s.input = s.input.AppendError("notin", list)
	}
	return s
}

//
// String Length Validations
func (s StringInput) Required() StringInput {
	s.input = s.input.AppendValidation("required", nil)
	if !validate.Required(s.stringData) {
		s.input = s.input.AppendError("required", nil)
	}
	return s
}
func (s StringInput) IsEmpty() StringInput {
	s.input = s.input.AppendValidation("empty", nil)
	if !validate.IsEmpty(s.stringData) {
		s.input = s.input.AppendError("empty", nil)
	}
	return s
}
func (s StringInput) IsNotEmpty() StringInput {
	// Add validaiton to the data.validations map
	s.input = s.input.AppendValidation("notempty", nil)
	if !validate.NotEmpty(s.stringData) {
		s.input = s.input.AppendError("notempty", nil)
	}
	return s
}
func (s StringInput) IsBetween(start, end int) StringInput {
	values := []string{strconv.Itoa(start), strconv.Itoa(end)}
	s.input = s.input.AppendValidation("between", values)
	if !validate.IsBetween(s.stringData, start, end) {
		s.input = s.input.AppendError("between", values)
	}
	return s
}
func (s StringInput) IsLessThan(lt int) StringInput {
	values := []string{strconv.Itoa(lt)}
	s.input = s.input.AppendValidation("lessthan", values)
	if !validate.IsLessThan(s.stringData, lt) {
		s.input = s.input.AppendError("lessthan", values)
	}
	return s
}
func (s StringInput) IsGreaterThan(gt int) StringInput {
	values := []string{strconv.Itoa(gt)}
	s.input = s.input.AppendValidation("greaterthan", values)
	if !validate.IsGreaterThan(s.stringData, gt) {
		s.input = s.input.AppendError("greaterthan", values)
	}
	return s
}

//
// Substring Validation
// WARNING: DOES NOT WORK FOR UTF8 MATCHING
// This will let through look-alikes
func (s StringInput) IsContaining(ss string) StringInput {
	values := []string{ss}
	s.input = s.input.AppendValidation("iscontaining", values)
	if !validate.IsContaining(s.stringData, ss) {
		s.input = s.input.AppendError("iscontaining", values)
	}
	return s
}
func (s StringInput) NotContaining(ss string) StringInput {
	values := []string{ss}
	s.input = s.input.AppendValidation("notcontaiyyning", values)
	if !validate.NotContaining(s.stringData, ss) {
		s.input = s.input.AppendError("notcontaining", values)
	}
	return s
}

//
// Regex Validation
func (s StringInput) IsRegexMatch(pattern string) StringInput {
	values := []string{pattern}
	s.input = s.input.AppendValidation("regexmatch", values)
	if !validate.IsRegexMatch(s.stringData, pattern) {
		s.input = s.input.AppendError("regexmatch", values)
	}
	return s
}
func (s StringInput) NoRegexMatch(pattern string) StringInput {
	values := []string{pattern}
	s.input = s.input.AppendValidation("noregexmatch", values)
	if !validate.NoRegexMatch(s.stringData, pattern) {
		s.input = s.input.AppendError("noregexmatch", values)
	}
	return s
}

//
// UTF8 Rune Validation
func (s StringInput) IsUTF8() StringInput {
	s.input = s.input.AppendValidation("utf8", nil)
	if !validate.IsUTF8(s.stringData) {
		s.input = s.input.AppendError("utf8", nil)
	}
	return s
}
func (s StringInput) NoUTF8() StringInput {
	s.input = s.input.AppendValidation("noutf8", nil)
	if !validate.NoUTF8(s.stringData) {
		s.input = s.input.AppendError("noutf8", nil)
	}
	return s
}
func (s StringInput) IsUppercase() StringInput {
	s.input = s.input.AppendValidation("uppercase", nil)
	if !validate.IsUppercase(s.stringData) {
		s.input = s.input.AppendError("uppercase", nil)
	}
	return s
}
func (s StringInput) NoUppercase() StringInput {
	s.input = s.input.AppendValidation("nouppercase", nil)
	if !validate.NoUppercase(s.stringData) {
		s.input = s.input.AppendError("nouppercase", nil)
	}
	return s
}
func (s StringInput) IsLowercase() StringInput {
	s.input = s.input.AppendValidation("lowercase", nil)
	if !validate.IsLowercase(s.stringData) {
		s.input = s.input.AppendError("lowercase", nil)
	}
	return s
}
func (s StringInput) NoLowercase() StringInput {
	s.input = s.input.AppendValidation("nolowercase", nil)
	if !validate.NoLowercase(s.stringData) {
		s.input = s.input.AppendError("nolowercase", nil)
	}
	return s
}
func (s StringInput) IsPrintable() StringInput {
	s.input = s.input.AppendValidation("printable", nil)
	if !validate.IsPrintable(s.stringData) {
		s.input = s.input.AppendError("printable", nil)
	}
	return s
}
func (s StringInput) NoPrintable() StringInput {
	s.input = s.input.AppendValidation("noprintable", nil)
	if !validate.NoPrintable(s.stringData) {
		s.input = s.input.AppendError("noprintable", nil)
	}
	return s
}
func (s StringInput) IsAlphabetic() StringInput {
	s.input = s.input.AppendValidation("alphabetic", nil)
	if !validate.IsAlphabetic(s.stringData) {
		s.input = s.input.AppendError("alphabetic", nil)
	}
	return s
}
func (s StringInput) NoAlphabetic() StringInput {
	s.input = s.input.AppendValidation("noalphabetic", nil)
	if !validate.NoAlphabetic(s.stringData) {
		s.input = s.input.AppendError("noalphabetic", nil)
	}
	return s
}
func (s StringInput) IsAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("alphanumeric", nil)
	if !validate.IsAlphanumeric(s.stringData) {
		s.input = s.input.AppendError("alphanumeric", nil)
	}
	return s
}
func (s StringInput) NoAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("noalphanumeric", nil)
	if !validate.NoAlphanumeric(s.stringData) {
		s.input = s.input.AppendError("noalphanumeric", nil)
	}
	return s
}
func (s StringInput) IsNumeric() StringInput {
	s.input = s.input.AppendValidation("numeric", nil)
	if !validate.IsNumeric(s.stringData) {
		s.input = s.input.AppendError("numeric", nil)
	}
	return s
}
func (s StringInput) NoNumeric() StringInput {
	s.input = s.input.AppendValidation("nonumeric", nil)
	if !validate.NoNumeric(s.stringData) {
		s.input = s.input.AppendError("nonumeric", nil)
	}
	return s
}
func (s StringInput) IsDigits() StringInput {
	s.input = s.input.AppendValidation("digits", nil)
	if !validate.IsDigits(s.stringData) {
		s.input = s.input.AppendError("digits", nil)
	}
	return s
}
func (s StringInput) NoDigits() StringInput {
	s.input = s.input.AppendValidation("nodigits", nil)
	if !validate.NoDigits(s.stringData) {
		s.input = s.input.AppendError("nodigits", nil)
	}
	return s
}
func (s StringInput) IsPunctuation() StringInput {
	s.input = s.input.AppendValidation("punctuation", nil)
	if !validate.IsPunctuation(s.stringData) {
		s.input = s.input.AppendError("punctuation", nil)
	}
	return s
}
func (s StringInput) NoPunctuation() StringInput {
	s.input = s.input.AppendValidation("nopunctuation", nil)
	if !validate.NoPunctuation(s.stringData) {
		s.input = s.input.AppendError("nopunctuation", nil)
	}
	return s
}
func (s StringInput) IsSymbols() StringInput {
	s.input = s.input.AppendValidation("symbols", nil)
	if !validate.IsSymbols(s.stringData) {
		s.input = s.input.AppendError("symbols", nil)
	}
	return s
}
func (s StringInput) NoSymbols() StringInput {
	s.input = s.input.AppendValidation("nosymbols", nil)
	if !validate.NoSymbols(s.stringData) {
		s.input = s.input.AppendError("nosymbols", nil)
	}
	return s
}
func (s StringInput) IsMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("markchars", nil)
	if !validate.IsMarkCharacters(s.stringData) {
		s.input = s.input.AppendError("markchars", nil)
	}
	return s
}
func (s StringInput) NoMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("nomarkchars", nil)
	if !validate.NoMarkCharacters(s.stringData) {
		s.input = s.input.AppendError("nomarkchars", nil)
	}
	return s
}
func (s StringInput) IsWhitespaces() StringInput {
	s.input = s.input.AppendValidation("spaces", nil)
	if !validate.IsWhitespaces(s.stringData) {
		s.input = s.input.AppendError("spaces", nil)
	}
	return s
}
func (s StringInput) NoWhitespaces() StringInput {
	s.input = s.input.AppendValidation("nospaces", nil)
	if !validate.NoWhitespaces(s.stringData) {
		s.input = s.input.AppendError("nospaces", nil)
	}
	return s
}
func (s StringInput) IsControlCharacters() StringInput {
	s.input = s.input.AppendValidation("controlchars", nil)
	if !validate.IsControlCharacters(s.stringData) {
		s.input = s.input.AppendError("controlchars", nil)
	}
	return s
}
func (s StringInput) NoControlCharacters() StringInput {
	s.input = s.input.AppendValidation("nocontrolchars", nil)
	if !validate.NoControlCharacters(s.stringData) {
		s.input = s.input.AppendError("nocontrolchars", nil)
	}
	return s
}
func (s StringInput) IsGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("graphicchars", nil)
	if !validate.IsGraphicCharacters(s.stringData) {
		s.input = s.input.AppendError("graphicchars", nil)
	}
	return s
}
func (s StringInput) NoGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("nographicchars", nil)
	if !validate.NoGraphicCharacters(s.stringData) {
		s.input = s.input.AppendError("nographicchars", nil)
	}
	return s
}
