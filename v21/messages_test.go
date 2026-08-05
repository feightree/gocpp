package v21

import (
	"encoding/json"
	"errors"
	"testing"

	ocpp "github.com/feightree/gocpp/ocpp"
)

func TestAdjustPeriodicEventStreamRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":0,"params":{"interval":0,"values":0}}`),
		},
		{
			name:    "rejects id out of range",
			input:   []byte(`{"id":-1,"params":{"interval":0,"values":0}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s AdjustPeriodicEventStreamRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AdjustPeriodicEventStreamRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAdjustPeriodicEventStreamResponseUnmarshalJSON(t *testing.T) {
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
			var s AdjustPeriodicEventStreamResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AdjustPeriodicEventStreamResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAFRRSignalRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"timestamp":"2024-01-01T00:00:00Z","signal":1}`),
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"signal":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s AFRRSignalRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AFRRSignalRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAFRRSignalResponseUnmarshalJSON(t *testing.T) {
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
			var s AFRRSignalResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AFRRSignalResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAuthorizeRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificate":"sample","idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"certificate":"sample","iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects certificate exceeding max length",
			input:   []byte(`{"certificate":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects iso15118CertificateHashData exceeding max items",
			input:   []byte(`{"certificate":"sample","idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"iso15118CertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}]}`),
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
			input: []byte(`{"certificateStatus":"Accepted","allowedEnergyTransfer":["AC_single_phase"],"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","language2":"sample","evseId":[0],"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
		},
		{
			name:    "rejects missing idTokenInfo",
			input:   []byte(`{"certificateStatus":"Accepted","allowedEnergyTransfer":["AC_single_phase"],"tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
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

func TestBatterySwapRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"eventType":"BatteryIn","requestId":1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"batteryData":[{"evseId":0,"serialNumber":"sample","soC":50.0,"soH":50.0,"productionDate":"2024-01-01T00:00:00Z","vendorInfo":"sample"}]}`),
		},
		{
			name:    "rejects missing eventType",
			input:   []byte(`{"requestId":1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"batteryData":[{"evseId":0,"serialNumber":"sample","soC":50.0,"soH":50.0,"productionDate":"2024-01-01T00:00:00Z","vendorInfo":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"eventType":"BatteryIn","requestId":1,"batteryData":[{"evseId":0,"serialNumber":"sample","soC":50.0,"soH":50.0,"productionDate":"2024-01-01T00:00:00Z","vendorInfo":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing batteryData",
			input:   []byte(`{"eventType":"BatteryIn","requestId":1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s BatterySwapRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s BatterySwapRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestBatterySwapResponseUnmarshalJSON(t *testing.T) {
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
			var s BatterySwapResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s BatterySwapResponse
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
			input: []byte(`{"reservationId":0}`),
		},
		{
			name:    "rejects reservationId out of range",
			input:   []byte(`{"reservationId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"certificateChain":"sample","certificateType":"ChargingStationCertificate","requestId":1}`),
		},
		{
			name:    "rejects missing certificateChain",
			input:   []byte(`{"certificateType":"ChargingStationCertificate","requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects certificateChain exceeding max length",
			input:   []byte(`{"certificateChain":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","certificateType":"ChargingStationCertificate","requestId":1}`),
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
			input: []byte(`{"operationalStatus":"Inoperative","evse":{"id":0,"connectorId":0}}`),
		},
		{
			name:    "rejects missing operationalStatus",
			input:   []byte(`{"evse":{"id":0,"connectorId":0}}`),
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

func TestChangeTransactionTariffRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1","tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{"tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing tariff",
			input:   []byte(`{"transactionId":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################","tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChangeTransactionTariffRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChangeTransactionTariffRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChangeTransactionTariffResponseUnmarshalJSON(t *testing.T) {
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
			var s ChangeTransactionTariffResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChangeTransactionTariffResponse
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
			input: []byte(`{"chargingProfileId":1,"chargingProfileCriteria":{"evseId":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":0}}`),
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

func TestClearDERControlRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"isDefault":true,"controlType":"EnterService","controlId":"A1"}`),
		},
		{
			name:    "rejects controlId exceeding max length",
			input:   []byte(`{"isDefault":true,"controlType":"EnterService","controlId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects controlId outside identifierString charset",
			input:   []byte(`{"isDefault":true,"controlType":"EnterService","controlId":"####################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearDERControlRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearDERControlRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearDERControlResponseUnmarshalJSON(t *testing.T) {
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
			var s ClearDERControlResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearDERControlResponse
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
			input: []byte(`{"id":0}`),
		},
		{
			name:    "rejects id out of range",
			input:   []byte(`{"id":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"chargingLimitSource":"sample","evseId":0}`),
		},
		{
			name:    "rejects missing chargingLimitSource",
			input:   []byte(`{"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects chargingLimitSource exceeding max length",
			input:   []byte(`{"chargingLimitSource":"xxxxxxxxxxxxxxxxxxxxx","evseId":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"chargingLimitSource":"sample","evseId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestClearTariffsRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"tariffIds":["sample"],"evseId":0}`),
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"tariffIds":["sample"],"evseId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearTariffsRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearTariffsRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearTariffsResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"clearTariffsResult":[{"tariffId":"sample","status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
		},
		{
			name:    "rejects missing clearTariffsResult",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearTariffsResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearTariffsResponse
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
			input: []byte(`{"id":[0]}`),
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
			input: []byte(`{"clearMonitoringResult":[{"status":"Accepted","id":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
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

func TestClosePeriodicEventStreamRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":0}`),
		},
		{
			name:    "rejects id out of range",
			input:   []byte(`{"id":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClosePeriodicEventStreamRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClosePeriodicEventStreamRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClosePeriodicEventStreamResponseUnmarshalJSON(t *testing.T) {
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
			var s ClosePeriodicEventStreamResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClosePeriodicEventStreamResponse
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
			input: []byte(`{"requestId":0,"report":true,"clear":true,"customerIdentifier":"sample","idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"customerCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
		},
		{
			name:    "rejects requestId out of range",
			input:   []byte(`{"requestId":-1,"report":true,"clear":true,"customerIdentifier":"sample","idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"customerCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects customerIdentifier exceeding max length",
			input:   []byte(`{"requestId":0,"report":true,"clear":true,"customerIdentifier":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"customerCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
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
			input: []byte(`{"messageId":"sample","data":{"key":"value"},"vendorId":"sample"}`),
		},
		{
			name:    "rejects missing vendorId",
			input:   []byte(`{"messageId":"sample","data":{"key":"value"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects messageId exceeding max length",
			input:   []byte(`{"messageId":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","data":{"key":"value"},"vendorId":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects vendorId exceeding max length",
			input:   []byte(`{"messageId":"sample","data":{"key":"value"},"vendorId":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
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
			input: []byte(`{"status":"Accepted","data":{"key":"value"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"data":{"key":"value"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
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
			input: []byte(`{"status":"Downloaded","requestId":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"requestId":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
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
			input: []byte(`{"iso15118SchemaVersion":"sample","action":"Install","exiRequest":"sample","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1"]}`),
		},
		{
			name:    "rejects missing iso15118SchemaVersion",
			input:   []byte(`{"action":"Install","exiRequest":"sample","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1"]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing action",
			input:   []byte(`{"iso15118SchemaVersion":"sample","exiRequest":"sample","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1"]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing exiRequest",
			input:   []byte(`{"iso15118SchemaVersion":"sample","action":"Install","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1"]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects iso15118SchemaVersion exceeding max length",
			input:   []byte(`{"iso15118SchemaVersion":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","action":"Install","exiRequest":"sample","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects exiRequest exceeding max length",
			input:   []byte(`{"iso15118SchemaVersion":"sample","action":"Install","exiRequest":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects maximumContractCertificateChains out of range",
			input:   []byte(`{"iso15118SchemaVersion":"sample","action":"Install","exiRequest":"sample","maximumContractCertificateChains":-1,"prioritizedEMAIDs":["A1"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects prioritizedEMAIDs exceeding max items",
			input:   []byte(`{"iso15118SchemaVersion":"sample","action":"Install","exiRequest":"sample","maximumContractCertificateChains":0,"prioritizedEMAIDs":["A1","A1","A1","A1","A1","A1","A1","A1","A1"]}`),
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
			input: []byte(`{"status":"Accepted","exiResponse":"sample","remainingContracts":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"exiResponse":"sample","remainingContracts":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing exiResponse",
			input:   []byte(`{"status":"Accepted","remainingContracts":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects exiResponse exceeding max length",
			input:   []byte(`{"status":"Accepted","exiResponse":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","remainingContracts":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects remainingContracts out of range",
			input:   []byte(`{"status":"Accepted","exiResponse":"sample","remainingContracts":-1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
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

func TestGetCertificateChainStatusRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateStatusRequests":[{"source":"CRL","urls":["sample"],"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}]}`),
		},
		{
			name:    "rejects missing certificateStatusRequests",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects certificateStatusRequests exceeding max items",
			input:   []byte(`{"certificateStatusRequests":[{"source":"CRL","urls":["sample"],"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","urls":["sample"],"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","urls":["sample"],"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","urls":["sample"],"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","urls":["sample"],"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetCertificateChainStatusRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetCertificateChainStatusRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetCertificateChainStatusResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateStatus":[{"source":"CRL","status":"Good","nextUpdate":"2024-01-01T00:00:00Z","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}]}`),
		},
		{
			name:    "rejects missing certificateStatus",
			input:   []byte(`{}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects certificateStatus exceeding max items",
			input:   []byte(`{"certificateStatus":[{"source":"CRL","status":"Good","nextUpdate":"2024-01-01T00:00:00Z","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","status":"Good","nextUpdate":"2024-01-01T00:00:00Z","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","status":"Good","nextUpdate":"2024-01-01T00:00:00Z","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","status":"Good","nextUpdate":"2024-01-01T00:00:00Z","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}},{"source":"CRL","status":"Good","nextUpdate":"2024-01-01T00:00:00Z","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetCertificateChainStatusResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetCertificateChainStatusResponse
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
			input:   []byte(`{"status":"Accepted","ocspResult":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
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
			input: []byte(`{"requestId":1,"evseId":0,"chargingProfile":{"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":0,"chargingProfileId":[1],"chargingLimitSource":["sample"]}}`),
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"requestId":1,"evseId":-1,"chargingProfile":{"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":0,"chargingProfileId":[1],"chargingLimitSource":["sample"]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"duration":1,"chargingRateUnit":"W","evseId":0}`),
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"duration":1,"chargingRateUnit":"W","evseId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"status":"Accepted","statusInfo":{"reasonCode":"sample","additionalInfo":"sample"},"schedule":{"evseId":0,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}]}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"},"schedule":{"evseId":0,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}]}}`),
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

func TestGetDERControlRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"isDefault":true,"controlType":"EnterService","controlId":"A1"}`),
		},
		{
			name:    "rejects controlId exceeding max length",
			input:   []byte(`{"requestId":1,"isDefault":true,"controlType":"EnterService","controlId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects controlId outside identifierString charset",
			input:   []byte(`{"requestId":1,"isDefault":true,"controlType":"EnterService","controlId":"####################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetDERControlRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetDERControlRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetDERControlResponseUnmarshalJSON(t *testing.T) {
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
			var s GetDERControlResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetDERControlResponse
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
			input: []byte(`{"id":[0],"requestId":1,"priority":"AlwaysFront","state":"Charging"}`),
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
			input: []byte(`{"logType":"DiagnosticsLog","requestId":1,"retries":0,"retryInterval":1,"log":{"remoteLocation":"sample","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}}`),
		},
		{
			name:    "rejects missing logType",
			input:   []byte(`{"requestId":1,"retries":0,"retryInterval":1,"log":{"remoteLocation":"sample","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing log",
			input:   []byte(`{"logType":"DiagnosticsLog","requestId":1,"retries":0,"retryInterval":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects retries out of range",
			input:   []byte(`{"logType":"DiagnosticsLog","requestId":1,"retries":-1,"retryInterval":1,"log":{"remoteLocation":"sample","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"requestId":1,"monitoringCriteria":["ThresholdMonitoring"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects monitoringCriteria exceeding max items",
			input:   []byte(`{"requestId":1,"monitoringCriteria":["ThresholdMonitoring","ThresholdMonitoring","ThresholdMonitoring","ThresholdMonitoring"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
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

func TestGetPeriodicEventStreamRequestUnmarshalJSON(t *testing.T) {
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
			var s GetPeriodicEventStreamRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetPeriodicEventStreamRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetPeriodicEventStreamResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"constantStreamData":[{"id":0,"variableMonitoringId":0,"params":{"interval":0,"values":0}}]}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetPeriodicEventStreamResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetPeriodicEventStreamResponse
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
			input: []byte(`{"requestId":1,"componentCriteria":["Active"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects componentCriteria exceeding max items",
			input:   []byte(`{"requestId":1,"componentCriteria":["Active","Active","Active","Active","Active"],"componentVariable":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
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

func TestGetTariffsRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":0}`),
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetTariffsRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetTariffsRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetTariffsResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","tariffAssignments":[{"tariffId":"sample","tariffKind":"DefaultTariff","validFrom":"2024-01-01T00:00:00Z","evseIds":[0],"idTokens":["A1"]}],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"tariffAssignments":[{"tariffId":"sample","tariffKind":"DefaultTariff","validFrom":"2024-01-01T00:00:00Z","evseIds":[0],"idTokens":["A1"]}],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetTariffsResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetTariffsResponse
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
			input: []byte(`{"getVariableData":[{"attributeType":"Actual","component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
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
			input: []byte(`{"getVariableResult":[{"attributeStatus":"Accepted","attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
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
			input:   []byte(`{"certificateType":"V2GRootCertificate","certificate":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
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
			input: []byte(`{"status":"BadMessage","requestId":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"requestId":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
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
			input: []byte(`{"evseId":0,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
		},
		{
			name:    "rejects missing meterValue",
			input:   []byte(`{"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestNotifyAllowedEnergyTransferRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1","allowedEnergyTransfer":["AC_single_phase"]}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{"allowedEnergyTransfer":["AC_single_phase"]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing allowedEnergyTransfer",
			input:   []byte(`{"transactionId":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","allowedEnergyTransfer":["AC_single_phase"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################","allowedEnergyTransfer":["AC_single_phase"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyAllowedEnergyTransferRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyAllowedEnergyTransferRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyAllowedEnergyTransferResponseUnmarshalJSON(t *testing.T) {
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
			var s NotifyAllowedEnergyTransferResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyAllowedEnergyTransferResponse
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
			input: []byte(`{"evseId":0,"chargingLimit":{"chargingLimitSource":"sample","isLocalGeneration":true,"isGridCritical":true},"chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}`),
		},
		{
			name:    "rejects missing chargingLimit",
			input:   []byte(`{"evseId":0,"chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1,"chargingLimit":{"chargingLimitSource":"sample","isLocalGeneration":true,"isGridCritical":true},"chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"data":"sample","tbc":true,"seqNo":0,"generatedAt":"2024-01-01T00:00:00Z","requestId":0}`),
		},
		{
			name:    "rejects missing data",
			input:   []byte(`{"tbc":true,"seqNo":0,"generatedAt":"2024-01-01T00:00:00Z","requestId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"data":"sample","tbc":true,"seqNo":0,"requestId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects data exceeding max length",
			input:   []byte(`{"data":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","tbc":true,"seqNo":0,"generatedAt":"2024-01-01T00:00:00Z","requestId":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects seqNo out of range",
			input:   []byte(`{"data":"sample","tbc":true,"seqNo":-1,"generatedAt":"2024-01-01T00:00:00Z","requestId":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects requestId out of range",
			input:   []byte(`{"data":"sample","tbc":true,"seqNo":0,"generatedAt":"2024-01-01T00:00:00Z","requestId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestNotifyDERAlarmRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"controlType":"EnterService","gridEventFault":"CurrentImbalance","alarmEnded":true,"timestamp":"2024-01-01T00:00:00Z","extraInfo":"sample"}`),
		},
		{
			name:    "rejects missing controlType",
			input:   []byte(`{"gridEventFault":"CurrentImbalance","alarmEnded":true,"timestamp":"2024-01-01T00:00:00Z","extraInfo":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"controlType":"EnterService","gridEventFault":"CurrentImbalance","alarmEnded":true,"extraInfo":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects extraInfo exceeding max length",
			input:   []byte(`{"controlType":"EnterService","gridEventFault":"CurrentImbalance","alarmEnded":true,"timestamp":"2024-01-01T00:00:00Z","extraInfo":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyDERAlarmRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyDERAlarmRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyDERAlarmResponseUnmarshalJSON(t *testing.T) {
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
			var s NotifyDERAlarmResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyDERAlarmResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyDERStartStopRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"controlId":"A1","started":true,"timestamp":"2024-01-01T00:00:00Z","supersededIds":["A1"]}`),
		},
		{
			name:    "rejects missing controlId",
			input:   []byte(`{"started":true,"timestamp":"2024-01-01T00:00:00Z","supersededIds":["A1"]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"controlId":"A1","started":true,"supersededIds":["A1"]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects controlId exceeding max length",
			input:   []byte(`{"controlId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","started":true,"timestamp":"2024-01-01T00:00:00Z","supersededIds":["A1"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects controlId outside identifierString charset",
			input:   []byte(`{"controlId":"####################################","started":true,"timestamp":"2024-01-01T00:00:00Z","supersededIds":["A1"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects supersededIds exceeding max items",
			input:   []byte(`{"controlId":"A1","started":true,"timestamp":"2024-01-01T00:00:00Z","supersededIds":["A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyDERStartStopRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyDERStartStopRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyDERStartStopResponseUnmarshalJSON(t *testing.T) {
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
			var s NotifyDERStartStopResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyDERStartStopResponse
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
			input: []byte(`{"requestId":1,"tbc":true,"messageInfo":[{"id":0,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"messageExtra":[{"format":"ASCII","language":"sample","content":"sample"}]}]}`),
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
			input: []byte(`{"evseId":1,"maxScheduleTuples":0,"timestamp":"2024-01-01T00:00:00Z","chargingNeeds":{"requestedEnergyTransfer":"AC_single_phase","availableEnergyTransfer":["AC_single_phase"],"controlMode":"ScheduledControl","mobilityNeedsMode":"EVCC","departureTime":"2024-01-01T00:00:00Z","v2xChargingParameters":{"minChargePower":1.0,"minChargePower_L2":1.0,"minChargePower_L3":1.0,"maxChargePower":1.0,"maxChargePower_L2":1.0,"maxChargePower_L3":1.0,"minDischargePower":1.0,"minDischargePower_L2":1.0,"minDischargePower_L3":1.0,"maxDischargePower":1.0,"maxDischargePower_L2":1.0,"maxDischargePower_L3":1.0,"minChargeCurrent":1.0,"maxChargeCurrent":1.0,"minDischargeCurrent":1.0,"maxDischargeCurrent":1.0,"minVoltage":1.0,"maxVoltage":1.0,"evTargetEnergyRequest":1.0,"evMinEnergyRequest":1.0,"evMaxEnergyRequest":1.0,"evMinV2XEnergyRequest":1.0,"evMaxV2XEnergyRequest":1.0,"targetSoC":50},"dcChargingParameters":{"evMaxCurrent":1.0,"evMaxVoltage":1.0,"evMaxPower":1.0,"evEnergyCapacity":1.0,"energyAmount":1.0,"stateOfCharge":50,"fullSoC":50,"bulkSoC":50},"acChargingParameters":{"energyAmount":1.0,"evMinCurrent":1.0,"evMaxCurrent":1.0,"evMaxVoltage":1.0},"evEnergyOffer":{"evPowerSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","evPowerScheduleEntries":[{"duration":1,"power":1.0}]},"evAbsolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","currency":"sam","priceAlgorithm":"sample","evAbsolutePriceScheduleEntries":[{"duration":1,"evPriceRule":[{"energyFee":1.0,"powerRangeStart":1.0}]}]}},"derChargingParameters":{"evSupportedDERControl":["EnterService"],"evOverExcitedMaxDischargePower":1.0,"evOverExcitedPowerFactor":1.0,"evUnderExcitedMaxDischargePower":1.0,"evUnderExcitedPowerFactor":1.0,"maxApparentPower":1.0,"maxChargeApparentPower":1.0,"maxChargeApparentPower_L2":1.0,"maxChargeApparentPower_L3":1.0,"maxDischargeApparentPower":1.0,"maxDischargeApparentPower_L2":1.0,"maxDischargeApparentPower_L3":1.0,"maxChargeReactivePower":1.0,"maxChargeReactivePower_L2":1.0,"maxChargeReactivePower_L3":1.0,"minChargeReactivePower":1.0,"minChargeReactivePower_L2":1.0,"minChargeReactivePower_L3":1.0,"maxDischargeReactivePower":1.0,"maxDischargeReactivePower_L2":1.0,"maxDischargeReactivePower_L3":1.0,"minDischargeReactivePower":1.0,"minDischargeReactivePower_L2":1.0,"minDischargeReactivePower_L3":1.0,"nominalVoltage":1.0,"nominalVoltageOffset":1.0,"maxNominalVoltage":1.0,"minNominalVoltage":1.0,"evInverterManufacturer":"sample","evInverterModel":"sample","evInverterSerialNumber":"sample","evInverterSwVersion":"sample","evInverterHwVersion":"sample","evIslandingDetectionMethod":["NoAntiIslandingSupport"],"evIslandingTripTime":1.0,"evMaximumLevel1DCInjection":1.0,"evDurationLevel1DCInjection":1.0,"evMaximumLevel2DCInjection":1.0,"evDurationLevel2DCInjection":1.0,"evReactiveSusceptance":1.0,"evSessionTotalDischargeEnergyAvailable":1.0}}}`),
		},
		{
			name:    "rejects missing chargingNeeds",
			input:   []byte(`{"evseId":1,"maxScheduleTuples":0,"timestamp":"2024-01-01T00:00:00Z"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":0,"maxScheduleTuples":0,"timestamp":"2024-01-01T00:00:00Z","chargingNeeds":{"requestedEnergyTransfer":"AC_single_phase","availableEnergyTransfer":["AC_single_phase"],"controlMode":"ScheduledControl","mobilityNeedsMode":"EVCC","departureTime":"2024-01-01T00:00:00Z","v2xChargingParameters":{"minChargePower":1.0,"minChargePower_L2":1.0,"minChargePower_L3":1.0,"maxChargePower":1.0,"maxChargePower_L2":1.0,"maxChargePower_L3":1.0,"minDischargePower":1.0,"minDischargePower_L2":1.0,"minDischargePower_L3":1.0,"maxDischargePower":1.0,"maxDischargePower_L2":1.0,"maxDischargePower_L3":1.0,"minChargeCurrent":1.0,"maxChargeCurrent":1.0,"minDischargeCurrent":1.0,"maxDischargeCurrent":1.0,"minVoltage":1.0,"maxVoltage":1.0,"evTargetEnergyRequest":1.0,"evMinEnergyRequest":1.0,"evMaxEnergyRequest":1.0,"evMinV2XEnergyRequest":1.0,"evMaxV2XEnergyRequest":1.0,"targetSoC":50},"dcChargingParameters":{"evMaxCurrent":1.0,"evMaxVoltage":1.0,"evMaxPower":1.0,"evEnergyCapacity":1.0,"energyAmount":1.0,"stateOfCharge":50,"fullSoC":50,"bulkSoC":50},"acChargingParameters":{"energyAmount":1.0,"evMinCurrent":1.0,"evMaxCurrent":1.0,"evMaxVoltage":1.0},"evEnergyOffer":{"evPowerSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","evPowerScheduleEntries":[{"duration":1,"power":1.0}]},"evAbsolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","currency":"sam","priceAlgorithm":"sample","evAbsolutePriceScheduleEntries":[{"duration":1,"evPriceRule":[{"energyFee":1.0,"powerRangeStart":1.0}]}]}},"derChargingParameters":{"evSupportedDERControl":["EnterService"],"evOverExcitedMaxDischargePower":1.0,"evOverExcitedPowerFactor":1.0,"evUnderExcitedMaxDischargePower":1.0,"evUnderExcitedPowerFactor":1.0,"maxApparentPower":1.0,"maxChargeApparentPower":1.0,"maxChargeApparentPower_L2":1.0,"maxChargeApparentPower_L3":1.0,"maxDischargeApparentPower":1.0,"maxDischargeApparentPower_L2":1.0,"maxDischargeApparentPower_L3":1.0,"maxChargeReactivePower":1.0,"maxChargeReactivePower_L2":1.0,"maxChargeReactivePower_L3":1.0,"minChargeReactivePower":1.0,"minChargeReactivePower_L2":1.0,"minChargeReactivePower_L3":1.0,"maxDischargeReactivePower":1.0,"maxDischargeReactivePower_L2":1.0,"maxDischargeReactivePower_L3":1.0,"minDischargeReactivePower":1.0,"minDischargeReactivePower_L2":1.0,"minDischargeReactivePower_L3":1.0,"nominalVoltage":1.0,"nominalVoltageOffset":1.0,"maxNominalVoltage":1.0,"minNominalVoltage":1.0,"evInverterManufacturer":"sample","evInverterModel":"sample","evInverterSerialNumber":"sample","evInverterSwVersion":"sample","evInverterHwVersion":"sample","evIslandingDetectionMethod":["NoAntiIslandingSupport"],"evIslandingTripTime":1.0,"evMaximumLevel1DCInjection":1.0,"evDurationLevel1DCInjection":1.0,"evMaximumLevel2DCInjection":1.0,"evDurationLevel2DCInjection":1.0,"evReactiveSusceptance":1.0,"evSessionTotalDischargeEnergyAvailable":1.0}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects maxScheduleTuples out of range",
			input:   []byte(`{"evseId":1,"maxScheduleTuples":-1,"timestamp":"2024-01-01T00:00:00Z","chargingNeeds":{"requestedEnergyTransfer":"AC_single_phase","availableEnergyTransfer":["AC_single_phase"],"controlMode":"ScheduledControl","mobilityNeedsMode":"EVCC","departureTime":"2024-01-01T00:00:00Z","v2xChargingParameters":{"minChargePower":1.0,"minChargePower_L2":1.0,"minChargePower_L3":1.0,"maxChargePower":1.0,"maxChargePower_L2":1.0,"maxChargePower_L3":1.0,"minDischargePower":1.0,"minDischargePower_L2":1.0,"minDischargePower_L3":1.0,"maxDischargePower":1.0,"maxDischargePower_L2":1.0,"maxDischargePower_L3":1.0,"minChargeCurrent":1.0,"maxChargeCurrent":1.0,"minDischargeCurrent":1.0,"maxDischargeCurrent":1.0,"minVoltage":1.0,"maxVoltage":1.0,"evTargetEnergyRequest":1.0,"evMinEnergyRequest":1.0,"evMaxEnergyRequest":1.0,"evMinV2XEnergyRequest":1.0,"evMaxV2XEnergyRequest":1.0,"targetSoC":50},"dcChargingParameters":{"evMaxCurrent":1.0,"evMaxVoltage":1.0,"evMaxPower":1.0,"evEnergyCapacity":1.0,"energyAmount":1.0,"stateOfCharge":50,"fullSoC":50,"bulkSoC":50},"acChargingParameters":{"energyAmount":1.0,"evMinCurrent":1.0,"evMaxCurrent":1.0,"evMaxVoltage":1.0},"evEnergyOffer":{"evPowerSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","evPowerScheduleEntries":[{"duration":1,"power":1.0}]},"evAbsolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","currency":"sam","priceAlgorithm":"sample","evAbsolutePriceScheduleEntries":[{"duration":1,"evPriceRule":[{"energyFee":1.0,"powerRangeStart":1.0}]}]}},"derChargingParameters":{"evSupportedDERControl":["EnterService"],"evOverExcitedMaxDischargePower":1.0,"evOverExcitedPowerFactor":1.0,"evUnderExcitedMaxDischargePower":1.0,"evUnderExcitedPowerFactor":1.0,"maxApparentPower":1.0,"maxChargeApparentPower":1.0,"maxChargeApparentPower_L2":1.0,"maxChargeApparentPower_L3":1.0,"maxDischargeApparentPower":1.0,"maxDischargeApparentPower_L2":1.0,"maxDischargeApparentPower_L3":1.0,"maxChargeReactivePower":1.0,"maxChargeReactivePower_L2":1.0,"maxChargeReactivePower_L3":1.0,"minChargeReactivePower":1.0,"minChargeReactivePower_L2":1.0,"minChargeReactivePower_L3":1.0,"maxDischargeReactivePower":1.0,"maxDischargeReactivePower_L2":1.0,"maxDischargeReactivePower_L3":1.0,"minDischargeReactivePower":1.0,"minDischargeReactivePower_L2":1.0,"minDischargeReactivePower_L3":1.0,"nominalVoltage":1.0,"nominalVoltageOffset":1.0,"maxNominalVoltage":1.0,"minNominalVoltage":1.0,"evInverterManufacturer":"sample","evInverterModel":"sample","evInverterSerialNumber":"sample","evInverterSwVersion":"sample","evInverterHwVersion":"sample","evIslandingDetectionMethod":["NoAntiIslandingSupport"],"evIslandingTripTime":1.0,"evMaximumLevel1DCInjection":1.0,"evDurationLevel1DCInjection":1.0,"evMaximumLevel2DCInjection":1.0,"evDurationLevel2DCInjection":1.0,"evReactiveSusceptance":1.0,"evSessionTotalDischargeEnergyAvailable":1.0}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"timeBase":"2024-01-01T00:00:00Z","evseId":1,"selectedChargingScheduleId":0,"powerToleranceAcceptance":true,"chargingSchedule":{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}}`),
		},
		{
			name:    "rejects missing timeBase",
			input:   []byte(`{"evseId":1,"selectedChargingScheduleId":0,"powerToleranceAcceptance":true,"chargingSchedule":{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingSchedule",
			input:   []byte(`{"timeBase":"2024-01-01T00:00:00Z","evseId":1,"selectedChargingScheduleId":0,"powerToleranceAcceptance":true}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"timeBase":"2024-01-01T00:00:00Z","evseId":0,"selectedChargingScheduleId":0,"powerToleranceAcceptance":true,"chargingSchedule":{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects selectedChargingScheduleId out of range",
			input:   []byte(`{"timeBase":"2024-01-01T00:00:00Z","evseId":1,"selectedChargingScheduleId":-1,"powerToleranceAcceptance":true,"chargingSchedule":{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":0,"eventData":[{"eventId":0,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":0,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":0,"eventNotificationType":"HardWiredNotification","severity":0,"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"tbc":true,"seqNo":0,"eventData":[{"eventId":0,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":0,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":0,"eventNotificationType":"HardWiredNotification","severity":0,"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing eventData",
			input:   []byte(`{"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects seqNo out of range",
			input:   []byte(`{"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":-1,"eventData":[{"eventId":0,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":0,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":0,"eventNotificationType":"HardWiredNotification","severity":0,"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"requestId":1,"tbc":true,"seqNo":0,"generatedAt":"2024-01-01T00:00:00Z","monitor":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":0,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":0,"eventNotificationType":"HardWiredNotification"}]}]}`),
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"requestId":1,"tbc":true,"seqNo":0,"monitor":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":0,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":0,"eventNotificationType":"HardWiredNotification"}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects seqNo out of range",
			input:   []byte(`{"requestId":1,"tbc":true,"seqNo":-1,"generatedAt":"2024-01-01T00:00:00Z","monitor":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":0,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":0,"eventNotificationType":"HardWiredNotification"}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestNotifyPeriodicEventStreamRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":0,"pending":0,"basetime":"2024-01-01T00:00:00Z","data":[{"t":1.0,"v":"sample"}]}`),
		},
		{
			name:    "rejects missing basetime",
			input:   []byte(`{"id":0,"pending":0,"data":[{"t":1.0,"v":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing data",
			input:   []byte(`{"id":0,"pending":0,"basetime":"2024-01-01T00:00:00Z"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects id out of range",
			input:   []byte(`{"id":-1,"pending":0,"basetime":"2024-01-01T00:00:00Z","data":[{"t":1.0,"v":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects pending out of range",
			input:   []byte(`{"id":0,"pending":-1,"basetime":"2024-01-01T00:00:00Z","data":[{"t":1.0,"v":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyPeriodicEventStreamRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyPeriodicEventStreamRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyPriorityChargingRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1","activated":true}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{"activated":true}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","activated":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################","activated":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyPriorityChargingRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyPriorityChargingRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyPriorityChargingResponseUnmarshalJSON(t *testing.T) {
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
			var s NotifyPriorityChargingResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyPriorityChargingResponse
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
			input: []byte(`{"requestId":1,"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":0,"reportData":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"maxElements":1,"valuesList":"sample","supportsMonitoring":true}}]}`),
		},
		{
			name:    "rejects missing generatedAt",
			input:   []byte(`{"requestId":1,"tbc":true,"seqNo":0,"reportData":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"maxElements":1,"valuesList":"sample","supportsMonitoring":true}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects seqNo out of range",
			input:   []byte(`{"requestId":1,"generatedAt":"2024-01-01T00:00:00Z","tbc":true,"seqNo":-1,"reportData":[{"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"maxElements":1,"valuesList":"sample","supportsMonitoring":true}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestNotifySettlementRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
		},
		{
			name:    "rejects missing pspRef",
			input:   []byte(`{"transactionId":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"transactionId":"A1","pspRef":"A1","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing settlementTime",
			input:   []byte(`{"transactionId":"A1","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects pspRef exceeding max length",
			input:   []byte(`{"transactionId":"A1","pspRef":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects pspRef outside identifierString charset",
			input:   []byte(`{"transactionId":"A1","pspRef":"###############################################################################################################################################################################################################################################################","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects statusInfo exceeding max length",
			input:   []byte(`{"transactionId":"A1","pspRef":"A1","status":"Settled","statusInfo":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects receiptId exceeding max length",
			input:   []byte(`{"transactionId":"A1","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","receiptUrl":"sample","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects receiptUrl exceeding max length",
			input:   []byte(`{"transactionId":"A1","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","vatNumber":"sample","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects vatNumber exceeding max length",
			input:   []byte(`{"transactionId":"A1","pspRef":"A1","status":"Settled","statusInfo":"sample","settlementAmount":1.0,"settlementTime":"2024-01-01T00:00:00Z","receiptId":"sample","receiptUrl":"sample","vatNumber":"xxxxxxxxxxxxxxxxxxxxx","vatCompany":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifySettlementRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifySettlementRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifySettlementResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"receiptUrl":"sample","receiptId":"sample"}`),
		},
		{
			name:    "rejects receiptUrl exceeding max length",
			input:   []byte(`{"receiptUrl":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","receiptId":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects receiptId exceeding max length",
			input:   []byte(`{"receiptUrl":"sample","receiptId":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifySettlementResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifySettlementResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyWebPaymentStartedRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":0,"timeout":1}`),
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1,"timeout":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NotifyWebPaymentStartedRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyWebPaymentStartedRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNotifyWebPaymentStartedResponseUnmarshalJSON(t *testing.T) {
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
			var s NotifyWebPaymentStartedResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NotifyWebPaymentStartedResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestOpenPeriodicEventStreamRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"constantStreamData":{"id":0,"variableMonitoringId":0,"params":{"interval":0,"values":0}}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s OpenPeriodicEventStreamRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s OpenPeriodicEventStreamRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestOpenPeriodicEventStreamResponseUnmarshalJSON(t *testing.T) {
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
			var s OpenPeriodicEventStreamResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s OpenPeriodicEventStreamResponse
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
			input: []byte(`{"location":"sample","retries":0,"checksum":"A1","requestId":0,"retryInterval":0}`),
		},
		{
			name:    "rejects missing location",
			input:   []byte(`{"retries":0,"checksum":"A1","requestId":0,"retryInterval":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing checksum",
			input:   []byte(`{"location":"sample","retries":0,"requestId":0,"retryInterval":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects location exceeding max length",
			input:   []byte(`{"location":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","retries":0,"checksum":"A1","requestId":0,"retryInterval":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects retries out of range",
			input:   []byte(`{"location":"sample","retries":-1,"checksum":"A1","requestId":0,"retryInterval":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects checksum exceeding max length",
			input:   []byte(`{"location":"sample","retries":0,"checksum":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","requestId":0,"retryInterval":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects checksum outside identifierString charset",
			input:   []byte(`{"location":"sample","retries":0,"checksum":"################################","requestId":0,"retryInterval":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects requestId out of range",
			input:   []byte(`{"location":"sample","retries":0,"checksum":"A1","requestId":-1,"retryInterval":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects retryInterval out of range",
			input:   []byte(`{"location":"sample","retries":0,"checksum":"A1","requestId":0,"retryInterval":-1}`),
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
			input: []byte(`{"status":"Idle","location":["sample"],"requestId":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"location":["sample"],"requestId":0,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects requestId out of range",
			input:   []byte(`{"status":"Idle","location":["sample"],"requestId":-1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestPullDynamicScheduleUpdateRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"chargingProfileId":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s PullDynamicScheduleUpdateRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s PullDynamicScheduleUpdateRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestPullDynamicScheduleUpdateResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","scheduleUpdate":{"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"scheduleUpdate":{"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s PullDynamicScheduleUpdateResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s PullDynamicScheduleUpdateResponse
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
			input: []byte(`{"requestId":1,"chargingLimitSource":"sample","tbc":true,"evseId":0,"chargingProfile":[{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}]}`),
		},
		{
			name:    "rejects missing chargingLimitSource",
			input:   []byte(`{"requestId":1,"tbc":true,"evseId":0,"chargingProfile":[{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingProfile",
			input:   []byte(`{"requestId":1,"chargingLimitSource":"sample","tbc":true,"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects chargingLimitSource exceeding max length",
			input:   []byte(`{"requestId":1,"chargingLimitSource":"xxxxxxxxxxxxxxxxxxxxx","tbc":true,"evseId":0,"chargingProfile":[{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"requestId":1,"chargingLimitSource":"sample","tbc":true,"evseId":-1,"chargingProfile":[{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestReportDERControlRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
		},
		{
			name:    "rejects fixedPFAbsorb exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects fixedPFInject exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects fixedVar exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects limitMaxDischarge exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}},{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects freqDroop exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}},{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects enterService exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}},{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects gradient exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}},{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects curve exceeding max items",
			input:   []byte(`{"requestId":1,"tbc":true,"fixedPFAbsorb":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedPFInject":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedPF":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"fixedVar":[{"id":"A1","isDefault":true,"isSuperseded":true,"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"limitMaxDischarge":[{"id":"A1","isDefault":true,"isSuperseded":true,"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}}],"freqDroop":[{"id":"A1","isDefault":true,"isSuperseded":true,"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0}}],"enterService":[{"id":"A1","enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0}}],"gradient":[{"id":"A1","gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}],"curve":[{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},{"id":"A1","curveType":"EnterService","isDefault":true,"isSuperseded":true,"curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReportDERControlRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReportDERControlRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReportDERControlResponseUnmarshalJSON(t *testing.T) {
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
			var s ReportDERControlResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReportDERControlResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRequestBatterySwapRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestId":1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s RequestBatterySwapRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RequestBatterySwapRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRequestBatterySwapResponseUnmarshalJSON(t *testing.T) {
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
			var s RequestBatterySwapResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RequestBatterySwapResponse
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
			input: []byte(`{"evseId":1,"remoteStartId":1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"chargingProfile":{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"evseId":1,"remoteStartId":1,"chargingProfile":{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":0,"remoteStartId":1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"chargingProfile":{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"reservationId":0,"reservationUpdateStatus":"Expired"}`),
		},
		{
			name:    "rejects missing reservationUpdateStatus",
			input:   []byte(`{"reservationId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects reservationId out of range",
			input:   []byte(`{"reservationId":-1,"reservationUpdateStatus":"Expired"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"id":0,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"sample","evseId":0,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
		},
		{
			name:    "rejects missing expiryDateTime",
			input:   []byte(`{"id":0,"connectorType":"sample","evseId":0,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"id":0,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"sample","evseId":0,"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects id out of range",
			input:   []byte(`{"id":-1,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"sample","evseId":0,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects connectorType exceeding max length",
			input:   []byte(`{"id":0,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"xxxxxxxxxxxxxxxxxxxxx","evseId":0,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"id":0,"expiryDateTime":"2024-01-01T00:00:00Z","connectorType":"sample","evseId":-1,"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"type":"Immediate","evseId":0}`),
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"type":"Immediate","evseId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			name:    "rejects missing timestamp",
			input:   []byte(`{"type":"sample","techInfo":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects type exceeding max length",
			input:   []byte(`{"type":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","timestamp":"2024-01-01T00:00:00Z","techInfo":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"versionNumber":1,"updateType":"Differential","localAuthorizationList":[{"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","language2":"sample","evseId":[0],"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}]}`),
		},
		{
			name:    "rejects missing updateType",
			input:   []byte(`{"versionNumber":1,"localAuthorizationList":[{"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","language2":"sample","evseId":[0],"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}]}`),
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
			input: []byte(`{"evseId":0,"chargingProfile":{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}}`),
		},
		{
			name:    "rejects missing chargingProfile",
			input:   []byte(`{"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1,"chargingProfile":{"id":1,"stackLevel":0,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","maxOfflineDuration":1,"invalidAfterOfflineDuration":true,"dynUpdateInterval":1,"dynUpdateTime":"2024-01-01T00:00:00Z","priceScheduleSignature":"sample","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"powerTolerance":1.0,"signatureId":0,"digestValue":"sample","useLocalTime":true,"randomizedDelay":0,"salesTariff":{"id":0,"salesTariffDescription":"sample","numEPriceLevels":0,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]},"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"numberPhases":2,"phaseToUse":2,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0,"preconditioningRequest":true,"evseSleep":true,"v2xBaseline":1.0,"operationMode":"Idle","v2xFreqWattCurve":[{"frequency":1.0,"power":1.0}],"v2xSignalWattCurve":[{"signal":1,"power":1.0}]}],"absolutePriceSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleID":0,"priceScheduleDescription":"sample","currency":"sam","language":"sample","priceAlgorithm":"sample","priceRuleStacks":[{"duration":1,"priceRule":[{"parkingFeePeriod":1,"carbonDioxideEmission":0,"renewableGenerationPercentage":50,"energyFee":{"exponent":1,"value":1},"parkingFee":{"exponent":1,"value":1},"powerRangeStart":{"exponent":1,"value":1}}]}],"taxRules":[{"taxRuleID":0,"taxRuleName":"sample","taxIncludedInPrice":true,"appliesToEnergyFee":true,"appliesToParkingFee":true,"appliesToOverstayFee":true,"appliesToMinimumMaximumCost":true,"taxRate":{"exponent":1,"value":1}}],"additionalSelectedServices":[{"serviceName":"sample","serviceFee":{"exponent":1,"value":1}}],"overstayRuleList":{"overstayTimeThreshold":1,"overstayPowerThreshold":{"exponent":1,"value":1},"overstayRule":[{"overstayRuleDescription":"sample","startTime":1,"overstayFeePeriod":1,"overstayFee":{"exponent":1,"value":1}}]},"minimumCost":{"exponent":1,"value":1},"maximumCost":{"exponent":1,"value":1}},"priceLevelSchedule":{"timeAnchor":"2024-01-01T00:00:00Z","priceScheduleId":0,"priceScheduleDescription":"sample","numberOfPriceLevels":0,"priceLevelScheduleEntries":[{"duration":1,"priceLevel":0}]},"limitAtSoC":{"soc":50,"limit":1.0}}]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestSetDefaultTariffRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":0,"tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
		},
		{
			name:    "rejects missing tariff",
			input:   []byte(`{"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1,"tariff":{"tariffId":"sample","currency":"sam","validFrom":"2024-01-01T00:00:00Z","description":[{"format":"ASCII","language":"sample","content":"sample"}],"energy":{"taxRates":[{"type":"sample","tax":1.0,"stack":0}],"prices":[{"priceKwh":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}]},"chargingTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"fixedFee":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"minCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"maxCost":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"prices":[{"priceMinute":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","minEnergy":1.0,"maxEnergy":1.0,"minCurrent":1.0,"maxCurrent":1.0,"minPower":1.0,"maxPower":1.0,"minTime":1,"maxTime":1,"minChargingTime":1,"maxChargingTime":1,"minIdleTime":1,"maxIdleTime":1}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationFixed":{"prices":[{"priceFixed":1.0,"conditions":{"startTimeOfDay":"sample","endTimeOfDay":"sample","dayOfWeek":["Monday"],"validFromDate":"sample","validToDate":"sample","evseKind":"AC","paymentBrand":"sample","paymentRecognition":"sample"}}],"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetDefaultTariffRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetDefaultTariffRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetDefaultTariffResponseUnmarshalJSON(t *testing.T) {
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
			var s SetDefaultTariffResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetDefaultTariffResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetDERControlRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"isDefault":true,"controlId":"A1","controlType":"EnterService","curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]},"fixedPFAbsorb":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedPFInject":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0},"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0},"gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}`),
		},
		{
			name:    "rejects missing controlId",
			input:   []byte(`{"isDefault":true,"controlType":"EnterService","curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]},"fixedPFAbsorb":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedPFInject":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0},"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0},"gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing controlType",
			input:   []byte(`{"isDefault":true,"controlId":"A1","curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]},"fixedPFAbsorb":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedPFInject":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0},"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0},"gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects controlId exceeding max length",
			input:   []byte(`{"isDefault":true,"controlId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","controlType":"EnterService","curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]},"fixedPFAbsorb":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedPFInject":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0},"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0},"gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects controlId outside identifierString charset",
			input:   []byte(`{"isDefault":true,"controlId":"####################################","controlType":"EnterService","curve":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]},"fixedPFAbsorb":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedPFInject":{"priority":0,"displacement":1.0,"excitation":true,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"fixedVar":{"priority":0,"setpoint":1.0,"unit":"Not_Applicable","startTime":"2024-01-01T00:00:00Z","duration":1.0},"limitMaxDischarge":{"priority":0,"pctMaxDischargePower":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"powerMonitoringMustTrip":{"priority":0,"yUnit":"Not_Applicable","responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0,"hysteresis":{"hysteresisHigh":1.0,"hysteresisLow":1.0,"hysteresisDelay":1.0,"hysteresisGradient":1.0},"voltageParams":{"hv10MinMeanValue":1.0,"hv10MinMeanTripDelay":1.0,"powerDuringCessation":"Active"},"reactivePowerParams":{"vRef":1.0,"autonomousVRefEnable":true,"autonomousVRefTimeConstant":1.0},"curveData":[{"x":1.0,"y":1.0}]}},"freqDroop":{"priority":0,"overFreq":1.0,"underFreq":1.0,"overDroop":1.0,"underDroop":1.0,"responseTime":1.0,"startTime":"2024-01-01T00:00:00Z","duration":1.0},"enterService":{"priority":0,"highVoltage":1.0,"lowVoltage":1.0,"highFreq":1.0,"lowFreq":1.0,"delay":1.0,"randomDelay":1.0,"rampRate":1.0},"gradient":{"priority":0,"gradient":1.0,"softGradient":1.0}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetDERControlRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetDERControlRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetDERControlResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","supersededIds":["A1"],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"supersededIds":["A1"],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects supersededIds exceeding max items",
			input:   []byte(`{"status":"Accepted","supersededIds":["A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1","A1"],"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetDERControlResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetDERControlResponse
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
			input: []byte(`{"message":{"id":0,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"messageExtra":[{"format":"ASCII","language":"sample","content":"sample"}]}}`),
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
			input: []byte(`{"severity":0}`),
		},
		{
			name:    "rejects severity out of range",
			input:   []byte(`{"severity":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"configurationSlot":1,"connectionData":{"ocppVersion":"OCPP12","ocppInterface":"Wired0","ocppTransport":"SOAP","messageTimeout":1,"ocppCsmsUrl":"sample","securityProfile":0,"identity":"sample","basicAuthPassword":"sample","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"PAP"}}}`),
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
			input: []byte(`{"setMonitoringData":[{"id":0,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":0,"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"periodicEventStream":{"interval":0,"values":0}}]}`),
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
			input: []byte(`{"setMonitoringResult":[{"id":0,"status":"Accepted","type":"UpperThreshold","severity":0,"component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
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
			input: []byte(`{"setVariableData":[{"attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"}}]}`),
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
			input: []byte(`{"setVariableResult":[{"attributeType":"Actual","attributeStatus":"Accepted","component":{"name":"A1","instance":"A1","evse":{"id":0,"connectorId":0}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}]}`),
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
			input: []byte(`{"csr":"sample","certificateType":"ChargingStationCertificate","requestId":1,"hashRootCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
		},
		{
			name:    "rejects missing csr",
			input:   []byte(`{"certificateType":"ChargingStationCertificate","requestId":1,"hashRootCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects csr exceeding max length",
			input:   []byte(`{"csr":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","certificateType":"ChargingStationCertificate","requestId":1,"hashRootCertificate":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}}`),
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
			input: []byte(`{"timestamp":"2024-01-01T00:00:00Z","connectorStatus":"Available","evseId":0,"connectorId":0}`),
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"connectorStatus":"Available","evseId":0,"connectorId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing connectorStatus",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","evseId":0,"connectorId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","connectorStatus":"Available","evseId":-1,"connectorId":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects connectorId out of range",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","connectorStatus":"Available","evseId":0,"connectorId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"AbnormalCondition","seqNo":0,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
		},
		{
			name:    "rejects missing eventType",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z","triggerReason":"AbnormalCondition","seqNo":0,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"eventType":"Ended","triggerReason":"AbnormalCondition","seqNo":0,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing triggerReason",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","seqNo":0,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing transactionInfo",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"AbnormalCondition","seqNo":0,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects seqNo out of range",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"AbnormalCondition","seqNo":-1,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects numberOfPhasesUsed out of range",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"AbnormalCondition","seqNo":0,"offline":true,"numberOfPhasesUsed":-1,"cableMaxCurrent":1,"reservationId":0,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects reservationId out of range",
			input:   []byte(`{"eventType":"Ended","timestamp":"2024-01-01T00:00:00Z","triggerReason":"AbnormalCondition","seqNo":0,"offline":true,"numberOfPhasesUsed":0,"cableMaxCurrent":1,"reservationId":-1,"preconditioningStatus":"Unknown","evseSleep":true,"meterValue":[{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"measurand":"Current.Export","context":"Interruption.Begin","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}],"idToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"evse":{"id":0,"connectorId":0},"transactionInfo":{"transactionId":"A1","chargingState":"EVConnected","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1,"operationMode":"Idle","tariffId":"sample","transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}},"costDetails":{"failureToCalculate":true,"failureReason":"sample","chargingPeriods":[{"tariffId":"sample","startPeriod":"2024-01-01T00:00:00Z","dimensions":[{"type":"Energy","volume":1.0}]}],"totalCost":{"currency":"sam","typeOfCost":"NormalCost","fixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"energy":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"chargingTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"idleTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"reservationTime":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]},"total":{"exclTax":1.0,"inclTax":1.0},"reservationFixed":{"exclTax":1.0,"inclTax":1.0,"taxRates":[{"type":"sample","tax":1.0,"stack":0}]}},"totalUsage":{"energy":1.0,"chargingTime":1,"idleTime":1,"reservationTime":1}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"totalCost":1.0,"chargingPriority":1,"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","language2":"sample","evseId":[0],"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"updatedPersonalMessage":{"format":"ASCII","language":"sample","content":"sample"},"updatedPersonalMessageExtra":[{"format":"ASCII","language":"sample","content":"sample"}],"transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}}`),
		},
		{
			name:    "rejects updatedPersonalMessageExtra exceeding max items",
			input:   []byte(`{"totalCost":1.0,"chargingPriority":1,"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","language2":"sample","evseId":[0],"groupIdToken":{"idToken":"A1","type":"sample","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"updatedPersonalMessage":{"format":"ASCII","language":"sample","content":"sample"},"updatedPersonalMessageExtra":[{"format":"ASCII","language":"sample","content":"sample"},{"format":"ASCII","language":"sample","content":"sample"},{"format":"ASCII","language":"sample","content":"sample"},{"format":"ASCII","language":"sample","content":"sample"},{"format":"ASCII","language":"sample","content":"sample"}],"transactionLimit":{"maxCost":1.0,"maxEnergy":1.0,"maxTime":1,"maxSoC":50}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"requestedMessage":"BootNotification","customTrigger":"sample","evse":{"id":0,"connectorId":0}}`),
		},
		{
			name:    "rejects missing requestedMessage",
			input:   []byte(`{"customTrigger":"sample","evse":{"id":0,"connectorId":0}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects customTrigger exceeding max length",
			input:   []byte(`{"requestedMessage":"BootNotification","customTrigger":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","evse":{"id":0,"connectorId":0}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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
			input: []byte(`{"evseId":0,"connectorId":0}`),
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"evseId":-1,"connectorId":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects connectorId out of range",
			input:   []byte(`{"evseId":0,"connectorId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestUpdateDynamicScheduleRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"chargingProfileId":1,"scheduleUpdate":{"limit":1.0,"limit_L2":1.0,"limit_L3":1.0,"dischargeLimit":-1.0,"dischargeLimit_L2":-1.0,"dischargeLimit_L3":-1.0,"setpoint":1.0,"setpoint_L2":1.0,"setpoint_L3":1.0,"setpointReactive":1.0,"setpointReactive_L2":1.0,"setpointReactive_L3":1.0}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UpdateDynamicScheduleRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UpdateDynamicScheduleRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUpdateDynamicScheduleResponseUnmarshalJSON(t *testing.T) {
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
			var s UpdateDynamicScheduleResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UpdateDynamicScheduleResponse
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
			input: []byte(`{"retries":0,"retryInterval":1,"requestId":1,"firmware":{"location":"sample","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}}`),
		},
		{
			name:    "rejects missing firmware",
			input:   []byte(`{"retries":0,"retryInterval":1,"requestId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects retries out of range",
			input:   []byte(`{"retries":-1,"retryInterval":1,"requestId":1,"firmware":{"location":"sample","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
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

func TestUsePriorityChargingRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1","activate":true}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{"activate":true}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","activate":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################","activate":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UsePriorityChargingRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UsePriorityChargingRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUsePriorityChargingResponseUnmarshalJSON(t *testing.T) {
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
			var s UsePriorityChargingResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UsePriorityChargingResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVatNumberValidationRequestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"vatNumber":"sample","evseId":0}`),
		},
		{
			name:    "rejects missing vatNumber",
			input:   []byte(`{"evseId":0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects vatNumber exceeding max length",
			input:   []byte(`{"vatNumber":"xxxxxxxxxxxxxxxxxxxxx","evseId":0}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"vatNumber":"sample","evseId":-1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VatNumberValidationRequest
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VatNumberValidationRequest
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVatNumberValidationResponseUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"vatNumber":"sample","evseId":0,"status":"Accepted","company":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing vatNumber",
			input:   []byte(`{"evseId":0,"status":"Accepted","company":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"vatNumber":"sample","evseId":0,"company":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects vatNumber exceeding max length",
			input:   []byte(`{"vatNumber":"xxxxxxxxxxxxxxxxxxxxx","evseId":0,"status":"Accepted","company":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects evseId out of range",
			input:   []byte(`{"vatNumber":"sample","evseId":-1,"status":"Accepted","company":{"name":"sample","address1":"sample","address2":"sample","city":"sample","postalCode":"sample","country":"sample"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VatNumberValidationResponse
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VatNumberValidationResponse
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}
