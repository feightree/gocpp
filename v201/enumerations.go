package v201

import (
	"encoding/json"
	"fmt"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// APNAuthenticationEnumType (3.1)
type APNAuthenticationEnumType string

const (
	// Use CHAP authentication
	APNAuthenticationEnumTypeCHAP APNAuthenticationEnumType = "CHAP"
	// Use no authentication
	APNAuthenticationEnumTypeNONE APNAuthenticationEnumType = "NONE"
	// Use PAP authentication
	APNAuthenticationEnumTypePAP APNAuthenticationEnumType = "PAP"
	// Sequentially try CHAP, PAP, NONE.
	APNAuthenticationEnumTypeAUTO APNAuthenticationEnumType = "AUTO"
)

func (s *APNAuthenticationEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch APNAuthenticationEnumType(raw) {
	case APNAuthenticationEnumTypeCHAP,
		APNAuthenticationEnumTypeNONE,
		APNAuthenticationEnumTypePAP,
		APNAuthenticationEnumTypeAUTO:
		*s = APNAuthenticationEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid APNAuthenticationEnumType", raw),
	)
}

// AttributeEnumType (3.2)
type AttributeEnumType string

const (
	// The actual value of the variable.
	AttributeEnumTypeActual AttributeEnumType = "Actual"
	// The target value for this variable.
	AttributeEnumTypeTarget AttributeEnumType = "Target"
	// The minimal allowed value for this variable
	AttributeEnumTypeMinSet AttributeEnumType = "MinSet"
	// Thne maximum allowed value for this variable
	AttributeEnumTypeMaxSet AttributeEnumType = "MaxSet"
)

func (s *AttributeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch AttributeEnumType(raw) {
	case AttributeEnumTypeActual,
		AttributeEnumTypeTarget,
		AttributeEnumTypeMinSet,
		AttributeEnumTypeMaxSet:
		*s = AttributeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid AttributeEnumType", raw),
	)
}

// AuthorizationStatusEnumType (3.3)
//
// Status of an authorization response.
type AuthorizationStatusEnumType string

const (
	// Identifier is allowed for charging.
	AuthorizationStatusEnumTypeAccepted AuthorizationStatusEnumType = "Accepted"
	// Identifier has been blocked. Not allowed for charging.
	AuthorizationStatusEnumTypeBlocked AuthorizationStatusEnumType = "Blocked"
	// Identifier is already involved in another transaction and multiple transactions are not allowed.
	// (Only relevant for the response to a transactionEventRequest(eventType=Started).)
	AuthorizationStatusEnumTypeConcurrentTx AuthorizationStatusEnumType = "ConcurrentTx"
	// Identifier has expired. Not allowed for charging.
	AuthorizationStatusEnumTypeExpired AuthorizationStatusEnumType = "Expired"
	// Identifier is invalid. Not allowed for charging.
	AuthorizationStatusEnumTypeInvalid AuthorizationStatusEnumType = "Invalid"
	// Identifier is valid, but EV Driver doesn’t have enough credit to start charging. Not allowed for
	// charging.
	AuthorizationStatusEnumTypeNoCredit AuthorizationStatusEnumType = "NoCredit"
	// Identifier is valid, but not allowed to charge at this type of EVSE.
	AuthorizationStatusEnumTypeNotAllowedTypeEVSE AuthorizationStatusEnumType = "NotAllowedTypeEVSE"
	// Identifier is valid, but not allowed to charge at this location.
	AuthorizationStatusEnumTypeNotAtThisLocation AuthorizationStatusEnumType = "NotAtThisLocation"
	// Identifier is valid, but not allowed to charge at this location at this time.
	AuthorizationStatusEnumTypeNotAtThisTime AuthorizationStatusEnumType = "NotAtThisTime"
	// Identifier is unknown. Not allowed for charging.
	AuthorizationStatusEnumTypeUnknown AuthorizationStatusEnumType = "Unknown"
)

func (s *AuthorizationStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch AuthorizationStatusEnumType(raw) {
	case AuthorizationStatusEnumTypeAccepted,
		AuthorizationStatusEnumTypeBlocked,
		AuthorizationStatusEnumTypeConcurrentTx,
		AuthorizationStatusEnumTypeExpired,
		AuthorizationStatusEnumTypeInvalid,
		AuthorizationStatusEnumTypeNoCredit,
		AuthorizationStatusEnumTypeNotAllowedTypeEVSE,
		AuthorizationStatusEnumTypeNotAtThisLocation,
		AuthorizationStatusEnumTypeNotAtThisTime,
		AuthorizationStatusEnumTypeUnknown:
		*s = AuthorizationStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid AuthorizationStatusEnumType", raw),
	)
}

// AuthorizeCertificateStatusEnumType (3.4)
//
// Status of the EV Contract certificate.
type AuthorizeCertificateStatusEnumType string

const (
	// Positive response
	AuthorizeCertificateStatusEnumTypeAccepted AuthorizeCertificateStatusEnumType = "Accepted"
	// <not used>
	AuthorizeCertificateStatusEnumTypeSignatureError AuthorizeCertificateStatusEnumType = "SignatureError"
	// If the contract certificate in the AuthorizeRequest is expired.
	AuthorizeCertificateStatusEnumTypeCertificateExpired AuthorizeCertificateStatusEnumType = "CertificateExpired"
	// If the Charging Station or CSMS determine (via a CRL or OCSP response) that the contract certificate
	// in the AuthorizeRequest is marked as revoked.
	AuthorizeCertificateStatusEnumTypeCertificateRevoked AuthorizeCertificateStatusEnumType = "CertificateRevoked"
	// <not used>
	AuthorizeCertificateStatusEnumTypeNoCertificateAvailable AuthorizeCertificateStatusEnumType = "NoCertificateAvailable"
	// If the contract certificate contained in the AuthorizeRequest message is not valid.
	AuthorizeCertificateStatusEnumTypeCertChainError AuthorizeCertificateStatusEnumType = "CertChainError"
	// If the EMAID provided by EVCC is invalid, unknown, expired or blocked.
	AuthorizeCertificateStatusEnumTypeContractCancelled AuthorizeCertificateStatusEnumType = "ContractCancelled"
)

func (s *AuthorizeCertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch AuthorizeCertificateStatusEnumType(raw) {
	case AuthorizeCertificateStatusEnumTypeAccepted,
		AuthorizeCertificateStatusEnumTypeSignatureError,
		AuthorizeCertificateStatusEnumTypeCertificateExpired,
		AuthorizeCertificateStatusEnumTypeCertificateRevoked,
		AuthorizeCertificateStatusEnumTypeNoCertificateAvailable,
		AuthorizeCertificateStatusEnumTypeCertChainError,
		AuthorizeCertificateStatusEnumTypeContractCancelled:
		*s = AuthorizeCertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid AuthorizeCertificateStatusEnumType", raw),
	)
}

// BootReasonEnumType (3.5)
type BootReasonEnumType string

const (
	// The Charging Station rebooted due to an application error.
	BootReasonEnumTypeApplicationReset BootReasonEnumType = "ApplicationReset"
	// The Charging Station rebooted due to a firmware update.
	BootReasonEnumTypeFirmwareUpdate BootReasonEnumType = "FirmwareUpdate"
	// The Charging Station rebooted due to a local reset command.
	BootReasonEnumTypeLocalReset BootReasonEnumType = "LocalReset"
	// The Charging Station powered up and registers itself with the CSMS.
	BootReasonEnumTypePowerUp BootReasonEnumType = "PowerUp"
	// The Charging Station rebooted due to a remote reset command.
	BootReasonEnumTypeRemoteReset BootReasonEnumType = "RemoteReset"
	// The Charging Station rebooted due to a scheduled reset command.
	BootReasonEnumTypeScheduledReset BootReasonEnumType = "ScheduledReset"
	// Requested by the CSMS via a TriggerMessage
	BootReasonEnumTypeTriggered BootReasonEnumType = "Triggered"
	// The boot reason is unknown.
	BootReasonEnumTypeUnknown BootReasonEnumType = "Unknown"
	// The Charging Station rebooted due to an elapsed watchdog timer.
	BootReasonEnumTypeWatchdog BootReasonEnumType = "Watchdog"
)

func (s *BootReasonEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch BootReasonEnumType(raw) {
	case BootReasonEnumTypeApplicationReset,
		BootReasonEnumTypeFirmwareUpdate,
		BootReasonEnumTypeLocalReset,
		BootReasonEnumTypePowerUp,
		BootReasonEnumTypeRemoteReset,
		BootReasonEnumTypeScheduledReset,
		BootReasonEnumTypeTriggered,
		BootReasonEnumTypeUnknown,
		BootReasonEnumTypeWatchdog:
		*s = BootReasonEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid BootReasonEnumType", raw),
	)
}

// CancelReservationStatusEnumType (3.6)
//
// Status in CancelReservationResponse.
type CancelReservationStatusEnumType string

const (
	// Reservation for the identifier has been canceled.
	CancelReservationStatusEnumTypeAccepted CancelReservationStatusEnumType = "Accepted"
	// Reservation could not be canceled, because there is no reservation active for the identifier.
	CancelReservationStatusEnumTypeRejected CancelReservationStatusEnumType = "Rejected"
)

func (s *CancelReservationStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CancelReservationStatusEnumType(raw) {
	case CancelReservationStatusEnumTypeAccepted,
		CancelReservationStatusEnumTypeRejected:
		*s = CancelReservationStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CancelReservationStatusEnumType", raw),
	)
}

// CertificateActionEnumType (3.7)
type CertificateActionEnumType string

const (
	// Install the provided certificate.
	CertificateActionEnumTypeInstall CertificateActionEnumType = "Install"
	// Update the provided certificate.
	CertificateActionEnumTypeUpdate CertificateActionEnumType = "Update"
)

func (s *CertificateActionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CertificateActionEnumType(raw) {
	case CertificateActionEnumTypeInstall,
		CertificateActionEnumTypeUpdate:
		*s = CertificateActionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CertificateActionEnumType", raw),
	)
}

// CertificateSignedStatusEnumType (3.8)
type CertificateSignedStatusEnumType string

const (
	// Signed certificate is valid.
	CertificateSignedStatusEnumTypeAccepted CertificateSignedStatusEnumType = "Accepted"
	// Signed certificate is invalid.
	CertificateSignedStatusEnumTypeRejected CertificateSignedStatusEnumType = "Rejected"
)

func (s *CertificateSignedStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CertificateSignedStatusEnumType(raw) {
	case CertificateSignedStatusEnumTypeAccepted,
		CertificateSignedStatusEnumTypeRejected:
		*s = CertificateSignedStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CertificateSignedStatusEnumType", raw),
	)
}

// CertificateSigningUseEnumType (3.9)
type CertificateSigningUseEnumType string

const (
	// Client side certificate used by the Charging Station to connect the the CSMS.
	CertificateSigningUseEnumTypeChargingStationCertificate CertificateSigningUseEnumType = "ChargingStationCertificate"
	// Use for certificate for 15118 connections. This means that the certificate should be derived from
	// the V2G root.
	CertificateSigningUseEnumTypeV2GCertificate CertificateSigningUseEnumType = "V2GCertificate"
)

func (s *CertificateSigningUseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CertificateSigningUseEnumType(raw) {
	case CertificateSigningUseEnumTypeChargingStationCertificate,
		CertificateSigningUseEnumTypeV2GCertificate:
		*s = CertificateSigningUseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CertificateSigningUseEnumType", raw),
	)
}

// ChangeAvailabilityStatusEnumType (3.10)
//
// Status returned in response to ChangeAvailabilityRequest.
type ChangeAvailabilityStatusEnumType string

const (
	// Request has been accepted and will be executed.
	ChangeAvailabilityStatusEnumTypeAccepted ChangeAvailabilityStatusEnumType = "Accepted"
	// Request has not been accepted and will not be executed.
	ChangeAvailabilityStatusEnumTypeRejected ChangeAvailabilityStatusEnumType = "Rejected"
	// Request has been accepted and will be executed when transaction(s) in progress have finished.
	ChangeAvailabilityStatusEnumTypeScheduled ChangeAvailabilityStatusEnumType = "Scheduled"
)

func (s *ChangeAvailabilityStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChangeAvailabilityStatusEnumType(raw) {
	case ChangeAvailabilityStatusEnumTypeAccepted,
		ChangeAvailabilityStatusEnumTypeRejected,
		ChangeAvailabilityStatusEnumTypeScheduled:
		*s = ChangeAvailabilityStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChangeAvailabilityStatusEnumType", raw),
	)
}

// ChargingLimitSourceEnumType (3.11)
//
// Enumeration for indicating from which source a charging limit originates.
type ChargingLimitSourceEnumType string

const (
	// Indicates that an Energy Management System has sent a charging limit.
	ChargingLimitSourceEnumTypeEMS ChargingLimitSourceEnumType = "EMS"
	// Indicates that an external source, not being an EMS or system operator, has sent a charging limit.
	ChargingLimitSourceEnumTypeOther ChargingLimitSourceEnumType = "Other"
	// Indicates that a System Operator (DSO or TSO) has sent a charging limit.
	ChargingLimitSourceEnumTypeSO ChargingLimitSourceEnumType = "SO"
	// Indicates that the CSO has set this charging profile.
	ChargingLimitSourceEnumTypeCSO ChargingLimitSourceEnumType = "CSO"
)

func (s *ChargingLimitSourceEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingLimitSourceEnumType(raw) {
	case ChargingLimitSourceEnumTypeEMS,
		ChargingLimitSourceEnumTypeOther,
		ChargingLimitSourceEnumTypeSO,
		ChargingLimitSourceEnumTypeCSO:
		*s = ChargingLimitSourceEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingLimitSourceEnumType", raw),
	)
}

// ChargingProfileKindEnumType (3.12)
//
// Kind of charging profile.
type ChargingProfileKindEnumType string

const (
	// Schedule periods are relative to a fixed point in time defined in the schedule. This requires that
	// startSchedule is set to a starting point in time.
	ChargingProfileKindEnumTypeAbsolute ChargingProfileKindEnumType = "Absolute"
	// The schedule restarts periodically at the first schedule period. To be most useful, this requires
	// that startSchedule is set to a starting point in time.
	ChargingProfileKindEnumTypeRecurring ChargingProfileKindEnumType = "Recurring"
	// Charging schedule periods should start when the EVSE is ready to deliver energy. i.e. when the EV
	// driver is authorized and the EV is connected. When a ChargingProfile is received for a transaction
	// that is already charging, then the charging schedule periods should remain relative to the
	// PowerPathClosed moment. No value for startSchedule should be supplied.
	ChargingProfileKindEnumTypeRelative ChargingProfileKindEnumType = "Relative"
)

func (s *ChargingProfileKindEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfileKindEnumType(raw) {
	case ChargingProfileKindEnumTypeAbsolute,
		ChargingProfileKindEnumTypeRecurring,
		ChargingProfileKindEnumTypeRelative:
		*s = ChargingProfileKindEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfileKindEnumType", raw),
	)
}

// ChargingProfilePurposeEnumType (3.13)
//
// Purpose of the charging profile.
type ChargingProfilePurposeEnumType string

const (
	// Additional constraints that will be incorporated into a local power schedule. Only valid for a
	// Charging Station. Therefore evse.Id MUST be 0 in the SetChargingProfileRequest message.
	ChargingProfilePurposeEnumTypeChargingStationExternalConstraints ChargingProfilePurposeEnumType = "ChargingStationExternalConstraints"
	// Configuration for the maximum power or current available for an entire Charging Station.
	ChargingProfilePurposeEnumTypeChargingStationMaxProfile ChargingProfilePurposeEnumType = "ChargingStationMaxProfile"
	// Default profile that can be configured in the Charging Station. When a new transaction is started,
	// this profile SHALL be used, unless it was a transaction that was started by a
	// RequestStartTransactionRequest with a ChargingProfile that is accepted by the Charging Station.
	ChargingProfilePurposeEnumTypeTxDefaultProfile ChargingProfilePurposeEnumType = "TxDefaultProfile"
	// Profile with constraints to be imposed by the Charging Station on the current transaction, or on a
	// new transaction when this is started via a RequestStartTransactionRequest with a ChargingProfile. A
	// profile with this purpose SHALL cease to be valid when the transaction terminates.
	ChargingProfilePurposeEnumTypeTxProfile ChargingProfilePurposeEnumType = "TxProfile"
)

func (s *ChargingProfilePurposeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfilePurposeEnumType(raw) {
	case ChargingProfilePurposeEnumTypeChargingStationExternalConstraints,
		ChargingProfilePurposeEnumTypeChargingStationMaxProfile,
		ChargingProfilePurposeEnumTypeTxDefaultProfile,
		ChargingProfilePurposeEnumTypeTxProfile:
		*s = ChargingProfilePurposeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfilePurposeEnumType", raw),
	)
}

// ChargingProfileStatusEnumType (3.14)
//
// Status returned in response to SetChargingProfileRequest.
type ChargingProfileStatusEnumType string

const (
	// Request has been accepted and will be executed.
	ChargingProfileStatusEnumTypeAccepted ChargingProfileStatusEnumType = "Accepted"
	// Request has not been accepted and will not be executed.
	ChargingProfileStatusEnumTypeRejected ChargingProfileStatusEnumType = "Rejected"
)

func (s *ChargingProfileStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfileStatusEnumType(raw) {
	case ChargingProfileStatusEnumTypeAccepted,
		ChargingProfileStatusEnumTypeRejected:
		*s = ChargingProfileStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfileStatusEnumType", raw),
	)
}

// ChargingRateUnitEnumType (3.15)
//
// Unit in which a charging schedule is defined.
type ChargingRateUnitEnumType string

const (
	// Watts (power). This is the TOTAL allowed charging power. If used for AC Charging, the phase current
	// should be calculated via: Current per phase = Power / (Line Voltage * Number of Phases). The "Line
	// Voltage" used in the calculation is not the measured voltage, but the set voltage for the area
	// (hence, 230 of 110 volt). The "Number of Phases" is the numberPhases from the
	// ChargingSchedulePeriod. It is usually more convenient to use this for DC charging. Note that if
	// numberPhases in a ChargingSchedulePeriod is absent, 3 SHALL be assumed.
	ChargingRateUnitEnumTypeW ChargingRateUnitEnumType = "W"
	// Amperes (current). The amount of Ampere per phase, not the sum of all phases. It is usually more
	// convenient to use this for AC charging.
	ChargingRateUnitEnumTypeA ChargingRateUnitEnumType = "A"
)

func (s *ChargingRateUnitEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingRateUnitEnumType(raw) {
	case ChargingRateUnitEnumTypeW,
		ChargingRateUnitEnumTypeA:
		*s = ChargingRateUnitEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingRateUnitEnumType", raw),
	)
}

// ChargingStateEnumType (3.16)
//
// The state of the charging process.
type ChargingStateEnumType string

const (
	// The contactor of the Connector is closed and energy is flowing to between EVSE and EV.
	ChargingStateEnumTypeCharging ChargingStateEnumType = "Charging"
	// There is a connection between EV and EVSE, in case the protocol used between EV and the Charging
	// Station can detect a connection, the protocol needs to detect this for the state to become active.
	// The connection can either be wired or wireless.
	ChargingStateEnumTypeEVConnected ChargingStateEnumType = "EVConnected"
	// When the EV is connected to the EVSE and the EVSE is offering energy but the EV is not taking any
	// energy.
	ChargingStateEnumTypeSuspendedEV ChargingStateEnumType = "SuspendedEV"
	// When the EV is connected to the EVSE but the EVSE is not offering energy to the EV, e.g. due to a
	// smart charging restriction, local supply power constraints, or when charging has stopped because of
	// the authorization status in the response to a transactionEventRequest indicating that charging is
	// not allowed etc.
	ChargingStateEnumTypeSuspendedEVSE ChargingStateEnumType = "SuspendedEVSE"
	// There is no connection between EV and EVSE.
	ChargingStateEnumTypeIdle ChargingStateEnumType = "Idle"
)

func (s *ChargingStateEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingStateEnumType(raw) {
	case ChargingStateEnumTypeCharging,
		ChargingStateEnumTypeEVConnected,
		ChargingStateEnumTypeSuspendedEV,
		ChargingStateEnumTypeSuspendedEVSE,
		ChargingStateEnumTypeIdle:
		*s = ChargingStateEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingStateEnumType", raw),
	)
}

// ClearCacheStatusEnumType (3.17)
//
// Status returned in response to ClearCacheRequest.
type ClearCacheStatusEnumType string

const (
	// Command has been executed.
	ClearCacheStatusEnumTypeAccepted ClearCacheStatusEnumType = "Accepted"
	// Command has not been executed.
	ClearCacheStatusEnumTypeRejected ClearCacheStatusEnumType = "Rejected"
)

func (s *ClearCacheStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearCacheStatusEnumType(raw) {
	case ClearCacheStatusEnumTypeAccepted,
		ClearCacheStatusEnumTypeRejected:
		*s = ClearCacheStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearCacheStatusEnumType", raw),
	)
}

// ClearChargingProfileStatusEnumType (3.18)
//
// Status returned in response to ClearChargingProfileRequest.
type ClearChargingProfileStatusEnumType string

const (
	// Request has been accepted and will be executed.
	ClearChargingProfileStatusEnumTypeAccepted ClearChargingProfileStatusEnumType = "Accepted"
	// No Charging Profile(s) were found matching the request.
	ClearChargingProfileStatusEnumTypeUnknown ClearChargingProfileStatusEnumType = "Unknown"
)

func (s *ClearChargingProfileStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearChargingProfileStatusEnumType(raw) {
	case ClearChargingProfileStatusEnumTypeAccepted,
		ClearChargingProfileStatusEnumTypeUnknown:
		*s = ClearChargingProfileStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearChargingProfileStatusEnumType", raw),
	)
}

// ClearMessageStatusEnumType (3.19)
//
// Result for a ClearDisplayMessageRequest as used in a ClearDisplayMessageResponse.
type ClearMessageStatusEnumType string

const (
	// Request successfully executed: message cleared.
	ClearMessageStatusEnumTypeAccepted ClearMessageStatusEnumType = "Accepted"
	// Given message (based on the id) not known.
	ClearMessageStatusEnumTypeUnknown ClearMessageStatusEnumType = "Unknown"
)

func (s *ClearMessageStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearMessageStatusEnumType(raw) {
	case ClearMessageStatusEnumTypeAccepted,
		ClearMessageStatusEnumTypeUnknown:
		*s = ClearMessageStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearMessageStatusEnumType", raw),
	)
}

// ClearMonitoringStatusEnumType (3.20)
type ClearMonitoringStatusEnumType string

const (
	// Monitor successfully cleared.
	ClearMonitoringStatusEnumTypeAccepted ClearMonitoringStatusEnumType = "Accepted"
	// Clearing of monitor rejected.
	ClearMonitoringStatusEnumTypeRejected ClearMonitoringStatusEnumType = "Rejected"
	// Monitor Id is not found.
	ClearMonitoringStatusEnumTypeNotFound ClearMonitoringStatusEnumType = "NotFound"
)

func (s *ClearMonitoringStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearMonitoringStatusEnumType(raw) {
	case ClearMonitoringStatusEnumTypeAccepted,
		ClearMonitoringStatusEnumTypeRejected,
		ClearMonitoringStatusEnumTypeNotFound:
		*s = ClearMonitoringStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearMonitoringStatusEnumType", raw),
	)
}

// ComponentCriterionEnumType (3.21)
type ComponentCriterionEnumType string

const (
	// Components that are active, i.e. having Active = 1
	ComponentCriterionEnumTypeActive ComponentCriterionEnumType = "Active"
	// Components that are available, i.e. having Available = 1
	ComponentCriterionEnumTypeAvailable ComponentCriterionEnumType = "Available"
	// Components that are enabled, i.e. having Enabled = 1
	ComponentCriterionEnumTypeEnabled ComponentCriterionEnumType = "Enabled"
	// Components that reported a problem, i.e. having Problem = 1
	ComponentCriterionEnumTypeProblem ComponentCriterionEnumType = "Problem"
)

func (s *ComponentCriterionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ComponentCriterionEnumType(raw) {
	case ComponentCriterionEnumTypeActive,
		ComponentCriterionEnumTypeAvailable,
		ComponentCriterionEnumTypeEnabled,
		ComponentCriterionEnumTypeProblem:
		*s = ComponentCriterionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ComponentCriterionEnumType", raw),
	)
}

// ConnectorEnumType (3.22)
//
// Allowed values of ConnectorCode.
//
// This enumeration does not attempt to include every possible power connector type worldwide as an
// individual type, but to specifically define those that are known to be in use (or likely to be in
// use) in the Charging Stations using the OCPP protocol. In particular, many of the very large number
// of domestic electrical sockets designs in use in many countries are excluded, unless there is
// evidence that they are or are likely to be approved for use on NOTE Charging Stations in some
// jurisdictions (e.g. as secondary connectors for charging light EVs such as electric scooters). These
// light connector types can be represented with the enumeration value Other1PhMax16A. Similarly, any
// single phase connector not otherwise enumerated that is rated for 16A or over should be reported as
// Other1PhOver16A. All 3 phase connector types not explicitly enumerated should be represented as
// Other3Ph.
type ConnectorEnumType string

const (
	// Combined Charging System 1 (captive cabled) a.k.a. Combo 1
	ConnectorEnumTypeCCCS1 ConnectorEnumType = "cCCS1"
	// Combined Charging System 2 (captive cabled) a.k.a. Combo 2
	ConnectorEnumTypeCCCS2 ConnectorEnumType = "cCCS2"
	// JARI G105-1993 (captive cabled) a.k.a. CHAdeMO
	ConnectorEnumTypeCG105 ConnectorEnumType = "cG105"
	// Tesla Connector (captive cabled)
	ConnectorEnumTypeCTesla ConnectorEnumType = "cTesla"
	// IEC62196-2 Type 1 connector (captive cabled) a.k.a. J1772
	ConnectorEnumTypeCType1 ConnectorEnumType = "cType1"
	// IEC62196-2 Type 2 connector (captive cabled) a.k.a. Mennekes connector
	ConnectorEnumTypeCType2 ConnectorEnumType = "cType2"
	// 16A 1 phase IEC60309 socket
	ConnectorEnumTypeS3091P16A ConnectorEnumType = "s309-1P-16A"
)

func (s *ConnectorEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ConnectorEnumType(raw) {
	case ConnectorEnumTypeCCCS1,
		ConnectorEnumTypeCCCS2,
		ConnectorEnumTypeCG105,
		ConnectorEnumTypeCTesla,
		ConnectorEnumTypeCType1,
		ConnectorEnumTypeCType2,
		ConnectorEnumTypeS3091P16A:
		*s = ConnectorEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ConnectorEnumType", raw),
	)
}

// ConnectorStatusEnumType (3.23)
//
// A status can be reported for the Connector of an EVSE of a Charging Station. States considered
// Operative are: Available, Reserved and Occupied. States considered Inoperative are: Unavailable,
// Faulted.
type ConnectorStatusEnumType string

const (
	// When a Connector becomes available for a new User (Operative)
	ConnectorStatusEnumTypeAvailable ConnectorStatusEnumType = "Available"
	// When a Connector becomes occupied, so it is not available for a new EV driver. (Operative)
	ConnectorStatusEnumTypeOccupied ConnectorStatusEnumType = "Occupied"
	// When a Connector becomes reserved as a result of ReserveNow command (Operative)
	ConnectorStatusEnumTypeReserved ConnectorStatusEnumType = "Reserved"
	// When a Connector becomes unavailable as the result of a Change Availability command or an event upon
	// which the Charging Station transitions to unavailable at its discretion. Upon receipt of
	// ChangeAvailability message command, the status MAY change immediately or the change MAY be
	// scheduled. When scheduled, StatusNotification SHALL be send when the availability change becomes
	// effective (Inoperative)
	ConnectorStatusEnumTypeUnavailable ConnectorStatusEnumType = "Unavailable"
	// When a Connector (or the EVSE or the entire Charging Station it belongs to) has reported an error
	// and is not available for energy delivery. (Inoperative).
	ConnectorStatusEnumTypeFaulted ConnectorStatusEnumType = "Faulted"
)

func (s *ConnectorStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ConnectorStatusEnumType(raw) {
	case ConnectorStatusEnumTypeAvailable,
		ConnectorStatusEnumTypeOccupied,
		ConnectorStatusEnumTypeReserved,
		ConnectorStatusEnumTypeUnavailable,
		ConnectorStatusEnumTypeFaulted:
		*s = ConnectorStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ConnectorStatusEnumType", raw),
	)
}

// CostKindEnumType (3.24)
type CostKindEnumType string

const (
	// Absolute value. Carbon Dioxide emissions, in grams per kWh.
	CostKindEnumTypeCarbonDioxideEmission CostKindEnumType = "CarbonDioxideEmission"
	// Relative value. Price per kWh, as percentage relative to the maximum price stated in any of all
	// tariffs indicated to the EV.
	CostKindEnumTypeRelativePricePercentage CostKindEnumType = "RelativePricePercentage"
	// Relative value. Percentage of renewable generation within total generation.
	CostKindEnumTypeRenewableGenerationPercentage CostKindEnumType = "RenewableGenerationPercentage"
)

func (s *CostKindEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CostKindEnumType(raw) {
	case CostKindEnumTypeCarbonDioxideEmission,
		CostKindEnumTypeRelativePricePercentage,
		CostKindEnumTypeRenewableGenerationPercentage:
		*s = CostKindEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CostKindEnumType", raw),
	)
}

// CustomerInformationStatusEnumType (3.25)
//
// Status in CancelReservationResponse.
type CustomerInformationStatusEnumType string

const (
	// The Charging Station accepted the message.
	CustomerInformationStatusEnumTypeAccepted CustomerInformationStatusEnumType = "Accepted"
	// When the Charging Station is in a state where it cannot process this request.
	CustomerInformationStatusEnumTypeRejected CustomerInformationStatusEnumType = "Rejected"
	// In a request to the Charging Station no reference to a customer is included.
	CustomerInformationStatusEnumTypeInvalid CustomerInformationStatusEnumType = "Invalid"
)

func (s *CustomerInformationStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CustomerInformationStatusEnumType(raw) {
	case CustomerInformationStatusEnumTypeAccepted,
		CustomerInformationStatusEnumTypeRejected,
		CustomerInformationStatusEnumTypeInvalid:
		*s = CustomerInformationStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CustomerInformationStatusEnumType", raw),
	)
}

// DataEnumType (3.26)
type DataEnumType string

const (
	// This variable is of the type string.
	DataEnumTypeString DataEnumType = "string"
	// This variable is of the type decimal.
	DataEnumTypeDecimal DataEnumType = "decimal"
	// This variable is of the type integer.
	DataEnumTypeInteger DataEnumType = "integer"
	// DateTime following the [RFC3339] specification.
	DataEnumTypeDateTime DataEnumType = "dateTime"
	// This variable is of the type boolean.
	DataEnumTypeBoolean DataEnumType = "boolean"
	// Supported/allowed values for a single choice, enumerated, text variable.
	DataEnumTypeOptionList DataEnumType = "OptionList"
	// Supported/allowed values for an ordered sequence variable.
	DataEnumTypeSequenceList DataEnumType = "SequenceList"
	// Supported/allowed values for a mathematical set variable.
	DataEnumTypeMemberList DataEnumType = "MemberList"
)

func (s *DataEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DataEnumType(raw) {
	case DataEnumTypeString,
		DataEnumTypeDecimal,
		DataEnumTypeInteger,
		DataEnumTypeDateTime,
		DataEnumTypeBoolean,
		DataEnumTypeOptionList,
		DataEnumTypeSequenceList,
		DataEnumTypeMemberList:
		*s = DataEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DataEnumType", raw),
	)
}

// DataTransferStatusEnumType (3.27)
//
// Status in DataTransferResponse.
type DataTransferStatusEnumType string

const (
	// Message has been accepted and the contained request is accepted.
	DataTransferStatusEnumTypeAccepted DataTransferStatusEnumType = "Accepted"
	// Message has been accepted but the contained request is rejected.
	DataTransferStatusEnumTypeRejected DataTransferStatusEnumType = "Rejected"
	// Message could not be interpreted due to unknown messageId string.
	DataTransferStatusEnumTypeUnknownMessageId DataTransferStatusEnumType = "UnknownMessageId"
	// Message could not be interpreted due to unknown vendorId string.
	DataTransferStatusEnumTypeUnknownVendorId DataTransferStatusEnumType = "UnknownVendorId"
)

func (s *DataTransferStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DataTransferStatusEnumType(raw) {
	case DataTransferStatusEnumTypeAccepted,
		DataTransferStatusEnumTypeRejected,
		DataTransferStatusEnumTypeUnknownMessageId,
		DataTransferStatusEnumTypeUnknownVendorId:
		*s = DataTransferStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DataTransferStatusEnumType", raw),
	)
}

// DeleteCertificateStatusEnumType (3.28)
type DeleteCertificateStatusEnumType string

const (
	// Normal successful completion (no errors).
	DeleteCertificateStatusEnumTypeAccepted DeleteCertificateStatusEnumType = "Accepted"
	// The Charging Station either failed to remove the certificate or rejected the request. A Charging
	// Station may reject the request to prevent the deletion of a certificate, if it is the last one from
	// its certificate type.
	DeleteCertificateStatusEnumTypeFailed DeleteCertificateStatusEnumType = "Failed"
	// Requested resource not found.
	DeleteCertificateStatusEnumTypeNotFound DeleteCertificateStatusEnumType = "NotFound"
)

func (s *DeleteCertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DeleteCertificateStatusEnumType(raw) {
	case DeleteCertificateStatusEnumTypeAccepted,
		DeleteCertificateStatusEnumTypeFailed,
		DeleteCertificateStatusEnumTypeNotFound:
		*s = DeleteCertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DeleteCertificateStatusEnumType", raw),
	)
}

// DisplayMessageStatusEnumType (3.29)
//
// Result for a SetDisplayMessageRequest as used in a SetDisplayMessageResponse.
type DisplayMessageStatusEnumType string

const (
	// Request to display message accepted.
	DisplayMessageStatusEnumTypeAccepted DisplayMessageStatusEnumType = "Accepted"
	// None of the formats in the given message are supported.
	DisplayMessageStatusEnumTypeNotSupportedMessageFormat DisplayMessageStatusEnumType = "NotSupportedMessageFormat"
	// Request cannot be handled.
	DisplayMessageStatusEnumTypeRejected DisplayMessageStatusEnumType = "Rejected"
	// The given MessagePriority not supported for displaying messages by Charging Station.
	DisplayMessageStatusEnumTypeNotSupportedPriority DisplayMessageStatusEnumType = "NotSupportedPriority"
	// The given MessageState not supported for displaying messages by Charging Station.
	DisplayMessageStatusEnumTypeNotSupportedState DisplayMessageStatusEnumType = "NotSupportedState"
	// Given Transaction not known/ongoing.
	DisplayMessageStatusEnumTypeUnknownTransaction DisplayMessageStatusEnumType = "UnknownTransaction"
)

func (s *DisplayMessageStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DisplayMessageStatusEnumType(raw) {
	case DisplayMessageStatusEnumTypeAccepted,
		DisplayMessageStatusEnumTypeNotSupportedMessageFormat,
		DisplayMessageStatusEnumTypeRejected,
		DisplayMessageStatusEnumTypeNotSupportedPriority,
		DisplayMessageStatusEnumTypeNotSupportedState,
		DisplayMessageStatusEnumTypeUnknownTransaction:
		*s = DisplayMessageStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DisplayMessageStatusEnumType", raw),
	)
}

// EnergyTransferModeEnumType (3.30)
//
// Enumeration of energy transfer modes.
type EnergyTransferModeEnumType string

const (
	// DC charging.
	EnergyTransferModeEnumTypeDC EnergyTransferModeEnumType = "DC"
	// AC single phase charging according to IEC 62196.
	EnergyTransferModeEnumTypeACSinglePhase EnergyTransferModeEnumType = "AC_single_phase"
	// AC two phase charging according to IEC 62196.
	EnergyTransferModeEnumTypeACTwoPhase EnergyTransferModeEnumType = "AC_two_phase"
	// AC three phase charging according to IEC 62196.
	EnergyTransferModeEnumTypeACThreePhase EnergyTransferModeEnumType = "AC_three_phase"
)

func (s *EnergyTransferModeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch EnergyTransferModeEnumType(raw) {
	case EnergyTransferModeEnumTypeDC,
		EnergyTransferModeEnumTypeACSinglePhase,
		EnergyTransferModeEnumTypeACTwoPhase,
		EnergyTransferModeEnumTypeACThreePhase:
		*s = EnergyTransferModeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid EnergyTransferModeEnumType", raw),
	)
}

// EventNotificationEnumType (3.31)
//
// Specifies the event notification type of the message.
type EventNotificationEnumType string

const (
	// The software implemented by the manufacturer triggered a hardwired notification.
	EventNotificationEnumTypeHardWiredNotification EventNotificationEnumType = "HardWiredNotification"
	// Triggered by a monitor, which is hardwired by the manufacturer.
	EventNotificationEnumTypeHardWiredMonitor EventNotificationEnumType = "HardWiredMonitor"
	// Triggered by a monitor, which is preconfigured by the manufacturer.
	EventNotificationEnumTypePreconfiguredMonitor EventNotificationEnumType = "PreconfiguredMonitor"
	// Triggered by a monitor, which is set with the setvariablemonitoringrequest message by the Charging
	// Station Operator.
	EventNotificationEnumTypeCustomMonitor EventNotificationEnumType = "CustomMonitor"
)

func (s *EventNotificationEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch EventNotificationEnumType(raw) {
	case EventNotificationEnumTypeHardWiredNotification,
		EventNotificationEnumTypeHardWiredMonitor,
		EventNotificationEnumTypePreconfiguredMonitor,
		EventNotificationEnumTypeCustomMonitor:
		*s = EventNotificationEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid EventNotificationEnumType", raw),
	)
}

// EventTriggerEnumType (3.32)
type EventTriggerEnumType string

const (
	// Monitored variable has passed an Lower or Upper Threshold. Also used as trigger type for a
	// HardWiredNotification.
	EventTriggerEnumTypeAlerting EventTriggerEnumType = "Alerting"
	// Delta Monitored Variable value has changed by more than specified amount
	EventTriggerEnumTypeDelta EventTriggerEnumType = "Delta"
	// Periodic Monitored Variable has been sampled for reporting at the specified interval
	EventTriggerEnumTypePeriodic EventTriggerEnumType = "Periodic"
)

func (s *EventTriggerEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch EventTriggerEnumType(raw) {
	case EventTriggerEnumTypeAlerting,
		EventTriggerEnumTypeDelta,
		EventTriggerEnumTypePeriodic:
		*s = EventTriggerEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid EventTriggerEnumType", raw),
	)
}

// FirmwareStatusEnumType (3.33)
//
// Status of a firmware download.
//
// A value with "Intermediate state" in the description, is an intermediate state, update process is
// not finished.
//
// A value with "Failure end state" in the description, is an end state, update process has stopped,
// update failed.
//
// A value with "Successful end state" in the description, is an end state, update process has stopped,
// update successful.
type FirmwareStatusEnumType string

const (
	// Intermediate state. New firmware has been downloaded by Charging Station.
	FirmwareStatusEnumTypeDownloaded FirmwareStatusEnumType = "Downloaded"
	// Failure end state. Charging Station failed to download firmware.
	FirmwareStatusEnumTypeDownloadFailed FirmwareStatusEnumType = "DownloadFailed"
	// Intermediate state. Firmware is being downloaded.
	FirmwareStatusEnumTypeDownloading FirmwareStatusEnumType = "Downloading"
	// Intermediate state. Downloading of new firmware has been scheduled.
	FirmwareStatusEnumTypeDownloadScheduled FirmwareStatusEnumType = "DownloadScheduled"
	// Intermediate state. Downloading has been paused.
	FirmwareStatusEnumTypeDownloadPaused FirmwareStatusEnumType = "DownloadPaused"
	// Charging Station is not performing firmware update related tasks. Status Idle SHALL only be used as
	// in a FirmwareStatusNotificationRequest that was triggered by TriggerMessageRequest.
	FirmwareStatusEnumTypeIdle FirmwareStatusEnumType = "Idle"
	// Failure end state. Installation of new firmware has failed.
	FirmwareStatusEnumTypeInstallationFailed FirmwareStatusEnumType = "InstallationFailed"
	// Intermediate state. Firmware is being installed.
	FirmwareStatusEnumTypeInstalling FirmwareStatusEnumType = "Installing"
	// Successful end state. New firmware has successfully been installed in Charging Station.
	FirmwareStatusEnumTypeInstalled FirmwareStatusEnumType = "Installed"
	// Intermediate state. Charging Station is about to reboot to activate new firmware. If sent before
	// installing the firmware, it indicates the Charging Station is about to reboot to start installing
	// new firmware. If sent after installing the new firmware, it indicates the Charging Station has
	// finished installing, but requires a reboot to activate the new firmware, which will be done
	// automatically when idle. This status MAY be omitted if a reboot is an integral part of the
	// installation and cannot be reported separately.
	FirmwareStatusEnumTypeInstallRebooting FirmwareStatusEnumType = "InstallRebooting"
	// Intermediate state. Installation of the downloaded firmware is scheduled to take place on
	// installDateTime given in UpdateFirmware request.
	FirmwareStatusEnumTypeInstallScheduled FirmwareStatusEnumType = "InstallScheduled"
	// Failure end state. Verification of the new firmware (e.g. using a checksum or some other means) has
	// failed and installation will not proceed. (Final failure state)
	FirmwareStatusEnumTypeInstallVerificationFailed FirmwareStatusEnumType = "InstallVerificationFailed"
	// Failure end state. The firmware signature is not valid.
	FirmwareStatusEnumTypeInvalidSignature FirmwareStatusEnumType = "InvalidSignature"
	// Intermediate state. Provide signature successfully verified.
	FirmwareStatusEnumTypeSignatureVerified FirmwareStatusEnumType = "SignatureVerified"
)

func (s *FirmwareStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch FirmwareStatusEnumType(raw) {
	case FirmwareStatusEnumTypeDownloaded,
		FirmwareStatusEnumTypeDownloadFailed,
		FirmwareStatusEnumTypeDownloading,
		FirmwareStatusEnumTypeDownloadScheduled,
		FirmwareStatusEnumTypeDownloadPaused,
		FirmwareStatusEnumTypeIdle,
		FirmwareStatusEnumTypeInstallationFailed,
		FirmwareStatusEnumTypeInstalling,
		FirmwareStatusEnumTypeInstalled,
		FirmwareStatusEnumTypeInstallRebooting,
		FirmwareStatusEnumTypeInstallScheduled,
		FirmwareStatusEnumTypeInstallVerificationFailed,
		FirmwareStatusEnumTypeInvalidSignature,
		FirmwareStatusEnumTypeSignatureVerified:
		*s = FirmwareStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid FirmwareStatusEnumType", raw),
	)
}

// GenericDeviceModelStatusEnumType (3.34)
type GenericDeviceModelStatusEnumType string

const (
	// Request has been accepted and will be executed.
	GenericDeviceModelStatusEnumTypeAccepted GenericDeviceModelStatusEnumType = "Accepted"
	// Request has not been accepted and will not be executed.
	GenericDeviceModelStatusEnumTypeRejected GenericDeviceModelStatusEnumType = "Rejected"
	// The content of the request message is not supported.
	GenericDeviceModelStatusEnumTypeNotSupported GenericDeviceModelStatusEnumType = "NotSupported"
	// If the combination of received criteria result in an empty result set.
	GenericDeviceModelStatusEnumTypeEmptyResultSet GenericDeviceModelStatusEnumType = "EmptyResultSet"
)

func (s *GenericDeviceModelStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GenericDeviceModelStatusEnumType(raw) {
	case GenericDeviceModelStatusEnumTypeAccepted,
		GenericDeviceModelStatusEnumTypeRejected,
		GenericDeviceModelStatusEnumTypeNotSupported,
		GenericDeviceModelStatusEnumTypeEmptyResultSet:
		*s = GenericDeviceModelStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GenericDeviceModelStatusEnumType", raw),
	)
}

// GenericStatusEnumType (3.35)
//
// Generic message response status
type GenericStatusEnumType string

const (
	// Request has been accepted and will be executed.
	GenericStatusEnumTypeAccepted GenericStatusEnumType = "Accepted"
	// Request has not been accepted and will not be executed.
	GenericStatusEnumTypeRejected GenericStatusEnumType = "Rejected"
)

func (s *GenericStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GenericStatusEnumType(raw) {
	case GenericStatusEnumTypeAccepted,
		GenericStatusEnumTypeRejected:
		*s = GenericStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GenericStatusEnumType", raw),
	)
}

// GetCertificateIdUseEnumType (3.36)
type GetCertificateIdUseEnumType string

const (
	// Use for certificate of the V2G Root.
	GetCertificateIdUseEnumTypeV2GRootCertificate GetCertificateIdUseEnumType = "V2GRootCertificate"
	// Use for certificate from an eMobility Service provider. To support PnC charging with contracts from
	// service providers that not derived their certificates from the V2G root.
	GetCertificateIdUseEnumTypeMORootCertificate GetCertificateIdUseEnumType = "MORootCertificate"
	// Root certificate for verification of the CSMS certificate.
	GetCertificateIdUseEnumTypeCSMSRootCertificate GetCertificateIdUseEnumType = "CSMSRootCertificate"
	// ISO 15118 V2G certificate chain (excluding the V2GRootCertificate).
	GetCertificateIdUseEnumTypeV2GCertificateChain GetCertificateIdUseEnumType = "V2GCertificateChain"
	// Root certificate for verification of the Manufacturer certificate.
	GetCertificateIdUseEnumTypeManufacturerRootCertificate GetCertificateIdUseEnumType = "ManufacturerRootCertificate"
)

func (s *GetCertificateIdUseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetCertificateIdUseEnumType(raw) {
	case GetCertificateIdUseEnumTypeV2GRootCertificate,
		GetCertificateIdUseEnumTypeMORootCertificate,
		GetCertificateIdUseEnumTypeCSMSRootCertificate,
		GetCertificateIdUseEnumTypeV2GCertificateChain,
		GetCertificateIdUseEnumTypeManufacturerRootCertificate:
		*s = GetCertificateIdUseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetCertificateIdUseEnumType", raw),
	)
}

// GetCertificateStatusEnumType (3.37)
type GetCertificateStatusEnumType string

const (
	// Successfully retrieved the OCSP certificate status.
	GetCertificateStatusEnumTypeAccepted GetCertificateStatusEnumType = "Accepted"
	// Failed to retrieve the OCSP certificate status.
	GetCertificateStatusEnumTypeFailed GetCertificateStatusEnumType = "Failed"
)

func (s *GetCertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetCertificateStatusEnumType(raw) {
	case GetCertificateStatusEnumTypeAccepted,
		GetCertificateStatusEnumTypeFailed:
		*s = GetCertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetCertificateStatusEnumType", raw),
	)
}

// GetChargingProfileStatusEnumType (3.38)
type GetChargingProfileStatusEnumType string

const (
	// Normal successful completion (no errors).
	GetChargingProfileStatusEnumTypeAccepted GetChargingProfileStatusEnumType = "Accepted"
	// No ChargingProfiles found that match the information in the GetChargingProfilesRequest.
	GetChargingProfileStatusEnumTypeNoProfiles GetChargingProfileStatusEnumType = "NoProfiles"
)

func (s *GetChargingProfileStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetChargingProfileStatusEnumType(raw) {
	case GetChargingProfileStatusEnumTypeAccepted,
		GetChargingProfileStatusEnumTypeNoProfiles:
		*s = GetChargingProfileStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetChargingProfileStatusEnumType", raw),
	)
}

// GetDisplayMessagesStatusEnumType (3.39)
type GetDisplayMessagesStatusEnumType string

const (
	// Request accepted, there are Display Messages found that match all the requested criteria. The
	// Charging Station will send NotifyDisplayMessagesRequest messages to report the requested Display
	// Messages.
	GetDisplayMessagesStatusEnumTypeAccepted GetDisplayMessagesStatusEnumType = "Accepted"
	// No messages found that match the given criteria.
	GetDisplayMessagesStatusEnumTypeUnknown GetDisplayMessagesStatusEnumType = "Unknown"
)

func (s *GetDisplayMessagesStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetDisplayMessagesStatusEnumType(raw) {
	case GetDisplayMessagesStatusEnumTypeAccepted,
		GetDisplayMessagesStatusEnumTypeUnknown:
		*s = GetDisplayMessagesStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetDisplayMessagesStatusEnumType", raw),
	)
}

// GetInstalledCertificateStatusEnumType (3.40)
type GetInstalledCertificateStatusEnumType string

const (
	// Normal successful completion (no errors).
	GetInstalledCertificateStatusEnumTypeAccepted GetInstalledCertificateStatusEnumType = "Accepted"
	// Requested resource not found.
	GetInstalledCertificateStatusEnumTypeNotFound GetInstalledCertificateStatusEnumType = "NotFound"
)

func (s *GetInstalledCertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetInstalledCertificateStatusEnumType(raw) {
	case GetInstalledCertificateStatusEnumTypeAccepted,
		GetInstalledCertificateStatusEnumTypeNotFound:
		*s = GetInstalledCertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetInstalledCertificateStatusEnumType", raw),
	)
}

// GetVariableStatusEnumType (3.41)
type GetVariableStatusEnumType string

const (
	// Variable successfully set.
	GetVariableStatusEnumTypeAccepted GetVariableStatusEnumType = "Accepted"
	// Request is rejected.
	GetVariableStatusEnumTypeRejected GetVariableStatusEnumType = "Rejected"
	// Component is not known.
	GetVariableStatusEnumTypeUnknownComponent GetVariableStatusEnumType = "UnknownComponent"
	// Variable is not known.
	GetVariableStatusEnumTypeUnknownVariable GetVariableStatusEnumType = "UnknownVariable"
	// The AttributeType is not supported.
	GetVariableStatusEnumTypeNotSupportedAttributeType GetVariableStatusEnumType = "NotSupportedAttributeType"
)

func (s *GetVariableStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GetVariableStatusEnumType(raw) {
	case GetVariableStatusEnumTypeAccepted,
		GetVariableStatusEnumTypeRejected,
		GetVariableStatusEnumTypeUnknownComponent,
		GetVariableStatusEnumTypeUnknownVariable,
		GetVariableStatusEnumTypeNotSupportedAttributeType:
		*s = GetVariableStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetVariableStatusEnumType", raw),
	)
}

// HashAlgorithmEnumType (3.42)
type HashAlgorithmEnumType string

const (
	// SHA-256 hash algorithm.
	HashAlgorithmEnumTypeSHA256 HashAlgorithmEnumType = "SHA256"
	// SHA-384 hash algorithm.
	HashAlgorithmEnumTypeSHA384 HashAlgorithmEnumType = "SHA384"
	// SHA-512 hash algorithm.
	HashAlgorithmEnumTypeSHA512 HashAlgorithmEnumType = "SHA512"
)

func (s *HashAlgorithmEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch HashAlgorithmEnumType(raw) {
	case HashAlgorithmEnumTypeSHA256,
		HashAlgorithmEnumTypeSHA384,
		HashAlgorithmEnumTypeSHA512:
		*s = HashAlgorithmEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid HashAlgorithmEnumType", raw),
	)
}

// IdTokenEnumType (3.43)
//
// Allowable values of the IdTokenType field.
type IdTokenEnumType string

const (
	// A centrally, in the CSMS (or other server) generated id (for example used for a remotely started
	// transaction that is activated by SMS). No format defined, might be a UUID.
	IdTokenEnumTypeCentral IdTokenEnumType = "Central"
	// Electro-mobility account id as defined in ISO 15118
	IdTokenEnumTypeEMAID IdTokenEnumType = "eMAID"
	// ISO 14443 UID of RFID card. It is represented as an array of 4 or 7 bytes in hexadecimal
	// representation.
	IdTokenEnumTypeISO14443 IdTokenEnumType = "ISO14443"
	// ISO 15693 UID of RFID card. It is represented as an array of 8 bytes in hexadecimal representation.
	IdTokenEnumTypeISO15693 IdTokenEnumType = "ISO15693"
	// User use a private key-code to authorize a charging transaction. For example: Pin-code.
	IdTokenEnumTypeKeyCode IdTokenEnumType = "KeyCode"
	// A locally generated id (e.g. internal id created by the Charging Station). No format defined, might
	// be a UUID
	IdTokenEnumTypeLocal IdTokenEnumType = "Local"
	// The MacAddress of the EVCC (Electric Vehicle Communication Controller) that is connected to the
	// EVSE. This is used as a token type when the MAC address is used for authorization ("Autocharge").
	IdTokenEnumTypeMacAddress IdTokenEnumType = "MacAddress"
	// Transaction is started and no authorization possible. Charging Station only has a start button or
	// mechanical key etc. IdToken field SHALL be left empty.
	IdTokenEnumTypeNoAuthorization IdTokenEnumType = "NoAuthorization"
)

func (s *IdTokenEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch IdTokenEnumType(raw) {
	case IdTokenEnumTypeCentral,
		IdTokenEnumTypeEMAID,
		IdTokenEnumTypeISO14443,
		IdTokenEnumTypeISO15693,
		IdTokenEnumTypeKeyCode,
		IdTokenEnumTypeLocal,
		IdTokenEnumTypeMacAddress,
		IdTokenEnumTypeNoAuthorization:
		*s = IdTokenEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid IdTokenEnumType", raw),
	)
}

// InstallCertificateStatusEnumType (3.44)
type InstallCertificateStatusEnumType string

const (
	// The installation of the certificate succeeded.
	InstallCertificateStatusEnumTypeAccepted InstallCertificateStatusEnumType = "Accepted"
	// The certificate is invalid and/or incorrect OR the CSO tries to install more certificates than
	// allowed.
	InstallCertificateStatusEnumTypeRejected InstallCertificateStatusEnumType = "Rejected"
	// The certificate is valid and correct, but there is another reason the installation did not succeed.
	InstallCertificateStatusEnumTypeFailed InstallCertificateStatusEnumType = "Failed"
)

func (s *InstallCertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch InstallCertificateStatusEnumType(raw) {
	case InstallCertificateStatusEnumTypeAccepted,
		InstallCertificateStatusEnumTypeRejected,
		InstallCertificateStatusEnumTypeFailed:
		*s = InstallCertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid InstallCertificateStatusEnumType", raw),
	)
}

// InstallCertificateUseEnumType (3.45)
type InstallCertificateUseEnumType string

const (
	// Use for certificate of the V2G Root, a V2G Charging Station Certificate MUST be derived from one of
	// the installed V2GRootCertificate certificates.
	InstallCertificateUseEnumTypeV2GRootCertificate InstallCertificateUseEnumType = "V2GRootCertificate"
	// Use for certificate from an eMobility Service provider. To support PnC charging with contracts from
	// service providers that not derived their certificates from the V2G root.
	InstallCertificateUseEnumTypeMORootCertificate InstallCertificateUseEnumType = "MORootCertificate"
	// Root certificate, used by the CA to sign the CSMS and Charging Station certificate.
	InstallCertificateUseEnumTypeCSMSRootCertificate InstallCertificateUseEnumType = "CSMSRootCertificate"
	// Root certificate for verification of the Manufacturer certificate.
	InstallCertificateUseEnumTypeManufacturerRootCertificate InstallCertificateUseEnumType = "ManufacturerRootCertificate"
)

func (s *InstallCertificateUseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch InstallCertificateUseEnumType(raw) {
	case InstallCertificateUseEnumTypeV2GRootCertificate,
		InstallCertificateUseEnumTypeMORootCertificate,
		InstallCertificateUseEnumTypeCSMSRootCertificate,
		InstallCertificateUseEnumTypeManufacturerRootCertificate:
		*s = InstallCertificateUseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid InstallCertificateUseEnumType", raw),
	)
}

// Iso15118EVCertificateStatusEnumType (3.46)
type Iso15118EVCertificateStatusEnumType string

const (
	// exiResponse included. This is no indication whether the update was successful, just that the message
	// was processed properly.
	Iso15118EVCertificateStatusEnumTypeAccepted Iso15118EVCertificateStatusEnumType = "Accepted"
	// Processing of the message was not successful, no exiResponse included.
	Iso15118EVCertificateStatusEnumTypeFailed Iso15118EVCertificateStatusEnumType = "Failed"
)

func (s *Iso15118EVCertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch Iso15118EVCertificateStatusEnumType(raw) {
	case Iso15118EVCertificateStatusEnumTypeAccepted,
		Iso15118EVCertificateStatusEnumTypeFailed:
		*s = Iso15118EVCertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid Iso15118EVCertificateStatusEnumType", raw),
	)
}

// LocationEnumType (3.47)
//
// Allowable values of the optional "location" field of a value element.
type LocationEnumType string

const (
	// Measurement inside body of Charging Station (e.g. Temperature).
	LocationEnumTypeBody LocationEnumType = "Body"
	// Measurement taken from cable between EV and Charging Station.
	LocationEnumTypeCable LocationEnumType = "Cable"
	// Measurement taken by EV.
	LocationEnumTypeEV LocationEnumType = "EV"
	// For the Charging Station (evseId = 0): measurement at network ("grid") inlet connection. For
	// measurements with evseId > 0, these are measurements taken at the EVSE inlet (This can be useful for
	// a DC charger).
	LocationEnumTypeInlet LocationEnumType = "Inlet"
	// Measurement at a Connector. Default value.
	LocationEnumTypeOutlet LocationEnumType = "Outlet"
)

func (s *LocationEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch LocationEnumType(raw) {
	case LocationEnumTypeBody,
		LocationEnumTypeCable,
		LocationEnumTypeEV,
		LocationEnumTypeInlet,
		LocationEnumTypeOutlet:
		*s = LocationEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid LocationEnumType", raw),
	)
}

// LogEnumType (3.48)
type LogEnumType string

const (
	// This contains the field definition of a diagnostics log file
	LogEnumTypeDiagnosticsLog LogEnumType = "DiagnosticsLog"
	// Sent by the CSMS to the Charging Station to request that the Charging Station uploads the security
	// log.
	LogEnumTypeSecurityLog LogEnumType = "SecurityLog"
)

func (s *LogEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch LogEnumType(raw) {
	case LogEnumTypeDiagnosticsLog,
		LogEnumTypeSecurityLog:
		*s = LogEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid LogEnumType", raw),
	)
}

// LogStatusEnumType (3.49)
//
// Generic message response status
type LogStatusEnumType string

const (
	// Accepted this log upload. This does not mean the log file is uploaded is successfully, the Charging
	// Station will now start the log file upload.
	LogStatusEnumTypeAccepted LogStatusEnumType = "Accepted"
	// Log update request rejected.
	LogStatusEnumTypeRejected LogStatusEnumType = "Rejected"
	// Accepted this log upload, but in doing this has canceled an ongoing log file upload.
	LogStatusEnumTypeAcceptedCanceled LogStatusEnumType = "AcceptedCanceled"
)

func (s *LogStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch LogStatusEnumType(raw) {
	case LogStatusEnumTypeAccepted,
		LogStatusEnumTypeRejected,
		LogStatusEnumTypeAcceptedCanceled:
		*s = LogStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid LogStatusEnumType", raw),
	)
}

// MeasurandEnumType (3.50)
//
// Allowable values of the optional "measurand" field of a Value element, as used in MeterValuesRequest
// and TransactionEventRequest with eventTypes Started, Ended and Updated. Default value of "measurand"
// is always "Energy.Active.Import.Register".
//
// Note 1: Two measurands (Current.Offered and Power.Offered) are available that are strictly speaking
// no measured values. They indicate the maximum amount of current/power that is being offered to the
// EV and are intended for use in smart charging applications.
//
// Note 2: Import is energy flow from the Grid to the Charging Station, EV or other load. Export is
// energy flow from the EV to the Charging Station and/or from the Charging Station to the Grid. Except
// in the case of a meter replacement, all "Register" values relating to a single charging transaction,
// or a non-transactional consumer (e.g. Charging Station internal power supply, overall supply) MUST
// be monotonically increasing in time.
//
// Note 3: The actual quantity of energy corresponding to a reported ".Register" value is computed as
// the register value in question minus the register value recorded/reported at the start of the
// transaction or other relevant starting reference point in time. For improved auditability,
// ".Register" values SHOULD be reported exactly as they are directly read from a non-volatile register
// in the electrical metering hardware, and SHOULD NOT be re-based to zero at the start of
// transactions. This allows any "missing energy" between sequential transactions, due to hardware
// fault, meter replacement, mis-wiring, fraud, etc. to be identified, by allowing the CSMS to confirm
// that the starting register value of any transaction is identical to the finishing register value of
// the preceding transaction on the same connector.
type MeasurandEnumType string

const (
	// Instantaneous current flow from EV
	MeasurandEnumTypeCurrentExport MeasurandEnumType = "Current.Export"
	// Instantaneous current flow to EV
	MeasurandEnumTypeCurrentImport MeasurandEnumType = "Current.Import"
	// Maximum current offered to EV
	MeasurandEnumTypeCurrentOffered MeasurandEnumType = "Current.Offered"
	// Numerical value read from the "active electrical energy" (Wh or kWh) register of the (most
	// authoritative) electrical meter measuring energy exported (to the grid).
	MeasurandEnumTypeEnergyActiveExportRegister MeasurandEnumType = "Energy.Active.Export.Register"
	// Numerical value read from the "active electrical energy" (Wh or kWh) register of the (most
	// authoritative) electrical meter measuring energy imported (from the grid supply).
	MeasurandEnumTypeEnergyActiveImportRegister MeasurandEnumType = "Energy.Active.Import.Register"
)

func (s *MeasurandEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MeasurandEnumType(raw) {
	case MeasurandEnumTypeCurrentExport,
		MeasurandEnumTypeCurrentImport,
		MeasurandEnumTypeCurrentOffered,
		MeasurandEnumTypeEnergyActiveExportRegister,
		MeasurandEnumTypeEnergyActiveImportRegister:
		*s = MeasurandEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MeasurandEnumType", raw),
	)
}

// MessageFormatEnumType (3.51)
//
// Format of a message to be displayed on the display of the Charging Station.
type MessageFormatEnumType string

const (
	// Message content is ASCII formatted, only printable ASCII allowed.
	MessageFormatEnumTypeASCII MessageFormatEnumType = "ASCII"
	// Message content is HTML formatted.
	MessageFormatEnumTypeHTML MessageFormatEnumType = "HTML"
	// Message content is URI that Charging Station should download and use to display. for example a HTML
	// page to be shown in a web-browser.
	MessageFormatEnumTypeURI MessageFormatEnumType = "URI"
	// Message content is UTF-8 formatted.
	MessageFormatEnumTypeUTF8 MessageFormatEnumType = "UTF8"
)

func (s *MessageFormatEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MessageFormatEnumType(raw) {
	case MessageFormatEnumTypeASCII,
		MessageFormatEnumTypeHTML,
		MessageFormatEnumTypeURI,
		MessageFormatEnumTypeUTF8:
		*s = MessageFormatEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageFormatEnumType", raw),
	)
}

// MessagePriorityEnumType (3.52)
//
// Priority with which a message should be displayed on a Charging Station.
type MessagePriorityEnumType string

const (
	// Show this message always in front. Highest priority, don’t cycle with other messages. When a newer
	// message with this MessagePriority is received, this message is replaced. No Charging Station own
	// message may override this message.
	MessagePriorityEnumTypeAlwaysFront MessagePriorityEnumType = "AlwaysFront"
	// Show this message in front of the normal cycle of messages. When more messages with this priority
	// are to be shown, they SHALL be cycled.
	MessagePriorityEnumTypeInFront MessagePriorityEnumType = "InFront"
	// Show this message in the cycle of messages.
	MessagePriorityEnumTypeNormalCycle MessagePriorityEnumType = "NormalCycle"
)

func (s *MessagePriorityEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MessagePriorityEnumType(raw) {
	case MessagePriorityEnumTypeAlwaysFront,
		MessagePriorityEnumTypeInFront,
		MessagePriorityEnumTypeNormalCycle:
		*s = MessagePriorityEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessagePriorityEnumType", raw),
	)
}

// MessageStateEnumType (3.53)
//
// State of the Charging Station during which a message SHALL be displayed.
type MessageStateEnumType string

const (
	// Message only to be shown while the Charging Station is charging.
	MessageStateEnumTypeCharging MessageStateEnumType = "Charging"
	// Message only to be shown while the Charging Station is in faulted state.
	MessageStateEnumTypeFaulted MessageStateEnumType = "Faulted"
	// Message only to be shown while the Charging Station is idle (not charging).
	MessageStateEnumTypeIdle MessageStateEnumType = "Idle"
	// Message only to be shown while the Charging Station is in unavailable state.
	MessageStateEnumTypeUnavailable MessageStateEnumType = "Unavailable"
)

func (s *MessageStateEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MessageStateEnumType(raw) {
	case MessageStateEnumTypeCharging,
		MessageStateEnumTypeFaulted,
		MessageStateEnumTypeIdle,
		MessageStateEnumTypeUnavailable:
		*s = MessageStateEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageStateEnumType", raw),
	)
}

// MessageTriggerEnumType (3.54)
//
// Type of request to be triggered by trigger messages.
type MessageTriggerEnumType string

const (
	// To trigger BootNotification.
	MessageTriggerEnumTypeBootNotification MessageTriggerEnumType = "BootNotification"
	// To trigger LogStatusNotification.
	MessageTriggerEnumTypeLogStatusNotification MessageTriggerEnumType = "LogStatusNotification"
	// To trigger FirmwareStatusNotification.
	MessageTriggerEnumTypeFirmwareStatusNotification MessageTriggerEnumType = "FirmwareStatusNotification"
	// To trigger Heartbeat.
	MessageTriggerEnumTypeHeartbeat MessageTriggerEnumType = "Heartbeat"
	// To trigger MeterValues.
	MessageTriggerEnumTypeMeterValues MessageTriggerEnumType = "MeterValues"
	// To trigger a SignCertificate with certificateType: ChargingStationCertificate.
	MessageTriggerEnumTypeSignChargingStationCertificate MessageTriggerEnumType = "SignChargingStationCertificate"
	// To trigger a SignCertificate with certificateType: V2GCertificate
	MessageTriggerEnumTypeSignV2GCertificate MessageTriggerEnumType = "SignV2GCertificate"
	// To trigger StatusNotification.
	MessageTriggerEnumTypeStatusNotification MessageTriggerEnumType = "StatusNotification"
	// To trigger TransactionEvent.
	MessageTriggerEnumTypeTransactionEvent MessageTriggerEnumType = "TransactionEvent"
	// To trigger a SignCertificate with certificateType: ChargingStationCertificate AND V2GCertificate
	MessageTriggerEnumTypeSignCombinedCertificate MessageTriggerEnumType = "SignCombinedCertificate"
	// To trigger PublishFirmwareStatusNotification.
	MessageTriggerEnumTypePublishFirmwareStatusNotification MessageTriggerEnumType = "PublishFirmwareStatusNotification"
)

func (s *MessageTriggerEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MessageTriggerEnumType(raw) {
	case MessageTriggerEnumTypeBootNotification,
		MessageTriggerEnumTypeLogStatusNotification,
		MessageTriggerEnumTypeFirmwareStatusNotification,
		MessageTriggerEnumTypeHeartbeat,
		MessageTriggerEnumTypeMeterValues,
		MessageTriggerEnumTypeSignChargingStationCertificate,
		MessageTriggerEnumTypeSignV2GCertificate,
		MessageTriggerEnumTypeStatusNotification,
		MessageTriggerEnumTypeTransactionEvent,
		MessageTriggerEnumTypeSignCombinedCertificate,
		MessageTriggerEnumTypePublishFirmwareStatusNotification:
		*s = MessageTriggerEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageTriggerEnumType", raw),
	)
}

// MonitorEnumType (3.55)
type MonitorEnumType string

const (
	// Triggers an event notice when the actual value of the Variable rises above monitorValue
	MonitorEnumTypeUpperThreshold MonitorEnumType = "UpperThreshold"
	// Triggers an event notice when the actual value of the Variable drops below monitorValue.
	MonitorEnumTypeLowerThreshold MonitorEnumType = "LowerThreshold"
	// Triggers an event notice when the actual value has changed more than plus or minus monitorValue
	// since the time that this monitor was set or since the last time this event notice was sent,
	// whichever was last. For variables that are not numeric, like boolean, string or enumerations, a
	// monitor of type Delta will trigger an event notice whenever the variable changes, regardless of the
	// value of monitorValue.
	MonitorEnumTypeDelta MonitorEnumType = "Delta"
	// Triggers an event notice every monitorValue seconds interval, starting from the time that this
	// monitor was set.
	MonitorEnumTypePeriodic MonitorEnumType = "Periodic"
	// Triggers an event notice every monitorValue seconds interval, starting from the nearest
	// clock-aligned interval after this monitor was set. For example, a monitorValue of 900 will trigger
	// event notices at 0, 15, 30 and 45 minutes after the hour, every hour.
	MonitorEnumTypePeriodicClockAligned MonitorEnumType = "PeriodicClockAligned"
)

func (s *MonitorEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MonitorEnumType(raw) {
	case MonitorEnumTypeUpperThreshold,
		MonitorEnumTypeLowerThreshold,
		MonitorEnumTypeDelta,
		MonitorEnumTypePeriodic,
		MonitorEnumTypePeriodicClockAligned:
		*s = MonitorEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MonitorEnumType", raw),
	)
}

// MonitoringBaseEnumType (3.56)
type MonitoringBaseEnumType string

const (
	// Activate all pre-configured monitors.
	MonitoringBaseEnumTypeAll MonitoringBaseEnumType = "All"
	// Activate the default monitoring settings as recommended by the manufacturer. This is a subset of all
	// pre- configured monitors.
	MonitoringBaseEnumTypeFactoryDefault MonitoringBaseEnumType = "FactoryDefault"
	// Clears all custom monitors and disables all pre-configured monitors.
	MonitoringBaseEnumTypeHardWiredOnly MonitoringBaseEnumType = "HardWiredOnly"
)

func (s *MonitoringBaseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MonitoringBaseEnumType(raw) {
	case MonitoringBaseEnumTypeAll,
		MonitoringBaseEnumTypeFactoryDefault,
		MonitoringBaseEnumTypeHardWiredOnly:
		*s = MonitoringBaseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MonitoringBaseEnumType", raw),
	)
}

// MonitoringCriterionEnumType (3.57)
type MonitoringCriterionEnumType string

const (
	// Report variables and components with a monitor of type UpperThreshold or LowerThreshold.
	MonitoringCriterionEnumTypeThresholdMonitoring MonitoringCriterionEnumType = "ThresholdMonitoring"
	// Report variables and components with a monitor of type Delta.
	MonitoringCriterionEnumTypeDeltaMonitoring MonitoringCriterionEnumType = "DeltaMonitoring"
	// Report variables and components with a monitor of type Periodic or PeriodicClockAligned.
	MonitoringCriterionEnumTypePeriodicMonitoring MonitoringCriterionEnumType = "PeriodicMonitoring"
)

func (s *MonitoringCriterionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MonitoringCriterionEnumType(raw) {
	case MonitoringCriterionEnumTypeThresholdMonitoring,
		MonitoringCriterionEnumTypeDeltaMonitoring,
		MonitoringCriterionEnumTypePeriodicMonitoring:
		*s = MonitoringCriterionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MonitoringCriterionEnumType", raw),
	)
}

// MutabilityEnumType (3.58)
type MutabilityEnumType string

const (
	// This variable is read-only.
	MutabilityEnumTypeReadOnly MutabilityEnumType = "ReadOnly"
	// This variable is write-only.
	MutabilityEnumTypeWriteOnly MutabilityEnumType = "WriteOnly"
	// This variable is read-write.
	MutabilityEnumTypeReadWrite MutabilityEnumType = "ReadWrite"
)

func (s *MutabilityEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MutabilityEnumType(raw) {
	case MutabilityEnumTypeReadOnly,
		MutabilityEnumTypeWriteOnly,
		MutabilityEnumTypeReadWrite:
		*s = MutabilityEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MutabilityEnumType", raw),
	)
}

// NotifyEVChargingNeedsStatusEnumType (3.59)
type NotifyEVChargingNeedsStatusEnumType string

const (
	// A schedule will be provided momentarily.
	NotifyEVChargingNeedsStatusEnumTypeAccepted NotifyEVChargingNeedsStatusEnumType = "Accepted"
	// Service not available.
	NotifyEVChargingNeedsStatusEnumTypeRejected NotifyEVChargingNeedsStatusEnumType = "Rejected"
	// The CSMS is gathering information to provide a schedule.
	NotifyEVChargingNeedsStatusEnumTypeProcessing NotifyEVChargingNeedsStatusEnumType = "Processing"
)

func (s *NotifyEVChargingNeedsStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch NotifyEVChargingNeedsStatusEnumType(raw) {
	case NotifyEVChargingNeedsStatusEnumTypeAccepted,
		NotifyEVChargingNeedsStatusEnumTypeRejected,
		NotifyEVChargingNeedsStatusEnumTypeProcessing:
		*s = NotifyEVChargingNeedsStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid NotifyEVChargingNeedsStatusEnumType", raw),
	)
}

// OCPPInterfaceEnumType (3.60)
//
// Enumeration of network interfaces.
type OCPPInterfaceEnumType string

const (
	// Use wired connection 0
	OCPPInterfaceEnumTypeWired0 OCPPInterfaceEnumType = "Wired0"
	// Use wired connection 1
	OCPPInterfaceEnumTypeWired1 OCPPInterfaceEnumType = "Wired1"
	// Use wired connection 2
	OCPPInterfaceEnumTypeWired2 OCPPInterfaceEnumType = "Wired2"
	// Use wired connection 3
	OCPPInterfaceEnumTypeWired3 OCPPInterfaceEnumType = "Wired3"
	// Use wireless connection 0
	OCPPInterfaceEnumTypeWireless0 OCPPInterfaceEnumType = "Wireless0"
	// Use wireless connection 1
	OCPPInterfaceEnumTypeWireless1 OCPPInterfaceEnumType = "Wireless1"
	// Use wireless connection 2
	OCPPInterfaceEnumTypeWireless2 OCPPInterfaceEnumType = "Wireless2"
	// Use wireless connection 3
	OCPPInterfaceEnumTypeWireless3 OCPPInterfaceEnumType = "Wireless3"
)

func (s *OCPPInterfaceEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OCPPInterfaceEnumType(raw) {
	case OCPPInterfaceEnumTypeWired0,
		OCPPInterfaceEnumTypeWired1,
		OCPPInterfaceEnumTypeWired2,
		OCPPInterfaceEnumTypeWired3,
		OCPPInterfaceEnumTypeWireless0,
		OCPPInterfaceEnumTypeWireless1,
		OCPPInterfaceEnumTypeWireless2,
		OCPPInterfaceEnumTypeWireless3:
		*s = OCPPInterfaceEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OCPPInterfaceEnumType", raw),
	)
}

// OCPPTransportEnumType (3.61)
//
// Enumeration of OCPP transport mechanisms. SOAP is currently not a valid value for OCPP 2.0.
type OCPPTransportEnumType string

const (
	// Use JSON over WebSockets for transport of OCPP PDU’s
	OCPPTransportEnumTypeJSON OCPPTransportEnumType = "JSON"
	// Use SOAP for transport of OCPP PDU’s
	OCPPTransportEnumTypeSOAP OCPPTransportEnumType = "SOAP"
)

func (s *OCPPTransportEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OCPPTransportEnumType(raw) {
	case OCPPTransportEnumTypeJSON,
		OCPPTransportEnumTypeSOAP:
		*s = OCPPTransportEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OCPPTransportEnumType", raw),
	)
}

// OCPPVersionEnumType (3.62)
//
// Enumeration of OCPP versions.
type OCPPVersionEnumType string

const (
	// OCPP version 1.2
	OCPPVersionEnumTypeOCPP12 OCPPVersionEnumType = "OCPP12"
	// OCPP version 1.5
	OCPPVersionEnumTypeOCPP15 OCPPVersionEnumType = "OCPP15"
	// OCPP version 1.6
	OCPPVersionEnumTypeOCPP16 OCPPVersionEnumType = "OCPP16"
)

func (s *OCPPVersionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OCPPVersionEnumType(raw) {
	case OCPPVersionEnumTypeOCPP12,
		OCPPVersionEnumTypeOCPP15,
		OCPPVersionEnumTypeOCPP16:
		*s = OCPPVersionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OCPPVersionEnumType", raw),
	)
}

// OperationalStatusEnumType (3.63)
//
// Requested availability change.
type OperationalStatusEnumType string

const (
	// Charging Station is not available for charging.
	OperationalStatusEnumTypeInoperative OperationalStatusEnumType = "Inoperative"
	// Charging Station is available for charging.
	OperationalStatusEnumTypeOperative OperationalStatusEnumType = "Operative"
)

func (s *OperationalStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OperationalStatusEnumType(raw) {
	case OperationalStatusEnumTypeInoperative,
		OperationalStatusEnumTypeOperative:
		*s = OperationalStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OperationalStatusEnumType", raw),
	)
}

// PhaseEnumType (3.64)
//
// Phase specifies how a measured value is to be interpreted. Please note that not all values of Phase
// are applicable to all Measurands.
type PhaseEnumType string

const (
	// Measured on L1
	PhaseEnumTypeL1 PhaseEnumType = "L1"
	// Measured on L2
	PhaseEnumTypeL2 PhaseEnumType = "L2"
	// Measured on L3
	PhaseEnumTypeL3 PhaseEnumType = "L3"
	// Measured on Neutral
	PhaseEnumTypeN PhaseEnumType = "N"
	// Measured on L1 with respect to Neutral conductor
	PhaseEnumTypeL1N PhaseEnumType = "L1-N"
	// Measured on L2 with respect to Neutral conductor
	PhaseEnumTypeL2N PhaseEnumType = "L2-N"
	// Measured on L3 with respect to Neutral conductor
	PhaseEnumTypeL3N PhaseEnumType = "L3-N"
	// Measured between L1 and L2
	PhaseEnumTypeL1L2 PhaseEnumType = "L1-L2"
	// Measured between L2 and L3
	PhaseEnumTypeL2L3 PhaseEnumType = "L2-L3"
	// Measured between L3 and L1
	PhaseEnumTypeL3L1 PhaseEnumType = "L3-L1"
)

func (s *PhaseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch PhaseEnumType(raw) {
	case PhaseEnumTypeL1,
		PhaseEnumTypeL2,
		PhaseEnumTypeL3,
		PhaseEnumTypeN,
		PhaseEnumTypeL1N,
		PhaseEnumTypeL2N,
		PhaseEnumTypeL3N,
		PhaseEnumTypeL1L2,
		PhaseEnumTypeL2L3,
		PhaseEnumTypeL3L1:
		*s = PhaseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PhaseEnumType", raw),
	)
}

// PublishFirmwareStatusEnumType (3.65)
//
// Status for when publishing a Firmware.
type PublishFirmwareStatusEnumType string

const (
	PublishFirmwareStatusEnumTypeIdle PublishFirmwareStatusEnumType = "Idle"
	// Intermediate state. Downloading of new firmware has been scheduled.
	PublishFirmwareStatusEnumTypeDownloadScheduled PublishFirmwareStatusEnumType = "DownloadScheduled"
	// Intermediate state. Firmware is being downloaded.
	PublishFirmwareStatusEnumTypeDownloading PublishFirmwareStatusEnumType = "Downloading"
	// Intermediate state. New firmware has been downloaded by Charging Station.
	PublishFirmwareStatusEnumTypeDownloaded PublishFirmwareStatusEnumType = "Downloaded"
	// The firmware has been successfully published.
	PublishFirmwareStatusEnumTypePublished PublishFirmwareStatusEnumType = "Published"
	// Failure end state. Charging Station failed to download firmware.
	PublishFirmwareStatusEnumTypeDownloadFailed PublishFirmwareStatusEnumType = "DownloadFailed"
	// Intermediate state. Downloading has been paused.
	PublishFirmwareStatusEnumTypeDownloadPaused PublishFirmwareStatusEnumType = "DownloadPaused"
)

func (s *PublishFirmwareStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch PublishFirmwareStatusEnumType(raw) {
	case PublishFirmwareStatusEnumTypeIdle,
		PublishFirmwareStatusEnumTypeDownloadScheduled,
		PublishFirmwareStatusEnumTypeDownloading,
		PublishFirmwareStatusEnumTypeDownloaded,
		PublishFirmwareStatusEnumTypePublished,
		PublishFirmwareStatusEnumTypeDownloadFailed,
		PublishFirmwareStatusEnumTypeDownloadPaused:
		*s = PublishFirmwareStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PublishFirmwareStatusEnumType", raw),
	)
}

// ReadingContextEnumType (3.66)
//
// Values of the context field.
type ReadingContextEnumType string

const (
	// Value taken at start of interruption.
	ReadingContextEnumTypeInterruptionBegin ReadingContextEnumType = "Interruption.Begin"
	// Value taken when resuming after interruption.
	ReadingContextEnumTypeInterruptionEnd ReadingContextEnumType = "Interruption.End"
	// Value for any other situations.
	ReadingContextEnumTypeOther ReadingContextEnumType = "Other"
	// Value taken at clock aligned interval.
	ReadingContextEnumTypeSampleClock ReadingContextEnumType = "Sample.Clock"
	// Value taken as periodic sample relative to start time of transaction.
	ReadingContextEnumTypeSamplePeriodic ReadingContextEnumType = "Sample.Periodic"
	// Value taken at start of transaction.
	ReadingContextEnumTypeTransactionBegin ReadingContextEnumType = "Transaction.Begin"
	// Value taken at end of transaction.
	ReadingContextEnumTypeTransactionEnd ReadingContextEnumType = "Transaction.End"
	// Value taken in response to TriggerMessageRequest.
	ReadingContextEnumTypeTrigger ReadingContextEnumType = "Trigger"
)

func (s *ReadingContextEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReadingContextEnumType(raw) {
	case ReadingContextEnumTypeInterruptionBegin,
		ReadingContextEnumTypeInterruptionEnd,
		ReadingContextEnumTypeOther,
		ReadingContextEnumTypeSampleClock,
		ReadingContextEnumTypeSamplePeriodic,
		ReadingContextEnumTypeTransactionBegin,
		ReadingContextEnumTypeTransactionEnd,
		ReadingContextEnumTypeTrigger:
		*s = ReadingContextEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReadingContextEnumType", raw),
	)
}

// ReasonEnumType (3.67)
//
// Reason for stopping a transaction.
type ReasonEnumType string

const (
	// The transaction was stopped because of the authorization status in the response to a
	// transactionEventRequest.
	ReasonEnumTypeDeAuthorized ReasonEnumType = "DeAuthorized"
	// Emergency stop button was used.
	ReasonEnumTypeEmergencyStop ReasonEnumType = "EmergencyStop"
	// EV charging session reached a locally enforced maximum energy transfer limit
	ReasonEnumTypeEnergyLimitReached ReasonEnumType = "EnergyLimitReached"
	// Disconnecting of cable, vehicle moved away from inductive charge unit.
	ReasonEnumTypeEVDisconnected ReasonEnumType = "EVDisconnected"
	// A GroundFault has occurred
	ReasonEnumTypeGroundFault ReasonEnumType = "GroundFault"
	// A Reset(Immediate) command was received.
	ReasonEnumTypeImmediateReset ReasonEnumType = "ImmediateReset"
	// Stopped locally on request of the EV Driver at the Charging Station. This is a regular termination
	// of a transaction. Examples: presenting an IdToken tag, pressing a button to stop.
	ReasonEnumTypeLocal ReasonEnumType = "Local"
	// A local credit limit enforced through the Charging Station has been exceeded.
	ReasonEnumTypeLocalOutOfCredit ReasonEnumType = "LocalOutOfCredit"
	// The transaction was stopped using a token with a MasterPassGroupId.
	ReasonEnumTypeMasterPass ReasonEnumType = "MasterPass"
	// Any other reason.
	ReasonEnumTypeOther ReasonEnumType = "Other"
	// A larger than intended electric current has occurred
	ReasonEnumTypeOvercurrentFault ReasonEnumType = "OvercurrentFault"
	// Complete loss of power.
	ReasonEnumTypePowerLoss ReasonEnumType = "PowerLoss"
	// Quality of power too low, e.g. voltage too low/high, phase imbalance, etc.
	ReasonEnumTypePowerQuality ReasonEnumType = "PowerQuality"
	// A locally initiated reset/reboot occurred. (for instance watchdog kicked in)
	ReasonEnumTypeReboot ReasonEnumType = "Reboot"
	// Stopped remotely on request of the CSMS. This is a regular termination of a transaction. Examples:
	// termination using a smartphone app, exceeding a (non local) prepaid credit.
	ReasonEnumTypeRemote ReasonEnumType = "Remote"
	// Electric vehicle has reported reaching a locally enforced maximum battery State of Charge (SOC)
	ReasonEnumTypeSOCLimitReached ReasonEnumType = "SOCLimitReached"
	// The transaction was stopped by the EV
	ReasonEnumTypeStoppedByEV ReasonEnumType = "StoppedByEV"
	// EV charging session reached a locally enforced time limit
	ReasonEnumTypeTimeLimitReached ReasonEnumType = "TimeLimitReached"
	// EV not connected within timeout
	ReasonEnumTypeTimeout ReasonEnumType = "Timeout"
)

func (s *ReasonEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReasonEnumType(raw) {
	case ReasonEnumTypeDeAuthorized,
		ReasonEnumTypeEmergencyStop,
		ReasonEnumTypeEnergyLimitReached,
		ReasonEnumTypeEVDisconnected,
		ReasonEnumTypeGroundFault,
		ReasonEnumTypeImmediateReset,
		ReasonEnumTypeLocal,
		ReasonEnumTypeLocalOutOfCredit,
		ReasonEnumTypeMasterPass,
		ReasonEnumTypeOther,
		ReasonEnumTypeOvercurrentFault,
		ReasonEnumTypePowerLoss,
		ReasonEnumTypePowerQuality,
		ReasonEnumTypeReboot,
		ReasonEnumTypeRemote,
		ReasonEnumTypeSOCLimitReached,
		ReasonEnumTypeStoppedByEV,
		ReasonEnumTypeTimeLimitReached,
		ReasonEnumTypeTimeout:
		*s = ReasonEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReasonEnumType", raw),
	)
}

// RecurrencyKindEnumType (3.68)
type RecurrencyKindEnumType string

const (
	// The schedule restarts every 24 hours, at the same time as in the startSchedule.
	RecurrencyKindEnumTypeDaily RecurrencyKindEnumType = "Daily"
	// The schedule restarts every 7 days, at the same time and day-of-the-week as in the startSchedule.
	RecurrencyKindEnumTypeWeekly RecurrencyKindEnumType = "Weekly"
)

func (s *RecurrencyKindEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch RecurrencyKindEnumType(raw) {
	case RecurrencyKindEnumTypeDaily,
		RecurrencyKindEnumTypeWeekly:
		*s = RecurrencyKindEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid RecurrencyKindEnumType", raw),
	)
}

// RegistrationStatusEnumType (3.69)
//
// Result of registration in response to BootNotificationRequest.
type RegistrationStatusEnumType string

const (
	// Charging Station is accepted by the CSMS.
	RegistrationStatusEnumTypeAccepted RegistrationStatusEnumType = "Accepted"
	// CSMS is not yet ready to accept the Charging Station. CSMS may send messages to retrieve information
	// or prepare the Charging Station.
	RegistrationStatusEnumTypePending RegistrationStatusEnumType = "Pending"
	// Charging Station is not accepted by CSMS. This may happen when the Charging Station id is not known
	// by CSMS.
	RegistrationStatusEnumTypeRejected RegistrationStatusEnumType = "Rejected"
)

func (s *RegistrationStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch RegistrationStatusEnumType(raw) {
	case RegistrationStatusEnumTypeAccepted,
		RegistrationStatusEnumTypePending,
		RegistrationStatusEnumTypeRejected:
		*s = RegistrationStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid RegistrationStatusEnumType", raw),
	)
}

// ReportBaseEnumType (3.70)
type ReportBaseEnumType string

const (
	// Required. A (configuration) report that lists all Components/Variables that can be set by the
	// operator.
	ReportBaseEnumTypeConfigurationInventory ReportBaseEnumType = "ConfigurationInventory"
	// Required. A (full) report that lists everything except monitoring settings.
	ReportBaseEnumTypeFullInventory ReportBaseEnumType = "FullInventory"
	// Optional. A (summary) report that lists Components/Variables relating to the Charging Station’s
	// current charging availability, and to any existing problem conditions.
	ReportBaseEnumTypeSummaryInventory ReportBaseEnumType = "SummaryInventory"
)

func (s *ReportBaseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReportBaseEnumType(raw) {
	case ReportBaseEnumTypeConfigurationInventory,
		ReportBaseEnumTypeFullInventory,
		ReportBaseEnumTypeSummaryInventory:
		*s = ReportBaseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReportBaseEnumType", raw),
	)
}

// RequestStartStopStatusEnumType (3.71)
//
// The result of a RequestStartTransactionRequest or RequestStopTransactionRequest.
type RequestStartStopStatusEnumType string

const (
	// Command will be executed.
	RequestStartStopStatusEnumTypeAccepted RequestStartStopStatusEnumType = "Accepted"
	// Command will not be executed.
	RequestStartStopStatusEnumTypeRejected RequestStartStopStatusEnumType = "Rejected"
)

func (s *RequestStartStopStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch RequestStartStopStatusEnumType(raw) {
	case RequestStartStopStatusEnumTypeAccepted,
		RequestStartStopStatusEnumTypeRejected:
		*s = RequestStartStopStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid RequestStartStopStatusEnumType", raw),
	)
}

// ReservationUpdateStatusEnumType (3.72)
type ReservationUpdateStatusEnumType string

const (
	// The reservation is expired.
	ReservationUpdateStatusEnumTypeExpired ReservationUpdateStatusEnumType = "Expired"
	// The reservation is removed.
	ReservationUpdateStatusEnumTypeRemoved ReservationUpdateStatusEnumType = "Removed"
)

func (s *ReservationUpdateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReservationUpdateStatusEnumType(raw) {
	case ReservationUpdateStatusEnumTypeExpired,
		ReservationUpdateStatusEnumTypeRemoved:
		*s = ReservationUpdateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReservationUpdateStatusEnumType", raw),
	)
}

// ReserveNowStatusEnumType (3.73)
//
// Status in ReserveNowResponse.
type ReserveNowStatusEnumType string

const (
	// Reservation has been made.
	ReserveNowStatusEnumTypeAccepted ReserveNowStatusEnumType = "Accepted"
	// Reservation has not been made, because evse, connectors or specified connector are in a faulted
	// state.
	ReserveNowStatusEnumTypeFaulted ReserveNowStatusEnumType = "Faulted"
	// Reservation has not been made. The evse or the specified connector is occupied.
	ReserveNowStatusEnumTypeOccupied ReserveNowStatusEnumType = "Occupied"
	// Reservation has not been made. Charging Station is not configured to accept reservations.
	ReserveNowStatusEnumTypeRejected ReserveNowStatusEnumType = "Rejected"
	// Reservation has not been made, because evse, connectors or specified connector are in an unavailable
	// state.
	ReserveNowStatusEnumTypeUnavailable ReserveNowStatusEnumType = "Unavailable"
)

func (s *ReserveNowStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReserveNowStatusEnumType(raw) {
	case ReserveNowStatusEnumTypeAccepted,
		ReserveNowStatusEnumTypeFaulted,
		ReserveNowStatusEnumTypeOccupied,
		ReserveNowStatusEnumTypeRejected,
		ReserveNowStatusEnumTypeUnavailable:
		*s = ReserveNowStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReserveNowStatusEnumType", raw),
	)
}

// ResetEnumType (3.74)
//
// Type of reset requested.
type ResetEnumType string

const (
	// Immediate reset of the Charging Station.
	ResetEnumTypeImmediate ResetEnumType = "Immediate"
	// Delay reset until no more transactions are active.
	ResetEnumTypeOnIdle ResetEnumType = "OnIdle"
)

func (s *ResetEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ResetEnumType(raw) {
	case ResetEnumTypeImmediate,
		ResetEnumTypeOnIdle:
		*s = ResetEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ResetEnumType", raw),
	)
}

// ResetStatusEnumType (3.75)
//
// Result of ResetRequest.
type ResetStatusEnumType string

const (
	// Command will be executed.
	ResetStatusEnumTypeAccepted ResetStatusEnumType = "Accepted"
	// Command will not be executed.
	ResetStatusEnumTypeRejected ResetStatusEnumType = "Rejected"
	// Reset command is scheduled, Charging Station is busy with a process that cannot be interrupted at
	// the moment. Reset will be executed when process is finished.
	ResetStatusEnumTypeScheduled ResetStatusEnumType = "Scheduled"
)

func (s *ResetStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ResetStatusEnumType(raw) {
	case ResetStatusEnumTypeAccepted,
		ResetStatusEnumTypeRejected,
		ResetStatusEnumTypeScheduled:
		*s = ResetStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ResetStatusEnumType", raw),
	)
}

// SendLocalListStatusEnumType (3.76)
//
// Type of update for SendLocalListRequest.
type SendLocalListStatusEnumType string

const (
	// Local Authorization List successfully updated.
	SendLocalListStatusEnumTypeAccepted SendLocalListStatusEnumType = "Accepted"
	// Failed to update the Local Authorization List.
	SendLocalListStatusEnumTypeFailed SendLocalListStatusEnumType = "Failed"
	// Version number in the request for a differential update is less or equal then version number of
	// current list.
	SendLocalListStatusEnumTypeVersionMismatch SendLocalListStatusEnumType = "VersionMismatch"
)

func (s *SendLocalListStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch SendLocalListStatusEnumType(raw) {
	case SendLocalListStatusEnumTypeAccepted,
		SendLocalListStatusEnumTypeFailed,
		SendLocalListStatusEnumTypeVersionMismatch:
		*s = SendLocalListStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid SendLocalListStatusEnumType", raw),
	)
}

// SetMonitoringStatusEnumType (3.77)
type SetMonitoringStatusEnumType string

const (
	// Monitor successfully set.
	SetMonitoringStatusEnumTypeAccepted SetMonitoringStatusEnumType = "Accepted"
	// Component is not known.
	SetMonitoringStatusEnumTypeUnknownComponent SetMonitoringStatusEnumType = "UnknownComponent"
	// Variable is not known.
	SetMonitoringStatusEnumTypeUnknownVariable SetMonitoringStatusEnumType = "UnknownVariable"
	// Requested monitor type is not supported.
	SetMonitoringStatusEnumTypeUnsupportedMonitorType SetMonitoringStatusEnumType = "UnsupportedMonitorType"
	// Request is rejected.
	SetMonitoringStatusEnumTypeRejected SetMonitoringStatusEnumType = "Rejected"
	// A monitor already exists for the given type/severity combination.
	SetMonitoringStatusEnumTypeDuplicate SetMonitoringStatusEnumType = "Duplicate"
)

func (s *SetMonitoringStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch SetMonitoringStatusEnumType(raw) {
	case SetMonitoringStatusEnumTypeAccepted,
		SetMonitoringStatusEnumTypeUnknownComponent,
		SetMonitoringStatusEnumTypeUnknownVariable,
		SetMonitoringStatusEnumTypeUnsupportedMonitorType,
		SetMonitoringStatusEnumTypeRejected,
		SetMonitoringStatusEnumTypeDuplicate:
		*s = SetMonitoringStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid SetMonitoringStatusEnumType", raw),
	)
}

// SetNetworkProfileStatusEnumType (3.78)
//
// Possible values of SetNetworkProfileStatus as used in SetNetworkProfileResponse.
type SetNetworkProfileStatusEnumType string

const (
	// Setting new data successful
	SetNetworkProfileStatusEnumTypeAccepted SetNetworkProfileStatusEnumType = "Accepted"
	// Setting new data rejected
	SetNetworkProfileStatusEnumTypeRejected SetNetworkProfileStatusEnumType = "Rejected"
	// Setting new data failed
	SetNetworkProfileStatusEnumTypeFailed SetNetworkProfileStatusEnumType = "Failed"
)

func (s *SetNetworkProfileStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch SetNetworkProfileStatusEnumType(raw) {
	case SetNetworkProfileStatusEnumTypeAccepted,
		SetNetworkProfileStatusEnumTypeRejected,
		SetNetworkProfileStatusEnumTypeFailed:
		*s = SetNetworkProfileStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid SetNetworkProfileStatusEnumType", raw),
	)
}

// SetVariableStatusEnumType (3.79)
type SetVariableStatusEnumType string

const (
	// Variable successfully set.
	SetVariableStatusEnumTypeAccepted SetVariableStatusEnumType = "Accepted"
	// Request is rejected.
	SetVariableStatusEnumTypeRejected SetVariableStatusEnumType = "Rejected"
	// Component is not known.
	SetVariableStatusEnumTypeUnknownComponent SetVariableStatusEnumType = "UnknownComponent"
	// Variable is not known.
	SetVariableStatusEnumTypeUnknownVariable SetVariableStatusEnumType = "UnknownVariable"
	// The AttributeType is not supported.
	SetVariableStatusEnumTypeNotSupportedAttributeType SetVariableStatusEnumType = "NotSupportedAttributeType"
	// A reboot is required.
	SetVariableStatusEnumTypeRebootRequired SetVariableStatusEnumType = "RebootRequired"
)

func (s *SetVariableStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch SetVariableStatusEnumType(raw) {
	case SetVariableStatusEnumTypeAccepted,
		SetVariableStatusEnumTypeRejected,
		SetVariableStatusEnumTypeUnknownComponent,
		SetVariableStatusEnumTypeUnknownVariable,
		SetVariableStatusEnumTypeNotSupportedAttributeType,
		SetVariableStatusEnumTypeRebootRequired:
		*s = SetVariableStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid SetVariableStatusEnumType", raw),
	)
}

// TransactionEventEnumType (3.80)
type TransactionEventEnumType string

const (
	// Last event of a transaction
	TransactionEventEnumTypeEnded TransactionEventEnumType = "Ended"
	// First event of a transaction.
	TransactionEventEnumTypeStarted TransactionEventEnumType = "Started"
	// Transaction event in between 'Started' and 'Ended'.
	TransactionEventEnumTypeUpdated TransactionEventEnumType = "Updated"
)

func (s *TransactionEventEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TransactionEventEnumType(raw) {
	case TransactionEventEnumTypeEnded,
		TransactionEventEnumTypeStarted,
		TransactionEventEnumTypeUpdated:
		*s = TransactionEventEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TransactionEventEnumType", raw),
	)
}

// TriggerMessageStatusEnumType (3.81)
//
// Status in TriggerMessageResponse.
type TriggerMessageStatusEnumType string

const (
	// Requested message will be sent.
	TriggerMessageStatusEnumTypeAccepted TriggerMessageStatusEnumType = "Accepted"
	// Requested message will not be sent.
	TriggerMessageStatusEnumTypeRejected TriggerMessageStatusEnumType = "Rejected"
	// Requested message cannot be sent because it is either not implemented or unknown.
	TriggerMessageStatusEnumTypeNotImplemented TriggerMessageStatusEnumType = "NotImplemented"
)

func (s *TriggerMessageStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TriggerMessageStatusEnumType(raw) {
	case TriggerMessageStatusEnumTypeAccepted,
		TriggerMessageStatusEnumTypeRejected,
		TriggerMessageStatusEnumTypeNotImplemented:
		*s = TriggerMessageStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TriggerMessageStatusEnumType", raw),
	)
}

// TriggerReasonEnumType (3.82)
//
// Reason that triggered a transactionEventRequest.
type TriggerReasonEnumType string

const (
	// Charging is authorized, by any means. Might be an RFID, or other authorization means.
	TriggerReasonEnumTypeAuthorized TriggerReasonEnumType = "Authorized"
	// Cable is plugged in and EVDetected.
	TriggerReasonEnumTypeCablePluggedIn TriggerReasonEnumType = "CablePluggedIn"
	// Rate of charging changed by more than LimitChangeSignificance.
	TriggerReasonEnumTypeChargingRateChanged TriggerReasonEnumType = "ChargingRateChanged"
	// Charging State changed.
	TriggerReasonEnumTypeChargingStateChanged TriggerReasonEnumType = "ChargingStateChanged"
	// The transaction was stopped because of the authorization status in the response to a
	// transactionEventRequest.
	TriggerReasonEnumTypeDeauthorized TriggerReasonEnumType = "Deauthorized"
	// Maximum energy of charging reached. For example: in a pre-paid charging solution
	TriggerReasonEnumTypeEnergyLimitReached TriggerReasonEnumType = "EnergyLimitReached"
	// Communication with EV lost, for example: cable disconnected.
	TriggerReasonEnumTypeEVCommunicationLost TriggerReasonEnumType = "EVCommunicationLost"
	// EV not connected before the connection is timed out.
	TriggerReasonEnumTypeEVConnectTimeout TriggerReasonEnumType = "EVConnectTimeout"
	// Needed to send a clock aligned meter value
	TriggerReasonEnumTypeMeterValueClock TriggerReasonEnumType = "MeterValueClock"
)

func (s *TriggerReasonEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TriggerReasonEnumType(raw) {
	case TriggerReasonEnumTypeAuthorized,
		TriggerReasonEnumTypeCablePluggedIn,
		TriggerReasonEnumTypeChargingRateChanged,
		TriggerReasonEnumTypeChargingStateChanged,
		TriggerReasonEnumTypeDeauthorized,
		TriggerReasonEnumTypeEnergyLimitReached,
		TriggerReasonEnumTypeEVCommunicationLost,
		TriggerReasonEnumTypeEVConnectTimeout,
		TriggerReasonEnumTypeMeterValueClock:
		*s = TriggerReasonEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TriggerReasonEnumType", raw),
	)
}

// UnlockStatusEnumType (3.83)
//
// Status in response to UnlockConnectorRequest.
type UnlockStatusEnumType string

const (
	// Connector has successfully been unlocked.
	UnlockStatusEnumTypeUnlocked UnlockStatusEnumType = "Unlocked"
	// Failed to unlock the connector.
	UnlockStatusEnumTypeUnlockFailed UnlockStatusEnumType = "UnlockFailed"
	// The connector is not unlocked, because there is still an authorized transaction ongoing.
	UnlockStatusEnumTypeOngoingAuthorizedTransaction UnlockStatusEnumType = "OngoingAuthorizedTransaction"
	// The specified connector is not known by the Charging Station.
	UnlockStatusEnumTypeUnknownConnector UnlockStatusEnumType = "UnknownConnector"
)

func (s *UnlockStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UnlockStatusEnumType(raw) {
	case UnlockStatusEnumTypeUnlocked,
		UnlockStatusEnumTypeUnlockFailed,
		UnlockStatusEnumTypeOngoingAuthorizedTransaction,
		UnlockStatusEnumTypeUnknownConnector:
		*s = UnlockStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UnlockStatusEnumType", raw),
	)
}

// UnpublishFirmwareStatusEnumType (3.84)
//
// Status for when publishing a Firmware.
type UnpublishFirmwareStatusEnumType string

const (
	// Intermediate state. Firmware is being downloaded.
	UnpublishFirmwareStatusEnumTypeDownloadOngoing UnpublishFirmwareStatusEnumType = "DownloadOngoing"
	// There is no published file.
	UnpublishFirmwareStatusEnumTypeNoFirmware UnpublishFirmwareStatusEnumType = "NoFirmware"
	// Successful end state. Firmware file no longer being published.
	UnpublishFirmwareStatusEnumTypeUnpublished UnpublishFirmwareStatusEnumType = "Unpublished"
)

func (s *UnpublishFirmwareStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UnpublishFirmwareStatusEnumType(raw) {
	case UnpublishFirmwareStatusEnumTypeDownloadOngoing,
		UnpublishFirmwareStatusEnumTypeNoFirmware,
		UnpublishFirmwareStatusEnumTypeUnpublished:
		*s = UnpublishFirmwareStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UnpublishFirmwareStatusEnumType", raw),
	)
}

// UpdateEnumType (3.85)
type UpdateEnumType string

const (
	// Indicates that the current Local Authorization List must be updated with the values in this message.
	UpdateEnumTypeDifferential UpdateEnumType = "Differential"
	// Indicates that the current Local Authorization List must be replaced by the values in this message.
	UpdateEnumTypeFull UpdateEnumType = "Full"
)

func (s *UpdateEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UpdateEnumType(raw) {
	case UpdateEnumTypeDifferential,
		UpdateEnumTypeFull:
		*s = UpdateEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UpdateEnumType", raw),
	)
}

// UpdateFirmwareStatusEnumType (3.86)
//
// Generic message response status
type UpdateFirmwareStatusEnumType string

const (
	// Accepted this firmware update request. This does not mean the firmware update is successful, the
	// Charging Station will now start the firmware update process.
	UpdateFirmwareStatusEnumTypeAccepted UpdateFirmwareStatusEnumType = "Accepted"
	// Firmware update request rejected.
	UpdateFirmwareStatusEnumTypeRejected UpdateFirmwareStatusEnumType = "Rejected"
	// Accepted this firmware update request, but in doing this has canceled an ongoing firmware update.
	UpdateFirmwareStatusEnumTypeAcceptedCanceled UpdateFirmwareStatusEnumType = "AcceptedCanceled"
	// The certificate is invalid.
	UpdateFirmwareStatusEnumTypeInvalidCertificate UpdateFirmwareStatusEnumType = "InvalidCertificate"
	// Failure end state. The Firmware Signing certificate has been revoked.
	UpdateFirmwareStatusEnumTypeRevokedCertificate UpdateFirmwareStatusEnumType = "RevokedCertificate"
)

func (s *UpdateFirmwareStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UpdateFirmwareStatusEnumType(raw) {
	case UpdateFirmwareStatusEnumTypeAccepted,
		UpdateFirmwareStatusEnumTypeRejected,
		UpdateFirmwareStatusEnumTypeAcceptedCanceled,
		UpdateFirmwareStatusEnumTypeInvalidCertificate,
		UpdateFirmwareStatusEnumTypeRevokedCertificate:
		*s = UpdateFirmwareStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UpdateFirmwareStatusEnumType", raw),
	)
}

// UploadLogStatusEnumType (3.87)
type UploadLogStatusEnumType string

const (
	// A badly formatted packet or other protocol incompatibility was detected.
	UploadLogStatusEnumTypeBadMessage UploadLogStatusEnumType = "BadMessage"
	// The Charging Station is not uploading a log file. Idle SHALL only be used when the message was
	// triggered by a TriggerMessageRequest.
	UploadLogStatusEnumTypeIdle UploadLogStatusEnumType = "Idle"
	// The server does not support the operation
	UploadLogStatusEnumTypeNotSupportedOperation UploadLogStatusEnumType = "NotSupportedOperation"
	// Insufficient permissions to perform the operation.
	UploadLogStatusEnumTypePermissionDenied UploadLogStatusEnumType = "PermissionDenied"
	// File has been uploaded successfully.
	UploadLogStatusEnumTypeUploaded UploadLogStatusEnumType = "Uploaded"
	// Failed to upload the requested file.
	UploadLogStatusEnumTypeUploadFailure UploadLogStatusEnumType = "UploadFailure"
	// File is being uploaded.
	UploadLogStatusEnumTypeUploading UploadLogStatusEnumType = "Uploading"
	// On-going log upload is canceled and new request to upload log has been accepted.
	UploadLogStatusEnumTypeAcceptedCanceled UploadLogStatusEnumType = "AcceptedCanceled"
)

func (s *UploadLogStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch UploadLogStatusEnumType(raw) {
	case UploadLogStatusEnumTypeBadMessage,
		UploadLogStatusEnumTypeIdle,
		UploadLogStatusEnumTypeNotSupportedOperation,
		UploadLogStatusEnumTypePermissionDenied,
		UploadLogStatusEnumTypeUploaded,
		UploadLogStatusEnumTypeUploadFailure,
		UploadLogStatusEnumTypeUploading,
		UploadLogStatusEnumTypeAcceptedCanceled:
		*s = UploadLogStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid UploadLogStatusEnumType", raw),
	)
}

// VPNEnumType (3.88)
//
// Enumeration of VPN Types.
type VPNEnumType string

const (
	// IKEv2 VPN
	VPNEnumTypeIKEv2 VPNEnumType = "IKEv2"
	// IPSec VPN
	VPNEnumTypeIPSec VPNEnumType = "IPSec"
	// L2TP VPN
	VPNEnumTypeL2TP VPNEnumType = "L2TP"
	// PPTP VPN
	VPNEnumTypePPTP VPNEnumType = "PPTP"
)

func (s *VPNEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch VPNEnumType(raw) {
	case VPNEnumTypeIKEv2,
		VPNEnumTypeIPSec,
		VPNEnumTypeL2TP,
		VPNEnumTypePPTP:
		*s = VPNEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid VPNEnumType", raw),
	)
}
