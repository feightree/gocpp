package ocpp16

import "errors"

var (
	ErrStringLength20  = errors.New("exceeds maximum length of 20 characters")
	ErrStringLength25  = errors.New("exceeds maximum length of 25 characters")
	ErrStringLength50  = errors.New("exceeds maximum length of 50 characters")
	ErrStringLength255 = errors.New("exceeds maximum length of 255 characters")
	ErrStringLength500 = errors.New("exceeds maximum length of 500 characters")
	ErrTypeString      = errors.New("must be of type 'string'")
	ErrNonVisibleASCII = errors.New("characters must fall within the visible ASCII range (hex 0x21 to 0x7E)")
	ErrInvalidEnum     = errors.New("invalid enum value")
)
