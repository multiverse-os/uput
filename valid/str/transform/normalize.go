package transtr

import (
//"unicode"
)

// TODO: Should find way to pass in what to UTF categories are desired on otherside
// without having 100 funcs
func Normalize(s string) string {
	for _, c := range s {
		if unicode.IsSpace(c) {
			c = " "
		} else if unicode.IsControl(c) {
			c = ""
		} else if !unicode.IsPrint(c) {
			c = ""
		}
	}
	// TODO: Does this work?
	return s
}
