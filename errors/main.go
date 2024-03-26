// Package errors custom.
//
// Custom errors from this project.
package errors

import (
	"errors"
)

var (
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrPhoneRequired      = errors.New("phone is required")
	ErrUserRequired       = errors.New("user is required")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidPassword    = errors.New("invalid password, must have at least 6 characters max lenght of 12, 1 uppercase, 1 lowercase, 1 number and 1 symbol")
	ErrPhoneNumberInvalid = errors.New("invalid length of phone number, must be 10 digits an be numeric")
	ErrTokenGeneration    = errors.New("error generating token")
)
