package validstr

import (
	validinput "lib/uput/valid/input"
)

const (
	_                           = iota       // Ignore 0
	In validinput.ValidationKey = iota + 128 // 4 << 0 = 00010000
	NotIn
	Required
	Empty
	NotEmpty
	Blank
	NotBlank
	Between
	LessThan
	GreaterThan
	Contains
	NotContaining
	RegexMatch
	NoRegexMatch
	UTF8
	NoUTF8
	Uppercase
	NoUppercase
	MinUppercase
	Lowercase
	NoLowercase
	MinLowercase
	Printable
	NoPrintable
	Alphabetic
	NoAlphabetic
	MinAlphabetic
	Alphanumeric
	NoAlphanumeric
	MinAlphanumeric
	Numeric
	NoNumeric
	MinNumeric
	Digits
	NoDigits
	MinDigits
	Punctuation
	NoPunctuation
	MinPunctuation
	Symbols
	NoSymbols
	MinSymbols
	Marks
	NoMarks
	Graphics
	NoGraphics
	Controls
	NoControls
	Spaces
	NoSpaces
	MinSpaces
)

func StringToValidationKey(keyString string) validinput.ValidationKey {
	return (ValidationKeyNames())[keyString]
}

func ValidationKeyNames() map[string]validinput.ValidationKey {
	return map[string]validinput.ValidationKey{
		"in":              In,
		"notin":           NotIn,
		"required":        Required,
		"empty":           Empty,
		"notempty":        NotEmpty,
		"blank":           Blank,
		"notblank":        NotBlank,
		"between":         Between,
		"lessthan":        LessThan,
		"greaterthan":     GreaterThan,
		"contains":        Contains,
		"notcontaining":   NotContaining,
		"regexmatch":      RegexMatch,
		"noregexmatch":    NoRegexMatch,
		"utf8":            UTF8,
		"noutf8":          NoUTF8,
		"uppercase":       Uppercase,
		"nouppercase":     NoUppercase,
		"minuppercase":    MinUppercase,
		"lowercase":       Lowercase,
		"nolowercase":     NoLowercase,
		"minlowercase":    MinLowercase,
		"printable":       Printable,
		"noprintable":     NoPrintable,
		"alphabetic":      Alphabetic,
		"noalphabetic":    NoAlphabetic,
		"minalphabetic":   MinAlphabetic,
		"alphanumeric":    Alphanumeric,
		"noalphanumeric":  NoAlphanumeric,
		"minalphanumeric": MinAlphanumeric,
		"numeric":         Numeric,
		"nonumeric":       NoNumeric,
		"minnumeric":      MinNumeric,
		"digits":          Digits,
		"nodigits":        NoDigits,
		"mindigits":       MinDigits,
		"punctuation":     Punctuation,
		"nopunctuation":   NoPunctuation,
		"minpunctuation":  MinPunctuation,
		"symbols":         Symbols,
		"nosymbols":       NoSymbols,
		"minsymbols":      MinSymbols,
		"marks":           Marks,
		"nomarks":         NoMarks,
		"graphics":        Graphics,
		"nographics":      NoGraphics,
		"controls":        Controls,
		"nocontrols":      NoControls,
		"spaces":          Spaces,
		"nospaces":        NoSpaces,
		"minspaces":       MinSpaces,
	}
}
