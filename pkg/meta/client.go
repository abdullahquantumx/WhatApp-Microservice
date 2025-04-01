// pkg/meta/client.go
package meta

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"messaging-microservice/pkg/utils"
)

// MessageResponse represents a response from the Meta WhatsApp API
type MessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		WaID string `json:"wa_id"`
	} `json:"contacts,omitempty"`
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages,omitempty"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

// Client defines the interface for WhatsApp API clients
type Client interface {
	SendTemplateMessage(ctx context.Context, to, templateName string, parameters map[string]interface{}) (*MessageResponse, error)
	ValidateWebhookSignature(signatureHeader, url string, body []byte) bool
}

// metaClient implements Client using Meta WhatsApp API
type metaClient struct {
	phoneNumberID string
	accessToken   string
	appSecret     string
	apiURL        string
	httpClient    *http.Client
	logger        utils.Logger
}

// NewClient creates a new Meta WhatsApp client
func NewClient(phoneNumberID, accessToken, appSecret string, logger utils.Logger) Client {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &metaClient{
		phoneNumberID: phoneNumberID,
		accessToken:   accessToken,
		appSecret:     appSecret,
		apiURL:        "https://graph.facebook.com/v18.0", // Using v18.0 as it's current as of writing
		httpClient:    httpClient,
		logger:        logger,
	}
}

// SendTemplateMessage sends a WhatsApp template message through Meta's API
func (c *metaClient) SendTemplateMessage(ctx context.Context, to, templateName string, parameters map[string]interface{}) (*MessageResponse, error) {
	// Normalize phone number (remove WhatsApp prefix if present)
	to = c.normalizePhoneNumber(to)

	// Build template components based on parameters
	components, err := c.buildTemplateComponents(parameters)
	if err != nil {
		return nil, err
	}

	// Prepare request payload
	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                to,
		"type":              "template",
		"template": map[string]interface{}{
			"name":       templateName,
			"language":   map[string]string{"code": "en_US"},
			"components": components,
		},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Create request
	url := fmt.Sprintf("%s/%s/messages", c.apiURL, c.phoneNumberID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check for error status code
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.logger.Error("Meta API error", "status", resp.StatusCode, "body", string(body))
		return nil, fmt.Errorf("meta API error: %d - %s", resp.StatusCode, string(body))
	}

	// Parse response
	var messageResponse MessageResponse
	if err := json.Unmarshal(body, &messageResponse); err != nil {
		return nil, err
	}

	// Check for error in response
	if messageResponse.Error != nil {
		return &messageResponse, fmt.Errorf("meta API error: %d - %s", messageResponse.Error.Code, messageResponse.Error.Message)
	}

	return &messageResponse, nil
}

// ValidateWebhookSignature validates the signature of a webhook from Meta
func (c *metaClient) ValidateWebhookSignature(signature string, _ string, body []byte) bool {
	if c.appSecret == "" || signature == "" {
		return false
	}

	// Extract X-Hub-Signature-256 value
	signatureParts := strings.Split(signature, "=")
	if len(signatureParts) != 2 || signatureParts[0] != "sha256" {
		return false
	}
	receivedSignature := signatureParts[1]

	// Compute HMAC with SHA256
	h := hmac.New(sha256.New, []byte(c.appSecret))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare signatures
	return receivedSignature == expectedSignature
}

// Helper methods

// normalizePhoneNumber removes the "whatsapp:" prefix if present
func (c *metaClient) normalizePhoneNumber(phoneNumber string) string {
	return strings.TrimPrefix(phoneNumber, "whatsapp:")
}

// buildTemplateComponents builds the components array for a template message
func (c *metaClient) buildTemplateComponents(parameters map[string]interface{}) ([]map[string]interface{}, error) {
	if len(parameters) == 0 {
		return nil, nil
	}

	// Convert parameters to component format
	var params []map[string]interface{}
	for _, value := range parameters {
		params = append(params, map[string]interface{}{
			"type": "text",
			"text": fmt.Sprintf("%v", value),
		})
	}

	// Create the body component with parameters
	components := []map[string]interface{}{
		{
			"type":       "body",
			"parameters": params,
		},
	}

	return components, nil
}

// GetMessageExternalID extracts the external message ID from the response
func (c *metaClient) GetMessageExternalID(response *MessageResponse) (string, error) {
	if response == nil {
		return "", errors.New("response is nil")
	}

	if len(response.Messages) > 0 && response.Messages[0].ID != "" {
		return response.Messages[0].ID, nil
	}

	return "", errors.New("no message ID found in response")
}