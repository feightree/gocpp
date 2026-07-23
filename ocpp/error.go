// Package ocpp provides common shared types and utilities for OCPP (Open Charge
// Point Protocol) implementations.
package ocpp

import (
	"errors"
	"fmt"
)

type ErrorCode string

const (
	// Requested Action is not known by receiver
	ErrNotImplemented ErrorCode = "NotImplemented"
	// Requested Action is recognized but not supported by the receiver
	ErrNotSupported ErrorCode = "NotSupported"
	// An internal error occurred and the receiver was not able to process the
	// requested Action successfully
	ErrInternalError ErrorCode = "InternalError"
	// Payload for Action is incomplete
	ErrProtocolError ErrorCode = "ProtocolError"
	// During the processing of Action a security issue occurred preventing
	// receiver from completing the Action successfully
	ErrSecurityError ErrorCode = "SecurityError"
	// Payload for Action is syntactically incorrect or not conform the PDU
	// structure for Action
	ErrFormationViolation ErrorCode = "FormationViolation"
	// Payload is syntactically correct but at least one field contains an
	// invalid value
	ErrPropertyConstraintViolation ErrorCode = "PropertyConstraintViolation"
	// Payload for Action is syntactically correct but at least one of the
	// fields violates occurence constraints
	ErrOccurenceConstraintViolation ErrorCode = "OccurenceConstraintViolation"
	// Payload for Action is syntactically correct but at least one of the
	// fields violates data type constraints (e.g. “somestring”: 12)
	ErrTypeConstraintViolation ErrorCode = "TypeConstraintViolation"
	// Any other error not covered by the previous ones
	ErrGenericError ErrorCode = "GenericError"
)

func (c ErrorCode) Error() string {
	return string(c)
}

type Error struct {
	Code    ErrorCode
	Field   string
	Message string
}

// NewError creates a initialized ValidationError
func NewError(code ErrorCode, field, message string) *Error {
	return &Error{
		Code:    code,
		Field:   field,
		Message: message,
	}
}

// Error implements the standard Go error interface
func (e *Error) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Field, e.Message)
	}

	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap exposes the underlying ErrorCode so that errors.Is can match
// against the package's sentinel error codes (e.g. ocpp.ErrInvalid).
func (e *Error) Unwrap() error {
	return e.Code
}

// WrapField wraps an error with a field prefix, if the error is an ocpp.Error
func WrapField(prefix string, err error) error {
	var e *Error

	if errors.As(err, &e) {
		field := prefix

		if e.Field != "" {
			field = prefix + "." + e.Field
		}

		return NewError(e.Code, field, e.Message)
	}

	return err
}
