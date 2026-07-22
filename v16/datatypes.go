package v16

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// AuthorizationData (7.1)
//
// Elements that constitute an entry of a Local Authorization List update.
type AuthorizationData struct {
	// Required. The identifier to which this authorization applies.
	IDTag IDToken `json:"idTag"`
	// Optional. (Required when UpdateType is Full) This contains information about
	// authorization status, expiry and parent id. For a Differential update the following
	// applies: If this element is present, then this entry SHALL be added or updated in
	// the Local Authorization List. If this element is absent, than the entry for this
	// idtag in the Local Authorization List SHALL be deleted.
	IDTagInfo *IDTagInfo `json:"idTagInfo,omitempty"`
}

func (s *AuthorizationData) UnmarshalJSON(data []byte) error {
	type Alias AuthorizationData
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = AuthorizationData(raw)
	return s.Validate()
}

func (s AuthorizationData) Validate() error {
	if s.IDTag == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idTag", "required field is missing")
	}

	if err := s.IDTag.Validate(); err != nil {
		return ocpp.WrapField("idTag", err)
	}

	if s.IDTagInfo != nil {
		if err := s.IDTagInfo.Validate(); err != nil {
			return ocpp.WrapField("idTagInfo", err)
		}
	}

	return nil
}

// AuthorizationStatus (7.2)
//
// Status in a response to an Authorize.req.
type AuthorizationStatus string

const (
	// Identifier is allowed for charging.
	AuthorizationStatusAccepted AuthorizationStatus = "Accepted"
	// Identifier has been blocked. Not allowed for charging.
	AuthorizationStatusBlocked AuthorizationStatus = "Blocked"
	// Identifier has expired. Not allowed for charging.
	AuthorizationStatusExpired AuthorizationStatus = "Expired"
	// Identifier is unknown. Not allowed for charging.
	AuthorizationStatusInvalid AuthorizationStatus = "Invalid"
	// Identifier is already involved in another transaction and multiple transactions are not allowed. (Only relevant for a
	// StartTransaction.req.)
	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

func (s *AuthorizationStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch AuthorizationStatus(raw) {
	case AuthorizationStatusAccepted,
		AuthorizationStatusBlocked,
		AuthorizationStatusExpired,
		AuthorizationStatusInvalid,
		AuthorizationStatusConcurrentTx:
		*s = AuthorizationStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid AuthorizationStatus", raw),
	)
}

// AvailabilityStatus (7.3)
//
// Status returned in response to ChangeAvailability.req.
type AvailabilityStatus string

const (
	// Request has been accepted and will be executed.
	AvailabilityStatusAccepted AvailabilityStatus = "Accepted"
	// Request has not been accepted and will not be executed.
	AvailabilityStatusRejected AvailabilityStatus = "Rejected"
	// Request has been accepted and will be executed when transaction(s) in progress have finished.
	AvailabilityStatusScheduled AvailabilityStatus = "Scheduled"
)

func (s *AvailabilityStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch AvailabilityStatus(raw) {
	case AvailabilityStatusAccepted,
		AvailabilityStatusRejected,
		AvailabilityStatusScheduled:
		*s = AvailabilityStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid AvailabilityStatus", raw),
	)
}

// AvailabilityType (7.4)
//
// Requested availability change in ChangeAvailability.req.
type AvailabilityType string

const (
	// Charge point is not available for charging.
	AvailabilityTypeInoperative AvailabilityType = "Inoperative"
	// Charge point is available for charging.
	AvailabilityTypeOperative AvailabilityType = "Operative"
)

func (s *AvailabilityType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch AvailabilityType(raw) {
	case AvailabilityTypeInoperative,
		AvailabilityTypeOperative:
		*s = AvailabilityType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid AvailabilityType", raw),
	)
}

// CancelReservationStatus (7.5)
//
// Status in CancelReservation.conf.
type CancelReservationStatus string

const (
	// Reservation for the identifier has been cancelled.
	CancelReservationStatusAccepted CancelReservationStatus = "Accepted"
	// Reservation could not be cancelled, because there is no reservation active for the identifier.
	CancelReservationStatusRejected CancelReservationStatus = "Rejected"
)

func (s *CancelReservationStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CancelReservationStatus(raw) {
	case CancelReservationStatusAccepted,
		CancelReservationStatusRejected:
		*s = CancelReservationStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CancelReservationStatus", raw),
	)
}

// ChargePointErrorCode (7.6)
//
// Charge Point status reported in StatusNotification.req.
type ChargePointErrorCode string

const (
	// Failure to lock or unlock connector.
	ChargePointErrorCodeConnectorLockFailure ChargePointErrorCode = "ConnectorLockFailure"
	// Communication failure with the vehicle, might be Mode 3 or other communication protocol problem. This is
	// not a real error in the sense that the Charge Point doesn’t need to go to the faulted state. Instead, it should to the SuspendedEVSE state.
	ChargePointErrorCodeEVCommunicationError ChargePointErrorCode = "EVCommunicationError"
	// Ground fault circuit interrupter has been activated.
	ChargePointErrorCodeGroundFailure ChargePointErrorCode = "GroundFailure"
	// Temperature inside Charge Point is too high.
	ChargePointErrorCodeHighTemperature ChargePointErrorCode = "HighTemperature"
	// Error in internal hard- or software component.
	ChargePointErrorCodeInternalError ChargePointErrorCode = "InternalError"
	// The authorization information received from the Central System is in conflict with the LocalAuthorizationList.
	ChargePointErrorCodeLocalListConflict ChargePointErrorCode = "LocalListConflict"
	// No error to report.
	ChargePointErrorCodeNoError ChargePointErrorCode = "NoError"
	// Other type of error. More information in vendorErrorCode.
	ChargePointErrorCodeOtherError ChargePointErrorCode = "OtherError"
	// Over current protection device has tripped.
	ChargePointErrorCodeOverCurrentFailure ChargePointErrorCode = "OverCurrentFailure"
	// Voltage has risen above an acceptable level.
	ChargePointErrorCodeOverVoltage ChargePointErrorCode = "OverVoltage"
	// Failure to read electrical/energy/power meter.
	ChargePointErrorCodePowerMeterFailure ChargePointErrorCode = "PowerMeterFailure"
	// Failure to control power switch.
	ChargePointErrorCodePowerSwitchFailure ChargePointErrorCode = "PowerSwitchFailure"
	// Failure with idTag reader.
	ChargePointErrorCodeReaderFailure ChargePointErrorCode = "ReaderFailure"
	// Unable to perform a reset.
	ChargePointErrorCodeResetFailure ChargePointErrorCode = "ResetFailure"
	// Voltage has dropped below an acceptable level.
	ChargePointErrorCodeUnderVoltage ChargePointErrorCode = "UnderVoltage"
	// Wireless communication device reports a weak signal.
	ChargePointErrorCodeWeakSignal ChargePointErrorCode = "WeakSignal"
)

func (s *ChargePointErrorCode) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargePointErrorCode(raw) {
	case ChargePointErrorCodeConnectorLockFailure,
		ChargePointErrorCodeEVCommunicationError,
		ChargePointErrorCodeGroundFailure,
		ChargePointErrorCodeHighTemperature,
		ChargePointErrorCodeInternalError,
		ChargePointErrorCodeLocalListConflict,
		ChargePointErrorCodeNoError,
		ChargePointErrorCodeOtherError,
		ChargePointErrorCodeOverCurrentFailure,
		ChargePointErrorCodeOverVoltage,
		ChargePointErrorCodePowerMeterFailure,
		ChargePointErrorCodePowerSwitchFailure,
		ChargePointErrorCodeReaderFailure,
		ChargePointErrorCodeResetFailure,
		ChargePointErrorCodeUnderVoltage,
		ChargePointErrorCodeWeakSignal:
		*s = ChargePointErrorCode(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargePointErrorCode", raw),
	)
}

// ChargePointStatus (7.7)
//
// Status reported in StatusNotification.req. A status can be reported for the Charge Point main controller
// (connectorId = 0) or for a specific connector. Status for the Charge Point main controller is a subset of the
// enumeration: Available, Unavailable or Faulted.
// States considered Operative are: Available, Preparing, Charging, SuspendedEVSE, SuspendedEV, Finishing, Reserved.
// States considered Inoperative are: Unavailable, Faulted.
type ChargePointStatus string

const (
	// When a Connector becomes available for a new user (Operative)
	ChargePointStatusAvailable ChargePointStatus = "Available"
	// When a Connector becomes no longer available for a new user but there is no ongoing Transaction (yet). Typically a Connector
	// is in preparing state when a user presents a tag, inserts a cable or a vehicle occupies the parking bay
	// (Operative)
	ChargePointStatusPreparing ChargePointStatus = "Preparing"
	// When the contactor of a Connector closes, allowing the vehicle to charge
	// (Operative)
	ChargePointStatusCharging ChargePointStatus = "Charging"
	// When the EV is connected to the EVSE but the EVSE is not offering energy to the EV, e.g. due to a smart charging restriction,
	// local supply power constraints, or as the result of StartTransaction.conf indicating that charging is not allowed etc.
	// (Operative)
	ChargePointStatusSuspendedEVSE ChargePointStatus = "SuspendedEVSE"
	// When the EV is connected to the EVSE and the EVSE is offering energy but the EV is not taking any energy.
	// (Operative)
	ChargePointStatusSuspendedEV ChargePointStatus = "SuspendedEV"
	// When a Transaction has stopped at a Connector, but the Connector is not yet available for a new user, e.g. the cable has not
	// been removed or the vehicle has not left the parking bay
	// (Operative)
	ChargePointStatusFinishing ChargePointStatus = "Finishing"
	// When a Connector becomes reserved as a result of a Reserve Now command
	// (Operative)
	ChargePointStatusReserved ChargePointStatus = "Reserved"
	// When a Connector becomes unavailable as the result of a Change Availability command or an event upon which the Charge
	// Point transitions to unavailable at its discretion. Upon receipt of a Change Availability command, the status MAY change
	// immediately or the change MAY be scheduled. When scheduled, the Status Notification shall be send when the availability
	// change becomes effective
	// (Inoperative)
	ChargePointStatusUnavailable ChargePointStatus = "Unavailable"
	// When a Charge Point or connector has reported an error and is not available for energy delivery . (Inoperative).
	ChargePointStatusFaulted ChargePointStatus = "Faulted"
)

func (s *ChargePointStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargePointStatus(raw) {
	case ChargePointStatusAvailable,
		ChargePointStatusPreparing,
		ChargePointStatusCharging,
		ChargePointStatusSuspendedEVSE,
		ChargePointStatusSuspendedEV,
		ChargePointStatusFinishing,
		ChargePointStatusReserved,
		ChargePointStatusUnavailable,
		ChargePointStatusFaulted:
		*s = ChargePointStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargePointStatus", raw),
	)
}

// ChargingProfile (7.8)
//
// A ChargingProfile consists of a ChargingSchedule, describing the amount of power or current that can be
// delivered per time interval.
type ChargingProfile struct {
	// Required. Unique identifier for this profile.
	ChargingProfileID int32 `json:"chargingProfileId"`
	// Optional. Only valid if ChargingProfilePurpose is set to TxProfile,
	// the transactionId MAY be used to match the profile to a transaction.
	TransactionID *int32 `json:"transactionId,omitempty"`
	// Required. Value determining level in hierarchy stack of profiles.
	// Higher values have precedence over lower values. Lowest level is
	// 0.
	StackLevel int32 `json:"stackLevel"`
	// Required. Defines the purpose of the schedule transferred by this
	// message.
	ChargingProfilePurpose ChargingProfilePurposeType `json:"chargingProfilePurpose"`
	// Required. Indicates the kind of schedule.
	ChargingProfileKind ChargingProfileKindType `json:"chargingProfileKind"`
	// Optional. Indicates the start point of a recurrence.
	RecurrencyKind *RecurrencyKindType `json:"recurrencyKind,omitempty"`
	// Optional. Point in time at which the profile starts to be valid. If
	// absent, the profile is valid as soon as it is received by the Charge
	// Point.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Optional. Point in time at which the profile stops to be valid. If
	// absent, the profile is valid until it is replaced by another profile.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// Required. Contains limits for the available power or current over
	// time.
	ChargingSchedule ChargingSchedule `json:"chargingSchedule"`
}

func (s *ChargingProfile) UnmarshalJSON(data []byte) error {
	type Alias ChargingProfile
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingProfile(a)
	return s.Validate()
}

func (s ChargingProfile) Validate() error {
	if s.StackLevel < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "stackLevel", "must be >= 0")
	}

	if s.ChargingProfilePurpose == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingProfilePurpose", "required field is missing")
	}

	if s.ChargingProfileKind == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingProfileKind", "required field is missing")
	}

	if s.RecurrencyKind != nil && *s.RecurrencyKind == "" {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "recurrencyKind", "should not be empty")
	}

	if s.TransactionID != nil && s.ChargingProfilePurpose != ChargingProfilePurposeTypeTxProfile {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "is only valid for TxProfile")
	}

	if err := s.ChargingSchedule.Validate(); err != nil {
		return ocpp.WrapField("chargingSchedule", err)
	}

	return nil
}

// ChargingProfileKindType (7.9)
//
// Kind of charging profile, as used in: ChargingProfile.
type ChargingProfileKindType string

const (
	// Schedule periods are relative to a fixed point in time defined in the schedule.
	ChargingProfileKindTypeAbsolute ChargingProfileKindType = "Absolute"
	// The schedule restarts periodically at the first schedule period.
	ChargingProfileKindTypeRecurring ChargingProfileKindType = "Recurring"
	// Schedule periods are relative to a situation-specific start point (such as the start of a Transaction) that is determined by the
	// charge point.
	ChargingProfileKindTypeRelative ChargingProfileKindType = "Relative"
)

func (s *ChargingProfileKindType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfileKindType(raw) {
	case ChargingProfileKindTypeAbsolute,
		ChargingProfileKindTypeRecurring,
		ChargingProfileKindTypeRelative:
		*s = ChargingProfileKindType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfileKindType", raw),
	)
}

// ChargingProfilePurposeType (7.10)
//
// Purpose of the charging profile, as used in: ChargingProfile.
type ChargingProfilePurposeType string

const (
	// Configuration for the maximum power or current available for an entire Charge Point.
	ChargingProfilePurposeTypeChargePointMaxProfile ChargingProfilePurposeType = "ChargePointMaxProfile"
	// Default profile that can be configured in the Charge Point. When a new transaction is started, this profile
	// SHALL be used, unless it was a transaction that was started by a RemoteStartTransaction.req with a
	// ChargeProfile that is accepted by the Charge Point.
	ChargingProfilePurposeTypeTxDefaultProfile ChargingProfilePurposeType = "TxDefaultProfile"
	// Profile with constraints to be imposed by the Charge Point on the current transaction, or on a new transaction
	// when this is started via a RemoteStartTransaction.req with a ChargeProfile. A profile with this purpose SHALL
	// cease to be valid when the transaction terminates.
	ChargingProfilePurposeTypeTxProfile ChargingProfilePurposeType = "TxProfile"
)

func (s *ChargingProfilePurposeType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfilePurposeType(raw) {
	case ChargingProfilePurposeTypeChargePointMaxProfile,
		ChargingProfilePurposeTypeTxDefaultProfile,
		ChargingProfilePurposeTypeTxProfile:
		*s = ChargingProfilePurposeType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfilePurposeType", raw),
	)
}

// ChargingProfileStatus (7.11)
//
// Status returned in response to SetChargingProfile.req.
type ChargingProfileStatus string

const (
	// Request has been accepted and will be executed.
	ChargingProfileStatusAccepted ChargingProfileStatus = "Accepted"
	// Request has not been accepted and will not be executed.
	ChargingProfileStatusRejected ChargingProfileStatus = "Rejected"
	// Charge Point indicates that the request is not supported.
	ChargingProfileStatusNotSupported ChargingProfileStatus = "NotSupported"
)

func (s *ChargingProfileStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfileStatus(raw) {
	case ChargingProfileStatusAccepted,
		ChargingProfileStatusRejected,
		ChargingProfileStatusNotSupported:
		*s = ChargingProfileStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfileStatus", raw),
	)
}

// ChargingRateUnitType (7.12)
//
// Unit in which a charging schedule is defined, as used in: GetCompositeSchedule.req and ChargingSchedule
type ChargingRateUnitType string

const (
	// Watts (power).
	// This is the TOTAL allowed charging power.
	// If used for AC Charging, the phase current should be calculated via: Current per phase = Power / (Line Voltage * Number of
	// Phases). The "Line Voltage" used in the calculation is not the measured voltage, but the set voltage for the area (hence, 230 of
	// 110 volt). The "Number of Phases" is the numberPhases from the ChargingSchedulePeriod.
	// It is usually more convenient to use this for DC charging.
	// Note that if numberPhases in a ChargingSchedulePeriod is absent, 3 SHALL be assumed.
	ChargingRateUnitTypeW ChargingRateUnitType = "W"
	// Amperes (current).
	// The amount of Ampere per phase, not the sum of all phases.
	// It is usually more convenient to use this for AC charging.
	ChargingRateUnitTypeA ChargingRateUnitType = "A"
)

func (s *ChargingRateUnitType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingRateUnitType(raw) {
	case ChargingRateUnitTypeW,
		ChargingRateUnitTypeA:
		*s = ChargingRateUnitType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingRateUnitType", raw),
	)
}

// ChargingSchedule (7.13)
//
// Charging schedule structure defines a list of charging periods, as used in: GetCompositeSchedule.conf and
// ChargingProfile.
type ChargingSchedule struct {
	// Optional. Duration of the charging schedule in seconds. If the
	// duration is left empty, the last period will continue indefinitely or
	// until end of the transaction in case startSchedule is absent.
	Duration *int32 `json:"duration,omitempty"`
	// Optional. Starting point of an absolute schedule. If absent the
	// schedule will be relative to start of charging.
	StartSchedule *time.Time `json:"startSchedule,omitempty"`
	// Required. The unit of measure Limit is expressed in.
	ChargingRateUnit ChargingRateUnitType `json:"chargingRateUnit"`
	// Required. List of ChargingSchedulePeriod elements defining
	// maximum power or current usage over time. The startPeriod of
	// the first ChargingSchedulePeriod SHALL always be 0.
	ChargingSchedulePeriod []ChargingSchedulePeriod `json:"chargingSchedulePeriod"`
	// Optional. Minimum charging rate supported by the electric
	// vehicle. The unit of measure is defined by the chargingRateUnit.
	// This parameter is intended to be used by a local smart charging
	// algorithm to optimize the power allocation for in the case a
	// charging process is inefficient at lower charging rates. Accepts at
	// most one digit fraction (e.g. 8.1)
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`
}

func (s *ChargingSchedule) UnmarshalJSON(data []byte) error {
	type Alias ChargingSchedule
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = ChargingSchedule(raw)
	return s.Validate()
}

func (s ChargingSchedule) Validate() error {
	if s.Duration != nil && *s.Duration < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "duration", "must be > 0")
	}

	if s.ChargingRateUnit == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingRateUnit", "required field is missing")
	}

	if len(s.ChargingSchedulePeriod) == 0 {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingSchedulePeriod", "must not be an empty array")
	}

	if s.ChargingSchedulePeriod[0].StartPeriod != 0 {
		return ocpp.NewError(
			ocpp.ErrPropertyConstraintViolation,
			"chargingSchedulePeriod[0]",
			"startPeriod must be 0",
		)
	}

	for i, p := range s.ChargingSchedulePeriod {
		if err := p.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingSchedulePeriod[%d]", i), err)
		}
	}

	if s.MinChargingRate != nil {
		if *s.MinChargingRate < 0 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "minChargingRate", "must be >= 0")
		}

		f := strconv.FormatFloat(*s.MinChargingRate, 'f', -1, 64)
		if strings.IndexByte(f, '.') != -1 && len(f)-strings.IndexByte(f, '.')-1 > 1 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "minChargingRate", "must have at most one digit fraction")
		}
	}

	return nil
}

// ChargingSchedulePeriod (7.14)
//
// Charging schedule period structure defines a time period in a charging schedule, as used in: ChargingSchedule.
type ChargingSchedulePeriod struct {
	// Required. Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	StartPeriod int32 `json:"startPeriod"`
	// Required. Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes or Watts. Accepts at most one digit
	// fraction (e.g. 8.1).
	Limit float64 `json:"limit"`
	// Optional. The number of phases that can be used for charging. If a number of
	// phases is needed, numberPhases=3 will be assumed unless another number is
	// given.
	NumberPhases *int32 `json:"numberPhases,omitempty"`
}

func (s *ChargingSchedulePeriod) UnmarshalJSON(data []byte) error {
	type Alias ChargingSchedulePeriod
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = ChargingSchedulePeriod(raw)
	return s.Validate()
}

func (s ChargingSchedulePeriod) Validate() error {
	if s.StartPeriod < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "startPeriod", "must be >= 0")
	}

	if s.Limit < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "limit", "must be >= 0")
	}

	f := strconv.FormatFloat(s.Limit, 'f', -1, 64)
	if strings.IndexByte(f, '.') != -1 && len(f)-strings.IndexByte(f, '.')-1 > 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "limit", "must have at most one digit fraction")
	}

	if s.NumberPhases != nil {
		switch *s.NumberPhases {
		case 1, 2, 3:
		default:
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "numberPhases", "must be 1, 2 or 3")
		}
	}

	return nil
}

// validateCiString checks that s is no longer than maxLen and contains only
// printable ASCII characters (0x20-0x7E), as required by the OCPP 1.6 CiString*
// types.
func validateCiString(s string, maxLen int, typeName string) error {
	if len(s) > maxLen {
		return ocpp.NewError(
			ocpp.ErrPropertyConstraintViolation,
			"",
			fmt.Sprintf("%s exceeds max length of %d", typeName, maxLen),
		)
	}

	for _, r := range s {
		if r < 0x20 || r > 0x7E {
			return ocpp.NewError(
				ocpp.ErrPropertyConstraintViolation,
				"",
				fmt.Sprintf("%s contains non-printable ASCII characters", typeName),
			)
		}
	}

	return nil
}

// CiString20Type (7.15)
//
// Generic case insensitive string of 20 characters.
type CiString20Type string

func NewCiString20Type(s string) (CiString20Type, error) {
	c := CiString20Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s CiString20Type) Validate() error {
	return validateCiString(string(s), 20, "CiString20Type")
}

func (s CiString20Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s CiString20Type) String() string {
	return string(s)
}

func (s *CiString20Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewCiString20Type(raw)

	if err != nil {
		return err
	}

	*s = c
	return nil
}

// CiString25Type (7.16)
//
// Generic case insensitive string of 25 characters.
type CiString25Type string

func NewCiString25Type(s string) (CiString25Type, error) {
	c := CiString25Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s CiString25Type) Validate() error {
	return validateCiString(string(s), 25, "CiString25Type")
}

func (s CiString25Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s CiString25Type) String() string {
	return string(s)
}

func (s *CiString25Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewCiString25Type(raw)

	if err != nil {
		return err
	}

	*s = c
	return nil
}

// CiString50Type (7.17)
//
// Generic case insensitive string of 50 characters.
type CiString50Type string

func NewCiString50Type(s string) (CiString50Type, error) {
	c := CiString50Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s CiString50Type) Validate() error {
	return validateCiString(string(s), 50, "CiString50Type")
}

func (s CiString50Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s CiString50Type) String() string {
	return string(s)
}

func (s *CiString50Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewCiString50Type(raw)

	if err != nil {
		return err
	}

	*s = c
	return nil
}

// CiString255Type (7.18)
//
// Generic case insensitive string of 255 characters.
type CiString255Type string

func NewCiString255Type(s string) (CiString255Type, error) {
	c := CiString255Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s CiString255Type) Validate() error {
	return validateCiString(string(s), 255, "CiString255Type")
}

func (s CiString255Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s CiString255Type) String() string {
	return string(s)
}

func (s *CiString255Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewCiString255Type(raw)

	if err != nil {
		return err
	}

	*s = c
	return nil
}

// CiString500Type (7.19)
//
// Generic case insensitive string of 500 characters.
type CiString500Type string

func NewCiString500Type(s string) (CiString500Type, error) {
	c := CiString500Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s CiString500Type) Validate() error {
	return validateCiString(string(s), 500, "CiString500Type")
}

func (s CiString500Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s CiString500Type) String() string {
	return string(s)
}

func (s *CiString500Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewCiString500Type(raw)

	if err != nil {
		return err
	}

	*s = c
	return nil
}

// ClearCacheStatus (7.20)
//
// Status returned in response to ClearCache.req.
type ClearCacheStatus string

const (
	// Command has been executed.
	ClearCacheStatusAccepted ClearCacheStatus = "Accepted"
	// Command has not been executed.
	ClearCacheStatusRejected ClearCacheStatus = "Rejected"
)

func (s *ClearCacheStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearCacheStatus(raw) {
	case ClearCacheStatusAccepted,
		ClearCacheStatusRejected:
		*s = ClearCacheStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearCacheStatus", raw),
	)
}

// ClearChargingProfileStatus (7.21)
//
// Status returned in response to ClearChargingProfile.req.
type ClearChargingProfileStatus string

const (
	// Request has been accepted and will be executed.
	ClearChargingProfileStatusAccepted ClearChargingProfileStatus = "Accepted"
	// No Charging Profile(s) were found matching the request.
	ClearChargingProfileStatusUnknown ClearChargingProfileStatus = "Unknown"
)

func (s *ClearChargingProfileStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearChargingProfileStatus(raw) {
	case ClearChargingProfileStatusAccepted,
		ClearChargingProfileStatusUnknown:
		*s = ClearChargingProfileStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearChargingProfileStatus", raw),
	)
}

// ConfigurationStatus (7.22)
//
// Status in ChangeConfiguration.conf.
type ConfigurationStatus string

const (
	// Configuration key is supported and setting has been changed.
	ConfigurationStatusAccepted ConfigurationStatus = "Accepted"
	// Configuration key is supported, but setting could not be changed.
	ConfigurationStatusRejected ConfigurationStatus = "Rejected"
	// Configuration key is supported and setting has been changed, but change will be available after reboot (Charge Point will not
	// reboot itself)
	ConfigurationStatusRebootRequired ConfigurationStatus = "RebootRequired"
	// Configuration key is not supported.
	ConfigurationStatusNotSupported ConfigurationStatus = "NotSupported"
)

func (s *ConfigurationStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ConfigurationStatus(raw) {
	case ConfigurationStatusAccepted,
		ConfigurationStatusRejected,
		ConfigurationStatusRebootRequired,
		ConfigurationStatusNotSupported:
		*s = ConfigurationStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ConfigurationStatus", raw),
	)
}

// DataTransferStatus (7.23)
//
// Status in DataTransfer.conf.
type DataTransferStatus string

const (
	// Message has been accepted and the contained request is accepted.
	DataTransferStatusAccepted DataTransferStatus = "Accepted"
	// Message has been accepted but the contained request is rejected.
	DataTransferStatusRejected DataTransferStatus = "Rejected"
	// Message could not be interpreted due to unknown messageId string.
	DataTransferStatusUnknownMessageID DataTransferStatus = "UnknownMessageId"
	// Message could not be interpreted due to unknown vendorId string.
	DataTransferStatusUnknownVendorID DataTransferStatus = "UnknownVendorId"
)

func (s *DataTransferStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DataTransferStatus(raw) {
	case DataTransferStatusAccepted,
		DataTransferStatusRejected,
		DataTransferStatusUnknownMessageID,
		DataTransferStatusUnknownVendorID:
		*s = DataTransferStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DataTransferStatus", raw),
	)
}

// DiagnosticsStatus (7.24)
//
// Status in DiagnosticsStatusNotification.req.
type DiagnosticsStatus string

const (
	// Charge Point is not performing diagnostics related tasks. Status Idle SHALL only be used as in a
	// DiagnosticsStatusNotification.req that was triggered by a TriggerMessage.req
	DiagnosticsStatusIdle DiagnosticsStatus = "Idle"
	// Diagnostics information has been uploaded.
	DiagnosticsStatusUploaded DiagnosticsStatus = "Uploaded"
	// Uploading of diagnostics failed.
	DiagnosticsStatusUploadFailed DiagnosticsStatus = "UploadFailed"
	// File is being uploaded.
	DiagnosticsStatusUploading DiagnosticsStatus = "Uploading"
)

func (s *DiagnosticsStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DiagnosticsStatus(raw) {
	case DiagnosticsStatusIdle,
		DiagnosticsStatusUploaded,
		DiagnosticsStatusUploadFailed,
		DiagnosticsStatusUploading:
		*s = DiagnosticsStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DiagnosticsStatus", raw),
	)
}

// FirmwareStatus (7.25)
//
// Status of a firmware download as reported in FirmwareStatusNotification.req.
type FirmwareStatus string

const (
	// New firmware has been downloaded by Charge Point.
	FirmwareStatusDownloaded FirmwareStatus = "Downloaded"
	// Charge point failed to download firmware.
	FirmwareStatusDownloadFailed FirmwareStatus = "DownloadFailed"
	// Firmware is being downloaded.
	FirmwareStatusDownloading FirmwareStatus = "Downloading"
	// Charge Point is not performing firmware update related tasks. Status Idle SHALL only be used as in a
	// FirmwareStatusNotification.req that was triggered by a TriggerMessage.req
	FirmwareStatusIdle FirmwareStatus = "Idle"
	// Installation of new firmware has failed.
	FirmwareStatusInstallationFailed FirmwareStatus = "InstallationFailed"
	// Firmware is being installed.
	FirmwareStatusInstalling FirmwareStatus = "Installing"
	// New firmware has successfully been installed in charge point.
	FirmwareStatusInstalled FirmwareStatus = "Installed"
)

func (s *FirmwareStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch FirmwareStatus(raw) {
	case FirmwareStatusDownloaded,
		FirmwareStatusDownloadFailed,
		FirmwareStatusDownloading,
		FirmwareStatusIdle,
		FirmwareStatusInstallationFailed,
		FirmwareStatusInstalling,
		FirmwareStatusInstalled:
		*s = FirmwareStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid FirmwareStatus", raw),
	)
}

// GetCompositeScheduleStatus (7.26)
//
// Status returned in response to GetCompositeSchedule.req.
type GetCompositeScheduleStatus string

const (
	// Request has been accepted and will be executed.
	GetCompositeScheduleStatusAccepted GetCompositeScheduleStatus = "Accepted"
	// Request has not been accepted and will not be executed.
	GetCompositeScheduleStatusRejected GetCompositeScheduleStatus = "Rejected"
)

func (s *GetCompositeScheduleStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetCompositeScheduleStatus(raw) {
	case GetCompositeScheduleStatusAccepted,
		GetCompositeScheduleStatusRejected:
		*s = GetCompositeScheduleStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetCompositeScheduleStatus", raw),
	)
}

// IDTagInfo (7.27)
//
// Contains status information about an identifier. It is returned in Authorize, Start Transaction and Stop
// Transaction responses.
//
// If expiryDate is not given, the status has no end date.
type IDTagInfo struct {
	// Optional. This contains the date at which idTag should be removed from the
	// Authorization Cache.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Optional. This contains the parent-identifier.
	ParentIDTag *IDToken `json:"parentIdTag,omitempty"`
	// Required. This contains whether the idTag has been accepted or not by the
	// Central System.
	Status AuthorizationStatus `json:"status"`
}

func (s *IDTagInfo) UnmarshalJSON(data []byte) error {
	type Alias IDTagInfo
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = IDTagInfo(raw)
	return s.Validate()
}

func (s IDTagInfo) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.ParentIDTag != nil {
		if *s.ParentIDTag == "" {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "parentIdTag", "should not be empty")
		}

		if err := s.ParentIDTag.Validate(); err != nil {
			return ocpp.WrapField("parentIdTag", err)
		}
	}

	return nil
}

// IDToken (7.28)
//
// Contains the identifier to use for authorization. It is a case insensitive string. In future releases this may become
// a complex type to support multiple forms of identifiers.
type IDToken = CiString20Type

// KeyValue (7.29)
//
// Contains information about a specific configuration key. It is returned in GetConfiguration.conf.
type KeyValue struct {
	// Required.
	Key CiString50Type `json:"key"`
	// Required. False if the value can be set with the ChangeConfiguration message.
	Readonly bool `json:"readonly"`
	// Optional. If key is known but not set, this field may be absent.
	Value *CiString500Type `json:"value,omitempty"`
}

func (s *KeyValue) UnmarshalJSON(data []byte) error {
	type Alias KeyValue
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = KeyValue(raw)
	return s.Validate()
}

func (s KeyValue) Validate() error {
	if s.Key == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "key", "required field is missing")
	}

	if err := s.Key.Validate(); err != nil {
		return ocpp.WrapField("key", err)
	}

	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return ocpp.WrapField("value", err)
		}
	}

	return nil
}

// Location (7.30)
//
// Allowable values of the optional "location" field of a value element in SampledValue.
type Location string

const (
	// Measurement inside body of Charge Point (e.g. Temperature)
	LocationBody Location = "Body"
	// Measurement taken from cable between EV and Charge Point
	LocationCable Location = "Cable"
	// Measurement taken by EV
	LocationEV Location = "EV"
	// Measurement at network (“grid”) inlet connection
	LocationInlet Location = "Inlet"
	// Measurement at a Connector. Default value
	LocationOutlet Location = "Outlet"
)

func (s *Location) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch Location(raw) {
	case LocationBody,
		LocationCable,
		LocationEV,
		LocationInlet,
		LocationOutlet:
		*s = Location(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid Location", raw),
	)
}

// Measurand (7.31)
//
// Allowable values of the optional "measurand" field of a Value element, as used in MeterValues.req and
// StopTransaction.req messages. Default value of "measurand" is always "Energy.Active.Import.Register"
type Measurand string

const (
	// Instantaneous current flow from EV
	MeasurandCurrentExport Measurand = "Current.Export"
	// Instantaneous current flow to EV
	MeasurandCurrentImport Measurand = "Current.Import"
	// Maximum current offered to EV
	MeasurandCurrentOffered Measurand = "Current.Offered"
	// Numerical value read from the "active electrical energy" (Wh or kWh) register of the (most authoritative)
	// electrical meter measuring energy exported (to the grid).
	MeasurandEnergyActiveExportRegister Measurand = "Energy.Active.Export.Register"
	// Numerical value read from the "active electrical energy" (Wh or kWh) register of the (most authoritative)
	// electrical meter measuring energy imported (from the grid supply).
	MeasurandEnergyActiveImportRegister Measurand = "Energy.Active.Import.Register"
	// Numerical value read from the "reactive electrical energy" (VARh or kVARh) register of the (most
	// authoritative) electrical meter measuring energy exported (to the grid).
	MeasurandEnergyReactiveExportRegister Measurand = "Energy.Reactive.Export.Register"
	// Numerical value read from the "reactive electrical energy" (VARh or kVARh) register of the (most
	// authoritative) electrical meter measuring energy imported (from the grid supply).
	MeasurandEnergyReactiveImportRegister Measurand = "Energy.Reactive.Import.Register"
	// Absolute amount of "active electrical energy" (Wh or kWh) exported (to the grid) during an associated time
	// "interval", specified by a Metervalues ReadingContext, and applicable interval duration configuration values
	// (in seconds) for "ClockAlignedDataInterval" and "MeterValueSampleInterval".
	MeasurandEnergyActiveExportInterval Measurand = "Energy.Active.Export.Interval"
	// Absolute amount of "active electrical energy" (Wh or kWh) imported (from the grid supply) during an
	// associated time "interval", specified by a Metervalues ReadingContext, and applicable interval duration
	// configuration values (in seconds) for "ClockAlignedDataInterval" and "MeterValueSampleInterval".
	MeasurandEnergyActiveImportInterval Measurand = "Energy.Active.Import.Interval"
	// Absolute amount of "reactive electrical energy" (VARh or kVARh) exported (to the grid) during an associated
	// time "interval", specified by a Metervalues ReadingContext, and applicable interval duration configuration
	// values (in seconds) for "ClockAlignedDataInterval" and "MeterValueSampleInterval".
	MeasurandEnergyReactiveExportInterval Measurand = "Energy.Reactive.Export.Interval"
	// Absolute amount of "reactive electrical energy" (VARh or kVARh) imported (from the grid supply) during an
	// associated time "interval", specified by a Metervalues ReadingContext, and applicable interval duration
	// configuration values (in seconds) for "ClockAlignedDataInterval" and "MeterValueSampleInterval".
	MeasurandEnergyReactiveImportInterval Measurand = "Energy.Reactive.Import.Interval"
	// Instantaneous reading of powerline frequency. NOTE: OCPP 1.6 does not have a UnitOfMeasure for
	// frequency, the UnitOfMeasure for any SampledValue with measurand: Frequency is Hertz.
	MeasurandFrequency Measurand = "Frequency"
	// Instantaneous active power exported by EV. (W or kW)
	MeasurandPowerActiveExport Measurand = "Power.Active.Export"
	// Instantaneous active power imported by EV. (W or kW)
	MeasurandPowerActiveImport Measurand = "Power.Active.Import"
	// Instantaneous power factor of total energy flow
	MeasurandPowerFactor Measurand = "Power.Factor"
	// Maximum power offered to EV
	MeasurandPowerOffered Measurand = "Power.Offered"
	// Instantaneous reactive power exported by EV. (var or kvar)
	MeasurandPowerReactiveExport Measurand = "Power.Reactive.Export"
	// Instantaneous reactive power imported by EV. (var or kvar)
	MeasurandPowerReactiveImport Measurand = "Power.Reactive.Import"
	// Fan speed in RPM
	MeasurandRPM Measurand = "RPM"
	// State of charge of charging vehicle in percentage
	MeasurandSoC Measurand = "SoC"
	// Temperature reading inside Charge Point.
	MeasurandTemperature Measurand = "Temperature"
	// Instantaneous AC RMS supply voltage
	MeasurandVoltage Measurand = "Voltage"
)

func (s *Measurand) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch Measurand(raw) {
	case MeasurandCurrentExport,
		MeasurandCurrentImport,
		MeasurandCurrentOffered,
		MeasurandEnergyActiveExportRegister,
		MeasurandEnergyActiveImportRegister,
		MeasurandEnergyReactiveExportRegister,
		MeasurandEnergyReactiveImportRegister,
		MeasurandEnergyActiveExportInterval,
		MeasurandEnergyActiveImportInterval,
		MeasurandEnergyReactiveExportInterval,
		MeasurandEnergyReactiveImportInterval,
		MeasurandFrequency,
		MeasurandPowerActiveExport,
		MeasurandPowerActiveImport,
		MeasurandPowerFactor,
		MeasurandPowerOffered,
		MeasurandPowerReactiveExport,
		MeasurandPowerReactiveImport,
		MeasurandRPM,
		MeasurandSoC,
		MeasurandTemperature,
		MeasurandVoltage:
		*s = Measurand(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid Measurand", raw),
	)
}

// MessageTrigger (7.32)
//
// Type of request to be triggered in a TriggerMessage.req.
type MessageTrigger string

const (
	// To trigger a BootNotification request
	MessageTriggerBootNotification MessageTrigger = "BootNotification"
	// To trigger a DiagnosticsStatusNotification request
	MessageTriggerDiagnosticsStatusNotification MessageTrigger = "DiagnosticsStatusNotification"
	// To trigger a FirmwareStatusNotification request
	MessageTriggerFirmwareStatusNotification MessageTrigger = "FirmwareStatusNotification"
	// To trigger a Heartbeat request
	MessageTriggerHeartbeat MessageTrigger = "Heartbeat"
	// To trigger a MeterValues request
	MessageTriggerMeterValues MessageTrigger = "MeterValues"
	// To trigger a StatusNotification request
	MessageTriggerStatusNotification MessageTrigger = "StatusNotification"
)

func (s *MessageTrigger) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MessageTrigger(raw) {
	case MessageTriggerBootNotification,
		MessageTriggerDiagnosticsStatusNotification,
		MessageTriggerFirmwareStatusNotification,
		MessageTriggerHeartbeat,
		MessageTriggerMeterValues,
		MessageTriggerStatusNotification:
		*s = MessageTrigger(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageTrigger", raw),
	)
}

// MeterValue (7.33)
//
// Collection of one or more sampled values in MeterValues.req and StopTransaction.req. All sampled values in a
// MeterValue are sampled at the same point in time.
type MeterValue struct {
	// Required. Timestamp for measured value(s).
	Timestamp time.Time `json:"timestamp"`
	// Required. One or more measured values
	SampledValue []SampledValue `json:"sampledValue"`
}

func (s *MeterValue) UnmarshalJSON(data []byte) error {
	type Alias MeterValue
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = MeterValue(raw)
	return s.Validate()
}

func (s MeterValue) Validate() error {
	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if len(s.SampledValue) == 0 {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "sampledValue", "must not be an empty array")
	}

	for i, v := range s.SampledValue {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("sampledValue[%d]", i), err)
		}
	}

	return nil
}

// Phase (7.34)
//
// Phase as used in SampledValue. Phase specifies how a measured value is to be interpreted. Please note that not
// all values of Phase are applicable to all Measurands.
type Phase string

const (
	// Measured on L1
	PhaseL1 Phase = "L1"
	// Measured on L2
	PhaseL2 Phase = "L2"
	// Measured on L3
	PhaseL3 Phase = "L3"
	// Measured on Neutral
	PhaseN Phase = "N"
	// Measured on L1 with respect to Neutral conductor
	PhaseL1N Phase = "L1-N"
	// Measured on L2 with respect to Neutral conductor
	PhaseL2N Phase = "L2-N"
	// Measured on L3 with respect to Neutral conductor
	PhaseL3N Phase = "L3-N"
	// Measured between L1 and L2
	PhaseL1L2 Phase = "L1-L2"
	// Measured between L2 and L3
	PhaseL2L3 Phase = "L2-L3"
	// Measured between L3 and L1
	PhaseL3L1 Phase = "L3-L1"
)

func (s *Phase) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch Phase(raw) {
	case PhaseL1,
		PhaseL2,
		PhaseL3,
		PhaseN,
		PhaseL1N,
		PhaseL2N,
		PhaseL3N,
		PhaseL1L2,
		PhaseL2L3,
		PhaseL3L1:
		*s = Phase(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid Phase", raw),
	)
}

// ReadingContext (7.35)
//
// Values of the context field of a value in SampledValue.
type ReadingContext string

const (
	// Value taken at start of interruption.
	ReadingContextInterruptionBegin ReadingContext = "Interruption.Begin"
	// Value taken when resuming after interruption.
	ReadingContextInterruptionEnd ReadingContext = "Interruption.End"
	// Value for any other situations.
	ReadingContextOther ReadingContext = "Other"
	// Value taken at clock aligned interval.
	ReadingContextSampleClock ReadingContext = "Sample.Clock"
	// Value taken as periodic sample relative to start time of transaction.
	ReadingContextSamplePeriodic ReadingContext = "Sample.Periodic"
	// Value taken at start of transaction.
	ReadingContextTransactionBegin ReadingContext = "Transaction.Begin"
	// Value taken at end of transaction.
	ReadingContextTransactionEnd ReadingContext = "Transaction.End"
	// Value taken in response to a TriggerMessage.req
	ReadingContextTrigger ReadingContext = "Trigger"
)

func (s *ReadingContext) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReadingContext(raw) {
	case ReadingContextInterruptionBegin,
		ReadingContextInterruptionEnd,
		ReadingContextOther,
		ReadingContextSampleClock,
		ReadingContextSamplePeriodic,
		ReadingContextTransactionBegin,
		ReadingContextTransactionEnd,
		ReadingContextTrigger:
		*s = ReadingContext(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReadingContext", raw),
	)
}

// Reason (7.36)
//
// Reason for stopping a transaction in StopTransaction.req.
type Reason string

const (
	// The transaction was stopped because of the authorization status in a StartTransaction.conf
	ReasonDeAuthorized Reason = "DeAuthorized"
	// Emergency stop button was used.
	ReasonEmergencyStop Reason = "EmergencyStop"
	// Disconnecting of cable, vehicle moved away from inductive charge unit.
	ReasonEVDisconnected Reason = "EVDisconnected"
	// A hard reset command was received.
	ReasonHardReset Reason = "HardReset"
	// Stopped locally on request of the user at the Charge Point. This is a regular termination of a transaction. Examples: presenting
	// an RFID tag, pressing a button to stop.
	ReasonLocal Reason = "Local"
	// Any other reason.
	ReasonOther Reason = "Other"
	// Complete loss of power.
	ReasonPowerLoss Reason = "PowerLoss"
	// A locally initiated reset/reboot occurred. (for instance watchdog kicked in)
	ReasonReboot Reason = "Reboot"
	// Stopped remotely on request of the user. This is a regular termination of a transaction. Examples: termination using a
	// smartphone app, exceeding a (non local) prepaid credit.
	ReasonRemote Reason = "Remote"
	// A soft reset command was received.
	ReasonSoftReset Reason = "SoftReset"
	// Central System sent an Unlock Connector command.
	ReasonUnlockCommand Reason = "UnlockCommand"
)

func (s *Reason) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch Reason(raw) {
	case ReasonDeAuthorized,
		ReasonEmergencyStop,
		ReasonEVDisconnected,
		ReasonHardReset,
		ReasonLocal,
		ReasonOther,
		ReasonPowerLoss,
		ReasonReboot,
		ReasonRemote,
		ReasonSoftReset,
		ReasonUnlockCommand:
		*s = Reason(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid Reason", raw),
	)
}

// RecurrencyKindType (7.37)
//
// Type of recurrence of a charging profile, as used in ChargingProfile.
type RecurrencyKindType string

const (
	// The schedule restarts every 24 hours, at the same time as in the startSchedule.
	RecurrencyKindTypeDaily RecurrencyKindType = "Daily"
	// The schedule restarts every 7 days, at the same time and day-of-the-week as in the startSchedule.
	RecurrencyKindTypeWeekly RecurrencyKindType = "Weekly"
)

func (s *RecurrencyKindType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch RecurrencyKindType(raw) {
	case RecurrencyKindTypeDaily,
		RecurrencyKindTypeWeekly:
		*s = RecurrencyKindType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid RecurrencyKindType", raw),
	)
}

// RegistrationStatus (7.38)
//
// Result of registration in response to BootNotification.req.
type RegistrationStatus string

const (
	// Charge point is accepted by Central System.
	RegistrationStatusAccepted RegistrationStatus = "Accepted"
	// Central System is not yet ready to accept the Charge Point. Central System may send messages to retrieve information or
	// prepare the Charge Point.
	RegistrationStatusPending RegistrationStatus = "Pending"
	// Charge point is not accepted by Central System. This may happen when the Charge Point id is not known by Central System.
	RegistrationStatusRejected RegistrationStatus = "Rejected"
)

func (s *RegistrationStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch RegistrationStatus(raw) {
	case RegistrationStatusAccepted,
		RegistrationStatusPending,
		RegistrationStatusRejected:
		*s = RegistrationStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid RegistrationStatus", raw),
	)
}

// RemoteStartStopStatus (7.39)
//
// The result of a RemoteStartTransaction.req or RemoteStopTransaction.req request.
type RemoteStartStopStatus string

const (
	// Command will be executed.
	RemoteStartStopStatusAccepted RemoteStartStopStatus = "Accepted"
	// Command will not be executed.
	RemoteStartStopStatusRejected RemoteStartStopStatus = "Rejected"
)

func (s *RemoteStartStopStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch RemoteStartStopStatus(raw) {
	case RemoteStartStopStatusAccepted,
		RemoteStartStopStatusRejected:
		*s = RemoteStartStopStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid RemoteStartStopStatus", raw),
	)
}

// ReservationStatus (7.40)
//
// Status in ReserveNow.conf.
type ReservationStatus string

const (
	// Reservation has been made.
	ReservationStatusAccepted ReservationStatus = "Accepted"
	// Reservation has not been made, because connectors or specified connector are in a faulted state.
	ReservationStatusFaulted ReservationStatus = "Faulted"
	// Reservation has not been made. All connectors or the specified connector are occupied.
	ReservationStatusOccupied ReservationStatus = "Occupied"
	// Reservation has not been made. Charge Point is not configured to accept reservations.
	ReservationStatusRejected ReservationStatus = "Rejected"
	// Reservation has not been made, because connectors or specified connector are in an unavailable state.
	ReservationStatusUnavailable ReservationStatus = "Unavailable"
)

func (s *ReservationStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReservationStatus(raw) {
	case ReservationStatusAccepted,
		ReservationStatusFaulted,
		ReservationStatusOccupied,
		ReservationStatusRejected,
		ReservationStatusUnavailable:
		*s = ReservationStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReservationStatus", raw),
	)
}

// ResetStatus (7.41)
//
// Result of Reset.req.
type ResetStatus string

const (
	// Command will be executed.
	ResetStatusAccepted ResetStatus = "Accepted"
	// Command will not be executed.
	ResetStatusRejected ResetStatus = "Rejected"
)

func (s *ResetStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ResetStatus(raw) {
	case ResetStatusAccepted,
		ResetStatusRejected:
		*s = ResetStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ResetStatus", raw),
	)
}

// ResetType (7.42)
//
// Type of reset requested by Reset.req.
type ResetType string

const (
	// Restart (all) the hardware, the Charge Point is not required to gracefully stop ongoing transaction. If possible the Charge Point
	// sends a StopTransaction.req for previously ongoing transactions after having restarted and having been accepted by the
	// Central System via a BootNotification.conf. This is a last resort solution for a not correctly functioning Charge Point, by sending
	// a "hard" reset, (queued) information might get lost.
	ResetTypeHard ResetType = "Hard"
	// Stop ongoing transactions gracefully and sending StopTransaction.req for every ongoing transaction. It should then restart the
	// application software (if possible, otherwise restart the processor/controller).
	ResetTypeSoft ResetType = "Soft"
)

func (s *ResetType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ResetType(raw) {
	case ResetTypeHard,
		ResetTypeSoft:
		*s = ResetType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ResetType", raw),
	)
}

// SampledValue (7.43)
//
// Single sampled value in MeterValues. Each value can be accompanied by optional fields.
type SampledValue struct {
	// Required. Value as a “Raw” (decimal) number or “SignedData”. Field Type is
	// “string” to allow for digitally signed data readings. Decimal numeric values are
	// also acceptable to allow fractional values for measurands such as Temperature
	// and Current.
	Value string `json:"value"`
	// Optional. Type of detail value: start, end or sample. Default = “Sample.Periodic”
	Context *ReadingContext `json:"context,omitempty"`
	// Optional. Raw or signed data. Default = “Raw”
	Format *ValueFormat `json:"format,omitempty"`
	// Optional. Type of measurement. Default = “Energy.Active.Import.Register”
	Measurand *Measurand `json:"measurand,omitempty"`
	// Optional. indicates how the measured value is to be interpreted. For instance
	// between L1 and neutral (L1-N) Please note that not all values of phase are
	// applicable to all Measurands. When phase is absent, the measured value is
	// interpreted as an overall value.
	Phase *Phase `json:"phase,omitempty"`
	// Optional. Location of measurement. Default=”Outlet”
	Location *Location `json:"location,omitempty"`
	// Optional. Unit of the value. Default = “Wh” if the (default) measurand is an
	// “Energy” type.
	Unit *UnitOfMeasure `json:"unit,omitempty"`
}

func (s *SampledValue) UnmarshalJSON(data []byte) error {
	type Alias SampledValue
	var raw Alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*s = SampledValue(raw)
	return s.Validate()
}

func (s SampledValue) Validate() error {
	if s.Value == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "value", "required field is missing")
	}

	return nil
}

// TriggerMessageStatus (7.44)
//
// Status in TriggerMessage.conf.
type TriggerMessageStatus string

const (
	// Requested notification will be sent.
	TriggerMessageStatusAccepted TriggerMessageStatus = "Accepted"
	// Requested notification will not be sent.
	TriggerMessageStatusRejected TriggerMessageStatus = "Rejected"
	// Requested notification cannot be sent because it is either not implemented or unknown.
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
)

func (s *TriggerMessageStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TriggerMessageStatus(raw) {
	case TriggerMessageStatusAccepted,
		TriggerMessageStatusRejected,
		TriggerMessageStatusNotImplemented:
		*s = TriggerMessageStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TriggerMessageStatus", raw),
	)
}

// UnitOfMeasure (7.45)
//
// Allowable values of the optional "unit" field of a Value element, as used in SampledValue. Default value of "unit"
// is always "Wh".
type UnitOfMeasure string

const (
	// Watt-hours (energy). Default.
	UnitOfMeasureWh UnitOfMeasure = "Wh"
	// kiloWatt-hours (energy).
	UnitOfMeasureKWh UnitOfMeasure = "kWh"
	// Var-hours (reactive energy).
	UnitOfMeasureVarh UnitOfMeasure = "varh"
	// kilovar-hours (reactive energy).
	UnitOfMeasureKvarh UnitOfMeasure = "kvarh"
	// Watts (power).
	UnitOfMeasureW UnitOfMeasure = "W"
	// kilowatts (power).
	UnitOfMeasureKW UnitOfMeasure = "kW"
	// VoltAmpere (apparent power).
	UnitOfMeasureVA UnitOfMeasure = "VA"
	// kiloVolt Ampere (apparent power).
	UnitOfMeasureKVA UnitOfMeasure = "kVA"
	// Vars (reactive power).
	UnitOfMeasureVar UnitOfMeasure = "var"
	// kilovars (reactive power).
	UnitOfMeasureKvar UnitOfMeasure = "kvar"
	// Amperes (current).
	UnitOfMeasureA UnitOfMeasure = "A"
	// Voltage (r.m.s. AC).
	UnitOfMeasureV UnitOfMeasure = "V"
	// Degrees (temperature).
	UnitOfMeasureCelsius UnitOfMeasure = "Celsius"
	// Degrees (temperature).
	UnitOfMeasureFahrenheit UnitOfMeasure = "Fahrenheit"
	// Degrees Kelvin (temperature).
	UnitOfMeasureK UnitOfMeasure = "K"
	// Percentage.
	UnitOfMeasurePercent UnitOfMeasure = "Percent"
)

func (s *UnitOfMeasure) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UnitOfMeasure(raw) {
	case UnitOfMeasureWh,
		UnitOfMeasureKWh,
		UnitOfMeasureVarh,
		UnitOfMeasureKvarh,
		UnitOfMeasureW,
		UnitOfMeasureKW,
		UnitOfMeasureVA,
		UnitOfMeasureKVA,
		UnitOfMeasureVar,
		UnitOfMeasureKvar,
		UnitOfMeasureA,
		UnitOfMeasureV,
		UnitOfMeasureCelsius,
		UnitOfMeasureFahrenheit,
		UnitOfMeasureK,
		UnitOfMeasurePercent:
		*s = UnitOfMeasure(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UnitOfMeasure", raw),
	)
}

// UnlockStatus (7.46)
//
// Status in response to UnlockConnector.req.
type UnlockStatus string

const (
	// Connector has successfully been unlocked.
	UnlockStatusUnlocked UnlockStatus = "Unlocked"
	// Failed to unlock the connector: The Charge Point has tried to unlock the connector and has detected that the connector is still
	// locked or the unlock mechanism failed.
	UnlockStatusUnlockFailed UnlockStatus = "UnlockFailed"
	// Charge Point has no connector lock, or ConnectorId is unknown.
	UnlockStatusNotSupported UnlockStatus = "NotSupported"
)

func (s *UnlockStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UnlockStatus(raw) {
	case UnlockStatusUnlocked,
		UnlockStatusUnlockFailed,
		UnlockStatusNotSupported:
		*s = UnlockStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UnlockStatus", raw),
	)
}

// UpdateStatus (7.47)
//
// Type of update for a SendLocalList.req.
type UpdateStatus string

const (
	// Local Authorization List successfully updated.
	UpdateStatusAccepted UpdateStatus = "Accepted"
	// Failed to update the Local Authorization List.
	UpdateStatusFailed UpdateStatus = "Failed"
	// Update of Local Authorization List is not supported by Charge Point.
	UpdateStatusNotSupported UpdateStatus = "NotSupported"
	// Version number in the request for a differential update is less or equal then version number of current list.
	UpdateStatusVersionMismatch UpdateStatus = "VersionMismatch"
)

func (s *UpdateStatus) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UpdateStatus(raw) {
	case UpdateStatusAccepted,
		UpdateStatusFailed,
		UpdateStatusNotSupported,
		UpdateStatusVersionMismatch:
		*s = UpdateStatus(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UpdateStatus", raw),
	)
}

// UpdateType (7.48)
//
// Type of update for a SendLocalList.req.
type UpdateType string

const (
	// Indicates that the current Local Authorization List must be updated with the values in this message.
	UpdateTypeDifferential UpdateType = "Differential"
	// Indicates that the current Local Authorization List must be replaced by the values in this message.
	UpdateTypeFull UpdateType = "Full"
)

func (s *UpdateType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UpdateType(raw) {
	case UpdateTypeDifferential,
		UpdateTypeFull:
		*s = UpdateType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UpdateType", raw),
	)
}

// ValueFormat (7.49)
//
// Format that specifies how the value element in SampledValue is to be interpreted.
type ValueFormat string

const (
	// Data is to be interpreted as integer/decimal numeric data.
	ValueFormatRaw ValueFormat = "Raw"
	// Data is represented as a signed binary data block, encoded as hex data.
	ValueFormatSignedData ValueFormat = "SignedData"
)

func (s *ValueFormat) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ValueFormat(raw) {
	case ValueFormatRaw,
		ValueFormatSignedData:
		*s = ValueFormat(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ValueFormat", raw),
	)
}
