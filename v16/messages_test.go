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
			name:    "rejects zero connectorId",
			input:   GetCompositeScheduleConf{Status: GetCompositeScheduleStatusAccepted, ConnectorID: ptr(int32(0))},
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
