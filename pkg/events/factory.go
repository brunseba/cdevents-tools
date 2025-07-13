package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cdevents/sdk-go/pkg/api"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	"github.com/google/uuid"
)

// CustomData represents custom data that can be added to events
// This follows the CDEvents spec: https://github.com/cdevents/spec/blob/v0.4.1/spec.md#cdevents-custom-data
type CustomData struct {
	Data interface{} `json:"customData,omitempty"`
	ContentType string `json:"customDataContentType,omitempty"`
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
// Note: The current CDEvents SDK v0.4.1 doesn't support direct custom data injection,
// so we handle custom data in the output formatters instead.
func (ef *EventFactory) applyCustomData(event api.CDEvent, customData *CustomData) {
	// Custom data is now handled in the output formatters
	// This function is kept for future SDK versions that may support direct custom data
}

// ParseCustomDataFromJSON parses custom data from JSON string
func ParseCustomDataFromJSON(jsonData string) (*CustomData, error) {
	if jsonData == "" {
		return nil, nil
	}

	// Parse the JSON data into a generic interface{}
	var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return nil, fmt.Errorf("failed to parse custom data JSON: %w", err)
	}

	return &CustomData{
		Data: data,
		ContentType: "application/json",
	}, nil
}

