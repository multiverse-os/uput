package validstr

import (
	"strconv"

	validinput "lib/uput/valid/input"
	validate "lib/uput/valid/str/is"

	// DEV
	"fmt"
	inputstatus "lib/uput/valid/input/status"
)

type StringInput struct {
	stringData string
	input      validinput.InputData
}

//
// Validation Input Function
// ==========================================================================
func If(s string) StringInput {
	if validinput.GlobalLocalizedText == nil {
		loadedTextCount := validinput.LoadGlobalLocalizedText((DefaultStringValidationText()))
		fmt.Println("[DEV] Loaded (", loadedTextCount, ") String Validation Localizations")
	}
	return StringInput{
		stringData: s,
		input:      validinput.New(s),
	}
}
func (s StringInput) isValid() bool {
	return (s.input.IsValid())
}

//
// Validation Output Function
// ==========================================================================
func (s StringInput) IsValid() (bool, string, []error) {
	onlyErrors := false
	statusJSON, err := inputstatus.GetStatus(s.input, onlyErrors).Encode(
		map[inputstatus.EncodeOption]string{
			inputstatus.Format: "json",
			inputstatus.Indent: "  ",
		},
	)
	if err == nil {
		fmt.Println(statusJSON)
	}
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
	s.input = s.input.SetLastValidationText(validinput.ValidationText{Error: message})
	return s
}
func (s StringInput) ValidationDescription(message string) StringInput {
	s.input = s.input.SetLastValidationText(validinput.ValidationText{Description: message})
	return s
}
func (s StringInput) ValidationText(description, message string) StringInput {
	s.input = s.input.SetLastValidationText(validinput.ValidationText{Description: message})
	return s
}
func (s StringInput) SetValidationText(key, message, description string) StringInput {
	s.input = s.input.SetValidationText((StringToValidationKey(key)), validinput.ValidationText{Description: description, Error: message})
	return s
}
func (s StringInput) ErrorMessages(errorMessages map[string]string) StringInput {
	for key, message := range errorMessages {
		s.input = s.input.SetValidationText((StringToValidationKey(key)), validinput.ValidationText{Error: message})
	}
	return s
}
func (s StringInput) ValidationDescriptions(descriptions map[string]string) StringInput {
	for key, description := range descriptions {
		s.input = s.input.SetValidationText((StringToValidationKey(key)), validinput.ValidationText{Description: description})
	}
	return s
}

//
// Chainable String Validations
// ==========================================================================

//
// String Slice Validations
func (s StringInput) IsIn(list []string) StringInput {
	s.input = s.input.AppendValidation(In, list, validate.IsInSlice(s.stringData, list))
	return s
}
func (s StringInput) NotIn(list []string) StringInput {
	s.input = s.input.AppendValidation(NotIn, list, !validate.IsInSlice(s.stringData, list))
	return s
}

//
// String Length Validations
func (s StringInput) Required() StringInput {
	s.input = s.input.AppendValidation(Required, nil, validate.IsNotEmpty(s.stringData))
	return s
}
func (s StringInput) IsEmpty() StringInput {
	s.input = s.input.AppendValidation(Empty, nil, validate.IsEmpty(s.stringData))
	return s
}
func (s StringInput) NotEmpty() StringInput {
	s.input = s.input.AppendValidation(NotEmpty, nil, validate.IsNotEmpty(s.stringData))
	return s
}
func (s StringInput) IsBlank() StringInput {
	s.input = s.input.AppendValidation(Blank, nil, validate.IsBlank(s.stringData))
	return s
}
func (s StringInput) IsNotBlank() StringInput {
	s.input = s.input.AppendValidation(NotBlank, nil, validate.IsNotBlank(s.stringData))
	return s
}
func (s StringInput) IsBetween(start, end int) StringInput {
	s.input = s.input.AppendValidation(Between, []string{strconv.Itoa(start), strconv.Itoa(end)}, validate.IsBetween(s.stringData, start, end))
	return s
}
func (s StringInput) IsLessThan(lt int) StringInput {
	s.input = s.input.AppendValidation(LessThan, []string{strconv.Itoa(lt)}, validate.IsLessThan(s.stringData, lt))
	return s
}
func (s StringInput) IsGreaterThan(gt int) StringInput {
	s.input = s.input.AppendValidation(GreaterThan, []string{strconv.Itoa(gt)}, validate.IsGreaterThan(s.stringData, gt))
	return s
}

//
// Substring Validation
// WARNING: DOES NOT WORK FOR UTF8 MATCHING
// This will let through look-alikes, like K
// and K for kelvin temperature.
func (s StringInput) Contains(ss string) StringInput {
	s.input = s.input.AppendValidation(Contains, []string{ss}, validate.Contains(s.stringData, ss))
	return s
}
func (s StringInput) NotContaining(ss string) StringInput {
	s.input = s.input.AppendValidation(NotContaining, []string{ss}, !validate.Contains(s.stringData, ss))
	return s
}

//
// Regex Validation
func (s StringInput) IsRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation(RegexMatch, []string{pattern}, validate.IsRegexMatch(s.stringData, pattern))
	return s
}
func (s StringInput) NoRegexMatch(pattern string) StringInput {
	s.input = s.input.AppendValidation(NoRegexMatch, []string{pattern}, !validate.IsRegexMatch(s.stringData, pattern))
	return s
}

//
// UTF8 Validation
func (s StringInput) IsUTF8() StringInput {
	s.input = s.input.AppendValidation(UTF8, nil, validate.IsUTF8(s.stringData))
	return s
}
func (s StringInput) NoUTF8() StringInput {
	s.input = s.input.AppendValidation(NoUTF8, nil, !validate.IsUTF8(s.stringData))
	return s
}

//
// UTF8 Rune Validation
func (s StringInput) IsAlphabetic() StringInput {
	s.input = s.input.AppendValidation(Alphabetic, nil, validate.Alphabetic(s.stringData, true, 0))
	return s
}
func (s StringInput) NoAlphabetic() StringInput {
	s.input = s.input.AppendValidation(NoAlphabetic, nil, validate.Alphabetic(s.stringData, false, 0))
	return s
}
func (s StringInput) MinAlphabeticCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinAlphabetic, nil, validate.Alphabetic(s.stringData, true, count))
	return s
}
func (s StringInput) IsAlphanumeric() StringInput {
	s.input = s.input.AppendValidation(Alphanumeric, nil, validate.Alphanumeric(s.stringData, true, 0))
	return s
}
func (s StringInput) NoAlphanumeric() StringInput {
	s.input = s.input.AppendValidation(NoAlphanumeric, nil, validate.Alphanumeric(s.stringData, false, 0))
	return s
}
func (s StringInput) IsNumeric() StringInput {
	s.input = s.input.AppendValidation(Numeric, nil, validate.Numeric(s.stringData, true, 0))
	return s
}
func (s StringInput) NoNumeric() StringInput {
	s.input = s.input.AppendValidation(NoNumeric, nil, validate.Numeric(s.stringData, false, 0))
	return s
}
func (s StringInput) MinNumericCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinNumeric, nil, validate.Numeric(s.stringData, true, count))
	return s
}
func (s StringInput) IsUppercase() StringInput {
	s.input = s.input.AppendValidation(Uppercase, nil, validate.Uppercase(s.stringData, true, 0))
	return s
}
func (s StringInput) NoUppercase() StringInput {
	s.input = s.input.AppendValidation(NoUppercase, nil, validate.Uppercase(s.stringData, false, 0))
	return s
}
func (s StringInput) MinUppercaseCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinUppercase, nil, validate.Uppercase(s.stringData, true, count))
	return s
}
func (s StringInput) IsLowercase() StringInput {
	s.input = s.input.AppendValidation(Lowercase, nil, validate.Lowercase(s.stringData, true, 0))
	return s
}
func (s StringInput) NoLowercase() StringInput {
	s.input = s.input.AppendValidation(NoLowercase, nil, validate.Lowercase(s.stringData, false, 0))
	return s
}
func (s StringInput) MinLowercaseCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinLowercase, nil, validate.Lowercase(s.stringData, true, count))
	return s
}
func (s StringInput) IsPrintable() StringInput {
	s.input = s.input.AppendValidation(Printable, nil, validate.Printable(s.stringData, true, 0))
	return s
}
func (s StringInput) NoPrintable() StringInput {
	s.input = s.input.AppendValidation(NoPrintable, nil, validate.Printable(s.stringData, false, 0))
	return s
}
func (s StringInput) IsPunctuation() StringInput {
	s.input = s.input.AppendValidation(Punctuation, nil, validate.Punctuation(s.stringData, true, 0))
	return s
}
func (s StringInput) NoPunctuation() StringInput {
	s.input = s.input.AppendValidation(NoPunctuation, nil, validate.Punctuation(s.stringData, false, 0))
	return s
}
func (s StringInput) MinPunctuationCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinPunctuation, nil, validate.Punctuation(s.stringData, true, count))
	return s
}
func (s StringInput) IsSymbols() StringInput {
	s.input = s.input.AppendValidation(Symbols, nil, validate.Symbols(s.stringData, true, 0))
	return s
}
func (s StringInput) NoSymbols() StringInput {
	s.input = s.input.AppendValidation(NoSymbols, nil, validate.Symbols(s.stringData, false, 0))
	return s
}
func (s StringInput) MinSymbolCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinSymbols, nil, validate.Symbols(s.stringData, true, count))
	return s
}
func (s StringInput) IsWhitespaces() StringInput {
	s.input = s.input.AppendValidation(Spaces, nil, validate.Whitespaces(s.stringData, true, 0))
	return s
}
func (s StringInput) NoWhitespaces() StringInput {
	s.input = s.input.AppendValidation(NoSpaces, nil, validate.Whitespaces(s.stringData, false, 0))
	return s
}
func (s StringInput) IsControlCharacters() StringInput {
	s.input = s.input.AppendValidation(Controls, nil, validate.ControlCharacters(s.stringData, true, 0))
	return s
}
func (s StringInput) NoControlCharacters() StringInput {
	s.input = s.input.AppendValidation(NoControls, nil, validate.ControlCharacters(s.stringData, false, 0))
	return s
}
func (s StringInput) IsGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation(Graphics, nil, validate.GraphicCharacters(s.stringData, true, 0))
	return s
}
func (s StringInput) NoGraphicCharacters() StringInput {
	s.input = s.input.AppendValidation(NoGraphics, nil, validate.GraphicCharacters(s.stringData, false, 0))
	return s
}
func (s StringInput) IsMarkCharacters() StringInput {
	s.input = s.input.AppendValidation(Marks, nil, validate.MarkCharacters(s.stringData, true, 0))
	return s
}
func (s StringInput) NoMarkCharacters() StringInput {
	s.input = s.input.AppendValidation(NoMarks, nil, validate.MarkCharacters(s.stringData, false, 0))
	return s
}
func (s StringInput) IsDigits() StringInput {
	s.input = s.input.AppendValidation(Digits, nil, validate.Digits(s.stringData, true, 0))
	return s
}
func (s StringInput) NoDigits() StringInput {
	s.input = s.input.AppendValidation(NoDigits, nil, validate.Digits(s.stringData, false, 0))
	return s
}
func (s StringInput) MinDigitCount(count int) StringInput {
	s.input = s.input.AppendValidation(MinDigits, nil, validate.Digits(s.stringData, true, count))
	return s
}
