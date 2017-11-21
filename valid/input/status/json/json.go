package jsonstatus

import (
	"bytes"
	"encoding/json"
	"fmt" // DEV

	validinput "lib/uput/valid/input"
	inputstatus "lib/uput/valid/input/status"
)

func PrintJSONValidationStatus(input validinput.InputData) {
	inputStatus := inputstatus.GetInputStatus(input)
	var indentedJSON bytes.Buffer
	inputJSON, _ := json.Marshal(inputStatus)
	err := json.Indent(&indentedJSON, []byte(inputJSON), "", "  ")
	if err == nil {
		fmt.Println(indentedJSON.String())
	}
}
