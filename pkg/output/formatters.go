package output

import (
	"encoding/json"
	"fmt"

	"github.com/cdevents/sdk-go/pkg/api"
	"gopkg.in/yaml.v3"
)

// FormatOutput formats the CDEvent based on the specified format
func FormatOutput(event api.CDEvent, format string) (string, error) {
	switch format {
	case "json":
		return formatJSON(event)
	case "yaml":
		return formatYAML(event)
	case "cloudevent":
		return formatCloudEvent(event)
	default:
		return "", fmt.Errorf("unsupported output format: %s", format)
	}
}

// formatJSON formats the event as JSON
func formatJSON(event api.CDEvent) (string, error) {
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal event to JSON: %w", err)
	}
	return string(data), nil
}

// formatYAML formats the event as YAML
func formatYAML(event api.CDEvent) (string, error) {
	data, err := yaml.Marshal(event)
	if err != nil {
		return "", fmt.Errorf("failed to marshal event to YAML: %w", err)
	}
	return string(data), nil
}

// formatCloudEvent formats the event as CloudEvent JSON
func formatCloudEvent(event api.CDEvent) (string, error) {
	ce, err := api.AsCloudEvent(event)
	if err != nil {
		return "", fmt.Errorf("failed to convert to CloudEvent: %w", err)
	}
	
	data, err := json.MarshalIndent(ce, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal CloudEvent to JSON: %w", err)
	}
	return string(data), nil
}

// FormatMultipleEvents formats multiple events
func FormatMultipleEvents(events []api.CDEvent, format string) (string, error) {
	switch format {
	case "json":
		return formatMultipleJSON(events)
	case "yaml":
		return formatMultipleYAML(events)
	case "cloudevent":
		return formatMultipleCloudEvents(events)
	default:
		return "", fmt.Errorf("unsupported output format: %s", format)
	}
}

// formatMultipleJSON formats multiple events as JSON array
func formatMultipleJSON(events []api.CDEvent) (string, error) {
	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal events to JSON: %w", err)
	}
	return string(data), nil
}

// formatMultipleYAML formats multiple events as YAML array
func formatMultipleYAML(events []api.CDEvent) (string, error) {
	data, err := yaml.Marshal(events)
	if err != nil {
		return "", fmt.Errorf("failed to marshal events to YAML: %w", err)
	}
	return string(data), nil
}

// formatMultipleCloudEvents formats multiple events as CloudEvents JSON array
func formatMultipleCloudEvents(events []api.CDEvent) (string, error) {
	var cloudEvents []interface{}
	
	for _, event := range events {
		ce, err := api.AsCloudEvent(event)
		if err != nil {
			return "", fmt.Errorf("failed to convert event to CloudEvent: %w", err)
		}
		cloudEvents = append(cloudEvents, ce)
	}
	
	data, err := json.MarshalIndent(cloudEvents, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal CloudEvents to JSON: %w", err)
	}
	return string(data), nil
}
