.PHONY: generate build test test-v lint clean tidy all

# Generate Go code from proto files
generate:
	rm -rf gen
	buf generate

# Build the project
build:
	go build ./...

# Run tests
test:
	go test ./...

# Run tests with verbose output
test-v:
	go test ./... -v

# Run linter (requires golangci-lint)
lint:
	golangci-lint run ./...

# Clean generated files
clean:
	rm -rf gen

# Tidy dependencies
tidy:
	go mod tidy

# Generate, tidy, and build
all: generate tidy build

# CI target: generate, build, test, and lint
ci: generate tidy build test lint
