# MCP Scan Makefile

# Variables
BINARY_NAME=mcp-scan
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME?=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS=-ldflags="-s -w -X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all build clean test coverage deps fmt lint security demo docker help install build-all

all: clean deps test build ## Build everything

# Build for current platform
build: ## Build the binary
	go build $(LDFLAGS) -o $(BINARY_NAME) .

# Build release binaries for all platforms
build-all: clean ## Build for all platforms
	@echo "Building for all platforms..."
	@mkdir -p build
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o build/$(BINARY_NAME)-windows-arm64.exe .
	@echo "Generating checksums..."
	@cd build && sha256sum * > checksums.txt

# Clean build artifacts
clean: ## Clean build artifacts
	rm -rf build/
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

# Run tests
test: ## Run tests
	go test -v -race ./...

coverage: ## Generate test coverage report
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

deps: ## Download dependencies
	go mod download
	go mod verify

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

security: ## Run security checks
	@which gosec > /dev/null || (echo "Installing gosec..." && go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest)
	gosec ./...
	@which govulncheck > /dev/null || (echo "Installing govulncheck..." && go install golang.org/x/vuln/cmd/govulncheck@latest)
	govulncheck ./...

# Run demo scan
demo: build ## Run demo scan
	./$(BINARY_NAME) demo

# Install locally
install: build ## Install binary to system
	sudo cp $(BINARY_NAME) /usr/local/bin/

docker: ## Build Docker image
	docker build -t mcpscan/mcp-scan:latest .

docker-run: ## Run in Docker container  
	docker run --rm -v $(PWD):/workspace mcpscan/mcp-scan:latest demo

release-prep: clean deps test lint security build-all ## Prepare for release
	@echo "Release preparation complete"
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"

# Development commands
dev-scan: build ## Run development scan
	./$(BINARY_NAME) -v

dev-json: build ## Generate JSON report
	./$(BINARY_NAME) -f json -o report.json

help: ## Show this help
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Default target
.DEFAULT_GOAL := help