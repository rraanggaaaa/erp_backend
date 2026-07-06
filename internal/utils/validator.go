package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidationErrors(err error) map[string]string {

	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {

		errors[e.Field()] = strings.ToLower(e.Tag())

	}

	return errors

}
