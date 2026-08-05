package v21

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	ocpp "github.com/feightree/gocpp/ocpp"
)

func TestAPNAuthenticationEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str APNAuthenticationEnumType
		data := fmt.Appendf(nil, `"%s"`, APNAuthenticationEnumTypePAP)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != APNAuthenticationEnumTypePAP {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, APNAuthenticationEnumTypePAP)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str APNAuthenticationEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str APNAuthenticationEnumType
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

func TestAttributeEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str AttributeEnumType
		data := fmt.Appendf(nil, `"%s"`, AttributeEnumTypeActual)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != AttributeEnumTypeActual {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, AttributeEnumTypeActual)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str AttributeEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str AttributeEnumType
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

func TestAuthorizationStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str AuthorizationStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, AuthorizationStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != AuthorizationStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, AuthorizationStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str AuthorizationStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str AuthorizationStatusEnumType
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

func TestAuthorizeCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str AuthorizeCertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, AuthorizeCertificateStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != AuthorizeCertificateStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, AuthorizeCertificateStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str AuthorizeCertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str AuthorizeCertificateStatusEnumType
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

func TestBatterySwapEventEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str BatterySwapEventEnumType
		data := fmt.Appendf(nil, `"%s"`, BatterySwapEventEnumTypeBatteryIn)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != BatterySwapEventEnumTypeBatteryIn {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, BatterySwapEventEnumTypeBatteryIn)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str BatterySwapEventEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str BatterySwapEventEnumType
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

func TestBootReasonEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str BootReasonEnumType
		data := fmt.Appendf(nil, `"%s"`, BootReasonEnumTypeApplicationReset)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != BootReasonEnumTypeApplicationReset {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, BootReasonEnumTypeApplicationReset)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str BootReasonEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str BootReasonEnumType
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

func TestCancelReservationStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CancelReservationStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, CancelReservationStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CancelReservationStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CancelReservationStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CancelReservationStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CancelReservationStatusEnumType
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

func TestCertificateActionEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CertificateActionEnumType
		data := fmt.Appendf(nil, `"%s"`, CertificateActionEnumTypeInstall)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CertificateActionEnumTypeInstall {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CertificateActionEnumTypeInstall)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CertificateActionEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CertificateActionEnumType
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

func TestCertificateSignedStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CertificateSignedStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, CertificateSignedStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CertificateSignedStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CertificateSignedStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CertificateSignedStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CertificateSignedStatusEnumType
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

func TestCertificateSigningUseEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CertificateSigningUseEnumType
		data := fmt.Appendf(nil, `"%s"`, CertificateSigningUseEnumTypeChargingStationCertificate)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CertificateSigningUseEnumTypeChargingStationCertificate {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CertificateSigningUseEnumTypeChargingStationCertificate)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CertificateSigningUseEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CertificateSigningUseEnumType
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

func TestCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, CertificateStatusEnumTypeGood)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CertificateStatusEnumTypeGood {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CertificateStatusEnumTypeGood)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CertificateStatusEnumType
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

func TestCertificateStatusSourceEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CertificateStatusSourceEnumType
		data := fmt.Appendf(nil, `"%s"`, CertificateStatusSourceEnumTypeCRL)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CertificateStatusSourceEnumTypeCRL {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CertificateStatusSourceEnumTypeCRL)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CertificateStatusSourceEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CertificateStatusSourceEnumType
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

func TestChangeAvailabilityStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChangeAvailabilityStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ChangeAvailabilityStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChangeAvailabilityStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChangeAvailabilityStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChangeAvailabilityStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChangeAvailabilityStatusEnumType
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

func TestChargingProfileKindEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingProfileKindEnumType
		data := fmt.Appendf(nil, `"%s"`, ChargingProfileKindEnumTypeAbsolute)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingProfileKindEnumTypeAbsolute {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingProfileKindEnumTypeAbsolute)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingProfileKindEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingProfileKindEnumType
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

func TestChargingProfilePurposeEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingProfilePurposeEnumType
		data := fmt.Appendf(nil, `"%s"`, ChargingProfilePurposeEnumTypeChargingStationExternalConstraints)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingProfilePurposeEnumTypeChargingStationExternalConstraints {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingProfilePurposeEnumTypeChargingStationExternalConstraints)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingProfilePurposeEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingProfilePurposeEnumType
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

func TestChargingProfileStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingProfileStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ChargingProfileStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingProfileStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingProfileStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingProfileStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingProfileStatusEnumType
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

func TestChargingRateUnitEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingRateUnitEnumType
		data := fmt.Appendf(nil, `"%s"`, ChargingRateUnitEnumTypeW)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingRateUnitEnumTypeW {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingRateUnitEnumTypeW)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingRateUnitEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingRateUnitEnumType
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

func TestChargingStateEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ChargingStateEnumType
		data := fmt.Appendf(nil, `"%s"`, ChargingStateEnumTypeEVConnected)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ChargingStateEnumTypeEVConnected {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ChargingStateEnumTypeEVConnected)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ChargingStateEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ChargingStateEnumType
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

func TestClearCacheStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ClearCacheStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ClearCacheStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ClearCacheStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ClearCacheStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ClearCacheStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ClearCacheStatusEnumType
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

func TestClearChargingProfileStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ClearChargingProfileStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ClearChargingProfileStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ClearChargingProfileStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ClearChargingProfileStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ClearChargingProfileStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ClearChargingProfileStatusEnumType
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

func TestClearMessageStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ClearMessageStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ClearMessageStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ClearMessageStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ClearMessageStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ClearMessageStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ClearMessageStatusEnumType
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

func TestClearMonitoringStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ClearMonitoringStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ClearMonitoringStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ClearMonitoringStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ClearMonitoringStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ClearMonitoringStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ClearMonitoringStatusEnumType
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

func TestComponentCriterionEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ComponentCriterionEnumType
		data := fmt.Appendf(nil, `"%s"`, ComponentCriterionEnumTypeActive)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ComponentCriterionEnumTypeActive {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ComponentCriterionEnumTypeActive)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ComponentCriterionEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ComponentCriterionEnumType
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

func TestConnectorStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ConnectorStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ConnectorStatusEnumTypeAvailable)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ConnectorStatusEnumTypeAvailable {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ConnectorStatusEnumTypeAvailable)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ConnectorStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ConnectorStatusEnumType
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

func TestControlModeEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ControlModeEnumType
		data := fmt.Appendf(nil, `"%s"`, ControlModeEnumTypeScheduledControl)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ControlModeEnumTypeScheduledControl {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ControlModeEnumTypeScheduledControl)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ControlModeEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ControlModeEnumType
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

func TestCostDimensionEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CostDimensionEnumType
		data := fmt.Appendf(nil, `"%s"`, CostDimensionEnumTypeEnergy)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CostDimensionEnumTypeEnergy {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CostDimensionEnumTypeEnergy)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CostDimensionEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CostDimensionEnumType
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

func TestCostKindEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CostKindEnumType
		data := fmt.Appendf(nil, `"%s"`, CostKindEnumTypeCarbonDioxideEmission)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CostKindEnumTypeCarbonDioxideEmission {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CostKindEnumTypeCarbonDioxideEmission)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CostKindEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CostKindEnumType
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

func TestCustomerInformationStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str CustomerInformationStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, CustomerInformationStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != CustomerInformationStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, CustomerInformationStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str CustomerInformationStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str CustomerInformationStatusEnumType
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

func TestDataEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DataEnumType
		data := fmt.Appendf(nil, `"%s"`, DataEnumTypeString)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DataEnumTypeString {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DataEnumTypeString)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DataEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DataEnumType
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

func TestDataTransferStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DataTransferStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, DataTransferStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DataTransferStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DataTransferStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DataTransferStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DataTransferStatusEnumType
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

func TestDayOfWeekEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DayOfWeekEnumType
		data := fmt.Appendf(nil, `"%s"`, DayOfWeekEnumTypeMonday)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DayOfWeekEnumTypeMonday {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DayOfWeekEnumTypeMonday)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DayOfWeekEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DayOfWeekEnumType
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

func TestDeleteCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DeleteCertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, DeleteCertificateStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DeleteCertificateStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DeleteCertificateStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DeleteCertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DeleteCertificateStatusEnumType
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

func TestDERControlEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DERControlEnumType
		data := fmt.Appendf(nil, `"%s"`, DERControlEnumTypeEnterService)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DERControlEnumTypeEnterService {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DERControlEnumTypeEnterService)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DERControlEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DERControlEnumType
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

func TestDERControlStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DERControlStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, DERControlStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DERControlStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DERControlStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DERControlStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DERControlStatusEnumType
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

func TestDERUnitEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DERUnitEnumType
		data := fmt.Appendf(nil, `"%s"`, DERUnitEnumTypeNotApplicable)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DERUnitEnumTypeNotApplicable {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DERUnitEnumTypeNotApplicable)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DERUnitEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DERUnitEnumType
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

func TestDisplayMessageStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str DisplayMessageStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, DisplayMessageStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != DisplayMessageStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, DisplayMessageStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str DisplayMessageStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str DisplayMessageStatusEnumType
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

func TestEnergyTransferModeEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str EnergyTransferModeEnumType
		data := fmt.Appendf(nil, `"%s"`, EnergyTransferModeEnumTypeACSinglePhase)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != EnergyTransferModeEnumTypeACSinglePhase {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, EnergyTransferModeEnumTypeACSinglePhase)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str EnergyTransferModeEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str EnergyTransferModeEnumType
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

func TestEventNotificationEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str EventNotificationEnumType
		data := fmt.Appendf(nil, `"%s"`, EventNotificationEnumTypeHardWiredNotification)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != EventNotificationEnumTypeHardWiredNotification {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, EventNotificationEnumTypeHardWiredNotification)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str EventNotificationEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str EventNotificationEnumType
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

func TestEventTriggerEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str EventTriggerEnumType
		data := fmt.Appendf(nil, `"%s"`, EventTriggerEnumTypeAlerting)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != EventTriggerEnumTypeAlerting {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, EventTriggerEnumTypeAlerting)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str EventTriggerEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str EventTriggerEnumType
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

func TestEVSEKindEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str EVSEKindEnumType
		data := fmt.Appendf(nil, `"%s"`, EVSEKindEnumTypeAC)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != EVSEKindEnumTypeAC {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, EVSEKindEnumTypeAC)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str EVSEKindEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str EVSEKindEnumType
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

func TestFirmwareStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str FirmwareStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, FirmwareStatusEnumTypeDownloaded)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != FirmwareStatusEnumTypeDownloaded {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, FirmwareStatusEnumTypeDownloaded)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str FirmwareStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str FirmwareStatusEnumType
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

func TestGenericDeviceModelStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GenericDeviceModelStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GenericDeviceModelStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GenericDeviceModelStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GenericDeviceModelStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GenericDeviceModelStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GenericDeviceModelStatusEnumType
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

func TestGenericStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GenericStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GenericStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GenericStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GenericStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GenericStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GenericStatusEnumType
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

func TestGetCertificateIdUseEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetCertificateIdUseEnumType
		data := fmt.Appendf(nil, `"%s"`, GetCertificateIdUseEnumTypeV2GRootCertificate)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetCertificateIdUseEnumTypeV2GRootCertificate {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetCertificateIdUseEnumTypeV2GRootCertificate)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetCertificateIdUseEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetCertificateIdUseEnumType
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

func TestGetCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetCertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GetCertificateStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetCertificateStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetCertificateStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetCertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetCertificateStatusEnumType
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

func TestGetChargingProfileStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetChargingProfileStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GetChargingProfileStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetChargingProfileStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetChargingProfileStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetChargingProfileStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetChargingProfileStatusEnumType
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

func TestGetDisplayMessagesStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetDisplayMessagesStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GetDisplayMessagesStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetDisplayMessagesStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetDisplayMessagesStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetDisplayMessagesStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetDisplayMessagesStatusEnumType
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

func TestGetInstalledCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetInstalledCertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GetInstalledCertificateStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetInstalledCertificateStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetInstalledCertificateStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetInstalledCertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetInstalledCertificateStatusEnumType
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

func TestGetVariableStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GetVariableStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, GetVariableStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GetVariableStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GetVariableStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GetVariableStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GetVariableStatusEnumType
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

func TestGridEventFaultEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str GridEventFaultEnumType
		data := fmt.Appendf(nil, `"%s"`, GridEventFaultEnumTypeCurrentImbalance)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != GridEventFaultEnumTypeCurrentImbalance {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, GridEventFaultEnumTypeCurrentImbalance)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str GridEventFaultEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str GridEventFaultEnumType
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

func TestHashAlgorithmEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str HashAlgorithmEnumType
		data := fmt.Appendf(nil, `"%s"`, HashAlgorithmEnumTypeSHA256)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != HashAlgorithmEnumTypeSHA256 {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, HashAlgorithmEnumTypeSHA256)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str HashAlgorithmEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str HashAlgorithmEnumType
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

func TestInstallCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str InstallCertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, InstallCertificateStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != InstallCertificateStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, InstallCertificateStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str InstallCertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str InstallCertificateStatusEnumType
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

func TestInstallCertificateUseEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str InstallCertificateUseEnumType
		data := fmt.Appendf(nil, `"%s"`, InstallCertificateUseEnumTypeV2GRootCertificate)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != InstallCertificateUseEnumTypeV2GRootCertificate {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, InstallCertificateUseEnumTypeV2GRootCertificate)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str InstallCertificateUseEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str InstallCertificateUseEnumType
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

func TestIslandingDetectionEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str IslandingDetectionEnumType
		data := fmt.Appendf(nil, `"%s"`, IslandingDetectionEnumTypeNoAntiIslandingSupport)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != IslandingDetectionEnumTypeNoAntiIslandingSupport {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, IslandingDetectionEnumTypeNoAntiIslandingSupport)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str IslandingDetectionEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str IslandingDetectionEnumType
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

func TestIso15118EVCertificateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str Iso15118EVCertificateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, Iso15118EVCertificateStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != Iso15118EVCertificateStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, Iso15118EVCertificateStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str Iso15118EVCertificateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str Iso15118EVCertificateStatusEnumType
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

func TestLocationEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str LocationEnumType
		data := fmt.Appendf(nil, `"%s"`, LocationEnumTypeBody)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != LocationEnumTypeBody {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, LocationEnumTypeBody)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str LocationEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str LocationEnumType
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

func TestLogEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str LogEnumType
		data := fmt.Appendf(nil, `"%s"`, LogEnumTypeDiagnosticsLog)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != LogEnumTypeDiagnosticsLog {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, LogEnumTypeDiagnosticsLog)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str LogEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str LogEnumType
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

func TestLogStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str LogStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, LogStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != LogStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, LogStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str LogStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str LogStatusEnumType
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

func TestMeasurandEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MeasurandEnumType
		data := fmt.Appendf(nil, `"%s"`, MeasurandEnumTypeCurrentExport)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MeasurandEnumTypeCurrentExport {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MeasurandEnumTypeCurrentExport)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MeasurandEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MeasurandEnumType
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

func TestMessageFormatEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MessageFormatEnumType
		data := fmt.Appendf(nil, `"%s"`, MessageFormatEnumTypeASCII)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MessageFormatEnumTypeASCII {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MessageFormatEnumTypeASCII)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MessageFormatEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MessageFormatEnumType
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

func TestMessagePriorityEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MessagePriorityEnumType
		data := fmt.Appendf(nil, `"%s"`, MessagePriorityEnumTypeAlwaysFront)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MessagePriorityEnumTypeAlwaysFront {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MessagePriorityEnumTypeAlwaysFront)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MessagePriorityEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MessagePriorityEnumType
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

func TestMessageStateEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MessageStateEnumType
		data := fmt.Appendf(nil, `"%s"`, MessageStateEnumTypeCharging)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MessageStateEnumTypeCharging {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MessageStateEnumTypeCharging)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MessageStateEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MessageStateEnumType
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

func TestMessageTriggerEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MessageTriggerEnumType
		data := fmt.Appendf(nil, `"%s"`, MessageTriggerEnumTypeBootNotification)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MessageTriggerEnumTypeBootNotification {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MessageTriggerEnumTypeBootNotification)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MessageTriggerEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MessageTriggerEnumType
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

func TestMobilityNeedsModeEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MobilityNeedsModeEnumType
		data := fmt.Appendf(nil, `"%s"`, MobilityNeedsModeEnumTypeEVCC)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MobilityNeedsModeEnumTypeEVCC {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MobilityNeedsModeEnumTypeEVCC)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MobilityNeedsModeEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MobilityNeedsModeEnumType
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

func TestMonitorEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MonitorEnumType
		data := fmt.Appendf(nil, `"%s"`, MonitorEnumTypeUpperThreshold)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MonitorEnumTypeUpperThreshold {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MonitorEnumTypeUpperThreshold)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MonitorEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MonitorEnumType
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

func TestMonitoringBaseEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MonitoringBaseEnumType
		data := fmt.Appendf(nil, `"%s"`, MonitoringBaseEnumTypeAll)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MonitoringBaseEnumTypeAll {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MonitoringBaseEnumTypeAll)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MonitoringBaseEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MonitoringBaseEnumType
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

func TestMonitoringCriterionEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MonitoringCriterionEnumType
		data := fmt.Appendf(nil, `"%s"`, MonitoringCriterionEnumTypeThresholdMonitoring)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MonitoringCriterionEnumTypeThresholdMonitoring {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MonitoringCriterionEnumTypeThresholdMonitoring)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MonitoringCriterionEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MonitoringCriterionEnumType
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

func TestMutabilityEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str MutabilityEnumType
		data := fmt.Appendf(nil, `"%s"`, MutabilityEnumTypeReadOnly)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != MutabilityEnumTypeReadOnly {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, MutabilityEnumTypeReadOnly)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str MutabilityEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str MutabilityEnumType
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

func TestNotifyAllowedEnergyTransferStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str NotifyAllowedEnergyTransferStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, NotifyAllowedEnergyTransferStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != NotifyAllowedEnergyTransferStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, NotifyAllowedEnergyTransferStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str NotifyAllowedEnergyTransferStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str NotifyAllowedEnergyTransferStatusEnumType
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

func TestNotifyEVChargingNeedsStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str NotifyEVChargingNeedsStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, NotifyEVChargingNeedsStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != NotifyEVChargingNeedsStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, NotifyEVChargingNeedsStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str NotifyEVChargingNeedsStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str NotifyEVChargingNeedsStatusEnumType
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

func TestOCPPInterfaceEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str OCPPInterfaceEnumType
		data := fmt.Appendf(nil, `"%s"`, OCPPInterfaceEnumTypeWired0)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != OCPPInterfaceEnumTypeWired0 {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, OCPPInterfaceEnumTypeWired0)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str OCPPInterfaceEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str OCPPInterfaceEnumType
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

func TestOCPPTransportEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str OCPPTransportEnumType
		data := fmt.Appendf(nil, `"%s"`, OCPPTransportEnumTypeSOAP)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != OCPPTransportEnumTypeSOAP {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, OCPPTransportEnumTypeSOAP)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str OCPPTransportEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str OCPPTransportEnumType
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

func TestOCPPVersionEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str OCPPVersionEnumType
		data := fmt.Appendf(nil, `"%s"`, OCPPVersionEnumTypeOCPP12)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != OCPPVersionEnumTypeOCPP12 {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, OCPPVersionEnumTypeOCPP12)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str OCPPVersionEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str OCPPVersionEnumType
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

func TestOperationalStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str OperationalStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, OperationalStatusEnumTypeInoperative)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != OperationalStatusEnumTypeInoperative {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, OperationalStatusEnumTypeInoperative)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str OperationalStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str OperationalStatusEnumType
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

func TestOperationModeEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str OperationModeEnumType
		data := fmt.Appendf(nil, `"%s"`, OperationModeEnumTypeIdle)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != OperationModeEnumTypeIdle {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, OperationModeEnumTypeIdle)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str OperationModeEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str OperationModeEnumType
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

func TestPaymentStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str PaymentStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, PaymentStatusEnumTypeSettled)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PaymentStatusEnumTypeSettled {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PaymentStatusEnumTypeSettled)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str PaymentStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str PaymentStatusEnumType
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

func TestPhaseEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str PhaseEnumType
		data := fmt.Appendf(nil, `"%s"`, PhaseEnumTypeL1)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PhaseEnumTypeL1 {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PhaseEnumTypeL1)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str PhaseEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str PhaseEnumType
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

func TestPowerDuringCessationEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str PowerDuringCessationEnumType
		data := fmt.Appendf(nil, `"%s"`, PowerDuringCessationEnumTypeActive)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PowerDuringCessationEnumTypeActive {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PowerDuringCessationEnumTypeActive)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str PowerDuringCessationEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str PowerDuringCessationEnumType
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

func TestPreconditioningStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str PreconditioningStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, PreconditioningStatusEnumTypeUnknown)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PreconditioningStatusEnumTypeUnknown {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PreconditioningStatusEnumTypeUnknown)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str PreconditioningStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str PreconditioningStatusEnumType
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

func TestPriorityChargingStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str PriorityChargingStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, PriorityChargingStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PriorityChargingStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PriorityChargingStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str PriorityChargingStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str PriorityChargingStatusEnumType
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

func TestPublishFirmwareStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str PublishFirmwareStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, PublishFirmwareStatusEnumTypeIdle)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != PublishFirmwareStatusEnumTypeIdle {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, PublishFirmwareStatusEnumTypeIdle)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str PublishFirmwareStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str PublishFirmwareStatusEnumType
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

func TestReadingContextEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReadingContextEnumType
		data := fmt.Appendf(nil, `"%s"`, ReadingContextEnumTypeInterruptionBegin)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReadingContextEnumTypeInterruptionBegin {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReadingContextEnumTypeInterruptionBegin)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReadingContextEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReadingContextEnumType
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

func TestReasonEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReasonEnumType
		data := fmt.Appendf(nil, `"%s"`, ReasonEnumTypeDeAuthorized)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReasonEnumTypeDeAuthorized {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReasonEnumTypeDeAuthorized)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReasonEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReasonEnumType
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

func TestRecurrencyKindEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str RecurrencyKindEnumType
		data := fmt.Appendf(nil, `"%s"`, RecurrencyKindEnumTypeDaily)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != RecurrencyKindEnumTypeDaily {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, RecurrencyKindEnumTypeDaily)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str RecurrencyKindEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str RecurrencyKindEnumType
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

func TestRegistrationStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str RegistrationStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, RegistrationStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != RegistrationStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, RegistrationStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str RegistrationStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str RegistrationStatusEnumType
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

func TestReportBaseEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReportBaseEnumType
		data := fmt.Appendf(nil, `"%s"`, ReportBaseEnumTypeConfigurationInventory)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReportBaseEnumTypeConfigurationInventory {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReportBaseEnumTypeConfigurationInventory)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReportBaseEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReportBaseEnumType
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

func TestRequestStartStopStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str RequestStartStopStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, RequestStartStopStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != RequestStartStopStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, RequestStartStopStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str RequestStartStopStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str RequestStartStopStatusEnumType
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

func TestReservationUpdateStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReservationUpdateStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ReservationUpdateStatusEnumTypeExpired)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReservationUpdateStatusEnumTypeExpired {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReservationUpdateStatusEnumTypeExpired)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReservationUpdateStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReservationUpdateStatusEnumType
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

func TestReserveNowStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ReserveNowStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ReserveNowStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ReserveNowStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ReserveNowStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ReserveNowStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ReserveNowStatusEnumType
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

func TestResetEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ResetEnumType
		data := fmt.Appendf(nil, `"%s"`, ResetEnumTypeImmediate)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ResetEnumTypeImmediate {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ResetEnumTypeImmediate)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ResetEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ResetEnumType
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

func TestResetStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str ResetStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, ResetStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != ResetStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, ResetStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str ResetStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str ResetStatusEnumType
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

func TestSendLocalListStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str SendLocalListStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, SendLocalListStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != SendLocalListStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, SendLocalListStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str SendLocalListStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str SendLocalListStatusEnumType
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

func TestSetMonitoringStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str SetMonitoringStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, SetMonitoringStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != SetMonitoringStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, SetMonitoringStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str SetMonitoringStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str SetMonitoringStatusEnumType
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

func TestSetNetworkProfileStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str SetNetworkProfileStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, SetNetworkProfileStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != SetNetworkProfileStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, SetNetworkProfileStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str SetNetworkProfileStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str SetNetworkProfileStatusEnumType
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

func TestSetVariableStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str SetVariableStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, SetVariableStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != SetVariableStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, SetVariableStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str SetVariableStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str SetVariableStatusEnumType
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

func TestTariffChangeStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TariffChangeStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, TariffChangeStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TariffChangeStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TariffChangeStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TariffChangeStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TariffChangeStatusEnumType
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

func TestTariffClearStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TariffClearStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, TariffClearStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TariffClearStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TariffClearStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TariffClearStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TariffClearStatusEnumType
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

func TestTariffCostEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TariffCostEnumType
		data := fmt.Appendf(nil, `"%s"`, TariffCostEnumTypeNormalCost)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TariffCostEnumTypeNormalCost {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TariffCostEnumTypeNormalCost)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TariffCostEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TariffCostEnumType
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

func TestTariffGetStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TariffGetStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, TariffGetStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TariffGetStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TariffGetStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TariffGetStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TariffGetStatusEnumType
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

func TestTariffKindEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TariffKindEnumType
		data := fmt.Appendf(nil, `"%s"`, TariffKindEnumTypeDefaultTariff)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TariffKindEnumTypeDefaultTariff {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TariffKindEnumTypeDefaultTariff)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TariffKindEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TariffKindEnumType
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

func TestTariffSetStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TariffSetStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, TariffSetStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TariffSetStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TariffSetStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TariffSetStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TariffSetStatusEnumType
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

func TestTransactionEventEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TransactionEventEnumType
		data := fmt.Appendf(nil, `"%s"`, TransactionEventEnumTypeEnded)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TransactionEventEnumTypeEnded {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TransactionEventEnumTypeEnded)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TransactionEventEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TransactionEventEnumType
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

func TestTriggerMessageStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TriggerMessageStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, TriggerMessageStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TriggerMessageStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TriggerMessageStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TriggerMessageStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TriggerMessageStatusEnumType
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

func TestTriggerReasonEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str TriggerReasonEnumType
		data := fmt.Appendf(nil, `"%s"`, TriggerReasonEnumTypeAbnormalCondition)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != TriggerReasonEnumTypeAbnormalCondition {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, TriggerReasonEnumTypeAbnormalCondition)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str TriggerReasonEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str TriggerReasonEnumType
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

func TestUnlockStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UnlockStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, UnlockStatusEnumTypeUnlocked)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UnlockStatusEnumTypeUnlocked {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UnlockStatusEnumTypeUnlocked)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UnlockStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UnlockStatusEnumType
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

func TestUnpublishFirmwareStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UnpublishFirmwareStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, UnpublishFirmwareStatusEnumTypeDownloadOngoing)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UnpublishFirmwareStatusEnumTypeDownloadOngoing {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UnpublishFirmwareStatusEnumTypeDownloadOngoing)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UnpublishFirmwareStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UnpublishFirmwareStatusEnumType
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

func TestUpdateEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UpdateEnumType
		data := fmt.Appendf(nil, `"%s"`, UpdateEnumTypeDifferential)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UpdateEnumTypeDifferential {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UpdateEnumTypeDifferential)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UpdateEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UpdateEnumType
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

func TestUpdateFirmwareStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UpdateFirmwareStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, UpdateFirmwareStatusEnumTypeAccepted)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UpdateFirmwareStatusEnumTypeAccepted {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UpdateFirmwareStatusEnumTypeAccepted)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UpdateFirmwareStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UpdateFirmwareStatusEnumType
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

func TestUploadLogStatusEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str UploadLogStatusEnumType
		data := fmt.Appendf(nil, `"%s"`, UploadLogStatusEnumTypeBadMessage)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != UploadLogStatusEnumTypeBadMessage {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, UploadLogStatusEnumTypeBadMessage)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str UploadLogStatusEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str UploadLogStatusEnumType
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

func TestVPNEnumTypeUnmarshalJSON(t *testing.T) {
	t.Run("accepts known value", func(t *testing.T) {
		var str VPNEnumType
		data := fmt.Appendf(nil, `"%s"`, VPNEnumTypeIKEv2)

		if err := json.Unmarshal(data, &str); err != nil {
			t.Fatalf("unexpected unmarshal error: %v", err)
		}

		if str != VPNEnumTypeIKEv2 {
			t.Errorf("unexpected error\ngot:  %q\nwant: %q", str, VPNEnumTypeIKEv2)
		}
	})

	t.Run("rejects unknown value", func(t *testing.T) {
		var str VPNEnumType
		data := []byte(`"Foo"`)

		err := json.Unmarshal(data, &str)

		if !errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
			t.Errorf("unexpected error\ngot:  %v\nwant: errors.Is match for %v", err, ocpp.ErrPropertyConstraintViolation)
		}
	})

	t.Run("rejects wrong JSON type", func(t *testing.T) {
		var str VPNEnumType
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
