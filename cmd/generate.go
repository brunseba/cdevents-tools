package cmd

import (
	"fmt"
	"os"

	"github.com/cdevents/cdevents-cli/pkg/events"
	"github.com/cdevents/cdevents-cli/pkg/output"
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
  cdevents-cli generate task started --id "task-101" --name "my-task" --custom "key=value"

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

	// Custom data flags
	cmd.Flags().StringSlice("custom", []string{}, "Custom data as key=value pairs")
	cmd.Flags().String("custom-json", "", "Custom data in JSON format")
	cmd.Flags().String("custom-yaml", "", "Custom data in YAML format")
}

// parseCustomData returns custom data parsed from the command
func parseCustomData(cmd *cobra.Command) (*events.CustomData, error) {
	customKeyValue, err := cmd.Flags().GetStringSlice("custom")
	if err != nil {
		return nil, err
	}

	customJSON, err := cmd.Flags().GetString("custom-json")
	if err != nil {
		return nil, err
	}

	customYAML, err := cmd.Flags().GetString("custom-yaml")
	if err != nil {
		return nil, err
	}

	// Parse custom data from JSON, YAML, or key=value pairs
	if customJSON != "" {
		return events.ParseCustomDataFromJSON(customJSON)
	}

	if customYAML != "" {
		return events.ParseCustomDataFromYAML(customYAML)
	}

	return events.ParseCustomDataFromKeyValue(customKeyValue)
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
	if cdEvent, ok := event.(api.CDEvent); ok {
		formatted, err := output.FormatOutput(cdEvent, format)
		if err != nil {
			return fmt.Errorf("failed to format output: %w", err)
		}
		fmt.Print(formatted)
		return nil
	}
	return fmt.Errorf("invalid event type")
}
