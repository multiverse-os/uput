package valid

import (
	"math"
)

type intValidator struct {
	min, max int64
}

func NewIntValidator() IntValidator {
	return &intValidator{
		min: math.MinInt64,
		max: math.MaxInt64,
	}
}

func (v *intValidator) Range(min, max int64) IntValidator {
	v.min = min
	v.max = max

	return v
}

func (v *intValidator) Validate(input int64) (int64, []error) {
	if input < v.min {
		return input, []error{ErrTooSmall}
	}
	if input > v.max {
		return input, []error{ErrTooLarge}
	}
	return input, nil
}
