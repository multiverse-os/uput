package is

import (
	"regexp"
	"strconv"
)

func (v Validator) Eq(check string, value interface{}, t string) bool {

	r := string(regexp.MustCompile("Eq\\((.*)\\)").ReplaceAllString(check, "$1"))

	if r == "" {
		return true
	}

	switch t {

	case "string":
		if r == value.(string) {
			return true
		}
	case "int":
		if n, err := strconv.Atoi(r); err == nil {
			if n == value.(int) {
				return true
			}
		}
	case "bool":
		if n, err := strconv.ParseBool(r); err == nil {
			if n == value.(bool) {
				return true
			}
		}
	case "float64", "float32":
		if n, err := strconv.ParseFloat(r, 64); err == nil {
			if n == value.(float64) {
				return true
			}
		}

	}

	return false

}

func (v Validator) Regex(check string, value interface{}, t string) bool {

	r := string(regexp.MustCompile("Regex\\(([\\s\\S]*)\\)").ReplaceAllString(check, "$1"))

	if t != "string" {
		return false
	}

	if r == "" {
		return true
	}

	m, _ := regexp.MatchString(r, value.(string))

	return m

}

func (v Validator) Minlen(check string, value interface{}, t string) bool {

	if t != "string" {
		return false
	}

	r := string(regexp.MustCompile("Minlen\\(([\\d.]*)\\)").ReplaceAllString(check, "$1"))

	l, err := strconv.Atoi(r)

	if err != nil {
		return false
	}

	return len(value.(string)) >= l

}

func (v Validator) Maxlen(check string, value interface{}, t string) bool {

	if t != "string" {
		return false
	}

	r := string(regexp.MustCompile("Maxlen\\(([\\d.]*)\\)").ReplaceAllString(check, "$1"))

	l, err := strconv.Atoi(r)

	if err != nil {
		return false
	}

	return len(value.(string)) <= l

}

func (v Validator) Min(check string, value interface{}, t string) bool {

	r := string(regexp.MustCompile("Min\\(([\\d.]*)\\)").ReplaceAllString(check, "$1"))

	if t == "int" {

		m, err := strconv.Atoi(r)

		if err != nil {
			return false
		}

		return value.(int) >= m

	} else if t == "int32" {

		m, err := strconv.Atoi(r)

		if err != nil {
			return false
		}

		return value.(int32) >= int32(m)

	} else if t == "int64" {

		m, err := strconv.Atoi(r)

		if err != nil {
			return false
		}

		return value.(int64) >= int64(m)

	} else if t == "float32" {

		m, err := strconv.ParseFloat(r, 32)

		if err != nil {
			return false
		}

		return value.(float32) >= float32(m)

	} else if t == "float64" {

		m, err := strconv.ParseFloat(r, 64)

		if err != nil {
			return false
		}

		return value.(float64) >= m

	}

	return false

}

func (v Validator) Max(check string, value interface{}, t string) bool {

	r := string(regexp.MustCompile("Max\\(([\\d.]*)\\)").ReplaceAllString(check, "$1"))

	if t == "int" {

		m, err := strconv.Atoi(r)

		if err != nil {
			return false
		}

		return value.(int) <= m

	} else if t == "int32" {

		m, err := strconv.Atoi(r)

		if err != nil {
			return false
		}

		return value.(int32) <= int32(m)

	} else if t == "int64" {

		m, err := strconv.Atoi(r)

		if err != nil {
			return false
		}

		return value.(int64) <= int64(m)

	} else if t == "float32" {

		m, err := strconv.ParseFloat(r, 32)

		if err != nil {
			return false
		}

		return value.(float32) <= float32(m)

	} else if t == "float64" {

		m, err := strconv.ParseFloat(r, 64)

		if err != nil {
			return false
		}

		return value.(float64) <= m

	}

	return false

}
