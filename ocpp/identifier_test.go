package ocpp

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateIdentifierString(t *testing.T) {
	t.Run("accepts valid identifier strings", func(t *testing.T) {
		tests := map[string]string{
			"empty":         "",
			"lowercase":     "abcxyz",
			"uppercase":     "ABCXYZ",
			"digits":        "0123456789",
			"symbols":       "*-_=:+|@.",
			"mixed":         "CS-001_a@b.c",
			"at max length": strings.Repeat("a", 20),
		}

		for name, s := range tests {
			t.Run(name, func(t *testing.T) {
				if err := ValidateIdentifierString(s, 20, "TestType"); err != nil {
					t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
				}
			})
		}
	})

	t.Run("rejects strings exceeding max length", func(t *testing.T) {
		err := ValidateIdentifierString(strings.Repeat("a", 21), 20, "TestType")

		if !errors.Is(err, ErrPropertyConstraintViolation) {
			t.Fatalf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ErrPropertyConstraintViolation)
		}

		want := "[PropertyConstraintViolation] TestType exceeds max length of 20"
		if err.Error() != want {
			t.Errorf("unexpected Error()\ngot:  %s\nwant: %s", err.Error(), want)
		}
	})

	t.Run("rejects characters outside the charset", func(t *testing.T) {
		tests := map[string]struct {
			s    string
			want string
		}{
			"space":     {"a b", `[PropertyConstraintViolation] TestType contains ' ', which is outside the identifierString charset`},
			"slash":     {"a/b", `[PropertyConstraintViolation] TestType contains '/', which is outside the identifierString charset`},
			"hash":      {"#", `[PropertyConstraintViolation] TestType contains '#', which is outside the identifierString charset`},
			"non-ASCII": {"café", `[PropertyConstraintViolation] TestType contains 'é', which is outside the identifierString charset`},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				err := ValidateIdentifierString(tt.s, 20, "TestType")

				if !errors.Is(err, ErrPropertyConstraintViolation) {
					t.Fatalf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ErrPropertyConstraintViolation)
				}

				if err.Error() != tt.want {
					t.Errorf("unexpected Error()\ngot:  %s\nwant: %s", err.Error(), tt.want)
				}
			})
		}
	})
}
