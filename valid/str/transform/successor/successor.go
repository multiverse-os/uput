package successor

import (
	"bytes"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

// Successor returns the successor to string.
//
// If there is one alphanumeric rune is found in string, increase the rune by 1.
// If increment generates a "carry", the rune to the left of it is incremented.
// This process repeats until there is no carry, adding an additional rune if necessary.
//
// If there is no alphanumeric rune, the rightmost rune will be increased by 1
// regardless whether the result is a valid rune or not.
//
// Only following characters are alphanumeric.
//     * a - z
//     * A - Z
//     * 0 - 9
//
// Samples (borrowed from ruby's String#succ document):
//     "abcd"      => "abce"
//     "THX1138"   => "THX1139"
//     "<<koala>>" => "<<koalb>>"
//     "1999zzz"   => "2000aaa"
//     "ZZZ9999"   => "AAAA0000"
//     "***"       => "**+"
func Successor(str string) string {
	if str == "" {
		return str
	}
	var r rune
	var i int
	carry := ' '
	runes := []rune(str)
	l := len(runes)
	lastAlphanumeric := l
	for i = l - 1; i >= 0; i-- {
		r = runes[i]
		if ('a' <= r && r <= 'y') ||
			('A' <= r && r <= 'Y') ||
			('0' <= r && r <= '8') {
			runes[i]++
			carry = ' '
			lastAlphanumeric = i
			break
		}
		switch r {
		case 'z':
			runes[i] = 'a'
			carry = 'a'
			lastAlphanumeric = i
		case 'Z':
			runes[i] = 'A'
			carry = 'A'
			lastAlphanumeric = i
		case '9':
			runes[i] = '0'
			carry = '0'
			lastAlphanumeric = i
		}
	}
	// Needs to add one character for carry.
	if i < 0 && carry != ' ' {
		buf := &bytes.Buffer{}
		buf.Grow(l + 4) // Reserve enough space for write.
		if lastAlphanumeric != 0 {
			buf.WriteString(str[:lastAlphanumeric])
		}
		buf.WriteRune(carry)
		for _, r = range runes[lastAlphanumeric:] {
			buf.WriteRune(r)
		}
		return buf.String()
	}
	// No alphanumeric character. Simply increase last rune's value.
	if lastAlphanumeric == l {
		runes[l-1]++
	}
	return string(runes)
}
