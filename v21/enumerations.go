package v21

import (
	"encoding/json"
	"fmt"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// APNAuthenticationEnumType (3.1)
type APNAuthenticationEnumType string

const (
	// Use PAP authentication
	APNAuthenticationEnumTypePAP APNAuthenticationEnumType = "PAP"
	// Use CHAP authentication
	APNAuthenticationEnumTypeCHAP APNAuthenticationEnumType = "CHAP"
	// Use no authentication
	APNAuthenticationEnumTypeNONE APNAuthenticationEnumType = "NONE"
	// Sequentially try CHAP, PAP, NONE.
	APNAuthenticationEnumTypeAUTO APNAuthenticationEnumType = "AUTO"
)

func (s *APNAuthenticationEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch APNAuthenticationEnumType(raw) {
	case APNAuthenticationEnumTypePAP,
		APNAuthenticationEnumTypeCHAP,
		APNAuthenticationEnumTypeNONE,
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
	// The maximum allowed value for this variable
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
	// Identifier is already involved in another transaction and multiple transactions are not allowed. (Only
	// relevant for the response to a TransactionEventRequest(eventType=Started).)
	AuthorizationStatusEnumTypeConcurrentTx AuthorizationStatusEnumType = "ConcurrentTx"
	// Identifier has expired. Not allowed for charging.
	AuthorizationStatusEnumTypeExpired AuthorizationStatusEnumType = "Expired"
	// Identifier is invalid. Not allowed for charging.
	AuthorizationStatusEnumTypeInvalid AuthorizationStatusEnumType = "Invalid"
	// Identifier is valid, but EV Driver doesn’t have enough credit to start charging. Not allowed for charging.
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
	// If the Charging Station or CSMS determine (via a CRL or OCSP response) that the contract certificate in the
	// AuthorizeRequest is marked as revoked.
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

// BatterySwapEventEnumType (3.5)
//
// (2.1) Battery in/out event at a swap station.
type BatterySwapEventEnumType string

const (
	// Battery (or set of batteries) is inserted.
	BatterySwapEventEnumTypeBatteryIn BatterySwapEventEnumType = "BatteryIn"
	// Battery (or set of batteries) is removed.
	BatterySwapEventEnumTypeBatteryOut BatterySwapEventEnumType = "BatteryOut"
	// The offered batteries have not been removed within timeout.
	BatterySwapEventEnumTypeBatteryOutTimeout BatterySwapEventEnumType = "BatteryOutTimeout"
)

func (s *BatterySwapEventEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch BatterySwapEventEnumType(raw) {
	case BatterySwapEventEnumTypeBatteryIn,
		BatterySwapEventEnumTypeBatteryOut,
		BatterySwapEventEnumTypeBatteryOutTimeout:
		*s = BatterySwapEventEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid BatterySwapEventEnumType", raw),
	)
}

// BootReasonEnumType (3.6)
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

// CancelReservationStatusEnumType (3.7)
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

// CertificateActionEnumType (3.8)
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

// CertificateSignedStatusEnumType (3.9)
type CertificateSignedStatusEnumType string

const (
	// Signed certificate is valid.
	CertificateSignedStatusEnumTypeAccepted CertificateSignedStatusEnumType = "Accepted"
	// Signed certificate is invalid or requestId is unknown.
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

// CertificateSigningUseEnumType (3.10)
type CertificateSigningUseEnumType string

const (
	// Client side certificate used by the Charging Station to connect the the CSMS.
	CertificateSigningUseEnumTypeChargingStationCertificate CertificateSigningUseEnumType = "ChargingStationCertificate"
	// Use for certificate for ISO 15118-2 connections. This means that the certificate should be derived from the
	// V2G root.
	CertificateSigningUseEnumTypeV2GCertificate CertificateSigningUseEnumType = "V2GCertificate"
	// (2.1) Use for certificate for ISO 15118-20 connections. This means that the certificate should be derived from
	// the V2G root.
	CertificateSigningUseEnumTypeV2G20Certificate CertificateSigningUseEnumType = "V2G20Certificate"
)

func (s *CertificateSigningUseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CertificateSigningUseEnumType(raw) {
	case CertificateSigningUseEnumTypeChargingStationCertificate,
		CertificateSigningUseEnumTypeV2GCertificate,
		CertificateSigningUseEnumTypeV2G20Certificate:
		*s = CertificateSigningUseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CertificateSigningUseEnumType", raw),
	)
}

// CertificateStatusEnumType (3.11)
//
// OCSP or CRL status of certificate.
type CertificateStatusEnumType string

const (
	// Certificate has not been revoked.
	CertificateStatusEnumTypeGood CertificateStatusEnumType = "Good"
	// Certificate has been revoked.
	CertificateStatusEnumTypeRevoked CertificateStatusEnumType = "Revoked"
	// Certificate is unknown.
	CertificateStatusEnumTypeUnknown CertificateStatusEnumType = "Unknown"
	// The request to OCSP responder or CRL distribution point failed.
	CertificateStatusEnumTypeFailed CertificateStatusEnumType = "Failed"
)

func (s *CertificateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CertificateStatusEnumType(raw) {
	case CertificateStatusEnumTypeGood,
		CertificateStatusEnumTypeRevoked,
		CertificateStatusEnumTypeUnknown,
		CertificateStatusEnumTypeFailed:
		*s = CertificateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CertificateStatusEnumType", raw),
	)
}

// CertificateStatusSourceEnumType (3.12)
//
// Source of certificate status, OCSP or CRL.
type CertificateStatusSourceEnumType string

const (
	// Checked in a certificate revocation list.
	CertificateStatusSourceEnumTypeCRL CertificateStatusSourceEnumType = "CRL"
	// Checked via OCSP request.
	CertificateStatusSourceEnumTypeOCSP CertificateStatusSourceEnumType = "OCSP"
)

func (s *CertificateStatusSourceEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CertificateStatusSourceEnumType(raw) {
	case CertificateStatusSourceEnumTypeCRL,
		CertificateStatusSourceEnumTypeOCSP:
		*s = CertificateStatusSourceEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CertificateStatusSourceEnumType", raw),
	)
}

// ChangeAvailabilityStatusEnumType (3.13)
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

// ChargingProfileKindEnumType (3.14)
//
// Kind of charging profile.
type ChargingProfileKindEnumType string

const (
	// Schedule periods are relative to a fixed point in time defined in the schedule. This requires that
	// startSchedule is set to a starting point in time.
	ChargingProfileKindEnumTypeAbsolute ChargingProfileKindEnumType = "Absolute"
	// The schedule restarts periodically at the first schedule period. To be most useful, this requires that
	// startSchedule is set to a starting point in time.
	ChargingProfileKindEnumTypeRecurring ChargingProfileKindEnumType = "Recurring"
	// Charging schedule periods start when the EVSE is ready to deliver energy. i.e. when the EV driver is
	// authorized and the EV is connected. When a ChargingProfile is received for a transaction that is already
	// charging, then the charging schedule periods remain relative to the PowerPathClosed moment. No value for
	// startSchedule must be supplied.
	ChargingProfileKindEnumTypeRelative ChargingProfileKindEnumType = "Relative"
	// (2.1) The schedule consists of only one charging schedule period, which is updated dynamically by CSMS.
	ChargingProfileKindEnumTypeDynamic ChargingProfileKindEnumType = "Dynamic"
)

func (s *ChargingProfileKindEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ChargingProfileKindEnumType(raw) {
	case ChargingProfileKindEnumTypeAbsolute,
		ChargingProfileKindEnumTypeRecurring,
		ChargingProfileKindEnumTypeRelative,
		ChargingProfileKindEnumTypeDynamic:
		*s = ChargingProfileKindEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfileKindEnumType", raw),
	)
}

// ChargingProfilePurposeEnumType (3.15)
//
// Purpose of the charging profile.
type ChargingProfilePurposeEnumType string

const (
	// Additional constraints from an external source (e.g. an EMS) that will be incorporated into a local power
	// schedule. When applied to evse.Id = 0 it sets a limit to the entire Charging Station. Note: In OCPP 2.0.1 this
	// purpose was only allowed on evse.Id = 0. In OCPP 2.1 it can be set to an individual EVSE.
	ChargingProfilePurposeEnumTypeChargingStationExternalConstraints ChargingProfilePurposeEnumType = "ChargingStationExternalConstraints"
	// Configuration for the maximum power or current available for an entire Charging Station.
	ChargingProfilePurposeEnumTypeChargingStationMaxProfile ChargingProfilePurposeEnumType = "ChargingStationMaxProfile"
	// Default profile that can be configured in the Charging Station. When a new transaction is started, this
	// profile SHALL be used, unless it was a transaction that was started by a RequestStartTransactionRequest with a
	// ChargingProfile that is accepted by the Charging Station.
	ChargingProfilePurposeEnumTypeTxDefaultProfile ChargingProfilePurposeEnumType = "TxDefaultProfile"
	// Profile with constraints to be imposed by the Charging Station on the current transaction, or on a new
	// transaction when this is started via a RequestStartTransactionRequest with a ChargingProfile. A profile with
	// this purpose SHALL cease to be valid when the transaction terminates.
	ChargingProfilePurposeEnumTypeTxProfile ChargingProfilePurposeEnumType = "TxProfile"
	// (2.1) This profile is used in place of a Tx(Default)Profile, when priority charging is requested, either
	// locally on Charging Station or via a request from CSMS.
	ChargingProfilePurposeEnumTypePriorityCharging ChargingProfilePurposeEnumType = "PriorityCharging"
	// (2.1) This profile adds capacity from local generation. Its capacity is added on top of other charging
	// profiles.
	ChargingProfilePurposeEnumTypeLocalGeneration ChargingProfilePurposeEnumType = "LocalGeneration"
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
		ChargingProfilePurposeEnumTypeTxProfile,
		ChargingProfilePurposeEnumTypePriorityCharging,
		ChargingProfilePurposeEnumTypeLocalGeneration:
		*s = ChargingProfilePurposeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ChargingProfilePurposeEnumType", raw),
	)
}

// ChargingProfileStatusEnumType (3.16)
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

// ChargingRateUnitEnumType (3.17)
//
// Unit in which a charging schedule is defined.
type ChargingRateUnitEnumType string

const (
	// Watts (power). This is the TOTAL allowed charging power. If used for AC Charging, the phase current should be
	// calculated via: Current per phase = Power / (Line Voltage * Number of Phases). The "Line Voltage" used in the
	// calculation is not the measured voltage, but the set voltage for the area (hence, 230 of 110 volt). The
	// "Number of Phases" is the numberPhases from the ChargingSchedulePeriod. It is usually more convenient to use
	// this for DC charging. Note that if numberPhases in a ChargingSchedulePeriod is absent, 3 SHALL be assumed.
	ChargingRateUnitEnumTypeW ChargingRateUnitEnumType = "W"
	// Amperes (current). The amount of Ampere per phase, not the sum of all phases. It is usually more convenient to
	// use this for AC charging.
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

// ChargingStateEnumType (3.18)
//
// The state of the charging process.
type ChargingStateEnumType string

const (
	// There is a connection between EV and EVSE, in case the protocol used between EV and the Charging Station can
	// detect a connection, the protocol needs to detect this for the state to become active. The connection can
	// either be wired or wireless. Authorization is required to proceed to state Charging.
	ChargingStateEnumTypeEVConnected ChargingStateEnumType = "EVConnected"
	// The contactor of the Connector is closed and energy is flowing to between EVSE and EV.
	ChargingStateEnumTypeCharging ChargingStateEnumType = "Charging"
	// When the EV is connected to the EVSE and the EVSE is offering energy but the EV is not taking any energy.
	ChargingStateEnumTypeSuspendedEV ChargingStateEnumType = "SuspendedEV"
	// When the EV is connected to the EVSE but the EVSE is not offering energy to the EV, e.g. due to a smart
	// charging restrictions or local supply power constraints.
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
	case ChargingStateEnumTypeEVConnected,
		ChargingStateEnumTypeCharging,
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

// ClearCacheStatusEnumType (3.19)
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

// ClearChargingProfileStatusEnumType (3.20)
//
// Status returned in response to ClearChargingProfileRequest.
type ClearChargingProfileStatusEnumType string

const (
	// Request has been accepted.
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

// ClearMessageStatusEnumType (3.21)
//
// Result for a ClearDisplayMessageRequest as used in a ClearDisplayMessageResponse.
type ClearMessageStatusEnumType string

const (
	// Request successfully executed: message cleared.
	ClearMessageStatusEnumTypeAccepted ClearMessageStatusEnumType = "Accepted"
	// Given message (based on the id) not known.
	ClearMessageStatusEnumTypeUnknown ClearMessageStatusEnumType = "Unknown"
	// (2.1) Request could not be executed.
	ClearMessageStatusEnumTypeRejected ClearMessageStatusEnumType = "Rejected"
)

func (s *ClearMessageStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ClearMessageStatusEnumType(raw) {
	case ClearMessageStatusEnumTypeAccepted,
		ClearMessageStatusEnumTypeUnknown,
		ClearMessageStatusEnumTypeRejected:
		*s = ClearMessageStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ClearMessageStatusEnumType", raw),
	)
}

// ClearMonitoringStatusEnumType (3.22)
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

// ComponentCriterionEnumType (3.23)
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

// ConnectorStatusEnumType (3.24)
//
// A status can be reported for the Connector of an EVSE of a Charging Station. States considered Operative are:
// Available, Reserved and Occupied. States considered Inoperative are: Unavailable, Faulted.
type ConnectorStatusEnumType string

const (
	// When a Connector becomes available for a new User (Operative)
	ConnectorStatusEnumTypeAvailable ConnectorStatusEnumType = "Available"
	// When a Connector becomes occupied, so it is not available for a new EV driver. (Operative)
	ConnectorStatusEnumTypeOccupied ConnectorStatusEnumType = "Occupied"
	// When a Connector becomes reserved as a result of ReserveNow command (Operative)
	ConnectorStatusEnumTypeReserved ConnectorStatusEnumType = "Reserved"
	// When a Connector becomes unavailable as the result of a Change Availability command or an event upon which the
	// Charging Station transitions to unavailable at its discretion. Upon receipt of ChangeAvailability message
	// command, the status MAY change immediately or the change MAY be scheduled. When scheduled, StatusNotification
	// SHALL be send when the availability change becomes effective (Inoperative)
	ConnectorStatusEnumTypeUnavailable ConnectorStatusEnumType = "Unavailable"
	// When a Connector (or the EVSE or the entire Charging Station it belongs to) has reported an error and is not
	// available for energy delivery. (Inoperative).
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

// ControlModeEnumType (3.25)
//
// (2.1) ISO 15118-20 service parameter for control mode
type ControlModeEnumType string

const (
	// Scheduled control mode, EVSE provides up to three schedules for EV to choose from. EV follows the selected
	// schedule.
	ControlModeEnumTypeScheduledControl ControlModeEnumType = "ScheduledControl"
	// Dynamic control mode, EVSE executes a single schedule by sending setpoints to EV at every interval.
	ControlModeEnumTypeDynamicControl ControlModeEnumType = "DynamicControl"
)

func (s *ControlModeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ControlModeEnumType(raw) {
	case ControlModeEnumTypeScheduledControl,
		ControlModeEnumTypeDynamicControl:
		*s = ControlModeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ControlModeEnumType", raw),
	)
}

// CostDimensionEnumType (3.26)
//
// Usage dimension for cost in charging period.
type CostDimensionEnumType string

const (
	// Total amount of energy (dis-)charged during this charging period, defined in Wh (kiloWatt-hours). When
	// negative, more energy was feed into the grid then charged into the EV.
	CostDimensionEnumTypeEnergy CostDimensionEnumType = "Energy"
	// Sum of the maximum current over all phases, reached during this charging period, defined in A (Ampere).
	CostDimensionEnumTypeMaxCurrent CostDimensionEnumType = "MaxCurrent"
	// Sum of the minimum current over all phases, reached during this charging period, when negative, current has
	// flowed from the EV to the grid. Defined in A (Ampere).
	CostDimensionEnumTypeMinCurrent CostDimensionEnumType = "MinCurrent"
	// Maximum power reached during this charging period: defined in W (Watt).
	CostDimensionEnumTypeMaxPower CostDimensionEnumType = "MaxPower"
	// Minimum power reached during this charging period: defined in W (Watt), when negative, the power has flowed
	// from the EV to the grid.
	CostDimensionEnumTypeMinPower CostDimensionEnumType = "MinPower"
	// Time not charging during this charging period: defined in seconds.
	CostDimensionEnumTypeIdleTIme CostDimensionEnumType = "IdleTIme"
	// Time charging during this charging period: defined in seconds.
	CostDimensionEnumTypeChargingTime CostDimensionEnumType = "ChargingTime"
)

func (s *CostDimensionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch CostDimensionEnumType(raw) {
	case CostDimensionEnumTypeEnergy,
		CostDimensionEnumTypeMaxCurrent,
		CostDimensionEnumTypeMinCurrent,
		CostDimensionEnumTypeMaxPower,
		CostDimensionEnumTypeMinPower,
		CostDimensionEnumTypeIdleTIme,
		CostDimensionEnumTypeChargingTime:
		*s = CostDimensionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid CostDimensionEnumType", raw),
	)
}

// CostKindEnumType (3.27)
type CostKindEnumType string

const (
	// Absolute value. Carbon Dioxide emissions, in grams per kWh
	CostKindEnumTypeCarbonDioxideEmission CostKindEnumType = "CarbonDioxideEmission"
	// Relative value. Percentage of renewable generation within total generation.
	CostKindEnumTypeRelativePricePercentage CostKindEnumType = "RelativePricePercentage"
	// Relative value. Price per kWh, as percentage relative to the maximum price stated in any of all tariffs
	// indicated to the EV.
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

// CustomerInformationStatusEnumType (3.28)
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

// DataEnumType (3.29)
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

// DataTransferStatusEnumType (3.30)
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

// DayOfWeekEnumType (3.31)
type DayOfWeekEnumType string

const (
	DayOfWeekEnumTypeMonday    DayOfWeekEnumType = "Monday"
	DayOfWeekEnumTypeTuesday   DayOfWeekEnumType = "Tuesday"
	DayOfWeekEnumTypeWednesday DayOfWeekEnumType = "Wednesday"
	DayOfWeekEnumTypeThursday  DayOfWeekEnumType = "Thursday"
	DayOfWeekEnumTypeFriday    DayOfWeekEnumType = "Friday"
	DayOfWeekEnumTypeSaturday  DayOfWeekEnumType = "Saturday"
	DayOfWeekEnumTypeSunday    DayOfWeekEnumType = "Sunday"
)

func (s *DayOfWeekEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DayOfWeekEnumType(raw) {
	case DayOfWeekEnumTypeMonday,
		DayOfWeekEnumTypeTuesday,
		DayOfWeekEnumTypeWednesday,
		DayOfWeekEnumTypeThursday,
		DayOfWeekEnumTypeFriday,
		DayOfWeekEnumTypeSaturday,
		DayOfWeekEnumTypeSunday:
		*s = DayOfWeekEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DayOfWeekEnumType", raw),
	)
}

// DeleteCertificateStatusEnumType (3.32)
type DeleteCertificateStatusEnumType string

const (
	// Normal successful completion (no errors).
	DeleteCertificateStatusEnumTypeAccepted DeleteCertificateStatusEnumType = "Accepted"
	// The Charging Station either failed to remove the certificate or rejected the request. A Charging Station may
	// reject the request to prevent the deletion of a certificate, if it is the last one of its certificate type.
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

// DERControlEnumType (3.33)
//
// Enumeration of DER controls
type DERControlEnumType string

const (
	// Enter Service parameters setting
	DERControlEnumTypeEnterService DERControlEnumType = "EnterService"
	// Frequency droop settings
	DERControlEnumTypeFreqDroop DERControlEnumType = "FreqDroop"
	// Frequency-Watt curve
	DERControlEnumTypeFreqWatt DERControlEnumType = "FreqWatt"
	// Fixed power factor when absorbing power setting
	DERControlEnumTypeFixedPFAbsorb DERControlEnumType = "FixedPFAbsorb"
	// Fixed power factor when injecting power setting
	DERControlEnumTypeFixedPFInject DERControlEnumType = "FixedPFInject"
	// Fixed reactive power setpoint
	DERControlEnumTypeFixedVar DERControlEnumType = "FixedVar"
	// Gradient settings
	DERControlEnumTypeGradients DERControlEnumType = "Gradients"
	// High Frequency Must Trip curve
	DERControlEnumTypeHFMustTrip DERControlEnumType = "HFMustTrip"
	// High Frequency May Trip curve (ride-through)
	DERControlEnumTypeHFMayTrip DERControlEnumType = "HFMayTrip"
	// High Voltage Must Trip curve
	DERControlEnumTypeHVMustTrip DERControlEnumType = "HVMustTrip"
	// High Voltage Momentary Cessation curve
	DERControlEnumTypeHVMomCess DERControlEnumType = "HVMomCess"
	// High Voltage May Trip curve (ride-through)
	DERControlEnumTypeHVMayTrip DERControlEnumType = "HVMayTrip"
	// Limit discharge power to percentage of rated discharge power
	DERControlEnumTypeLimitMaxDischarge DERControlEnumType = "LimitMaxDischarge"
	// Low Frequency Must Trip curve
	DERControlEnumTypeLFMustTrip DERControlEnumType = "LFMustTrip"
	// Low Voltage Must Trip curve
	DERControlEnumTypeLVMustTrip DERControlEnumType = "LVMustTrip"
	// Low Voltage Momentary Cessation curve
	DERControlEnumTypeLVMomCess DERControlEnumType = "LVMomCess"
	// Low Voltage May Trip curve (ride-through)
	DERControlEnumTypeLVMayTrip DERControlEnumType = "LVMayTrip"
	// Power Monitoring curve according to VDE-AR-N 4105 section 5.5.2
	DERControlEnumTypePowerMonitoringMustTrip DERControlEnumType = "PowerMonitoringMustTrip"
	// Volt-Var curve
	DERControlEnumTypeVoltVar DERControlEnumType = "VoltVar"
	// Volt-Watt curve
	DERControlEnumTypeVoltWatt DERControlEnumType = "VoltWatt"
	// Watt-PowerFactor curve
	DERControlEnumTypeWattPF DERControlEnumType = "WattPF"
	// Watt-Var curve
	DERControlEnumTypeWattVar DERControlEnumType = "WattVar"
)

func (s *DERControlEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DERControlEnumType(raw) {
	case DERControlEnumTypeEnterService,
		DERControlEnumTypeFreqDroop,
		DERControlEnumTypeFreqWatt,
		DERControlEnumTypeFixedPFAbsorb,
		DERControlEnumTypeFixedPFInject,
		DERControlEnumTypeFixedVar,
		DERControlEnumTypeGradients,
		DERControlEnumTypeHFMustTrip,
		DERControlEnumTypeHFMayTrip,
		DERControlEnumTypeHVMustTrip,
		DERControlEnumTypeHVMomCess,
		DERControlEnumTypeHVMayTrip,
		DERControlEnumTypeLimitMaxDischarge,
		DERControlEnumTypeLFMustTrip,
		DERControlEnumTypeLVMustTrip,
		DERControlEnumTypeLVMomCess,
		DERControlEnumTypeLVMayTrip,
		DERControlEnumTypePowerMonitoringMustTrip,
		DERControlEnumTypeVoltVar,
		DERControlEnumTypeVoltWatt,
		DERControlEnumTypeWattPF,
		DERControlEnumTypeWattVar:
		*s = DERControlEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DERControlEnumType", raw),
	)
}

// DERControlStatusEnumType (3.34)
type DERControlStatusEnumType string

const (
	// Operation successful
	DERControlStatusEnumTypeAccepted DERControlStatusEnumType = "Accepted"
	// Operation failed
	DERControlStatusEnumTypeRejected DERControlStatusEnumType = "Rejected"
	// Type of DER setting or curve is not supported
	DERControlStatusEnumTypeNotSupported DERControlStatusEnumType = "NotSupported"
	// Type or Id in clear/get request was not found
	DERControlStatusEnumTypeNotFound DERControlStatusEnumType = "NotFound"
)

func (s *DERControlStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DERControlStatusEnumType(raw) {
	case DERControlStatusEnumTypeAccepted,
		DERControlStatusEnumTypeRejected,
		DERControlStatusEnumTypeNotSupported,
		DERControlStatusEnumTypeNotFound:
		*s = DERControlStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DERControlStatusEnumType", raw),
	)
}

// DERUnitEnumType (3.35)
type DERUnitEnumType string

const (
	// No unit applicable (e.g. for ride-through curves)
	DERUnitEnumTypeNotApplicable DERUnitEnumType = "Not_Applicable"
	// Percentage of configured active power
	DERUnitEnumTypePctMaxW DERUnitEnumType = "PctMaxW"
	// Percentage of configured reactive power
	DERUnitEnumTypePctMaxVar DERUnitEnumType = "PctMaxVar"
	// Percentage of available reserve active power
	DERUnitEnumTypePctWAvail DERUnitEnumType = "PctWAvail"
	// Percentage of available reserve reactive power
	DERUnitEnumTypePctVarAvail DERUnitEnumType = "PctVarAvail"
	// Percentage of effective voltage
	DERUnitEnumTypePctEffectiveV DERUnitEnumType = "PctEffectiveV"
)

func (s *DERUnitEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch DERUnitEnumType(raw) {
	case DERUnitEnumTypeNotApplicable,
		DERUnitEnumTypePctMaxW,
		DERUnitEnumTypePctMaxVar,
		DERUnitEnumTypePctWAvail,
		DERUnitEnumTypePctVarAvail,
		DERUnitEnumTypePctEffectiveV:
		*s = DERUnitEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DERUnitEnumType", raw),
	)
}

// DisplayMessageStatusEnumType (3.36)
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
	// (2.1) Message contains one or more languages that are not supported by Charging Station.
	DisplayMessageStatusEnumTypeLanguageNotSupported DisplayMessageStatusEnumType = "LanguageNotSupported"
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
		DisplayMessageStatusEnumTypeUnknownTransaction,
		DisplayMessageStatusEnumTypeLanguageNotSupported:
		*s = DisplayMessageStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid DisplayMessageStatusEnumType", raw),
	)
}

// EnergyTransferModeEnumType (3.37)
//
// Enumeration of energy transfer modes.
type EnergyTransferModeEnumType string

const (
	// AC single phase charging according to IEC 62196.
	EnergyTransferModeEnumTypeACSinglePhase EnergyTransferModeEnumType = "AC_single_phase"
	// AC two phase charging according to IEC 62196.
	EnergyTransferModeEnumTypeACTwoPhase EnergyTransferModeEnumType = "AC_two_phase"
	// AC three phase charging according to IEC 62196.
	EnergyTransferModeEnumTypeACThreePhase EnergyTransferModeEnumType = "AC_three_phase"
	// DC charging.
	EnergyTransferModeEnumTypeDC EnergyTransferModeEnumType = "DC"
	// (2.1) AC bidirectional (no DER control), ISO 15118-20
	EnergyTransferModeEnumTypeACBPT EnergyTransferModeEnumType = "AC_BPT"
	// (2.1) AC bidirectional with DER control, ISO 15118-20 (amendment to -20)
	EnergyTransferModeEnumTypeACBPTDER EnergyTransferModeEnumType = "AC_BPT_DER"
	// (2.1) AC charging-only with DER control, ISO 15118-20 (amendment to -20) Note: at time of writing (July 2024)
	// not yet defined for ISO 15118-20.
	EnergyTransferModeEnumTypeACDER EnergyTransferModeEnumType = "AC_DER"
	// (2.1) DC bidirectional power transfer, ISO 15118-20
	EnergyTransferModeEnumTypeDCBPT EnergyTransferModeEnumType = "DC_BPT"
	// (2.1) DC via ACDP connector (pantograph), ISO 15118-20
	EnergyTransferModeEnumTypeDCACDP EnergyTransferModeEnumType = "DC_ACDP"
	// (2.1) DC bidirectional via ACDP connector (pantograph), ISO 15118-20
	EnergyTransferModeEnumTypeDCACDPBPT EnergyTransferModeEnumType = "DC_ACDP_BPT"
	// (2.1) Wireless power transfer, ISO 15118-20
	EnergyTransferModeEnumTypeWPT EnergyTransferModeEnumType = "WPT"
)

func (s *EnergyTransferModeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch EnergyTransferModeEnumType(raw) {
	case EnergyTransferModeEnumTypeACSinglePhase,
		EnergyTransferModeEnumTypeACTwoPhase,
		EnergyTransferModeEnumTypeACThreePhase,
		EnergyTransferModeEnumTypeDC,
		EnergyTransferModeEnumTypeACBPT,
		EnergyTransferModeEnumTypeACBPTDER,
		EnergyTransferModeEnumTypeACDER,
		EnergyTransferModeEnumTypeDCBPT,
		EnergyTransferModeEnumTypeDCACDP,
		EnergyTransferModeEnumTypeDCACDPBPT,
		EnergyTransferModeEnumTypeWPT:
		*s = EnergyTransferModeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid EnergyTransferModeEnumType", raw),
	)
}

// EventNotificationEnumType (3.38)
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
	// Triggered by a monitor, which is set with the setvariablemonitoringrequest message by the Charging Station
	// Operator.
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

// EventTriggerEnumType (3.39)
type EventTriggerEnumType string

const (
	// Monitored variable has passed a Lower or Upper Threshold. Also used as trigger type for a
	// HardwiredNotification.
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

// EVSEKindEnumType (3.40)
type EVSEKindEnumType string

const (
	// AC current EVSE
	EVSEKindEnumTypeAC EVSEKindEnumType = "AC"
	// DC current EVSE
	EVSEKindEnumTypeDC EVSEKindEnumType = "DC"
)

func (s *EVSEKindEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch EVSEKindEnumType(raw) {
	case EVSEKindEnumTypeAC,
		EVSEKindEnumTypeDC:
		*s = EVSEKindEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid EVSEKindEnumType", raw),
	)
}

// FirmwareStatusEnumType (3.41)
//
// Status of a firmware download.
//
// A value with "Intermediate state" in the description, is an intermediate state, update process is not
// finished.
//
// A value with "Failure end state" in the description, is an end state, update process has stopped, update
// failed.
//
// A value with "Successful end state" in the description, is an end state, update process has stopped, update
// successful.
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
	// Charging Station is not performing firmware update related tasks. Status Idle SHALL only be used as in a
	// FirmwareStatusNotificationRequest that was triggered by TriggerMessageRequest.
	FirmwareStatusEnumTypeIdle FirmwareStatusEnumType = "Idle"
	// Failure end state. Installation of new firmware has failed.
	FirmwareStatusEnumTypeInstallationFailed FirmwareStatusEnumType = "InstallationFailed"
	// Intermediate state. Firmware is being installed.
	FirmwareStatusEnumTypeInstalling FirmwareStatusEnumType = "Installing"
	// Successful end state. New firmware has successfully been installed in Charging Station.
	FirmwareStatusEnumTypeInstalled FirmwareStatusEnumType = "Installed"
	// Intermediate state. If sent before installing the firmware, it indicates the Charging Station is about to
	// reboot to start installing new firmware. If sent after installing the new firmware, it indicates the Charging
	// Station has finished installing, but requires a reboot to activate the new firmware, which will be done
	// automatically when idle. This status MAY be omitted if a reboot is an integral part of the installation and
	// cannot be reported separately.
	FirmwareStatusEnumTypeInstallRebooting FirmwareStatusEnumType = "InstallRebooting"
	// Intermediate state. Installation of the downloaded firmware is scheduled to take place on installDateTime
	// given in UpdateFirmware request.
	FirmwareStatusEnumTypeInstallScheduled FirmwareStatusEnumType = "InstallScheduled"
	// Failure end state. Verification of the new firmware (e.g. using a checksum or some other means) has failed and
	// installation will not proceed. (Final failure state)
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

// GenericDeviceModelStatusEnumType (3.42)
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

// GenericStatusEnumType (3.43)
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

// GetCertificateIdUseEnumType (3.44)
type GetCertificateIdUseEnumType string

const (
	// Use for certificate of the ISO 15118 V2G Root.
	GetCertificateIdUseEnumTypeV2GRootCertificate GetCertificateIdUseEnumType = "V2GRootCertificate"
	// Use for certificate from an eMobility Service provider. To support PnC charging with contracts from service
	// providers that not derived their certificates from the V2G root.
	GetCertificateIdUseEnumTypeMORootCertificate GetCertificateIdUseEnumType = "MORootCertificate"
	// Root certificate for verification of the CSMS certificate.
	GetCertificateIdUseEnumTypeCSMSRootCertificate GetCertificateIdUseEnumType = "CSMSRootCertificate"
	// ISO 15118 V2G certificate chain (excluding the V2GRootCertificate).
	GetCertificateIdUseEnumTypeV2GCertificateChain GetCertificateIdUseEnumType = "V2GCertificateChain"
	// Root certificate for verification of the Manufacturer certificate.
	GetCertificateIdUseEnumTypeManufacturerRootCertificate GetCertificateIdUseEnumType = "ManufacturerRootCertificate"
	// (2.1) OEM root certificate for 2-way TLS with EV.
	GetCertificateIdUseEnumTypeOEMRootCertificate GetCertificateIdUseEnumType = "OEMRootCertificate"
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
		GetCertificateIdUseEnumTypeManufacturerRootCertificate,
		GetCertificateIdUseEnumTypeOEMRootCertificate:
		*s = GetCertificateIdUseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GetCertificateIdUseEnumType", raw),
	)
}

// GetCertificateStatusEnumType (3.45)
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

// GetChargingProfileStatusEnumType (3.46)
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

// GetDisplayMessagesStatusEnumType (3.47)
type GetDisplayMessagesStatusEnumType string

const (
	// Request accepted, there are Display Messages found that match all the requested criteria. The Charging Station
	// will send NotifyDisplayMessagesRequest messages to report the requested Display Messages.
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

// GetInstalledCertificateStatusEnumType (3.48)
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

// GetVariableStatusEnumType (3.49)
type GetVariableStatusEnumType string

const (
	// Variable successfully retrieved.
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

// GridEventFaultEnumType (3.50)
type GridEventFaultEnumType string

const (
	// Current imbalance detected
	GridEventFaultEnumTypeCurrentImbalance GridEventFaultEnumType = "CurrentImbalance"
	// A local emergency detected
	GridEventFaultEnumTypeLocalEmergency GridEventFaultEnumType = "LocalEmergency"
	// Low input power detected
	GridEventFaultEnumTypeLowInputPower GridEventFaultEnumType = "LowInputPower"
	// Overcurrent detected
	GridEventFaultEnumTypeOverCurrent GridEventFaultEnumType = "OverCurrent"
	// Over frequency detected
	GridEventFaultEnumTypeOverFrequency GridEventFaultEnumType = "OverFrequency"
	// Over voltage detected
	GridEventFaultEnumTypeOverVoltage GridEventFaultEnumType = "OverVoltage"
	// Phase rotation detected
	GridEventFaultEnumTypePhaseRotation GridEventFaultEnumType = "PhaseRotation"
	// A remote emergency detected
	GridEventFaultEnumTypeRemoteEmergency GridEventFaultEnumType = "RemoteEmergency"
	// Under frequency detected
	GridEventFaultEnumTypeUnderFrequency GridEventFaultEnumType = "UnderFrequency"
	// Under voltage detected
	GridEventFaultEnumTypeUnderVoltage GridEventFaultEnumType = "UnderVoltage"
	// Voltage imbalance detected
	GridEventFaultEnumTypeVoltageImbalance GridEventFaultEnumType = "VoltageImbalance"
)

func (s *GridEventFaultEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch GridEventFaultEnumType(raw) {
	case GridEventFaultEnumTypeCurrentImbalance,
		GridEventFaultEnumTypeLocalEmergency,
		GridEventFaultEnumTypeLowInputPower,
		GridEventFaultEnumTypeOverCurrent,
		GridEventFaultEnumTypeOverFrequency,
		GridEventFaultEnumTypeOverVoltage,
		GridEventFaultEnumTypePhaseRotation,
		GridEventFaultEnumTypeRemoteEmergency,
		GridEventFaultEnumTypeUnderFrequency,
		GridEventFaultEnumTypeUnderVoltage,
		GridEventFaultEnumTypeVoltageImbalance:
		*s = GridEventFaultEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid GridEventFaultEnumType", raw),
	)
}

// HashAlgorithmEnumType (3.51)
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

// InstallCertificateStatusEnumType (3.52)
type InstallCertificateStatusEnumType string

const (
	// The installation of the certificate succeeded.
	InstallCertificateStatusEnumTypeAccepted InstallCertificateStatusEnumType = "Accepted"
	// The certificate is invalid and/or incorrect OR the CSO tries to install more certificates than allowed.
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

// InstallCertificateUseEnumType (3.53)
type InstallCertificateUseEnumType string

const (
	// Use for certificate of the ISO 15118 V2G Root. A V2G Charging Station Certificate MUST be derived from one of
	// the installed V2GRootCertificate certificates.
	InstallCertificateUseEnumTypeV2GRootCertificate InstallCertificateUseEnumType = "V2GRootCertificate"
	// Use for certificate from an eMobility Service provider. To support PnC charging with contracts from service
	// providers that not derived their certificates from the V2G root.
	InstallCertificateUseEnumTypeMORootCertificate InstallCertificateUseEnumType = "MORootCertificate"
	// Root certificate for verification of the Manufacturer certificate.
	InstallCertificateUseEnumTypeManufacturerRootCertificate InstallCertificateUseEnumType = "ManufacturerRootCertificate"
	// Root certificate, used by the CA to sign the CSMS and Charging Station certificate.
	InstallCertificateUseEnumTypeCSMSRootCertificate InstallCertificateUseEnumType = "CSMSRootCertificate"
	// (2.1) OEM root certificate for 2-way TLS with EV.
	InstallCertificateUseEnumTypeOEMRootCertificate InstallCertificateUseEnumType = "OEMRootCertificate"
)

func (s *InstallCertificateUseEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch InstallCertificateUseEnumType(raw) {
	case InstallCertificateUseEnumTypeV2GRootCertificate,
		InstallCertificateUseEnumTypeMORootCertificate,
		InstallCertificateUseEnumTypeManufacturerRootCertificate,
		InstallCertificateUseEnumTypeCSMSRootCertificate,
		InstallCertificateUseEnumTypeOEMRootCertificate:
		*s = InstallCertificateUseEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid InstallCertificateUseEnumType", raw),
	)
}

// IslandingDetectionEnumType (3.54)
//
// Methods of islanding detection
type IslandingDetectionEnumType string

const (
	// No anti-island detection supported
	IslandingDetectionEnumTypeNoAntiIslandingSupport IslandingDetectionEnumType = "NoAntiIslandingSupport"
	// RoCoF - Rate of Change of Frequency
	IslandingDetectionEnumTypeRoCoF IslandingDetectionEnumType = "RoCoF"
	// Under/over voltage (UVP/OVP)
	IslandingDetectionEnumTypeUVPOVP IslandingDetectionEnumType = "UVP_OVP"
	// Under/over frequency (UFP/OFP)
	IslandingDetectionEnumTypeUFPOFP IslandingDetectionEnumType = "UFP_OFP"
	// Voltage Vector Shift
	IslandingDetectionEnumTypeVoltageVectorShift IslandingDetectionEnumType = "VoltageVectorShift"
	// Zero Crossing Detection
	IslandingDetectionEnumTypeZeroCrossingDetection IslandingDetectionEnumType = "ZeroCrossingDetection"
	// Other passive anti-island detection method supported
	IslandingDetectionEnumTypeOtherPassive IslandingDetectionEnumType = "OtherPassive"
	// Impedance measurement
	IslandingDetectionEnumTypeImpedanceMeasurement IslandingDetectionEnumType = "ImpedanceMeasurement"
	// Impedance detection at a specific frequency
	IslandingDetectionEnumTypeImpedanceAtFrequency IslandingDetectionEnumType = "ImpedanceAtFrequency"
	// Slip-mode frequency shift
	IslandingDetectionEnumTypeSlipModeFrequencyShift IslandingDetectionEnumType = "SlipModeFrequencyShift"
	// Frequency bias/Sandia frequency shift
	IslandingDetectionEnumTypeSandiaFrequencyShift IslandingDetectionEnumType = "SandiaFrequencyShift"
	// Sandia voltage shift
	IslandingDetectionEnumTypeSandiaVoltageShift IslandingDetectionEnumType = "SandiaVoltageShift"
	// Frequency jump
	IslandingDetectionEnumTypeFrequencyJump IslandingDetectionEnumType = "FrequencyJump"
	// RCL Q factor
	IslandingDetectionEnumTypeRCLQFactor IslandingDetectionEnumType = "RCLQFactor"
	// Other active anti-island detection method supported
	IslandingDetectionEnumTypeOtherActive IslandingDetectionEnumType = "OtherActive"
)

func (s *IslandingDetectionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch IslandingDetectionEnumType(raw) {
	case IslandingDetectionEnumTypeNoAntiIslandingSupport,
		IslandingDetectionEnumTypeRoCoF,
		IslandingDetectionEnumTypeUVPOVP,
		IslandingDetectionEnumTypeUFPOFP,
		IslandingDetectionEnumTypeVoltageVectorShift,
		IslandingDetectionEnumTypeZeroCrossingDetection,
		IslandingDetectionEnumTypeOtherPassive,
		IslandingDetectionEnumTypeImpedanceMeasurement,
		IslandingDetectionEnumTypeImpedanceAtFrequency,
		IslandingDetectionEnumTypeSlipModeFrequencyShift,
		IslandingDetectionEnumTypeSandiaFrequencyShift,
		IslandingDetectionEnumTypeSandiaVoltageShift,
		IslandingDetectionEnumTypeFrequencyJump,
		IslandingDetectionEnumTypeRCLQFactor,
		IslandingDetectionEnumTypeOtherActive:
		*s = IslandingDetectionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid IslandingDetectionEnumType", raw),
	)
}

// Iso15118EVCertificateStatusEnumType (3.55)
type Iso15118EVCertificateStatusEnumType string

const (
	// exiResponse included. This is no indication whether the update was successful, just that the message was
	// processed properly.
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

// LocationEnumType (3.56)
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
	// For the Charging Station (evseId = 0): measurement at network ("grid") inlet connection of the station. For
	// measurements with evseId > 0, these are measurements taken at the EVSE inlet (This can be useful for a DC
	// charger).
	LocationEnumTypeInlet LocationEnumType = "Inlet"
	// Measurement at a Connector. Default value.
	LocationEnumTypeOutlet LocationEnumType = "Outlet"
	// (2.1) Measurement taken from an upstream local grid meter of the premise. This can be useful for charging
	// stations that are connected "behind the meter" of a building, and that are able to read the building energy
	// meter.
	LocationEnumTypeUpstream LocationEnumType = "Upstream"
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
		LocationEnumTypeOutlet,
		LocationEnumTypeUpstream:
		*s = LocationEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid LocationEnumType", raw),
	)
}

// LogEnumType (3.57)
type LogEnumType string

const (
	// This contains the field definition of a diagnostics log file
	LogEnumTypeDiagnosticsLog LogEnumType = "DiagnosticsLog"
	// Sent by the CSMS to the Charging Station to request that the Charging Station uploads the security log.
	LogEnumTypeSecurityLog LogEnumType = "SecurityLog"
	// (2.1) The log of sampled measurements from the DataCollector component.
	LogEnumTypeDataCollectorLog LogEnumType = "DataCollectorLog"
)

func (s *LogEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch LogEnumType(raw) {
	case LogEnumTypeDiagnosticsLog,
		LogEnumTypeSecurityLog,
		LogEnumTypeDataCollectorLog:
		*s = LogEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid LogEnumType", raw),
	)
}

// LogStatusEnumType (3.58)
//
// Generic message response status
type LogStatusEnumType string

const (
	// Accepted this log upload. This does not mean the log file is uploaded is successfully, the Charging Station
	// will now start the log file upload.
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

// MeasurandEnumType (3.59)
//
// Allowable values of the optional "measurand" field of a Value element, as used in MeterValuesRequest and
// TransactionEventRequest with eventTypes Started, Ended and Updated. Default value of "measurand" is always
// "Energy.Active.Import.Register".
//
// Note 1: Two measurands (Current.Offered and Power.Offered) are available that are strictly speaking no
// measured values. They indicate the maximum amount of current/power that is being offered to the EV and are
// intended for use in smart charging applications. The measurands with .Setpoint are not measured values, but
// are a value that should be followed as closely as possible.
//
// Note 2: Import is energy flow from the Grid to the Charging Station, EV or other load. Export is energy flow
// from the EV to the Charging Station and/or from the Charging Station to the Grid. Except in the case of a
// meter replacement, all "Register" values relating to a single charging transaction, or a non-transactional
// consumer (e.g. Charging Station internal power supply, overall supply) MUST be monotonically increasing in
// time.
//
// Note 3: The actual quantity of energy corresponding to a reported ".Register" value is computed as the
// register value in question minus the register value recorded/reported at the start of the transaction or other
// relevant starting reference point in time. For improved auditability, ".Register" values SHOULD be reported
// exactly as they are directly read from a non-volatile register in the electrical metering hardware, and SHOULD
// NOT be re-based to zero at the start of transactions. This allows any "missing energy" between sequential
// transactions, due to hardware fault, meter replacement, mis-wiring, fraud, etc. to be identified, by allowing
// the CSMS to confirm that the starting register value of any transaction is identical to the finishing register
// value of the preceding transaction on the same connector.
//
// Note 4: Measurands that have a direction as part of the name (.Import or .Export) have non-negative values.
// Measurands with .Setpoint, .Net or .Residual in the name can have negative values.
//
// Note 5: Measurands starting with Display represent the optional ISO 15118-20 DisplayParameters.
type MeasurandEnumType string

const (
	// Instantaneous current flow from EV
	MeasurandEnumTypeCurrentExport MeasurandEnumType = "Current.Export"
	// (2.1) Maximum current EV is offered to export. Min(EV, EVSE)
	MeasurandEnumTypeCurrentExportOffered MeasurandEnumType = "Current.Export.Offered"
	// (2.1) Minimum current EV can discharge with. Max(EV, EVSE)
	MeasurandEnumTypeCurrentExportMinimum MeasurandEnumType = "Current.Export.Minimum"
	// Instantaneous current flow to EV
	MeasurandEnumTypeCurrentImport MeasurandEnumType = "Current.Import"
	// (2.1) Maximum current offered to EV.
	MeasurandEnumTypeCurrentImportOffered MeasurandEnumType = "Current.Import.Offered"
	// (2.1) Minimum current EV can be charged with. Max(EV, EVSE).
	MeasurandEnumTypeCurrentImportMinimum MeasurandEnumType = "Current.Import.Minimum"
	// Maximum current offered to EV. Synonymous to Current.Import.Offered.
	MeasurandEnumTypeCurrentOffered MeasurandEnumType = "Current.Offered"
	// (2.1) Current state of charge of the EV battery.
	MeasurandEnumTypeDisplayPresentSOC MeasurandEnumType = "Display.PresentSOC"
	// (2.1) Minimum State of Charge EV needs after charging of the EV battery the EV to keep throughout the charging
	// session.
	MeasurandEnumTypeDisplayMinimumSOC MeasurandEnumType = "Display.MinimumSOC"
	// (2.1) Target State of Charge of the EV battery EV needs after charging.
	MeasurandEnumTypeDisplayTargetSOC MeasurandEnumType = "Display.TargetSOC"
	// (2.1) The SOC at which the EV will prohibit any further charging.
	MeasurandEnumTypeDisplayMaximumSOC MeasurandEnumType = "Display.MaximumSOC"
	// (2.1) The remaining time it takes to reach the minimum SOC. It is communicated as the offset in seconds from
	// the point in time this value was received from EV.
	MeasurandEnumTypeDisplayRemainingTimeToMinimumSOC MeasurandEnumType = "Display.RemainingTimeToMinimumSOC"
	// (2.1) The remaining time it takes to reach the TargetSOC. It is communicated as the offset in seconds from the
	// point in time this value was received from EV.
	MeasurandEnumTypeDisplayRemainingTimeToTargetSOC MeasurandEnumType = "Display.RemainingTimeToTargetSOC"
	// (2.1) The remaining time it takes to reach the maximum SOC. It is communicated as the offset in seconds from
	// the point in time this value was received from EV.
	MeasurandEnumTypeDisplayRemainingTimeToMaximumSOC MeasurandEnumType = "Display.RemainingTimeToMaximumSOC"
	// (2.1) Indication if the charging is complete from EV point of view (value = 1).
	MeasurandEnumTypeDisplayChargingComplete MeasurandEnumType = "Display.ChargingComplete"
	// (2.1) The calculated amount of electrical Energy in Wh stored in the battery when the displayed SOC equals 100
	// %.
	MeasurandEnumTypeDisplayBatteryEnergyCapacity MeasurandEnumType = "Display.BatteryEnergyCapacity"
	// (2.1) Inlet temperature too high to accept specific operating condition.
	MeasurandEnumTypeDisplayInletHot MeasurandEnumType = "Display.InletHot"
	// Absolute amount of "active electrical energy" (Wh or kWh) exported (to the grid) during an associated time
	// "interval", specified by a Metervalues ReadingContext, and applicable interval duration configuration values
	// (in seconds) for ClockAlignedDataInterval and TxnMeterValueSampleInterval.
	MeasurandEnumTypeEnergyActiveExportInterval MeasurandEnumType = "Energy.Active.Export.Interval"
	// Numerical value read from the "active electrical energy" (Wh or kWh) register of the (most authoritative)
	// electrical meter measuring energy exported (to the grid).
	MeasurandEnumTypeEnergyActiveExportRegister MeasurandEnumType = "Energy.Active.Export.Register"
	// Absolute amount of "active electrical energy" (Wh or kWh) imported (from the grid supply) during an associated
	// time "interval", specified by a Metervalues ReadingContext, and applicable interval duration configuration
	// values (in seconds) for ClockAlignedDataInterval and TxnMeterValueSampleInterval.
	MeasurandEnumTypeEnergyActiveImportInterval MeasurandEnumType = "Energy.Active.Import.Interval"
	// Numerical value read from the "active electrical energy" (Wh or kWh) register of the (most authoritative)
	// electrical meter measuring energy imported (from the grid supply).
	MeasurandEnumTypeEnergyActiveImportRegister MeasurandEnumType = "Energy.Active.Import.Register"
	// (2.1) Calculated energy loss after energy meter. Will be reset to 0 at start of transaction. Unit is Wh.
	MeasurandEnumTypeEnergyActiveImportCableLoss MeasurandEnumType = "Energy.Active.Import.CableLoss"
	// (2.1) Cumulative amount of imported energy that was from local generation. Value will be cumulative during a
	// transaction, but is allowed to be reset to 0 at start of a transaction.
	MeasurandEnumTypeEnergyActiveImportLocalGenerationRegister MeasurandEnumType = "Energy.Active.Import.LocalGeneration.Register"
	// Numerical value read from the "net active electrical energy" (Wh or kWh) register.
	MeasurandEnumTypeEnergyActiveNet MeasurandEnumType = "Energy.Active.Net"
	// (2.1) Energy during interval when Setpoint would be followed exactly, as calculated by Charging Station.
	// Relevant when Setpoint changes frequently during an interval as result of LocalLoadBalancing or
	// LocalFrequencyControl. Can be negative if energy was exported.
	MeasurandEnumTypeEnergyActiveSetpointInterval MeasurandEnumType = "Energy.Active.Setpoint.Interval"
	// Numerical value read from the "apparent electrical export energy" (VAh or kVAh) register.
	MeasurandEnumTypeEnergyApparentExport MeasurandEnumType = "Energy.Apparent.Export"
	// Numerical value read from the "apparent electrical import energy" (VAh or kVAh) register.
	MeasurandEnumTypeEnergyApparentImport MeasurandEnumType = "Energy.Apparent.Import"
	// Numerical value read from the "apparent electrical energy" (VAh or kVAh) register.
	MeasurandEnumTypeEnergyApparentNet MeasurandEnumType = "Energy.Apparent.Net"
	// Absolute amount of "reactive electrical energy" (varh or kvarh) exported (to the grid) during an associated
	// time "interval", specified by a Metervalues ReadingContext, and applicable interval duration configuration
	// values (in seconds) for ClockAlignedDataInterval and TxnMeterValueSampleInterval.
	MeasurandEnumTypeEnergyReactiveExportInterval MeasurandEnumType = "Energy.Reactive.Export.Interval"
	// Numerical value read from the "reactive electrical energy" (varh or kvarh) register of the (most
	// authoritative) electrical meter measuring energy exported (to the grid).
	MeasurandEnumTypeEnergyReactiveExportRegister MeasurandEnumType = "Energy.Reactive.Export.Register"
	// Absolute amount of "reactive electrical energy" (varh or kvarh) imported (from the grid supply) during an
	// associated time "interval", specified by a Metervalues ReadingContext, and applicable interval duration
	// configuration values (in seconds) for ClockAlignedDataInterval and TxnMeterValueSampleInterval.
	MeasurandEnumTypeEnergyReactiveImportInterval MeasurandEnumType = "Energy.Reactive.Import.Interval"
	// Numerical value read from the "reactive electrical energy" (varh or kvarh) register of the (most
	// authoritative) electrical meter measuring energy imported (from the grid supply).
	MeasurandEnumTypeEnergyReactiveImportRegister MeasurandEnumType = "Energy.Reactive.Import.Register"
	// Numerical value read from the "net reactive electrical energy" (varh or kvarh) register.
	MeasurandEnumTypeEnergyReactiveNet MeasurandEnumType = "Energy.Reactive.Net"
	// (2.1) Energy to requested state of charge. (Wh)
	MeasurandEnumTypeEnergyRequestTarget MeasurandEnumType = "EnergyRequest.Target"
	// (2.1) Energy to minimum allowed state of charge. (Wh)
	MeasurandEnumTypeEnergyRequestMinimum MeasurandEnumType = "EnergyRequest.Minimum"
	// (2.1) Energy to maximum allowed state of charge. (Wh)
	MeasurandEnumTypeEnergyRequestMaximum MeasurandEnumType = "EnergyRequest.Maximum"
	// (2.1) Energy to minimum state of charge for cycling (V2X) activity. Positive value means that current state of
	// charge is below V2X range. (Wh)
	MeasurandEnumTypeEnergyRequestMinimumV2X MeasurandEnumType = "EnergyRequest.Minimum.V2X"
	// (2.1) Energy to maximum state of charge for cycling (V2X) activity. Negative value means that current state of
	// charge is above V2X range. (Wh)
	MeasurandEnumTypeEnergyRequestMaximumV2X MeasurandEnumType = "EnergyRequest.Maximum.V2X"
	// (2.1) Energy to end of bulk charging. (Wh)
	MeasurandEnumTypeEnergyRequestBulk MeasurandEnumType = "EnergyRequest.Bulk"
	// Instantaneous reading of powerline frequency.
	MeasurandEnumTypeFrequency MeasurandEnumType = "Frequency"
	// Instantaneous active power exported by EV. (W or kW)
	MeasurandEnumTypePowerActiveExport MeasurandEnumType = "Power.Active.Export"
	// Instantaneous active power imported by EV. (W or kW)
	MeasurandEnumTypePowerActiveImport MeasurandEnumType = "Power.Active.Import"
	// (2.1) Power setpoint for charging or discharging (negative for discharging), that should be followed as close
	// as possible.
	MeasurandEnumTypePowerActiveSetpoint MeasurandEnumType = "Power.Active.Setpoint"
	// (2.1) Difference between the given charging setpoint and the actual power measured. Can be negative.
	MeasurandEnumTypePowerActiveResidual MeasurandEnumType = "Power.Active.Residual"
	// (2.1) Minimum power the EV can be discharged with. Max(EV, EVSE)
	MeasurandEnumTypePowerExportMinimum MeasurandEnumType = "Power.Export.Minimum"
	// (2.1) Power offered to EV for discharging. Min(EV, EVSE)
	MeasurandEnumTypePowerExportOffered MeasurandEnumType = "Power.Export.Offered"
	// Instantaneous power factor of total energy flow
	MeasurandEnumTypePowerFactor MeasurandEnumType = "Power.Factor"
	// (2.1) Power offered to EV for charging. Min(EV, EVSE)
	MeasurandEnumTypePowerImportOffered MeasurandEnumType = "Power.Import.Offered"
	// (2.1) Minimum power the EV can be charged with. Max(EV, EVSE)
	MeasurandEnumTypePowerImportMinimum MeasurandEnumType = "Power.Import.Minimum"
	// Maximum power offered to EV. Synonymous to Power.Import.Offered.
	MeasurandEnumTypePowerOffered MeasurandEnumType = "Power.Offered"
	// Instantaneous reactive power exported by EV. (var or kvar)
	MeasurandEnumTypePowerReactiveExport MeasurandEnumType = "Power.Reactive.Export"
	// Instantaneous reactive power imported by EV. (var or kvar)
	MeasurandEnumTypePowerReactiveImport MeasurandEnumType = "Power.Reactive.Import"
	// State of charge of charging vehicle in percentage
	MeasurandEnumTypeSoC MeasurandEnumType = "SoC"
	// Instantaneous DC or AC RMS supply voltage. For location = Inlet and evseId = 0: voltage at charging station
	// grid connection. For location = Outlet and evseId > 0: voltage at EVSE outlet towards the EV.
	MeasurandEnumTypeVoltage MeasurandEnumType = "Voltage"
	// (2.1) Minimum voltage the EV can be charged or discharged with. Max(EV, EVSE)
	MeasurandEnumTypeVoltageMinimum MeasurandEnumType = "Voltage.Minimum"
	// (2.1) Maximum voltage the EV can be charged or discharged with. Min(EV, EVSE)
	MeasurandEnumTypeVoltageMaximum MeasurandEnumType = "Voltage.Maximum"
)

func (s *MeasurandEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MeasurandEnumType(raw) {
	case MeasurandEnumTypeCurrentExport,
		MeasurandEnumTypeCurrentExportOffered,
		MeasurandEnumTypeCurrentExportMinimum,
		MeasurandEnumTypeCurrentImport,
		MeasurandEnumTypeCurrentImportOffered,
		MeasurandEnumTypeCurrentImportMinimum,
		MeasurandEnumTypeCurrentOffered,
		MeasurandEnumTypeDisplayPresentSOC,
		MeasurandEnumTypeDisplayMinimumSOC,
		MeasurandEnumTypeDisplayTargetSOC,
		MeasurandEnumTypeDisplayMaximumSOC,
		MeasurandEnumTypeDisplayRemainingTimeToMinimumSOC,
		MeasurandEnumTypeDisplayRemainingTimeToTargetSOC,
		MeasurandEnumTypeDisplayRemainingTimeToMaximumSOC,
		MeasurandEnumTypeDisplayChargingComplete,
		MeasurandEnumTypeDisplayBatteryEnergyCapacity,
		MeasurandEnumTypeDisplayInletHot,
		MeasurandEnumTypeEnergyActiveExportInterval,
		MeasurandEnumTypeEnergyActiveExportRegister,
		MeasurandEnumTypeEnergyActiveImportInterval,
		MeasurandEnumTypeEnergyActiveImportRegister,
		MeasurandEnumTypeEnergyActiveImportCableLoss,
		MeasurandEnumTypeEnergyActiveImportLocalGenerationRegister,
		MeasurandEnumTypeEnergyActiveNet,
		MeasurandEnumTypeEnergyActiveSetpointInterval,
		MeasurandEnumTypeEnergyApparentExport,
		MeasurandEnumTypeEnergyApparentImport,
		MeasurandEnumTypeEnergyApparentNet,
		MeasurandEnumTypeEnergyReactiveExportInterval,
		MeasurandEnumTypeEnergyReactiveExportRegister,
		MeasurandEnumTypeEnergyReactiveImportInterval,
		MeasurandEnumTypeEnergyReactiveImportRegister,
		MeasurandEnumTypeEnergyReactiveNet,
		MeasurandEnumTypeEnergyRequestTarget,
		MeasurandEnumTypeEnergyRequestMinimum,
		MeasurandEnumTypeEnergyRequestMaximum,
		MeasurandEnumTypeEnergyRequestMinimumV2X,
		MeasurandEnumTypeEnergyRequestMaximumV2X,
		MeasurandEnumTypeEnergyRequestBulk,
		MeasurandEnumTypeFrequency,
		MeasurandEnumTypePowerActiveExport,
		MeasurandEnumTypePowerActiveImport,
		MeasurandEnumTypePowerActiveSetpoint,
		MeasurandEnumTypePowerActiveResidual,
		MeasurandEnumTypePowerExportMinimum,
		MeasurandEnumTypePowerExportOffered,
		MeasurandEnumTypePowerFactor,
		MeasurandEnumTypePowerImportOffered,
		MeasurandEnumTypePowerImportMinimum,
		MeasurandEnumTypePowerOffered,
		MeasurandEnumTypePowerReactiveExport,
		MeasurandEnumTypePowerReactiveImport,
		MeasurandEnumTypeSoC,
		MeasurandEnumTypeVoltage,
		MeasurandEnumTypeVoltageMinimum,
		MeasurandEnumTypeVoltageMaximum:
		*s = MeasurandEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MeasurandEnumType", raw),
	)
}

// MessageFormatEnumType (3.60)
//
// Format of a message to be displayed on the display of the Charging Station.
type MessageFormatEnumType string

const (
	// Message content is ASCII formatted, only 7-bit printable ASCII allowed.
	MessageFormatEnumTypeASCII MessageFormatEnumType = "ASCII"
	// Message content is HTML formatted.
	MessageFormatEnumTypeHTML MessageFormatEnumType = "HTML"
	// Message content is URI that Charging Station should download and use to display. for example a HTML page to be
	// shown in a web-browser.
	MessageFormatEnumTypeURI MessageFormatEnumType = "URI"
	// Message content is UTF-8 formatted.
	MessageFormatEnumTypeUTF8 MessageFormatEnumType = "UTF8"
	// Message content is a text (usually a URL) that Charging Station will display as a QR code on the display.
	// Note: this is not a dynamic QR code and should not be used for payments.
	MessageFormatEnumTypeQRCODE MessageFormatEnumType = "QRCODE"
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
		MessageFormatEnumTypeUTF8,
		MessageFormatEnumTypeQRCODE:
		*s = MessageFormatEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageFormatEnumType", raw),
	)
}

// MessagePriorityEnumType (3.61)
//
// Priority with which a message should be displayed on a Charging Station.
type MessagePriorityEnumType string

const (
	// Show this message always in front. Highest priority, don’t cycle with other messages. When a newer message
	// with this MessagePriority is received, this message is replaced. No Charging Station own message may override
	// this message.
	MessagePriorityEnumTypeAlwaysFront MessagePriorityEnumType = "AlwaysFront"
	// Show this message in front of the normal cycle of messages. When more messages with this priority are to be
	// shown, they SHALL be cycled.
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

// MessageStateEnumType (3.62)
//
// State of the Charging Station during which a message SHALL be displayed.
type MessageStateEnumType string

const (
	// Message only to be shown while the Charging Station is charging.
	MessageStateEnumTypeCharging MessageStateEnumType = "Charging"
	// Message only to be shown while the Charging Station is in faulted state.
	MessageStateEnumTypeFaulted MessageStateEnumType = "Faulted"
	// Message only to be shown while the Charging Station is idle (no transaction active).
	MessageStateEnumTypeIdle MessageStateEnumType = "Idle"
	// Message only to be shown while the Charging Station is in unavailable state.
	MessageStateEnumTypeUnavailable MessageStateEnumType = "Unavailable"
	// (2.1) Message only to be shown when Charging Station (or EV) has suspending the charging during a transaction.
	MessageStateEnumTypeSuspended MessageStateEnumType = "Suspended"
	// (2.1) Message only to be shown while the EV is discharging.
	MessageStateEnumTypeDischarging MessageStateEnumType = "Discharging"
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
		MessageStateEnumTypeUnavailable,
		MessageStateEnumTypeSuspended,
		MessageStateEnumTypeDischarging:
		*s = MessageStateEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageStateEnumType", raw),
	)
}

// MessageTriggerEnumType (3.63)
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
	// To trigger a SignCertificate with typeOfCertificate: ChargingStationCertificate.
	MessageTriggerEnumTypeSignChargingStationCertificate MessageTriggerEnumType = "SignChargingStationCertificate"
	// To trigger a SignCertificate with typeOfCertificate: V2GCertificate
	MessageTriggerEnumTypeSignV2GCertificate MessageTriggerEnumType = "SignV2GCertificate"
	// (2.1) Same as SignV2GCertificate, but this triggers Charging Station explicitly to only sign V2G certificate
	// for ISO 15118-20.
	MessageTriggerEnumTypeSignV2G20Certificate MessageTriggerEnumType = "SignV2G20Certificate"
	// To trigger StatusNotification.
	MessageTriggerEnumTypeStatusNotification MessageTriggerEnumType = "StatusNotification"
	// To trigger TransactionEvent.
	MessageTriggerEnumTypeTransactionEvent MessageTriggerEnumType = "TransactionEvent"
	// To trigger a SignCertificate with typeOfCertificate: ChargingStationCertificate AND V2GCertificate
	MessageTriggerEnumTypeSignCombinedCertificate MessageTriggerEnumType = "SignCombinedCertificate"
	// To trigger PublishFirmwareStatusNotification.
	MessageTriggerEnumTypePublishFirmwareStatusNotification MessageTriggerEnumType = "PublishFirmwareStatusNotification"
	// (2.1) To trigger the message referred to in customTrigger field.
	MessageTriggerEnumTypeCustomTrigger MessageTriggerEnumType = "CustomTrigger"
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
		MessageTriggerEnumTypeSignV2G20Certificate,
		MessageTriggerEnumTypeStatusNotification,
		MessageTriggerEnumTypeTransactionEvent,
		MessageTriggerEnumTypeSignCombinedCertificate,
		MessageTriggerEnumTypePublishFirmwareStatusNotification,
		MessageTriggerEnumTypeCustomTrigger:
		*s = MessageTriggerEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MessageTriggerEnumType", raw),
	)
}

// MobilityNeedsModeEnumType (3.64)
//
// (2.1) ISO 15118-20 service parameter for mobility needs mode.
type MobilityNeedsModeEnumType string

const (
	// Only EV determines min/target SOC and departure time.
	MobilityNeedsModeEnumTypeEVCC MobilityNeedsModeEnumType = "EVCC"
	// Charging station or CSMS may also update min/target SOC and departure time.
	MobilityNeedsModeEnumTypeEVCCSECC MobilityNeedsModeEnumType = "EVCC_SECC"
)

func (s *MobilityNeedsModeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch MobilityNeedsModeEnumType(raw) {
	case MobilityNeedsModeEnumTypeEVCC,
		MobilityNeedsModeEnumTypeEVCCSECC:
		*s = MobilityNeedsModeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MobilityNeedsModeEnumType", raw),
	)
}

// MonitorEnumType (3.65)
type MonitorEnumType string

const (
	// Triggers an event notice when the actual value of the Variable rises above value.
	MonitorEnumTypeUpperThreshold MonitorEnumType = "UpperThreshold"
	// Triggers an event notice when the actual value of the Variable drops below value.
	MonitorEnumTypeLowerThreshold MonitorEnumType = "LowerThreshold"
	// Triggers an event notice when the actual value has changed more than plus or minus value since the time that
	// this monitor was set or since the last time this event notice was sent, whichever was last. For variables that
	// are not numeric, like boolean, string or enumerations, a monitor of type Delta will trigger an event notice
	// whenever the variable changes, regardless of the value of value.
	MonitorEnumTypeDelta MonitorEnumType = "Delta"
	// Triggers an event notice every value seconds interval, starting from the time that this monitor was set.
	MonitorEnumTypePeriodic MonitorEnumType = "Periodic"
	// Triggers an event notice every value seconds interval, starting from the nearest clock-aligned interval after
	// this monitor was set. For example, a value of 900 will trigger event notices at 0, 15, 30 and 45 minutes after
	// the hour, every hour.
	MonitorEnumTypePeriodicClockAligned MonitorEnumType = "PeriodicClockAligned"
	// (2.1) Triggers an event notice when the actual value differs from the target value more than plus or minus
	// value since the time that this monitor was set or since the last time this event notice was sent, whichever
	// was last. Behavior of this type of monitor for a variable that is not numeric, is not defined. Example: when
	// target = 100, value = 10, then an event is triggered when actual < 90 or actual > 110.
	MonitorEnumTypeTargetDelta MonitorEnumType = "TargetDelta"
	// (2.1) Triggers an event notice when the actual value differs from the target value more than plus or minus
	// (value * target value) since the time that this monitor was set or since the last time this event notice was
	// sent, whichever was last. Behavior of this type of monitor for a variable that is not numeric, is not defined.
	// Example: when target = 100, value = 0.1, then an event is triggered when actual < 90 or actual > 110.
	MonitorEnumTypeTargetDeltaRelative MonitorEnumType = "TargetDeltaRelative"
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
		MonitorEnumTypePeriodicClockAligned,
		MonitorEnumTypeTargetDelta,
		MonitorEnumTypeTargetDeltaRelative:
		*s = MonitorEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid MonitorEnumType", raw),
	)
}

// MonitoringBaseEnumType (3.66)
type MonitoringBaseEnumType string

const (
	// Activate all pre-configured monitors while leaving custom monitors intact, including those that overrule a
	// pre-configured monitor.
	MonitoringBaseEnumTypeAll MonitoringBaseEnumType = "All"
	// (Re)activate the default monitors of the charging station and remove all custom monitors.
	MonitoringBaseEnumTypeFactoryDefault MonitoringBaseEnumType = "FactoryDefault"
	// Removes all custom monitors and disables all pre-configured monitors.
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

// MonitoringCriterionEnumType (3.67)
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

// MutabilityEnumType (3.68)
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

// NotifyAllowedEnergyTransferStatusEnumType (3.69)
//
// (2.1) Status result of a NotifyAllowedEnergyTransferRequest
type NotifyAllowedEnergyTransferStatusEnumType string

const (
	// Request has been accepted.
	NotifyAllowedEnergyTransferStatusEnumTypeAccepted NotifyAllowedEnergyTransferStatusEnumType = "Accepted"
	// Request has been rejected. Should not occur, unless there are some technical problems.
	NotifyAllowedEnergyTransferStatusEnumTypeRejected NotifyAllowedEnergyTransferStatusEnumType = "Rejected"
)

func (s *NotifyAllowedEnergyTransferStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch NotifyAllowedEnergyTransferStatusEnumType(raw) {
	case NotifyAllowedEnergyTransferStatusEnumTypeAccepted,
		NotifyAllowedEnergyTransferStatusEnumTypeRejected:
		*s = NotifyAllowedEnergyTransferStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid NotifyAllowedEnergyTransferStatusEnumType", raw),
	)
}

// NotifyEVChargingNeedsStatusEnumType (3.70)
type NotifyEVChargingNeedsStatusEnumType string

const (
	// A schedule will be provided momentarily.
	NotifyEVChargingNeedsStatusEnumTypeAccepted NotifyEVChargingNeedsStatusEnumType = "Accepted"
	// (2.1) Service not available. No charging profile can be provided. For an ISO 15118-20 session this is used to
	// convey that the requested energy transfer type is not possible.
	NotifyEVChargingNeedsStatusEnumTypeRejected NotifyEVChargingNeedsStatusEnumType = "Rejected"
	// The CSMS is gathering information to provide a schedule.
	NotifyEVChargingNeedsStatusEnumTypeProcessing NotifyEVChargingNeedsStatusEnumType = "Processing"
	// (2.1) CSMS will not provide a charging profile at this time. CS should not wait for it. For an ISO 15118-20
	// session this value is used instead of Rejected to differentiate between the situation where no charging
	// profile is available (NoChargingProfile) and requested energy transfer type is not available (Rejected).
	NotifyEVChargingNeedsStatusEnumTypeNoChargingProfile NotifyEVChargingNeedsStatusEnumType = "NoChargingProfile"
)

func (s *NotifyEVChargingNeedsStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch NotifyEVChargingNeedsStatusEnumType(raw) {
	case NotifyEVChargingNeedsStatusEnumTypeAccepted,
		NotifyEVChargingNeedsStatusEnumTypeRejected,
		NotifyEVChargingNeedsStatusEnumTypeProcessing,
		NotifyEVChargingNeedsStatusEnumTypeNoChargingProfile:
		*s = NotifyEVChargingNeedsStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid NotifyEVChargingNeedsStatusEnumType", raw),
	)
}

// OCPPInterfaceEnumType (3.71)
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
	// (2.1) Use any interface.
	OCPPInterfaceEnumTypeAny OCPPInterfaceEnumType = "Any"
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
		OCPPInterfaceEnumTypeWireless3,
		OCPPInterfaceEnumTypeAny:
		*s = OCPPInterfaceEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OCPPInterfaceEnumType", raw),
	)
}

// OCPPTransportEnumType (3.72)
//
// Enumeration of OCPP transport mechanisms. SOAP is currently not a valid value for OCPP 2.0.
type OCPPTransportEnumType string

const (
	// Use SOAP for transport of OCPP PDU’s
	OCPPTransportEnumTypeSOAP OCPPTransportEnumType = "SOAP"
	// Use JSON over WebSockets for transport of OCPP PDU’s
	OCPPTransportEnumTypeJSON OCPPTransportEnumType = "JSON"
)

func (s *OCPPTransportEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OCPPTransportEnumType(raw) {
	case OCPPTransportEnumTypeSOAP,
		OCPPTransportEnumTypeJSON:
		*s = OCPPTransportEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OCPPTransportEnumType", raw),
	)
}

// OCPPVersionEnumType (3.73)
//
// Enumeration of OCPP versions.
type OCPPVersionEnumType string

const (
	// OCPP version 1.2
	OCPPVersionEnumTypeOCPP12 OCPPVersionEnumType = "OCPP12"
	// OCPP version 1.5
	OCPPVersionEnumTypeOCPP15 OCPPVersionEnumType = "OCPP15"
	// OCPP version 1.6, websocket subprotocol: ocpp1.6
	OCPPVersionEnumTypeOCPP16 OCPPVersionEnumType = "OCPP16"
	// No longer in use. The OCPP 2.0 release of OCPP has been withdrawn. The value OCPP20 is treated as OCPP2.0.1.
	OCPPVersionEnumTypeOCPP20 OCPPVersionEnumType = "OCPP20"
	// OCPP version 2.0.1, websocket subprotocol: ocpp2.0.1
	OCPPVersionEnumTypeOCPP201 OCPPVersionEnumType = "OCPP201"
	// (2.1) OCPP version 2.1, websocket subprotocol: ocpp2.1
	OCPPVersionEnumTypeOCPP21 OCPPVersionEnumType = "OCPP21"
)

func (s *OCPPVersionEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OCPPVersionEnumType(raw) {
	case OCPPVersionEnumTypeOCPP12,
		OCPPVersionEnumTypeOCPP15,
		OCPPVersionEnumTypeOCPP16,
		OCPPVersionEnumTypeOCPP20,
		OCPPVersionEnumTypeOCPP201,
		OCPPVersionEnumTypeOCPP21:
		*s = OCPPVersionEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OCPPVersionEnumType", raw),
	)
}

// OperationalStatusEnumType (3.74)
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

// OperationModeEnumType (3.75)
//
// (2.1) Operation mode for (bi-)directional charging during a charging schedule period.
type OperationModeEnumType string

const (
	// Minimize energy consumption by having the EV either on standby or in sleep.
	OperationModeEnumTypeIdle OperationModeEnumType = "Idle"
	// Classic charging or smart charging mode. (default)
	OperationModeEnumTypeChargingOnly OperationModeEnumType = "ChargingOnly"
	// Control of setpoint by CSMS or some secondary actor that relays through the CSMS.
	OperationModeEnumTypeCentralSetpoint OperationModeEnumType = "CentralSetpoint"
	// Control of setpoint by an external actor directly on the Charging Station.
	OperationModeEnumTypeExternalSetpoint OperationModeEnumType = "ExternalSetpoint"
	// Control of (dis)charging limits by an external actor on the Charging Station.
	OperationModeEnumTypeExternalLimits OperationModeEnumType = "ExternalLimits"
	// Frequency support with control by CSMS or some secondary actor that relays through the CSMS.
	OperationModeEnumTypeCentralFrequency OperationModeEnumType = "CentralFrequency"
	// Frequency support with control in the Charging Station.
	OperationModeEnumTypeLocalFrequency OperationModeEnumType = "LocalFrequency"
	// Load-balancing performed by the Charging Station.
	OperationModeEnumTypeLocalLoadBalancing OperationModeEnumType = "LocalLoadBalancing"
)

func (s *OperationModeEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch OperationModeEnumType(raw) {
	case OperationModeEnumTypeIdle,
		OperationModeEnumTypeChargingOnly,
		OperationModeEnumTypeCentralSetpoint,
		OperationModeEnumTypeExternalSetpoint,
		OperationModeEnumTypeExternalLimits,
		OperationModeEnumTypeCentralFrequency,
		OperationModeEnumTypeLocalFrequency,
		OperationModeEnumTypeLocalLoadBalancing:
		*s = OperationModeEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid OperationModeEnumType", raw),
	)
}

// PaymentStatusEnumType (3.76)
//
// (2.1) Status of the settlement of an ad hoc payment.
type PaymentStatusEnumType string

const (
	// Settled successfully by the PSP.
	PaymentStatusEnumTypeSettled PaymentStatusEnumType = "Settled"
	// No billable part of the OCPP transaction, cancelation sent to the PSP.
	PaymentStatusEnumTypeCanceled PaymentStatusEnumType = "Canceled"
	// Rejected by the PSP.
	PaymentStatusEnumTypeRejected PaymentStatusEnumType = "Rejected"
	// Sent after the final attempt that fails due to communication problems.
	PaymentStatusEnumTypeFailed PaymentStatusEnumType = "Failed"
)

func (s *PaymentStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch PaymentStatusEnumType(raw) {
	case PaymentStatusEnumTypeSettled,
		PaymentStatusEnumTypeCanceled,
		PaymentStatusEnumTypeRejected,
		PaymentStatusEnumTypeFailed:
		*s = PaymentStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PaymentStatusEnumType", raw),
	)
}

// PhaseEnumType (3.77)
//
// Phase specifies how a measured value is to be interpreted. Please note that not all values of Phase are
// applicable to all Measurands.
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

// PowerDuringCessationEnumType (3.78)
type PowerDuringCessationEnumType string

const (
	// Active power
	PowerDuringCessationEnumTypeActive PowerDuringCessationEnumType = "Active"
	// Reactive power
	PowerDuringCessationEnumTypeReactive PowerDuringCessationEnumType = "Reactive"
)

func (s *PowerDuringCessationEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch PowerDuringCessationEnumType(raw) {
	case PowerDuringCessationEnumTypeActive,
		PowerDuringCessationEnumTypeReactive:
		*s = PowerDuringCessationEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PowerDuringCessationEnumType", raw),
	)
}

// PreconditioningStatusEnumType (3.79)
//
// (2.1) Preconditioning status of the battery
type PreconditioningStatusEnumType string

const (
	// No information available on the status of preconditioning
	PreconditioningStatusEnumTypeUnknown PreconditioningStatusEnumType = "Unknown"
	// The battery is preconditioned and ready to react directly on a given setpoint for charging (and discharging
	// when available).
	PreconditioningStatusEnumTypeReady PreconditioningStatusEnumType = "Ready"
	// Busy with preconditioning the BMS. When done will move to status Ready.
	PreconditioningStatusEnumTypeNotReady PreconditioningStatusEnumType = "NotReady"
	// The battery is not preconditioned and not able to directly react to given setpoint.
	PreconditioningStatusEnumTypePreconditioning PreconditioningStatusEnumType = "Preconditioning"
)

func (s *PreconditioningStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch PreconditioningStatusEnumType(raw) {
	case PreconditioningStatusEnumTypeUnknown,
		PreconditioningStatusEnumTypeReady,
		PreconditioningStatusEnumTypeNotReady,
		PreconditioningStatusEnumTypePreconditioning:
		*s = PreconditioningStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PreconditioningStatusEnumType", raw),
	)
}

// PriorityChargingStatusEnumType (3.80)
//
// (2.1) Status of a UsePriorityChargingRequest
type PriorityChargingStatusEnumType string

const (
	// Request has been accepted.
	PriorityChargingStatusEnumTypeAccepted PriorityChargingStatusEnumType = "Accepted"
	// Request has been rejected.
	PriorityChargingStatusEnumTypeRejected PriorityChargingStatusEnumType = "Rejected"
	// No priority charging profile present.
	PriorityChargingStatusEnumTypeNoProfile PriorityChargingStatusEnumType = "NoProfile"
)

func (s *PriorityChargingStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch PriorityChargingStatusEnumType(raw) {
	case PriorityChargingStatusEnumTypeAccepted,
		PriorityChargingStatusEnumTypeRejected,
		PriorityChargingStatusEnumTypeNoProfile:
		*s = PriorityChargingStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PriorityChargingStatusEnumType", raw),
	)
}

// PublishFirmwareStatusEnumType (3.81)
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
	// Failure end state. The firmware checksum is not matching.
	PublishFirmwareStatusEnumTypeInvalidChecksum PublishFirmwareStatusEnumType = "InvalidChecksum"
	// Intermediate state. The Firmware checksum is successfully verified.
	PublishFirmwareStatusEnumTypeChecksumVerified PublishFirmwareStatusEnumType = "ChecksumVerified"
	// Publishing the new firmware has failed.
	PublishFirmwareStatusEnumTypePublishFailed PublishFirmwareStatusEnumType = "PublishFailed"
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
		PublishFirmwareStatusEnumTypeDownloadPaused,
		PublishFirmwareStatusEnumTypeInvalidChecksum,
		PublishFirmwareStatusEnumTypeChecksumVerified,
		PublishFirmwareStatusEnumTypePublishFailed:
		*s = PublishFirmwareStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid PublishFirmwareStatusEnumType", raw),
	)
}

// ReadingContextEnumType (3.82)
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

// ReasonEnumType (3.83)
//
// Reason for stopping a transaction.
//
// Each reason has a (Failed) or (Successful) label, that indicates whether this situation refers to a failed or
// successful charging session according to Charging Station. There may be situations, though, where a
// transaction is ended normally (e.g. stoppedReason = Local), but no energy was delivered because of a failure
// in EVSE or EV.
type ReasonEnumType string

const (
	// The transaction was stopped because of the authorization status in the response to a TransactionEventRequest.
	// (Failed)
	ReasonEnumTypeDeAuthorized ReasonEnumType = "DeAuthorized"
	// Emergency stop button was used. (Failed)
	ReasonEnumTypeEmergencyStop ReasonEnumType = "EmergencyStop"
	// (2.1) Deprecated, because it stops energy transfer, not the transaction. EV charging session reached a locally
	// enforced maximum energy transfer limit. (Successful)
	ReasonEnumTypeEnergyLimitReached ReasonEnumType = "EnergyLimitReached"
	// Disconnecting of cable, vehicle moved away from inductive charge unit. (Successful)
	ReasonEnumTypeEVDisconnected ReasonEnumType = "EVDisconnected"
	// A GroundFault has occurred. (Failed)
	ReasonEnumTypeGroundFault ReasonEnumType = "GroundFault"
	// A Reset(Immediate) command was received. (Failed)
	ReasonEnumTypeImmediateReset ReasonEnumType = "ImmediateReset"
	// The transaction was stopped using a token that belongs to the MasterPassGroupId. (Successful)
	ReasonEnumTypeMasterPass ReasonEnumType = "MasterPass"
	// Stopped locally on request of the EV Driver at the Charge Point. This is a regular termination of a
	// transaction. Examples: presenting an IdToken tag, pressing a button to stop. (Successful)
	ReasonEnumTypeLocal ReasonEnumType = "Local"
	// (2.1) Deprecated, because it stops energy transfer, not the transaction. A local credit limit enforced through
	// the Charging Station has been exceeded. (Successful)
	ReasonEnumTypeLocalOutOfCredit ReasonEnumType = "LocalOutOfCredit"
	// Any other reason. (Failed)
	ReasonEnumTypeOther ReasonEnumType = "Other"
	// A larger than intended electric current has occurred. (Failed)
	ReasonEnumTypeOvercurrentFault ReasonEnumType = "OvercurrentFault"
	// Complete loss of power. (Failed)
	ReasonEnumTypePowerLoss ReasonEnumType = "PowerLoss"
	// Quality of power too low, e.g. voltage too low/high, phase imbalance, etc. (Failed)
	ReasonEnumTypePowerQuality ReasonEnumType = "PowerQuality"
	// A locally initiated reset/reboot occurred. (for instance watchdog kicked in). (Failed)
	ReasonEnumTypeReboot ReasonEnumType = "Reboot"
	// Stopped remotely on request of the CSMS. This is a regular termination of a transaction. Examples: termination
	// using a smartphone app, exceeding a (non local) prepaid credit. (Successful)
	ReasonEnumTypeRemote ReasonEnumType = "Remote"
	// (2.1) Deprecated, because it stops energy transfer, not the transaction. Electric vehicle has reported
	// reaching a locally enforced maximum battery State of Charge (SOC). (Successful)
	ReasonEnumTypeSOCLimitReached ReasonEnumType = "SOCLimitReached"
	// The transaction was stopped by the EV. (Successful)
	ReasonEnumTypeStoppedByEV ReasonEnumType = "StoppedByEV"
	// (2.1) Deprecated, because it stops energy transfer, not the transaction. EV charging session reached a locally
	// enforced time limit. (Successful)
	ReasonEnumTypeTimeLimitReached ReasonEnumType = "TimeLimitReached"
	// EV not connected within timeout. (Failed)
	ReasonEnumTypeTimeout ReasonEnumType = "Timeout"
	// (2.1) CSMS cannot accept the requested energy transfer type. (Failed)
	ReasonEnumTypeReqEnergyTransferRejected ReasonEnumType = "ReqEnergyTransferRejected"
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
		ReasonEnumTypeMasterPass,
		ReasonEnumTypeLocal,
		ReasonEnumTypeLocalOutOfCredit,
		ReasonEnumTypeOther,
		ReasonEnumTypeOvercurrentFault,
		ReasonEnumTypePowerLoss,
		ReasonEnumTypePowerQuality,
		ReasonEnumTypeReboot,
		ReasonEnumTypeRemote,
		ReasonEnumTypeSOCLimitReached,
		ReasonEnumTypeStoppedByEV,
		ReasonEnumTypeTimeLimitReached,
		ReasonEnumTypeTimeout,
		ReasonEnumTypeReqEnergyTransferRejected:
		*s = ReasonEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReasonEnumType", raw),
	)
}

// RecurrencyKindEnumType (3.84)
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

// RegistrationStatusEnumType (3.85)
//
// Result of registration in response to BootNotificationRequest.
type RegistrationStatusEnumType string

const (
	// Charging Station is accepted by the CSMS.
	RegistrationStatusEnumTypeAccepted RegistrationStatusEnumType = "Accepted"
	// CSMS is not yet ready to accept the Charging Station. CSMS may send messages to retrieve information or
	// prepare the Charging Station.
	RegistrationStatusEnumTypePending RegistrationStatusEnumType = "Pending"
	// Charging Station is not accepted by CSMS. This may happen when the Charging Station id is not known by CSMS.
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

// ReportBaseEnumType (3.86)
type ReportBaseEnumType string

const (
	// Required. A (configuration) report that lists all Components/Variables that can be set by the operator.
	ReportBaseEnumTypeConfigurationInventory ReportBaseEnumType = "ConfigurationInventory"
	// Required. A (full) report that lists everything except monitoring settings.
	ReportBaseEnumTypeFullInventory ReportBaseEnumType = "FullInventory"
	// Optional. A (summary) report that lists Components/Variables relating to the Charging Station’s current
	// charging availability, and to any existing problem conditions. For the Charging Station Component: -
	// AvailabilityState. For each EVSE Component: - AvailabilityState. For each Connector Component: -
	// AvailabilityState (if known and different from EVSE). For all Components in an abnormal State: - Active
	// (Problem, Tripped, Overload, Fallback) variables. - Any other diagnostically relevant Variables of the
	// Components. - Include TechCode and TechInfo where available. All monitored Component.Variables in Critical or
	// Alert state shall also be included. - Charging Stations that do not have Monitoring implemented are NOT
	// REQUIRED to include Connector Availability, monitoring alerts, and MAY limit problem reporting detail to just
	// the active Problem boolean Variable.
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

// RequestStartStopStatusEnumType (3.87)
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

// ReservationUpdateStatusEnumType (3.88)
type ReservationUpdateStatusEnumType string

const (
	// The reservation is expired.
	ReservationUpdateStatusEnumTypeExpired ReservationUpdateStatusEnumType = "Expired"
	// The reservation is removed.
	ReservationUpdateStatusEnumTypeRemoved ReservationUpdateStatusEnumType = "Removed"
	// (2.1) The reservation was used, but no transaction was started.
	ReservationUpdateStatusEnumTypeNoTransaction ReservationUpdateStatusEnumType = "NoTransaction"
)

func (s *ReservationUpdateStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ReservationUpdateStatusEnumType(raw) {
	case ReservationUpdateStatusEnumTypeExpired,
		ReservationUpdateStatusEnumTypeRemoved,
		ReservationUpdateStatusEnumTypeNoTransaction:
		*s = ReservationUpdateStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ReservationUpdateStatusEnumType", raw),
	)
}

// ReserveNowStatusEnumType (3.89)
//
// Status in ReserveNowResponse.
type ReserveNowStatusEnumType string

const (
	// Reservation has been made.
	ReserveNowStatusEnumTypeAccepted ReserveNowStatusEnumType = "Accepted"
	// Reservation has not been made, because evse, connectors or specified connector are in a faulted state.
	ReserveNowStatusEnumTypeFaulted ReserveNowStatusEnumType = "Faulted"
	// Reservation has not been made. The evse or the specified connector is occupied.
	ReserveNowStatusEnumTypeOccupied ReserveNowStatusEnumType = "Occupied"
	// Reservation has not been made. Charging Station is not configured to accept reservations.
	ReserveNowStatusEnumTypeRejected ReserveNowStatusEnumType = "Rejected"
	// Reservation has not been made, because evse, connectors or specified connector are in an unavailable state.
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

// ResetEnumType (3.90)
//
// Type of reset requested.
type ResetEnumType string

const (
	// Immediate reset of the Charging Station or EVSE.
	ResetEnumTypeImmediate ResetEnumType = "Immediate"
	// Delay reset until no more transactions are active.
	ResetEnumTypeOnIdle ResetEnumType = "OnIdle"
	// (2.1) Immediate reset and resume transaction(s) afterwards
	ResetEnumTypeImmediateAndResume ResetEnumType = "ImmediateAndResume"
)

func (s *ResetEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch ResetEnumType(raw) {
	case ResetEnumTypeImmediate,
		ResetEnumTypeOnIdle,
		ResetEnumTypeImmediateAndResume:
		*s = ResetEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid ResetEnumType", raw),
	)
}

// ResetStatusEnumType (3.91)
//
// Result of ResetRequest.
type ResetStatusEnumType string

const (
	// Command will be executed.
	ResetStatusEnumTypeAccepted ResetStatusEnumType = "Accepted"
	// Command will not be executed.
	ResetStatusEnumTypeRejected ResetStatusEnumType = "Rejected"
	// Reset command is scheduled, Charging Station is busy with a process that cannot be interrupted at the moment.
	// Reset will be executed when process is finished.
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

// SendLocalListStatusEnumType (3.92)
//
// Type of update for SendLocalListRequest.
type SendLocalListStatusEnumType string

const (
	// Local Authorization List successfully updated.
	SendLocalListStatusEnumTypeAccepted SendLocalListStatusEnumType = "Accepted"
	// Failed to update the Local Authorization List.
	SendLocalListStatusEnumTypeFailed SendLocalListStatusEnumType = "Failed"
	// Version number in the request for a differential update is less or equal then version number of current list.
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

// SetMonitoringStatusEnumType (3.93)
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

// SetNetworkProfileStatusEnumType (3.94)
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

// SetVariableStatusEnumType (3.95)
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

// TariffChangeStatusEnumType (3.96)
type TariffChangeStatusEnumType string

const (
	// Tariff has been accepted.
	TariffChangeStatusEnumTypeAccepted TariffChangeStatusEnumType = "Accepted"
	// Tariff has been rejected. More info in statusInfo.
	TariffChangeStatusEnumTypeRejected TariffChangeStatusEnumType = "Rejected"
	// Tariff has too many elements and cannot be processed.
	TariffChangeStatusEnumTypeTooManyElements TariffChangeStatusEnumType = "TooManyElements"
	// A condition is not supported, or conditions are not supported at all.
	TariffChangeStatusEnumTypeConditionNotSupported TariffChangeStatusEnumType = "ConditionNotSupported"
	// Transaction does not exist or has already ended
	TariffChangeStatusEnumTypeTxNotFound TariffChangeStatusEnumType = "TxNotFound"
	// Cannot change currency during a transaction
	TariffChangeStatusEnumTypeNoCurrencyChange TariffChangeStatusEnumType = "NoCurrencyChange"
)

func (s *TariffChangeStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TariffChangeStatusEnumType(raw) {
	case TariffChangeStatusEnumTypeAccepted,
		TariffChangeStatusEnumTypeRejected,
		TariffChangeStatusEnumTypeTooManyElements,
		TariffChangeStatusEnumTypeConditionNotSupported,
		TariffChangeStatusEnumTypeTxNotFound,
		TariffChangeStatusEnumTypeNoCurrencyChange:
		*s = TariffChangeStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TariffChangeStatusEnumType", raw),
	)
}

// TariffClearStatusEnumType (3.97)
type TariffClearStatusEnumType string

const (
	// Clearing tariff has been accepted.
	TariffClearStatusEnumTypeAccepted TariffClearStatusEnumType = "Accepted"
	// Clearing tariff has been rejected. More info in statusInfo.
	TariffClearStatusEnumTypeRejected TariffClearStatusEnumType = "Rejected"
	// No tariff for EVSE of IdToken
	TariffClearStatusEnumTypeNoTariff TariffClearStatusEnumType = "NoTariff"
)

func (s *TariffClearStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TariffClearStatusEnumType(raw) {
	case TariffClearStatusEnumTypeAccepted,
		TariffClearStatusEnumTypeRejected,
		TariffClearStatusEnumTypeNoTariff:
		*s = TariffClearStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TariffClearStatusEnumType", raw),
	)
}

// TariffCostEnumType (3.98)
//
// Type of cost: normal cost calculation, or limited to a min or max value.
type TariffCostEnumType string

const (
	// Cost is result of normal cost calculation.
	TariffCostEnumTypeNormalCost TariffCostEnumType = "NormalCost"
	// Cost is the minimum cost for this tariff.
	TariffCostEnumTypeMinCost TariffCostEnumType = "MinCost"
	// Cost is the maximum cost for this tariff.
	TariffCostEnumTypeMaxCost TariffCostEnumType = "MaxCost"
)

func (s *TariffCostEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TariffCostEnumType(raw) {
	case TariffCostEnumTypeNormalCost,
		TariffCostEnumTypeMinCost,
		TariffCostEnumTypeMaxCost:
		*s = TariffCostEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TariffCostEnumType", raw),
	)
}

// TariffGetStatusEnumType (3.99)
type TariffGetStatusEnumType string

const (
	// Tariff has been accepted.
	TariffGetStatusEnumTypeAccepted TariffGetStatusEnumType = "Accepted"
	// Tariff has been rejected. More info in statusInfo.
	TariffGetStatusEnumTypeRejected TariffGetStatusEnumType = "Rejected"
	// No tariff present on Charging Station or EVSE.
	TariffGetStatusEnumTypeNoTariff TariffGetStatusEnumType = "NoTariff"
)

func (s *TariffGetStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TariffGetStatusEnumType(raw) {
	case TariffGetStatusEnumTypeAccepted,
		TariffGetStatusEnumTypeRejected,
		TariffGetStatusEnumTypeNoTariff:
		*s = TariffGetStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TariffGetStatusEnumType", raw),
	)
}

// TariffKindEnumType (3.100)
type TariffKindEnumType string

const (
	// Default tariff
	TariffKindEnumTypeDefaultTariff TariffKindEnumType = "DefaultTariff"
	// Driver-specific tariff
	TariffKindEnumTypeDriverTariff TariffKindEnumType = "DriverTariff"
)

func (s *TariffKindEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TariffKindEnumType(raw) {
	case TariffKindEnumTypeDefaultTariff,
		TariffKindEnumTypeDriverTariff:
		*s = TariffKindEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TariffKindEnumType", raw),
	)
}

// TariffSetStatusEnumType (3.101)
type TariffSetStatusEnumType string

const (
	// Tariff has been accepted.
	TariffSetStatusEnumTypeAccepted TariffSetStatusEnumType = "Accepted"
	// Tariff has been rejected. More info in statusInfo.
	TariffSetStatusEnumTypeRejected TariffSetStatusEnumType = "Rejected"
	// Tariff has too many elements and cannot be processed.
	TariffSetStatusEnumTypeTooManyElements TariffSetStatusEnumType = "TooManyElements"
	// A condition is not supported, or conditions are not supported at all.
	TariffSetStatusEnumTypeConditionNotSupported TariffSetStatusEnumType = "ConditionNotSupported"
	// TariffId already exists in Charging Station.
	TariffSetStatusEnumTypeDuplicateTariffId TariffSetStatusEnumType = "DuplicateTariffId"
)

func (s *TariffSetStatusEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TariffSetStatusEnumType(raw) {
	case TariffSetStatusEnumTypeAccepted,
		TariffSetStatusEnumTypeRejected,
		TariffSetStatusEnumTypeTooManyElements,
		TariffSetStatusEnumTypeConditionNotSupported,
		TariffSetStatusEnumTypeDuplicateTariffId:
		*s = TariffSetStatusEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TariffSetStatusEnumType", raw),
	)
}

// TransactionEventEnumType (3.102)
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

// TriggerMessageStatusEnumType (3.103)
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

// TriggerReasonEnumType (3.104)
//
// Reason that triggered a transactionEventRequest.
type TriggerReasonEnumType string

const (
	// An Abnormal Error or Fault Condition has occurred.
	TriggerReasonEnumTypeAbnormalCondition TriggerReasonEnumType = "AbnormalCondition"
	// Charging is authorized, by any means. Might be an RFID, or other authorization means.
	TriggerReasonEnumTypeAuthorized TriggerReasonEnumType = "Authorized"
	// Cable is plugged in and EVDetected.
	TriggerReasonEnumTypeCablePluggedIn TriggerReasonEnumType = "CablePluggedIn"
	// Rate of charging changed by more than LimitChangeSignificance by an external actor (e.g. an EMS).
	TriggerReasonEnumTypeChargingRateChanged TriggerReasonEnumType = "ChargingRateChanged"
	// Charging State changed.
	TriggerReasonEnumTypeChargingStateChanged TriggerReasonEnumType = "ChargingStateChanged"
	// (2.1) Maximum cost has been reached, as defined by transactionLimit.maxCost.
	TriggerReasonEnumTypeCostLimitReached TriggerReasonEnumType = "CostLimitReached"
	// The transaction was stopped because of the authorization status in the response to a transactionEventRequest.
	TriggerReasonEnumTypeDeauthorized TriggerReasonEnumType = "Deauthorized"
	// Maximum energy of charging reached as defined by transactionLimit.maxEnergy.
	TriggerReasonEnumTypeEnergyLimitReached TriggerReasonEnumType = "EnergyLimitReached"
	// Communication with EV lost, for example: cable disconnected.
	TriggerReasonEnumTypeEVCommunicationLost TriggerReasonEnumType = "EVCommunicationLost"
	// EV not connected before the connection is timed out.
	TriggerReasonEnumTypeEVConnectTimeout TriggerReasonEnumType = "EVConnectTimeout"
	// EV departed. For example: When a departing EV triggers a parking bay detector.
	TriggerReasonEnumTypeEVDeparted TriggerReasonEnumType = "EVDeparted"
	// EV detected. For example: When an arriving EV triggers a parking bay detector.
	TriggerReasonEnumTypeEVDetected TriggerReasonEnumType = "EVDetected"
	// (2.1) Limit of cost/time/energy/SoC for transaction has set or changed.
	TriggerReasonEnumTypeLimitSet TriggerReasonEnumType = "LimitSet"
	// Needed to send a clock aligned meter value
	TriggerReasonEnumTypeMeterValueClock TriggerReasonEnumType = "MeterValueClock"
	// Needed to send a periodic meter value
	TriggerReasonEnumTypeMeterValuePeriodic TriggerReasonEnumType = "MeterValuePeriodic"
	// (2.1) V2X operation mode has changed (at start of a new charging schedule period).
	TriggerReasonEnumTypeOperationModeChanged TriggerReasonEnumType = "OperationModeChanged"
	// A RequestStartTransactionRequest has been sent.
	TriggerReasonEnumTypeRemoteStart TriggerReasonEnumType = "RemoteStart"
	// A RequestStopTransactionRequest has been sent.
	TriggerReasonEnumTypeRemoteStop TriggerReasonEnumType = "RemoteStop"
	// CSMS sent a Reset Charging Station command.
	TriggerReasonEnumTypeResetCommand TriggerReasonEnumType = "ResetCommand"
	// (2.1) Trigger used when TransactionEvent is sent (only) to report a running cost update.
	TriggerReasonEnumTypeRunningCost TriggerReasonEnumType = "RunningCost"
	// Signed data is received from the energy meter.
	TriggerReasonEnumTypeSignedDataReceived TriggerReasonEnumType = "SignedDataReceived"
	// (2.1) State of charge limit has been reached, as defined by transactionLimit.maxSoC.
	TriggerReasonEnumTypeSoCLimitReached TriggerReasonEnumType = "SoCLimitReached"
	// An EV Driver has been authorized to stop charging. For example: By swiping an RFID card.
	TriggerReasonEnumTypeStopAuthorized TriggerReasonEnumType = "StopAuthorized"
	// (2.1) Tariff for transaction has changed.
	TriggerReasonEnumTypeTariffChanged TriggerReasonEnumType = "TariffChanged"
	// (2.1) Trigger to notify that EV Driver has not accepted the tariff for transaction. idToken becomes
	// deauthorized.
	TriggerReasonEnumTypeTariffNotAccepted TriggerReasonEnumType = "TariffNotAccepted"
	// (2.1) Maximum time of charging reached, as defined by transactionLimit.maxTime.
	TriggerReasonEnumTypeTimeLimitReached TriggerReasonEnumType = "TimeLimitReached"
	// Requested by the CSMS via a TriggerMessageRequest.
	TriggerReasonEnumTypeTrigger TriggerReasonEnumType = "Trigger"
	// (2.1) Transaction has resumed after reset or power outage.
	TriggerReasonEnumTypeTxResumed TriggerReasonEnumType = "TxResumed"
	// CSMS sent an Unlock Connector command.
	TriggerReasonEnumTypeUnlockCommand TriggerReasonEnumType = "UnlockCommand"
)

func (s *TriggerReasonEnumType) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch TriggerReasonEnumType(raw) {
	case TriggerReasonEnumTypeAbnormalCondition,
		TriggerReasonEnumTypeAuthorized,
		TriggerReasonEnumTypeCablePluggedIn,
		TriggerReasonEnumTypeChargingRateChanged,
		TriggerReasonEnumTypeChargingStateChanged,
		TriggerReasonEnumTypeCostLimitReached,
		TriggerReasonEnumTypeDeauthorized,
		TriggerReasonEnumTypeEnergyLimitReached,
		TriggerReasonEnumTypeEVCommunicationLost,
		TriggerReasonEnumTypeEVConnectTimeout,
		TriggerReasonEnumTypeEVDeparted,
		TriggerReasonEnumTypeEVDetected,
		TriggerReasonEnumTypeLimitSet,
		TriggerReasonEnumTypeMeterValueClock,
		TriggerReasonEnumTypeMeterValuePeriodic,
		TriggerReasonEnumTypeOperationModeChanged,
		TriggerReasonEnumTypeRemoteStart,
		TriggerReasonEnumTypeRemoteStop,
		TriggerReasonEnumTypeResetCommand,
		TriggerReasonEnumTypeRunningCost,
		TriggerReasonEnumTypeSignedDataReceived,
		TriggerReasonEnumTypeSoCLimitReached,
		TriggerReasonEnumTypeStopAuthorized,
		TriggerReasonEnumTypeTariffChanged,
		TriggerReasonEnumTypeTariffNotAccepted,
		TriggerReasonEnumTypeTimeLimitReached,
		TriggerReasonEnumTypeTrigger,
		TriggerReasonEnumTypeTxResumed,
		TriggerReasonEnumTypeUnlockCommand:
		*s = TriggerReasonEnumType(raw)
		return nil
	}

	return ocpp.NewError(
		ocpp.ErrPropertyConstraintViolation,
		"",
		fmt.Sprintf("%s is an invalid TriggerReasonEnumType", raw),
	)
}

// UnlockStatusEnumType (3.105)
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

// UnpublishFirmwareStatusEnumType (3.106)
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

// UpdateEnumType (3.107)
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

// UpdateFirmwareStatusEnumType (3.108)
//
// Generic message response status
type UpdateFirmwareStatusEnumType string

const (
	// Accepted this firmware update request. This does not mean the firmware update is successful, the Charging
	// Station will now start the firmware update process.
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

// UploadLogStatusEnumType (3.109)
type UploadLogStatusEnumType string

const (
	// A badly formatted packet or other protocol incompatibility was detected.
	UploadLogStatusEnumTypeBadMessage UploadLogStatusEnumType = "BadMessage"
	// The Charging Station is not uploading a log file. Idle SHALL only be used when the message was triggered by a
	// TriggerMessageRequest.
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

// VPNEnumType (3.110)
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
