package valid

import (
//"lib/uput/valid/map" validmap
)

//
// Maps will contain one of the already defined types
// so only define messages specific to map
var mapErrMessages = map[string]string{
	"empty":            "is not empty",
	"notempty":         "is empty",
	"keyexists":        "key does not exist",
	"between":          "length not between",
	"minimum":          "length below minimum",
	"maximum":          "length above maximum",
	"greaterthan":      "length greater than",
	"greaterthanequal": "length greater than or equal to",
	"lessthan":         "length less than",
	"lessthanequal":    "length less than or equal to",
}

// TODO: Add ErrorMessages for whatever the value type is
