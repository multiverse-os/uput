package modcase

import (
	"bytes"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

// ToSnakeCase can convert all upper case characters in a string to
// underscore format.
//
// Some samples.
//     "FirstName"  => "first_name"
//     "HTTPServer" => "http_server"
//     "NoHTTPS"    => "no_https"
//     "GO_PATH"    => "go_path"
//     "GO PATH"    => "go_path"      // space is converted to underscore.
//     "GO-PATH"    => "go_path"      // hyphen is converted to underscore.
func ToSnakeCase(str string) string {
	if len(str) == 0 {
		return ""
	}
	buf := &bytes.Buffer{}
	var prev, r0, r1 rune
	var size int
	r0 = '_'
	for len(str) > 0 {
		prev = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]
		switch {
		case r0 == utf8.RuneError:
			buf.WriteByte(byte(str[0]))
		case unicode.IsUpper(r0):
			if prev != '_' {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r0))
			if len(str) == 0 {
				break
			}
			r0, size = utf8.DecodeRuneInString(str)
			str = str[size:]
			if !unicode.IsUpper(r0) {
				buf.WriteRune(r0)
				break
			}
			// find next non-upper-case character and insert `_` properly.
			// it's designed to convert `HTTPServer` to `http_server`.
			// if there are more than 2 adjacent upper case characters in a word,
			// treat them as an abbreviation plus a normal word.
			for len(str) > 0 {
				r1 = r0
				r0, size = utf8.DecodeRuneInString(str)
				str = str[size:]
				if r0 == utf8.RuneError {
					buf.WriteRune(unicode.ToLower(r1))
					buf.WriteByte(byte(str[0]))
					break
				}
				if !unicode.IsUpper(r0) {
					if r0 == '_' || r0 == ' ' || r0 == '-' {
						r0 = '_'
						buf.WriteRune(unicode.ToLower(r1))
					} else {
						buf.WriteRune('_')
						buf.WriteRune(unicode.ToLower(r1))
						buf.WriteRune(r0)
					}
					break
				}
				buf.WriteRune(unicode.ToLower(r1))
			}
			if len(str) == 0 || r0 == '_' {
				buf.WriteRune(unicode.ToLower(r0))
				break
			}
		default:
			if r0 == ' ' || r0 == '-' {
				r0 = '_'
			}
			buf.WriteRune(r0)
		}
	}
	return buf.String()
}

// SwapCase will swap characters case from upper to lower or lower to upper.
func SwapCase(str string) string {
	var r rune
	var size int
	buf := &bytes.Buffer{}
	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)
		switch {
		case unicode.IsUpper(r):
			buf.WriteRune(unicode.ToLower(r))
		case unicode.IsLower(r):
			buf.WriteRune(unicode.ToUpper(r))
		default:
			buf.WriteRune(r)
		}
		str = str[size:]
	}
	return buf.String()
}
