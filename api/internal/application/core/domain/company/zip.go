package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyZip = errors.New("empty zip code supplied")
)

type Zip string

func NewZip(zc string) (Zip, error) {
	zc = strings.TrimSpace(zc)
	if zc == "" {
		return "", ErrEmptyEmail
	}

	// Add custom validation here

	// Validate it is a valid username.
	return Zip(zc), nil
}