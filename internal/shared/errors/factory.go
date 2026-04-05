package errors

func NewBadRequest(details string) AppError {
	return AppError{
		Code:    400,
		Message: "bad request",
		Details: details,
	}
}

func NewUnauthorized(details string) AppError {
	return AppError{
		Code:    401,
		Message: "unauthorized",
		Details: details,
	}
}

func NewNotFound(details string) AppError {
	return AppError{
		Code:    404,
		Message: "resource not found",
		Details: details,
	}
}

func NewTooManyRequests(details string) AppError {
	return AppError{
		Code:    429,
		Message: "too many requests",
		Details: details,
	}
}

func NewInternal(details string) AppError {
	return AppError{
		Code:    500,
		Message: "internal server error",
		Details: details,
	}
}

func NewValidationError(fields []FieldError) AppError {
	return AppError{
		Code:    400,
		Message: "bad request",
		Errors:  fields,
	}
}
