package valid

import (
//"lib/uput/valid/int"
)

var intErrMessages = map[string]string{
	"isequal":          "is not equal",
	"notequal":         "is equal",
	"isin":             "not included in",
	"iszero":           "is not zero",
	"notzero":          "is zero",
	"isodd":            "is not odd",
	"notodd":           "is odd",
	"iseven":           "is not even",
	"noteven":          "is even",
	"ispositive":       "not positive",
	"notpositive":      "not negative",
	"isnegative":       "not negative",
	"notnegative":      "not positive",
	"greaterthan":      "greater than",
	"minimum":          "below minimum",
	"greaterthanequal": "greater than or equal to",
	"lessthan":         "less than",
	"maximum":          "above maximum",
	"lessthanequal":    "less than or equal to",
	"isbetween":        "is between",
}
