.PHONY: build clean test run install dev lint fmt vet deps help

# Variables
BINARY_NAME=oracle
BUILD_DIR=dist
MAIN_PATH=./main.go

# Default target
all: build

# Build the application
build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)
	rm -rf dist/

# Run tests
test:
	go test -v ./...

# Run the application
run:
	go run $(MAIN_PATH)

# Install the application to GOPATH/bin
install:
	go install $(MAIN_PATH)

# Development mode with live reload (requires air)
dev:
	air

# Lint the code
lint:
	golangci-lint run

# Format the code
fmt:
	go fmt ./...

# Vet the code
vet:
	go vet ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Install development tools
dev-deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Release using goreleaser (dry run)
release-dry:
	goreleaser release --snapshot --clean

# Create a new release tag
tag:
	@read -p "Enter version (e.g., v1.0.0): " version; \
	git tag -a $$version -m "Release $$version"; \
	git push origin $$version

# Help
help:
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  build-all   - Build for multiple platforms"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo "  run         - Run the application"
	@echo "  install     - Install to GOPATH/bin"
	@echo "  dev         - Development mode with live reload"
	@echo "  lint        - Lint the code"
	@echo "  fmt         - Format the code"
	@echo "  vet         - Vet the code"
	@echo "  deps        - Download dependencies"
	@echo "  dev-deps    - Install development tools"
	@echo "  release-dry - Test release with goreleaser"
	@echo "  tag         - Create a new release tag"
	@echo "  help        - Show this help"
