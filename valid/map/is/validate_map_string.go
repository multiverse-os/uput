package valid

type mapStringValidator struct {
	mapValidatorEntry
	validator StringValidator
}

func NewMapStringValidator(root *mapValidator, key string) *mapStringValidator {
	return &mapStringValidator{
		NewMapValidatorEntry(root, key),
		NewStringValidator(),
	}
}

func (v *mapStringValidator) TrimSpace() MapStringValidator {
	v.validator = v.validator.TrimSpace()
	return v
}

func (v *mapStringValidator) NotEmpty() MapStringValidator {
	v.validator = v.validator.NotEmpty()
	return v
}

func (v *mapStringValidator) Regex(regex string) MapStringValidator {
	v.validator = v.validator.Regex(regex)
	return v
}

func (v *mapStringValidator) Func(f StringValidatorFunc) MapStringValidator {
	v.validator = v.validator.Func(f)
	return v
}

func (v *mapStringValidator) AsInt() MapStringIntValidator {
	validator := NewMapStringIntValidator(v.mapValidatorEntry, v.validator.AsInt())
	validator.root.validators[v.key] = validator
	return validator
}

func (v *mapStringValidator) AsBool() MapStringBoolValidator {
	validator := NewMapStringBoolValidator(v.root, v.key, v.validator.AsBool())
	validator.root.validators[v.key] = validator
	return validator
}

func (v *mapStringValidator) Required() MapStringValidator {
	v.required = true
	return v
}

func (v *mapStringValidator) Default(defaut interface{}) MapStringValidator {
	v.defaut = defaut
	return v
}

func (v *mapStringValidator) validate(input interface{}) (interface{}, []error) {
	s, ok := input.(string)
	if ok {
		return v.validator.Validate(s)
	}
	return input, []error{ErrNotAString}
}
