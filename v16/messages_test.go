package v16

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/feightree/gocpp/ocpp"
)

func ptr[T any](v T) *T { return &v }

func TestAuthorizeReqUnmarshal(t *testing.T) {
	t.Run("rejects oversized idTag", func(t *testing.T) {
		var v AuthorizeReq
		err := json.Unmarshal([]byte(`{"idTag":"ThisIsOver20Characters"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects non-printable ASCII in idTag", func(t *testing.T) {
		var v AuthorizeReq
		err := json.Unmarshal([]byte(`{"idTag":"café"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v AuthorizeReq
		err := json.Unmarshal([]byte(`{"idTag":"ABC123"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTag != "ABC123" {
			t.Errorf("got %q, want %q", v.IDTag, "ABC123")
		}
	})
}

func TestAuthorizeReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   AuthorizeReq
		wantErr error
	}{
		{
			name:    "rejects empty idTag",
			input:   AuthorizeReq{IDTag: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized idTag",
			input:   AuthorizeReq{IDTag: "ThisIsOver20Characters"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects non-printable idTag characters",
			input:   AuthorizeReq{IDTag: "café"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts valid input",
			input:   AuthorizeReq{IDTag: "ABC123"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizeConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v AuthorizeConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{"status":"BadStatus"}}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v AuthorizeConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{"status":"Accepted"}}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTagInfo.Status != AuthorizationStatusAccepted {
			t.Errorf("got %q, want %q", v.IDTagInfo.Status, AuthorizationStatusAccepted)
		}
	})
}

func TestAuthorizeConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   AuthorizeConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   AuthorizeConf{IDTagInfo: IDTagInfo{Status: AuthorizationStatusAccepted}},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   AuthorizeConf{IDTagInfo: IDTagInfo{Status: ""}},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty parentIdTag",
			input:   AuthorizeConf{IDTagInfo: IDTagInfo{Status: AuthorizationStatusAccepted, ParentIDTag: ptr(IDToken(""))}},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized parentIdTag",
			input:   AuthorizeConf{IDTagInfo: IDTagInfo{Status: AuthorizationStatusAccepted, ParentIDTag: ptr(IDToken("ThisIsOver20Characters"))}},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: %v", err, tt.wantErr)
			}
		})
	}
}

func TestBootNotificationReqUnmarshal(t *testing.T) {
	t.Run("rejects oversized chargePointModel", func(t *testing.T) {
		var v BootNotificationReq
		err := json.Unmarshal([]byte(`{"chargePointModel":"ThisIsOver20Characters","chargePointVendor":"VendorY"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v BootNotificationReq
		err := json.Unmarshal([]byte(`{"chargePointModel":"ModelX","chargePointVendor":"VendorY"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ChargePointModel != "ModelX" {
			t.Errorf("got %q, want %q", v.ChargePointModel, "ModelX")
		}

		if v.ChargePointVendor != "VendorY" {
			t.Errorf("got %q, want %q", v.ChargePointVendor, "VendorY")
		}
	})
}

func TestBootNotificationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   BootNotificationReq
		wantErr error
	}{
		{
			name: "accepts valid input",
			input: BootNotificationReq{
				ChargePointModel:      "ModelX",
				ChargePointVendor:     "VendorY",
				ChargeBoxSerialNumber: ptr(CiString25Type("ChargeBoxSerialNumber")),
			},
			wantErr: nil,
		},
		{
			name:    "rejects empty chargePointModel",
			input:   BootNotificationReq{ChargePointModel: "", ChargePointVendor: "VendorY"},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty chargePointVendor",
			input:   BootNotificationReq{ChargePointModel: "ModelX", ChargePointVendor: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized chargePointModel",
			input:   BootNotificationReq{ChargePointModel: "ThisIsOver20Characters", ChargePointVendor: "VendorY"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects oversized chargePointVendor",
			input:   BootNotificationReq{ChargePointModel: "ModelX", ChargePointVendor: "ThisIsOver20Characters"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects invalid optional field",
			input: BootNotificationReq{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "VendorY",
				FirmwareVersion:   ptr(CiString50Type("ThisStringIsOverFiftyCharactersLongAndShouldBeRejected")),
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: %v", err, tt.wantErr)
			}
		})
	}
}

func TestBootNotificationConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v BootNotificationConf
		err := json.Unmarshal([]byte(`{"currentTime":"2024-01-01T00:00:00Z","interval":30,"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v BootNotificationConf
		err := json.Unmarshal([]byte(`{"currentTime":"2024-01-01T00:00:00Z","interval":30,"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != RegistrationStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, RegistrationStatusAccepted)
		}

		if v.Interval != 30 {
			t.Errorf("got %d, want %d", v.Interval, 30)
		}
	})
}

func TestBootNotificationConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   BootNotificationConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   BootNotificationConf{CurrentTime: time.Now(), Interval: 30, Status: RegistrationStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects zero currentTime",
			input:   BootNotificationConf{Interval: 30, Status: RegistrationStatusAccepted},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects zero interval",
			input:   BootNotificationConf{CurrentTime: time.Now(), Interval: 0, Status: RegistrationStatusAccepted},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects empty status",
			input:   BootNotificationConf{CurrentTime: time.Now(), Interval: 30, Status: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: %v", err, tt.wantErr)
			}
		})
	}
}
