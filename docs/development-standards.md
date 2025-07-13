# Development Standards

This document outlines the development standards and practices used in the CDEvents CLI project. These standards ensure consistency, maintainability, and quality across the codebase.

## Code Quality Standards

### Language and Environment
- **Language**: Go (Golang) 1.21 or later
- **Build System**: Go modules with semantic versioning
- **Runtime**: Cross-platform support (Linux, macOS, Windows)
- **Containerization**: Docker with multi-architecture support

### Code Style and Formatting
- **Formatting**: Use `gofmt` for consistent code formatting
- **Linting**: Use `go vet` for static code analysis
- **Conventions**: Follow idiomatic Go conventions and best practices
- **Documentation**: Use Go doc comments for all public APIs
- **Error Handling**: Proper error handling with descriptive messages

### Architecture Standards
- **Structure**: Clean architecture with separation of concerns
- **Packages**: Organize code in logical packages (`cmd`, `pkg/events`, `pkg/transport`, `pkg/output`)
- **Interfaces**: Use interfaces for abstraction and testability
- **Dependency Injection**: Minimize dependencies and use dependency injection where appropriate

## Testing Standards

### Coverage Requirements
- **Minimum Coverage**: 80% overall code coverage
- **Package Coverage**: 
  - `cmd`: 87.0% (achieved)
  - `pkg/events`: 78.1% (achieved)
  - `pkg/output`: 78.7% (achieved)
  - `pkg/transport`: 90.9% (achieved)

### Testing Framework
- **Unit Tests**: Use Go's built-in `testing` package
- **Test Structure**: Follow table-driven tests where appropriate
- **Mocking**: Create mock implementations for external dependencies
- **Integration Tests**: Test CLI commands end-to-end

### Test Categories
1. **Unit Tests**: Test individual functions and methods
2. **Integration Tests**: Test command-line interface and workflows
3. **Mock Tests**: Test with mock transports and external services
4. **Edge Case Tests**: Test error conditions and boundary cases

## Documentation Standards

### User Documentation
- **Format**: Markdown with MkDocs Material theme
- **Structure**: Organized in logical sections (Getting Started, CLI Reference, Examples)
- **Examples**: Provide practical examples for all features
- **Accuracy**: Keep documentation synchronized with code changes

### Code Documentation
- **Go Doc**: Document all public functions, types, and packages
- **Comments**: Write clear, concise comments for complex logic
- **CLI Help**: Provide comprehensive help text for all commands
- **Error Messages**: Use descriptive error messages with suggested solutions

## Version Control Standards

### Commit Message Format
Follow [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

#### Types
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `chore`: Maintenance tasks
- `ci`: CI/CD changes

#### Examples
```
feat: add support for Kafka transport
fix: resolve custom data parsing issue
docs: update CLI reference documentation
test: enforce code coverage for cmd package
```

### Branch Strategy
- **Main Branch**: `main` - stable, production-ready code
- **Feature Branches**: `feature/description` - new features
- **Bug Fix Branches**: `fix/description` - bug fixes
- **Documentation Branches**: `docs/description` - documentation updates

## Build and Deployment Standards

### Build Process
- **Multi-Stage Builds**: Use Docker multi-stage builds for efficient images
- **Multi-Architecture**: Support `linux/amd64` and `linux/arm64`
- **Static Binaries**: Build statically linked binaries for easy deployment
- **Version Tagging**: Use semantic versioning for releases

### Dependency Management
- **Go Modules**: Use Go modules for dependency management
- **Pinned Versions**: Pin specific versions for reproducible builds
- **Minimal Dependencies**: Minimize external dependencies
- **Security**: Regularly update dependencies for security patches

## Quality Assurance Standards

### Code Review Process
- **Peer Review**: All code changes require peer review
- **Review Checklist**: Use standardized review checklist
- **Automated Checks**: CI/CD pipeline runs automated checks
- **Approval**: At least one approval from a maintainer

### Continuous Integration
- **Automated Testing**: Run full test suite on every commit
- **Code Coverage**: Enforce minimum code coverage requirements
- **Static Analysis**: Run linting and security checks
- **Build Verification**: Verify builds across supported platforms

## Security Standards

### Code Security
- **Input Validation**: Validate all user inputs
- **Error Handling**: Avoid exposing sensitive information in errors
- **Dependencies**: Use tools like `govulncheck` for vulnerability scanning
- **Secrets**: Never commit secrets or sensitive data

### Container Security
- **Base Images**: Use minimal, secure base images
- **Non-Root**: Run containers as non-root users
- **Scanning**: Scan images for vulnerabilities
- **Minimal Attack Surface**: Include only necessary components

## Performance Standards

### Efficiency
- **Resource Usage**: Optimize for minimal memory and CPU usage
- **Startup Time**: Fast CLI startup and command execution
- **Concurrent Processing**: Use goroutines for concurrent operations where appropriate
- **Caching**: Implement caching for repeated operations

### Scalability
- **Event Processing**: Handle large volumes of events efficiently
- **Transport Layer**: Support multiple concurrent transport connections
- **Memory Management**: Proper memory management to prevent leaks

## Compliance Standards

### CDEvents Specification
- **Spec Compliance**: Follow CDEvents specification v0.4.1
- **Event Structure**: Ensure generated events conform to specification
- **Backwards Compatibility**: Maintain compatibility with previous versions
- **Validation**: Validate events against CDEvents schema

### CloudEvents Standard
- **CloudEvents v1.0**: Full compatibility with CloudEvents specification v1.0
- **Transport Format**: Events are transportable as CloudEvents using `api.AsCloudEvent()`
- **Binary Encoding**: Support for CloudEvents binary content mode over HTTP
- **Event Mapping**: Proper mapping between CDEvents and CloudEvents attributes:
  - CDEvents `id` → CloudEvents `id`
  - CDEvents `source` → CloudEvents `source`
  - CDEvents `type` → CloudEvents `type`
  - CDEvents `subject.id` → CloudEvents `subject`
  - CDEvents `timestamp` → CloudEvents `time`
  - CDEvents event data → CloudEvents `data` (JSON format)
- **Custom Data**: CloudEvents data field includes full CDEvents with custom data
- **Content Type**: Uses `application/json` for CloudEvents `datacontenttype`

### Open Source Standards
- **License**: Apache License 2.0
- **Contributing**: Clear contribution guidelines
- **Code of Conduct**: CNCF Code of Conduct
- **Community**: Active community engagement

## Tools and Automation

### Development Tools
- **IDE**: VS Code with Go extension (recommended)
- **Debugging**: Use Go debugger (delve)
- **Profiling**: Use Go profiling tools for performance analysis
- **Documentation**: MkDocs for documentation generation

### CI/CD Tools
- **Version Control**: Git with conventional commits
- **Container Registry**: Docker Hub for image distribution
- **Testing**: Go testing framework with coverage reporting
- **Documentation**: Automated documentation updates

## Monitoring and Metrics

### Code Quality Metrics
- **Code Coverage**: Track and maintain high code coverage
- **Cyclomatic Complexity**: Monitor code complexity
- **Technical Debt**: Regular refactoring to reduce technical debt
- **Performance**: Monitor build times and execution performance

### Usage Metrics
- **CLI Usage**: Track command usage patterns
- **Error Rates**: Monitor error rates and patterns
- **Performance**: Track execution times and resource usage
- **Adoption**: Monitor project adoption and growth

## Maintenance Standards

### Regular Tasks
- **Dependency Updates**: Regular dependency updates
- **Security Patches**: Timely application of security patches
- **Documentation Updates**: Keep documentation current
- **Performance Monitoring**: Regular performance reviews

### Release Process
- **Versioning**: Semantic versioning for releases
- **Release Notes**: Comprehensive release notes
- **Testing**: Thorough testing before releases
- **Rollback**: Rollback procedures for problematic releases

## Resources

### Documentation
- [CDEvents Specification](https://github.com/cdevents/spec)
- [Go Documentation](https://golang.org/doc/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [MkDocs Material](https://squidfunk.github.io/mkdocs-material/)

### Tools
- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [MkDocs](https://www.mkdocs.org/)
- [GitHub Actions](https://github.com/features/actions)

---

*These standards are living documents and may be updated as the project evolves. All contributors are expected to follow these standards to maintain code quality and consistency.*
