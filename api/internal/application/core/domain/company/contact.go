package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyContact = errors.New("empty contact no. supplied")
)

type Contact string

func NewContactNo(cn string) (Contact, error) {
	cn = strings.TrimSpace(cn)
	if cn == "" {
		return "", ErrEmptyContact
	}

	// Add custom validation here

	// Validate it is a valid username.
	return Contact(cn), nil
}