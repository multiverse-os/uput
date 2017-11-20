package validatestr

import (
	"fmt" // DEV
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

func Test_IsGreaterThan(t *testing.T) {
	if !IsGreaterThan("bigger", 3) {
		t.Errorf("!IsGreaterThan() or NotGreaterThan('bigger', 3) returned true, but should be false")
	}
	if IsGreaterThan("hi", 5) {
		t.Errorf("IsGreaterThan('hi', 5) returned true, should be false")
	}
}

func Test_Contains(t *testing.T) {
	if !Contains("bigger", "big") {
		t.Errorf("!Contains or NotContains('bigger', 'big') returned true, but should be false")
	}
	if Contains("hi", "red") {
		t.Errorf("Contains('hi', 'red') returned true, should be false")
	}
}

func Test_IsRegexMatch(t *testing.T) {
	if !IsRegexMatch("seafood", "sea.*") {
		t.Errorf("!IsRegexMatch or NotRegexMatch('bigger', '/[big]/') returned true, but should be false")
	}
	if IsRegexMatch("foodersea*", "bar.*") {
		t.Errorf("IsRegexMatch returned true, should be false")
	}
}

func Test_IsUTF8(t *testing.T) {
	if !IsUTF8("Hello, 世界") {
		t.Errorf("!IsUTF8() NotUTF8 returned true with valid UTF8 string")
	}
	if IsUTF8("\x91\x80\x80\x80") {
		t.Errorf("IsUTF8() returned true for string that is NOT UTF8.")
	}
}

func Test_Alphabetic(t *testing.T) {
	if Alphabetic("19fi-!", true, 0) {
		t.Errorf("Alphabetic('19fi-!', true, 0) failed, non-alphabetic string returned true")
	} else {
		fmt.Println("AlphabeticCount(3) for string should be 19fi-!, should be FALSE")
	}
	fmt.Println("OnlyNonAlphabetic('918jk1j', false, 0): ", Alphabetic("918jk1j", false, 0))
	if Alphabetic("918jk1j", false, 0) {
		t.Errorf("Alphabetic('918jkfj', false, 0) or OnlyNonAlphabetic() returned true, with string containing alphabetic characters")
	} else {
		fmt.Println("OnlyNONAlphabetic() for string 918jk1j, should be FALSE")
	}
	if Alphabetic("jk1222233", false, 3) {
		fmt.Println("AtleastMinNONAlphabeticCount(3) for string jk1222233, should be TRUE")
	} else {
		t.Errorf("Alphabetic('jk1222233', false, 3) or MinNonAlphabeticCount(3) returned true when given a 4 char length alphabetic string")
	}
	if Alphabetic("jkfg", true, 3) {
		fmt.Println("AtleastMinAlphabeticCount(3) for string jkfg, should be TRUE")
	} else {
		t.Errorf("Alphabetic('jkfj', true, 3) or MinAlphabeticCount(3) returned true when given a 4 char length alphabetic string")
	}
	if Alphabetic("112233", true, 2) {
		t.Errorf("Alphabetic('112233', true, 2) or MinAlphabeticCount(2) returned true when given a 4 char length NON-alphabetic string")
	}
	if Alphabetic("19fi01", false, 0) {
		t.Errorf("Alphabetic('19fi-!', false, 0) returned true, when it should be false, since its not ONLY no alphabetic")
	}
}

//func Test_Alphanumeric(t *testing.T) {
//	if Alphanumeric("19fi-!", true, 0) {
//		t.Errorf("Alphanumeric() failed, non-alphanumeric string returned true")
//	}
//	if !Alphanumeric("918jkfj", true, 0) {
//		t.Errorf("!Alphanumeric() or NoAlphanumeric() returned true when given a alphanumeric string")
//	}
//	if Alphanumeric("19fi01", false, 0) {
//		t.Errorf("Alphanumeric('19fi-!', false, 0) returned true, when false should imply no alphanumeric")
//	}
//}

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
