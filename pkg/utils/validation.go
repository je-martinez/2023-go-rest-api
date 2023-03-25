package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStructErrors(err error) []ErrorField {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorField, len(ve))
		for i, fe := range ve {
			out[i] = ErrorField{fe.Field(), msgForTag(fe)}
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
