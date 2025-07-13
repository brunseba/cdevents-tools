package transport_test

import (
	"context"
	"fmt"
	"testing"
	"github.com/cdevents/sdk-go/pkg/api"
	"github.com/cdevents/cdevents-cli/pkg/transport"
	cdeventsv04 "github.com/cdevents/sdk-go/pkg/api/v04"
)

func TestNewHTTPTransport(t *testing.T) {
	transport, err := transport.NewHTTPTransport("http://example.com")
	if err != nil {
		t.Fatalf("unexpected error creating HTTP transport: %v", err)
	}
	if transport == nil {
		t.Fatalf("HTTP transport should not be nil")
	}
}

func TestNewConsoleTransport(t *testing.T) {
	consoleTransport := transport.NewConsoleTransport("json")
	if consoleTransport == nil {
		t.Fatalf("console transport should not be nil")
	}
}

func TestNewFileTransport(t *testing.T) {
	fileTransport := transport.NewFileTransport("/tmp/test.json", "json")
	if fileTransport == nil {
		t.Fatalf("file transport should not be nil")
	}
}

func TestNewKafkaTransport(t *testing.T) {
	_, err := transport.NewKafkaTransport([]string{"localhost:9092"}, "test-topic")
	if err == nil {
		t.Fatalf("expected error for unimplemented Kafka transport")
	}
}

func TestNewTransportFactory(t *testing.T) {
	factory := transport.NewTransportFactory()
	if factory == nil {
		t.Fatalf("transport factory should not be nil")
	}
}

func TestTransportFactory_CreateTransport(t *testing.T) {
	factory := transport.NewTransportFactory()

	testCases := []struct {
		name      string
		target    string
		shouldErr bool
	}{
		{"empty target", "", false},
		{"console target", "console", false},
		{"http target", "http://example.com", false},
		{"https target", "https://example.com", false},
		{"file target", "file:///tmp/test.json", false},
		{"kafka target", "kafka://localhost:9092", true},
		{"unsupported target", "ftp://example.com", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			transport, err := factory.CreateTransport(tc.target)

			if tc.shouldErr {
				if err == nil {
					t.Errorf("expected error for target: %s", tc.target)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if transport == nil {
				t.Fatalf("transport should not be nil")
			}
		})
	}
}

func TestNewMultiTransport(t *testing.T) {
	console := transport.NewConsoleTransport("json")
	file := transport.NewFileTransport("/tmp/test.json", "json")
	
	multiTransport := transport.NewMultiTransport(console, file)
	if multiTransport == nil {
		t.Fatalf("multi transport should not be nil")
	}
}

func TestConsoleTransport_Send(t *testing.T) {
	consoleTransport := transport.NewConsoleTransport("json")
	
	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")

	err = consoleTransport.Send(context.Background(), event)
	if err != nil {
		t.Fatalf("unexpected error sending to console: %v", err)
	}
}

func TestFileTransport_Send(t *testing.T) {
	fileTransport := transport.NewFileTransport("/tmp/test.json", "json")
	
	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")

	err = fileTransport.Send(context.Background(), event)
	if err != nil {
		t.Fatalf("unexpected error sending to file: %v", err)
	}
}

func TestMultiTransport_Send(t *testing.T) {
	console := transport.NewConsoleTransport("json")
	file := transport.NewFileTransport("/tmp/test.json", "json")
	multiTransport := transport.NewMultiTransport(console, file)
	
	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")

	err = multiTransport.Send(context.Background(), event)
	if err != nil {
		t.Fatalf("unexpected error sending to multi transport: %v", err)
	}
}

func TestHTTPTransport_Send(t *testing.T) {
	// Test sending event via HTTP transport
	// Note: This will fail because httpbin is not running in test environment
	// but it tests the code path
	httpTransport, err := transport.NewHTTPTransport("http://httpbin.org/post")
	if err != nil {
		t.Fatalf("failed to create HTTP transport: %v", err)
	}

	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// This will likely fail due to network issues in test environment
	// but we're testing the code path
	err = httpTransport.Send(context.Background(), event)
	// We don't assert on the error because it depends on network connectivity
	_ = err // Suppress unused variable warning
}

func TestMultiTransport_SendWithErrors(t *testing.T) {
	// Create a mock failing transport
	failingTransport := &failingTransport{}
	console := transport.NewConsoleTransport("json")
	multiTransport := transport.NewMultiTransport(failingTransport, console)
	
	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	err = multiTransport.Send(context.Background(), event)
	// Multi transport should return error if any transport fails
	if err == nil {
		t.Errorf("expected error from multi transport with failing transport")
	}
}

// failingTransport is a mock transport that always fails
type failingTransport struct{}

func (t *failingTransport) Send(ctx context.Context, event api.CDEvent) error {
	return fmt.Errorf("mock transport failure")
}

func TestWithHTTPHeaders(t *testing.T) {
	// Test the HTTP headers option
	headers := map[string]string{"Authorization": "Bearer token"}
	option := transport.WithHTTPHeaders(headers)

	// This should not panic
	httpTransport, err := transport.NewHTTPTransport("http://example.com", option)
	if err != nil {
		t.Fatalf("failed to create HTTP transport with headers: %v", err)
	}
	if httpTransport == nil {
		t.Fatalf("HTTP transport should not be nil")
	}
}

func TestKafkaTransport_Send(t *testing.T) {
	// Since Kafka transport always returns error, test that path
	// We can't instantiate it directly, but we can test the error path
	kafkaTransport := &struct {
		brokers []string
		topic   string
	}{
		brokers: []string{"localhost:9092"},
		topic:   "test-topic",
	}
	
	// Create a test event
	event, err := cdeventsv04.NewPipelineRunQueuedEvent()
	if err != nil {
		t.Fatalf("failed to create test event: %v", err)
	}
	event.SetId("test-id")
	event.SetSource("test-source")
	event.SetSubjectId("pipeline-123")
	event.SetSubjectPipelineName("test-pipeline")

	// We can't call Send directly as KafkaTransport is not exported
	// but we can test that NewKafkaTransport returns error
	_, err = transport.NewKafkaTransport(kafkaTransport.brokers, kafkaTransport.topic)
	if err == nil {
		t.Fatalf("expected error for unimplemented Kafka transport")
	}
}
