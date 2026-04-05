package errors

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationErrors(err error) []FieldError {
	var errorsList []FieldError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return errorsList
	}

	for _, fieldErr := range validationErrors {
		field := strings.ToLower(fieldErr.Field())

		var message string

		switch fieldErr.Tag() {
			
		case "required":
			message = "is required"
		case "email":
			message = "must be a valid email"
		case "min":
			message = "is too short"
		case "max":
			message = "is too long"
		default:
			message = "is invalid"
		}

		errorsList = append(errorsList, FieldError{
			Field:   field,
			Message: message,
		})
	}

	return errorsList
}