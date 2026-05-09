# Contributing

Contributions are welcome. Please read this document before opening a pull request.

## Getting started

**Prerequisites:** Go 1.26+, Make, Homebrew (macOS)

```sh
git clone https://github.com/feightree/gocpp
cd gocpp
make install
```

## Making changes

- Keep changes focused. One concern per pull request.
- All exported types, functions, and constants must have doc comments.
- Run `make audit` and `make test` before submitting. Both must pass.

## Submitting a pull request

Open a pull request against `main`. Include:

- A clear description of what the change does and why.
- Reference to any relevant spec section (e.g. `OCPP 1.6 §6.3`).
- Whether AI tools were used (see below).

## Code style

- Formatting is enforced by `gofmt`. Run `make tidy` before committing.
- Linting is enforced by `golangci-lint`. Run `make lint` to check.
- Follow standard Go conventions: https://go.dev/doc/effective_go

## AI usage policy

AI tools are permitted. Please disclose in your pull request description:

- Whether AI was used (yes / no / partially)
- Which tool(s), if applicable (e.g. Claude, Copilot)
- That you have read and understood all code you are submitting

Pull request review discussion should come from you, not pasted from an AI.

## Spec compliance

This library implements [OCPP 1.6](https://www.openchargealliance.org/protocols/ocpp-16/).
Changes must remain compliant with the specification. If you believe the spec
is ambiguous or incorrect, open an issue to discuss before implementing.
