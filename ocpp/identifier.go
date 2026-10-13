package ocpp

import (
	"fmt"
	"strings"
)

// ValidateIdentifierString validates that s is no longer than maxLen and
// contains only characters from the identifierString charset defined in
// OCPP 2.0.1 and 2.1: a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
func ValidateIdentifierString(s string, maxLen int, typeName string) error {
	if len(s) > maxLen {
		return NewError(ErrPropertyConstraintViolation, "", fmt.Sprintf("%s exceeds max length of %d", typeName, maxLen))
	}

	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
		case strings.ContainsRune("*-_=:+|@.", r):
		default:
			return NewError(ErrPropertyConstraintViolation, "", fmt.Sprintf("%s contains %q, which is outside the identifierString charset", typeName, r))
		}
	}

	return nil
}
