# Docker Guide

Learn how to deploy CDEvents CLI with Docker.

## Building the Docker Image

To build the Docker image for the CDEvents CLI:

```bash
docker-compose build
```

The Dockerfile is configured to build multi-platform binaries using Go and create an Alpine-based container image.

## Running the CLI with Docker

You can run the CDEvents CLI in a container:

```bash
docker run --rm cdevents-cli:latest [command] [flags]
```

### Examples

Generate a pipeline started event:

```bash
docker run --rm cdevents-cli:latest generate pipeline started --id "pipeline-123" --name "my-pipeline"
```

Send an event to an HTTP endpoint:

```bash
docker run --rm cdevents-cli:latest send --target http://localhost:8080/events pipeline started --id "pipeline-123" --name "my-pipeline"
```

## Using docker-compose

Use `docker-compose` for more complex scenarios, such as including HTTP endpoints or other services:

```yaml
version: '3.8'

services:
  cdevents-cli:
    build:
      context: .
      dockerfile: Dockerfile
    image: cdevents-cli:latest

  # Example HTTP endpoint
  httpbin:
    image: kennethreitz/httpbin:latest
    container_name: httpbin
    ports:
      - "8080:80"
```

Run the setup:

```bash
docker-compose up
```

## Configuration with Docker

You can pass environment variables to Docker containers:

```bash
docker run --rm \
  -e CDEVENTS_SOURCE="my-docker-source" \
  -e CDEVENTS_TARGET="http://example.com/events" \
  cdevents-cli:latest send pipeline started --id "pipeline-123" --name "my-pipeline"
```

### Using Volume Mounts

Mount volumes to store output files or logs:

```bash
docker run --rm \
  -v $(pwd)/output:/output \
  cdevents-cli:latest generate pipeline started --id "pipeline-123" --name "my-pipeline" --output file:///output/event.json
```

## Deployment Strategies

- **Local Development**: Test and develop using `docker-compose` for isolated environments.
- **CI/CD Systems**: Integrate the Docker image with Jenkins, GitLab CI, or GitHub Actions.
- **Kubernetes**: Deploy the Docker image in a Kubernetes cluster to handle event generation at scale.

## Troubleshooting Docker

### Common Issues

1. **Permission Denied**: Ensure Docker has the necessary permissions to access volumes or network.
2. **Image Not Found**: Make sure the image is built before running.
3. **Networking Issues**: Check Docker network settings if containers are unable to communicate.

## Security Considerations

- Run containers as non-root wherever possible.
- Limit network access to only the necessary endpoints.
- Regularly update the base image to include security patches.
