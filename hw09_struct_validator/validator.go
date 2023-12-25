package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	ErrRegexp = errors.New("value does not match regular expression")
	ErrLen    = errors.New("length of string does not match expected")
	ErrMax    = errors.New("value is larger than max")
	ErrMin    = errors.New("value is less than min")
	ErrIn     = errors.New("actual value is not allowed")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var result string
	for _, val := range v {
		result += val.Field + ": " + val.Err.Error() + "\n"
	}
	return result
}

func Validate(iv interface{}) error {
	v := reflect.ValueOf(iv)

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("%T is not a pointer to struct", iv)
	}
	t := v.Type()
	er := new(ValidationErrors)
	var err error
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv := v.Field(i)
		var tags []string
		tag := field.Tag.Get("validate")
		if strings.Contains(tag, "|") {
			tags = strings.Split(tag, "|")
		} else {
			tags = append(tags, tag)
		}
		if len(tag) != 0 {
			er, err = validateByKind(field.Name, fv, tags, *er)
			if err != nil {
				return err
			}
		}
	}
	return er
}

func validateByKind(field string, value reflect.Value, tags []string, er ValidationErrors) (*ValidationErrors, error) {
	var err error
	switch {
	case value.Kind() == reflect.String:
		val := value.String()
		er, err = typeSwitch(field, val, tags, er)
	case value.Kind() == reflect.Int:
		val := int(value.Int())
		er, err = typeSwitch(field, val, tags, er)
	case value.Kind() == reflect.Int64:
		val := int(value.Int())
		er, err = typeSwitch(field, val, tags, er)
	case value.Kind() == reflect.Slice:
		er, err = typeSwitch(field, value.Interface(), tags, er)
	}
	return &er, err
}

func typeSwitch(fieldName string, val interface{}, tags []string, er ValidationErrors) (ValidationErrors, error) {
	var err error
	switch h := val.(type) {
	case int:
		er, err = getValidateIntByTag(fieldName, h, tags, er)
	case string:
		er, err = getValidateStringByTag(fieldName, h, tags, er)
	case []string:
		for _, v := range h {
			er, err = getValidateStringByTag(fieldName, v, tags, er)
		}
	case []int:
		for _, v := range h {
			er, err = getValidateIntByTag(fieldName, v, tags, er)
		}
	}
	return er, err
}

// base constructors

func getValidateStringByTag(field string, value string, tags []string, er ValidationErrors) (ValidationErrors, error) {
	var err error
	for _, tag := range tags {
		tagValue := strings.Split(tag, ":")[1]
		switch {
		case strings.HasPrefix(tag, "len:"):
			er, err = getLenValidate(field, value, tagValue, er)
		case strings.HasPrefix(tag, "in:"):
			er = getInValidate(field, value, tagValue, er)
		case strings.HasPrefix(tag, "regexp:"):
			er, err = getRegexpValidate(field, value, tagValue, er)
		}
	}
	return er, err
}

func getValidateIntByTag(field string, value int, tags []string, er ValidationErrors) (ValidationErrors, error) {
	var err error
	for _, tag := range tags {
		tagValue := strings.Split(tag, ":")[1]
		switch {
		case strings.HasPrefix(tag, "in:"):
			i := strconv.Itoa(value)
			er = getInValidate(field, i, tagValue, er)
		case strings.HasPrefix(tag, "min:"):
			er, err = getMinValidate(field, value, tagValue, er)
		case strings.HasPrefix(tag, "max:"):
			er, err = getMaxValidate(field, value, tagValue, er)
		}
	}
	return er, err
}

// validators

func getLenValidate(fieldName string, value string, tagValue string, er ValidationErrors) (ValidationErrors, error) {
	var e ValidationError
	i, err := strconv.Atoi(tagValue)
	if err != nil {
		return er, fmt.Errorf("atoi error: %w", err)
	}
	if len(value) != i {
		e.Field = fieldName
		e.Err = ErrLen
		return append(er, e), nil
	}
	return er, nil
}

// strings

func getInValidate(fieldName string, value string, tagValue string, er ValidationErrors) ValidationErrors {
	var e ValidationError
	dict := strings.Split(tagValue, ",")
	var ok bool
	for _, v := range dict {
		if v == value {
			ok = true
		}
	}
	if !ok {
		e.Field = fieldName
		e.Err = ErrIn
		return append(er, e)
	}
	return er
}

func getRegexpValidate(fieldName string, value string, tagValue string, er ValidationErrors) (ValidationErrors, error) {
	var e ValidationError
	matched, err := regexp.Match(tagValue, []byte(value))
	if err != nil {
		return er, fmt.Errorf("match error: %w", err)
	}
	if !matched {
		e.Field = fieldName
		e.Err = ErrRegexp
		return append(er, e), nil
	}
	return er, nil
}

// ints

func getMinValidate(fieldName string, value int, tagValue string, er ValidationErrors) (ValidationErrors, error) {
	var e ValidationError
	i, err := strconv.Atoi(tagValue)
	if err != nil {
		return er, fmt.Errorf("atoi error: %w", err)
	}
	if i > value {
		e.Field = fieldName
		e.Err = ErrMin
		return append(er, e), nil
	}
	return er, nil
}

func getMaxValidate(fieldName string, value int, tagValue string, er ValidationErrors) (ValidationErrors, error) {
	var e ValidationError
	i, err := strconv.Atoi(tagValue)
	if err != nil {
		return er, fmt.Errorf("atoi error: %w", err)
	}
	if i < value {
		e.Field = fieldName
		e.Err = ErrMax
		return append(er, e), nil
	}
	return er, nil
}
