package v16

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

func TestAuthorizationDataUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts all fields",
			input: []byte(`{"idTag":"ABC123","idTagInfo":{"status":"Accepted"}}`),
		},
		{
			name:  "accepts without optional idTagInfo",
			input: []byte(`{"idTag":"ABC123"}`),
		},
		{
			name:    "rejects missing idTag",
			input:   []byte(`{"idTagInfo":{"status":"Accepted"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty idTag",
			input:   []byte(`{"idTag":"","idTagInfo":{"status":"Accepted"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects idTag exceeding max length",
			input:   []byte(`{"idTag":"AAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects invalid idTagInfo",
			input:   []byte(`{"idTag":"ABC123","idTagInfo":{}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a AuthorizationData

			err := json.Unmarshal(tt.input, &a)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizationDataValidate(t *testing.T) {
	parentIDTag := IDToken("PARENT01")

	tests := []struct {
		name    string
		input   AuthorizationData
		wantErr error
	}{
		{
			name: "accepts all fields",
			input: AuthorizationData{
				IDTag:     IDToken("ABC123"),
				IDTagInfo: &IDTagInfo{Status: AuthorizationStatusAccepted, ParentIDTag: &parentIDTag},
			},
		},
		{
			name:  "accepts without optional idTagInfo",
			input: AuthorizationData{IDTag: IDToken("ABC123")},
		},
		{
			name:    "rejects missing idTag",
			input:   AuthorizationData{},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects idTag with invalid characters",
			input:   AuthorizationData{IDTag: IDToken("foo\x7fbar")},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects invalid idTagInfo",
			input: AuthorizationData{
				IDTag:     IDToken("ABC123"),
				IDTagInfo: &IDTagInfo{},
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizationStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str AuthorizationStatus
		data := fmt.Appendf(nil, `"%s"`, AuthorizationStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != AuthorizationStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, AuthorizationStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str AuthorizationStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str AuthorizationStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestAvailabilityStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str AvailabilityStatus
		data := fmt.Appendf(nil, `"%s"`, AvailabilityStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != AvailabilityStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, AvailabilityStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str AvailabilityStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str AvailabilityStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestAvailabilityTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str AvailabilityType
		data := fmt.Appendf(nil, `"%s"`, AvailabilityTypeOperative)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != AvailabilityTypeOperative {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, AvailabilityTypeOperative)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str AvailabilityType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str AvailabilityType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestCancelReservationStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CancelReservationStatus
		data := fmt.Appendf(nil, `"%s"`, CancelReservationStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CancelReservationStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CancelReservationStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CancelReservationStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CancelReservationStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargePointErrorCodeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargePointErrorCode
		data := fmt.Appendf(nil, `"%s"`, ChargePointErrorCodeGroundFailure)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargePointErrorCodeGroundFailure {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargePointErrorCodeGroundFailure)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargePointErrorCode
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargePointErrorCode
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargePointStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargePointStatus
		data := fmt.Appendf(nil, `"%s"`, ChargePointStatusFinishing)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargePointStatusFinishing {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargePointStatusFinishing)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargePointStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargePointStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargingProfileUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:  "accepts required fields only",
			input: `{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}`,
		},
		{
			name:  "accepts with all optional fields",
			input: `{"chargingProfileId":1,"transactionId":42,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Recurring","recurrencyKind":"Daily","validFrom":"2026-01-01T00:00:00Z","validTo":"2026-12-31T23:59:59Z","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}`,
		},
		{
			name:    "rejects negative stackLevel",
			input:   `{"chargingProfileId":1,"stackLevel":-1,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}`,
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing chargingProfilePurpose",
			input:   `{"chargingProfileId":1,"stackLevel":0,"chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingProfileKind",
			input:   `{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId when purpose is not TxProfile",
			input:   `{"chargingProfileId":1,"transactionId":42,"stackLevel":0,"chargingProfilePurpose":"ChargePointMaxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32}]}}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects invalid chargingSchedule",
			input:   `{"chargingProfileId":1,"stackLevel":0,"chargingProfilePurpose":"TxProfile","chargingProfileKind":"Absolute","chargingSchedule":{"chargingRateUnit":"W","chargingSchedulePeriod":[]}}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingProfile
			err := json.Unmarshal([]byte(tt.input), &s)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestChargingProfileValidate(t *testing.T) {
	txID := int32(42)
	rk := RecurrencyKindTypeDaily
	rkEmpty := RecurrencyKindType("")
	validSchedule := ChargingSchedule{
		ChargingRateUnit:       ChargingRateUnitTypeW,
		ChargingSchedulePeriod: []ChargingSchedulePeriod{{StartPeriod: 0, Limit: 32}},
	}

	tests := []struct {
		name    string
		input   ChargingProfile
		wantErr error
	}{
		{
			name: "accepts required fields only",
			input: ChargingProfile{
				ChargingProfileID:      1,
				StackLevel:             0,
				ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
				ChargingProfileKind:    ChargingProfileKindTypeAbsolute,
				ChargingSchedule:       validSchedule,
			},
		},
		{
			name: "accepts with all optional fields",
			input: ChargingProfile{
				ChargingProfileID:      1,
				TransactionID:          &txID,
				StackLevel:             0,
				ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
				ChargingProfileKind:    ChargingProfileKindTypeRecurring,
				RecurrencyKind:         &rk,
				ChargingSchedule:       validSchedule,
			},
		},
		{
			name: "rejects negative stackLevel",
			input: ChargingProfile{
				ChargingProfileID:      1,
				StackLevel:             -1,
				ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
				ChargingProfileKind:    ChargingProfileKindTypeAbsolute,
				ChargingSchedule:       validSchedule,
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects missing chargingProfilePurpose",
			input: ChargingProfile{
				ChargingProfileID:   1,
				StackLevel:          0,
				ChargingProfileKind: ChargingProfileKindTypeAbsolute,
				ChargingSchedule:    validSchedule,
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects missing chargingProfileKind",
			input: ChargingProfile{
				ChargingProfileID:      1,
				StackLevel:             0,
				ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
				ChargingSchedule:       validSchedule,
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects empty recurrencyKind pointer",
			input: ChargingProfile{
				ChargingProfileID:      1,
				StackLevel:             0,
				ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
				ChargingProfileKind:    ChargingProfileKindTypeRecurring,
				RecurrencyKind:         &rkEmpty,
				ChargingSchedule:       validSchedule,
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects transactionId when purpose is not TxProfile",
			input: ChargingProfile{
				ChargingProfileID:      1,
				TransactionID:          &txID,
				StackLevel:             0,
				ChargingProfilePurpose: ChargingProfilePurposeTypeChargePointMaxProfile,
				ChargingProfileKind:    ChargingProfileKindTypeAbsolute,
				ChargingSchedule:       validSchedule,
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects invalid chargingSchedule",
			input: ChargingProfile{
				ChargingProfileID:      1,
				StackLevel:             0,
				ChargingProfilePurpose: ChargingProfilePurposeTypeTxProfile,
				ChargingProfileKind:    ChargingProfileKindTypeAbsolute,
				ChargingSchedule:       ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW},
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestChargingProfileKindTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingProfileKindType
		data := fmt.Appendf(nil, `"%s"`, ChargingProfileKindTypeRelative)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingProfileKindTypeRelative {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingProfileKindTypeRelative)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingProfileKindType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingProfileKindType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargingProfilePurposeTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingProfilePurposeType
		data := fmt.Appendf(nil, `"%s"`, ChargingProfilePurposeTypeTxProfile)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingProfilePurposeTypeTxProfile {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingProfilePurposeTypeTxProfile)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingProfilePurposeType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingProfilePurposeType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargingProfileStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingProfileStatus
		data := fmt.Appendf(nil, `"%s"`, ChargingProfileStatusNotSupported)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingProfileStatusNotSupported {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingProfileStatusNotSupported)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingProfileStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingProfileStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargingRateUnitTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingRateUnitType
		data := fmt.Appendf(nil, `"%s"`, ChargingRateUnitTypeW)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingRateUnitTypeW {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingRateUnitTypeW)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingRateUnitType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingRateUnitType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestChargingScheduleUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:  "accepts required fields only",
			input: `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":8.1}]}`,
		},
		{
			name:  "accepts with duration",
			input: `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":8.1}],"duration":3600}`,
		},
		{
			name:  "accepts with minChargingRate",
			input: `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":8.1}],"minChargingRate":6.0}`,
		},
		{
			name:  "accepts multiple periods",
			input: `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":32},{"startPeriod":3600,"limit":8.1}]}`,
		},
		{
			name:    "rejects missing chargingRateUnit",
			input:   `{"chargingSchedulePeriod":[{"startPeriod":0,"limit":8.1}]}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty chargingSchedulePeriod",
			input:   `{"chargingRateUnit":"W","chargingSchedulePeriod":[]}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects first startPeriod not zero",
			input:   `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":60,"limit":8.1}]}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects duration of zero",
			input:   `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":8.1}],"duration":0}`,
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects minChargingRate with two digit fraction",
			input:   `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":8.1}],"minChargingRate":6.12}`,
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects invalid period in list",
			input:   `{"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":0,"limit":-1}]}`,
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingSchedule
			err := json.Unmarshal([]byte(tt.input), &s)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestChargingScheduleValidate(t *testing.T) {
	validPeriod := ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.1}
	duration := int32(3600)
	zeroDuration := int32(0)
	minRate := float64(6.0)
	negMinRate := float64(-1.0)
	badMinRate := float64(6.12)

	tests := []struct {
		name    string
		input   ChargingSchedule
		wantErr error
	}{
		{
			name:  "accepts required fields only",
			input: ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}},
		},
		{
			name:  "accepts with duration",
			input: ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}, Duration: &duration},
		},
		{
			name:  "accepts with minChargingRate",
			input: ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}, MinChargingRate: &minRate},
		},
		{
			name: "accepts multiple periods",
			input: ChargingSchedule{
				ChargingRateUnit: ChargingRateUnitTypeW,
				ChargingSchedulePeriod: []ChargingSchedulePeriod{
					{StartPeriod: 0, Limit: 32},
					{StartPeriod: 3600, Limit: 8.1},
				},
			},
		},
		{
			name:    "rejects missing chargingRateUnit",
			input:   ChargingSchedule{ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects empty chargingSchedulePeriod",
			input:   ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{}},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects first startPeriod not zero",
			input:   ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{{StartPeriod: 60, Limit: 8.1}}},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects duration less than 1",
			input:   ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}, Duration: &zeroDuration},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects negative minChargingRate",
			input:   ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}, MinChargingRate: &negMinRate},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects minChargingRate with two digit fraction",
			input:   ChargingSchedule{ChargingRateUnit: ChargingRateUnitTypeW, ChargingSchedulePeriod: []ChargingSchedulePeriod{validPeriod}, MinChargingRate: &badMinRate},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name: "rejects invalid period in list",
			input: ChargingSchedule{
				ChargingRateUnit:       ChargingRateUnitTypeW,
				ChargingSchedulePeriod: []ChargingSchedulePeriod{{StartPeriod: 0, Limit: -1}},
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestChargingSchedulePeriodUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{name: "accepts valid period", input: `{"startPeriod":0,"limit":8.1}`},
		{name: "accepts zero limit", input: `{"startPeriod":0,"limit":0}`},
		{name: "accepts integer limit", input: `{"startPeriod":3600,"limit":32}`},
		{name: "accepts numberPhases 1", input: `{"startPeriod":0,"limit":8.1,"numberPhases":1}`},
		{name: "accepts numberPhases 3", input: `{"startPeriod":0,"limit":8.1,"numberPhases":3}`},
		{name: "rejects negative startPeriod", input: `{"startPeriod":-1,"limit":8.1}`, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects negative limit", input: `{"startPeriod":0,"limit":-1.0}`, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects limit with two digit fraction", input: `{"startPeriod":0,"limit":8.12}`, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects numberPhases 0", input: `{"startPeriod":0,"limit":8.1,"numberPhases":0}`, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects numberPhases 4", input: `{"startPeriod":0,"limit":8.1,"numberPhases":4}`, wantErr: ocpp.ErrPropertyConstraintViolation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingSchedulePeriod
			err := json.Unmarshal([]byte(tt.input), &s)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestChargingSchedulePeriodValidate(t *testing.T) {
	np1 := int32(1)
	np3 := int32(3)
	np0 := int32(0)
	np4 := int32(4)

	tests := []struct {
		name    string
		input   ChargingSchedulePeriod
		wantErr error
	}{
		{name: "accepts valid period", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.1}},
		{name: "accepts zero limit", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 0}},
		{name: "accepts integer limit", input: ChargingSchedulePeriod{StartPeriod: 3600, Limit: 32}},
		{name: "accepts numberPhases 1", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.1, NumberPhases: &np1}},
		{name: "accepts numberPhases 3", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.1, NumberPhases: &np3}},
		{name: "rejects negative startPeriod", input: ChargingSchedulePeriod{StartPeriod: -1, Limit: 8.1}, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects negative limit", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: -1.0}, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects limit with two digit fraction", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.12}, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects numberPhases 0", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.1, NumberPhases: &np0}, wantErr: ocpp.ErrPropertyConstraintViolation},
		{name: "rejects numberPhases 4", input: ChargingSchedulePeriod{StartPeriod: 0, Limit: 8.1, NumberPhases: &np4}, wantErr: ocpp.ErrPropertyConstraintViolation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestNewCiString20Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString20Type
		equals     string
		err        error
		errStr     string
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"foooooooooooooooooooo"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString20Type exceeds max length of 20",
		},
		{
			name:       "should accept space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString20Type("\x20"),
			err:        nil,
		},
		{
			name:       "should accept lower boundary printable ASCII (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString20Type("\x20"),
			err:        nil,
		},
		{
			name:       "should error if containing non-printable ASCII (0x1F)",
			inputBytes: []byte(`"\u001F"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString20Type contains non-printable ASCII characters",
		},
		{
			name:       "should accept upper boundary visible ASCII (0x7E)",
			inputBytes: []byte(`"` + "\x7E" + `"`),
			want:       CiString20Type("\x7E"),
			err:        nil,
		},
		{
			name:       "should error if containing DEL character (0x7F)",
			inputBytes: []byte(`"` + "\x7F" + `"`),
			err:        ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foo"`),
			want:       CiString20Type("foo"),
			equals:     "FoO",
			err:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var str CiString20Type
			err := json.Unmarshal(tt.inputBytes, &str)

			if !errors.Is(err, tt.err) {
				t.Errorf("unexpected unmarshaling error\ngot:  %v\nwant: %v", err, tt.err)
			}

			if tt.errStr != "" && err != nil && err.Error() != tt.errStr {
				t.Errorf("unexpected error string\ngot:  %v\nwant: %v", err.Error(), tt.errStr)
			}

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}

			if tt.equals != "" && !str.Equals(tt.equals) {
				t.Errorf("expected %s to CI equal %s", str, tt.equals)
			}
		})
	}
}

func TestNewCiString25Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString25Type
		equals     string
		err        error
		errStr     string
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"fooooooooooooooooooooooooo"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString25Type exceeds max length of 25",
		},
		{
			name:       "should accept space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString25Type("\x20"),
			err:        nil,
		},
		{
			name:       "should error if containing non-printable ASCII (0x1F)",
			inputBytes: []byte(`"\u001F"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString25Type contains non-printable ASCII characters",
		},
		{
			name:       "should accept lower boundary printable ASCII (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString25Type("\x20"),
			err:        nil,
		},
		{
			name:       "should accept upper boundary visible ASCII (0x7E)",
			inputBytes: []byte(`"` + "\x7E" + `"`),
			want:       CiString25Type("\x7E"),
			err:        nil,
		},
		{
			name:       "should error if containing DEL character (0x7F)",
			inputBytes: []byte(`"` + "\x7F" + `"`),
			err:        ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foooooooooooooooooooooooo"`),
			want:       CiString25Type("foooooooooooooooooooooooo"),
			equals:     "foooooooooooOoooooooooooo",
			err:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var str CiString25Type
			err := json.Unmarshal(tt.inputBytes, &str)

			if !errors.Is(err, tt.err) {
				t.Errorf("unexpected unmarshaling error\ngot:  %v\nwant: %v", err, tt.err)
			}

			if tt.errStr != "" && err != nil && err.Error() != tt.errStr {
				t.Errorf("unexpected error string\ngot:  %v\nwant: %v", err.Error(), tt.errStr)
			}

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}

			if tt.equals != "" && !str.Equals(tt.equals) {
				t.Errorf("expected %s to CI equal %s", str, tt.equals)
			}
		})
	}
}

func TestNewCiString50Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString50Type
		equals     string
		err        error
		errStr     string
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"foooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString50Type exceeds max length of 50",
		},
		{
			name:       "should accept space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString50Type("\x20"),
			err:        nil,
		},
		{
			name:       "should error if containing non-printable ASCII (0x1F)",
			inputBytes: []byte(`"\u001F"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString50Type contains non-printable ASCII characters",
		},
		{
			name:       "should accept lower boundary printable ASCII (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString50Type("\x20"),
			err:        nil,
		},
		{
			name:       "should accept upper boundary visible ASCII (0x7E)",
			inputBytes: []byte(`"` + "\x7E" + `"`),
			want:       CiString50Type("\x7E"),
			err:        nil,
		},
		{
			name:       "should error if containing DEL character (0x7F)",
			inputBytes: []byte(`"` + "\x7F" + `"`),
			err:        ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foooooooooooooooooooooooo"`),
			want:       CiString50Type("foooooooooooooooooooooooo"),
			equals:     "foooooooooooOoooooooooooo",
			err:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var str CiString50Type
			err := json.Unmarshal(tt.inputBytes, &str)

			if !errors.Is(err, tt.err) {
				t.Errorf("unexpected unmarshaling error\ngot:  %v\nwant: %v", err, tt.err)
			}

			if tt.errStr != "" && err != nil && err.Error() != tt.errStr {
				t.Errorf("unexpected error string\ngot:  %v\nwant: %v", err.Error(), tt.errStr)
			}

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}

			if tt.equals != "" && !str.Equals(tt.equals) {
				t.Errorf("expected %s to CI equal %s", str, tt.equals)
			}
		})
	}
}

func TestNewCiString255Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString255Type
		equals     string
		err        error
		errStr     string
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"fooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString255Type exceeds max length of 255",
		},
		{
			name:       "should accept space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString255Type("\x20"),
			err:        nil,
		},
		{
			name:       "should error if containing non-printable ASCII (0x1F)",
			inputBytes: []byte(`"\u001F"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString255Type contains non-printable ASCII characters",
		},
		{
			name:       "should accept lower boundary printable ASCII (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString255Type("\x20"),
			err:        nil,
		},
		{
			name:       "should accept upper boundary visible ASCII (0x7E)",
			inputBytes: []byte(`"` + "\x7E" + `"`),
			want:       CiString255Type("\x7E"),
			err:        nil,
		},
		{
			name:       "should error if containing DEL character (0x7F)",
			inputBytes: []byte(`"` + "\x7F" + `"`),
			err:        ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foooooooooooooooooooooooo"`),
			want:       CiString255Type("foooooooooooooooooooooooo"),
			equals:     "foooooooooooOoooooooooooo",
			err:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var str CiString255Type
			err := json.Unmarshal(tt.inputBytes, &str)

			if !errors.Is(err, tt.err) {
				t.Errorf("unexpected unmarshaling error\ngot:  %v\nwant: %v", err, tt.err)
			}

			if tt.errStr != "" && err != nil && err.Error() != tt.errStr {
				t.Errorf("unexpected error string\ngot:  %v\nwant: %v", err.Error(), tt.errStr)
			}

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}

			if tt.equals != "" && !str.Equals(tt.equals) {
				t.Errorf("expected %s to CI equal %s", str, tt.equals)
			}
		})
	}
}

func TestNewCiString500Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString500Type
		equals     string
		err        error
		errStr     string
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"foooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString500Type exceeds max length of 500",
		},
		{
			name:       "should accept space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString500Type("\x20"),
			err:        nil,
		},
		{
			name:       "should error if containing non-printable ASCII (0x1F)",
			inputBytes: []byte(`"\u001F"`),
			err:        ocpp.ErrPropertyConstraintViolation,
			errStr:     "[PropertyConstraintViolation] CiString500Type contains non-printable ASCII characters",
		},
		{
			name:       "should accept lower boundary printable ASCII (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			want:       CiString500Type("\x20"),
			err:        nil,
		},
		{
			name:       "should accept upper boundary visible ASCII (0x7E)",
			inputBytes: []byte(`"` + "\x7E" + `"`),
			want:       CiString500Type("\x7E"),
			err:        nil,
		},
		{
			name:       "should error if containing DEL character (0x7F)",
			inputBytes: []byte(`"` + "\x7F" + `"`),
			err:        ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foooooooooooooooooooooooo"`),
			want:       CiString500Type("foooooooooooooooooooooooo"),
			equals:     "foooooooooooOoooooooooooo",
			err:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var str CiString500Type
			err := json.Unmarshal(tt.inputBytes, &str)

			if !errors.Is(err, tt.err) {
				t.Errorf("unexpected unmarshaling error\ngot:  %v\nwant: %v", err, tt.err)
			}

			if tt.errStr != "" && err != nil && err.Error() != tt.errStr {
				t.Errorf("unexpected error string\ngot:  %v\nwant: %v", err.Error(), tt.errStr)
			}

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}

			if tt.equals != "" && !str.Equals(tt.equals) {
				t.Errorf("expected %s to CI equal %s", str, tt.equals)
			}
		})
	}
}

func TestClearCacheStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ClearCacheStatus
		data := fmt.Appendf(nil, `"%s"`, ClearCacheStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ClearCacheStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ClearCacheStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ClearCacheStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ClearCacheStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestClearChargingProfileStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ClearChargingProfileStatus
		data := fmt.Appendf(nil, `"%s"`, ClearChargingProfileStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ClearChargingProfileStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ClearChargingProfileStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ClearChargingProfileStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ClearChargingProfileStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestConfigurationStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ConfigurationStatus
		data := fmt.Appendf(nil, `"%s"`, ConfigurationStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ConfigurationStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ConfigurationStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ConfigurationStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ConfigurationStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestDataTransferStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DataTransferStatus
		data := fmt.Appendf(nil, `"%s"`, DataTransferStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DataTransferStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DataTransferStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DataTransferStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DataTransferStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestDiagnosticsStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DiagnosticsStatus
		data := fmt.Appendf(nil, `"%s"`, DiagnosticsStatusUploaded)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DiagnosticsStatusUploaded {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DiagnosticsStatusUploaded)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DiagnosticsStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DiagnosticsStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestFirmwareStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str FirmwareStatus
		data := fmt.Appendf(nil, `"%s"`, FirmwareStatusDownloading)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != FirmwareStatusDownloading {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, FirmwareStatusDownloading)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str FirmwareStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str FirmwareStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestGetCompositeScheduleStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetCompositeScheduleStatus
		data := fmt.Appendf(nil, `"%s"`, GetCompositeScheduleStatusRejected)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetCompositeScheduleStatusRejected {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetCompositeScheduleStatusRejected)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetCompositeScheduleStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetCompositeScheduleStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestIDTagInfoUnmarshalJSON(t *testing.T) {
	t.Run("rejects invalid expiryDate", func(t *testing.T) {
		got := []byte(`{
			"status": "Accepted",
			"expiryDate": "",
			"parentIdTag": ""
		}`)
		var i IDTagInfo
		var parseErr *time.ParseError

		err := json.Unmarshal(got, &i)
		if !errors.As(err, &parseErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: %v", err, parseErr)
		}
	})

	t.Run("accepts with only required fields", func(t *testing.T) {
		got := []byte(`{
			"status": "Accepted"
		}`)
		var i IDTagInfo

		if err := json.Unmarshal(got, &i); err != nil {
			t.Errorf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("accepts all fields", func(t *testing.T) {
		got := []byte(`{
			"status": "Accepted",
			"expiryDate": "2026-06-30T14:30:00Z",
			"parentIdTag": "PARENT01"
		}`)
		var i IDTagInfo

		if err := json.Unmarshal(got, &i); err != nil {
			t.Errorf("unexpected unmarshal error: %v", err)
		}
	})
}

func TestIDTagInfoValidate(t *testing.T) {
	now := time.Now()
	parentIDTag := IDToken("PARENT01")
	emptyParentIDTag := IDToken("")
	invalidParentIDTag := IDToken("foo\x7fbar")

	tests := []struct {
		name    string
		input   IDTagInfo
		wantErr error
	}{
		{
			name: "accepts all fields",
			input: IDTagInfo{
				Status:      AuthorizationStatusAccepted,
				ExpiryDate:  &now,
				ParentIDTag: &parentIDTag,
			},
		},
		{
			name: "accepts with only required fields",
			input: IDTagInfo{
				Status: AuthorizationStatusAccepted,
			},
		},
		{
			name:    "rejects missing status",
			input:   IDTagInfo{},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects empty parentIdTag",
			input: IDTagInfo{
				Status:      AuthorizationStatusAccepted,
				ParentIDTag: &emptyParentIDTag,
			},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name: "rejects parentIdTag with invalid characters",
			input: IDTagInfo{
				Status:      AuthorizationStatusAccepted,
				ParentIDTag: &invalidParentIDTag,
			},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyValueUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:  "accepts required fields only",
			input: `{"key":"SomeKey","readonly":false}`,
		},
		{
			name:  "accepts with optional value",
			input: `{"key":"SomeKey","readonly":true,"value":"SomeValue"}`,
		},
		{
			name:    "rejects missing key",
			input:   `{"readonly":false}`,
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects key with non-printable ASCII",
			input:   `{"key":"foobar","readonly":false}`,
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects value with non-printable ASCII",
			input:   `{"key":"SomeKey","readonly":false,"value":"foobar"}`,
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s KeyValue
			err := json.Unmarshal([]byte(tt.input), &s)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestKeyValueValidate(t *testing.T) {
	value := CiString500Type("SomeValue")

	tests := []struct {
		name    string
		input   KeyValue
		wantErr error
	}{
		{
			name:  "accepts required fields only",
			input: KeyValue{Key: "SomeKey", Readonly: false},
		},
		{
			name:  "accepts with optional value",
			input: KeyValue{Key: "SomeKey", Readonly: true, Value: &value},
		},
		{
			name:    "rejects missing key",
			input:   KeyValue{},
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects key with non-printable ASCII",
			input:   KeyValue{Key: "foo\x7fbar"},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects value with non-printable ASCII",
			input:   KeyValue{Key: "SomeKey", Value: func() *CiString500Type { v := CiString500Type("foo\x7fbar"); return &v }()},
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error\ngot:  %v\nwant: nil", err)
			}
		})
	}
}

func TestLocationUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str Location
		data := fmt.Appendf(nil, `"%s"`, LocationCable)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != LocationCable {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, LocationCable)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str Location
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str Location
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestMeasurandUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str Measurand
		data := fmt.Appendf(nil, `"%s"`, MeasurandEnergyActiveExportRegister)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MeasurandEnergyActiveExportRegister {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MeasurandEnergyActiveExportRegister)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str Measurand
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str Measurand
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestMessageTriggerUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MessageTrigger
		data := fmt.Appendf(nil, `"%s"`, MessageTriggerBootNotification)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MessageTriggerBootNotification {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MessageTriggerBootNotification)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MessageTrigger
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MessageTrigger
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestPhaseUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str Phase
		data := fmt.Appendf(nil, `"%s"`, PhaseL1N)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PhaseL1N {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PhaseL1N)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str Phase
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str Phase
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestReadingContextUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReadingContext
		data := fmt.Appendf(nil, `"%s"`, ReadingContextOther)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReadingContextOther {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReadingContextOther)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReadingContext
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReadingContext
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestReasonUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str Reason
		data := fmt.Appendf(nil, `"%s"`, ReasonOther)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReasonOther {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReasonOther)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str Reason
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str Reason
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestRecurrencyKindTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str RecurrencyKindType
		data := fmt.Appendf(nil, `"%s"`, RecurrencyKindTypeWeekly)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != RecurrencyKindTypeWeekly {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, RecurrencyKindTypeWeekly)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str RecurrencyKindType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str RecurrencyKindType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestRegistrationStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str RegistrationStatus
		data := fmt.Appendf(nil, `"%s"`, RegistrationStatusPending)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != RegistrationStatusPending {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, RegistrationStatusPending)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str RegistrationStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str RegistrationStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestRemoteStartStopStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str RemoteStartStopStatus
		data := fmt.Appendf(nil, `"%s"`, RemoteStartStopStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != RemoteStartStopStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, RemoteStartStopStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str RemoteStartStopStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str RemoteStartStopStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestReservationStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReservationStatus
		data := fmt.Appendf(nil, `"%s"`, ReservationStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReservationStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReservationStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReservationStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReservationStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestResetStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ResetStatus
		data := fmt.Appendf(nil, `"%s"`, ResetStatusAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ResetStatusAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ResetStatusAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ResetStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ResetStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestResetTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ResetType
		data := fmt.Appendf(nil, `"%s"`, ResetTypeHard)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ResetTypeHard {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ResetTypeHard)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ResetType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ResetType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestTriggerMessageStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TriggerMessageStatus
		data := fmt.Appendf(nil, `"%s"`, TriggerMessageStatusNotImplemented)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TriggerMessageStatusNotImplemented {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TriggerMessageStatusNotImplemented)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TriggerMessageStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TriggerMessageStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestUnitOfMeasureUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UnitOfMeasure
		data := fmt.Appendf(nil, `"%s"`, UnitOfMeasureKWh)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UnitOfMeasureKWh {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UnitOfMeasureKWh)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UnitOfMeasure
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UnitOfMeasure
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestUnlockStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UnlockStatus
		data := fmt.Appendf(nil, `"%s"`, UnlockStatusNotSupported)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UnlockStatusNotSupported {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UnlockStatusNotSupported)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UnlockStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UnlockStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestUpdateStatusUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UpdateStatus
		data := fmt.Appendf(nil, `"%s"`, UpdateStatusFailed)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UpdateStatusFailed {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UpdateStatusFailed)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UpdateStatus
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UpdateStatus
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestUpdateTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UpdateType
		data := fmt.Appendf(nil, `"%s"`, UpdateTypeDifferential)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UpdateTypeDifferential {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UpdateTypeDifferential)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UpdateType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UpdateType
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}

func TestValueFormatUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ValueFormat
		data := fmt.Appendf(nil, `"%s"`, ValueFormatSignedData)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ValueFormatSignedData {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ValueFormatSignedData)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ValueFormat
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ValueFormat
		data := []byte(`1`)

		err := json.Unmarshal(data, &str)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}

		if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: not an ocpp.ErrPropertyConstraintViolation", err)
		}
	})
}
