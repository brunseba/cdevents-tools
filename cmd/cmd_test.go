package cmd_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/brunseba/cdevents-tools/cmd"
	"github.com/brunseba/cdevents-tools/pkg/events"
	"github.com/cdevents/sdk-go/pkg/api"
)

func TestRootCommand(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct {
		name string
		args []string
		expectError bool
	}{
		{
			name: "help flag",
			args: []string{"cdevents-cli", "--help"},
			expectError: false,
		},
		{
			name: "version flag",
			args: []string{"cdevents-cli", "--version"},
			expectError: false,
		},
		{
			name: "verbose flag",
			args: []string{"cdevents-cli", "--verbose", "--help"},
			expectError: false,
		},
		{
			name: "invalid flag",
			args: []string{"cdevents-cli", "--invalid"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := cmd.Execute()
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestGenerateCommands(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct {
		name string
		args []string
		expectError bool
	}{
		{
			name: "generate pipeline started",
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: false,
		},
		{
			name: "generate pipeline finished with outcome",
			args: []string{"cdevents-cli", "generate", "pipeline", "finished", "--id", "123", "--name", "test-pipeline", "--outcome", "success"},
			expectError: false,
		},
		{
			name: "generate build started",
			args: []string{"cdevents-cli", "generate", "build", "started", "--id", "456", "--name", "test-build"},
			expectError: false,
		},
		{
			name: "generate build finished with custom data",
			args: []string{"cdevents-cli", "generate", "build", "finished", "--id", "456", "--name", "test-build", "--custom-json", `{"key":"value"}`},
			expectError: false,
		},
		{
			name: "generate task started",
			args: []string{"cdevents-cli", "generate", "task", "started", "--id", "789", "--name", "test-task"},
			expectError: false,
		},
		{
			name: "generate service deployed",
			args: []string{"cdevents-cli", "generate", "service", "deployed", "--id", "101", "--name", "test-service"},
			expectError: false,
		},
		{
			name: "generate test finished",
			args: []string{"cdevents-cli", "generate", "test", "finished", "--id", "111", "--name", "test-case"},
			expectError: true, // "finished" is not a valid test event type
		},
		{
			name: "generate with yaml output",
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--output", "yaml"},
			expectError: false,
		},
		{
			name: "generate with cloudevent output",
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--output", "cloudevent"},
			expectError: false,
		},
		{
			name: "generate pipeline missing id",
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--name", "test-pipeline"},
			expectError: false, // CLI uses empty string as default
		},
		{
			name: "generate pipeline missing name",
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123"},
			expectError: false, // CLI uses empty string as default
		},
		{
			name: "generate pipeline invalid event type",
			args: []string{"cdevents-cli", "generate", "pipeline", "invalid", "--id", "123", "--name", "test-pipeline"},
			expectError: true,
		},
		{
			name: "generate with invalid custom json",
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--custom-json", `{"invalid json`},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := cmd.Execute()
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestSendCommands(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct {
		name string
		args []string
		expectError bool
	}{
		{
			name: "send to console",
			args: []string{"cdevents-cli", "send", "--target", "console", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: false,
		},
		{
			name: "send to file",
			args: []string{"cdevents-cli", "send", "--target", "file://test.json", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: false,
		},
		{
			name: "send with retries",
			args: []string{"cdevents-cli", "send", "--target", "console", "--retries", "5", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: false,
		},
		{
			name: "send with timeout",
			args: []string{"cdevents-cli", "send", "--target", "console", "--timeout", "10s", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: false,
		},
		{
			name: "send with headers",
			args: []string{"cdevents-cli", "send", "--target", "console", "--headers", "Authorization=Bearer token", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: false,
		},
		{
			name: "send invalid target",
			args: []string{"cdevents-cli", "send", "--target", "invalid://target", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
			expectError: true,
		},
		{
			name: "send missing id",
			args: []string{"cdevents-cli", "send", "--target", "console", "pipeline", "started", "--name", "test-pipeline"},
			expectError: false, // CLI uses empty string as default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := cmd.Execute()
			if tt.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestSendEventWithRetry(t *testing.T) {
	// Create a mock transport that fails first attempts
	mockTransport := &MockTransport{
		failCount: 2, // fail first 2 attempts
	}

	// Create a mock event
	factory := events.NewEventFactory("test-source")
	event, err := factory.CreatePipelineRunEvent("started", "test-id", "test-name", "", "", "", nil)
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	ctx := context.Background()
	
	// Test successful retry
	err = cmd.SendEventWithRetry(ctx, mockTransport, event, 3)
	if err != nil {
		t.Errorf("expected success after retries, got error: %v", err)
	}
	
	// Test failure after max retries
	mockTransport.failCount = 10 // fail all attempts
	err = cmd.SendEventWithRetry(ctx, mockTransport, event, 3)
	if err == nil {
		t.Error("expected failure after max retries, got success")
	}
}

// MockTransport is a mock implementation of transport.Transport for testing
type MockTransport struct {
	failCount int
	callCount int
}

func (m *MockTransport) Send(ctx context.Context, event api.CDEvent) error {
	m.callCount++
	if m.callCount <= m.failCount {
		return errors.New("mock transport failure")
	}
	return nil
}

func (m *MockTransport) String() string {
	return "mock-transport"
}

func TestConfigInit(t *testing.T) {
	// Test config initialization by setting a config file
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Create a temporary config file
	tempFile, err := os.CreateTemp("", "cdevents-cli-test-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some config content
	_, err = tempFile.WriteString("verbose: true\noutput: yaml\n")
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}
	tempFile.Close()

	// Test with config file
	os.Args = []string{"cdevents-cli", "--config", tempFile.Name(), "--help"}
	err = cmd.Execute()
	if err != nil {
		t.Errorf("unexpected error with config file: %v", err)
	}
}

func TestInvalidEventType(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = []string{"cdevents-cli", "generate", "pipeline", "invalid-event-type", "--id", "123", "--name", "test"}
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error for invalid event type, got none")
	}
}

func TestCustomDataParsing(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Test with valid custom JSON
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--custom-json", `{"key":"value","num":42}`}
	err := cmd.Execute()
	if err != nil {
		t.Errorf("unexpected error with valid custom JSON: %v", err)
	}

	// Test with invalid custom JSON
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--custom-json", `{"invalid":json`}
	err = cmd.Execute()
	if err == nil {
		t.Error("expected error with invalid custom JSON, got none")
	}
}

func TestOutputFormats(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	formats := []string{"json", "yaml", "cloudevent"}
	
	for _, format := range formats {
		t.Run("format_"+format, func(t *testing.T) {
			os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--output", format}
			err := cmd.Execute()
			if err != nil {
				t.Errorf("unexpected error with format %s: %v", format, err)
			}
		})
	}
}

func TestErrorHandling(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Test with no arguments to subcommand
	os.Args = []string{"cdevents-cli", "generate", "pipeline"}
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error with no event type, got none")
	}

	// Test with invalid subcommand
	os.Args = []string{"cdevents-cli", "invalid-command"}
	err = cmd.Execute()
	if err == nil {
		t.Error("expected error with invalid command, got none")
	}
}

func TestEdgeCases(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Test with empty custom JSON
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--custom-json", ""}
	err := cmd.Execute()
	if err != nil {
		t.Errorf("unexpected error with empty custom JSON: %v", err)
	}

	// Test with URL parameter
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--url", "https://example.com"}
	err = cmd.Execute()
	if err != nil {
		t.Errorf("unexpected error with URL parameter: %v", err)
	}

	// Test with error details
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "finished", "--id", "123", "--name", "test", "--outcome", "failure", "--errors", "Build failed"}
	err = cmd.Execute()
	if err != nil {
		t.Errorf("unexpected error with error details: %v", err)
	}
}
