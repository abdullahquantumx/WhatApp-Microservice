// test/message_service_test.go
package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/your-org/whatsapp-microservice/internal/service"
	"github.com/your-org/whatsapp-microservice/pkg/twilio"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// Mock repositories and clients
type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) CreateMessage(ctx context.Context, message *service.Message) (int64, error) {
	args := m.Called(ctx, message)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockMessageRepository) GetMessageByID(ctx context.Context, id int64) (*service.Message, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.Message), args.Error(1)
}

func (m *MockMessageRepository) GetMessageByExternalID(ctx context.Context, externalID string) (*service.Message, error) {
	args := m.Called(ctx, externalID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*service.Message), args.Error(1)
}

func (m *MockMessageRepository) ListMessages(ctx context.Context, orderID, customerID, phoneNumber string, limit, offset int) ([]*service.Message, error) {
	args := m.Called(ctx, orderID, customerID, phoneNumber, limit, offset)
	return args.Get(0).([]*service.Message), args.Error(1)
}

func (m *MockMessageRepository) UpdateMessageStatus(ctx context.Context, id int64, status, errorMessage, externalID string) error {
	args := m.Called(ctx, id, status, errorMessage, externalID)
	return args.Error(0)
}

type MockWhatsAppClient struct {
	mock.Mock
}

func (m *MockWhatsAppClient) SendTemplateMessage(ctx context.Context, to, templateName string, parameters map[string]interface{}) (*twilio.MessageResponse, error) {
	args := m.Called(ctx, to, templateName, parameters)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*twilio.MessageResponse), args.Error(1)
}

func (m *MockWhatsAppClient) ValidateWebhookSignature(signatureHeader, url string, body []byte) bool {
	args := m.Called(signatureHeader, url, body)
	return args.Bool(0)
}

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce(ctx context.Context, value []byte) error {
	args := m.Called(ctx, value)
	return args.Error(0)
}

func (m *MockProducer) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Mock logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Info(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Warn(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Error(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

func (m *MockLogger) Fatal(msg string, keysAndValues ...interface{}) {
	m.Called(msg, keysAndValues)
}

// Test SendTemplateMessage
func TestSendTemplateMessage(t *testing.T) {
	// Create mocks
	mockRepo := new(MockMessageRepository)
	mockWhatsApp := new(MockWhatsAppClient)
	mockProducer := new(MockProducer)
	mockLogger := new(MockLogger)

	// Create test data
	phoneNumber := "+1234567890"
	templateID := "order_confirmation"
	parameters := map[string]interface{}{
		"order_id": "ORD-12345",
	}
	orderID := "ORD-12345"
	customerID := "CUST-6789"

	// Set up mock expectations
	mockRepo.On("CreateMessage", mock.Anything, mock.MatchedBy(func(m *service.Message) bool {
		return m.PhoneNumber == "whatsapp:+1234567890" && m.TemplateID == templateID
	})).Return(1, nil)

	mockProducer.On("Produce", mock.Anything, mock.Anything).Return(nil)

	// Set up logger expectations
	mockLogger.On("Error", mock.Anything, mock.Anything).Maybe()
	mockLogger.On("Info", mock.Anything, mock.Anything).Maybe()

	// Create service
	svc := service.NewMessageService(mockRepo, mockWhatsApp, mockProducer, mockLogger)

	// Test
	ctx := context.Background()
	msg, err := svc.SendTemplateMessage(ctx, phoneNumber, templateID, parameters, orderID, customerID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, msg)
	assert.Equal(t, int64(1), msg.ID)
	assert.Equal(t, "whatsapp:+1234567890", msg.PhoneNumber)
	assert.Equal(t, templateID, msg.TemplateID)
	assert.Equal(t, "queued", msg.Status)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
	mockProducer.AssertExpectations(t)
}

// Test SendTemplateMessage with repository error
func TestSendTemplateMessageRepositoryError(t *testing.T) {
	// Create mocks
	mockRepo := new(MockMessageRepository)
	mockWhatsApp := new(MockWhatsAppClient)
	mockProducer := new(MockProducer)
	mockLogger := new(MockLogger)

	// Create test data
	phoneNumber := "+1234567890"
	templateID := "order_confirmation"
	parameters := map[string]interface{}{
		"order_id": "ORD-12345",
	}
	orderID := "ORD-12345"
	customerID := "CUST-6789"

	// Set up mock expectations with error
	mockRepo.On("CreateMessage", mock.Anything, mock.Anything).Return(0, errors.New("database error"))

	// Set up logger expectations
	mockLogger.On("Error", mock.Anything, mock.Anything).Maybe()

	// Create service
	svc := service.NewMessageService(mockRepo, mockWhatsApp, mockProducer, mockLogger)

	// Test
	ctx := context.Background()
	msg, err := svc.SendTemplateMessage(ctx, phoneNumber, templateID, parameters, orderID, customerID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, msg)
	assert.Contains(t, err.Error(), "database error")

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
	mockProducer.AssertNotCalled(t, "Produce", mock.Anything, mock.Anything)
}
