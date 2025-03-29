// pkg/twilio/client.go
package twilio

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// MessageResponse represents a response from the Twilio API
type MessageResponse struct {
	SID                 string `json:"sid"`
	AccountSID          string `json:"account_sid"`
	MessagingServiceSID string `json:"messaging_service_sid"`
	Status              string `json:"status"`
	ErrorCode           string `json:"error_code,omitempty"`
	ErrorMessage        string `json:"error_message,omitempty"`
	DateCreated         string `json:"date_created"`
	DateUpdated         string `json:"date_updated"`
}

// Client defines the interface for WhatsApp API clients
type Client interface {
	SendTemplateMessage(ctx context.Context, to, templateName string, parameters map[string]interface{}) (*MessageResponse, error)
	ValidateWebhookSignature(signatureHeader, url string, body []byte) bool
}

// twilioClient implements Client using Twilio
type twilioClient struct {
	accountSID  string
	authToken   string
	fromNumber  string
	httpClient  *http.Client
	logger      utils.Logger
	apiEndpoint string
}

// NewClient creates a new Twilio client
func NewClient(accountSID, authToken, fromNumber string, logger utils.Logger) Client {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &twilioClient{
		accountSID:  accountSID,
		authToken:   authToken,
		fromNumber:  fromNumber,
		httpClient:  httpClient,
		logger:      logger,
		apiEndpoint: "https://api.twilio.com/2010-04-01",
	}
}

// SendTemplateMessage sends a WhatsApp template message
func (c *twilioClient) SendTemplateMessage(ctx context.Context, to, templateName string, parameters map[string]interface{}) (*MessageResponse, error) {
	// Normalize phone number if needed
	if !utils.HasWhatsAppPrefix(to) {
		to = "whatsapp:" + to
	}

	// Build Twilio API endpoint
	endpoint := fmt.Sprintf("%s/Accounts/%s/Messages.json", c.apiEndpoint, c.accountSID)

	// Create form data
	formData := url.Values{}
	formData.Set("To", to)
	formData.Set("From", c.fromNumber)

	// Build template body
	body, err := c.buildTemplateBody(templateName, parameters)
	if err != nil {
		return nil, err
	}
	formData.Set("Body", body)

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.SetBasicAuth(c.accountSID, c.authToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for error
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		var twilioError struct {
			Code     int    `json:"code"`
			Message  string `json:"message"`
			MoreInfo string `json:"more_info"`
			Status   int    `json:"status"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&twilioError); err != nil {
			return nil, fmt.Errorf("twilio API error: %d", resp.StatusCode)
		}

		return nil, fmt.Errorf("twilio API error: %s", twilioError.Message)
	}

	// Parse response
	var messageResponse MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
		return nil, err
	}

	return &messageResponse, nil
}

// buildTemplateBody builds the WhatsApp template body
func (c *twilioClient) buildTemplateBody(templateName string, parameters map[string]interface{}) (string, error) {
	// For Twilio, template format is different from pure WhatsApp API
	// This is a simplified implementation
	// You may need to adjust this based on your actual templates

	switch templateName {
	case "order_confirmation":
		orderID, ok := parameters["order_id"].(string)
		if !ok {
			return "", errors.New("order_id parameter is required")
		}
		return fmt.Sprintf("Your order %s has been confirmed. Thank you for your purchase!", orderID), nil

	case "shipment_dispatched":
		orderID, ok := parameters["order_id"].(string)
		if !ok {
			return "", errors.New("order_id parameter is required")
		}
		trackingID, ok := parameters["tracking_id"].(string)
		if !ok {
			return "", errors.New("tracking_id parameter is required")
		}
		eta, ok := parameters["estimated_delivery"].(string)
		if !ok {
			return "", errors.New("estimated_delivery parameter is required")
		}
		return fmt.Sprintf("Your order %s has been dispatched! Track your package with tracking ID %s. Estimated delivery: %s", orderID, trackingID, eta), nil

	case "delivery_eta":
		orderID, ok := parameters["order_id"].(string)
		if !ok {
			return "", errors.New("order_id parameter is required")
		}
		eta, ok := parameters["estimated_delivery"].(string)
		if !ok {
			return "", errors.New("estimated_delivery parameter is required")
		}
		return fmt.Sprintf("Your order %s is scheduled for delivery on %s", orderID, eta), nil

	case "delivery_confirmation":
		orderID, ok := parameters["order_id"].(string)
		if !ok {
			return "", errors.New("order_id parameter is required")
		}
		return fmt.Sprintf("Your order %s has been delivered. Thank you for choosing our service!", orderID), nil

	case "delay_notification":
		orderID, ok := parameters["order_id"].(string)
		if !ok {
			return "", errors.New("order_id parameter is required")
		}
		reason, ok := parameters["reason"].(string)
		if !ok {
			return "", errors.New("reason parameter is required")
		}
		newEta, ok := parameters["new_estimated_delivery"].(string)
		if !ok {
			return "", errors.New("new_estimated_delivery parameter is required")
		}
		return fmt.Sprintf("We're sorry, but your order %s has been delayed due to %s. New estimated delivery: %s", orderID, reason, newEta), nil

	default:
		return "", fmt.Errorf("unknown template: %s", templateName)
	}
}

// ValidateWebhookSignature validates the signature of a Twilio webhook
func (c *twilioClient) ValidateWebhookSignature(signatureHeader, url string, body []byte) bool {
	// In a real implementation, you would use Twilio's SDK to validate the signature
	// This is a placeholder implementation
	// For production, you should use the official Twilio helper library

	// Example validation logic would look like:
	// import "github.com/twilio/twilio-go/client"
	// return client.ValidateSignature(c.authToken, signatureHeader, url, body)

	// For now, just log that we would validate
	c.logger.Info("Validating webhook signature", "url", url, "signature", signatureHeader)
	
	// Always return true for this example
	// In production, replace with actual validation
	return true
}