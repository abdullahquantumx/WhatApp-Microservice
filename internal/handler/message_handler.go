// internal/handler/message_handler.go
package handler

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/your-org/whatsapp-microservice/internal/service"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
	pb "github.com/your-org/whatsapp-microservice/proto"
)

// GrpcMessageHandler handles gRPC requests for WhatsApp messages
type GrpcMessageHandler struct {
	pb.UnimplementedWhatsAppServiceServer
	messageService service.MessageService
	logger         utils.Logger
}

// NewGrpcMessageHandler creates a new gRPC message handler
func NewGrpcMessageHandler(messageService service.MessageService, logger utils.Logger) *GrpcMessageHandler {
	return &GrpcMessageHandler{
		messageService: messageService,
		logger:         logger,
	}
}

// SendTemplateMessage sends a WhatsApp template message
func (h *GrpcMessageHandler) SendTemplateMessage(ctx context.Context, req *pb.SendTemplateMessageRequest) (*pb.SendTemplateMessageResponse, error) {
	// Validate request
	if req.PhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "phone_number is required")
	}
	if req.TemplateId == "" {
		return nil, status.Error(codes.InvalidArgument, "template_id is required")
	}

	// Convert parameters from proto map to regular map
	parameters := make(map[string]interface{})
	for key, value := range req.Parameters {
		parameters[key] = value
	}

	// Call service
	msg, err := h.messageService.SendTemplateMessage(ctx, req.PhoneNumber, req.TemplateId, parameters, req.OrderId, req.CustomerId)
	if err != nil {
		h.logger.Error("Failed to send template message", "error", err)
		return nil, status.Error(codes.Internal, "failed to send message: "+err.Error())
	}

	// Create response
	resp := &pb.SendTemplateMessageResponse{
		MessageId:  msg.ID,
		Status:     msg.Status,
		ExternalId: msg.ExternalID,
	}

	return resp, nil
}

// GetMessage retrieves a message by ID
func (h *GrpcMessageHandler) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.MessageResponse, error) {
	// Call service
	msg, err := h.messageService.GetMessageByID(ctx, req.MessageId)
	if err != nil {
		h.logger.Error("Failed to get message", "error", err, "message_id", req.MessageId)
		return nil, status.Error(codes.NotFound, "message not found")
	}

	// Convert to proto response
	resp := convertMessageToProto(msg)
	return resp, nil
}

// ListMessages retrieves a list of messages
func (h *GrpcMessageHandler) ListMessages(ctx context.Context, req *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	// Set default limit if not provided
	limit := int(req.Limit)
	if limit <= 0 {
		limit = 10
	}

	// Call service
	messages, err := h.messageService.ListMessages(ctx, req.OrderId, req.CustomerId, req.PhoneNumber, limit, int(req.Offset))
	if err != nil {
		h.logger.Error("Failed to list messages", "error", err)
		return nil, status.Error(codes.Internal, "failed to list messages: "+err.Error())
	}

	// Count total (in a real implementation, this would be a separate query)
	totalCount := len(messages)

	// Convert to proto response
	protoMessages := make([]*pb.MessageResponse, 0, len(messages))
	for _, msg := range messages {
		protoMessages = append(protoMessages, convertMessageToProto(msg))
	}

	// Create response
	resp := &pb.ListMessagesResponse{
		Messages:   protoMessages,
		TotalCount: int32(totalCount),
	}

	return resp, nil
}

// Helper function to convert a service.Message to pb.MessageResponse
func convertMessageToProto(msg *service.Message) *pb.MessageResponse {
	// Convert parameters from map[string]interface{} to map[string]string
	parameters := make(map[string]string)
	for key, value := range msg.Parameters {
		if strValue, ok := value.(string); ok {
			parameters[key] = strValue
		} else {
			// For non-string values, convert to string
			parameters[key] = utils.AnyToString(value)
		}
	}

	return &pb.MessageResponse{
		Id:           msg.ID,
		PhoneNumber:  msg.PhoneNumber,
		TemplateId:   msg.TemplateID,
		Parameters:   parameters,
		OrderId:      msg.OrderID,
		CustomerId:   msg.CustomerID,
		Status:       msg.Status,
		ErrorMessage: msg.ErrorMessage,
		ExternalId:   msg.ExternalID,
		CreatedAt:    msg.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    msg.UpdatedAt.Format(time.RFC3339),
	}
}