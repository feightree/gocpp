package v21

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	ocpp "github.com/feightree/gocpp/ocpp"
)

// validateStringLen checks that s does not exceed maxLen characters, as
// required by OCPP 2.1's "string[0..N]" datatypes (Part 2 §2.1.4).
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

// validateNonPositive checks that v is <= 0.
func validateNonPositive[T int32 | float64](v T, field string) error {
	if v > 0 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, field, "must be <= 0")
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

// IdentifierString255Type (Part 2 §2.1.4)
//
// A case-insensitive identifierString of at most 255 characters, restricted
// to the character set a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
type IdentifierString255Type string

func NewIdentifierString255Type(s string) (IdentifierString255Type, error) {
	c := IdentifierString255Type(s)
	if err := c.Validate(); err != nil {
		return "", err
	}

	return c, nil
}

func (s IdentifierString255Type) Validate() error {
	return validateIdentifierString(string(s), 255, "IdentifierString255Type")
}

func (s IdentifierString255Type) Equals(other string) bool {
	return strings.EqualFold(string(s), other)
}

func (s IdentifierString255Type) String() string {
	return string(s)
}

func (s *IdentifierString255Type) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	c, err := NewIdentifierString255Type(raw)
	if err != nil {
		return err
	}

	*s = c
	return nil
}

// validateIdentifierString checks that s is no longer than maxLen and
// contains only characters from the identifierString charset defined in
// OCPP 2.1 Part 2 §2.1.4: a-z, A-Z, 0-9, "*", "-", "_", "=", ":", "+", "|", "@", ".".
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

// AbsolutePriceScheduleType (2.1)
//
// The AbsolutePriceScheduleType is modeled after the same type that is defined in ISO 15118-20, such that if it
// is supplied by an EMSP as a signed EXI message, the conversion from EXI to JSON (in OCPP) and back to EXI (for
// ISO 15118-20) does not change the digest and therefore does not invalidate the signature.
type AbsolutePriceScheduleType struct {
	// Required. Starting point of price schedule.
	TimeAnchor time.Time `json:"timeAnchor"`
	// Required. Unique ID of price schedule
	PriceScheduleID int32 `json:"priceScheduleID"`
	// Optional. Description of the price schedule.
	PriceScheduleDescription *string `json:"priceScheduleDescription,omitempty"`
	// Required. Currency according to ISO 4217.
	Currency string `json:"currency"`
	// Required. String that indicates what language is used for the human readable strings in the price schedule.
	// Based on ISO 639.
	Language string `json:"language"`
	// Required. A string in URN notation which shall uniquely identify an algorithm that defines how to compute an
	// energy fee sum for a specific power profile based on the EnergyFee information from the PriceRule elements.
	PriceAlgorithm string `json:"priceAlgorithm"`
	// Required. A set of pricing rules for parking and energy costs.
	PriceRuleStacks []PriceRuleStackType `json:"priceRuleStacks"`
	// Optional. Optional. Describes the applicable tax rule(s) for this price schedule.
	TaxRules []TaxRuleType `json:"taxRules,omitempty"`
	// Optional. Optional. A set of prices for optional services (e.g. valet, carwash).
	AdditionalSelectedServices []AdditionalSelectedServicesType `json:"additionalSelectedServices,omitempty"`
	// Optional. Optional. A set of overstay rules that allows for escalation of charges after the overstay is
	// triggered.
	OverstayRuleList *OverstayRuleListType `json:"overstayRuleList,omitempty"`
	// Optional. Optional. Minimum amount to be billed for the overall charging session (e.g. including energy,
	// parking, and overstay).
	MinimumCost *RationalNumberType `json:"minimumCost,omitempty"`
	// Optional. Optional. Maximum amount to be billed for the overall charging session (e.g. including energy,
	// parking, and overstay).
	MaximumCost *RationalNumberType `json:"maximumCost,omitempty"`
}

func (s *AbsolutePriceScheduleType) UnmarshalJSON(data []byte) error {
	type Alias AbsolutePriceScheduleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AbsolutePriceScheduleType(a)
	return s.Validate()
}

func (s AbsolutePriceScheduleType) Validate() error {
	if s.TimeAnchor.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timeAnchor", "required field is missing")
	}

	if err := validateNonNegative(s.PriceScheduleID, "priceScheduleID"); err != nil {
		return err
	}

	if s.PriceScheduleDescription != nil {
		if err := validateStringLen(*s.PriceScheduleDescription, 160, "priceScheduleDescription"); err != nil {
			return err
		}
	}

	if s.Currency == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currency", "required field is missing")
	}

	if err := validateStringLen(s.Currency, 3, "currency"); err != nil {
		return err
	}

	if s.Language == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "language", "required field is missing")
	}

	if err := validateStringLen(s.Language, 8, "language"); err != nil {
		return err
	}

	if s.PriceAlgorithm == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "priceAlgorithm", "required field is missing")
	}

	if err := validateStringLen(s.PriceAlgorithm, 2000, "priceAlgorithm"); err != nil {
		return err
	}

	if err := validateSliceLen(s.PriceRuleStacks, 1, 1024, "priceRuleStacks"); err != nil {
		return err
	}

	for i, v := range s.PriceRuleStacks {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("priceRuleStacks[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.TaxRules, 0, 10, "taxRules"); err != nil {
		return err
	}

	for i, v := range s.TaxRules {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("taxRules[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.AdditionalSelectedServices, 0, 5, "additionalSelectedServices"); err != nil {
		return err
	}

	for i, v := range s.AdditionalSelectedServices {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("additionalSelectedServices[%d]", i), err)
		}
	}

	if s.OverstayRuleList != nil {
		if err := s.OverstayRuleList.Validate(); err != nil {
			return ocpp.WrapField("overstayRuleList", err)
		}
	}

	if s.MinimumCost != nil {
		if err := s.MinimumCost.Validate(); err != nil {
			return ocpp.WrapField("minimumCost", err)
		}
	}

	if s.MaximumCost != nil {
		if err := s.MaximumCost.Validate(); err != nil {
			return ocpp.WrapField("maximumCost", err)
		}
	}

	return nil
}

// ACChargingParametersType (2.2)
//
// EV AC charging parameters for ISO 15118-2
type ACChargingParametersType struct {
	// Required. Amount of energy requested (in Wh). This includes energy required for preconditioning. Relates to:
	// ISO 15118-2: AC_EVChargeParameterType: EAmount ISO 15118-20: Dynamic/Scheduled_SEReqControlModeType:
	// EVTargetEnergyRequest
	EnergyAmount float64 `json:"energyAmount"`
	// Required. Minimum current (amps) supported by the electric vehicle (per phase). Relates to: ISO 15118-2:
	// AC_EVChargeParameterType: EVMinCurrent
	EVMinCurrent float64 `json:"evMinCurrent"`
	// Required. Maximum current (amps) supported by the electric vehicle (per phase). Includes cable capacity.
	// Relates to: ISO 15118-2: AC_EVChargeParameterType: EVMaxCurrent
	EVMaxCurrent float64 `json:"evMaxCurrent"`
	// Required. Maximum voltage supported by the electric vehicle. Relates to: ISO 15118-2:
	// AC_EVChargeParameterType: EVMaxVoltage
	EVMaxVoltage float64 `json:"evMaxVoltage"`
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

// AdditionalInfoType (2.3)
//
// Contains a case insensitive identifier to use for the authorization and the type of authorization to support
// multiple forms of identifiers.
type AdditionalInfoType struct {
	// Required. (2.1) This field specifies the additional IdToken.
	AdditionalIDToken IdentifierString255Type `json:"additionalIdToken"`
	// Required. additionalInfo can be used to send extra information to CSMS in addition to the regular
	// authorization with IdToken. AdditionalInfo contains one or more custom types, which need to be agreed upon by
	// all parties involved. When the type is not supported, the CSMS/Charging Station MAY ignore the additionalInfo.
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

// AdditionalSelectedServicesType (2.4)
//
// Part of ISO 15118-20 price schedule.
type AdditionalSelectedServicesType struct {
	// Required. Human readable string to identify this service.
	ServiceName string `json:"serviceName"`
	// Required. Cost of the service.
	ServiceFee RationalNumberType `json:"serviceFee"`
}

func (s *AdditionalSelectedServicesType) UnmarshalJSON(data []byte) error {
	type Alias AdditionalSelectedServicesType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AdditionalSelectedServicesType(a)
	return s.Validate()
}

func (s AdditionalSelectedServicesType) Validate() error {
	if s.ServiceName == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "serviceName", "required field is missing")
	}

	if err := validateStringLen(s.ServiceName, 80, "serviceName"); err != nil {
		return err
	}

	if err := s.ServiceFee.Validate(); err != nil {
		return ocpp.WrapField("serviceFee", err)
	}

	return nil
}

// AddressType (2.5)
//
// (2.1) A generic address format.
type AddressType struct {
	// Required. Name of person/company
	Name string `json:"name"`
	// Required. Address line 1
	Address1 string `json:"address1"`
	// Optional. Address line 2
	Address2 *string `json:"address2,omitempty"`
	// Required. City
	City string `json:"city"`
	// Optional. Postal code
	PostalCode *string `json:"postalCode,omitempty"`
	// Required. Country name
	Country string `json:"country"`
}

func (s *AddressType) UnmarshalJSON(data []byte) error {
	type Alias AddressType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = AddressType(a)
	return s.Validate()
}

func (s AddressType) Validate() error {
	if s.Name == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "name", "required field is missing")
	}

	if err := validateStringLen(s.Name, 50, "name"); err != nil {
		return err
	}

	if s.Address1 == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "address1", "required field is missing")
	}

	if err := validateStringLen(s.Address1, 100, "address1"); err != nil {
		return err
	}

	if s.Address2 != nil {
		if err := validateStringLen(*s.Address2, 100, "address2"); err != nil {
			return err
		}
	}

	if s.City == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "city", "required field is missing")
	}

	if err := validateStringLen(s.City, 100, "city"); err != nil {
		return err
	}

	if s.PostalCode != nil {
		if err := validateStringLen(*s.PostalCode, 20, "postalCode"); err != nil {
			return err
		}
	}

	if s.Country == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "country", "required field is missing")
	}

	if err := validateStringLen(s.Country, 50, "country"); err != nil {
		return err
	}

	return nil
}

// APNType (2.6)
//
// Collection of configuration data needed to make a data-connection over a cellular network.
//
// When asking a GSM modem to dial in, it is possible to specify which mobile operator should be used. This can
// be done with the mobile country code (MCC) in combination with a mobile network code (MNC). Example: If your
// preferred network is Vodafone Netherlands, the MCC=204 and the MNC=04 which means the key NOTE
// PreferredNetwork = 20404 Some modems allows to specify a preferred network, which means, if this network is
// not available, a different network is used. If you specify UseOnlyPreferredNetwork and this network is not
// available, the modem will not dial in.
type APNType struct {
	// Required. The Access Point Name as an URL.
	APN string `json:"apn"`
	// Optional. APN username.
	APNUserName *string `json:"apnUserName,omitempty"`
	// Optional. (2.1) APN Password.
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

	if err := validateStringLen(s.APN, 2000, "apn"); err != nil {
		return err
	}

	if s.APNUserName != nil {
		if err := validateStringLen(*s.APNUserName, 50, "apnUserName"); err != nil {
			return err
		}
	}

	if s.APNPassword != nil {
		if err := validateStringLen(*s.APNPassword, 64, "apnPassword"); err != nil {
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

// AuthorizationData (2.7)
//
// Contains the identifier to use for authorization.
type AuthorizationData struct {
	// Optional. Required when UpdateType is Full. This contains information about authorization status, expiry and
	// group id. For a Differential update the following applies: If this element is present, then this entry SHALL
	// be added or updated in the Local Authorization List. If this element is absent, the entry for this IdToken in
	// the Local Authorization List SHALL be deleted.
	IDTokenInfo *IDTokenInfoType `json:"idTokenInfo,omitempty"`
	// Required. This contains the identifier which needs to be stored for authorization.
	IDToken IDTokenType `json:"idToken"`
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

// BatteryDataType (2.8)
type BatteryDataType struct {
	// Required. Slot number where battery is inserted or removed.
	EVSEID int32 `json:"evseId"`
	// Required. Serial number of battery.
	SerialNumber string `json:"serialNumber"`
	// Required. State of charge
	SOC float64 `json:"soC"`
	// Required. State of health
	SOH float64 `json:"soH"`
	// Optional. Production date of battery.
	ProductionDate *time.Time `json:"productionDate,omitempty"`
	// Optional. Vendor-specific info from battery in undefined format.
	VendorInfo *string `json:"vendorInfo,omitempty"`
}

func (s *BatteryDataType) UnmarshalJSON(data []byte) error {
	type Alias BatteryDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = BatteryDataType(a)
	return s.Validate()
}

func (s BatteryDataType) Validate() error {
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

	if s.SerialNumber == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "serialNumber", "required field is missing")
	}

	if err := validateStringLen(s.SerialNumber, 50, "serialNumber"); err != nil {
		return err
	}

	if s.SOC < 0 || s.SOC > 100 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "soC", "must be between 0 and 100")
	}

	if s.SOH < 0 || s.SOH > 100 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "soH", "must be between 0 and 100")
	}

	if s.VendorInfo != nil {
		if err := validateStringLen(*s.VendorInfo, 500, "vendorInfo"); err != nil {
			return err
		}
	}

	return nil
}

// CertificateHashDataChainType (2.9)
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

// CertificateHashDataType (2.10)
type CertificateHashDataType struct {
	// Required. Used algorithms for the hashes provided.
	HashAlgorithm HashAlgorithmEnumType `json:"hashAlgorithm"`
	// Required. The hash of the issuer’s distinguished name (DN), that must be calculated over the DER encoding of
	// the issuer’s name field in the certificate being checked.
	IssuerNameHash IdentifierString128Type `json:"issuerNameHash"`
	// Required. The hash of the DER encoded public key: the value (excluding tag and length) of the subject public
	// key field in the issuer’s certificate.
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

// CertificateStatusRequestInfoType (2.11)
//
// Data necessary to request the revocation status of a certificate.
type CertificateStatusRequestInfoType struct {
	// Required. Source of status: OCSP, CRL
	Source CertificateStatusSourceEnumType `json:"source"`
	// Required. URL(s) of source.
	Urls []string `json:"urls"`
	// Required. Hash data of certificate.
	CertificateHashData CertificateHashDataType `json:"certificateHashData"`
}

func (s *CertificateStatusRequestInfoType) UnmarshalJSON(data []byte) error {
	type Alias CertificateStatusRequestInfoType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CertificateStatusRequestInfoType(a)
	return s.Validate()
}

func (s CertificateStatusRequestInfoType) Validate() error {
	if s.Source == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "source", "required field is missing")
	}

	if err := validateSliceLen(s.Urls, 1, 5, "urls"); err != nil {
		return err
	}

	for i, v := range s.Urls {
		if err := validateStringLen(v, 2000, fmt.Sprintf("urls[%d]", i)); err != nil {
			return err
		}
	}

	if err := s.CertificateHashData.Validate(); err != nil {
		return ocpp.WrapField("certificateHashData", err)
	}

	return nil
}

// CertificateStatusType (2.12)
//
// Revocation status of certificate
type CertificateStatusType struct {
	// Required. Source of status: OCSP, CRL
	Source CertificateStatusSourceEnumType `json:"source"`
	// Required. Status of certificate: good, revoked or unknown.
	Status CertificateStatusEnumType `json:"status"`
	// Required.
	NextUpdate time.Time `json:"nextUpdate"`
	// Required. Hash data of the certificate.
	CertificateHashData CertificateHashDataType `json:"certificateHashData"`
}

func (s *CertificateStatusType) UnmarshalJSON(data []byte) error {
	type Alias CertificateStatusType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CertificateStatusType(a)
	return s.Validate()
}

func (s CertificateStatusType) Validate() error {
	if s.Source == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "source", "required field is missing")
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.NextUpdate.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "nextUpdate", "required field is missing")
	}

	if err := s.CertificateHashData.Validate(); err != nil {
		return ocpp.WrapField("certificateHashData", err)
	}

	return nil
}

// ChargingLimitType (2.13)
type ChargingLimitType struct {
	// Required. Represents the source of the charging limit. Values defined in appendix as
	// ChargingLimitSourceEnumStringType.
	ChargingLimitSource string `json:"chargingLimitSource"`
	// Optional. (2.1) True when the reported limit concerns local generation that is providing extra capacity,
	// instead of a limitation.
	IsLocalGeneration *bool `json:"isLocalGeneration,omitempty"`
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

	if err := validateStringLen(s.ChargingLimitSource, 20, "chargingLimitSource"); err != nil {
		return err
	}

	return nil
}

// ChargingNeedsType (2.14)
type ChargingNeedsType struct {
	// Required. Mode of energy transfer requested by the EV.
	RequestedEnergyTransfer EnergyTransferModeEnumType `json:"requestedEnergyTransfer"`
	// Optional. (2.1) Modes of energy transfer that are marked as available by EV.
	AvailableEnergyTransfer []EnergyTransferModeEnumType `json:"availableEnergyTransfer,omitempty"`
	// Optional. (2.1) Indicates whether EV wants to operate in Dynamic or Scheduled mode. When absent, Scheduled
	// mode is assumed for backwards compatibility. ISO 15118-20: ServiceSelectionReq(SelectedEnergyTransferService)
	ControlMode *ControlModeEnumType `json:"controlMode,omitempty"`
	// Optional. (2.1) Value of EVCC indicates that EV determines min/target SOC and departure time. A value of
	// EVCC_SECC indicates that charging station or CSMS may also update min/target SOC and departure time. ISO
	// 15118-20: ServiceSelectionReq(SelectedEnergyTransferService)
	MobilityNeedsMode *MobilityNeedsModeEnumType `json:"mobilityNeedsMode,omitempty"`
	// Optional. Estimated departure time of the EV. ISO 15118-2: AC/DC_EVChargeParameterType: DepartureTime ISO
	// 15118-20: Dynamic/Scheduled_SEReqControlModeType: DepartureTIme
	DepartureTime *time.Time `json:"departureTime,omitempty"`
	// Optional. (2.1) The list of charging parameters that apply to an ISO 15118-20 session or any other session
	// that supports bidirectional charging.
	V2XChargingParameters *V2XChargingParametersType `json:"v2xChargingParameters,omitempty"`
	// Optional. EV DC charging parameters
	DCChargingParameters *DCChargingParametersType `json:"dcChargingParameters,omitempty"`
	// Optional. EV AC charging parameters.
	ACChargingParameters *ACChargingParametersType `json:"acChargingParameters,omitempty"`
	// Optional. (2.1) Discharging and associated price offered by EV. Schedule periods during which EV is willing to
	// discharge have a negative value for power.
	EVEnergyOffer *EVEnergyOfferType `json:"evEnergyOffer,omitempty"`
	// Optional. (2.1) Additional charging parameters for ISO 15118-20 AC bidirectional sessions with DER control
	// (AC_BPT_DER)
	DERChargingParameters *DERChargingParametersType `json:"derChargingParameters,omitempty"`
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

	if s.V2XChargingParameters != nil {
		if err := s.V2XChargingParameters.Validate(); err != nil {
			return ocpp.WrapField("v2xChargingParameters", err)
		}
	}

	if s.DCChargingParameters != nil {
		if err := s.DCChargingParameters.Validate(); err != nil {
			return ocpp.WrapField("dcChargingParameters", err)
		}
	}

	if s.ACChargingParameters != nil {
		if err := s.ACChargingParameters.Validate(); err != nil {
			return ocpp.WrapField("acChargingParameters", err)
		}
	}

	if s.EVEnergyOffer != nil {
		if err := s.EVEnergyOffer.Validate(); err != nil {
			return ocpp.WrapField("evEnergyOffer", err)
		}
	}

	if s.DERChargingParameters != nil {
		if err := s.DERChargingParameters.Validate(); err != nil {
			return ocpp.WrapField("derChargingParameters", err)
		}
	}

	return nil
}

// ChargingPeriodType (2.15)
//
// A ChargingPeriodType consists of a start time, and a list of possible values that influence this period, for
// example: amount of energy charged this period, maximum current during this period etc.
type ChargingPeriodType struct {
	// Optional. Unique identifier of the Tariff that was used to calculate cost. If not provided, then cost was
	// calculated by some other means.
	TariffID *string `json:"tariffId,omitempty"`
	// Required. Start timestamp of charging period. A period ends when the next period starts. The last period ends
	// when the session ends.
	StartPeriod time.Time `json:"startPeriod"`
	// Optional. List of volume per cost dimension for this charging period.
	Dimensions []CostDimensionType `json:"dimensions,omitempty"`
}

func (s *ChargingPeriodType) UnmarshalJSON(data []byte) error {
	type Alias ChargingPeriodType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingPeriodType(a)
	return s.Validate()
}

func (s ChargingPeriodType) Validate() error {
	if s.TariffID != nil {
		if err := validateStringLen(*s.TariffID, 60, "tariffId"); err != nil {
			return err
		}
	}

	if s.StartPeriod.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "startPeriod", "required field is missing")
	}

	for i, v := range s.Dimensions {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("dimensions[%d]", i), err)
		}
	}

	return nil
}

// ChargingProfileCriterionType (2.16)
//
// A ChargingProfileCriterionType is a filter for charging profiles to be selected by a
// GetChargingProfilesRequest.
type ChargingProfileCriterionType struct {
	// Optional. Defines the purpose of the schedule transferred by this profile
	ChargingProfilePurpose *ChargingProfilePurposeEnumType `json:"chargingProfilePurpose,omitempty"`
	// Optional. Value determining level in hierarchy stack of profiles. Higher values have precedence over lower
	// values. Lowest level is 0.
	StackLevel *int32 `json:"stackLevel,omitempty"`
	// Optional. List of all the chargingProfileIds requested. Any ChargingProfile that matches one of these profiles
	// will be reported. If omitted, the Charging Station SHALL not filter on chargingProfileId. This field SHALL NOT
	// contain more ids than set in ChargingProfileEntries.maxLimit
	ChargingProfileID []int32 `json:"chargingProfileId,omitempty"`
	// Optional. For which charging limit sources, charging profiles SHALL be reported. If omitted, the Charging
	// Station SHALL not filter on chargingLimitSource. Values defined in Appendix as
	// ChargingLimitSourceEnumStringType.
	ChargingLimitSource []string `json:"chargingLimitSource,omitempty"`
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
	if s.StackLevel != nil {
		if err := validateNonNegative(*s.StackLevel, "stackLevel"); err != nil {
			return err
		}
	}

	if err := validateSliceLen(s.ChargingLimitSource, 0, 4, "chargingLimitSource"); err != nil {
		return err
	}

	for i, v := range s.ChargingLimitSource {
		if err := validateStringLen(v, 20, fmt.Sprintf("chargingLimitSource[%d]", i)); err != nil {
			return err
		}
	}

	return nil
}

// ChargingProfileType (2.17)
//
// A ChargingProfile consists of 1 to 3 ChargingSchedules with a list of ChargingSchedulePeriods, describing the
// amount of power or current that can be delivered per time interval.
type ChargingProfileType struct {
	// Required. Id of ChargingProfile. Unique within charging station. Id can have a negative value. This is useful
	// to distinguish charging profiles from an external actor (external constraints) from charging profiles received
	// from CSMS.
	ID int32 `json:"id"`
	// Required. Value determining level in hierarchy stack of profiles. Higher values have precedence over lower
	// values. Lowest level is 0.
	StackLevel int32 `json:"stackLevel"`
	// Required. Defines the purpose of the schedule transferred by this profile
	ChargingProfilePurpose ChargingProfilePurposeEnumType `json:"chargingProfilePurpose"`
	// Required. Indicates the kind of schedule.
	ChargingProfileKind ChargingProfileKindEnumType `json:"chargingProfileKind"`
	// Optional. Indicates the start point of a recurrence.
	RecurrencyKind *RecurrencyKindEnumType `json:"recurrencyKind,omitempty"`
	// Optional. Point in time at which the profile starts to be valid. If absent, the profile is valid as soon as it
	// is received by the Charging Station.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Optional. Point in time at which the profile stops to be valid. If absent, the profile is valid until it is
	// replaced by another profile.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// Optional. SHALL only be included if ChargingProfilePurpose is set to TxProfile in a SetChargingProfileRequest.
	// The transactionId is used to match the profile to a specific transaction.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Optional. (2.1) Period in seconds that this charging profile remains valid after the Charging Station has gone
	// offline. After this period the charging profile becomes invalid for as long as it is offline and the Charging
	// Station reverts back to a valid profile with a lower stack level. If invalidAfterOfflineDuration is true, then
	// this charging profile will become permanently invalid. A value of 0 means that the charging profile is
	// immediately invalid while offline. When the field is absent, then no timeout applies and the charging profile
	// remains valid when offline.
	MaxOfflineDuration *int32 `json:"maxOfflineDuration,omitempty"`
	// Optional. (2.1) When set to true this charging profile will not be valid anymore after being offline for more
	// than maxOfflineDuration. When absent defaults to false.
	InvalidAfterOfflineDuration *bool `json:"invalidAfterOfflineDuration,omitempty"`
	// Optional. (2.1) Interval in seconds after receipt of last update, when to request a profile update by sending
	// a PullDynamicScheduleUpdateRequest message. A value of 0 or no value means that no update interval applies.
	// Only relevant in a dynamic charging profile.
	DynUpdateInterval *int32 `json:"dynUpdateInterval,omitempty"`
	// Optional. (2.1) Time at which limits or setpoints in this charging profile were last updated by a
	// PullDynamicScheduleUpdateRequest or UpdateDynamicScheduleRequest or by an external actor. Only relevant in a
	// dynamic charging profile.
	DynUpdateTime *time.Time `json:"dynUpdateTime,omitempty"`
	// Optional. (2.1) ISO 15118-20 signature for all price schedules in chargingSchedules. Note: for 256-bit
	// elliptic curves (like secp256k1) the ECDSA signature is 512 bits (64 bytes) and for 521-bit curves (like
	// secp521r1) the signature is 1042 bits. This equals 131 bytes, which can be encoded as base64 in 176 bytes.
	PriceScheduleSignature *string `json:"priceScheduleSignature,omitempty"`
	// Required. Schedule that contains limits for the available power or current over time. In order to support ISO
	// 15118 schedule negotiation, it supports at most three schedules with associated tariff to choose from. Having
	// multiple chargingSchedules is only allowed for charging profiles of purpose TxProfile in the context of an ISO
	// 15118 charging session. For ISO 15118 Dynamic Control Mode only one chargingSchedule shall be provided.
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
	if err := validateNonNegative(s.StackLevel, "stackLevel"); err != nil {
		return err
	}

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

	if s.PriceScheduleSignature != nil {
		if err := validateStringLen(*s.PriceScheduleSignature, 256, "priceScheduleSignature"); err != nil {
			return err
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

// ChargingSchedulePeriodType (2.18)
//
// Charging schedule period structure defines a time period in a charging schedule. It is used in:
// CompositeScheduleType and in ChargingScheduleType. When used in a NotifyEVChargingScheduleRequest only
// startPeriod, limit, limit_L2, limit_L3 are relevant.
type ChargingSchedulePeriodType struct {
	// Required. Start of the period, in seconds from the start of schedule. The value of StartPeriod also defines
	// the stop time of the previous period.
	StartPeriod int32 `json:"startPeriod"`
	// Optional. Optional only when not required by the operationMode, as in CentralSetpoint, ExternalSetpoint,
	// ExternalLimits, LocalFrequency, LocalLoadBalancing. Charging rate limit during the schedule period, in the
	// applicable chargingRateUnit. This SHOULD be a non- negative value; a negative value is only supported for
	// backwards compatibility with older systems that use a negative value to specify a discharging limit. When
	// using chargingRateUnit = W, this field represents the sum of the power of all phases, unless values are
	// provided for L2 and L3, in which case this field represents phase L1.
	Limit *float64 `json:"limit,omitempty"`
	// Optional. (2.1) Charging rate limit on phase L2 in the applicable chargingRateUnit.
	LimitL2 *float64 `json:"limit_L2,omitempty"`
	// Optional. (2.1) Charging rate limit on phase L3 in the applicable chargingRateUnit.
	LimitL3 *float64 `json:"limit_L3,omitempty"`
	// Optional. The number of phases that can be used for charging. For a DC EVSE this field should be omitted. For
	// an AC EVSE a default value of numberPhases = 3 will be assumed if the field is absent.
	NumberPhases *int32 `json:"numberPhases,omitempty"`
	// Optional. Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching the phase connected to
	// the EV, i.e. ACPhaseSwitchingSupported is defined and true. It’s not allowed unless both conditions above are
	// true. If both conditions are true, and phaseToUse is omitted, the Charging Station / EVSE will make the
	// selection on its own.
	PhaseToUse *int32 `json:"phaseToUse,omitempty"`
	// Optional. (2.1) Limit in chargingRateUnit that the EV is allowed to discharge with. Note, these are negative
	// values in order to be consistent with setpoint, which can be positive and negative. For AC this field
	// represents the sum of all phases, unless values are provided for L2 and L3, in which case this field
	// represents phase L1.
	DischargeLimit *float64 `json:"dischargeLimit,omitempty"`
	// Optional. (2.1) Limit in chargingRateUnit on phase L2 that the EV is allowed to discharge with.
	DischargeLimitL2 *float64 `json:"dischargeLimit_L2,omitempty"`
	// Optional. (2.1) Limit in chargingRateUnit on phase L3 that the EV is allowed to discharge with.
	DischargeLimitL3 *float64 `json:"dischargeLimit_L3,omitempty"`
	// Optional. (2.1) Setpoint in chargingRateUnit that the EV should follow as close as possible. Use negative
	// values for discharging. When a limit and/or dischargeLimit are given the overshoot when following setpoint
	// must remain within these values. This field represents the sum of all phases, unless values are provided for
	// L2 and L3, in which case this field represents phase L1.
	Setpoint *float64 `json:"setpoint,omitempty"`
	// Optional. (2.1) Setpoint in chargingRateUnit that the EV should follow on phase L2 as close as possible.
	SetpointL2 *float64 `json:"setpoint_L2,omitempty"`
	// Optional. (2.1) Setpoint in chargingRateUnit that the EV should follow on phase L3 as close as possible.
	SetpointL3 *float64 `json:"setpoint_L3,omitempty"`
	// Optional. (2.1) Setpoint for reactive power (or current) in chargingRateUnit that the EV should follow as
	// closely as possible. Positive values for inductive, negative for capacitive reactive power or current. This
	// field represents the sum of all phases, unless values are provided for L2 and L3, in which case this field
	// represents phase L1.
	SetpointReactive *float64 `json:"setpointReactive,omitempty"`
	// Optional. (2.1) Setpoint for reactive power (or current) in chargingRateUnit that the EV should follow on
	// phase L2 as closely as possible.
	SetpointReactiveL2 *float64 `json:"setpointReactive_L2,omitempty"`
	// Optional. (2.1) Setpoint for reactive power (or current) in chargingRateUnit that the EV should follow on
	// phase L3 as closely as possible.
	SetpointReactiveL3 *float64 `json:"setpointReactive_L3,omitempty"`
	// Optional. (2.1) If true, the EV should attempt to keep the BMS preconditioned for this time interval.
	PreconditioningRequest *bool `json:"preconditioningRequest,omitempty"`
	// Optional. (2.1) If true, the EVSE must turn off power electronics/modules associated with this transaction.
	// Default value when absent is false.
	EVSESleep *bool `json:"evseSleep,omitempty"`
	// Optional. (2.1) Power value that, when present, is used as a baseline on top of which values from
	// v2xFreqWattCurve and v2xSignalWattCurve are added.
	V2XBaseline *float64 `json:"v2xBaseline,omitempty"`
	// Optional. (2.1) Charging operation mode to use during this time interval. When absent defaults to
	// ChargingOnly.
	OperationMode *OperationModeEnumType `json:"operationMode,omitempty"`
	// Optional. (2.1) Only required when operationMode = LocalFrequency. When used it must contain at least two
	// coordinates to specify a power-frequency table to use during this period. The table determines the value of
	// setpoint power for a given frequency. chargingRateUnit must be W for LocalFrequency control.
	V2XFreqWattCurve []V2XFreqWattPointType `json:"v2xFreqWattCurve,omitempty"`
	// Optional. (2.1) Only used, but not required, when operationMode = LocalFrequency. When used it must contain at
	// least two coordinates to specify a signal- frequency curve to use during this period. The curve determines the
	// value of setpoint power for a given signal. chargingRateUnit must be W for LocalFrequency control.
	V2XSignalWattCurve []V2XSignalWattPointType `json:"v2xSignalWattCurve,omitempty"`
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
	if s.NumberPhases != nil {
		if *s.NumberPhases < 0 || *s.NumberPhases > 3 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "numberPhases", "must be between 0 and 3")
		}
	}

	if s.PhaseToUse != nil {
		if *s.PhaseToUse < 0 || *s.PhaseToUse > 3 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "phaseToUse", "must be between 0 and 3")
		}
	}

	if s.DischargeLimit != nil {
		if err := validateNonPositive(*s.DischargeLimit, "dischargeLimit"); err != nil {
			return err
		}
	}

	if s.DischargeLimitL2 != nil {
		if err := validateNonPositive(*s.DischargeLimitL2, "dischargeLimit_L2"); err != nil {
			return err
		}
	}

	if s.DischargeLimitL3 != nil {
		if err := validateNonPositive(*s.DischargeLimitL3, "dischargeLimit_L3"); err != nil {
			return err
		}
	}

	if err := validateSliceLen(s.V2XFreqWattCurve, 0, 20, "v2xFreqWattCurve"); err != nil {
		return err
	}

	for i, v := range s.V2XFreqWattCurve {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("v2xFreqWattCurve[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.V2XSignalWattCurve, 0, 20, "v2xSignalWattCurve"); err != nil {
		return err
	}

	for i, v := range s.V2XSignalWattCurve {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("v2xSignalWattCurve[%d]", i), err)
		}
	}

	return nil
}

// ChargingScheduleType (2.19)
//
// Charging schedule structure defines a list of charging periods, as used in: NotifyEVChargingScheduleRequest
// and ChargingProfileType. When used in a NotifyEVChargingScheduleRequest only duration and
// chargingSchedulePeriod are relevant and chargingRateUnit must be 'W'. An ISO 15118-20 session may provide
// either an absolutePriceSchedule or a priceLevelSchedule. An ISO 15118-2 session can only provide
// a_salesTariff_ element. The field digestValue is used when price schedule or sales tariff are signed.
type ChargingScheduleType struct {
	// Required.
	ID int32 `json:"id"`
	// Optional. Starting point of an absolute schedule or recurring schedule.
	StartSchedule *time.Time `json:"startSchedule,omitempty"`
	// Optional. Duration of the charging schedule in seconds. If the duration is left empty, the last period will
	// continue indefinitely or until end of the transaction in case startSchedule is absent.
	Duration *int32 `json:"duration,omitempty"`
	// Required. The unit of measure in which limits and setpoints are expressed.
	ChargingRateUnit ChargingRateUnitEnumType `json:"chargingRateUnit"`
	// Optional. Minimum charging rate supported by the EV. The unit of measure is defined by the chargingRateUnit.
	// This parameter is intended to be used by a local smart charging algorithm to optimize the power allocation for
	// in the case a charging process is inefficient at lower charging rates.
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`
	// Optional. (2.1) Power tolerance when following EVPowerProfile.
	PowerTolerance *float64 `json:"powerTolerance,omitempty"`
	// Optional. (2.1) Id of this element for referencing in a signature.
	SignatureID *int32 `json:"signatureId,omitempty"`
	// Optional. (2.1) Base64 encoded hash (SHA256 for ISO 15118-2, SHA512 for ISO 15118-20) of the EXI price
	// schedule element. Used in signature.
	DigestValue *string `json:"digestValue,omitempty"`
	// Optional. (2.1) Defaults to false. When true, disregard time zone offset in dateTime fields of
	// ChargingScheduleType and use unqualified local time at Charging Station instead. This allows the same Absolute
	// or Recurring charging profile to be used in both summer and winter time.
	UseLocalTime *bool `json:"useLocalTime,omitempty"`
	// Optional. (2.1) Defaults to 0. When randomizedDelay not equals zero, then the start of each
	// ChargingSchedulePeriodType is delayed by a randomly chosen number of seconds between 0 and randomizedDelay.
	// Only allowed for TxProfile and TxDefaultProfile.
	RandomizedDelay *int32 `json:"randomizedDelay,omitempty"`
	// Optional. Sales tariff for charging associated with this schedule.
	SalesTariff *SalesTariffType `json:"salesTariff,omitempty"`
	// Required. List of ChargingSchedulePeriod elements defining maximum power or current usage over time. The
	// maximum number of periods, that is supported by the Charging Station, if less than 1024, is set by device
	// model variable SmartChargingCtrlr.PeriodsPerSchedule.
	ChargingSchedulePeriod []ChargingSchedulePeriodType `json:"chargingSchedulePeriod"`
	// Optional. (2.1) The ISO 15118-20 absolute price schedule.
	AbsolutePriceSchedule *AbsolutePriceScheduleType `json:"absolutePriceSchedule,omitempty"`
	// Optional. (2.1) The ISO 15118-20 price level schedule
	PriceLevelSchedule *PriceLevelScheduleType `json:"priceLevelSchedule,omitempty"`
	// Optional. (2.1) When present and SoC of EV is greater than or equal to soc, then charging limit or setpoint
	// will be capped to the value of limit.
	LimitAtSOC *LimitAtSOCType `json:"limitAtSoC,omitempty"`
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

	if s.SignatureID != nil {
		if err := validateNonNegative(*s.SignatureID, "signatureId"); err != nil {
			return err
		}
	}

	if s.DigestValue != nil {
		if err := validateStringLen(*s.DigestValue, 88, "digestValue"); err != nil {
			return err
		}
	}

	if s.RandomizedDelay != nil {
		if err := validateNonNegative(*s.RandomizedDelay, "randomizedDelay"); err != nil {
			return err
		}
	}

	if s.SalesTariff != nil {
		if err := s.SalesTariff.Validate(); err != nil {
			return ocpp.WrapField("salesTariff", err)
		}
	}

	if err := validateSliceLen(s.ChargingSchedulePeriod, 1, 1024, "chargingSchedulePeriod"); err != nil {
		return err
	}

	for i, v := range s.ChargingSchedulePeriod {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingSchedulePeriod[%d]", i), err)
		}
	}

	if s.AbsolutePriceSchedule != nil {
		if err := s.AbsolutePriceSchedule.Validate(); err != nil {
			return ocpp.WrapField("absolutePriceSchedule", err)
		}
	}

	if s.PriceLevelSchedule != nil {
		if err := s.PriceLevelSchedule.Validate(); err != nil {
			return ocpp.WrapField("priceLevelSchedule", err)
		}
	}

	if s.LimitAtSOC != nil {
		if err := s.LimitAtSOC.Validate(); err != nil {
			return ocpp.WrapField("limitAtSoC", err)
		}
	}

	return nil
}

// ChargingScheduleUpdateType (2.20)
//
// Updates to a ChargingSchedulePeriodType for dynamic charging profiles.
type ChargingScheduleUpdateType struct {
	// Optional. Optional only when not required by the operationMode, as in CentralSetpoint, ExternalSetpoint,
	// ExternalLimits, LocalFrequency, LocalLoadBalancing. Charging rate limit during the schedule period, in the
	// applicable chargingRateUnit. This SHOULD be a non- negative value; a negative value is only supported for
	// backwards compatibility with older systems that use a negative value to specify a discharging limit. For AC
	// this field represents the sum of all phases, unless values are provided for L2 and L3, in which case this
	// field represents phase L1.
	Limit *float64 `json:"limit,omitempty"`
	// Optional. (2.1) Charging rate limit on phase L2 in the applicable chargingRateUnit.
	LimitL2 *float64 `json:"limit_L2,omitempty"`
	// Optional. (2.1) Charging rate limit on phase L3 in the applicable chargingRateUnit.
	LimitL3 *float64 `json:"limit_L3,omitempty"`
	// Optional. (2.1) Limit in chargingRateUnit that the EV is allowed to discharge with. Note, these are negative
	// values in order to be consistent with setpoint, which can be positive and negative. For AC this field
	// represents the sum of all phases, unless values are provided for L2 and L3, in which case this field
	// represents phase L1.
	DischargeLimit *float64 `json:"dischargeLimit,omitempty"`
	// Optional. (2.1) Limit in chargingRateUnit on phase L2 that the EV is allowed to discharge with.
	DischargeLimitL2 *float64 `json:"dischargeLimit_L2,omitempty"`
	// Optional. (2.1) Limit in chargingRateUnit on phase L3 that the EV is allowed to discharge with.
	DischargeLimitL3 *float64 `json:"dischargeLimit_L3,omitempty"`
	// Optional. (2.1) Setpoint in chargingRateUnit that the EV should follow as close as possible. Use negative
	// values for discharging. When a limit and/or dischargeLimit are given the overshoot when following setpoint
	// must remain within these values. This field represents the sum of all phases, unless values are provided for
	// L2 and L3, in which case this field represents phase L1.
	Setpoint *float64 `json:"setpoint,omitempty"`
	// Optional. (2.1) Setpoint in chargingRateUnit that the EV should follow on phase L2 as close as possible.
	SetpointL2 *float64 `json:"setpoint_L2,omitempty"`
	// Optional. (2.1) Setpoint in chargingRateUnit that the EV should follow on phase L3 as close as possible.
	SetpointL3 *float64 `json:"setpoint_L3,omitempty"`
	// Optional. (2.1) Setpoint for reactive power (or current) in chargingRateUnit that the EV should follow as
	// closely as possible. Positive values for inductive, negative for capacitive reactive power or current. This
	// field represents the sum of all phases, unless values are provided for L2 and L3, in which case this field
	// represents phase L1.
	SetpointReactive *float64 `json:"setpointReactive,omitempty"`
	// Optional. (2.1) Setpoint for reactive power (or current) in chargingRateUnit that the EV should follow on
	// phase L2 as closely as possible.
	SetpointReactiveL2 *float64 `json:"setpointReactive_L2,omitempty"`
	// Optional. (2.1) Setpoint for reactive power (or current) in chargingRateUnit that the EV should follow on
	// phase L3 as closely as possible.
	SetpointReactiveL3 *float64 `json:"setpointReactive_L3,omitempty"`
}

func (s *ChargingScheduleUpdateType) UnmarshalJSON(data []byte) error {
	type Alias ChargingScheduleUpdateType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ChargingScheduleUpdateType(a)
	return s.Validate()
}

func (s ChargingScheduleUpdateType) Validate() error {
	if s.DischargeLimit != nil {
		if err := validateNonPositive(*s.DischargeLimit, "dischargeLimit"); err != nil {
			return err
		}
	}

	if s.DischargeLimitL2 != nil {
		if err := validateNonPositive(*s.DischargeLimitL2, "dischargeLimit_L2"); err != nil {
			return err
		}
	}

	if s.DischargeLimitL3 != nil {
		if err := validateNonPositive(*s.DischargeLimitL3, "dischargeLimit_L3"); err != nil {
			return err
		}
	}

	return nil
}

// ChargingStationType (2.21)
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

// ClearChargingProfileType (2.22)
//
// A ClearChargingProfileType is a filter for charging profiles to be cleared by ClearChargingProfileRequest.
type ClearChargingProfileType struct {
	// Optional. Specifies the id of the EVSE for which to clear charging profiles. An evseId of zero (0) specifies
	// the charging profile for the overall Charging Station. Absence of this parameter means the clearing applies to
	// all charging profiles that match the other criteria in the request.
	EVSEID *int32 `json:"evseId,omitempty"`
	// Optional. Specifies to purpose of the charging profiles that will be cleared, if they meet the other criteria
	// in the request.
	ChargingProfilePurpose *ChargingProfilePurposeEnumType `json:"chargingProfilePurpose,omitempty"`
	// Optional. Specifies the stackLevel for which charging profiles will be cleared, if they meet the other
	// criteria in the request.
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
	if s.EVSEID != nil {
		if err := validateNonNegative(*s.EVSEID, "evseId"); err != nil {
			return err
		}
	}

	if s.StackLevel != nil {
		if err := validateNonNegative(*s.StackLevel, "stackLevel"); err != nil {
			return err
		}
	}

	return nil
}

// ClearMonitoringResultType (2.23)
type ClearMonitoringResultType struct {
	// Required. Result of the clear request for this monitor, identified by its Id.
	Status ClearMonitoringStatusEnumType `json:"status"`
	// Required. Id of the monitor of which a clear was requested.
	ID int32 `json:"id"`
	// Optional. Element providing more information about the status.
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

	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if s.StatusInfo != nil {
		if err := s.StatusInfo.Validate(); err != nil {
			return ocpp.WrapField("statusInfo", err)
		}
	}

	return nil
}

// ClearTariffsResultType (2.24)
type ClearTariffsResultType struct {
	// Optional. Id of tariff for which status is reported. If no tariffs were found, then this field is absent, and
	// status will be NoTariff.
	TariffID *string `json:"tariffId,omitempty"`
	// Required.
	Status TariffClearStatusEnumType `json:"status"`
	// Optional. Additional info on status
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

func (s *ClearTariffsResultType) UnmarshalJSON(data []byte) error {
	type Alias ClearTariffsResultType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ClearTariffsResultType(a)
	return s.Validate()
}

func (s ClearTariffsResultType) Validate() error {
	if s.TariffID != nil {
		if err := validateStringLen(*s.TariffID, 60, "tariffId"); err != nil {
			return err
		}
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

// ComponentType (2.25)
//
// A physical or logical component
type ComponentType struct {
	// Required. Name of the component. Name should be taken from the list of standardized component names whenever
	// possible. Case Insensitive. strongly advised to use Camel Case.
	Name IdentifierString50Type `json:"name"`
	// Optional. Name of instance in case the component exists as multiple instances. Case Insensitive. strongly
	// advised to use Camel Case.
	Instance *IdentifierString50Type `json:"instance,omitempty"`
	// Optional. Specifies the EVSE when component is located at EVSE level, also specifies the connector when
	// component is located at Connector level.
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

// ComponentVariableType (2.26)
//
// Class to report components, variables and variable attributes and characteristics.
type ComponentVariableType struct {
	// Required. Component for which a report of Variable is requested.
	Component ComponentType `json:"component"`
	// Optional. Variable for which the report is requested.
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

// CompositeScheduleType (2.27)
type CompositeScheduleType struct {
	// Required.
	EVSEID int32 `json:"evseId"`
	// Required.
	Duration int32 `json:"duration"`
	// Required.
	ScheduleStart time.Time `json:"scheduleStart"`
	// Required.
	ChargingRateUnit ChargingRateUnitEnumType `json:"chargingRateUnit"`
	// Required. List of ChargingSchedulePeriod elements defining maximum power or current over time.
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
	if err := validateNonNegative(s.EVSEID, "evseId"); err != nil {
		return err
	}

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

// ConstantStreamDataType (2.28)
type ConstantStreamDataType struct {
	// Required. Uniquely identifies the stream
	ID int32 `json:"id"`
	// Required. Id of monitor used to report his event. It can be a preconfigured or hardwired monitor.
	VariableMonitoringID int32 `json:"variableMonitoringId"`
	// Required. Max time and items parameters
	Params PeriodicEventStreamParamsType `json:"params"`
}

func (s *ConstantStreamDataType) UnmarshalJSON(data []byte) error {
	type Alias ConstantStreamDataType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ConstantStreamDataType(a)
	return s.Validate()
}

func (s ConstantStreamDataType) Validate() error {
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if err := validateNonNegative(s.VariableMonitoringID, "variableMonitoringId"); err != nil {
		return err
	}

	if err := s.Params.Validate(); err != nil {
		return ocpp.WrapField("params", err)
	}

	return nil
}

// ConsumptionCostType (2.29)
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

// CostDetailsType (2.30)
//
// CostDetailsType contains the cost as calculated by Charging Station based on provided TariffType.
//
// NOTE          Reservation is not shown as a chargingPeriod, because it took place outside of the transaction.
type CostDetailsType struct {
	// Optional. If set to true, then Charging Station has failed to calculate the cost.
	FailureToCalculate *bool `json:"failureToCalculate,omitempty"`
	// Optional. Optional human-readable reason text in case of failure to calculate.
	FailureReason *string `json:"failureReason,omitempty"`
	// Optional. List of Charging Periods that make up this charging session. A finished session has of 1 or more
	// periods, where each period has a different list of dimensions that determined the price. When sent as a
	// running cost update during a transaction chargingPeriods are omitted.
	ChargingPeriods []ChargingPeriodType `json:"chargingPeriods,omitempty"`
	// Required. Total sum of all the costs of this transaction in the specified currency.
	TotalCost TotalCostType `json:"totalCost"`
	// Required. Total usage of energy and time
	TotalUsage TotalUsageType `json:"totalUsage"`
}

func (s *CostDetailsType) UnmarshalJSON(data []byte) error {
	type Alias CostDetailsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CostDetailsType(a)
	return s.Validate()
}

func (s CostDetailsType) Validate() error {
	if s.FailureReason != nil {
		if err := validateStringLen(*s.FailureReason, 500, "failureReason"); err != nil {
			return err
		}
	}

	for i, v := range s.ChargingPeriods {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("chargingPeriods[%d]", i), err)
		}
	}

	if err := s.TotalCost.Validate(); err != nil {
		return ocpp.WrapField("totalCost", err)
	}

	if err := s.TotalUsage.Validate(); err != nil {
		return ocpp.WrapField("totalUsage", err)
	}

	return nil
}

// CostDimensionType (2.31)
//
// Volume consumed of cost dimension.
type CostDimensionType struct {
	// Required. Type of cost dimension: energy, power, time, etc.
	Type CostDimensionEnumType `json:"type"`
	// Required. Volume of the dimension consumed, measured according to the dimension type.
	Volume float64 `json:"volume"`
}

func (s *CostDimensionType) UnmarshalJSON(data []byte) error {
	type Alias CostDimensionType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = CostDimensionType(a)
	return s.Validate()
}

func (s CostDimensionType) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	return nil
}

// CostType (2.32)
type CostType struct {
	// Required. The kind of cost referred to in the message element amount
	CostKind CostKindEnumType `json:"costKind"`
	// Required. The estimated or actual cost per kWh
	Amount int32 `json:"amount"`
	// Optional. Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The final value is
	// determined by: amount * 10 ^ amountMultiplier
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

// DCChargingParametersType (2.33)
//
// EV DC charging parameters for ISO 15118-2
type DCChargingParametersType struct {
	// Required. Maximum current (in A) supported by the electric vehicle. Includes cable capacity. Relates to: ISO
	// 15118-2: DC_EVChargeParameterType:EVMaximumCurrentLimit
	EVMaxCurrent float64 `json:"evMaxCurrent"`
	// Required. Maximum voltage supported by the electric vehicle. Relates to: ISO 15118-2:
	// DC_EVChargeParameterType: EVMaximumVoltageLimit
	EVMaxVoltage float64 `json:"evMaxVoltage"`
	// Optional. Maximum power (in W) supported by the electric vehicle. Required for DC charging. Relates to: ISO
	// 15118-2: DC_EVChargeParameterType: EVMaximumPowerLimit
	EVMaxPower *float64 `json:"evMaxPower,omitempty"`
	// Optional. Capacity of the electric vehicle battery (in Wh). Relates to: ISO 15118-2: DC_EVChargeParameterType:
	// EVEnergyCapacity
	EVEnergyCapacity *float64 `json:"evEnergyCapacity,omitempty"`
	// Optional. Amount of energy requested (in Wh). This inludes energy required for preconditioning. Relates to:
	// ISO 15118-2: DC_EVChargeParameterType: EVEnergyRequest
	EnergyAmount *float64 `json:"energyAmount,omitempty"`
	// Optional. Energy available in the battery (in percent of the battery capacity) Relates to: ISO 15118-2:
	// DC_EVChargeParameterType: DC_EVStatus: EVRESSSOC
	StateOfCharge *int32 `json:"stateOfCharge,omitempty"`
	// Optional. Percentage of SoC at which the EV considers the battery fully charged. (possible values: 0 - 100)
	// Relates to: ISO 15118-2: DC_EVChargeParameterType: FullSOC
	FullSOC *int32 `json:"fullSoC,omitempty"`
	// Optional. Percentage of SoC at which the EV considers a fast charging process to end. (possible values: 0 -
	// 100) Relates to: ISO 15118-2: DC_EVChargeParameterType: BulkSOC
	BulkSOC *int32 `json:"bulkSoC,omitempty"`
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

	if s.FullSOC != nil {
		if *s.FullSOC < 0 || *s.FullSOC > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "fullSoC", "must be between 0 and 100")
		}
	}

	if s.BulkSOC != nil {
		if *s.BulkSOC < 0 || *s.BulkSOC > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "bulkSoC", "must be between 0 and 100")
		}
	}

	return nil
}

// DERChargingParametersType (2.34)
//
// (2.1) DERChargingParametersType is used in ChargingNeedsType during an ISO 15118-20 session for AC_BPT_DER to
// report the inverter settings related to DER control that were agreed between EVSE and EV.
//
// Fields starting with "ev" contain values from the EV. Other fields contain a value that is supported by both
// EV and EVSE.
//
// DERChargingParametersType type is only relevant in case of an ISO 15118-20 AC_BPT_DER/AC_DER charging session.
//
// NOTE        All these fields have values greater or equal to zero (i.e. are non-negative)
type DERChargingParametersType struct {
	// Optional. DER control functions supported by EV. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:DERCon
	// trolFunctions (bitmap)
	EVSupportedDERControl []DERControlEnumType `json:"evSupportedDERControl,omitempty"`
	// Optional. Rated maximum injected active power by EV, at specified over-excited power factor
	// (overExcitedPowerFactor). It can also be defined as the rated maximum discharge power at the rated minimum
	// injected reactive power value. This means that if the EV is providing reactive power support, and it is
	// requested to discharge at max power (e.g. to satisfy an EMS request), the EV may override the request and
	// discharge up to overExcitedMaximumDischargePower to meet the minimum reactive power requirements. Corresponds
	// to the WOvPF attribute in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVOverExcitedMaximumDischargePower
	EVOverExcitedMaxDischargePower *float64 `json:"evOverExcitedMaxDischargePower,omitempty"`
	// Optional. EV power factor when injecting (over excited) the minimum reactive power. Corresponds to the OvPF
	// attribute in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType: EVOverExcitedPowerFactor
	EVOverExcitedPowerFactor *float64 `json:"evOverExcitedPowerFactor,omitempty"`
	// Optional. Rated maximum injected active power by EV supported at specified under-excited power factor
	// (EVUnderExcitedPowerFactor). It can also be defined as the rated maximum dischargePower at the rated minimum
	// absorbed reactive power value. This means that if the EV is providing reactive power support, and it is
	// requested to discharge at max power (e.g. to satisfy an EMS request), the EV may override the request and
	// discharge up to underExcitedMaximumDischargePower to meet the minimum reactive power requirements. This
	// corresponds to the WUnPF attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVUnderExcitedMaximumDischargePower
	EVUnderExcitedMaxDischargePower *float64 `json:"evUnderExcitedMaxDischargePower,omitempty"`
	// Optional. EV power factor when injecting (under excited) the minimum reactive power. Corresponds to the OvPF
	// attribute in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType: EVUnderExcitedPowerFactor
	EVUnderExcitedPowerFactor *float64 `json:"evUnderExcitedPowerFactor,omitempty"`
	// Optional. Rated maximum total apparent power, defined by min(EV, EVSE) in va. Corresponds to the VAMaxRtg in
	// IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType: EVMaximumApparentPower
	MaxApparentPower *float64 `json:"maxApparentPower,omitempty"`
	// Optional. Rated maximum absorbed apparent power, defined by min(EV, EVSE) in va. This field represents the sum
	// of all phases, unless values are provided for L2 and L3, in which case this field represents phase L1.
	// Corresponds to the ChaVAMaxRtg in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumChargeApparentPower
	MaxChargeApparentPower *float64 `json:"maxChargeApparentPower,omitempty"`
	// Optional. Rated maximum absorbed apparent power on phase L2, defined by min(EV, EVSE) in va. Corresponds to
	// the ChaVAMaxRtg in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumChargeApparentPower_L2
	MaxChargeApparentPowerL2 *float64 `json:"maxChargeApparentPower_L2,omitempty"`
	// Optional. Rated maximum absorbed apparent power on phase L3, defined by min(EV, EVSE) in va. Corresponds to
	// the ChaVAMaxRtg in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumChargeApparentPower_L3
	MaxChargeApparentPowerL3 *float64 `json:"maxChargeApparentPower_L3,omitempty"`
	// Optional. Rated maximum injected apparent power, defined by min(EV, EVSE) in va. This field represents the sum
	// of all phases, unless values are provided for L2 and L3, in which case this field represents phase L1.
	// Corresponds to the DisVAMaxRtg in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargeApparentPower
	MaxDischargeApparentPower *float64 `json:"maxDischargeApparentPower,omitempty"`
	// Optional. Rated maximum injected apparent power on phase L2, defined by min(EV, EVSE) in va. Corresponds to
	// the DisVAMaxRtg in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargeApparentPower_L2
	MaxDischargeApparentPowerL2 *float64 `json:"maxDischargeApparentPower_L2,omitempty"`
	// Optional. Rated maximum injected apparent power on phase L3, defined by min(EV, EVSE) in va. Corresponds to
	// the DisVAMaxRtg in IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargeApparentPower_L3
	MaxDischargeApparentPowerL3 *float64 `json:"maxDischargeApparentPower_L3,omitempty"`
	// Optional. Rated maximum absorbed reactive power, defined by min(EV, EVSE), in vars. This field represents the
	// sum of all phases, unless values are provided for L2 and L3, in which case this field represents phase L1.
	// Corresponds to the AvarMax attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumChargeReactivePower
	MaxChargeReactivePower *float64 `json:"maxChargeReactivePower,omitempty"`
	// Optional. Rated maximum absorbed reactive power, defined by min(EV, EVSE), in vars on phase L2. Corresponds to
	// the AvarMax attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumChargeReactivePower_L2
	MaxChargeReactivePowerL2 *float64 `json:"maxChargeReactivePower_L2,omitempty"`
	// Optional. Rated maximum absorbed reactive power, defined by min(EV, EVSE), in vars on phase L3. Corresponds to
	// the AvarMax attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumChargeReactivePower_L3
	MaxChargeReactivePowerL3 *float64 `json:"maxChargeReactivePower_L3,omitempty"`
	// Optional. Rated minimum absorbed reactive power, defined by max(EV, EVSE), in vars. This field represents the
	// sum of all phases, unless values are provided for L2 and L3, in which case this field represents phase L1. ISO
	// 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumChargeReactivePower
	MinChargeReactivePower *float64 `json:"minChargeReactivePower,omitempty"`
	// Optional. Rated minimum absorbed reactive power, defined by max(EV, EVSE), in vars on phase L2. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumChargeReactivePower_L2
	MinChargeReactivePowerL2 *float64 `json:"minChargeReactivePower_L2,omitempty"`
	// Optional. Rated minimum absorbed reactive power, defined by max(EV, EVSE), in vars on phase L3. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumChargeReactivePower_L3
	MinChargeReactivePowerL3 *float64 `json:"minChargeReactivePower_L3,omitempty"`
	// Optional. Rated maximum injected reactive power, defined by min(EV, EVSE), in vars. This field represents the
	// sum of all phases, unless values are provided for L2 and L3, in which case this field represents phase L1.
	// Corresponds to the IvarMax attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargeReactivePower
	MaxDischargeReactivePower *float64 `json:"maxDischargeReactivePower,omitempty"`
	// Optional. Rated maximum injected reactive power, defined by min(EV, EVSE), in vars on phase L2. Corresponds to
	// the IvarMax attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargeReactivePower_L2
	MaxDischargeReactivePowerL2 *float64 `json:"maxDischargeReactivePower_L2,omitempty"`
	// Optional. Rated maximum injected reactive power, defined by min(EV, EVSE), in vars on phase L3. Corresponds to
	// the IvarMax attribute in the IEC 61850. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargeReactivePower_L3
	MaxDischargeReactivePowerL3 *float64 `json:"maxDischargeReactivePower_L3,omitempty"`
	// Optional. Rated minimum injected reactive power, defined by max(EV, EVSE), in vars. This field represents the
	// sum of all phases, unless values are provided for L2 and L3, in which case this field represents phase L1. ISO
	// 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumDischargeReactivePower
	MinDischargeReactivePower *float64 `json:"minDischargeReactivePower,omitempty"`
	// Optional. Rated minimum injected reactive power, defined by max(EV, EVSE), in var on phase L2. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumDischargeReactivePower_L2
	MinDischargeReactivePowerL2 *float64 `json:"minDischargeReactivePower_L2,omitempty"`
	// Optional. Rated minimum injected reactive power, defined by max(EV, EVSE), in var on phase L3. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumDischargeReactivePower_L3
	MinDischargeReactivePowerL3 *float64 `json:"minDischargeReactivePower_L3,omitempty"`
	// Optional. Line voltage supported by EVSE and EV. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVNominalVoltage
	NominalVoltage *float64 `json:"nominalVoltage,omitempty"`
	// Optional. The nominal AC voltage (rms) offset between the Charging Station’s electrical connection point and
	// the utility’s point of common coupling. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVNominalVoltageOffset
	NominalVoltageOffset *float64 `json:"nominalVoltageOffset,omitempty"`
	// Optional. Maximum AC rms voltage, as defined by min(EV, EVSE) to operate with. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMaximumNominalVoltage
	MaxNominalVoltage *float64 `json:"maxNominalVoltage,omitempty"`
	// Optional. Minimum AC rms voltage, as defined by max(EV, EVSE) to operate with. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMinimumNominalVoltage
	MinNominalVoltage *float64 `json:"minNominalVoltage,omitempty"`
	// Optional. Manufacturer of the EV inverter. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVInverterManufacturer
	EVInverterManufacturer *string `json:"evInverterManufacturer,omitempty"`
	// Optional. Model name of the EV inverter. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVInverterModel
	EVInverterModel *string `json:"evInverterModel,omitempty"`
	// Optional. Serial number of the EV inverter. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVInverterSerialNumber
	EVInverterSerialNumber *string `json:"evInverterSerialNumber,omitempty"`
	// Optional. Software version of EV inverter. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVInverterSwVersion
	EVInverterSwVersion *string `json:"evInverterSwVersion,omitempty"`
	// Optional. Hardware version of EV inverter. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVInverterHwVersion
	EVInverterHwVersion *string `json:"evInverterHwVersion,omitempty"`
	// Optional. Type of islanding detection method. Only mandatory when islanding detection is required at the site,
	// as set in the ISO 15118 Service Details configuration. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVIslandingDetectionMethod
	EVIslandingDetectionMethod []IslandingDetectionEnumType `json:"evIslandingDetectionMethod,omitempty"`
	// Optional. Time after which EV will trip if an island has been detected. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVIslandingTripTime
	EVIslandingTripTime *float64 `json:"evIslandingTripTime,omitempty"`
	// Optional. Maximum injected DC current allowed at level 1 charging. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMaximumLevel1DCInjection
	EVMaximumLevel1DCInjection *float64 `json:"evMaximumLevel1DCInjection,omitempty"`
	// Optional. Maximum allowed duration of DC injection at level 1 charging. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVDurationLevel1DCInjection
	EVDurationLevel1DCInjection *float64 `json:"evDurationLevel1DCInjection,omitempty"`
	// Optional. Maximum injected DC current allowed at level 2 charging. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVMaximumLevel2DCInjection
	EVMaximumLevel2DCInjection *float64 `json:"evMaximumLevel2DCInjection,omitempty"`
	// Optional. Maximum allowed duration of DC injection at level 2 charging. ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVDurationLevel2DCInjection
	EVDurationLevel2DCInjection *float64 `json:"evDurationLevel2DCInjection,omitempty"`
	// Optional. Measure of the susceptibility of the circuit to reactance, in Siemens (S). ISO 15118-20:
	// DER_BPT_AC_CPDReqEnergyTransferModeType: EVReactiveSusceptance
	EVReactiveSusceptance *float64 `json:"evReactiveSusceptance,omitempty"`
	// Optional. Total energy value, in Wh, that EV is allowed to provide during the entire V2G session. The value is
	// independent of the V2X Cycling area. Once this value reaches the value of 0, the EV may block any attempt to
	// discharge in order to protect the battery health. ISO 15118-20: DER_BPT_AC_CPDReqEnergyTransferModeType:
	// EVSessionTotalDischargeEnergyAvailable
	EVSessionTotalDischargeEnergyAvailable *float64 `json:"evSessionTotalDischargeEnergyAvailable,omitempty"`
}

func (s *DERChargingParametersType) UnmarshalJSON(data []byte) error {
	type Alias DERChargingParametersType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DERChargingParametersType(a)
	return s.Validate()
}

func (s DERChargingParametersType) Validate() error {
	if s.EVInverterManufacturer != nil {
		if err := validateStringLen(*s.EVInverterManufacturer, 50, "evInverterManufacturer"); err != nil {
			return err
		}
	}

	if s.EVInverterModel != nil {
		if err := validateStringLen(*s.EVInverterModel, 50, "evInverterModel"); err != nil {
			return err
		}
	}

	if s.EVInverterSerialNumber != nil {
		if err := validateStringLen(*s.EVInverterSerialNumber, 50, "evInverterSerialNumber"); err != nil {
			return err
		}
	}

	if s.EVInverterSwVersion != nil {
		if err := validateStringLen(*s.EVInverterSwVersion, 50, "evInverterSwVersion"); err != nil {
			return err
		}
	}

	if s.EVInverterHwVersion != nil {
		if err := validateStringLen(*s.EVInverterHwVersion, 50, "evInverterHwVersion"); err != nil {
			return err
		}
	}

	return nil
}

// DERCurveGetType (2.35)
type DERCurveGetType struct {
	// Required. Id of DER curve
	ID IdentifierString36Type `json:"id"`
	// Required. Type of DER curve
	CurveType DERControlEnumType `json:"curveType"`
	// Required. True if this is a default curve
	IsDefault bool `json:"isDefault"`
	// Required. True if this setting is superseded by a higher priority setting (i.e. lower value of priority)
	IsSuperseded bool `json:"isSuperseded"`
	// Required. Parameters defining the DER curve
	Curve DERCurveType `json:"curve"`
}

func (s *DERCurveGetType) UnmarshalJSON(data []byte) error {
	type Alias DERCurveGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DERCurveGetType(a)
	return s.Validate()
}

func (s DERCurveGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if s.CurveType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "curveType", "required field is missing")
	}

	if err := s.Curve.Validate(); err != nil {
		return ocpp.WrapField("curve", err)
	}

	return nil
}

// DERCurvePointsType (2.36)
type DERCurvePointsType struct {
	// Required. The data value of the X-axis (independent) variable, depending on the curve type.
	X float64 `json:"x"`
	// Required. The data value of the Y-axis (dependent) variable, depending on the DERUnitEnumType of the curve. If
	// y is power factor, then a positive value means DER is absorbing reactive power (under-excited), a negative
	// value when DER is injecting reactive power (over-excited).
	Y float64 `json:"y"`
}

func (s *DERCurvePointsType) UnmarshalJSON(data []byte) error {
	type Alias DERCurvePointsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DERCurvePointsType(a)
	return s.Validate()
}

func (s DERCurvePointsType) Validate() error {
	_ = s
	return nil
}

// DERCurveType (2.37)
type DERCurveType struct {
	// Required. Priority of curve (0=highest)
	Priority int32 `json:"priority"`
	// Required. Unit of the Y-axis of DER curve
	YUnit DERUnitEnumType `json:"yUnit"`
	// Optional. Open loop response time, the time to ramp up to 90% of the new target in response to the change in
	// voltage, in seconds. A value of 0 is used to mean no limit. When not present, the device should follow its
	// default behavior.
	ResponseTime *float64 `json:"responseTime,omitempty"`
	// Optional. Point in time when this curve will become activated. Only absent when default is true.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Optional. Duration in seconds that this curve will be active. Only absent when default is true.
	Duration *float64 `json:"duration,omitempty"`
	// Optional. Hysteresis parameters for curve.
	Hysteresis *HysteresisType `json:"hysteresis,omitempty"`
	// Optional. Additional parameters for voltage curves.
	VoltageParams *VoltageParamsType `json:"voltageParams,omitempty"`
	// Optional. Additional parameters for VoltVar curve.
	ReactivePowerParams *ReactivePowerParamsType `json:"reactivePowerParams,omitempty"`
	// Required. Coordinates of the DER curve. X-axis is determined by curveType. Y-axis is determined by yUnit.
	CurveData []DERCurvePointsType `json:"curveData"`
}

func (s *DERCurveType) UnmarshalJSON(data []byte) error {
	type Alias DERCurveType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = DERCurveType(a)
	return s.Validate()
}

func (s DERCurveType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	if s.YUnit == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "yUnit", "required field is missing")
	}

	if s.Hysteresis != nil {
		if err := s.Hysteresis.Validate(); err != nil {
			return ocpp.WrapField("hysteresis", err)
		}
	}

	if s.VoltageParams != nil {
		if err := s.VoltageParams.Validate(); err != nil {
			return ocpp.WrapField("voltageParams", err)
		}
	}

	if s.ReactivePowerParams != nil {
		if err := s.ReactivePowerParams.Validate(); err != nil {
			return ocpp.WrapField("reactivePowerParams", err)
		}
	}

	if err := validateSliceLen(s.CurveData, 1, 10, "curveData"); err != nil {
		return err
	}

	for i, v := range s.CurveData {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("curveData[%d]", i), err)
		}
	}

	return nil
}

// EnterServiceGetType (2.38)
type EnterServiceGetType struct {
	// Required. Id of setting
	ID IdentifierString36Type `json:"id"`
	// Required. Enter Service settings
	EnterService EnterServiceType `json:"enterService"`
}

func (s *EnterServiceGetType) UnmarshalJSON(data []byte) error {
	type Alias EnterServiceGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EnterServiceGetType(a)
	return s.Validate()
}

func (s EnterServiceGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if err := s.EnterService.Validate(); err != nil {
		return ocpp.WrapField("enterService", err)
	}

	return nil
}

// EnterServiceType (2.39)
type EnterServiceType struct {
	// Required. Priority of setting (0=highest)
	Priority int32 `json:"priority"`
	// Required. Enter service voltage high
	HighVoltage float64 `json:"highVoltage"`
	// Required. Enter service voltage low
	LowVoltage float64 `json:"lowVoltage"`
	// Required. Enter service frequency high
	HighFreq float64 `json:"highFreq"`
	// Required. Enter service frequency low
	LowFreq float64 `json:"lowFreq"`
	// Optional. Enter service delay
	Delay *float64 `json:"delay,omitempty"`
	// Optional. Enter service randomized delay
	RandomDelay *float64 `json:"randomDelay,omitempty"`
	// Optional. Enter service ramp rate in seconds
	RampRate *float64 `json:"rampRate,omitempty"`
}

func (s *EnterServiceType) UnmarshalJSON(data []byte) error {
	type Alias EnterServiceType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EnterServiceType(a)
	return s.Validate()
}

func (s EnterServiceType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	return nil
}

// EVAbsolutePriceScheduleEntryType (2.40)
//
// (2.1) An entry in price schedule over time for which EV is willing to discharge.
type EVAbsolutePriceScheduleEntryType struct {
	// Required. The amount of seconds of this entry.
	Duration int32 `json:"duration"`
	// Required. A set of pricing rules for energy costs.
	EVPriceRule []EVPriceRuleType `json:"evPriceRule"`
}

func (s *EVAbsolutePriceScheduleEntryType) UnmarshalJSON(data []byte) error {
	type Alias EVAbsolutePriceScheduleEntryType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVAbsolutePriceScheduleEntryType(a)
	return s.Validate()
}

func (s EVAbsolutePriceScheduleEntryType) Validate() error {
	if err := validateSliceLen(s.EVPriceRule, 1, 8, "evPriceRule"); err != nil {
		return err
	}

	for i, v := range s.EVPriceRule {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("evPriceRule[%d]", i), err)
		}
	}

	return nil
}

// EVAbsolutePriceScheduleType (2.41)
//
// (2.1) Price schedule of EV energy offer.
type EVAbsolutePriceScheduleType struct {
	// Required. Starting point in time of the EVEnergyOffer.
	TimeAnchor time.Time `json:"timeAnchor"`
	// Required. Currency code according to ISO 4217.
	Currency string `json:"currency"`
	// Required. ISO 15118-20 URN of price algorithm: Power, PeakPower, StackedEnergy.
	PriceAlgorithm string `json:"priceAlgorithm"`
	// Required. Schedule of prices for which EV is willing to discharge.
	EVAbsolutePriceScheduleEntries []EVAbsolutePriceScheduleEntryType `json:"evAbsolutePriceScheduleEntries"`
}

func (s *EVAbsolutePriceScheduleType) UnmarshalJSON(data []byte) error {
	type Alias EVAbsolutePriceScheduleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVAbsolutePriceScheduleType(a)
	return s.Validate()
}

func (s EVAbsolutePriceScheduleType) Validate() error {
	if s.TimeAnchor.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timeAnchor", "required field is missing")
	}

	if s.Currency == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currency", "required field is missing")
	}

	if err := validateStringLen(s.Currency, 3, "currency"); err != nil {
		return err
	}

	if s.PriceAlgorithm == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "priceAlgorithm", "required field is missing")
	}

	if err := validateStringLen(s.PriceAlgorithm, 2000, "priceAlgorithm"); err != nil {
		return err
	}

	if err := validateSliceLen(s.EVAbsolutePriceScheduleEntries, 1, 1024, "evAbsolutePriceScheduleEntries"); err != nil {
		return err
	}

	for i, v := range s.EVAbsolutePriceScheduleEntries {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("evAbsolutePriceScheduleEntries[%d]", i), err)
		}
	}

	return nil
}

// EVEnergyOfferType (2.42)
//
// (2.1) A schedule of the energy amount over time that EV is willing to discharge. A negative value indicates
// the willingness to discharge under specific conditions, a positive value indicates that the EV currently is
// not able to offer energy to discharge.
type EVEnergyOfferType struct {
	// Required. Power schedule offered for discharging.
	EVPowerSchedule EVPowerScheduleType `json:"evPowerSchedule"`
	// Optional. Price schedule for which EV is willing to discharge.
	EVAbsolutePriceSchedule *EVAbsolutePriceScheduleType `json:"evAbsolutePriceSchedule,omitempty"`
}

func (s *EVEnergyOfferType) UnmarshalJSON(data []byte) error {
	type Alias EVEnergyOfferType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVEnergyOfferType(a)
	return s.Validate()
}

func (s EVEnergyOfferType) Validate() error {
	if err := s.EVPowerSchedule.Validate(); err != nil {
		return ocpp.WrapField("evPowerSchedule", err)
	}

	if s.EVAbsolutePriceSchedule != nil {
		if err := s.EVAbsolutePriceSchedule.Validate(); err != nil {
			return ocpp.WrapField("evAbsolutePriceSchedule", err)
		}
	}

	return nil
}

// EventDataType (2.43)
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
	// Required. Actual value (attributeType Actual) of the variable. The Configuration Variable ReportingValueSize
	// can be used to limit GetVariableResult.attributeValue, VariableAttribute.value and EventData.actualValue. The
	// max size of these values will always remain equal.
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
	// Optional. (2.1) Severity associated with the monitor in variableMonitoringId or with the hardwired
	// notification.
	Severity *int32 `json:"severity,omitempty"`
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
	if err := validateNonNegative(s.EventID, "eventId"); err != nil {
		return err
	}

	if s.Timestamp.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timestamp", "required field is missing")
	}

	if s.Trigger == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "trigger", "required field is missing")
	}

	if s.Cause != nil {
		if err := validateNonNegative(*s.Cause, "cause"); err != nil {
			return err
		}
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

	if s.VariableMonitoringID != nil {
		if err := validateNonNegative(*s.VariableMonitoringID, "variableMonitoringId"); err != nil {
			return err
		}
	}

	if s.EventNotificationType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "eventNotificationType", "required field is missing")
	}

	if s.Severity != nil {
		if err := validateNonNegative(*s.Severity, "severity"); err != nil {
			return err
		}
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	return nil
}

// EVPowerScheduleEntryType (2.44)
//
// (2.1) An entry in schedule of the energy amount over time that EV is willing to discharge. A negative value
// indicates the willingness to discharge under specific conditions, a positive value indicates that the EV
// currently is not able to offer energy to discharge.
type EVPowerScheduleEntryType struct {
	// Required. The duration of this entry.
	Duration int32 `json:"duration"`
	// Required. Defines maximum amount of power for the duration of this EVPowerScheduleEntry to be discharged from
	// the EV battery through EVSE power outlet. Negative values are used for discharging.
	Power float64 `json:"power"`
}

func (s *EVPowerScheduleEntryType) UnmarshalJSON(data []byte) error {
	type Alias EVPowerScheduleEntryType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVPowerScheduleEntryType(a)
	return s.Validate()
}

func (s EVPowerScheduleEntryType) Validate() error {
	_ = s
	return nil
}

// EVPowerScheduleType (2.45)
//
// (2.1) Schedule of EV energy offer.
type EVPowerScheduleType struct {
	// Required. The time that defines the starting point for the EVEnergyOffer.
	TimeAnchor time.Time `json:"timeAnchor"`
	// Required. List of EVPowerScheduleEntries.
	EVPowerScheduleEntries []EVPowerScheduleEntryType `json:"evPowerScheduleEntries"`
}

func (s *EVPowerScheduleType) UnmarshalJSON(data []byte) error {
	type Alias EVPowerScheduleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVPowerScheduleType(a)
	return s.Validate()
}

func (s EVPowerScheduleType) Validate() error {
	if s.TimeAnchor.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timeAnchor", "required field is missing")
	}

	if err := validateSliceLen(s.EVPowerScheduleEntries, 1, 1024, "evPowerScheduleEntries"); err != nil {
		return err
	}

	for i, v := range s.EVPowerScheduleEntries {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("evPowerScheduleEntries[%d]", i), err)
		}
	}

	return nil
}

// EVPriceRuleType (2.46)
//
// (2.1) An entry in price schedule over time for which EV is willing to discharge.
type EVPriceRuleType struct {
	// Required. Cost per kWh.
	EnergyFee float64 `json:"energyFee"`
	// Required. The EnergyFee applies between this value and the value of the PowerRangeStart of the subsequent
	// EVPriceRule. If the power is below this value, the EnergyFee of the previous EVPriceRule applies. Negative
	// values are used for discharging.
	PowerRangeStart float64 `json:"powerRangeStart"`
}

func (s *EVPriceRuleType) UnmarshalJSON(data []byte) error {
	type Alias EVPriceRuleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = EVPriceRuleType(a)
	return s.Validate()
}

func (s EVPriceRuleType) Validate() error {
	_ = s
	return nil
}

// EVSEType (2.47)
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
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if s.ConnectorID != nil {
		if err := validateNonNegative(*s.ConnectorID, "connectorId"); err != nil {
			return err
		}
	}

	return nil
}

// FirmwareType (2.48)
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

	if err := validateStringLen(s.Location, 2000, "location"); err != nil {
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

// FixedPFGetType (2.49)
type FixedPFGetType struct {
	// Required. Id of setting.
	ID IdentifierString36Type `json:"id"`
	// Required. True if setting is a default control.
	IsDefault bool `json:"isDefault"`
	// Required. True if this setting is superseded by a lower priority setting.
	IsSuperseded bool `json:"isSuperseded"`
	// Required. FixedPF for AbsorbW or InjectW
	FixedPF FixedPFType `json:"fixedPF"`
}

func (s *FixedPFGetType) UnmarshalJSON(data []byte) error {
	type Alias FixedPFGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FixedPFGetType(a)
	return s.Validate()
}

func (s FixedPFGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if err := s.FixedPF.Validate(); err != nil {
		return ocpp.WrapField("fixedPF", err)
	}

	return nil
}

// FixedPFType (2.50)
type FixedPFType struct {
	// Required. Priority of setting (0=highest)
	Priority int32 `json:"priority"`
	// Required. Power factor, cos(phi), as value between 0..1.
	Displacement float64 `json:"displacement"`
	// Required. True when absorbing reactive power (under- excited), false when injecting reactive power (over-
	// excited).
	Excitation bool `json:"excitation"`
	// Optional. Time when this setting becomes active
	StartTime *time.Time `json:"startTime,omitempty"`
	// Optional. Duration in seconds that this setting is active.
	Duration *float64 `json:"duration,omitempty"`
}

func (s *FixedPFType) UnmarshalJSON(data []byte) error {
	type Alias FixedPFType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FixedPFType(a)
	return s.Validate()
}

func (s FixedPFType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	return nil
}

// FixedVarGetType (2.51)
type FixedVarGetType struct {
	// Required. Id of setting
	ID IdentifierString36Type `json:"id"`
	// Required. True if setting is a default control.
	IsDefault bool `json:"isDefault"`
	// Required. True if this setting is superseded by a lower priority setting
	IsSuperseded bool `json:"isSuperseded"`
	// Required. Fixed Var setpoint
	FixedVar FixedVarType `json:"fixedVar"`
}

func (s *FixedVarGetType) UnmarshalJSON(data []byte) error {
	type Alias FixedVarGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FixedVarGetType(a)
	return s.Validate()
}

func (s FixedVarGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if err := s.FixedVar.Validate(); err != nil {
		return ocpp.WrapField("fixedVar", err)
	}

	return nil
}

// FixedVarType (2.52)
type FixedVarType struct {
	// Required. Priority of setting (0=highest)
	Priority int32 `json:"priority"`
	// Required. The value specifies a target var output interpreted as a signed percentage (-100 to 100). A negative
	// value refers to charging, whereas a positive one refers to discharging. The value type is determined by the
	// unit field.
	Setpoint float64 `json:"setpoint"`
	// Required. Unit of the setpoint.
	Unit DERUnitEnumType `json:"unit"`
	// Optional. Time when this setting becomes active.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Optional. Duration in seconds that this setting is active.
	Duration *float64 `json:"duration,omitempty"`
}

func (s *FixedVarType) UnmarshalJSON(data []byte) error {
	type Alias FixedVarType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FixedVarType(a)
	return s.Validate()
}

func (s FixedVarType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	if s.Unit == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "unit", "required field is missing")
	}

	return nil
}

// FreqDroopGetType (2.53)
type FreqDroopGetType struct {
	// Required. Id of setting
	ID IdentifierString36Type `json:"id"`
	// Required. True if setting is a default control.
	IsDefault bool `json:"isDefault"`
	// Required. True if this setting is superseded by a higher priority setting (i.e. lower value of priority)
	IsSuperseded bool `json:"isSuperseded"`
	// Required. FreqDroop parameters
	FreqDroop FreqDroopType `json:"freqDroop"`
}

func (s *FreqDroopGetType) UnmarshalJSON(data []byte) error {
	type Alias FreqDroopGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FreqDroopGetType(a)
	return s.Validate()
}

func (s FreqDroopGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if err := s.FreqDroop.Validate(); err != nil {
		return ocpp.WrapField("freqDroop", err)
	}

	return nil
}

// FreqDroopType (2.54)
type FreqDroopType struct {
	// Required. Priority of setting (0=highest)
	Priority int32 `json:"priority"`
	// Required. Over-frequency start of droop
	OverFreq float64 `json:"overFreq"`
	// Required. Under-frequency start of droop
	UnderFreq float64 `json:"underFreq"`
	// Required. Over-frequency droop per unit, oFDroop
	OverDroop float64 `json:"overDroop"`
	// Required. Under-frequency droop per unit, uFDroop
	UnderDroop float64 `json:"underDroop"`
	// Required. Open loop response time in seconds
	ResponseTime float64 `json:"responseTime"`
	// Optional. Time when this setting becomes active
	StartTime *time.Time `json:"startTime,omitempty"`
	// Optional. Duration in seconds that this setting is active
	Duration *float64 `json:"duration,omitempty"`
}

func (s *FreqDroopType) UnmarshalJSON(data []byte) error {
	type Alias FreqDroopType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = FreqDroopType(a)
	return s.Validate()
}

func (s FreqDroopType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	return nil
}

// GetVariableDataType (2.55)
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

// GetVariableResultType (2.56)
//
// Class to hold results of GetVariables request.
type GetVariableResultType struct {
	// Required.
	AttributeStatus GetVariableStatusEnumType `json:"attributeStatus"`
	// Optional.
	AttributeType *AttributeEnumType `json:"attributeType,omitempty"`
	// Optional. Value of requested attribute type of component- variable. This field can only be empty when the
	// given status is NOT accepted. The Configuration Variable ReportingValueSize can be used to limit
	// GetVariableResult.attributeValue, VariableAttribute.value and EventData.actualValue. The max size of these
	// values will always remain equal.
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

// GradientGetType (2.57)
type GradientGetType struct {
	// Required. Id of setting
	ID IdentifierString36Type `json:"id"`
	// Required. Gradient setting
	Gradient GradientType `json:"gradient"`
}

func (s *GradientGetType) UnmarshalJSON(data []byte) error {
	type Alias GradientGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GradientGetType(a)
	return s.Validate()
}

func (s GradientGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if err := s.Gradient.Validate(); err != nil {
		return ocpp.WrapField("gradient", err)
	}

	return nil
}

// GradientType (2.58)
type GradientType struct {
	// Required. Id of setting
	Priority int32 `json:"priority"`
	// Required. Default ramp rate in seconds (0 if not applicable)
	Gradient float64 `json:"gradient"`
	// Required. Soft-start ramp rate in seconds (0 if not applicable)
	SoftGradient float64 `json:"softGradient"`
}

func (s *GradientType) UnmarshalJSON(data []byte) error {
	type Alias GradientType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = GradientType(a)
	return s.Validate()
}

func (s GradientType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	return nil
}

// HysteresisType (2.59)
type HysteresisType struct {
	// Optional. High value for return to normal operation after a grid event, in absolute value. This value adopts
	// the same unit as defined by yUnit
	HysteresisHigh *float64 `json:"hysteresisHigh,omitempty"`
	// Optional. Low value for return to normal operation after a grid event, in absolute value. This value adopts
	// the same unit as defined by yUnit
	HysteresisLow *float64 `json:"hysteresisLow,omitempty"`
	// Optional. Delay in seconds, once grid parameter within HysteresisLow and HysteresisHigh, for the EV to return
	// to normal operation after a grid event.
	HysteresisDelay *float64 `json:"hysteresisDelay,omitempty"`
	// Optional. Set default rate of change (ramp rate %/s) for the EV to return to normal operation after a grid
	// event
	HysteresisGradient *float64 `json:"hysteresisGradient,omitempty"`
}

func (s *HysteresisType) UnmarshalJSON(data []byte) error {
	type Alias HysteresisType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = HysteresisType(a)
	return s.Validate()
}

func (s HysteresisType) Validate() error {
	_ = s
	return nil
}

// IDTokenInfoType (2.60)
//
// Contains status information about an identifier. It is advised to not stop charging for a token that expires
// during charging, as ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the status has
// no end date.
type IDTokenInfoType struct {
	// Required. Current status of the ID Token.
	Status AuthorizationStatusEnumType `json:"status"`
	// Optional. Date and Time after which the token must be considered invalid.
	CacheExpiryDateTime *time.Time `json:"cacheExpiryDateTime,omitempty"`
	// Optional. Priority from a business point of view. Default priority is 0, The range is from -9 to 9. Higher
	// values indicate a higher priority. The chargingPriority in TransactionEventResponse overrules this one.
	ChargingPriority *int32 `json:"chargingPriority,omitempty"`
	// Optional. Preferred user interface language of identifier user. Contains a language code as defined in
	// [RFC5646].
	Language1 *string `json:"language1,omitempty"`
	// Optional. Second preferred user interface language of identifier user. Don’t use when language1 is omitted,
	// has to be different from language1. Contains a language code as defined in [RFC5646].
	Language2 *string `json:"language2,omitempty"`
	// Optional. Only used when the IdToken is only valid for one or more specific EVSEs, not for the entire Charging
	// Station.
	EVSEID []int32 `json:"evseId,omitempty"`
	// Optional. This contains the group identifier.
	GroupIDToken *IDTokenType `json:"groupIdToken,omitempty"`
	// Optional. Personal message that can be shown to the EV Driver and can be used for tariff information, user
	// greetings etc.
	PersonalMessage *MessageContentType `json:"personalMessage,omitempty"`
}

func (s *IDTokenInfoType) UnmarshalJSON(data []byte) error {
	type Alias IDTokenInfoType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = IDTokenInfoType(a)
	return s.Validate()
}

func (s IDTokenInfoType) Validate() error {
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

	for i, v := range s.EVSEID {
		if err := validateNonNegative(v, fmt.Sprintf("evseId[%d]", i)); err != nil {
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

// IDTokenType (2.61)
//
// Contains a case insensitive identifier to use for the authorization and the type of authorization to support
// multiple forms of identifiers.
type IDTokenType struct {
	// Required. (2.1) IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can for example also
	// contain a UUID.
	IDToken IdentifierString255Type `json:"idToken"`
	// Required. (2.1) Enumeration of possible idToken types. Values defined in Appendix as IdTokenEnumStringType.
	Type string `json:"type"`
	// Optional. AdditionalInfo can be used to send extra information which can be validated by the CSMS in addition
	// to the regular authorization with IdToken. AdditionalInfo contains one or more custom types, which need to be
	// agreed upon by all parties involved. When AdditionalInfo is NOT implemented or a not supported
	// AdditionalInfo.type is used, the CSMS/Charging Station MAY ignore the AdditionalInfo.
	AdditionalInfo []AdditionalInfoType `json:"additionalInfo,omitempty"`
}

func (s *IDTokenType) UnmarshalJSON(data []byte) error {
	type Alias IDTokenType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = IDTokenType(a)
	return s.Validate()
}

func (s IDTokenType) Validate() error {
	if s.IDToken == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "idToken", "required field is missing")
	}

	if err := s.IDToken.Validate(); err != nil {
		return ocpp.WrapField("idToken", err)
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateStringLen(s.Type, 20, "type"); err != nil {
		return err
	}

	for i, v := range s.AdditionalInfo {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("additionalInfo[%d]", i), err)
		}
	}

	return nil
}

// LimitAtSOCType (2.62)
type LimitAtSOCType struct {
	// Required. The SoC value beyond which the charging rate limit should be applied.
	SOC int32 `json:"soc"`
	// Required. Charging rate limit beyond the SoC value. The unit is defined by chargingSchedule.chargingRateUnit.
	Limit float64 `json:"limit"`
}

func (s *LimitAtSOCType) UnmarshalJSON(data []byte) error {
	type Alias LimitAtSOCType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = LimitAtSOCType(a)
	return s.Validate()
}

func (s LimitAtSOCType) Validate() error {
	if s.SOC < 0 || s.SOC > 100 {
		return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "soc", "must be between 0 and 100")
	}

	return nil
}

// LimitMaxDischargeGetType (2.63)
type LimitMaxDischargeGetType struct {
	// Required. Id of setting
	ID IdentifierString36Type `json:"id"`
	// Required. True if setting is a default control.
	IsDefault bool `json:"isDefault"`
	// Required. True if this setting is superseded by a higher priority setting (i.e. lower value of priority)
	IsSuperseded bool `json:"isSuperseded"`
	// Required. Maximum discharge power as percentage or rated capability
	LimitMaxDischarge LimitMaxDischargeType `json:"limitMaxDischarge"`
}

func (s *LimitMaxDischargeGetType) UnmarshalJSON(data []byte) error {
	type Alias LimitMaxDischargeGetType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = LimitMaxDischargeGetType(a)
	return s.Validate()
}

func (s LimitMaxDischargeGetType) Validate() error {
	if s.ID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "id", "required field is missing")
	}

	if err := s.ID.Validate(); err != nil {
		return ocpp.WrapField("id", err)
	}

	if err := s.LimitMaxDischarge.Validate(); err != nil {
		return ocpp.WrapField("limitMaxDischarge", err)
	}

	return nil
}

// LimitMaxDischargeType (2.64)
type LimitMaxDischargeType struct {
	// Required. Priority of setting (0=highest)
	Priority int32 `json:"priority"`
	// Optional. Only for PowerMonitoring. The value specifies a percentage (0 to 100) of the rated maximum discharge
	// power of EV. The PowerMonitoring curve becomes active when power exceeds this percentage.
	PctMaxDischargePower *float64 `json:"pctMaxDischargePower,omitempty"`
	// Optional. Time when this setting becomes active
	StartTime *time.Time `json:"startTime,omitempty"`
	// Optional. Duration in seconds that this setting is active
	Duration *float64 `json:"duration,omitempty"`
	// Optional. The curve is an interpolation of data points where the x-axis values are time in seconds and the y-
	// axis values refer to the percentage value of the rated EVMaximumDischargePower, reported in the
	// ChargeParameterDiscoveryRequest message. The value lies between 0 and 100. The curve is activated when the
	// power value measured via the ExternalMeter value reported in the ChargeLoopRes is higher than the
	// pctMaxDischargePower defined above. If the power does not stay within the defined curve for the respective
	// time period, the EV must trip.
	PowerMonitoringMustTrip *DERCurveType `json:"powerMonitoringMustTrip,omitempty"`
}

func (s *LimitMaxDischargeType) UnmarshalJSON(data []byte) error {
	type Alias LimitMaxDischargeType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = LimitMaxDischargeType(a)
	return s.Validate()
}

func (s LimitMaxDischargeType) Validate() error {
	if err := validateNonNegative(s.Priority, "priority"); err != nil {
		return err
	}

	if s.PowerMonitoringMustTrip != nil {
		if err := s.PowerMonitoringMustTrip.Validate(); err != nil {
			return ocpp.WrapField("powerMonitoringMustTrip", err)
		}
	}

	return nil
}

// LogParametersType (2.65)
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

	if err := validateStringLen(s.RemoteLocation, 2000, "remoteLocation"); err != nil {
		return err
	}

	return nil
}

// MessageContentType (2.66)
//
// Contains message details, for a message to be displayed on a Charging Station.
type MessageContentType struct {
	// Required. Format of the message.
	Format MessageFormatEnumType `json:"format"`
	// Optional. Message language identifier. Contains a language code as defined in [RFC5646].
	Language *string `json:"language,omitempty"`
	// Required. (2.1) Required. Message contents. Maximum length supported by Charging Station is given in
	// OCPPCommCtrlr.FieldLength["MessageContentType.content"]. Maximum length defaults to 1024.
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

	if err := validateStringLen(s.Content, 1024, "content"); err != nil {
		return err
	}

	return nil
}

// MessageInfoType (2.67)
//
// Contains message details, for a message to be displayed on a Charging Station.
type MessageInfoType struct {
	// Required. Unique id within an exchange context. It is defined within the OCPP context as a positive Integer
	// value (greater or equal to zero).
	ID int32 `json:"id"`
	// Required. With what priority should this message be shown
	Priority MessagePriorityEnumType `json:"priority"`
	// Optional. During what state should this message be shown. When omitted this message should be shown in any
	// state of the Charging Station.
	State *MessageStateEnumType `json:"state,omitempty"`
	// Optional. From what date-time should this message be shown. If omitted: directly.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Optional. Until what date-time should this message be shown, after this date/time this message SHALL be
	// removed.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Optional. During which transaction shall this message be shown. Message SHALL be removed by the Charging
	// Station after transaction has ended.
	TransactionID *IdentifierString36Type `json:"transactionId,omitempty"`
	// Required. Contains message details for the message to be displayed on a Charging Station.
	Message MessageContentType `json:"message"`
	// Optional. When a Charging Station has multiple Displays, this field can be used to define to which Display
	// this message belongs.
	Display *ComponentType `json:"display,omitempty"`
	// Optional. (2.1) Contains message details for extra languages to be displayed on a Charging Station.
	MessageExtra []MessageContentType `json:"messageExtra,omitempty"`
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
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

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

	if err := validateSliceLen(s.MessageExtra, 0, 4, "messageExtra"); err != nil {
		return err
	}

	for i, v := range s.MessageExtra {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("messageExtra[%d]", i), err)
		}
	}

	return nil
}

// MeterValueType (2.68)
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

// ModemType (2.69)
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

// MonitoringDataType (2.70)
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

// NetworkConnectionProfileType (2.71)
//
// The NetworkConnectionProfile defines the functional and technical parameters of a communication link.
type NetworkConnectionProfileType struct {
	// Optional. (2.1) This field is ignored, since the OCPP version to use is determined during the websocket
	// handshake. The field is only kept for backwards compatibility with the OCPP 2.0.1 JSON schema.
	OCPPVersion *OCPPVersionEnumType `json:"ocppVersion,omitempty"`
	// Required. Applicable Network Interface. Charging Station is allowed to use a different network interface to
	// connect if the given one does not work.
	OCPPInterface OCPPInterfaceEnumType `json:"ocppInterface"`
	// Required. Defines the transport protocol (e.g. SOAP or JSON). Note: SOAP is not supported in OCPP 2.x, but is
	// supported by earlier versions of OCPP.
	OCPPTransport OCPPTransportEnumType `json:"ocppTransport"`
	// Required. Duration in seconds before a message send by the Charging Station via this network connection times-
	// out. The best setting depends on the underlying network and response times of the CSMS. If you are looking for
	// a some guideline: use 30 seconds as a starting point.
	MessageTimeout int32 `json:"messageTimeout"`
	// Required. URL of the CSMS(s) that this Charging Station communicates with, without the Charging Station
	// identity part. The SecurityCtrlr.Identity field is appended to ocppCsmsUrl to provide the full websocket URL.
	OCPPCSMSURL string `json:"ocppCsmsUrl"`
	// Required. This field specifies the security profile used when connecting to the CSMS with this
	// NetworkConnectionProfile.
	SecurityProfile int32 `json:"securityProfile"`
	// Optional. (2.1) Charging Station identity to be used as the basic authentication username.
	Identity *string `json:"identity,omitempty"`
	// Optional. (2.1) BasicAuthPassword to use for security profile 1 or 2.
	BasicAuthPassword *string `json:"basicAuthPassword,omitempty"`
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
	if s.OCPPInterface == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppInterface", "required field is missing")
	}

	if s.OCPPTransport == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppTransport", "required field is missing")
	}

	if s.OCPPCSMSURL == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "ocppCsmsUrl", "required field is missing")
	}

	if err := validateStringLen(s.OCPPCSMSURL, 2000, "ocppCsmsUrl"); err != nil {
		return err
	}

	if err := validateNonNegative(s.SecurityProfile, "securityProfile"); err != nil {
		return err
	}

	if s.Identity != nil {
		if err := validateStringLen(*s.Identity, 48, "identity"); err != nil {
			return err
		}
	}

	if s.BasicAuthPassword != nil {
		if err := validateStringLen(*s.BasicAuthPassword, 64, "basicAuthPassword"); err != nil {
			return err
		}
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

// OCSPRequestDataType (2.72)
//
// Information about a certificate for an OCSP check.
type OCSPRequestDataType struct {
	// Required. Used algorithms for the hashes provided.
	HashAlgorithm HashAlgorithmEnumType `json:"hashAlgorithm"`
	// Required. The hash of the issuer’s distinguished name (DN), that must be calculated over the DER encoding of
	// the issuer’s name field in the certificate being checked.
	IssuerNameHash IdentifierString128Type `json:"issuerNameHash"`
	// Required. The hash of the DER encoded public key: the value (excluding tag and length) of the subject public
	// key field in the issuer’s certificate.
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

	if err := validateStringLen(s.ResponderURL, 2000, "responderURL"); err != nil {
		return err
	}

	return nil
}

// OverstayRuleListType (2.73)
//
// Part of ISO 15118-20 price schedule.
type OverstayRuleListType struct {
	// Optional. Time till overstay is applied in seconds.
	OverstayTimeThreshold *int32 `json:"overstayTimeThreshold,omitempty"`
	// Optional. Power threshold in W at which the overstay applies.
	OverstayPowerThreshold *RationalNumberType `json:"overstayPowerThreshold,omitempty"`
	// Required. Overstay rules that will be applied.
	OverstayRule []OverstayRuleType `json:"overstayRule"`
}

func (s *OverstayRuleListType) UnmarshalJSON(data []byte) error {
	type Alias OverstayRuleListType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = OverstayRuleListType(a)
	return s.Validate()
}

func (s OverstayRuleListType) Validate() error {
	if s.OverstayPowerThreshold != nil {
		if err := s.OverstayPowerThreshold.Validate(); err != nil {
			return ocpp.WrapField("overstayPowerThreshold", err)
		}
	}

	if err := validateSliceLen(s.OverstayRule, 1, 5, "overstayRule"); err != nil {
		return err
	}

	for i, v := range s.OverstayRule {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("overstayRule[%d]", i), err)
		}
	}

	return nil
}

// OverstayRuleType (2.74)
//
// Part of ISO 15118-20 price schedule.
type OverstayRuleType struct {
	// Optional. Human readable string to identify the overstay rule.
	OverstayRuleDescription *string `json:"overstayRuleDescription,omitempty"`
	// Required. Time in seconds after trigger of the parent Overstay Rules for this particular fee to apply.
	StartTime int32 `json:"startTime"`
	// Required. Time till overstay will be reapplied
	OverstayFeePeriod int32 `json:"overstayFeePeriod"`
	// Required. Fee that applies to this overstay.
	OverstayFee RationalNumberType `json:"overstayFee"`
}

func (s *OverstayRuleType) UnmarshalJSON(data []byte) error {
	type Alias OverstayRuleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = OverstayRuleType(a)
	return s.Validate()
}

func (s OverstayRuleType) Validate() error {
	if s.OverstayRuleDescription != nil {
		if err := validateStringLen(*s.OverstayRuleDescription, 32, "overstayRuleDescription"); err != nil {
			return err
		}
	}

	if err := s.OverstayFee.Validate(); err != nil {
		return ocpp.WrapField("overstayFee", err)
	}

	return nil
}

// PeriodicEventStreamParamsType (2.75)
type PeriodicEventStreamParamsType struct {
	// Optional. Time in seconds after which stream data is sent.
	Interval *int32 `json:"interval,omitempty"`
	// Optional. Number of items to be sent together in stream.
	Values *int32 `json:"values,omitempty"`
}

func (s *PeriodicEventStreamParamsType) UnmarshalJSON(data []byte) error {
	type Alias PeriodicEventStreamParamsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PeriodicEventStreamParamsType(a)
	return s.Validate()
}

func (s PeriodicEventStreamParamsType) Validate() error {
	if s.Interval != nil {
		if err := validateNonNegative(*s.Interval, "interval"); err != nil {
			return err
		}
	}

	if s.Values != nil {
		if err := validateNonNegative(*s.Values, "values"); err != nil {
			return err
		}
	}

	return nil
}

// PriceLevelScheduleEntryType (2.76)
//
// Part of ISO 15118-20 price schedule.
type PriceLevelScheduleEntryType struct {
	// Required. The amount of seconds that define the duration of this given PriceLevelScheduleEntry.
	Duration int32 `json:"duration"`
	// Required. Defines the price level of this PriceLevelScheduleEntry (referring to NumberOfPriceLevels). Small
	// values for the PriceLevel represent a cheaper PriceLevelScheduleEntry. Large values for the PriceLevel
	// represent a more expensive PriceLevelScheduleEntry.
	PriceLevel int32 `json:"priceLevel"`
}

func (s *PriceLevelScheduleEntryType) UnmarshalJSON(data []byte) error {
	type Alias PriceLevelScheduleEntryType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PriceLevelScheduleEntryType(a)
	return s.Validate()
}

func (s PriceLevelScheduleEntryType) Validate() error {
	if err := validateNonNegative(s.PriceLevel, "priceLevel"); err != nil {
		return err
	}

	return nil
}

// PriceLevelScheduleType (2.77)
//
// The PriceLevelScheduleType is modeled after the same type that is defined in ISO 15118-20, such that if it is
// supplied by an EMSP as a signed EXI message, the conversion from EXI to JSON (in OCPP) and back to EXI (for
// ISO 15118-20) does not change the digest and therefore does not invalidate the signature.
type PriceLevelScheduleType struct {
	// Required. Starting point of this price schedule.
	TimeAnchor time.Time `json:"timeAnchor"`
	// Required. Unique ID of this price schedule.
	PriceScheduleID int32 `json:"priceScheduleId"`
	// Optional. Description of the price schedule.
	PriceScheduleDescription *string `json:"priceScheduleDescription,omitempty"`
	// Required. Defines the overall number of distinct price level elements used across all PriceLevelSchedules.
	NumberOfPriceLevels int32 `json:"numberOfPriceLevels"`
	// Required. List of entries of the schedule.
	PriceLevelScheduleEntries []PriceLevelScheduleEntryType `json:"priceLevelScheduleEntries"`
}

func (s *PriceLevelScheduleType) UnmarshalJSON(data []byte) error {
	type Alias PriceLevelScheduleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PriceLevelScheduleType(a)
	return s.Validate()
}

func (s PriceLevelScheduleType) Validate() error {
	if s.TimeAnchor.IsZero() {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "timeAnchor", "required field is missing")
	}

	if err := validateNonNegative(s.PriceScheduleID, "priceScheduleId"); err != nil {
		return err
	}

	if s.PriceScheduleDescription != nil {
		if err := validateStringLen(*s.PriceScheduleDescription, 32, "priceScheduleDescription"); err != nil {
			return err
		}
	}

	if err := validateNonNegative(s.NumberOfPriceLevels, "numberOfPriceLevels"); err != nil {
		return err
	}

	if err := validateSliceLen(s.PriceLevelScheduleEntries, 1, 100, "priceLevelScheduleEntries"); err != nil {
		return err
	}

	for i, v := range s.PriceLevelScheduleEntries {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("priceLevelScheduleEntries[%d]", i), err)
		}
	}

	return nil
}

// PriceRuleStackType (2.78)
//
// Part of ISO 15118-20 price schedule.
type PriceRuleStackType struct {
	// Required. Duration of the stack of price rules. he amount of seconds that define the duration of the given
	// PriceRule(s).
	Duration int32 `json:"duration"`
	// Required. Contains the price rules.
	PriceRule []PriceRuleType `json:"priceRule"`
}

func (s *PriceRuleStackType) UnmarshalJSON(data []byte) error {
	type Alias PriceRuleStackType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PriceRuleStackType(a)
	return s.Validate()
}

func (s PriceRuleStackType) Validate() error {
	if err := validateSliceLen(s.PriceRule, 1, 8, "priceRule"); err != nil {
		return err
	}

	for i, v := range s.PriceRule {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("priceRule[%d]", i), err)
		}
	}

	return nil
}

// PriceRuleType (2.79)
//
// Part of ISO 15118-20 price schedule.
type PriceRuleType struct {
	// Optional. The duration of the parking fee period (in seconds). When the time enters into a ParkingFeePeriod,
	// the ParkingFee will apply to the session. .
	ParkingFeePeriod *int32 `json:"parkingFeePeriod,omitempty"`
	// Optional. Number of grams of CO2 per kWh.
	CarbonDioxideEmission *int32 `json:"carbonDioxideEmission,omitempty"`
	// Optional. Percentage of the power that is created by renewable resources.
	RenewableGenerationPercentage *int32 `json:"renewableGenerationPercentage,omitempty"`
	// Required. Cost per kWh. Use zero for free energy.
	EnergyFee RationalNumberType `json:"energyFee"`
	// Optional. Optional. Cost of parking. Mandatory whenever a parking fee applies.
	ParkingFee *RationalNumberType `json:"parkingFee,omitempty"`
	// Required. For values 0 and above, this is the power level above which this price rule applies. If there is
	// another PriceRule with a higher value, and the current power is above that value, then that other PriceRule
	// applies.For negative values, this is the power level below which this price rule applies. If there is another
	// PriceRule with a lower value, and the current power is below that value, then that other PriceRule applies.
	PowerRangeStart RationalNumberType `json:"powerRangeStart"`
}

func (s *PriceRuleType) UnmarshalJSON(data []byte) error {
	type Alias PriceRuleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PriceRuleType(a)
	return s.Validate()
}

func (s PriceRuleType) Validate() error {
	if s.CarbonDioxideEmission != nil {
		if err := validateNonNegative(*s.CarbonDioxideEmission, "carbonDioxideEmission"); err != nil {
			return err
		}
	}

	if s.RenewableGenerationPercentage != nil {
		if *s.RenewableGenerationPercentage < 0 || *s.RenewableGenerationPercentage > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "renewableGenerationPercentage", "must be between 0 and 100")
		}
	}

	if err := s.EnergyFee.Validate(); err != nil {
		return ocpp.WrapField("energyFee", err)
	}

	if s.ParkingFee != nil {
		if err := s.ParkingFee.Validate(); err != nil {
			return ocpp.WrapField("parkingFee", err)
		}
	}

	if err := s.PowerRangeStart.Validate(); err != nil {
		return ocpp.WrapField("powerRangeStart", err)
	}

	return nil
}

// PriceType (2.80)
//
// Price with and without tax. At least one of exclTax, inclTax must be present.
type PriceType struct {
	// Optional. Price/cost excluding tax. Can be absent if inclTax is present.
	ExclTax *float64 `json:"exclTax,omitempty"`
	// Optional. Price/cost including tax. Can be absent if exclTax is present.
	InclTax *float64 `json:"inclTax,omitempty"`
	// Optional. Tax percentages that were used to calculate inclTax from exclTax (for displaying/printing on
	// invoices).
	TaxRates []TaxRateType `json:"taxRates,omitempty"`
}

func (s *PriceType) UnmarshalJSON(data []byte) error {
	type Alias PriceType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = PriceType(a)
	return s.Validate()
}

func (s PriceType) Validate() error {
	if err := validateSliceLen(s.TaxRates, 0, 5, "taxRates"); err != nil {
		return err
	}

	for i, v := range s.TaxRates {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("taxRates[%d]", i), err)
		}
	}

	return nil
}

// RationalNumberType (2.81)
//
// Part of ISO 15118-20 price schedule.
type RationalNumberType struct {
	// Required. The exponent to base 10 (dec)
	Exponent int32 `json:"exponent"`
	// Required. Value which shall be multiplied.
	Value int32 `json:"value"`
}

func (s *RationalNumberType) UnmarshalJSON(data []byte) error {
	type Alias RationalNumberType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = RationalNumberType(a)
	return s.Validate()
}

func (s RationalNumberType) Validate() error {
	_ = s
	return nil
}

// ReactivePowerParamsType (2.82)
type ReactivePowerParamsType struct {
	// Optional. Only for VoltVar curve: The nominal ac voltage (rms) adjustment to the voltage curve points for
	// Volt-Var curves (percentage).
	VRef *float64 `json:"vRef,omitempty"`
	// Optional. Only for VoltVar: Enable/disable autonomous VRef adjustment
	AutonomousVRefEnable *bool `json:"autonomousVRefEnable,omitempty"`
	// Optional. Only for VoltVar: Adjustment range for VRef time constant
	AutonomousVRefTimeConstant *float64 `json:"autonomousVRefTimeConstant,omitempty"`
}

func (s *ReactivePowerParamsType) UnmarshalJSON(data []byte) error {
	type Alias ReactivePowerParamsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = ReactivePowerParamsType(a)
	return s.Validate()
}

func (s ReactivePowerParamsType) Validate() error {
	_ = s
	return nil
}

// RelativeTimeIntervalType (2.83)
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

// ReportDataType (2.84)
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

// SalesTariffEntryType (2.85)
type SalesTariffEntryType struct {
	// Optional. Defines the price level of this SalesTariffEntry (referring to NumEPriceLevels). Small values for
	// the EPriceLevel represent a cheaper TariffEntry. Large values for the EPriceLevel represent a more expensive
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

// SalesTariffType (2.86)
//
// A SalesTariff provided by a Mobility Operator (EMSP) . NOTE: This dataType is based on dataTypes from ISO
// 15118-2.
type SalesTariffType struct {
	// Required. SalesTariff identifier used to identify one sales tariff. An SAID remains a unique identifier for
	// one schedule throughout a charging session.
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
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if s.SalesTariffDescription != nil {
		if err := validateStringLen(*s.SalesTariffDescription, 32, "salesTariffDescription"); err != nil {
			return err
		}
	}

	if s.NumEPriceLevels != nil {
		if err := validateNonNegative(*s.NumEPriceLevels, "numEPriceLevels"); err != nil {
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

// SampledValueType (2.87)
//
// Single sampled value in MeterValues. Each value can be accompanied by optional fields.
//
// To save on mobile data usage, default values of all of the optional fields are such that. The value without
// any additional fields will be interpreted, as a register reading of active import energy in Wh (Watt-hour)
// units.
type SampledValueType struct {
	// Required. Indicates the measured value.
	Value float64 `json:"value"`
	// Optional. Type of measurement. Default = "Energy.Active.Import.Register"
	Measurand *MeasurandEnumType `json:"measurand,omitempty"`
	// Optional. Type of detail value: start, end or sample. Default = "Sample.Periodic"
	Context *ReadingContextEnumType `json:"context,omitempty"`
	// Optional. Indicates how the measured value is to be interpreted. For instance between L1 and neutral (L1-N)
	// Please note that not all values of phase are applicable to all Measurands. When phase is absent, the measured
	// value is interpreted as an overall value.
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

// SetMonitoringDataType (2.88)
//
// Class to hold parameters of SetVariableMonitoring request.
type SetMonitoringDataType struct {
	// Optional. An id SHALL only be given to replace an existing monitor. The Charging Station handles the
	// generation of id’s for new monitors.
	ID *int32 `json:"id,omitempty"`
	// Optional. Monitor only active when a transaction is ongoing on a component relevant to this transaction.
	// Default = false.
	Transaction *bool `json:"transaction,omitempty"`
	// Required. Value for threshold or delta monitoring. For Periodic or PeriodicClockAligned this is the interval
	// in seconds.
	Value float64 `json:"value"`
	// Required. The type of this monitor, e.g. a threshold, delta or periodic monitor.
	Type MonitorEnumType `json:"type"`
	// Required. The severity that will be assigned to an event that is triggered by this monitor. The severity range
	// is 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels have the following
	// meaning: 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and action should be
	// taken immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue regular
	// operations due to Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station is
	// unable to continue regular operations due to software or minor hardware issues. Action is required. 3-Critical
	// Indicates a critical error. Action is required. 4-Error Indicates a non-urgent error. Action is required.
	// 5-Alert Indicates an alert event. Default severity for any type of monitoring event. 6-Warning Indicates a
	// warning event. Action may be required. 7-Notice Indicates an unusual event. No immediate action is required.
	// 8-Informational Indicates a regular operational event. May be used for reporting, measuring throughput, etc.
	// No action is required. 9-Debug Indicates information useful to developers for debugging, not useful during
	// operations.
	Severity int32 `json:"severity"`
	// Required. Component for which monitor is set.
	Component ComponentType `json:"component"`
	// Required. Variable for which monitor is set.
	Variable VariableType `json:"variable"`
	// Optional. (2.1) Optional. When present, events from a monitor will be sent via a periodic event stream. Used
	// for monitors of type Periodic, PeriodicClockAligned or Delta.
	PeriodicEventStream *PeriodicEventStreamParamsType `json:"periodicEventStream,omitempty"`
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
	if s.ID != nil {
		if err := validateNonNegative(*s.ID, "id"); err != nil {
			return err
		}
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateNonNegative(s.Severity, "severity"); err != nil {
		return err
	}

	if err := s.Component.Validate(); err != nil {
		return ocpp.WrapField("component", err)
	}

	if err := s.Variable.Validate(); err != nil {
		return ocpp.WrapField("variable", err)
	}

	if s.PeriodicEventStream != nil {
		if err := s.PeriodicEventStream.Validate(); err != nil {
			return ocpp.WrapField("periodicEventStream", err)
		}
	}

	return nil
}

// SetMonitoringResultType (2.89)
//
// Class to hold result of SetVariableMonitoring request.
type SetMonitoringResultType struct {
	// Optional. Id given to the VariableMonitor by the Charging Station. The Id is only returned when status is
	// accepted. Installed VariableMonitors should have unique id’s but the id’s of removed Installed monitors should
	// have unique id’s but the id’s of removed monitors MAY be reused.
	ID *int32 `json:"id,omitempty"`
	// Required. Status is OK if a value could be returned. Otherwise this will indicate the reason why a value could
	// not be returned.
	Status SetMonitoringStatusEnumType `json:"status"`
	// Required. The type of this monitor, e.g. a threshold, delta or periodic monitor.
	Type MonitorEnumType `json:"type"`
	// Required. The severity that will be assigned to an event that is triggered by this monitor. The severity range
	// is 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels have the following
	// meaning: 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and action should be
	// taken immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue regular
	// operations due to Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station is
	// unable to continue regular operations due to software or minor hardware issues. Action is required. 3-Critical
	// Indicates a critical error. Action is required. 4-Error Indicates a non-urgent error. Action is required.
	// 5-Alert Indicates an alert event. Default severity for any type of monitoring event. 6-Warning Indicates a
	// warning event. Action may be required. 7-Notice Indicates an unusual event. No immediate action is required.
	// 8-Informational Indicates a regular operational event. May be used for reporting, measuring throughput, etc.
	// No action is required. 9-Debug Indicates information useful to developers for debugging, not useful during
	// operations.
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
	if s.ID != nil {
		if err := validateNonNegative(*s.ID, "id"); err != nil {
			return err
		}
	}

	if s.Status == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "status", "required field is missing")
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateNonNegative(s.Severity, "severity"); err != nil {
		return err
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

// SetVariableDataType (2.90)
type SetVariableDataType struct {
	// Optional. Type of attribute: Actual, Target, MinSet, MaxSet. Default is Actual when omitted.
	AttributeType *AttributeEnumType `json:"attributeType,omitempty"`
	// Required. Value to be assigned to attribute of variable. This value is allowed to be an empty string (""). The
	// Configuration Variable ConfigurationValueSize can be used to limit SetVariableData.attributeValue and
	// VariableCharacteristics.valuesList. The max size of these values will always remain equal.
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

	if err := validateStringLen(s.AttributeValue, 2500, "attributeValue"); err != nil {
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

// SetVariableResultType (2.91)
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

// SignedMeterValueType (2.92)
//
// Represent a signed version of the meter value.
type SignedMeterValueType struct {
	// Required. Base64 encoded, contains the signed data from the meter in the format specified in encodingMethod,
	// which might contain more then just the meter value. It can contain information like timestamps, reference to a
	// customer etc.
	SignedMeterData string `json:"signedMeterData"`
	// Optional. (2.1) Method used to create the digital signature. Optional, if already included in signedMeterData.
	// Standard values for this are defined in Appendix as SigningMethodEnumStringType.
	SigningMethod *string `json:"signingMethod,omitempty"`
	// Required. Format used by the energy meter to encode the meter data. For example: OCMF or EDL.
	EncodingMethod string `json:"encodingMethod"`
	// Optional. (2.1) Base64 encoded, sending depends on configuration variable PublicKeyWithSignedMeterValue.
	PublicKey *string `json:"publicKey,omitempty"`
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

	if err := validateStringLen(s.SignedMeterData, 32768, "signedMeterData"); err != nil {
		return err
	}

	if s.SigningMethod != nil {
		if err := validateStringLen(*s.SigningMethod, 50, "signingMethod"); err != nil {
			return err
		}
	}

	if s.EncodingMethod == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "encodingMethod", "required field is missing")
	}

	if err := validateStringLen(s.EncodingMethod, 50, "encodingMethod"); err != nil {
		return err
	}

	if s.PublicKey != nil {
		if err := validateStringLen(*s.PublicKey, 2500, "publicKey"); err != nil {
			return err
		}
	}

	return nil
}

// StatusInfoType (2.93)
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
		if err := validateStringLen(*s.AdditionalInfo, 1024, "additionalInfo"); err != nil {
			return err
		}
	}

	return nil
}

// StreamDataElementType (2.94)
type StreamDataElementType struct {
	// Required. Offset relative to basetime of this message. basetime + t is timestamp of recorded value.
	Offset float64 `json:"t"`
	// Required.
	Value string `json:"v"`
}

func (s *StreamDataElementType) UnmarshalJSON(data []byte) error {
	type Alias StreamDataElementType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = StreamDataElementType(a)
	return s.Validate()
}

func (s StreamDataElementType) Validate() error {
	if s.Value == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "v", "required field is missing")
	}

	if err := validateStringLen(s.Value, 2500, "v"); err != nil {
		return err
	}

	return nil
}

// TariffAssignmentType (2.95)
//
// Shows assignment of tariffs to EVSE or IdToken.
type TariffAssignmentType struct {
	// Required. Tariff id.
	TariffID string `json:"tariffId"`
	// Required. Kind of tariff (driver/default)
	TariffKind TariffKindEnumType `json:"tariffKind"`
	// Optional. Date/time when this tariff become active.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Optional.
	EVSEIds []int32 `json:"evseIds,omitempty"`
	// Optional. IdTokens related to tariff
	IDTokens []IdentifierString255Type `json:"idTokens,omitempty"`
}

func (s *TariffAssignmentType) UnmarshalJSON(data []byte) error {
	type Alias TariffAssignmentType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffAssignmentType(a)
	return s.Validate()
}

func (s TariffAssignmentType) Validate() error {
	if s.TariffID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "tariffId", "required field is missing")
	}

	if err := validateStringLen(s.TariffID, 60, "tariffId"); err != nil {
		return err
	}

	if s.TariffKind == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "tariffKind", "required field is missing")
	}

	for i, v := range s.EVSEIds {
		if err := validateNonNegative(v, fmt.Sprintf("evseIds[%d]", i)); err != nil {
			return err
		}
	}

	for i, v := range s.IDTokens {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("idTokens[%d]", i), err)
		}
	}

	return nil
}

// TariffConditionsFixedType (2.96)
//
// These conditions describe if a FixedPrice applies at start of the transaction.
//
// When more than one restriction is set, they are to be treated as a logical AND. All need to be valid before
// this price is active.
//
// startTimeOfDay and endTimeOfDay are in local time, because it is the time in the tariff as it is shown to the
// EV NOTE        driver at the Charging Station. A Charging Station will convert this to the internal time zone
// that it uses (which is recommended to be UTC, see section Generic chapter 3.1) when performing cost
// calculation.
type TariffConditionsFixedType struct {
	// Optional. Start time of day in local time. Format as per RFC 3339: time-hour ":" time-minute Must be in 24h
	// format with leading zeros. Hour/Minute separator: ":" Regex: ([0-1][0-9]|2[0-3]):[0-5][0-9]
	StartTimeOfDay *string `json:"startTimeOfDay,omitempty"`
	// Optional. End time of day in local time. Same syntax as startTimeOfDay. If end time < start time then the
	// period wraps around to the next day. To stop at end of the day use: 00:00.
	EndTimeOfDay *string `json:"endTimeOfDay,omitempty"`
	// Optional. Day(s) of the week this is tariff applies.
	DayOfWeek []DayOfWeekEnumType `json:"dayOfWeek,omitempty"`
	// Optional. Start date in local time, for example: 2015-12- 24. Valid from this day (inclusive). Format as per
	// RFC 3339: full-date Regex: ([12][0-9]{3})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])
	ValidFromDate *string `json:"validFromDate,omitempty"`
	// Optional. End date in local time, for example: 2015-12-27. Valid until this day (exclusive). Same syntax as
	// validFromDate.
	ValidToDate *string `json:"validToDate,omitempty"`
	// Optional. Type of EVSE (AC, DC) this tariff applies to.
	EVSEKind *EVSEKindEnumType `json:"evseKind,omitempty"`
	// Optional. For which payment brand this (adhoc) tariff applies. Can be used to add a surcharge for certain
	// payment brands. Based on value of additionalIdToken from idToken.additionalInfo.type = "PaymentBrand".
	PaymentBrand *string `json:"paymentBrand,omitempty"`
	// Optional. Type of adhoc payment, e.g. CC, Debit. Based on value of additionalIdToken from
	// idToken.additionalInfo.type = "PaymentRecognition".
	PaymentRecognition *string `json:"paymentRecognition,omitempty"`
}

func (s *TariffConditionsFixedType) UnmarshalJSON(data []byte) error {
	type Alias TariffConditionsFixedType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffConditionsFixedType(a)
	return s.Validate()
}

func (s TariffConditionsFixedType) Validate() error {
	if err := validateSliceLen(s.DayOfWeek, 0, 7, "dayOfWeek"); err != nil {
		return err
	}

	if s.PaymentBrand != nil {
		if err := validateStringLen(*s.PaymentBrand, 20, "paymentBrand"); err != nil {
			return err
		}
	}

	if s.PaymentRecognition != nil {
		if err := validateStringLen(*s.PaymentRecognition, 20, "paymentRecognition"); err != nil {
			return err
		}
	}

	return nil
}

// TariffConditionsType (2.97)
//
// These conditions describe if and when a TariffEnergyType or TariffTimeType applies during a transaction.
//
// When more than one restriction is set, they are to be treated as a logical AND. All need to be valid before
// this price is active.
//
// For reverse energy flow (discharging) negative values of energy, power and current are used.
//
// minXXX (where XXX = Kwh/A/Kw) must be read as "closest to zero", and maxXXX as "furthest from zero". For NOTE
// example, a charging power range from 10 kW to 50 kWh is given by minPower = 10000 and maxPower = 50000, and a
// discharging power range from -10 kW to -50 kW is given by minPower = -10 and maxPower = -50.
//
// startTimeOfDay and endTimeOfDay are in local time, because it is the time in the tariff as it is shown to the
// EV NOTE        driver at the Charging Station. A Charging Station will convert this to the internal time zone
// that it uses (which is recommended to be UTC, see section Generic chapter 3.1) when performing cost
// calculation.
type TariffConditionsType struct {
	// Optional. Start time of day in local time. Format as per RFC 3339: time-hour ":" time-minute Must be in 24h
	// format with leading zeros. Hour/Minute separator: ":" Regex: ([0-1][0-9]|2[0-3]):[0-5][0-9]
	StartTimeOfDay *string `json:"startTimeOfDay,omitempty"`
	// Optional. End time of day in local time. Same syntax as startTimeOfDay. If end time < start time then the
	// period wraps around to the next day. To stop at end of the day use: 00:00.
	EndTimeOfDay *string `json:"endTimeOfDay,omitempty"`
	// Optional. Day(s) of the week this is tariff applies.
	DayOfWeek []DayOfWeekEnumType `json:"dayOfWeek,omitempty"`
	// Optional. Start date in local time, for example: 2015-12- 24. Valid from this day (inclusive). Format as per
	// RFC 3339: full-date Regex: ([12][0-9]{3})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])
	ValidFromDate *string `json:"validFromDate,omitempty"`
	// Optional. End date in local time, for example: 2015-12-27. Valid until this day (exclusive). Same syntax as
	// validFromDate.
	ValidToDate *string `json:"validToDate,omitempty"`
	// Optional. Type of EVSE (AC, DC) this tariff applies to.
	EVSEKind *EVSEKindEnumType `json:"evseKind,omitempty"`
	// Optional. Minimum consumed energy in Wh, for example 20000 Wh. Valid from this amount of energy (inclusive)
	// being used.
	MinEnergy *float64 `json:"minEnergy,omitempty"`
	// Optional. Maximum consumed energy in Wh, for example 50000 Wh. Valid until this amount of energy (exclusive)
	// being used.
	MaxEnergy *float64 `json:"maxEnergy,omitempty"`
	// Optional. Sum of the minimum current (in Amperes) over all phases, for example 5 A. When the EV is charging
	// with more than, or equal to, the defined amount of current, this price is/becomes active. If the charging
	// current is or becomes lower, this price is not or no longer valid and becomes inactive. This is NOT about the
	// minimum current over the entire transaction.
	MinCurrent *float64 `json:"minCurrent,omitempty"`
	// Optional. Sum of the maximum current (in Amperes) over all phases, for example 20 A. When the EV is charging
	// with less than the defined amount of current, this price becomes/is active. If the charging current is or
	// becomes higher, this price is not or no longer valid and becomes inactive. This is NOT about the maximum
	// current over the entire transaction.
	MaxCurrent *float64 `json:"maxCurrent,omitempty"`
	// Optional. Minimum power in W, for example 5000 W. When the EV is charging with more than, or equal to, the
	// defined amount of power, this price is/becomes active. If the charging power is or becomes lower, this price
	// is not or no longer valid and becomes inactive. This is NOT about the minimum power over the entire
	// transaction.
	MinPower *float64 `json:"minPower,omitempty"`
	// Optional. Maximum power in W, for example 20000 W. When the EV is charging with less than the defined amount
	// of power, this price becomes/is active. If the charging power is or becomes higher, this price is not or no
	// longer valid and becomes inactive. This is NOT about the maximum power over the entire transaction.
	MaxPower *float64 `json:"maxPower,omitempty"`
	// Optional. Minimum duration in seconds the transaction (charging & idle) MUST last (inclusive). When the
	// duration of a transaction is longer than the defined value, this price is or becomes active. Before that
	// moment, this price is not yet active.
	MinTime *int32 `json:"minTime,omitempty"`
	// Optional. Maximum duration in seconds the transaction (charging & idle) MUST last (exclusive). When the
	// duration of a transaction is shorter than the defined value, this price is or becomes active. After that
	// moment, this price is no longer active.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// Optional. Minimum duration in seconds the charging MUST last (inclusive). When the duration of a charging is
	// longer than the defined value, this price is or becomes active. Before that moment, this price is not yet
	// active.
	MinChargingTime *int32 `json:"minChargingTime,omitempty"`
	// Optional. Maximum duration in seconds the charging MUST last (exclusive). When the duration of a charging is
	// shorter than the defined value, this price is or becomes active. After that moment, this price is no longer
	// active.
	MaxChargingTime *int32 `json:"maxChargingTime,omitempty"`
	// Optional. Minimum duration in seconds the idle period (i.e. not charging) MUST last (inclusive). When the
	// duration of the idle time is longer than the defined value, this price is or becomes active. Before that
	// moment, this price is not yet active.
	MinIdleTime *int32 `json:"minIdleTime,omitempty"`
	// Optional. Maximum duration in seconds the idle period (i.e. not charging) MUST last (exclusive). When the
	// duration of idle time is shorter than the defined value, this price is or becomes active. After that moment,
	// this price is no longer active.
	MaxIdleTime *int32 `json:"maxIdleTime,omitempty"`
}

func (s *TariffConditionsType) UnmarshalJSON(data []byte) error {
	type Alias TariffConditionsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffConditionsType(a)
	return s.Validate()
}

func (s TariffConditionsType) Validate() error {
	if err := validateSliceLen(s.DayOfWeek, 0, 7, "dayOfWeek"); err != nil {
		return err
	}

	return nil
}

// TariffEnergyPriceType (2.98)
//
// Tariff with optional conditions for an energy price.
type TariffEnergyPriceType struct {
	// Required. Price per kWh (excl. tax) for this element.
	PriceKwh float64 `json:"priceKwh"`
	// Optional. Conditions when this tariff element price is applicable. When absent always applicable,
	Conditions *TariffConditionsType `json:"conditions,omitempty"`
}

func (s *TariffEnergyPriceType) UnmarshalJSON(data []byte) error {
	type Alias TariffEnergyPriceType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffEnergyPriceType(a)
	return s.Validate()
}

func (s TariffEnergyPriceType) Validate() error {
	if s.Conditions != nil {
		if err := s.Conditions.Validate(); err != nil {
			return ocpp.WrapField("conditions", err)
		}
	}

	return nil
}

// TariffEnergyType (2.99)
//
// Price elements and tax for energy
type TariffEnergyType struct {
	// Optional. Applicable tax percentages for this tariff dimension. If omitted, no tax is applicable. Not
	// providing a tax is different from 0% tax, which would be a value of 0.0 here.
	TaxRates []TaxRateType `json:"taxRates,omitempty"`
	// Required. Element tariff price and conditions
	Prices []TariffEnergyPriceType `json:"prices"`
}

func (s *TariffEnergyType) UnmarshalJSON(data []byte) error {
	type Alias TariffEnergyType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffEnergyType(a)
	return s.Validate()
}

func (s TariffEnergyType) Validate() error {
	if err := validateSliceLen(s.TaxRates, 0, 5, "taxRates"); err != nil {
		return err
	}

	for i, v := range s.TaxRates {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("taxRates[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.Prices, 1, -1, "prices"); err != nil {
		return err
	}

	for i, v := range s.Prices {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("prices[%d]", i), err)
		}
	}

	return nil
}

// TariffFixedPriceType (2.100)
//
// Tariff with optional conditions for a fixed price.
type TariffFixedPriceType struct {
	// Required. Fixed price for this element e.g. a start fee.
	PriceFixed float64 `json:"priceFixed"`
	// Optional. Conditions when this tariff element price is applicable. When absent always applicable,
	Conditions *TariffConditionsFixedType `json:"conditions,omitempty"`
}

func (s *TariffFixedPriceType) UnmarshalJSON(data []byte) error {
	type Alias TariffFixedPriceType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffFixedPriceType(a)
	return s.Validate()
}

func (s TariffFixedPriceType) Validate() error {
	if s.Conditions != nil {
		if err := s.Conditions.Validate(); err != nil {
			return ocpp.WrapField("conditions", err)
		}
	}

	return nil
}

// TariffFixedType (2.101)
type TariffFixedType struct {
	// Required.
	Prices []TariffFixedPriceType `json:"prices"`
	// Optional. Applicable tax percentages for this tariff dimension. If omitted, no tax is applicable. Not
	// providing a tax is different from 0% tax, which would be a value of 0.0 here.
	TaxRates []TaxRateType `json:"taxRates,omitempty"`
}

func (s *TariffFixedType) UnmarshalJSON(data []byte) error {
	type Alias TariffFixedType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffFixedType(a)
	return s.Validate()
}

func (s TariffFixedType) Validate() error {
	if err := validateSliceLen(s.Prices, 1, -1, "prices"); err != nil {
		return err
	}

	for i, v := range s.Prices {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("prices[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.TaxRates, 0, 5, "taxRates"); err != nil {
		return err
	}

	for i, v := range s.TaxRates {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("taxRates[%d]", i), err)
		}
	}

	return nil
}

// TariffTimePriceType (2.102)
//
// Tariff with optional conditions for a time duration price.
type TariffTimePriceType struct {
	// Required. Price per minute (excl. tax) for this element.
	PriceMinute float64 `json:"priceMinute"`
	// Optional. Conditions when this tariff element price is applicable. When absent always applicable,
	Conditions *TariffConditionsType `json:"conditions,omitempty"`
}

func (s *TariffTimePriceType) UnmarshalJSON(data []byte) error {
	type Alias TariffTimePriceType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffTimePriceType(a)
	return s.Validate()
}

func (s TariffTimePriceType) Validate() error {
	if s.Conditions != nil {
		if err := s.Conditions.Validate(); err != nil {
			return ocpp.WrapField("conditions", err)
		}
	}

	return nil
}

// TariffTimeType (2.103)
//
// Price elements and tax for time
type TariffTimeType struct {
	// Required. Element tariff price and conditions
	Prices []TariffTimePriceType `json:"prices"`
	// Optional. Applicable tax percentages for this tariff dimension. If omitted, no tax is applicable. Not
	// providing a tax is different from 0% tax, which would be a value of 0.0 here.
	TaxRates []TaxRateType `json:"taxRates,omitempty"`
}

func (s *TariffTimeType) UnmarshalJSON(data []byte) error {
	type Alias TariffTimeType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffTimeType(a)
	return s.Validate()
}

func (s TariffTimeType) Validate() error {
	if err := validateSliceLen(s.Prices, 1, -1, "prices"); err != nil {
		return err
	}

	for i, v := range s.Prices {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("prices[%d]", i), err)
		}
	}

	if err := validateSliceLen(s.TaxRates, 0, 5, "taxRates"); err != nil {
		return err
	}

	for i, v := range s.TaxRates {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("taxRates[%d]", i), err)
		}
	}

	return nil
}

// TariffType (2.104)
//
// A tariff is described by fields with prices for: energy, charging time, idle time, fixed fee, reservation
// time, reservation fixed fee. Each of these fields may have (optional) conditions that specify when a price is
// applicable. The description contains a human-readable explanation of the tariff to be shown to the user. The
// other fields are parameters that define the tariff. These are used by the charging station to calculate the
// price.
type TariffType struct {
	// Required. Unique id of tariff
	TariffID string `json:"tariffId"`
	// Required. Currency code according to ISO 4217
	Currency string `json:"currency"`
	// Optional. Time when this tariff becomes active. When absent, it is immediately active.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Optional. List of multi-language tariff information texts to be shown to the user.
	Description []MessageContentType `json:"description,omitempty"`
	// Optional. Energy tariff
	Energy *TariffEnergyType `json:"energy,omitempty"`
	// Optional. Charging time tariff
	ChargingTime *TariffTimeType `json:"chargingTime,omitempty"`
	// Optional. Idle time tariff
	IdleTime *TariffTimeType `json:"idleTime,omitempty"`
	// Optional. Fixed fee tariff
	FixedFee *TariffFixedType `json:"fixedFee,omitempty"`
	// Optional. The minimal cost for a transaction with this tariff including and excluding taxes. Minimum can be
	// including tax or excluding tax, or both.
	MinCost *PriceType `json:"minCost,omitempty"`
	// Optional. The maximum cost for a transaction with this tariff. Maximum can be including tax or excluding tax,
	// or both.
	MaxCost *PriceType `json:"maxCost,omitempty"`
	// Optional. Reservation time tariff
	ReservationTime *TariffTimeType `json:"reservationTime,omitempty"`
	// Optional. Fixed fee for a reservation
	ReservationFixed *TariffFixedType `json:"reservationFixed,omitempty"`
}

func (s *TariffType) UnmarshalJSON(data []byte) error {
	type Alias TariffType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TariffType(a)
	return s.Validate()
}

func (s TariffType) Validate() error {
	if s.TariffID == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "tariffId", "required field is missing")
	}

	if err := validateStringLen(s.TariffID, 60, "tariffId"); err != nil {
		return err
	}

	if s.Currency == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currency", "required field is missing")
	}

	if err := validateStringLen(s.Currency, 3, "currency"); err != nil {
		return err
	}

	if err := validateSliceLen(s.Description, 0, 10, "description"); err != nil {
		return err
	}

	for i, v := range s.Description {
		if err := v.Validate(); err != nil {
			return ocpp.WrapField(fmt.Sprintf("description[%d]", i), err)
		}
	}

	if s.Energy != nil {
		if err := s.Energy.Validate(); err != nil {
			return ocpp.WrapField("energy", err)
		}
	}

	if s.ChargingTime != nil {
		if err := s.ChargingTime.Validate(); err != nil {
			return ocpp.WrapField("chargingTime", err)
		}
	}

	if s.IdleTime != nil {
		if err := s.IdleTime.Validate(); err != nil {
			return ocpp.WrapField("idleTime", err)
		}
	}

	if s.FixedFee != nil {
		if err := s.FixedFee.Validate(); err != nil {
			return ocpp.WrapField("fixedFee", err)
		}
	}

	if s.MinCost != nil {
		if err := s.MinCost.Validate(); err != nil {
			return ocpp.WrapField("minCost", err)
		}
	}

	if s.MaxCost != nil {
		if err := s.MaxCost.Validate(); err != nil {
			return ocpp.WrapField("maxCost", err)
		}
	}

	if s.ReservationTime != nil {
		if err := s.ReservationTime.Validate(); err != nil {
			return ocpp.WrapField("reservationTime", err)
		}
	}

	if s.ReservationFixed != nil {
		if err := s.ReservationFixed.Validate(); err != nil {
			return ocpp.WrapField("reservationFixed", err)
		}
	}

	return nil
}

// TaxRateType (2.105)
//
// Tax percentage
type TaxRateType struct {
	// Required. Type of this tax, e.g. "Federal ", "State", for information on receipt.
	Type string `json:"type"`
	// Required. Tax percentage
	Tax float64 `json:"tax"`
	// Optional. Stack level for this type of tax. Default value, when absent, is 0. stack = 0: tax on net price;
	// stack = 1: tax added on top of stack 0; stack = 2: tax added on top of stack 1, etc.
	Stack *int32 `json:"stack,omitempty"`
}

func (s *TaxRateType) UnmarshalJSON(data []byte) error {
	type Alias TaxRateType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TaxRateType(a)
	return s.Validate()
}

func (s TaxRateType) Validate() error {
	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateStringLen(s.Type, 20, "type"); err != nil {
		return err
	}

	if s.Stack != nil {
		if err := validateNonNegative(*s.Stack, "stack"); err != nil {
			return err
		}
	}

	return nil
}

// TaxRuleType (2.106)
//
// Part of ISO 15118-20 price schedule.
type TaxRuleType struct {
	// Required. Id for the tax rule.
	TaxRuleID int32 `json:"taxRuleID"`
	// Optional. Human readable string to identify the tax rule.
	TaxRuleName *string `json:"taxRuleName,omitempty"`
	// Optional. Indicates whether the tax is included in any price or not.
	TaxIncludedInPrice *bool `json:"taxIncludedInPrice,omitempty"`
	// Required. Indicates whether this tax applies to Energy Fees.
	AppliesToEnergyFee bool `json:"appliesToEnergyFee"`
	// Required. Indicates whether this tax applies to Parking Fees.
	AppliesToParkingFee bool `json:"appliesToParkingFee"`
	// Required. Indicates whether this tax applies to Overstay Fees.
	AppliesToOverstayFee bool `json:"appliesToOverstayFee"`
	// Required. Indicates whether this tax applies to Minimum/Maximum Cost.
	AppliesToMinimumMaximumCost bool `json:"appliesToMinimumMaximumCost"`
	// Required. Percentage of the total amount of applying fee (energy, parking, overstay, MinimumCost and/or
	// MaximumCost).
	TaxRate RationalNumberType `json:"taxRate"`
}

func (s *TaxRuleType) UnmarshalJSON(data []byte) error {
	type Alias TaxRuleType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TaxRuleType(a)
	return s.Validate()
}

func (s TaxRuleType) Validate() error {
	if err := validateNonNegative(s.TaxRuleID, "taxRuleID"); err != nil {
		return err
	}

	if s.TaxRuleName != nil {
		if err := validateStringLen(*s.TaxRuleName, 100, "taxRuleName"); err != nil {
			return err
		}
	}

	if err := s.TaxRate.Validate(); err != nil {
		return ocpp.WrapField("taxRate", err)
	}

	return nil
}

// TotalCostType (2.107)
//
// This contains the cost calculated during a transaction. It is used both for running cost and final cost of the
// transaction.
type TotalCostType struct {
	// Required. Currency of the costs in ISO 4217 Code.
	Currency string `json:"currency"`
	// Required. Type of cost: normal or the minimum or maximum cost.
	TypeOfCost TariffCostEnumType `json:"typeOfCost"`
	// Optional. Total sum of all flat fees in the specified currency, except for TariffFixedPriceTypes with
	// conditions.isReservation = true (counted in reservation).
	Fixed *PriceType `json:"fixed,omitempty"`
	// Optional. Total sum of all the cost of all the energy used, in the specified currency.
	Energy *PriceType `json:"energy,omitempty"`
	// Optional. Total sum of all the cost related to duration of charging during this transaction, in the specified
	// currency.
	ChargingTime *PriceType `json:"chargingTime,omitempty"`
	// Optional. Total sum of all the cost related to idle time of this transaction, including fixed price
	// components, in the specified currency.
	IdleTime *PriceType `json:"idleTime,omitempty"`
	// Optional. Sum of all time-based cost related to reservation, i.e. TariffType.reservationTime, in the specified
	// currency.
	ReservationTime *PriceType `json:"reservationTime,omitempty"`
	// Required. Total of associated cost elements for fixed, energy, chargingTime, idleTime and reservation.
	Total TotalPriceType `json:"total"`
	// Optional. Sum of fixed cost related to reservation, i.e. TariffType.reservationFixed, in the specified
	// currency.
	ReservationFixed *PriceType `json:"reservationFixed,omitempty"`
}

func (s *TotalCostType) UnmarshalJSON(data []byte) error {
	type Alias TotalCostType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TotalCostType(a)
	return s.Validate()
}

func (s TotalCostType) Validate() error {
	if s.Currency == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "currency", "required field is missing")
	}

	if err := validateStringLen(s.Currency, 3, "currency"); err != nil {
		return err
	}

	if s.TypeOfCost == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "typeOfCost", "required field is missing")
	}

	if s.Fixed != nil {
		if err := s.Fixed.Validate(); err != nil {
			return ocpp.WrapField("fixed", err)
		}
	}

	if s.Energy != nil {
		if err := s.Energy.Validate(); err != nil {
			return ocpp.WrapField("energy", err)
		}
	}

	if s.ChargingTime != nil {
		if err := s.ChargingTime.Validate(); err != nil {
			return ocpp.WrapField("chargingTime", err)
		}
	}

	if s.IdleTime != nil {
		if err := s.IdleTime.Validate(); err != nil {
			return ocpp.WrapField("idleTime", err)
		}
	}

	if s.ReservationTime != nil {
		if err := s.ReservationTime.Validate(); err != nil {
			return ocpp.WrapField("reservationTime", err)
		}
	}

	if err := s.Total.Validate(); err != nil {
		return ocpp.WrapField("total", err)
	}

	if s.ReservationFixed != nil {
		if err := s.ReservationFixed.Validate(); err != nil {
			return ocpp.WrapField("reservationFixed", err)
		}
	}

	return nil
}

// TotalPriceType (2.108)
//
// Total cost with and without tax. Contains the total of energy, charging time, idle time, fixed and reservation
// costs including and/or excluding tax.
type TotalPriceType struct {
	// Optional. Price/cost excluding tax. Can be absent if inclTax is present.
	ExclTax *float64 `json:"exclTax,omitempty"`
	// Optional. Price/cost including tax. Can be absent if exclTax is present.
	InclTax *float64 `json:"inclTax,omitempty"`
}

func (s *TotalPriceType) UnmarshalJSON(data []byte) error {
	type Alias TotalPriceType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TotalPriceType(a)
	return s.Validate()
}

func (s TotalPriceType) Validate() error {
	_ = s
	return nil
}

// TotalUsageType (2.109)
//
// This contains the calculated usage of energy, charging time and idle time during a transaction.
type TotalUsageType struct {
	// Required.
	Energy float64 `json:"energy"`
	// Required. Total duration of the charging session (including the duration of charging and not charging), in
	// seconds.
	ChargingTime int32 `json:"chargingTime"`
	// Required. Total duration of the charging session where the EV was not charging (no energy was transferred
	// between EVSE and EV), in seconds.
	IdleTime int32 `json:"idleTime"`
	// Optional. Total time of reservation in seconds.
	ReservationTime *int32 `json:"reservationTime,omitempty"`
}

func (s *TotalUsageType) UnmarshalJSON(data []byte) error {
	type Alias TotalUsageType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TotalUsageType(a)
	return s.Validate()
}

func (s TotalUsageType) Validate() error {
	_ = s
	return nil
}

// TransactionLimitType (2.110)
//
// Cost, energy, time or SoC limit for a transaction.
type TransactionLimitType struct {
	// Optional. Maximum allowed cost of transaction in currency of tariff.
	MaxCost *float64 `json:"maxCost,omitempty"`
	// Optional. Maximum allowed energy in Wh to charge in transaction.
	MaxEnergy *float64 `json:"maxEnergy,omitempty"`
	// Optional. Maximum duration of transaction in seconds from start to end.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// Optional. Maximum State of Charge of EV in percentage.
	MaxSOC *int32 `json:"maxSoC,omitempty"`
}

func (s *TransactionLimitType) UnmarshalJSON(data []byte) error {
	type Alias TransactionLimitType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = TransactionLimitType(a)
	return s.Validate()
}

func (s TransactionLimitType) Validate() error {
	if s.MaxSOC != nil {
		if *s.MaxSOC < 0 || *s.MaxSOC > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "maxSoC", "must be between 0 and 100")
		}
	}

	return nil
}

// TransactionType (2.111)
type TransactionType struct {
	// Required. This contains the Id of the transaction.
	TransactionID IdentifierString36Type `json:"transactionId"`
	// Optional. Current charging state, is required when state has changed. Omitted when there is no communication
	// between EVSE and EV, because no cable is plugged in.
	ChargingState *ChargingStateEnumType `json:"chargingState,omitempty"`
	// Optional. Contains the total time that energy flowed from EVSE to EV during the transaction (in seconds). Note
	// that timeSpentCharging is smaller or equal to the duration of the transaction.
	TimeSpentCharging *int32 `json:"timeSpentCharging,omitempty"`
	// Optional. The stoppedReason is the reason/event that initiated the process of stopping the transaction. It
	// will normally be the user stopping authorization via card (Local or MasterPass) or app (Remote), but it can
	// also be CSMS revoking authorization (DeAuthorized), or disconnecting the EV when TxStopPoint = EVConnected
	// (EVDisconnected). Most other reasons are related to technical faults or energy limitations. MAY only be
	// omitted when stoppedReason is "Local"
	StoppedReason *ReasonEnumType `json:"stoppedReason,omitempty"`
	// Optional. The ID given to remote start request (RequestStartTransactionRequest. This enables to CSMS to match
	// the started transaction to the given start request.
	RemoteStartID *int32 `json:"remoteStartId,omitempty"`
	// Optional. (2.1) The operationMode that is currently in effect for the transaction.
	OperationMode *OperationModeEnumType `json:"operationMode,omitempty"`
	// Optional. (2.1) Id of tariff in use for transaction
	TariffID *string `json:"tariffId,omitempty"`
	// Optional. (2.1) Maximum cost/energy/time allowed for this transaction.
	TransactionLimit *TransactionLimitType `json:"transactionLimit,omitempty"`
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

	if s.TariffID != nil {
		if err := validateStringLen(*s.TariffID, 60, "tariffId"); err != nil {
			return err
		}
	}

	if s.TransactionLimit != nil {
		if err := s.TransactionLimit.Validate(); err != nil {
			return ocpp.WrapField("transactionLimit", err)
		}
	}

	return nil
}

// UnitOfMeasureType (2.112)
//
// Represents a UnitOfMeasure with a multiplier
type UnitOfMeasureType struct {
	// Optional. Unit of the value. Default = "Wh" if the (default) measurand is an "Energy" type. This field SHALL
	// use a value from the list Standardized Units of Measurements in Part 2 Appendices. If an applicable unit is
	// available in that list, otherwise a "custom" unit might be used.
	Unit *string `json:"unit,omitempty"`
	// Optional. Multiplier, this value represents the exponent to base 10. I.e. multiplier 3 means 10 raised to the
	// 3rd power. Default is 0. The multiplier only multiplies the value of the measurand. It does not specify a
	// conversion between units, for example, kW and W.
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

// V2XChargingParametersType (2.113)
//
// Charging parameters for ISO 15118-20, also supporting V2X charging/discharging.+ All values are greater or
// equal to zero, with the exception of EVMinEnergyRequest, EVMaxEnergyRequest, EVTargetEnergyRequest,
// EVMinV2XEnergyRequest and EVMaxV2XEnergyRequest.
type V2XChargingParametersType struct {
	// Optional. Minimum charge power in W, defined by max(EV, EVSE). This field represents the sum of all phases,
	// unless values are provided for L2 and L3, in which case this field represents phase L1. Relates to: ISO
	// 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType: EVMinimumChargePower
	MinChargePower *float64 `json:"minChargePower,omitempty"`
	// Optional. Minimum charge power on phase L2 in W, defined by max(EV, EVSE). Relates to: ISO 15118-20:
	// BPT_AC/DC_CPDReqEnergyTransferModeType: EVMinimumChargePower_L2
	MinChargePowerL2 *float64 `json:"minChargePower_L2,omitempty"`
	// Optional. Minimum charge power on phase L3 in W, defined by max(EV, EVSE). Relates to: ISO 15118-20:
	// BPT_AC/DC_CPDReqEnergyTransferModeType: EVMinimumChargePower_L3
	MinChargePowerL3 *float64 `json:"minChargePower_L3,omitempty"`
	// Optional. Maximum charge (absorbed) power in W, defined by min(EV, EVSE) at unity power factor. This field
	// represents the sum of all phases, unless values are provided for L2 and L3, in which case this field
	// represents phase L1. It corresponds to the ChaWMax attribute in the IEC 61850. It is usually equivalent to the
	// rated apparent power of the EV when discharging (ChaVAMax) in IEC 61850. Relates to: ISO 15118-20:
	// BPT_AC/DC_CPDReqEnergyTransferModeType: EVMaximumChargePower
	MaxChargePower *float64 `json:"maxChargePower,omitempty"`
	// Optional. Maximum charge power on phase L2 in W, defined by min(EV, EVSE) Relates to: ISO 15118-20:
	// BPT_AC/DC_CPDReqEnergyTransferModeType: EVMaximumChargePower_L2
	MaxChargePowerL2 *float64 `json:"maxChargePower_L2,omitempty"`
	// Optional. Maximum charge power on phase L3 in W, defined by min(EV, EVSE) Relates to: ISO 15118-20:
	// BPT_AC/DC_CPDReqEnergyTransferModeType: EVMaximumChargePower_L3
	MaxChargePowerL3 *float64 `json:"maxChargePower_L3,omitempty"`
	// Optional. Minimum discharge (injected) power in W, defined by max(EV, EVSE) at unity power factor. Value >= 0.
	// This field represents the sum of all phases, unless values are provided for L2 and L3, in which case this
	// field represents phase L1. It corresponds to the WMax attribute in the IEC 61850. It is usually equivalent to
	// the rated apparent power of the EV when discharging (VAMax attribute in the IEC 61850). Relates to: ISO
	// 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType: EVMinimumDischargePower
	MinDischargePower *float64 `json:"minDischargePower,omitempty"`
	// Optional. Minimum discharge power on phase L2 in W, defined by max(EV, EVSE). Value >= 0. Relates to: ISO
	// 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType: EVMinimumDischargePower_L2
	MinDischargePowerL2 *float64 `json:"minDischargePower_L2,omitempty"`
	// Optional. Minimum discharge power on phase L3 in W, defined by max(EV, EVSE). Value >= 0. Relates to: ISO
	// 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType: EVMinimumDischargePower_L3
	MinDischargePowerL3 *float64 `json:"minDischargePower_L3,omitempty"`
	// Optional. Maximum discharge (injected) power in W, defined by min(EV, EVSE) at unity power factor. Value >= 0.
	// This field represents the sum of all phases, unless values are provided for L2 and L3, in which case this
	// field represents phase L1. Relates to: ISO 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType:
	// EVMaximumDischargePower
	MaxDischargePower *float64 `json:"maxDischargePower,omitempty"`
	// Optional. Maximum discharge power on phase L2 in W, defined by min(EV, EVSE). Value >= 0. Relates to: ISO
	// 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType: EVMaximumDischargePowe_L2
	MaxDischargePowerL2 *float64 `json:"maxDischargePower_L2,omitempty"`
	// Optional. Maximum discharge power on phase L3 in W, defined by min(EV, EVSE). Value >= 0. Relates to: ISO
	// 15118-20: BPT_AC/DC_CPDReqEnergyTransferModeType: EVMaximumDischargePower_L3
	MaxDischargePowerL3 *float64 `json:"maxDischargePower_L3,omitempty"`
	// Optional. Minimum charge current in A, defined by max(EV, EVSE) Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: EVMinimumChargeCurrent
	MinChargeCurrent *float64 `json:"minChargeCurrent,omitempty"`
	// Optional. Maximum charge current in A, defined by min(EV, EVSE) Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: EVMaximumChargeCurrent
	MaxChargeCurrent *float64 `json:"maxChargeCurrent,omitempty"`
	// Optional. Minimum discharge current in A, defined by max(EV, EVSE). Value >= 0. Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: EVMinimumDischargeCurrent
	MinDischargeCurrent *float64 `json:"minDischargeCurrent,omitempty"`
	// Optional. Maximum discharge current in A, defined by min(EV, EVSE). Value >= 0. Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: EVMaximumDischargeCurrent
	MaxDischargeCurrent *float64 `json:"maxDischargeCurrent,omitempty"`
	// Optional. Minimum voltage in V, defined by max(EV, EVSE) Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: EVMinimumVoltage
	MinVoltage *float64 `json:"minVoltage,omitempty"`
	// Optional. Maximum voltage in V, defined by min(EV, EVSE) Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: EVMaximumVoltage
	MaxVoltage *float64 `json:"maxVoltage,omitempty"`
	// Optional. Energy to requested state of charge in Wh Relates to: ISO 15118-20:
	// Dynamic/Scheduled_SEReqControlModeType: EVTargetEnergyRequest
	EVTargetEnergyRequest *float64 `json:"evTargetEnergyRequest,omitempty"`
	// Optional. Energy to minimum allowed state of charge in Wh Relates to: ISO 15118-20:
	// Dynamic/Scheduled_SEReqControlModeType: EVMinimumEnergyRequest
	EVMinEnergyRequest *float64 `json:"evMinEnergyRequest,omitempty"`
	// Optional. Energy to maximum state of charge in Wh Relates to: ISO 15118-20:
	// Dynamic/Scheduled_SEReqControlModeType: EVMaximumEnergyRequest
	EVMaxEnergyRequest *float64 `json:"evMaxEnergyRequest,omitempty"`
	// Optional. Energy (in Wh) to minimum state of charge for cycling (V2X) activity. Positive value means that
	// current state of charge is below V2X range. Relates to: ISO 15118-20: Dynamic_SEReqControlModeType:
	// EVMinimumV2XEnergyRequest
	EVMinV2XEnergyRequest *float64 `json:"evMinV2XEnergyRequest,omitempty"`
	// Optional. Energy (in Wh) to maximum state of charge for cycling (V2X) activity. Negative value indicates that
	// current state of charge is above V2X range. Relates to: ISO 15118-20: Dynamic_SEReqControlModeType:
	// EVMaximumV2XEnergyRequest
	EVMaxV2XEnergyRequest *float64 `json:"evMaxV2XEnergyRequest,omitempty"`
	// Optional. Target state of charge at departure as percentage. Relates to: ISO 15118-20:
	// BPT_DC_CPDReqEnergyTransferModeType: TargetSOC
	TargetSOC *int32 `json:"targetSoC,omitempty"`
}

func (s *V2XChargingParametersType) UnmarshalJSON(data []byte) error {
	type Alias V2XChargingParametersType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = V2XChargingParametersType(a)
	return s.Validate()
}

func (s V2XChargingParametersType) Validate() error {
	if s.TargetSOC != nil {
		if *s.TargetSOC < 0 || *s.TargetSOC > 100 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "targetSoC", "must be between 0 and 100")
		}
	}

	return nil
}

// V2XFreqWattPointType (2.114)
//
// (2.1) A point of a frequency-watt curve.
type V2XFreqWattPointType struct {
	// Required. Net frequency in Hz.
	Frequency float64 `json:"frequency"`
	// Required. Power in W to charge (positive) or discharge (negative) at specified frequency.
	Power float64 `json:"power"`
}

func (s *V2XFreqWattPointType) UnmarshalJSON(data []byte) error {
	type Alias V2XFreqWattPointType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = V2XFreqWattPointType(a)
	return s.Validate()
}

func (s V2XFreqWattPointType) Validate() error {
	_ = s
	return nil
}

// V2XSignalWattPointType (2.115)
//
// (2.1) A point of a signal-watt curve.
type V2XSignalWattPointType struct {
	// Required. Signal value from an AFRRSignalRequest.
	Signal int32 `json:"signal"`
	// Required. Power in W to charge (positive) or discharge (negative) at specified frequency.
	Power float64 `json:"power"`
}

func (s *V2XSignalWattPointType) UnmarshalJSON(data []byte) error {
	type Alias V2XSignalWattPointType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = V2XSignalWattPointType(a)
	return s.Validate()
}

func (s V2XSignalWattPointType) Validate() error {
	_ = s
	return nil
}

// VariableAttributeType (2.116)
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
	// Optional. If true, value will be persistent across system reboots or power down. Default when omitted is
	// false.
	Persistent *bool `json:"persistent,omitempty"`
	// Optional. If true, value that will never be changed by the Charging Station at runtime. Default when omitted
	// is false.
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

// VariableCharacteristicsType (2.117)
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
	// Optional. (2.1) Maximum number of elements from valuesList that are supported as attributeValue.
	MaxElements *int32 `json:"maxElements,omitempty"`
	// Optional. Mandatory when dataType = OptionList, MemberList or SequenceList. In that case valuesList specifies
	// the allowed values for the type. The length of this field can be limited by
	// DeviceDataCtrlr.ConfigurationValueSize. * OptionList: The (Actual) Variable value must be a single value from
	// the reported (CSV) enumeration list. * MemberList: The (Actual) Variable value may be an (unordered) (sub-)set
	// of the reported (CSV) valid values list. * SequenceList: The (Actual) Variable value may be an ordered
	// (priority, etc) (sub-)set of the reported (CSV) valid values. This is a comma separated list. The
	// Configuration Variable ConfigurationValueSize can be used to limit SetVariableData.attributeValue and
	// VariableCharacteristics.valuesList. The max size of these values will always remain equal.
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

	if s.MaxElements != nil {
		if *s.MaxElements < 1 {
			return ocpp.NewError(ocpp.ErrPropertyConstraintViolation, "maxElements", "must be >= 1")
		}
	}

	if s.ValuesList != nil {
		if err := validateStringLen(*s.ValuesList, 1000, "valuesList"); err != nil {
			return err
		}
	}

	return nil
}

// VariableMonitoringType (2.118)
//
// A monitoring setting for a variable.
type VariableMonitoringType struct {
	// Required. Identifies the monitor.
	ID int32 `json:"id"`
	// Required. Monitor only active when a transaction is ongoing on a component relevant to this transaction.
	Transaction bool `json:"transaction"`
	// Required. Value for threshold or delta monitoring. For Periodic or PeriodicClockAligned this is the interval
	// in seconds.
	Value float64 `json:"value"`
	// Required. The type of this monitor, e.g. a threshold, delta or periodic monitor.
	Type MonitorEnumType `json:"type"`
	// Required. The severity that will be assigned to an event that is triggered by this monitor. The severity range
	// is 0-9, with 0 as the highest and 9 as the lowest severity level. The severity levels have the following
	// meaning: 0-Danger Indicates lives are potentially in danger. Urgent attention is needed and action should be
	// taken immediately. 1-Hardware Failure Indicates that the Charging Station is unable to continue regular
	// operations due to Hardware issues. Action is required. 2-System Failure Indicates that the Charging Station is
	// unable to continue regular operations due to software or minor hardware issues. Action is required. 3-Critical
	// Indicates a critical error. Action is required. 4-Error Indicates a non-urgent error. Action is required.
	// 5-Alert Indicates an alert event. Default severity for any type of monitoring event. 6-Warning Indicates a
	// warning event. Action may be required. 7-Notice Indicates an unusual event. No immediate action is required.
	// 8-Informational Indicates a regular operational event. May be used for reporting, measuring throughput, etc.
	// No action is required. 9-Debug Indicates information useful to developers for debugging, not useful during
	// operations.
	Severity int32 `json:"severity"`
	// Required. (2.1) Type of monitor.
	EventNotificationType EventNotificationEnumType `json:"eventNotificationType"`
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
	if err := validateNonNegative(s.ID, "id"); err != nil {
		return err
	}

	if s.Type == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "type", "required field is missing")
	}

	if err := validateNonNegative(s.Severity, "severity"); err != nil {
		return err
	}

	if s.EventNotificationType == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "eventNotificationType", "required field is missing")
	}

	return nil
}

// VariableType (2.119)
//
// Reference key to a component-variable.
type VariableType struct {
	// Required. Name of the variable. Name should be taken from the list of standardized variable names whenever
	// possible. Case Insensitive. strongly advised to use Camel Case.
	Name IdentifierString50Type `json:"name"`
	// Optional. Name of instance in case the variable exists as multiple instances. Case Insensitive. strongly
	// advised to use Camel Case.
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

// VoltageParamsType (2.120)
type VoltageParamsType struct {
	// Optional. EN 50549-1 chapter 4.9.3.4 Voltage threshold for the 10 min time window mean value monitoring. The
	// 10 min mean is recalculated up to every 3 s. If the present voltage is above this threshold for more than the
	// time defined by hv10MinMeanValue, the EV must trip. This value is mandatory if hv10MinMeanTripDelay is set.
	Hv10MinMeanValue *float64 `json:"hv10MinMeanValue,omitempty"`
	// Optional. Time for which the voltage is allowed to stay above the 10 min mean value. After this time, the EV
	// must trip. This value is mandatory if OverVoltageMeanValue10min is set.
	Hv10MinMeanTripDelay *float64 `json:"hv10MinMeanTripDelay,omitempty"`
	// Optional. Parameter is only sent, if the EV has to feed-in power or reactive power during fault-ride through
	// (FRT) as defined by HVMomCess curve and LVMomCess curve.
	PowerDuringCessation *PowerDuringCessationEnumType `json:"powerDuringCessation,omitempty"`
}

func (s *VoltageParamsType) UnmarshalJSON(data []byte) error {
	type Alias VoltageParamsType
	var a Alias

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	*s = VoltageParamsType(a)
	return s.Validate()
}

func (s VoltageParamsType) Validate() error {
	_ = s
	return nil
}

// VPNType (2.121)
//
// VPN Configuration settings
type VPNType struct {
	// Required. VPN Server Address
	Server string `json:"server"`
	// Required. VPN User
	User string `json:"user"`
	// Optional. VPN group.
	Group *string `json:"group,omitempty"`
	// Required. (2.1) VPN Password.
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

	if err := validateStringLen(s.Server, 2000, "server"); err != nil {
		return err
	}

	if s.User == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "user", "required field is missing")
	}

	if err := validateStringLen(s.User, 50, "user"); err != nil {
		return err
	}

	if s.Group != nil {
		if err := validateStringLen(*s.Group, 50, "group"); err != nil {
			return err
		}
	}

	if s.Password == "" {
		return ocpp.NewError(ocpp.ErrOccurenceConstraintViolation, "password", "required field is missing")
	}

	if err := validateStringLen(s.Password, 64, "password"); err != nil {
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
