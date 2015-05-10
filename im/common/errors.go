package common

import "errors"

var (
	InvalidMessageError = errors.New("Invalid message")
)

// Server error
type ServerError struct {
	err string
}

// Error string
func (this *ServerError) Error() string {
	return this.err
}

// Protocol error
type ProtocolError struct {
	err string
}

// Error string
func (this *ProtocolError) Error() string {
	return this.err
}

// Configuration error
type ConfigurationError struct {
	err string
}

// Error string
func (this *ConfigurationError) Error() string {
	return this.err
}
