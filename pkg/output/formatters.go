package output

import (
	"encoding/json"
	"fmt"

	"github.com/cdevents/sdk-go/pkg/api"
	"gopkg.in/yaml.v3"
)

// CustomData represents custom data that can be added to events
// This follows the CDEvents spec: https://github.com/cdevents/spec/blob/v0.4.1/spec.md#cdevents-custom-data
type CustomData struct {
	Data interface{} `json:"customData,omitempty"`
	ContentType string `json:"customDataContentType,omitempty"`
}

// FormatOutput formats the CDEvent based on the specified format
func FormatOutput(event api.CDEvent, format string) (string, error) {
	return FormatOutputWithCustomData(event, nil, format)
}

// FormatOutputWithCustomData formats the CDEvent with custom data based on the specified format
func FormatOutputWithCustomData(event api.CDEvent, customData *CustomData, format string) (string, error) {
	switch format {
	case "json":
		return formatJSONWithCustomData(event, customData)
	case "yaml":
		return formatYAMLWithCustomData(event, customData)
	case "cloudevent":
		return formatCloudEventWithCustomData(event, customData)
	default:
		return "", fmt.Errorf("unsupported output format: %s", format)
	}
}

// formatJSON formats the event as JSON
func formatJSON(event api.CDEvent) (string, error) {
	return formatJSONWithCustomData(event, nil)
}

// formatJSONWithCustomData formats the event as JSON with custom data
func formatJSONWithCustomData(event api.CDEvent, customData *CustomData) (string, error) {
	// Marshal the event to get its JSON representation
	eventData, err := json.Marshal(event)
	if err != nil {
		return "", fmt.Errorf("failed to marshal event: %w", err)
	}

	// Parse the event JSON to a map
	var eventMap map[string]interface{}
	if err := json.Unmarshal(eventData, &eventMap); err != nil {
		return "", fmt.Errorf("failed to unmarshal event: %w", err)
	}

	// Add custom data according to CDEvents spec at the root level
	if customData != nil {
		if customData.Data != nil {
			eventMap["customData"] = customData.Data
		}
		if customData.ContentType != "" {
			eventMap["customDataContentType"] = customData.ContentType
		}
	}

	// Marshal back to JSON with custom data
	data, err := json.MarshalIndent(eventMap, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal event with custom data to JSON: %w", err)
	}
	return string(data), nil
}

// formatYAML formats the event as YAML
func formatYAML(event api.CDEvent) (string, error) {
	return formatYAMLWithCustomData(event, nil)
}

// formatYAMLWithCustomData formats the event as YAML with custom data
func formatYAMLWithCustomData(event api.CDEvent, customData *CustomData) (string, error) {
	if customData == nil {
		data, err := yaml.Marshal(event)
		if err != nil {
			return "", fmt.Errorf("failed to marshal event to YAML: %w", err)
		}
		return string(data), nil
	}

	// Similar to JSON, but for YAML
	eventData, err := json.Marshal(event)
	if err != nil {
		return "", fmt.Errorf("failed to marshal event: %w", err)
	}

	var eventMap map[string]interface{}
	if err := json.Unmarshal(eventData, &eventMap); err != nil {
		return "", fmt.Errorf("failed to unmarshal event: %w", err)
	}

	// Add custom data according to CDEvents spec at the root level
	if customData.Data != nil {
		eventMap["customData"] = customData.Data
	}
	if customData.ContentType != "" {
		eventMap["customDataContentType"] = customData.ContentType
	}

	data, err := yaml.Marshal(eventMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal event with custom data to YAML: %w", err)
	}
	return string(data), nil
}

// formatCloudEvent formats the event as CloudEvent JSON
func formatCloudEvent(event api.CDEvent) (string, error) {
	return formatCloudEventWithCustomData(event, nil)
}

// formatCloudEventWithCustomData formats the event as CloudEvent JSON with custom data
func formatCloudEventWithCustomData(event api.CDEvent, customData *CustomData) (string, error) {
	ce, err := api.AsCloudEvent(event)
	if err != nil {
		return "", fmt.Errorf("failed to convert to CloudEvent: %w", err)
	}

	if customData != nil {
		// Add custom data to the CloudEvent data field
		ceData, err := json.Marshal(ce)
		if err != nil {
			return "", fmt.Errorf("failed to marshal CloudEvent: %w", err)
		}

		var ceMap map[string]interface{}
		if err := json.Unmarshal(ceData, &ceMap); err != nil {
			return "", fmt.Errorf("failed to unmarshal CloudEvent: %w", err)
		}

		// Add custom data to the CloudEvent data according to CDEvents spec
		if data, ok := ceMap["data"].(map[string]interface{}); ok {
			if customData.Data != nil {
				data["customData"] = customData.Data
			}
			if customData.ContentType != "" {
				data["customDataContentType"] = customData.ContentType
			}
		}

		data, err := json.MarshalIndent(ceMap, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to marshal CloudEvent with custom data to JSON: %w", err)
		}
		return string(data), nil
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
