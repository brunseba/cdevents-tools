# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v1.0.0] - 2025-07-13

### Added
- Initial release of CDEvents CLI tool
- Support for generating CDEvents for pipeline, build, task, service, and test events
- Multiple output formats: JSON, YAML, and CloudEvents
- Multiple transport options: HTTP, console, and file
- Custom data support via `--custom-json` flag
- CloudEvents v1.0 compatibility with binary encoding
- Comprehensive test coverage (82.3%)
- Schema validation against CDEvents v0.4.1 specification
- Docker support with multi-platform binaries
- Quality metrics with coverage, complexity, and performance monitoring
- Extensive documentation with CLI reference and usage examples
- Make targets for easy development and testing
- Support for retry logic and configurable timeouts
- HTTP headers support for secure transport
- Comprehensive input/output examples

### Features
- **Event Generation**: Pipeline, build, task, service, and test events
- **Output Formats**: JSON, YAML, CloudEvents
- **Transport Methods**: HTTP, console, file
- **Custom Data**: JSON format with data, labels, annotations, and links
- **CloudEvents Integration**: Full v1.0 specification support
- **Schema Validation**: CDEvents v0.4.1 compliance
- **Quality Assurance**: 82.3% test coverage, complexity monitoring
- **Documentation**: Comprehensive CLI reference and examples
- **Docker Support**: Multi-platform containerized deployment
- **CI/CD Integration**: Jenkins, GitHub Actions, GitLab CI examples

### Technical Details
- Go 1.21 support
- CDEvents SDK v0.4.1 integration
- CloudEvents SDK v2.15.2 integration
- Cobra CLI framework
- Viper configuration management
- Comprehensive error handling and logging
- Retry mechanisms with exponential backoff
- HTTP transport with proper headers
- File transport with atomic writes
- Console transport with colored output

### Quality Metrics
- **Code Coverage**: 82.3% overall
- **Build Time**: 2.116s
- **Test Execution**: 7.7s
- **Cyclomatic Complexity**: 5 functions >10 complexity
- **Docker-based Quality Analysis**: Reproducible metrics
- **Linting**: golangci-lint with comprehensive rules
- **Schema Validation**: gojsonschema for CDEvents compliance

### Documentation
- Complete CLI reference with examples
- Getting started guide
- Input/output examples for all event types
- Docker deployment guide
- Quality metrics documentation
- Contributing guidelines
- Development standards

[v1.0.0]: https://github.com/brunseba/cdevents-tools/releases/tag/v1.0.0
