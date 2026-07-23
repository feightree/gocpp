package v16

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// AuthorizeReq (6.1)
//
// This contains the field definition of the Authorize.req PDU sent by the Charge Point to the Central System. See
// also Authorize
type AuthorizeReq struct {
	// Required. This contains the identifier that needs to be authorized.
	IDTag IDToken `json:"idTag"`
}

func (s *AuthorizeReq) UnmarshalJSON(data []byte) error {
	type Alias AuthorizeReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AuthorizeReq(a)
	return s.Validate()
}

func (s AuthorizeReq) Validate() error {
	if s.IDTag == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idTag", "required field is missing")
	}

	if err := s.IDTag.Validate(); err != nil {
		return ocpp.WrapField("idTag", err)
	}

	return nil
}

// AuthorizeConf (6.2)
//
// This contains the field definition of the Authorize.conf PDU sent by the Central System to the Charge Point in
// response to a Authorize.req PDU. See also Authorize
type AuthorizeConf struct {
	// Required. This contains information about authorization status, expiry parent id.
	IDTagInfo IDTagInfo `json:"idTagInfo"`
}

func (s *AuthorizeConf) UnmarshalJSON(data []byte) error {
	type Alias AuthorizeConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AuthorizeConf(a)
	return s.Validate()
}

func (s AuthorizeConf) Validate() error {
	if err := s.IDTagInfo.Validate(); err != nil {
		return ocpp.WrapField("idTagInfo", err)
	}

	return nil
}

// BootNotificationReq (6.3)
//
// This contains the field definition of the BootNotification.req PDU sent by the Charge Point to the Central System.
// See also Boot Notification
type BootNotificationReq struct {
	// Optional. This contains a value that identifies the serial number of
	// the Charge Box inside the Charge Point. Deprecated, will be
	// removed in future version
	ChargeBoxSerialNumber *CiString25Type `json:"chargeBoxSerialNumber,omitempty"`
	// Required. This contains a value that identifies the model of the
	// ChargePoint.
	ChargePointModel CiString20Type `json:"chargePointModel"`
	// Optional. This contains a value that identifies the serial number of
	// the Charge Point.
	ChargePointSerialNumber *CiString25Type `json:"chargePointSerialNumber,omitempty"`
	// Required. This contains a value that identifies the vendor of the
	// ChargePoint.
	ChargePointVendor CiString20Type `json:"chargePointVendor"`
	// Optional. This contains the firmware version of the Charge Point.
	FirmwareVersion *CiString50Type `json:"firmwareVersion,omitempty"`
	// Optional. This contains the ICCID of the modem’s SIM card.
	ICCID *CiString20Type `json:"iccid,omitempty"`
	// Optional. This contains the IMSI of the modem’s SIM card.
	IMSI *CiString20Type `json:"imsi,omitempty"`
	// Optional. This contains the serial number of the main electrical
	// meter of the Charge Point.
	MeterSerialNumber *CiString25Type `json:"meterSerialNumber,omitempty"`
	// Optional. This contains the type of the main electrical meter of
	// the Charge Point.
	MeterType *CiString25Type `json:"meterType,omitempty"`
}

func (s *BootNotificationReq) UnmarshalJSON(data []byte) error {
	type Alias BootNotificationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BootNotificationReq(a)
	return s.Validate()
}

func (s BootNotificationReq) Validate() error {
	if s.ChargeBoxSerialNumber != nil {
		if err := s.ChargeBoxSerialNumber.Validate(); err != nil {
			return ocpp.WrapField("chargeBoxSerialNumber", err)
		}
	}

	if s.ChargePointModel == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargePointModel", "required field is missing")
	}

	if err := s.ChargePointModel.Validate(); err != nil {
		return ocpp.WrapField("chargePointModel", err)
	}

	if s.ChargePointSerialNumber != nil {
		if err := s.ChargePointSerialNumber.Validate(); err != nil {
			return ocpp.WrapField("chargePointSerialNumber", err)
		}
	}

	if s.ChargePointVendor == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargePointVendor", "required field is missing")
	}

	if err := s.ChargePointVendor.Validate(); err != nil {
		return ocpp.WrapField("chargePointVendor", err)
	}

	if s.FirmwareVersion != nil {
		if err := s.FirmwareVersion.Validate(); err != nil {
			return ocpp.WrapField("firmwareVersion", err)
		}
	}

	if s.ICCID != nil {
		if err := s.ICCID.Validate(); err != nil {
			return ocpp.WrapField("iccid", err)
		}
	}

	if s.IMSI != nil {
		if err := s.IMSI.Validate(); err != nil {
			return ocpp.WrapField("imsi", err)
		}
	}

	if s.MeterSerialNumber != nil {
		if err := s.MeterSerialNumber.Validate(); err != nil {
			return ocpp.WrapField("meterSerialNumber", err)
		}
	}

	if s.MeterType != nil {
		if err := s.MeterType.Validate(); err != nil {
			return ocpp.WrapField("meterType", err)
		}
	}

	return nil
}

// BootNotificationConf (6.4)
//
// This contains the field definition of the BootNotification.conf PDU sent by the Central System to the Charge Point
// in response to a BootNotification.req PDU. See also Boot Notification
type BootNotificationConf struct {
	// Required. This contains the Central System’s current time.
	CurrentTime time.Time `json:"currentTime"`
	// Required. When RegistrationStatus is Accepted, this contains the heartbeat
	// interval in seconds. If the Central System returns something other than
	// Accepted, the value of the interval field indicates the minimum wait time before
	// sending a next BootNotification request.
	Interval int32 `json:"interval"`
	// Required. This contains whether the Charge Point has been registered within System Central.
	Status RegistrationStatus `json:"status"`
}

func (s *BootNotificationConf) UnmarshalJSON(data []byte) error {
	type Alias BootNotificationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BootNotificationConf(a)
	return s.Validate()
}

func (s BootNotificationConf) Validate() error {
	if s.CurrentTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currentTime", "required field is missing")
	}

	if s.Interval < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "interval", "must be > 0")
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// CancelReservationReq (6.5)
//
// This contains the field definition of the CancelReservation.req PDU sent by the Central System to the Charge
// Point. See also Cancel Reservation
type CancelReservationReq struct {
	// Required. Id of the reservation to cancel.
	ReservationID int32 `json:"reservationId"`
}

func (s *CancelReservationReq) UnmarshalJSON(data []byte) error {
	type Alias CancelReservationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CancelReservationReq(a)
	return s.Validate()
}

func (s CancelReservationReq) Validate() error {
	if s.ReservationID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "reservationId", "must be > 0")
	}

	return nil
}

// CancelReservationConf (6.6)
//
// This contains the field definition of the CancelReservation.conf PDU sent by the Charge Point to the Central
// System in response to a CancelReservation.req PDU. See also Cancel Reservation
type CancelReservationConf struct {
	// Required. This indicates the success or failure of the cancelling of
	// a reservation by Central System.
	Status CancelReservationStatus `json:"status"`
}

func (s *CancelReservationConf) UnmarshalJSON(data []byte) error {
	type Alias CancelReservationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CancelReservationConf(a)
	return s.Validate()
}

func (s CancelReservationConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// ChangeAvailabilityReq (6.7)
//
// This contains the field definition of the ChangeAvailability.req PDU sent by the Central System to the Charge
// Point. See also Change Availability
type ChangeAvailabilityReq struct {
	// Required. The id of the connector for which availability needs to change. Id (zero) is used if the availability of the Charge Point and all its connectors needs
	// to change.
	ConnectorID int32 `json:"connectorId"`
	// Required. This contains the type of availability change that the Charge Point
	// should perform.
	Type AvailabilityType `json:"type"`
}

func (s *ChangeAvailabilityReq) UnmarshalJSON(data []byte) error {
	type Alias ChangeAvailabilityReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeAvailabilityReq(a)
	return s.Validate()
}

func (s ChangeAvailabilityReq) Validate() error {
	if s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	return nil
}

// ChangeAvailabilityConf (6.8)
//
// This contains the field definition of the ChangeAvailability.conf PDU return by Charge Point to Central System.
// See also Change Availability
type ChangeAvailabilityConf struct {
	// Required. This indicates whether the Charge Point is able to perform the
	// availability change.
	Status AvailabilityStatus `json:"status"`
}

func (s *ChangeAvailabilityConf) UnmarshalJSON(data []byte) error {
	type Alias ChangeAvailabilityConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeAvailabilityConf(a)
	return s.Validate()
}

func (s ChangeAvailabilityConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// ChangeConfigurationReq (6.9)
//
// This contains the field definition of the ChangeConfiguration.req PDU sent by Central System to Charge Point. It
// is RECOMMENDED that the content and meaning of the 'key' and 'value' fields is agreed upon between Charge
// Point and Central System. See also Change Configuration
type ChangeConfigurationReq struct {
	// Required. The name of the configuration setting to change.
	// See for standard configuration key names and associated values
	Key CiString50Type `json:"key"`
	// Required. The new value as string for the setting.
	// See for standard configuration key names and associated values.
	Value CiString500Type `json:"value"`
}

func (s *ChangeConfigurationReq) UnmarshalJSON(data []byte) error {
	type Alias ChangeConfigurationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeConfigurationReq(a)
	return s.Validate()
}

func (s ChangeConfigurationReq) Validate() error {
	if s.Key == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "key", "required field is missing")
	}

	if err := s.Key.Validate(); err != nil {
		return ocpp.WrapField("key", err)
	}

	if err := s.Value.Validate(); err != nil {
		return ocpp.WrapField("value", err)
	}

	return nil
}

// ChangeConfigurationConf (6.10)
//
// This contains the field definition of the ChangeConfiguration.conf PDU returned from Charge Point to Central
// System. See also Change Configuration
type ChangeConfigurationConf struct {
	// Required. Returns whether configuration change has been accepted.
	Status ConfigurationStatus `json:"status"`
}

func (s *ChangeConfigurationConf) UnmarshalJSON(data []byte) error {
	type Alias ChangeConfigurationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeConfigurationConf(a)
	return s.Validate()
}

func (s ChangeConfigurationConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// ClearCacheReq (6.11)
//
// This contains the field definition of the ClearCache.req PDU sent by the Central System to the Charge Point. See
// also Clear Cache
//
// No fields are defined.
type ClearCacheReq struct{}

func (s *ClearCacheReq) UnmarshalJSON(data []byte) error {
	type Alias ClearCacheReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearCacheReq(a)
	return s.Validate()
}

func (s ClearCacheReq) Validate() error {
	return nil
}

// ClearCacheConf (6.12)
//
// This contains the field definition of the ClearCache.conf PDU sent by the Charge Point to the Central System in
// response to a ClearCache.req PDU. See also Clear Cache
type ClearCacheConf struct {
	// Required. Accepted if the Charge Point has executed the request, otherwise rejected.
	Status ClearCacheStatus `json:"status"`
}

func (s *ClearCacheConf) UnmarshalJSON(data []byte) error {
	type Alias ClearCacheConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearCacheConf(a)
	return s.Validate()
}

func (s ClearCacheConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// ClearChargingProfileReq (6.13)
//
// This contains the field definition of the ClearChargingProfile.req PDU sent by the Central System to the Charge
// Point.
//
// The Central System can use this message to clear (remove) either a specific charging profile (denoted by id) or a
// selection of charging profiles that match with the values of the optional connectorId, stackLevel and
// chargingProfilePurpose fields. See also Clear Charging Profile
type ClearChargingProfileReq struct {
	// Optional. The ID of the charging profile to clear.
	ID *int32 `json:"id,omitempty"`
	// Optional. Specifies the ID of the connector for which to clear
	// charging profiles. A connectorId of zero (0) specifies the charging
	// profile for the overall Charge Point. Absence of this parameter
	// means the clearing applies to all charging profiles that match the
	// other criteria in the request.
	ConnectorID *int32 `json:"connectorId,omitempty"`
	// Optional. Specifies to purpose of the charging profiles that will be
	// cleared, if they meet the other criteria in the request.
	ChargingProfilePurpose *ChargingProfilePurposeType `json:"chargingProfilePurpose,omitempty"`
	// Optional. specifies the stackLevel for which charging profiles will
	// be cleared, if they meet the other criteria in the request
	StackLevel *int32 `json:"stackLevel,omitempty"`
}

func (s *ClearChargingProfileReq) UnmarshalJSON(data []byte) error {
	type Alias ClearChargingProfileReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearChargingProfileReq(a)
	return s.Validate()
}

func (s ClearChargingProfileReq) Validate() error {
	if s.ConnectorID != nil && *s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if s.StackLevel != nil && *s.StackLevel < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "stackLevel", "must be >= 0")
	}

	return nil
}

// ClearChargingProfileConf (6.14)
//
// This contains the field definition of the ClearChargingProfile.conf PDU sent by the Charge Point to the Central
// System in response to a ClearChargingProfile.req PDU. See also Clear Charging Profile
type ClearChargingProfileConf struct {
	// Required. Indicates if the Charge Point was able to execute the request.
	Status ClearChargingProfileStatus `json:"status"`
}

func (s *ClearChargingProfileConf) UnmarshalJSON(data []byte) error {
	type Alias ClearChargingProfileConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearChargingProfileConf(a)
	return s.Validate()
}

func (s ClearChargingProfileConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// DataTransferReq (6.15)
//
// This contains the field definition of the DataTransfer.req PDU sent either by the Central System to the Charge
// Point or vice versa. See also Data Transfer
type DataTransferReq struct {
	// Required. This identifies the Vendor specific implementation
	VendorID CiString255Type `json:"vendorId"`
	// Optional. Additional identification field
	MessageID *CiString50Type `json:"messageId,omitempty"`
	// Optional. Data without specified length or format.
	Data *string `json:"data,omitempty"`
}

func (s *DataTransferReq) UnmarshalJSON(data []byte) error {
	type Alias DataTransferReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DataTransferReq(a)
	return s.Validate()
}

func (s DataTransferReq) Validate() error {
	if s.VendorID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "vendorId", "required field is missing")
	}

	if err := s.VendorID.Validate(); err != nil {
		return ocpp.WrapField("vendorId", err)
	}

	if s.MessageID != nil {
		if err := s.MessageID.Validate(); err != nil {
			return ocpp.WrapField("messageId", err)
		}
	}

	return nil
}

// DataTransferConf (6.16)
//
// This contains the field definition of the DataTransfer.conf PDU sent by the Charge Point to the Central System or
// vice versa in response to a DataTransfer.req PDU. See also Data Transfer
type DataTransferConf struct {
	// Required. This indicates the success or failure of the data transfer.
	Status DataTransferStatus `json:"status"`
	// Optional. Data in response to request.
	Data *string `json:"data,omitempty"`
}

func (s *DataTransferConf) UnmarshalJSON(data []byte) error {
	type Alias DataTransferConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DataTransferConf(a)
	return s.Validate()
}

func (s DataTransferConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// DiagnosticsStatusNotificationReq (6.17)
//
// This contains the field definition of the DiagnosticsStatusNotification.req PDU sent by the Charge Point to the
// Central System. See also Diagnostics Status Notification
type DiagnosticsStatusNotificationReq struct {
	// Required. This contains the status of the diagnostics upload.
	Status DiagnosticsStatus `json:"status"`
}

func (s *DiagnosticsStatusNotificationReq) UnmarshalJSON(data []byte) error {
	type Alias DiagnosticsStatusNotificationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DiagnosticsStatusNotificationReq(a)
	return s.Validate()
}

func (s DiagnosticsStatusNotificationReq) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// DiagnosticsStatusNotificationConf (6.18)
//
// This contains the field definition of the DiagnosticsStatusNotification.conf PDU sent by the Central System to the
// Charge Point in response to a DiagnosticsStatusNotification.req PDU. See also Diagnostics Status Notification
// No fields are defined.
type DiagnosticsStatusNotificationConf struct{}

func (s *DiagnosticsStatusNotificationConf) UnmarshalJSON(data []byte) error {
	type Alias DiagnosticsStatusNotificationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DiagnosticsStatusNotificationConf(a)
	return s.Validate()
}

func (s DiagnosticsStatusNotificationConf) Validate() error {
	return nil
}

// FirmwareStatusNotificationReq (6.19)
//
// This contains the field definition of the FirmwareStatusNotification.req PDU sent by the Charge Point to the
// Central System. See also Firmware Status Notification
type FirmwareStatusNotificationReq struct {
	// Required. This contains the progress status of the firmware installation.
	Status FirmwareStatus `json:"status"`
}

func (s *FirmwareStatusNotificationReq) UnmarshalJSON(data []byte) error {
	type Alias FirmwareStatusNotificationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FirmwareStatusNotificationReq(a)
	return s.Validate()
}

func (s FirmwareStatusNotificationReq) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// FirmwareStatusNotificationConf (6.20)
//
// This contains the field definition of the FirmwareStatusNotification.conf PDU sent by the Central System to the
// Charge Point in response to a FirmwareStatusNotification.req PDU. See also Firmware Status Notification
// No fields are defined.
type FirmwareStatusNotificationConf struct{}

func (s *FirmwareStatusNotificationConf) UnmarshalJSON(data []byte) error {
	type Alias FirmwareStatusNotificationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FirmwareStatusNotificationConf(a)
	return s.Validate()
}

func (s FirmwareStatusNotificationConf) Validate() error {
	return nil
}

// GetCompositeScheduleReq (6.21)
//
// This contains the field definition of the GetCompositeSchedule.req PDU sent by the Central System to the
// Charge Point. See also Get Composite Schedule
type GetCompositeScheduleReq struct {
	// Required. The ID of the Connector for which the schedule is
	// requested. When ConnectorId=0, the Charge Point will calculate
	// the expected consumption for the grid connection.
	ConnectorID int32 `json:"connectorId"`
	// Required. Time in seconds. length of requested schedule
	Duration int32 `json:"duration"`
	// Optional. Can be used to force a power or current profile
	ChargingRateUnit *ChargingRateUnitType `json:"chargingRateUnit,omitempty"`
}

func (s *GetCompositeScheduleReq) UnmarshalJSON(data []byte) error {
	type Alias GetCompositeScheduleReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCompositeScheduleReq(a)
	return s.Validate()
}

func (s GetCompositeScheduleReq) Validate() error {
	if s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if s.Duration < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "duration", "must be > 0")
	}

	return nil
}

// GetCompositeScheduleConf (6.22)
//
// This contains the field definition of the GetCompositeSchedule.conf PDU sent by the Charge Point to the Central
// System in response to a GetCompositeSchedule.req PDU. See also Get Composite Schedule
type GetCompositeScheduleConf struct {
	// Required. Status of the request. The Charge Point will indicate if it
	// was able to process the request
	Status GetCompositeScheduleStatus `json:"status"`
	// Optional. The charging schedule contained in this notification
	// applies to a Connector.
	ConnectorID *int32 `json:"connectorId,omitempty"`
	// Optional. Time. Periods contained in the charging profile are
	// relative to this point in time.
	// If status is "Rejected", this field may be absent.
	ScheduleStart *time.Time `json:"scheduleStart,omitempty"`
	// Optional. Planned Composite Charging Schedule, the energy
	// consumption over time. Always relative to ScheduleStart.
	// If status is "Rejected", this field may be absent.
	ChargingSchedule *ChargingSchedule `json:"chargingSchedule,omitempty"`
}

func (s *GetCompositeScheduleConf) UnmarshalJSON(data []byte) error {
	type Alias GetCompositeScheduleConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCompositeScheduleConf(a)
	return s.Validate()
}

func (s GetCompositeScheduleConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.ConnectorID != nil && *s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	return nil
}

// GetConfigurationReq (6.23)
//
// This contains the field definition of the GetConfiguration.req PDU sent by the Central System to the Charge
// Point. See also Get Configuration
type GetConfigurationReq struct {
	// Optional. List of keys for which the configuration value is requested.
	Key []CiString50Type `json:"key,omitempty"`
}

func (s *GetConfigurationReq) UnmarshalJSON(data []byte) error {
	type Alias GetConfigurationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetConfigurationReq(a)
	return s.Validate()
}

func (s GetConfigurationReq) Validate() error {
	for i, p := range s.Key {
		if err := p.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("key[%d]", i), err)
		}
	}

	return nil
}

// GetConfigurationConf (6.24)
//
// This contains the field definition of the GetConfiguration.conf PDU sent by the Charge Point to the Central
// System in response to a GetConfiguration.req. See also Get Configuration
type GetConfigurationConf struct {
	// Optional. List of requested or known keys
	ConfigurationKey []KeyValue `json:"configurationKey,omitempty"`
	// Optional. Requested keys that are unknown
	UnknownKey []CiString50Type `json:"unknownKey,omitempty"`
}

func (s *GetConfigurationConf) UnmarshalJSON(data []byte) error {
	type Alias GetConfigurationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetConfigurationConf(a)
	return s.Validate()
}

func (s GetConfigurationConf) Validate() error {
	for i, p := range s.ConfigurationKey {
		if err := p.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("configurationKey[%d]", i), err)
		}
	}

	for i, p := range s.UnknownKey {
		if err := p.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("unknownKey[%d]", i), err)
		}
	}

	return nil
}

// GetDiagnosticsReq (6.25)
//
// This contains the field definition of the GetDiagnostics.req PDU sent by the Central System to the Charge Point.
// See also Get Diagnostics
type GetDiagnosticsReq struct {
	// Required. This contains the location (directory) where the diagnostics file shall
	// be uploaded to.
	Location string `json:"location"`
	// Optional. This specifies how many times Charge Point must try to upload the
	// diagnostics before giving up. If this field is not present, it is left to Charge Point
	// to decide how many times it wants to retry.
	Retries *int32 `json:"retries,omitempty"`
	// Optional. The interval in seconds after which a retry may be attempted. If this
	// field is not present, it is left to Charge Point to decide how long to wait between
	// attempts.
	RetryInterval *int32 `json:"retryInterval,omitempty"`
	// Optional. This contains the date and time of the oldest logging information to
	// include in the diagnostics.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Optional. This contains the date and time of the latest logging information to
	// include in the diagnostics.
	StopTime *time.Time `json:"stopTime,omitempty"`
}

func (s *GetDiagnosticsReq) UnmarshalJSON(data []byte) error {
	type Alias GetDiagnosticsReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetDiagnosticsReq(a)
	return s.Validate()
}

func (s GetDiagnosticsReq) Validate() error {
	if s.Location == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "location", "required field is missing")
	}

	if s.Retries != nil && *s.Retries < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "retries", "must be >= 0")
	}

	if s.RetryInterval != nil && *s.RetryInterval < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "retryInterval", "must be >= 1")
	}

	return nil
}

// GetDiagnosticsConf (6.26)
//
// This contains the field definition of the GetDiagnostics.conf PDU sent by the Charge Point to the Central System
// in response to a GetDiagnostics.req PDU. See also Get Diagnostics
type GetDiagnosticsConf struct {
	// Optional. This contains the name of the file with diagnostic information that will
	// be uploaded. This field is not present when no diagnostic information is
	// available.
	FileName *CiString255Type `json:"fileName,omitempty"`
}

func (s *GetDiagnosticsConf) UnmarshalJSON(data []byte) error {
	type Alias GetDiagnosticsConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetDiagnosticsConf(a)
	return s.Validate()
}

func (s GetDiagnosticsConf) Validate() error {
	if s.FileName != nil {
		if err := s.FileName.Validate(); err != nil {
			return ocpp.WrapField("fileName", err)
		}
	}

	return nil
}

// GetLocalListVersionReq (6.27)
//
// This contains the field definition of the GetLocalListVersion.req PDU sent by the Central System to the Charge
// Point. See also Get Local List Version
//
// No fields are defined.
type GetLocalListVersionReq struct{}

func (s *GetLocalListVersionReq) UnmarshalJSON(data []byte) error {
	type Alias GetLocalListVersionReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetLocalListVersionReq(a)
	return s.Validate()
}

func (s GetLocalListVersionReq) Validate() error {
	return nil
}

// GetLocalListVersionConf (6.28)
//
// This contains the field definition of the GetLocalListVersion.conf PDU sent by the Charge Point to Central System
// in response to a GetLocalListVersion.req PDU. See also Get Local List Version
type GetLocalListVersionConf struct {
	// Required. This contains the current version number of the local authorization list in the Charge Point.
	ListVersion int32 `json:"listVersion"`
}

func (s *GetLocalListVersionConf) UnmarshalJSON(data []byte) error {
	type Alias GetLocalListVersionConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetLocalListVersionConf(a)
	return s.Validate()
}

func (s GetLocalListVersionConf) Validate() error {
	return nil
}

// HeartbeatReq (6.29)
//
// This contains the field definition of the Heartbeat.req PDU sent by the Charge Point to the Central System. See
// also Heartbeat
//
// No fields are defined.
type HeartbeatReq struct{}

func (s *HeartbeatReq) UnmarshalJSON(data []byte) error {
	type Alias HeartbeatReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = HeartbeatReq(a)
	return s.Validate()
}

func (s HeartbeatReq) Validate() error {
	return nil
}

// HeartbeatConf (6.30)
//
// This contains the field definition of the Heartbeat.conf PDU sent by the Central System to the Charge Point in
// response to a Heartbeat.req PDU. See also Heartbeat
type HeartbeatConf struct {
	// Required. This contains the current time of the Central System.
	CurrentTime time.Time `json:"currentTime"`
}

func (s *HeartbeatConf) UnmarshalJSON(data []byte) error {
	type Alias HeartbeatConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = HeartbeatConf(a)
	return s.Validate()
}

func (s HeartbeatConf) Validate() error {
	if s.CurrentTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currentTime", "required field is missing")
	}

	return nil
}

// MeterValuesReq (6.31)
//
// This contains the field definition of the MeterValues.req PDU sent by the Charge Point to the Central System. See
// also Meter Values
type MeterValuesReq struct {
	// Required. This contains a number (>0) designating a connector of the Charge Point.
	//
	// ‘0’ (zero) is used to designate the main powermeter.
	ConnectorID int32 `json:"connectorId"`
	// Optional. The transaction to which these meter samples are related.
	TransactionID *int32 `json:"transactionId,omitempty"`
	// Required. The sampled meter values with timestamps.
	MeterValue []MeterValue `json:"meterValue"`
}

func (s *MeterValuesReq) UnmarshalJSON(data []byte) error {
	type Alias MeterValuesReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MeterValuesReq(a)
	return s.Validate()
}

func (s MeterValuesReq) Validate() error {
	if s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if len(s.MeterValue) == 0 {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "meterValue", "must not be an empty array")
	}

	for i, v := range s.MeterValue {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("meterValue[%d]", i), err)
		}
	}

	return nil
}

// MeterValuesConf (6.32)
//
// This contains the field definition of the MeterValues.conf PDU sent by the Central System to the Charge Point in
// response to a MeterValues.req PDU. See also Meter Values
//
// No fields are defined.
type MeterValuesConf struct{}

func (s *MeterValuesConf) UnmarshalJSON(data []byte) error {
	type Alias MeterValuesConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MeterValuesConf(a)
	return s.Validate()
}

func (s MeterValuesConf) Validate() error {
	return nil
}

// RemoteStartTransactionReq (6.33)
//
// This contains the field definitions of the RemoteStartTransaction.req PDU sent to Charge Point by Central
// System. See also Remote Start Transaction
type RemoteStartTransactionReq struct {
	// Optional. Number of the connector on which to start the transaction.
	// connectorId SHALL be > 0
	ConnectorID *int32 `json:"connectorId,omitempty"`
	// Required. The identifier that Charge Point must use to start a transaction.
	IDTag IDToken `json:"idTag"`
	// Optional. Charging Profile to be used by the Charge Point for the requested
	// transaction. ChargingProfilePurpose MUST be set to TxProfile
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
}

func (s *RemoteStartTransactionReq) UnmarshalJSON(data []byte) error {
	type Alias RemoteStartTransactionReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RemoteStartTransactionReq(a)
	return s.Validate()
}

func (s RemoteStartTransactionReq) Validate() error {
	if s.ConnectorID != nil && *s.ConnectorID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be > 0")
	}

	if s.IDTag == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idTag", "required field is missing")
	}

	if err := s.IDTag.Validate(); err != nil {
		return ocpp.WrapField("idTag", err)
	}

	if s.ChargingProfile != nil {
		if err := s.ChargingProfile.Validate(); err != nil {
			return ocpp.WrapField("chargingProfile", err)
		}

		if s.ChargingProfile.ChargingProfilePurpose != ChargingProfilePurposeTypeTxProfile {
			return ocpp.NewError(
				ocpp.ErrPropertyConstraintViolation,
				"chargingProfile.chargingProfilePurpose",
				"must be TxProfile",
			)
		}
	}

	return nil
}

// RemoteStartTransactionConf (6.34)
//
// This contains the field definitions of the RemoteStartTransaction.conf PDU sent from Charge Point to Central
// System. See also Remote Start Transaction
type RemoteStartTransactionConf struct {
	// Required. Status indicating whether Charge Point accepts the request to start a transaction.
	Status RemoteStartStopStatus `json:"status"`
}

func (s *RemoteStartTransactionConf) UnmarshalJSON(data []byte) error {
	type Alias RemoteStartTransactionConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RemoteStartTransactionConf(a)
	return s.Validate()
}

func (s RemoteStartTransactionConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// RemoteStopTransactionReq (6.35)
//
// This contains the field definitions of the RemoteStopTransaction.req PDU sent to Charge Point by Central
// System. See also Remote Stop Transaction
type RemoteStopTransactionReq struct {
	// Required. The identifier of the transaction which Charge Point is requested to stop.
	TransactionID int32 `json:"transactionId"`
}

func (s *RemoteStopTransactionReq) UnmarshalJSON(data []byte) error {
	type Alias RemoteStopTransactionReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RemoteStopTransactionReq(a)
	return s.Validate()
}

func (s RemoteStopTransactionReq) Validate() error {
	if s.TransactionID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "transactionId", "must be > 0")
	}

	return nil
}

// RemoteStopTransactionConf (6.36)
//
// This contains the field definitions of the RemoteStopTransaction.conf PDU sent from Charge Point to Central
// System. See also Remote Stop Transaction
type RemoteStopTransactionConf struct {
	// Required. Status indicating whether Charge Point accepts the request to stop a transaction.
	Status RemoteStartStopStatus `json:"status"`
}

func (s *RemoteStopTransactionConf) UnmarshalJSON(data []byte) error {
	type Alias RemoteStopTransactionConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RemoteStopTransactionConf(a)
	return s.Validate()
}

func (s RemoteStopTransactionConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// ReserveNowReq (6.37)
//
// This contains the field definition of the ReserveNow.req PDU sent by the Central System to the Charge Point. See
// also Reserve Now
type ReserveNowReq struct {
	// Required. This contains the id of the connector to be reserved. A value of 0
	// means that the reservation is not for a specific connector.
	ConnectorID int32 `json:"connectorId"`
	// Required. This contains the date and time when the reservation ends.
	ExpiryDate time.Time `json:"expiryDate"`
	// Required. The identifier for which the Charge Point has to reserve a connector.
	IDTag IDToken `json:"idTag"`
	// Optional. The parent idTag.
	ParentIDTag *IDToken `json:"parentIdTag,omitempty"`
	// Required. Unique id for this reservation.
	ReservationID int32 `json:"reservationId"`
}

func (s *ReserveNowReq) UnmarshalJSON(data []byte) error {
	type Alias ReserveNowReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReserveNowReq(a)
	return s.Validate()
}

func (s ReserveNowReq) Validate() error {
	if s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if s.ExpiryDate.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "expiryDate", "required field is missing")
	}

	if s.IDTag == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idTag", "required field is missing")
	}

	if err := s.IDTag.Validate(); err != nil {
		return ocpp.WrapField("idTag", err)
	}

	if s.ParentIDTag != nil {
		if err := s.ParentIDTag.Validate(); err != nil {
			return ocpp.WrapField("parentIdTag", err)
		}
	}

	if s.ReservationID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "reservationId", "must be > 0")
	}

	return nil
}

// ReserveNowConf (6.38)
//
// This contains the field definition of the ReserveNow.conf PDU sent by the Charge Point to the Central System in
// response to a ReserveNow.req PDU. See also Reserve Now
type ReserveNowConf struct {
	// Required. This indicates the success or failure of the reservation.
	Status ReservationStatus `json:"status"`
}

func (s *ReserveNowConf) UnmarshalJSON(data []byte) error {
	type Alias ReserveNowConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReserveNowConf(a)
	return s.Validate()
}

func (s ReserveNowConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// ResetReq (6.39)
//
// This contains the field definition of the Reset.req PDU sent by the Central System to the Charge Point. See also
// Reset
type ResetReq struct {
	// Required. This contains the type of reset that the Charge Point should perform.
	Type ResetType `json:"type"`
}

func (s *ResetReq) UnmarshalJSON(data []byte) error {
	type Alias ResetReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ResetReq(a)
	return s.Validate()
}

func (s ResetReq) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	return nil
}

// ResetConf (6.40)
//
// This contains the field definition of the Reset.conf PDU sent by the Charge Point to the Central System in
// response to a Reset.req PDU. See also Reset
type ResetConf struct {
	// Required. This indicates whether the Charge Point is able to perform the reset.
	Status ResetStatus `json:"status"`
}

func (s *ResetConf) UnmarshalJSON(data []byte) error {
	type Alias ResetConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ResetConf(a)
	return s.Validate()
}

func (s ResetConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// SendLocalListReq (6.41)
//
// This contains the field definition of the SendLocalList.req PDU sent by the Central System to the Charge Point.
//
// If no (empty) localAuthorizationList is given and the updateType is Full, all identifications are removed from the
// list. Requesting a Differential update without (empty) localAuthorizationList will have no effect on the list. All
// idTags in the localAuthorizationList MUST be unique, no duplicate values are allowed. See also Send Local List
type SendLocalListReq struct {
	// Required. In case of a full update this is the version number of the
	// full list. In case of a differential update it is the version number of
	// the list after the update has been applied.
	ListVersion int32 `json:"listVersion"`
	// Optional. In case of a full update this contains the list of values
	// that form the new local authorization list. In case of a differential
	// update it contains the changes to be applied to the local
	// authorization list in the Charge Point. Maximum number of
	// AuthorizationData elements is available in the configuration key:
	// SendLocalListMaxLength
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty"`
	// Required. This contains the type of update (full or differential) of
	// this request.
	UpdateType UpdateType `json:"updateType"`
}

func (s *SendLocalListReq) UnmarshalJSON(data []byte) error {
	type Alias SendLocalListReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SendLocalListReq(a)
	return s.Validate()
}

func (s SendLocalListReq) Validate() error {
	seen := make(map[string]struct{}, len(s.LocalAuthorizationList))

	for i, v := range s.LocalAuthorizationList {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("localAuthorizationList[%d]", i), err)
		}

		key := strings.ToLower(v.IDTag.String())
		if _, dup := seen[key]; dup {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, fmt.Sprintf("localAuthorizationList[%d].idTag", i), "must be unique")
		}

		seen[key] = struct{}{}
	}

	if s.UpdateType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "updateType", "required field is missing")
	}

	return nil
}

// SendLocalListConf (6.42)
//
// This contains the field definition of the SendLocalList.conf PDU sent by the Charge Point to the Central System in
// response to a SendLocalList.req PDU. See also Send Local List
type SendLocalListConf struct {
	// Required. This indicates whether the Charge Point has successfully received and
	// applied the update of the local authorization list.
	Status UpdateStatus `json:"status"`
}

func (s *SendLocalListConf) UnmarshalJSON(data []byte) error {
	type Alias SendLocalListConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SendLocalListConf(a)
	return s.Validate()
}

func (s SendLocalListConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// SetChargingProfileReq (6.43)
//
// This contains the field definition of the SetChargingProfile.req PDU sent by the Central System to the Charge
// Point.
//
// The Central System uses this message to send charging profiles to a Charge Point. See also Set Charging Profile
type SetChargingProfileReq struct {
	// Required. The connector to which the charging profile applies. If connectorId = 0,
	// the message contains an overall limit for the Charge Point.
	ConnectorID int32 `json:"connectorId"`
	// Required. The charging profile to be set at the Charge Point.
	CsChargingProfiles ChargingProfile `json:"csChargingProfiles"`
}

func (s *SetChargingProfileReq) UnmarshalJSON(data []byte) error {
	type Alias SetChargingProfileReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetChargingProfileReq(a)
	return s.Validate()
}

func (s SetChargingProfileReq) Validate() error {
	if s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if err := s.CsChargingProfiles.Validate(); err != nil {
		return ocpp.WrapField("csChargingProfiles", err)
	}

	if s.CsChargingProfiles.ChargingProfilePurpose == ChargingProfilePurposeTypeChargePointMaxProfile && s.ConnectorID != 0 {
		return ocpp.NewError(
			ocpp.ErrPropertyConstraintViolation,
			"connectorId",
			"must be 0 when csChargingProfiles.chargingProfilePurpose is ChargePointMaxProfile",
		)
	}

	return nil
}

// SetChargingProfileConf (6.44)
//
// This contains the field definition of the SetChargingProfile.conf PDU sent by the Charge Point to the Central
// System in response to a SetChargingProfile.req PDU. See also Set Charging Profile
type SetChargingProfileConf struct {
	// Required. Returns whether the Charge Point has been able to process the
	// message successfully. This does not guarantee the schedule will be followed to
	// the letter. There might be other constraints the Charge Point may need to take
	// into account.
	Status ChargingProfileStatus `json:"status"`
}

func (s *SetChargingProfileConf) UnmarshalJSON(data []byte) error {
	type Alias SetChargingProfileConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetChargingProfileConf(a)
	return s.Validate()
}

func (s SetChargingProfileConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// StartTransactionReq (6.45)
//
// This section contains the field definition of the StartTransaction.req PDU sent by the Charge Point to the Central
// System. See also Start Transaction
type StartTransactionReq struct {
	// Required. This identifies which connector of the Charge Point is used.
	ConnectorID int32 `json:"connectorId"`
	// Required. This contains the identifier for which a transaction has to be started.
	IDTag IDToken `json:"idTag"`
	// Required. This contains the meter value in Wh for the connector at start of the
	// transaction.
	MeterStart int32 `json:"meterStart"`
	// Optional. This contains the id of the reservation that terminates as a result of
	// this transaction.
	ReservationID *int32 `json:"reservationId,omitempty"`
	// Required. This contains the date and time on which the transaction is started.
	Timestamp time.Time `json:"timestamp"`
}

func (s *StartTransactionReq) UnmarshalJSON(data []byte) error {
	type Alias StartTransactionReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StartTransactionReq(a)
	return s.Validate()
}

func (s StartTransactionReq) Validate() error {
	if s.ConnectorID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be > 0")
	}

	if s.IDTag == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idTag", "required field is missing")
	}

	if err := s.IDTag.Validate(); err != nil {
		return ocpp.WrapField("idTag", err)
	}

	if s.MeterStart < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "meterStart", "must be >= 0")
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	return nil
}

// StartTransactionConf (6.46)
//
// This contains the field definition of the StartTransaction.conf PDU sent by the Central System to the Charge Point
// in response to a StartTransaction.req PDU. See also Start Transaction
type StartTransactionConf struct {
	// Required. This contains information about authorization status, expiry and
	// parent id.
	IDTagInfo IDTagInfo `json:"idTagInfo"`
	// Required. This contains the transaction id supplied by the Central System.
	TransactionID int32 `json:"transactionId"`
}

func (s *StartTransactionConf) UnmarshalJSON(data []byte) error {
	type Alias StartTransactionConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StartTransactionConf(a)
	return s.Validate()
}

func (s StartTransactionConf) Validate() error {
	if err := s.IDTagInfo.Validate(); err != nil {
		return ocpp.WrapField("idTagInfo", err)
	}

	if s.TransactionID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "transactionId", "must be > 0")
	}

	return nil
}

// StatusNotificationReq (6.47)
//
// This contains the field definition of the StatusNotification.req PDU sent by the Charge Point to the Central
// System. See also Status Notification
type StatusNotificationReq struct {
	// Required. The id of the connector for which the status is reported.
	// Id '0' (zero) is used if the status is for the Charge Point main
	// controller.
	ConnectorID int32 `json:"connectorId"`
	// Required. This contains the error code reported by the Charge
	// Point.
	ErrorCode ChargePointErrorCode `json:"errorCode"`
	// Optional. Additional free format information related to the error.
	Info *CiString50Type `json:"info,omitempty"`
	// Required. This contains the current status of the Charge Point.
	Status ChargePointStatus `json:"status"`
	// Optional. The time for which the status is reported. If absent time
	// of receipt of the message will be assumed.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Optional. This identifies the vendor-specific implementation.
	VendorID *CiString255Type `json:"vendorId,omitempty"`
	// Optional. This contains the vendor-specific error code.
	VendorErrorCode *CiString50Type `json:"vendorErrorCode,omitempty"`
}

func (s *StatusNotificationReq) UnmarshalJSON(data []byte) error {
	type Alias StatusNotificationReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StatusNotificationReq(a)
	return s.Validate()
}

func (s StatusNotificationReq) Validate() error {
	if s.ConnectorID < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be >= 0")
	}

	if s.ErrorCode == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "errorCode", "required field is missing")
	}

	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return ocpp.WrapField("info", err)
		}
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.VendorID != nil {
		if err := s.VendorID.Validate(); err != nil {
			return ocpp.WrapField("vendorId", err)
		}
	}

	if s.VendorErrorCode != nil {
		if err := s.VendorErrorCode.Validate(); err != nil {
			return ocpp.WrapField("vendorErrorCode", err)
		}
	}

	return nil
}

// StatusNotificationConf (6.48)
//
// This contains the field definition of the StatusNotification.conf PDU sent by the Central System to the Charge
// Point in response to an StatusNotification.req PDU. See also Status Notification
//
// No fields are defined.
type StatusNotificationConf struct{}

func (s *StatusNotificationConf) UnmarshalJSON(data []byte) error {
	type Alias StatusNotificationConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StatusNotificationConf(a)
	return s.Validate()
}

func (s StatusNotificationConf) Validate() error {
	return nil
}

// StopTransactionReq (6.49)
//
// This contains the field definition of the StopTransaction.req PDU sent by the Charge Point to the Central System.
// See also Stop Transaction
type StopTransactionReq struct {
	// Optional. This contains the identifier which requested to stop the charging. It is
	// optional because a Charge Point may terminate charging without the presence
	// of an idTag, e.g. in case of a reset. A Charge Point SHALL send the idTag if known.
	IDTag *IDToken `json:"idTag,omitempty"`
	// Required. This contains the meter value in Wh for the connector at end of the
	// transaction.
	MeterStop int32 `json:"meterStop"`
	// Required. This contains the date and time on which the transaction is stopped.
	Timestamp time.Time `json:"timestamp"`
	// Required. This contains the transaction-id as received by the
	// StartTransaction.conf.
	TransactionID int32 `json:"transactionId"`
	// Optional. This contains the reason why the transaction was stopped. MAY only
	// be omitted when the Reason is "Local".
	Reason *Reason `json:"reason,omitempty"`
	// Optional. This contains transaction usage details relevant for billing purposes.
	TransactionData []MeterValue `json:"transactionData,omitempty"`
}

func (s *StopTransactionReq) UnmarshalJSON(data []byte) error {
	type Alias StopTransactionReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StopTransactionReq(a)
	return s.Validate()
}

func (s StopTransactionReq) Validate() error {
	if s.IDTag != nil {
		if err := s.IDTag.Validate(); err != nil {
			return ocpp.WrapField("idTag", err)
		}
	}

	if s.MeterStop < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "meterStop", "must be >= 0")
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.TransactionID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "transactionId", "must be > 0")
	}

	for i, v := range s.TransactionData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("transactionData[%d]", i), err)
		}
	}

	return nil
}

// StopTransactionConf (6.50)
//
// This contains the field definition of the StopTransaction.conf PDU sent by the Central System to the Charge Point
// in response to a StopTransaction.req PDU. See also Stop Transaction
type StopTransactionConf struct {
	// Optional. This contains information about authorization status, expiry and
	// parent id. It is optional, because a transaction may have been stopped without
	// an identifier.
	IDTagInfo *IDTagInfo `json:"idTagInfo,omitempty"`
}

func (s *StopTransactionConf) UnmarshalJSON(data []byte) error {
	type Alias StopTransactionConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StopTransactionConf(a)
	return s.Validate()
}

func (s StopTransactionConf) Validate() error {
	if s.IDTagInfo != nil {
		if err := s.IDTagInfo.Validate(); err != nil {
			return ocpp.WrapField("idTagInfo", err)
		}
	}

	return nil
}

// TriggerMessageReq (6.51)
//
// This contains the field definition of the TriggerMessage.req PDU sent by the Central System to the Charge Point.
// See also Trigger Message
type TriggerMessageReq struct {
	// Required.
	RequestedMessage MessageTrigger `json:"requestedMessage"`
	// Optional. Only filled in when request applies to a specific connector.
	ConnectorID *int32 `json:"connectorId,omitempty"`
}

func (s *TriggerMessageReq) UnmarshalJSON(data []byte) error {
	type Alias TriggerMessageReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TriggerMessageReq(a)
	return s.Validate()
}

func (s TriggerMessageReq) Validate() error {
	if s.RequestedMessage == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "requestedMessage", "required field is missing")
	}

	if s.ConnectorID != nil && *s.ConnectorID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be > 0")
	}

	return nil
}

// TriggerMessageConf (6.52)
//
// This contains the field definition of the TriggerMessage.conf PDU sent by the Charge Point to the Central System
// in response to a TriggerMessage.req PDU. See also Trigger Message
type TriggerMessageConf struct {
	// Required. Indicates whether the Charge Point will send the requested
	// notification or not.
	Status TriggerMessageStatus `json:"status"`
}

func (s *TriggerMessageConf) UnmarshalJSON(data []byte) error {
	type Alias TriggerMessageConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TriggerMessageConf(a)
	return s.Validate()
}

func (s TriggerMessageConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// UnlockConnectorReq (6.53)
//
// This contains the field definition of the UnlockConnector.req PDU sent by the Central System to the Charge
// Point. See also Unlock Connector
type UnlockConnectorReq struct {
	// Required. This contains the identifier of the connector to be unlocked.
	ConnectorID int32 `json:"connectorId"`
}

func (s *UnlockConnectorReq) UnmarshalJSON(data []byte) error {
	type Alias UnlockConnectorReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnlockConnectorReq(a)
	return s.Validate()
}

func (s UnlockConnectorReq) Validate() error {
	if s.ConnectorID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "connectorId", "must be > 0")
	}

	return nil
}

// UnlockConnectorConf (6.54)
//
// This contains the field definition of the UnlockConnector.conf PDU sent by the Charge Point to the Central
// System in response to an UnlockConnector.req PDU. See also Unlock Connector
type UnlockConnectorConf struct {
	// Required. This indicates whether the Charge Point has unlocked the connector.
	Status UnlockStatus `json:"status"`
}

func (s *UnlockConnectorConf) UnmarshalJSON(data []byte) error {
	type Alias UnlockConnectorConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnlockConnectorConf(a)
	return s.Validate()
}

func (s UnlockConnectorConf) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// UpdateFirmwareReq (6.55)
//
// This contains the field definition of the UpdateFirmware.req PDU sent by the Central System to the Charge Point.
// See also Update Firmware
type UpdateFirmwareReq struct {
	// Required. This contains a string containing a URI pointing to a location from
	// which to retrieve the firmware.
	Location string `json:"location"`
	// Optional. This specifies how many times Charge Point must try to download the
	// firmware before giving up. If this field is not present, it is left to Charge Point to
	// decide how many times it wants to retry.
	Retries *int32 `json:"retries,omitempty"`
	// Required. This contains the date and time after which the Charge Point is
	// allowed to retrieve the (new) firmware.
	RetrieveDate time.Time `json:"retrieveDate"`
	// Optional. The interval in seconds after which a retry may be attempted. If this
	// field is not present, it is left to Charge Point to decide how long to wait between
	// attempts.
	RetryInterval *int32 `json:"retryInterval,omitempty"`
}

func (s *UpdateFirmwareReq) UnmarshalJSON(data []byte) error {
	type Alias UpdateFirmwareReq
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UpdateFirmwareReq(a)
	return s.Validate()
}

func (s UpdateFirmwareReq) Validate() error {
	if s.Location == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "location", "required field is missing")
	}

	if s.Retries != nil && *s.Retries < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "retries", "must be >= 0")
	}

	if s.RetrieveDate.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "retrieveDate", "required field is missing")
	}

	if s.RetryInterval != nil && *s.RetryInterval < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "retryInterval", "must be >= 1")
	}

	return nil
}

// UpdateFirmwareConf (6.56)
//
// This contains the field definition of the UpdateFirmware.conf PDU sent by the Charge Point to the Central
// System in response to a UpdateFirmware.req PDU. See also Update Firmware
//
// No fields are defined.
type UpdateFirmwareConf struct{}

func (s *UpdateFirmwareConf) UnmarshalJSON(data []byte) error {
	type Alias UpdateFirmwareConf
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UpdateFirmwareConf(a)
	return s.Validate()
}

func (s UpdateFirmwareConf) Validate() error {
	return nil
}
