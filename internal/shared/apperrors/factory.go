package apperrors

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

func NewConflict(msg string) AppError {
	return AppError{
		Code:    409,
		Message: msg,
	}
}

func NewTooManyRequests(details string) AppError {
	return AppError{
		Code:    429,
		Message: "too many requests",
		Details: details,
	}
}

func NewInternal() AppError {
	return AppError{
		Code:    500,
		Message: "internal server error",
	}
}

func NewValidationError(fields []FieldError) AppError {
	return AppError{
		Code:    400,
		Message: "bad request",
		Errors:  fields,
	}
}
