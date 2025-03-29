// internal/handler/webhook_handler.go
package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-org/whatsapp-microservice/internal/service"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
	pb "github.com/your-org/whatsapp-microservice/proto"
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
	// Read the raw body
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error("Failed to read webhook body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Validate webhook signature if needed
	// For Twilio, signature is in X-Twilio-Signature header
	signature := c.GetHeader("X-Twilio-Signature")
	
	// Process the webhook
	if err := h.webhookService.ProcessWebhook(c.Request.Context(), body, signature, c.Request.URL.String()); err != nil {
		h.logger.Error("Failed to process webhook", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process webhook"})
		return
	}

	// Return 200 OK to acknowledge receipt
	c.Status(http.StatusOK)
}

// Also add gRPC webhook handler if needed
// This could be used by internal services to programmatically trigger webhook-like events

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