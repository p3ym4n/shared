# Makefile
#
# - config - tooling versions and variable configs
# - test - run tests
# - run - run application
# - tools - golang useful tools
# - development - scripts for development

# === CONFIG =======================================================
cover_profile_filename := cover.out

# === BUILD =======================================================
# Generate mocks for interfaces
.PHONY: generate-mocks
generate-mocks:
	go generate ./...

# === TEST =======================================================
# Run all tests.
test:
	go test ./...

# Run all tests with verbose output.
test-verbose:
	go test -v ./...

# Test how much of a package’s code is exercised by running the package’s tests.
test-cover:
	go test -cover -coverprofile=$(cover_profile_filename) ./...

# Test data race and package coverage
test-pre-commit:
	go test -race -cover -coverprofile=$(cover_profile_filename) ./...

# === RUN =======================================================
# Run application using linters: it runs linters in parallel, uses caching, supports yaml config, etc.
lint:
	golangci-lint run ./... --timeout=3m

# === TOOLS =======================================================
# Get a decorated HTML presentation of cover file: showing the covered (green), uncovered (red), and uninstrumented (grey) source.
tool-read-cover:
	go tool cover -html=$(cover_profile_filename)

# Format go code
tool-go-format:
	 go fmt ./...

# Clean up unused dependencies (fix go.mod and go.sum either)
tool-mod-tidy:
	go mod tidy

# === DEVELOPMENT =======================================================
dev-pre-commit: generate-mocks tool-mod-tidy tool-go-format lint test-pre-commit
