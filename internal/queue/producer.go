// internal/queue/producer.go
package queue

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
)

// Producer defines the interface for message producers
type Producer interface {
	Produce(ctx context.Context, value []byte) error
	Close() error
}

// kafkaProducer implements Producer using Kafka
type kafkaProducer struct {
	writer *kafka.Writer
	logger utils.Logger
}

// NewProducer creates a new Kafka producer
func NewProducer(brokers []string, topic string, logger utils.Logger) (Producer, error) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
	}

	return &kafkaProducer{
		writer: writer,
		logger: logger,
	}, nil
}

// Produce sends a message to Kafka
func (p *kafkaProducer) Produce(ctx context.Context, value []byte) error {
	msg := kafka.Message{
		Value: value,
		Time:  time.Now(),
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		p.logger.Error("Failed to write message to Kafka", "error", err)
		return err
	}

	return nil
}

// Close closes the Kafka writer
func (p *kafkaProducer) Close() error {
	return p.writer.Close()
}