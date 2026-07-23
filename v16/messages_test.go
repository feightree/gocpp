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
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestClearCacheReqValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := ClearCacheReq{}.Validate()

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

func TestDataTransferReqUnmarshal(t *testing.T) {
	t.Run("rejects oversized vendorId", func(t *testing.T) {
		var v DataTransferReq
		err := json.Unmarshal([]byte(`{"vendorId":"oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects non-printable ASCII in vendorId", func(t *testing.T) {
		var v DataTransferReq
		err := json.Unmarshal([]byte(`{"vendorId":"café"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects oversized messageId", func(t *testing.T) {
		var v DataTransferReq
		err := json.Unmarshal([]byte(`{"vendorId":"ACME","messageId":"ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required field", func(t *testing.T) {
		var v DataTransferReq
		err := json.Unmarshal([]byte(`{"vendorId":"ACME"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.VendorID != "ACME" {
			t.Errorf("got %q, want %q", v.VendorID, "ACME")
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v DataTransferReq
		err := json.Unmarshal([]byte(`{"vendorId":"ACME","messageId":"MsgType1","data":"payload"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.VendorID != "ACME" {
			t.Errorf("got %q, want %q", v.VendorID, "ACME")
		}

		if v.MessageID == nil || *v.MessageID != "MsgType1" {
			t.Errorf("got %v, want %q", v.MessageID, "MsgType1")
		}

		if v.Data == nil || *v.Data != "payload" {
			t.Errorf("got %v, want %q", v.Data, "payload")
		}
	})
}

func TestDataTransferReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   DataTransferReq
		wantErr error
	}{
		{
			name:    "rejects empty vendorId",
			input:   DataTransferReq{VendorID: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized vendorId",
			input:   DataTransferReq{VendorID: "oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects non-printable vendorId characters",
			input:   DataTransferReq{VendorID: "café"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects oversized messageId",
			input: DataTransferReq{
				VendorID:  "ACME",
				MessageID: ptr(CiString50Type("ThisStringIsOverFiftyCharactersLongAndShouldBeRejected")),
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts valid input with only required field",
			input:   DataTransferReq{VendorID: "ACME"},
			wantErr: nil,
		},
		{
			name: "accepts valid input with all fields",
			input: DataTransferReq{
				VendorID:  "ACME",
				MessageID: ptr(CiString50Type("MsgType1")),
				Data:      ptr("payload"),
			},
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

func TestDataTransferConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v DataTransferConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v DataTransferConf
		err := json.Unmarshal([]byte(`{"status":"Accepted","data":"payload"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != DataTransferStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, DataTransferStatusAccepted)
		}

		if v.Data == nil || *v.Data != "payload" {
			t.Errorf("got %v, want %q", v.Data, "payload")
		}
	})
}

func TestDataTransferConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   DataTransferConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   DataTransferConf{Status: DataTransferStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   DataTransferConf{Status: ""},
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

func TestDiagnosticsStatusNotificationReqUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v DiagnosticsStatusNotificationReq
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v DiagnosticsStatusNotificationReq
		err := json.Unmarshal([]byte(`{"status":"Uploaded"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != DiagnosticsStatusUploaded {
			t.Errorf("got %q, want %q", v.Status, DiagnosticsStatusUploaded)
		}
	})
}

func TestDiagnosticsStatusNotificationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   DiagnosticsStatusNotificationReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   DiagnosticsStatusNotificationReq{Status: DiagnosticsStatusUploaded},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   DiagnosticsStatusNotificationReq{Status: ""},
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

func TestDiagnosticsStatusNotificationConfUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v DiagnosticsStatusNotificationConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestDiagnosticsStatusNotificationConfValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := DiagnosticsStatusNotificationConf{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestFirmwareStatusNotificationReqUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v FirmwareStatusNotificationReq
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v FirmwareStatusNotificationReq
		err := json.Unmarshal([]byte(`{"status":"Installed"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != FirmwareStatusInstalled {
			t.Errorf("got %q, want %q", v.Status, FirmwareStatusInstalled)
		}
	})
}

func TestFirmwareStatusNotificationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   FirmwareStatusNotificationReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   FirmwareStatusNotificationReq{Status: FirmwareStatusInstalled},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   FirmwareStatusNotificationReq{Status: ""},
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

func TestFirmwareStatusNotificationConfUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v FirmwareStatusNotificationConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestFirmwareStatusNotificationConfValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := FirmwareStatusNotificationConf{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestGetCompositeScheduleReqUnmarshal(t *testing.T) {
	t.Run("rejects invalid chargingRateUnit", func(t *testing.T) {
		var v GetCompositeScheduleReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"duration":60,"chargingRateUnit":"BadUnit"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v GetCompositeScheduleReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"duration":60}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID != 1 {
			t.Errorf("got %d, want %d", v.ConnectorID, 1)
		}

		if v.Duration != 60 {
			t.Errorf("got %d, want %d", v.Duration, 60)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v GetCompositeScheduleReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"duration":60,"chargingRateUnit":"W"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ChargingRateUnit == nil || *v.ChargingRateUnit != ChargingRateUnitTypeW {
			t.Errorf("got %v, want %q", v.ChargingRateUnit, ChargingRateUnitTypeW)
		}
	})
}

func TestGetCompositeScheduleReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetCompositeScheduleReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   GetCompositeScheduleReq{ConnectorID: 1, Duration: 60},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   GetCompositeScheduleReq{ConnectorID: 0, Duration: 60},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   GetCompositeScheduleReq{ConnectorID: -1, Duration: 60},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero duration",
			input:   GetCompositeScheduleReq{ConnectorID: 1, Duration: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative duration",
			input:   GetCompositeScheduleReq{ConnectorID: 1, Duration: -1},
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

func TestGetCompositeScheduleConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v GetCompositeScheduleConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects invalid nested chargingSchedule", func(t *testing.T) {
		var v GetCompositeScheduleConf
		err := json.Unmarshal([]byte(`{"status":"Accepted","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[]}}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("valid input with only required field", func(t *testing.T) {
		var v GetCompositeScheduleConf
		err := json.Unmarshal([]byte(`{"status":"Rejected"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != GetCompositeScheduleStatusRejected {
			t.Errorf("got %q, want %q", v.Status, GetCompositeScheduleStatusRejected)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v GetCompositeScheduleConf
		err := json.Unmarshal([]byte(`{"status":"Accepted","connectorId":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":10}]}}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID == nil || *v.ConnectorID != 1 {
			t.Errorf("got %v, want %d", v.ConnectorID, 1)
		}

		if v.ScheduleStart == nil || !v.ScheduleStart.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.ScheduleStart, "2024-01-01T00:00:00Z")
		}

		if v.ChargingSchedule == nil {
			t.Errorf("got nil, want non-nil chargingSchedule")
		}
	})
}

func TestGetCompositeScheduleConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetCompositeScheduleConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   GetCompositeScheduleConf{Status: GetCompositeScheduleStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   GetCompositeScheduleConf{Status: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "accepts positive connectorId",
			input:   GetCompositeScheduleConf{Status: GetCompositeScheduleStatusAccepted, ConnectorID: ptr(int32(1))},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   GetCompositeScheduleConf{Status: GetCompositeScheduleStatusAccepted, ConnectorID: ptr(int32(0))},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   GetCompositeScheduleConf{Status: GetCompositeScheduleStatusAccepted, ConnectorID: ptr(int32(-1))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts zero-value scheduleStart",
			input:   GetCompositeScheduleConf{Status: GetCompositeScheduleStatusAccepted, ScheduleStart: &time.Time{}},
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

func TestGetConfigurationReqUnmarshal(t *testing.T) {
	t.Run("rejects oversized key", func(t *testing.T) {
		var v GetConfigurationReq
		err := json.Unmarshal([]byte(`{"key":["ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"]}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with no keys", func(t *testing.T) {
		var v GetConfigurationReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Key != nil {
			t.Errorf("got %v, want nil", v.Key)
		}
	})

	t.Run("valid input with keys", func(t *testing.T) {
		var v GetConfigurationReq
		err := json.Unmarshal([]byte(`{"key":["ConfigKey1","ConfigKey2"]}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(v.Key) != 2 || v.Key[0] != "ConfigKey1" || v.Key[1] != "ConfigKey2" {
			t.Errorf("got %v, want %v", v.Key, []string{"ConfigKey1", "ConfigKey2"})
		}
	})
}

func TestGetConfigurationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetConfigurationReq
		wantErr error
	}{
		{
			name:    "accepts nil keys",
			input:   GetConfigurationReq{},
			wantErr: nil,
		},
		{
			name:    "accepts valid keys",
			input:   GetConfigurationReq{Key: []CiString50Type{"ConfigKey1", "ConfigKey2"}},
			wantErr: nil,
		},
		{
			name:    "rejects oversized key",
			input:   GetConfigurationReq{Key: []CiString50Type{"ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"}},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects non-printable ASCII in key",
			input:   GetConfigurationReq{Key: []CiString50Type{"café"}},
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

func TestGetConfigurationConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid nested configurationKey", func(t *testing.T) {
		var v GetConfigurationConf
		err := json.Unmarshal([]byte(`{"configurationKey":[{"key":"","readonly":false}]}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects oversized unknownKey", func(t *testing.T) {
		var v GetConfigurationConf
		err := json.Unmarshal([]byte(`{"unknownKey":["ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"]}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with no fields", func(t *testing.T) {
		var v GetConfigurationConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v GetConfigurationConf
		err := json.Unmarshal([]byte(`{"configurationKey":[{"key":"ConfigKey1","readonly":false,"value":"SomeValue"}],"unknownKey":["UnknownKey1"]}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(v.ConfigurationKey) != 1 || v.ConfigurationKey[0].Key != "ConfigKey1" {
			t.Errorf("got %v, want key %q", v.ConfigurationKey, "ConfigKey1")
		}

		if len(v.UnknownKey) != 1 || v.UnknownKey[0] != "UnknownKey1" {
			t.Errorf("got %v, want %v", v.UnknownKey, []string{"UnknownKey1"})
		}
	})
}

func TestGetConfigurationConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetConfigurationConf
		wantErr error
	}{
		{
			name:    "accepts empty conf",
			input:   GetConfigurationConf{},
			wantErr: nil,
		},
		{
			name: "accepts valid configurationKey and unknownKey",
			input: GetConfigurationConf{
				ConfigurationKey: []KeyValue{{Key: "ConfigKey1", Readonly: false}},
				UnknownKey:       []CiString50Type{"UnknownKey1"},
			},
			wantErr: nil,
		},
		{
			name: "rejects invalid configurationKey entry",
			input: GetConfigurationConf{
				ConfigurationKey: []KeyValue{{Key: "", Readonly: false}},
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects oversized unknownKey entry",
			input: GetConfigurationConf{
				UnknownKey: []CiString50Type{"ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"},
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

func TestGetDiagnosticsReqUnmarshal(t *testing.T) {
	t.Run("rejects missing location", func(t *testing.T) {
		var v GetDiagnosticsReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects negative retries", func(t *testing.T) {
		var v GetDiagnosticsReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retries":-1}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects zero retryInterval", func(t *testing.T) {
		var v GetDiagnosticsReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retryInterval":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required field", func(t *testing.T) {
		var v GetDiagnosticsReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Location != "ftp://example.com/" {
			t.Errorf("got %q, want %q", v.Location, "ftp://example.com/")
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v GetDiagnosticsReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retries":3,"retryInterval":60,"startTime":"2024-01-01T00:00:00Z","stopTime":"2024-01-02T00:00:00Z"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Retries == nil || *v.Retries != 3 {
			t.Errorf("got %v, want %d", v.Retries, 3)
		}

		if v.RetryInterval == nil || *v.RetryInterval != 60 {
			t.Errorf("got %v, want %d", v.RetryInterval, 60)
		}

		if v.StartTime == nil || !v.StartTime.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.StartTime, "2024-01-01T00:00:00Z")
		}

		if v.StopTime == nil || !v.StopTime.Equal(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.StopTime, "2024-01-02T00:00:00Z")
		}
	})
}

func TestGetDiagnosticsReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetDiagnosticsReq
		wantErr error
	}{
		{
			name:    "accepts valid input with only required field",
			input:   GetDiagnosticsReq{Location: "ftp://example.com/"},
			wantErr: nil,
		},
		{
			name: "accepts valid input with all fields",
			input: GetDiagnosticsReq{
				Location:      "ftp://example.com/",
				Retries:       ptr(int32(3)),
				RetryInterval: ptr(int32(60)),
				StartTime:     ptr(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
				StopTime:      ptr(time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
			wantErr: nil,
		},
		{
			name:    "rejects empty location",
			input:   GetDiagnosticsReq{Location: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "accepts zero retries",
			input:   GetDiagnosticsReq{Location: "ftp://example.com/", Retries: ptr(int32(0))},
			wantErr: nil,
		},
		{
			name:    "rejects negative retries",
			input:   GetDiagnosticsReq{Location: "ftp://example.com/", Retries: ptr(int32(-1))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero retryInterval",
			input:   GetDiagnosticsReq{Location: "ftp://example.com/", RetryInterval: ptr(int32(0))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative retryInterval",
			input:   GetDiagnosticsReq{Location: "ftp://example.com/", RetryInterval: ptr(int32(-1))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts positive retryInterval",
			input:   GetDiagnosticsReq{Location: "ftp://example.com/", RetryInterval: ptr(int32(60))},
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

func TestGetDiagnosticsConfUnmarshal(t *testing.T) {
	t.Run("rejects oversized fileName", func(t *testing.T) {
		var v GetDiagnosticsConf
		err := json.Unmarshal([]byte(`{"fileName":"oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with no fields", func(t *testing.T) {
		var v GetDiagnosticsConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.FileName != nil {
			t.Errorf("got %v, want nil", v.FileName)
		}
	})

	t.Run("valid input with fileName", func(t *testing.T) {
		var v GetDiagnosticsConf
		err := json.Unmarshal([]byte(`{"fileName":"diagnostics-2024-01-01.log"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.FileName == nil || *v.FileName != "diagnostics-2024-01-01.log" {
			t.Errorf("got %v, want %q", v.FileName, "diagnostics-2024-01-01.log")
		}
	})
}

func TestGetDiagnosticsConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetDiagnosticsConf
		wantErr error
	}{
		{
			name:    "accepts empty conf",
			input:   GetDiagnosticsConf{},
			wantErr: nil,
		},
		{
			name:    "accepts valid fileName",
			input:   GetDiagnosticsConf{FileName: ptr(CiString255Type("diagnostics-2024-01-01.log"))},
			wantErr: nil,
		},
		{
			name:    "rejects oversized fileName",
			input:   GetDiagnosticsConf{FileName: ptr(CiString255Type("oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects non-printable ASCII in fileName",
			input:   GetDiagnosticsConf{FileName: ptr(CiString255Type("café"))},
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

func TestGetLocalListVersionReqUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v GetLocalListVersionReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestGetLocalListVersionReqValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := GetLocalListVersionReq{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestGetLocalListVersionConfUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v GetLocalListVersionConf
		err := json.Unmarshal([]byte(`{"listVersion":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ListVersion != 1 {
			t.Errorf("got %d, want %d", v.ListVersion, 1)
		}
	})

	t.Run("valid input with sentinel listVersion", func(t *testing.T) {
		var v GetLocalListVersionConf
		err := json.Unmarshal([]byte(`{"listVersion":-1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ListVersion != -1 {
			t.Errorf("got %d, want %d", v.ListVersion, -1)
		}
	})
}

func TestGetLocalListVersionConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   GetLocalListVersionConf
		wantErr error
	}{
		{
			name:    "accepts positive listVersion",
			input:   GetLocalListVersionConf{ListVersion: 1},
			wantErr: nil,
		},
		{
			name:    "accepts zero listVersion",
			input:   GetLocalListVersionConf{ListVersion: 0},
			wantErr: nil,
		},
		{
			name:    "accepts sentinel listVersion of -1",
			input:   GetLocalListVersionConf{ListVersion: -1},
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

func TestHeartbeatReqUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v HeartbeatReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestHeartbeatReqValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := HeartbeatReq{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestHeartbeatConfUnmarshal(t *testing.T) {
	t.Run("rejects missing currentTime", func(t *testing.T) {
		var v HeartbeatConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v HeartbeatConf
		err := json.Unmarshal([]byte(`{"currentTime":"2024-01-01T00:00:00Z"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !v.CurrentTime.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.CurrentTime, "2024-01-01T00:00:00Z")
		}
	})
}

func TestHeartbeatConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   HeartbeatConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   HeartbeatConf{CurrentTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			wantErr: nil,
		},
		{
			name:    "rejects zero currentTime",
			input:   HeartbeatConf{},
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

func TestMeterValuesReqUnmarshal(t *testing.T) {
	t.Run("rejects negative connectorId", func(t *testing.T) {
		var v MeterValuesReq
		err := json.Unmarshal([]byte(`{"connectorId":-1,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":"100"}]}]}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects empty meterValue array", func(t *testing.T) {
		var v MeterValuesReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"meterValue":[]}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects invalid nested meterValue", func(t *testing.T) {
		var v MeterValuesReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[]}]}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v MeterValuesReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":"100"}]}]}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID != 0 {
			t.Errorf("got %d, want %d", v.ConnectorID, 0)
		}

		if len(v.MeterValue) != 1 || len(v.MeterValue[0].SampledValue) != 1 || v.MeterValue[0].SampledValue[0].Value != "100" {
			t.Errorf("got %v, want a single meterValue with sampledValue[0].value %q", v.MeterValue, "100")
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v MeterValuesReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"transactionId":42,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":"100"}]}]}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.TransactionID == nil || *v.TransactionID != 42 {
			t.Errorf("got %v, want %d", v.TransactionID, 42)
		}
	})
}

func TestMeterValuesReqValidate(t *testing.T) {
	validMeterValue := MeterValue{
		Timestamp:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		SampledValue: []SampledValue{{Value: "100"}},
	}

	tests := []struct {
		name    string
		input   MeterValuesReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   MeterValuesReq{ConnectorID: 1, MeterValue: []MeterValue{validMeterValue}},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   MeterValuesReq{ConnectorID: 0, MeterValue: []MeterValue{validMeterValue}},
			wantErr: nil,
		},
		{
			name:    "accepts optional transactionId",
			input:   MeterValuesReq{ConnectorID: 1, TransactionID: ptr(int32(42)), MeterValue: []MeterValue{validMeterValue}},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   MeterValuesReq{ConnectorID: -1, MeterValue: []MeterValue{validMeterValue}},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects empty meterValue array",
			input:   MeterValuesReq{ConnectorID: 1, MeterValue: []MeterValue{}},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects invalid meterValue entry",
			input:   MeterValuesReq{ConnectorID: 1, MeterValue: []MeterValue{{Timestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}}},
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

func TestMeterValuesConfUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v MeterValuesConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestMeterValuesConfValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := MeterValuesConf{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestRemoteStartTransactionReqUnmarshal(t *testing.T) {
	t.Run("rejects zero connectorId", func(t *testing.T) {
		var v RemoteStartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"idTag":"ABC123"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects missing idTag", func(t *testing.T) {
		var v RemoteStartTransactionReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects chargingProfile with non-TxProfile purpose", func(t *testing.T) {
		var v RemoteStartTransactionReq
		err := json.Unmarshal([]byte(`{"idTag":"ABC123","chargingProfile":{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"ChargePointMaxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v RemoteStartTransactionReq
		err := json.Unmarshal([]byte(`{"idTag":"ABC123"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTag != "ABC123" {
			t.Errorf("got %q, want %q", v.IDTag, "ABC123")
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v RemoteStartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"idTag":"ABC123","chargingProfile":{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID == nil || *v.ConnectorID != 1 {
			t.Errorf("got %v, want %d", v.ConnectorID, 1)
		}

		if v.ChargingProfile == nil || v.ChargingProfile.ChargingProfilePurpose != ChargingProfilePurposeTypeTxProfile {
			t.Errorf("got %v, want chargingProfilePurpose %q", v.ChargingProfile, ChargingProfilePurposeTypeTxProfile)
		}
	})
}

func TestRemoteStartTransactionReqValidate(t *testing.T) {
	validProfile := ChargingProfile{
		ChargingProfileID:      1,
		StackLevel:             0,
		ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
		ChargingProfileKind:    ChargingProfileKindTypeAbsolute,
		ChargingSchedule: ChargingSchedule{
			ChargingRateUnit:       ChargingRateUnitTypeW,
			ChargingSchedulePeriod: []ChargingSchedulePeriod{{StartPeriod: 0, Limit: 32}},
		},
	}

	wrongPurposeProfile := validProfile
	wrongPurposeProfile.ChargingProfilePurpose = ChargingProfilePurposeTypeChargePointMaxProfile

	invalidProfile := validProfile
	invalidProfile.StackLevel = -1

	tests := []struct {
		name    string
		input   RemoteStartTransactionReq
		wantErr error
	}{
		{
			name:    "accepts valid input with only required fields",
			input:   RemoteStartTransactionReq{IDTag: "ABC123"},
			wantErr: nil,
		},
		{
			name:    "accepts positive connectorId",
			input:   RemoteStartTransactionReq{ConnectorID: ptr(int32(1)), IDTag: "ABC123"},
			wantErr: nil,
		},
		{
			name:    "rejects zero connectorId",
			input:   RemoteStartTransactionReq{ConnectorID: ptr(int32(0)), IDTag: "ABC123"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative connectorId",
			input:   RemoteStartTransactionReq{ConnectorID: ptr(int32(-1)), IDTag: "ABC123"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects empty idTag",
			input:   RemoteStartTransactionReq{IDTag: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized idTag",
			input:   RemoteStartTransactionReq{IDTag: "ThisIsOver20Characters"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts valid chargingProfile with TxProfile purpose",
			input:   RemoteStartTransactionReq{IDTag: "ABC123", ChargingProfile: &validProfile},
			wantErr: nil,
		},
		{
			name:    "rejects chargingProfile with non-TxProfile purpose",
			input:   RemoteStartTransactionReq{IDTag: "ABC123", ChargingProfile: &wrongPurposeProfile},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects invalid nested chargingProfile",
			input:   RemoteStartTransactionReq{IDTag: "ABC123", ChargingProfile: &invalidProfile},
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

func TestRemoteStartTransactionConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v RemoteStartTransactionConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v RemoteStartTransactionConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != RemoteStartStopStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, RemoteStartStopStatusAccepted)
		}
	})
}

func TestRemoteStartTransactionConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   RemoteStartTransactionConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   RemoteStartTransactionConf{Status: RemoteStartStopStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   RemoteStartTransactionConf{Status: ""},
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

func TestRemoteStopTransactionReqUnmarshal(t *testing.T) {
	t.Run("rejects zero transactionId", func(t *testing.T) {
		var v RemoteStopTransactionReq
		err := json.Unmarshal([]byte(`{"transactionId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v RemoteStopTransactionReq
		err := json.Unmarshal([]byte(`{"transactionId":42}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.TransactionID != 42 {
			t.Errorf("got %d, want %d", v.TransactionID, 42)
		}
	})
}

func TestRemoteStopTransactionReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   RemoteStopTransactionReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   RemoteStopTransactionReq{TransactionID: 42},
			wantErr: nil,
		},
		{
			name:    "rejects zero transactionId",
			input:   RemoteStopTransactionReq{TransactionID: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative transactionId",
			input:   RemoteStopTransactionReq{TransactionID: -1},
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

func TestRemoteStopTransactionConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v RemoteStopTransactionConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v RemoteStopTransactionConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != RemoteStartStopStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, RemoteStartStopStatusAccepted)
		}
	})
}

func TestRemoteStopTransactionConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   RemoteStopTransactionConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   RemoteStopTransactionConf{Status: RemoteStartStopStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   RemoteStopTransactionConf{Status: ""},
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

func TestReserveNowReqUnmarshal(t *testing.T) {
	t.Run("rejects negative connectorId", func(t *testing.T) {
		var v ReserveNowReq
		err := json.Unmarshal([]byte(`{"connectorId":-1,"expiryDate":"2024-01-01T00:00:00Z","idTag":"ABC123","reservationId":1}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects missing expiryDate", func(t *testing.T) {
		var v ReserveNowReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"idTag":"ABC123","reservationId":1}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects missing idTag", func(t *testing.T) {
		var v ReserveNowReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"expiryDate":"2024-01-01T00:00:00Z","reservationId":1}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects zero reservationId", func(t *testing.T) {
		var v ReserveNowReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"expiryDate":"2024-01-01T00:00:00Z","idTag":"ABC123","reservationId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v ReserveNowReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"expiryDate":"2024-01-01T00:00:00Z","idTag":"ABC123","reservationId":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTag != "ABC123" {
			t.Errorf("got %q, want %q", v.IDTag, "ABC123")
		}

		if v.ReservationID != 1 {
			t.Errorf("got %d, want %d", v.ReservationID, 1)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v ReserveNowReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"expiryDate":"2024-01-01T00:00:00Z","idTag":"ABC123","parentIdTag":"PARENT01","reservationId":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ParentIDTag == nil || *v.ParentIDTag != "PARENT01" {
			t.Errorf("got %v, want %q", v.ParentIDTag, "PARENT01")
		}

		if !v.ExpiryDate.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.ExpiryDate, "2024-01-01T00:00:00Z")
		}
	})
}

func TestReserveNowReqValidate(t *testing.T) {
	validExpiry := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		input   ReserveNowReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ReserveNowReq{ConnectorID: 1, ExpiryDate: validExpiry, IDTag: "ABC123", ReservationID: 1},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "ABC123", ReservationID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   ReserveNowReq{ConnectorID: -1, ExpiryDate: validExpiry, IDTag: "ABC123", ReservationID: 1},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero expiryDate",
			input:   ReserveNowReq{ConnectorID: 0, IDTag: "ABC123", ReservationID: 1},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty idTag",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "", ReservationID: 1},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized idTag",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "ThisIsOver20Characters", ReservationID: 1},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts valid parentIdTag",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "ABC123", ParentIDTag: ptr(IDToken("PARENT01")), ReservationID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects oversized parentIdTag",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "ABC123", ParentIDTag: ptr(IDToken("ThisIsOver20Characters")), ReservationID: 1},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero reservationId",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "ABC123", ReservationID: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative reservationId",
			input:   ReserveNowReq{ConnectorID: 0, ExpiryDate: validExpiry, IDTag: "ABC123", ReservationID: -1},
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

func TestReserveNowConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v ReserveNowConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ReserveNowConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != ReservationStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, ReservationStatusAccepted)
		}
	})
}

func TestReserveNowConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ReserveNowConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ReserveNowConf{Status: ReservationStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   ReserveNowConf{Status: ""},
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

func TestResetReqUnmarshal(t *testing.T) {
	t.Run("rejects invalid type", func(t *testing.T) {
		var v ResetReq
		err := json.Unmarshal([]byte(`{"type":"BadType"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ResetReq
		err := json.Unmarshal([]byte(`{"type":"Hard"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Type != ResetTypeHard {
			t.Errorf("got %q, want %q", v.Type, ResetTypeHard)
		}
	})
}

func TestResetReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ResetReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ResetReq{Type: ResetTypeSoft},
			wantErr: nil,
		},
		{
			name:    "rejects empty type",
			input:   ResetReq{Type: ""},
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

func TestResetConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v ResetConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v ResetConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != ResetStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, ResetStatusAccepted)
		}
	})
}

func TestResetConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   ResetConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   ResetConf{Status: ResetStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   ResetConf{Status: ""},
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

func TestSendLocalListReqUnmarshal(t *testing.T) {
	t.Run("rejects missing updateType", func(t *testing.T) {
		var v SendLocalListReq
		err := json.Unmarshal([]byte(`{"listVersion":1}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects invalid nested localAuthorizationList entry", func(t *testing.T) {
		var v SendLocalListReq
		err := json.Unmarshal([]byte(`{"listVersion":1,"updateType":"Full","localAuthorizationList":[{"idTag":""}]}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects duplicate idTag", func(t *testing.T) {
		var v SendLocalListReq
		err := json.Unmarshal([]byte(`{"listVersion":1,"updateType":"Full","localAuthorizationList":[{"idTag":"ABC123"},{"idTag":"abc123"}]}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v SendLocalListReq
		err := json.Unmarshal([]byte(`{"listVersion":1,"updateType":"Full"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ListVersion != 1 {
			t.Errorf("got %d, want %d", v.ListVersion, 1)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v SendLocalListReq
		err := json.Unmarshal([]byte(`{"listVersion":1,"updateType":"Full","localAuthorizationList":[{"idTag":"ABC123"},{"idTag":"DEF456"}]}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(v.LocalAuthorizationList) != 2 || v.LocalAuthorizationList[0].IDTag != "ABC123" || v.LocalAuthorizationList[1].IDTag != "DEF456" {
			t.Errorf("got %v, want idTags ABC123 and DEF456", v.LocalAuthorizationList)
		}
	})
}

func TestSendLocalListReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   SendLocalListReq
		wantErr error
	}{
		{
			name:    "accepts valid input with only required fields",
			input:   SendLocalListReq{ListVersion: 1, UpdateType: UpdateTypeFull},
			wantErr: nil,
		},
		{
			name: "accepts unique idTags",
			input: SendLocalListReq{
				ListVersion:            1,
				UpdateType:             UpdateTypeFull,
				LocalAuthorizationList: []AuthorizationData{{IDTag: "ABC123"}, {IDTag: "DEF456"}},
			},
			wantErr: nil,
		},
		{
			name:    "rejects missing updateType",
			input:   SendLocalListReq{ListVersion: 1},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects invalid localAuthorizationList entry",
			input: SendLocalListReq{
				ListVersion:            1,
				UpdateType:             UpdateTypeFull,
				LocalAuthorizationList: []AuthorizationData{{IDTag: ""}},
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects duplicate idTag",
			input: SendLocalListReq{
				ListVersion:            1,
				UpdateType:             UpdateTypeFull,
				LocalAuthorizationList: []AuthorizationData{{IDTag: "ABC123"}, {IDTag: "ABC123"}},
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects duplicate idTag case-insensitively",
			input: SendLocalListReq{
				ListVersion:            1,
				UpdateType:             UpdateTypeFull,
				LocalAuthorizationList: []AuthorizationData{{IDTag: "ABC123"}, {IDTag: "abc123"}},
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

func TestSendLocalListConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v SendLocalListConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v SendLocalListConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != UpdateStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, UpdateStatusAccepted)
		}
	})
}

func TestSendLocalListConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   SendLocalListConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   SendLocalListConf{Status: UpdateStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   SendLocalListConf{Status: ""},
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

func TestSetChargingProfileReqUnmarshal(t *testing.T) {
	t.Run("rejects negative connectorId", func(t *testing.T) {
		var v SetChargingProfileReq
		err := json.Unmarshal([]byte(`{"connectorId":-1,"csChargingProfiles":{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects invalid nested csChargingProfiles", func(t *testing.T) {
		var v SetChargingProfileReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"csChargingProfiles":{"chargingProfileId":1,"stackLevel":-1,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects ChargePointMaxProfile with non-zero connectorId", func(t *testing.T) {
		var v SetChargingProfileReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"csChargingProfiles":{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"ChargePointMaxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v SetChargingProfileReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"csChargingProfiles":{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID != 1 {
			t.Errorf("got %d, want %d", v.ConnectorID, 1)
		}

		if v.CsChargingProfiles.ChargingProfilePurpose != ChargingProfilePurposeTypeTxProfile {
			t.Errorf("got %q, want %q", v.CsChargingProfiles.ChargingProfilePurpose, ChargingProfilePurposeTypeTxProfile)
		}
	})

	t.Run("valid input with ChargePointMaxProfile at connectorId 0", func(t *testing.T) {
		var v SetChargingProfileReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"csChargingProfiles":{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"ChargePointMaxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestSetChargingProfileReqValidate(t *testing.T) {
	txProfile := ChargingProfile{
		ChargingProfileID:      1,
		StackLevel:             0,
		ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
		ChargingProfileKind:    ChargingProfileKindTypeAbsolute,
		ChargingSchedule: ChargingSchedule{
			ChargingRateUnit:       ChargingRateUnitTypeW,
			ChargingSchedulePeriod: []ChargingSchedulePeriod{{StartPeriod: 0, Limit: 32}},
		},
	}

	maxProfile := txProfile
	maxProfile.ChargingProfilePurpose = ChargingProfilePurposeTypeChargePointMaxProfile

	invalidProfile := txProfile
	invalidProfile.StackLevel = -1

	tests := []struct {
		name    string
		input   SetChargingProfileReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   SetChargingProfileReq{ConnectorID: 1, CsChargingProfiles: txProfile},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   SetChargingProfileReq{ConnectorID: 0, CsChargingProfiles: txProfile},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   SetChargingProfileReq{ConnectorID: -1, CsChargingProfiles: txProfile},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects invalid nested csChargingProfiles",
			input:   SetChargingProfileReq{ConnectorID: 0, CsChargingProfiles: invalidProfile},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts ChargePointMaxProfile at connectorId 0",
			input:   SetChargingProfileReq{ConnectorID: 0, CsChargingProfiles: maxProfile},
			wantErr: nil,
		},
		{
			name:    "rejects ChargePointMaxProfile at non-zero connectorId",
			input:   SetChargingProfileReq{ConnectorID: 1, CsChargingProfiles: maxProfile},
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

func TestSetChargingProfileConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v SetChargingProfileConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v SetChargingProfileConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != ChargingProfileStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, ChargingProfileStatusAccepted)
		}
	})
}

func TestSetChargingProfileConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   SetChargingProfileConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   SetChargingProfileConf{Status: ChargingProfileStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   SetChargingProfileConf{Status: ""},
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

func TestStartTransactionReqUnmarshal(t *testing.T) {
	t.Run("rejects zero connectorId", func(t *testing.T) {
		var v StartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"idTag":"ABC123","meterStart":0,"timestamp":"2024-01-01T00:00:00Z"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects negative meterStart", func(t *testing.T) {
		var v StartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"idTag":"ABC123","meterStart":-1,"timestamp":"2024-01-01T00:00:00Z"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects missing timestamp", func(t *testing.T) {
		var v StartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"idTag":"ABC123","meterStart":0}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v StartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"idTag":"ABC123","meterStart":0,"timestamp":"2024-01-01T00:00:00Z"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID != 1 {
			t.Errorf("got %d, want %d", v.ConnectorID, 1)
		}

		if !v.Timestamp.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.Timestamp, "2024-01-01T00:00:00Z")
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v StartTransactionReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"idTag":"ABC123","meterStart":100,"reservationId":5,"timestamp":"2024-01-01T00:00:00Z"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ReservationID == nil || *v.ReservationID != 5 {
			t.Errorf("got %v, want %d", v.ReservationID, 5)
		}
	})
}

func TestStartTransactionReqValidate(t *testing.T) {
	validTimestamp := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		input   StartTransactionReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "ABC123", MeterStart: 0, Timestamp: validTimestamp},
			wantErr: nil,
		},
		{
			name:    "rejects zero connectorId",
			input:   StartTransactionReq{ConnectorID: 0, IDTag: "ABC123", MeterStart: 0, Timestamp: validTimestamp},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative connectorId",
			input:   StartTransactionReq{ConnectorID: -1, IDTag: "ABC123", MeterStart: 0, Timestamp: validTimestamp},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects empty idTag",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "", MeterStart: 0, Timestamp: validTimestamp},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects oversized idTag",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "ThisIsOver20Characters", MeterStart: 0, Timestamp: validTimestamp},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts zero meterStart",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "ABC123", MeterStart: 0, Timestamp: validTimestamp},
			wantErr: nil,
		},
		{
			name:    "rejects negative meterStart",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "ABC123", MeterStart: -1, Timestamp: validTimestamp},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero timestamp",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "ABC123", MeterStart: 0},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "accepts optional reservationId",
			input:   StartTransactionReq{ConnectorID: 1, IDTag: "ABC123", MeterStart: 0, ReservationID: ptr(int32(5)), Timestamp: validTimestamp},
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

func TestStartTransactionConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid nested idTagInfo", func(t *testing.T) {
		var v StartTransactionConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{},"transactionId":1}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects zero transactionId", func(t *testing.T) {
		var v StartTransactionConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{"status":"Accepted"},"transactionId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v StartTransactionConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{"status":"Accepted"},"transactionId":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTagInfo.Status != AuthorizationStatusAccepted {
			t.Errorf("got %q, want %q", v.IDTagInfo.Status, AuthorizationStatusAccepted)
		}

		if v.TransactionID != 1 {
			t.Errorf("got %d, want %d", v.TransactionID, 1)
		}
	})
}

func TestStartTransactionConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   StartTransactionConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   StartTransactionConf{IDTagInfo: IDTagInfo{Status: AuthorizationStatusAccepted}, TransactionID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects invalid idTagInfo",
			input:   StartTransactionConf{IDTagInfo: IDTagInfo{Status: ""}, TransactionID: 1},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects zero transactionId",
			input:   StartTransactionConf{IDTagInfo: IDTagInfo{Status: AuthorizationStatusAccepted}, TransactionID: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative transactionId",
			input:   StartTransactionConf{IDTagInfo: IDTagInfo{Status: AuthorizationStatusAccepted}, TransactionID: -1},
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

func TestStatusNotificationReqUnmarshal(t *testing.T) {
	t.Run("rejects negative connectorId", func(t *testing.T) {
		var v StatusNotificationReq
		err := json.Unmarshal([]byte(`{"connectorId":-1,"errorCode":"NoError","status":"Available"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects invalid errorCode", func(t *testing.T) {
		var v StatusNotificationReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"errorCode":"BadError","status":"Available"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects invalid status", func(t *testing.T) {
		var v StatusNotificationReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"errorCode":"NoError","status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects oversized info", func(t *testing.T) {
		var v StatusNotificationReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"errorCode":"NoError","status":"Available","info":"ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v StatusNotificationReq
		err := json.Unmarshal([]byte(`{"connectorId":0,"errorCode":"NoError","status":"Available"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ErrorCode != ChargePointErrorCodeNoError {
			t.Errorf("got %q, want %q", v.ErrorCode, ChargePointErrorCodeNoError)
		}

		if v.Status != ChargePointStatusAvailable {
			t.Errorf("got %q, want %q", v.Status, ChargePointStatusAvailable)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v StatusNotificationReq
		err := json.Unmarshal([]byte(`{"connectorId":1,"errorCode":"NoError","info":"info text","status":"Available","timestamp":"2024-01-01T00:00:00Z","vendorId":"ACME","vendorErrorCode":"VE01"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Info == nil || *v.Info != "info text" {
			t.Errorf("got %v, want %q", v.Info, "info text")
		}

		if v.Timestamp == nil || !v.Timestamp.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.Timestamp, "2024-01-01T00:00:00Z")
		}

		if v.VendorID == nil || *v.VendorID != "ACME" {
			t.Errorf("got %v, want %q", v.VendorID, "ACME")
		}

		if v.VendorErrorCode == nil || *v.VendorErrorCode != "VE01" {
			t.Errorf("got %v, want %q", v.VendorErrorCode, "VE01")
		}
	})
}

func TestStatusNotificationReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   StatusNotificationReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   StatusNotificationReq{ConnectorID: 1, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable},
			wantErr: nil,
		},
		{
			name:    "accepts zero connectorId",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable},
			wantErr: nil,
		},
		{
			name:    "rejects negative connectorId",
			input:   StatusNotificationReq{ConnectorID: -1, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects empty errorCode",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: "", Status: ChargePointStatusAvailable},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty status",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ""},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "accepts valid info",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable, Info: ptr(CiString50Type("info text"))},
			wantErr: nil,
		},
		{
			name:    "rejects oversized info",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable, Info: ptr(CiString50Type("ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts valid vendorId",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable, VendorID: ptr(CiString255Type("ACME"))},
			wantErr: nil,
		},
		{
			name:    "rejects non-printable vendorId",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable, VendorID: ptr(CiString255Type("café"))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts valid vendorErrorCode",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable, VendorErrorCode: ptr(CiString50Type("VE01"))},
			wantErr: nil,
		},
		{
			name:    "rejects oversized vendorErrorCode",
			input:   StatusNotificationReq{ConnectorID: 0, ErrorCode: ChargePointErrorCodeNoError, Status: ChargePointStatusAvailable, VendorErrorCode: ptr(CiString50Type("ThisStringIsOverFiftyCharactersLongAndShouldBeRejected"))},
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

func TestStatusNotificationConfUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v StatusNotificationConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestStatusNotificationConfValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := StatusNotificationConf{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestStopTransactionReqUnmarshal(t *testing.T) {
	t.Run("rejects negative meterStop", func(t *testing.T) {
		var v StopTransactionReq
		err := json.Unmarshal([]byte(`{"meterStop":-1,"timestamp":"2024-01-01T00:00:00Z","transactionId":1}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects missing timestamp", func(t *testing.T) {
		var v StopTransactionReq
		err := json.Unmarshal([]byte(`{"meterStop":0,"transactionId":1}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects zero transactionId", func(t *testing.T) {
		var v StopTransactionReq
		err := json.Unmarshal([]byte(`{"meterStop":0,"timestamp":"2024-01-01T00:00:00Z","transactionId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects invalid nested transactionData", func(t *testing.T) {
		var v StopTransactionReq
		err := json.Unmarshal([]byte(`{"meterStop":0,"timestamp":"2024-01-01T00:00:00Z","transactionId":1,"transactionData":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[]}]}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v StopTransactionReq
		err := json.Unmarshal([]byte(`{"meterStop":100,"timestamp":"2024-01-01T00:00:00Z","transactionId":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.MeterStop != 100 {
			t.Errorf("got %d, want %d", v.MeterStop, 100)
		}

		if v.TransactionID != 1 {
			t.Errorf("got %d, want %d", v.TransactionID, 1)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v StopTransactionReq
		err := json.Unmarshal([]byte(`{"idTag":"ABC123","meterStop":100,"timestamp":"2024-01-01T00:00:00Z","transactionId":1,"reason":"Local","transactionData":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":"100"}]}]}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTag == nil || *v.IDTag != "ABC123" {
			t.Errorf("got %v, want %q", v.IDTag, "ABC123")
		}

		if v.Reason == nil || *v.Reason != ReasonLocal {
			t.Errorf("got %v, want %q", v.Reason, ReasonLocal)
		}

		if len(v.TransactionData) != 1 {
			t.Errorf("got %v, want a single transactionData entry", v.TransactionData)
		}
	})
}

func TestStopTransactionReqValidate(t *testing.T) {
	validTimestamp := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		input   StopTransactionReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   StopTransactionReq{MeterStop: 0, Timestamp: validTimestamp, TransactionID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects oversized idTag",
			input:   StopTransactionReq{IDTag: ptr(IDToken("ThisIsOver20Characters")), MeterStop: 0, Timestamp: validTimestamp, TransactionID: 1},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts zero meterStop",
			input:   StopTransactionReq{MeterStop: 0, Timestamp: validTimestamp, TransactionID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects negative meterStop",
			input:   StopTransactionReq{MeterStop: -1, Timestamp: validTimestamp, TransactionID: 1},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero timestamp",
			input:   StopTransactionReq{MeterStop: 0, TransactionID: 1},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects zero transactionId",
			input:   StopTransactionReq{MeterStop: 0, Timestamp: validTimestamp, TransactionID: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative transactionId",
			input:   StopTransactionReq{MeterStop: 0, Timestamp: validTimestamp, TransactionID: -1},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "accepts valid transactionData",
			input: StopTransactionReq{
				MeterStop:     0,
				Timestamp:     validTimestamp,
				TransactionID: 1,
				TransactionData: []MeterValue{
					{Timestamp: validTimestamp, SampledValue: []SampledValue{{Value: "100"}}},
				},
			},
			wantErr: nil,
		},
		{
			name: "rejects invalid transactionData entry",
			input: StopTransactionReq{
				MeterStop:       0,
				Timestamp:       validTimestamp,
				TransactionID:   1,
				TransactionData: []MeterValue{{Timestamp: validTimestamp}},
			},
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

func TestStopTransactionConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid nested idTagInfo", func(t *testing.T) {
		var v StopTransactionConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{}}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("valid input with no fields", func(t *testing.T) {
		var v StopTransactionConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTagInfo != nil {
			t.Errorf("got %v, want nil", v.IDTagInfo)
		}
	})

	t.Run("valid input with idTagInfo", func(t *testing.T) {
		var v StopTransactionConf
		err := json.Unmarshal([]byte(`{"idTagInfo":{"status":"Accepted"}}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.IDTagInfo == nil || v.IDTagInfo.Status != AuthorizationStatusAccepted {
			t.Errorf("got %v, want status %q", v.IDTagInfo, AuthorizationStatusAccepted)
		}
	})
}

func TestStopTransactionConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   StopTransactionConf
		wantErr error
	}{
		{
			name:    "accepts nil idTagInfo",
			input:   StopTransactionConf{},
			wantErr: nil,
		},
		{
			name:    "accepts valid idTagInfo",
			input:   StopTransactionConf{IDTagInfo: &IDTagInfo{Status: AuthorizationStatusAccepted}},
			wantErr: nil,
		},
		{
			name:    "rejects invalid idTagInfo",
			input:   StopTransactionConf{IDTagInfo: &IDTagInfo{Status: ""}},
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

func TestTriggerMessageReqUnmarshal(t *testing.T) {
	t.Run("rejects missing requestedMessage", func(t *testing.T) {
		var v TriggerMessageReq
		err := json.Unmarshal([]byte(`{}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects zero connectorId", func(t *testing.T) {
		var v TriggerMessageReq
		err := json.Unmarshal([]byte(`{"requestedMessage":"Heartbeat","connectorId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v TriggerMessageReq
		err := json.Unmarshal([]byte(`{"requestedMessage":"Heartbeat"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.RequestedMessage != MessageTriggerHeartbeat {
			t.Errorf("got %q, want %q", v.RequestedMessage, MessageTriggerHeartbeat)
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v TriggerMessageReq
		err := json.Unmarshal([]byte(`{"requestedMessage":"StatusNotification","connectorId":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID == nil || *v.ConnectorID != 1 {
			t.Errorf("got %v, want %d", v.ConnectorID, 1)
		}
	})
}

func TestTriggerMessageReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   TriggerMessageReq
		wantErr error
	}{
		{
			name:    "accepts valid input with only required fields",
			input:   TriggerMessageReq{RequestedMessage: MessageTriggerHeartbeat},
			wantErr: nil,
		},
		{
			name:    "accepts positive connectorId",
			input:   TriggerMessageReq{RequestedMessage: MessageTriggerStatusNotification, ConnectorID: ptr(int32(1))},
			wantErr: nil,
		},
		{
			name:    "rejects missing requestedMessage",
			input:   TriggerMessageReq{},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects zero connectorId",
			input:   TriggerMessageReq{RequestedMessage: MessageTriggerHeartbeat, ConnectorID: ptr(int32(0))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative connectorId",
			input:   TriggerMessageReq{RequestedMessage: MessageTriggerHeartbeat, ConnectorID: ptr(int32(-1))},
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

func TestTriggerMessageConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v TriggerMessageConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v TriggerMessageConf
		err := json.Unmarshal([]byte(`{"status":"Accepted"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != TriggerMessageStatusAccepted {
			t.Errorf("got %q, want %q", v.Status, TriggerMessageStatusAccepted)
		}
	})
}

func TestTriggerMessageConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   TriggerMessageConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   TriggerMessageConf{Status: TriggerMessageStatusAccepted},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   TriggerMessageConf{Status: ""},
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

func TestUnlockConnectorReqUnmarshal(t *testing.T) {
	t.Run("rejects zero connectorId", func(t *testing.T) {
		var v UnlockConnectorReq
		err := json.Unmarshal([]byte(`{"connectorId":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v UnlockConnectorReq
		err := json.Unmarshal([]byte(`{"connectorId":1}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.ConnectorID != 1 {
			t.Errorf("got %d, want %d", v.ConnectorID, 1)
		}
	})
}

func TestUnlockConnectorReqValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   UnlockConnectorReq
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   UnlockConnectorReq{ConnectorID: 1},
			wantErr: nil,
		},
		{
			name:    "rejects zero connectorId",
			input:   UnlockConnectorReq{ConnectorID: 0},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative connectorId",
			input:   UnlockConnectorReq{ConnectorID: -1},
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

func TestUnlockConnectorConfUnmarshal(t *testing.T) {
	t.Run("rejects invalid status", func(t *testing.T) {
		var v UnlockConnectorConf
		err := json.Unmarshal([]byte(`{"status":"BadStatus"}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var v UnlockConnectorConf
		err := json.Unmarshal([]byte(`{"status":"Unlocked"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Status != UnlockStatusUnlocked {
			t.Errorf("got %q, want %q", v.Status, UnlockStatusUnlocked)
		}
	})
}

func TestUnlockConnectorConfValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   UnlockConnectorConf
		wantErr error
	}{
		{
			name:    "accepts valid input",
			input:   UnlockConnectorConf{Status: UnlockStatusUnlocked},
			wantErr: nil,
		},
		{
			name:    "rejects empty status",
			input:   UnlockConnectorConf{Status: ""},
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

func TestUpdateFirmwareReqUnmarshal(t *testing.T) {
	t.Run("rejects missing location", func(t *testing.T) {
		var v UpdateFirmwareReq
		err := json.Unmarshal([]byte(`{"retrieveDate":"2024-01-01T00:00:00Z"}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects negative retries", func(t *testing.T) {
		var v UpdateFirmwareReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retrieveDate":"2024-01-01T00:00:00Z","retries":-1}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects missing retrieveDate", func(t *testing.T) {
		var v UpdateFirmwareReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/"}`), &v)

		if !errors.Is(err, ocpp.ErrOccurenceConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrOccurenceConstraintViolation)
		}
	})

	t.Run("rejects zero retryInterval", func(t *testing.T) {
		var v UpdateFirmwareReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retrieveDate":"2024-01-01T00:00:00Z","retryInterval":0}`), &v)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("got %v, want %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("valid input with only required fields", func(t *testing.T) {
		var v UpdateFirmwareReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retrieveDate":"2024-01-01T00:00:00Z"}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Location != "ftp://example.com/" {
			t.Errorf("got %q, want %q", v.Location, "ftp://example.com/")
		}

		if !v.RetrieveDate.Equal(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("got %v, want %v", v.RetrieveDate, "2024-01-01T00:00:00Z")
		}
	})

	t.Run("valid input with all fields", func(t *testing.T) {
		var v UpdateFirmwareReq
		err := json.Unmarshal([]byte(`{"location":"ftp://example.com/","retrieveDate":"2024-01-01T00:00:00Z","retries":3,"retryInterval":60}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if v.Retries == nil || *v.Retries != 3 {
			t.Errorf("got %v, want %d", v.Retries, 3)
		}

		if v.RetryInterval == nil || *v.RetryInterval != 60 {
			t.Errorf("got %v, want %d", v.RetryInterval, 60)
		}
	})
}

func TestUpdateFirmwareReqValidate(t *testing.T) {
	validRetrieveDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		input   UpdateFirmwareReq
		wantErr error
	}{
		{
			name:    "accepts valid input with only required fields",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/", RetrieveDate: validRetrieveDate},
			wantErr: nil,
		},
		{
			name: "accepts valid input with all fields",
			input: UpdateFirmwareReq{
				Location:      "ftp://example.com/",
				RetrieveDate:  validRetrieveDate,
				Retries:       ptr(int32(3)),
				RetryInterval: ptr(int32(60)),
			},
			wantErr: nil,
		},
		{
			name:    "rejects empty location",
			input:   UpdateFirmwareReq{Location: "", RetrieveDate: validRetrieveDate},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "accepts zero retries",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/", RetrieveDate: validRetrieveDate, Retries: ptr(int32(0))},
			wantErr: nil,
		},
		{
			name:    "rejects negative retries",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/", RetrieveDate: validRetrieveDate, Retries: ptr(int32(-1))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects zero retrieveDate",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/"},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects zero retryInterval",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/", RetrieveDate: validRetrieveDate, RetryInterval: ptr(int32(0))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative retryInterval",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/", RetrieveDate: validRetrieveDate, RetryInterval: ptr(int32(-1))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "accepts positive retryInterval",
			input:   UpdateFirmwareReq{Location: "ftp://example.com/", RetrieveDate: validRetrieveDate, RetryInterval: ptr(int32(60))},
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

func TestUpdateFirmwareConfUnmarshal(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		var v UpdateFirmwareConf
		err := json.Unmarshal([]byte(`{}`), &v)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestUpdateFirmwareConfValidate(t *testing.T) {
	t.Run("accepts valid input", func(t *testing.T) {
		err := UpdateFirmwareConf{}.Validate()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
