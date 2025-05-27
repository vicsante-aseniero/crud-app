package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyName = errors.New("empty name supplied")
)

type Name string

func NewCompanyName(cn string) (Name, error) {
	cn = strings.TrimSpace(cn)
	if cn == "" {
		return "", ErrEmptyName
	}

	// Must be less than 15 chars


	// Check if Name is only alpha
	

	// Validate it is a valid username.
	return Name(cn), nil
}