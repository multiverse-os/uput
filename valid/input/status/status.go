package validstatus

import (
	"lib/uput/valid/input"
)

type InputStatus struct {
	InputData              interface{} `json:"input_data" yaml:"input_data" xml:"input_data,attr"`
	DataType               string      `json:"data_type" yaml:"input_data" xml:"data_type,attr"`
	ValidationCount        int         `json:"validation_count" yaml:"validation_count" xml:"validation_count,attr"`
	ValidationDescriptions []string    `json:"validations" yaml:"validations" xml:"validations,attr"`
	ErrorCount             int         `json:"error_count" yaml:"error_count" xml:"error_count,attr"`
	ErrorMessages          []string    `json:"errors" yaml:"errors" xml:"errors,attr"`
}

func ValidationStatus(input validinput.InputData) InputStatus {
	return InputStatus{
		InputDataType:          input.DataType.String(),
		InputData:              input.Data,
		ValidationCount:        len(input.Validations),
		ValidationDescriptions: (input.ValidationDescriptions()),
		ErrorCount:             len(input.InputErrors()),
		ErrorMessages:          (input.ErrorMessages()),
	}
}
