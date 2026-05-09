package ocpp16

import "time"

// 6.1. Authorize.req
//
// This contains the field definition of the Authorize.req PDU sent by the Charge Point to the Central System. See
// also Authorize
type AuthorizeReq struct {
	// Required. This contains the identifier that needs to be authorized.
	IDTag IDToken `json:"idTag"`
}

// 6.2. Authorize.conf
//
// This contains the field definition of the Authorize.conf PDU sent by the Central System to the Charge Point in
// response to a Authorize.req PDU. See also Authorize
type AuthorizeConf struct {
	// Required. This contains information about authorization status, expiry parent id.
	IDTagInfo IDTagInfo `json:"idTagInfo"`
}

// 6.3. BootNotification.req
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

// 6.4. BootNotification.conf
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

// 6.5. CancelReservation.req
//
// This contains the field definition of the CancelReservation.req PDU sent by the Central System to the Charge
// Point. See also Cancel Reservation
type CancelReservationReq struct {
	// Required. Id of the reservation to cancel.
	ReservationID int32 `json:"reservationId"`
}

// 6.6. CancelReservation.conf
//
// This contains the field definition of the CancelReservation.conf PDU sent by the Charge Point to the Central
// System in response to a CancelReservation.req PDU. See also Cancel Reservation
type CancelReservationConf struct {
	// Required. This indicates the success or failure of the cancelling of
	// a reservation by Central System.
	Status CancelReservationStatus `json:"status"`
}

// 6.7. ChangeAvailability.req
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

// 6.8. ChangeAvailability.conf
//
// This contains the field definition of the ChangeAvailability.conf PDU return by Charge Point to Central System.
// See also Change Availability
type ChangeAvailabilityConf struct {
	// Required. This indicates whether the Charge Point is able to perform the
	// availability change.
	Status AvailabilityStatus `json:"status"`
}

// 6.9. ChangeConfiguration.req
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

// 6.10. ChangeConfiguration.conf
//
// This contains the field definition of the ChangeConfiguration.conf PDU returned from Charge Point to Central
// System. See also Change Configuration
type ChangeConfigurationConf struct {
	// Required. Returns whether configuration change has been accepted.
	Status ConfigurationStatus `json:"status"`
}

// 6.11. ClearCache.req
//
// This contains the field definition of the ClearCache.req PDU sent by the Central System to the Charge Point. See
// also Clear Cache
//
// No fields are defined.
type ClearCacheReq struct{}

// 6.12. ClearCache.conf
//
// This contains the field definition of the ClearCache.conf PDU sent by the Charge Point to the Central System in
// response to a ClearCache.req PDU. See also Clear Cache
type ClearCacheConf struct {
	// Required. Accepted if the Charge Point has executed the request, otherwise rejected.
	Status ClearCacheStatus `json:"status"`
}

// 6.13. ClearChargingProfile.req
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

// 6.14. ClearChargingProfile.conf
//
// This contains the field definition of the ClearChargingProfile.conf PDU sent by the Charge Point to the Central
// System in response to a ClearChargingProfile.req PDU. See also Clear Charging Profile
type ClearChargingProfileConf struct {
	// Required. Indicates if the Charge Point was able to execute the request.
	Status ClearChargingProfileStatus `json:"status"`
}

// 6.15. DataTransfer.req
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

// 6.16. DataTransfer.conf
//
// This contains the field definition of the DataTransfer.conf PDU sent by the Charge Point to the Central System or
// vice versa in response to a DataTransfer.req PDU. See also Data Transfer
type DataTransferConf struct {
	// Required. This indicates the success or failure of the data transfer.
	Status DataTransferStatus `json:"status"`
	// Optional. Data in response to request.
	Data *string `json:"data,omitempty"`
}

// 6.17. DiagnosticsStatusNotification.req
//
// This contains the field definition of the DiagnosticsStatusNotification.req PDU sent by the Charge Point to the
// Central System. See also Diagnostics Status Notification
type DiagnosticsStatusNotificationReq struct {
	// Required. This contains the status of the diagnostics upload.
	Status DiagnosticsStatus `json:"status"`
}

// 6.18. DiagnosticsStatusNotification.conf
//
// This contains the field definition of the DiagnosticsStatusNotification.conf PDU sent by the Central System to the
// Charge Point in response to a DiagnosticsStatusNotification.req PDU. See also Diagnostics Status Notification
// No fields are defined.
type DiagnosticsStatusNotificationConf struct{}

// 6.19. FirmwareStatusNotification.req
//
// This contains the field definition of the FirmwareStatusNotification.req PDU sent by the Charge Point to the
// Central System. See also Firmware Status Notification
type FirmwareStatusNotificationReq struct {
	// Required. This contains the progress status of the firmware installation.
	Status FirmwareStatus `json:"status"`
}

// 6.20. FirmwareStatusNotification.conf
//
// This contains the field definition of the FirmwareStatusNotification.conf PDU sent by the Central System to the
// Charge Point in response to a FirmwareStatusNotification.req PDU. See also Firmware Status Notification
// No fields are defined.
type FirmwareStatusNotificationConf struct{}

// 6.21. GetCompositeSchedule.req
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

// 6.22. GetCompositeSchedule.conf
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

// 6.23. GetConfiguration.req
//
// This contains the field definition of the GetConfiguration.req PDU sent by the Central System to the Charge
// Point. See also Get Configuration
type GetConfigurationReq struct {
	// Optional. List of keys for which the configuration value is requested.
	Key []CiString50Type `json:"key,omitempty"`
}

// 6.24. GetConfiguration.conf
//
// This contains the field definition of the GetConfiguration.conf PDU sent by the Charge Point to the Central
// System in response to a GetConfiguration.req. See also Get Configuration
type GetConfigurationConf struct {
	// Optional. List of requested or known keys
	ConfigurationKey []KeyValue `json:"configurationKey,omitempty"`
	// Optional. Requested keys that are unknown
	UnknownKey []CiString50Type `json:"unknownKey,omitempty"`
}

// 6.25. GetDiagnostics.req
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

// 6.26. GetDiagnostics.conf
//
// This contains the field definition of the GetDiagnostics.conf PDU sent by the Charge Point to the Central System
// in response to a GetDiagnostics.req PDU. See also Get Diagnostics
type GetDiagnosticsConf struct {
	// Optional. This contains the name of the file with diagnostic information that will
	// be uploaded. This field is not present when no diagnostic information is
	// available.
	FileName *CiString255Type `json:"fileName,omitempty"`
}

// 6.27. GetLocalListVersion.req
//
// This contains the field definition of the GetLocalListVersion.req PDU sent by the Central System to the Charge
// Point. See also Get Local List Version
//
// No fields are defined.
type GetLocalListVersionReq struct{}

// 6.28. GetLocalListVersion.conf
//
// This contains the field definition of the GetLocalListVersion.conf PDU sent by the Charge Point to Central System
// in response to a GetLocalListVersion.req PDU. See also Get Local List Version
type GetLocalListVersionConf struct {
	// Required. This contains the current version number of the local authorization list in the Charge Point.
	ListVersion int32 `json:"listVersion"`
}

// 6.29. Heartbeat.req
//
// This contains the field definition of the Heartbeat.req PDU sent by the Charge Point to the Central System. See
// also Heartbeat
//
// No fields are defined.
type HeartbeatReq struct{}

// 6.30. Heartbeat.conf
//
// This contains the field definition of the Heartbeat.conf PDU sent by the Central System to the Charge Point in
// response to a Heartbeat.req PDU. See also Heartbeat
type HeartbeatConf struct {
	// Required. This contains the current time of the Central System.
	CurrentTime time.Time `json:"currentTime"`
}

// 6.31. MeterValues.req
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

// 6.32. MeterValues.conf
//
// This contains the field definition of the MeterValues.conf PDU sent by the Central System to the Charge Point in
// response to a MeterValues.req PDU. See also Meter Values
//
// No fields are defined.
type MeterValuesConf struct{}

// 6.33. RemoteStartTransaction.req
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

// 6.34. RemoteStartTransaction.conf
//
// This contains the field definitions of the RemoteStartTransaction.conf PDU sent from Charge Point to Central
// System. See also Remote Start Transaction
type RemoteStartTransactionConf struct {
	// Required. Status indicating whether Charge Point accepts the request to start a transaction.
	Status RemoteStartStopStatus `json:"status"`
}

// 6.35. RemoteStopTransaction.req
//
// This contains the field definitions of the RemoteStopTransaction.req PDU sent to Charge Point by Central
// System. See also Remote Stop Transaction
type RemoteStopTransactionReq struct {
	// Required. The identifier of the transaction which Charge Point is requested to stop.
	TransactionID int32 `json:"transactionId"`
}

// 6.36. RemoteStopTransaction.conf
//
// This contains the field definitions of the RemoteStopTransaction.conf PDU sent from Charge Point to Central
// System. See also Remote Stop Transaction
type RemoteStopTransactionConf struct {
	// Required. Status indicating whether Charge Point accepts the request to stop a transaction.
	Status RemoteStartStopStatus `json:"status"`
}

// 6.37. ReserveNow.req
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

// 6.38. ReserveNow.conf
//
// This contains the field definition of the ReserveNow.conf PDU sent by the Charge Point to the Central System in
// response to a ReserveNow.req PDU. See also Reserve Now
type ReserveNowConf struct {
	// Required. This indicates the success or failure of the reservation.
	Status ReservationStatus `json:"status"`
}

// 6.39. Reset.req
//
// This contains the field definition of the Reset.req PDU sent by the Central System to the Charge Point. See also
// Reset
type ResetReq struct {
	// Required. This contains the type of reset that the Charge Point should perform.
	Type ResetType `json:"type"`
}

// 6.40. Reset.conf
//
// This contains the field definition of the Reset.conf PDU sent by the Charge Point to the Central System in
// response to a Reset.req PDU. See also Reset
type ResetConf struct {
	// Required. This indicates whether the Charge Point is able to perform the reset.
	Status ResetStatus `json:"status"`
}

// 6.41. SendLocalList.req
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

// 6.42. SendLocalList.conf
//
// This contains the field definition of the SendLocalList.conf PDU sent by the Charge Point to the Central System in
// response to a SendLocalList.req PDU. See also Send Local List
type SendLocalListConf struct {
	// Required. This indicates whether the Charge Point has successfully received and
	// applied the update of the local authorization list.
	Status UpdateStatus `json:"status"`
}

// 6.43. SetChargingProfile.req
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

// 6.44. SetChargingProfile.conf
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

// 6.45. StartTransaction.req
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

// 6.46. StartTransaction.conf
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

// 6.47. StatusNotification.req
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

// 6.48. StatusNotification.conf
//
// This contains the field definition of the StatusNotification.conf PDU sent by the Central System to the Charge
// Point in response to an StatusNotification.req PDU. See also Status Notification
//
// No fields are defined.
type StatusNotificationConf struct{}

// 6.49. StopTransaction.req
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

// 6.50. StopTransaction.conf
//
// This contains the field definition of the StopTransaction.conf PDU sent by the Central System to the Charge Point
// in response to a StopTransaction.req PDU. See also Stop Transaction
type StopTransactionConf struct {
	// Optional. This contains information about authorization status, expiry and
	// parent id. It is optional, because a transaction may have been stopped without
	// an identifier.
	IDTagInfo *IDTagInfo `json:"idTagInfo,omitempty"`
}

// 6.51. TriggerMessage.req
//
// This contains the field definition of the TriggerMessage.req PDU sent by the Central System to the Charge Point.
// See also Trigger Message
type TriggerMessageReq struct {
	// Required.
	RequestedMessage MessageTrigger `json:"requestedMessage"`
	// Optional. Only filled in when request applies to a specific connector.
	ConnectorID *int32 `json:"connectorId,omitempty"`
}

// 6.52. TriggerMessage.conf
//
// This contains the field definition of the TriggerMessage.conf PDU sent by the Charge Point to the Central System
// in response to a TriggerMessage.req PDU. See also Trigger Message
type TriggerMessageConf struct {
	// Required. Indicates whether the Charge Point will send the requested
	// notification or not.
	Status TriggerMessageStatus `json:"status"`
}

// 6.53. UnlockConnector.req
//
// This contains the field definition of the UnlockConnector.req PDU sent by the Central System to the Charge
// Point. See also Unlock Connector
type UnlockConnectorReq struct {
	// Required. This contains the identifier of the connector to be unlocked.
	ConnectorID int32 `json:"connectorId"`
}

// 6.54. UnlockConnector.conf
//
// This contains the field definition of the UnlockConnector.conf PDU sent by the Charge Point to the Central
// System in response to an UnlockConnector.req PDU. See also Unlock Connector
type UnlockConnectorConf struct {
	// Required. This indicates whether the Charge Point has unlocked the connector.
	Status UnlockStatus `json:"status"`
}

// 6.55. UpdateFirmware.req
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

// 6.56. UpdateFirmware.conf
//
// This contains the field definition of the UpdateFirmware.conf PDU sent by the Charge Point to the Central
// System in response to a UpdateFirmware.req PDU. See also Update Firmware
//
// No fields are defined.
type UpdateFirmwareConf struct{}
