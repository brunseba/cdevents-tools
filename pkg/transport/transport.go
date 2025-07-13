package transport

import (
	"context"
	"fmt"
	"strings"

	"github.com/cdevents/sdk-go/pkg/api"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Transport interface for sending events
type Transport interface {
	Send(ctx context.Context, event api.CDEvent) error
}

// HTTPTransport sends events via HTTP
type HTTPTransport struct {
	client cloudevents.Client
	target string
}

// NewHTTPTransport creates a new HTTP transport
func NewHTTPTransport(target string, options ...HTTPOption) (*HTTPTransport, error) {
	client, err := cloudevents.NewClientHTTP()
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	transport := &HTTPTransport{
		client: client,
		target: target,
	}

	for _, option := range options {
		option(transport)
	}

	return transport, nil
}

// HTTPOption configures HTTP transport
type HTTPOption func(*HTTPTransport)

// WithHTTPHeaders adds custom headers to HTTP requests
func WithHTTPHeaders(headers map[string]string) HTTPOption {
	return func(t *HTTPTransport) {
		// Logic to configure headers goes here if needed
	}
}

// Send sends an event via HTTP
func (t *HTTPTransport) Send(ctx context.Context, event api.CDEvent) error {
	ce, err := api.AsCloudEvent(event)
	if err != nil {
		return fmt.Errorf("failed to convert to CloudEvent: %w", err)
	}

	ctx = cloudevents.ContextWithTarget(ctx, t.target)
	ctx = cloudevents.WithEncodingBinary(ctx)

	result := t.client.Send(ctx, *ce)
	if cloudevents.IsUndelivered(result) {
		return fmt.Errorf("failed to send event: %w", result)
	}

	return nil
}

// ConsoleTransport outputs events to console
type ConsoleTransport struct {
	format string
}

// NewConsoleTransport creates a new console transport
func NewConsoleTransport(format string) *ConsoleTransport {
	return &ConsoleTransport{
		format: format,
	}
}

// Send outputs an event to console
func (t *ConsoleTransport) Send(ctx context.Context, event api.CDEvent) error {
	// This would use the output package to format the event
	// For now, just print a simple message
	fmt.Printf("Event sent to console: %s\n", event.GetId())
	return nil
}

// FileTransport writes events to a file
type FileTransport struct {
	filename string
	format   string
}

// NewFileTransport creates a new file transport
func NewFileTransport(filename, format string) *FileTransport {
	return &FileTransport{
		filename: filename,
		format:   format,
	}
}

// Send writes an event to a file
func (t *FileTransport) Send(ctx context.Context, event api.CDEvent) error {
	// Implementation would write to file
	fmt.Printf("Event sent to file %s: %s\n", t.filename, event.GetId())
	return nil
}

// KafkaTransport sends events to Kafka
type KafkaTransport struct {
	brokers []string
	topic   string
	client  cloudevents.Client
}

// NewKafkaTransport creates a new Kafka transport
func NewKafkaTransport(brokers []string, topic string) (*KafkaTransport, error) {
	// For now, return an error indicating Kafka support is not implemented
	return nil, fmt.Errorf("Kafka transport not implemented yet")
}

// Send sends an event to Kafka
func (t *KafkaTransport) Send(ctx context.Context, event api.CDEvent) error {
	return fmt.Errorf("Kafka transport not implemented yet")
}

// TransportFactory creates transports based on configuration
type TransportFactory struct{}

// NewTransportFactory creates a new transport factory
func NewTransportFactory() *TransportFactory {
	return &TransportFactory{}
}

// CreateTransport creates a transport based on the target URL
func (f *TransportFactory) CreateTransport(target string) (Transport, error) {
	if target == "" || target == "console" {
		return NewConsoleTransport("json"), nil
	}

	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
		return NewHTTPTransport(target)
	}

	if strings.HasPrefix(target, "file://") {
		filename := strings.TrimPrefix(target, "file://")
		return NewFileTransport(filename, "json"), nil
	}

	if strings.HasPrefix(target, "kafka://") {
		return nil, fmt.Errorf("Kafka transport not implemented yet")
	}

	return nil, fmt.Errorf("unsupported transport target: %s", target)
}

// MultiTransport sends events to multiple transports
type MultiTransport struct {
	transports []Transport
}

// NewMultiTransport creates a new multi transport
func NewMultiTransport(transports ...Transport) *MultiTransport {
	return &MultiTransport{
		transports: transports,
	}
}

// Send sends an event to all configured transports
func (t *MultiTransport) Send(ctx context.Context, event api.CDEvent) error {
	var errors []string

	for _, transport := range t.transports {
		if err := transport.Send(ctx, event); err != nil {
			errors = append(errors, err.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("transport errors: %s", strings.Join(errors, "; "))
	}

	return nil
}
