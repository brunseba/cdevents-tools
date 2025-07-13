package cmd

import (
	"fmt"

	"github.com/cdevents/cdevents-cli/pkg/events"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sendPipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Send pipeline events",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := events.NewEventFactory(getDefaultSource())
		eventType := args[0]

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

		target := viper.GetString("target")
		retries := viper.GetInt("retries")
		timeout := viper.GetDuration("timeout")

		return sendEvent(event, target, retries, timeout)
	},
}

func init() {
	addCommonGenerateFlags(sendPipelineCmd)
	sendCmd.AddCommand(sendPipelineCmd)
}
