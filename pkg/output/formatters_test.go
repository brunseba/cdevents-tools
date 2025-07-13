package output_test

import (
	"strings"
	"testing"
	"github.com/cdevents/sdk-go/pkg/api"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	"github.com/brunseba/cdevents-tools/pkg/output"
)

func TestFormatJSONWithCustomData(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	// Set basic properties
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	customData := output.CustomData{
		Data: map[string]interface{}{
			"mydata": "value123",
		},
		ContentType: "application/json",
	}

	formatted, err := output.FormatOutputWithCustomData(event, &customData, "json")
	if err != nil {
		t.Fatalf("failed to format output with custom data: %v", err)
	}

	// Check that the formatted output contains the custom data
	if !strings.Contains(formatted, `"customData"`) {
		t.Errorf("formatted output should contain customData field")
	}
	if !strings.Contains(formatted, `"customDataContentType"`) {
		t.Errorf("formatted output should contain customDataContentType field")
	}
	if !strings.Contains(formatted, `"mydata": "value123"`) {
		t.Errorf("formatted output should contain custom data values")
	}
}

func TestFormatYAMLWithCustomDataError(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test with custom data that has complex nested structures
	customData := output.CustomData{
		Data: map[string]interface{}{
			"complex": map[string]interface{}{
				"nested": []string{"array", "of", "values"},
				"number": 42,
				"boolean": true,
			},
		},
		ContentType: "application/json",
	}

	formatted, err := output.FormatOutputWithCustomData(event, &customData, "yaml")
	if err != nil {
		t.Fatalf("failed to format YAML with complex custom data: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted YAML should not be empty")
	}

	// Check that complex data is present
	if !strings.Contains(formatted, "complex:") {
		t.Errorf("YAML should contain complex custom data")
	}
}

func TestFormatCloudEventWithoutCustomData(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test CloudEvent formatting without custom data (different code path)
	formatted, err := output.FormatOutputWithCustomData(event, nil, "cloudevent")
	if err != nil {
		t.Fatalf("failed to format CloudEvent without custom data: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted CloudEvent should not be empty")
	}

	// Should not contain custom data fields
	if strings.Contains(formatted, "customData") {
		t.Errorf("CloudEvent should not contain customData when none provided")
	}
}

// Test individual format functions
func TestFormatJSON(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test formatJSON function by calling it through FormatOutput
	formatted, err := output.FormatOutput(event, "json")
	if err != nil {
		t.Fatalf("failed to format JSON: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted JSON should not be empty")
	}
}

func TestFormatYAML(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test formatYAML function by calling it through FormatOutput
	formatted, err := output.FormatOutput(event, "yaml")
	if err != nil {
		t.Fatalf("failed to format YAML: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted YAML should not be empty")
	}
}

func TestFormatCloudEvent(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test formatCloudEvent function by calling it through FormatOutput
	formatted, err := output.FormatOutput(event, "cloudevent")
	if err != nil {
		t.Fatalf("failed to format CloudEvent: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted CloudEvent should not be empty")
	}
}

func TestFormatYAMLWithoutCustomData(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test YAML formatting without custom data (different code path)
	formatted, err := output.FormatOutputWithCustomData(event, nil, "yaml")
	if err != nil {
		t.Fatalf("failed to format YAML without custom data: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted YAML should not be empty")
	}

	// Should not contain custom data fields
	if strings.Contains(formatted, "customData") {
		t.Errorf("YAML should not contain customData when none provided")
	}
}

func TestFormatYAMLWithCustomData(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	// Set basic properties
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	customData := output.CustomData{
		Data: map[string]interface{}{
			"mydata": "value123",
		},
		ContentType: "application/json",
	}

	formatted, err := output.FormatOutputWithCustomData(event, &customData, "yaml")
	if err != nil {
		t.Fatalf("failed to format output with custom data: %v", err)
	}

	// Check that the formatted output contains the custom data
	if !strings.Contains(formatted, "customData:") {
		t.Errorf("formatted YAML output should contain customData field")
	}
	if !strings.Contains(formatted, "customDataContentType:") {
		t.Errorf("formatted YAML output should contain customDataContentType field")
	}
	if !strings.Contains(formatted, "mydata: value123") {
		t.Errorf("formatted YAML output should contain custom data values")
	}
}

func TestFormatCloudEventWithCustomData(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	// Set basic properties
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	customData := output.CustomData{
		Data: map[string]interface{}{
			"mydata": "value123",
		},
		ContentType: "application/json",
	}

	formatted, err := output.FormatOutputWithCustomData(event, &customData, "cloudevent")
	if err != nil {
		t.Fatalf("failed to format output with custom data: %v", err)
	}

	// Check that the formatted output contains the custom data
	if !strings.Contains(formatted, "\"customData\":") {
		t.Errorf("CloudEvent output should contain customData field")
	}
	if !strings.Contains(formatted, "\"customDataContentType\":") {
		t.Errorf("CloudEvent output should contain customDataContentType field")
	}
	if !strings.Contains(formatted, "\"mydata\": \"value123\"") {
		t.Errorf("CloudEvent output should contain custom data values")
	}
}

// Test cases without custom data
func TestFormatOutputWithoutCustomData(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	// Set basic properties
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// Test all formats without custom data
	formats := []string{"json", "yaml", "cloudevent"}
	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			formatted, err := output.FormatOutputWithCustomData(event, nil, format)
			if err != nil {
				t.Fatalf("failed to format output without custom data: %v", err)
			}

			if formatted == "" {
				t.Errorf("formatted output should not be empty")
			}

			// Should not contain custom data fields
			if strings.Contains(formatted, "customData") {
				t.Errorf("output should not contain customData when none provided")
			}
		})
	}
}

func TestFormatOutput(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	// Set basic properties
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	formatted, err := output.FormatOutput(event, "json")
	if err != nil {
		t.Fatalf("failed to format output: %v", err)
	}

	if formatted == "" {
		t.Errorf("formatted output should not be empty")
	}
}

func TestFormatOutputUnsupportedFormat(t *testing.T) {
	// Create a real CDEvent
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}

	_, err = output.FormatOutputWithCustomData(event, nil, "unsupported")
	if err == nil {
		t.Fatalf("expected error for unsupported format")
	}
}

func TestFormatMultipleEvents(t *testing.T) {
	// Create multiple test events
	event1, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event 1: %v", err)
	}
	event1.SetId("test-id-1")
	event1.SetSource("test-source")
	event1.SetSubjectId("pipeline-123")
	event1.SetSubjectPipelineName("test-pipeline")

	event2, err := cdeventsv04.NewPipelineRunStartedEvent()
	if err != nil {
		t.Fatalf("failed to create test event 2: %v", err)
	}
	event2.SetId("test-id-2")
	event2.SetSource("test-source")
	event2.SetSubjectId("pipeline-456")
	event2.SetSubjectPipelineName("test-pipeline-2")

	events := []api.CDEvent{event1, event2}

	// Test all formats
	formats := []string{"json", "yaml", "cloudevent"}
	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			formatted, err := output.FormatMultipleEvents(events, format)
			if err != nil {
				t.Fatalf("failed to format multiple events: %v", err)
			}

			if formatted == "" {
				t.Errorf("formatted output should not be empty")
			}
		})
	}
}

func TestFormatMultipleEventsUnsupportedFormat(t *testing.T) {
	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	events := []api.CDEvent{event}

	_, err = output.FormatMultipleEvents(events, "unsupported")
	if err == nil {
		t.Fatalf("expected error for unsupported format")
	}
}

