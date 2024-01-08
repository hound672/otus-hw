package application

import (
	"github.com/bufbuild/protovalidate-go"
)

func NewValidator() (*protovalidate.Validator, error) {
	return protovalidate.New()
}
