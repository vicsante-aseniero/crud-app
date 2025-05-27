package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyEmail = errors.New("empty email supplied")
)

type Email string

func NewEmail(em string) (Email, error) {
	em = strings.TrimSpace(em)
	if em == "" {
		return "", ErrEmptyEmail
	}

	// Add custom validation here

	// Validate it is a valid username.
	return Email(em), nil
}