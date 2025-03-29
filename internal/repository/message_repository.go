// internal/repository/message_repository.go
package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// MessageModel represents a message in the database
type MessageModel struct {
	ID           int64          `db:"id"`
	PhoneNumber  string         `db:"phone_number"`
	TemplateID   string         `db:"template_id"`
	Parameters   string         `db:"parameters"`
	OrderID      sql.NullString `db:"order_id"`
	CustomerID   sql.NullString `db:"customer_id"`
	Status       string         `db:"status"`
	ErrorMessage sql.NullString `db:"error_message"`
	ExternalID   sql.NullString `db:"external_id"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}

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

// MessageRepository defines the interface for database operations
type MessageRepository interface {
	CreateMessage(ctx context.Context, message *Message) (int64, error)
	GetMessageByID(ctx context.Context, id int64) (*Message, error)
	GetMessageByExternalID(ctx context.Context, externalID string) (*Message, error)
	ListMessages(ctx context.Context, orderID, customerID, phoneNumber string, limit, offset int) ([]*Message, error)
	UpdateMessageStatus(ctx context.Context, id int64, status, errorMessage, externalID string) error
}

// messageRepository implements MessageRepository
type messageRepository struct {
	db     *sqlx.DB
	logger utils.Logger
}

// NewMessageRepository creates a new message repository
func NewMessageRepository(db *sqlx.DB, logger utils.Logger) MessageRepository {
	return &messageRepository{
		db:     db,
		logger: logger,
	}
}

// CreateMessage creates a new message
func (r *messageRepository) CreateMessage(ctx context.Context, message *Message) (int64, error) {
	// Convert parameters to JSON
	paramsJSON, err := json.Marshal(message.Parameters)
	if err != nil {
		return 0, err
	}

	// Create model
	model := MessageModel{
		PhoneNumber:  message.PhoneNumber,
		TemplateID:   message.TemplateID,
		Parameters:   string(paramsJSON),
		Status:       message.Status,
		CreatedAt:    message.CreatedAt,
		UpdatedAt:    message.UpdatedAt,
	}

	// Set nullable fields
	if message.OrderID != "" {
		model.OrderID = sql.NullString{String: message.OrderID, Valid: true}
	}
	if message.CustomerID != "" {
		model.CustomerID = sql.NullString{String: message.CustomerID, Valid: true}
	}
	if message.ErrorMessage != "" {
		model.ErrorMessage = sql.NullString{String: message.ErrorMessage, Valid: true}
	}
	if message.ExternalID != "" {
		model.ExternalID = sql.NullString{String: message.ExternalID, Valid: true}
	}

	// Insert into database
	query := `
		INSERT INTO messages (
			phone_number, template_id, parameters, 
			order_id, customer_id, status, 
			error_message, external_id, created_at, updated_at
		) VALUES (
			:phone_number, :template_id, :parameters, 
			:order_id, :customer_id, :status, 
			:error_message, :external_id, :created_at, :updated_at
		) RETURNING id
	`

	rows, err := r.db.NamedQueryContext(ctx, query, model)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("no id returned after insert")
	}

	return id, nil
}

// GetMessageByID retrieves a message by ID
func (r *messageRepository) GetMessageByID(ctx context.Context, id int64) (*Message, error) {
	query := `
		SELECT id, phone_number, template_id, parameters, 
			order_id, customer_id, status, 
			error_message, external_id, created_at, updated_at
		FROM messages
		WHERE id = $1
	`

	var model MessageModel
	if err := r.db.GetContext(ctx, &model, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("message not found")
		}
		return nil, err
	}

	// Convert to Message
	return modelToMessage(&model)
}

// GetMessageByExternalID retrieves a message by external ID
func (r *messageRepository) GetMessageByExternalID(ctx context.Context, externalID string) (*Message, error) {
	query := `
		SELECT id, phone_number, template_id, parameters, 
			order_id, customer_id, status, 
			error_message, external_id, created_at, updated_at
		FROM messages
		WHERE external_id = $1
	`

	var model MessageModel
	if err := r.db.GetContext(ctx, &model, query, externalID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("message not found")
		}
		return nil, err
	}

	// Convert to Message
	return modelToMessage(&model)
}

// ListMessages retrieves a list of messages
func (r *messageRepository) ListMessages(ctx context.Context, orderID, customerID, phoneNumber string, limit, offset int) ([]*Message, error) {
	// Build query
	query := `
		SELECT id, phone_number, template_id, parameters, 
			order_id, customer_id, status, 
			error_message, external_id, created_at, updated_at
		FROM messages
		WHERE 1=1
	`
	
	// Add filters
	args := []interface{}{}
	argIndex := 1

	if orderID != "" {
		query += " AND order_id = $" + utils.GetPlaceholderIndex(argIndex)
		args = append(args, orderID)
		argIndex++
	}

	if customerID != "" {
		query += " AND customer_id = $" + utils.GetPlaceholderIndex(argIndex)
		args = append(args, customerID)
		argIndex++
	}

	if phoneNumber != "" {
		query += " AND phone_number = $" + utils.GetPlaceholderIndex(argIndex)
		args = append(args, phoneNumber)
		argIndex++
	}

	// Add pagination
	query += " ORDER BY created_at DESC LIMIT $" + utils.GetPlaceholderIndex(argIndex) + " OFFSET $" + utils.GetPlaceholderIndex(argIndex+1)
	args = append(args, limit, offset)

	// Execute query
	var models []MessageModel
	if err := r.db.SelectContext(ctx, &models, query, args...); err != nil {
		return nil, err
	}

	// Convert to Messages
	messages := make([]*Message, 0, len(models))
	for _, model := range models {
		msg, err := modelToMessage(&model)
		if err != nil {
			r.logger.Error("Failed to convert model to message", "error", err)
			continue
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// UpdateMessageStatus updates the status of a message
func (r *messageRepository) UpdateMessageStatus(ctx context.Context, id int64, status, errorMessage, externalID string) error {
	query := `
		UPDATE messages
		SET status = $1, updated_at = $2
	`
	args := []interface{}{status, time.Now()}
	argIndex := 3

	// Add error message if provided
	if errorMessage != "" {
		query += ", error_message = $" + utils.GetPlaceholderIndex(argIndex)
		args = append(args, errorMessage)
		argIndex++
	}

	// Add external ID if provided
	if externalID != "" {
		query += ", external_id = $" + utils.GetPlaceholderIndex(argIndex)
		args = append(args, externalID)
		argIndex++
	}

	// Add where clause
	query += " WHERE id = $" + utils.GetPlaceholderIndex(argIndex)
	args = append(args, id)

	// Execute query
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

// Helper function to convert model to message
func modelToMessage(model *MessageModel) (*Message, error) {
	// Parse parameters JSON
	var parameters map[string]interface{}
	if err := json.Unmarshal([]byte(model.Parameters), &parameters); err != nil {
		return nil, err
	}

	// Create message
	message := &Message{
		ID:          model.ID,
		PhoneNumber: model.PhoneNumber,
		TemplateID:  model.TemplateID,
		Parameters:  parameters,
		Status:      model.Status,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	// Set nullable fields
	if model.OrderID.Valid {
		message.OrderID = model.OrderID.String
	}
	if model.CustomerID.Valid {
		message.CustomerID = model.CustomerID.String
	}
	if model.ErrorMessage.Valid {
		message.ErrorMessage = model.ErrorMessage.String
	}
	if model.ExternalID.Valid {
		message.ExternalID = model.ExternalID.String
	}

	return message, nil
}