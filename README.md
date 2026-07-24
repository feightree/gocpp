# gocpp

Go types and message definitions for the [Open Charge Point Protocol (OCPP) 1.6](https://www.openchargealliance.org/protocols/ocpp-16/), as defined by the Open Charge Alliance.

## What this is

`gocpp` provides spec-compliant Go types for OCPP 1.6 messages and data types, suitable for use as a foundation when building charge point or central system implementations. All 56 OCPP 1.6 message types (Authorize through UpdateFirmware) and their supporting data types are implemented, with validation on unmarshal and full test coverage.

## What this is not

This library does not provide a WebSocket transport, message routing, or any networking. It is purely a typed representation of the protocol.

## Packages

- `v16` — OCPP 1.6 message and data types (enums, request/response structs).
- `ocpp` — shared, version-agnostic error types used across `v16` (and future protocol versions).

## Installation

```sh
go get github.com/feightree/gocpp
```

Requires Go 1.26+.

## Usage

```go
import "github.com/feightree/gocpp/v16"

// Unmarshal an incoming message
var req v16.BootNotificationReq
if err := json.Unmarshal(data, &req); err != nil {
    // invalid payload or out-of-spec field values
}

// Construct an outgoing message
conf := v16.BootNotificationConf{
    CurrentTime: time.Now().UTC(),
    Interval:    60,
    Status:      v16.RegistrationStatusAccepted,
}

// Marshal does not validate — call Validate explicitly for messages you build by hand
if err := conf.Validate(); err != nil {
    // constructed message violates the spec
}

data, err := json.Marshal(conf)
```

Every type validates itself on unmarshal — enum fields must match a known
value, required fields must be present, and constrained fields (string
lengths, numeric ranges, cross-field rules like `TxProfile`-only charging
profiles) are checked. Failures return an `*ocpp.Error` wrapping an
`ocpp.ErrorCode` (e.g. `ocpp.ErrPropertyConstraintViolation`,
`ocpp.ErrOccurenceConstraintViolation`), matching the error codes OCPP itself
uses in `CALLERROR` messages. Use `errors.Is` to detect a specific code.

```go
import (
    "errors"

    "github.com/feightree/gocpp/ocpp"
    "github.com/feightree/gocpp/v16"
)

var status v16.AuthorizationStatus
err := json.Unmarshal([]byte(`"Bogus"`), &status)
if errors.Is(err, ocpp.ErrPropertyConstraintViolation) {
    // received an unknown enum value
}
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE.md).
