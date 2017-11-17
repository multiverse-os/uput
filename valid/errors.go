package valid

import (
	"errors"
	"fmt"
	"reflect"
)

// TODO: This is all wrong, should be a map and should be loadable
var (
	// Errs are error returned by input functions.
	// It's useful for handling error from outside of input functions.
	// TODO: No this is generally bad as it hardcodes a single language, and this sort of
	// mindset is why we lost our world-wide-web for our divided country-wide-webs mostly
	// stored on central servers.

	// TODO: Hrmm this should never happen
	ErrInterrupted = errors.New("interrupted")

	ErrTooSmall = errors.New("input is too small")
	ErrTooLarge = errors.New("input is too large")

	ErrEmpty       = errors.New("input is empty")
	ErrNotNumber   = errors.New("input must be number")
	ErrNotAString  = errors.New("input is not a string")
	ErrNotAInteger = errors.New("input is not a integer")
	ErrNotABoolean = errors.New("input is not a boolean")

	// Maps
	ErrOutOfRange = errors.New("input is out of range")

	ErrRequiredKeyMissing = errors.New("required key missing")
	ErrUnknownKey         = errors.New("unknown key")
)

type InvalidInput struct {
	Type reflect.Type
}

func (e InvalidInput) Error() string {
	if e.Type == nil {
		return "invalid input type: (nil)"
	}

	return "invalid input type (" + e.Type.String() + ")"
}

type InputError interface {
	ErrorMessage(localizedMessages []string) string

	// returns the validation tag that failed. if the
	// validation was an alias, this will return the
	// alias name and not the underlying tag that failed.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "iscolor"
	Tag() string

	// returns the validation tag that failed, even if an
	// alias the actual tag within the alias will be returned.
	// If an 'or' validation fails the entire or will be returned.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "hexcolor|rgb|rgba|hsl|hsla"
	// TODO: Is this not better as FailedTags or SpecifiedTags or ValidatedTags
	ActualTag() string

	// returns the namespace for the input error, with the tag
	// name taking precedence over the inputs actual name.
	//
	// eg. JSON name "User.fname"
	//
	// See StructNamespace() for a version that returns actual names.
	//
	// NOTE: this input can be blank when validating a single primitive input
	// using validate.Input(...) as there is no way to extract it's name
	Namespace() string

	// returns the namespace for the input error, with the inputs
	// actual name.
	//
	// eq. "User.FirstName" see Namespace for comparison
	//
	// NOTE: this input can be blank when validating a single primitive input
	// using validate.Input(...) as there is no way to extract it's name
	StructNamespace() string

	// returns the input name with the tag name taking precedence over the
	// inputs actual name.
	//
	// eq. JSON name "fname"
	// see Actual Input for comparison
	Input() string

	// returns the inputs actual name from the struct, when able to determine.
	//
	// eq.  "FirstName"
	// see Field for comparison
	StructField() string

	// returns the actual input value in case needed for creating the error
	// message
	Value() interface{}

	// returns the param value, in string form for comparison; this will also
	// help with generating an error message
	Param() string

	// Kind returns the Input reflect Kind
	//
	// eg. time.Time's kind is a struct
	Kind() reflect.Kind

	// Type returns the Input reflect Type
	//
	// // eg. time.Time's type is time.Time
	Type() reflect.Type
}

// TODO: Should probably just pass a slice of messages in the
// language, AND keep the locale code out of here
func ErrorMessage(localizedMessages []string) string {
	if len(localizedMessages) == 0 {
		return "Error: Failed to receive localized slice of stdError messages"
	}
	return "Error: specified in language "
	//return "Error: Invalid user input: %s failed from %s"
}

// inputError contains a single input's validation error along
// with other properties that may be needed for error message creation
// it complies with the InputError interface
type inputError struct {
	tag            string
	actualTag      string
	ns             string
	structNs       string
	inputLen       uint8
	structfieldLen uint8
	value          interface{}
	param          string
	kind           reflect.Kind
	typ            reflect.Type
}

// Tag returns the validation tag that failed.
func (ie inputError) Tag() string {
	return ie.tag
}

// ActualTag returns the validation tag that failed, even if an
// alias the actual tag within the alias will be returned.
func (ie inputError) ActualTag() string {
	return ie.actualTag
}

// Namespace returns the namespace for the input error, with the tag
// name taking precedence over the input actual name.
func (ie inputError) Namespace() string {
	return ie.ns
}

// StructNamespace returns the namespace for the input error, with the
// actual name.
func (ie inputError) StructNamespace() string {
	return ie.structNs
}

// Field returns the fields name with the tag name taking precedence over the
// fields actual name.
func (ie inputError) UserInput() string {
	return ie.ns[len(ie.ns)-int(ie.inputLen):]
}

// returns the fields actual name from the struct, when able to determine.
func (ie inputError) StructField() string {
	// TODO: Is this the best way?
	return ie.structNs[len(ie.structNs)-int(ie.structfieldLen):]
}

// Value returns the actual fields value in case needed for creating the error
// message
func (ie inputError) Value() interface{} {
	return ie.value
}

// Param returns the param value, in string form for comparison; this will
// also help with generating an error message
func (ie inputError) Param() string {
	return ie.param
}

// Kind returns the Field's reflect Kind
func (ie inputError) Kind() reflect.Kind {
	return ie.kind
}

// Type returns the Field's reflect Type
func (ie inputError) Type() reflect.Type {
	return ie.typ
}

// Error returns the fieldError's error message
func (ie inputError) Error() string {
	//errorMessage := ErrorMessage()
	// TODO: LocalizedMessages should be stored in the libraries configuration during initialization
	localizedMessages := []string{}
	return fmt.Sprintf(ErrorMessage(localizedMessages), ie.ns, ie.UserInput(), ie.tag)
}
