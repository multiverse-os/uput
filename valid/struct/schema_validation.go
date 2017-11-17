package valid

import (
	"errors"
	"reflect"
)

type (
	Validatable interface {
		Validate() error
	}

	Rule interface {
		Validate(value interface{}) error
	}

	RuleFunc func(value interface{}) error
)

var (
	ErrorTag = "json"

	Skip = &skipRule{}

	validatableType = reflect.TypeOf((*Validatable)(nil)).Elem()
)

// validateMap validates a map of validatable elements
func validateMap(rv reflect.Value) (errs []error) {
	for _, key := range rv.MapKeys() {
		if mv := rv.MapIndex(key).Interface(); mv != nil {
			if err := mv.(Validatable).Validate(); err != nil {
				errs = append(errs, errors.New("put the right error in here"))
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// validateMap validates a slice/array of validatable elements
func validateSlice(rv reflect.Value) (errs []error) {
	l := rv.Len()
	for i := 0; i < l; i++ {
		if ev := rv.Index(i).Interface(); ev != nil {
			if err := ev.(Validatable).Validate(); err != nil {
				errs = append(errs, errors.New("put right error here"))
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

type skipRule struct{}

func (r *skipRule) Validate(interface{}) error {
	return nil
}

type inlineRule struct {
	f RuleFunc
}

func (r *inlineRule) Validate(value interface{}) error {
	return r.f(value)
}

// By wraps a RuleFunc into a Rule.
func By(f RuleFunc) Rule {
	return &inlineRule{f}
}

// ----------------------------------------
// Above comes from ozzo
// ----------------------------------------

// Similar method to validator, may not go with this
type StringInput string

// TODO: Could have two versions of this lib, basically one that returns
// just the validated string and any possible errors to simplify everything
// as much as possible and allow for very quick and slim use.

// The other returning the userInput object with all of that jazz

// TODO: uinput library implies its more than just valdiation
// it should be able to chain in additional thins like:
// * Sanitation of values
// * Default fallback
// * Trimming/modding/coversion/etc of data

// TODO: Can these be conbined into a single one?
// Validate Basic Datatypes
//type isValid func(input reflect.Value, dataType reflect.Type) (bool, error)

// Start with reflect.ValueOf(data).IsValid()

//type isStringValid func(s string) (bool, error)
//type isIntValid func(n int) (bool, error)
//type isUintValid func(n uint) (bool, error)

// Validate Slice Datatypes
//type isStringSliceValid func(s []string) (bool, error)
//type isIntSliceValid func(n []int) (bool, error)
//type isUintSliceValid func(n []uint) (bool, error)

// Validate String
//func ValidateString(userInput string, validate isStringValid) (string, error) {
//	fmt.Println("[userInput:Validation] string:", userInput)
//	isValid, err := validate(userInput)
//	if isValid {
//		return userInput, nil
//	} else {
//		return "", err
//	}
//}
//
// Common String Validations
//func Maximum(userInput string) (bool, error) {
//	fmt.Println("[UserInput:Validation] string.IsEmpty?", userInput)
//	if userInput == "" {
//		return false, errors.New("is empty")
//	} else {
//		return true, nil
//	}
//}
//
//func IsEmpty(userInput string) (bool, error) {
//	fmt.Println("[UserInput:Validation] string.IsEmpty?", userInput)
//	if userInput == "" {
//		return false, errors.New("is empty")
//	} else {
//		return true, nil
//	}
//}
//
//func IsNotEmpty(userInput string) (bool, error) {
//	fmt.Println("[UserInput:Validation] string.IsNotEmpty?", userInput)
//	if userInput != "" {
//		return false, errors.New("not empty")
//	} else {
//		return true, nil
//	}
//}
