package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStructErrors(err error) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = fmt.Sprintf("%s: %s", fe.Field(), msgForTag(fe))
		}
		return out
	}
	return nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error()
}
