# Makefile
#
# - config - tooling versions and variable configs
# - test - run tests
# - run - run application
# - tools - golang useful tools
# - development - scripts for development

# === CONFIG =======================================================
cover_profile_filename := cover.out
dirs = $(subst ./,,$(dir $(shell find ./* | grep "go\.mod" | sort )))
mockdirs = $(shell grep -irl . -e '//go:generate' --exclude Makefile)

# === BUILD =======================================================
# Generate mocks for interfaces
.PHONY: generate-mocks
generate-mocks:
	for i in $(mockdirs); do \
		go generate $$i ; \
	done; \

# === TEST =======================================================
# Run all tests.
.PHONY: test
test:
	for i in $(dirs); do \
		go test $$i...; \
	done; \

## Run all tests with verbose output.
.PHONY: test-verbose
test-verbose:
	for i in $(dirs); do \
		go test -v $$i...; \
	done; \

# Test how much of a package’s code is exercised by running the package’s tests.
.PHONY: test-cover
test-cover:
	for i in $(dirs); do \
		go test -cover -coverprofile=$(cover_profile_filename) $$i...; \
	done; \

# Test data race and package coverage
.PHONY: test-pre-commit
test-pre-commit:
	for i in $(dirs); do \
		go test -race -cover -coverprofile=$(cover_profile_filename) $$i...; \
	done; \

# === RUN =======================================================
# Run application using linters: it runs linters in parallel, uses caching, supports yaml config, etc.
.PHONY: lint
lint:
	for i in $(dirs); do \
		cd $$i ; \
  		golangci-lint run --timeout=3m ; \
	done; \

# === TOOLS =======================================================
# Format go code
.PHONY: tool-go-format
tool-go-format:
	for i in $(dirs); do \
		go fmt $$i... ; \
	done; \

# Clean up unused dependencies (fix go.mod and go.sum either)
.PHONY: tool-mod-tidy
tool-mod-tidy:
	for i in $(dirs); do \
  		cd ./$$i ; \
		go mod tidy ; \
	done; \

# === DEVELOPMENT =======================================================
.PHONY: dev-pre-commit
dev-pre-commit: generate-mocks tool-mod-tidy tool-go-format lint test-pre-commit
