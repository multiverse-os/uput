package inputstatus

import (
	validinput "lib/uput/valid/input"
)

type InputStatus struct {
	InputData              interface{} `json:input_data`
	InputDataType          string      `json:data_type`
	ValidationCount        int         `json:validation_count`
	ValidationDescriptions []string    `json:validations`
	ErrorCount             int         `json:error_count`
	ErrorMessages          []string    `json:errors`
}

func GetInputStatus(input validinput.InputData) InputStatus {
	return InputStatus{
		InputDataType:          input.DataTypeName,
		InputData:              input.Data,
		ValidationCount:        len(input.Validations),
		ValidationDescriptions: input.Validations,
		ErrorCount:             len(input.InputErrors()),
		ErrorMessages:          input.ErrorMessages,
	}
}
