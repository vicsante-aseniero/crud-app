package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyWebsite = errors.New("empty website supplied")
)

type Website string

func NewWebsite(ws string) (Website, error) {
	ws = strings.TrimSpace(ws)
	if ws == "" {
		return "", ErrEmptyWebsite
	}

	// Add custom validation here


	// Validate it is a valid username.
	return Website(ws), nil
}