package errors

type AppError struct {
	Errors []string `json:"errors"`
}

func NewAppError(message string) AppError {
	errorResult := make([]string, 1)
	errorResult[0] = message

	return AppError{Errors: errorResult}
}

func (appErr *AppError) HasErrors() bool {
	if appErr == nil {
		return false
	}
	return len(appErr.Errors) >= 1
}

func (appErr *AppError) Concat(other *AppError) *AppError {
	if other.Errors == nil {
		return appErr
	}

	var errors []string
	errors = append(append(errors, appErr.Errors...), other.Errors...)

	return newAppErrors(errors)
}

func newAppErrors(messages []string) *AppError {
	return &AppError{Errors: messages}
}
