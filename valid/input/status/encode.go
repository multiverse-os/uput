package validstatus

import (
	"unicode"

	validinput "lib/uput/valid/input"
)

type encodeOption int

const (
	Format encodeOption = iota
	Indent
)

func Status(status validinput.InputStatus, options map[encodeOption]string) string {
	var exists bool
	switch options[Format] {
	case "json" || "JSON":
		indention, exists := options[Idention]
		if !exists {
			indention = ""
		} else {
			for _, c := range indention {
				if !unicode.IsSpace(c) {
					indention = ""
					break
				}
			}
		}
		return jsonStatus(status, indention)
	}
	// TODO: Add YAML encoding of status
	//case "yaml" || "YAML":
	// TODO: Add XML encoding of status
	//case "xml" || "XML":
}

func Input(input validinput.InputData, options map[string]string) string {

}
