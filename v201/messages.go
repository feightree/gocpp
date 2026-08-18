package v201

import (
	"encoding/json"
	"fmt"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// AuthorizeRequest (1.1.1)
//
// This contains the field definition of the AuthorizeRequest PDU sent by the Charging Station to the CSMS.
type AuthorizeRequest struct {
	// Optional. The X.509 certificate chain presented by EV and encoded in PEM format. Order of certificates in chain is
	// from leaf up to (but excluding) root certificate. Only needed in case of central contract validation when Charging
	// Station cannot validate the contract certificate.
	Certificate *string `json:"certificate,omitempty"`
	// Required. This contains the identifier that needs to be authorized.
	IDToken IdTokenType `json:"idToken"`
	// Optional. Contains the information needed to verify the EV Contract Certificate via OCSP. Not needed if certificate
	// is provided.
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
		if err := validateStringLen(*s.Certificate, 5500, "certificate"); err != nil {
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

// AuthorizeResponse (1.1.2)
//
// This contains the field definition of the AuthorizeResponse PDU sent by the CSMS to the Charging Station in
// response to an AuthorizeRequest.
type AuthorizeResponse struct {
	// Optional. Certificate status information. - if all certificates are valid: return 'Accepted'. - if one of the
	// certificates was revoked, return 'CertificateRevoked'.
	CertificateStatus *AuthorizeCertificateStatusEnumType `json:"certificateStatus,omitempty"`
	// Required. This contains information about authorization status, expiry and group id.
	IDTokenInfo IdTokenInfoType `json:"idTokenInfo"`
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

	return nil
}

// BootNotificationRequest (1.2.1)
//
// This contains the field definition of the BootNotificationRequest PDU sent by the Charging Station to the CSMS.
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

// BootNotificationResponse (1.2.2)
//
// This contains the field definition of the BootNotificationResponse PDU sent by the CSMS to the Charging Station in
// response to a BootNotificationRequest.
type BootNotificationResponse struct {
	// Required. This contains the CSMS’s current time.
	CurrentTime time.Time `json:"currentTime"`
	// Required. When Status is Accepted, this contains the heartbeat interval in seconds. If the CSMS returns something
	// other than Accepted, the value of the interval field indicates the minimum wait time before sending a next
	// BootNotification request.
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

// CancelReservationRequest (1.3.1)
//
// This contains the field definition of the CancelReservationRequest PDU sent by the CSMS to the Charging Station.
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
	_ = s
	return nil
}

// CancelReservationResponse (1.3.2)
//
// This contains the field definition of the CancelReservationResponse PDU sent by the Charging Station to the CSMS in
// response to a CancelReservationRequest.
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

// CertificateSignedRequest (1.4.1)
//
// This contains the field definition of the CertificateSignedRequest PDU sent by the CSMS to the Charging Station.
type CertificateSignedRequest struct {
	// Required. The signed PEM encoded X.509 certificate. This SHALL also contain the necessary sub CA certificates, when
	// applicable. The order of the bundle follows the certificate chain, starting from the leaf certificate. The
	// Configuration Variable MaxCertificateChainSize can be used to limit the maximum size of this field.
	CertificateChain string `json:"certificateChain"`
	// Optional. Indicates the type of the signed certificate that is returned. When omitted the certificate is used for
	// both the 15118 connection (if implemented) and the Charging Station to CSMS connection. This field is required when
	// a certificateType was included in the SignCertificateRequest that requested this certificate to be signed AND both
	// the 15118 connection and the Charging Station connection are implemented.
	CertificateType *CertificateSigningUseEnumType `json:"certificateType,omitempty"`
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

// CertificateSignedResponse (1.4.2)
//
// This contains the field definition of the CertificateSignedResponse PDU sent by the Charging Station to the CSMS in
// response to a CertificateSignedRequest.
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

// ChangeAvailabilityRequest (1.5.1)
//
// This contains the field definition of the ChangeAvailabilityRequest PDU sent by the CSMS to the Charging Station.
type ChangeAvailabilityRequest struct {
	// Required. This contains the type of availability change that the Charging Station should perform.
	OperationalStatus OperationalStatusEnumType `json:"operationalStatus"`
	// Optional. Contains Id’s to designate a specific EVSE/connector by index numbers. When omitted, the message refers
	// to the Charging Station as a whole.
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

// ChangeAvailabilityResponse (1.5.2)
//
// This contains the field definition of the ChangeAvailabilityResponse PDU sent by the Charging Station to the CSMS.
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

// ClearCacheRequest (1.6.1)
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

// ClearCacheResponse (1.6.2)
//
// This contains the field definition of the ClearCacheResponse PDU sent by the Charging Station to the CSMS in
// response to a
//
// ClearCacheRequest.
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

// ClearChargingProfileRequest (1.7.1)
//
// This contains the field definition of the ClearChargingProfileRequest PDU sent by the CSMS to the Charging Station.
// The CSMS can use this message to clear (remove) either a specific charging profile (denoted by id) or a selection
// of charging profiles that match with the values of the optional evse, stackLevel and ChargingProfilePurpose fields.
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

// ClearChargingProfileResponse (1.7.2)
//
// This contains the field definition of the ClearChargingProfileResponse PDU sent by the Charging Station to the CSMS
// in response to a ClearChargingProfileRequest.
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

// ClearDisplayMessageRequest (1.8.1)
//
// This contains the field definition of the ClearDisplayMessageRequest PDU sent by the CSMS to the Charging Station.
// The CSMS asks the Charging Station to clear a display message that has been configured in the Charging Station to
// be cleared/removed. See also O05 - Clear a Display Message.
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
	_ = s
	return nil
}

// ClearDisplayMessageResponse (1.8.2)
//
// This contains the field definition of the ClearDisplayMessageResponse PDU sent by the Charging Station to the CSMS
// in a response to a ClearDisplayMessageRequest. See also O05 - Clear a Display Message.
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

// ClearedChargingLimitRequest (1.9.1)
//
// This contains the field definition of the ClearedChargingLimitRequest PDU sent by the Charging Station to the CSMS.
type ClearedChargingLimitRequest struct {
	// Required. Source of the charging limit.
	ChargingLimitSource ChargingLimitSourceEnumType `json:"chargingLimitSource"`
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

	return nil
}

// ClearedChargingLimitResponse (1.9.2)
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

// ClearVariableMonitoringRequest (1.10.1)
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

	return nil
}

// ClearVariableMonitoringResponse (1.10.2)
//
// This contains the field definition of the ClearVariableMonitoringResponse PDU sent by the Charging Station to the
// CSMS.
type ClearVariableMonitoringResponse struct {
	// Required. List of result statuses per monitor.
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

// CostUpdatedRequest (1.11.1)
//
// This contains the field definition of the CostUpdatedRequest PDU sent by the CSMS to the Charging Station. With
// this request the CSMS can send the current cost of a transaction to a Charging Station.
type CostUpdatedRequest struct {
	// Required. Current total cost, based on the information known by the CSMS, of the transaction including taxes. In
	// the currency configured with the configuration Variable: [Currency]
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

// CostUpdatedResponse (1.11.2)
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

// CustomerInformationRequest (1.12.1)
//
// This contains the field definition of the CustomerInformationRequest PDU sent by the CSMS to the Charging Station.
type CustomerInformationRequest struct {
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Required. Flag indicating whether the Charging Station should return NotifyCustomerInformationRequest messages
	// containing information about the customer referred to.
	Report bool `json:"report"`
	// Required. Flag indicating whether the Charging Station should clear all information about the customer referred to.
	Clear bool `json:"clear"`
	// Optional. A (e.g. vendor specific) identifier of the customer this request refers to. This field contains a custom
	// identifier other than IdToken and Certificate. One of the possible identifiers (customerIdentifier, customerIdToken
	// or customerCertificate) should be in the request message.
	CustomerIdentifier *string `json:"customerIdentifier,omitempty"`
	// Optional. The IdToken of the customer this request refers to. One of the possible identifiers (customerIdentifier,
	// customerIdToken or customerCertificate) should be in the request message.
	IDToken *IdTokenType `json:"idToken,omitempty"`
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

// CustomerInformationResponse (1.12.2)
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

// DataTransferRequest (1.13.1)
//
// This contains the field definition of the DataTransferRequest PDU sent either by the CSMS to the Charging Station
// or vice versa.
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

// DataTransferResponse (1.13.2)
//
// This contains the field definition of the DataTransferResponse PDU sent by the Charging Station to the CSMS or vice
// versa in response to a DataTransferRequest.
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

// DeleteCertificateRequest (1.14.1)
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

// DeleteCertificateResponse (1.14.2)
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

// FirmwareStatusNotificationRequest (1.15.1)
//
// This contains the field definition of the FirmwareStatusNotificationRequest PDU sent by the Charging Station to the
// CSMS.
type FirmwareStatusNotificationRequest struct {
	// Required. This contains the progress status of the firmware installation.
	Status FirmwareStatusEnumType `json:"status"`
	// Optional. The request id that was provided in the UpdateFirmwareRequest that started this firmware update. This
	// field is mandatory, unless the message was triggered by a TriggerMessageRequest AND there is no firmware update
	// ongoing.
	RequestID *int32 `json:"requestId,omitempty"`
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

	return nil
}

// FirmwareStatusNotificationResponse (1.15.2)
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

// Get15118EVCertificateRequest (1.16.1)
//
// This message is sent by the Charging Station to the CSMS if an ISO 15118 vehicle selects the service Certificate
// installation. NOTE: This message is based on CertificateInstallationReq Res from ISO 15118 2.
type Get15118EVCertificateRequest struct {
	// Required. Schema version currently used for the 15118 session between EV and Charging Station. Needed for parsing
	// of the EXI stream by the CSMS.
	Iso15118SchemaVersion string `json:"iso15118SchemaVersion"`
	// Required. Defines whether certificate needs to be installed or updated.
	Action CertificateActionEnumType `json:"action"`
	// Required. Raw CertificateInstallationReq request from EV, Base64 encoded.
	ExiRequest string `json:"exiRequest"`
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

	if err := validateStringLen(s.ExiRequest, 5600, "exiRequest"); err != nil {
		return err
	}

	return nil
}

// Get15118EVCertificateResponse (1.16.2)
//
// Response message from CSMS to Charging Station containing the status and optionally new certificate. NOTE: This
// message is based on CertificateInstallationReq Res from ISO 15118-2.
type Get15118EVCertificateResponse struct {
	// Required. Indicates whether the message was processed properly.
	Status Iso15118EVCertificateStatusEnumType `json:"status"`
	// Required. Raw CertificateInstallationRes response for the EV, Base64 encoded. The Charging Station can let the CSMS
	// know it supports a higher field size by reporting this using the device model as
	// OCPPCommCtrlr.FieldLength["Get15118EVCertificateRes ponse.exiResponse"] = <New max length>
	ExiResponse string `json:"exiResponse"`
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

	if err := validateStringLen(s.ExiResponse, 5600, "exiResponse"); err != nil {
		return err
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetBaseReportRequest (1.17.1)
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

// GetBaseReportResponse (1.17.2)
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

// GetCertificateStatusRequest (1.18.1)
//
// This contains the field definition of the GetCertificateStatusRequest PDU sent by the Charging Station to the CSMS.
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

// GetCertificateStatusResponse (1.18.2)
//
// This contains the field definition of the GetCertificateStatusResponse PDU sent by the CSMS to the Charging
// Station.
type GetCertificateStatusResponse struct {
	// Required. This indicates whether the charging station was able to retrieve the OCSP certificate status.
	Status GetCertificateStatusEnumType `json:"status"`
	// Optional. OCSPResponse class as defined in IETF RFC 6960. DER encoded (as defined in IETF RFC 6960), and then
	// base64 encoded. MAY only be omitted when status is not Accepted.
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
		if err := validateStringLen(*s.OCSPResult, 5500, "ocspResult"); err != nil {
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

// GetChargingProfilesRequest (1.19.1)
//
// The message GetChargingProfilesRequest can be used by the CSMS to request installed charging profiles from the
// Charging Station. The charging profiles will then be reported by the Charging Station via
// ReportChargingProfilesRequest messages.
type GetChargingProfilesRequest struct {
	// Required. Reference identification that is to be used by the Charging Station in the ReportChargingProfilesRequest
	// when provided.
	RequestID int32 `json:"requestId"`
	// Optional. For which EVSE installed charging profiles SHALL be reported. If 0, only charging profiles installed on
	// the Charging Station itself (the grid connection) SHALL be reported. If omitted, all installed charging profiles
	// SHALL be reported. Reported charging profiles SHALL match the criteria in field chargingProfile.
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
	if err := s.ChargingProfile.Validate(); err != nil {
		return ocpp.WrapField("chargingProfile", err)
	}

	return nil
}

// GetChargingProfilesResponse (1.19.2)
//
// This contains the field definition of the GetChargingProfilesResponse PDU sent by the Charging Station to the CSMS
// in response to a GetChargingProfilesRequest.
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

// GetCompositeScheduleRequest (1.20.1)
//
// This contains the field definition of the GetCompositeScheduleRequest PDU sent by the CSMS to the Charging Station.
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
	_ = s
	return nil
}

// GetCompositeScheduleResponse (1.20.2)
//
// This contains the field definition of the GetCompositeScheduleResponse PDU sent by the Charging Station to the CSMS
// in response to a GetCompositeScheduleRequest .
type GetCompositeScheduleResponse struct {
	// Required. The Charging Station will indicate if it was able to process the request
	Status GenericStatusEnumType `json:"status"`
	// Optional. This field contains the calculated composite schedule. It may only be omitted when this message contains
	// status Rejected.
	Schedule *CompositeScheduleType `json:"schedule,omitempty"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
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

	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return ocpp.WrapField("schedule", err)
		}
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// GetDisplayMessagesRequest (1.21.1)
type GetDisplayMessagesRequest struct {
	// Optional. If provided the Charging Station shall return Display Messages of the given ids. This field SHALL NOT
	// contain more ids than set in NumberOfDisplayMessages.maxLimit
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
	_ = s
	return nil
}

// GetDisplayMessagesResponse (1.21.2)
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

// GetInstalledCertificateIdsRequest (1.22.1)
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

// GetInstalledCertificateIdsResponse (1.22.2)
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

// GetLocalListVersionRequest (1.23.1)
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

// GetLocalListVersionResponse (1.23.2)
//
// This contains the field definition of the GetLocalListVersionResponse PDU sent by the Charging Station to CSMS in
// response to a GetLocalListVersionRequest.
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

// GetLogRequest (1.24.1)
//
// This contains the field definition of the GetLogRequest PDU sent by the CSMS to the Charging Station.
type GetLogRequest struct {
	// Required. This contains the type of log file that the Charging Station should send.
	LogType LogEnumType `json:"logType"`
	// Required. The Id of this request
	RequestID int32 `json:"requestId"`
	// Optional. This specifies how many times the Charging Station must retry to upload the log before giving up. If this
	// field is not present, it is left to Charging Station to decide how many times it wants to retry. If the value is 0,
	// it means: no retries.
	Retries *int32 `json:"retries,omitempty"`
	// Optional. The interval in seconds after which a retry may be attempted. If this field is not present, it is left to
	// Charging Station to decide how long to wait between attempts.
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

	if err := s.Log.Validate(); err != nil {
		return ocpp.WrapField("log", err)
	}

	return nil
}

// GetLogResponse (1.24.2)
//
// This contains the field definition of the GetLogResponse PDU sent by the Charging Station to the CSMS in response
// to a GetLogRequest.
type GetLogResponse struct {
	// Required. This field indicates whether the Charging Station was able to accept the request.
	Status LogStatusEnumType `json:"status"`
	// Optional. This contains the name of the log file that will be uploaded. This field is not present when no logging
	// information is available.
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

// GetMonitoringReportRequest (1.25.1)
//
// This contains the field definition of the GetMonitoringReportRequest PDU sent by the CSMS to the Charging Station.
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

// GetMonitoringReportResponse (1.25.2)
//
// This contains the field definition of the GetMonitoringReportResponse PDU sent by the Charging Station to the CSMS.
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

// GetReportRequest (1.26.1)
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

// GetReportResponse (1.26.2)
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

// GetTransactionStatusRequest (1.27.1)
//
// With this message, the CSMS can ask the Charging Station whether it has transaction-related messages waiting to be
// delivered to the CSMS. When a transactionId is provided, only messages for a specific transaction are asked for.
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

// GetTransactionStatusResponse (1.27.2)
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

// GetVariablesRequest (1.28.1)
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

// GetVariablesResponse (1.28.2)
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

// HeartbeatRequest (1.29.1)
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

// HeartbeatResponse (1.29.2)
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

// InstallCertificateRequest (1.30.1)
//
// Used by the CSMS to request installation of a certificate on a Charging Station. Note: This message is not for
// installing a TLS client certificate in a charging station. The CertificateSignedRequest mechanism is used for that.
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

	if err := validateStringLen(s.Certificate, 5500, "certificate"); err != nil {
		return err
	}

	return nil
}

// InstallCertificateResponse (1.30.2)
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

// LogStatusNotificationRequest (1.31.1)
//
// This contains the field definition of the LogStatusNotificationRequest PDU sent by the Charging Station to the
// CSMS.
type LogStatusNotificationRequest struct {
	// Required. This contains the status of the log upload.
	Status UploadLogStatusEnumType `json:"status"`
	// Optional. The request id that was provided in GetLogRequest that started this log upload. This field is mandatory,
	// unless the message was triggered by a TriggerMessageRequest AND there is no log upload ongoing.
	RequestID *int32 `json:"requestId,omitempty"`
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

	return nil
}

// LogStatusNotificationResponse (1.31.2)
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

// MeterValuesRequest (1.32.1)
//
// This contains the field definition of the MeterValuesRequest PDU sent by the Charging Station to the CSMS. This
// message might be removed in a future version of OCPP. It will be replaced by Device Management Monitoring events.
type MeterValuesRequest struct {
	// Required. This contains a number (>0) designating an EVSE of the Charging Station. ‘0’ (zero) is used to designate
	// the main power meter.
	EVSEID int32 `json:"evseId"`
	// Required. The sampled meter values with timestamps.
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

// MeterValuesResponse (1.32.2)
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

// NotifyChargingLimitRequest (1.33.1)
//
// The message NotifyChargingLimitRequest can be used to communicate a charging limit, set by an external system on
// the Charging Station (Not installed by the CSO via SetChargingProfileRequest), to the CSMS.
type NotifyChargingLimitRequest struct {
	// Optional. The charging schedule contained in this notification applies to an EVSE. evseId must be > 0.
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

// NotifyChargingLimitResponse (1.33.2)
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

// NotifyCustomerInformationRequest (1.34.1)
//
// This contains the field definition of the NotifyCustomerInformationRequest PDU sent by the Charging Station to the
// CSMS.
type NotifyCustomerInformationRequest struct {
	// Required. (Part of) the requested data. No format specified in which the data is returned. Should be human
	// readable.
	Data string `json:"data"`
	// Optional. “to be continued” indicator. Indicates whether another part of the monitoringData follows in an upcoming
	// notifyMonitoringReportRequest message. Default value when omitted is false.
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

	if s.GeneratedAt.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "generatedAt", "required field is missing")
	}

	return nil
}

// NotifyCustomerInformationResponse (1.34.2)
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

// NotifyDisplayMessagesRequest (1.35.1)
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

// NotifyDisplayMessagesResponse (1.35.2)
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

// NotifyEVChargingNeedsRequest (1.36.1)
//
// The Charging Station uses this message to communicate the charging needs as calculated by the EV to the CSMS.
type NotifyEVChargingNeedsRequest struct {
	// Optional. Contains the maximum schedule tuples the car supports per schedule.
	MaxScheduleTuples *int32 `json:"maxScheduleTuples,omitempty"`
	// Required. Defines the EVSE and connector to which the EV is connected. EvseId may not be 0.
	EVSEID int32 `json:"evseId"`
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
	if err := s.ChargingNeeds.Validate(); err != nil {
		return ocpp.WrapField("chargingNeeds", err)
	}

	return nil
}

// NotifyEVChargingNeedsResponse (1.36.2)
//
// Response to a NotifyEVChargingNeedsRequest.
type NotifyEVChargingNeedsResponse struct {
	// Required. Returns whether the CSMS has been able to process the message successfully. It does not imply that the
	// evChargingNeeds can be met with the current charging profile.
	Status NotifyEVChargingNeedsStatusEnumType `json:"status"`
	// Optional. Detailed status information.
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

// NotifyEVChargingScheduleRequest (1.37.1)
//
// The Charging Station uses this message to communicate the charging schedule as calculated by the EV to the CSMS.
type NotifyEVChargingScheduleRequest struct {
	// Required. Periods contained in the charging profile are relative to this point in time.
	TimeBase time.Time `json:"timeBase"`
	// Required. The charging schedule contained in this notification applies to an EVSE. EvseId must be > 0.
	EVSEID int32 `json:"evseId"`
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

	if err := s.ChargingSchedule.Validate(); err != nil {
		return ocpp.WrapField("chargingSchedule", err)
	}

	return nil
}

// NotifyEVChargingScheduleResponse (1.37.2)
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

// NotifyEventRequest (1.38.1)
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
	// Required. List of EventData. An EventData element contains only the Component, Variable and VariableMonitoring data
	// that caused the event. The list of EventData will usally contain one eventData element, but the Charging Station
	// may decide to group multiple events in one notification. For example, when multiple events triggered at the same
	// time.
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

// NotifyEventResponse (1.38.2)
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

// NotifyMonitoringReportRequest (1.39.1)
//
// This contains the field definition of the NotifyMonitoringRequest PDU sent by the Charging Station to the CSMS.
type NotifyMonitoringReportRequest struct {
	// Required. The id of the GetMonitoringRequest that requested this report.
	RequestID int32 `json:"requestId"`
	// Optional. “to be continued” indicator. Indicates whether another part of the monitoringData follows in an upcoming
	// notifyMonitoringReportRequest message. Default value when omitted is false.
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

// NotifyMonitoringReportResponse (1.39.2)
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

// NotifyReportRequest (1.40.1)
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

	for i, v := range s.ReportData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("reportData[%d]", i), err)
		}
	}

	return nil
}

// NotifyReportResponse (1.40.2)
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

// PublishFirmwareRequest (1.41.1)
//
// This contains the field definition of the PublishFirmwareRequest PDU sent by the CSMS to the Local Controller.
type PublishFirmwareRequest struct {
	// Required. This contains a string containing a URI pointing to a location from which to retrieve the firmware.
	Location string `json:"location"`
	// Optional. This specifies how many times the Charging Station must retry to download the firmware before giving up.
	// If this field is not present, it is left to Charging Station to decide how many times it wants to retry. If the
	// value is 0, it means: no retries.
	Retries *int32 `json:"retries,omitempty"`
	// Required. The MD5 checksum over the entire firmware file as a hexadecimal string of length 32.
	Checksum IdentifierString32Type `json:"checksum"`
	// Required. The Id of the request.
	RequestID int32 `json:"requestId"`
	// Optional. The interval in seconds after which a retry may be attempted. If this field is not present, it is left to
	// Charging Station to decide how long to wait between attempts.
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

	if err := validateStringLen(s.Location, 512, "location"); err != nil {
		return err
	}

	if s.Checksum == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "checksum", "required field is missing")
	}

	if err := s.Checksum.Validate(); err != nil {
		return ocpp.WrapField("checksum", err)
	}

	return nil
}

// PublishFirmwareResponse (1.41.2)
//
// This contains the field definition of the PublishFirmwareResponse PDU sent by the Local Controller to the CSMS in
// response to a PublishFirmwareRequest.
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

// PublishFirmwareStatusNotificationRequest (1.42.1)
//
// This contains the field definition of the PublishFirmwareStatusNotificationRequest PDU sent by the Charging Station
// to the CSMS.
type PublishFirmwareStatusNotificationRequest struct {
	// Required. This contains the progress status of the publishfirmware installation.
	Status PublishFirmwareStatusEnumType `json:"status"`
	// Optional. Required if status is Published. Can be multiple URI’s, if the Local Controller supports e.g. HTTP,
	// HTTPS, and FTP.
	Location []string `json:"location,omitempty"`
	// Optional. The request id that was provided in the PublishFirmwareRequest which triggered this action.
	RequestID *int32 `json:"requestId,omitempty"`
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
		if err := validateStringLen(v, 512, fmt.Sprintf("location[%d]", i)); err != nil {
			return err
		}
	}

	return nil
}

// PublishFirmwareStatusNotificationResponse (1.42.2)
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

// ReportChargingProfilesRequest (1.43.1)
//
// Reports charging profiles installed in the Charging Station, as requested via a GetChargingProfilesRequest message.
// The charging profile report can be split over multiple ReportChargingProfilesRequest messages, this can be because
// charging profiles for different charging sources need to be reported, or because there is just to much data for one
// message.
type ReportChargingProfilesRequest struct {
	// Required. Id used to match the GetChargingProfilesRequest message with the resulting ReportChargingProfilesRequest
	// messages. When the CSMS provided a requestId in the GetChargingProfilesRequest, this field SHALL contain the same
	// value.
	RequestID int32 `json:"requestId"`
	// Required. Source that has installed this charging profile.
	ChargingLimitSource ChargingLimitSourceEnumType `json:"chargingLimitSource"`
	// Optional. To Be Continued. Default value when omitted: false. false indicates that there are no further messages as
	// part of this report.
	Tbc *bool `json:"tbc,omitempty"`
	// Required. The evse to which the charging profile applies. If evseId = 0, the message contains an overall limit for
	// the Charging Station.
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

// ReportChargingProfilesResponse (1.43.2)
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

// RequestStartTransactionRequest (1.44.1)
//
// This contains the field definitions of the RequestStartTransactionRequest PDU sent to Charging Station by CSMS.
type RequestStartTransactionRequest struct {
	// Optional. Number of the EVSE on which to start the transaction. EvseId SHALL be > 0
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. Id given by the server to this start request. The Charging Station will return this in the
	// TransactionEventRequest, letting the server know which transaction was started for this request.
	RemoteStartID int32 `json:"remoteStartId"`
	// Required. The identifier that the Charging Station must use to start a transaction.
	IDToken IdTokenType `json:"idToken"`
	// Optional. Charging Profile to be used by the Charging Station for the requested transaction. ChargingProfilePurpose
	// MUST be set to TxProfile
	ChargingProfile *ChargingProfileType `json:"chargingProfile,omitempty"`
	// Optional. The groupIdToken is only relevant when the transaction is to be started on an EVSE for which a
	// reservation for groupIdToken is active, and the configuration variable AuthorizeRemoteStart = false (otherwise the
	// AuthorizeResponse could return the groupIdToken).
	GroupIDToken *IdTokenType `json:"groupIdToken,omitempty"`
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

// RequestStartTransactionResponse (1.44.2)
//
// This contains the field definitions of the RequestStartTransactionResponse PDU sent from Charging Station to CSMS.
type RequestStartTransactionResponse struct {
	// Required. Status indicating whether the Charging Station accepts the request to start a transaction.
	Status RequestStartStopStatusEnumType `json:"status"`
	// Optional. When the transaction was already started by the Charging Station before the
	// RequestStartTransactionRequest was received, for example: cable plugged in first. This contains the transactionId
	// of the already started transaction.
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

// RequestStopTransactionRequest (1.45.1)
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

// RequestStopTransactionResponse (1.45.2)
//
// This contains the field definitions of the RequestStopTransactionResponse PDU sent from Charging Station to CSMS.
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

// ReservationStatusUpdateRequest (1.46.1)
//
// This contains the field definition of the ReservationStatusUpdateRequest PDU sent by the Charging Station to the
// CSMS.
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
	if s.ReservationUpdateStatus == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "reservationUpdateStatus", "required field is missing")
	}

	return nil
}

// ReservationStatusUpdateResponse (1.46.2)
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

// ReserveNowRequest (1.47.1)
//
// This contains the field definition of the ReserveNowRequest PDU sent by the CSMS to the Charging Station.
type ReserveNowRequest struct {
	// Required. Id of reservation.
	ID int32 `json:"id"`
	// Required. Date and time at which the reservation expires.
	ExpiryDateTime time.Time `json:"expiryDateTime"`
	// Optional. This field specifies the connector type.
	ConnectorType *ConnectorEnumType `json:"connectorType,omitempty"`
	// Optional. This contains ID of the evse to be reserved.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Required. The identifier for which the reservation is made.
	IDToken IdTokenType `json:"idToken"`
	// Optional. The group identifier for which the reservation is made.
	GroupIDToken *IdTokenType `json:"groupIdToken,omitempty"`
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
	if s.ExpiryDateTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "expiryDateTime", "required field is missing")
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

// ReserveNowResponse (1.47.2)
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

// ResetRequest (1.48.1)
//
// This contains the field definition of the ResetRequest PDU sent by the CSMS to the Charging Station.
type ResetRequest struct {
	// Required. This contains the type of reset that the Charging Station or EVSE should perform.
	Type ResetEnumType `json:"type"`
	// Optional. This contains the ID of a specific EVSE that needs to be reset, instead of the entire Charging Station.
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

	return nil
}

// ResetResponse (1.48.2)
//
// This contains the field definition of the ResetResponse PDU sent by the Charging Station to the CSMS in response to
// ResetRequest.
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

// SecurityEventNotificationRequest (1.49.1)
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

// SecurityEventNotificationResponse (1.49.2)
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

// SendLocalListRequest (1.50.1)
//
// This contains the field definition of the SendLocalListRequest PDU sent by the CSMS to the Charging Station. If no
// (empty) localAuthorizationList is given and the updateType is Full, all IdTokens are removed from the list.
// Requesting a Differential update without or with empty localAuthorizationList will have no effect on the list. All
// IdTokens in the localAuthorizationList MUST be unique, no duplicate values are allowed.
type SendLocalListRequest struct {
	// Required. In case of a full update this is the version number of the full list. In case of a differential update it
	// is the version number of the list after the update has been applied.
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

// SendLocalListResponse (1.50.2)
//
// This contains the field definition of the SendLocalListResponse PDU sent by the Charging Station to the CSMS in
// response to SendLocalListRequest PDU.
type SendLocalListResponse struct {
	// Required. This indicates whether the Charging Station has successfully received and applied the update of the Local
	// Authorization List.
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

// SetChargingProfileRequest (1.51.1)
//
// This contains the field definition of the SetChargingProfileRequest PDU sent by the CSMS to the Charging Station.
// The CSMS uses this message to send charging profiles to a Charging Station.
type SetChargingProfileRequest struct {
	// Required. For TxDefaultProfile an evseId=0 applies the profile to each individual evse. For
	// ChargingStationMaxProfile and ChargingStationExternalConstraints an evseId=0 contains an overal limit for the whole
	// Charging Station.
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
	if err := s.ChargingProfile.Validate(); err != nil {
		return ocpp.WrapField("chargingProfile", err)
	}

	return nil
}

// SetChargingProfileResponse (1.51.2)
//
// This contains the field definition of the SetChargingProfileResponse PDU sent by the Charging Station to the CSMS
// in response to SetChargingProfileRequest PDU.
type SetChargingProfileResponse struct {
	// Required. Returns whether the Charging Station has been able to process the message successfully. This does not
	// guarantee the schedule will be followed to the letter. There might be other constraints the Charging Station may
	// need to take into account.
	Status ChargingProfileStatusEnumType `json:"status"`
	// Optional. Detailed status information.
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

// SetDisplayMessageRequest (1.52.1)
//
// This contains the field definition of the SetDisplayMessageRequest PDU sent by the CSMS to the Charging Station.
// The CSMS asks the Charging Station to configure a new display message that the Charging Station will display (in
// the future). See also O01 - Set Display Message, O02 - Set Display Message for Transaction and O06 - Replace
// Display Message
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

// SetDisplayMessageResponse (1.52.2)
//
// This contains the field definition of the SetDisplayMessageResponse PDU sent by the Charging Station to the CSMS in
// a response to a SetDisplayMessageRequest. See also O01 - Set Display Message, O02 - Set Display Message for
// Transaction and O06 - Replace Display Message
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

// SetMonitoringBaseRequest (1.53.1)
//
// This contains the field definition of the SetMonitoringBaseRequest PDU sent by the CSMS to the Charging Station.
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

// SetMonitoringBaseResponse (1.53.2)
//
// This contains the field definition of the SetMonitoringBaseResponse PDU sent by the Charging Station to the CSMS in
// response to a SetMonitoringBaseRequest.
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

// SetMonitoringLevelRequest (1.54.1)
//
// This contains the field definition of the SetMonitoringLevelRequest PDU sent by the CSMS to the Charging Station.
type SetMonitoringLevelRequest struct {
	// Required. The Charging Station SHALL only report events with a severity number lower than or equal to this
	// severity. The severity range is 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels
	// have the following meaning: 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue
	// regular operations due to Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station
	// is unable to continue regular operations due to software or minor hardware issues. Action is required. 3-Critical
	// Indicates a critical error. Action is required. 4-Error Indicates a non-urgent error. Action is required. 5-Alert
	// Indicates an alert event. Default severity for any type of monitoring event. 6-Warning Indicates a warning event.
	// Action may be required. 7-Notice Indicates an unusual event. No immediate action is required. 8-Informational
	// Indicates a regular operational event. May be used for reporting, measuring throughput, etc. No action is required.
	// 9-Debug Indicates information useful to developers for debugging, not useful during operations.
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
	_ = s
	return nil
}

// SetMonitoringLevelResponse (1.54.2)
//
// This contains the field definition of the SetMonitoringLevelResponse PDU sent by the Charging Station to the CSMS
// in response to a SetMonitoringLevelRequest.
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

// SetNetworkProfileRequest (1.55.1)
//
// With this message the CSMS gains the ability to configure the connection data (e.g. CSMS URL, OCPP version, APN,
// etc) on a Charging Station.
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

// SetNetworkProfileResponse (1.55.2)
//
// This contains the field definition of the SetNetworkProfileResponse PDU sent by the Charging Station to the CSMS in
// response to a SetNetworkProfileRequest.
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

// SetVariableMonitoringRequest (1.56.1)
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

// SetVariableMonitoringResponse (1.56.2)
//
// This contains the field definition of the SetVariableMonitoringResponse PDU sent by the Charging Station to the
// CSMS in response to a SetVariableMonitoringRequest.
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

// SetVariablesRequest (1.57.1)
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

// SetVariablesResponse (1.57.2)
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

// SignCertificateRequest (1.58.1)
//
// Sent by the Charging Station to the CSMS to request that the Certificate Authority signs the public key into a
// certificate.
type SignCertificateRequest struct {
	// Required. The Charging Station SHALL send the public key in form of a Certificate Signing Request (CSR) as
	// described in RFC 2986 [22] and then PEM encoded, using the SignCertificateRequest message.
	Csr string `json:"csr"`
	// Optional. Indicates the type of certificate that is to be signed. When omitted the certificate is to be used for
	// both the 15118 connection (if implemented) and the Charging Station to CSMS connection.
	CertificateType *CertificateSigningUseEnumType `json:"certificateType,omitempty"`
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

	return nil
}

// SignCertificateResponse (1.58.2)
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

// StatusNotificationRequest (1.59.1)
//
// This contains the field definition of the StatusNotificationRequest PDU sent by the Charging Station to the CSMS.
// This message might be removed in a future version of OCPP. It will be replaced by Device Management Monitoring
// events.
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

	return nil
}

// StatusNotificationResponse (1.59.2)
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

// TransactionEventRequest (1.60.1)
//
// This section contains the field definition of the TransactionEventRequest PDU sent by the Charging Station to the
// CSMS. For each of the eventTypes; Started, Updated and Ended, the corresponding cardinality is specified.
type TransactionEventRequest struct {
	// Required. This contains the type of this event. The first TransactionEvent of a transaction SHALL contain:
	// "Started" The last TransactionEvent of a transaction SHALL contain: "Ended" All others SHALL contain: "Updated"
	EventType TransactionEventEnumType `json:"eventType"`
	// Required. The date and time at which this transaction event occurred.
	Timestamp time.Time `json:"timestamp"`
	// Required. Reason the Charging Station sends this message to the CSMS
	TriggerReason TriggerReasonEnumType `json:"triggerReason"`
	// Required. Incremental sequence number, helps with determining if all messages of a transaction have been received.
	SeqNo int32 `json:"seqNo"`
	// Optional. Indication that this transaction event happened when the Charging Station was offline. Default = false,
	// meaning: the event occurred when the Charging Station was online.
	Offline *bool `json:"offline,omitempty"`
	// Optional. If the Charging Station is able to report the number of phases used, then it SHALL provide it. When
	// omitted the CSMS may be able to determine the number of phases used via device management.
	NumberOfPhasesUsed *int32 `json:"numberOfPhasesUsed,omitempty"`
	// Optional. The maximum current of the connected cable in Ampere (A).
	CableMaxCurrent *int32 `json:"cableMaxCurrent,omitempty"`
	// Optional. This contains the Id of the reservation that terminates as a result of this transaction.
	ReservationID *int32 `json:"reservationId,omitempty"`
	// Required. Contains transaction specific information.
	TransactionInfo TransactionType `json:"transactionInfo"`
	// Optional. This contains the identifier for which a transaction is (or will be) started or stopped. Is required when
	// the EV Driver becomes authorized for this transaction and when the EV Driver ends authorization. The IdToken should
	// only be sent once in a TransactionEventRequest for every authorization (for starting or for stopping) done for this
	// transaction, so that CSMS can return the idTokenInfo in the TransactionEventResponse. idToken should not be present
	// in the TransactionEventRequest when a transaction is ended by a RequestStopTransactionRequest or a ResetRequest.
	IDToken *IdTokenType `json:"idToken,omitempty"`
	// Optional. This identifies which evse (and connector) of the Charging Station is used.
	EVSE *EVSEType `json:"evse,omitempty"`
	// Optional. This contains the relevant meter values. Depending on the EventType of this TransactionEvent the
	// following Configuration Variable is used to configure the content: Started: SampledDataTxStartedMeasurands Updated:
	// SampledDataTxUpdatedMeasurands Ended: SampledDataTxEndedMeasurands & AlignedDataTxEndedMeasurands
	MeterValue []MeterValueType `json:"meterValue,omitempty"`
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

	if err := s.TransactionInfo.Validate(); err != nil {
		return ocpp.WrapField("transactionInfo", err)
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

	for i, v := range s.MeterValue {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("meterValue[%d]", i), err)
		}
	}

	return nil
}

// TransactionEventResponse (1.60.2)
//
// This contains the field definition of the TransactionEventResponse PDU sent by the CSMS to the Charging Station in
// response to a
//
// TransactionEventRequest.
type TransactionEventResponse struct {
	// Optional. When eventType of TransactionEventRequest is Updated, then this value contains the running cost. When
	// eventType of TransactionEventRequest is Ended, then this contains the final total cost of this transaction,
	// including taxes, in the currency configured with the Configuration Variable: Currency. Absence of this value does
	// not imply that the transaction was free. To indicate a free transaction, the CSMS SHALL send a value of 0.00.
	TotalCost *float64 `json:"totalCost,omitempty"`
	// Optional. Priority from a business point of view. Default priority is 0, The range is from -9 to 9. Higher values
	// indicate a higher priority. The chargingPriority in TransactionEventResponse is temporarily, so it may not be set
	// in the IdTokenInfoType afterwards. Also the chargingPriority in TransactionEventResponse overrules the one in
	// IdTokenInfoType.
	ChargingPriority *int32 `json:"chargingPriority,omitempty"`
	// Optional. This contains information about authorization status, expiry and group id. Is required when the
	// transactionEventRequest contained an idToken.
	IDTokenInfo *IdTokenInfoType `json:"idTokenInfo,omitempty"`
	// Optional. This can contain updated personal message that can be shown to the EV Driver. This can be used to provide
	// updated tariff information .
	UpdatedPersonalMessage *MessageContentType `json:"updatedPersonalMessage,omitempty"`
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

	return nil
}

// TriggerMessageRequest (1.61.1)
//
// This contains the field definition of the TriggerMessageRequest PDU sent by the CSMS to the Charging Station.
type TriggerMessageRequest struct {
	// Required. Type of message to be triggered.
	RequestedMessage MessageTriggerEnumType `json:"requestedMessage"`
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

	if s.EVSE != nil {
		if err := s.EVSE.Validate(); err != nil {
			return ocpp.WrapField("evse", err)
		}
	}

	return nil
}

// TriggerMessageResponse (1.61.2)
//
// This contains the field definition of the TriggerMessageResponse PDU sent by the Charging Station to the CSMS in
// response to TriggerMessageResponse.
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

// UnlockConnectorRequest (1.62.1)
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
	_ = s
	return nil
}

// UnlockConnectorResponse (1.62.2)
//
// This contains the field definition of the UnlockConnectorResponse PDU sent by the Charging Station to the CSMS in
// response to an UnlockConnectorRequest.
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

// UnpublishFirmwareRequest (1.63.1)
//
// This contains the field definition of the UnpublishFirmwareRequest PDU sent by the CSMS to the Charging Station.
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

// UnpublishFirmwareResponse (1.63.2)
//
// This contains the field definition of the UnpublishFirmwareResponse PDU sent by the Charging Station to the CSMS in
// response to a UnpublishFirmwareRequest.
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

// UpdateFirmwareRequest (1.64.1)
//
// This contains the field definition of the UpdateFirmwareRequest PDU sent by the CSMS to the Charging Station.
type UpdateFirmwareRequest struct {
	// Optional. This specifies how many times the Charging Station must retry to download the firmware before giving up.
	// If this field is not present, it is left to Charging Station to decide how many times it wants to retry. If the
	// value is 0, it means: no retries.
	Retries *int32 `json:"retries,omitempty"`
	// Optional. The interval in seconds after which a retry may be attempted. If this field is not present, it is left to
	// Charging Station to decide how long to wait between attempts.
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
	if err := s.Firmware.Validate(); err != nil {
		return ocpp.WrapField("firmware", err)
	}

	return nil
}

// UpdateFirmwareResponse (1.64.2)
//
// This contains the field definition of the UpdateFirmwareResponse PDU sent by the Charging Station to the CSMS in
// response to an UpdateFirmwareRequest.
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
