package v21

import (
	"encoding/json"
	"fmt"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// AdjustPeriodicEventStreamRequest (1.1.1)
type AdjustPeriodicEventStreamRequest struct {
	// Required.
	ID int32 `json:"id"`
	// Required. Updated rate of sending data
	Params PeriodicEventStreamParamsType `json:"params"`
}

func (s *AdjustPeriodicEventStreamRequest) UnmarshalJSON(data []byte) error {
	type Alias AdjustPeriodicEventStreamRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AdjustPeriodicEventStreamRequest(a)
	return s.Validate()
}

func (s AdjustPeriodicEventStreamRequest) Validate() error {
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if err := s.Params.Validate(); err != nil {
		return ocpp.WrapField("params", err)
	}

	return nil
}

// AdjustPeriodicEventStreamResponse (1.1.2)
type AdjustPeriodicEventStreamResponse struct {
	// Required. Status of operation.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status information
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *AdjustPeriodicEventStreamResponse) UnmarshalJSON(data []byte) error {
	type Alias AdjustPeriodicEventStreamResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AdjustPeriodicEventStreamResponse(a)
	return s.Validate()
}

func (s AdjustPeriodicEventStreamResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// AFRRSignalRequest (1.2.1)
//
// (2.1) This message passes an aFRR signal on to the charging station. Charging station uses the value of signal
// to select a matching power value from the v2xSignalWattCurve in the ChargingSchedulePeriod.
type AFRRSignalRequest struct {
	// Required. Time when signal becomes active.
	Timestamp time.Time `json:"timestamp"`
	// Required. Value of signal in v2xSignalWattCurve.
	Signal int32 `json:"signal"`
}

func (s *AFRRSignalRequest) UnmarshalJSON(data []byte) error {
	type Alias AFRRSignalRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AFRRSignalRequest(a)
	return s.Validate()
}

func (s AFRRSignalRequest) Validate() error {
	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	return nil
}

// AFRRSignalResponse (1.2.2)
//
// (2.1) Response stating whether signal was accepted. Response will be Accepted if a v2xSignalWattCurve_ element
// exists in the ChargingSchedulePeriodType for that point in time.
type AFRRSignalResponse struct {
	// Required.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Additional information on status.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *AFRRSignalResponse) UnmarshalJSON(data []byte) error {
	type Alias AFRRSignalResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AFRRSignalResponse(a)
	return s.Validate()
}

func (s AFRRSignalResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// AuthorizeRequest (1.3.1)
//
// This contains the field definition of the AuthorizeRequest PDU sent by the Charging Station to the CSMS.
type AuthorizeRequest struct {
	// Optional. (2.1) The X.509 certificate chain presented by EV and encoded in PEM format. Order of certificates
	// in chain is from leaf up to (but excluding) root certificate. Only needed in case of central contract
	// validation when Charging Station cannot validate the contract certificate.
	Certificate *string `json:"certificate,omitempty"`
	// Required. This contains the identifier that needs to be authorized.
	IDToken IDTokenType `json:"idToken"`
	// Optional. Contains the information needed to verify the EV Contract Certificate via OCSP. Not needed if
	// certificate is provided.
	Iso15118CertificateHashData []OCSPRequestDataType `json:"iso15118CertificateHashData,omitempty"`
}

func (s *AuthorizeRequest) UnmarshalJSON(data []byte) error {
	type Alias AuthorizeRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AuthorizeRequest(a)
	return s.Validate()
}

func (s AuthorizeRequest) Validate() error {
	if s.Certificate != nil {
		if err := validateStringLen(*s.Certificate, 10000, "certificate"); err != nil {
			return err
		}
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	if err := validateSliceLen(s.Iso15118CertificateHashData, 0, 4, "iso15118CertificateHashData"); err != nil {
		return err
	}

	for i, v := range s.Iso15118CertificateHashData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("iso15118CertificateHashData[%d]", i), err)
		}
	}

	return nil
}

// AuthorizeResponse (1.3.2)
//
// This contains the field definition of the AuthorizeResponse PDU sent by the CSMS to the Charging Station in
// response to an AuthorizeRequest.
type AuthorizeResponse struct {
	// Optional. Certificate status information. - if all certificates are valid: return 'Accepted'. - if one of the
	// certificates was revoked, return 'CertificateRevoked'.
	CertificateStatus *AuthorizeCertificateStatusEnumType `json:"certificateStatus,omitempty"`
	// Optional. (2.1) List of allowed energy transfer modes the EV can choose from. If omitted this defaults to
	// charging only.
	AllowedEnergyTransfer []EnergyTransferModeEnumType `json:"allowedEnergyTransfer,omitempty"`
	// Required. This contains information about authorization status, expiry and group id.
	IDTokenInfo IDTokenInfoType `json:"idTokenInfo"`
	// Optional. (2.1) Tariff for this idToken.
	Tariff *TariffType `json:"tariff,omitempty"`
}

func (s *AuthorizeResponse) UnmarshalJSON(data []byte) error {
	type Alias AuthorizeResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AuthorizeResponse(a)
	return s.Validate()
}

func (s AuthorizeResponse) Validate() error {
	if err := s.IDTokenInfo.Validate(); err != nil {
		return ocpp.WrapField("idTokenInfo", err)
	}

	if s.Tariff != nil {
		if err := s.Tariff.Validate(); err != nil {
			return ocpp.WrapField("tariff", err)
		}
	}

	return nil
}

// BatterySwapRequest (1.4.1)
//
// (2.1) Message sent by Charging Station when a battery is swapped in or out of a battery swap station.
type BatterySwapRequest struct {
	// Required. Battery in/out
	EventType BatterySwapEventEnumType `json:"eventType"`
	// Required. RequestId to correlate BatteryIn/Out events and optional RequestBatterySwapRequest.
	RequestID int32 `json:"requestId"`
	// Required. Id token of EV Driver
	IDToken IDTokenType `json:"idToken"`
	// Required. Info on batteries inserted or taken out.
	BatteryData []BatteryDataType `json:"batteryData"`
}

func (s *BatterySwapRequest) UnmarshalJSON(data []byte) error {
	type Alias BatterySwapRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BatterySwapRequest(a)
	return s.Validate()
}

func (s BatterySwapRequest) Validate() error {
	if s.EventType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "eventType", "required field is missing")
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	if err := validateSliceLen(s.BatteryData, 1, -1, "batteryData"); err != nil {
		return err
	}

	for i, v := range s.BatteryData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("batteryData[%d]", i), err)
		}
	}

	return nil
}

// BatterySwapResponse (1.4.2)
//
// (2.1) Empty response by CSMS to confirm receipt of BatterySwapRequest.
//
// No fields are defined.
type BatterySwapResponse struct {
}

func (s *BatterySwapResponse) UnmarshalJSON(data []byte) error {
	type Alias BatterySwapResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BatterySwapResponse(a)
	return s.Validate()
}

func (s BatterySwapResponse) Validate() error {
	_ = s
	return nil
}

// BootNotificationRequest (1.5.1)
//
// This contains the field definition of the BootNotificationRequest PDU sent by the Charging Station to the
// CSMS.
type BootNotificationRequest struct {
	// Required. This contains the reason for sending this message to the CSMS.
	Reason BootReasonEnumType `json:"reason"`
	// Required. Identifies the Charging Station
	ChargingStation ChargingStationType `json:"chargingStation"`
}

func (s *BootNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias BootNotificationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BootNotificationRequest(a)
	return s.Validate()
}

func (s BootNotificationRequest) Validate() error {
	if s.Reason == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "reason", "required field is missing")
	}

	if err := s.ChargingStation.Validate(); err != nil {
		return ocpp.WrapField("chargingStation", err)
	}

	return nil
}

// BootNotificationResponse (1.5.2)
//
// This contains the field definition of the BootNotificationResponse PDU sent by the CSMS to the Charging
// Station in response to a BootNotificationRequest.
type BootNotificationResponse struct {
	// Required. This contains the CSMS’s current time.
	CurrentTime time.Time `json:"currentTime"`
	// Required. When Status is Accepted, this contains the heartbeat interval in seconds. If the CSMS returns
	// something other than Accepted, the value of the interval field indicates the minimum wait time before sending
	// a next BootNotification request.
	Interval int32 `json:"interval"`
	// Required. This contains whether the Charging Station has been registered within the CSMS.
	Status RegistrationStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *BootNotificationResponse) UnmarshalJSON(data []byte) error {
	type Alias BootNotificationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BootNotificationResponse(a)
	return s.Validate()
}

func (s BootNotificationResponse) Validate() error {
	if s.CurrentTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currentTime", "required field is missing")
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// CancelReservationRequest (1.6.1)
//
// This contains the field definition of the CancelReservationRequest PDU sent by the CSMS to the Charging
// Station.
type CancelReservationRequest struct {
	// Required. Id of the reservation to cancel.
	ReservationID int32 `json:"reservationId"`
}

func (s *CancelReservationRequest) UnmarshalJSON(data []byte) error {
	type Alias CancelReservationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CancelReservationRequest(a)
	return s.Validate()
}

func (s CancelReservationRequest) Validate() error {
	if err := validateNonNegative(s.ReservationID, "reservationId"); err != nil {
		return err
	}

	return nil
}

// CancelReservationResponse (1.6.2)
//
// This contains the field definition of the CancelReservationResponse PDU sent by the Charging Station to the
// CSMS in response to a CancelReservationRequest.
type CancelReservationResponse struct {
	// Required. This indicates the success or failure of the canceling of a reservation by CSMS.
	Status CancelReservationStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *CancelReservationResponse) UnmarshalJSON(data []byte) error {
	type Alias CancelReservationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CancelReservationResponse(a)
	return s.Validate()
}

func (s CancelReservationResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// CertificateSignedRequest (1.7.1)
//
// This contains the field definition of the CertificateSignedRequest PDU sent by the CSMS to the Charging
// Station.
type CertificateSignedRequest struct {
	// Required. The signed PEM encoded X.509 certificate. This SHALL also contain the necessary sub CA certificates,
	// when applicable. The order of the bundle follows the certificate chain, starting from the leaf certificate.
	// The Configuration Variable MaxCertificateChainSize can be used to limit the maximum size of this field.
	CertificateChain string `json:"certificateChain"`
	// Optional. Indicates the type of the signed certificate that is returned. When omitted the certificate is used
	// for both the 15118 connection (if implemented) and the Charging Station to CSMS connection. This field is
	// required when a typeOfCertificate was included in the SignCertificateRequest that requested this certificate
	// to be signed AND both the 15118 connection and the Charging Station connection are implemented.
	CertificateType *CertificateSigningUseEnumType `json:"certificateType,omitempty"`
	// Optional. (2.1) RequestId to correlate this message with the SignCertificateRequest.
	RequestID *int32 `json:"requestId,omitempty"`
}

func (s *CertificateSignedRequest) UnmarshalJSON(data []byte) error {
	type Alias CertificateSignedRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CertificateSignedRequest(a)
	return s.Validate()
}

func (s CertificateSignedRequest) Validate() error {
	if s.CertificateChain == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "certificateChain", "required field is missing")
	}

	if err := validateStringLen(s.CertificateChain, 10000, "certificateChain"); err != nil {
		return err
	}

	return nil
}

// CertificateSignedResponse (1.7.2)
//
// This contains the field definition of the CertificateSignedResponse PDU sent by the Charging Station to the
// CSMS in response to a CertificateSignedRequest.
type CertificateSignedResponse struct {
	// Required. Returns whether certificate signing has been accepted, otherwise rejected.
	Status CertificateSignedStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *CertificateSignedResponse) UnmarshalJSON(data []byte) error {
	type Alias CertificateSignedResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CertificateSignedResponse(a)
	return s.Validate()
}

func (s CertificateSignedResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ChangeAvailabilityRequest (1.8.1)
//
// This contains the field definition of the ChangeAvailabilityRequest PDU sent by the CSMS to the Charging
// Station.
type ChangeAvailabilityRequest struct {
	// Required. This contains the type of availability change that the Charging Station should perform.
	OperationalStatus OperationalStatusEnumType `json:"operationalStatus"`
	// Optional. Contains Id’s to designate a specific EVSE/connector by index numbers. When omitted, the message
	// refers to the Charging Station as a whole.
	EVSE *EVSEType `json:"evse,omitempty"`
}

func (s *ChangeAvailabilityRequest) UnmarshalJSON(data []byte) error {
	type Alias ChangeAvailabilityRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeAvailabilityRequest(a)
	return s.Validate()
}

func (s ChangeAvailabilityRequest) Validate() error {
	if s.OperationalStatus == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "operationalStatus", "required field is missing")
	}

	if s.EVSE != nil {
		if err := s.EVSE.Validate(); err != nil {
			return ocpp.WrapField("evse", err)
		}
	}

	return nil
}

// ChangeAvailabilityResponse (1.8.2)
//
// This contains the field definition of the ChangeAvailabilityResponse PDU sent by the Charging Station to the
// CSMS.
type ChangeAvailabilityResponse struct {
	// Required. This indicates whether the Charging Station is able to perform the availability change.
	Status ChangeAvailabilityStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ChangeAvailabilityResponse) UnmarshalJSON(data []byte) error {
	type Alias ChangeAvailabilityResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeAvailabilityResponse(a)
	return s.Validate()
}

func (s ChangeAvailabilityResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ChangeTransactionTariffRequest (1.9.1)
type ChangeTransactionTariffRequest struct {
	// Required. Transaction id for new tariff.
	TransactionID IdentifierString36Type `json:"transactionId"`
	// Required. New tariff to use for transaction.
	Tariff TariffType `json:"tariff"`
}

func (s *ChangeTransactionTariffRequest) UnmarshalJSON(data []byte) error {
	type Alias ChangeTransactionTariffRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeTransactionTariffRequest(a)
	return s.Validate()
}

func (s ChangeTransactionTariffRequest) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	if err := s.Tariff.Validate(); err != nil {
		return ocpp.WrapField("tariff", err)
	}

	return nil
}

// ChangeTransactionTariffResponse (1.9.2)
type ChangeTransactionTariffResponse struct {
	// Required. Status of the operation
	Status TariffChangeStatusEnumType `json:"status"`
	// Optional. Detailed status information
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ChangeTransactionTariffResponse) UnmarshalJSON(data []byte) error {
	type Alias ChangeTransactionTariffResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChangeTransactionTariffResponse(a)
	return s.Validate()
}

func (s ChangeTransactionTariffResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ClearCacheRequest (1.10.1)
//
// This contains the field definition of the ClearCacheRequest PDU sent by the CSMS to the Charging Station. No
// fields are defined.
//
// No fields are defined.
type ClearCacheRequest struct {
}

func (s *ClearCacheRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearCacheRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearCacheRequest(a)
	return s.Validate()
}

func (s ClearCacheRequest) Validate() error {
	_ = s
	return nil
}

// ClearCacheResponse (1.10.2)
//
// This contains the field definition of the ClearCacheResponse PDU sent by the Charging Station to the CSMS in
// response to a ClearCacheRequest.
type ClearCacheResponse struct {
	// Required. Accepted if the Charging Station has executed the request, otherwise rejected.
	Status ClearCacheStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ClearCacheResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearCacheResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearCacheResponse(a)
	return s.Validate()
}

func (s ClearCacheResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ClearChargingProfileRequest (1.11.1)
//
// This contains the field definition of the ClearChargingProfileRequest PDU sent by the CSMS to the Charging
// Station. The CSMS can use this message to clear (remove) either a specific charging profile (denoted by id) or
// a selection of charging profiles that match with the values of the optional evse, stackLevel and
// ChargingProfilePurpose fields.
type ClearChargingProfileRequest struct {
	// Optional. The Id of the charging profile to clear.
	ChargingProfileID *int32 `json:"chargingProfileId,omitempty"`
	// Optional. Specifies the charging profile.
	ChargingProfileCriteria *ClearChargingProfileType `json:"chargingProfileCriteria,omitempty"`
}

func (s *ClearChargingProfileRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearChargingProfileRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearChargingProfileRequest(a)
	return s.Validate()
}

func (s ClearChargingProfileRequest) Validate() error {
	if s.ChargingProfileCriteria != nil {
		if err := s.ChargingProfileCriteria.Validate(); err != nil {
			return ocpp.WrapField("chargingProfileCriteria", err)
		}
	}

	return nil
}

// ClearChargingProfileResponse (1.11.2)
//
// This contains the field definition of the ClearChargingProfileResponse PDU sent by the Charging Station to the
// CSMS in response to a ClearChargingProfileRequest.
type ClearChargingProfileResponse struct {
	// Required. Indicates if the Charging Station was able to execute the request.
	Status ClearChargingProfileStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ClearChargingProfileResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearChargingProfileResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearChargingProfileResponse(a)
	return s.Validate()
}

func (s ClearChargingProfileResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ClearDERControlRequest (1.12.1)
type ClearDERControlRequest struct {
	// Required. True: clearing default DER controls. False: clearing scheduled controls.
	IsDefault bool `json:"isDefault"`
	// Optional. Name of control settings to clear. Not used when controlId is provided.
	ControlType *DERControlEnumType `json:"controlType,omitempty"`
	// Optional. Id of control setting to clear. When omitted all settings for controlType are cleared.
	ControlID *IdentifierString36Type `json:"controlId,omitempty"`
}

func (s *ClearDERControlRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearDERControlRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearDERControlRequest(a)
	return s.Validate()
}

func (s ClearDERControlRequest) Validate() error {
	if s.ControlID != nil {
		if err := s.ControlID.Validate(); err != nil {
			return ocpp.WrapField("controlId", err)
		}
	}

	return nil
}

// ClearDERControlResponse (1.12.2)
type ClearDERControlResponse struct {
	// Required. Result of operation.
	Status DERControlStatusEnumType `json:"status"`
	// Optional. Detailed status information
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ClearDERControlResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearDERControlResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearDERControlResponse(a)
	return s.Validate()
}

func (s ClearDERControlResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ClearDisplayMessageRequest (1.13.1)
//
// This contains the field definition of the ClearDisplayMessageRequest PDU sent by the CSMS to the Charging
// Station. The CSMS asks the Charging Station to clear a display message that has been configured in the
// Charging Station to be cleared/removed. See also O05 - Clear a Display Message.
type ClearDisplayMessageRequest struct {
	// Required. Id of the message that SHALL be removed from the Charging Station.
	ID int32 `json:"id"`
}

func (s *ClearDisplayMessageRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearDisplayMessageRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearDisplayMessageRequest(a)
	return s.Validate()
}

func (s ClearDisplayMessageRequest) Validate() error {
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	return nil
}

// ClearDisplayMessageResponse (1.13.2)
//
// This contains the field definition of the ClearDisplayMessageResponse PDU sent by the Charging Station to the
// CSMS in a response to a ClearDisplayMessageRequest. See also O05 - Clear a Display Message.
type ClearDisplayMessageResponse struct {
	// Required. Returns whether the Charging Station has been able to remove the message.
	Status ClearMessageStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ClearDisplayMessageResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearDisplayMessageResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearDisplayMessageResponse(a)
	return s.Validate()
}

func (s ClearDisplayMessageResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ClearedChargingLimitRequest (1.14.1)
//
// This contains the field definition of the ClearedChargingLimitRequest PDU sent by the Charging Station to the
// CSMS.
type ClearedChargingLimitRequest struct {
	// Required. Source of the charging limit. Allowed values defined in Appendix as
	// ChargingLimitSourceEnumStringType.
	ChargingLimitSource string `json:"chargingLimitSource"`
	// Optional. EVSE Identifier.
	EVSEID *int32 `json:"evseId,omitempty"`
}

func (s *ClearedChargingLimitRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearedChargingLimitRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearedChargingLimitRequest(a)
	return s.Validate()
}

func (s ClearedChargingLimitRequest) Validate() error {
	if s.ChargingLimitSource == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingLimitSource", "required field is missing")
	}

	if err := validateStringLen(s.ChargingLimitSource, 20, "chargingLimitSource"); err != nil {
		return err
	}

	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	return nil
}

// ClearedChargingLimitResponse (1.14.2)
//
// This contains the field definition of the ClearedChargingLimitResponse PDU sent by the CSMS to the Charging
// Station. No fields are defined.
//
// No fields are defined.
type ClearedChargingLimitResponse struct {
}

func (s *ClearedChargingLimitResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearedChargingLimitResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearedChargingLimitResponse(a)
	return s.Validate()
}

func (s ClearedChargingLimitResponse) Validate() error {
	_ = s
	return nil
}

// ClearTariffsRequest (1.15.1)
type ClearTariffsRequest struct {
	// Optional. List of tariff Ids to clear. When absent clears all tariffs at evseId.
	TariffIds []string `json:"tariffIds,omitempty"`
	// Optional. When present only clear tariffs matching tariffIds at EVSE evseId.
	EVSEID *int32 `json:"evseId,omitempty"`
}

func (s *ClearTariffsRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearTariffsRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearTariffsRequest(a)
	return s.Validate()
}

func (s ClearTariffsRequest) Validate() error {
	for i, v := range s.TariffIds {
		if err := validateStringLen(v, 60, fmt.Sprintf("tariffIds[%d]", i)); err != nil {
			return err
		}
	}

	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	return nil
}

// ClearTariffsResponse (1.15.2)
type ClearTariffsResponse struct {
	// Required. Result per tariff.
	ClearTariffsResult []ClearTariffsResultType `json:"clearTariffsResult"`
}

func (s *ClearTariffsResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearTariffsResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearTariffsResponse(a)
	return s.Validate()
}

func (s ClearTariffsResponse) Validate() error {
	if err := validateSliceLen(s.ClearTariffsResult, 1, -1, "clearTariffsResult"); err != nil {
		return err
	}

	for i, v := range s.ClearTariffsResult {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("clearTariffsResult[%d]", i), err)
		}
	}

	return nil
}

// ClearVariableMonitoringRequest (1.16.1)
//
// This contains the field definition of the ClearVariableMonitoringRequest PDU sent by the CSMS to the Charging
// Station.
type ClearVariableMonitoringRequest struct {
	// Required. List of the monitors to be cleared, identified by there Id.
	ID []int32 `json:"id"`
}

func (s *ClearVariableMonitoringRequest) UnmarshalJSON(data []byte) error {
	type Alias ClearVariableMonitoringRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearVariableMonitoringRequest(a)
	return s.Validate()
}

func (s ClearVariableMonitoringRequest) Validate() error {
	if err := validateSliceLen(s.ID, 1, -1, "id"); err != nil {
		return err
	}

	for i, v := range s.ID {
		if err := validateNonNegative(v, fmt.Sprintf("id[%d]", i)); err != nil {
			return err
		}
	}

	return nil
}

// ClearVariableMonitoringResponse (1.16.2)
//
// This contains the field definition of the ClearVariableMonitoringResponse PDU sent by the Charging Station to
// the CSMS.
type ClearVariableMonitoringResponse struct {
	// Required. List of status per monitor.
	ClearMonitoringResult []ClearMonitoringResultType `json:"clearMonitoringResult"`
}

func (s *ClearVariableMonitoringResponse) UnmarshalJSON(data []byte) error {
	type Alias ClearVariableMonitoringResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearVariableMonitoringResponse(a)
	return s.Validate()
}

func (s ClearVariableMonitoringResponse) Validate() error {
	if err := validateSliceLen(s.ClearMonitoringResult, 1, -1, "clearMonitoringResult"); err != nil {
		return err
	}

	for i, v := range s.ClearMonitoringResult {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("clearMonitoringResult[%d]", i), err)
		}
	}

	return nil
}

// ClosePeriodicEventStreamRequest (1.17.1)
type ClosePeriodicEventStreamRequest struct {
	// Required. Id of stream to close.
	ID int32 `json:"id"`
}

func (s *ClosePeriodicEventStreamRequest) UnmarshalJSON(data []byte) error {
	type Alias ClosePeriodicEventStreamRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClosePeriodicEventStreamRequest(a)
	return s.Validate()
}

func (s ClosePeriodicEventStreamRequest) Validate() error {
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	return nil
}

// ClosePeriodicEventStreamResponse (1.17.2)
//
// No fields are defined.
type ClosePeriodicEventStreamResponse struct {
}

func (s *ClosePeriodicEventStreamResponse) UnmarshalJSON(data []byte) error {
	type Alias ClosePeriodicEventStreamResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClosePeriodicEventStreamResponse(a)
	return s.Validate()
}

func (s ClosePeriodicEventStreamResponse) Validate() error {
	_ = s
	return nil
}

// CostUpdatedRequest (1.18.1)
//
// This contains the field definition of the CostUpdatedRequest PDU sent by the CSMS to the Charging Station.
// With this request the CSMS can send the current cost of a transaction to a Charging Station.
type CostUpdatedRequest struct {
	// Required. Current total cost, based on the information known by the CSMS, of the transaction including taxes.
	// In the currency configured with the configuration Variable: [Currency]
	TotalCost float64 `json:"totalCost"`
	// Required. Transaction Id of the transaction the current cost are asked for.
	TransactionID IdentifierString36Type `json:"transactionId"`
}

func (s *CostUpdatedRequest) UnmarshalJSON(data []byte) error {
	type Alias CostUpdatedRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CostUpdatedRequest(a)
	return s.Validate()
}

func (s CostUpdatedRequest) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	return nil
}

// CostUpdatedResponse (1.18.2)
//
// This contains the field definition of the CostUpdatedResponse PDU sent by the Charging Station to the CSMS in
// response to CostUpdatedRequest. No fields are defined.
//
// No fields are defined.
type CostUpdatedResponse struct {
}

func (s *CostUpdatedResponse) UnmarshalJSON(data []byte) error {
	type Alias CostUpdatedResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CostUpdatedResponse(a)
	return s.Validate()
}

func (s CostUpdatedResponse) Validate() error {
	_ = s
	return nil
}

// CustomerInformationRequest (1.19.1)
type CustomerInformationRequest struct {
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Required. Flag indicating whether the Charging Station should return NotifyCustomerInformationRequest messages
	// containing information about the customer referred to.
	Report bool `json:"report"`
	// Required. Flag indicating whether the Charging Station should clear all information about the customer
	// referred to.
	Clear bool `json:"clear"`
	// Optional. A (e.g. vendor specific) identifier of the customer this request refers to. This field contains a
	// custom identifier other than IdToken and Certificate. One of the possible identifiers (customerIdentifier,
	// customerIdToken or customerCertificate) should be in the request message.
	CustomerIdentifier *string `json:"customerIdentifier,omitempty"`
	// Optional. The IdToken of the customer this request refers to. One of the possible identifiers
	// (customerIdentifier, customerIdToken or customerCertificate) should be in the request message.
	IDToken *IDTokenType `json:"idToken,omitempty"`
	// Optional. The Certificate of the customer this request refers to. One of the possible identifiers
	// (customerIdentifier, customerIdToken or customerCertificate) should be in the request message.
	CustomerCertificate *CertificateHashDataType `json:"customerCertificate,omitempty"`
}

func (s *CustomerInformationRequest) UnmarshalJSON(data []byte) error {
	type Alias CustomerInformationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CustomerInformationRequest(a)
	return s.Validate()
}

func (s CustomerInformationRequest) Validate() error {
	if err := validateNonNegative(s.RequestID, "requestId"); err != nil {
		return err
	}

	if s.CustomerIdentifier != nil {
		if err := validateStringLen(*s.CustomerIdentifier, 64, "customerIdentifier"); err != nil {
			return err
		}
	}

	if s.IDToken != nil {
		if err := s.IDToken.Validate(); err != nil {
			return ocpp.WrapField("idToken", err)
		}
	}

	if s.CustomerCertificate != nil {
		if err := s.CustomerCertificate.Validate(); err != nil {
			return ocpp.WrapField("customerCertificate", err)
		}
	}

	return nil
}

// CustomerInformationResponse (1.19.2)
type CustomerInformationResponse struct {
	// Required. Indicates whether the request was accepted.
	Status CustomerInformationStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *CustomerInformationResponse) UnmarshalJSON(data []byte) error {
	type Alias CustomerInformationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CustomerInformationResponse(a)
	return s.Validate()
}

func (s CustomerInformationResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// DataTransferRequest (1.20.1)
//
// This contains the field definition of the DataTransferRequest PDU sent either by the CSMS to the Charging
// Station or vice versa.
type DataTransferRequest struct {
	// Optional. May be used to indicate a specific message or implementation.
	MessageID *string `json:"messageId,omitempty"`
	// Optional. Data without specified length or format. This needs to be decided by both parties (Open to
	// implementation).
	Data json.RawMessage `json:"data,omitempty"`
	// Required. This identifies the Vendor specific implementation
	VendorID string `json:"vendorId"`
}

func (s *DataTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias DataTransferRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DataTransferRequest(a)
	return s.Validate()
}

func (s DataTransferRequest) Validate() error {
	if s.MessageID != nil {
		if err := validateStringLen(*s.MessageID, 50, "messageId"); err != nil {
			return err
		}
	}

	if s.VendorID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "vendorId", "required field is missing")
	}

	if err := validateStringLen(s.VendorID, 255, "vendorId"); err != nil {
		return err
	}

	return nil
}

// DataTransferResponse (1.20.2)
//
// This contains the field definition of the DataTransferResponse PDU sent by the Charging Station to the CSMS or
// vice versa in response to a DataTransferRequest.
type DataTransferResponse struct {
	// Required. This indicates the success or failure of the data transfer.
	Status DataTransferStatusEnumType `json:"status"`
	// Optional. Data without specified length or format, in response to request.
	Data json.RawMessage `json:"data,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *DataTransferResponse) UnmarshalJSON(data []byte) error {
	type Alias DataTransferResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DataTransferResponse(a)
	return s.Validate()
}

func (s DataTransferResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// DeleteCertificateRequest (1.21.1)
//
// Used by the CSMS to request deletion of an installed certificate on a Charging Station.
type DeleteCertificateRequest struct {
	// Required. Indicates the certificate of which deletion is requested.
	CertificateHashData CertificateHashDataType `json:"certificateHashData"`
}

func (s *DeleteCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias DeleteCertificateRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DeleteCertificateRequest(a)
	return s.Validate()
}

func (s DeleteCertificateRequest) Validate() error {
	if err := s.CertificateHashData.Validate(); err != nil {
		return ocpp.WrapField("certificateHashData", err)
	}

	return nil
}

// DeleteCertificateResponse (1.21.2)
//
// Response to a DeleteCertificateRequest.
type DeleteCertificateResponse struct {
	// Required. Charging Station indicates if it can process the request.
	Status DeleteCertificateStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *DeleteCertificateResponse) UnmarshalJSON(data []byte) error {
	type Alias DeleteCertificateResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DeleteCertificateResponse(a)
	return s.Validate()
}

func (s DeleteCertificateResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// FirmwareStatusNotificationRequest (1.22.1)
//
// This contains the field definition of the FirmwareStatusNotificationRequest PDU sent by the Charging Station
// to the CSMS.
type FirmwareStatusNotificationRequest struct {
	// Required. This contains the progress status of the firmware installation.
	Status FirmwareStatusEnumType `json:"status"`
	// Optional. The request id that was provided in the UpdateFirmwareRequest that started this firmware update.
	// This field is mandatory, unless the message was triggered by a TriggerMessageRequest AND there is no firmware
	// update ongoing.
	RequestID *int32 `json:"requestId,omitempty"`
	// Optional. (2.1) Detailed status info
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *FirmwareStatusNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias FirmwareStatusNotificationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FirmwareStatusNotificationRequest(a)
	return s.Validate()
}

func (s FirmwareStatusNotificationRequest) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// FirmwareStatusNotificationResponse (1.22.2)
//
// This contains the field definition of the FirmwareStatusNotificationResponse PDU sent by the CSMS to the
// Charging Station in response to a FirmwareStatusNotificationRequest. No fields are defined.
//
// No fields are defined.
type FirmwareStatusNotificationResponse struct {
}

func (s *FirmwareStatusNotificationResponse) UnmarshalJSON(data []byte) error {
	type Alias FirmwareStatusNotificationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FirmwareStatusNotificationResponse(a)
	return s.Validate()
}

func (s FirmwareStatusNotificationResponse) Validate() error {
	_ = s
	return nil
}

// Get15118EVCertificateRequest (1.23.1)
//
// This message is sent by the Charging Station to the CSMS if an ISO 15118 vehicle selects the service
// Certificate installation. NOTE: This message is based on CertificateInstallationReq Res from ISO 15118 2.
type Get15118EVCertificateRequest struct {
	// Required. Schema version currently used for the 15118 session between EV and Charging Station. Needed for
	// parsing of the EXI stream by the CSMS.
	Iso15118SchemaVersion string `json:"iso15118SchemaVersion"`
	// Required. Defines whether certificate needs to be installed or updated.
	Action CertificateActionEnumType `json:"action"`
	// Required. (2.1) Raw CertificateInstallationReq request from EV, Base64 encoded. Extended to support ISO
	// 15118-20 certificates. The minimum supported length is 11000. If a longer exiRequest is supported, then the
	// supported length must be communicated in variable OCPPCommCtrlr.FieldLength[
	// "Get15118EVCertificateRequest.exiRequest" ].
	ExiRequest string `json:"exiRequest"`
	// Optional. (2.1) Absent during ISO 15118-2 session. Required during ISO 15118-20 session. Maximum number of
	// contracts that EV wants to install.
	MaximumContractCertificateChains *int32 `json:"maximumContractCertificateChains,omitempty"`
	// Optional. (2.1) Absent during ISO 15118-2 session. Optional during ISO 15118-20 session. List of EMAIDs for
	// which contract certificates must be requested first, in case there are more certificates than allowed by
	// maximumContractCertificateChains.
	PrioritizedEMAIDs []IdentifierString255Type `json:"prioritizedEMAIDs,omitempty"`
}

func (s *Get15118EVCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias Get15118EVCertificateRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = Get15118EVCertificateRequest(a)
	return s.Validate()
}

func (s Get15118EVCertificateRequest) Validate() error {
	if s.Iso15118SchemaVersion == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "iso15118SchemaVersion", "required field is missing")
	}

	if err := validateStringLen(s.Iso15118SchemaVersion, 50, "iso15118SchemaVersion"); err != nil {
		return err
	}

	if s.Action == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "action", "required field is missing")
	}

	if s.ExiRequest == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "exiRequest", "required field is missing")
	}

	if err := validateStringLen(s.ExiRequest, 11000, "exiRequest"); err != nil {
		return err
	}

	if s.MaximumContractCertificateChains != nil {
		if err := validateNonNegative(*s.MaximumContractCertificateChains, "maximumContractCertificateChains"); err != nil {
			return err
		}
	}

	if err := validateSliceLen(s.PrioritizedEMAIDs, 0, 8, "prioritizedEMAIDs"); err != nil {
		return err
	}

	for i, v := range s.PrioritizedEMAIDs {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("prioritizedEMAIDs[%d]", i), err)
		}
	}

	return nil
}

// Get15118EVCertificateResponse (1.23.2)
//
// Response message from CSMS to Charging Station containing the status and optionally new certificate. NOTE:
// This message is based on CertificateInstallationReq Res from ISO 15118-2.
type Get15118EVCertificateResponse struct {
	// Required. Indicates whether the message was processed properly.
	Status Iso15118EVCertificateStatusEnumType `json:"status"`
	// Required. (2/1) Raw CertificateInstallationRes response for the EV, Base64 encoded. Extended to support ISO
	// 15118-20 certificates. The minimum supported length is 17000. If a longer exiResponse is supported, then the
	// supported length must be communicated in variable OCPPCommCtrlr.FieldLength[
	// "Get15118EVCertificateResponse.exiResponse" ].
	ExiResponse string `json:"exiResponse"`
	// Optional. (2.1) Number of contracts that can be retrieved with additional requests.
	RemainingContracts *int32 `json:"remainingContracts,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *Get15118EVCertificateResponse) UnmarshalJSON(data []byte) error {
	type Alias Get15118EVCertificateResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = Get15118EVCertificateResponse(a)
	return s.Validate()
}

func (s Get15118EVCertificateResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.ExiResponse == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "exiResponse", "required field is missing")
	}

	if err := validateStringLen(s.ExiResponse, 17000, "exiResponse"); err != nil {
		return err
	}

	if s.RemainingContracts != nil {
		if err := validateNonNegative(*s.RemainingContracts, "remainingContracts"); err != nil {
			return err
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetBaseReportRequest (1.24.1)
//
// This contains the field definition of the GetBaseReportRequest PDU sent by the CSMS to the Charging Station.
type GetBaseReportRequest struct {
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Required. This field specifies the report base.
	ReportBase ReportBaseEnumType `json:"reportBase"`
}

func (s *GetBaseReportRequest) UnmarshalJSON(data []byte) error {
	type Alias GetBaseReportRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetBaseReportRequest(a)
	return s.Validate()
}

func (s GetBaseReportRequest) Validate() error {
	if s.ReportBase == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "reportBase", "required field is missing")
	}

	return nil
}

// GetBaseReportResponse (1.24.2)
//
// This contains the field definition of the GetBaseReportResponse PDU sent by the Charging Station to the CSMS.
type GetBaseReportResponse struct {
	// Required. This indicates whether the Charging Station is able to accept this request.
	Status GenericDeviceModelStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetBaseReportResponse) UnmarshalJSON(data []byte) error {
	type Alias GetBaseReportResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetBaseReportResponse(a)
	return s.Validate()
}

func (s GetBaseReportResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetCertificateChainStatusRequest (1.25.1)
type GetCertificateChainStatusRequest struct {
	// Required. Certificate to check revocation status for.
	CertificateStatusRequests []CertificateStatusRequestInfoType `json:"certificateStatusRequests"`
}

func (s *GetCertificateChainStatusRequest) UnmarshalJSON(data []byte) error {
	type Alias GetCertificateChainStatusRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCertificateChainStatusRequest(a)
	return s.Validate()
}

func (s GetCertificateChainStatusRequest) Validate() error {
	if err := validateSliceLen(s.CertificateStatusRequests, 1, 4, "certificateStatusRequests"); err != nil {
		return err
	}

	for i, v := range s.CertificateStatusRequests {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("certificateStatusRequests[%d]", i), err)
		}
	}

	return nil
}

// GetCertificateChainStatusResponse (1.25.2)
type GetCertificateChainStatusResponse struct {
	// Required. Status of the certificate revocation check.
	CertificateStatus []CertificateStatusType `json:"certificateStatus"`
}

func (s *GetCertificateChainStatusResponse) UnmarshalJSON(data []byte) error {
	type Alias GetCertificateChainStatusResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCertificateChainStatusResponse(a)
	return s.Validate()
}

func (s GetCertificateChainStatusResponse) Validate() error {
	if err := validateSliceLen(s.CertificateStatus, 1, 4, "certificateStatus"); err != nil {
		return err
	}

	for i, v := range s.CertificateStatus {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("certificateStatus[%d]", i), err)
		}
	}

	return nil
}

// GetCertificateStatusRequest (1.26.1)
//
// This contains the field definition of the GetCertificateStatusRequest PDU sent by the Charging Station to the
// CSMS.
type GetCertificateStatusRequest struct {
	// Required. Indicates the certificate of which the status is requested.
	OCSPRequestData OCSPRequestDataType `json:"ocspRequestData"`
}

func (s *GetCertificateStatusRequest) UnmarshalJSON(data []byte) error {
	type Alias GetCertificateStatusRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCertificateStatusRequest(a)
	return s.Validate()
}

func (s GetCertificateStatusRequest) Validate() error {
	if err := s.OCSPRequestData.Validate(); err != nil {
		return ocpp.WrapField("ocspRequestData", err)
	}

	return nil
}

// GetCertificateStatusResponse (1.26.2)
//
// This contains the field definition of the GetCertificateStatusResponse PDU sent by the CSMS to the Charging
// Station.
type GetCertificateStatusResponse struct {
	// Required. This indicates whether the charging station was able to retrieve the OCSP certificate status.
	Status GetCertificateStatusEnumType `json:"status"`
	// Optional. (2.1) OCSPResponse class as defined in IETF RFC 6960. DER encoded (as defined in IETF RFC 6960), and
	// then base64 encoded. MAY only be omitted when status is not Accepted. The minimum supported length is 18000.
	// If a longer ocspResult is supported, then the supported length must be communicated in variable
	// OCPPCommCtrlr.FieldLength[ "GetCertificateStatusResponse.ocspResult" ].
	OCSPResult *string `json:"ocspResult,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetCertificateStatusResponse) UnmarshalJSON(data []byte) error {
	type Alias GetCertificateStatusResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCertificateStatusResponse(a)
	return s.Validate()
}

func (s GetCertificateStatusResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.OCSPResult != nil {
		if err := validateStringLen(*s.OCSPResult, 18000, "ocspResult"); err != nil {
			return err
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetChargingProfilesRequest (1.27.1)
//
// The message GetChargingProfilesRequest can be used by the CSMS to request installed charging profiles from the
// Charging Station. The charging profiles will then be reported by the Charging Station via
// ReportChargingProfilesRequest messages.
type GetChargingProfilesRequest struct {
	// Required. Reference identification that is to be used by the Charging Station in the
	// ReportChargingProfilesRequest when provided.
	RequestID int32 `json:"requestId"`
	// Optional. For which EVSE installed charging profiles SHALL be reported. If 0, only charging profiles installed
	// on the Charging Station itself (the grid connection) SHALL be reported. If omitted, all installed charging
	// profiles SHALL be reported. Reported charging profiles SHALL match the criteria in field chargingProfile.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. Specifies the charging profile.
	ChargingProfile ChargingProfileCriterionType `json:"chargingProfile"`
}

func (s *GetChargingProfilesRequest) UnmarshalJSON(data []byte) error {
	type Alias GetChargingProfilesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetChargingProfilesRequest(a)
	return s.Validate()
}

func (s GetChargingProfilesRequest) Validate() error {
	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	if err := s.ChargingProfile.Validate(); err != nil {
		return ocpp.WrapField("chargingProfile", err)
	}

	return nil
}

// GetChargingProfilesResponse (1.27.2)
//
// This contains the field definition of the GetChargingProfilesResponse PDU sent by the Charging Station to the
// CSMS in response to a GetChargingProfilesRequest.
type GetChargingProfilesResponse struct {
	// Required. This indicates whether the Charging Station is able to process this request and will send
	// ReportChargingProfilesRequest messages.
	Status GetChargingProfileStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetChargingProfilesResponse) UnmarshalJSON(data []byte) error {
	type Alias GetChargingProfilesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetChargingProfilesResponse(a)
	return s.Validate()
}

func (s GetChargingProfilesResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetCompositeScheduleRequest (1.28.1)
//
// This contains the field definition of the GetCompositeScheduleRequest PDU sent by the CSMS to the Charging
// Station.
type GetCompositeScheduleRequest struct {
	// Required. Length of the requested schedule in seconds.
	Duration int32 `json:"duration"`
	// Optional. Can be used to force a power or current profile.
	ChargingRateUnit *ChargingRateUnitEnumType `json:"chargingRateUnit,omitempty"`
	// Required. The ID of the EVSE for which the schedule is requested. When evseid=0, the Charging Station will
	// calculate the expected consumption for the grid connection.
	EVSEID int32 `json:"evseId"`
}

func (s *GetCompositeScheduleRequest) UnmarshalJSON(data []byte) error {
	type Alias GetCompositeScheduleRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCompositeScheduleRequest(a)
	return s.Validate()
}

func (s GetCompositeScheduleRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	return nil
}

// GetCompositeScheduleResponse (1.28.2)
//
// This contains the field definition of the GetCompositeScheduleResponse PDU sent by the Charging Station to the
// CSMS in response to a GetCompositeScheduleRequest .
type GetCompositeScheduleResponse struct {
	// Required. The Charging Station will indicate if it was able to process the request
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
	// Optional. This field contains the calculated composite schedule. It may only be omitted when this message
	// contains status Rejected.
	Schedule *CompositeScheduleType `json:"schedule,omitempty"`
}

func (s *GetCompositeScheduleResponse) UnmarshalJSON(data []byte) error {
	type Alias GetCompositeScheduleResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetCompositeScheduleResponse(a)
	return s.Validate()
}

func (s GetCompositeScheduleResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return ocpp.WrapField("schedule", err)
		}
	}

	return nil
}

// GetDERControlRequest (1.29.1)
type GetDERControlRequest struct {
	// Required. RequestId to be used in ReportDERControlRequest.
	RequestID int32 `json:"requestId"`
	// Optional. True: get a default DER control. False: get a scheduled control.
	IsDefault *bool `json:"isDefault,omitempty"`
	// Optional. Type of control settings to retrieve. Not used when controlId is provided.
	ControlType *DERControlEnumType `json:"controlType,omitempty"`
	// Optional. Id of setting to get. When omitted all settings for controlType are retrieved.
	ControlID *IdentifierString36Type `json:"controlId,omitempty"`
}

func (s *GetDERControlRequest) UnmarshalJSON(data []byte) error {
	type Alias GetDERControlRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetDERControlRequest(a)
	return s.Validate()
}

func (s GetDERControlRequest) Validate() error {
	if s.ControlID != nil {
		if err := s.ControlID.Validate(); err != nil {
			return ocpp.WrapField("controlId", err)
		}
	}

	return nil
}

// GetDERControlResponse (1.29.2)
type GetDERControlResponse struct {
	// Required. Result of operation.
	Status DERControlStatusEnumType `json:"status"`
	// Optional. Detailed status info.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetDERControlResponse) UnmarshalJSON(data []byte) error {
	type Alias GetDERControlResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetDERControlResponse(a)
	return s.Validate()
}

func (s GetDERControlResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetDisplayMessagesRequest (1.30.1)
type GetDisplayMessagesRequest struct {
	// Optional. If provided the Charging Station shall return Display Messages of the given ids. This field SHALL
	// NOT contain more ids than set in NumberOfDisplayMessages.maxLimit
	ID []int32 `json:"id,omitempty"`
	// Required. The Id of this request.
	RequestID int32 `json:"requestId"`
	// Optional. If provided the Charging Station shall return Display Messages with the given priority only.
	Priority *MessagePriorityEnumType `json:"priority,omitempty"`
	// Optional. If provided the Charging Station shall return Display Messages with the given state only.
	State *MessageStateEnumType `json:"state,omitempty"`
}

func (s *GetDisplayMessagesRequest) UnmarshalJSON(data []byte) error {
	type Alias GetDisplayMessagesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetDisplayMessagesRequest(a)
	return s.Validate()
}

func (s GetDisplayMessagesRequest) Validate() error {
	for i, v := range s.ID {
		if err := validateNonNegative(v, fmt.Sprintf("id[%d]", i)); err != nil {
			return err
		}
	}

	return nil
}

// GetDisplayMessagesResponse (1.30.2)
type GetDisplayMessagesResponse struct {
	// Required. Indicates if the Charging Station has Display Messages that match the request criteria in the
	// GetDisplayMessagesRequest
	Status GetDisplayMessagesStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetDisplayMessagesResponse) UnmarshalJSON(data []byte) error {
	type Alias GetDisplayMessagesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetDisplayMessagesResponse(a)
	return s.Validate()
}

func (s GetDisplayMessagesResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetInstalledCertificateIdsRequest (1.31.1)
//
// Used by the CSMS to request an overview of the installed certificates on a Charging Station.
type GetInstalledCertificateIdsRequest struct {
	// Optional. Indicates the type of certificates requested. When omitted, all certificate types are requested.
	CertificateType []GetCertificateIdUseEnumType `json:"certificateType,omitempty"`
}

func (s *GetInstalledCertificateIdsRequest) UnmarshalJSON(data []byte) error {
	type Alias GetInstalledCertificateIdsRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetInstalledCertificateIdsRequest(a)
	return s.Validate()
}

func (s GetInstalledCertificateIdsRequest) Validate() error {
	_ = s
	return nil
}

// GetInstalledCertificateIdsResponse (1.31.2)
//
// Response to a GetInstalledCertificateIDsRequest.
type GetInstalledCertificateIdsResponse struct {
	// Required. Charging Station indicates if it can process the request.
	Status GetInstalledCertificateStatusEnumType `json:"status"`
	// Optional. The Charging Station includes the Certificate information for each available certificate.
	CertificateHashDataChain []CertificateHashDataChainType `json:"certificateHashDataChain,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetInstalledCertificateIdsResponse) UnmarshalJSON(data []byte) error {
	type Alias GetInstalledCertificateIdsResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetInstalledCertificateIdsResponse(a)
	return s.Validate()
}

func (s GetInstalledCertificateIdsResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	for i, v := range s.CertificateHashDataChain {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("certificateHashDataChain[%d]", i), err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetLocalListVersionRequest (1.32.1)
//
// This contains the field definition of the GetLocalListVersionRequest PDU sent by the CSMS to the Charging
// Station. No fields are defined.
//
// No fields are defined.
type GetLocalListVersionRequest struct {
}

func (s *GetLocalListVersionRequest) UnmarshalJSON(data []byte) error {
	type Alias GetLocalListVersionRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetLocalListVersionRequest(a)
	return s.Validate()
}

func (s GetLocalListVersionRequest) Validate() error {
	_ = s
	return nil
}

// GetLocalListVersionResponse (1.32.2)
//
// This contains the field definition of the GetLocalListVersionResponse PDU sent by the Charging Station to CSMS
// in response to a GetLocalListVersionRequest.
type GetLocalListVersionResponse struct {
	// Required. This contains the current version number of the local authorization list in the Charging Station.
	VersionNumber int32 `json:"versionNumber"`
}

func (s *GetLocalListVersionResponse) UnmarshalJSON(data []byte) error {
	type Alias GetLocalListVersionResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetLocalListVersionResponse(a)
	return s.Validate()
}

func (s GetLocalListVersionResponse) Validate() error {
	_ = s
	return nil
}

// GetLogRequest (1.33.1)
//
// This contains the field definition of the GetLogRequest PDU sent by the CSMS to the Charging Station.
type GetLogRequest struct {
	// Required. This contains the type of log file that the Charging Station should send.
	LogType LogEnumType `json:"logType"`
	// Required. The Id of this request
	RequestID int32 `json:"requestId"`
	// Optional. This specifies how many times the Charging Station must retry to upload the log before giving up. If
	// this field is not present, it is left to Charging Station to decide how many times it wants to retry. If the
	// value is 0, it means: no retries.
	Retries *int32 `json:"retries,omitempty"`
	// Optional. The interval in seconds after which a retry may be attempted. If this field is not present, it is
	// left to Charging Station to decide how long to wait between attempts.
	RetryInterval *int32 `json:"retryInterval,omitempty"`
	// Required. This field specifies the requested log and the location to which the log should be sent.
	Log LogParametersType `json:"log"`
}

func (s *GetLogRequest) UnmarshalJSON(data []byte) error {
	type Alias GetLogRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetLogRequest(a)
	return s.Validate()
}

func (s GetLogRequest) Validate() error {
	if s.LogType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "logType", "required field is missing")
	}

	if s.Retries != nil {
		if err := validateNonNegative(*s.Retries, "retries"); err != nil {
			return err
		}
	}

	if err := s.Log.Validate(); err != nil {
		return ocpp.WrapField("log", err)
	}

	return nil
}

// GetLogResponse (1.33.2)
//
// This contains the field definition of the GetLogResponse PDU sent by the Charging Station to the CSMS in
// response to a GetLogRequest.
type GetLogResponse struct {
	// Required. This field indicates whether the Charging Station was able to accept the request.
	Status LogStatusEnumType `json:"status"`
	// Optional. This contains the name of the log file that will be uploaded. This field is not present when no
	// logging information is available.
	Filename *string `json:"filename,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetLogResponse) UnmarshalJSON(data []byte) error {
	type Alias GetLogResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetLogResponse(a)
	return s.Validate()
}

func (s GetLogResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.Filename != nil {
		if err := validateStringLen(*s.Filename, 255, "filename"); err != nil {
			return err
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetMonitoringReportRequest (1.34.1)
//
// This contains the field definition of the GetMonitoringReportRequest PDU sent by the CSMS to the Charging
// Station.
type GetMonitoringReportRequest struct {
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Optional. This field contains criteria for components for which a monitoring report is requested
	MonitoringCriteria []MonitoringCriterionEnumType `json:"monitoringCriteria,omitempty"`
	// Optional. This field specifies the components and variables for which a monitoring report is requested.
	ComponentVariable []ComponentVariableType `json:"componentVariable,omitempty"`
}

func (s *GetMonitoringReportRequest) UnmarshalJSON(data []byte) error {
	type Alias GetMonitoringReportRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetMonitoringReportRequest(a)
	return s.Validate()
}

func (s GetMonitoringReportRequest) Validate() error {
	if err := validateSliceLen(s.MonitoringCriteria, 0, 3, "monitoringCriteria"); err != nil {
		return err
	}

	for i, v := range s.ComponentVariable {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("componentVariable[%d]", i), err)
		}
	}

	return nil
}

// GetMonitoringReportResponse (1.34.2)
//
// This contains the field definition of the GetMonitoringReportResponse PDU sent by the Charging Station to the
// CSMS.
type GetMonitoringReportResponse struct {
	// Required. This field indicates whether the Charging Station was able to accept the request.
	Status GenericDeviceModelStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetMonitoringReportResponse) UnmarshalJSON(data []byte) error {
	type Alias GetMonitoringReportResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetMonitoringReportResponse(a)
	return s.Validate()
}

func (s GetMonitoringReportResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetPeriodicEventStreamRequest (1.35.1)
//
// No fields are defined.
type GetPeriodicEventStreamRequest struct {
}

func (s *GetPeriodicEventStreamRequest) UnmarshalJSON(data []byte) error {
	type Alias GetPeriodicEventStreamRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetPeriodicEventStreamRequest(a)
	return s.Validate()
}

func (s GetPeriodicEventStreamRequest) Validate() error {
	_ = s
	return nil
}

// GetPeriodicEventStreamResponse (1.35.2)
type GetPeriodicEventStreamResponse struct {
	// Optional. List of constant part of streams
	ConstantStreamData []ConstantStreamDataType `json:"constantStreamData,omitempty"`
}

func (s *GetPeriodicEventStreamResponse) UnmarshalJSON(data []byte) error {
	type Alias GetPeriodicEventStreamResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetPeriodicEventStreamResponse(a)
	return s.Validate()
}

func (s GetPeriodicEventStreamResponse) Validate() error {
	for i, v := range s.ConstantStreamData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("constantStreamData[%d]", i), err)
		}
	}

	return nil
}

// GetReportRequest (1.36.1)
//
// This contains the field definition of the GetReportRequest PDU sent by the CSMS to the Charging Station.
type GetReportRequest struct {
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Optional. This field contains criteria for components for which a report is requested
	ComponentCriteria []ComponentCriterionEnumType `json:"componentCriteria,omitempty"`
	// Optional. This field specifies the components and variables for which a report is requested.
	ComponentVariable []ComponentVariableType `json:"componentVariable,omitempty"`
}

func (s *GetReportRequest) UnmarshalJSON(data []byte) error {
	type Alias GetReportRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetReportRequest(a)
	return s.Validate()
}

func (s GetReportRequest) Validate() error {
	if err := validateSliceLen(s.ComponentCriteria, 0, 4, "componentCriteria"); err != nil {
		return err
	}

	for i, v := range s.ComponentVariable {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("componentVariable[%d]", i), err)
		}
	}

	return nil
}

// GetReportResponse (1.36.2)
//
// The response to a GetReportRequest, sent by the Charging Station to the CSMS.
type GetReportResponse struct {
	// Required. This field indicates whether the Charging Station was able to accept the request.
	Status GenericDeviceModelStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetReportResponse) UnmarshalJSON(data []byte) error {
	type Alias GetReportResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetReportResponse(a)
	return s.Validate()
}

func (s GetReportResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetTariffsRequest (1.37.1)
type GetTariffsRequest struct {
	// Required. EVSE id to get tariff from. When evseId = 0, this gets tariffs from all EVSEs.
	EVSEID int32 `json:"evseId"`
}

func (s *GetTariffsRequest) UnmarshalJSON(data []byte) error {
	type Alias GetTariffsRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetTariffsRequest(a)
	return s.Validate()
}

func (s GetTariffsRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	return nil
}

// GetTariffsResponse (1.37.2)
type GetTariffsResponse struct {
	// Required. Status of operation
	Status TariffGetStatusEnumType `json:"status"`
	// Optional. Installed default and user-specific tariffs per EVSE
	TariffAssignments []TariffAssignmentType `json:"tariffAssignments,omitempty"`
	// Optional. Details status information
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *GetTariffsResponse) UnmarshalJSON(data []byte) error {
	type Alias GetTariffsResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetTariffsResponse(a)
	return s.Validate()
}

func (s GetTariffsResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	for i, v := range s.TariffAssignments {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("tariffAssignments[%d]", i), err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetTransactionStatusRequest (1.38.1)
//
// With this message, the CSMS can ask the Charging Station whether it has transaction-related messages waiting
// to be delivered to the CSMS. When a transactionId is provided, only messages for a specific transaction are
// asked for.
type GetTransactionStatusRequest struct {
	// Optional. The Id of the transaction for which the status is requested.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
}

func (s *GetTransactionStatusRequest) UnmarshalJSON(data []byte) error {
	type Alias GetTransactionStatusRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetTransactionStatusRequest(a)
	return s.Validate()
}

func (s GetTransactionStatusRequest) Validate() error {
	if s.TransactionID != nil {
		if err := s.TransactionID.Validate(); err != nil {
			return ocpp.WrapField("transactionId", err)
		}
	}

	return nil
}

// GetTransactionStatusResponse (1.38.2)
//
// The response to a GetTransactionStatusRequest, sent by the Charging Station to the CSMS.
type GetTransactionStatusResponse struct {
	// Optional. Whether the transaction is still ongoing.
	OngoingIndicator *bool `json:"ongoingIndicator,omitempty"`
	// Required. Whether there are still message to be delivered.
	MessagesInQueue bool `json:"messagesInQueue"`
}

func (s *GetTransactionStatusResponse) UnmarshalJSON(data []byte) error {
	type Alias GetTransactionStatusResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetTransactionStatusResponse(a)
	return s.Validate()
}

func (s GetTransactionStatusResponse) Validate() error {
	_ = s
	return nil
}

// GetVariablesRequest (1.39.1)
//
// This contains the field definition of the GetVariablesRequest PDU sent by the CSMS to the Charging Station.
type GetVariablesRequest struct {
	// Required. List of requested variables.
	GetVariableData []GetVariableDataType `json:"getVariableData"`
}

func (s *GetVariablesRequest) UnmarshalJSON(data []byte) error {
	type Alias GetVariablesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetVariablesRequest(a)
	return s.Validate()
}

func (s GetVariablesRequest) Validate() error {
	if err := validateSliceLen(s.GetVariableData, 1, -1, "getVariableData"); err != nil {
		return err
	}

	for i, v := range s.GetVariableData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("getVariableData[%d]", i), err)
		}
	}

	return nil
}

// GetVariablesResponse (1.39.2)
//
// This contains the field definition of the GetVariablesResponse PDU sent by the CSMS to the Charging Station in
// response to GetVariablesRequest.
type GetVariablesResponse struct {
	// Required. List of requested variables and their values.
	GetVariableResult []GetVariableResultType `json:"getVariableResult"`
}

func (s *GetVariablesResponse) UnmarshalJSON(data []byte) error {
	type Alias GetVariablesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetVariablesResponse(a)
	return s.Validate()
}

func (s GetVariablesResponse) Validate() error {
	if err := validateSliceLen(s.GetVariableResult, 1, -1, "getVariableResult"); err != nil {
		return err
	}

	for i, v := range s.GetVariableResult {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("getVariableResult[%d]", i), err)
		}
	}

	return nil
}

// HeartbeatRequest (1.40.1)
//
// This contains the field definition of the HeartbeatRequest PDU sent by the Charging Station to the CSMS. No
// fields are defined.
//
// No fields are defined.
type HeartbeatRequest struct {
}

func (s *HeartbeatRequest) UnmarshalJSON(data []byte) error {
	type Alias HeartbeatRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = HeartbeatRequest(a)
	return s.Validate()
}

func (s HeartbeatRequest) Validate() error {
	_ = s
	return nil
}

// HeartbeatResponse (1.40.2)
//
// This contains the field definition of the HeartbeatResponse PDU sent by the CSMS to the Charging Station in
// response to a HeartbeatRequest.
type HeartbeatResponse struct {
	// Required. Contains the current time of the CSMS.
	CurrentTime time.Time `json:"currentTime"`
}

func (s *HeartbeatResponse) UnmarshalJSON(data []byte) error {
	type Alias HeartbeatResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = HeartbeatResponse(a)
	return s.Validate()
}

func (s HeartbeatResponse) Validate() error {
	if s.CurrentTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currentTime", "required field is missing")
	}

	return nil
}

// InstallCertificateRequest (1.41.1)
//
// Used by the CSMS to request installation of a certificate on a Charging Station. Note: This message is not for
// installing a TLS client certificate in a charging station. The CertificateSignedRequest mechanism is used for
// that.
type InstallCertificateRequest struct {
	// Required. Indicates the certificate type that is sent.
	CertificateType InstallCertificateUseEnumType `json:"certificateType"`
	// Required. A PEM encoded X.509 certificate.
	Certificate string `json:"certificate"`
}

func (s *InstallCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias InstallCertificateRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = InstallCertificateRequest(a)
	return s.Validate()
}

func (s InstallCertificateRequest) Validate() error {
	if s.CertificateType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "certificateType", "required field is missing")
	}

	if s.Certificate == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "certificate", "required field is missing")
	}

	if err := validateStringLen(s.Certificate, 10000, "certificate"); err != nil {
		return err
	}

	return nil
}

// InstallCertificateResponse (1.41.2)
//
// The response to a InstallCertificateRequest, sent by the Charging Station to the CSMS.
type InstallCertificateResponse struct {
	// Required. Charging Station indicates if installation was successful.
	Status InstallCertificateStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *InstallCertificateResponse) UnmarshalJSON(data []byte) error {
	type Alias InstallCertificateResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = InstallCertificateResponse(a)
	return s.Validate()
}

func (s InstallCertificateResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// LogStatusNotificationRequest (1.42.1)
//
// This contains the field definition of the LogStatusNotificationRequest PDU sent by the Charging Station to the
// CSMS.
type LogStatusNotificationRequest struct {
	// Required. This contains the status of the log upload.
	Status UploadLogStatusEnumType `json:"status"`
	// Optional. The request id that was provided in GetLogRequest that started this log upload. This field is
	// mandatory, unless the message was triggered by a TriggerMessageRequest AND there is no log upload ongoing.
	RequestID *int32 `json:"requestId,omitempty"`
	// Optional. (2.1) Detailed status info
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *LogStatusNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias LogStatusNotificationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = LogStatusNotificationRequest(a)
	return s.Validate()
}

func (s LogStatusNotificationRequest) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// LogStatusNotificationResponse (1.42.2)
//
// This contains the field definition of the LogStatusNotificationResponse PDU sent by the CSMS to the Charging
// Station in response to LogStatusNotificationRequest. No fields are defined.
//
// No fields are defined.
type LogStatusNotificationResponse struct {
}

func (s *LogStatusNotificationResponse) UnmarshalJSON(data []byte) error {
	type Alias LogStatusNotificationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = LogStatusNotificationResponse(a)
	return s.Validate()
}

func (s LogStatusNotificationResponse) Validate() error {
	_ = s
	return nil
}

// MeterValuesRequest (1.43.1)
//
// This contains the field definition of the MeterValuesRequest PDU sent by the Charging Station to the CSMS.
// This message might be removed in a future version of OCPP. It will be replaced by Device Management Monitoring
// events.
type MeterValuesRequest struct {
	// Required. This contains a number (>0) designating an EVSE of the Charging Station. ‘0’ (zero) is used to
	// designate the main power meter.
	EVSEID int32 `json:"evseId"`
	// Required. The sampled meter values with timestamps. The following Configuration Variables are used to
	// configure which measurands are sent: - AlignedDataMeasurands - AlignedDataUpstreamMeasurands
	MeterValue []MeterValueType `json:"meterValue"`
}

func (s *MeterValuesRequest) UnmarshalJSON(data []byte) error {
	type Alias MeterValuesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MeterValuesRequest(a)
	return s.Validate()
}

func (s MeterValuesRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if err := validateSliceLen(s.MeterValue, 1, -1, "meterValue"); err != nil {
		return err
	}

	for i, v := range s.MeterValue {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("meterValue[%d]", i), err)
		}
	}

	return nil
}

// MeterValuesResponse (1.43.2)
//
// This contains the field definition of the MeterValuesResponse PDU sent by the CSMS to the Charging Station in
// response to a MeterValuesRequest PDU. This message might be removed in a future version of OCPP. It will be
// replaced by Device Management Monitoring events.
//
// No fields are defined.
//
// No fields are defined.
type MeterValuesResponse struct {
}

func (s *MeterValuesResponse) UnmarshalJSON(data []byte) error {
	type Alias MeterValuesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MeterValuesResponse(a)
	return s.Validate()
}

func (s MeterValuesResponse) Validate() error {
	_ = s
	return nil
}

// NotifyAllowedEnergyTransferRequest (1.44.1)
//
// (2.1) Message sent by CSMS to update the list of authorized energy services, e.g. to allow bidirectional
// charging for a charging session that is already in progress. One example is that the EV has already started a
// transaction in charging-only mode and meanwhile the CSMS has found that this EV is authorized by some
// secondary actor, such as an aggregating party, to use bidirectional charging. This message is then used to
// give the EV the opportunity to change energy service from charging-only to bidirectional charging.
//
// Another example is that the CSMS wishes to change the active energy service. This is done by updating the list
// of authorized energy services and omitting the currently active energy service.
type NotifyAllowedEnergyTransferRequest struct {
	// Required. The transaction for which the allowed energy transfer is allowed.
	TransactionID IdentifierString36Type `json:"transactionId"`
	// Required. Modes of energy transfer that are accepted by CSMS.
	AllowedEnergyTransfer []EnergyTransferModeEnumType `json:"allowedEnergyTransfer"`
}

func (s *NotifyAllowedEnergyTransferRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyAllowedEnergyTransferRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyAllowedEnergyTransferRequest(a)
	return s.Validate()
}

func (s NotifyAllowedEnergyTransferRequest) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	if err := validateSliceLen(s.AllowedEnergyTransfer, 1, -1, "allowedEnergyTransfer"); err != nil {
		return err
	}

	return nil
}

// NotifyAllowedEnergyTransferResponse (1.44.2)
//
// (2.1) Status of NotifyAllowedEnergyServicesRequest. Request should normally not be rejected, unless there are
// some technical problems.
type NotifyAllowedEnergyTransferResponse struct {
	// Required.
	Status NotifyAllowedEnergyTransferStatusEnumType `json:"status"`
	// Optional.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *NotifyAllowedEnergyTransferResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyAllowedEnergyTransferResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyAllowedEnergyTransferResponse(a)
	return s.Validate()
}

func (s NotifyAllowedEnergyTransferResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// NotifyChargingLimitRequest (1.45.1)
//
// The message NotifyChargingLimitRequest can be used to communicate a charging limit, set by an external system
// on the Charging Station (Not installed by the CSO via SetChargingProfileRequest), to the CSMS.
type NotifyChargingLimitRequest struct {
	// Optional. The EVSE to which the charging limit is set. If absent or when zero, it applies to the entire
	// Charging Station.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. This contains the source of the charging limit and whether it is grid critical.
	ChargingLimit ChargingLimitType `json:"chargingLimit"`
	// Optional. Contains limits for the available power or current over time, as set by the external source.
	ChargingSchedule []ChargingScheduleType `json:"chargingSchedule,omitempty"`
}

func (s *NotifyChargingLimitRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyChargingLimitRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyChargingLimitRequest(a)
	return s.Validate()
}

func (s NotifyChargingLimitRequest) Validate() error {
	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	if err := s.ChargingLimit.Validate(); err != nil {
		return ocpp.WrapField("chargingLimit", err)
	}

	for i, v := range s.ChargingSchedule {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingSchedule[%d]", i), err)
		}
	}

	return nil
}

// NotifyChargingLimitResponse (1.45.2)
//
// The NotifyChargingLimitResponse message is sent by the CSMS to the Charging Station in response to a
// NotifyChargingLimitsRequest. No fields are defined.
//
// No fields are defined.
type NotifyChargingLimitResponse struct {
}

func (s *NotifyChargingLimitResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyChargingLimitResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyChargingLimitResponse(a)
	return s.Validate()
}

func (s NotifyChargingLimitResponse) Validate() error {
	_ = s
	return nil
}

// NotifyCustomerInformationRequest (1.46.1)
type NotifyCustomerInformationRequest struct {
	// Required. (Part of) the requested data. No format specified in which the data is returned. Should be human
	// readable.
	Data string `json:"data"`
	// Optional. “to be continued” indicator. Indicates whether another part of the monitoringData follows in an
	// upcoming notifyMonitoringReportRequest message. Default value when omitted is false.
	Tbc *bool `json:"tbc,omitempty"`
	// Required. Sequence number of this message. First message starts at 0.
	SeqNo int32 `json:"seqNo"`
	// Required. Timestamp of the moment this message was generated at the Charging Station.
	GeneratedAt time.Time `json:"generatedAt"`
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
}

func (s *NotifyCustomerInformationRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyCustomerInformationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyCustomerInformationRequest(a)
	return s.Validate()
}

func (s NotifyCustomerInformationRequest) Validate() error {
	if s.Data == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "data", "required field is missing")
	}

	if err := validateStringLen(s.Data, 512, "data"); err != nil {
		return err
	}

	if err := validateNonNegative(s.SeqNo, "seqNo"); err != nil {
		return err
	}

	if s.GeneratedAt.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "generatedAt", "required field is missing")
	}

	if err := validateNonNegative(s.RequestID, "requestId"); err != nil {
		return err
	}

	return nil
}

// NotifyCustomerInformationResponse (1.46.2)
//
// No fields are defined.
type NotifyCustomerInformationResponse struct {
}

func (s *NotifyCustomerInformationResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyCustomerInformationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyCustomerInformationResponse(a)
	return s.Validate()
}

func (s NotifyCustomerInformationResponse) Validate() error {
	_ = s
	return nil
}

// NotifyDERAlarmRequest (1.47.1)
type NotifyDERAlarmRequest struct {
	// Required. Name of DER control, e.g. LFMustTrip
	ControlType DERControlEnumType `json:"controlType"`
	// Optional. Type of grid event that caused this
	GridEventFault *GridEventFaultEnumType `json:"gridEventFault,omitempty"`
	// Optional. True when error condition has ended. Absent or false when alarm has started.
	AlarmEnded *bool `json:"alarmEnded,omitempty"`
	// Required. Time of start or end of alarm.
	Timestamp time.Time `json:"timestamp"`
	// Optional. Optional info provided by EV.
	ExtraInfo *string `json:"extraInfo,omitempty"`
}

func (s *NotifyDERAlarmRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyDERAlarmRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyDERAlarmRequest(a)
	return s.Validate()
}

func (s NotifyDERAlarmRequest) Validate() error {
	if s.ControlType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "controlType", "required field is missing")
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.ExtraInfo != nil {
		if err := validateStringLen(*s.ExtraInfo, 200, "extraInfo"); err != nil {
			return err
		}
	}

	return nil
}

// NotifyDERAlarmResponse (1.47.2)
//
// This message has no fields.
//
// No fields are defined.
type NotifyDERAlarmResponse struct {
}

func (s *NotifyDERAlarmResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyDERAlarmResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyDERAlarmResponse(a)
	return s.Validate()
}

func (s NotifyDERAlarmResponse) Validate() error {
	_ = s
	return nil
}

// NotifyDERStartStopRequest (1.48.1)
type NotifyDERStartStopRequest struct {
	// Required. Id of the started or stopped DER control. Corresponds to the controlId of the SetDERControlRequest.
	ControlID IdentifierString36Type `json:"controlId"`
	// Required. True if DER control has started. False if it has ended.
	Started bool `json:"started"`
	// Required. Time of start or end of event.
	Timestamp time.Time `json:"timestamp"`
	// Optional. List of controlIds that are superseded as a result of this control starting.
	SupersededIds []IdentifierString36Type `json:"supersededIds,omitempty"`
}

func (s *NotifyDERStartStopRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyDERStartStopRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyDERStartStopRequest(a)
	return s.Validate()
}

func (s NotifyDERStartStopRequest) Validate() error {
	if s.ControlID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "controlId", "required field is missing")
	}

	if err := s.ControlID.Validate(); err != nil {
		return ocpp.WrapField("controlId", err)
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if err := validateSliceLen(s.SupersededIds, 0, 24, "supersededIds"); err != nil {
		return err
	}

	for i, v := range s.SupersededIds {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("supersededIds[%d]", i), err)
		}
	}

	return nil
}

// NotifyDERStartStopResponse (1.48.2)
//
// This message has no fields.
//
// No fields are defined.
type NotifyDERStartStopResponse struct {
}

func (s *NotifyDERStartStopResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyDERStartStopResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyDERStartStopResponse(a)
	return s.Validate()
}

func (s NotifyDERStartStopResponse) Validate() error {
	_ = s
	return nil
}

// NotifyDisplayMessagesRequest (1.49.1)
//
// This contains the field definition of the NotifyDisplayMessagesRequest PDU sent by the Charging Station to the
// CSMS.
type NotifyDisplayMessagesRequest struct {
	// Required. The id of the GetDisplayMessagesRequest that requested this message.
	RequestID int32 `json:"requestId"`
	// Optional. "to be continued" indicator. Indicates whether another part of the report follows in an upcoming
	// NotifyDisplayMessagesRequest message. Default value when omitted is false.
	Tbc *bool `json:"tbc,omitempty"`
	// Optional. The requested display message as configured in the Charging Station.
	MessageInfo []MessageInfoType `json:"messageInfo,omitempty"`
}

func (s *NotifyDisplayMessagesRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyDisplayMessagesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyDisplayMessagesRequest(a)
	return s.Validate()
}

func (s NotifyDisplayMessagesRequest) Validate() error {
	for i, v := range s.MessageInfo {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("messageInfo[%d]", i), err)
		}
	}

	return nil
}

// NotifyDisplayMessagesResponse (1.49.2)
//
// The NotifyDisplayMessagesResponse message is sent by the CSMS to the Charging Station in response to a
// NotifyDisplayMessagesRequest. No fields are defined.
//
// No fields are defined.
type NotifyDisplayMessagesResponse struct {
}

func (s *NotifyDisplayMessagesResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyDisplayMessagesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyDisplayMessagesResponse(a)
	return s.Validate()
}

func (s NotifyDisplayMessagesResponse) Validate() error {
	_ = s
	return nil
}

// NotifyEVChargingNeedsRequest (1.50.1)
//
// The Charging Station uses this message to communicate the charging needs as calculated by the EV to the CSMS.
type NotifyEVChargingNeedsRequest struct {
	// Required. Defines the EVSE and connector to which the EV is connected. EvseId may not be 0.
	EVSEID int32 `json:"evseId"`
	// Optional. Contains the maximum elements the EV supports for: - ISO 15118-2: schedule tuples in SASchedule
	// (both Pmax and Tariff). - ISO 15118-20: PowerScheduleEntry, PriceRule and PriceLevelScheduleEntries.
	MaxScheduleTuples *int32 `json:"maxScheduleTuples,omitempty"`
	// Optional. (2.1) Time when EV charging needs were received. Field can be added when charging station was
	// offline when charging needs were received.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Required. The characteristics of the energy delivery required.
	ChargingNeeds ChargingNeedsType `json:"chargingNeeds"`
}

func (s *NotifyEVChargingNeedsRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyEVChargingNeedsRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyEVChargingNeedsRequest(a)
	return s.Validate()
}

func (s NotifyEVChargingNeedsRequest) Validate() error {
	if s.EVSEID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "evseId", "must be >= 1")
	}

	if s.MaxScheduleTuples != nil {
		if err := validateNonNegative(*s.MaxScheduleTuples, "maxScheduleTuples"); err != nil {
			return err
		}
	}

	if err := s.ChargingNeeds.Validate(); err != nil {
		return ocpp.WrapField("chargingNeeds", err)
	}

	return nil
}

// NotifyEVChargingNeedsResponse (1.50.2)
//
// Response to a NotifyEVChargingNeedsRequest.
type NotifyEVChargingNeedsResponse struct {
	// Required. Returns whether the CSMS has been able to process the message successfully. It does not imply that
	// the evChargingNeeds can be met with the current charging profile.
	Status NotifyEVChargingNeedsStatusEnumType `json:"status"`
	// Optional.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *NotifyEVChargingNeedsResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyEVChargingNeedsResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyEVChargingNeedsResponse(a)
	return s.Validate()
}

func (s NotifyEVChargingNeedsResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// NotifyEVChargingScheduleRequest (1.51.1)
//
// The Charging Station uses this message to communicate the charging schedule as calculated by the EV to the
// CSMS.
type NotifyEVChargingScheduleRequest struct {
	// Required. Periods contained in the charging profile are relative to this point in time.
	TimeBase time.Time `json:"timeBase"`
	// Required. The charging schedule contained in this notification applies to an EVSE. EvseId must be > 0.
	EVSEID int32 `json:"evseId"`
	// Optional. (2.1) Id of the chargingSchedule that EV selected from the provided ChargingProfile.
	SelectedChargingScheduleID *int32 `json:"selectedChargingScheduleId,omitempty"`
	// Optional. (2.1) True when power tolerance is accepted by EV. This value is taken from
	// EVPowerProfile.PowerToleranceAcceptance in the ISO 15118-20 PowerDeliverReq message..
	PowerToleranceAcceptance *bool `json:"powerToleranceAcceptance,omitempty"`
	// Required. Planned energy consumption of the EV over time. Always relative to timeBase.
	ChargingSchedule ChargingScheduleType `json:"chargingSchedule"`
}

func (s *NotifyEVChargingScheduleRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyEVChargingScheduleRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyEVChargingScheduleRequest(a)
	return s.Validate()
}

func (s NotifyEVChargingScheduleRequest) Validate() error {
	if s.TimeBase.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timeBase", "required field is missing")
	}

	if s.EVSEID < 1 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "evseId", "must be >= 1")
	}

	if s.SelectedChargingScheduleID != nil {
		if err := validateNonNegative(*s.SelectedChargingScheduleID, "selectedChargingScheduleId"); err != nil {
			return err
		}
	}

	if err := s.ChargingSchedule.Validate(); err != nil {
		return ocpp.WrapField("chargingSchedule", err)
	}

	return nil
}

// NotifyEVChargingScheduleResponse (1.51.2)
//
// Response to a NotifyEVChargingScheduleRequest message.
type NotifyEVChargingScheduleResponse struct {
	// Required. Returns whether the CSMS has been able to process the message successfully. It does not imply any
	// approval of the charging schedule.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *NotifyEVChargingScheduleResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyEVChargingScheduleResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyEVChargingScheduleResponse(a)
	return s.Validate()
}

func (s NotifyEVChargingScheduleResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// NotifyEventRequest (1.52.1)
//
// This contains the field definition of the NotifyEventRequest PDU sent by the Charging Station to the CSMS.
type NotifyEventRequest struct {
	// Required. Timestamp of the moment this message was generated at the Charging Station.
	GeneratedAt time.Time `json:"generatedAt"`
	// Optional. “to be continued” indicator. Indicates whether another part of the report follows in an upcoming
	// notifyEventRequest message. Default value when omitted is false.
	Tbc *bool `json:"tbc,omitempty"`
	// Required. Sequence number of this message. First message starts at 0.
	SeqNo int32 `json:"seqNo"`
	// Required. List of EventData. An EventData element contains only the Component, Variable and VariableMonitoring
	// data that caused the event. The list of EventData will usally contain one eventData element, but the Charging
	// Station may decide to group multiple events in one notification. For example, when multiple events triggered
	// at the same time.
	EventData []EventDataType `json:"eventData"`
}

func (s *NotifyEventRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyEventRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyEventRequest(a)
	return s.Validate()
}

func (s NotifyEventRequest) Validate() error {
	if s.GeneratedAt.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "generatedAt", "required field is missing")
	}

	if err := validateNonNegative(s.SeqNo, "seqNo"); err != nil {
		return err
	}

	if err := validateSliceLen(s.EventData, 1, -1, "eventData"); err != nil {
		return err
	}

	for i, v := range s.EventData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("eventData[%d]", i), err)
		}
	}

	return nil
}

// NotifyEventResponse (1.52.2)
//
// Response to NotifyEventRequest. No fields are defined.
//
// No fields are defined.
type NotifyEventResponse struct {
}

func (s *NotifyEventResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyEventResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyEventResponse(a)
	return s.Validate()
}

func (s NotifyEventResponse) Validate() error {
	_ = s
	return nil
}

// NotifyMonitoringReportRequest (1.53.1)
//
// This contains the field definition of the NotifyMonitoringRequest PDU sent by the Charging Station to the
// CSMS.
type NotifyMonitoringReportRequest struct {
	// Required. The id of the GetMonitoringRequest that requested this report.
	RequestID int32 `json:"requestId"`
	// Optional. “to be continued” indicator. Indicates whether another part of the monitoringData follows in an
	// upcoming notifyMonitoringReportRequest message. Default value when omitted is false.
	Tbc *bool `json:"tbc,omitempty"`
	// Required. Sequence number of this message. First message starts at 0.
	SeqNo int32 `json:"seqNo"`
	// Required. Timestamp of the moment this message was generated at the Charging Station.
	GeneratedAt time.Time `json:"generatedAt"`
	// Optional. List of MonitoringData containing monitoring settings.
	Monitor []MonitoringDataType `json:"monitor,omitempty"`
}

func (s *NotifyMonitoringReportRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyMonitoringReportRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyMonitoringReportRequest(a)
	return s.Validate()
}

func (s NotifyMonitoringReportRequest) Validate() error {
	if err := validateNonNegative(s.SeqNo, "seqNo"); err != nil {
		return err
	}

	if s.GeneratedAt.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "generatedAt", "required field is missing")
	}

	for i, v := range s.Monitor {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("monitor[%d]", i), err)
		}
	}

	return nil
}

// NotifyMonitoringReportResponse (1.53.2)
//
// Response to a NotifyMonitoringRequest message. No fields are defined.
//
// No fields are defined.
type NotifyMonitoringReportResponse struct {
}

func (s *NotifyMonitoringReportResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyMonitoringReportResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyMonitoringReportResponse(a)
	return s.Validate()
}

func (s NotifyMonitoringReportResponse) Validate() error {
	_ = s
	return nil
}

// NotifyPeriodicEventStreamRequest (1.54)
//
// This is a message of messageType SEND. It does not have a response.
type NotifyPeriodicEventStreamRequest struct {
	// Required. Id of stream.
	ID int32 `json:"id"`
	// Required. Number of data elements still pending to be sent.
	Pending int32 `json:"pending"`
	// Required. Base timestamp to add to time offset of values.
	Basetime time.Time `json:"basetime"`
	// Required. Variable part of stream data
	Data []StreamDataElementType `json:"data"`
}

func (s *NotifyPeriodicEventStreamRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyPeriodicEventStreamRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyPeriodicEventStreamRequest(a)
	return s.Validate()
}

func (s NotifyPeriodicEventStreamRequest) Validate() error {
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if err := validateNonNegative(s.Pending, "pending"); err != nil {
		return err
	}

	if s.Basetime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "basetime", "required field is missing")
	}

	if err := validateSliceLen(s.Data, 1, -1, "data"); err != nil {
		return err
	}

	for i, v := range s.Data {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("data[%d]", i), err)
		}
	}

	return nil
}

// NotifyPriorityChargingRequest (1.55.1)
//
// (2.1) Message sent by Charging Station to notify CSMS that it has switched to the priority charging profile,
// that allows for the maximum possible current or power under current circumstances. Message contains a
// transactionId, because it only applies to the transaction in progress.
type NotifyPriorityChargingRequest struct {
	// Required. The transaction for which priority charging is requested.
	TransactionID IdentifierString36Type `json:"transactionId"`
	// Required. True if priority charging was activated. False if it has stopped using the priority charging
	// profile.
	Activated bool `json:"activated"`
}

func (s *NotifyPriorityChargingRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyPriorityChargingRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyPriorityChargingRequest(a)
	return s.Validate()
}

func (s NotifyPriorityChargingRequest) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	return nil
}

// NotifyPriorityChargingResponse (1.55.2)
//
// (2.1) This response message has an empty body.
//
// No fields are defined.
type NotifyPriorityChargingResponse struct {
}

func (s *NotifyPriorityChargingResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyPriorityChargingResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyPriorityChargingResponse(a)
	return s.Validate()
}

func (s NotifyPriorityChargingResponse) Validate() error {
	_ = s
	return nil
}

// NotifyReportRequest (1.56.1)
//
// This contains the field definition of the NotifyReportRequest PDU sent by the Charging Station to the CSMS.
type NotifyReportRequest struct {
	// Required. The id of the GetReportRequest or GetBaseReportRequest that requested this report
	RequestID int32 `json:"requestId"`
	// Required. Timestamp of the moment this message was generated at the Charging Station.
	GeneratedAt time.Time `json:"generatedAt"`
	// Optional. “to be continued” indicator. Indicates whether another part of the report follows in an upcoming
	// notifyReportRequest message. Default value when omitted is false.
	Tbc *bool `json:"tbc,omitempty"`
	// Required. Sequence number of this message. First message starts at 0.
	SeqNo int32 `json:"seqNo"`
	// Optional. List of ReportData.
	ReportData []ReportDataType `json:"reportData,omitempty"`
}

func (s *NotifyReportRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyReportRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyReportRequest(a)
	return s.Validate()
}

func (s NotifyReportRequest) Validate() error {
	if s.GeneratedAt.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "generatedAt", "required field is missing")
	}

	if err := validateNonNegative(s.SeqNo, "seqNo"); err != nil {
		return err
	}

	for i, v := range s.ReportData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("reportData[%d]", i), err)
		}
	}

	return nil
}

// NotifyReportResponse (1.56.2)
//
// Response to a NotifyReportRequest message. No fields are defined.
//
// No fields are defined.
type NotifyReportResponse struct {
}

func (s *NotifyReportResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyReportResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyReportResponse(a)
	return s.Validate()
}

func (s NotifyReportResponse) Validate() error {
	_ = s
	return nil
}

// NotifySettlementRequest (1.57.1)
type NotifySettlementRequest struct {
	// Optional. The transactionId that the settlement belongs to. Can be empty if the payment transaction is
	// canceled prior to the start of the OCPP transaction.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Required. The payment reference received from the payment terminal and is used as the value for idToken.
	PSPRef IdentifierString255Type `json:"pspRef"`
	// Required. The status of the settlement attempt.
	Status PaymentStatusEnumType `json:"status"`
	// Optional. Additional information from payment terminal/payment process.
	StatusInfo *string `json:"statusInfo,omitempty"`
	// Required. The amount that was settled, or attempted to be settled (in case of failure).
	SettlementAmount float64 `json:"settlementAmount"`
	// Required. The time when the settlement was done.
	SettlementTime time.Time `json:"settlementTime"`
	// Optional.
	ReceiptID *string `json:"receiptId,omitempty"`
	// Optional. The receipt URL, to be used if the receipt is generated by the payment terminal or the CS.
	ReceiptURL *string `json:"receiptUrl,omitempty"`
	// Optional. VAT number for a company receipt.
	VATNumber *string `json:"vatNumber,omitempty"`
	// Optional. Company address associated with VAT number.
	VATCompany *AddressType `json:"vatCompany,omitempty"`
}

func (s *NotifySettlementRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifySettlementRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifySettlementRequest(a)
	return s.Validate()
}

func (s NotifySettlementRequest) Validate() error {
	if s.TransactionID != nil {
		if err := s.TransactionID.Validate(); err != nil {
			return ocpp.WrapField("transactionId", err)
		}
	}

	if s.PSPRef == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "pspRef", "required field is missing")
	}

	if err := s.PSPRef.Validate(); err != nil {
		return ocpp.WrapField("pspRef", err)
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := validateStringLen(*s.StatusInfo, 500, "statusInfo"); err != nil {
			return err
		}
	}

	if s.SettlementTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "settlementTime", "required field is missing")
	}

	if s.ReceiptID != nil {
		if err := validateStringLen(*s.ReceiptID, 50, "receiptId"); err != nil {
			return err
		}
	}

	if s.ReceiptURL != nil {
		if err := validateStringLen(*s.ReceiptURL, 2000, "receiptUrl"); err != nil {
			return err
		}
	}

	if s.VATNumber != nil {
		if err := validateStringLen(*s.VATNumber, 20, "vatNumber"); err != nil {
			return err
		}
	}

	if s.VATCompany != nil {
		if err := s.VATCompany.Validate(); err != nil {
			return ocpp.WrapField("vatCompany", err)
		}
	}

	return nil
}

// NotifySettlementResponse (1.57.2)
type NotifySettlementResponse struct {
	// Optional. The receipt URL if receipt generated by CSMS. The Charging Station can QR encode it and show it to
	// the EV Driver.
	ReceiptURL *string `json:"receiptUrl,omitempty"`
	// Optional. The receipt id if the receipt is generated by CSMS.
	ReceiptID *string `json:"receiptId,omitempty"`
}

func (s *NotifySettlementResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifySettlementResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifySettlementResponse(a)
	return s.Validate()
}

func (s NotifySettlementResponse) Validate() error {
	if s.ReceiptURL != nil {
		if err := validateStringLen(*s.ReceiptURL, 2000, "receiptUrl"); err != nil {
			return err
		}
	}

	if s.ReceiptID != nil {
		if err := validateStringLen(*s.ReceiptID, 50, "receiptId"); err != nil {
			return err
		}
	}

	return nil
}

// NotifyWebPaymentStartedRequest (1.58.1)
type NotifyWebPaymentStartedRequest struct {
	// Required. EVSE id for which transaction is requested.
	EVSEID int32 `json:"evseId"`
	// Required. Timeout value in seconds after which no result of web payment process (e.g. QR code scanning) is to
	// be expected anymore.
	Timeout int32 `json:"timeout"`
}

func (s *NotifyWebPaymentStartedRequest) UnmarshalJSON(data []byte) error {
	type Alias NotifyWebPaymentStartedRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyWebPaymentStartedRequest(a)
	return s.Validate()
}

func (s NotifyWebPaymentStartedRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	return nil
}

// NotifyWebPaymentStartedResponse (1.58.2)
//
// No fields are defined.
type NotifyWebPaymentStartedResponse struct {
}

func (s *NotifyWebPaymentStartedResponse) UnmarshalJSON(data []byte) error {
	type Alias NotifyWebPaymentStartedResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NotifyWebPaymentStartedResponse(a)
	return s.Validate()
}

func (s NotifyWebPaymentStartedResponse) Validate() error {
	_ = s
	return nil
}

// OpenPeriodicEventStreamRequest (1.59.1)
type OpenPeriodicEventStreamRequest struct {
	// Required. Constant part of stream data
	ConstantStreamData ConstantStreamDataType `json:"constantStreamData"`
}

func (s *OpenPeriodicEventStreamRequest) UnmarshalJSON(data []byte) error {
	type Alias OpenPeriodicEventStreamRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = OpenPeriodicEventStreamRequest(a)
	return s.Validate()
}

func (s OpenPeriodicEventStreamRequest) Validate() error {
	if err := s.ConstantStreamData.Validate(); err != nil {
		return ocpp.WrapField("constantStreamData", err)
	}

	return nil
}

// OpenPeriodicEventStreamResponse (1.59.2)
type OpenPeriodicEventStreamResponse struct {
	// Required. Result of request.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status info
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *OpenPeriodicEventStreamResponse) UnmarshalJSON(data []byte) error {
	type Alias OpenPeriodicEventStreamResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = OpenPeriodicEventStreamResponse(a)
	return s.Validate()
}

func (s OpenPeriodicEventStreamResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// PublishFirmwareRequest (1.60.1)
//
// This contains the field definition of the PublishFirmwareRequest PDU sent by the CSMS to the Local Controller.
type PublishFirmwareRequest struct {
	// Required. This contains a string containing a URI pointing to a location from which to retrieve the firmware.
	Location string `json:"location"`
	// Optional. This specifies how many times Charging Station must retry to download the firmware before giving up.
	// If this field is not present, it is left to Charging Station to decide how many times it wants to retry. If
	// the value is 0, it means: no retries.
	Retries *int32 `json:"retries,omitempty"`
	// Required. The MD5 checksum over the entire firmware file as a hexadecimal string of length 32.
	Checksum IdentifierString32Type `json:"checksum"`
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Optional. The interval in seconds after which a retry may be attempted. If this field is not present, it is
	// left to Charging Station to decide how long to wait between attempts.
	RetryInterval *int32 `json:"retryInterval,omitempty"`
}

func (s *PublishFirmwareRequest) UnmarshalJSON(data []byte) error {
	type Alias PublishFirmwareRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PublishFirmwareRequest(a)
	return s.Validate()
}

func (s PublishFirmwareRequest) Validate() error {
	if s.Location == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "location", "required field is missing")
	}

	if err := validateStringLen(s.Location, 2000, "location"); err != nil {
		return err
	}

	if s.Retries != nil {
		if err := validateNonNegative(*s.Retries, "retries"); err != nil {
			return err
		}
	}

	if s.Checksum == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "checksum", "required field is missing")
	}

	if err := s.Checksum.Validate(); err != nil {
		return ocpp.WrapField("checksum", err)
	}

	if err := validateNonNegative(s.RequestID, "requestId"); err != nil {
		return err
	}

	if s.RetryInterval != nil {
		if err := validateNonNegative(*s.RetryInterval, "retryInterval"); err != nil {
			return err
		}
	}

	return nil
}

// PublishFirmwareResponse (1.60.2)
//
// This contains the field definition of the PublishFirmwareResponse PDU sent by the Local Controller to the CSMS
// in response to a PublishFirmwareRequest.
type PublishFirmwareResponse struct {
	// Required. Indicates whether the request was accepted.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *PublishFirmwareResponse) UnmarshalJSON(data []byte) error {
	type Alias PublishFirmwareResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PublishFirmwareResponse(a)
	return s.Validate()
}

func (s PublishFirmwareResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// PublishFirmwareStatusNotificationRequest (1.61.1)
//
// This contains the field definition of the PublishFirmwareStatusNotificationRequest PDU sent by the Charging
// Station to the CSMS.
type PublishFirmwareStatusNotificationRequest struct {
	// Required. This contains the progress status of the publishfirmware installation.
	Status PublishFirmwareStatusEnumType `json:"status"`
	// Optional. Required if status is Published. Can be multiple URI’s, if the Local Controller supports e.g. HTTP,
	// HTTPS, and FTP.
	Location []string `json:"location,omitempty"`
	// Optional. The request id that was provided in the PublishFirmwareRequest which triggered this action.
	RequestID *int32 `json:"requestId,omitempty"`
	// Optional. (2.1) Detailed status info
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *PublishFirmwareStatusNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias PublishFirmwareStatusNotificationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PublishFirmwareStatusNotificationRequest(a)
	return s.Validate()
}

func (s PublishFirmwareStatusNotificationRequest) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	for i, v := range s.Location {
		if err := validateStringLen(v, 2000, fmt.Sprintf("location[%d]", i)); err != nil {
			return err
		}
	}

	if s.RequestID != nil {
		if err := validateNonNegative(*s.RequestID, "requestId"); err != nil {
			return err
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// PublishFirmwareStatusNotificationResponse (1.61.2)
//
// This contains the field definition of the PublishFirmwareStatusNotificationResponse PDU sent by the CSMS to
// the Charging station in response to a PublishFirmwareStatusNotificationRequest.
//
// No fields are defined.
type PublishFirmwareStatusNotificationResponse struct {
}

func (s *PublishFirmwareStatusNotificationResponse) UnmarshalJSON(data []byte) error {
	type Alias PublishFirmwareStatusNotificationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PublishFirmwareStatusNotificationResponse(a)
	return s.Validate()
}

func (s PublishFirmwareStatusNotificationResponse) Validate() error {
	_ = s
	return nil
}

// PullDynamicScheduleUpdateRequest (1.62.1)
//
// (2.1) This message is sent by a Charging Station to request an update of setpoints and/or limits of the
// charging profile with given chargingProfileId.
type PullDynamicScheduleUpdateRequest struct {
	// Required. Id of charging profile to update.
	ChargingProfileID int32 `json:"chargingProfileId"`
}

func (s *PullDynamicScheduleUpdateRequest) UnmarshalJSON(data []byte) error {
	type Alias PullDynamicScheduleUpdateRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PullDynamicScheduleUpdateRequest(a)
	return s.Validate()
}

func (s PullDynamicScheduleUpdateRequest) Validate() error {
	_ = s
	return nil
}

// PullDynamicScheduleUpdateResponse (1.62.2)
//
// (2.1) If no data can be provided by CSMS, then the response will only contain status.
type PullDynamicScheduleUpdateResponse struct {
	// Required. Result of request.
	Status ChargingProfileStatusEnumType `json:"status"`
	// Optional. Updated charging schedule period values.
	ScheduleUpdate *ChargingScheduleUpdateType `json:"scheduleUpdate,omitempty"`
	// Optional. Additional info about status
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *PullDynamicScheduleUpdateResponse) UnmarshalJSON(data []byte) error {
	type Alias PullDynamicScheduleUpdateResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PullDynamicScheduleUpdateResponse(a)
	return s.Validate()
}

func (s PullDynamicScheduleUpdateResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.ScheduleUpdate != nil {
		if err := s.ScheduleUpdate.Validate(); err != nil {
			return ocpp.WrapField("scheduleUpdate", err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ReportChargingProfilesRequest (1.63.1)
//
// Reports charging profiles installed in the Charging Station, as requested via a GetChargingProfilesRequest
// message. The charging profile report can be split over multiple ReportChargingProfilesRequest messages, this
// can be because charging profiles for different charging sources need to be reported, or because there is just
// to much data for one message.
type ReportChargingProfilesRequest struct {
	// Required. Id used to match the GetChargingProfilesRequest message with the resulting
	// ReportChargingProfilesRequest messages. When the CSMS provided a requestId in the GetChargingProfilesRequest,
	// this field SHALL contain the same value.
	RequestID int32 `json:"requestId"`
	// Required. Source that has installed this charging profile. Values defined in Appendix as
	// ChargingLimitSourceEnumStringType.
	ChargingLimitSource string `json:"chargingLimitSource"`
	// Optional. To Be Continued. Default value when omitted: false. false indicates that there are no further
	// messages as part of this report.
	Tbc *bool `json:"tbc,omitempty"`
	// Required. The evse to which the charging profile applies. If evseId = 0, the message contains an overall limit
	// for the Charging Station.
	EVSEID int32 `json:"evseId"`
	// Required. The charging profile as configured in the Charging Station.
	ChargingProfile []ChargingProfileType `json:"chargingProfile"`
}

func (s *ReportChargingProfilesRequest) UnmarshalJSON(data []byte) error {
	type Alias ReportChargingProfilesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReportChargingProfilesRequest(a)
	return s.Validate()
}

func (s ReportChargingProfilesRequest) Validate() error {
	if s.ChargingLimitSource == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingLimitSource", "required field is missing")
	}

	if err := validateStringLen(s.ChargingLimitSource, 20, "chargingLimitSource"); err != nil {
		return err
	}

	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if err := validateSliceLen(s.ChargingProfile, 1, -1, "chargingProfile"); err != nil {
		return err
	}

	for i, v := range s.ChargingProfile {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingProfile[%d]", i), err)
		}
	}

	return nil
}

// ReportChargingProfilesResponse (1.63.2)
//
// The ReportChargingProfilesResponse message is sent by the CSMS to the Charging Station in response to a
// ReportChargingProfilesRequest. No fields are defined.
//
// No fields are defined.
type ReportChargingProfilesResponse struct {
}

func (s *ReportChargingProfilesResponse) UnmarshalJSON(data []byte) error {
	type Alias ReportChargingProfilesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReportChargingProfilesResponse(a)
	return s.Validate()
}

func (s ReportChargingProfilesResponse) Validate() error {
	_ = s
	return nil
}

// ReportDERControlRequest (1.64.1)
//
// Reports DER controls requested by a GetDERControlRequest message. The report may consist of more than one
// message.
type ReportDERControlRequest struct {
	// Required. RequestId from GetDERControlRequest.
	RequestID int32 `json:"requestId"`
	// Optional. To Be Continued. Default value when omitted: false. False indicates that there are no further
	// messages as part of this report.
	Tbc *bool `json:"tbc,omitempty"`
	// Optional. Fixed power factor setpoint when absorbing active power
	FixedPFAbsorb []FixedPFGetType `json:"fixedPFAbsorb,omitempty"`
	// Optional. Fixed power factor setpoint when injecting active power
	FixedPFInject []FixedPFGetType `json:"fixedPFInject,omitempty"`
	// Optional. Fixed reactive power setting
	FixedVar []FixedVarGetType `json:"fixedVar,omitempty"`
	// Optional. Limit maximum discharge as percentage of rated capability
	LimitMaxDischarge []LimitMaxDischargeGetType `json:"limitMaxDischarge,omitempty"`
	// Optional. Frequency-Watt parameterized mode
	FreqDroop []FreqDroopGetType `json:"freqDroop,omitempty"`
	// Optional. Enter service after trip parameters
	EnterService []EnterServiceGetType `json:"enterService,omitempty"`
	// Optional. Gradient settings
	Gradient []GradientGetType `json:"gradient,omitempty"`
	// Optional. Voltage/Frequency/Active/Reactive curve
	Curve []DERCurveGetType `json:"curve,omitempty"`
}

func (s *ReportDERControlRequest) UnmarshalJSON(data []byte) error {
	type Alias ReportDERControlRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReportDERControlRequest(a)
	return s.Validate()
}

func (s ReportDERControlRequest) Validate() error {
	if err := validateSliceLen(s.FixedPFAbsorb, 0, 24, "fixedPFAbsorb"); err != nil {
		return err
	}

	for i, v := range s.FixedPFAbsorb {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("fixedPFAbsorb[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.FixedPFInject, 0, 24, "fixedPFInject"); err != nil {
		return err
	}

	for i, v := range s.FixedPFInject {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("fixedPFInject[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.FixedVar, 0, 24, "fixedVar"); err != nil {
		return err
	}

	for i, v := range s.FixedVar {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("fixedVar[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.LimitMaxDischarge, 0, 24, "limitMaxDischarge"); err != nil {
		return err
	}

	for i, v := range s.LimitMaxDischarge {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("limitMaxDischarge[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.FreqDroop, 0, 24, "freqDroop"); err != nil {
		return err
	}

	for i, v := range s.FreqDroop {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("freqDroop[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.EnterService, 0, 24, "enterService"); err != nil {
		return err
	}

	for i, v := range s.EnterService {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("enterService[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.Gradient, 0, 24, "gradient"); err != nil {
		return err
	}

	for i, v := range s.Gradient {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("gradient[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.Curve, 0, 24, "curve"); err != nil {
		return err
	}

	for i, v := range s.Curve {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("curve[%d]", i), err)
		}
	}

	return nil
}

// ReportDERControlResponse (1.64.2)
//
// This is an empty message sent by CSMS in response to a ReportDERControlRequest message.
//
// No fields are defined.
type ReportDERControlResponse struct {
}

func (s *ReportDERControlResponse) UnmarshalJSON(data []byte) error {
	type Alias ReportDERControlResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReportDERControlResponse(a)
	return s.Validate()
}

func (s ReportDERControlResponse) Validate() error {
	_ = s
	return nil
}

// RequestBatterySwapRequest (1.65.1)
type RequestBatterySwapRequest struct {
	// Required. Request id to match with BatterySwapRequest.
	RequestID int32 `json:"requestId"`
	// Required. Id token of EV driver.
	IDToken IDTokenType `json:"idToken"`
}

func (s *RequestBatterySwapRequest) UnmarshalJSON(data []byte) error {
	type Alias RequestBatterySwapRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RequestBatterySwapRequest(a)
	return s.Validate()
}

func (s RequestBatterySwapRequest) Validate() error {
	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	return nil
}

// RequestBatterySwapResponse (1.65.2)
type RequestBatterySwapResponse struct {
	// Required. Accepted or rejected the request.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Additional info on status
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *RequestBatterySwapResponse) UnmarshalJSON(data []byte) error {
	type Alias RequestBatterySwapResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RequestBatterySwapResponse(a)
	return s.Validate()
}

func (s RequestBatterySwapResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// RequestStartTransactionRequest (1.66.1)
//
// This contains the field definitions of the RequestStartTransactionRequest PDU sent to Charging Station by
// CSMS.
type RequestStartTransactionRequest struct {
	// Optional. Number of the EVSE on which to start the transaction. EvseId SHALL be > 0
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. Id given by the server to this start request. The Charging Station will return this in the
	// TransactionEventRequest, letting the server know which transaction was started for this request. Use to start
	// a transaction.
	RemoteStartID int32 `json:"remoteStartId"`
	// Required. The identifier that the Charging Station must use to start a transaction.
	IDToken IDTokenType `json:"idToken"`
	// Optional. Charging Profile to be used by the Charging Station for the requested transaction.
	// ChargingProfilePurpose MUST be set to TxProfile
	ChargingProfile *ChargingProfileType `json:"chargingProfile,omitempty"`
	// Optional. The groupIdToken is only relevant when the transaction is to be started on an EVSE for which a
	// reservation for groupIdToken is active, and the configuration variable AuthorizeRemoteStart = false (otherwise
	// the AuthorizeResponse could return the groupIdToken).
	GroupIDToken *IDTokenType `json:"groupIdToken,omitempty"`
}

func (s *RequestStartTransactionRequest) UnmarshalJSON(data []byte) error {
	type Alias RequestStartTransactionRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RequestStartTransactionRequest(a)
	return s.Validate()
}

func (s RequestStartTransactionRequest) Validate() error {
	if s.EVSEID != nil {
		if *s.EVSEID < 1 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "evseId", "must be >= 1")
		}
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	if s.ChargingProfile != nil {
		if err := s.ChargingProfile.Validate(); err != nil {
			return ocpp.WrapField("chargingProfile", err)
		}
	}

	if s.GroupIDToken != nil {
		if err := s.GroupIDToken.Validate(); err != nil {
			return ocpp.WrapField("groupIdToken", err)
		}
	}

	return nil
}

// RequestStartTransactionResponse (1.66.2)
//
// This contains the field definitions of the RequestStartTransactionResponse PDU sent from Charging Station to
// CSMS.
type RequestStartTransactionResponse struct {
	// Required. Status indicating whether the Charging Station accepts the request to start a transaction.
	Status RequestStartStopStatusEnumType `json:"status"`
	// Optional. When the transaction was already started by the Charging Station before the
	// RequestStartTransactionRequest was received, for example: cable plugged in first. This contains the
	// transactionId of the already started transaction.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *RequestStartTransactionResponse) UnmarshalJSON(data []byte) error {
	type Alias RequestStartTransactionResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RequestStartTransactionResponse(a)
	return s.Validate()
}

func (s RequestStartTransactionResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.TransactionID != nil {
		if err := s.TransactionID.Validate(); err != nil {
			return ocpp.WrapField("transactionId", err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// RequestStopTransactionRequest (1.67.1)
//
// This contains the field definitions of the RequestStopTransactionRequest PDU sent to Charging Station by CSMS.
type RequestStopTransactionRequest struct {
	// Required. The identifier of the transaction which the Charging Station is requested to stop.
	TransactionID IdentifierString36Type `json:"transactionId"`
}

func (s *RequestStopTransactionRequest) UnmarshalJSON(data []byte) error {
	type Alias RequestStopTransactionRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RequestStopTransactionRequest(a)
	return s.Validate()
}

func (s RequestStopTransactionRequest) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	return nil
}

// RequestStopTransactionResponse (1.67.2)
//
// This contains the field definitions of the RequestStopTransactionResponse PDU sent from Charging Station to
// CSMS.
type RequestStopTransactionResponse struct {
	// Required. Status indicating whether Charging Station accepts the request to stop a transaction.
	Status RequestStartStopStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *RequestStopTransactionResponse) UnmarshalJSON(data []byte) error {
	type Alias RequestStopTransactionResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RequestStopTransactionResponse(a)
	return s.Validate()
}

func (s RequestStopTransactionResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ReservationStatusUpdateRequest (1.68.1)
//
// This contains the field definition of the ReservationStatusUpdateRequest PDU sent by the Charging Station to
// the CSMS.
type ReservationStatusUpdateRequest struct {
	// Required. The ID of the reservation.
	ReservationID int32 `json:"reservationId"`
	// Required. The updated reservation status.
	ReservationUpdateStatus ReservationUpdateStatusEnumType `json:"reservationUpdateStatus"`
}

func (s *ReservationStatusUpdateRequest) UnmarshalJSON(data []byte) error {
	type Alias ReservationStatusUpdateRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReservationStatusUpdateRequest(a)
	return s.Validate()
}

func (s ReservationStatusUpdateRequest) Validate() error {
	if err := validateNonNegative(s.ReservationID, "reservationId"); err != nil {
		return err
	}

	if s.ReservationUpdateStatus == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "reservationUpdateStatus", "required field is missing")
	}

	return nil
}

// ReservationStatusUpdateResponse (1.68.2)
//
// This contains the field definition of the ReservationStatusUpdateResponse PDU sent by the CSMS to the Charging
// Station in response to a ReservationStatusUpdateRequest. No fields are defined.
//
// No fields are defined.
type ReservationStatusUpdateResponse struct {
}

func (s *ReservationStatusUpdateResponse) UnmarshalJSON(data []byte) error {
	type Alias ReservationStatusUpdateResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReservationStatusUpdateResponse(a)
	return s.Validate()
}

func (s ReservationStatusUpdateResponse) Validate() error {
	_ = s
	return nil
}

// ReserveNowRequest (1.69.1)
//
// This contains the field definition of the ReserveNowRequest PDU sent by the CSMS to the Charging Station.
type ReserveNowRequest struct {
	// Required. Id of reservation.
	ID int32 `json:"id"`
	// Required. Date and time at which the reservation expires.
	ExpiryDateTime time.Time `json:"expiryDateTime"`
	// Optional. This field specifies the connector type. Values defined in Appendix as ConnectorEnumStringType.
	ConnectorType *string `json:"connectorType,omitempty"`
	// Optional. This contains ID of the evse to be reserved.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. The identifier for which the reservation is made.
	IDToken IDTokenType `json:"idToken"`
	// Optional. The group identifier for which the reservation is made.
	GroupIDToken *IDTokenType `json:"groupIdToken,omitempty"`
}

func (s *ReserveNowRequest) UnmarshalJSON(data []byte) error {
	type Alias ReserveNowRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReserveNowRequest(a)
	return s.Validate()
}

func (s ReserveNowRequest) Validate() error {
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if s.ExpiryDateTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "expiryDateTime", "required field is missing")
	}

	if s.ConnectorType != nil {
		if err := validateStringLen(*s.ConnectorType, 20, "connectorType"); err != nil {
			return err
		}
	}

	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	if s.GroupIDToken != nil {
		if err := s.GroupIDToken.Validate(); err != nil {
			return ocpp.WrapField("groupIdToken", err)
		}
	}

	return nil
}

// ReserveNowResponse (1.69.2)
//
// This contains the field definition of the ReserveNowResponse PDU sent by the Charging Station to the CSMS in
// response to ReserveNowRequest PDU.
type ReserveNowResponse struct {
	// Required. This indicates the success or failure of the reservation.
	Status ReserveNowStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ReserveNowResponse) UnmarshalJSON(data []byte) error {
	type Alias ReserveNowResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReserveNowResponse(a)
	return s.Validate()
}

func (s ReserveNowResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ResetRequest (1.70.1)
//
// This contains the field definition of the ResetRequest PDU sent by the CSMS to the Charging Station.
type ResetRequest struct {
	// Required. This contains the type of reset that the Charging Station or EVSE should perform.
	Type ResetEnumType `json:"type"`
	// Optional. This contains the ID of a specific EVSE that needs to be reset, instead of the entire Charging
	// Station.
	EVSEID *int32 `json:"evseId,omitempty"`
}

func (s *ResetRequest) UnmarshalJSON(data []byte) error {
	type Alias ResetRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ResetRequest(a)
	return s.Validate()
}

func (s ResetRequest) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	return nil
}

// ResetResponse (1.70.2)
//
// This contains the field definition of the ResetResponse PDU sent by the Charging Station to the CSMS in
// response to ResetRequest.
type ResetResponse struct {
	// Required. This indicates whether the Charging Station is able to perform the reset.
	Status ResetStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ResetResponse) UnmarshalJSON(data []byte) error {
	type Alias ResetResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ResetResponse(a)
	return s.Validate()
}

func (s ResetResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SecurityEventNotificationRequest (1.71.1)
//
// Sent by the Charging Station to the CSMS in case of a security event.
type SecurityEventNotificationRequest struct {
	// Required. Type of the security event. This value should be taken from the Security events list.
	Type string `json:"type"`
	// Required. Date and time at which the event occurred.
	Timestamp time.Time `json:"timestamp"`
	// Optional. Additional information about the occurred security event.
	TechInfo *string `json:"techInfo,omitempty"`
}

func (s *SecurityEventNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias SecurityEventNotificationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SecurityEventNotificationRequest(a)
	return s.Validate()
}

func (s SecurityEventNotificationRequest) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateStringLen(s.Type, 50, "type"); err != nil {
		return err
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.TechInfo != nil {
		if err := validateStringLen(*s.TechInfo, 255, "techInfo"); err != nil {
			return err
		}
	}

	return nil
}

// SecurityEventNotificationResponse (1.71.2)
//
// Sent by the CSMS to the Charging Station to confirm the receipt of a SecurityEventNotificationRequest message.
// No fields are defined.
//
// No fields are defined.
type SecurityEventNotificationResponse struct {
}

func (s *SecurityEventNotificationResponse) UnmarshalJSON(data []byte) error {
	type Alias SecurityEventNotificationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SecurityEventNotificationResponse(a)
	return s.Validate()
}

func (s SecurityEventNotificationResponse) Validate() error {
	_ = s
	return nil
}

// SendLocalListRequest (1.72.1)
//
// This contains the field definition of the SendLocalListRequest PDU sent by the CSMS to the Charging Station.
// If no (empty) localAuthorizationList is given and the updateType is Full, all IdTokens are removed from the
// list. Requesting a Differential update without or with empty localAuthorizationList will have no effect on the
// list. All IdTokens in the localAuthorizationList MUST be unique, no duplicate values are allowed.
type SendLocalListRequest struct {
	// Required. In case of a full update this is the version number of the full list. In case of a differential
	// update it is the version number of the list after the update has been applied.
	VersionNumber int32 `json:"versionNumber"`
	// Required. This contains the type of update (full or differential) of this request.
	UpdateType UpdateEnumType `json:"updateType"`
	// Optional. This contains the Local Authorization List entries.
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty"`
}

func (s *SendLocalListRequest) UnmarshalJSON(data []byte) error {
	type Alias SendLocalListRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SendLocalListRequest(a)
	return s.Validate()
}

func (s SendLocalListRequest) Validate() error {
	if s.UpdateType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "updateType", "required field is missing")
	}

	for i, v := range s.LocalAuthorizationList {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("localAuthorizationList[%d]", i), err)
		}
	}

	return nil
}

// SendLocalListResponse (1.72.2)
//
// This contains the field definition of the SendLocalListResponse PDU sent by the Charging Station to the CSMS
// in response to SendLocalListRequest PDU.
type SendLocalListResponse struct {
	// Required. This indicates whether the Charging Station has successfully received and applied the update of the
	// Local Authorization List.
	Status SendLocalListStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SendLocalListResponse) UnmarshalJSON(data []byte) error {
	type Alias SendLocalListResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SendLocalListResponse(a)
	return s.Validate()
}

func (s SendLocalListResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetChargingProfileRequest (1.73.1)
//
// This contains the field definition of the SetChargingProfileRequest PDU sent by the CSMS to the Charging
// Station. The CSMS uses this message to send charging profiles to a Charging Station.
type SetChargingProfileRequest struct {
	// Required. For TxDefaultProfile an evseId=0 applies the profile to each individual evse. For
	// ChargingStationMaxProfile and ChargingStationExternalConstraints an evseId=0 contains an overal limit for the
	// whole Charging Station.
	EVSEID int32 `json:"evseId"`
	// Required. The charging profile to be set at the Charging Station.
	ChargingProfile ChargingProfileType `json:"chargingProfile"`
}

func (s *SetChargingProfileRequest) UnmarshalJSON(data []byte) error {
	type Alias SetChargingProfileRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetChargingProfileRequest(a)
	return s.Validate()
}

func (s SetChargingProfileRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if err := s.ChargingProfile.Validate(); err != nil {
		return ocpp.WrapField("chargingProfile", err)
	}

	return nil
}

// SetChargingProfileResponse (1.73.2)
//
// This contains the field definition of the SetChargingProfileResponse PDU sent by the Charging Station to the
// CSMS in response to SetChargingProfileRequest PDU.
type SetChargingProfileResponse struct {
	// Required. Returns whether the Charging Station has been able to process the message successfully. This does
	// not guarantee the schedule will be followed to the letter. There might be other constraints the Charging
	// Station may need to take into account.
	Status ChargingProfileStatusEnumType `json:"status"`
	// Optional.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetChargingProfileResponse) UnmarshalJSON(data []byte) error {
	type Alias SetChargingProfileResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetChargingProfileResponse(a)
	return s.Validate()
}

func (s SetChargingProfileResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetDefaultTariffRequest (1.74.1)
type SetDefaultTariffRequest struct {
	// Required. EVSE that tariff applies to. When evseId = 0, then tarriff applies to all EVSEs.
	EVSEID int32 `json:"evseId"`
	// Required. Tariff structure.
	Tariff TariffType `json:"tariff"`
}

func (s *SetDefaultTariffRequest) UnmarshalJSON(data []byte) error {
	type Alias SetDefaultTariffRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetDefaultTariffRequest(a)
	return s.Validate()
}

func (s SetDefaultTariffRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if err := s.Tariff.Validate(); err != nil {
		return ocpp.WrapField("tariff", err)
	}

	return nil
}

// SetDefaultTariffResponse (1.74.2)
type SetDefaultTariffResponse struct {
	// Required.
	Status TariffSetStatusEnumType `json:"status"`
	// Optional. Detailed info on status
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetDefaultTariffResponse) UnmarshalJSON(data []byte) error {
	type Alias SetDefaultTariffResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetDefaultTariffResponse(a)
	return s.Validate()
}

func (s SetDefaultTariffResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetDERControlRequest (1.75.1)
type SetDERControlRequest struct {
	// Required. True if this is a default DER control
	IsDefault bool `json:"isDefault"`
	// Required. Unique id of this control, e.g. UUID
	ControlID IdentifierString36Type `json:"controlId"`
	// Required. Type of control. Determines which setting field below is used.
	ControlType DERControlEnumType `json:"controlType"`
	// Optional. Voltage/Frequency/Active/Reactive curve
	Curve *DERCurveType `json:"curve,omitempty"`
	// Optional. Fixed power factor setpoint when absorbing active power
	FixedPFAbsorb *FixedPFType `json:"fixedPFAbsorb,omitempty"`
	// Optional. Fixed power factor setpoint when injecting active power
	FixedPFInject *FixedPFType `json:"fixedPFInject,omitempty"`
	// Optional. Fixed reactive power
	FixedVar *FixedVarType `json:"fixedVar,omitempty"`
	// Optional. Limit maximum discharge as percentage of rated capability
	LimitMaxDischarge *LimitMaxDischargeType `json:"limitMaxDischarge,omitempty"`
	// Optional. Frequency-Watt parameterized mode
	FreqDroop *FreqDroopType `json:"freqDroop,omitempty"`
	// Optional. Enter service after trip parameters (default control only)
	EnterService *EnterServiceType `json:"enterService,omitempty"`
	// Optional. Gradient (default ramp rate) settings (default control only)
	Gradient *GradientType `json:"gradient,omitempty"`
}

func (s *SetDERControlRequest) UnmarshalJSON(data []byte) error {
	type Alias SetDERControlRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetDERControlRequest(a)
	return s.Validate()
}

func (s SetDERControlRequest) Validate() error {
	if s.ControlID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "controlId", "required field is missing")
	}

	if err := s.ControlID.Validate(); err != nil {
		return ocpp.WrapField("controlId", err)
	}

	if s.ControlType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "controlType", "required field is missing")
	}

	if s.Curve != nil {
		if err := s.Curve.Validate(); err != nil {
			return ocpp.WrapField("curve", err)
		}
	}

	if s.FixedPFAbsorb != nil {
		if err := s.FixedPFAbsorb.Validate(); err != nil {
			return ocpp.WrapField("fixedPFAbsorb", err)
		}
	}

	if s.FixedPFInject != nil {
		if err := s.FixedPFInject.Validate(); err != nil {
			return ocpp.WrapField("fixedPFInject", err)
		}
	}

	if s.FixedVar != nil {
		if err := s.FixedVar.Validate(); err != nil {
			return ocpp.WrapField("fixedVar", err)
		}
	}

	if s.LimitMaxDischarge != nil {
		if err := s.LimitMaxDischarge.Validate(); err != nil {
			return ocpp.WrapField("limitMaxDischarge", err)
		}
	}

	if s.FreqDroop != nil {
		if err := s.FreqDroop.Validate(); err != nil {
			return ocpp.WrapField("freqDroop", err)
		}
	}

	if s.EnterService != nil {
		if err := s.EnterService.Validate(); err != nil {
			return ocpp.WrapField("enterService", err)
		}
	}

	if s.Gradient != nil {
		if err := s.Gradient.Validate(); err != nil {
			return ocpp.WrapField("gradient", err)
		}
	}

	return nil
}

// SetDERControlResponse (1.75.2)
type SetDERControlResponse struct {
	// Required. Result of operation.
	Status DERControlStatusEnumType `json:"status"`
	// Optional. List of controlIds that are superseded as a result of setting this control.
	SupersededIds []IdentifierString36Type `json:"supersededIds,omitempty"`
	// Optional. Additional details on status
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetDERControlResponse) UnmarshalJSON(data []byte) error {
	type Alias SetDERControlResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetDERControlResponse(a)
	return s.Validate()
}

func (s SetDERControlResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if err := validateSliceLen(s.SupersededIds, 0, 24, "supersededIds"); err != nil {
		return err
	}

	for i, v := range s.SupersededIds {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("supersededIds[%d]", i), err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetDisplayMessageRequest (1.76.1)
//
// This contains the field definition of the SetDisplayMessageRequest PDU sent by the CSMS to the Charging
// Station. The CSMS asks the Charging Station to configure a new display message that the Charging Station will
// display (in the future). See also O01 - Set Display Message, O02 - Set Display Message for Transaction and O06
// - Replace Display Message
type SetDisplayMessageRequest struct {
	// Required. Message to be configured in the Charging Station, to be displayed.
	Message MessageInfoType `json:"message"`
}

func (s *SetDisplayMessageRequest) UnmarshalJSON(data []byte) error {
	type Alias SetDisplayMessageRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetDisplayMessageRequest(a)
	return s.Validate()
}

func (s SetDisplayMessageRequest) Validate() error {
	if err := s.Message.Validate(); err != nil {
		return ocpp.WrapField("message", err)
	}

	return nil
}

// SetDisplayMessageResponse (1.76.2)
//
// This contains the field definition of the SetDisplayMessageResponse PDU sent by the Charging Station to the
// CSMS in a response to a SetDisplayMessageRequest. See also O01 - Set Display Message, O02 - Set Display
// Message for Transaction and O06 - Replace Display Message
type SetDisplayMessageResponse struct {
	// Required. This indicates whether the Charging Station is able to display the message.
	Status DisplayMessageStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetDisplayMessageResponse) UnmarshalJSON(data []byte) error {
	type Alias SetDisplayMessageResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetDisplayMessageResponse(a)
	return s.Validate()
}

func (s SetDisplayMessageResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetMonitoringBaseRequest (1.77.1)
//
// This contains the field definition of the SetMonitoringBaseRequest PDU sent by the CSMS to the Charging
// Station.
type SetMonitoringBaseRequest struct {
	// Required. Specify which monitoring base will be set
	MonitoringBase MonitoringBaseEnumType `json:"monitoringBase"`
}

func (s *SetMonitoringBaseRequest) UnmarshalJSON(data []byte) error {
	type Alias SetMonitoringBaseRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetMonitoringBaseRequest(a)
	return s.Validate()
}

func (s SetMonitoringBaseRequest) Validate() error {
	if s.MonitoringBase == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "monitoringBase", "required field is missing")
	}

	return nil
}

// SetMonitoringBaseResponse (1.77.2)
//
// This contains the field definition of the SetMonitoringBaseResponse PDU sent by the Charging Station to the
// CSMS in response to a SetMonitoringBaseRequest.
type SetMonitoringBaseResponse struct {
	// Required. Indicates whether the Charging Station was able to accept the request.
	Status GenericDeviceModelStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetMonitoringBaseResponse) UnmarshalJSON(data []byte) error {
	type Alias SetMonitoringBaseResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetMonitoringBaseResponse(a)
	return s.Validate()
}

func (s SetMonitoringBaseResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetMonitoringLevelRequest (1.78.1)
//
// This contains the field definition of the SetMonitoringLevelRequest PDU sent by the CSMS to the Charging
// Station.
type SetMonitoringLevelRequest struct {
	// Required. The Charging Station SHALL only report events with a severity number lower than or equal to this
	// severity. The severity range is 0-9, with 0 as the highest and 9 as the lowest severity level. The severity
	// levels have the following meaning: 0-Danger Indicates lives are potentially in danger. Urgent attention is
	// needed and action should be taken immediately. 1-Hardware Failure Indicates that the Charging Station is
	// unable to continue regular operations due to Hardware issues. Action is required. 2-System Failure Indicates
	// that the Charging Station is unable to continue regular operations due to software or minor hardware issues.
	// Action is required. 3-Critical Indicates a critical error. Action is required. 4-Error Indicates a non-urgent
	// error. Action is required. 5-Alert Indicates an alert event. Default severity for any type of monitoring
	// event. 6-Warning Indicates a warning event. Action may be required. 7-Notice Indicates an unusual event. No
	// immediate action is required. 8-Informational Indicates a regular operational event. May be used for
	// reporting, measuring throughput, etc. No action is required. 9-Debug Indicates information useful to
	// developers for debugging, not useful during operations.
	Severity int32 `json:"severity"`
}

func (s *SetMonitoringLevelRequest) UnmarshalJSON(data []byte) error {
	type Alias SetMonitoringLevelRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetMonitoringLevelRequest(a)
	return s.Validate()
}

func (s SetMonitoringLevelRequest) Validate() error {
	if err := validateNonNegative(s.Severity, "severity"); err != nil {
		return err
	}

	return nil
}

// SetMonitoringLevelResponse (1.78.2)
//
// This contains the field definition of the SetMonitoringLevelResponse PDU sent by the Charging Station to the
// CSMS in response to a SetMonitoringLevelRequest.
type SetMonitoringLevelResponse struct {
	// Required. Indicates whether the Charging Station was able to accept the request.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetMonitoringLevelResponse) UnmarshalJSON(data []byte) error {
	type Alias SetMonitoringLevelResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetMonitoringLevelResponse(a)
	return s.Validate()
}

func (s SetMonitoringLevelResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetNetworkProfileRequest (1.79.1)
//
// With this message the CSMS gains the ability to configure the connection data (e.g. CSMS URL, OCPP version,
// APN, etc) on a Charging Station.
type SetNetworkProfileRequest struct {
	// Required. Slot in which the configuration should be stored.
	ConfigurationSlot int32 `json:"configurationSlot"`
	// Required. Connection details.
	ConnectionData NetworkConnectionProfileType `json:"connectionData"`
}

func (s *SetNetworkProfileRequest) UnmarshalJSON(data []byte) error {
	type Alias SetNetworkProfileRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetNetworkProfileRequest(a)
	return s.Validate()
}

func (s SetNetworkProfileRequest) Validate() error {
	if err := s.ConnectionData.Validate(); err != nil {
		return ocpp.WrapField("connectionData", err)
	}

	return nil
}

// SetNetworkProfileResponse (1.79.2)
//
// This contains the field definition of the SetNetworkProfileResponse PDU sent by the Charging Station to the
// CSMS in response to a SetNetworkProfileRequest.
type SetNetworkProfileResponse struct {
	// Required. Result of operation.
	Status SetNetworkProfileStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetNetworkProfileResponse) UnmarshalJSON(data []byte) error {
	type Alias SetNetworkProfileResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetNetworkProfileResponse(a)
	return s.Validate()
}

func (s SetNetworkProfileResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetVariableMonitoringRequest (1.80.1)
//
// This contains the field definition of the SetVariableMonitoringRequest PDU sent by the CSMS to the Charging
// Station.
type SetVariableMonitoringRequest struct {
	// Required. List of MonitoringData containing monitoring settings.
	SetMonitoringData []SetMonitoringDataType `json:"setMonitoringData"`
}

func (s *SetVariableMonitoringRequest) UnmarshalJSON(data []byte) error {
	type Alias SetVariableMonitoringRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetVariableMonitoringRequest(a)
	return s.Validate()
}

func (s SetVariableMonitoringRequest) Validate() error {
	if err := validateSliceLen(s.SetMonitoringData, 1, -1, "setMonitoringData"); err != nil {
		return err
	}

	for i, v := range s.SetMonitoringData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("setMonitoringData[%d]", i), err)
		}
	}

	return nil
}

// SetVariableMonitoringResponse (1.80.2)
//
// This contains the field definition of the SetVariableMonitoringResponse PDU sent by the Charging Station to
// the CSMS in response to a SetVariableMonitoringRequest.
type SetVariableMonitoringResponse struct {
	// Required. List of result statuses per monitor.
	SetMonitoringResult []SetMonitoringResultType `json:"setMonitoringResult"`
}

func (s *SetVariableMonitoringResponse) UnmarshalJSON(data []byte) error {
	type Alias SetVariableMonitoringResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetVariableMonitoringResponse(a)
	return s.Validate()
}

func (s SetVariableMonitoringResponse) Validate() error {
	if err := validateSliceLen(s.SetMonitoringResult, 1, -1, "setMonitoringResult"); err != nil {
		return err
	}

	for i, v := range s.SetMonitoringResult {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("setMonitoringResult[%d]", i), err)
		}
	}

	return nil
}

// SetVariablesRequest (1.81.1)
//
// This contains the field definition of the SetVariablesRequest PDU sent by the CSMS to the Charging Station.
type SetVariablesRequest struct {
	// Required. List of Component-Variable pairs and attribute values to set.
	SetVariableData []SetVariableDataType `json:"setVariableData"`
}

func (s *SetVariablesRequest) UnmarshalJSON(data []byte) error {
	type Alias SetVariablesRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetVariablesRequest(a)
	return s.Validate()
}

func (s SetVariablesRequest) Validate() error {
	if err := validateSliceLen(s.SetVariableData, 1, -1, "setVariableData"); err != nil {
		return err
	}

	for i, v := range s.SetVariableData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("setVariableData[%d]", i), err)
		}
	}

	return nil
}

// SetVariablesResponse (1.81.2)
//
// This contains the field definition of the SetVariablesResponse PDU sent by the Charging Station to the CSMS in
// response to a SetVariablesRequest.
type SetVariablesResponse struct {
	// Required. List of result statuses per Component-Variable.
	SetVariableResult []SetVariableResultType `json:"setVariableResult"`
}

func (s *SetVariablesResponse) UnmarshalJSON(data []byte) error {
	type Alias SetVariablesResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetVariablesResponse(a)
	return s.Validate()
}

func (s SetVariablesResponse) Validate() error {
	if err := validateSliceLen(s.SetVariableResult, 1, -1, "setVariableResult"); err != nil {
		return err
	}

	for i, v := range s.SetVariableResult {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("setVariableResult[%d]", i), err)
		}
	}

	return nil
}

// SignCertificateRequest (1.82.1)
//
// Sent by the Charging Station to the CSMS to request that the Certificate Authority signs the public key into a
// certificate.
type SignCertificateRequest struct {
	// Required. The Charging Station SHALL send the public key in form of a Certificate Signing Request (CSR) as
	// described in RFC 2986 [22] and then PEM encoded, using the SignCertificateRequest message.
	Csr string `json:"csr"`
	// Optional. Indicates the type of certificate that is to be signed. When omitted the certificate is to be used
	// for both the 15118 connection (if implemented) and the Charging Station to CSMS connection.
	CertificateType *CertificateSigningUseEnumType `json:"certificateType,omitempty"`
	// Optional. (2.1) RequestId to match this message with the CertificateSignedRequest.
	RequestID *int32 `json:"requestId,omitempty"`
	// Optional. (2.1) The hash of the root certificate to identify the PKI to use.
	HashRootCertificate *CertificateHashDataType `json:"hashRootCertificate,omitempty"`
}

func (s *SignCertificateRequest) UnmarshalJSON(data []byte) error {
	type Alias SignCertificateRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SignCertificateRequest(a)
	return s.Validate()
}

func (s SignCertificateRequest) Validate() error {
	if s.Csr == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "csr", "required field is missing")
	}

	if err := validateStringLen(s.Csr, 5500, "csr"); err != nil {
		return err
	}

	if s.HashRootCertificate != nil {
		if err := s.HashRootCertificate.Validate(); err != nil {
			return ocpp.WrapField("hashRootCertificate", err)
		}
	}

	return nil
}

// SignCertificateResponse (1.82.2)
//
// Sent by the CSMS to the Charging Station in response to the SignCertificateRequest message.
type SignCertificateResponse struct {
	// Required. Specifies whether the CSMS can process the request.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SignCertificateResponse) UnmarshalJSON(data []byte) error {
	type Alias SignCertificateResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SignCertificateResponse(a)
	return s.Validate()
}

func (s SignCertificateResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// StatusNotificationRequest (1.83.1)
//
// This contains the field definition of the StatusNotificationRequest PDU sent by the Charging Station to the
// CSMS. This message might be removed in a future version of OCPP. It will be replaced by Device Management
// Monitoring events.
type StatusNotificationRequest struct {
	// Required. The time for which the status is reported.
	Timestamp time.Time `json:"timestamp"`
	// Required. This contains the current status of the Connector.
	ConnectorStatus ConnectorStatusEnumType `json:"connectorStatus"`
	// Required. The id of the EVSE to which the connector belongs for which the the status is reported.
	EVSEID int32 `json:"evseId"`
	// Required. The id of the connector within the EVSE for which the status is reported.
	ConnectorID int32 `json:"connectorId"`
}

func (s *StatusNotificationRequest) UnmarshalJSON(data []byte) error {
	type Alias StatusNotificationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StatusNotificationRequest(a)
	return s.Validate()
}

func (s StatusNotificationRequest) Validate() error {
	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.ConnectorStatus == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "connectorStatus", "required field is missing")
	}

	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if err := validateNonNegative(s.ConnectorID, "connectorId"); err != nil {
		return err
	}

	return nil
}

// StatusNotificationResponse (1.83.2)
//
// This contains the field definition of StatusNotificationResponse sent by the CSMS to the Charging Station in
// response to a StatusNotificationRequest. This message might be removed in a future version of OCPP. It will be
// replaced by Device Management Monitoring events.
//
// No fields are defined.
//
// No fields are defined.
type StatusNotificationResponse struct {
}

func (s *StatusNotificationResponse) UnmarshalJSON(data []byte) error {
	type Alias StatusNotificationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StatusNotificationResponse(a)
	return s.Validate()
}

func (s StatusNotificationResponse) Validate() error {
	_ = s
	return nil
}

// TransactionEventRequest (1.84.1)
//
// This section contains the field definition of the TransactionEventRequest PDU sent by the Charging Station to
// the CSMS. For each of the eventTypes; Started, Updated and Ended, the corresponding cardinality is specified.
type TransactionEventRequest struct {
	// Required. This contains the type of this event. The first TransactionEvent of a transaction SHALL contain:
	// "Started" The last TransactionEvent of a transaction SHALL contain: "Ended" All others SHALL contain:
	// "Updated"
	EventType TransactionEventEnumType `json:"eventType"`
	// Required. The date and time at which this transaction event occurred.
	Timestamp time.Time `json:"timestamp"`
	// Required. Reason the Charging Station sends this message to the CSMS
	TriggerReason TriggerReasonEnumType `json:"triggerReason"`
	// Required. Incremental sequence number, helps with determining if all messages of a transaction have been
	// received.
	SeqNo int32 `json:"seqNo"`
	// Optional. Indication that this transaction event happened when the Charging Station was offline. Default =
	// false, meaning: the event occurred when the Charging Station was online.
	Offline *bool `json:"offline,omitempty"`
	// Optional. If the Charging Station is able to report the number of phases used, then it SHALL provide it. When
	// omitted the CSMS may be able to determine the number of phases used as follows: 1: The numberPhases in the
	// currently used ChargingSchedule. 2: The number of phases provided via device management.
	NumberOfPhasesUsed *int32 `json:"numberOfPhasesUsed,omitempty"`
	// Optional. The maximum current of the connected cable in Ampere (A).
	CableMaxCurrent *int32 `json:"cableMaxCurrent,omitempty"`
	// Optional. This contains the Id of the reservation that terminates as a result of this transaction.
	ReservationID *int32 `json:"reservationId,omitempty"`
	// Optional. (2.1) The current preconditioning status of the BMS in the EV. Default value is Unknown.
	PreconditioningStatus *PreconditioningStatusEnumType `json:"preconditioningStatus,omitempty"`
	// Optional. (2.1) True when EVSE electronics are in sleep mode for this transaction. Default value (when absent)
	// is false.
	EVSESleep *bool `json:"evseSleep,omitempty"`
	// Optional. This contains the meter values relevant to the transaction. Depending on the eventType of this
	// TransactionEventRequest the following Configuration Variable is used to configure which measurands are used:
	// Started: SampledDataTxStartedMeasurands Updated: SampledDataTxUpdatedMeasurands and AlignedDataMeasurands
	// Ended: SampledDataTxEndedMeasurands and AlignedDataTxEndedMeasurands
	MeterValue []MeterValueType `json:"meterValue,omitempty"`
	// Optional. This contains the identifier for which a transaction is (or will be) started or stopped. Is required
	// when the EV Driver becomes authorized for this transaction and when the EV Driver ends authorization. The
	// IdToken should only be sent once in a TransactionEventRequest for every authorization (for starting or for
	// stopping) done for this transaction, so that CSMS can return the idTokenInfo in the TransactionEventResponse.
	// idToken should not be present in the TransactionEventRequest when a transaction is ended by a
	// RequestStopTransactionRequest or a ResetRequest.
	IDToken *IDTokenType `json:"idToken,omitempty"`
	// Optional. This identifies which evse (and connector) of the Charging Station is used.
	EVSE *EVSEType `json:"evse,omitempty"`
	// Required. Contains transaction specific information.
	TransactionInfo TransactionType `json:"transactionInfo"`
	// Optional. (2.1) Optional. Only required in TransactionEventRequest('Ended') and only if Charging Station
	// calculated cost locally.
	CostDetails *CostDetailsType `json:"costDetails,omitempty"`
}

func (s *TransactionEventRequest) UnmarshalJSON(data []byte) error {
	type Alias TransactionEventRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TransactionEventRequest(a)
	return s.Validate()
}

func (s TransactionEventRequest) Validate() error {
	if s.EventType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "eventType", "required field is missing")
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.TriggerReason == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "triggerReason", "required field is missing")
	}

	if err := validateNonNegative(s.SeqNo, "seqNo"); err != nil {
		return err
	}

	if s.NumberOfPhasesUsed != nil {
		if err := validateNonNegative(*s.NumberOfPhasesUsed, "numberOfPhasesUsed"); err != nil {
			return err
		}
	}

	if s.ReservationID != nil {
		if err := validateNonNegative(*s.ReservationID, "reservationId"); err != nil {
			return err
		}
	}

	for i, v := range s.MeterValue {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("meterValue[%d]", i), err)
		}
	}

	if s.IDToken != nil {
		if err := s.IDToken.Validate(); err != nil {
			return ocpp.WrapField("idToken", err)
		}
	}

	if s.EVSE != nil {
		if err := s.EVSE.Validate(); err != nil {
			return ocpp.WrapField("evse", err)
		}
	}

	if err := s.TransactionInfo.Validate(); err != nil {
		return ocpp.WrapField("transactionInfo", err)
	}

	if s.CostDetails != nil {
		if err := s.CostDetails.Validate(); err != nil {
			return ocpp.WrapField("costDetails", err)
		}
	}

	return nil
}

// TransactionEventResponse (1.84.2)
//
// This contains the field definition of the TransactionEventResponse PDU sent by the CSMS to the Charging
// Station in response to a TransactionEventRequest.
type TransactionEventResponse struct {
	// Optional. When eventType of TransactionEventRequest is Updated, then this value contains the running cost.
	// When eventType of TransactionEventRequest is Ended, then this contains the final total cost of this
	// transaction, including taxes, in the currency configured with the Configuration Variable: Currency. Absence of
	// this value does not imply that the transaction was free. To indicate a free transaction, the CSMS SHALL send a
	// value of 0.00.
	TotalCost *float64 `json:"totalCost,omitempty"`
	// Optional. Priority from a business point of view. Default priority is 0, The range is from -9 to 9. Higher
	// values indicate a higher priority. The chargingPriority in TransactionEventResponse is temporarily, so it may
	// not be set in the IdTokenInfoType afterwards. Also the chargingPriority in TransactionEventResponse has a
	// higher priority than the one in IdTokenInfoType.
	ChargingPriority *int32 `json:"chargingPriority,omitempty"`
	// Optional. This contains information about authorization status, expiry and group id. Is required when the
	// transactionEventRequest contained an idToken.
	IDTokenInfo *IDTokenInfoType `json:"idTokenInfo,omitempty"`
	// Optional. This can contain updated personal message that can be shown to the EV Driver. This can be used to
	// provide updated tariff information .
	UpdatedPersonalMessage *MessageContentType `json:"updatedPersonalMessage,omitempty"`
	// Optional. (2.1) Additional languages besides the default language in updatedPersonalMessage.
	UpdatedPersonalMessageExtra []MessageContentType `json:"updatedPersonalMessageExtra,omitempty"`
	// Optional. (2.1) Maximum cost/energy/time limit allowed for this transaction.
	TransactionLimit *TransactionLimitType `json:"transactionLimit,omitempty"`
}

func (s *TransactionEventResponse) UnmarshalJSON(data []byte) error {
	type Alias TransactionEventResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TransactionEventResponse(a)
	return s.Validate()
}

func (s TransactionEventResponse) Validate() error {
	if s.IDTokenInfo != nil {
		if err := s.IDTokenInfo.Validate(); err != nil {
			return ocpp.WrapField("idTokenInfo", err)
		}
	}

	if s.UpdatedPersonalMessage != nil {
		if err := s.UpdatedPersonalMessage.Validate(); err != nil {
			return ocpp.WrapField("updatedPersonalMessage", err)
		}
	}

	if err := validateSliceLen(s.UpdatedPersonalMessageExtra, 0, 4, "updatedPersonalMessageExtra"); err != nil {
		return err
	}

	for i, v := range s.UpdatedPersonalMessageExtra {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("updatedPersonalMessageExtra[%d]", i), err)
		}
	}

	if s.TransactionLimit != nil {
		if err := s.TransactionLimit.Validate(); err != nil {
			return ocpp.WrapField("transactionLimit", err)
		}
	}

	return nil
}

// TriggerMessageRequest (1.85.1)
//
// This contains the field definition of the TriggerMessageRequest PDU sent by the CSMS to the Charging Station.
type TriggerMessageRequest struct {
	// Required. Type of message to be triggered.
	RequestedMessage MessageTriggerEnumType `json:"requestedMessage"`
	// Optional. (2.1) When requestedMessage = CustomTrigger this will trigger sending the corresponding message in
	// field customTrigger, if supported by Charging Station.
	CustomTrigger *string `json:"customTrigger,omitempty"`
	// Optional. Can be used to specifiy the EVSE and Connector if required for the message which needs to be sent.
	EVSE *EVSEType `json:"evse,omitempty"`
}

func (s *TriggerMessageRequest) UnmarshalJSON(data []byte) error {
	type Alias TriggerMessageRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TriggerMessageRequest(a)
	return s.Validate()
}

func (s TriggerMessageRequest) Validate() error {
	if s.RequestedMessage == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "requestedMessage", "required field is missing")
	}

	if s.CustomTrigger != nil {
		if err := validateStringLen(*s.CustomTrigger, 50, "customTrigger"); err != nil {
			return err
		}
	}

	if s.EVSE != nil {
		if err := s.EVSE.Validate(); err != nil {
			return ocpp.WrapField("evse", err)
		}
	}

	return nil
}

// TriggerMessageResponse (1.85.2)
//
// This contains the field definition of the TriggerMessageResponse PDU sent by the Charging Station to the CSMS
// in response to TriggerMessageResponse.
type TriggerMessageResponse struct {
	// Required. Indicates whether the Charging Station will send the requested notification or not.
	Status TriggerMessageStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *TriggerMessageResponse) UnmarshalJSON(data []byte) error {
	type Alias TriggerMessageResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TriggerMessageResponse(a)
	return s.Validate()
}

func (s TriggerMessageResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// UnlockConnectorRequest (1.86.1)
//
// This contains the field definition of the UnlockConnectorRequest PDU sent by the CSMS to the Charging Station.
type UnlockConnectorRequest struct {
	// Required. This contains the identifier of the EVSE for which a connector needs to be unlocked.
	EVSEID int32 `json:"evseId"`
	// Required. This contains the identifier of the connector that needs to be unlocked.
	ConnectorID int32 `json:"connectorId"`
}

func (s *UnlockConnectorRequest) UnmarshalJSON(data []byte) error {
	type Alias UnlockConnectorRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnlockConnectorRequest(a)
	return s.Validate()
}

func (s UnlockConnectorRequest) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if err := validateNonNegative(s.ConnectorID, "connectorId"); err != nil {
		return err
	}

	return nil
}

// UnlockConnectorResponse (1.86.2)
//
// This contains the field definition of the UnlockConnectorResponse PDU sent by the Charging Station to the CSMS
// in response to an UnlockConnectorRequest.
type UnlockConnectorResponse struct {
	// Required. This indicates whether the Charging Station has unlocked the connector.
	Status UnlockStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *UnlockConnectorResponse) UnmarshalJSON(data []byte) error {
	type Alias UnlockConnectorResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnlockConnectorResponse(a)
	return s.Validate()
}

func (s UnlockConnectorResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// UnpublishFirmwareRequest (1.87.1)
//
// This contains the field definition of the UnpublishFirmwareRequest PDU sent by the CSMS to the Charging
// Station.
type UnpublishFirmwareRequest struct {
	// Required. The MD5 checksum over the entire firmware file as a hexadecimal string of length 32.
	Checksum IdentifierString32Type `json:"checksum"`
}

func (s *UnpublishFirmwareRequest) UnmarshalJSON(data []byte) error {
	type Alias UnpublishFirmwareRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnpublishFirmwareRequest(a)
	return s.Validate()
}

func (s UnpublishFirmwareRequest) Validate() error {
	if s.Checksum == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "checksum", "required field is missing")
	}

	if err := s.Checksum.Validate(); err != nil {
		return ocpp.WrapField("checksum", err)
	}

	return nil
}

// UnpublishFirmwareResponse (1.87.2)
//
// This contains the field definition of the UnpublishFirmwareResponse PDU sent by the Charging Station to the
// CSMS in response to a UnpublishFirmwareRequest.
type UnpublishFirmwareResponse struct {
	// Required. Indicates whether the Local Controller succeeded in unpublishing the firmware.
	Status UnpublishFirmwareStatusEnumType `json:"status"`
}

func (s *UnpublishFirmwareResponse) UnmarshalJSON(data []byte) error {
	type Alias UnpublishFirmwareResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnpublishFirmwareResponse(a)
	return s.Validate()
}

func (s UnpublishFirmwareResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	return nil
}

// UpdateDynamicScheduleRequest (1.88.1)
//
// (2.1) This message is used to update a setpoint or limit in a dynamic charging profile.
type UpdateDynamicScheduleRequest struct {
	// Required. Id of charging profile to update.
	ChargingProfileID int32 `json:"chargingProfileId"`
	// Required. Updated values for charging schedule period.
	ScheduleUpdate ChargingScheduleUpdateType `json:"scheduleUpdate"`
}

func (s *UpdateDynamicScheduleRequest) UnmarshalJSON(data []byte) error {
	type Alias UpdateDynamicScheduleRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UpdateDynamicScheduleRequest(a)
	return s.Validate()
}

func (s UpdateDynamicScheduleRequest) Validate() error {
	if err := s.ScheduleUpdate.Validate(); err != nil {
		return ocpp.WrapField("scheduleUpdate", err)
	}

	return nil
}

// UpdateDynamicScheduleResponse (1.88.2)
//
// (2.1) Returns whether the Charging Station has been able to process the message successfully.
type UpdateDynamicScheduleResponse struct {
	// Required. Returns whether message was processed successfully.
	Status ChargingProfileStatusEnumType `json:"status"`
	// Optional. Detailed status info.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *UpdateDynamicScheduleResponse) UnmarshalJSON(data []byte) error {
	type Alias UpdateDynamicScheduleResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UpdateDynamicScheduleResponse(a)
	return s.Validate()
}

func (s UpdateDynamicScheduleResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// UpdateFirmwareRequest (1.89.1)
//
// This contains the field definition of the UpdateFirmwareRequest PDU sent by the CSMS to the Charging Station.
type UpdateFirmwareRequest struct {
	// Optional. This specifies how many times Charging Station must retry to download the firmware before giving up.
	// If this field is not present, it is left to Charging Station to decide how many times it wants to retry. If
	// the value is 0, it means: no retries.
	Retries *int32 `json:"retries,omitempty"`
	// Optional. The interval in seconds after which a retry may be attempted. If this field is not present, it is
	// left to Charging Station to decide how long to wait between attempts.
	RetryInterval *int32 `json:"retryInterval,omitempty"`
	// Required. The Id of this request
	RequestID int32 `json:"requestId"`
	// Required. Specifies the firmware to be updated on the Charging Station.
	Firmware FirmwareType `json:"firmware"`
}

func (s *UpdateFirmwareRequest) UnmarshalJSON(data []byte) error {
	type Alias UpdateFirmwareRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UpdateFirmwareRequest(a)
	return s.Validate()
}

func (s UpdateFirmwareRequest) Validate() error {
	if s.Retries != nil {
		if err := validateNonNegative(*s.Retries, "retries"); err != nil {
			return err
		}
	}

	if err := s.Firmware.Validate(); err != nil {
		return ocpp.WrapField("firmware", err)
	}

	return nil
}

// UpdateFirmwareResponse (1.89.2)
//
// This contains the field definition of the UpdateFirmwareResponse PDU sent by the Charging Station to the CSMS
// in response to an UpdateFirmwareRequest.
type UpdateFirmwareResponse struct {
	// Required. This field indicates whether the Charging Station was able to accept the request.
	Status UpdateFirmwareStatusEnumType `json:"status"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *UpdateFirmwareResponse) UnmarshalJSON(data []byte) error {
	type Alias UpdateFirmwareResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UpdateFirmwareResponse(a)
	return s.Validate()
}

func (s UpdateFirmwareResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// UsePriorityChargingRequest (1.90.1)
//
// (2.1) Message sent by CSMS to tell Charging Station to switch to the priority charging profile, that allows
// for the maximum possible current or power under current circumstances. Message contains a transactionId,
// because it only applies to the transaction in progress.
type UsePriorityChargingRequest struct {
	// Required. The transaction for which priority charging is requested.
	TransactionID IdentifierString36Type `json:"transactionId"`
	// Required. True to request priority charging. False to request stopping priority charging.
	Activate bool `json:"activate"`
}

func (s *UsePriorityChargingRequest) UnmarshalJSON(data []byte) error {
	type Alias UsePriorityChargingRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UsePriorityChargingRequest(a)
	return s.Validate()
}

func (s UsePriorityChargingRequest) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	return nil
}

// UsePriorityChargingResponse (1.90.2)
//
// (2.1) Status of the UsePriorityChargingRequest.
type UsePriorityChargingResponse struct {
	// Required. Result of the request.
	Status PriorityChargingStatusEnumType `json:"status"`
	// Optional.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *UsePriorityChargingResponse) UnmarshalJSON(data []byte) error {
	type Alias UsePriorityChargingResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UsePriorityChargingResponse(a)
	return s.Validate()
}

func (s UsePriorityChargingResponse) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// VatNumberValidationRequest (1.91.1)
type VatNumberValidationRequest struct {
	// Required. VAT number to check.
	VATNumber string `json:"vatNumber"`
	// Optional. EVSE id for which check is done
	EVSEID *int32 `json:"evseId,omitempty"`
}

func (s *VatNumberValidationRequest) UnmarshalJSON(data []byte) error {
	type Alias VatNumberValidationRequest
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VatNumberValidationRequest(a)
	return s.Validate()
}

func (s VatNumberValidationRequest) Validate() error {
	if s.VATNumber == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "vatNumber", "required field is missing")
	}

	if err := validateStringLen(s.VATNumber, 20, "vatNumber"); err != nil {
		return err
	}

	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	return nil
}

// VatNumberValidationResponse (1.91.2)
type VatNumberValidationResponse struct {
	// Required. VAT number that was requested.
	VATNumber string `json:"vatNumber"`
	// Optional. EVSE id for which check was requested.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. Result of operation.
	Status GenericStatusEnumType `json:"status"`
	// Optional. Company address associated with vatNumber.
	Company *AddressType `json:"company,omitempty"`
	// Optional. Additional info on status
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *VatNumberValidationResponse) UnmarshalJSON(data []byte) error {
	type Alias VatNumberValidationResponse
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VatNumberValidationResponse(a)
	return s.Validate()
}

func (s VatNumberValidationResponse) Validate() error {
	if s.VATNumber == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "vatNumber", "required field is missing")
	}

	if err := validateStringLen(s.VATNumber, 20, "vatNumber"); err != nil {
		return err
	}

	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.Company != nil {
		if err := s.Company.Validate(); err != nil {
			return ocpp.WrapField("company", err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}
