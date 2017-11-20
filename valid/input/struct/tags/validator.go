package validstruct

var ErrosListHolder ErrorsList

func Valid(s interface{}) bool {

	ErrosListHolder = ErrorsList{}

	fields := getValidationTags(s)

	if len(fields) == 0 {
		return true
	}

	//fmt.Println(">>IN ", len(ErrosListHolder))
	//fmt.Println(">>IN ", ErrosListHolder.Any())

	for _, val := range fields {
		switch filterValidationTag("tname", val.validaton) {
		case "presence":
			validatePresence(val)
		case "min", "max":
			validateMinMax(val)
		case "format":
			validateFormat(val)
		}
	}

	//fmt.Println(">>OUT ", len(ErrosListHolder))
	//fmt.Println(">>OUT ", ErrosListHolder.Any())

	if ErrosListHolder.Any() {
		return false
	}

	return true

}
