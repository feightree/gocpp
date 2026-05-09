# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit:
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-ST1020,-ST1021,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## lint: run golangci-lint
.PHONY: lint
lint:
	golangci-lint run

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## benchmark: run all benchmarks
.PHONY: benchmark
benchmark:
	go test -bench=. -benchmem -run=^$$ ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## update: update all packages
.PHONY: update
update:
	go get -u ./...
	go mod tidy

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fix ./...
	go fmt ./...
