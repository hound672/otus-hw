package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const symbolToEscape string = "\\"

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	var symbolToRepeat string
	var result strings.Builder
	var isNextSymbolEscaped bool

	for _, symbolRune := range source {
		currentSymbol := string(symbolRune)
		switch {
		case isNextSymbolEscaped:
			if !(unicode.IsDigit(symbolRune) || currentSymbol == symbolToEscape) {
				return "", ErrInvalidString
			}
			symbolToRepeat = currentSymbol
			isNextSymbolEscaped = false

		case currentSymbol == symbolToEscape:
			result.WriteString(symbolToRepeat)
			symbolToRepeat = ""
			isNextSymbolEscaped = true

		case unicode.IsDigit(symbolRune):
			if symbolToRepeat == "" {
				return "", ErrInvalidString
			}
			repeatCount, err := strconv.Atoi(currentSymbol)
			if err != nil {
				return "", err
			}
			result.WriteString(strings.Repeat(symbolToRepeat, repeatCount))
			symbolToRepeat = ""

		default:
			result.WriteString(symbolToRepeat)
			symbolToRepeat = currentSymbol
		}
	}
	if isNextSymbolEscaped {
		return "", ErrInvalidString
	}
	result.WriteString(symbolToRepeat)
	return result.String(), nil
}
