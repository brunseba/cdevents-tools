# CDEvents CLI

A command-line tool for generating and sending CDEvents into CI/CD toolchains using CloudEvents as transport.

## Overview

CDEvents CLI is a powerful tool designed to integrate with your CI/CD pipeline by generating and transmitting standardized CDEvents. It helps create interoperability between different CI/CD tools and provides observability into your software delivery process.

## Key Features

- ✅ **Generate CDEvents**: Create various types of CDEvents (pipeline, task, build, service, test)
- ✅ **Custom Data Support**: Add custom data, labels, annotations, and links to events
- ✅ **Multiple Input Formats**: Support JSON, YAML, and key=value pairs for custom data
- ✅ **Multiple Transports**: Send events via HTTP, console output, or file
- ✅ **CloudEvents Compatible**: Full CloudEvents v1.0 specification support with binary encoding
- ✅ **Standard Compliance**: Follows CDEvents v0.4.1 and CloudEvents v1.0 standards
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

## CloudEvents Compatibility

This CLI tool provides full compatibility with the [CloudEvents v1.0 specification](https://cloudevents.io/), ensuring seamless integration with CloudEvents-compatible systems:

### Transport Layer
- **HTTP Binary Mode**: Events are sent as CloudEvents using binary content mode over HTTP
- **Event Mapping**: Automatic mapping between CDEvents and CloudEvents attributes
- **Standard Headers**: Proper CloudEvents headers (ce-specversion, ce-type, ce-source, etc.)
- **JSON Payload**: CDEvents data is embedded in CloudEvents `data` field as JSON

### Event Transformation

| CDEvents Attribute | CloudEvents Attribute | Description |
|-------------------|----------------------|-------------|
| `id` | `id` | Unique event identifier |
| `source` | `source` | Event source URI |
| `type` | `type` | Event type (e.g., `dev.cdevents.pipeline.started.0.2.0`) |
| `subject.id` | `subject` | Subject identifier |
| `timestamp` | `time` | Event timestamp |
| Event data | `data` | Full CDEvents payload including custom data |

### Output Formats
- **Native CDEvents**: JSON/YAML output in CDEvents format
- **CloudEvents**: JSON output in CloudEvents envelope format
- **Transport Ready**: Events ready for CloudEvents-compatible message brokers

### Integration Benefits
- **Ecosystem Compatibility**: Works with any CloudEvents-compatible system
- **Message Brokers**: Compatible with Kafka, Pulsar, RabbitMQ (with CloudEvents plugins)
- **Cloud Platforms**: Native support for AWS EventBridge, Google Cloud Pub/Sub, Azure Event Grid
- **Observability**: Compatible with OpenTelemetry and other observability tools

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
