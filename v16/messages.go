package v16

import (
	"encoding/json"
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
	type raw AuthorizeReq
	var r raw

	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	*s = AuthorizeReq(r)
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
	type raw AuthorizeConf
	var r raw

	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	*s = AuthorizeConf(r)
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
	type raw BootNotificationReq
	var r raw

	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	*s = BootNotificationReq(r)
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
	type raw BootNotificationConf
	var r raw

	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	*s = BootNotificationConf(r)
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
