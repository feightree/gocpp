package v201

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// validateStringLen checks that s does not exceed maxLen characters, as
// required by OCPP 2.0.1's "string[0..N]" datatypes (Part 2 §2.1.4).
func validateStringLen(s string, maxLen int, field string) error {
	if len(s) > maxLen {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, field, fmt.Sprintf("must not exceed %d characters", maxLen))
	}

	return nil
}

// validateSliceLen checks that len(s) falls within [min, max]. A negative
// max means unbounded.
func validateSliceLen[T any](s []T, min, max int, field string) error {
	if len(s) < min {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, field, fmt.Sprintf("must contain at least %d item(s)", min))
	}

	if max >= 0 && len(s) > max {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, field, fmt.Sprintf("must contain at most %d item(s)", max))
	}

	return nil
}

// validateNonNegative checks that v is >= 0.
func validateNonNegative[T int32 | float64](v T, field string) error {
	if v < 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, field, "must be >= 0")
	}

	return nil
}

func validateIdentifierString(s string, maxLen int, typeName string) error {
	if len(s) > maxLen {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "", fmt.Sprintf("%s exceeds max length of %d", typeName, maxLen))
	}

	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
		case strings.ContainsRune("*-_=:+|@.", r):
		default:
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "", fmt.Sprintf("%s contains characters outside the identifierString charset", typeName))
		}
	}

	return nil
}

// IdentifierString6Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 6 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString6Type string

func NewIdentifierString6Type(s string) (IdentifierString6Type, error) {
	c := IdentifierString6Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString6Type) Validate() error {
	return validateIdentifierString(string(s), 6, "IdentifierString6Type")
}

func (s IdentifierString6Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString6Type) String() string {
	return string(s)
}

func (s *IdentifierString6Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString6Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// IdentifierString20Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 20 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString20Type string

func NewIdentifierString20Type(s string) (IdentifierString20Type, error) {
	c := IdentifierString20Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString20Type) Validate() error {
	return validateIdentifierString(string(s), 20, "IdentifierString20Type")
}

func (s IdentifierString20Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString20Type) String() string {
	return string(s)
}

func (s *IdentifierString20Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString20Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// IdentifierString32Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 32 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString32Type string

func NewIdentifierString32Type(s string) (IdentifierString32Type, error) {
	c := IdentifierString32Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString32Type) Validate() error {
	return validateIdentifierString(string(s), 32, "IdentifierString32Type")
}

func (s IdentifierString32Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString32Type) String() string {
	return string(s)
}

func (s *IdentifierString32Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString32Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// IdentifierString36Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 36 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString36Type string

func NewIdentifierString36Type(s string) (IdentifierString36Type, error) {
	c := IdentifierString36Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString36Type) Validate() error {
	return validateIdentifierString(string(s), 36, "IdentifierString36Type")
}

func (s IdentifierString36Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString36Type) String() string {
	return string(s)
}

func (s *IdentifierString36Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString36Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// IdentifierString40Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 40 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString40Type string

func NewIdentifierString40Type(s string) (IdentifierString40Type, error) {
	c := IdentifierString40Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString40Type) Validate() error {
	return validateIdentifierString(string(s), 40, "IdentifierString40Type")
}

func (s IdentifierString40Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString40Type) String() string {
	return string(s)
}

func (s *IdentifierString40Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString40Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// IdentifierString50Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 50 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString50Type string

func NewIdentifierString50Type(s string) (IdentifierString50Type, error) {
	c := IdentifierString50Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString50Type) Validate() error {
	return validateIdentifierString(string(s), 50, "IdentifierString50Type")
}

func (s IdentifierString50Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString50Type) String() string {
	return string(s)
}

func (s *IdentifierString50Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString50Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// IdentifierString128Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 128 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString128Type string

func NewIdentifierString128Type(s string) (IdentifierString128Type, error) {
	c := IdentifierString128Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString128Type) Validate() error {
	return validateIdentifierString(string(s), 128, "IdentifierString128Type")
}

func (s IdentifierString128Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString128Type) String() string {
	return string(s)
}

func (s *IdentifierString128Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString128Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// ACChargingParametersType (2.1)
//
// EV AC charging parameters.
type ACChargingParametersType struct {
	// Required. Amount of energy requested (in Wh). This includes energy required for preconditioning.
	EnergyAmount int32 `json:"energyAmount"`
	// Required. Minimum current (amps) supported by the electric vehicle (per phase).
	EVMinCurrent int32 `json:"evMinCurrent"`
	// Required. Maximum current (amps) supported by the electric vehicle (per phase). Includes cable capacity.
	EVMaxCurrent int32 `json:"evMaxCurrent"`
	// Required. Maximum voltage supported by the electric vehicle
	EVMaxVoltage int32 `json:"evMaxVoltage"`
}

func (s *ACChargingParametersType) UnmarshalJSON(data []byte) error {
	type Alias ACChargingParametersType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ACChargingParametersType(a)
	return s.Validate()
}

func (s ACChargingParametersType) Validate() error {
	_ = s
	return nil
}

// AdditionalInfoType (2.2)
//
// Contains a case insensitive identifier to use for the authorization and the type of authorization to support
// multiple forms of identifiers.
type AdditionalInfoType struct {
	// Required. This field specifies the additional IdToken.
	AdditionalIDToken IdentifierString36Type `json:"additionalIdToken"`
	// Required. This defines the type of the additionalIdToken. This is a custom type, so the implementation needs to be
	// agreed upon by all involved parties.
	Type string `json:"type"`
}

func (s *AdditionalInfoType) UnmarshalJSON(data []byte) error {
	type Alias AdditionalInfoType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AdditionalInfoType(a)
	return s.Validate()
}

func (s AdditionalInfoType) Validate() error {
	if s.AdditionalIDToken == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "additionalIdToken", "required field is missing")
	}

	if err := s.AdditionalIDToken.Validate(); err != nil {
		return ocpp.WrapField("additionalIdToken", err)
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateStringLen(s.Type, 50, "type"); err != nil {
		return err
	}

	return nil
}

// APNType (2.3)
//
// Collection of configuration data needed to make a data-connection over a cellular network.
//
// When asking a GSM modem to dial in, it is possible to specify which mobile operator should be used. This can be
// done with the mobile country code (MCC) in combination with a mobile network code (MNC). Example: If your preferred
// network is Vodafone Netherlands, the MCC=204 and the MNC=04 which means the key NOTE PreferredNetwork = 20404 Some
// modems allows to specify a preferred network, which means, if this network is not available, a different network is
// used. If you specify UseOnlyPreferredNetwork and this network is not available, the modem will not dial in.
type APNType struct {
	// Required. The Access Point Name as an URL.
	APN string `json:"apn"`
	// Optional. APN username.
	APNUserName *string `json:"apnUserName,omitempty"`
	// Optional. APN Password.
	APNPassword *string `json:"apnPassword,omitempty"`
	// Optional. SIM card pin code.
	SimPin *int32 `json:"simPin,omitempty"`
	// Optional. Preferred network, written as MCC and MNC concatenated. See note.
	PreferredNetwork *IdentifierString6Type `json:"preferredNetwork,omitempty"`
	// Optional. Default: false. Use only the preferred Network, do not dial in when not available. See Note.
	UseOnlyPreferredNetwork *bool `json:"useOnlyPreferredNetwork,omitempty"`
	// Required. Authentication method.
	APNAuthentication APNAuthenticationEnumType `json:"apnAuthentication"`
}

func (s *APNType) UnmarshalJSON(data []byte) error {
	type Alias APNType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = APNType(a)
	return s.Validate()
}

func (s APNType) Validate() error {
	if s.APN == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "apn", "required field is missing")
	}

	if err := validateStringLen(s.APN, 512, "apn"); err != nil {
		return err
	}

	if s.APNUserName != nil {
		if err := validateStringLen(*s.APNUserName, 20, "apnUserName"); err != nil {
			return err
		}
	}

	if s.APNPassword != nil {
		if err := validateStringLen(*s.APNPassword, 20, "apnPassword"); err != nil {
			return err
		}
	}

	if s.PreferredNetwork != nil {
		if err := s.PreferredNetwork.Validate(); err != nil {
			return ocpp.WrapField("preferredNetwork", err)
		}
	}

	if s.APNAuthentication == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "apnAuthentication", "required field is missing")
	}

	return nil
}

// AuthorizationData (2.4)
//
// Contains the identifier to use for authorization.
type AuthorizationData struct {
	// Optional. Required when UpdateType is Full. This contains information about authorization status, expiry and group
	// id. For a Differential update the following applies: If this element is present, then this entry SHALL be added or
	// updated in the Local Authorization List. If this element is absent, the entry for this IdToken in the Local
	// Authorization List SHALL be deleted.
	IDTokenInfo *IdTokenInfoType `json:"idTokenInfo,omitempty"`
	// Required. This contains the identifier which needs to be stored for authorization.
	IDToken IdTokenType `json:"idToken"`
}

func (s *AuthorizationData) UnmarshalJSON(data []byte) error {
	type Alias AuthorizationData
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AuthorizationData(a)
	return s.Validate()
}

func (s AuthorizationData) Validate() error {
	if s.IDTokenInfo != nil {
		if err := s.IDTokenInfo.Validate(); err != nil {
			return ocpp.WrapField("idTokenInfo", err)
		}
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	return nil
}

// CertificateHashDataChainType (2.5)
type CertificateHashDataChainType struct {
	// Required. Indicates the type of the requested certificate(s).
	CertificateType GetCertificateIdUseEnumType `json:"certificateType"`
	// Required. Information to identify a certificate.
	CertificateHashData CertificateHashDataType `json:"certificateHashData"`
	// Optional. Information to identify the child certificate(s).
	ChildCertificateHashData []CertificateHashDataType `json:"childCertificateHashData,omitempty"`
}

func (s *CertificateHashDataChainType) UnmarshalJSON(data []byte) error {
	type Alias CertificateHashDataChainType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CertificateHashDataChainType(a)
	return s.Validate()
}

func (s CertificateHashDataChainType) Validate() error {
	if s.CertificateType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "certificateType", "required field is missing")
	}

	if err := s.CertificateHashData.Validate(); err != nil {
		return ocpp.WrapField("certificateHashData", err)
	}

	if err := validateSliceLen(s.ChildCertificateHashData, 0, 4, "childCertificateHashData"); err != nil {
		return err
	}

	for i, v := range s.ChildCertificateHashData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("childCertificateHashData[%d]", i), err)
		}
	}

	return nil
}

// CertificateHashDataType (2.6)
type CertificateHashDataType struct {
	// Required. Used algorithms for the hashes provided.
	HashAlgorithm HashAlgorithmEnumType `json:"hashAlgorithm"`
	// Required. The hash of the issuer’s distinguished name (DN), that must be calculated over the DER encoding of the
	// issuer’s name field in the certificate being checked.
	IssuerNameHash IdentifierString128Type `json:"issuerNameHash"`
	// Required. The hash of the DER encoded public key: the value (excluding tag and length) of the subject public key
	// field in the issuer’s certificate.
	IssuerKeyHash string `json:"issuerKeyHash"`
	// Required. The string representation of the hexadecimal value of the serial number without the prefix "0x" and
	// without leading zeroes.
	SerialNumber IdentifierString40Type `json:"serialNumber"`
}

func (s *CertificateHashDataType) UnmarshalJSON(data []byte) error {
	type Alias CertificateHashDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CertificateHashDataType(a)
	return s.Validate()
}

func (s CertificateHashDataType) Validate() error {
	if s.HashAlgorithm == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "hashAlgorithm", "required field is missing")
	}

	if s.IssuerNameHash == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "issuerNameHash", "required field is missing")
	}

	if err := s.IssuerNameHash.Validate(); err != nil {
		return ocpp.WrapField("issuerNameHash", err)
	}

	if s.IssuerKeyHash == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "issuerKeyHash", "required field is missing")
	}

	if err := validateStringLen(s.IssuerKeyHash, 128, "issuerKeyHash"); err != nil {
		return err
	}

	if s.SerialNumber == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "serialNumber", "required field is missing")
	}

	if err := s.SerialNumber.Validate(); err != nil {
		return ocpp.WrapField("serialNumber", err)
	}

	return nil
}

// ChargingLimitType (2.7)
type ChargingLimitType struct {
	// Required. Represents the source of the charging limit.
	ChargingLimitSource ChargingLimitSourceEnumType `json:"chargingLimitSource"`
	// Optional. Indicates whether the charging limit is critical for the grid.
	IsGridCritical *bool `json:"isGridCritical,omitempty"`
}

func (s *ChargingLimitType) UnmarshalJSON(data []byte) error {
	type Alias ChargingLimitType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingLimitType(a)
	return s.Validate()
}

func (s ChargingLimitType) Validate() error {
	if s.ChargingLimitSource == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingLimitSource", "required field is missing")
	}

	return nil
}

// ChargingNeedsType (2.8)
type ChargingNeedsType struct {
	// Required. Mode of energy transfer requested by the EV.
	RequestedEnergyTransfer EnergyTransferModeEnumType `json:"requestedEnergyTransfer"`
	// Optional. Estimated departure time of the EV.
	DepartureTime *time.Time `json:"departureTime,omitempty"`
	// Optional. EV AC charging parameters.
	AcChargingParameters *ACChargingParametersType `json:"acChargingParameters,omitempty"`
	// Optional. EV DC charging parameters
	DcChargingParameters *DCChargingParametersType `json:"dcChargingParameters,omitempty"`
}

func (s *ChargingNeedsType) UnmarshalJSON(data []byte) error {
	type Alias ChargingNeedsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingNeedsType(a)
	return s.Validate()
}

func (s ChargingNeedsType) Validate() error {
	if s.RequestedEnergyTransfer == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "requestedEnergyTransfer", "required field is missing")
	}

	if s.AcChargingParameters != nil {
		if err := s.AcChargingParameters.Validate(); err != nil {
			return ocpp.WrapField("acChargingParameters", err)
		}
	}

	if s.DcChargingParameters != nil {
		if err := s.DcChargingParameters.Validate(); err != nil {
			return ocpp.WrapField("dcChargingParameters", err)
		}
	}

	return nil
}

// ChargingProfileCriterionType (2.9)
//
// A ChargingProfile consists of ChargingSchedule, describing the amount of power or current that can be delivered per
// time interval.
type ChargingProfileCriterionType struct {
	// Optional. Defines the purpose of the schedule transferred by this profile
	ChargingProfilePurpose *ChargingProfilePurposeEnumType `json:"chargingProfilePurpose,omitempty"`
	// Optional. Value determining level in hierarchy stack of profiles. Higher values have precedence over lower values.
	// Lowest level is 0.
	StackLevel *int32 `json:"stackLevel,omitempty"`
	// Optional. List of all the chargingProfileIds requested. Any ChargingProfile that matches one of these profiles will
	// be reported. If omitted, the Charging Station SHALL not filter on chargingProfileId. This field SHALL NOT contain
	// more ids than set in ChargingProfileEntries.maxLimit
	ChargingProfileID []int32 `json:"chargingProfileId,omitempty"`
	// Optional. For which charging limit sources, charging profiles SHALL be reported. If omitted, the Charging Station
	// SHALL not filter on chargingLimitSource.
	ChargingLimitSource []ChargingLimitSourceEnumType `json:"chargingLimitSource,omitempty"`
}

func (s *ChargingProfileCriterionType) UnmarshalJSON(data []byte) error {
	type Alias ChargingProfileCriterionType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingProfileCriterionType(a)
	return s.Validate()
}

func (s ChargingProfileCriterionType) Validate() error {
	if err := validateSliceLen(s.ChargingLimitSource, 0, 4, "chargingLimitSource"); err != nil {
		return err
	}

	return nil
}

// ChargingProfileType (2.10)
//
// A ChargingProfile consists of ChargingSchedule, describing the amount of power or current that can be delivered per
// time interval.
type ChargingProfileType struct {
	// Required. Id of ChargingProfile.
	ID int32 `json:"id"`
	// Required. Value determining level in hierarchy stack of profiles. Higher values have precedence over lower values.
	// Lowest level is 0.
	StackLevel int32 `json:"stackLevel"`
	// Required. Defines the purpose of the schedule transferred by this profile
	ChargingProfilePurpose ChargingProfilePurposeEnumType `json:"chargingProfilePurpose"`
	// Required. Indicates the kind of schedule.
	ChargingProfileKind ChargingProfileKindEnumType `json:"chargingProfileKind"`
	// Optional. Indicates the start point of a recurrence.
	RecurrencyKind *RecurrencyKindEnumType `json:"recurrencyKind,omitempty"`
	// Optional. Point in time at which the profile starts to be valid. If absent, the profile is valid as soon as it is
	// received by the Charging Station.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Optional. Point in time at which the profile stops to be valid. If absent, the profile is valid until it is
	// replaced by another profile.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// Optional. SHALL only be included when ChargingProfilePurpose is set to TxProfile in a SetChargingProfileRequest.
	// The transactionId is used to match the profile to a specific transaction.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Required. Schedule that contains limits for the available power or current over time. In order to support ISO 15118
	// schedule negotiation, it supports at most three schedules with associated tariff to choose from. Having multiple
	// chargingSchedules is only allowed for charging profiles of purpose TxProfile in the context of an ISO 15118
	// charging session.
	ChargingSchedule []ChargingScheduleType `json:"chargingSchedule"`
}

func (s *ChargingProfileType) UnmarshalJSON(data []byte) error {
	type Alias ChargingProfileType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingProfileType(a)
	return s.Validate()
}

func (s ChargingProfileType) Validate() error {
	if s.ChargingProfilePurpose == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingProfilePurpose", "required field is missing")
	}

	if s.ChargingProfileKind == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingProfileKind", "required field is missing")
	}

	if s.TransactionID != nil {
		if err := s.TransactionID.Validate(); err != nil {
			return ocpp.WrapField("transactionId", err)
		}
	}

	if err := validateSliceLen(s.ChargingSchedule, 1, 3, "chargingSchedule"); err != nil {
		return err
	}

	for i, v := range s.ChargingSchedule {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingSchedule[%d]", i), err)
		}
	}

	return nil
}

// ChargingSchedulePeriodType (2.11)
//
// Charging schedule period structure defines a time period in a charging schedule.
type ChargingSchedulePeriodType struct {
	// Required. Start of the period, in seconds from the start of schedule. The value of StartPeriod also defines the
	// stop time of the previous period.
	StartPeriod int32 `json:"startPeriod"`
	// Required. Charging rate limit during the schedule period, in the applicable chargingRateUnit, for example in
	// Amperes (A) or Watts (W). Accepts at most one digit fraction (e.g. 8.1).
	Limit float64 `json:"limit"`
	// Optional. The number of phases that can be used for charging. For a DC EVSE this field should be omitted. For an AC
	// EVSE a default value of numberPhases = 3 will be assumed if the field is absent.
	NumberPhases *int32 `json:"numberPhases,omitempty"`
	// Optional. Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching the phase connected to the
	// EV, i.e. ACPhaseSwitchingSupported is defined and true. It’s not allowed unless both conditions above are true. If
	// both conditions are true, and phaseToUse is omitted, the Charging Station / EVSE will make the selection on its
	// own.
	PhaseToUse *int32 `json:"phaseToUse,omitempty"`
}

func (s *ChargingSchedulePeriodType) UnmarshalJSON(data []byte) error {
	type Alias ChargingSchedulePeriodType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingSchedulePeriodType(a)
	return s.Validate()
}

func (s ChargingSchedulePeriodType) Validate() error {
	_ = s
	return nil
}

// ChargingScheduleType (2.12)
//
// Charging schedule structure defines a list of charging periods, as used in: GetCompositeSchedule.conf and
// ChargingProfile.
type ChargingScheduleType struct {
	// Required. Identifies the ChargingSchedule.
	ID int32 `json:"id"`
	// Optional. Starting point of an absolute or recurring schedule.
	StartSchedule *time.Time `json:"startSchedule,omitempty"`
	// Optional. Duration of the charging schedule in seconds. If the duration is left empty, the last period will
	// continue indefinitely or until end of the transaction if chargingProfilePurpose = TxProfile.
	Duration *int32 `json:"duration,omitempty"`
	// Required. The unit of measure Limit is expressed in.
	ChargingRateUnit ChargingRateUnitEnumType `json:"chargingRateUnit"`
	// Optional. Minimum charging rate supported by the EV. The unit of measure is defined by the chargingRateUnit. This
	// parameter is intended to be used by a local smart charging algorithm to optimize the power allocation for in the
	// case a charging process is inefficient at lower charging rates. Accepts at most one digit fraction (e.g. 8.1)
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`
	// Required. List of ChargingSchedulePeriod elements defining maximum power or current usage over time. The maximum
	// number of periods, that is supported by the Charging Station, if less than 1024, is set by device model variable
	// SmartChargingCtrlr.PeriodsPerSchedule.
	ChargingSchedulePeriod []ChargingSchedulePeriodType `json:"chargingSchedulePeriod"`
	// Optional. Sales tariff associated with this charging schedule.
	SalesTariff *SalesTariffType `json:"salesTariff,omitempty"`
}

func (s *ChargingScheduleType) UnmarshalJSON(data []byte) error {
	type Alias ChargingScheduleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingScheduleType(a)
	return s.Validate()
}

func (s ChargingScheduleType) Validate() error {
	if s.ChargingRateUnit == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingRateUnit", "required field is missing")
	}

	if err := validateSliceLen(s.ChargingSchedulePeriod, 1, 1024, "chargingSchedulePeriod"); err != nil {
		return err
	}

	for i, v := range s.ChargingSchedulePeriod {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingSchedulePeriod[%d]", i), err)
		}
	}

	if s.SalesTariff != nil {
		if err := s.SalesTariff.Validate(); err != nil {
			return ocpp.WrapField("salesTariff", err)
		}
	}

	return nil
}

// ChargingStationType (2.13)
//
// The physical system where an Electrical Vehicle (EV) can be charged.
type ChargingStationType struct {
	// Optional. Vendor-specific device identifier.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Required. Defines the model of the device.
	Model string `json:"model"`
	// Required. Identifies the vendor (not necessarily in a unique manner).
	VendorName string `json:"vendorName"`
	// Optional. This contains the firmware version of the Charging Station.
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	// Optional. Defines the functional parameters of a communication link.
	Modem *ModemType `json:"modem,omitempty"`
}

func (s *ChargingStationType) UnmarshalJSON(data []byte) error {
	type Alias ChargingStationType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingStationType(a)
	return s.Validate()
}

func (s ChargingStationType) Validate() error {
	if s.SerialNumber != nil {
		if err := validateStringLen(*s.SerialNumber, 25, "serialNumber"); err != nil {
			return err
		}
	}

	if s.Model == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "model", "required field is missing")
	}

	if err := validateStringLen(s.Model, 20, "model"); err != nil {
		return err
	}

	if s.VendorName == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "vendorName", "required field is missing")
	}

	if err := validateStringLen(s.VendorName, 50, "vendorName"); err != nil {
		return err
	}

	if s.FirmwareVersion != nil {
		if err := validateStringLen(*s.FirmwareVersion, 50, "firmwareVersion"); err != nil {
			return err
		}
	}

	if s.Modem != nil {
		if err := s.Modem.Validate(); err != nil {
			return ocpp.WrapField("modem", err)
		}
	}

	return nil
}

// ClearChargingProfileType (2.14)
//
// A ChargingProfile consists of a ChargingSchedule, describing the amount of power or current that can be delivered
// per time interval.
type ClearChargingProfileType struct {
	// Optional. Specifies the id of the EVSE for which to clear charging profiles. An evseId of zero (0) specifies the
	// charging profile for the overall Charging Station. Absence of this parameter means the clearing applies to all
	// charging profiles that match the other criteria in the request.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Optional. Specifies to purpose of the charging profiles that will be cleared, if they meet the other criteria in
	// the request.
	ChargingProfilePurpose *ChargingProfilePurposeEnumType `json:"chargingProfilePurpose,omitempty"`
	// Optional. Specifies the stackLevel for which charging profiles will be cleared, if they meet the other criteria in
	// the request.
	StackLevel *int32 `json:"stackLevel,omitempty"`
}

func (s *ClearChargingProfileType) UnmarshalJSON(data []byte) error {
	type Alias ClearChargingProfileType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearChargingProfileType(a)
	return s.Validate()
}

func (s ClearChargingProfileType) Validate() error {
	_ = s
	return nil
}

// ClearMonitoringResultType (2.15)
type ClearMonitoringResultType struct {
	// Required. Result of the clear request for this monitor, identified by its Id.
	Status ClearMonitoringStatusEnumType `json:"status"`
	// Required. Id of the monitor of which a clear was requested.
	ID int32 `json:"id"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ClearMonitoringResultType) UnmarshalJSON(data []byte) error {
	type Alias ClearMonitoringResultType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearMonitoringResultType(a)
	return s.Validate()
}

func (s ClearMonitoringResultType) Validate() error {
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

// ComponentType (2.16)
//
// A physical or logical component
type ComponentType struct {
	// Required. Name of the component. Name should be taken from the list of standardized component names whenever
	// possible. Case Insensitive. strongly advised to use Camel Case.
	Name IdentifierString50Type `json:"name"`
	// Optional. Name of instance in case the component exists as multiple instances. Case Insensitive. strongly advised
	// to use Camel Case.
	Instance *IdentifierString50Type `json:"instance,omitempty"`
	// Optional. Specifies the EVSE when component is located at EVSE level, also specifies the connector when component
	// is located at Connector level.
	EVSE *EVSEType `json:"evse,omitempty"`
}

func (s *ComponentType) UnmarshalJSON(data []byte) error {
	type Alias ComponentType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ComponentType(a)
	return s.Validate()
}

func (s ComponentType) Validate() error {
	if s.Name == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "name", "required field is missing")
	}

	if err := s.Name.Validate(); err != nil {
		return ocpp.WrapField("name", err)
	}

	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return ocpp.WrapField("instance", err)
		}
	}

	if s.EVSE != nil {
		if err := s.EVSE.Validate(); err != nil {
			return ocpp.WrapField("evse", err)
		}
	}

	return nil
}

// ComponentVariableType (2.17)
//
// Class to report components, variables and variable attributes and characteristics.
type ComponentVariableType struct {
	// Required. Component for which a report of Variable is requested.
	Component ComponentType `json:"component"`
	// Optional. Variable(s) for which the report is requested.
	Variable *VariableType `json:"variable,omitempty"`
}

func (s *ComponentVariableType) UnmarshalJSON(data []byte) error {
	type Alias ComponentVariableType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ComponentVariableType(a)
	return s.Validate()
}

func (s ComponentVariableType) Validate() error {
	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if s.Variable != nil {
		if err := s.Variable.Validate(); err != nil {
			return ocpp.WrapField("variable", err)
		}
	}

	return nil
}

// CompositeScheduleType (2.18)
type CompositeScheduleType struct {
	// Required. The ID of the EVSE for which the schedule is requested. When evseid=0, the Charging Station calculated
	// the expected consumption for the grid connection.
	EVSEID int32 `json:"evseId"`
	// Required. Duration of the schedule in seconds.
	Duration int32 `json:"duration"`
	// Required. Date and time at which the schedule becomes active. All time measurements within the schedule are
	// relative to this timestamp.
	ScheduleStart time.Time `json:"scheduleStart"`
	// Required. The unit of measure Limit is expressed in.
	ChargingRateUnit ChargingRateUnitEnumType `json:"chargingRateUnit"`
	// Required. List of ChargingSchedulePeriod elements defining maximum power or current usage over time.
	ChargingSchedulePeriod []ChargingSchedulePeriodType `json:"chargingSchedulePeriod"`
}

func (s *CompositeScheduleType) UnmarshalJSON(data []byte) error {
	type Alias CompositeScheduleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CompositeScheduleType(a)
	return s.Validate()
}

func (s CompositeScheduleType) Validate() error {
	if s.ScheduleStart.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "scheduleStart", "required field is missing")
	}

	if s.ChargingRateUnit == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "chargingRateUnit", "required field is missing")
	}

	if err := validateSliceLen(s.ChargingSchedulePeriod, 1, -1, "chargingSchedulePeriod"); err != nil {
		return err
	}

	for i, v := range s.ChargingSchedulePeriod {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingSchedulePeriod[%d]", i), err)
		}
	}

	return nil
}

// ConsumptionCostType (2.19)
type ConsumptionCostType struct {
	// Required. The lowest level of consumption that defines the starting point of this consumption block. The block
	// interval extends to the start of the next interval.
	StartValue float64 `json:"startValue"`
	// Required. This field contains the cost details.
	Cost []CostType `json:"cost"`
}

func (s *ConsumptionCostType) UnmarshalJSON(data []byte) error {
	type Alias ConsumptionCostType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ConsumptionCostType(a)
	return s.Validate()
}

func (s ConsumptionCostType) Validate() error {
	if err := validateSliceLen(s.Cost, 1, 3, "cost"); err != nil {
		return err
	}

	for i, v := range s.Cost {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("cost[%d]", i), err)
		}
	}

	return nil
}

// CostType (2.20)
type CostType struct {
	// Required. The kind of cost referred to in the message element amount
	CostKind CostKindEnumType `json:"costKind"`
	// Required. The estimated or actual cost per kWh
	Amount int32 `json:"amount"`
	// Optional. Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The final value is determined
	// by: amount * 10 ^ amountMultiplier
	AmountMultiplier *int32 `json:"amountMultiplier,omitempty"`
}

func (s *CostType) UnmarshalJSON(data []byte) error {
	type Alias CostType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CostType(a)
	return s.Validate()
}

func (s CostType) Validate() error {
	if s.CostKind == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "costKind", "required field is missing")
	}

	return nil
}

// DCChargingParametersType (2.21)
//
// EV DC charging parameters
type DCChargingParametersType struct {
	// Required. Maximum current (amps) supported by the electric vehicle. Includes cable capacity.
	EVMaxCurrent int32 `json:"evMaxCurrent"`
	// Required. Maximum voltage supported by the electric vehicle
	EVMaxVoltage int32 `json:"evMaxVoltage"`
	// Optional. Amount of energy requested (in Wh). This inludes energy required for preconditioning.
	EnergyAmount *int32 `json:"energyAmount,omitempty"`
	// Optional. Maximum power (in W) supported by the electric vehicle. Required for DC charging.
	EVMaxPower *int32 `json:"evMaxPower,omitempty"`
	// Optional. Energy available in the battery (in percent of the battery capacity)
	StateOfCharge *int32 `json:"stateOfCharge,omitempty"`
	// Optional. Capacity of the electric vehicle battery (in Wh)
	EVEnergyCapacity *int32 `json:"evEnergyCapacity,omitempty"`
	// Optional. Percentage of SoC at which the EV considers the battery fully charged. (possible values: 0 - 100)
	FullSoC *int32 `json:"fullSoC,omitempty"`
	// Optional. Percentage of SoC at which the EV considers a fast charging process to end. (possible values: 0 - 100)
	BulkSoC *int32 `json:"bulkSoC,omitempty"`
}

func (s *DCChargingParametersType) UnmarshalJSON(data []byte) error {
	type Alias DCChargingParametersType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DCChargingParametersType(a)
	return s.Validate()
}

func (s DCChargingParametersType) Validate() error {
	if s.StateOfCharge != nil {
		if *s.StateOfCharge < 0 || *s.StateOfCharge > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "stateOfCharge", "must be between 0 and 100")
		}
	}

	if s.FullSoC != nil {
		if *s.FullSoC < 0 || *s.FullSoC > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "fullSoC", "must be between 0 and 100")
		}
	}

	if s.BulkSoC != nil {
		if *s.BulkSoC < 0 || *s.BulkSoC > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "bulkSoC", "must be between 0 and 100")
		}
	}

	return nil
}

// EventDataType (2.22)
//
// Class to report an event notification for a component-variable.
type EventDataType struct {
	// Required. Identifies the event. This field can be referred to as a cause by other events.
	EventID int32 `json:"eventId"`
	// Required. Timestamp of the moment the report was generated.
	Timestamp time.Time `json:"timestamp"`
	// Required. Type of trigger for this event, e.g. exceeding a threshold value.
	Trigger EventTriggerEnumType `json:"trigger"`
	// Optional. Refers to the Id of an event that is considered to be the cause for this event.
	Cause *int32 `json:"cause,omitempty"`
	// Required. Actual value (attributeType Actual) of the variable. The Configuration Variable ReportingValueSize can be
	// used to limit GetVariableResult.attributeValue, VariableAttribute.value and EventData.actualValue. The max size of
	// these values will always remain equal.
	ActualValue string `json:"actualValue"`
	// Optional. Technical (error) code as reported by component.
	TechCode *string `json:"techCode,omitempty"`
	// Optional. Technical detail information as reported by component.
	TechInfo *string `json:"techInfo,omitempty"`
	// Optional. Cleared is set to true to report the clearing of a monitored situation, i.e. a 'return to normal'.
	Cleared *bool `json:"cleared,omitempty"`
	// Optional. If an event notification is linked to a specific transaction, this field can be used to specify its
	// transactionId.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Optional. Identifies the VariableMonitoring which triggered the event.
	VariableMonitoringID *int32 `json:"variableMonitoringId,omitempty"`
	// Required. Specifies the event notification type of the message.
	EventNotificationType EventNotificationEnumType `json:"eventNotificationType"`
	// Required. Component for which event is notified.
	Component ComponentType `json:"component"`
	// Required. Variable for which event is notified.
	Variable VariableType `json:"variable"`
}

func (s *EventDataType) UnmarshalJSON(data []byte) error {
	type Alias EventDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EventDataType(a)
	return s.Validate()
}

func (s EventDataType) Validate() error {
	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.Trigger == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "trigger", "required field is missing")
	}

	if s.ActualValue == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "actualValue", "required field is missing")
	}

	if err := validateStringLen(s.ActualValue, 2500, "actualValue"); err != nil {
		return err
	}

	if s.TechCode != nil {
		if err := validateStringLen(*s.TechCode, 50, "techCode"); err != nil {
			return err
		}
	}

	if s.TechInfo != nil {
		if err := validateStringLen(*s.TechInfo, 500, "techInfo"); err != nil {
			return err
		}
	}

	if s.TransactionID != nil {
		if err := s.TransactionID.Validate(); err != nil {
			return ocpp.WrapField("transactionId", err)
		}
	}

	if s.EventNotificationType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "eventNotificationType", "required field is missing")
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	return nil
}

// EVSEType (2.23)
//
// Electric Vehicle Supply Equipment
type EVSEType struct {
	// Required. EVSE Identifier. This contains a number (> 0) designating an EVSE of the Charging Station.
	ID int32 `json:"id"`
	// Optional. An id to designate a specific connector (on an EVSE) by connector index number.
	ConnectorID *int32 `json:"connectorId,omitempty"`
}

func (s *EVSEType) UnmarshalJSON(data []byte) error {
	type Alias EVSEType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVSEType(a)
	return s.Validate()
}

func (s EVSEType) Validate() error {
	_ = s
	return nil
}

// FirmwareType (2.24)
//
// Represents a copy of the firmware that can be loaded/updated on the Charging Station.
type FirmwareType struct {
	// Required. URI defining the origin of the firmware.
	Location string `json:"location"`
	// Required. Date and time at which the firmware shall be retrieved.
	RetrieveDateTime time.Time `json:"retrieveDateTime"`
	// Optional. Date and time at which the firmware shall be installed.
	InstallDateTime *time.Time `json:"installDateTime,omitempty"`
	// Optional. Certificate with which the firmware was signed. PEM encoded X.509 certificate.
	SigningCertificate *string `json:"signingCertificate,omitempty"`
	// Optional. Base64 encoded firmware signature.
	Signature *string `json:"signature,omitempty"`
}

func (s *FirmwareType) UnmarshalJSON(data []byte) error {
	type Alias FirmwareType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FirmwareType(a)
	return s.Validate()
}

func (s FirmwareType) Validate() error {
	if s.Location == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "location", "required field is missing")
	}

	if err := validateStringLen(s.Location, 512, "location"); err != nil {
		return err
	}

	if s.RetrieveDateTime.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "retrieveDateTime", "required field is missing")
	}

	if s.SigningCertificate != nil {
		if err := validateStringLen(*s.SigningCertificate, 5500, "signingCertificate"); err != nil {
			return err
		}
	}

	if s.Signature != nil {
		if err := validateStringLen(*s.Signature, 800, "signature"); err != nil {
			return err
		}
	}

	return nil
}

// GetVariableDataType (2.25)
//
// Class to hold parameters for GetVariables request.
type GetVariableDataType struct {
	// Optional. Attribute type for which value is requested. When absent, default Actual is assumed.
	AttributeType *AttributeEnumType `json:"attributeType,omitempty"`
	// Required. Component for which the Variable is requested.
	Component ComponentType `json:"component"`
	// Required. Variable for which the attribute value is requested.
	Variable VariableType `json:"variable"`
}

func (s *GetVariableDataType) UnmarshalJSON(data []byte) error {
	type Alias GetVariableDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetVariableDataType(a)
	return s.Validate()
}

func (s GetVariableDataType) Validate() error {
	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	return nil
}

// GetVariableResultType (2.26)
//
// Class to hold results of GetVariables request.
type GetVariableResultType struct {
	// Required. Result status of getting the variable.
	AttributeStatus GetVariableStatusEnumType `json:"attributeStatus"`
	// Optional. Attribute type for which value is requested. When absent, default Actual is assumed.
	AttributeType *AttributeEnumType `json:"attributeType,omitempty"`
	// Optional. Value of requested attribute type of component- variable. This field can only be empty when the given
	// status is NOT accepted. The Configuration Variable ReportingValueSize can be used to limit
	// GetVariableResult.attributeValue, VariableAttribute.value and EventData.actualValue. The max size of these values
	// will always remain equal.
	AttributeValue *string `json:"attributeValue,omitempty"`
	// Required. Component for which the Variable is requested.
	Component ComponentType `json:"component"`
	// Required. Variable for which the attribute value is requested.
	Variable VariableType `json:"variable"`
	// Optional. Detailed attribute status information.
	AttributeStatusInfo *StatusInfoType `json:"attributeStatusInfo,omitempty"`
}

func (s *GetVariableResultType) UnmarshalJSON(data []byte) error {
	type Alias GetVariableResultType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GetVariableResultType(a)
	return s.Validate()
}

func (s GetVariableResultType) Validate() error {
	if s.AttributeStatus == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "attributeStatus", "required field is missing")
	}

	if s.AttributeValue != nil {
		if err := validateStringLen(*s.AttributeValue, 2500, "attributeValue"); err != nil {
			return err
		}
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	if s.AttributeStatusInfo != nil {
		if err := s.AttributeStatusInfo.Validate(); err != nil {
			return ocpp.WrapField("attributeStatusInfo", err)
		}
	}

	return nil
}

// IdTokenInfoType (2.27)
//
// Contains status information about an identifier. It is advised to not stop charging for a token that expires during
// charging, as ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the status has no end date.
type IdTokenInfoType struct {
	// Required. Current status of the ID Token.
	Status AuthorizationStatusEnumType `json:"status"`
	// Optional. Date and Time after which the token must be considered invalid.
	CacheExpiryDateTime *time.Time `json:"cacheExpiryDateTime,omitempty"`
	// Optional. Priority from a business point of view. Default priority is 0, The range is from -9 to 9. Higher values
	// indicate a higher priority. The chargingPriority in TransactionEventResponse overrules this one.
	ChargingPriority *int32 `json:"chargingPriority,omitempty"`
	// Optional. Preferred user interface language of identifier user. Contains a language code as defined in [RFC5646].
	Language1 *string `json:"language1,omitempty"`
	// Optional. Only used when the IdToken is only valid for one or more specific EVSEs, not for the entire Charging
	// Station.
	EVSEID []int32 `json:"evseId,omitempty"`
	// Optional. Second preferred user interface language of identifier user. Don’t use when language1 is omitted, has to
	// be different from language1. Contains a language code as defined in [RFC5646].
	Language2 *string `json:"language2,omitempty"`
	// Optional. This contains the group identifier.
	GroupIDToken *IdTokenType `json:"groupIdToken,omitempty"`
	// Optional. Personal message that can be shown to the EV Driver and can be used for tariff information, user
	// greetings etc.
	PersonalMessage *MessageContentType `json:"personalMessage,omitempty"`
}

func (s *IdTokenInfoType) UnmarshalJSON(data []byte) error {
	type Alias IdTokenInfoType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = IdTokenInfoType(a)
	return s.Validate()
}

func (s IdTokenInfoType) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.Language1 != nil {
		if err := validateStringLen(*s.Language1, 8, "language1"); err != nil {
			return err
		}
	}

	if s.Language2 != nil {
		if err := validateStringLen(*s.Language2, 8, "language2"); err != nil {
			return err
		}
	}

	if s.GroupIDToken != nil {
		if err := s.GroupIDToken.Validate(); err != nil {
			return ocpp.WrapField("groupIdToken", err)
		}
	}

	if s.PersonalMessage != nil {
		if err := s.PersonalMessage.Validate(); err != nil {
			return ocpp.WrapField("personalMessage", err)
		}
	}

	return nil
}

// IdTokenType (2.28)
//
// Contains a case insensitive identifier to use for the authorization and the type of authorization to support
// multiple forms of identifiers.
type IdTokenType struct {
	// Required. IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can for example also contain a
	// UUID.
	IDToken IdentifierString36Type `json:"idToken"`
	// Required. Enumeration of possible idToken types.
	Type IdTokenEnumType `json:"type"`
	// Optional. AdditionalInfo can be used to send extra information which can be validated by the CSMS in addition to
	// the regular authorization with IdToken. AdditionalInfo contains one or more custom types, which need to be agreed
	// upon by all parties involved. When AdditionalInfo is NOT implemented or a not supported AdditionalInfo.type is
	// used, the CSMS/Charging Station MAY ignore the AdditionalInfo.
	AdditionalInfo []AdditionalInfoType `json:"additionalInfo,omitempty"`
}

func (s *IdTokenType) UnmarshalJSON(data []byte) error {
	type Alias IdTokenType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = IdTokenType(a)
	return s.Validate()
}

func (s IdTokenType) Validate() error {
	if s.IDToken == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idToken", "required field is missing")
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	for i, v := range s.AdditionalInfo {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("additionalInfo[%d]", i), err)
		}
	}

	return nil
}

// LogParametersType (2.29)
//
// Generic class for the configuration of logging entries.
type LogParametersType struct {
	// Required. The URL of the location at the remote system where the log should be stored.
	RemoteLocation string `json:"remoteLocation"`
	// Optional. This contains the date and time of the oldest logging information to include in the diagnostics.
	OldestTimestamp *time.Time `json:"oldestTimestamp,omitempty"`
	// Optional. This contains the date and time of the latest logging information to include in the diagnostics.
	LatestTimestamp *time.Time `json:"latestTimestamp,omitempty"`
}

func (s *LogParametersType) UnmarshalJSON(data []byte) error {
	type Alias LogParametersType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = LogParametersType(a)
	return s.Validate()
}

func (s LogParametersType) Validate() error {
	if s.RemoteLocation == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "remoteLocation", "required field is missing")
	}

	if err := validateStringLen(s.RemoteLocation, 512, "remoteLocation"); err != nil {
		return err
	}

	return nil
}

// MessageContentType (2.30)
//
// Contains message details, for a message to be displayed on a Charging Station.
type MessageContentType struct {
	// Required. Format of the message.
	Format MessageFormatEnumType `json:"format"`
	// Optional. Message language identifier. Contains a language code as defined in [RFC5646].
	Language *string `json:"language,omitempty"`
	// Required. Message contents.
	Content string `json:"content"`
}

func (s *MessageContentType) UnmarshalJSON(data []byte) error {
	type Alias MessageContentType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MessageContentType(a)
	return s.Validate()
}

func (s MessageContentType) Validate() error {
	if s.Format == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "format", "required field is missing")
	}

	if s.Language != nil {
		if err := validateStringLen(*s.Language, 8, "language"); err != nil {
			return err
		}
	}

	if s.Content == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "content", "required field is missing")
	}

	if err := validateStringLen(s.Content, 512, "content"); err != nil {
		return err
	}

	return nil
}

// MessageInfoType (2.31)
//
// Contains message details, for a message to be displayed on a Charging Station.
type MessageInfoType struct {
	// Required. Unique id within an exchange context. It is defined within the OCPP context as a positive Integer value
	// (greater or equal to zero).
	ID int32 `json:"id"`
	// Required. With what priority should this message be shown
	Priority MessagePriorityEnumType `json:"priority"`
	// Optional. During what state should this message be shown. When omitted this message should be shown in any state of
	// the Charging Station.
	State *MessageStateEnumType `json:"state,omitempty"`
	// Optional. From what date-time should this message be shown. If omitted: directly.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Optional. Until what date-time should this message be shown, after this date/time this message SHALL be removed.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Optional. During which transaction shall this message be shown. Message SHALL be removed by the Charging Station
	// after transaction has ended.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Required. Contains message details for the message to be displayed on a Charging Station.
	Message MessageContentType `json:"message"`
	// Optional. When a Charging Station has multiple Displays, this field can be used to define to which Display this
	// message belongs.
	Display *ComponentType `json:"display,omitempty"`
}

func (s *MessageInfoType) UnmarshalJSON(data []byte) error {
	type Alias MessageInfoType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MessageInfoType(a)
	return s.Validate()
}

func (s MessageInfoType) Validate() error {
	if s.Priority == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "priority", "required field is missing")
	}

	if s.TransactionID != nil {
		if err := s.TransactionID.Validate(); err != nil {
			return ocpp.WrapField("transactionId", err)
		}
	}

	if err := s.Message.Validate(); err != nil {
		return ocpp.WrapField("message", err)
	}

	if s.Display != nil {
		if err := s.Display.Validate(); err != nil {
			return ocpp.WrapField("display", err)
		}
	}

	return nil
}

// MeterValueType (2.32)
//
// Collection of one or more sampled values in MeterValuesRequest and TransactionEvent. All sampled values in a
// MeterValue are sampled at the same point in time.
type MeterValueType struct {
	// Required. Timestamp for measured value(s).
	Timestamp time.Time `json:"timestamp"`
	// Required. One or more measured values
	SampledValue []SampledValueType `json:"sampledValue"`
}

func (s *MeterValueType) UnmarshalJSON(data []byte) error {
	type Alias MeterValueType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MeterValueType(a)
	return s.Validate()
}

func (s MeterValueType) Validate() error {
	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if err := validateSliceLen(s.SampledValue, 1, -1, "sampledValue"); err != nil {
		return err
	}

	for i, v := range s.SampledValue {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("sampledValue[%d]", i), err)
		}
	}

	return nil
}

// ModemType (2.33)
//
// Defines parameters required for initiating and maintaining wireless communication with other devices.
type ModemType struct {
	// Optional. This contains the ICCID of the modem’s SIM card.
	Iccid *IdentifierString20Type `json:"iccid,omitempty"`
	// Optional. This contains the IMSI of the modem’s SIM card.
	Imsi *IdentifierString20Type `json:"imsi,omitempty"`
}

func (s *ModemType) UnmarshalJSON(data []byte) error {
	type Alias ModemType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ModemType(a)
	return s.Validate()
}

func (s ModemType) Validate() error {
	if s.Iccid != nil {
		if err := s.Iccid.Validate(); err != nil {
			return ocpp.WrapField("iccid", err)
		}
	}

	if s.Imsi != nil {
		if err := s.Imsi.Validate(); err != nil {
			return ocpp.WrapField("imsi", err)
		}
	}

	return nil
}

// MonitoringDataType (2.34)
//
// Class to hold parameters of SetVariableMonitoring request.
type MonitoringDataType struct {
	// Required. Component for which monitoring report was requested.
	Component ComponentType `json:"component"`
	// Required. Variable for which monitoring report was requested.
	Variable VariableType `json:"variable"`
	// Required. List of monitors for this Component-Variable pair.
	VariableMonitoring []VariableMonitoringType `json:"variableMonitoring"`
}

func (s *MonitoringDataType) UnmarshalJSON(data []byte) error {
	type Alias MonitoringDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = MonitoringDataType(a)
	return s.Validate()
}

func (s MonitoringDataType) Validate() error {
	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	if err := validateSliceLen(s.VariableMonitoring, 1, -1, "variableMonitoring"); err != nil {
		return err
	}

	for i, v := range s.VariableMonitoring {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("variableMonitoring[%d]", i), err)
		}
	}

	return nil
}

// NetworkConnectionProfileType (2.35)
//
// The NetworkConnectionProfile defines the functional and technical parameters of a communication link.
type NetworkConnectionProfileType struct {
	// Required. Defines the OCPP version used for this communication function. This field is ignored, since the OCPP
	// version to use is determined during the websocket handshake.
	OCPPVersion OCPPVersionEnumType `json:"ocppVersion"`
	// Required. Defines the transport protocol (e.g. SOAP or JSON). Note: SOAP is not supported in OCPP 2.0, but is
	// supported by other versions of OCPP.
	OCPPTransport OCPPTransportEnumType `json:"ocppTransport"`
	// Required. URL of the CSMS(s) that this Charging Station communicates with, without the Charging Station identity
	// part. The SecurityCtrlr.Identity field is appended to ocppCsmsUrl to provide the full websocket URL
	OCPPCSMSURL string `json:"ocppCsmsUrl"`
	// Required. Duration in seconds before a message send by the Charging Station via this network connection times- out.
	// The best setting depends on the underlying network and response times of the CSMS. If you are looking for a some
	// guideline: use 30 seconds as a starting point.
	MessageTimeout int32 `json:"messageTimeout"`
	// Required. This field specifies the security profile used when connecting to the CSMS with this
	// NetworkConnectionProfile.
	SecurityProfile int32 `json:"securityProfile"`
	// Required. Applicable Network Interface. Charging Station is allowed to use a different network interface to connect
	// if the given one does not work
	OCPPInterface OCPPInterfaceEnumType `json:"ocppInterface"`
	// Optional. Settings to be used to set up the VPN connection
	VPN *VPNType `json:"vpn,omitempty"`
	// Optional. Collection of configuration data needed to make a data-connection over a cellular network.
	APN *APNType `json:"apn,omitempty"`
}

func (s *NetworkConnectionProfileType) UnmarshalJSON(data []byte) error {
	type Alias NetworkConnectionProfileType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = NetworkConnectionProfileType(a)
	return s.Validate()
}

func (s NetworkConnectionProfileType) Validate() error {
	if s.OCPPVersion == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppVersion", "required field is missing")
	}

	if s.OCPPTransport == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppTransport", "required field is missing")
	}

	if s.OCPPCSMSURL == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppCsmsUrl", "required field is missing")
	}

	if err := validateStringLen(s.OCPPCSMSURL, 512, "ocppCsmsUrl"); err != nil {
		return err
	}

	if s.OCPPInterface == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppInterface", "required field is missing")
	}

	if s.VPN != nil {
		if err := s.VPN.Validate(); err != nil {
			return ocpp.WrapField("vpn", err)
		}
	}

	if s.APN != nil {
		if err := s.APN.Validate(); err != nil {
			return ocpp.WrapField("apn", err)
		}
	}

	return nil
}

// OCSPRequestDataType (2.36)
type OCSPRequestDataType struct {
	// Required. Used algorithms for the hashes provided.
	HashAlgorithm HashAlgorithmEnumType `json:"hashAlgorithm"`
	// Required. The hash of the issuer’s distinguished name (DN), that must be calculated over the DER encoding of the
	// issuer’s name field in the certificate being checked.
	IssuerNameHash IdentifierString128Type `json:"issuerNameHash"`
	// Required. The hash of the DER encoded public key: the value (excluding tag and length) of the subject public key
	// field in the issuer’s certificate.
	IssuerKeyHash string `json:"issuerKeyHash"`
	// Required. The string representation of the hexadecimal value of the serial number without the prefix "0x" and
	// without leading zeroes.
	SerialNumber IdentifierString40Type `json:"serialNumber"`
	// Required. This contains the responder URL (Case insensitive).
	ResponderURL string `json:"responderURL"`
}

func (s *OCSPRequestDataType) UnmarshalJSON(data []byte) error {
	type Alias OCSPRequestDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = OCSPRequestDataType(a)
	return s.Validate()
}

func (s OCSPRequestDataType) Validate() error {
	if s.HashAlgorithm == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "hashAlgorithm", "required field is missing")
	}

	if s.IssuerNameHash == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "issuerNameHash", "required field is missing")
	}

	if err := s.IssuerNameHash.Validate(); err != nil {
		return ocpp.WrapField("issuerNameHash", err)
	}

	if s.IssuerKeyHash == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "issuerKeyHash", "required field is missing")
	}

	if err := validateStringLen(s.IssuerKeyHash, 128, "issuerKeyHash"); err != nil {
		return err
	}

	if s.SerialNumber == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "serialNumber", "required field is missing")
	}

	if err := s.SerialNumber.Validate(); err != nil {
		return ocpp.WrapField("serialNumber", err)
	}

	if s.ResponderURL == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "responderURL", "required field is missing")
	}

	if err := validateStringLen(s.ResponderURL, 512, "responderURL"); err != nil {
		return err
	}

	return nil
}

// RelativeTimeIntervalType (2.37)
type RelativeTimeIntervalType struct {
	// Required. Start of the interval, in seconds from NOW.
	Start int32 `json:"start"`
	// Optional. Duration of the interval, in seconds.
	Duration *int32 `json:"duration,omitempty"`
}

func (s *RelativeTimeIntervalType) UnmarshalJSON(data []byte) error {
	type Alias RelativeTimeIntervalType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RelativeTimeIntervalType(a)
	return s.Validate()
}

func (s RelativeTimeIntervalType) Validate() error {
	_ = s
	return nil
}

// ReportDataType (2.38)
//
// Class to report components, variables and variable attributes and characteristics.
type ReportDataType struct {
	// Required. Component for which a report of Variable is requested.
	Component ComponentType `json:"component"`
	// Required. Variable for which report is requested.
	Variable VariableType `json:"variable"`
	// Required. Attribute data of a variable.
	VariableAttribute []VariableAttributeType `json:"variableAttribute"`
	// Optional. Fixed read-only parameters of a variable.
	VariableCharacteristics *VariableCharacteristicsType `json:"variableCharacteristics,omitempty"`
}

func (s *ReportDataType) UnmarshalJSON(data []byte) error {
	type Alias ReportDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReportDataType(a)
	return s.Validate()
}

func (s ReportDataType) Validate() error {
	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	if err := validateSliceLen(s.VariableAttribute, 1, 4, "variableAttribute"); err != nil {
		return err
	}

	for i, v := range s.VariableAttribute {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("variableAttribute[%d]", i), err)
		}
	}

	if s.VariableCharacteristics != nil {
		if err := s.VariableCharacteristics.Validate(); err != nil {
			return ocpp.WrapField("variableCharacteristics", err)
		}
	}

	return nil
}

// SalesTariffEntryType (2.39)
type SalesTariffEntryType struct {
	// Optional. Defines the price level of this SalesTariffEntry (referring to NumEPriceLevels). Small values for the
	// EPriceLevel represent a cheaper TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	EPriceLevel *int32 `json:"ePriceLevel,omitempty"`
	// Required. Defines the time interval the SalesTariffEntry is valid for, based upon relative times.
	RelativeTimeInterval RelativeTimeIntervalType `json:"relativeTimeInterval"`
	// Optional. Defines additional means for further relative price information and/or alternative costs.
	ConsumptionCost []ConsumptionCostType `json:"consumptionCost,omitempty"`
}

func (s *SalesTariffEntryType) UnmarshalJSON(data []byte) error {
	type Alias SalesTariffEntryType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SalesTariffEntryType(a)
	return s.Validate()
}

func (s SalesTariffEntryType) Validate() error {
	if s.EPriceLevel != nil {
		if err := validateNonNegative(*s.EPriceLevel, "ePriceLevel"); err != nil {
			return err
		}
	}

	if err := s.RelativeTimeInterval.Validate(); err != nil {
		return ocpp.WrapField("relativeTimeInterval", err)
	}

	if err := validateSliceLen(s.ConsumptionCost, 0, 3, "consumptionCost"); err != nil {
		return err
	}

	for i, v := range s.ConsumptionCost {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("consumptionCost[%d]", i), err)
		}
	}

	return nil
}

// SalesTariffType (2.40)
//
// NOTE      This dataType is based on dataTypes from ISO 15118-2.
type SalesTariffType struct {
	// Required. SalesTariff identifier used to identify one sales tariff. An SAID remains a unique identifier for one
	// schedule throughout a charging session.
	ID int32 `json:"id"`
	// Optional. A human readable title/short description of the sales tariff e.g. for HMI display purposes.
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty"`
	// Optional. Defines the overall number of distinct price levels used across all provided SalesTariff elements.
	NumEPriceLevels *int32 `json:"numEPriceLevels,omitempty"`
	// Required. Encapsulating element describing all relevant details for one time interval of the SalesTariff. The
	// number of SalesTariffEntry elements is limited by the parameter maxScheduleTuples.
	SalesTariffEntry []SalesTariffEntryType `json:"salesTariffEntry"`
}

func (s *SalesTariffType) UnmarshalJSON(data []byte) error {
	type Alias SalesTariffType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SalesTariffType(a)
	return s.Validate()
}

func (s SalesTariffType) Validate() error {
	if s.SalesTariffDescription != nil {
		if err := validateStringLen(*s.SalesTariffDescription, 32, "salesTariffDescription"); err != nil {
			return err
		}
	}

	if err := validateSliceLen(s.SalesTariffEntry, 1, 1024, "salesTariffEntry"); err != nil {
		return err
	}

	for i, v := range s.SalesTariffEntry {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("salesTariffEntry[%d]", i), err)
		}
	}

	return nil
}

// SampledValueType (2.41)
//
// Single sampled value in MeterValues. Each value can be accompanied by optional fields.
//
// To save on mobile data usage, default values of all of the optional fields are such that. The value without any
// additional fields will be interpreted, as a register reading of active import energy in Wh (Watt-hour) units.
type SampledValueType struct {
	// Required. Indicates the measured value.
	Value float64 `json:"value"`
	// Optional. Type of detail value: start, end or sample. Default = "Sample.Periodic"
	Context *ReadingContextEnumType `json:"context,omitempty"`
	// Optional. Type of measurement. Default = "Energy.Active.Import.Register"
	Measurand *MeasurandEnumType `json:"measurand,omitempty"`
	// Optional. Indicates how the measured value is to be interpreted. For instance between L1 and neutral (L1-N) Please
	// note that not all values of phase are applicable to all Measurands. When phase is absent, the measured value is
	// interpreted as an overall value.
	Phase *PhaseEnumType `json:"phase,omitempty"`
	// Optional. Indicates where the measured value has been sampled. Default = "Outlet"
	Location *LocationEnumType `json:"location,omitempty"`
	// Optional. Contains the MeterValueSignature with sign/encoding method information.
	SignedMeterValue *SignedMeterValueType `json:"signedMeterValue,omitempty"`
	// Optional. Represents a UnitOfMeasure including a multiplier
	UnitOfMeasure *UnitOfMeasureType `json:"unitOfMeasure,omitempty"`
}

func (s *SampledValueType) UnmarshalJSON(data []byte) error {
	type Alias SampledValueType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SampledValueType(a)
	return s.Validate()
}

func (s SampledValueType) Validate() error {
	if s.SignedMeterValue != nil {
		if err := s.SignedMeterValue.Validate(); err != nil {
			return ocpp.WrapField("signedMeterValue", err)
		}
	}

	if s.UnitOfMeasure != nil {
		if err := s.UnitOfMeasure.Validate(); err != nil {
			return ocpp.WrapField("unitOfMeasure", err)
		}
	}

	return nil
}

// SetMonitoringDataType (2.42)
//
// Class to hold parameters of SetVariableMonitoring request.
type SetMonitoringDataType struct {
	// Optional. An id SHALL only be given to replace an existing monitor. The Charging Station handles the generation of
	// id’s for new monitors.
	ID *int32 `json:"id,omitempty"`
	// Optional. Monitor only active when a transaction is ongoing on a component relevant to this transaction. Default =
	// false.
	Transaction *bool `json:"transaction,omitempty"`
	// Required. Value for threshold or delta monitoring. For Periodic or PeriodicClockAligned this is the interval in
	// seconds.
	Value float64 `json:"value"`
	// Required. The type of this monitor, e.g. a threshold, delta or periodic monitor.
	Type MonitorEnumType `json:"type"`
	// Required. The severity that will be assigned to an event that is triggered by this monitor. The severity range is
	// 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels have the following meaning:
	// 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and action should be taken
	// immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue regular operations due to
	// Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station is unable to continue
	// regular operations due to software or minor hardware issues. Action is required. 3-Critical Indicates a critical
	// error. Action is required. 4-Error Indicates a non-urgent error. Action is required. 5-Alert Indicates an alert
	// event. Default severity for any type of monitoring event. 6-Warning Indicates a warning event. Action may be
	// required. 7-Notice Indicates an unusual event. No immediate action is required. 8-Informational Indicates a regular
	// operational event. May be used for reporting, measuring throughput, etc. No action is required. 9-Debug Indicates
	// information useful to developers for debugging, not useful during operations.
	Severity int32 `json:"severity"`
	// Required. Component for which monitor is set.
	Component ComponentType `json:"component"`
	// Required. Variable for which monitor is set.
	Variable VariableType `json:"variable"`
}

func (s *SetMonitoringDataType) UnmarshalJSON(data []byte) error {
	type Alias SetMonitoringDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetMonitoringDataType(a)
	return s.Validate()
}

func (s SetMonitoringDataType) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	return nil
}

// SetMonitoringResultType (2.43)
//
// Class to hold result of SetVariableMonitoring request.
type SetMonitoringResultType struct {
	// Optional. Id given to the VariableMonitor by the Charging Station. The Id is only returned when status is accepted.
	// Installed VariableMonitors should have unique id’s but the id’s of removed Installed monitors should have unique
	// id’s but the id’s of removed monitors MAY be reused.
	ID *int32 `json:"id,omitempty"`
	// Required. Status is OK if a value could be returned. Otherwise this will indicate the reason why a value could not
	// be returned.
	Status SetMonitoringStatusEnumType `json:"status"`
	// Required. The type of this monitor, e.g. a threshold, delta or periodic monitor.
	Type MonitorEnumType `json:"type"`
	// Required. The severity that will be assigned to an event that is triggered by this monitor. The severity range is
	// 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels have the following meaning:
	// 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and action should be taken
	// immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue regular operations due to
	// Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station is unable to continue
	// regular operations due to software or minor hardware issues. Action is required. 3-Critical Indicates a critical
	// error. Action is required. 4-Error Indicates a non-urgent error. Action is required. 5-Alert Indicates an alert
	// event. Default severity for any type of monitoring event. 6-Warning Indicates a warning event. Action may be
	// required. 7-Notice Indicates an unusual event. No immediate action is required. 8-Informational Indicates a regular
	// operational event. May be used for reporting, measuring throughput, etc. No action is required. 9-Debug Indicates
	// information useful to developers for debugging, not useful during operations.
	Severity int32 `json:"severity"`
	// Required. Component for which status is returned.
	Component ComponentType `json:"component"`
	// Required. Variable for which status is returned.
	Variable VariableType `json:"variable"`
	// Optional. Detailed status information.
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *SetMonitoringResultType) UnmarshalJSON(data []byte) error {
	type Alias SetMonitoringResultType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetMonitoringResultType(a)
	return s.Validate()
}

func (s SetMonitoringResultType) Validate() error {
	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// SetVariableDataType (2.44)
type SetVariableDataType struct {
	// Optional. Type of attribute: Actual, Target, MinSet, MaxSet. Default is Actual when omitted.
	AttributeType *AttributeEnumType `json:"attributeType,omitempty"`
	// Required. Value to be assigned to attribute of variable. The value is allowed to be an empty string (""). The
	// Configuration Variable ConfigurationValueSize can be used to limit SetVariableData.attributeValue and
	// VariableCharacteristics.valueList. The max size of these values will always remain equal.
	AttributeValue string `json:"attributeValue"`
	// Required. The component for which the variable data is set.
	Component ComponentType `json:"component"`
	// Required. Specifies the that needs to be set.
	Variable VariableType `json:"variable"`
}

func (s *SetVariableDataType) UnmarshalJSON(data []byte) error {
	type Alias SetVariableDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetVariableDataType(a)
	return s.Validate()
}

func (s SetVariableDataType) Validate() error {
	if s.AttributeValue == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "attributeValue", "required field is missing")
	}

	if err := validateStringLen(s.AttributeValue, 1000, "attributeValue"); err != nil {
		return err
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	return nil
}

// SetVariableResultType (2.45)
type SetVariableResultType struct {
	// Optional. Type of attribute: Actual, Target, MinSet, MaxSet. Default is Actual when omitted.
	AttributeType *AttributeEnumType `json:"attributeType,omitempty"`
	// Required. Result status of setting the variable.
	AttributeStatus SetVariableStatusEnumType `json:"attributeStatus"`
	// Required. The component for which result is returned.
	Component ComponentType `json:"component"`
	// Required. The variable for which the result is returned.
	Variable VariableType `json:"variable"`
	// Optional. Detailed attribute status information.
	AttributeStatusInfo *StatusInfoType `json:"attributeStatusInfo,omitempty"`
}

func (s *SetVariableResultType) UnmarshalJSON(data []byte) error {
	type Alias SetVariableResultType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SetVariableResultType(a)
	return s.Validate()
}

func (s SetVariableResultType) Validate() error {
	if s.AttributeStatus == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "attributeStatus", "required field is missing")
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	if s.AttributeStatusInfo != nil {
		if err := s.AttributeStatusInfo.Validate(); err != nil {
			return ocpp.WrapField("attributeStatusInfo", err)
		}
	}

	return nil
}

// SignedMeterValueType (2.46)
//
// Represent a signed version of the meter value.
type SignedMeterValueType struct {
	// Required. Base64 encoded, contains the signed data which might contain more then just the meter value. It can
	// contain information like timestamps, reference to a customer etc.
	SignedMeterData string `json:"signedMeterData"`
	// Required. Method used to create the digital signature.
	SigningMethod string `json:"signingMethod"`
	// Required. Method used to encode the meter values before applying the digital signature algorithm.
	EncodingMethod string `json:"encodingMethod"`
	// Required. Base64 encoded, sending depends on configuration variable PublicKeyWithSignedMeterValue.
	PublicKey string `json:"publicKey"`
}

func (s *SignedMeterValueType) UnmarshalJSON(data []byte) error {
	type Alias SignedMeterValueType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = SignedMeterValueType(a)
	return s.Validate()
}

func (s SignedMeterValueType) Validate() error {
	if s.SignedMeterData == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "signedMeterData", "required field is missing")
	}

	if err := validateStringLen(s.SignedMeterData, 2500, "signedMeterData"); err != nil {
		return err
	}

	if s.SigningMethod == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "signingMethod", "required field is missing")
	}

	if err := validateStringLen(s.SigningMethod, 50, "signingMethod"); err != nil {
		return err
	}

	if s.EncodingMethod == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "encodingMethod", "required field is missing")
	}

	if err := validateStringLen(s.EncodingMethod, 50, "encodingMethod"); err != nil {
		return err
	}

	if s.PublicKey == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "publicKey", "required field is missing")
	}

	if err := validateStringLen(s.PublicKey, 2500, "publicKey"); err != nil {
		return err
	}

	return nil
}

// StatusInfoType (2.47)
//
// Element providing more information about the status.
type StatusInfoType struct {
	// Required. A predefined code for the reason why the status is returned in this response. The string is case-
	// insensitive.
	ReasonCode string `json:"reasonCode"`
	// Optional. Additional text to provide detailed information.
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
}

func (s *StatusInfoType) UnmarshalJSON(data []byte) error {
	type Alias StatusInfoType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StatusInfoType(a)
	return s.Validate()
}

func (s StatusInfoType) Validate() error {
	if s.ReasonCode == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "reasonCode", "required field is missing")
	}

	if err := validateStringLen(s.ReasonCode, 20, "reasonCode"); err != nil {
		return err
	}

	if s.AdditionalInfo != nil {
		if err := validateStringLen(*s.AdditionalInfo, 512, "additionalInfo"); err != nil {
			return err
		}
	}

	return nil
}

// TransactionType (2.48)
type TransactionType struct {
	// Required. This contains the Id of the transaction.
	TransactionID IdentifierString36Type `json:"transactionId"`
	// Optional. Current charging state, is required when state has changed.
	ChargingState *ChargingStateEnumType `json:"chargingState,omitempty"`
	// Optional. Contains the total time that energy flowed from EVSE to EV during the transaction (in seconds). Note that
	// timeSpentCharging is smaller or equal to the duration of the transaction.
	TimeSpentCharging *int32 `json:"timeSpentCharging,omitempty"`
	// Optional. The <i>stoppedReason </i>is the reason/event that initiated the process of stopping the transaction. It
	// will normally be the user stopping authorization via card (Local or MasterPass) or app (Remote), but it can also be
	// CSMS revoking authorization (DeAuthorized), or disconnecting the EV when TxStopPoint = EVConnected
	// (EVDisconnected). Most other reasons are related to technical faults or energy limitations. MAY only be omitted
	// when <i>stoppedReason </i>is "Local"
	StoppedReason *ReasonEnumType `json:"stoppedReason,omitempty"`
	// Optional. The ID given to remote start request (RequestStartTransactionRequest. This enables to CSMS to match the
	// started transaction to the given start request.
	RemoteStartID *int32 `json:"remoteStartId,omitempty"`
}

func (s *TransactionType) UnmarshalJSON(data []byte) error {
	type Alias TransactionType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TransactionType(a)
	return s.Validate()
}

func (s TransactionType) Validate() error {
	if s.TransactionID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "transactionId", "required field is missing")
	}

	if err := s.TransactionID.Validate(); err != nil {
		return ocpp.WrapField("transactionId", err)
	}

	return nil
}

// UnitOfMeasureType (2.49)
//
// Represents a UnitOfMeasure with a multiplier
type UnitOfMeasureType struct {
	// Optional. Unit of the value. Default = "Wh" if the (default) measurand is an "Energy" type. This field SHALL use a
	// value from the list Standardized Units of Measurements in Part 2 Appendices. If an applicable unit is available in
	// that list, otherwise a "custom" unit might be used.
	Unit *string `json:"unit,omitempty"`
	// Optional. Multiplier, this value represents the exponent to base 10. I.e. multiplier 3 means 10 raised to the 3rd
	// power. Default is 0.
	Multiplier *int32 `json:"multiplier,omitempty"`
}

func (s *UnitOfMeasureType) UnmarshalJSON(data []byte) error {
	type Alias UnitOfMeasureType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = UnitOfMeasureType(a)
	return s.Validate()
}

func (s UnitOfMeasureType) Validate() error {
	if s.Unit != nil {
		if err := validateStringLen(*s.Unit, 20, "unit"); err != nil {
			return err
		}
	}

	return nil
}

// VariableAttributeType (2.50)
//
// Attribute data of a variable.
type VariableAttributeType struct {
	// Optional. Attribute: Actual, MinSet, MaxSet, etc. Defaults to Actual if absent.
	Type *AttributeEnumType `json:"type,omitempty"`
	// Optional. Value of the attribute. May only be omitted when mutability is set to 'WriteOnly'. The Configuration
	// Variable ReportingValueSize can be used to limit GetVariableResult.attributeValue, VariableAttribute.value and
	// EventData.actualValue. The max size of these values will always remain equal.
	Value *string `json:"value,omitempty"`
	// Optional. Defines the mutability of this attribute. Default is ReadWrite when omitted.
	Mutability *MutabilityEnumType `json:"mutability,omitempty"`
	// Optional. If true, value will be persistent across system reboots or power down. Default when omitted is false.
	Persistent *bool `json:"persistent,omitempty"`
	// Optional. If true, value that will never be changed by the Charging Station at runtime. Default when omitted is
	// false.
	Constant *bool `json:"constant,omitempty"`
}

func (s *VariableAttributeType) UnmarshalJSON(data []byte) error {
	type Alias VariableAttributeType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VariableAttributeType(a)
	return s.Validate()
}

func (s VariableAttributeType) Validate() error {
	if s.Value != nil {
		if err := validateStringLen(*s.Value, 2500, "value"); err != nil {
			return err
		}
	}

	return nil
}

// VariableCharacteristicsType (2.51)
//
// Fixed read-only parameters of a variable.
type VariableCharacteristicsType struct {
	// Optional. Unit of the variable. When the transmitted value has a unit, this field SHALL be included.
	Unit *string `json:"unit,omitempty"`
	// Required. Data type of this variable.
	DataType DataEnumType `json:"dataType"`
	// Optional. Minimum possible value of this variable.
	MinLimit *float64 `json:"minLimit,omitempty"`
	// Optional. Maximum possible value of this variable. When the datatype of this Variable is String, OptionList,
	// SequenceList or MemberList, this field defines the maximum length of the (CSV) string.
	MaxLimit *float64 `json:"maxLimit,omitempty"`
	// Optional. Mandatory when dataType = OptionList, MemberList or SequenceList. valuesList specifies the allowed values
	// for the type. * OptionList: The (Actual) Variable value must be a single value from the reported (CSV) enumeration
	// list. * MemberList: The (Actual) Variable value may be an (unordered) (sub-)set of the reported (CSV) valid values
	// list. * SequenceList: The (Actual) Variable value may be an ordered (priority, etc) (sub-)set of the reported (CSV)
	// valid values. This is a comma separated list. The Configuration Variable ConfigurationValueSize can be used to
	// limit SetVariableData.attributeValue and VariableCharacteristics.valueList. The max size of these values will
	// always remain equal.
	ValuesList *string `json:"valuesList,omitempty"`
	// Required. Flag indicating if this variable supports monitoring.
	SupportsMonitoring bool `json:"supportsMonitoring"`
}

func (s *VariableCharacteristicsType) UnmarshalJSON(data []byte) error {
	type Alias VariableCharacteristicsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VariableCharacteristicsType(a)
	return s.Validate()
}

func (s VariableCharacteristicsType) Validate() error {
	if s.Unit != nil {
		if err := validateStringLen(*s.Unit, 16, "unit"); err != nil {
			return err
		}
	}

	if s.DataType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "dataType", "required field is missing")
	}

	if s.ValuesList != nil {
		if err := validateStringLen(*s.ValuesList, 1000, "valuesList"); err != nil {
			return err
		}
	}

	return nil
}

// VariableMonitoringType (2.52)
//
// A monitoring setting for a variable.
type VariableMonitoringType struct {
	// Required. Identifies the monitor.
	ID int32 `json:"id"`
	// Required. Monitor only active when a transaction is ongoing on a component relevant to this transaction.
	Transaction bool `json:"transaction"`
	// Required. Value for threshold or delta monitoring. For Periodic or PeriodicClockAligned this is the interval in
	// seconds.
	Value float64 `json:"value"`
	// Required. The type of this monitor, e.g. a threshold, delta or periodic monitor.
	Type MonitorEnumType `json:"type"`
	// Required. The severity that will be assigned to an event that is triggered by this monitor. The severity range is
	// 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels have the following meaning:
	// 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and action should be taken
	// immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue regular operations due to
	// Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station is unable to continue
	// regular operations due to software or minor hardware issues. Action is required. 3-Critical Indicates a critical
	// error. Action is required. 4-Error Indicates a non-urgent error. Action is required. 5-Alert Indicates an alert
	// event. Default severity for any type of monitoring event. 6-Warning Indicates a warning event. Action may be
	// required. 7-Notice Indicates an unusual event. No immediate action is required. 8-Informational Indicates a regular
	// operational event. May be used for reporting, measuring throughput, etc. No action is required. 9-Debug Indicates
	// information useful to developers for debugging, not useful during operations.
	Severity int32 `json:"severity"`
}

func (s *VariableMonitoringType) UnmarshalJSON(data []byte) error {
	type Alias VariableMonitoringType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VariableMonitoringType(a)
	return s.Validate()
}

func (s VariableMonitoringType) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	return nil
}

// VariableType (2.53)
//
// Reference key to a component-variable.
type VariableType struct {
	// Required. Name of the variable. Name should be taken from the list of standardized variable names whenever
	// possible. Case Insensitive. strongly advised to use Camel Case.
	Name IdentifierString50Type `json:"name"`
	// Optional. Name of instance in case the variable exists as multiple instances. Case Insensitive. strongly advised to
	// use Camel Case.
	Instance *IdentifierString50Type `json:"instance,omitempty"`
}

func (s *VariableType) UnmarshalJSON(data []byte) error {
	type Alias VariableType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VariableType(a)
	return s.Validate()
}

func (s VariableType) Validate() error {
	if s.Name == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "name", "required field is missing")
	}

	if err := s.Name.Validate(); err != nil {
		return ocpp.WrapField("name", err)
	}

	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return ocpp.WrapField("instance", err)
		}
	}

	return nil
}

// VPNType (2.54)
//
// VPN Configuration settings
type VPNType struct {
	// Required. VPN Server Address
	Server string `json:"server"`
	// Required. VPN User
	User string `json:"user"`
	// Optional. VPN group.
	Group *string `json:"group,omitempty"`
	// Required. VPN Password.
	Password string `json:"password"`
	// Required. VPN shared secret.
	Key string `json:"key"`
	// Required. Type of VPN
	Type VPNEnumType `json:"type"`
}

func (s *VPNType) UnmarshalJSON(data []byte) error {
	type Alias VPNType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VPNType(a)
	return s.Validate()
}

func (s VPNType) Validate() error {
	if s.Server == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "server", "required field is missing")
	}

	if err := validateStringLen(s.Server, 512, "server"); err != nil {
		return err
	}

	if s.User == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "user", "required field is missing")
	}

	if err := validateStringLen(s.User, 20, "user"); err != nil {
		return err
	}

	if s.Group != nil {
		if err := validateStringLen(*s.Group, 20, "group"); err != nil {
			return err
		}
	}

	if s.Password == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "password", "required field is missing")
	}

	if err := validateStringLen(s.Password, 20, "password"); err != nil {
		return err
	}

	if s.Key == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "key", "required field is missing")
	}

	if err := validateStringLen(s.Key, 255, "key"); err != nil {
		return err
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	return nil
}
