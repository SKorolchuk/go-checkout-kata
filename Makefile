# ==============================================================================
# Build

fmt:
	go fmt ./...

vet:
	go vet -v ./...

clean:
	$(RM) -r bin
	go clean ./...

build: clean fmt vet
	go build -o bin/checkout-cli ./cmd/checkout-cli

run-example:
	go run ./cmd/checkout-cli/main.go --scan-series AABCAA

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
