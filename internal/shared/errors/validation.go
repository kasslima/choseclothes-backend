package errors

import (
	"encoding/json"
	"errors"
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

func ParseBindError(err error) AppError {
	if err.Error() == "EOF" {
		return NewBadRequest("body is required")
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		return NewBadRequest("invalid field type in request body")
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return NewBadRequest("invalid JSON format")
	}

	return NewBadRequest("invalid request body")
}
