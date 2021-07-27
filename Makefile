# ==============================================================================
# Build

fmt:
	go fmt ./...

vet:
	go vet -v ./...

clean:
	go clean

# Define targets for commands
bin/checkout-cli:
	go build -o bin/checkout-cli ./cmd/checkout-cli

build: clean fmt vet bin/checkout-cli

# ==============================================================================
# Running tests

tests:
	go test ./... -count=1 -v

lint:
	staticcheck -checks=all ./...

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all
