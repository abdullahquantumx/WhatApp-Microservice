// internal/service/webhook_service.go
package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/your-org/whatsapp-microservice/internal/queue"
	"github.com/your-org/whatsapp-microservice/internal/repository"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// WebhookService defines the interface for webhook operations
type WebhookService interface {
	ProcessWebhook(ctx context.Context, body []byte, signature, url string) error
}

// webhookService implements WebhookService
type webhookService struct {
	repo     repository.MessageRepository
	producer queue.Producer
	logger   utils.Logger
}

// NewWebhookService creates a new webhook service
func NewWebhookService(repo repository.MessageRepository, producer queue.Producer, logger utils.Logger) WebhookService {
	return &webhookService{
		repo:     repo,
		producer: producer,
		logger:   logger,
	}
}

// TwilioWebhookPayload represents a webhook payload from Twilio
type TwilioWebhookPayload struct {
	SmsSid            string `json:"SmsSid"`
	SmsStatus         string `json:"SmsStatus"`
	MessageStatus     string `json:"MessageStatus"`
	To                string `json:"To"`
	MessageSid        string `json:"MessageSid"`
	AccountSid        string `json:"AccountSid"`
	From              string `json:"From"`
	ApiVersion        string `json:"ApiVersion"`
	ErrorCode         string `json:"ErrorCode,omitempty"`
	ErrorMessage      string `json:"ErrorMessage,omitempty"`
	ChannelToAddress  string `json:"ChannelToAddress,omitempty"`
	ChannelPrefix     string `json:"ChannelPrefix,omitempty"`
	ChannelInstallSid string `json:"ChannelInstallSid,omitempty"`
}

// WebhookEvent represents a parsed webhook event
type WebhookEvent struct {
	ExternalID   string `json:"external_id"`
	Status       string `json:"status"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	PhoneNumber  string `json:"phone_number"`
}

// ProcessWebhook processes an incoming webhook
func (s *webhookService) ProcessWebhook(ctx context.Context, body []byte, signature, url string) error {
	// Validate signature
	// For Twilio, you would use their library to validate the signature
	// This is a simplified version
	if signature == "" {
		return errors.New("missing webhook signature")
	}

	// Parse webhook payload
	var twilioPayload TwilioWebhookPayload
	if err := json.Unmarshal(body, &twilioPayload); err != nil {
		s.logger.Error("Failed to unmarshal webhook payload", "error", err)
		return err
	}

	// Map status
	status := mapTwilioStatus(twilioPayload.MessageStatus)

	// Create webhook event
	event := WebhookEvent{
		ExternalID:   twilioPayload.MessageSid,
		Status:       status,
		ErrorCode:    twilioPayload.ErrorCode,
		ErrorMessage: twilioPayload.ErrorMessage,
		PhoneNumber:  twilioPayload.To,
	}

	// Handle webhook asynchronously
	eventData, err := json.Marshal(event)
	if err != nil {
		s.logger.Error("Failed to marshal webhook event", "error", err)
		return err
	}

	if err := s.producer.Produce(ctx, eventData); err != nil {
		s.logger.Error("Failed to produce webhook event to queue", "error", err)
		return err
	}

	// Also update message status directly for immediate feedback
	msg, err := s.repo.GetMessageByExternalID(ctx, event.ExternalID)
	if err == nil {
		if err := s.repo.UpdateMessageStatus(ctx, msg.ID, status, event.ErrorMessage, event.ExternalID); err != nil {
			s.logger.Error("Failed to update message status", "error", err)
		}
	}

	return nil
}

// mapTwilioStatus maps Twilio status to internal status
func mapTwilioStatus(twilioStatus string) string {
	switch twilioStatus {
	case "queued":
		return "queued"
	case "sending":
		return "sending"
	case "sent":
		return "sent"
	case "delivered":
		return "delivered"
	case "read":
		return "read"
	case "failed":
		return "failed"
	case "undelivered":
		return "failed"
	default:
		return "unknown"
	}
}