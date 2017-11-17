package validstr

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Catching look-alikes
// string Normalization and UTF8 comparisons
//
// Can you tell the difference between 'K' ("\u004B") and 'K' (Kelvin sign "\u212A")?
// It is easy to overlook the sometimes minute differences between variants of the same
// underlying character. It is generally a good idea to disallow such variants in
// *identifiers* (Usernames, Roles, etc) or anything where deceiving users with such
// look-alikes can pose a security hazard.
// TODO: Add this to the _test.go file to ensure this is always being checked for

type characterType int

// TODO: To use rangeMaps that exist within unicode, this needs to be mapped to
// rangeMap values. Can this exist within InputData or initialize in this within
// IfString to avoid memory usage when not being used?
const (
	Alphabetic characterType = iota
	Numeric
	Alphanumeric
	Digit
	Printable
	Punctuation
	Lower
	Upper
	Space
	Symbol
	Control
	Graphic
	Mark
)

// Slice
func IsInSlice(s string, lo []string) bool {
	for _, option := range lo {
		if option == s {
			return true
		}
	}
	return false
}
func NotInSlice(s string, lo []string) bool {
	for _, option := range lo {
		if option == s {
			return false
		}
	}
	return true
}

// String Length
func Required(s string) bool              { return (len(s) > 0) }
func IsEmpty(s string) bool               { return (s == "") }
func IsNotEmpty(s string) bool            { return (s != "") }
func IsBetween(s string, gt, lt int) bool { return (len(s) > gt || len(s) < lt) }
func IsLessThan(s string, lt int) bool    { return (len(s) < lt) }
func IsGreaterThan(s string, gt int) bool { return (len(s) > gt) }

// string.Contains
func IsContaining(s, ss string) bool  { return strings.Contains(s, ss) }
func NotContaining(s, ss string) bool { return !strings.Contains(s, ss) }

// regexp.MatchString
func IsRegexMatch(s, pattern string) (match bool) {
	match, _ = regexp.MatchString(pattern, s)
	return match
}
func NoRegexMatch(s, pattern string) (match bool) {
	match, _ = regexp.MatchString(pattern, s)
	return match
}

// UTF Rune Validations
func IsUTF8(s string) bool              { return utf8.ValidString(s) }
func NoUTF8(s string) bool              { return !utf8.ValidString(s) }
func IsPrintable(s string) bool         { return IsStringType(true, s, Printable) }
func NoPrintable(s string) bool         { return IsStringType(false, s, Printable) }
func IsAlphabetic(s string) bool        { return IsStringType(true, s, Alphabetic) }
func NoAlphabetic(s string) bool        { return IsStringType(false, s, Alphabetic) }
func IsNumeric(s string) bool           { return IsStringType(true, s, Numeric) }
func NoNumeric(s string) bool           { return IsStringType(false, s, Numeric) }
func IsAlphaNumeric(s string) bool      { return IsStringType(true, s, Alphanumeric) }
func NoAlphaNumeric(s string) bool      { return IsStringType(false, s, Alphanumeric) }
func IsDigits(s string) bool            { return IsStringType(true, s, Digit) }
func NoDigits(s string) bool            { return IsStringType(false, s, Digit) }
func IsPunctuation(s string) bool       { return IsStringType(true, s, Punctuation) }
func NoPunctuation(s string) bool       { return IsStringType(false, s, Punctuation) }
func IsLowercase(s string) bool         { return IsStringType(true, s, Lower) }
func NoLowercase(s string) bool         { return IsStringType(false, s, Lower) }
func IsUppercase(s string) bool         { return IsStringType(true, s, Upper) }
func NoUppercase(s string) bool         { return IsStringType(false, s, Upper) }
func IsWhitespaces(s string) bool       { return IsStringType(true, s, Space) }
func NoWhitespaces(s string) bool       { return IsStringType(false, s, Space) }
func IsSymbols(s string) bool           { return IsStringType(true, s, Symbol) }
func NoSymbols(s string) bool           { return IsStringType(false, s, Symbol) }
func IsControlCharacters(s string) bool { return IsStringType(true, s, Control) }
func NoControlCharacters(s string) bool { return IsStringType(false, s, Control) }
func IsGraphicCharacters(s string) bool { return IsStringType(true, s, Graphic) }
func NoGraphicCharacters(s string) bool { return IsStringType(false, s, Graphic) }
func IsMarkCharacters(s string) bool    { return IsStringType(true, s, Mark) }
func NoMarkCharacters(s string) bool    { return IsStringType(false, s, Mark) }

func IsStringType(is bool, s string, cType characterType) bool {
	// TODO: Id prefer to switch to a system that uses Is(rangeMap) rangeMap, so
	// a broader one that accepts []rangeMap to let developers choose whatever combination
	if cType == Alphabetic {
		for _, c := range s {
			if is && !unicode.IsLetter(c) {
				return false
			} else if !is && unicode.IsLetter(c) {
				return false
			}
		}
	} else if cType == Alphanumeric {
		for _, c := range s {
			if is && !unicode.IsLetter(c) && !unicode.IsNumber(c) {
				return false
			} else if !is && (unicode.IsLetter(c) || unicode.IsNumber(c)) {
				return false
			}
		}
	} else if cType == Numeric {
		for _, c := range s {
			if is && !unicode.IsNumber(c) {
				return false
			} else if !is && unicode.IsNumber(c) {
				return false
			}
		}
	} else if cType == Punctuation {
		for _, c := range s {
			if is && !unicode.IsPunct(c) {
				return false
			} else if !is && unicode.IsPunct(c) {
				return false
			}
		}
	} else if cType == Lower {
		for _, c := range s {
			if is && !unicode.IsLower(c) {
				return false
			} else if !is && unicode.IsLower(c) {
				return false
			}
		}
	} else if cType == Upper {
		for _, c := range s {
			if is && !unicode.IsUpper(c) {
				return false
			} else if !is && unicode.IsUpper(c) {
				return false
			}
		}
	} else if cType == Printable {
		for _, c := range s {
			if is && !unicode.IsPrint(c) {
				return false
			} else if !is && unicode.IsPrint(c) {
				return false
			}
		}
	} else if cType == Space {
		for _, c := range s {
			if is && !unicode.IsSpace(c) {
				return false
			} else if !is && unicode.IsSpace(c) {
				return false
			}
		}
	} else if cType == Symbol {
		for _, c := range s {
			if is && !unicode.IsSymbol(c) {
				return false
			} else if !is && unicode.IsSymbol(c) {
				return false
			}
		}
	} else if cType == Control {
		for _, c := range s {
			if is && !unicode.IsControl(c) {
				return false
			} else if !is && unicode.IsControl(c) {
				return false
			}
		}
	} else if cType == Graphic {
		for _, c := range s {
			if is && !unicode.IsGraphic(c) {
				return false
			} else if !is && unicode.IsGraphic(c) {
				return false
			}
		}
	} else if cType == Mark {
		for _, c := range s {
			if is && !unicode.IsMark(c) {
				return false
			} else if !is && unicode.IsMark(c) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
