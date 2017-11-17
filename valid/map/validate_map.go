package valid

var ignoreValue struct{} // Dummy value to indicate do not use the returned value

type mapValidator struct {
	validators    map[string]genericMapValidator
	failOnUnknown bool
}

type mapValidatorEntry struct {
	root     *mapValidator
	key      string
	required bool
	defaut   interface{}
}

func NewMapValidatorEntry(root *mapValidator, key string) mapValidatorEntry {
	return mapValidatorEntry{root, key, false, ignoreValue}
}

func NewMapValidator() MapValidator {
	return &mapValidator{
		map[string]genericMapValidator{},
		false,
	}
}

func (v *mapValidator) FailOnUnknown() MapValidator {
	v.failOnUnknown = true
	return v
}

func (v *mapValidator) Key(key string) MapStringValidator {
	validator := NewMapStringValidator(v, key)
	v.validators[key] = validator
	return validator
}

func (v *mapValidator) IntKey(key string) MapIntValidator {
	validator := NewMapIntValidator(v, key)
	v.validators[key] = validator
	return validator
}

func (v *mapValidator) Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error) {

	// TODO check that all the default values actually validate, and panic on first run if they don't.

	errs := make(map[string][]error)

	// Find missing keys, so we can fail/set defaults
	for key, validator := range v.validators {
		if _, found := input[key]; !found {
			value, new_errs := validator.validateMissing()
			if value != ignoreValue {
				input[key] = value
			}
			if new_errs != nil {
				errs[key] = append(errs[key], new_errs...)
			}
		}
	}

	// Now validate all values
	for key, value := range input {
		validator, found := v.validators[key]
		if found {
			value, new_errs := validator.validate(value)
			if value != ignoreValue {
				input[key] = value
			}
			if new_errs != nil {
				errs[key] = append(errs[key], new_errs...)
			}

		} else if v.failOnUnknown {
			errs[key] = append(errs[key], ErrUnknownKey)
			delete(input, key)
		}
	}

	return input, errs
}

func (e *mapValidatorEntry) Key(key string) MapStringValidator {
	return e.root.Key(key)
}

func (e *mapValidatorEntry) IntKey(key string) MapIntValidator {
	return e.root.IntKey(key)
}

func (e *mapValidatorEntry) Validate(input map[string]interface{}) (map[string]interface{}, map[string][]error) {
	return e.root.Validate(input)
}

func (e *mapValidatorEntry) validateMissing() (interface{}, []error) {
	if e.required {
		return ignoreValue, []error{ErrRequiredKeyMissing}
	}
	return e.defaut, nil
}
