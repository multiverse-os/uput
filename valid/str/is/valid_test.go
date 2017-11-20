package validatestr

import (
	"fmt"
	"testing"
)

func Test_IsInSlice(t *testing.T) {
	if !IsInSlice("string", []string{"string", "sub", "test"}) {
		t.Errorf("IsInSlice(): failed to find value that was in []string.")
	}
	if IsInSlice("string", []string{"clam", "john", "test"}) {
		t.Errorf("IsInSlice() failed by returning false for value not in []string.")
	}
}

func Test_NotEmpty(t *testing.T) {
	if NotEmpty("") {
		t.Errorf("NotEmpty() returned false for value that was empty")
	}
	if !NotEmpty("mario test") {
		t.Errorf("!NotEmpty() or IsEmpty returned true for value that was not empty")
	}
}

func Test_IsBetween(t *testing.T) {
	if !IsBetween("broiler", 3, 12) {
		t.Errorf("!IsBetween() or NotBetween(3, 12) returned TRUE for value string(broiler) with length that IS between 3-12. Should be false")
	}
	if IsBetween("okay", 10, 15) {
		fmt.Println("okay has length between 10-15:", IsBetween("okay", 10, 15))
		t.Errorf("IsBetween() returned true for value string(okay) with length that was NOT between 10-15")
	}
}

func Test_IsLessThan(t *testing.T) {
	if IsLessThan("bigger", 3) {
		t.Errorf("IsLessThan() returned true for string with length not less than 3")
	}
	if !IsLessThan("hi", 5) {
		t.Errorf("!IsLessThan() or NotLessThan 5 with string(hi) returned true")
	}
}

//func IsGreaterThan(s string, gt int) bool { return (len(s) > gt) }
//
////
//// strings.Contains substring value
//func Contains(s, ss string) bool { return strings.Contains(s, ss) }
//
////
//// regexp.Match pattern value
//func IsRegexMatch(s, pattern string) (match bool) {
//	match, _ = regexp.MatchString(pattern, s)
//	return match
//}
//
////
//// UTF8 Validation
//func IsUTF8(s string) bool { return utf8.ValidString(s) }
//
////
//// UTF Rune Validations
//func Alphabetic(s string, is bool, count uint8) bool { return isType(s, IsLetter, is, count) }
//func Alphanumeric(s string, is bool, count uint8) bool {
//	return isType(s, IsLetterOrNumber, is, count)
//}
//func Numeric(s string, is bool, count uint8) bool           { return isType(s, IsNumber, is, count) }
//func Punctuation(s string, is bool, count uint8) bool       { return isType(s, IsPunct, is, count) }
//func Lowercase(s string, is bool, count uint8) bool         { return isType(s, IsLower, is, count) }
//func Uppercase(s string, is bool, count uint8) bool         { return isType(s, IsUpper, is, count) }
//func Printable(s string, is bool, count uint8) bool         { return isType(s, IsPrint, is, count) }
//func Whitespaces(s string, is bool, count uint8) bool       { return isType(s, IsSpace, is, count) }
//func Symbols(s string, is bool, count uint8) bool           { return isType(s, IsSymbol, is, count) }
//func ControlCharacters(s string, is bool, count uint8) bool { return isType(s, IsControl, is, count) }
//func GraphicCharacters(s string, is bool, count uint8) bool { return isType(s, IsGraphic, is, count) }
//func MarkCharacters(s string, is bool, count uint8) bool    { return isType(s, IsMark, is, count) }
//func Digits(s string, is bool, count uint8) bool            { return isType(s, IsDigit, is, count) }
