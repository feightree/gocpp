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

		if err.Error() != "[GenericError] foo: bar" {
			t.Errorf("unexpected Error()\ngot:  %s\nwant: %s", err.Error(), "[GenericError] foo: bar")
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

func TestWrapField(t *testing.T) {
	t.Run("wraps an ocpp error with no field name", func(t *testing.T) {
		err := NewError(ErrGenericError, "", "message")
		wrap := WrapField("myfield", err)

		if wrap.Error() != "[GenericError] myfield: message" {
			t.Errorf("unexpected wrapped error\ngot:  %s\nwant: %s", wrap.Error(), "[GenericError] myfield: message")
		}
	})

	t.Run("wraps an ocpp error with a field name", func(t *testing.T) {
		err := NewError(ErrGenericError, "field", "message")
		wrap := WrapField("myfield", err)

		if wrap.Error() != "[GenericError] myfield.field: message" {
			t.Errorf("unexpected wrapped error\ngot:  %s\nwant: %s", wrap.Error(), "[GenericError] myfield.field: message")
		}
	})

	t.Run("doesnt wrap an non-ocp error", func(t *testing.T) {
		err := errors.New("error!")
		wrap := WrapField("myfield", err)

		if wrap.Error() != "error!" {
			t.Errorf("unexpected wrapped error\ngot:  %s\nwant: %s", wrap.Error(), "error!")
		}
	})
}
