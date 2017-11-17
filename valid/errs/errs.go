package validerrs

import (
	"fmt" // Just fot debug
	"strconv"
	//"lib/uput/valid/errors"
)

// Development Printing (remove later, don't assume logging style)
func PrintErrors(errs []error) {
	// TODO: Obviously should just be marshalling to JSON and printing
	// but this is temporary anyways
	if len(errs) > 0 {
		fmt.Println("{")
		fmt.Println("  \"error_count\": " + strconv.Itoa(len(errs)) + ",")
		fmt.Println("  \"errors\": {")
		for _, err := range errs {
			fmt.Println("    \"string\": \"" + err.Error() + "\",")
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}
