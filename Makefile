.PHONY: help build run test clean lint fmt install-deps

help:
	@echo "ClawCLI - AI-Powered CLI Assistant"
	@echo ""
	@echo "Available commands:"
	@echo "  make build        - Build the binary"
	@echo "  make run          - Run the application"
	@echo "  make test         - Run all tests"
	@echo "  make test-verbose - Run tests with verbose output"
	@echo "  make coverage     - Run tests with coverage report"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make lint         - Run linter (requires golangci-lint)"
	@echo "  make fmt          - Format code"
	@echo "  make deps         - Download dependencies"
	@echo "  make install      - Build and install binary"
	@echo "  make help         - Show this help message"

build:
	@echo "Building ClawCLI..."
	go build -o clawcli.exe -v

run: build
	@echo "Running ClawCLI..."
	./clawcli.exe

test:
	@echo "Running tests..."
	go test ./...

test-verbose:
	@echo "Running tests (verbose)..."
	go test -v ./...

coverage:
	@echo "Running tests with coverage..."
	go test -cover ./...
	@echo ""
	@echo "For detailed coverage report:"
	@echo "  go test -coverprofile=coverage.out ./..."
	@echo "  go tool cover -html=coverage.out"

clean:
	@echo "Cleaning build artifacts..."
	rm -f clawcli.exe clawcli coverage.out

lint:
	@echo "Running linter..."
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; exit 1; }
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	go fmt ./...

deps:
	@echo "Downloading dependencies..."
	go mod tidy
	go mod download

install: test build
	@echo "Installation complete!"
	@echo "Binary: ./clawcli.exe"

# For development with file watching (requires fswatch or similar)
watch:
	@command -v fswatch >/dev/null 2>&1 || { echo "fswatch not installed. Install with: brew install fswatch"; exit 1; }
	fswatch -r . --ignore build --ignore dist | xargs -I{} make test

.DEFAULT_GOAL := help
