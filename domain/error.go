package domain

type Error struct {
	errorType ErrorType
	message   string
}

func NewErrorType(errorType ErrorType, message string) *Error {
	return &Error{
		errorType: errorType,
		message:   message,
	}
}

func (e *Error) Type() ErrorType {
	return e.errorType
}

func (e *Error) Message() string {
	return e.message
}