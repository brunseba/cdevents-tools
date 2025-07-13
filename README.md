# CDEvents CLI

A command-line tool for generating and sending CDEvents into CI/CD toolchains using CloudEvents as transport.

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)
[![Docker](https://img.shields.io/badge/docker-available-blue.svg)](https://hub.docker.com/r/cdevents/cdevents-cli)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

## Overview

CDEvents CLI is a powerful tool designed to integrate with your CI/CD pipeline by generating and transmitting standardized CDEvents. It helps create interoperability between different CI/CD tools and provides observability into your software delivery process.

## Features

- ✅ **Generate CDEvents**: Create various types of CDEvents (pipeline, task, build, service, test)
- ✅ **Custom Data Support**: Add custom data, labels, annotations, and links to events
- ✅ **Multiple Input Formats**: Support JSON, YAML, and key=value pairs for custom data
- ✅ **Multiple Transports**: Send events via HTTP, console output, or file
- ✅ **CloudEvents Compatible**: Full CloudEvents specification support
- ✅ **Flexible Configuration**: Command-line flags and configuration files
- ✅ **Docker Support**: Containerized deployment with multi-platform binaries
- ✅ **Retry Logic**: Built-in retry mechanisms with configurable timeouts
- ✅ **CI/CD Integration**: Easy integration with Jenkins, GitHub Actions, GitLab CI, etc.

## Quick Start

### Docker (Recommended)

```bash
# Build the Docker image
docker-compose build

# Generate a pipeline started event
docker run --rm cdevents-cli:latest generate pipeline started \
  --id "pipeline-123" \
  --name "my-pipeline"

# Send an event to an HTTP endpoint
docker run --rm cdevents-cli:latest send \
  --target http://localhost:8080/events \
  pipeline started \
  --id "pipeline-123" \
  --name "my-pipeline"
```

### Local Installation

```bash
# Clone the repository
git clone https://github.com/cdevents/cdevents-cli.git
cd cdevents-cli

# Install dependencies
go mod tidy

# Build the binary
go build -o cdevents-cli

# Run the CLI
./cdevents-cli generate pipeline started \
  --id "pipeline-123" \
  --name "my-pipeline"
```

## Event Types Supported

| Category | Event Types | Description |
|----------|-------------|-------------|
| **Pipeline** | `queued`, `started`, `finished` | Pipeline execution events |
| **Task** | `started`, `finished` | Individual task execution events |
| **Build** | `queued`, `started`, `finished` | Build process events |
| **Service** | `deployed`, `published`, `removed`, `rolledback`, `upgraded` | Service deployment events |
| **Test** | `testcase-queued`, `testcase-started`, `testcase-finished`, `testcase-skipped`, `testsuite-queued`, `testsuite-started`, `testsuite-finished`, `testoutput-published` | Testing events |

## Usage Examples

### Generate Events

```bash
# Generate a pipeline started event
cdevents-cli generate pipeline started --id "pipeline-123" --name "my-pipeline"

# Generate a build finished event with outcome
cdevents-cli generate build finished --id "build-456" --name "my-build" --outcome "success"

# Generate a service deployed event
cdevents-cli generate service deployed --id "service-789" --name "my-service" --environment "production"

# Generate events with custom data
cdevents-cli generate pipeline started --id "pipeline-123" --name "my-pipeline" \
  --custom "build_number=456" --custom "branch=main"

# Generate events with JSON custom data
cdevents-cli generate build finished --id "build-456" --name "my-build" \
  --custom-json '{"data": {"duration": 120, "tests_passed": 150}, "labels": {"team": "backend"}}'
```

### Send Events

```bash
# Send to console (default)
cdevents-cli send pipeline started --id "pipeline-123" --name "my-pipeline"

# Send to HTTP endpoint
cdevents-cli send --target http://localhost:8080/events pipeline started --id "pipeline-123" --name "my-pipeline"

# Send to file
cdevents-cli send --target file://events.json pipeline started --id "pipeline-123" --name "my-pipeline"
```

### CI/CD Integration

#### Jenkins

```groovy
pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'cdevents-cli send --target http://events.company.com pipeline started --id "${BUILD_ID}" --name "${JOB_NAME}"'
                sh 'make build'
                sh 'cdevents-cli send --target http://events.company.com build finished --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "success"'
            }
        }
    }
}
```

#### GitHub Actions

```yaml
- name: Send pipeline started event
  run: |
    docker run --rm cdevents-cli:latest send \
      --target http://events.company.com \
      pipeline started \
      --id "${{ github.run_id }}" \
      --name "${{ github.workflow }}"
```

## Configuration

### Configuration File

Create `~/.cdevents-cli.yaml`:

```yaml
source: "my-ci-system"
target: "http://events.example.com"
output: "json"
retries: 3
timeout: 30s
```

### Environment Variables

```bash
export CDEVENTS_SOURCE="my-ci-system"
export CDEVENTS_TARGET="http://events.example.com"
export CDEVENTS_OUTPUT="json"
```

## Documentation

- [Getting Started](https://cdevents.github.io/cdevents-cli/getting-started/) - Detailed setup and usage instructions
- [CLI Reference](https://cdevents.github.io/cdevents-cli/cli-reference/) - Complete command reference
- [Examples](https://cdevents.github.io/cdevents-cli/examples/) - Real-world usage examples
- [Docker Guide](https://cdevents.github.io/cdevents-cli/docker/) - Docker deployment instructions

## Development

- [Contributing Guide](https://cdevents.github.io/cdevents-cli/contributing/) - How to contribute to the project
- [Development Standards](https://cdevents.github.io/cdevents-cli/development-standards/) - Development standards and practices

## Contributing

We welcome contributions! Please see our [Contributing Guide](https://cdevents.github.io/cdevents-cli/contributing/) for details and review our [Development Standards](https://cdevents.github.io/cdevents-cli/development-standards/) for coding practices and quality requirements.

## Community

- **Slack**: Join the [CDEvents Slack](https://cdeliveryfdn.slack.com/archives/C030SKZ0F4K)
- **GitHub**: [CDEvents CLI Repository](https://github.com/cdevents/cdevents-cli)
- **Documentation**: [CDEvents Official Docs](https://cdevents.dev/docs/)

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Related Projects

- [CDEvents Specification](https://github.com/cdevents/spec)
- [CDEvents Go SDK](https://github.com/cdevents/sdk-go)
- [CloudEvents](https://cloudevents.io/)

## Acknowledgments

CDEvents CLI is built on top of the excellent work by the CDEvents and CloudEvents communities. Special thanks to all contributors and the Continuous Delivery Foundation.
