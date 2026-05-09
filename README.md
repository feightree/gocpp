# gocpp

Go types and message definitions for the [Open Charge Point Protocol (OCPP) 1.6](https://www.openchargealliance.org/protocols/ocpp-16/), as defined by the Open Charge Alliance.

## What this is

`gocpp` provides spec-compliant Go types for OCPP 1.6 messages and data types, suitable for use as a foundation when building charge point or central system implementations.

## What this is not

This library does not provide a WebSocket transport, message routing, or any networking. It is purely a typed representation of the protocol.

## Installation

```sh
go get github.com/feightree/gocpp
```

Requires Go 1.26+.

## Usage

```go
import "github.com/feightree/gocpp/ocpp16"

// Unmarshal an incoming message
var req ocpp16.BootNotificationReq
if err := json.Unmarshal(data, &req); err != nil {
    // invalid payload or out-of-spec field values
}

// Construct an outgoing message
conf := ocpp16.BootNotificationConf{
    CurrentTime: time.Now().UTC(),
    Interval:    60,
    Status:      ocpp16.RegistrationStatusAccepted,
}
data, err := json.Marshal(conf)
```

Enum fields are validated on unmarshal. An invalid value returns an error
wrapping `ocpp16.ErrInvalidEnum`, which can be detected with `errors.Is`.

```go
var status ocpp16.AuthorizationStatus
err := json.Unmarshal([]byte(`"Bogus"`), &status)
if errors.Is(err, ocpp16.ErrInvalidEnum) {
    // received an unknown enum value
}
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE.md).
