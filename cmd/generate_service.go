package cmd

import (
	"fmt"

	"github.com/cdevents/cdevents-cli/pkg/events"
	"github.com/spf13/cobra"
)

var generateServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Generate service events",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := events.NewEventFactory(getDefaultSource())
		eventType := args[0]

		// Parse custom data
		customData, err := parseCustomData(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse custom data: %w", err)
		}

		event, err := factory.CreateServiceEvent(
			eventType,
			cmd.Flag("id").Value.String(),
			cmd.Flag("name").Value.String(),
			cmd.Flag("environment").Value.String(),
			cmd.Flag("url").Value.String(),
			customData,
		)
		if err != nil {
			return fmt.Errorf("failed to create service event: %w", err)
		}

		format := cmd.Flag("output").Value.String()
		return outputEvent(event, format)
	},
}

func init() {
	addCommonGenerateFlags(generateServiceCmd)
	generateServiceCmd.Flags().StringP("environment", "e", "", "Environment ID")
	generateCmd.AddCommand(generateServiceCmd)
}
