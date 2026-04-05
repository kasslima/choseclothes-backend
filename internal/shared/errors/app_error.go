package errors

type AppError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Details string       `json:"details,omitempty"`
	Errors  []FieldError `json:"errors,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}