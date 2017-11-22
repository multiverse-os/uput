package validstatus

import (
	"encoding/json"

	validinput "lib/uput/valid/input"
)

func jsonStatus(status validinput.InputStatus, indention string) (string, err error) {
	outputBytes, err := json.Marshal(status)
	if err == nil {
		err = json.Indent(&indentedJSON, []byte(inputJSON), "", indention)
		if err == nil {
			return string(outputBytes), nil
		}
	}
	return nil, err
}
