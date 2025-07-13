package cmd_test

import (
	"os"
	"testing"

	"github.com/cdevents/cdevents-cli/cmd"
)

func TestExecuteFunction(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Set up args for help
	os.Args = []string{"cdevents-cli", "--help"}
	
	// This should not panic or return error for help
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error running help: %v", err)
	}
}

func TestGeneratePipelineCommand(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Set up args for generate pipeline
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline"}
	
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing generate pipeline: %v", err)
	}
}

func TestSendCommand(t *testing.T) {
	// Save original args
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	
	// Set up args for send command
	os.Args = []string{"cdevents-cli", "send", "--target", "console", "pipeline", "started", "--id", "123", "--name", "test-pipeline"}
	
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing send: %v", err)
	}
}
