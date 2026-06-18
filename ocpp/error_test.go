package ocpp

import (
	"errors"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Run("returns a valid error with a field", func(t *testing.T) {
		err := NewError(ErrGenericError, "foo", "bar")

		if err.Code != ErrGenericError {
			t.Errorf("unexpected code\ngot:  %v\nwant: %v", err.Code, ErrGenericError)
		}

		if err.Field != "foo" {
			t.Errorf("unexpected field\ngot:  %v\nwant: %v", err.Field, "foo")
		}

		if err.Message != "bar" {
			t.Errorf("unexpected message\ngot:  %s\nwant: %s", err.Message, "bar")
		}

		if err.Error() != "[GenericError] field 'foo': bar" {
			t.Errorf("unexpected Error()\ngot:  %s\nwant: %s", err.Error(), "[GenericError] field 'foo': bar")
		}

		if !errors.Is(err, ErrGenericError) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ErrGenericError)
		}
	})

	t.Run("returns a valid error without a field", func(t *testing.T) {
		err := NewError(ErrInternalError, "", "bar")

		if err.Code != ErrInternalError {
			t.Errorf("unexpected code\ngot:  %v\nwant: %v", err.Code, ErrInternalError)
		}

		if err.Field != "" {
			t.Errorf("unexpected field\ngot:  %v\nwant: %v", err.Field, "")
		}

		if err.Message != "bar" {
			t.Errorf("unexpected message\ngot:  %s\nwant: %s", err.Message, "bar")
		}

		if err.Error() != "[InternalError] bar" {
			t.Errorf("unexpected Error()\ngot:  %s\nwant: %s", err.Error(), "[InternalError] bar")
		}

		if !errors.Is(err, ErrInternalError) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ErrInternalError)
		}
	})
}
