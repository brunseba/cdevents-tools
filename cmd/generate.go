package cmd

import (
	"fmt"
	"os"

	"github.com/brunseba/cdevents-tools/pkg/events"
	"github.com/brunseba/cdevents-tools/pkg/output"
	"github.com/cdevents/sdk-go/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate CDEvents",
	Long: `Generate CDEvents for various CI/CD activities.

Supported event types:
- pipeline: Pipeline run events (queued, started, finished)
- task: Task run events (started, finished)
- build: Build events (queued, started, finished)
- service: Service deployment events (deployed, published, removed, rolledback, upgraded)
- test: Test events (testcase-queued, testcase-started, testcase-finished, etc.)

Examples:
  # Generate a pipeline started event
  cdevents-cli generate pipeline started --id "pipeline-123" --name "my-pipeline"
  
# Generate a build finished event with outcome
  cdevents-cli generate build finished --id "build-456" --name "my-build" --outcome "success"

# Generate a task started event with custom data
  cdevents-cli generate task started --id "task-101" --name "my-task" --custom-json '{"key":"value"}'

# Generate a service deployed event
  cdevents-cli generate service deployed --id "service-789" --name "my-service" --environment "prod"`,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

// Common flags for all generate commands
func addCommonGenerateFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("id", "i", "", "Subject ID (required)")
	cmd.Flags().StringP("name", "n", "", "Subject name (required)")
	cmd.Flags().StringP("source", "s", "", "Event source (defaults to hostname)")
	cmd.Flags().StringP("url", "u", "", "Subject URL")
	cmd.Flags().StringP("outcome", "", "", "Outcome for finished events (success, failure, error, cancel)")
	cmd.Flags().StringP("errors", "", "", "Error details for failed events")
	
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("name")

	// Custom data flag
	cmd.Flags().String("custom-json", "", "Custom data in JSON format")
}

// parseCustomData returns custom data parsed from JSON only
func parseCustomData(cmd *cobra.Command) (*events.CustomData, error) {
	customJSON, err := cmd.Flags().GetString("custom-json")
	if err != nil {
		return nil, err
	}

	// Parse custom data from JSON only
	if customJSON != "" {
		return events.ParseCustomDataFromJSON(customJSON)
	}

	return nil, nil
}
func getDefaultSource() string {
	if source := viper.GetString("source"); source != "" {
		return source
	}
	
	hostname, err := os.Hostname()
	if err != nil {
		return "cdevents-cli"
	}
	return fmt.Sprintf("cdevents-cli/%s", hostname)
}

// outputEvent formats and outputs the event
func outputEvent(event interface{}, format string) error {
	return outputEventWithCustomData(event, nil, format)
}

// outputEventWithCustomData formats and outputs the event with custom data
func outputEventWithCustomData(event interface{}, customData *events.CustomData, format string) error {
	if cdEvent, ok := event.(api.CDEvent); ok {
		// Convert events.CustomData to output.CustomData
		var outputCustomData *output.CustomData
		if customData != nil {
			outputCustomData = &output.CustomData{
				Data:        customData.Data,
				ContentType: customData.ContentType,
			}
		}

		formatted, err := output.FormatOutputWithCustomData(cdEvent, outputCustomData, format)
		if err != nil {
			return fmt.Errorf("failed to format output: %w", err)
		}
		fmt.Print(formatted)
		return nil
	}
	return fmt.Errorf("invalid event type")
}
