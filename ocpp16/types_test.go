package ocpp16

import (
	"encoding/json"
	"errors"
	"testing"
)

// unmarshalEnum unmarshals data into dest and checks that the resulting error
// matches wantErr via errors.Is.
func unmarshalEnum[T comparable](t *testing.T, data []byte, dest *T, wantErr error) {
	t.Helper()
	err := json.Unmarshal(data, dest)
	if !errors.Is(err, wantErr) {
		t.Errorf("unexpected error\ngot:  %v\nwant: %v", err, wantErr)
	}
}

func TestAuthorizationStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s AuthorizationStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != AuthorizationStatusAccepted {
			t.Errorf("got %q, want %q", s, AuthorizationStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s AuthorizationStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s AuthorizationStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestAvailabilityStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s AvailabilityStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != AvailabilityStatusAccepted {
			t.Errorf("got %q, want %q", s, AvailabilityStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s AvailabilityStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s AvailabilityStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestAvailabilityTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s AvailabilityType
		unmarshalEnum(t, []byte(`"Inoperative"`), &s, nil)
		if s != AvailabilityTypeInoperative {
			t.Errorf("got %q, want %q", s, AvailabilityTypeInoperative)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s AvailabilityType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s AvailabilityType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestCancelReservationStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s CancelReservationStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != CancelReservationStatusAccepted {
			t.Errorf("got %q, want %q", s, CancelReservationStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s CancelReservationStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s CancelReservationStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestChargePointErrorCodeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ChargePointErrorCode
		unmarshalEnum(t, []byte(`"ConnectorLockFailure"`), &s, nil)
		if s != ChargePointErrorCodeConnectorLockFailure {
			t.Errorf("got %q, want %q", s, ChargePointErrorCodeConnectorLockFailure)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ChargePointErrorCode
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ChargePointErrorCode
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestChargePointStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ChargePointStatus
		unmarshalEnum(t, []byte(`"Available"`), &s, nil)
		if s != ChargePointStatusAvailable {
			t.Errorf("got %q, want %q", s, ChargePointStatusAvailable)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ChargePointStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ChargePointStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestChargingProfileKindTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ChargingProfileKindType
		unmarshalEnum(t, []byte(`"Absolute"`), &s, nil)
		if s != ChargingProfileKindTypeAbsolute {
			t.Errorf("got %q, want %q", s, ChargingProfileKindTypeAbsolute)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ChargingProfileKindType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ChargingProfileKindType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestChargingProfilePurposeTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ChargingProfilePurposeType
		unmarshalEnum(t, []byte(`"ChargePointMaxProfile"`), &s, nil)
		if s != ChargingProfilePurposeTypeChargePointMaxProfile {
			t.Errorf("got %q, want %q", s, ChargingProfilePurposeTypeChargePointMaxProfile)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ChargingProfilePurposeType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ChargingProfilePurposeType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestChargingProfileStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ChargingProfileStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != ChargingProfileStatusAccepted {
			t.Errorf("got %q, want %q", s, ChargingProfileStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ChargingProfileStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ChargingProfileStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestChargingRateUnitTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ChargingRateUnitType
		unmarshalEnum(t, []byte(`"W"`), &s, nil)
		if s != ChargingRateUnitTypeW {
			t.Errorf("got %q, want %q", s, ChargingRateUnitTypeW)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ChargingRateUnitType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ChargingRateUnitType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestClearCacheStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ClearCacheStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != ClearCacheStatusAccepted {
			t.Errorf("got %q, want %q", s, ClearCacheStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ClearCacheStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ClearCacheStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestClearChargingProfileStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ClearChargingProfileStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != ClearChargingProfileStatusAccepted {
			t.Errorf("got %q, want %q", s, ClearChargingProfileStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ClearChargingProfileStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ClearChargingProfileStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestConfigurationStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ConfigurationStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != ConfigurationStatusAccepted {
			t.Errorf("got %q, want %q", s, ConfigurationStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ConfigurationStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ConfigurationStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestDataTransferStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s DataTransferStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != DataTransferStatusAccepted {
			t.Errorf("got %q, want %q", s, DataTransferStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s DataTransferStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s DataTransferStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestDiagnosticsStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s DiagnosticsStatus
		unmarshalEnum(t, []byte(`"Idle"`), &s, nil)
		if s != DiagnosticsStatusIdle {
			t.Errorf("got %q, want %q", s, DiagnosticsStatusIdle)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s DiagnosticsStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s DiagnosticsStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestFirmwareStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s FirmwareStatus
		unmarshalEnum(t, []byte(`"Downloaded"`), &s, nil)
		if s != FirmwareStatusDownloaded {
			t.Errorf("got %q, want %q", s, FirmwareStatusDownloaded)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s FirmwareStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s FirmwareStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestGetCompositeScheduleStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s GetCompositeScheduleStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != GetCompositeScheduleStatusAccepted {
			t.Errorf("got %q, want %q", s, GetCompositeScheduleStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s GetCompositeScheduleStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s GetCompositeScheduleStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestLocationUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s Location
		unmarshalEnum(t, []byte(`"Body"`), &s, nil)
		if s != LocationBody {
			t.Errorf("got %q, want %q", s, LocationBody)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s Location
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s Location
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestMeasurandUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s Measurand
		unmarshalEnum(t, []byte(`"Current.Export"`), &s, nil)
		if s != MeasurandCurrentExport {
			t.Errorf("got %q, want %q", s, MeasurandCurrentExport)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s Measurand
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s Measurand
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestMessageTriggerUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s MessageTrigger
		unmarshalEnum(t, []byte(`"BootNotification"`), &s, nil)
		if s != MessageTriggerBootNotification {
			t.Errorf("got %q, want %q", s, MessageTriggerBootNotification)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s MessageTrigger
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s MessageTrigger
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestPhaseUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s Phase
		unmarshalEnum(t, []byte(`"L1"`), &s, nil)
		if s != PhaseL1 {
			t.Errorf("got %q, want %q", s, PhaseL1)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s Phase
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s Phase
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestReadingContextUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ReadingContext
		unmarshalEnum(t, []byte(`"Interruption.Begin"`), &s, nil)
		if s != ReadingContextInterruptionBegin {
			t.Errorf("got %q, want %q", s, ReadingContextInterruptionBegin)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ReadingContext
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ReadingContext
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestReasonUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s Reason
		unmarshalEnum(t, []byte(`"DeAuthorized"`), &s, nil)
		if s != ReasonDeAuthorized {
			t.Errorf("got %q, want %q", s, ReasonDeAuthorized)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s Reason
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s Reason
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestRecurrencyKindTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s RecurrencyKindType
		unmarshalEnum(t, []byte(`"Daily"`), &s, nil)
		if s != RecurrencyKindTypeDaily {
			t.Errorf("got %q, want %q", s, RecurrencyKindTypeDaily)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s RecurrencyKindType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s RecurrencyKindType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestRegistrationStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s RegistrationStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != RegistrationStatusAccepted {
			t.Errorf("got %q, want %q", s, RegistrationStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s RegistrationStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s RegistrationStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestRemoteStartStopStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s RemoteStartStopStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != RemoteStartStopStatusAccepted {
			t.Errorf("got %q, want %q", s, RemoteStartStopStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s RemoteStartStopStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s RemoteStartStopStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestReservationStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ReservationStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != ReservationStatusAccepted {
			t.Errorf("got %q, want %q", s, ReservationStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ReservationStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ReservationStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestResetStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ResetStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != ResetStatusAccepted {
			t.Errorf("got %q, want %q", s, ResetStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ResetStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ResetStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestResetTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ResetType
		unmarshalEnum(t, []byte(`"Hard"`), &s, nil)
		if s != ResetTypeHard {
			t.Errorf("got %q, want %q", s, ResetTypeHard)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ResetType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ResetType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestTriggerMessageStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s TriggerMessageStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != TriggerMessageStatusAccepted {
			t.Errorf("got %q, want %q", s, TriggerMessageStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s TriggerMessageStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s TriggerMessageStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestUnitOfMeasureUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s UnitOfMeasure
		unmarshalEnum(t, []byte(`"Wh"`), &s, nil)
		if s != UnitOfMeasureWh {
			t.Errorf("got %q, want %q", s, UnitOfMeasureWh)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s UnitOfMeasure
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s UnitOfMeasure
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestUnlockStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s UnlockStatus
		unmarshalEnum(t, []byte(`"Unlocked"`), &s, nil)
		if s != UnlockStatusUnlocked {
			t.Errorf("got %q, want %q", s, UnlockStatusUnlocked)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s UnlockStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s UnlockStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestUpdateStatusUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s UpdateStatus
		unmarshalEnum(t, []byte(`"Accepted"`), &s, nil)
		if s != UpdateStatusAccepted {
			t.Errorf("got %q, want %q", s, UpdateStatusAccepted)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s UpdateStatus
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s UpdateStatus
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestUpdateTypeUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s UpdateType
		unmarshalEnum(t, []byte(`"Differential"`), &s, nil)
		if s != UpdateTypeDifferential {
			t.Errorf("got %q, want %q", s, UpdateTypeDifferential)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s UpdateType
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s UpdateType
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestValueFormatUnmarshalJSON(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		var s ValueFormat
		unmarshalEnum(t, []byte(`"Raw"`), &s, nil)
		if s != ValueFormatRaw {
			t.Errorf("got %q, want %q", s, ValueFormatRaw)
		}
	})
	t.Run("invalid value", func(t *testing.T) {
		var s ValueFormat
		unmarshalEnum(t, []byte(`"Bogus"`), &s, ErrInvalidEnum)
	})
	t.Run("non-string", func(t *testing.T) {
		var s ValueFormat
		unmarshalEnum(t, []byte(`1`), &s, ErrTypeString)
	})
}

func TestNewCiString20Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString20Type
		err        error
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"foooooooooooooooooooo"`),
			err:        ErrStringLength20,
		},
		{
			name:       "should error if not string",
			inputBytes: []byte(`1`),
			err:        ErrTypeString,
		},
		{
			name:       "should error if containing space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "should accept lower boundary visible ASCII (0x21)",
			inputBytes: []byte(`"` + "\x21" + `"`),
			want:       CiString20Type("\x21"),
			err:        nil,
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
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foo"`),
			want:       CiString20Type("foo"),
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

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}
		})
	}
}

func TestNewCiString25Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString25Type
		err        error
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"fooooooooooooooooooooooooo"`),
			err:        ErrStringLength25,
		},
		{
			name:       "should error if not string",
			inputBytes: []byte(`1`),
			err:        ErrTypeString,
		},
		{
			name:       "should error if containing space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "should accept lower boundary visible ASCII (0x21)",
			inputBytes: []byte(`"` + "\x21" + `"`),
			want:       CiString25Type("\x21"),
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
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"foooooooooooooooooooooooo"`),
			want:       CiString25Type("foooooooooooooooooooooooo"),
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

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}
		})
	}
}

func TestNewCiString50Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString50Type
		err        error
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"foooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			err:        ErrStringLength50,
		},
		{
			name:       "should error if not string",
			inputBytes: []byte(`1`),
			err:        ErrTypeString,
		},
		{
			name:       "should error if containing space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "should accept lower boundary visible ASCII (0x21)",
			inputBytes: []byte(`"` + "\x21" + `"`),
			want:       CiString50Type("\x21"),
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
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"fooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			want:       CiString50Type("fooooooooooooooooooooooooooooooooooooooooooooooooo"),
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

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}
		})
	}
}

func TestNewCiString255Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString255Type
		err        error
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"fooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			err:        ErrStringLength255,
		},
		{
			name:       "should error if not string",
			inputBytes: []byte(`1`),
			err:        ErrTypeString,
		},
		{
			name:       "should error if containing space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "should accept lower boundary visible ASCII (0x21)",
			inputBytes: []byte(`"` + "\x21" + `"`),
			want:       CiString255Type("\x21"),
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
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"fooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			want:       CiString255Type("fooooooooooooooooooooooooooooooooooooooooooooooooo"),
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

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}
		})
	}
}

func TestNewCiString500Type(t *testing.T) {
	tests := []struct {
		name       string
		inputBytes []byte
		want       CiString500Type
		err        error
	}{
		{
			name:       "should error if string too long",
			inputBytes: []byte(`"foooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			err:        ErrStringLength500,
		},
		{
			name:       "should error if not string",
			inputBytes: []byte(`1`),
			err:        ErrTypeString,
		},
		{
			name:       "should error if containing space character (0x20)",
			inputBytes: []byte(`"` + "\x20" + `"`),
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "should accept lower boundary visible ASCII (0x21)",
			inputBytes: []byte(`"` + "\x21" + `"`),
			want:       CiString500Type("\x21"),
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
			err:        ErrNonVisibleASCII,
		},
		{
			name:       "can unmarshal string",
			inputBytes: []byte(`"fooooooooooooooooooooooooooooooooooooooooooooooooo"`),
			want:       CiString500Type("fooooooooooooooooooooooooooooooooooooooooooooooooo"),
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

			if str != tt.want {
				t.Errorf("unexpected string\ngot:  %v\nwant: %v", str, tt.want)
			}
		})
	}
}
