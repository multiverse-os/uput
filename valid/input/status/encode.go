package validstatus

import (
	"encoding/json"
	"unicode"
)

type EncodeOption int

const (
	Format EncodeOption = iota
	Indent
)

func (s InputStatus) Encode(options map[EncodeOption]string) (string, error) {
	switch options[Format] {
	case "json":
		indent, exists := options[Indent]
		if !exists {
			indent = ""
		} else {
			for _, c := range indent {
				if !unicode.IsSpace(c) {
					indent = ""
					break
				}
			}
		}
		return s.encodeJSON(indent)
	default:
		return "", nil
	}
	// TODO: Add YAML encoding of status
	//case "yaml" || "YAML":
	// TODO: Add XML encoding of status
	//case "xml" || "XML":
}

func (s InputStatus) encodeJSON(indent string) (string, error) {
	output, err := json.MarshalIndent(s, "", indent)
	if err == nil {
		//err = json.Indent(&output, []byte(inputJSON), "", indent)
		//if err == nil {
		return string(output), err
		//}
	}
	return "", err
}
