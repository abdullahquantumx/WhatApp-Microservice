// internal/queue/producer.go
package queue

import (
    "context"
    "time"

    "github.com/segmentio/kafka-go"
    "messaging-microservice/pkg/utils"
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

// WriterCreator is a function type for creating Kafka writers
// Used for testing to inject mocks
type WriterCreator func(brokers []string, topic string, logger utils.Logger) (interface{}, error)

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

// NewProducerWithWriter creates a new producer with a custom writer creator
// This function is primarily used for testing
func NewProducerWithWriter(brokers []string, topic string, logger utils.Logger, writerCreator WriterCreator) (Producer, error) {
    writer, err := writerCreator(brokers, topic, logger)
    if err != nil {
        return nil, err
    }

    return &kafkaProducer{
        writer: writer.(*kafka.Writer),
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