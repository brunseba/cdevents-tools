# Contributing

We welcome contributions to the CDEvents CLI project! This guide will help you get started.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker and Docker Compose
- Git
- Make (optional, for build automation)

### Development Setup

1. **Fork the repository**

   Fork the project on GitHub and clone your fork:

   ```bash
   git clone https://github.com/YOUR-USERNAME/cdevents-cli.git
   cd cdevents-cli
   ```

2. **Set up the upstream remote**

   ```bash
   git remote add upstream https://github.com/cdevents/cdevents-cli.git
   ```

3. **Install dependencies**

   ```bash
   go mod tidy
   ```

4. **Build the project**

   ```bash
   go build -o cdevents-cli
   ```

5. **Run tests**

   ```bash
   go test ./...
   ```

## Making Changes

### Code Style

- Follow standard Go conventions
- Use `gofmt` to format your code
- Run `go vet` to check for common errors
- Add comments for public functions and types
- Write tests for new functionality

### Commit Messages

Write clear, descriptive commit messages:

```
feat: add support for Kafka transport

- Implement KafkaTransport struct
- Add configuration options for brokers and topics
- Include retry logic for failed deliveries
- Add unit tests for Kafka functionality

Closes #123
```

Use conventional commit format:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for adding tests
- `refactor:` for code refactoring
- `chore:` for maintenance tasks

### Testing

#### Unit Tests

Write unit tests for all new code:

```bash
go test ./pkg/...
```

#### Integration Tests

Test with real services where possible:

```bash
# Start test services
docker-compose up -d httpbin

# Run integration tests
go test -tags=integration ./...
```

#### Manual Testing

Test the CLI manually:

```bash
# Test generation
./cdevents-cli generate pipeline started --id "test" --name "test"

# Test sending
./cdevents-cli send --target http://localhost:8080/post pipeline started --id "test" --name "test"
```

### Documentation

Update documentation when making changes:

1. **Code Documentation**: Add or update Go doc comments
2. **CLI Documentation**: Update command help text
3. **User Documentation**: Update markdown files in the `docs/` directory
4. **README**: Update the main README if needed

Generate documentation:

```bash
# Generate CLI documentation
./cdevents-cli --help > docs/cli-help.txt

# Serve documentation locally
mkdocs serve
```

## Pull Request Process

1. **Create a feature branch**

   ```bash
   git checkout -b feature/my-new-feature
   ```

2. **Make your changes**

   - Write code
   - Add tests
   - Update documentation
   - Ensure tests pass

3. **Test thoroughly**

   ```bash
   # Run all tests
   go test ./...
   
   # Test with Docker
   docker-compose build
   docker run --rm cdevents-cli:latest --help
   
   # Test documentation
   mkdocs serve
   ```

4. **Commit and push**

   ```bash
   git add .
   git commit -m "feat: add my new feature"
   git push origin feature/my-new-feature
   ```

5. **Create a pull request**

   - Go to GitHub and create a PR
   - Fill out the PR template
   - Link to any relevant issues
   - Request reviews from maintainers

### Pull Request Checklist

- [ ] Code follows Go conventions
- [ ] Tests are included and passing
- [ ] Documentation is updated
- [ ] Commit messages follow conventional format
- [ ] PR description is clear and complete
- [ ] All CI checks pass

## Types of Contributions

### Bug Reports

When reporting bugs, please include:

- Steps to reproduce the issue
- Expected vs actual behavior
- CDEvents CLI version
- Operating system and Go version
- Any relevant error messages or logs

### Feature Requests

For new features:

- Explain the use case
- Describe the proposed solution
- Consider alternatives
- Discuss impact on existing functionality

### Code Contributions

Areas where we welcome contributions:

- New transport implementations (Kafka, NATS, etc.)
- Additional event types
- Performance improvements
- Bug fixes
- Documentation improvements
- Test coverage improvements

### Documentation

Help improve our documentation:

- Fix typos and grammar
- Add examples
- Improve clarity
- Add translations
- Update outdated information

## Development Standards

This project follows a set of standards to ensure high code quality and consistency:

### Coding Standards
- **Language**: Go (Golang)
- **Version**: 1.21 or later
- **Practices**: Follow best practices including idiomatic Go conventions
- **Formatting**: Use `gofmt` for code formatting
- **Linting**: Use `go vet` for static code analysis

### Commit Standards
- **Format**: Conventional Commits (e.g., `feat`, `fix`, `docs`, `test`)
- **Example**: `feat: add new event type`
- **Structure**: Subject line followed by details and closing issues if any

### Testing Standards
- **Unit Tests**: Required for all functional changes
- **Coverage**: Aim for above 80% coverage
- **Tools**: Use built-in Go testing framework (`testing`) and coverage tools
- **Formats**: Test outputs should be in standard formats (e.g., JSON)

### Documentation Standards
- **Style**: Follow Markdown format
- **CLI Docs**: Update command help and examples
- **Project Docs**: Update user and contributor guides

### Review Standards
- **Reviews**: All changes must be peer-reviewed
- **Approvals**: At least one approval by another maintainer
- **Checks**: Ensure all CI checks are passing before merging

---

### Project Structure

```
cdevents-cli/
├── cmd/                    # CLI commands
├── pkg/
│   ├── events/            # Event factory and types
│   ├── transport/         # Transport implementations
│   └── output/            # Output formatters
├── docs/                  # Documentation
├── examples/              # Example configurations
├── Dockerfile            # Container image
├── docker-compose.yml    # Development environment
└── mkdocs.yml           # Documentation configuration
```

### Adding New Commands

1. Create a new file in `cmd/` directory
2. Implement the command using Cobra
3. Add the command to the parent command
4. Add tests in `cmd/` directory
5. Update documentation

### Adding New Event Types

1. Add event creation logic to `pkg/events/factory.go`
2. Add command support in `cmd/generate_*.go`
3. Add tests for the new event type
4. Update documentation

### Adding New Transports

1. Implement the `Transport` interface in `pkg/transport/`
2. Add the transport to `TransportFactory`
3. Add configuration options
4. Add tests
5. Update documentation

## Testing

### Running Tests

```bash
# Unit tests
go test ./...

# With coverage
go test -cover ./...

# Integration tests
go test -tags=integration ./...

# Race detection
go test -race ./...
```

### Writing Tests

Example test structure:

```go
func TestPipelineEventGeneration(t *testing.T) {
    factory := events.NewEventFactory("test-source")
    
    event, err := factory.CreatePipelineRunEvent(
        "started",
        "pipeline-123",
        "my-pipeline",
        "",
        "",
        "",
    )
    
    assert.NoError(t, err)
    assert.NotNil(t, event)
    assert.Equal(t, "pipeline-123", event.GetSubjectId())
}
```

## Community

### Communication

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For general questions and discussions
- **Slack**: Join the [CDEvents Slack](https://cdeliveryfdn.slack.com/archives/C030SKZ0F4K)

### Code of Conduct

This project follows the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).

## Recognition

Contributors are recognized in:

- The project's README
- Release notes
- Community calls and presentations

## Getting Help

If you need help:

1. Check the documentation
2. Search existing issues
3. Ask in GitHub Discussions
4. Join the Slack channel
5. Create a new issue

## License

By contributing to CDEvents CLI, you agree that your contributions will be licensed under the Apache License 2.0.

## Maintainers

Current maintainers:

- [List of maintainers to be added]

To become a maintainer:

1. Contribute regularly to the project
2. Demonstrate good judgment in code reviews
3. Help with community support
4. Be nominated by existing maintainers
