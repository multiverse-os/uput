package normalize

import (
	"unicode"
)

func DeleteWhiteSpace(str string) string {
	if str == "" {
		return str
	}
	sz := len(str)
	var chs bytes.Buffer
	count := 0
	for i := 0; i < sz; i++ {
		ch := rune(str[i])
		if !unicode.IsSpace(ch) {
			chs.WriteRune(ch)
			count++
		}
	}
	if count == sz {
		return str
	}
	return chs.String()
}
