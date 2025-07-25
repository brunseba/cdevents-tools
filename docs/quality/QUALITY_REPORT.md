# Quality Report

> This report is automatically generated by the quality analysis pipeline.
> Last updated: Sun Jul 13 21:43:35 UTC 2025

# Code Quality Report

## Coverage Analysis
- **Overall Coverage**: 82.3%
- **Threshold**: 70%
- **Status**: ✅ PASS

## Performance Metrics
- **Build Time**: 1s
- **Binary Size**: 17.6M
- **Test Execution Time**: 2s

## Complexity Analysis
- **High Complexity Functions**: 14 functions > 10 complexity
- **Status**: ⚠️ REVIEW NEEDED

## Linting Results
- **Status**: ⚠️ ISSUES FOUND

## Files Generated
- `coverage.out` - Coverage profile
- `coverage.html` - HTML coverage report
- `coverage_detailed.txt` - Detailed coverage report
- `complexity.txt` - High complexity functions
- `complexity_all.txt` - All functions complexity
- `lint.json` - Linting results (JSON)
- `lint.txt` - Linting results (text)
- `cdevents-cli` - Built binary

## Coverage by Package
github.com/brunseba/cdevents-tools/cmd/generate.go:41:			init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate.go:46:			addCommonGenerateFlags		100.0%
github.com/brunseba/cdevents-tools/cmd/generate.go:62:			parseCustomData			83.3%
github.com/brunseba/cdevents-tools/cmd/generate.go:75:			getDefaultSource		66.7%
github.com/brunseba/cdevents-tools/cmd/generate.go:88:			outputEvent			100.0%
github.com/brunseba/cdevents-tools/cmd/generate.go:93:			outputEventWithCustomData	80.0%
github.com/brunseba/cdevents-tools/cmd/generate_build.go:42:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate_pipeline.go:42:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate_service.go:41:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate_task.go:43:		init				100.0%

## How to Generate This Report

This report can be regenerated using:

```bash
# Using Make (recommended)
make quality-docker

# Or using Docker directly
docker build -f Dockerfile.quality -t cdevents-cli-quality .
docker run --rm -v $(pwd)/reports:/app/reports cdevents-cli-quality
```

## Report Files

The quality analysis generates the following files in `reports/`:

- `coverage.out` - Coverage profile for tooling
- `coverage.html` - Interactive HTML coverage report
- `coverage_detailed.txt` - Detailed text coverage report
- `complexity.txt` - Functions with high complexity
- `complexity_all.txt` - Complete complexity analysis
- `lint.json` - Linting results in JSON format
- `lint.txt` - Human-readable linting results
- `quality_report.md` - This quality summary
- `cdevents-cli` - Built binary for testing

## Viewing Reports

### Coverage Report
```bash
# View in browser
open reports/coverage.html

# Or view text summary
cat reports/coverage_detailed.txt
```

### Complexity Analysis
```bash
# View high complexity functions
cat reports/complexity.txt

# View all functions complexity
cat reports/complexity_all.txt
```

### Linting Results
```bash
# View linting issues
cat reports/lint.txt

# View JSON format for tooling
cat reports/lint.json
```
