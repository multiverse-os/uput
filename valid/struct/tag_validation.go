package valid

import (
	"errors"
	"reflect"
)

type InputStruct struct {
	// Struct
	dataField reflect.StructField
}

//var FieldValidations = map[string]structValidMap{
//	// Validate Existence
//	"nil":      valid.Nil,
//	"required": validators.Nil,
//	// Validate Numericality
//	"min":     valid.Min,
//	"max":     valid.Max,
//	"between": valid.Max,
//	// Validate Length
//	"minLength":     valid.MinLen,
//	"maxLength":     valid.MaxLen,
//	"lengthBetween": valid.MaxLen,
//	// Validate Type
//	"string":  valid.MaxLen,
//	"boolean": valid.MaxLen,
//	"int":     valid.MaxLen,
//	"int8":    valid.MaxLen,
//	"int32":   valid.MaxLen,
//	"int64":   valid.MaxLen,
//	"uint":    valid.MaxLen,
//	"uint8":   valid.MaxLen,
//	"uint32":  valid.MaxLen,
//	"uint64":  valid.MaxLen,
//}

//func (inputStruct InputStruct) validateField(field reflect.StructField, v reflect.Value) error {
//	if required, ok := inputStruct.dataField.Tag.Lookup("required"); ok {
//		if required == "true" {
//			value := v.FieldByName(field.Name)
//			isOk := false
//
//			switch inputStruct.dataField.Type.Name() {
//			case "string":
//				isOk = value.String() != ""
//				break
//			case "boolean":
//				isOk = value.Bool()
//				break
//			}
//
//			if !isOk {
//				return errors.New("%s is required" + inputStruct.dataField.Name)
//			}
//		}
//	}
//
//	return nil
//}
