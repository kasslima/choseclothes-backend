package errors

type AppError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Details string       `json:"details,omitempty"`
	Errors  []FieldError `json:"errors,omitempty"`
}

func (e AppError) Error() string {
	if e.Details != "" {
		return e.Details
	}
	return e.Message
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
