package shuffle

import (
	"bytes"
	"math/rand"
	"unicode"
	"unicode/utf8"
)

// Shuffle randomizes runes in a string and returns the result.
// It uses default random source in `math/rand`.
func Shuffle(str string) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	index := 0
	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1)
		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}
	return string(runes)
}

// ShuffleSource randomizes runes in a string with given random source.
func ShuffleSource(str string, src rand.Source) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	index := 0
	r := rand.New(src)
	for i := len(runes) - 1; i > 0; i-- {
		index = r.Intn(i + 1)
		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}
	return string(runes)
}
