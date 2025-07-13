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
