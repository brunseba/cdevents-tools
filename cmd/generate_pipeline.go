package cmd

import (
	"fmt"

	"github.com/cdevents/cdevents-cli/pkg/events"
	"github.com/spf13/cobra"
)

var generatePipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Generate pipeline events",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := events.NewEventFactory(getDefaultSource())
		eventType := args[0]

		// Parse custom data
		customData, err := parseCustomData(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse custom data: %w", err)
		}

		event, err := factory.CreatePipelineRunEvent(
			eventType,
			cmd.Flag("id").Value.String(),
			cmd.Flag("name").Value.String(),
			cmd.Flag("outcome").Value.String(),
			cmd.Flag("errors").Value.String(),
			cmd.Flag("url").Value.String(),
			customData,
		)
		if err != nil {
			return fmt.Errorf("failed to create pipeline event: %w", err)
		}

		format := cmd.Flag("output").Value.String()
		return outputEventWithCustomData(event, customData, format)
	},
}

func init() {
	addCommonGenerateFlags(generatePipelineCmd)
	generateCmd.AddCommand(generatePipelineCmd)
}
