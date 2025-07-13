# CDEvents CLI Makefile

.PHONY: quality quality-docker test coverage build clean install help

# Default target
help:
	@echo "CDEvents CLI - Available commands:"
	@echo "  make quality         - Run quality metrics locally"
	@echo "  make quality-docker  - Run quality metrics in Docker"
	@echo "  make test           - Run tests"
	@echo "  make coverage       - Run tests with coverage"
	@echo "  make build          - Build the binary"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make install        - Install dependencies"

# Install dependencies
install:
	go mod tidy
	go mod download

# Run tests
test:
	go test ./... -v

# Run tests with coverage
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Build the binary
build:
	go build -o cdevents-cli .

# Clean build artifacts
clean:
	rm -f cdevents-cli
	rm -f coverage.out coverage.html
	rm -rf reports/

# Run quality metrics locally (requires tools to be installed)
quality:
	./scripts/run_quality_metrics.sh

# Run quality metrics in Docker
quality-docker:
	@echo "Building Docker image for quality metrics..."
	docker build -f Dockerfile.quality -t cdevents-cli-quality .
	@echo "Running quality metrics in Docker..."
	docker run --rm -v $(PWD)/reports:/app/reports cdevents-cli-quality

# Run quality metrics and extract reports
quality-extract:
	@echo "Building Docker image for quality metrics..."
	docker build -f Dockerfile.quality -t cdevents-cli-quality .
	@echo "Running quality metrics in Docker..."
	docker run --rm -v $(PWD)/reports:/app/reports cdevents-cli-quality
	@echo "Quality reports extracted to ./reports/"
