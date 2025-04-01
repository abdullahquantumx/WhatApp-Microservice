// internal/service/webhook_service.go
package service

import (
	"context"
	"encoding/json"
	"errors"

	"messaging-microservice/internal/queue"
	"messaging-microservice/internal/repository"
	"messaging-microservice/pkg/utils"
)

// WebhookService defines the interface for webhook operations
type WebhookService interface {
	ProcessWebhook(ctx context.Context, body []byte, signature, url string) error
	UpdateMessageStatus(ctx context.Context, externalID, status, errorMessage string) error
	GetVerifyToken() string
}

// webhookService implements WebhookService
type webhookService struct {
	repo       repository.MessageRepository
	producer   queue.Producer
	logger     utils.Logger
	verifyToken string
}

// NewWebhookService creates a new webhook service
func NewWebhookService(repo repository.MessageRepository, producer queue.Producer, logger utils.Logger, verifyToken string) WebhookService {
	return &webhookService{
		repo:       repo,
		producer:   producer,
		logger:     logger,
		verifyToken: verifyToken,
	}
}

// MetaWebhookPayload represents the root structure of a Meta webhook payload
type MetaWebhookPayload struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Value struct {
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberID      string `json:"phone_number_id"`
				} `json:"metadata"`
				Statuses []struct {
					ID          string `json:"id"`
					RecipientID string `json:"recipient_id"`
					Status      string `json:"status"`
					Timestamp   string `json:"timestamp"`
					Errors      []struct {
						Code    int    `json:"code"`
						Title   string `json:"title"`
						Message string `json:"message"`
					} `json:"errors,omitempty"`
				} `json:"statuses,omitempty"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
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
	// This would need to be implemented with your Meta client
	if signature == "" {
		return errors.New("missing webhook signature")
	}

	// Parse webhook payload
	var metaPayload MetaWebhookPayload
	if err := json.Unmarshal(body, &metaPayload); err != nil {
		s.logger.Error("Failed to unmarshal webhook payload", "error", err)
		return err
	}

	// Check if it's a valid WhatsApp webhook
	if metaPayload.Object != "whatsapp_business_account" {
		s.logger.Warn("Received non-WhatsApp webhook", "object", metaPayload.Object)
		return nil // Not an error, just not relevant for us
	}

	// Process each status update
	for _, entry := range metaPayload.Entry {
		for _, change := range entry.Changes {
			for _, status := range change.Value.Statuses {
				// Map status
				mappedStatus := mapMetaStatus(status.Status)
				
				// Extract error info
				var errorMessage string
				if len(status.Errors) > 0 {
					errorMessage = status.Errors[0].Message
				}

				// Create webhook event
				event := WebhookEvent{
					ExternalID:   status.ID,
					Status:       mappedStatus,
					ErrorMessage: errorMessage,
					PhoneNumber:  status.RecipientID,
				}

				// Handle webhook asynchronously
				eventData, err := json.Marshal(event)
				if err != nil {
					s.logger.Error("Failed to marshal webhook event", "error", err)
					continue
				}

				if err := s.producer.Produce(ctx, eventData); err != nil {
					s.logger.Error("Failed to produce webhook event to queue", "error", err)
					continue
				}

				// Also update message status directly for immediate feedback
				s.UpdateMessageStatus(ctx, event.ExternalID, event.Status, event.ErrorMessage)
			}
		}
	}

	return nil
}

// UpdateMessageStatus updates the status of a message
func (s *webhookService) UpdateMessageStatus(ctx context.Context, externalID, status, errorMessage string) error {
	if externalID == "" {
		return errors.New("external ID is required")
	}

	msg, err := s.repo.GetMessageByExternalID(ctx, externalID)
	if err != nil {
		return err
	}

	return s.repo.UpdateMessageStatus(ctx, msg.ID, status, errorMessage, externalID)
}

// GetVerifyToken returns the verification token for webhook setup
func (s *webhookService) GetVerifyToken() string {
	return s.verifyToken
}

// mapMetaStatus maps Meta status to internal status
func mapMetaStatus(metaStatus string) string {
	switch metaStatus {
	case "sent":
		return "sent"
	case "delivered":
		return "delivered"
	case "read":
		return "read"
	case "failed":
		return "failed"
	default:
		return "unknown"
	}
}