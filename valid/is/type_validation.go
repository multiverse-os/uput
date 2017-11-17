package validation

// Type Validation
func ValidateType(value interface{}, expectedType string) bool {
	valueTypeString := reflect.TypeOf(value).String()
	return CheckRegexp(valueTypeString, `^`+expectedType+`((32)|(64))?$`)
}

// Regex Validation
func CheckRegexp(value interface{}, rex string) bool {
	val, ok := value.(string)
	if !ok {
		val = convert.MustString(value)
	}
	reg := regexp.MustCompile(rex)
	return reg.MatchString(val)
}

// TODO: These should just be combined, into single using func with map
func ValidateEmail(value interface{}) bool {
	return CheckRegexp(value, emailRegex)
}

func ValidatePhone(value interface{}) bool {
	return CheckRegexp(value, phoneRegex)
}

// Go Generic Type Validation
func ValidateIntRange(value interface{}, min, max int) bool {
	if val, err := convert.ToInt(value); err != nil {
		return false
	} else {
		return min <= val && val <= max
	}
}

func ValidateFloat64Range(value interface{}, min, max float64) bool {
	if !ValidateRealNumber(value) {
		return false
	}

	if val, err := convert.ToFloat64(value); err != nil {
		return false
	} else {
		return min <= val && val <= max
	}
}

func CheckRealNumber(value interface{}) bool {
	val := convert.MustString(value)
	return CheckRegexp(val, regularNumber)
}

// String, Slice, Array
func CheckLen(value interface{}, length int) bool {
	switch value.(type) {
	case string:
		return CheckRegexp(value, fmt.Sprintf("^.{%d}$", length))
	default:
		{
			refValue := reflect.ValueOf(value)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					return refValue.Len() == length
				}
			}
		}
	}

	return false
}

// Numeric
func CheckMin(value interface{}, min float64) bool {
	switch value.(type) {
	case string:
		{
			res, err := convert.ToFloat64(value)
			if err == nil {
				return res > float64(min)
			}
		}
	case int, int8, int16, int32, int64:
		{
			return convert.MustInt64(value) >= int64(min)
		}
	case uint, uint8, uint16, uint32, uint64:
		{
			return convert.MustUint64(value) >= uint64(min)
		}
	case float32, float64:
		{
			return convert.MustFloat64(value) >= float64(min)
		}
	default:
		{
			return CheckMin(convert.MustString(value), min)
		}
	}
	return false
}

func CheckMax(value interface{}, max float64) bool {
	switch value.(type) {
	case string:
		{
			res, err := convert.ToFloat64(value)
			if err == nil {
				return res <= float64(max)
			}
		}
	case int, int8, int16, int32, int64:
		{
			return convert.MustInt64(value) <= int64(max)
		}
	case uint, uint8, uint16, uint32, uint64:
		{
			return convert.MustUint64(value) <= uint64(max)
		}
	case float32, float64:
		{
			return convert.MustFloat64(value) <= float64(max)
		}
	default:
		{
			return CheckMin(convert.MustString(value), max)
		}
	}
	return false
}

// Size = String, Slice, Array
func CheckMaxSize(value interface{}, maxSize int) bool {
	switch value.(type) {
	case string:
		{
			return CheckRegexp(value, fmt.Sprintf("^.{0,%d}$", maxSize))
		}
	default:
		{
			refValue := reflect.ValueOf(value)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					return refValue.Len() <= maxSize
				}
			}
		}
	}

	return false
}

func CheckMinSize(value interface{}, minSize int) bool {
	switch value.(type) {
	case string:
		{
			return CheckRegexp(value, fmt.Sprintf("^.{%d,}$", minSize))
		}
	default:
		{
			refValue := reflect.ValueOf(value)
			switch refValue.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				{
					return refValue.Len() >= minSize
				}
			}
		}
	}

	return false
}

func IsValidBoolean(s string) bool {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "0", "f", "F", "false", "FALSE", "False":
		return true
	}
	return false
}

func IsValidNumber(s string) bool {
	// This function implements the JSON numbers grammar.
	// See https://tools.ietf.org/html/rfc7159#section-6
	// and http://json.org/number.gif

	if s == "" {
		return false
	}

	// Optional -
	if s[0] == '-' {
		s = s[1:]
		if s == "" {
			return false
		}
	}

	// Digits
	switch {
	default:
		return false

	case s[0] == '0':
		s = s[1:]

	case '1' <= s[0] && s[0] <= '9':
		s = s[1:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// . followed by 1 or more digits.
	if len(s) >= 2 && s[0] == '.' && '0' <= s[1] && s[1] <= '9' {
		s = s[2:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// e or E followed by an optional - or + and
	// 1 or more digits.
	if len(s) >= 2 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
			if s == "" {
				return false
			}
		}
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
		}
	}

	// Make sure we are at the end.
	return s == ""
}

func IsArray(value interface{}) bool {
	refValue := reflect.ValueOf(value)
	switch refValue.Kind() {
	case reflect.Slice, reflect.Array:
		{
			return true
		}
	default:
		return false
	}

}
