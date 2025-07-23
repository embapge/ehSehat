package utils

type badRequestError struct {
	message string
}

func (e *badRequestError) Error() string {
	return e.message
}

func NewBadRequestError(message string) error {
	return &badRequestError{message: message}
}
