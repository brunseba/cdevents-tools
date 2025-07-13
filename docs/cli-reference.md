# CLI Reference

Complete command reference for CDEvents CLI.

## Global Options

These options are available for all commands:

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--config` | | Config file path | `$HOME/.cdevents-cli.yaml` |
| `--output` | `-o` | Output format (json, yaml, cloudevent) | `json` |
| `--verbose` | `-v` | Verbose output | `false` |
| `--help` | `-h` | Show help | |
| `--version` | | Show version | |

## Root Command

```bash
cdevents-cli [command] [flags]
```

A CLI tool for generating and sending CDEvents into CI/CD toolchains using CloudEvents as transport.

## Commands

### generate

Generate CDEvents for various CI/CD activities.

```bash
cdevents-cli generate [event-type] [sub-command] [flags]
```

#### Common Generate Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--id` | `-i` | Subject ID | ✅ |
| `--name` | `-n` | Subject name | ✅ |
| `--source` | `-s` | Event source | |
| `--url` | `-u` | Subject URL | |
| `--outcome` | | Outcome (success, failure, error, cancel) | |
| `--errors` | | Error details | |
| `--custom-json` | | Custom data in JSON format | |

#### Pipeline Events

Generate pipeline run events.

```bash
cdevents-cli generate pipeline [queued|started|finished] [flags]
```

**Examples:**

```bash
# Generate pipeline queued event
cdevents-cli generate pipeline queued --id "pipeline-123" --name "my-pipeline"

# Generate pipeline started event
cdevents-cli generate pipeline started --id "pipeline-123" --name "my-pipeline" --url "https://ci.example.com/pipeline-123"

# Generate pipeline finished event (success)
cdevents-cli generate pipeline finished --id "pipeline-123" --name "my-pipeline" --outcome "success"

# Generate pipeline finished event (failure)
cdevents-cli generate pipeline finished --id "pipeline-123" --name "my-pipeline" --outcome "failure" --errors "Build step failed"

# Generate pipeline with JSON custom data
cdevents-cli generate pipeline finished --id "pipeline-123" --name "my-pipeline" \
  --custom-json '{"data": {"duration": 300, "tests_passed": 150, "build_number": 456, "branch": "feature/new-api"}, "labels": {"team": "backend"}}'
```

#### Build Events

Generate build events.

```bash
cdevents-cli generate build [queued|started|finished] [flags]
```

**Examples:**

```bash
# Generate build queued event
cdevents-cli generate build queued --id "build-456" --name "frontend-build"

# Generate build started event
cdevents-cli generate build started --id "build-456" --name "frontend-build"

# Generate build finished event
cdevents-cli generate build finished --id "build-456" --name "frontend-build" --outcome "success"
```

#### Task Events

Generate task run events.

```bash
cdevents-cli generate task [started|finished] [flags]
```

**Additional Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--pipeline` | `-p` | Pipeline ID |

**Examples:**

```bash
# Generate task started event
cdevents-cli generate task started --id "task-789" --name "unit-tests" --pipeline "pipeline-123"

# Generate task finished event
cdevents-cli generate task finished --id "task-789" --name "unit-tests" --pipeline "pipeline-123" --outcome "success"
```

#### Service Events

Generate service deployment events.

```bash
cdevents-cli generate service [deployed|published|removed|rolledback|upgraded] [flags]
```

**Additional Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--environment` | `-e` | Environment ID |

**Examples:**

```bash
# Generate service deployed event
cdevents-cli generate service deployed --id "service-001" --name "web-app" --environment "production"

# Generate service published event
cdevents-cli generate service published --id "service-001" --name "web-app" --environment "production"

# Generate service removed event
cdevents-cli generate service removed --id "service-001" --name "web-app" --environment "staging"
```

#### Test Events

Generate test events.

```bash
cdevents-cli generate test [event-type] [flags]
```

**Test Event Types:**

- `testcase-queued`
- `testcase-started`
- `testcase-finished`
- `testcase-skipped`
- `testsuite-queued`
- `testsuite-started`
- `testsuite-finished`
- `testoutput-published`

**Examples:**

```bash
# Generate test case started event
cdevents-cli generate test testcase-started --id "test-001" --name "unit-tests"

# Generate test case finished event
cdevents-cli generate test testcase-finished --id "test-001" --name "unit-tests" --outcome "success"

# Generate test suite finished event
cdevents-cli generate test testsuite-finished --id "testsuite-001" --name "integration-tests" --outcome "failure" --errors "Database connection failed"
```

### send

Send CDEvents to various targets.

```bash
cdevents-cli send [flags] [event-type] [sub-command] [flags]
```

#### Send Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--target` | `-t` | Target destination | `console` |
| `--retries` | `-r` | Number of retry attempts | `3` |
| `--timeout` | | Request timeout | `30s` |
| `--headers` | `-H` | HTTP headers (key=value format) | |

#### Target Formats

| Format | Description | Example |
|--------|-------------|---------|
| `console` | Output to console | `console` |
| `http://...` | HTTP endpoint | `http://localhost:8080/events` |
| `https://...` | HTTPS endpoint | `https://events.example.com/webhook` |
| `file://...` | File output | `file://events.json` |

#### Send Examples

```bash
# Send to console (default)
cdevents-cli send pipeline started --id "pipeline-123" --name "my-pipeline"

# Send to HTTP endpoint
cdevents-cli send --target http://localhost:8080/events pipeline started --id "pipeline-123" --name "my-pipeline"

# Send to HTTPS endpoint with custom headers
cdevents-cli send --target https://events.example.com/webhook --headers "Authorization=Bearer token123" pipeline started --id "pipeline-123" --name "my-pipeline"

# Send to file
cdevents-cli send --target file://pipeline-events.json pipeline started --id "pipeline-123" --name "my-pipeline"

# Send with retry configuration
cdevents-cli send --target http://localhost:8080/events --retries 5 --timeout 60s pipeline started --id "pipeline-123" --name "my-pipeline"
```

## Custom Data

CDEvents CLI supports adding custom data to events in JSON format:

### JSON Format

```bash
# Add structured data in JSON format
cdevents-cli generate build finished --id "build-456" --name "my-build" \
  --custom-json '{
    "data": {
      "duration": 120,
      "artifacts": ["app.js", "styles.css"],
      "coverage": 85.5
    },
    "labels": {
      "team": "frontend",
      "environment": "production"
    },
    "annotations": {
      "build.tool": "webpack",
      "build.version": "5.0.0"
    },
    "links": [
      {
        "name": "build-logs",
        "url": "https://ci.example.com/build/456/logs",
        "type": "logs"
      }
    ]
  }'
```

### Custom Data Structure

The custom data structure supports:

- **data**: Arbitrary key-value pairs with any JSON-compatible values
- **labels**: String key-value pairs for categorization
- **annotations**: String key-value pairs for metadata
- **links**: Array of related links with name, URL, and optional type

## Configuration

### Configuration File

The CLI looks for configuration in the following order:

1. `--config` flag
2. `$HOME/.cdevents-cli.yaml`
3. Environment variables
4. Default values

Example configuration file:

```yaml
# Default event source
source: "my-ci-system"

# Default target for send commands
target: "http://events.example.com"

# Default output format
output: "json"

# Default retry settings
retries: 3
timeout: 30s

# Default headers for HTTP requests
headers:
  - "Authorization=Bearer token123"
  - "X-Custom-Header=value"
```

### Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `CDEVENTS_SOURCE` | Default event source | `my-ci-system` |
| `CDEVENTS_TARGET` | Default target | `http://events.example.com` |
| `CDEVENTS_OUTPUT` | Default output format | `json` |
| `CDEVENTS_RETRIES` | Default retry count | `3` |
| `CDEVENTS_TIMEOUT` | Default timeout | `30s` |

## Output Formats

### JSON (Default)

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
    "id": "pipeline-123",
    "source": "cdevents-cli/hostname",
    "type": "pipelineRun",
    "content": {
      "pipelineName": "my-pipeline"
    }
  }
}
```

### YAML

```yaml
context:
  version: "0.4.0"
  id: "abc123-def456-ghi789"
  source: "cdevents-cli/hostname"
  type: "dev.cdevents.pipelinerun.started.0.2.0"
  timestamp: "2023-12-01T12:00:00Z"
subject:
  id: "pipeline-123"
  source: "cdevents-cli/hostname"
  type: "pipelineRun"
  content:
    pipelineName: "my-pipeline"
```

### CloudEvent

```json
{
  "specversion": "1.0",
  "type": "dev.cdevents.pipelinerun.started.0.2.0",
  "source": "cdevents-cli/hostname",
  "id": "abc123-def456-ghi789",
  "time": "2023-12-01T12:00:00Z",
  "datacontenttype": "application/json",
  "data": {
    "context": {
      "version": "0.4.0",
      "id": "abc123-def456-ghi789",
      "source": "cdevents-cli/hostname",
      "type": "dev.cdevents.pipelinerun.started.0.2.0",
      "timestamp": "2023-12-01T12:00:00Z"
    },
    "subject": {
      "id": "pipeline-123",
      "source": "cdevents-cli/hostname",
      "type": "pipelineRun",
      "content": {
        "pipelineName": "my-pipeline"
      }
    }
  }
}
```

## Event Types Reference

### Pipeline Events

| Event Type | Description | Required Fields | Optional Fields |
|------------|-------------|----------------|----------------|
| `queued` | Pipeline queued for execution | `id`, `name` | `url` |
| `started` | Pipeline execution started | `id`, `name` | `url` |
| `finished` | Pipeline execution completed | `id`, `name` | `url`, `outcome`, `errors` |

### Build Events

| Event Type | Description | Required Fields | Optional Fields |
|------------|-------------|----------------|----------------|
| `queued` | Build queued for execution | `id`, `name` | `url` |
| `started` | Build execution started | `id`, `name` | `url` |
| `finished` | Build execution completed | `id`, `name` | `url`, `outcome`, `errors` |

### Task Events

| Event Type | Description | Required Fields | Optional Fields |
|------------|-------------|----------------|----------------|
| `started` | Task execution started | `id`, `name` | `pipeline`, `url` |
| `finished` | Task execution completed | `id`, `name` | `pipeline`, `url`, `outcome`, `errors` |

### Service Events

| Event Type | Description | Required Fields | Optional Fields |
|------------|-------------|----------------|----------------|
| `deployed` | Service deployed | `id`, `name` | `environment`, `url` |
| `published` | Service published | `id`, `name` | `environment`, `url` |
| `removed` | Service removed | `id`, `name` | `environment`, `url` |
| `rolledback` | Service rolled back | `id`, `name` | `environment`, `url` |
| `upgraded` | Service upgraded | `id`, `name` | `environment`, `url` |

### Test Events

| Event Type | Description | Required Fields | Optional Fields |
|------------|-------------|----------------|----------------|
| `testcase-queued` | Test case queued | `id`, `name` | `url` |
| `testcase-started` | Test case started | `id`, `name` | `url` |
| `testcase-finished` | Test case completed | `id`, `name` | `url`, `outcome`, `errors` |
| `testcase-skipped` | Test case skipped | `id`, `name` | `url` |
| `testsuite-queued` | Test suite queued | `id`, `name` | `url` |
| `testsuite-started` | Test suite started | `id`, `name` | `url` |
| `testsuite-finished` | Test suite completed | `id`, `name` | `url`, `outcome`, `errors` |
| `testoutput-published` | Test output published | `id`, `name` | `url` |

## Error Handling

The CLI returns appropriate exit codes:

- `0`: Success
- `1`: General error
- `2`: Invalid arguments
- `3`: Network/transport error
- `4`: Configuration error

## Debugging

Enable verbose output for debugging:

```bash
cdevents-cli --verbose generate pipeline started --id "test" --name "test"
```

This will show:
- Configuration loading
- Event creation details
- Transport information
- HTTP request/response details
- Error stack traces
