package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
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

type (
	ValidationErrors []ValidationError
	validatorFuncInt = func(value int) *ValidationError
	validatorFuncStr = func(value string) *ValidationError
)

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
	resValidationErrors := make(ValidationErrors, 0)

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
			valErrs, err := validateByKind(field.Name, fv, tags)
			if err != nil {
				return err
			}
			resValidationErrors = append(resValidationErrors, valErrs...)
		}
	}

	if len(resValidationErrors) == 0 {
		return nil
	}
	return resValidationErrors
}

func validateByKind(field string, value reflect.Value, tags []string) (ValidationErrors, error) {
	switch {
	case value.Kind() == reflect.String:
		val := value.String()
		return typeSwitch(field, val, tags)
	case value.Kind() == reflect.Int:
		val := int(value.Int())
		return typeSwitch(field, val, tags)
	case value.Kind() == reflect.Int64:
		val := int(value.Int())
		return typeSwitch(field, val, tags)
	case value.Kind() == reflect.Slice:
		return typeSwitch(field, value.Interface(), tags)
	}

	return nil, fmt.Errorf("unknown type: %v", value.Kind())
}

//nolint:gocognit // separated funcs will be harder to read
func typeSwitch(fieldName string, val interface{}, tags []string) (ValidationErrors, error) {
	var validationErrors ValidationErrors

	switch h := val.(type) {
	case int:
		validators, err := getValidateIntByTag(fieldName, tags)
		if err != nil {
			return nil, err
		}
		for _, validator := range validators {
			if err := validator(h); err != nil {
				validationErrors = append(validationErrors, *err)
			}
		}
	case string:
		validators, err := getValidateStringByTag(fieldName, tags)
		if err != nil {
			return nil, err
		}
		for _, validator := range validators {
			if err := validator(h); err != nil {
				validationErrors = append(validationErrors, *err)
			}
		}
	case []string:
		validators, err := getValidateStringByTag(fieldName, tags)
		if err != nil {
			return nil, err
		}

		for _, v := range h {
			for _, validator := range validators {
				if err := validator(v); err != nil {
					validationErrors = append(validationErrors, *err)
				}
			}
		}
	case []int:
		validators, err := getValidateIntByTag(fieldName, tags)
		if err != nil {
			return nil, err
		}

		for _, v := range h {
			for _, validator := range validators {
				if err := validator(v); err != nil {
					validationErrors = append(validationErrors, *err)
				}
			}
		}
	}
	return validationErrors, nil
}

// base constructors

func getValidateStringByTag(field string, tags []string) ([]validatorFuncStr, error) {
	validators := make([]validatorFuncStr, 0, len(tags))

	for _, tag := range tags {
		tagValue := strings.Split(tag, ":")[1]
		switch {
		case strings.HasPrefix(tag, "len:"):
			validator, err := getLenValidate(field, tagValue)
			if err != nil {
				return nil, err
			}
			validators = append(validators, validator)
		case strings.HasPrefix(tag, "in:"):
			validator, err := getInStrValidate(field, tagValue)
			if err != nil {
				return nil, err
			}
			validators = append(validators, validator)
		case strings.HasPrefix(tag, "regexp:"):
			validator, err := getRegexpValidate(field, tagValue)
			if err != nil {
				return nil, err
			}
			validators = append(validators, validator)
		}
	}

	return validators, nil
}

func getValidateIntByTag(field string, tags []string) ([]validatorFuncInt, error) {
	validators := make([]validatorFuncInt, 0, len(tags))

	for _, tag := range tags {
		tagValue := strings.Split(tag, ":")[1]
		switch {
		case strings.HasPrefix(tag, "in:"):
			validator, err := getInIntValidate(field, tagValue)
			if err != nil {
				return nil, err
			}

			validators = append(validators, validator)
		case strings.HasPrefix(tag, "min:"):
			validator, err := getMinValidate(field, tagValue)
			if err != nil {
				return nil, err
			}

			validators = append(validators, validator)
		case strings.HasPrefix(tag, "max:"):
			validator, err := getMaxValidate(field, tagValue)
			if err != nil {
				return nil, err
			}

			validators = append(validators, validator)
		}
	}
	return validators, nil
}
