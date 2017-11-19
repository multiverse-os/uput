package main

import (
	"fmt"
	validstr "lib/uput/validstr"
)

func main() {
	fmt.Println("Validating with string datatype specific subpackage, which relies only on the generic string validations subpackage.")
	isValid, validatedInput, errs := validstr.If("testvalue").NotEmpty().IsBetween(2, 64).IsContaining("test").IsValid()
	if isValid {
		fmt.Println("Validated string: ", validatedInput)
	} else {
		fmt.Println("Failed to validate, ", len(errs), " number of validation errors")
	}
}
