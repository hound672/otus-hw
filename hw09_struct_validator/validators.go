package hw09structvalidator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// validators

// strings

func getLenValidate(fieldName string, tagValue string) (validatorFuncStr, error) {
	i, err := strconv.Atoi(tagValue)
	if err != nil {
		return nil, fmt.Errorf("atoi error: %w", err)
	}

	validatorFunc := func(val string) *ValidationError {
		if len(val) != i {
			e := &ValidationError{}
			e.Field = fieldName
			e.Err = ErrLen
			return e
		}
		return nil
	}

	return validatorFunc, nil
}

//nolint:unparam // just to math the interface
func getInStrValidate(fieldName string, tagValue string) (validatorFuncStr, error) {
	dict := strings.Split(tagValue, ",")

	validatorFunc := func(val string) *ValidationError {
		e := &ValidationError{}
		var ok bool

		for _, v := range dict {
			if v == val {
				ok = true
			}
		}

		if !ok {
			e.Field = fieldName
			e.Err = ErrIn
			return e
		}

		return nil
	}

	return validatorFunc, nil
}

func getRegexpValidate(fieldName string, tagValue string) (validatorFuncStr, error) {
	reg, err := regexp.Compile(tagValue)
	if err != nil {
		return nil, fmt.Errorf("regexp.Compile error: %w", err)
	}

	validatorFunc := func(val string) *ValidationError {
		e := &ValidationError{}

		if !reg.Match([]byte(val)) {
			e.Field = fieldName
			e.Err = ErrRegexp
			return e
		}

		return nil
	}

	return validatorFunc, nil
}

// ints

func getMinValidate(fieldName string, tagValue string) (validatorFuncInt, error) {
	i, err := strconv.Atoi(tagValue)
	if err != nil {
		return nil, fmt.Errorf("atoi error: %w", err)
	}

	validatorFunc := func(val int) *ValidationError {
		e := &ValidationError{}
		if i > val {
			e.Field = fieldName
			e.Err = ErrMin
			return e
		}
		return nil
	}

	return validatorFunc, nil
}

func getMaxValidate(fieldName string, tagValue string) (validatorFuncInt, error) {
	i, err := strconv.Atoi(tagValue)
	if err != nil {
		return nil, fmt.Errorf("atoi error: %w", err)
	}

	validatorFunc := func(val int) *ValidationError {
		e := &ValidationError{}
		if i < val {
			e.Field = fieldName
			e.Err = ErrMax
			return e
		}
		return nil
	}

	return validatorFunc, nil
}

func getInIntValidate(fieldName string, tagValue string) (validatorFuncInt, error) {
	dict := strings.Split(tagValue, ",")
	references := make([]int, 0, len(dict))
	for _, v := range dict {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("convert %s to int", v)
		}
		references = append(references, r)
	}

	validatorFunc := func(val int) *ValidationError {
		e := &ValidationError{}
		var ok bool

		for _, v := range references {
			if v == val {
				ok = true
			}
		}

		if !ok {
			e.Field = fieldName
			e.Err = ErrIn
			return e
		}

		return nil
	}

	return validatorFunc, nil
}
