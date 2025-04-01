// test/queue_test.go
package test

import (
	"context"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"messaging-microservice/internal/queue"
	"messaging-microservice/pkg/utils"
)

// Mock logger for queue tests
type MockQueueLogger struct {
	mock.Mock
}

func (m *MockQueueLogger) Debug(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockQueueLogger) Info(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockQueueLogger) Warn(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockQueueLogger) Error(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockQueueLogger) Fatal(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

// MockKafkaWriter mocks the Kafka writer
type MockKafkaWriter struct {
	mock.Mock
}

func (m *MockKafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	args := m.Called(ctx, msgs)
	return args.Error(0)
}

func (m *MockKafkaWriter) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Test producer
func TestProducer(t *testing.T) {
	// Create context with cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create mock logger
	mockLogger := new(MockQueueLogger)
	mockLogger.On("Error", mock.Anything, mock.Anything).Maybe()
	mockLogger.On("Info", mock.Anything, mock.Anything).Maybe()

	// Create mock Kafka writer
	mockKafkaWriter := new(MockKafkaWriter)
	mockKafkaWriter.On("WriteMessages", mock.Anything, mock.Anything).Return(nil)
	mockKafkaWriter.On("Close").Return(nil)

	// Test data
	testData := []byte(`{"message_id": 1, "phone_number": "+1234567890"}`)

	// Test function to override writer creation
	writerCreator := func(brokers []string, topic string, logger utils.Logger) (interface{}, error) {
		return mockKafkaWriter, nil
	}

	// Create producer with mock Kafka writer
	producer, err := queue.NewProducerWithWriter([]string{"localhost:9092"}, "test-topic", mockLogger, writerCreator)
	assert.NoError(t, err)
	assert.NotNil(t, producer)

	// Test produce
	err = producer.Produce(ctx, testData)
	assert.NoError(t, err)

	// Test close
	err = producer.Close()
	assert.NoError(t, err)

	// Verify mock expectations
	mockKafkaWriter.AssertExpectations(t)
}