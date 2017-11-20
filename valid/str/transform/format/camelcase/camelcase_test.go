package camelcase_test

import (
	"testing"
)

func Test_CamelCase(t *testing.T) {

	tests := []struct {
		in, out string
	}{
		{in: "fizz buzz bazz", out: "FizzBuzzBazz"},
		{in: "fizz_buzz_bazz", out: "FizzBuzzBazz"},
		{in: "fizz-buzz-bazz", out: "FizzBuzzBazz"},
		{in: "Fizz Buzz Bazz", out: "FizzBuzzBazz"},
		{in: "São Ñino França Ávido Caça", out: "SaoNinoFrancaAvidoCaca"},
	}

	for _, test := range tests {
		cc := CamelCase(test.in)
		if cc != test.out {
			t.Errorf("\nExpected: %s\nGot:      %s", test.out, cc)
		}
	}
}
