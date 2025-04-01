// internal/domain/message.go
package domain

import "time"

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