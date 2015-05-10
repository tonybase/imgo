package model

import "errors"

var (
	InvalidMessageError = errors.New("Invalid message")
)

// Database error
type DatabaseError struct {
	err string
}

// Error string
func (this *DatabaseError) Error() string {
	return this.err
}