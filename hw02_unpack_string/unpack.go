package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type State int

const (
	waitRune = iota
	waitRepeatCount
	waitEscaped
)

const escaped string = "\\"

func Unpack(source string) (string, error) {
	runes := []rune(source)
	currentState := waitRune

	result := &strings.Builder{}

	var repeatString string
	for _, el := range runes {

		switch {
		case currentState == waitEscaped:
			if !(unicode.IsDigit(el) || string(el) == escaped) {
				return "", ErrInvalidString
			}
			repeatString = string(el)
			currentState = waitRune
		case string(el) == escaped:
			currentState = waitEscaped
		case currentState == waitRune:
			if unicode.IsDigit(el) {
				return "", ErrInvalidString
			}
			repeatString = string(el)
			currentState = waitRepeatCount
		case currentState == waitRepeatCount:
			if !unicode.IsDigit(el) {
				result.WriteString(repeatString)
				repeatString = string(el)
				break
			}

			repeatCount, err := strconv.Atoi(string(el))
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(string(repeatString), repeatCount))
			currentState = waitRune
		}
	}
	result.WriteString(repeatString)

	return result.String(), nil
}
