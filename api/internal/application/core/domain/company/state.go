package company

import (
	"errors"
	"strings"
)

var (
	ErrEmptyState = errors.New("empty state supplied")
)

type State string

func NewState(sc string) (State, error) {
	sc = strings.TrimSpace(sc)
	if sc == "" {
		return "", ErrEmptyState
	}

	// Add custom validation here

	// Validate it is a valid username.
	return State(sc), nil
}