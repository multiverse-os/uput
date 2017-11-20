package valid

import (
//"lib/uinput/valid/time"
)

var timeErrorMessages = map[string]string{
	"isbefore":   "not before",
	"notbefore":  "is before",
	"isafter":    "not after",
	"notafter":   "is after",
	"isbetween":  "not between",
	"notbetween": "is between",
	"ispast":     "not in past",
	"notpast":    "in past",
	"isfuture":   "not in future",
	"notfuture":  "in future",
}
