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
