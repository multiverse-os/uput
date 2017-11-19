package main

import (
	"fmt"
	"lib/uput/valid"
)

func main() {
	fmt.Println("Validating with base valid package, includes all validation data type subpackages...")
	isValid, validatedInput, errs := valid.IfString("testvalue").NotEmpty().IsBetween(2, 64).IsContaining("test").IsValid()
	if isValid {
		fmt.Println("Validated string: ", validatedInput)
	} else {
		fmt.Println("Failed to validate, ", len(errs), " number of validation errors")
	}
}
