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

func TestCancelReservationReqUnmarshal(t *testing.T) {
	t.Run("rejects zero reservationId", func(t *testing.T) {
		var v CancelReservationReq
		err := json.Unmarshal([]byte(`{"reservationId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v CancelReservationReq
		err := json.Unmarshal([]byte(`{"reservationId":42}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ReservationID != 42 {
			t.Errorf("got %d, want %d", v.ReservationID, 42)
		}
	})
}

func TestCancelReservationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   CancelReservationReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   CancelReservationReq{ReservationID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects zero reservationId",
			input:   CancelReservationReq{ReservationID: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative reservationId",
			input:   CancelReservationReq{ReservationID: -1},
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

func TestCancelReservationConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v CancelReservationConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v CancelReservationConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != CancelReservationStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, CancelReservationStatusAccepted)
		}
	})
}

func TestCancelReservationConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   CancelReservationConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   CancelReservationConf{Status: CancelReservationStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   CancelReservationConf{Status: ""},
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

func TestChangeAvailabilityReqUnmarshal(t *testing.T) {
	t.Run("rejects invalid type", func(t *testing.T) {
		var v ChangeAvailabilityReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"type":"BadType"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ChangeAvailabilityReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"type":"Operative"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID != 1 {
			t.Errorf("got %d, want %d", v.ConnectorID, 1)
		}

		if v.Type != AvailabilityTypeOperative {
			t.Errorf("got %q, want %q", v.Type, AvailabilityTypeOperative)
		}
	})
}

func TestChangeAvailabilityReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ChangeAvailabilityReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ChangeAvailabilityReq{ConnectorID: 1, Type: AvailabilityTypeOperative},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   ChangeAvailabilityReq{ConnectorID: 0, Type: AvailabilityTypeOperative},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   ChangeAvailabilityReq{ConnectorID: -1, Type: AvailabilityTypeOperative},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects empty type",
			input:   ChangeAvailabilityReq{ConnectorID: 1, Type: ""},
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

func TestChangeAvailabilityConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v ChangeAvailabilityConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ChangeAvailabilityConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != AvailabilityStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, AvailabilityStatusAccepted)
		}
	})
}

func TestChangeAvailabilityConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ChangeAvailabilityConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ChangeAvailabilityConf{Status: AvailabilityStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   ChangeAvailabilityConf{Status: ""},
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

func TestChangeConfigurationReqUnmarshal(t *testing.T) {
	t.Run("rejects missing key", func(t *testing.T) {
		var v ChangeConfigurationReq
		err := json.Unmarshal([]byte(`{"value":"someValue"}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects oversized key", func(t *testing.T) {
		var v ChangeConfigurationReq
		err := json.Unmarshal([]byte(`{"key":"ThisKeyIsOverFiftyCharactersLongAndShouldBeRejected","value":"someValue"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("accepts empty value", func(t *testing.T) {
		var v ChangeConfigurationReq
		err := json.Unmarshal([]byte(`{"key":"someKey","value":""}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ChangeConfigurationReq
		err := json.Unmarshal([]byte(`{"key":"someKey","value":"someValue"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Key != "someKey" {
			t.Errorf("got %q, want %q", v.Key, "someKey")
		}

		if v.Value != "someValue" {
			t.Errorf("got %q, want %q", v.Value, "someValue")
		}
	})
}

func TestChangeConfigurationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ChangeConfigurationReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ChangeConfigurationReq{Key: "someKey", Value: "someValue"},
			wantErr: nil,
		},
		{
			name:    "accepts empty value",
			input:   ChangeConfigurationReq{Key: "someKey", Value: ""},
			wantErr: nil,
		},
		{
			name:    "rejects empty key",
			input:   ChangeConfigurationReq{Key: "", Value: "someValue"},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized key",
			input:   ChangeConfigurationReq{Key: "ThisKeyIsOverFiftyCharactersLongAndShouldBeRejected", Value: "someValue"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects oversized value",
			input:   ChangeConfigurationReq{Key: "someKey", Value: CiString500Type("x" + string(make([]byte, 500)))},
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

func TestChangeConfigurationConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v ChangeConfigurationConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ChangeConfigurationConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != ConfigurationStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, ConfigurationStatusAccepted)
		}
	})
}

func TestChangeConfigurationConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ChangeConfigurationConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ChangeConfigurationConf{Status: ConfigurationStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   ChangeConfigurationConf{Status: ""},
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

func TestClearCacheReqUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v ClearCacheReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestClearCacheConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v ClearCacheConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ClearCacheConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != ClearCacheStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, ClearCacheStatusAccepted)
		}
	})
}

func TestClearCacheConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ClearCacheConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ClearCacheConf{Status: ClearCacheStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   ClearCacheConf{Status: ""},
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

func TestClearChargingProfileReqUnmarshal(t *testing.T) {
	t.Run("rejects invalid chargingProfilePurpose", func(t *testing.T) {
		var v ClearChargingProfileReq
		err := json.Unmarshal([]byte(`{"chargingProfilePurpose":"BadPurpose"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with no fields", func(t *testing.T) {
		var v ClearChargingProfileReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v ClearChargingProfileReq
		err := json.Unmarshal([]byte(`{"id":1,"connectorId":0,"chargingProfilePurpose":"TxProfile","stackLevel":0}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestClearChargingProfileReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ClearChargingProfileReq
		wantErr error
	}{
		{
			name:    "accepts empty request",
			input:   ClearChargingProfileReq{},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   ClearChargingProfileReq{ConnectorID: ptr(int32(0))},
			wantErr: nil,
		},
		{
			name:    "accepts zero stackLevel",
			input:   ClearChargingProfileReq{StackLevel: ptr(int32(0))},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   ClearChargingProfileReq{ConnectorID: ptr(int32(-1))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative stackLevel",
			input:   ClearChargingProfileReq{StackLevel: ptr(int32(-1))},
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

func TestClearChargingProfileConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v ClearChargingProfileConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ClearChargingProfileConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != ClearChargingProfileStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, ClearChargingProfileStatusAccepted)
		}
	})
}

func TestClearChargingProfileConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ClearChargingProfileConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ClearChargingProfileConf{Status: ClearChargingProfileStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   ClearChargingProfileConf{Status: ""},
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
