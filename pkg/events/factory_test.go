package events_test

import (
	"encoding/json"
	"testing"

	"github.com/brunseba/cdevents-tools/pkg/events"
)

func TestValidateNewPipelineRunStartedEvent(t *testing.T) {
	// Test that we can create a valid pipeline run started event
	eventFactory := events.NewEventFactory("test-source")
	customData := &events.CustomData{
		Data: map[string]interface{}{"key": "value"},
		ContentType: "application/json",
	}
	event, err := eventFactory.CreatePipelineRunEvent("started", "pipeline-123", "test-pipeline", "status", "", "https://example.com", customData)
	if err != nil {
		t.Fatalf("Failed to create event: %v", err)
	}

	// Verify the event can be marshaled to JSON
	eventBytes, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal event: %v", err)
	}

	// Verify the JSON is not empty
	if len(eventBytes) == 0 {
		t.Fatalf("Event marshaled to empty JSON")
	}

	// Verify we can unmarshal it back
	var eventMap map[string]interface{}
	err = json.Unmarshal(eventBytes, &eventMap)
	if err != nil {
		t.Fatalf("Failed to unmarshal event JSON: %v", err)
	}

	// Verify required fields exist
	if _, ok := eventMap["context"]; !ok {
		t.Errorf("Event missing context field")
	}
	if _, ok := eventMap["subject"]; !ok {
		t.Errorf("Event missing subject field")
	}
}

func TestNewEventFactory(t *testing.T) {
	factory := events.NewEventFactory("test-source")
	if factory == nil {
		t.Fatalf("NewEventFactory should return a non-nil factory")
	}
}

func TestCreatePipelineRunEventWithCustomData(t *testing.T) {
	factory := events.NewEventFactory("test-source")
	customData := &events.CustomData{
		Data: map[string]interface{}{"key": "value"},
		ContentType: "application/json",
	}

	event, err := factory.CreatePipelineRunEvent(
		"finished",
		"pipeline-123",
		"test-pipeline",
		"success",
		"error details",
		"https://example.com",
		customData,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}
}

func TestCreateTaskRunEventWithCustomData(t *testing.T) {
	factory := events.NewEventFactory("test-source")
	customData := &events.CustomData{
		Data: map[string]interface{}{"key": "value"},
		ContentType: "application/json",
	}

	event, err := factory.CreateTaskRunEvent(
		"finished",
		"task-123",
		"test-task",
		"pipeline-456",
		"failure",
		"task failed",
		"https://example.com",
		customData,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}
}

func TestCreateBuildEventWithCustomData(t *testing.T) {
	factory := events.NewEventFactory("test-source")
	customData := &events.CustomData{
		Data: map[string]interface{}{"key": "value"},
		ContentType: "application/json",
	}

	event, err := factory.CreateBuildEvent(
		"finished",
		"build-123",
		"test-build",
		"error",
		"build failed",
		"https://example.com",
		customData,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}
}

func TestCreateServiceEventWithCustomData(t *testing.T) {
	factory := events.NewEventFactory("test-source")
	customData := &events.CustomData{
		Data: map[string]interface{}{"key": "value"},
		ContentType: "application/json",
	}

	event, err := factory.CreateServiceEvent(
		"deployed",
		"service-123",
		"test-service",
		"env-123",
		"https://example.com",
		customData,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}
}

func TestCreateTestEventWithCustomData(t *testing.T) {
	factory := events.NewEventFactory("test-source")
	customData := &events.CustomData{
		Data: map[string]interface{}{"key": "value"},
		ContentType: "application/json",
	}

	event, err := factory.CreateTestEvent(
		"testcase-finished",
		"test-123",
		"test-name",
		"success",
		"",
		"https://example.com",
		customData,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}
}

func TestCreateTestEventWithoutOutcome(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	// Test events that don't have outcome/errors
	eventTypes := []string{"testcase-queued", "testcase-started", "testcase-skipped", "testsuite-queued", "testsuite-started", "testoutput-published"}
	for _, eventType := range eventTypes {
		t.Run(eventType, func(t *testing.T) {
			event, err := factory.CreateTestEvent(
				eventType,
				"test-123",
				"test-name",
				"", // no outcome
				"", // no errors
				"https://example.com",
				nil,
			)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if event == nil {
				t.Fatalf("event should not be nil")
			}
		})
	}
}

func TestCreateEventsWithoutOptionalFields(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	// Test pipeline event without URL and outcome/errors
	event, err := factory.CreatePipelineRunEvent(
		"queued",
		"pipeline-123",
		"test-pipeline",
		"", // no outcome
		"", // no errors
		"", // no URL
		nil,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}

	// Test task event without URL and pipeline run ID
	event, err = factory.CreateTaskRunEvent(
		"started",
		"task-123",
		"test-task",
		"", // no pipeline run ID
		"", // no outcome
		"", // no errors
		"", // no URL
		nil,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}

	// Test service event without URL and environment ID
	event, err = factory.CreateServiceEvent(
		"published",
		"service-123",
		"test-service",
		"", // no environment ID
		"", // no URL
		nil,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}
}

func TestCreatePipelineRunEvent(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	testCases := []struct {
		name      string
		eventType string
		shouldErr bool
	}{
		{"queued event", "queued", false},
		{"started event", "started", false},
		{"finished event", "finished", false},
		{"invalid event", "invalid", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := factory.CreatePipelineRunEvent(
				tc.eventType,
				"pipeline-123",
				"test-pipeline",
				"success",
				"",
				"https://example.com",
				nil,
			)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for event type %s", tc.eventType)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if event == nil {
				t.Fatalf("event should not be nil")
			}

			if event.GetId() == "" {
				t.Errorf("event ID should not be empty")
			}

			if event.GetSource() != "test-source" {
				t.Errorf("expected source 'test-source', got '%s'", event.GetSource())
			}
		})
	}
}

func TestCreateTaskRunEvent(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	testCases := []struct {
		name      string
		eventType string
		shouldErr bool
	}{
		{"started event", "started", false},
		{"finished event", "finished", false},
		{"invalid event", "invalid", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := factory.CreateTaskRunEvent(
				tc.eventType,
				"task-123",
				"test-task",
				"pipeline-123",
				"success",
				"",
				"https://example.com",
				nil,
			)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for event type %s", tc.eventType)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if event == nil {
				t.Fatalf("event should not be nil")
			}

			if event.GetId() == "" {
				t.Errorf("event ID should not be empty")
			}

			if event.GetSource() != "test-source" {
				t.Errorf("expected source 'test-source', got '%s'", event.GetSource())
			}
		})
	}
}

func TestCreateBuildEvent(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	testCases := []struct {
		name      string
		eventType string
		shouldErr bool
	}{
		{"queued event", "queued", false},
		{"started event", "started", false},
		{"finished event", "finished", false},
		{"invalid event", "invalid", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := factory.CreateBuildEvent(
				tc.eventType,
				"build-123",
				"test-build",
				"success",
				"",
				"https://example.com",
				nil,
			)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for event type %s", tc.eventType)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if event == nil {
				t.Fatalf("event should not be nil")
			}

			if event.GetId() == "" {
				t.Errorf("event ID should not be empty")
			}

			if event.GetSource() != "test-source" {
				t.Errorf("expected source 'test-source', got '%s'", event.GetSource())
			}
		})
	}
}

func TestCreateServiceEvent(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	testCases := []struct {
		name      string
		eventType string
		shouldErr bool
	}{
		{"deployed event", "deployed", false},
		{"published event", "published", false},
		{"removed event", "removed", false},
		{"rolledback event", "rolledback", false},
		{"upgraded event", "upgraded", false},
		{"invalid event", "invalid", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := factory.CreateServiceEvent(
				tc.eventType,
				"service-123",
				"test-service",
				"env-123",
				"https://example.com",
				nil,
			)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for event type %s", tc.eventType)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if event == nil {
				t.Fatalf("event should not be nil")
			}

			if event.GetId() == "" {
				t.Errorf("event ID should not be empty")
			}

			if event.GetSource() != "test-source" {
				t.Errorf("expected source 'test-source', got '%s'", event.GetSource())
			}
		})
	}
}

func TestApplyCustomDataFunction(t *testing.T) {
	// This function doesn't do anything in the current implementation
	// but we can test that it doesn't panic
	factory := events.NewEventFactory("test-source")
	event, err := factory.CreatePipelineRunEvent(
		"started",
		"pipeline-123",
		"test-pipeline",
		"",
		"",
		"",
		&events.CustomData{
			Data: map[string]interface{}{"key": "value"},
			ContentType: "application/json",
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event == nil {
		t.Fatalf("event should not be nil")
	}

	// The applyCustomData function is called internally when custom data is provided
	// This test ensures the code path is covered
}

func TestCreateTestEvent(t *testing.T) {
	factory := events.NewEventFactory("test-source")

	testCases := []struct {
		name      string
		eventType string
		shouldErr bool
	}{
		{"testcase-queued event", "testcase-queued", false},
		{"testcase-started event", "testcase-started", false},
		{"testcase-finished event", "testcase-finished", false},
		{"testcase-skipped event", "testcase-skipped", false},
		{"testsuite-queued event", "testsuite-queued", false},
		{"testsuite-started event", "testsuite-started", false},
		{"testsuite-finished event", "testsuite-finished", false},
		{"testoutput-published event", "testoutput-published", false},
		{"invalid event", "invalid", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			event, err := factory.CreateTestEvent(
				tc.eventType,
				"test-123",
				"test-name",
				"success",
				"",
				"https://example.com",
				nil,
			)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for event type %s", tc.eventType)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if event == nil {
				t.Fatalf("event should not be nil")
			}

			if event.GetId() == "" {
				t.Errorf("event ID should not be empty")
			}

			if event.GetSource() != "test-source" {
				t.Errorf("expected source 'test-source', got '%s'", event.GetSource())
			}
		})
	}
}

func TestParseCustomDataFromJSON(t *testing.T) {
	testCases := []struct {
		name      string
		jsonData  string
		shouldErr bool
	}{
		{"empty string", "", false},
		{"valid json", `{"key": "value"}`, false},
		{"invalid json", `{invalid}`, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			customData, err := events.ParseCustomDataFromJSON(tc.jsonData)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for JSON data: %s", tc.jsonData)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tc.jsonData == "" {
				if customData != nil {
					t.Errorf("expected nil for empty string")
				}
			} else {
				if customData == nil {
					t.Errorf("expected non-nil custom data")
				}
				if customData.ContentType != "application/json" {
					t.Errorf("expected content type 'application/json', got '%s'", customData.ContentType)
				}
			}
		})
	}
}
