# Getting Started

This guide will help you get up and running with CDEvents CLI quickly.

## Prerequisites

- Docker (recommended) or Go 1.21+
- Basic understanding of CI/CD concepts
- Familiarity with CloudEvents (optional)

## Installation

### Option 1: Docker (Recommended)

The easiest way to get started is using Docker:

```bash
# Clone the repository
git clone https://github.com/cdevents/cdevents-cli.git
cd cdevents-cli

# Build the Docker image
docker-compose build

# Verify installation
docker run --rm cdevents-cli:latest --version
```

### Option 2: Local Build

If you prefer to build locally:

```bash
# Clone the repository
git clone https://github.com/cdevents/cdevents-cli.git
cd cdevents-cli

# Install dependencies
go mod tidy

# Build the binary
go build -o cdevents-cli

# Verify installation
./cdevents-cli --version
```

## Basic Usage

### Generating Your First Event

Let's generate a simple pipeline started event:

```bash
# Using Docker
docker run --rm cdevents-cli:latest generate pipeline started \
  --id "my-first-pipeline" \
  --name "Getting Started Pipeline"

# Using local binary
./cdevents-cli generate pipeline started \
  --id "my-first-pipeline" \
  --name "Getting Started Pipeline"
```

This will output a CDEvent in JSON format:

```json
{
  "context": {
    "version": "0.4.0",
    "id": "abc123-def456-ghi789",
    "source": "cdevents-cli/hostname",
    "type": "dev.cdevents.pipelinerun.started.0.2.0",
    "timestamp": "2023-12-01T12:00:00Z"
  },
  "subject": {
    "id": "my-first-pipeline",
    "source": "cdevents-cli/hostname",
    "type": "pipelineRun",
    "content": {
      "pipelineName": "Getting Started Pipeline"
    }
  }
}
```

### Output Formats

You can specify different output formats:

```bash
# JSON format (default)
cdevents-cli generate pipeline started --id "test" --name "test" --output json

# YAML format
cdevents-cli generate pipeline started --id "test" --name "test" --output yaml

# CloudEvent format
cdevents-cli generate pipeline started --id "test" --name "test" --output cloudevent
```

### Sending Events

To send events to an HTTP endpoint:

```bash
# Start a test HTTP server (in another terminal)
docker run -p 8080:80 kennethreitz/httpbin

# Send an event
cdevents-cli send --target http://localhost:8080/post \
  pipeline started \
  --id "pipeline-123" \
  --name "my-pipeline"
```

## Configuration

### Environment Variables

You can configure default values using environment variables:

```bash
export CDEVENTS_SOURCE="my-ci-system"
export CDEVENTS_TARGET="http://events.example.com"
export CDEVENTS_OUTPUT="json"
```

### Configuration File

Create a configuration file at `~/.cdevents-cli.yaml`:

```yaml
source: "my-ci-system"
target: "http://events.example.com"
output: "json"
retries: 3
timeout: 30s
```

## Event Types Overview

### Pipeline Events

```bash
# Pipeline queued
cdevents-cli generate pipeline queued --id "pipeline-1" --name "Build Pipeline"

# Pipeline started
cdevents-cli generate pipeline started --id "pipeline-1" --name "Build Pipeline"

# Pipeline finished (success)
cdevents-cli generate pipeline finished --id "pipeline-1" --name "Build Pipeline" --outcome "success"

# Pipeline finished (failure)
cdevents-cli generate pipeline finished --id "pipeline-1" --name "Build Pipeline" --outcome "failure" --errors "Build failed"
```

### Build Events

```bash
# Build queued
cdevents-cli generate build queued --id "build-1" --name "Frontend Build"

# Build started
cdevents-cli generate build started --id "build-1" --name "Frontend Build"

# Build finished
cdevents-cli generate build finished --id "build-1" --name "Frontend Build" --outcome "success"
```

### Service Events

```bash
# Service deployed
cdevents-cli generate service deployed --id "service-1" --name "web-app" --environment "production"

# Service published
cdevents-cli generate service published --id "service-1" --name "web-app" --environment "production"
```

### Test Events

```bash
# Test case started
cdevents-cli generate test testcase-started --id "test-1" --name "Unit Tests"

# Test case finished
cdevents-cli generate test testcase-finished --id "test-1" --name "Unit Tests" --outcome "success"

# Test suite finished
cdevents-cli generate test testsuite-finished --id "testsuite-1" --name "Integration Tests" --outcome "failure"
```

### Custom Data

You can add custom data to events using the `--custom-json` flag:

```bash
# Using JSON format
cdevents-cli generate pipeline started --id "pipeline-1" --name "My Pipeline" \
  --custom-json '{"data": {"build_number": 123, "branch": "main", "commit": "abc123"}, "labels": {"team": "backend"}}'

# Build event with custom data
cdevents-cli generate build finished --id "build-1" --name "Frontend Build" \
  --custom-json '{"data": {"duration": 120, "artifacts": ["app.js", "styles.css"]}, "labels": {"team": "frontend", "env": "prod"}}'

# Service deployment with custom data
cdevents-cli generate service deployed --id "service-1" --name "API Service" \
  --custom-json '{
    "data": {
      "version": "1.2.3",
      "replicas": 3,
      "health_check": "ok"
    },
    "labels": {
      "team": "backend",
      "tier": "api"
    },
    "links": [
      {
        "name": "deployment",
        "url": "https://k8s.example.com/deployment/api-service",
        "type": "kubernetes"
      }
    ]
'
```

## Quick Examples with Input/Output

Here are some quick examples showing exact commands and their outputs:

### Generate a Simple Pipeline Event

**Command:**
```bash
cdevents-cli generate pipeline started --id "my-pipeline-1" --name "Backend Build"
```

**Output:**
```json
{
  "context": {
    "version": "0.4.1",
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "source": "cdevents-cli/hostname",
    "type": "dev.cdevents.pipelinerun.started.0.2.0",
    "timestamp": "2024-01-15T14:30:00Z"
  },
  "subject": {
    "id": "my-pipeline-1",
    "source": "cdevents-cli/hostname",
    "type": "pipelineRun",
    "content": {
      "pipelineName": "Backend Build",
      "url": ""
    }
  }
}
```

### Generate Event with Custom Data

**Command:**
```bash
cdevents-cli generate build finished --id "build-456" --name "Frontend Build" \
  --outcome "success" \
  --custom-json '{
    "data": {
      "duration_seconds": 120,
      "test_coverage": 85.5,
      "artifact_size_mb": 12.3
    },
    "labels": {
      "team": "frontend",
      "environment": "staging"
    }
  }'
```

**Output:**
```json
{
  "context": {
    "version": "0.4.1",
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "source": "cdevents-cli/hostname",
    "type": "dev.cdevents.build.finished.0.2.0",
    "timestamp": "2024-01-15T14:32:00Z"
  },
  "subject": {
    "id": "build-456",
    "source": "cdevents-cli/hostname",
    "type": "build",
    "content": {
      "outcome": "success",
      "url": ""
    }
  },
  "customData": {
    "data": {
      "duration_seconds": 120,
      "test_coverage": 85.5,
      "artifact_size_mb": 12.3
    },
    "labels": {
      "team": "frontend",
      "environment": "staging"
    }
  },
  "customDataContentType": "application/json"
}
```

### Generate Event in YAML Format

**Command:**
```bash
cdevents-cli generate test testcase-finished --id "test-123" --name "Unit Tests" \
  --outcome "success" \
  --output yaml
```

**Output:**
```yaml
context:
  version: "0.4.1"
  id: "550e8400-e29b-41d4-a716-446655440002"
  source: "cdevents-cli/hostname"
  type: "dev.cdevents.testcaserun.finished.0.2.0"
  timestamp: "2024-01-15T14:35:00Z"
subject:
  id: "test-123"
  source: "cdevents-cli/hostname"
  type: "testCaseRun"
  content:
    outcome: "success"
    environment:
      id: "default"
      source: "cdevents-cli/hostname"
    testCase:
      id: "Unit Tests"
      name: "Unit Tests"
      type: "unit"
      uri: ""
```

### Generate Event in CloudEvent Format

**Command:**
```bash
cdevents-cli generate service deployed --id "service-789" --name "API Service" \
  --environment "production" \
  --output cloudevent
```

**Output:**
```json
{
  "specversion": "1.0",
  "id": "550e8400-e29b-41d4-a716-446655440003",
  "source": "cdevents-cli/hostname",
  "type": "dev.cdevents.service.deployed.0.2.0",
  "subject": "service-789",
  "time": "2024-01-15T14:37:00Z",
  "datacontenttype": "application/json",
  "data": {
    "context": {
      "version": "0.4.1",
      "id": "550e8400-e29b-41d4-a716-446655440003",
      "source": "cdevents-cli/hostname",
      "type": "dev.cdevents.service.deployed.0.2.0",
      "timestamp": "2024-01-15T14:37:00Z"
    },
    "subject": {
      "id": "service-789",
      "source": "cdevents-cli/hostname",
      "type": "service",
      "content": {
        "artifactId": "",
        "environment": {
          "id": "production",
          "source": "cdevents-cli/hostname",
          "uri": ""
        }
      }
    }
  }
}
```

### Send Event to Console

**Command:**
```bash
cdevents-cli send pipeline finished --id "pipeline-final" --name "Deployment" \
  --outcome "success" \
  --source "ci-system"
```

**Output:**
```
Event sent to console: pipeline-final
```

### Send Event to File

**Command:**
```bash
cdevents-cli send --target "file://my-events.json" \
  task started --id "task-001" --name "Database Migration" \
  --pipeline "deploy-pipeline"
```

**Output:**
```
Event sent to file my-events.json: task-001
```

**File Contents (my-events.json):**
```json
{
  "context": {
    "version": "0.4.1",
    "id": "550e8400-e29b-41d4-a716-446655440004",
    "source": "cdevents-cli/hostname",
    "type": "dev.cdevents.taskrun.started.0.2.0",
    "timestamp": "2024-01-15T14:40:00Z"
  },
  "subject": {
    "id": "task-001",
    "source": "cdevents-cli/hostname",
    "type": "taskRun",
    "content": {
      "taskName": "Database Migration",
      "pipelineRun": {
        "id": "deploy-pipeline",
        "source": "cdevents-cli/hostname"
      },
      "url": ""
    }
  }
}
```

## Integration Examples

### Jenkins Pipeline

```groovy
pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                // Send pipeline started event
                sh 'cdevents-cli send --target http://events.company.com pipeline started --id "${BUILD_ID}" --name "${JOB_NAME}"'
                
                // Your build steps here
                sh 'make build'
                
                // Send build finished event
                sh 'cdevents-cli send --target http://events.company.com build finished --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "success"'
            }
        }
    }
    
    post {
        success {
            sh 'cdevents-cli send --target http://events.company.com pipeline finished --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "success"'
        }
        failure {
            sh 'cdevents-cli send --target http://events.company.com pipeline finished --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "failure"'
        }
    }
}
```

### GitHub Actions

```yaml
name: CI Pipeline

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Send pipeline started event
      run: |
        docker run --rm cdevents-cli:latest send \
          --target http://events.company.com \
          pipeline started \
          --id "${{ github.run_id }}" \
          --name "${{ github.workflow }}"
    
    - name: Build
      run: make build
    
    - name: Send build finished event
      run: |
        docker run --rm cdevents-cli:latest send \
          --target http://events.company.com \
          build finished \
          --id "${{ github.run_id }}" \
          --name "${{ github.workflow }}" \
          --outcome "success"
```

### GitLab CI

```yaml
stages:
  - build
  - test
  - deploy

variables:
  EVENTS_ENDPOINT: "http://events.company.com"

before_script:
  - docker pull cdevents-cli:latest

build:
  stage: build
  script:
    - docker run --rm cdevents-cli:latest send --target $EVENTS_ENDPOINT pipeline started --id "$CI_PIPELINE_ID" --name "$CI_PROJECT_NAME"
    - make build
    - docker run --rm cdevents-cli:latest send --target $EVENTS_ENDPOINT build finished --id "$CI_PIPELINE_ID" --name "$CI_PROJECT_NAME" --outcome "success"
```

## Troubleshooting

### Common Issues

1. **Connection refused**: Check if your target endpoint is accessible
2. **Authentication errors**: Ensure your HTTP endpoint doesn't require authentication or configure headers
3. **Invalid event format**: Verify your event parameters are correct

### Debug Mode

Enable verbose output for debugging:

```bash
cdevents-cli --verbose generate pipeline started --id "test" --name "test"
```

### Testing Connectivity

Test your HTTP endpoint:

```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -d '{"test": "connectivity"}'
```

## Next Steps

- Explore the [CLI Reference](cli-reference.md) for complete command documentation
- Check out [Examples](examples.md) for more real-world scenarios
- Learn about [Docker deployment](docker.md) options
- Read the [Contributing Guide](contributing.md) to contribute to the project
