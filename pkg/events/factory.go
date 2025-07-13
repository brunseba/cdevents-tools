package events

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cdevents/sdk-go/pkg/api"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

// CustomData represents custom data that can be added to events
type CustomData struct {
	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`
	Links []CustomLink `json:"links,omitempty" yaml:"links,omitempty"`
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

// CustomLink represents a custom link that can be added to events
type CustomLink struct {
	Name string `json:"name"`
	URL string `json:"url"`
	Type string `json:"type,omitempty"`
}

// EventFactory creates CDEvents with common functionality
type EventFactory struct {
	defaultSource string
}

// NewEventFactory creates a new EventFactory
func NewEventFactory(defaultSource string) *EventFactory {
	return &EventFactory{
		defaultSource: defaultSource,
	}
}

// CreatePipelineRunEvent creates a pipeline run event
func (ef *EventFactory) CreatePipelineRunEvent(eventType, pipelineID, pipelineName, outcome, errors, url string, customData *CustomData) (api.CDEvent, error) {
	var event api.CDEvent
	var err error

	switch eventType {
	case "queued":
		event, err = cdeventsv04.NewPipelineRunQueuedEvent()
	case "started":
		event, err = cdeventsv04.NewPipelineRunStartedEvent()
	case "finished":
		event, err = cdeventsv04.NewPipelineRunFinishedEvent()
	default:
		return nil, fmt.Errorf("unsupported pipeline run event type: %s", eventType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline run event: %w", err)
	}

	// Set common fields
	event.SetId(uuid.New().String())
	event.SetSource(ef.defaultSource)
	event.SetTimestamp(time.Now())
	event.SetSubjectId(pipelineID)
	
	// Set pipeline-specific fields
	if pipelineRunEvent, ok := event.(interface {
		SetSubjectPipelineName(string)
		SetSubjectUrl(string)
	}); ok {
		pipelineRunEvent.SetSubjectPipelineName(pipelineName)
		if url != "" {
			pipelineRunEvent.SetSubjectUrl(url)
		}
	}

	// Set outcome and errors for finished events
	if eventType == "finished" {
		if finishedEvent, ok := event.(interface {
			SetSubjectOutcome(string)
			SetSubjectErrors(string)
		}); ok {
			if outcome != "" {
				finishedEvent.SetSubjectOutcome(outcome)
			}
			if errors != "" {
				finishedEvent.SetSubjectErrors(errors)
			}
		}
	}

	// Apply custom data if provided
	if customData != nil {
		ef.applyCustomData(event, customData)
	}

	return event, nil
}

// CreateTaskRunEvent creates a task run event
func (ef *EventFactory) CreateTaskRunEvent(eventType, taskID, taskName, pipelineRunID, outcome, errors, url string, customData *CustomData) (api.CDEvent, error) {
	var event api.CDEvent
	var err error

	switch eventType {
	case "started":
		event, err = cdeventsv04.NewTaskRunStartedEvent()
	case "finished":
		event, err = cdeventsv04.NewTaskRunFinishedEvent()
	default:
		return nil, fmt.Errorf("unsupported task run event type: %s", eventType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create task run event: %w", err)
	}

	// Set common fields
	event.SetId(uuid.New().String())
	event.SetSource(ef.defaultSource)
	event.SetTimestamp(time.Now())
	event.SetSubjectId(taskID)

	// Set task-specific fields
	if taskRunEvent, ok := event.(interface {
		SetSubjectTaskName(string)
		SetSubjectUrl(string)
		SetSubjectPipelineRun(map[string]interface{})
	}); ok {
		taskRunEvent.SetSubjectTaskName(taskName)
		if url != "" {
			taskRunEvent.SetSubjectUrl(url)
		}
		if pipelineRunID != "" {
			taskRunEvent.SetSubjectPipelineRun(map[string]interface{}{
				"id": pipelineRunID,
			})
		}
	}

	// Set outcome and errors for finished events
	if eventType == "finished" {
		if finishedEvent, ok := event.(interface {
			SetSubjectOutcome(string)
			SetSubjectErrors(string)
		}); ok {
			if outcome != "" {
				finishedEvent.SetSubjectOutcome(outcome)
			}
			if errors != "" {
				finishedEvent.SetSubjectErrors(errors)
			}
		}
	}

	// Apply custom data if provided
	if customData != nil {
		ef.applyCustomData(event, customData)
	}

	return event, nil
}

// CreateBuildEvent creates a build event
func (ef *EventFactory) CreateBuildEvent(eventType, buildID, buildName, outcome, errors, url string, customData *CustomData) (api.CDEvent, error) {
	var event api.CDEvent
	var err error

	switch eventType {
	case "queued":
		event, err = cdeventsv04.NewBuildQueuedEvent()
	case "started":
		event, err = cdeventsv04.NewBuildStartedEvent()
	case "finished":
		event, err = cdeventsv04.NewBuildFinishedEvent()
	default:
		return nil, fmt.Errorf("unsupported build event type: %s", eventType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create build event: %w", err)
	}

	// Set common fields
	event.SetId(uuid.New().String())
	event.SetSource(ef.defaultSource)
	event.SetTimestamp(time.Now())
	event.SetSubjectId(buildID)

	// Set build-specific fields  
	if buildEvent, ok := event.(interface {
		SetSubjectBuildName(string)
		SetSubjectUrl(string)
	}); ok {
		buildEvent.SetSubjectBuildName(buildName)
		if url != "" {
			buildEvent.SetSubjectUrl(url)
		}
	}

	// Set outcome and errors for finished events
	if eventType == "finished" {
		if finishedEvent, ok := event.(interface {
			SetSubjectOutcome(string)
			SetSubjectErrors(string)
		}); ok {
			if outcome != "" {
				finishedEvent.SetSubjectOutcome(outcome)
			}
			if errors != "" {
				finishedEvent.SetSubjectErrors(errors)
			}
		}
	}

	// Apply custom data if provided
	if customData != nil {
		ef.applyCustomData(event, customData)
	}

	return event, nil
}

// CreateServiceEvent creates a service deployment event
func (ef *EventFactory) CreateServiceEvent(eventType, serviceID, serviceName, environmentID, url string, customData *CustomData) (api.CDEvent, error) {
	var event api.CDEvent
	var err error

	switch eventType {
	case "deployed":
		event, err = cdeventsv04.NewServiceDeployedEvent()
	case "published":
		event, err = cdeventsv04.NewServicePublishedEvent()
	case "removed":
		event, err = cdeventsv04.NewServiceRemovedEvent()
	case "rolledback":
		event, err = cdeventsv04.NewServiceRolledbackEvent()
	case "upgraded":
		event, err = cdeventsv04.NewServiceUpgradedEvent()
	default:
		return nil, fmt.Errorf("unsupported service event type: %s", eventType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create service event: %w", err)
	}

	// Set common fields
	event.SetId(uuid.New().String())
	event.SetSource(ef.defaultSource)
	event.SetTimestamp(time.Now())
	event.SetSubjectId(serviceID)

	// Set service-specific fields
	if serviceEvent, ok := event.(interface {
		SetSubjectServiceName(string)
		SetSubjectUrl(string)
		SetSubjectEnvironment(map[string]interface{})
	}); ok {
		serviceEvent.SetSubjectServiceName(serviceName)
		if url != "" {
			serviceEvent.SetSubjectUrl(url)
		}
		if environmentID != "" {
			serviceEvent.SetSubjectEnvironment(map[string]interface{}{
				"id": environmentID,
			})
		}
	}

	// Apply custom data if provided
	if customData != nil {
		ef.applyCustomData(event, customData)
	}

	return event, nil
}

// CreateTestEvent creates a test event
func (ef *EventFactory) CreateTestEvent(eventType, testID, testName, outcome, errors, url string, customData *CustomData) (api.CDEvent, error) {
	var event api.CDEvent
	var err error

	switch eventType {
	case "testcase-queued":
		event, err = cdeventsv04.NewTestCaseRunQueuedEvent()
	case "testcase-started":
		event, err = cdeventsv04.NewTestCaseRunStartedEvent()
	case "testcase-finished":
		event, err = cdeventsv04.NewTestCaseRunFinishedEvent()
	case "testcase-skipped":
		event, err = cdeventsv04.NewTestCaseRunSkippedEvent()
	case "testsuite-queued":
		event, err = cdeventsv04.NewTestSuiteRunQueuedEvent()
	case "testsuite-started":
		event, err = cdeventsv04.NewTestSuiteRunStartedEvent()
	case "testsuite-finished":
		event, err = cdeventsv04.NewTestSuiteRunFinishedEvent()
	case "testoutput-published":
		event, err = cdeventsv04.NewTestOutputPublishedEvent()
	default:
		return nil, fmt.Errorf("unsupported test event type: %s", eventType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create test event: %w", err)
	}

	// Set common fields
	event.SetId(uuid.New().String())
	event.SetSource(ef.defaultSource)
	event.SetTimestamp(time.Now())
	event.SetSubjectId(testID)

	// Set test-specific fields based on event type
	switch eventType {
	case "testcase-queued", "testcase-started", "testcase-finished", "testcase-skipped":
		if testEvent, ok := event.(interface {
			SetSubjectTestCaseName(string)
			SetSubjectUrl(string)
		}); ok {
			testEvent.SetSubjectTestCaseName(testName)
			if url != "" {
				testEvent.SetSubjectUrl(url)
			}
		}
	case "testsuite-queued", "testsuite-started", "testsuite-finished":
		if testEvent, ok := event.(interface {
			SetSubjectTestSuiteName(string)
			SetSubjectUrl(string)
		}); ok {
			testEvent.SetSubjectTestSuiteName(testName)
			if url != "" {
				testEvent.SetSubjectUrl(url)
			}
		}
	}

	// Set outcome and errors for finished events
	if eventType == "testcase-finished" || eventType == "testsuite-finished" {
		if finishedEvent, ok := event.(interface {
			SetSubjectOutcome(string)
			SetSubjectErrors(string)
		}); ok {
			if outcome != "" {
				finishedEvent.SetSubjectOutcome(outcome)
			}
			if errors != "" {
				finishedEvent.SetSubjectErrors(errors)
			}
		}
	}

	// Apply custom data if provided
	if customData != nil {
		ef.applyCustomData(event, customData)
	}

	return event, nil
}

// applyCustomData applies custom data to a CDEvent
func (ef *EventFactory) applyCustomData(event api.CDEvent, customData *CustomData) {
	if customData == nil {
		return
	}

	// Apply custom data to the event's custom data extension
	if customData.Data != nil {
		if customEvent, ok := event.(interface {
			SetCustomData(map[string]interface{})
		}); ok {
			customEvent.SetCustomData(customData.Data)
		}
	}

	// Apply labels and annotations as custom data
	if len(customData.Labels) > 0 || len(customData.Annotations) > 0 {
		if customEvent, ok := event.(interface {
			SetCustomDataEntry(string, interface{})
		}); ok {
			if len(customData.Labels) > 0 {
				customEvent.SetCustomDataEntry("labels", customData.Labels)
			}
			if len(customData.Annotations) > 0 {
				customEvent.SetCustomDataEntry("annotations", customData.Annotations)
			}
		}
	}

	// Apply custom links
	if len(customData.Links) > 0 {
		if customEvent, ok := event.(interface {
			SetCustomDataEntry(string, interface{})
		}); ok {
			customEvent.SetCustomDataEntry("links", customData.Links)
		}
	}
}

// ParseCustomDataFromJSON parses custom data from JSON string
func ParseCustomDataFromJSON(jsonData string) (*CustomData, error) {
	if jsonData == "" {
		return nil, nil
	}

	var customData CustomData
	if err := json.Unmarshal([]byte(jsonData), &customData); err != nil {
		return nil, fmt.Errorf("failed to parse custom data JSON: %w", err)
	}

	return &customData, nil
}

// ParseCustomDataFromYAML parses custom data from YAML string
func ParseCustomDataFromYAML(yamlData string) (*CustomData, error) {
	if yamlData == "" {
		return nil, nil
	}

	var customData CustomData
	if err := yaml.Unmarshal([]byte(yamlData), &customData); err != nil {
		return nil, fmt.Errorf("failed to parse custom data YAML: %w", err)
	}

	return &customData, nil
}

// ParseCustomDataFromKeyValue parses custom data from key=value pairs
func ParseCustomDataFromKeyValue(keyValues []string) (*CustomData, error) {
	if len(keyValues) == 0 {
		return nil, nil
	}

	customData := &CustomData{
		Data: make(map[string]interface{}),
	}

	for _, kv := range keyValues {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid key=value format: %s", kv)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Try to parse as JSON first
		var jsonValue interface{}
		if err := json.Unmarshal([]byte(value), &jsonValue); err == nil {
			customData.Data[key] = jsonValue
		} else {
			// Fallback to string value
			customData.Data[key] = value
		}
	}

	return customData, nil
}

// ParseLabelsFromKeyValue parses labels from key=value pairs
func ParseLabelsFromKeyValue(keyValues []string) (map[string]string, error) {
	if len(keyValues) == 0 {
		return nil, nil
	}

	labels := make(map[string]string)
	for _, kv := range keyValues {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid key=value format: %s", kv)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		labels[key] = value
	}

	return labels, nil
}
