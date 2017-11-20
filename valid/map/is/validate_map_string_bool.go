package valid

type mapStringBoolValidator struct {
	mapValidatorEntry
	validator StringBoolValidator
}

func NewMapStringBoolValidator(root *mapValidator, key string, validator StringBoolValidator) *mapStringBoolValidator {
	return &mapStringBoolValidator{
		NewMapValidatorEntry(root, key),
		validator,
	}
}

func (v *mapStringBoolValidator) True(t string) MapStringBoolValidator {
	v.validator.True(t)
	return v
}

func (v *mapStringBoolValidator) False(f string) MapStringBoolValidator {
	v.validator.False(f)
	return v
}

func (v *mapStringBoolValidator) validate(input interface{}) (interface{}, []error) {
	s, ok := input.(string)
	if ok || input == nil {
		return v.validator.Validate(s)
	}
	return input, []error{ErrNotAString}
}
