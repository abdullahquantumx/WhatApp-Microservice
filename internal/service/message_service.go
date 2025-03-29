// internal/service/message_service.go
package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/your-org/whatsapp-microservice/internal/queue"
	"github.com/your-org/whatsapp-microservice/internal/repository"
	"github.com/your-org/whatsapp-microservice/pkg/twilio"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// Message represents a WhatsApp message
type Message struct {
	ID           int64                  `json:"id"`
	PhoneNumber  string                 `json:"phone_number"`
	TemplateID   string                 `json:"template_id"`
	Parameters   map[string]interface{} `json:"parameters"`
	OrderID      string                 `json:"order_id"`
	CustomerID   string                 `json:"customer_id"`
	Status       string                 `json:"status"`
	ErrorMessage string                 `json:"error_message,omitempty"`
	ExternalID   string                 `json:"external_id,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

// QueueMessage represents a message in the queue
type QueueMessage struct {
	MessageID   int64                  `json:"message_id"`
	PhoneNumber string                 `json:"phone_number"`
	TemplateID  string                 `json:"template_id"`
	Parameters  map[string]interface{} `json:"parameters"`
	OrderID     string                 `json:"order_id"`
	CustomerID  string                 `json:"customer_id"`
}

// MessageService defines the interface for message operations
type MessageService interface {
	SendTemplateMessage(ctx context.Context, phoneNumber, templateID string, parameters map[string]interface{}, orderID, customerID string) (*Message, error)
	GetMessageByID(ctx context.Context, id int64) (*Message, error)
	ListMessages(ctx context.Context, orderID, customerID, phoneNumber string, limit, offset int) ([]*Message, error)
	UpdateMessageStatus(ctx context.Context, externalID, status, errorMessage string) error
	ProcessQueueMessage(ctx context.Context, data []byte) error
}

// messageService implements MessageService
type messageService struct {
	repo      repository.MessageRepository
	whatsapp  twilio.Client
	producer  queue.Producer
	logger    utils.Logger
	isAsync   bool
}

// NewMessageService creates a new message service
func NewMessageService(repo repository.MessageRepository, whatsapp twilio.Client, producer queue.Producer, logger utils.Logger) MessageService {
	return &messageService{
		repo:     repo,
		whatsapp: whatsapp,
		producer: producer,
		logger:   logger,
		isAsync:  true, // Default to async processing
	}
}

// SendTemplateMessage sends a WhatsApp template message
func (s *messageService) SendTemplateMessage(ctx context.Context, phoneNumber, templateID string, parameters map[string]interface{}, orderID, customerID string) (*Message, error) {
	// Normalize phone number if needed
	if !utils.HasWhatsAppPrefix(phoneNumber) {
		phoneNumber = "whatsapp:" + phoneNumber
	}

	// Create message record
	msg := &Message{
		PhoneNumber: phoneNumber,
		TemplateID:  templateID,
		Parameters:  parameters,
		OrderID:     orderID,
		CustomerID:  customerID,
		Status:      "queued",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save to database
	msgID, err := s.repo.CreateMessage(ctx, msg)
	if err != nil {
		return nil, err
	}
	msg.ID = msgID

	if s.isAsync {
		// Queue for async processing
		queueMsg := QueueMessage{
			MessageID:   msg.ID,
			PhoneNumber: msg.PhoneNumber,
			TemplateID:  msg.TemplateID,
			Parameters:  msg.Parameters,
			OrderID:     msg.OrderID,
			CustomerID:  msg.CustomerID,
		}

		// Convert to JSON
		data, err := json.Marshal(queueMsg)
		if err != nil {
			s.logger.Error("Failed to marshal queue message", "error", err)
			return msg, nil // Return success but log error
		}

		// Send to queue
		if err := s.producer.Produce(ctx, data); err != nil {
			s.logger.Error("Failed to produce message to queue", "error", err)
			// Update message status
			if updateErr := s.repo.UpdateMessageStatus(ctx, msg.ID, "failed", "Failed to queue message: "+err.Error(), ""); updateErr != nil {
				s.logger.Error("Failed to update message status", "error", updateErr)
			}
			return nil, err
		}
	} else {
		// Send immediately
		if err := s.sendMessage(ctx, msg); err != nil {
			return nil, err
		}
	}

	return msg, nil
}

// ProcessQueueMessage processes a message from the queue
func (s *messageService) ProcessQueueMessage(ctx context.Context, data []byte) error {
	var queueMsg QueueMessage
	if err := json.Unmarshal(data, &queueMsg); err != nil {
		s.logger.Error("Failed to unmarshal queue message", "error", err)
		return err
	}

	// Get message from database
	msg, err := s.GetMessageByID(ctx, queueMsg.MessageID)
	if err != nil {
		s.logger.Error("Failed to get message from database", "error", err)
		return err
	}

	// Send message
	if err := s.sendMessage(ctx, msg); err != nil {
		s.logger.Error("Failed to send message", "error", err)
		return err
	}

	return nil
}

// sendMessage sends a WhatsApp message
func (s *messageService) sendMessage(ctx context.Context, msg *Message) error {
	// Update status to processing
	if err := s.repo.UpdateMessageStatus(ctx, msg.ID, "processing", "", ""); err != nil {
		return err
	}

	// Send message
	resp, err := s.whatsapp.SendTemplateMessage(ctx, msg.PhoneNumber, msg.TemplateID, msg.Parameters)
	if err != nil {
		// Update status to failed
		updateErr := s.repo.UpdateMessageStatus(ctx, msg.ID, "failed", err.Error(), "")
		if updateErr != nil {
			s.logger.Error("Failed to update message status", "error", updateErr)
		}
		return err
	}

	// Update status to sent
	if err := s.repo.UpdateMessageStatus(ctx, msg.ID, "sent", "", resp.SID); err != nil {
		return err
	}

	return nil
}

// GetMessageByID retrieves a message by ID
func (s *messageService) GetMessageByID(ctx context.Context, id int64) (*Message, error) {
	return s.repo.GetMessageByID(ctx, id)
}

// ListMessages retrieves a list of messages
func (s *messageService) ListMessages(ctx context.Context, orderID, customerID, phoneNumber string, limit, offset int) ([]*Message, error) {
	return s.repo.ListMessages(ctx, orderID, customerID, phoneNumber, limit, offset)
}

// UpdateMessageStatus updates the status of a message
func (s *messageService) UpdateMessageStatus(ctx context.Context, externalID, status, errorMessage string) error {
	if externalID == "" {
		return errors.New("external ID is required")
	}

	msg, err := s.repo.GetMessageByExternalID(ctx, externalID)
	if err != nil {
		return err
	}

	return s.repo.UpdateMessageStatus(ctx, msg.ID, status, errorMessage, externalID)
}