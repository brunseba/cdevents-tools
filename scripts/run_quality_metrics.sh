#!/bin/bash

set -e
# Continue on test failures but track them
set +e

echo "======================================"
echo "🔍 CDEvents CLI - Code Quality Metrics"
echo "======================================"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Create reports directory
mkdir -p reports

echo -e "${BLUE}📊 Step 1: Running Tests with Coverage${NC}"
echo "--------------------------------------"

# Run tests with coverage
echo "Running all tests with coverage analysis..."
time go test ./... -coverprofile=reports/coverage.out -covermode=atomic -v
TEST_EXIT_CODE=$?

if [ -f "reports/coverage.out" ]; then
    echo -e "${GREEN}✅ Coverage report generated successfully${NC}"
else
    echo -e "${RED}❌ Failed to generate coverage report${NC}"
    # Create empty coverage file to continue
    touch reports/coverage.out
fi

if [ $TEST_EXIT_CODE -ne 0 ]; then
    echo -e "${YELLOW}⚠️  Some tests failed, but continuing with analysis${NC}"
fi

echo ""
echo -e "${BLUE}📈 Step 2: Code Coverage Analysis${NC}"
echo "--------------------------------------"

# Generate coverage summary
echo "Coverage Summary:"
go tool cover -func=reports/coverage.out | tail -1

# Generate detailed coverage report
echo "Generating detailed coverage report..."
go tool cover -func=reports/coverage.out > reports/coverage_detailed.txt

# Generate HTML coverage report
echo "Generating HTML coverage report..."
go tool cover -html=reports/coverage.out -o reports/coverage.html

# Extract coverage percentage
COVERAGE=$(go tool cover -func=reports/coverage.out | tail -1 | awk '{print $3}' | sed 's/%//')
echo "Overall Coverage: ${COVERAGE}%"

# Coverage threshold check
THRESHOLD=70
if (( $(echo "$COVERAGE >= $THRESHOLD" | bc -l) )); then
    echo -e "${GREEN}✅ Coverage meets threshold (${THRESHOLD}%)${NC}"
else
    echo -e "${YELLOW}⚠️  Coverage below threshold (${THRESHOLD}%)${NC}"
fi

echo ""
echo -e "${BLUE}🔄 Step 3: Cyclomatic Complexity Analysis${NC}"
echo "--------------------------------------"

# Check cyclomatic complexity
echo "Analyzing cyclomatic complexity..."
gocyclo -over 10 . > reports/complexity.txt || true

if [ -s "reports/complexity.txt" ]; then
    echo -e "${YELLOW}⚠️  Functions with high complexity (>10):${NC}"
    cat reports/complexity.txt
else
    echo -e "${GREEN}✅ All functions have acceptable complexity (<=10)${NC}"
fi

# Generate complexity report for all functions
echo "Generating complete complexity report..."
gocyclo . > reports/complexity_all.txt

echo ""
echo -e "${BLUE}⏱️  Step 4: Performance Metrics${NC}"
echo "--------------------------------------"

# Measure build time
echo "Measuring build time..."
BUILD_START=$(date +%s.%N)
go build -o reports/cdevents-cli .
BUILD_END=$(date +%s.%N)
BUILD_TIME=$(echo "$BUILD_END - $BUILD_START" | bc)
echo "Build time: ${BUILD_TIME}s"

# Measure binary size
BINARY_SIZE=$(du -h reports/cdevents-cli | cut -f1)
echo "Binary size: ${BINARY_SIZE}"

# Test execution time
echo "Measuring test execution time..."
TEST_START=$(date +%s.%N)
go test ./... -short > /dev/null 2>&1
TEST_END=$(date +%s.%N)
TEST_TIME=$(echo "$TEST_END - $TEST_START" | bc)
echo "Test execution time: ${TEST_TIME}s"

echo ""
echo -e "${BLUE}📋 Step 5: Code Quality Linting${NC}"
echo "--------------------------------------"

# Run golangci-lint
echo "Running golangci-lint..."
golangci-lint run --out-format=json > reports/lint.json || true
golangci-lint run > reports/lint.txt || true

if [ -s "reports/lint.txt" ]; then
    echo -e "${YELLOW}⚠️  Linting issues found:${NC}"
    head -20 reports/lint.txt
else
    echo -e "${GREEN}✅ No linting issues found${NC}"
fi

echo ""
echo -e "${BLUE}📊 Step 6: Generating Quality Report${NC}"
echo "--------------------------------------"

# Generate comprehensive quality report
cat > reports/quality_report.md << EOF
# Code Quality Report

## Coverage Analysis
- **Overall Coverage**: ${COVERAGE}%
- **Threshold**: ${THRESHOLD}%
- **Status**: $(if (( $(echo "$COVERAGE >= $THRESHOLD" | bc -l) )); then echo "✅ PASS"; else echo "⚠️ BELOW THRESHOLD"; fi)

## Performance Metrics
- **Build Time**: ${BUILD_TIME}s
- **Binary Size**: ${BINARY_SIZE}
- **Test Execution Time**: ${TEST_TIME}s

## Complexity Analysis
- **High Complexity Functions**: $(wc -l < reports/complexity.txt) functions > 10 complexity
- **Status**: $(if [ -s "reports/complexity.txt" ]; then echo "⚠️ REVIEW NEEDED"; else echo "✅ GOOD"; fi)

## Linting Results
- **Status**: $(if [ -s "reports/lint.txt" ]; then echo "⚠️ ISSUES FOUND"; else echo "✅ CLEAN"; fi)

## Files Generated
- \`coverage.out\` - Coverage profile
- \`coverage.html\` - HTML coverage report
- \`coverage_detailed.txt\` - Detailed coverage report
- \`complexity.txt\` - High complexity functions
- \`complexity_all.txt\` - All functions complexity
- \`lint.json\` - Linting results (JSON)
- \`lint.txt\` - Linting results (text)
- \`cdevents-cli\` - Built binary

## Coverage by Package
EOF

# Add package coverage details
go tool cover -func=reports/coverage.out | grep -E "(\.go:.*%)" | head -10 >> reports/quality_report.md

echo ""
echo -e "${GREEN}🎉 Quality Analysis Complete!${NC}"
echo "--------------------------------------"
echo "📁 Reports generated in: ./reports/"
echo "📊 View HTML coverage: ./reports/coverage.html"
echo "📋 Quality summary: ./reports/quality_report.md"

# Copy quality report to docs folder for MkDocs
echo -e "${BLUE}📚 Step 7: Organizing Quality Documentation${NC}"
echo "--------------------------------------"
if [ -f "reports/quality_report.md" ]; then
    # Create docs/quality directory if it doesn't exist
    mkdir -p docs/quality
    
    # Create a proper documentation version with header
    cat > docs/quality/QUALITY_REPORT.md << DOCEOF
# Quality Report

> This report is automatically generated by the quality analysis pipeline.
> Last updated: $(date)

$(cat reports/quality_report.md)

## How to Generate This Report

This report can be regenerated using:

\`\`\`bash
# Using Make (recommended)
make quality-docker

# Or using Docker directly
docker build -f Dockerfile.quality -t cdevents-cli-quality .
docker run --rm -v \$(pwd)/reports:/app/reports cdevents-cli-quality
\`\`\`

## Report Files

The quality analysis generates the following files in \`reports/\`:

- \`coverage.out\` - Coverage profile for tooling
- \`coverage.html\` - Interactive HTML coverage report
- \`coverage_detailed.txt\` - Detailed text coverage report
- \`complexity.txt\` - Functions with high complexity
- \`complexity_all.txt\` - Complete complexity analysis
- \`lint.json\` - Linting results in JSON format
- \`lint.txt\` - Human-readable linting results
- \`quality_report.md\` - This quality summary
- \`cdevents-cli\` - Built binary for testing

## Viewing Reports

### Coverage Report
\`\`\`bash
# View in browser
open reports/coverage.html

# Or view text summary
cat reports/coverage_detailed.txt
\`\`\`

### Complexity Analysis
\`\`\`bash
# View high complexity functions
cat reports/complexity.txt

# View all functions complexity
cat reports/complexity_all.txt
\`\`\`

### Linting Results
\`\`\`bash
# View linting issues
cat reports/lint.txt

# View JSON format for tooling
cat reports/lint.json
\`\`\`
DOCEOF
    
    # Copy coverage report to docs/quality/
    if [ -f "reports/coverage.html" ]; then
        cp reports/coverage.html docs/quality/
        echo -e "${GREEN}✅ Coverage HTML report copied to docs/quality/coverage.html${NC}"
    fi
    
    # Create coverage summary for documentation
    if [ -f "reports/coverage_detailed.txt" ]; then
        cat > docs/quality/coverage-summary.md << COVEOF
# Coverage Summary

> Last updated: $(date)

## Overall Coverage: ${COVERAGE}%

### Coverage by Package

\`\`\`
$(cat reports/coverage_detailed.txt)
\`\`\`

### Coverage Threshold
- **Target**: ${THRESHOLD}%
- **Current**: ${COVERAGE}%
- **Status**: $(if (( $(echo "$COVERAGE >= $THRESHOLD" | bc -l) )); then echo "✅ PASS"; else echo "⚠️ BELOW THRESHOLD"; fi)

### View Interactive Report
[Open Coverage Report](coverage.html)
COVEOF
        echo -e "${GREEN}✅ Coverage summary created in docs/quality/coverage-summary.md${NC}"
    fi
    
    # Create complexity report for documentation
    if [ -f "reports/complexity.txt" ]; then
        cat > docs/quality/complexity-report.md << COMPLEXEOF
# Complexity Analysis

> Last updated: $(date)

## High Complexity Functions (>10)

$(if [ -s "reports/complexity.txt" ]; then
    echo "\`\`\`"
    cat reports/complexity.txt
    echo "\`\`\`"
    echo ""
    echo "**Total functions with high complexity:** $(wc -l <reports/complexity.txt)"
else
    echo "✅ No functions with complexity >10 found!"
fi)

## All Functions Complexity

<details>
<summary>View all functions complexity</summary>

\`\`\`
$(head -50 reports/complexity_all.txt)
\`\`\`

</details>

## Recommendations

- Functions with complexity >10 should be reviewed for refactoring
- Consider breaking down complex functions into smaller, more focused functions
- Use early returns and guard clauses to reduce nesting
- Extract common logic into helper functions
COMPLEXEOF
        echo -e "${GREEN}✅ Complexity report created in docs/quality/complexity-report.md${NC}"
    fi
    
    # Create linting report for documentation
    if [ -f "reports/lint.txt" ]; then
        cat > docs/quality/linting-report.md << LINTEOF
# Linting Report

> Last updated: $(date)

## Linting Status

**Status**: $(if [ -s "reports/lint.txt" ]; then echo "⚠️ ISSUES FOUND"; else echo "✅ CLEAN"; fi)

## Issues Found

$(if [ -s "reports/lint.txt" ]; then
    echo "\`\`\`"
    cat reports/lint.txt
    echo "\`\`\`"
else
    echo "✅ No linting issues found!"
fi)

## Linter Configuration

The project uses golangci-lint with the following modern linters:

- **revive** (replaces deprecated golint)
- **unused** (replaces deprecated deadcode, structcheck, varcheck)
- **exportloopref** (replaces deprecated scopelint)
- **staticcheck** for advanced static analysis
- **govet** for Go compiler checks
- **errcheck** for error handling verification

## Running Linting

\`\`\`bash
# Run linting locally
golangci-lint run

# Or using Docker
make quality-docker
\`\`\`
LINTEOF
        echo -e "${GREEN}✅ Linting report created in docs/quality/linting-report.md${NC}"
    fi
    
    # Create performance metrics for documentation
    cat > docs/quality/performance-metrics.md << PERFEOF
# Performance Metrics

> Last updated: $(date)

## Build Performance

- **Build Time**: ${BUILD_TIME}s
- **Binary Size**: ${BINARY_SIZE}
- **Test Execution Time**: ${TEST_TIME}s

## Performance Benchmarks

| Metric | Target | Current | Status |
|--------|--------|---------|--------|
| Build Time | <5s | ${BUILD_TIME}s | $(if (( $(echo "$BUILD_TIME < 5" | bc -l) )); then echo "✅ PASS"; else echo "⚠️ SLOW"; fi) |
| Test Time | <10s | ${TEST_TIME}s | $(if (( $(echo "$TEST_TIME < 10" | bc -l) )); then echo "✅ PASS"; else echo "⚠️ SLOW"; fi) |
| Binary Size | <20MB | ${BINARY_SIZE} | $(if [[ "${BINARY_SIZE}" =~ ^[0-9]+\.[0-9]+M$ ]] && (( $(echo "${BINARY_SIZE%M} < 20" | bc -l) )); then echo "✅ PASS"; else echo "⚠️ LARGE"; fi) |

## Performance Trends

Regular monitoring of these metrics helps identify performance regressions:

- Monitor build time trends
- Track binary size growth
- Optimize test execution
- Profile performance bottlenecks

## Optimization Tips

### Build Time
- Use Go module proxy for faster dependency resolution
- Optimize Docker build layers
- Use build caching

### Test Performance
- Use parallel test execution
- Optimize test setup/teardown
- Use test fixtures efficiently

### Binary Size
- Use build constraints for optional features
- Strip debug information for production builds
- Use UPX compression if needed
PERFEOF
    echo -e "${GREEN}✅ Performance metrics created in docs/quality/performance-metrics.md${NC}"
    
    echo -e "${GREEN}✅ All quality documentation organized in docs/quality/${NC}"
else
    echo -e "${YELLOW}⚠️  Quality report not found, skipping docs copy${NC}"
fi
echo ""

# Display final summary
echo -e "${BLUE}📋 Quality Summary${NC}"
echo "=================="
echo "Coverage: ${COVERAGE}%"
echo "Build Time: ${BUILD_TIME}s"
echo "Binary Size: ${BINARY_SIZE}"
echo "Test Time: ${TEST_TIME}s"
echo "High Complexity Functions: $(wc -l < reports/complexity.txt)"
echo ""

# Exit with appropriate code
if (( $(echo "$COVERAGE >= $THRESHOLD" | bc -l) )) && [ ! -s "reports/complexity.txt" ]; then
    echo -e "${GREEN}✅ All quality checks passed!${NC}"
    exit 0
else
    echo -e "${YELLOW}⚠️  Some quality checks need attention${NC}"
    exit 1
fi
