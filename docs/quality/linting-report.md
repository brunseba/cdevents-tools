# Linting Report

> Last updated: Sun Jul 13 21:43:35 UTC 2025

## Linting Status

**Status**: ⚠️ ISSUES FOUND

## Issues Found

```
pkg/transport/transport.go:46:22: unused-parameter: parameter 'headers' seems to be unused, consider removing or renaming it as _ (revive)
func WithHTTPHeaders(headers map[string]string) HTTPOption {
                     ^
pkg/transport/transport.go:83:33: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (t *ConsoleTransport) Send(ctx context.Context, event api.CDEvent) error {
                                ^
pkg/transport/transport.go:105:30: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (t *FileTransport) Send(ctx context.Context, event api.CDEvent) error {
                             ^
pkg/transport/transport.go:119:24: unused-parameter: parameter 'brokers' seems to be unused, consider removing or renaming it as _ (revive)
func NewKafkaTransport(brokers []string, topic string) (*KafkaTransport, error) {
                       ^
pkg/transport/transport.go:125:31: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (t *KafkaTransport) Send(ctx context.Context, event api.CDEvent) error {
                              ^
pkg/transport/transport.go:130:6: exported: type name will be used as transport.TransportFactory by other packages, and that stutters; consider calling this Factory (revive)
type TransportFactory struct{}
     ^
pkg/transport/transport.go:113:2: field `brokers` is unused (unused)
	brokers []string
	^
pkg/transport/transport.go:114:2: field `topic` is unused (unused)
	topic   string
	^
pkg/transport/transport.go:115:2: field `client` is unused (unused)
	client  cloudevents.Client
	^
pkg/transport/transport.go:8:2: import 'github.com/cdevents/sdk-go/pkg/api' is not allowed from list 'Main' (depguard)
	"github.com/cdevents/sdk-go/pkg/api"
	^
pkg/transport/transport.go:9:2: import 'github.com/cloudevents/sdk-go/v2' is not allowed from list 'Main' (depguard)
	cloudevents "github.com/cloudevents/sdk-go/v2"
	^
pkg/transport/transport.go:121:14: ST1005: error strings should not be capitalized (stylecheck)
	return nil, fmt.Errorf("Kafka transport not implemented yet")
	            ^
pkg/transport/transport.go:126:9: ST1005: error strings should not be capitalized (stylecheck)
	return fmt.Errorf("Kafka transport not implemented yet")
	       ^
pkg/transport/transport.go:153:15: ST1005: error strings should not be capitalized (stylecheck)
		return nil, fmt.Errorf("Kafka transport not implemented yet")
		            ^
cmd/generate_build.go:1: 1-45 lines are duplicate of `cmd/generate_test.go:1-45` (dupl)
package cmd

import (
	"fmt"

	"github.com/brunseba/cdevents-tools/pkg/events"
	"github.com/spf13/cobra"
)

var generateBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate build events",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := events.NewEventFactory(getDefaultSource())
		eventType := args[0]

		// Parse custom data
		customData, err := parseCustomData(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse custom data: %w", err)
		}

		event, err := factory.CreateBuildEvent(
			eventType,
			cmd.Flag("id").Value.String(),
			cmd.Flag("name").Value.String(),
			cmd.Flag("outcome").Value.String(),
			cmd.Flag("errors").Value.String(),
			cmd.Flag("url").Value.String(),
			customData,
		)
		if err != nil {
			return fmt.Errorf("failed to create build event: %w", err)
		}

		format := cmd.Flag("output").Value.String()
		return outputEvent(event, format)
	},
}

func init() {
	addCommonGenerateFlags(generateBuildCmd)
	generateCmd.AddCommand(generateBuildCmd)
}
cmd/generate_test.go:1: 1-45 lines are duplicate of `cmd/generate_build.go:1-45` (dupl)
package cmd

import (
	"fmt"

	"github.com/brunseba/cdevents-tools/pkg/events"
	"github.com/spf13/cobra"
)

var generateTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Generate test events",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := events.NewEventFactory(getDefaultSource())
		eventType := args[0]

		// Parse custom data
		customData, err := parseCustomData(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse custom data: %w", err)
		}

		event, err := factory.CreateTestEvent(
			eventType,
			cmd.Flag("id").Value.String(),
			cmd.Flag("name").Value.String(),
			cmd.Flag("outcome").Value.String(),
			cmd.Flag("errors").Value.String(),
			cmd.Flag("url").Value.String(),
			customData,
		)
		if err != nil {
			return fmt.Errorf("failed to create test event: %w", err)
		}

		format := cmd.Flag("output").Value.String()
		return outputEvent(event, format)
	},
}

func init() {
	addCommonGenerateFlags(generateTestCmd)
	generateCmd.AddCommand(generateTestCmd)
}
cmd/generate.go:54:22: Error return value of `cmd.MarkFlagRequired` is not checked (errcheck)
	cmd.MarkFlagRequired("id")
	                    ^
cmd/generate.go:55:22: Error return value of `cmd.MarkFlagRequired` is not checked (errcheck)
	cmd.MarkFlagRequired("name")
	                    ^
cmd/root.go:45:17: Error return value of `viper.BindPFlag` is not checked (errcheck)
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	               ^
cmd/root.go:46:17: Error return value of `viper.BindPFlag` is not checked (errcheck)
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	               ^
cmd/send.go:42:17: Error return value of `viper.BindPFlag` is not checked (errcheck)
	viper.BindPFlag("target", sendCmd.PersistentFlags().Lookup("target"))
	               ^
cmd/send.go:43:17: Error return value of `viper.BindPFlag` is not checked (errcheck)
	viper.BindPFlag("retries", sendCmd.PersistentFlags().Lookup("retries"))
	               ^
cmd/send.go:44:17: Error return value of `viper.BindPFlag` is not checked (errcheck)
	viper.BindPFlag("timeout", sendCmd.PersistentFlags().Lookup("timeout"))
	               ^
cmd/send.go:45:17: Error return value of `viper.BindPFlag` is not checked (errcheck)
	viper.BindPFlag("headers", sendCmd.PersistentFlags().Lookup("headers"))
	               ^
cmd/generate.go:41:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/generate_build.go:42:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/generate_pipeline.go:42:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/generate_service.go:41:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/generate_task.go:43:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/root.go:36:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/send.go:32:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/send_pipeline.go:45:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/generate_test.go:42:1: don't use `init` function (gochecknoinits)
func init() {
^
cmd/send.go:56:2: importShadow: shadow of imported from 'github.com/brunseba/cdevents-tools/pkg/transport' package 'transport' (gocritic)
	transport, err := factory.CreateTransport(target)
	^
cmd/send.go:72:46: importShadow: shadow of imported from 'github.com/brunseba/cdevents-tools/pkg/transport' package 'transport' (gocritic)
func SendEventWithRetry(ctx context.Context, transport transport.Transport, event api.CDEvent, maxRetries int) error {
                                             ^
cmd/generate.go:53: File is not `gofmt`-ed with `-s` (gofmt)
	
cmd/root.go:43: File is not `gofmt`-ed with `-s` (gofmt)
	
cmd/send.go:34: File is not `gofmt`-ed with `-s` (gofmt)
	
cmd/generate.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)

cmd/generate_build.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
cmd/generate_pipeline.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
cmd/generate_service.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
cmd/generate_task.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
cmd/send.go:8: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/transport"
cmd/send_pipeline.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
cmd/generate_test.go:6: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
cmd/send.go:36: line is 122 characters (lll)
	sendCmd.PersistentFlags().StringP("target", "t", "console", "Target to send events to (console, http://..., file://...)")
cmd/generate.go:7:2: import 'github.com/brunseba/cdevents-tools/pkg/events' is not allowed from list 'Main' (depguard)
	"github.com/brunseba/cdevents-tools/pkg/events"
	^
cmd/generate.go:8:2: import 'github.com/brunseba/cdevents-tools/pkg/output' is not allowed from list 'Main' (depguard)
	"github.com/brunseba/cdevents-tools/pkg/output"
	^
cmd/generate.go:11:2: import 'github.com/spf13/cobra' is not allowed from list 'Main' (depguard)
	"github.com/spf13/cobra"
	^
cmd/generate.go:12:2: import 'github.com/spf13/viper' is not allowed from list 'Main' (depguard)
	"github.com/spf13/viper"
	^
cmd/root.go:7:2: import 'github.com/spf13/cobra' is not allowed from list 'Main' (depguard)
	"github.com/spf13/cobra"
	^
cmd/root.go:8:2: import 'github.com/spf13/viper' is not allowed from list 'Main' (depguard)
	"github.com/spf13/viper"
	^
cmd/send.go:9:2: import 'github.com/cdevents/sdk-go/pkg/api' is not allowed from list 'Main' (depguard)
	"github.com/cdevents/sdk-go/pkg/api"
	^
cmd/send.go:10:2: import 'github.com/spf13/cobra' is not allowed from list 'Main' (depguard)
	"github.com/spf13/cobra"
	^
cmd/send_pipeline.go:7:2: import 'github.com/spf13/cobra' is not allowed from list 'Main' (depguard)
	"github.com/spf13/cobra"
	^
pkg/transport/transport_test.go:13:2: importShadow: shadow of imported from 'github.com/brunseba/cdevents-tools/pkg/transport' package 'transport' (gocritic)
	transport, err := transport.NewHTTPTransport("http://example.com")
	^
pkg/transport/transport_test.go:69:4: importShadow: shadow of imported from 'github.com/brunseba/cdevents-tools/pkg/transport' package 'transport' (gocritic)
			transport, err := factory.CreateTransport(tc.target)
			^
pkg/transport/transport_test.go:6: File is not `gofmt`-ed with `-s` (gofmt)
	"testing"
	"github.com/cdevents/sdk-go/pkg/api"
pkg/transport/transport_test.go:204:33: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (t *failingTransport) Send(ctx context.Context, event api.CDEvent) error {
                                ^
pkg/transport/transport_test.go:7:2: import 'github.com/cdevents/sdk-go/pkg/api' is not allowed from list 'Main' (depguard)
	"github.com/cdevents/sdk-go/pkg/api"
	^
pkg/events/factory.go:33: 33-90 lines are duplicate of `pkg/events/factory.go:157-214` (dupl)
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
pkg/events/factory.go:157: 157-214 lines are duplicate of `pkg/events/factory.go:33-90` (dupl)
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
pkg/events/factory.go:272: Function 'CreateTestEvent' has too many statements (47 > 40) (funlen)
func (ef *EventFactory) CreateTestEvent(eventType, testID, testName, outcome, errors, url string, customData *CustomData) (api.CDEvent, error) {
pkg/events/factory.go:40:7: string `started` has 3 occurrences, make it a constant (goconst)
	case "started":
	     ^
pkg/events/factory.go:289:7: string `testsuite-finished` has 3 occurrences, make it a constant (goconst)
	case "testsuite-finished":
	     ^
pkg/events/factory.go:281:7: string `testcase-finished` has 3 occurrences, make it a constant (goconst)
	case "testcase-finished":
	     ^
pkg/events/factory.go:42:7: string `finished` has 6 occurrences, make it a constant (goconst)
	case "finished":
	     ^
pkg/events/factory.go:93:1: cyclomatic complexity 12 of func `(*EventFactory).CreateTaskRunEvent` is high (> 10) (gocyclo)
func (ef *EventFactory) CreateTaskRunEvent(eventType, taskID, taskName, pipelineRunID, outcome, errors, url string, customData *CustomData) (api.CDEvent, error) {
^
pkg/events/factory.go:217:1: cyclomatic complexity 11 of func `(*EventFactory).CreateServiceEvent` is high (> 10) (gocyclo)
func (ef *EventFactory) CreateServiceEvent(eventType, serviceID, serviceName, environmentID, url string, customData *CustomData) (api.CDEvent, error) {
^
pkg/events/factory.go:16: File is not `gofmt`-ed with `-s` (gofmt)
	Data interface{} `json:"customData,omitempty"`
	ContentType string `json:"customDataContentType,omitempty"`
pkg/events/factory.go:357:41: unused-parameter: parameter 'event' seems to be unused, consider removing or renaming it as _ (revive)
func (ef *EventFactory) applyCustomData(event api.CDEvent, customData *CustomData) {
                                        ^
pkg/events/factory.go:8:2: import 'github.com/cdevents/sdk-go/pkg/api' is not allowed from list 'Main' (depguard)
	"github.com/cdevents/sdk-go/pkg/api"
	^
pkg/events/factory.go:9:2: import 'github.com/cdevents/sdk-go/pkg/api/v04' is not allowed from list 'Main' (depguard)
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
	^
pkg/events/factory.go:10:2: import 'github.com/google/uuid' is not allowed from list 'Main' (depguard)
	"github.com/google/uuid"
	^
pkg/output/formatters.go:14: File is not `gofmt`-ed with `-s` (gofmt)
	Data interface{} `json:"customData,omitempty"`
	ContentType string `json:"customDataContentType,omitempty"`
pkg/output/formatters.go:38:6: func `formatJSON` is unused (unused)
func formatJSON(event api.CDEvent) (string, error) {
     ^
pkg/output/formatters.go:75:6: func `formatYAML` is unused (unused)
func formatYAML(event api.CDEvent) (string, error) {
     ^
pkg/output/formatters.go:116:6: func `formatCloudEvent` is unused (unused)
func formatCloudEvent(event api.CDEvent) (string, error) {
     ^
pkg/output/formatters.go:7:2: import 'github.com/cdevents/sdk-go/pkg/api' is not allowed from list 'Main' (depguard)
	"github.com/cdevents/sdk-go/pkg/api"
	^
pkg/output/formatters.go:52:5: shadow: declaration of "err" shadows declaration at line 45 (govet)
	if err := json.Unmarshal(eventData, &eventMap); err != nil {
	   ^
pkg/output/formatters.go:96:5: shadow: declaration of "err" shadows declaration at line 90 (govet)
	if err := json.Unmarshal(eventData, &eventMap); err != nil {
	   ^
pkg/output/formatters.go:129:11: shadow: declaration of "err" shadows declaration at line 122 (govet)
		ceData, err := json.Marshal(ce)
		        ^
pkg/output/formatters.go:135:6: shadow: declaration of "err" shadows declaration at line 129 (govet)
		if err := json.Unmarshal(ceData, &ceMap); err != nil {
		   ^
main.go:7:2: import 'github.com/brunseba/cdevents-tools/cmd' is not allowed from list 'Main' (depguard)
	"github.com/brunseba/cdevents-tools/cmd"
	^
cmd/cmd_test.go:20: File is not `gofmt`-ed with `-s` (gofmt)
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
cmd/integration_test.go:14: File is not `gofmt`-ed with `-s` (gofmt)
	
cmd/cmd_test.go:8: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)

cmd/cmd_test.go:77: line is 136 characters (lll)
			args: []string{"cdevents-cli", "generate", "pipeline", "finished", "--id", "123", "--name", "test-pipeline", "--outcome", "success"},
cmd/cmd_test.go:87: line is 142 characters (lll)
			args: []string{"cdevents-cli", "generate", "build", "finished", "--id", "456", "--name", "test-build", "--custom-json", `{"key":"value"}`},
cmd/cmd_test.go:107: line is 131 characters (lll)
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--output", "yaml"},
cmd/cmd_test.go:112: line is 137 characters (lll)
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--output", "cloudevent"},
cmd/cmd_test.go:132: line is 146 characters (lll)
			args: []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--custom-json", `{"invalid json`},
cmd/cmd_test.go:163: line is 130 characters (lll)
			args: []string{"cdevents-cli", "send", "--target", "console", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
cmd/cmd_test.go:168: line is 139 characters (lll)
			args: []string{"cdevents-cli", "send", "--target", "file://test.json", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
cmd/cmd_test.go:173: line is 148 characters (lll)
			args: []string{"cdevents-cli", "send", "--target", "console", "--retries", "5", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
cmd/cmd_test.go:178: line is 150 characters (lll)
			args: []string{"cdevents-cli", "send", "--target", "console", "--timeout", "10s", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
cmd/cmd_test.go:183: line is 173 characters (lll)
			args: []string{"cdevents-cli", "send", "--target", "console", "--headers", "Authorization=Bearer token", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
cmd/cmd_test.go:188: line is 139 characters (lll)
			args: []string{"cdevents-cli", "send", "--target", "invalid://target", "pipeline", "started", "--id", "123", "--name", "test-pipeline"},
cmd/cmd_test.go:302: line is 148 characters (lll)
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--custom-json", `{"key":"value","num":42}`}
cmd/cmd_test.go:309: line is 139 characters (lll)
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--custom-json", `{"invalid":json`}
cmd/cmd_test.go:324: line is 134 characters (lll)
			os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test-pipeline", "--output", format}
cmd/cmd_test.go:357: line is 124 characters (lll)
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--custom-json", ""}
cmd/cmd_test.go:364: line is 135 characters (lll)
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "started", "--id", "123", "--name", "test", "--url", "https://example.com"}
cmd/cmd_test.go:371: line is 156 characters (lll)
	os.Args = []string{"cdevents-cli", "generate", "pipeline", "finished", "--id", "123", "--name", "test", "--outcome", "failure", "--errors", "Build failed"}
cmd/integration_test.go:45: line is 131 characters (lll)
	os.Args = []string{"cdevents-cli", "send", "--target", "console", "pipeline", "started", "--id", "123", "--name", "test-pipeline"}
cmd/cmd_test.go:247:30: unused-parameter: parameter 'ctx' seems to be unused, consider removing or renaming it as _ (revive)
func (m *MockTransport) Send(ctx context.Context, event api.CDEvent) error {
                             ^
cmd/cmd_test.go:9:2: import 'github.com/brunseba/cdevents-tools/cmd' is not allowed from list 'Main' (depguard)
	"github.com/brunseba/cdevents-tools/cmd"
	^
cmd/cmd_test.go:10:2: import 'github.com/brunseba/cdevents-tools/pkg/events' is not allowed from list 'Main' (depguard)
	"github.com/brunseba/cdevents-tools/pkg/events"
	^
cmd/integration_test.go:7:2: import 'github.com/brunseba/cdevents-tools/cmd' is not allowed from list 'Main' (depguard)
	"github.com/brunseba/cdevents-tools/cmd"
	^
pkg/events/factory_test.go:269: 269-319 lines are duplicate of `pkg/events/factory_test.go:373-423` (dupl)
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
pkg/events/factory_test.go:373: 373-423 lines are duplicate of `pkg/events/factory_test.go:269-319` (dupl)
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
pkg/events/factory_test.go:283: 283-318 lines are duplicate of `pkg/events/factory_test.go:387-422` (dupl)
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
pkg/events/factory_test.go:387: 387-422 lines are duplicate of `pkg/events/factory_test.go:526-561` (dupl)
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
pkg/events/factory_test.go:526: 526-561 lines are duplicate of `pkg/events/factory_test.go:283-318` (dupl)
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
pkg/events/factory_test.go:314:28: string `test-source` has 5 occurrences, make it a constant (goconst)
			if event.GetSource() != "test-source" {
			                        ^
pkg/events/factory_test.go:7: File is not `gofmt`-ed with `-s` (gofmt)
	"github.com/xeipuuv/gojsonschema"
pkg/events/schema_validation_test.go:149: File is not `gofmt`-ed with `-s` (gofmt)
		
pkg/events/schema_validation_test.go:8: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"github.com/brunseba/cdevents-tools/pkg/events"
pkg/events/factory_test.go:17: line is 143 characters (lll)
	event, err := eventFactory.CreatePipelineRunEvent("started", "pipeline-123", "test-pipeline", "status", "", "https://example.com", customData)
pkg/events/factory_test.go:184: line is 147 characters (lll)
	eventTypes := []string{"testcase-queued", "testcase-started", "testcase-skipped", "testsuite-queued", "testsuite-started", "testoutput-published"}
pkg/events/schema_validation_test.go:15: line is 134 characters (lll)
	event, err := factory.CreatePipelineRunEvent("started", "pipeline-123", "test-pipeline", "", "", "https://example.com/pipeline", nil)
pkg/events/schema_validation_test.go:26: line is 142 characters (lll)
	event, err := factory.CreatePipelineRunEvent("finished", "pipeline-123", "test-pipeline", "success", "", "https://example.com/pipeline", nil)
pkg/events/schema_validation_test.go:48: line is 127 characters (lll)
	event, err := factory.CreateBuildEvent("finished", "build-123", "test-build", "success", "", "https://example.com/build", nil)
pkg/events/schema_validation_test.go:59: line is 134 characters (lll)
	event, err := factory.CreateTaskRunEvent("started", "task-123", "test-task", "pipeline-456", "", "", "https://example.com/task", nil)
pkg/events/schema_validation_test.go:70: line is 142 characters (lll)
	event, err := factory.CreateTaskRunEvent("finished", "task-123", "test-task", "pipeline-456", "success", "", "https://example.com/task", nil)
pkg/events/schema_validation_test.go:115: line is 126 characters (lll)
	event, err := factory.CreateBuildEvent("started", "build-123", "test-build", "", "", "https://example.com/build", customData)
pkg/events/factory_test.go:598:19: SA5011: possible nil pointer dereference (staticcheck)
				if customData.ContentType != "application/json" {
				              ^
pkg/events/factory_test.go:595:8: SA5011(related information): this check suggests that the pointer can be nil (staticcheck)
				if customData == nil {
				   ^
pkg/output/formatters_test.go:11: 11-46 lines are duplicate of `pkg/output/formatters_test.go:207-242` (dupl)
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
pkg/output/formatters_test.go:207: 207-242 lines are duplicate of `pkg/output/formatters_test.go:244-279` (dupl)
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
pkg/output/formatters_test.go:244: 244-279 lines are duplicate of `pkg/output/formatters_test.go:11-46` (dupl)
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
pkg/output/formatters_test.go:86: 86-111 lines are duplicate of `pkg/output/formatters_test.go:180-205` (dupl)
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
pkg/output/formatters_test.go:180: 180-205 lines are duplicate of `pkg/output/formatters_test.go:86-111` (dupl)
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
pkg/output/formatters_test.go:4: File is not `gofmt`-ed with `-s` (gofmt)
	"strings"
	"testing"
pkg/output/formatters_test.go:5: File is not `goimports`-ed with -local github.com/brunseba/cdevents-tools (goimports)
	"testing"
pkg/output/formatters_test.go:6:2: import 'github.com/cdevents/sdk-go/pkg/api' is not allowed from list 'Main' (depguard)
	"github.com/cdevents/sdk-go/pkg/api"
	^
```

## Linter Configuration

The project uses golangci-lint with the following modern linters:

- **revive** (replaces deprecated golint)
- **unused** (replaces deprecated deadcode, structcheck, varcheck)
- **exportloopref** (replaces deprecated scopelint)
- **staticcheck** for advanced static analysis
- **govet** for Go compiler checks
- **errcheck** for error handling verification

## Running Linting

```bash
# Run linting locally
golangci-lint run

# Or using Docker
make quality-docker
```
