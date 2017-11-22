package main

import (
	"lib/uput/valid/input"
)

// TODO: Should these bee constant? This will likely fail when the same thing is implemented
// in int. Maybe each set should start at different starting points.
const (
	In validinput.ValidationKey = iota
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
	Spaces
	NoSpaces
	MinSpaces
)
