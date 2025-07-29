package app

import "errors"

var (
	ErrNotFound           = errors.New("data not found")
	ErrInvalidInput       = errors.New("invalid input")
	ErrMissingFields      = errors.New("missing required fields")
	ErrMissingID          = errors.New("missing ID")
	ErrInternal           = errors.New("internal server error")
	ErrEmailAlreadyExists = errors.New("email is already registered")
)
