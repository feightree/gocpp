package v201

import (
	"encoding/json"
	"errors"
	"testing"

	ocpp "github.com/feightree/gocpp/ocpp"
)

func TestIdentifierString6TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString6Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString6Type
		data := []byte(`"AAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString6Type
		data := []byte(`"######"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestIdentifierString20TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString20Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString20Type
		data := []byte(`"AAAAAAAAAAAAAAAAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString20Type
		data := []byte(`"####################"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestIdentifierString32TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString32Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString32Type
		data := []byte(`"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString32Type
		data := []byte(`"################################"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestIdentifierString36TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString36Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString36Type
		data := []byte(`"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString36Type
		data := []byte(`"####################################"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestIdentifierString40TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString40Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString40Type
		data := []byte(`"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString40Type
		data := []byte(`"########################################"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestIdentifierString50TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString50Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString50Type
		data := []byte(`"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString50Type
		data := []byte(`"##################################################"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestIdentifierString128TypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts valid value", func(t *testing.T) {
		var s IdentifierString128Type
		data := []byte(`"AAA"`)

		if err := json.Unmarshal(data, &s); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}
	})

	t.Run("rejects value exceeding max length", func(t *testing.T) {
		var s IdentifierString128Type
		data := []byte(`"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects value outside charset", func(t *testing.T) {
		var s IdentifierString128Type
		data := []byte(`"################################################################################################################################"`)

		err := json.Unmarshal(data, &s)
		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})
}

func TestACChargingParametersTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"energyAmount":1,"evMinCurrent":1,"evMaxCurrent":1,"evMaxVoltage":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ACChargingParametersType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ACChargingParametersType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAdditionalInfoTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"additionalIdToken":"A1","type":"sample"}`),
		},
		{
			name:    "rejects missing additionalIdToken",
			input:   []byte(`{"type":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects additionalIdToken exceeding max length",
			input:   []byte(`{"additionalIdToken":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","type":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects additionalIdToken outside identifierString charset",
			input:   []byte(`{"additionalIdToken":"####################################","type":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"additionalIdToken":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects type exceeding max length",
			input:   []byte(`{"additionalIdToken":"A1","type":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s AdditionalInfoType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AdditionalInfoType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAPNTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
		},
		{
			name:    "rejects missing apn",
			input:   []byte(`{"apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects apn exceeding max length",
			input:   []byte(`{"apn":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects apnUserName exceeding max length",
			input:   []byte(`{"apn":"sample","apnUserName":"xxxxxxxxxxxxxxxxxxxxx","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects apnPassword exceeding max length",
			input:   []byte(`{"apn":"sample","apnUserName":"sample","apnPassword":"xxxxxxxxxxxxxxxxxxxxx","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects preferredNetwork exceeding max length",
			input:   []byte(`{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"AAAAAAA","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects preferredNetwork outside identifierString charset",
			input:   []byte(`{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"######","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing apnAuthentication",
			input:   []byte(`{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s APNType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s APNType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestAuthorizationDataUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}},"idToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}}`),
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"idTokenInfo":{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s AuthorizationData
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s AuthorizationData
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCertificateHashDataChainTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"certificateType":"V2GRootCertificate","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},"childCertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}]}`),
		},
		{
			name:    "rejects missing certificateType",
			input:   []byte(`{"certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},"childCertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing certificateHashData",
			input:   []byte(`{"certificateType":"V2GRootCertificate","childCertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects childCertificateHashData exceeding max items",
			input:   []byte(`{"certificateType":"V2GRootCertificate","certificateHashData":{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},"childCertificateHashData":[{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"},{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CertificateHashDataChainType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CertificateHashDataChainType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCertificateHashDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}`),
		},
		{
			name:    "rejects missing hashAlgorithm",
			input:   []byte(`{"issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing issuerNameHash",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerKeyHash":"sample","serialNumber":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects issuerNameHash exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","issuerKeyHash":"sample","serialNumber":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects issuerNameHash outside identifierString charset",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"################################################################################################################################","issuerKeyHash":"sample","serialNumber":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing issuerKeyHash",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","serialNumber":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects issuerKeyHash exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","serialNumber":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing serialNumber",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects serialNumber exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects serialNumber outside identifierString charset",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"########################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CertificateHashDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CertificateHashDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingLimitTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"chargingLimitSource":"EMS","isGridCritical":true}`),
		},
		{
			name:    "rejects missing chargingLimitSource",
			input:   []byte(`{"isGridCritical":true}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingLimitType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingLimitType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingNeedsTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"requestedEnergyTransfer":"DC","departureTime":"2024-01-01T00:00:00Z","acChargingParameters":{"energyAmount":1,"evMinCurrent":1,"evMaxCurrent":1,"evMaxVoltage":1},"dcChargingParameters":{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":50,"evEnergyCapacity":1,"fullSoC":50,"bulkSoC":50}}`),
		},
		{
			name:    "rejects missing requestedEnergyTransfer",
			input:   []byte(`{"departureTime":"2024-01-01T00:00:00Z","acChargingParameters":{"energyAmount":1,"evMinCurrent":1,"evMaxCurrent":1,"evMaxVoltage":1},"dcChargingParameters":{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":50,"evEnergyCapacity":1,"fullSoC":50,"bulkSoC":50}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingNeedsType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingNeedsType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingProfileCriterionTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":1,"chargingProfileId":[1],"chargingLimitSource":["EMS"]}`),
		},
		{
			name:    "rejects chargingLimitSource exceeding max items",
			input:   []byte(`{"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":1,"chargingProfileId":[1],"chargingLimitSource":["EMS","EMS","EMS","EMS","EMS"]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingProfileCriterionType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingProfileCriterionType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingProfileTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
		},
		{
			name:    "rejects missing chargingProfilePurpose",
			input:   []byte(`{"id":1,"stackLevel":1,"chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingProfileKind",
			input:   []byte(`{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"####################################","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing chargingSchedule",
			input:   []byte(`{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects chargingSchedule exceeding max items",
			input:   []byte(`{"id":1,"stackLevel":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","chargingProfileKind":"Absolute","recurrencyKind":"Daily","validFrom":"2024-01-01T00:00:00Z","validTo":"2024-01-01T00:00:00Z","transactionId":"A1","chargingSchedule":[{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}},{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}},{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}},{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingProfileType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingProfileType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingSchedulePeriodTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingSchedulePeriodType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingSchedulePeriodType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingScheduleTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}`),
		},
		{
			name:    "rejects missing chargingRateUnit",
			input:   []byte(`{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingSchedulePeriod",
			input:   []byte(`{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects chargingSchedulePeriod exceeding max items",
			input:   []byte(`{"id":1,"startSchedule":"2024-01-01T00:00:00Z","duration":1,"chargingRateUnit":"W","minChargingRate":1.0,"chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1},{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}],"salesTariff":{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingScheduleType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingScheduleType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestChargingStationTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"serialNumber":"sample","model":"sample","vendorName":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}`),
		},
		{
			name:    "rejects serialNumber exceeding max length",
			input:   []byte(`{"serialNumber":"xxxxxxxxxxxxxxxxxxxxxxxxxx","model":"sample","vendorName":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing model",
			input:   []byte(`{"serialNumber":"sample","vendorName":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects model exceeding max length",
			input:   []byte(`{"serialNumber":"sample","model":"xxxxxxxxxxxxxxxxxxxxx","vendorName":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing vendorName",
			input:   []byte(`{"serialNumber":"sample","model":"sample","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects vendorName exceeding max length",
			input:   []byte(`{"serialNumber":"sample","model":"sample","vendorName":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","firmwareVersion":"sample","modem":{"iccid":"A1","imsi":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects firmwareVersion exceeding max length",
			input:   []byte(`{"serialNumber":"sample","model":"sample","vendorName":"sample","firmwareVersion":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","modem":{"iccid":"A1","imsi":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ChargingStationType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ChargingStationType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearChargingProfileTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"chargingProfilePurpose":"ChargingStationExternalConstraints","stackLevel":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearChargingProfileType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearChargingProfileType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestClearMonitoringResultTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","id":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"id":1,"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ClearMonitoringResultType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ClearMonitoringResultType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestComponentTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}`),
		},
		{
			name:    "rejects missing name",
			input:   []byte(`{"instance":"A1","evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects name exceeding max length",
			input:   []byte(`{"name":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","instance":"A1","evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects name outside identifierString charset",
			input:   []byte(`{"name":"##################################################","instance":"A1","evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects instance exceeding max length",
			input:   []byte(`{"name":"A1","instance":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects instance outside identifierString charset",
			input:   []byte(`{"name":"A1","instance":"##################################################","evse":{"id":1,"connectorId":1}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ComponentType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ComponentType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestComponentVariableTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ComponentVariableType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ComponentVariableType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCompositeScheduleTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evseId":1,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}]}`),
		},
		{
			name:    "rejects missing scheduleStart",
			input:   []byte(`{"evseId":1,"duration":1,"chargingRateUnit":"W","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingRateUnit",
			input:   []byte(`{"evseId":1,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingSchedulePeriod":[{"startPeriod":1,"limit":1.0,"numberPhases":1,"phaseToUse":1}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing chargingSchedulePeriod",
			input:   []byte(`{"evseId":1,"duration":1,"scheduleStart":"2024-01-01T00:00:00Z","chargingRateUnit":"W"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CompositeScheduleType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CompositeScheduleType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestConsumptionCostTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}`),
		},
		{
			name:    "rejects missing cost",
			input:   []byte(`{"startValue":1.0}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects cost exceeding max items",
			input:   []byte(`{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1},{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1},{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1},{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ConsumptionCostType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ConsumptionCostType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestCostTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}`),
		},
		{
			name:    "rejects missing costKind",
			input:   []byte(`{"amount":1,"amountMultiplier":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s CostType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s CostType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestDCChargingParametersTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":50,"evEnergyCapacity":1,"fullSoC":50,"bulkSoC":50}`),
		},
		{
			name:    "rejects stateOfCharge out of range",
			input:   []byte(`{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":101,"evEnergyCapacity":1,"fullSoC":50,"bulkSoC":50}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects fullSoC out of range",
			input:   []byte(`{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":50,"evEnergyCapacity":1,"fullSoC":101,"bulkSoC":50}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects bulkSoC out of range",
			input:   []byte(`{"evMaxCurrent":1,"evMaxVoltage":1,"energyAmount":1,"evMaxPower":1,"stateOfCharge":50,"evEnergyCapacity":1,"fullSoC":50,"bulkSoC":101}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s DCChargingParametersType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s DCChargingParametersType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestEventDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"eventId":1,"trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing trigger",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing actualValue",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects actualValue exceeding max length",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects techCode exceeding max length",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects techInfo exceeding max length",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"####################################","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing eventNotificationType",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"eventId":1,"timestamp":"2024-01-01T00:00:00Z","trigger":"Alerting","cause":1,"actualValue":"sample","techCode":"sample","techInfo":"sample","cleared":true,"transactionId":"A1","variableMonitoringId":1,"eventNotificationType":"HardWiredNotification","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s EventDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s EventDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestEVSETypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"connectorId":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s EVSEType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s EVSEType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestFirmwareTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"location":"sample","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}`),
		},
		{
			name:    "rejects missing location",
			input:   []byte(`{"retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects location exceeding max length",
			input:   []byte(`{"location":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing retrieveDateTime",
			input:   []byte(`{"location":"sample","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects signingCertificate exceeding max length",
			input:   []byte(`{"location":"sample","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","signature":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects signature exceeding max length",
			input:   []byte(`{"location":"sample","retrieveDateTime":"2024-01-01T00:00:00Z","installDateTime":"2024-01-01T00:00:00Z","signingCertificate":"sample","signature":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s FirmwareType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s FirmwareType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetVariableDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"attributeType":"Actual","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"attributeType":"Actual","variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"attributeType":"Actual","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetVariableDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetVariableDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestGetVariableResultTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"attributeStatus":"Accepted","attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing attributeStatus",
			input:   []byte(`{"attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects attributeValue exceeding max length",
			input:   []byte(`{"attributeStatus":"Accepted","attributeType":"Actual","attributeValue":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"attributeStatus":"Accepted","attributeType":"Actual","attributeValue":"sample","variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"attributeStatus":"Accepted","attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s GetVariableResultType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s GetVariableResultType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestIdTokenInfoTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects language1 exceeding max length",
			input:   []byte(`{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"xxxxxxxxx","evseId":[1],"language2":"sample","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects language2 exceeding max length",
			input:   []byte(`{"status":"Accepted","cacheExpiryDateTime":"2024-01-01T00:00:00Z","chargingPriority":1,"language1":"sample","evseId":[1],"language2":"xxxxxxxxx","groupIdToken":{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]},"personalMessage":{"format":"ASCII","language":"sample","content":"sample"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s IdTokenInfoType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s IdTokenInfoType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestIdTokenTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"idToken":"A1","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}`),
		},
		{
			name:    "rejects missing idToken",
			input:   []byte(`{"type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects idToken exceeding max length",
			input:   []byte(`{"idToken":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects idToken outside identifierString charset",
			input:   []byte(`{"idToken":"####################################","type":"Central","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"idToken":"A1","additionalInfo":[{"additionalIdToken":"A1","type":"sample"}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s IdTokenType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s IdTokenType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestLogParametersTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"remoteLocation":"sample","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}`),
		},
		{
			name:    "rejects missing remoteLocation",
			input:   []byte(`{"oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects remoteLocation exceeding max length",
			input:   []byte(`{"remoteLocation":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","oldestTimestamp":"2024-01-01T00:00:00Z","latestTimestamp":"2024-01-01T00:00:00Z"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s LogParametersType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s LogParametersType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestMessageContentTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"format":"ASCII","language":"sample","content":"sample"}`),
		},
		{
			name:    "rejects missing format",
			input:   []byte(`{"language":"sample","content":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects language exceeding max length",
			input:   []byte(`{"format":"ASCII","language":"xxxxxxxxx","content":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing content",
			input:   []byte(`{"format":"ASCII","language":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects content exceeding max length",
			input:   []byte(`{"format":"ASCII","language":"sample","content":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s MessageContentType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s MessageContentType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestMessageInfoTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
		},
		{
			name:    "rejects missing priority",
			input:   []byte(`{"id":1,"state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"id":1,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"id":1,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"####################################","message":{"format":"ASCII","language":"sample","content":"sample"},"display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing message",
			input:   []byte(`{"id":1,"priority":"AlwaysFront","state":"Charging","startDateTime":"2024-01-01T00:00:00Z","endDateTime":"2024-01-01T00:00:00Z","transactionId":"A1","display":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s MessageInfoType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s MessageInfoType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestMeterValueTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"timestamp":"2024-01-01T00:00:00Z","sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}`),
		},
		{
			name:    "rejects missing timestamp",
			input:   []byte(`{"sampledValue":[{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing sampledValue",
			input:   []byte(`{"timestamp":"2024-01-01T00:00:00Z"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s MeterValueType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s MeterValueType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestModemTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"iccid":"A1","imsi":"A1"}`),
		},
		{
			name:    "rejects iccid exceeding max length",
			input:   []byte(`{"iccid":"AAAAAAAAAAAAAAAAAAAAA","imsi":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects iccid outside identifierString charset",
			input:   []byte(`{"iccid":"####################","imsi":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects imsi exceeding max length",
			input:   []byte(`{"iccid":"A1","imsi":"AAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects imsi outside identifierString charset",
			input:   []byte(`{"iccid":"A1","imsi":"####################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ModemType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ModemType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestMonitoringDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1}]}`),
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"variable":{"name":"A1","instance":"A1"},"variableMonitoring":[{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variableMonitoring":[{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1}]}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variableMonitoring",
			input:   []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s MonitoringDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s MonitoringDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestNetworkConnectionProfileTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"ocppVersion":"OCPP12","ocppTransport":"JSON","ocppCsmsUrl":"sample","messageTimeout":1,"securityProfile":1,"ocppInterface":"Wired0","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}`),
		},
		{
			name:    "rejects missing ocppVersion",
			input:   []byte(`{"ocppTransport":"JSON","ocppCsmsUrl":"sample","messageTimeout":1,"securityProfile":1,"ocppInterface":"Wired0","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing ocppTransport",
			input:   []byte(`{"ocppVersion":"OCPP12","ocppCsmsUrl":"sample","messageTimeout":1,"securityProfile":1,"ocppInterface":"Wired0","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing ocppCsmsUrl",
			input:   []byte(`{"ocppVersion":"OCPP12","ocppTransport":"JSON","messageTimeout":1,"securityProfile":1,"ocppInterface":"Wired0","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects ocppCsmsUrl exceeding max length",
			input:   []byte(`{"ocppVersion":"OCPP12","ocppTransport":"JSON","ocppCsmsUrl":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","messageTimeout":1,"securityProfile":1,"ocppInterface":"Wired0","vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing ocppInterface",
			input:   []byte(`{"ocppVersion":"OCPP12","ocppTransport":"JSON","ocppCsmsUrl":"sample","messageTimeout":1,"securityProfile":1,"vpn":{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"},"apn":{"apn":"sample","apnUserName":"sample","apnPassword":"sample","simPin":1,"preferredNetwork":"A1","useOnlyPreferredNetwork":true,"apnAuthentication":"CHAP"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s NetworkConnectionProfileType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s NetworkConnectionProfileType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestOCSPRequestDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}`),
		},
		{
			name:    "rejects missing hashAlgorithm",
			input:   []byte(`{"issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing issuerNameHash",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects issuerNameHash exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects issuerNameHash outside identifierString charset",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"################################################################################################################################","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing issuerKeyHash",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","serialNumber":"A1","responderURL":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects issuerKeyHash exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","serialNumber":"A1","responderURL":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing serialNumber",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","responderURL":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects serialNumber exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","responderURL":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects serialNumber outside identifierString charset",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"########################################","responderURL":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing responderURL",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects responderURL exceeding max length",
			input:   []byte(`{"hashAlgorithm":"SHA256","issuerNameHash":"A1","issuerKeyHash":"sample","serialNumber":"A1","responderURL":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s OCSPRequestDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s OCSPRequestDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestRelativeTimeIntervalTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"start":1,"duration":1}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s RelativeTimeIntervalType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s RelativeTimeIntervalType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestReportDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}`),
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variableAttribute",
			input:   []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects variableAttribute exceeding max items",
			input:   []byte(`{"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"variableAttribute":[{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true},{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true},{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true},{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true},{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}],"variableCharacteristics":{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s ReportDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s ReportDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSalesTariffEntryTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}`),
		},
		{
			name:    "rejects ePriceLevel out of range",
			input:   []byte(`{"ePriceLevel":-1,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects consumptionCost exceeding max items",
			input:   []byte(`{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]},{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]},{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]},{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SalesTariffEntryType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SalesTariffEntryType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSalesTariffTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}`),
		},
		{
			name:    "rejects salesTariffDescription exceeding max length",
			input:   []byte(`{"id":1,"salesTariffDescription":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing salesTariffEntry",
			input:   []byte(`{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects salesTariffEntry exceeding max items",
			input:   []byte(`{"id":1,"salesTariffDescription":"sample","numEPriceLevels":1,"salesTariffEntry":[{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]},{"ePriceLevel":0,"relativeTimeInterval":{"start":1,"duration":1},"consumptionCost":[{"startValue":1.0,"cost":[{"costKind":"CarbonDioxideEmission","amount":1,"amountMultiplier":1}]}]}]}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SalesTariffType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SalesTariffType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSampledValueTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"value":1.0,"context":"Interruption.Begin","measurand":"Current.Export","phase":"L1","location":"Body","signedMeterValue":{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"},"unitOfMeasure":{"unit":"sample","multiplier":1}}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SampledValueType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SampledValueType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetMonitoringDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"id":1,"transaction":true,"value":1.0,"severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1,"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetMonitoringDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetMonitoringDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetMonitoringResultTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"status":"Accepted","type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing status",
			input:   []byte(`{"id":1,"type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"id":1,"status":"Accepted","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"id":1,"status":"Accepted","type":"UpperThreshold","severity":1,"variable":{"name":"A1","instance":"A1"},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"id":1,"status":"Accepted","type":"UpperThreshold","severity":1,"component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"statusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetMonitoringResultType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetMonitoringResultType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetVariableDataTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
		},
		{
			name:    "rejects missing attributeValue",
			input:   []byte(`{"attributeType":"Actual","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects attributeValue exceeding max length",
			input:   []byte(`{"attributeType":"Actual","attributeValue":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"attributeType":"Actual","attributeValue":"sample","variable":{"name":"A1","instance":"A1"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"attributeType":"Actual","attributeValue":"sample","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetVariableDataType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetVariableDataType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSetVariableResultTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"attributeType":"Actual","attributeStatus":"Accepted","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
		},
		{
			name:    "rejects missing attributeStatus",
			input:   []byte(`{"attributeType":"Actual","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing component",
			input:   []byte(`{"attributeType":"Actual","attributeStatus":"Accepted","variable":{"name":"A1","instance":"A1"},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects missing variable",
			input:   []byte(`{"attributeType":"Actual","attributeStatus":"Accepted","component":{"name":"A1","instance":"A1","evse":{"id":1,"connectorId":1}},"attributeStatusInfo":{"reasonCode":"sample","additionalInfo":"sample"}}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SetVariableResultType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SetVariableResultType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestSignedMeterValueTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"}`),
		},
		{
			name:    "rejects missing signedMeterData",
			input:   []byte(`{"signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects signedMeterData exceeding max length",
			input:   []byte(`{"signedMeterData":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","signingMethod":"sample","encodingMethod":"sample","publicKey":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing signingMethod",
			input:   []byte(`{"signedMeterData":"sample","encodingMethod":"sample","publicKey":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects signingMethod exceeding max length",
			input:   []byte(`{"signedMeterData":"sample","signingMethod":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","encodingMethod":"sample","publicKey":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing encodingMethod",
			input:   []byte(`{"signedMeterData":"sample","signingMethod":"sample","publicKey":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects encodingMethod exceeding max length",
			input:   []byte(`{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","publicKey":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing publicKey",
			input:   []byte(`{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects publicKey exceeding max length",
			input:   []byte(`{"signedMeterData":"sample","signingMethod":"sample","encodingMethod":"sample","publicKey":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SignedMeterValueType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s SignedMeterValueType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestStatusInfoTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"reasonCode":"sample","additionalInfo":"sample"}`),
		},
		{
			name:    "rejects missing reasonCode",
			input:   []byte(`{"additionalInfo":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects reasonCode exceeding max length",
			input:   []byte(`{"reasonCode":"xxxxxxxxxxxxxxxxxxxxx","additionalInfo":"sample"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects additionalInfo exceeding max length",
			input:   []byte(`{"reasonCode":"sample","additionalInfo":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s StatusInfoType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s StatusInfoType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestTransactionTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"transactionId":"A1","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1}`),
		},
		{
			name:    "rejects missing transactionId",
			input:   []byte(`{"chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects transactionId exceeding max length",
			input:   []byte(`{"transactionId":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects transactionId outside identifierString charset",
			input:   []byte(`{"transactionId":"####################################","chargingState":"Charging","timeSpentCharging":1,"stoppedReason":"DeAuthorized","remoteStartId":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s TransactionType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s TransactionType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestUnitOfMeasureTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"unit":"sample","multiplier":1}`),
		},
		{
			name:    "rejects unit exceeding max length",
			input:   []byte(`{"unit":"xxxxxxxxxxxxxxxxxxxxx","multiplier":1}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s UnitOfMeasureType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s UnitOfMeasureType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVariableAttributeTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"type":"Actual","value":"sample","mutability":"ReadOnly","persistent":true,"constant":true}`),
		},
		{
			name:    "rejects value exceeding max length",
			input:   []byte(`{"type":"Actual","value":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","mutability":"ReadOnly","persistent":true,"constant":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VariableAttributeType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VariableAttributeType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVariableCharacteristicsTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}`),
		},
		{
			name:    "rejects unit exceeding max length",
			input:   []byte(`{"unit":"xxxxxxxxxxxxxxxxx","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing dataType",
			input:   []byte(`{"unit":"sample","minLimit":1.0,"maxLimit":1.0,"valuesList":"sample","supportsMonitoring":true}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects valuesList exceeding max length",
			input:   []byte(`{"unit":"sample","dataType":"string","minLimit":1.0,"maxLimit":1.0,"valuesList":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","supportsMonitoring":true}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VariableCharacteristicsType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VariableCharacteristicsType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVariableMonitoringTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"id":1,"transaction":true,"value":1.0,"type":"UpperThreshold","severity":1}`),
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"id":1,"transaction":true,"value":1.0,"severity":1}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VariableMonitoringType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VariableMonitoringType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVariableTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"name":"A1","instance":"A1"}`),
		},
		{
			name:    "rejects missing name",
			input:   []byte(`{"instance":"A1"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects name exceeding max length",
			input:   []byte(`{"name":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA","instance":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects name outside identifierString charset",
			input:   []byte(`{"name":"##################################################","instance":"A1"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects instance exceeding max length",
			input:   []byte(`{"name":"A1","instance":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects instance outside identifierString charset",
			input:   []byte(`{"name":"A1","instance":"##################################################"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VariableType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VariableType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}

func TestVPNTypeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr error
	}{
		{
			name:  "accepts valid input",
			input: []byte(`{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"}`),
		},
		{
			name:    "rejects missing server",
			input:   []byte(`{"user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects server exceeding max length",
			input:   []byte(`{"server":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","user":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing user",
			input:   []byte(`{"server":"sample","group":"sample","password":"sample","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects user exceeding max length",
			input:   []byte(`{"server":"sample","user":"xxxxxxxxxxxxxxxxxxxxx","group":"sample","password":"sample","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects group exceeding max length",
			input:   []byte(`{"server":"sample","user":"sample","group":"xxxxxxxxxxxxxxxxxxxxx","password":"sample","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing password",
			input:   []byte(`{"server":"sample","user":"sample","group":"sample","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects password exceeding max length",
			input:   []byte(`{"server":"sample","user":"sample","group":"sample","password":"xxxxxxxxxxxxxxxxxxxxx","key":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing key",
			input:   []byte(`{"server":"sample","user":"sample","group":"sample","password":"sample","type":"IKEv2"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
		{
			name:    "rejects key exceeding max length",
			input:   []byte(`{"server":"sample","user":"sample","group":"sample","password":"sample","key":"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx","type":"IKEv2"}`),
			wantErr: ocpp.ErrPropertyConstraintViolation,
		},
		{
			name:    "rejects missing type",
			input:   []byte(`{"server":"sample","user":"sample","group":"sample","password":"sample","key":"sample"}`),
			wantErr: ocpp.ErrOccurenceConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s VPNType
			err := json.Unmarshal(tt.input, &s)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, tt.wantErr)
			}
		})
	}

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var s VPNType
		data := []byte(`1`)

		err := json.Unmarshal(data, &s)

		var typeErr *json.UnmarshalTypeError
		if !errors.As(err, &typeErr) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.As match for %v", err, typeErr)
		}
	})
}
