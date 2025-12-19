# Warp AI Agent Rules - CDEvents CLI

Project-specific rules for AI assistants working on the CDEvents CLI tool repository.

---

## Project Overview

**Repository**: cdevents-tools  
**Purpose**: Command-line tool for generating and sending CDEvents into CI/CD toolchains using CloudEvents as transport  
**Language**: Go 1.21+  
**Stack**: Cobra CLI, Viper config, CDEvents SDK Go v0.4.1, CloudEvents SDK v2.15.2  
**Deployment**: Docker multi-platform, GitHub Actions CI/CD  
**Standards**: CDEvents v0.4.1, CloudEvents v1.0

---

## Development Standards

### Go Language Standards

1. **Go Version**: 1.21+ (as specified in go.mod)
2. **Package Management**: Go modules (go.mod/go.sum)
3. **Code Style**: Follow standard Go conventions
   - Use `gofmt` for formatting
   - Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
   - Use golangci-lint for linting (config in `.golangci.yml`)

4. **Testing Requirements**:
   - Maintain ≥70% test coverage (current: 82.3%)
   - Write table-driven tests for multiple scenarios
   - Include integration tests for CLI commands
   - Test files must end with `_test.go`
   - Run tests: `go test ./...`

5. **Code Quality Metrics**:
   - Maximum cyclomatic complexity: 10 per function
   - No functions longer than 100 lines
   - All exported functions must have doc comments
   - Use meaningful variable names (no single-letter vars except loop counters)

### Project Structure

```
cdevents-tools/
├── cmd/                    # Cobra CLI commands
│   ├── root.go            # Root command definition
│   ├── generate*.go       # Generate event commands
│   ├── send*.go           # Send event commands
│   └── *_test.go          # Command tests
├── pkg/                    # Reusable packages
│   ├── events/            # Event generation logic
│   ├── transport/         # Transport implementations
│   └── config/            # Configuration management
├── scripts/                # Build and quality scripts
├── docs/                   # MkDocs documentation
├── reports/                # Quality analysis reports
├── examples/               # Usage examples
├── main.go                 # Application entry point
├── Dockerfile              # Production Docker image
├── Dockerfile.quality      # Quality analysis container
├── Makefile                # Build and task automation
├── mkdocs.yml              # Documentation configuration
└── docker-compose.yml      # Docker composition
```

---

## Code Organization Rules

### Command Structure (cmd/)

1. **Root Command** (`root.go`):
   - Define global flags and configuration
   - Initialize Viper for config management
   - Set up persistent flags

2. **Subcommands** (one file per major command):
   - `generate*.go` - Event generation commands
   - `send*.go` - Event sending commands
   - Follow Cobra best practices for command organization

3. **Command Naming**:
   - Use kebab-case for command names: `pipeline-started`, `build-finished`
   - Group related commands under parent commands
   - Keep command hierarchy shallow (max 2-3 levels)

4. **Command Tests**:
   - Each command file should have corresponding `*_test.go`
   - Test both success and error cases
   - Mock external dependencies (HTTP, file I/O)

### Package Organization (pkg/)

1. **events/** - Event generation and validation
   - CDEvents creation logic
   - Custom data handling (JSON, YAML, key=value)
   - Event schema validation

2. **transport/** - Transport layer implementations
   - HTTP transport with retry logic
   - File transport
   - Console transport
   - Transport interface for extensibility

3. **config/** - Configuration management
   - Config file loading (~/.cdevents-cli.yaml)
   - Environment variable support
   - Flag precedence handling

### File Naming Conventions

- Go files: `snake_case.go`
- Test files: `*_test.go`
- Documentation: `kebab-case.md`
- Scripts: `snake_case.sh`

---

## Git Workflow

### Conventional Commits (REQUIRED)

Use semantic commit messages following [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

**Types**:
- `feat`: New feature (generates, sends, transport)
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Maintenance (dependencies, build)
- `ci`: CI/CD changes

**Scopes**:
- `cli`: CLI command changes
- `events`: Event generation logic
- `transport`: Transport implementations
- `config`: Configuration handling
- `docker`: Docker-related changes
- `docs`: Documentation
- `quality`: Quality metrics and tooling

**Examples**:
```bash
feat(events): add support for test suite events
fix(transport): retry logic for HTTP timeouts
docs(readme): update installation instructions
test(cli): add integration tests for generate command
chore(deps): update CDEvents SDK to v0.4.2
```

**Co-authorship**:
```
feat(events): add custom annotations support

- Add --annotation flag to commands
- Support multiple annotations per event
- Update documentation

Co-Authored-By: Warp <agent@warp.dev>
```

### Branching Strategy

- `main` - production-ready code
- `develop` - integration branch for features
- `feature/*` - feature branches
- `fix/*` - bug fix branches
- `release/*` - release preparation branches

### Pre-commit Requirements

**Before committing, ensure**:
1. Code is formatted: `gofmt -w .`
2. Linting passes: `golangci-lint run`
3. Tests pass: `go test ./...`
4. Test coverage maintained: `go test -cover ./...`
5. Build succeeds: `go build`

### Quality Gates

Run before pushing:
```bash
# Full quality check
make quality-docker

# Or individual checks
make test
make lint
make coverage
```

---

## Docker Standards

### Production Dockerfile

**Location**: `Dockerfile`

**Standards**:
- Multi-stage builds for minimal image size
- Use official Go Alpine base image
- Non-root user execution
- Layer caching optimization
- Multi-platform support (linux/amd64, linux/arm64)

**Build**:
```bash
# Local build
docker build -t cdevents-cli:latest .

# Multi-platform build
docker buildx build --platform linux/amd64,linux/arm64 -t cdevents-cli:latest .
```

### Quality Dockerfile

**Location**: `Dockerfile.quality`

**Purpose**: Reproducible quality analysis environment

**Tools Installed**:
- gocyclo - Cyclomatic complexity analysis
- golangci-lint - Comprehensive linting
- go coverage tools
- Custom quality scripts

**Usage**:
```bash
make quality-docker
```

### Docker Compose

**Location**: `docker-compose.yml`

**Services**:
- `cdevents-cli` - Main CLI service
- Volume mounts for local development

---

## Testing Standards

### Test Coverage Requirements

**Targets**:
- Overall coverage: ≥70% (current: 82.3%)
- New code: ≥80% coverage
- Critical paths: 100% coverage

**Coverage Commands**:
```bash
# Run tests with coverage
go test -coverprofile=coverage.out ./...

# View coverage in browser
go tool cover -html=coverage.out

# Coverage summary
go tool cover -func=coverage.out
```

### Test Organization

**Table-Driven Tests** (preferred):
```go
func TestGeneratePipelineEvent(t *testing.T) {
    tests := []struct {
        name        string
        id          string
        eventName   string
        wantErr     bool
        errContains string
    }{
        {
            name:      "valid pipeline started event",
            id:        "pipeline-123",
            eventName: "my-pipeline",
            wantErr:   false,
        },
        {
            name:        "missing id",
            id:          "",
            eventName:   "my-pipeline",
            wantErr:     true,
            errContains: "id is required",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

**Integration Tests**:
- Location: `cmd/integration_test.go`
- Test full CLI command execution
- Use temporary directories for file operations
- Mock external HTTP endpoints

**Test Helpers**:
- Create reusable test utilities in `*_test.go` files
- Use `testing.T.Helper()` for helper functions
- Provide meaningful error messages

---

## Documentation Standards

### MkDocs Documentation

**Location**: `docs/`

**Structure**:
```
docs/
├── index.md                    # Home page
├── getting-started.md          # Installation and quick start
├── cli-reference.md            # Complete CLI command reference
├── examples.md                 # Usage examples
├── docker.md                   # Docker deployment
├── contributing.md             # Contributing guidelines
├── development-standards.md    # Development practices
├── quality-metrics.md          # Quality analysis
└── quality/                    # Quality reports
    ├── QUALITY_REPORT.md
    ├── coverage-summary.md
    ├── complexity-report.md
    ├── linting-report.md
    └── performance-metrics.md
```

**Configuration**: `mkdocs.yml`
- Theme: Material
- Features: Search, navigation, code highlighting
- Extensions: Mermaid diagrams support

**Building Documentation**:
```bash
# Install dependencies
pip install -r requirements.txt

# Serve locally
mkdocs serve

# Build static site
mkdocs build
```

### Code Documentation

**Function Comments** (required for exported functions):
```go
// GeneratePipelineEvent creates a new CDEvent for pipeline execution.
// It validates the input parameters and returns a CloudEvent with CDEvent payload.
//
// Parameters:
//   - id: Unique identifier for the pipeline
//   - name: Human-readable pipeline name
//   - eventType: Type of pipeline event (started, finished, queued)
//
// Returns:
//   - *cloudevents.Event: Generated CloudEvent
//   - error: Error if validation fails
func GeneratePipelineEvent(id, name, eventType string) (*cloudevents.Event, error) {
    // Implementation
}
```

**Package Comments** (required for each package):
```go
// Package events provides functionality for generating CDEvents.
// It supports pipeline, task, build, service, and test events.
package events
```

---

## Quality Standards & Metrics

### Quality Metrics Tracking

**Location**: `reports/` directory

**Generated Reports**:
1. `quality_report.md` - Comprehensive quality summary
2. `coverage.html` - Interactive HTML coverage report
3. `coverage.out` - Raw coverage data

**Quality Dashboard** (in docs):
- Coverage trends
- Complexity analysis
- Linting results
- Performance metrics

### Quality Gates (Automated)

**Coverage Gate**: Fails if < 70%
**Complexity Gate**: Warns if function complexity > 10
**Linting Gate**: Must pass golangci-lint
**Performance Gate**: Build time < 5s, test time < 10s

### Running Quality Analysis

**Docker-based (Recommended)**:
```bash
make quality-docker
```

**Manual**:
```bash
# Run tests with coverage
go test -coverprofile=coverage.out ./...

# Run linting
golangci-lint run

# Check complexity
gocyclo -over 10 .

# Generate reports
./scripts/run_quality_metrics.sh
```

---

## Build & Release

### Makefile Targets

**Common Targets**:
```bash
make build          # Build binary
make test           # Run tests
make coverage       # Generate coverage report
make lint           # Run linters
make docker-build   # Build Docker image
make quality-docker # Run quality analysis
make clean          # Clean build artifacts
```

### Versioning

**Version File**: `version.txt`

**Semantic Versioning**: MAJOR.MINOR.PATCH
- MAJOR: Breaking changes
- MINOR: New features (backward compatible)
- PATCH: Bug fixes

**Release Process**:
1. Update `version.txt`
2. Update `CHANGELOG.md`
3. Create git tag: `git tag -a vX.Y.Z -m "Release vX.Y.Z"`
4. Push tag: `git push origin vX.Y.Z`
5. GitHub Actions builds and publishes Docker image

---

## CI/CD Integration

### GitHub Actions

**Location**: `.github/workflows/`

**Workflows**:
- `ci.yml` - Build, test, lint on PR
- `release.yml` - Build and publish on tag
- `quality.yml` - Quality metrics reporting

**Required Checks**:
- All tests pass
- Linting clean
- Coverage ≥70%
- Build succeeds for all platforms

---

## Dependencies Management

### External Dependencies

**Core Dependencies**:
- `github.com/cdevents/sdk-go` - CDEvents SDK
- `github.com/cloudevents/sdk-go/v2` - CloudEvents SDK
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management

**Update Strategy**:
- Review security advisories monthly
- Update patch versions immediately
- Test thoroughly before updating minor/major versions
- Document breaking changes in CHANGELOG.md

**Dependency Commands**:
```bash
# Update dependencies
go get -u ./...
go mod tidy

# Verify dependencies
go mod verify

# View dependency tree
go mod graph
```

---

## AI Assistant Behavior Guidelines

### When Working on This Project

1. **Code Standards**:
   - Always format Go code with `gofmt`
   - Follow Go naming conventions (camelCase for unexported, PascalCase for exported)
   - Add doc comments for all exported functions
   - Keep functions under 100 lines
   - Maintain cyclomatic complexity < 10

2. **Testing Requirements**:
   - Write tests for all new functions
   - Use table-driven tests for multiple scenarios
   - Maintain ≥70% test coverage
   - Test both success and error paths
   - Mock external dependencies

3. **Command Implementation**:
   - Use Cobra best practices
   - Add both short and long descriptions
   - Include usage examples in help text
   - Validate input parameters
   - Return meaningful error messages

4. **Documentation Updates**:
   - Update CLI reference when adding commands
   - Add examples for new features
   - Update README.md for significant changes
   - Keep CHANGELOG.md current

5. **Quality Checks**:
   - Run `make quality-docker` before committing
   - Fix all linting warnings
   - Ensure tests pass
   - Check coverage hasn't dropped

### What to Avoid

- ❌ Single-letter variable names (except loop counters)
- ❌ Functions longer than 100 lines
- ❌ Unexported functions without doc comments
- ❌ Hard-coded values (use constants or config)
- ❌ Ignoring errors (`_ = someFunc()`)
- ❌ Committing without running tests
- ❌ Breaking existing CLI commands (backward compatibility)
- ❌ Adding dependencies without justification
- ❌ Reducing test coverage
- ❌ Skipping quality analysis

### When Adding New Features

1. **Design Phase**:
   - Review existing commands and patterns
   - Check CDEvents specification compliance
   - Consider backward compatibility
   - Plan test coverage

2. **Implementation Phase**:
   - Create command in appropriate `cmd/*` file
   - Implement logic in `pkg/` packages
   - Add table-driven tests
   - Update documentation

3. **Quality Phase**:
   - Run `make test`
   - Run `make lint`
   - Check coverage: `go test -cover ./...`
   - Run `make quality-docker`

4. **Documentation Phase**:
   - Update CLI reference
   - Add usage examples
   - Update README.md if needed
   - Add entry to CHANGELOG.md

5. **Review Phase**:
   - Test CLI commands manually
   - Verify Docker build works
   - Check all quality gates pass
   - Commit with conventional commit message

---

## CDEvents & CloudEvents Standards

### CDEvents Compliance

**Version**: v0.4.1

**Supported Event Types**:
- Pipeline: queued, started, finished
- Task: started, finished
- Build: queued, started, finished
- Service: deployed, published, removed, rolledback, upgraded
- Test: testcase-*, testsuite-*, testoutput-published

**Event Structure**:
- Must include context fields (type, source, id)
- Must follow CDEvents subject schema
- Support custom data, labels, annotations
- Follow CloudEvents v1.0 specification

### CloudEvents Integration

**Transport Encoding**: Binary mode (preferred)
**Content Type**: application/cloudevents+json
**Required Attributes**: id, source, type, specversion

---

## Error Handling

### Error Patterns

**Command Errors**:
```go
// Use cobra.Command.SilenceUsage to prevent usage on validation errors
cmd.SilenceUsage = true

// Return descriptive errors
return fmt.Errorf("failed to generate event: %w", err)
```

**Validation Errors**:
```go
// Check required parameters
if id == "" {
    return fmt.Errorf("id is required")
}

// Validate event types
validTypes := []string{"started", "finished", "queued"}
if !contains(validTypes, eventType) {
    return fmt.Errorf("invalid event type %q, must be one of: %v", eventType, validTypes)
}
```

**HTTP Transport Errors**:
```go
// Implement retry logic with exponential backoff
// Log retry attempts
// Return final error after max retries
```

---

## Configuration Management

### Configuration Files

**Location**: `~/.cdevents-cli.yaml`

**Example**:
```yaml
source: "my-ci-system"
target: "http://events.example.com"
output: "json"
retries: 3
timeout: 30s
```

### Environment Variables

**Prefix**: `CDEVENTS_`

**Examples**:
- `CDEVENTS_SOURCE`
- `CDEVENTS_TARGET`
- `CDEVENTS_OUTPUT`

### Precedence Order

1. Command-line flags (highest priority)
2. Environment variables
3. Configuration file
4. Default values (lowest priority)

---

## Deployment

### Docker Deployment

**Production Image**:
```bash
# Build
docker build -t cdevents-cli:v1.0.0 .

# Run
docker run --rm cdevents-cli:v1.0.0 generate pipeline started --id test-123
```

**Docker Compose**:
```bash
docker-compose up
```

### Binary Distribution

**Build for Multiple Platforms**:
```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o cdevents-cli-linux-amd64

# macOS
GOOS=darwin GOARCH=amd64 go build -o cdevents-cli-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o cdevents-cli-darwin-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o cdevents-cli-windows-amd64.exe
```

---

## Community & Support

**CDEvents Community**:
- Slack: [CDEvents Slack Channel](https://cdeliveryfdn.slack.com/archives/C030SKZ0F4K)
- GitHub: [CDEvents Organization](https://github.com/cdevents)
- Docs: [CDEvents Documentation](https://cdevents.dev/docs/)

**Related Projects**:
- [CDEvents Specification](https://github.com/cdevents/spec)
- [CDEvents Go SDK](https://github.com/cdevents/sdk-go)
- [CloudEvents](https://cloudevents.io/)

---

## Key Repository Information

**URLs**:
- Repository: https://github.com/brunseba/cdevents-tools
- Issues: https://github.com/brunseba/cdevents-tools/issues
- Docker Hub: https://hub.docker.com/r/cdevents/cdevents-cli

**Technology Stack**:
- Go: 1.21+
- Cobra: CLI framework
- Viper: Configuration
- CDEvents SDK: v0.4.1
- CloudEvents SDK: v2.15.2

**Quality Targets**:
- Test Coverage: ≥70%
- Build Time: <5s
- Test Execution: <10s
- Cyclomatic Complexity: <10 per function

---

**Last Updated**: December 2025  
**Version**: 1.0
