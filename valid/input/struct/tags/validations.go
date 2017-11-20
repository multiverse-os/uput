package validstruct

import (
	"log"
	"strconv"
)

type validationData struct {
	fieldName string
	validaton string
	value     interface{}
}

func validatePresence(data validationData) {
	if data.value == "" {
		ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Can not be blank"})
		return
	}
}

func validateFormat(data validationData) {

	tv := filterValidationTag("tvalue", data.validaton)

	switch tv.(string) {
	case "email":
		match := emailRegex.MatchString(data.value.(string))
		if !match {
			ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Format must match"})
			return
		}
	}
}

func validateMinMax(data validationData) {

	vl := filterValidationTag("tvalue", data.validaton)
	refValue, err := strconv.Atoi(vl.(string))
	if err != nil {
		ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must be a number"})
		log.Println(ErrosListHolder)
		return
	}

	switch filterValidationTag("tname", data.validaton) {
	case "min":
		if data.value.(int) < refValue {
			ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must be greater than " + vl.(string)})
			log.Println(ErrosListHolder)
			return
		}
	case "max":
		if data.value.(int) > refValue {
			ErrosListHolder.AppendError(ErrorMessage{data.fieldName, "Must be less than " + vl.(string)})
			log.Println(ErrosListHolder)
			return
		}
	}

}
