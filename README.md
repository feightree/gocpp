# gocpp

Go types and message definitions for the [Open Charge Point Protocol (OCPP)](https://www.openchargealliance.org/), versions [1.6](https://openchargealliance.org/protocols/open-charge-point-protocol/#ExploreOCPP1.6), [2.0.1](https://openchargealliance.org/protocols/open-charge-point-protocol/#ExploreOCPP2.0.1) and [2.1](https://openchargealliance.org/protocols/open-charge-point-protocol/#ExploreOCPP2.1), as defined by the Open Charge Alliance.

## What this is

`gocpp` provides spec-compliant Go types for OCPP messages and data types, suitable for use as a foundation when building charge point or central system implementations, for any of the three protocol versions:

- **OCPP 1.6** (`v16`): all 56 message types (Authorize through UpdateFirmware) and their supporting data types.
- **OCPP 2.0.1** (`v201`): all 64 actions (128 request/response types), 54 datatypes, and 88 enumerations.
- **OCPP 2.1** (`v21`): all 91 actions (181 request/response types — one, `NotifyPeriodicEventStream`, is a one-way message with no response), 121 datatypes, and 110 enumerations.

All three are implemented with validation on unmarshal and full test coverage.

## What this is not

This library does not provide a WebSocket transport, message routing, or any networking. It is purely a typed representation of the protocol.

## Packages

- `v16` — OCPP 1.6 message and data types (enums, request/response structs, suffixed `Req`/`Conf` per OCPP 1.6's own terminology).
- `v201` — OCPP 2.0.1 message and data types (enums, request/response structs, suffixed `Request`/`Response` per OCPP 2.0.1's own terminology).
- `v21` — OCPP 2.1 message and data types (enums, request/response structs, suffixed `Request`/`Response` per OCPP 2.1's own terminology).
- `ocpp` — shared, version-agnostic error types used across `v16`, `v201`, and `v21`.

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

`v201` and `v21` follow the same design (`UnmarshalJSON` validates,
`Validate` is exposed for hand-built messages you marshal yourself), just
with `Request`/`Response` naming instead of `Req`/`Conf`, e.g.
`v21.BootNotificationRequest` / `v21.BootNotificationResponse` — each
package matches the terminology its own OCPP version uses.

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
