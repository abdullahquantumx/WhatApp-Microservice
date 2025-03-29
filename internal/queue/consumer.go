// internal/queue/consumer.go
package queue

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// MessageHandler is a function to handle consumed messages
type MessageHandler func(context.Context, []byte) error

// Consumer defines the interface for message consumers
type Consumer interface {
	Consume(ctx context.Context, handler MessageHandler) error
	Close() error
}

// kafkaConsumer implements Consumer using Kafka
type kafkaConsumer struct {
	reader *kafka.Reader
	logger utils.Logger
}

// NewConsumer creates a new Kafka consumer
func NewConsumer(brokers []string, topic, groupID string, logger utils.Logger) (Consumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        groupID,
		MinBytes:       10e3,   // 10KB
		MaxBytes:       10e6,   // 10MB
		MaxWait:        time.Second,
		StartOffset:    kafka.FirstOffset,
		CommitInterval: time.Second,
	})

	return &kafkaConsumer{
		reader: reader,
		logger: logger,
	}, nil
}

// Consume consumes messages from Kafka
func (c *kafkaConsumer) Consume(ctx context.Context, handler MessageHandler) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			// Check if context was canceled
			if ctx.Err() != nil {
				return ctx.Err()
			}

			c.logger.Error("Failed to read message from Kafka", "error", err)
			continue
		}

		c.logger.Info("Received message from Kafka", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset)

		// Handle message
		if err := handler(ctx, msg.Value); err != nil {
			c.logger.Error("Failed to handle message", "error", err)
			// Continue processing other messages even if one fails
			// In a production system, you might want to handle retries, DLQ, etc.
		}
	}
}

// Close closes the Kafka reader
func (c *kafkaConsumer) Close() error {
	return c.reader.Close()
}