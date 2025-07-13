# CDEvents CLI

A command-line tool for generating and sending CDEvents into CI/CD toolchains using CloudEvents as transport.

## Overview

CDEvents CLI is a powerful tool designed to integrate with your CI/CD pipeline by generating and transmitting standardized CDEvents. It helps create interoperability between different CI/CD tools and provides observability into your software delivery process.

## Key Features

- ✅ **Generate CDEvents**: Create various types of CDEvents (pipeline, task, build, service, test)
- ✅ **Custom Data Support**: Add custom data, labels, annotations, and links to events
- ✅ **Multiple Input Formats**: Support JSON, YAML, and key=value pairs for custom data
- ✅ **Multiple Transports**: Send events via HTTP, console output, or file
- ✅ **CloudEvents Compatible**: Full CloudEvents specification support
- ✅ **Flexible Configuration**: Command-line flags and configuration files
- ✅ **Docker Support**: Containerized deployment with multi-platform binaries
- ✅ **Retry Logic**: Built-in retry mechanisms with configurable timeouts
- ✅ **CI/CD Integration**: Easy integration with Jenkins, GitHub Actions, GitLab CI, etc.

## What are CDEvents?

CDEvents is a common specification for Continuous Delivery events, enabling interoperability in the complete software production ecosystem. It's an incubated project at the [Continuous Delivery Foundation](https://cd.foundation) (CDF).

CDEvents extends CloudEvents by introducing purpose and semantics to events, providing:

- **Standardized Event Types**: Common vocabulary for CI/CD events
- **Event Linking**: Ability to link related events together
- **Observability**: Enhanced visibility into your delivery pipeline
- **Interoperability**: Common format across different tools

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

## Transport Options

- **HTTP**: Send events to HTTP endpoints using CloudEvents format
- **Console**: Output events to console in JSON, YAML, or CloudEvent format
- **File**: Write events to files for later processing

## Next Steps

- [Getting Started](getting-started.md) - Detailed setup and usage instructions
- [CLI Reference](cli-reference.md) - Complete command reference
- [Examples](examples.md) - Real-world usage examples
- [Docker Guide](docker.md) - Docker deployment instructions
- [Contributing](contributing.md) - How to contribute to the project
- [Development Standards](development-standards.md) - Development standards and practices

## Community

- **Slack**: Join the [CDEvents Slack](https://cdeliveryfdn.slack.com/archives/C030SKZ0F4K)
- **GitHub**: [CDEvents CLI Repository](https://github.com/cdevents/cdevents-cli)
- **Documentation**: [CDEvents Official Docs](https://cdevents.dev/docs/)

## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.
