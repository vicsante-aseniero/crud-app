package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyCountry = errors.New("empty country supplied")
)

type Country string

func NewCountry(cc string) (Country, error) {
	cc = strings.TrimSpace(cc)
	if cc == "" {
		return "", ErrEmptyCountry
	}

	// Add custom validation here

	// Validate it is a valid username.
	return Country(cc), nil
}