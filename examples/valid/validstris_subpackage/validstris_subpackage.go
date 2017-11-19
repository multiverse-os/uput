package main

import (
	"fmt"
	validstris "lib/uget/valid/str/is"
)

func main() {
	fmt.Println("Validating with just generic string validations subpackage, no chaining support, only bool returned checks")
	isValid := validstris.NotEmpty("test")
	fmt.Println("validstris.NotEmpty(\"test\") returns: ", isValid)
}
