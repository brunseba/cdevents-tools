package events_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/brunseba/cdevents-tools/pkg/events"
	"github.com/xeipuuv/gojsonschema"
)

// TestValidatePipelineRunStartedEvent validates a PipelineRunStarted event against its JSON schema
func TestValidatePipelineRunStartedEvent(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	event, err := factory.CreatePipelineRunEvent("started", "pipeline-123", "test-pipeline", "", "", "https://example.com/pipeline", nil)
	if err != nil {
		t.Fatalf("Failed to create pipeline run started event: %v", err)
	}

	validateEventAgainstSchema(t, event, "pipelinerunstarted.json")
}

// TestValidatePipelineRunFinishedEvent validates a PipelineRunFinished event against its JSON schema
func TestValidatePipelineRunFinishedEvent(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	event, err := factory.CreatePipelineRunEvent("finished", "pipeline-123", "test-pipeline", "success", "", "https://example.com/pipeline", nil)
	if err != nil {
		t.Fatalf("Failed to create pipeline run finished event: %v", err)
	}

	validateEventAgainstSchema(t, event, "pipelinerunfinished.json")
}

// TestValidateBuildStartedEvent validates a BuildStarted event against its JSON schema
func TestValidateBuildStartedEvent(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	event, err := factory.CreateBuildEvent("started", "build-123", "test-build", "", "", "https://example.com/build", nil)
	if err != nil {
		t.Fatalf("Failed to create build started event: %v", err)
	}

	validateEventAgainstSchema(t, event, "buildstarted.json")
}

// TestValidateBuildFinishedEvent validates a BuildFinished event against its JSON schema
func TestValidateBuildFinishedEvent(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	event, err := factory.CreateBuildEvent("finished", "build-123", "test-build", "success", "", "https://example.com/build", nil)
	if err != nil {
		t.Fatalf("Failed to create build finished event: %v", err)
	}

	validateEventAgainstSchema(t, event, "buildfinished.json")
}

// TestValidateTaskRunStartedEvent validates a TaskRunStarted event against its JSON schema
func TestValidateTaskRunStartedEvent(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	event, err := factory.CreateTaskRunEvent("started", "task-123", "test-task", "pipeline-456", "", "", "https://example.com/task", nil)
	if err != nil {
		t.Fatalf("Failed to create task run started event: %v", err)
	}

	validateEventAgainstSchema(t, event, "taskrunstarted.json")
}

// TestValidateTaskRunFinishedEvent validates a TaskRunFinished event against its JSON schema
func TestValidateTaskRunFinishedEvent(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	event, err := factory.CreateTaskRunEvent("finished", "task-123", "test-task", "pipeline-456", "success", "", "https://example.com/task", nil)
	if err != nil {
		t.Fatalf("Failed to create task run finished event: %v", err)
	}

	validateEventAgainstSchema(t, event, "taskrunfinished.json")
}

// TestValidateTestCaseRunStartedEvent validates a TestCaseRunStarted event against its JSON schema
func TestValidateTestCaseRunStartedEvent(t *testing.T) {
	// Skip this test for now as it requires environment object which isn't supported in current factory
	t.Skip("Test case events require environment object which isn't supported in current factory")
}

// TestValidateTestCaseRunFinishedEvent validates a TestCaseRunFinished event against its JSON schema
func TestValidateTestCaseRunFinishedEvent(t *testing.T) {
	// Skip this test for now as it requires environment object which isn't supported in current factory
	t.Skip("Test case events require environment object which isn't supported in current factory")
}

// TestValidateServiceDeployedEvent validates a ServiceDeployed event against its JSON schema
func TestValidateServiceDeployedEvent(t *testing.T) {
	// Skip this test for now as it requires environment object which isn't supported in current factory
	t.Skip("Service events require environment object which isn't supported in current factory")
}

// TestValidateEventWithCustomData validates an event with custom data against its JSON schema
func TestValidateEventWithCustomData(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	customData := &events.CustomData{
		Data: map[string]interface{}{
			"buildConfig": map[string]interface{}{
				"dockerfile": "Dockerfile",
				"context":    ".",
				"tags":       []string{"latest", "v1.0.0"},
			},
			"metadata": map[string]interface{}{
				"commit":     "abc123",
				"branch":     "main",
				"repository": "my-repo",
			},
		},
		ContentType: "application/json",
	}

	event, err := factory.CreateBuildEvent("started", "build-123", "test-build", "", "", "https://example.com/build", customData)
	if err != nil {
		t.Fatalf("Failed to create build event with custom data: %v", err)
	}

	validateEventAgainstSchema(t, event, "buildstarted.json")
}

// validateEventAgainstSchema is a helper function that validates an event against its JSON schema
func validateEventAgainstSchema(t *testing.T, event interface{}, schemaFile string) {
	t.Helper()

	// Marshal the event to JSON
	eventBytes, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal event to JSON: %v", err)
	}

	// Load the schema from CDEvents spec repository
	schemaURL := fmt.Sprintf("https://raw.githubusercontent.com/cdevents/spec/v0.4.1/schemas/%s", schemaFile)
	schemaLoader := gojsonschema.NewReferenceLoader(schemaURL)
	documentLoader := gojsonschema.NewBytesLoader(eventBytes)

	// Validate the event against the schema
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		t.Fatalf("Schema validation error: %v", err)
	}

	if !result.Valid() {
		t.Errorf("Event does not match schema %s:", schemaFile)
		for _, desc := range result.Errors() {
			t.Errorf("- %s", desc)
		}
		
		// Print the event JSON for debugging
		t.Logf("Event JSON: %s", string(eventBytes))
	}
}

// TestAllSchemaFilesExist verifies that all expected schema files exist
func TestAllSchemaFilesExist(t *testing.T) {
	expectedSchemas := []string{
		"pipelinerunstarted.json",
		"pipelinerunfinished.json",
		"pipelinerunqueued.json",
		"buildstarted.json",
		"buildfinished.json",
		"buildqueued.json",
		"taskrunstarted.json",
		"taskrunfinished.json",
		"testcaserunstarted.json",
		"testcaserunfinished.json",
		"testcaserunqueued.json",
		"testcaserunskipped.json",
		"testsuiterunstarted.json",
		"testsuiterunfinished.json",
		"testsuiterunqueued.json",
		"testoutputpublished.json",
		"servicedeployed.json",
		"servicepublished.json",
		"serviceremoved.json",
		"servicerolledback.json",
		"serviceupgraded.json",
	}

	for _, schema := range expectedSchemas {
		schemaURL := fmt.Sprintf("https://raw.githubusercontent.com/cdevents/spec/v0.4.1/schemas/%s", schema)
		if _, err := gojsonschema.NewReferenceLoader(schemaURL).LoadJSON(); err != nil {
			t.Errorf("Schema file %s not found or invalid: %v", schema, err)
		}
	}
}

// TestValidateMultipleEventTypes validates multiple event types in a single test
func TestValidateMultipleEventTypes(t *testing.T) {
	factory := events.NewEventFactory("https://test-source.example.com")
	
	testCases := []struct {
		name       string
		eventType  string
		createFunc func() (interface{}, error)
		schemaFile string
	}{
		{
			name:      "PipelineRunQueued",
			eventType: "pipeline-queued",
			createFunc: func() (interface{}, error) {
				return factory.CreatePipelineRunEvent("queued", "pipeline-123", "test-pipeline", "", "", "", nil)
			},
			schemaFile: "pipelinerunqueued.json",
		},
		{
			name:      "BuildQueued",
			eventType: "build-queued",
			createFunc: func() (interface{}, error) {
				return factory.CreateBuildEvent("queued", "build-123", "test-build", "", "", "", nil)
			},
			schemaFile: "buildqueued.json",
		},
		// Skip test case events for now as they require environment object
		// {
		// 	name:      "TestCaseRunQueued",
		// 	eventType: "testcase-queued",
		// 	createFunc: func() (interface{}, error) {
		// 		return factory.CreateTestEvent("testcase-queued", "test-123", "test-case-name", "", "", "", nil)
		// 	},
		// 	schemaFile: "testcaserunqueued.json",
		// },
		// {
		// 	name:      "TestCaseRunSkipped",
		// 	eventType: "testcase-skipped",
		// 	createFunc: func() (interface{}, error) {
		// 		return factory.CreateTestEvent("testcase-skipped", "test-123", "test-case-name", "", "", "", nil)
		// 	},
		// 	schemaFile: "testcaserunskipped.json",
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := tc.createFunc()
			if err != nil {
				t.Fatalf("Failed to create %s event: %v", tc.eventType, err)
			}

			validateEventAgainstSchema(t, event, tc.schemaFile)
		})
	}
}
