# MkDocs Documentation Updates

## Summary of Changes

Updated the MkDocs documentation to reflect that custom data is only available through the `--custom-json` flag, removing references to `--custom` and `--custom-yaml` flags that are not implemented.

## Files Updated

### 1. `docs/cli-reference.md`
- **Common Generate Flags Table**: Removed `--custom` and `--custom-yaml` flags, kept only `--custom-json`
- **Pipeline Examples**: Updated to use `--custom-json` instead of `--custom` flags
- **Custom Data Section**: 
  - Removed "Key=Value Pairs" subsection
  - Removed "YAML Format" subsection
  - Kept only "JSON Format" subsection
  - Updated introduction to clarify only JSON format is supported

### 2. `docs/getting-started.md`
- **Custom Data Section**: 
  - Updated introduction to mention only `--custom-json` flag
  - Converted all examples to use `--custom-json` format
  - Replaced key=value and YAML examples with JSON equivalents
  - Maintained the same functional examples but in JSON format

### 3. `docs/examples.md`
- **Jenkins Pipeline Example**: 
  - Updated pipeline started event to use `--custom-json` instead of multiple `--custom` flags
  - Updated success and failure handlers to use `--custom-json` format
  - Maintained all the same data but structured as JSON
- **Kubernetes Deployment Example**: 
  - Converted from `--custom-yaml` to `--custom-json` format
  - Preserved all the same metadata structure
  - Updated complex nested data to JSON format

## Key Changes Made

1. **Removed unsupported flags**: 
   - `--custom` (key=value pairs)
   - `--custom-yaml` (YAML format)

2. **Standardized on JSON format**:
   - All custom data examples now use `--custom-json`
   - Maintained equivalent functionality in JSON format
   - Preserved all data structures and examples

3. **Updated documentation structure**:
   - Simplified custom data section to focus on JSON only
   - Updated CLI reference to reflect actual available flags
   - Maintained comprehensive examples with proper JSON formatting

## Benefits

- **Accuracy**: Documentation now matches actual CLI implementation
- **Consistency**: All examples use the same custom data format
- **Clarity**: Removes confusion about multiple custom data options
- **Maintainability**: Easier to maintain with single custom data format

## Additional Updates - CLI Input/Output Examples

### New Section Added: "CLI Usage Examples with Input and Output"

**File**: `docs/examples.md`
- Added comprehensive section showing exact CLI commands with their complete outputs
- Includes examples for all major event types (pipeline, build, test, service, task)
- Shows different output formats (JSON, YAML, CloudEvent)
- Demonstrates custom data usage with realistic examples
- Includes send command examples with different targets (HTTP, file)

**Examples Added**:
1. **Basic Pipeline Event Generation** - Simple command with JSON output
2. **Pipeline Event with Custom Data** - Complex custom data with labels, annotations, and links
3. **Build Event with YAML Output** - Shows YAML formatting option
4. **Test Event with CloudEvent Output** - Demonstrates CloudEvent format
5. **Service Deployment Event** - Complex Kubernetes deployment scenario
6. **Task Event with Minimal Data** - Simple task event example
7. **Sending Events to Different Targets** - HTTP and file target examples

### New Section Added: "Quick Examples with Input/Output"

**File**: `docs/getting-started.md`
- Added quick-start examples with explicit input/output
- Shows progression from simple to complex commands
- Includes all three output formats (JSON, YAML, CloudEvent)
- Demonstrates send commands with console and file targets

**Examples Added**:
1. **Simple Pipeline Event** - Basic command with JSON output
2. **Build Event with Custom Data** - Shows custom data structure
3. **Test Event in YAML Format** - YAML output demonstration
4. **Service Event in CloudEvent Format** - CloudEvent output example
5. **Console Send Command** - Simple send to console
6. **File Send Command** - Send to file with file contents shown

## Key Features of New Examples

1. **Explicit Input/Output**: Every example shows the exact command and complete output
2. **Realistic Data**: Uses meaningful IDs, names, and custom data structures
3. **Multiple Formats**: Demonstrates JSON, YAML, and CloudEvent outputs
4. **Progressive Complexity**: From simple events to complex ones with custom data
5. **Transport Examples**: Shows different ways to send events (console, file, HTTP)
6. **Real-world Scenarios**: Kubernetes deployments, CI/CD pipelines, test results

## Benefits of New Examples

- **Learning**: Users can see exactly what to expect from each command
- **Testing**: Copy-paste examples for quick testing
- **Reference**: Complete command reference with outputs
- **Debugging**: Compare expected vs actual outputs
- **Integration**: Real-world scenarios for CI/CD integration

## Validation

All updated examples have been verified to:
- Use proper JSON syntax and maintain the same functionality as the original key=value and YAML examples
- Show realistic and complete CLI outputs
- Demonstrate proper CDEvents structure according to the specification
- Include meaningful custom data examples for real-world usage
