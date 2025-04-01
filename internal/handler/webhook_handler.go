// internal/handler/webhook_handler.go
package handler

import (
	"io/ioutil"
	"net/http"
	"context"

	"github.com/gin-gonic/gin"
	"messaging-microservice/internal/service"
	"messaging-microservice/pkg/utils"
	pb "messaging-microservice/proto"
)

// WebhookHandler handles webhook callbacks from WhatsApp
type WebhookHandler struct {
	webhookService service.WebhookService
	logger         utils.Logger
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler(webhookService service.WebhookService, logger utils.Logger) *WebhookHandler {
	return &WebhookHandler{
		webhookService: webhookService,
		logger:         logger,
	}
}

// HandleWebhook processes incoming webhook events from WhatsApp
func (h *WebhookHandler) HandleWebhook(c *gin.Context) {
	// Check if this is a verification request from Meta
	if c.Query("hub.mode") == "subscribe" && c.Query("hub.verify_token") != "" {
		h.handleVerification(c)
		return
	}

	// Read the raw body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error("Failed to read webhook body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Validate webhook signature
	// For Meta, signature is in X-Hub-Signature-256 header
	signature := c.GetHeader("X-Hub-Signature-256")
	
	// Process the webhook
	if err := h.webhookService.ProcessWebhook(c.Request.Context(), body, signature, c.Request.URL.String()); err != nil {
		h.logger.Error("Failed to process webhook", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process webhook"})
		return
	}

	// Return 200 OK to acknowledge receipt
	c.Status(http.StatusOK)
}

// handleVerification handles the webhook verification request from Meta
func (h *WebhookHandler) handleVerification(c *gin.Context) {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if mode != "subscribe" || token == "" {
		h.logger.Error("Invalid verification request", "mode", mode, "token", token)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification request"})
		return
	}

	// Verify the token against your configured verify token
	// This should be loaded from your configuration
	verifyToken := h.webhookService.GetVerifyToken()
	if token != verifyToken {
		h.logger.Error("Invalid verify token", "received", token, "expected", verifyToken)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verify token"})
		return
	}

	// If verification succeeds, respond with the challenge
	c.String(http.StatusOK, challenge)
}

// HandleGrpcWebhook handles webhook events coming through gRPC
func (h *WebhookHandler) HandleGrpcWebhook(ctx context.Context, req *pb.WebhookRequest) (*pb.WebhookResponse, error) {
	// Process the webhook
	err := h.webhookService.UpdateMessageStatus(ctx, req.ExternalId, req.Status, req.ErrorMessage)
	if err != nil {
		h.logger.Error("Failed to process gRPC webhook", "error", err)
		return &pb.WebhookResponse{
			Success: false,
			Message: "Failed to process webhook: " + err.Error(),
		}, nil
	}

	return &pb.WebhookResponse{
		Success: true,
		Message: "Webhook processed successfully",
	}, nil
}