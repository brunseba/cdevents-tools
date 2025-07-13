# GitHub Actions Workflows

This directory contains GitHub Actions workflows for the CDEvents CLI project.

## Workflows

### 📚 Documentation (`docs.yml`)

Automatically builds and deploys MkDocs documentation to GitHub Pages.

**Triggers:**
- Push to `main` branch
- Pull requests to `main` branch
- Manual workflow dispatch

**Features:**
- Builds MkDocs documentation with Material theme
- Generates quality reports automatically
- Deploys to GitHub Pages
- Updates quality metrics in real-time

**Access:** [Project Documentation](https://brunseba.github.io/cdevents-tools/)

### 🧪 Continuous Integration (`ci.yml`)

Runs quality checks, tests, and builds on every push and pull request.

**Triggers:**
- Push to `main` branch
- Pull requests to `main` branch

**Features:**
- Runs Go tests with coverage reporting
- Enforces 70% minimum code coverage
- Runs golangci-lint for code quality
- Checks cyclomatic complexity
- Builds the project
- Caches Go modules for faster builds

**Quality Gates:**
- ✅ Code coverage ≥ 70%
- ✅ No linting errors
- ✅ Build success
- ⚠️ Complexity monitoring

## Setup Requirements

### GitHub Pages Setup

1. Go to repository **Settings** → **Pages**
2. Set **Source** to "GitHub Actions"
3. The documentation will be available at: `https://brunseba.github.io/cdevents-tools/`

### Environment Variables

No special environment variables are required. The workflows use:

- `GITHUB_TOKEN` (automatically provided)
- Go modules cache
- Python pip cache

## Local Development

To run the same quality checks locally:

```bash
# Install dependencies
go mod download

# Run tests with coverage
go test ./... -coverprofile=coverage.out -covermode=atomic

# Check coverage
go tool cover -func=coverage.out

# Run linting
golangci-lint run

# Build documentation
mkdocs serve
```

## Quality Reports

The documentation workflow automatically generates:

- 📊 **Coverage Report**: Interactive HTML coverage analysis
- 📈 **Quality Metrics**: Code quality and performance metrics
- 🔍 **Complexity Analysis**: Function complexity analysis
- 🔧 **Linting Report**: Code quality issues and suggestions

All reports are automatically updated and published to the documentation site.

## Contributing

When contributing:

1. Ensure all CI checks pass
2. Maintain code coverage above 70%
3. Address any linting issues
4. Keep function complexity reasonable
5. Update documentation as needed

The workflows will automatically check your contributions and provide feedback through GitHub status checks.
