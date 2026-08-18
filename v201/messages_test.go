package v201

import (
	"encoding/json"
	"errors"
	"testing"

	ocpp "github.com/feightree/gocpp/ocpp"
)

func TestAuthorizeRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificate":"sample","idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
		},
		{
			name:    "rejects certificate exceeding max length",
			input:   []byte(`{"certificate":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"certificate":"sample","iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects iso15118CertificateHashData exceeding max items",
			input:   []byte(`{"certificate":"sample","idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s AuthorizeRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AuthorizeRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAuthorizeResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateStatus":"Accepted","idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}}}`),
		},
		{
			name:    "rejects missing idTokenInfo",
			input:   []byte(`{"certificateStatus":"Accepted"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s AuthorizeResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AuthorizeResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestBootNotificationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"reason":"ApplicationReset","chargingStation":{"serialNumber":"sample","model":"sample","vendorName":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}}`),
		},
		{
			name:    "rejects missing reason",
			input:   []byte(`{"chargingStation":{"serialNumber":"sample","model":"sample","vendorName":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingStation",
			input:   []byte(`{"reason":"ApplicationReset"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s BootNotificationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s BootNotificationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestBootNotificationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"currentTime":"2024-01-01T00:00:00Z","interval":1,"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing currentTime",
			input:   []byte(`{"interval":1,"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"currentTime":"2024-01-01T00:00:00Z","interval":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s BootNotificationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s BootNotificationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCancelReservationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"reservationId":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CancelReservationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CancelReservationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCancelReservationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CancelReservationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CancelReservationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCertificateSignedRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateChain":"sample","certificateType":"ChargingStationCertificate"}`),
		},
		{
			name:    "rejects missing certificateChain",
			input:   []byte(`{"certificateType":"ChargingStationCertificate"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects certificateChain exceeding max length",
			input:   []byte(`{"certificateChain":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","certificateType":"ChargingStationCertificate"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CertificateSignedRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CertificateSignedRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCertificateSignedResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CertificateSignedResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CertificateSignedResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChangeAvailabilityRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"operationalStatus":"Inoperative","evse":{"id":1,"connectorId":1}}`),
		},
		{
			name:    "rejects missing operationalStatus",
			input:   []byte(`{"evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChangeAvailabilityRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChangeAvailabilityRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChangeAvailabilityResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChangeAvailabilityResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChangeAvailabilityResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearCacheRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearCacheRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearCacheRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearCacheResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearCacheResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearCacheResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearChargingProfileRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"chargingProfileId":1,"chargingProfileCriteria":{"evseId":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":1}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearChargingProfileRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearChargingProfileRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearChargingProfileResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearChargingProfileResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearChargingProfileResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearDisplayMessageRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearDisplayMessageRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearDisplayMessageRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearDisplayMessageResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearDisplayMessageResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearDisplayMessageResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearedChargingLimitRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"chargingLimitSource":"EMS","evseId":1}`),
		},
		{
			name:    "rejects missing chargingLimitSource",
			input:   []byte(`{"evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearedChargingLimitRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearedChargingLimitRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearedChargingLimitResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearedChargingLimitResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearedChargingLimitResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearVariableMonitoringRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":[1]}`),
		},
		{
			name:    "rejects missing id",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearVariableMonitoringRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearVariableMonitoringRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearVariableMonitoringResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"clearMonitoringResult":[{"status":"Accepted","id":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
		},
		{
			name:    "rejects missing clearMonitoringResult",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearVariableMonitoringResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearVariableMonitoringResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCostUpdatedRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"totalCost":1.0,"transactionId":"A1"}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{"totalCost":1.0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"totalCost":1.0,"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"totalCost":1.0,"transactionId":"####################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CostUpdatedRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CostUpdatedRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCostUpdatedResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CostUpdatedResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CostUpdatedResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCustomerInformationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"report":true,"clear":true,"customerIdentifier":"sample","idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"customerCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
		},
		{
			name:    "rejects customerIdentifier exceeding max length",
			input:   []byte(`{"requestId":1,"report":true,"clear":true,"customerIdentifier":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"customerCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CustomerInformationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CustomerInformationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCustomerInformationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CustomerInformationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CustomerInformationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestDataTransferRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"messageId":"sample","data":"sample","vendorId":"sample"}`),
		},
		{
			name:    "rejects messageId exceeding max length",
			input:   []byte(`{"messageId":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","data":"sample","vendorId":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing vendorId",
			input:   []byte(`{"messageId":"sample","data":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects vendorId exceeding max length",
			input:   []byte(`{"messageId":"sample","data":"sample","vendorId":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s DataTransferRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s DataTransferRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestDataTransferResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","data":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"data":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s DataTransferResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s DataTransferResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestDeleteCertificateRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
		},
		{
			name:    "rejects missing certificateHashData",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s DeleteCertificateRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s DeleteCertificateRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestDeleteCertificateResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s DeleteCertificateResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s DeleteCertificateResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestFirmwareStatusNotificationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Downloaded","requestId":1}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s FirmwareStatusNotificationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s FirmwareStatusNotificationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestFirmwareStatusNotificationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s FirmwareStatusNotificationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s FirmwareStatusNotificationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGet15118EVCertificateRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"iso15118SchemaVersion":"sample","action":"Install","exiRequest":"sample"}`),
		},
		{
			name:    "rejects missing iso15118SchemaVersion",
			input:   []byte(`{"action":"Install","exiRequest":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects iso15118SchemaVersion exceeding max length",
			input:   []byte(`{"iso15118SchemaVersion":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","action":"Install","exiRequest":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing action",
			input:   []byte(`{"iso15118SchemaVersion":"sample","exiRequest":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing exiRequest",
			input:   []byte(`{"iso15118SchemaVersion":"sample","action":"Install"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects exiRequest exceeding max length",
			input:   []byte(`{"iso15118SchemaVersion":"sample","action":"Install","exiRequest":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Get15118EVCertificateRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s Get15118EVCertificateRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGet15118EVCertificateResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","exiResponse":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"exiResponse":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing exiResponse",
			input:   []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects exiResponse exceeding max length",
			input:   []byte(`{"status":"Accepted","exiResponse":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Get15118EVCertificateResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s Get15118EVCertificateResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetBaseReportRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"reportBase":"ConfigurationInventory"}`),
		},
		{
			name:    "rejects missing reportBase",
			input:   []byte(`{"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetBaseReportRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetBaseReportRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetBaseReportResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetBaseReportResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetBaseReportResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetCertificateStatusRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"ocspRequestData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}}`),
		},
		{
			name:    "rejects missing ocspRequestData",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetCertificateStatusRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetCertificateStatusRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetCertificateStatusResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","ocspResult":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"ocspResult":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects ocspResult exceeding max length",
			input:   []byte(`{"status":"Accepted","ocspResult":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetCertificateStatusResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetCertificateStatusResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetChargingProfilesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"evseId":1,"chargingProfile":{"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":1,"chargingProfileId":[1],"chargingLimitSource":["EMS"]}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetChargingProfilesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetChargingProfilesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetChargingProfilesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetChargingProfilesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetChargingProfilesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetCompositeScheduleRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"duration":1,"chargingRateUnit":"W","evseId":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetCompositeScheduleRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetCompositeScheduleRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetCompositeScheduleResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","schedule":{"evseId":1,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}]},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"schedule":{"evseId":1,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}]},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetCompositeScheduleResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetCompositeScheduleResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetDisplayMessagesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":[1],"requestId":1,"priority":"AlwaysFront","state":"Charging"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetDisplayMessagesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetDisplayMessagesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetDisplayMessagesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetDisplayMessagesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetDisplayMessagesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetInstalledCertificateIdsRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateType":["V2GRootCertificate"]}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetInstalledCertificateIdsRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetInstalledCertificateIdsRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetInstalledCertificateIdsResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","certificateHashDataChain":[{"certificateType":"V2GRootCertificate","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},"childCertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}]}],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"certificateHashDataChain":[{"certificateType":"V2GRootCertificate","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},"childCertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}]}],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetInstalledCertificateIdsResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetInstalledCertificateIdsResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetLocalListVersionRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetLocalListVersionRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetLocalListVersionRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetLocalListVersionResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"versionNumber":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetLocalListVersionResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetLocalListVersionResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetLogRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"logType":"DiagnosticsLog","requestId":1,"retries":1,"retryInterval":1,"log":{"remoteLocation":"sample","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}}`),
		},
		{
			name:    "rejects missing logType",
			input:   []byte(`{"requestId":1,"retries":1,"retryInterval":1,"log":{"remoteLocation":"sample","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing log",
			input:   []byte(`{"logType":"DiagnosticsLog","requestId":1,"retries":1,"retryInterval":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetLogRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetLogRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetLogResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","filename":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"filename":"sample","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects filename exceeding max length",
			input:   []byte(`{"status":"Accepted","filename":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetLogResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetLogResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetMonitoringReportRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"monitoringCriteria":["ThresholdMonitoring"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects monitoringCriteria exceeding max items",
			input:   []byte(`{"requestId":1,"monitoringCriteria":["ThresholdMonitoring","ThresholdMonitoring","ThresholdMonitoring","ThresholdMonitoring"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetMonitoringReportRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetMonitoringReportRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetMonitoringReportResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetMonitoringReportResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetMonitoringReportResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetReportRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"componentCriteria":["Active"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects componentCriteria exceeding max items",
			input:   []byte(`{"requestId":1,"componentCriteria":["Active","Active","Active","Active","Active"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetReportRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetReportRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetReportResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetReportResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetReportResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetTransactionStatusRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1"}`),
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetTransactionStatusRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetTransactionStatusRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetTransactionStatusResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"ongoingIndicator":true,"messagesInQueue":true}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetTransactionStatusResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetTransactionStatusResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetVariablesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"getVariableData":[{"attributeType":"Actual","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects missing getVariableData",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetVariablesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetVariablesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetVariablesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"getVariableResult":[{"attributeStatus":"Accepted","attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
		},
		{
			name:    "rejects missing getVariableResult",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetVariablesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetVariablesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestHeartbeatRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s HeartbeatRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s HeartbeatRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestHeartbeatResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"currentTime":"2024-01-01T00:00:00Z"}`),
		},
		{
			name:    "rejects missing currentTime",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s HeartbeatResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s HeartbeatResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestInstallCertificateRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateType":"V2GRootCertificate","certificate":"sample"}`),
		},
		{
			name:    "rejects missing certificateType",
			input:   []byte(`{"certificate":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing certificate",
			input:   []byte(`{"certificateType":"V2GRootCertificate"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects certificate exceeding max length",
			input:   []byte(`{"certificateType":"V2GRootCertificate","certificate":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s InstallCertificateRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s InstallCertificateRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestInstallCertificateResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s InstallCertificateResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s InstallCertificateResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestLogStatusNotificationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"BadMessage","requestId":1}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s LogStatusNotificationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s LogStatusNotificationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestLogStatusNotificationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s LogStatusNotificationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s LogStatusNotificationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestMeterValuesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
		},
		{
			name:    "rejects missing meterValue",
			input:   []byte(`{"evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s MeterValuesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s MeterValuesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestMeterValuesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s MeterValuesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s MeterValuesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyChargingLimitRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"chargingLimit":{"chargingLimitSource":"EMS","isGridCritical":true},"chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
		},
		{
			name:    "rejects missing chargingLimit",
			input:   []byte(`{"evseId":1,"chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyChargingLimitRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyChargingLimitRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyChargingLimitResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyChargingLimitResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyChargingLimitResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyCustomerInformationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"data":"sample","tbc":true,"seqNo":1,"generatedAt":"2024-01-01T00:00:00Z","requestId":1}`),
		},
		{
			name:    "rejects missing data",
			input:   []byte(`{"tbc":true,"seqNo":1,"generatedAt":"2024-01-01T00:00:00Z","requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects data exceeding max length",
			input:   []byte(`{"data":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","tbc":true,"seqNo":1,"generatedAt":"2024-01-01T00:00:00Z","requestId":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"data":"sample","tbc":true,"seqNo":1,"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyCustomerInformationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyCustomerInformationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyCustomerInformationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyCustomerInformationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyCustomerInformationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyDisplayMessagesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"tbc":true,"messageInfo":[{"id":1,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}]}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyDisplayMessagesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyDisplayMessagesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyDisplayMessagesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyDisplayMessagesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyDisplayMessagesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyEVChargingNeedsRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"maxScheduleTuples":1,"evseId":1,"chargingNeeds":{"requestedEnergyTransfer":"DC","departureTime":"2024-01-01T00:00:00Z","acChargingParameters":{"energyAmount":1,"evMinCurrent":1,"evMaxCurrent":1,"evMaxVoltage":1},"dcChargingParameters":{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":50,"evEnergyCapacity":1,"fullSoC":50,"bulkSoC":50}}}`),
		},
		{
			name:    "rejects missing chargingNeeds",
			input:   []byte(`{"maxScheduleTuples":1,"evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyEVChargingNeedsRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyEVChargingNeedsRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyEVChargingNeedsResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyEVChargingNeedsResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyEVChargingNeedsResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyEVChargingScheduleRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"timeBase":"2024-01-01T00:00:00Z","evseId":1,"chargingSchedule":{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}}`),
		},
		{
			name:    "rejects missing timeBase",
			input:   []byte(`{"evseId":1,"chargingSchedule":{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingSchedule",
			input:   []byte(`{"timeBase":"2024-01-01T00:00:00Z","evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyEVChargingScheduleRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyEVChargingScheduleRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyEVChargingScheduleResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyEVChargingScheduleResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyEVChargingScheduleResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyEventRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":1,"eventData":[{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"tbc":true,"seqNo":1,"eventData":[{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing eventData",
			input:   []byte(`{"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyEventRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyEventRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyEventResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyEventResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyEventResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyMonitoringReportRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"tbc":true,"seqNo":1,"generatedAt":"2024-01-01T00:00:00Z","monitor":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1}]}]}`),
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"requestId":1,"tbc":true,"seqNo":1,"monitor":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyMonitoringReportRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyMonitoringReportRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyMonitoringReportResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyMonitoringReportResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyMonitoringReportResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyReportRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":1,"reportData":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}]}`),
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"requestId":1,"tbc":true,"seqNo":1,"reportData":[{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyReportRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyReportRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyReportResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyReportResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyReportResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestPublishFirmwareRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"location":"sample","retries":1,"checksum":"A1","requestId":1,"retryInterval":1}`),
		},
		{
			name:    "rejects missing location",
			input:   []byte(`{"retries":1,"checksum":"A1","requestId":1,"retryInterval":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects location exceeding max length",
			input:   []byte(`{"location":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","retries":1,"checksum":"A1","requestId":1,"retryInterval":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing checksum",
			input:   []byte(`{"location":"sample","retries":1,"requestId":1,"retryInterval":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects checksum exceeding max length",
			input:   []byte(`{"location":"sample","retries":1,"checksum":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","requestId":1,"retryInterval":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects checksum outside identifierString charset",
			input:   []byte(`{"location":"sample","retries":1,"checksum":"################################","requestId":1,"retryInterval":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s PublishFirmwareRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s PublishFirmwareRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestPublishFirmwareResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s PublishFirmwareResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s PublishFirmwareResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestPublishFirmwareStatusNotificationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Idle","location":["sample"],"requestId":1}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"location":["sample"],"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s PublishFirmwareStatusNotificationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s PublishFirmwareStatusNotificationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestPublishFirmwareStatusNotificationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s PublishFirmwareStatusNotificationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s PublishFirmwareStatusNotificationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReportChargingProfilesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"chargingLimitSource":"EMS","tbc":true,"evseId":1,"chargingProfile":[{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}]}`),
		},
		{
			name:    "rejects missing chargingLimitSource",
			input:   []byte(`{"requestId":1,"tbc":true,"evseId":1,"chargingProfile":[{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingProfile",
			input:   []byte(`{"requestId":1,"chargingLimitSource":"EMS","tbc":true,"evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReportChargingProfilesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReportChargingProfilesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReportChargingProfilesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReportChargingProfilesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReportChargingProfilesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRequestStartTransactionRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"remoteStartId":1,"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"chargingProfile":{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]},"groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"evseId":1,"remoteStartId":1,"chargingProfile":{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]},"groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s RequestStartTransactionRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RequestStartTransactionRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRequestStartTransactionResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","transactionId":"A1","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"transactionId":"A1","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"status":"Accepted","transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"status":"Accepted","transactionId":"####################################","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s RequestStartTransactionResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RequestStartTransactionResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRequestStopTransactionRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1"}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s RequestStopTransactionRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RequestStopTransactionRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRequestStopTransactionResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s RequestStopTransactionResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RequestStopTransactionResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReservationStatusUpdateRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"reservationId":1,"reservationUpdateStatus":"Expired"}`),
		},
		{
			name:    "rejects missing reservationUpdateStatus",
			input:   []byte(`{"reservationId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReservationStatusUpdateRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReservationStatusUpdateRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReservationStatusUpdateResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReservationStatusUpdateResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReservationStatusUpdateResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReserveNowRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"cCCS1","evseId":1,"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
		},
		{
			name:    "rejects missing expiryDateTime",
			input:   []byte(`{"id":1,"connectorType":"cCCS1","evseId":1,"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"id":1,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"cCCS1","evseId":1,"groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReserveNowRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReserveNowRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReserveNowResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReserveNowResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReserveNowResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestResetRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"type":"Immediate","evseId":1}`),
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ResetRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ResetRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestResetResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ResetResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ResetResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSecurityEventNotificationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"type":"sample","timestamp":"2024-01-01T00:00:00Z","techInfo":"sample"}`),
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","techInfo":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects type exceeding max length",
			input:   []byte(`{"type":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","timestamp":"2024-01-01T00:00:00Z","techInfo":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"type":"sample","techInfo":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects techInfo exceeding max length",
			input:   []byte(`{"type":"sample","timestamp":"2024-01-01T00:00:00Z","techInfo":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SecurityEventNotificationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SecurityEventNotificationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSecurityEventNotificationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SecurityEventNotificationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SecurityEventNotificationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSendLocalListRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"versionNumber":1,"updateType":"Differential","localAuthorizationList":[{"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}]}`),
		},
		{
			name:    "rejects missing updateType",
			input:   []byte(`{"versionNumber":1,"localAuthorizationList":[{"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SendLocalListRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SendLocalListRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSendLocalListResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SendLocalListResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SendLocalListResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetChargingProfileRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"chargingProfile":{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}}`),
		},
		{
			name:    "rejects missing chargingProfile",
			input:   []byte(`{"evseId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetChargingProfileRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetChargingProfileRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetChargingProfileResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetChargingProfileResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetChargingProfileResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetDisplayMessageRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"message":{"id":1,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}}`),
		},
		{
			name:    "rejects missing message",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetDisplayMessageRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetDisplayMessageRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetDisplayMessageResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetDisplayMessageResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetDisplayMessageResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetMonitoringBaseRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"monitoringBase":"All"}`),
		},
		{
			name:    "rejects missing monitoringBase",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetMonitoringBaseRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetMonitoringBaseRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetMonitoringBaseResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetMonitoringBaseResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetMonitoringBaseResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetMonitoringLevelRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"severity":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetMonitoringLevelRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetMonitoringLevelRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetMonitoringLevelResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetMonitoringLevelResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetMonitoringLevelResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetNetworkProfileRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"configurationSlot":1,"connectionData":{"ocppVersion":"OCPP12","ocppTransport":"JSON","ocppCsmsUrl":"sample","messageTimeout":1,"securityProfile":1,"ocppInterface":"Wired0","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}}`),
		},
		{
			name:    "rejects missing connectionData",
			input:   []byte(`{"configurationSlot":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetNetworkProfileRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetNetworkProfileRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetNetworkProfileResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetNetworkProfileResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetNetworkProfileResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetVariableMonitoringRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"setMonitoringData":[{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects missing setMonitoringData",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetVariableMonitoringRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetVariableMonitoringRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetVariableMonitoringResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"setMonitoringResult":[{"id":1,"status":"Accepted","type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
		},
		{
			name:    "rejects missing setMonitoringResult",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetVariableMonitoringResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetVariableMonitoringResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetVariablesRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"setVariableData":[{"attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects missing setVariableData",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetVariablesRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetVariablesRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetVariablesResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"setVariableResult":[{"attributeType":"Actual","attributeStatus":"Accepted","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
		},
		{
			name:    "rejects missing setVariableResult",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetVariablesResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetVariablesResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSignCertificateRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"csr":"sample","certificateType":"ChargingStationCertificate"}`),
		},
		{
			name:    "rejects missing csr",
			input:   []byte(`{"certificateType":"ChargingStationCertificate"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects csr exceeding max length",
			input:   []byte(`{"csr":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","certificateType":"ChargingStationCertificate"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SignCertificateRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SignCertificateRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSignCertificateResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SignCertificateResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SignCertificateResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestStatusNotificationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"timestamp":"2024-01-01T00:00:00Z","connectorStatus":"Available","evseId":1,"connectorId":1}`),
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"connectorStatus":"Available","evseId":1,"connectorId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing connectorStatus",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","evseId":1,"connectorId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s StatusNotificationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s StatusNotificationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestStatusNotificationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s StatusNotificationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s StatusNotificationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestTransactionEventRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"Authorized","seqNo":1,"offline":true,"numberOfPhasesUsed":1,"cableMaxCurrent":1,"reservationId":1,"transactionInfo":{"transactionId":"A1","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":1,"connectorId":1},"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
		},
		{
			name:    "rejects missing eventType",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","triggerReason":"Authorized","seqNo":1,"offline":true,"numberOfPhasesUsed":1,"cableMaxCurrent":1,"reservationId":1,"transactionInfo":{"transactionId":"A1","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":1,"connectorId":1},"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"eventType":"Ended","triggerReason":"Authorized","seqNo":1,"offline":true,"numberOfPhasesUsed":1,"cableMaxCurrent":1,"reservationId":1,"transactionInfo":{"transactionId":"A1","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":1,"connectorId":1},"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing triggerReason",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","seqNo":1,"offline":true,"numberOfPhasesUsed":1,"cableMaxCurrent":1,"reservationId":1,"transactionInfo":{"transactionId":"A1","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":1,"connectorId":1},"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing transactionInfo",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"Authorized","seqNo":1,"offline":true,"numberOfPhasesUsed":1,"cableMaxCurrent":1,"reservationId":1,"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":1,"connectorId":1},"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s TransactionEventRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s TransactionEventRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestTransactionEventResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"totalCost":1.0,"chargingPriority":1,"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"updatedPersonalMessage":{"format":"ASCII","language":"sample","content":"sample"}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s TransactionEventResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s TransactionEventResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestTriggerMessageRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestedMessage":"BootNotification","evse":{"id":1,"connectorId":1}}`),
		},
		{
			name:    "rejects missing requestedMessage",
			input:   []byte(`{"evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s TriggerMessageRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s TriggerMessageRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestTriggerMessageResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s TriggerMessageResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s TriggerMessageResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUnlockConnectorRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"connectorId":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UnlockConnectorRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UnlockConnectorRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUnlockConnectorResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Unlocked","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UnlockConnectorResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UnlockConnectorResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUnpublishFirmwareRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"checksum":"A1"}`),
		},
		{
			name:    "rejects missing checksum",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects checksum exceeding max length",
			input:   []byte(`{"checksum":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects checksum outside identifierString charset",
			input:   []byte(`{"checksum":"################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UnpublishFirmwareRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UnpublishFirmwareRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUnpublishFirmwareResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"DownloadOngoing"}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UnpublishFirmwareResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UnpublishFirmwareResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUpdateFirmwareRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"retries":1,"retryInterval":1,"requestId":1,"firmware":{"location":"sample","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}}`),
		},
		{
			name:    "rejects missing firmware",
			input:   []byte(`{"retries":1,"retryInterval":1,"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UpdateFirmwareRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UpdateFirmwareRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUpdateFirmwareResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UpdateFirmwareResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UpdateFirmwareResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}
