package validatestr

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

//
// 'Valid String' Subpackage
//=================================================================
// This subpackage provides the boolean checks for the validations
// it can be used individually without the the rest of the more
// complex validation system or calling in validstr or valid will
// use this, so only import if you are using only this subpackage.

// Catching look-alikes
// string Normalization and UTF8 comparisons
//
// Can you tell the difference between 'K' ("\u004B") and 'K' (Kelvin sign "\u212A")?
// It is easy to overlook the sometimes minute differences between variants of the same
// underlying character. It is generally a good idea to disallow such variants in
// *identifiers* (Usernames, Roles, etc) or anything where deceiving users with such
// look-alikes can pose a security hazard.
// TODO: Add this to the _test.go file to ensure this is always being checked for
type runeType int

const (
	IsLetter = iota
	IsNumber
	IsLetterOrNumber
	IsPrint
	IsPunct
	IsLower
	IsUpper
	IsSpace
	IsSymbol
	IsControl
	IsGraphic
	IsMark
	IsDigit
)

//
// Check Functions
//==================================================================================
// Slice
func IsInSlice(s string, lo []string) bool {
	for _, option := range lo {
		if option == s {
			return true
		}
	}
	return false
}

//
// String Length
func IsEmpty(s string) bool                   { return len(s) == 0 }
func IsNotEmpty(s string) bool                { return !IsEmpty(s) }
func IsBlank(s string) bool                   { return len(strings.TrimSpace(s)) == 0 }
func IsNotBlank(text string) bool             { return !IsBlank(text) }
func IsBetween(s string, start, end int) bool { return (len(s) >= start && len(s) <= end) }
func IsLessThan(s string, lt int) bool        { return (len(s) < lt) }
func IsGreaterThan(s string, gt int) bool     { return (len(s) > gt) }

//
// strings.Contains substring value
func Contains(s, ss string) bool { return strings.Contains(s, ss) }

//
// regexp.Match pattern value
func IsRegexMatch(s, pattern string) (match bool) {
	match, _ = regexp.MatchString(pattern, s)
	return match
}

//
// UTF8 Validation
func IsUTF8(s string) bool { return utf8.ValidString(s) }

//
// UTF Rune Validations
func Alphabetic(s string, is bool, count int) bool { return isType(s, IsLetter, is, count) }
func Alphanumeric(s string, is bool, count int) bool {
	return isType(s, IsLetterOrNumber, is, count)
}
func Numeric(s string, is bool, count int) bool           { return isType(s, IsNumber, is, count) }
func Punctuation(s string, is bool, count int) bool       { return isType(s, IsPunct, is, count) }
func Lowercase(s string, is bool, count int) bool         { return isType(s, IsLower, is, count) }
func Uppercase(s string, is bool, count int) bool         { return isType(s, IsUpper, is, count) }
func Printable(s string, is bool, count int) bool         { return isType(s, IsPrint, is, count) }
func Whitespaces(s string, is bool, count int) bool       { return isType(s, IsSpace, is, count) }
func Symbols(s string, is bool, count int) bool           { return isType(s, IsSymbol, is, count) }
func ControlCharacters(s string, is bool, count int) bool { return isType(s, IsControl, is, count) }
func GraphicCharacters(s string, is bool, count int) bool { return isType(s, IsGraphic, is, count) }
func MarkCharacters(s string, is bool, count int) bool    { return isType(s, IsMark, is, count) }
func Digits(s string, is bool, count int) bool            { return isType(s, IsDigit, is, count) }

func runeOfType(r rune, rType runeType) bool {
	switch rType {
	case IsLetter:
		return unicode.IsLetter(r)
	case IsNumber:
		return unicode.IsNumber(r)
	case IsLetterOrNumber:
		return (unicode.IsLetter(r) || unicode.IsNumber(r))
	case IsPunct:
		return unicode.IsPunct(r)
	case IsUpper:
		return unicode.IsUpper(r)
	case IsLower:
		return unicode.IsLower(r)
	case IsPrint:
		return unicode.IsPrint(r)
	case IsSpace:
		return unicode.IsSpace(r)
	case IsSymbol:
		return unicode.IsSymbol(r)
	case IsControl:
		return unicode.IsControl(r)
	case IsGraphic:
		return unicode.IsGraphic(r)
	case IsMark:
		return unicode.IsMark(r)
	case IsDigit:
		return unicode.IsDigit(r)
	default:
		return false
	}
}

//
// Reference
//====================================================
//
//For '\b':
//	is control rune
//	is not printable rune
//For '5':
//	is digit rune
//	is graphic rune
//	is number rune
//	is printable rune
//For 'Ὂ':
//	is graphic rune
//	is letter rune
//	is printable rune
//	is upper case rune
//For 'g':
//	is graphic rune
//	is letter rune
//	is lower case rune
//	is printable rune
//For '̀':
//	is graphic rune
//	is mark rune
//	is printable rune
//For '9':
//	is digit rune
//	is graphic rune
//	is number rune
//	is printable rune
//For '!':
//	is graphic rune
//	is printable rune
//	is punct rune
//For ' ':
//	is graphic rune
//	is printable rune
//	is space rune
//For '℃':
//	is graphic rune
//	is printable rune
//	is symbol rune
//For 'ᾭ':
//	is graphic rune
//	is letter rune
//	is printable rune
//	is title case rune
//For 'G':
//	is graphic rune
//	is letter rune
//	is printable rune
//	is upper case rune

func isType(s string, rType runeType, is bool, count int) bool {
	typeCount := 0
	if count == 0 {
		count = len(s)
	}
	for _, r := range s {
		if is == runeOfType(r, rType) {
			if typeCount < count {
				typeCount++
			} else if typeCount == count {
				return true
			}
		}
	}
	return (typeCount >= count)
}
