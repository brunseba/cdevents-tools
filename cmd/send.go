package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/cdevents/cdevents-cli/pkg/transport"
	"github.com/cdevents/sdk-go/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send CDEvents to various targets",
	Long: `Send CDEvents to various targets such as HTTP endpoints, files, or console.

This command combines event generation and transmission in one step.

Examples:
  # Send a pipeline started event via HTTP
  cdevents-cli send --target http://localhost:8080/events pipeline started --id "pipeline-123" --name "my-pipeline"
  
  # Send a build finished event to console
  cdevents-cli send --target console build finished --id "build-456" --name "my-build" --outcome "success"
  
  # Send a service deployed event to a file
  cdevents-cli send --target file://events.json service deployed --id "service-789" --name "my-service"`,
}

func init() {
	rootCmd.AddCommand(sendCmd)
	
	// Add transport-specific flags
	sendCmd.PersistentFlags().StringP("target", "t", "console", "Target to send events to (console, http://..., file://...)")
	sendCmd.PersistentFlags().IntP("retries", "r", 3, "Number of retry attempts")
	sendCmd.PersistentFlags().DurationP("timeout", "", 30*time.Second, "Request timeout")
	sendCmd.PersistentFlags().StringSliceP("headers", "H", []string{}, "HTTP headers (format: key=value)")
	
	// Bind flags to viper
	viper.BindPFlag("target", sendCmd.PersistentFlags().Lookup("target"))
	viper.BindPFlag("retries", sendCmd.PersistentFlags().Lookup("retries"))
	viper.BindPFlag("timeout", sendCmd.PersistentFlags().Lookup("timeout"))
	viper.BindPFlag("headers", sendCmd.PersistentFlags().Lookup("headers"))
}

// sendEvent sends an event using the specified transport
func sendEvent(event interface{}, target string, retries int, timeout time.Duration) error {
	cdEvent, ok := event.(api.CDEvent)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	factory := transport.NewTransportFactory()
	transport, err := factory.CreateTransport(target)
	if err != nil {
		return fmt.Errorf("failed to create transport: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if retries > 0 {
	return SendEventWithRetry(ctx, transport, cdEvent, retries)
	}

	return transport.Send(ctx, cdEvent)
}

// SendEventWithRetry sends an event with retry logic
func SendEventWithRetry(ctx context.Context, transport transport.Transport, event api.CDEvent, maxRetries int) error {
	var lastErr error

	for i := 0; i <= maxRetries; i++ {
		if err := transport.Send(ctx, event); err != nil {
			lastErr = err
			if i < maxRetries {
				// Could add exponential backoff here
				continue
			}
		} else {
			return nil
		}
	}

	return fmt.Errorf("failed to send event after %d retries: %w", maxRetries, lastErr)
}
