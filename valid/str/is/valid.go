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

// TODO: To use rangeMaps that exist within unicode, this needs to be mapped to
// rangeMap values. Can this exist within InputData or initialize in this within
// IfString to avoid memory usage when not being used?
type characterType int

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

func Validations() map[string]interface{} {
	var inputString, inputSubstring string
	var stringSlice []string
	var startValue, endValue int
	return map[string]interface{}{
		"isin":           IsInSlice(inputString, stringSlice),
		"notin":          NotInSlice(inputString, stringSlice),
		"required":       Required(inputString),
		"empty":          IsEmpty(inputString),
		"notempty":       NotEmpty(inputString),
		"between":        IsBetween(inputString, startValue, endValue),
		"lessthan":       IsLessThan(inputString, endValue),
		"greaterthan":    IsGreaterThan(inputString, startValue),
		"containing":     IsContaining(inputString, inputSubstring),
		"notcontaining":  NotContaining(inputString, inputSubstring),
		"regexmatch":     IsRegexMatch(inputString, inputSubstring),
		"noregexmatch":   NoRegexMatch(inputString, inputSubstring),
		"utf8":           IsUTF8(inputString),
		"noutf8":         NoUTF8(inputString),
		"uppercase":      IsUppercase(inputString),
		"nouppercase":    NoUppercase(inputString),
		"lowercase":      IsLowercase(inputString),
		"nolowercase":    NoLowercase(inputString),
		"printable":      IsPrintable(inputString),
		"noprintable":    NoPrintable(inputString),
		"alphabetic":     IsAlphabetic(inputString),
		"noalphabetic":   NoAlphabetic(inputString),
		"alphanumeric":   IsAlphanumeric(inputString),
		"noalphanumeric": NoAlphanumeric(inputString),
		"numeric":        IsNumeric(inputString),
		"nonumeric":      NoNumeric(inputString),
		"digits":         IsDigits(inputString),
		"nodigits":       NoDigits(inputString),
		"punctuation":    IsPunctuation(inputString),
		"nopunctuation":  NoPunctuation(inputString),
		"symbols":        IsSymbols(inputString),
		"nosymbols":      NoSymbols(inputString),
		"marks":          IsMarkCharacters(inputString),
		"nomarks":        NoMarkCharacters(inputString),
		"graphics":       IsGraphicCharacters(inputString),
		"nographics":     NoGraphicCharacters(inputString),
		"spaces":         IsWhitespaces(inputString),
		"nospaces":       NoWhitespaces(inputString),
	}
}

func Validation(key string) interface{} {
	return (Validations())[key]
}

//
// Check Functions
//===========================================================
// These functions can be used individually if desired

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
func NotEmpty(s string) bool              { return (s != "") }
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
func NoPrintable(s string) bool         { return StringContainsType(false, s, Printable) }
func IsPrintable(s string) bool         { return StringContainsType(true, s, Printable) }
func IsAlphabetic(s string) bool        { return StringContainsType(true, s, Alphabetic) }
func NoAlphabetic(s string) bool        { return StringContainsType(false, s, Alphabetic) }
func IsNumeric(s string) bool           { return StringContainsType(true, s, Numeric) }
func NoNumeric(s string) bool           { return StringContainsType(false, s, Numeric) }
func IsAlphanumeric(s string) bool      { return StringContainsType(true, s, Alphanumeric) }
func NoAlphanumeric(s string) bool      { return StringContainsType(false, s, Alphanumeric) }
func IsDigits(s string) bool            { return StringContainsType(true, s, Digit) }
func NoDigits(s string) bool            { return StringContainsType(false, s, Digit) }
func IsPunctuation(s string) bool       { return StringContainsType(true, s, Punctuation) }
func NoPunctuation(s string) bool       { return StringContainsType(false, s, Punctuation) }
func IsLowercase(s string) bool         { return StringContainsType(true, s, Lower) }
func NoLowercase(s string) bool         { return StringContainsType(false, s, Lower) }
func IsUppercase(s string) bool         { return StringContainsType(true, s, Upper) }
func NoUppercase(s string) bool         { return StringContainsType(false, s, Upper) }
func IsWhitespaces(s string) bool       { return StringContainsType(true, s, Space) }
func NoWhitespaces(s string) bool       { return StringContainsType(false, s, Space) }
func IsSymbols(s string) bool           { return StringContainsType(true, s, Symbol) }
func NoSymbols(s string) bool           { return StringContainsType(false, s, Symbol) }
func IsControlCharacters(s string) bool { return StringContainsType(true, s, Control) }
func NoControlCharacters(s string) bool { return StringContainsType(false, s, Control) }
func IsGraphicCharacters(s string) bool { return StringContainsType(true, s, Graphic) }
func NoGraphicCharacters(s string) bool { return StringContainsType(false, s, Graphic) }
func IsMarkCharacters(s string) bool    { return StringContainsType(true, s, Mark) }
func NoMarkCharacters(s string) bool    { return StringContainsType(false, s, Mark) }

func StringContainsType(is bool, s string, cType characterType) bool {
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
