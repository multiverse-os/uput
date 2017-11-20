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
// ==========================================================================
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
	return (s.input.IsValid())
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
	s.input.ValidationText = s.input.UpdateLastValidationText(validinput.ValidationText{Error: message})
	fmt.Println("test message: ", message)
	return s
}

func (s StringInput) UpdateValidationText(key string, text validinput.ValidationText) StringInput {
	s.input.ValidationText = s.input.SetValidationText(key, text)
	return s
}

func (s StringInput) UpdateErrorMessages(errorMessages map[string]string) StringInput {
	s.input.ValidationText = s.input.SetAllTextOfType("Error", errorMessages)
	return s
}

func (s StringInput) UpdateValidationDescriptions(descriptions map[string]string) StringInput {
	s.input.ValidationText = s.input.SetAllTextOfType("Description", descriptions)
	return s
}

//
// Chainable String Validations
// ==========================================================================

//
// String Slice Validations
func (s StringInput) IsIn(list []string) StringInput {
	s.input = s.input.AppendValidation("isin", list, validate.IsInSlice(s.stringData, list))
	return s
}
func (s StringInput) NotIn(list []string) StringInput {
	s.input = s.input.AppendValidation("notin", list, !validate.IsInSlice(s.stringData, list))
	return s
}

//
// String Length Validations
func (s StringInput) Required() StringInput {
	s.input = s.input.AppendValidation("required", nil, validate.NotEmpty(s.stringData))
	return s
}
func (s StringInput) NotEmpty() StringInput {
	s.input = s.input.AppendValidation("notempty", nil, validate.NotEmpty(s.stringData))
	return s
}
func (s StringInput) IsBetween(start, end int) StringInput {
	s.input = s.input.AppendValidation("between", []string{strconv.Itoa(start), strconv.Itoa(end)}, validate.IsBetween(s.stringData, start, end))
	return s
}
func (s StringInput) IsLessThan(lt int) StringInput {
	s.input = s.input.AppendValidation("lessthan", []string{strconv.Itoa(lt)}, validate.IsLessThan(s.stringData, lt))
	return s
}
func (s StringInput) IsGreaterThan(gt int) StringInput {
	s.input = s.input.AppendValidation("greaterthan", []string{strconv.Itoa(gt)}, validate.IsGreaterThan(s.stringData, gt))
	return s
}

//
// Substring Validation
// WARNING: DOES NOT WORK FOR UTF8 MATCHING
// This will let through look-alikes, like K
// and K for kelvin temperature.
func (s StringInput) Contains(ss string) StringInput {
	s.input = s.input.AppendValidation("contains", []string{ss}, validate.Contains(s.stringData, ss))
	return s
}
func (s StringInput) NotContaining(ss string) StringInput {
	s.input = s.input.AppendValidation("notcontaining", []string{ss}, !validate.Contains(s.stringData, ss))
	return s
}

//
// Regex Validation
func (s StringInput) IsRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation("regexmatch", []string{pattern}, validate.IsRegexMatch(s.stringData, pattern))
	return s
}
func (s StringInput) NoRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation("noregexmatch", []string{pattern}, !validate.IsRegexMatch(s.stringData, pattern))
	return s
}

//
// UTF8 Validation
func (s StringInput) IsUTF8() StringInput {
	s.input = s.input.AppendValidation("utf8", nil, validate.IsUTF8(s.stringData))
	return s
}
func (s StringInput) NoUTF8() StringInput {
	s.input = s.input.AppendValidation("noutf8", nil, !validate.IsUTF8(s.stringData))
	return s
}

//
// UTF8 Rune Validation
func (s StringInput) IsAlphabetic() StringInput {
	s.input = s.input.AppendValidation("alphabetic", nil, validate.Alphabetic(s.stringData, true, 0))
	return s
}
func (s StringInput) NoAlphabetic() StringInput {
	s.input = s.input.AppendValidation("noalphabetic", nil, validate.Alphabetic(s.stringData, false, 0))
	return s
}
func (s StringInput) MinAlphabeticCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("alphabetic", nil, validate.Alphabetic(s.stringData, true, count))
	return s
}
func (s StringInput) IsAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("alphanumeric", nil, validate.Alphanumeric(s.stringData, true, 0))
	return s
}
func (s StringInput) NoAlphanumeric() StringInput {
	s.input = s.input.AppendValidation("noalphanumeric", nil, validate.Alphanumeric(s.stringData, false, 0))
	return s
}
func (s StringInput) IsNumeric() StringInput {
	s.input = s.input.AppendValidation("numeric", nil, validate.Numeric(s.stringData, true, 0))
	return s
}
func (s StringInput) NoNumeric() StringInput {
	s.input = s.input.AppendValidation("nonumeric", nil, validate.Numeric(s.stringData, false, 0))
	return s
}
func (s StringInput) MinNumericCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("numeric", nil, validate.Numeric(s.stringData, true, count))
	return s
}
func (s StringInput) IsUppercase() StringInput {
	s.input = s.input.AppendValidation("uppercase", nil, validate.Uppercase(s.stringData, true, 0))
	return s
}
func (s StringInput) NoUppercase() StringInput {
	s.input = s.input.AppendValidation("nouppercase", nil, validate.Uppercase(s.stringData, false, 0))
	return s
}
func (s StringInput) MinUppercaseCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("uppercase", nil, validate.Uppercase(s.stringData, true, count))
	return s
}
func (s StringInput) IsLowercase() StringInput {
	s.input = s.input.AppendValidation("lowercase", nil, validate.Lowercase(s.stringData, true, 0))
	return s
}
func (s StringInput) NoLowercase() StringInput {
	s.input = s.input.AppendValidation("nolowercase", nil, validate.Lowercase(s.stringData, false, 0))
	return s
}
func (s StringInput) MinLowercaseCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("lowercase", nil, validate.Lowercase(s.stringData, true, count))
	return s
}
func (s StringInput) IsPrintable() StringInput {
	s.input = s.input.AppendValidation("printable", nil, validate.Printable(s.stringData, true, 0))
	return s
}
func (s StringInput) NoPrintable() StringInput {
	s.input = s.input.AppendValidation("noprintable", nil, validate.Printable(s.stringData, false, 0))
	return s
}
func (s StringInput) IsPunctuation() StringInput {
	s.input = s.input.AppendValidation("punctuation", nil, validate.Punctuation(s.stringData, true, 0))
	return s
}
func (s StringInput) NoPunctuation() StringInput {
	s.input = s.input.AppendValidation("nopunctuation", nil, validate.Punctuation(s.stringData, false, 0))
	return s
}
func (s StringInput) MinPunctuationCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("nopunctuation", nil, validate.Punctuation(s.stringData, true, count))
	return s
}
func (s StringInput) IsSymbols() StringInput {
	s.input = s.input.AppendValidation("symbols", nil, validate.Symbols(s.stringData, true, 0))
	return s
}
func (s StringInput) NoSymbols() StringInput {
	s.input = s.input.AppendValidation("nosymbols", nil, validate.Symbols(s.stringData, false, 0))
	return s
}
func (s StringInput) MinSymbolCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("symbols", nil, validate.Symbols(s.stringData, true, count))
	return s
}
func (s StringInput) IsWhitespaces() StringInput {
	s.input = s.input.AppendValidation("spaces", nil, validate.Whitespaces(s.stringData, true, 0))
	return s
}
func (s StringInput) NoWhitespaces() StringInput {
	s.input = s.input.AppendValidation("nospaces", nil, validate.Whitespaces(s.stringData, false, 0))
	return s
}
func (s StringInput) IsControlCharacters() StringInput {
	s.input = s.input.AppendValidation("controlchars", nil, validate.ControlCharacters(s.stringData, true, 0))
	return s
}
func (s StringInput) NoControlCharacters() StringInput {
	s.input = s.input.AppendValidation("nocontrolchars", nil, validate.ControlCharacters(s.stringData, false, 0))
	return s
}
func (s StringInput) IsGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("graphicchars", nil, validate.GraphicCharacters(s.stringData, true, 0))
	return s
}
func (s StringInput) NoGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation("nographicchars", nil, validate.GraphicCharacters(s.stringData, false, 0))
	return s
}
func (s StringInput) IsMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("markchars", nil, validate.MarkCharacters(s.stringData, true, 0))
	return s
}
func (s StringInput) NoMarkCharacters() StringInput {
	s.input = s.input.AppendValidation("nomarkchars", nil, validate.MarkCharacters(s.stringData, false, 0))
	return s
}
func (s StringInput) IsDigits() StringInput {
	s.input = s.input.AppendValidation("digits", nil, validate.Digits(s.stringData, true, 0))
	return s
}
func (s StringInput) NoDigits() StringInput {
	s.input = s.input.AppendValidation("nodigits", nil, validate.Digits(s.stringData, false, 0))
	return s
}
func (s StringInput) MinDigitCount(count uint8) StringInput {
	s.input = s.input.AppendValidation("nodigits", nil, validate.Digits(s.stringData, true, count))
	return s
}
