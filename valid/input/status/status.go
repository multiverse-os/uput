package validstatus

import (
	"lib/uput/valid/input"
)

type InputStatus struct {
	Data                   interface{} `json:"input_data" yaml:"input_data" xml:"input_data,attr"`
	Kind                   string      `json:"data_type" yaml:"input_data" xml:"data_type,attr"`
	ValidationCount        int         `json:"validation_count,omitempty" yaml:"validation_count,omitempty" xml:"validation_count,attr,omitempty"`
	ValidationDescriptions []string    `json:"validations,omitempty" yaml:"validations,omitempty" xml:"validations,attr,omitempty" `
	ErrorCount             int         `json:"error_count" yaml:"error_count" xml:"error_count,attr"`
	ErrorMessages          []string    `json:"errors" yaml:"errors" xml:"errors,attr"`
}

func GetStatus(input validinput.InputData, onlyErrors bool) (status InputStatus) {
	status = InputStatus{
		Data:          input.Data,
		Kind:          input.Kind.String(),
		ErrorCount:    len(input.InputErrors()),
		ErrorMessages: (input.ErrorMessages()),
	}
	if !onlyErrors {
		status.ValidationCount = len(input.Validations)
		status.ValidationDescriptions = (input.ValidationDescriptions())
	}
	return status
}
