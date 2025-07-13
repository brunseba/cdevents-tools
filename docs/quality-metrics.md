# Code Quality Metrics

This document explains how to run and interpret code quality metrics for the CDEvents CLI project.

## Overview

The project uses Docker-based quality analysis to ensure consistent results across different environments. The quality metrics include:

- **Code Coverage**: Tracks test coverage across all packages
- **Cyclomatic Complexity**: Monitors function complexity to ensure maintainability
- **Performance Metrics**: Measures build time and test execution performance
- **Code Linting**: Ensures code quality and consistency

## Quick Start

### Prerequisites

- Docker installed and running
- Make (optional, for convenience)

### Running Quality Metrics

```bash
# Using Make (recommended)
make quality-docker

# Or using Docker directly
docker build -f Dockerfile.quality -t cdevents-cli-quality .
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality
```

### Viewing Results

After running the quality analysis, check the `reports/` directory:

```bash
ls -la reports/
# coverage.out       - Coverage profile
# coverage.html      - HTML coverage report
# quality_report.md  - Comprehensive quality report
```

## Available Commands

### Using Make

```bash
# Show available commands
make help

# Run quality metrics in Docker
make quality-docker

# Run tests with coverage locally
make coverage

# Build the project
make build

# Clean build artifacts
make clean
```

### Using Docker

```bash
# Build the quality analysis image
docker build -f Dockerfile.quality -t cdevents-cli-quality .

# Run full quality analysis
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality

# Run specific analysis
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality bash -c "go tool cover -func=reports/coverage.out"
```

## Quality Metrics Explained

### Code Coverage

**Target**: >70% overall coverage

```bash
# View coverage summary
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality bash -c "go tool cover -func=reports/coverage.out | tail -1"

# Open HTML coverage report
open reports/coverage.html
```

**Package-level coverage**:
- `pkg/transport`: 90.9% (Excellent)
- `pkg/events`: 78.1% (Good)
- `pkg/output`: 78.7% (Good)
- `cmd`: 87.0% (Very Good)

### Cyclomatic Complexity

**Target**: <10 complexity per function

```bash
# Check functions with high complexity
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality bash -c "gocyclo -over 10 ."
```

**Current status**: 5 functions with complexity >10 (acceptable)

### Performance Metrics

**Build Time Target**: <5s  
**Test Execution Target**: <10s

```bash
# Measure build time
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality bash -c "time go build -o /tmp/cdevents-cli ."

# Measure test time
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality bash -c "time go test ./... -short"
```

### Code Linting

```bash
# Run linting
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality bash -c "golangci-lint run"
```

## Interpreting Results

### Coverage Analysis

- **✅ Green (80%+)**: Excellent coverage
- **🟡 Yellow (70-80%)**: Good coverage
- **🔴 Red (<70%)**: Needs improvement

### Complexity Analysis

- **✅ Good**: All functions <10 complexity
- **🟡 Acceptable**: 1-10 functions with 10-20 complexity
- **🔴 Review Needed**: Functions >20 complexity

### Performance Benchmarks

- **✅ Excellent**: Build <3s, Tests <5s
- **🟡 Good**: Build <5s, Tests <10s
- **🔴 Slow**: Build >5s, Tests >10s

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Quality Metrics

on: [push, pull_request]

jobs:
  quality:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Run Quality Metrics
        run: |
          docker build -f Dockerfile.quality -t cdevents-cli-quality .
          docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality
      - name: Upload Coverage
        uses: actions/upload-artifact@v3
        with:
          name: quality-reports
          path: reports/
```

### Quality Gates

The quality analysis enforces these gates:

- **Coverage**: Must be >70%
- **Complexity**: Functions should be <20
- **Build Time**: Should be <5s
- **Test Time**: Should be <10s

## Troubleshooting

### Common Issues

1. **Docker Build Fails**
   ```bash
   # Clean Docker cache
   docker system prune -a
   ```

2. **Permission Issues**
   ```bash
   # Fix permissions
   sudo chown -R $USER:$USER reports/
   ```

3. **Coverage File Not Found**
   ```bash
   # Ensure tests run successfully
   docker run --rm cdevents-cli-quality bash -c "go test ./... -v"
   ```

### Local Development

For local development without Docker:

```bash
# Install tools
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run locally
./scripts/run_quality_metrics.sh
```

## Configuration

### Customizing Thresholds

Edit `scripts/run_quality_metrics.sh`:

```bash
# Coverage threshold
THRESHOLD=80  # Change from 70 to 80

# Complexity threshold
gocyclo -over 15 .  # Change from 10 to 15
```

### Linting Configuration

Customize `.golangci.yml` for linting rules:

```yaml
linters-settings:
  gocyclo:
    min-complexity: 15  # Adjust complexity threshold
  lll:
    line-length: 120   # Line length limit
```

## Best Practices

1. **Run quality metrics before every commit**
2. **Address failing tests immediately**
3. **Monitor coverage trends over time**
4. **Refactor high-complexity functions**
5. **Use quality gates in CI/CD**

## Resources

- [Go Testing Documentation](https://golang.org/pkg/testing/)
- [Cyclomatic Complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity)
- [GolangCI-Lint Documentation](https://golangci-lint.run/)
- [Docker Best Practices](https://docs.docker.com/develop/best-practices/)

---

For questions or improvements to the quality metrics, please open an issue or submit a pull request.
