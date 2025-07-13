# CDEvents CLI Docs

## Overview

CDEvents CLI is a command-line tool for generating and sending CDEvents 
into CI/CD toolchains using CloudEvents as transport.

### Features

- Generate various CDEvents (pipeline, task, build, deployment, etc.)
- Send events via HTTP, or outputs to console and file
- Supports CloudEvents format
- Retries and timeout configuration
- Integration with CI/CD systems

### Getting Started

1. **Build the Docker Image**

   Use the following command to build the Docker image:

   ```sh
   docker-compose build
   ```

2. **Running the CLI**

   Run the CLI with Docker:

   ```sh
   docker run --rm cdevents-cli:latest pipeline started --id "pipeline-123" --name "my-pipeline"
   ```

3. **Sending Events**

   Use the send command to send generated events to various targets:

   ```sh
   docker run --rm cdevents-cli:latest send --target http://localhost:8080/events pipeline started --id "pipeline-123" --name "my-pipeline"
   ```

### Configuration

CDEvents CLI can be configured using command-line flags or environment variables.

- **Flags:** Use `--help` with any command to see available flags.
- **Environment Variables:** Configure default values using environment variables:
   - `CDEVENTS_SOURCE` to set the default source.

### Development

Contributions are welcome! Here's how to set up a local development environment:

1. **Clone the Repository**

   ```sh
   git clone <repository-url>
   ```

2. **Install Dependencies**

   ```sh
   go mod tidy
   ```

3. **Run Tests**

   ```sh
   go test ./...
   ```
