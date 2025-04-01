// config/config.go
package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the service
type Config struct {
	// Server configuration
	HTTPPort        string
	GRPCPort        string
	Environment     string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration

	// Database configuration
	DatabaseURL          string
	DatabaseMaxOpenConns int
	DatabaseMaxIdleConns int

	// Meta WhatsApp configuration
	MetaPhoneNumberID string
	MetaAccessToken   string
	MetaAppSecret     string
	MetaVerifyToken   string

	// Kafka configuration
	KafkaBrokers []string
	KafkaTopic   string
	KafkaGroupID string

	// JWT configuration
	JWTSecret     string
	JWTExpiration time.Duration

	// Template IDs for WhatsApp
	OrderConfirmationTemplateID    string
	ShipmentDispatchedTemplateID   string
	DeliveryETATemplateID          string
	DeliveryConfirmationTemplateID string
	DelayNotificationTemplateID    string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		HTTPPort:        getEnv("HTTP_PORT", "8080"),
		GRPCPort:        getEnv("GRPC_PORT", "9090"),
		Environment:     getEnv("ENVIRONMENT", "development"),
		ReadTimeout:     getEnvAsDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
		ShutdownTimeout: getEnvAsDuration("SHUTDOWN_TIMEOUT", 10*time.Second),

		DatabaseURL:          getEnv("DATABASE_URL", ""),
		DatabaseMaxOpenConns: getEnvAsInt("DATABASE_MAX_OPEN_CONNS", 20),
		DatabaseMaxIdleConns: getEnvAsInt("DATABASE_MAX_IDLE_CONNS", 5),

		MetaPhoneNumberID: getEnv("META_PHONE_NUMBER_ID", ""),
		MetaAccessToken:   getEnv("META_ACCESS_TOKEN", ""),
		MetaAppSecret:     getEnv("META_APP_SECRET", ""),
		MetaVerifyToken:   getEnv("META_VERIFY_TOKEN", ""),

		KafkaBrokers: strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ","),
		KafkaTopic:   getEnv("KAFKA_TOPIC", "whatsapp-messages"),
		KafkaGroupID: getEnv("KAFKA_GROUP_ID", "whatsapp-microservice"),

		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiration: getEnvAsDuration("JWT_EXPIRATION", 24*time.Hour),

		OrderConfirmationTemplateID:    getEnv("ORDER_CONFIRMATION_TEMPLATE_ID", ""),
		ShipmentDispatchedTemplateID:   getEnv("SHIPMENT_DISPATCHED_TEMPLATE_ID", ""),
		DeliveryETATemplateID:          getEnv("DELIVERY_ETA_TEMPLATE_ID", ""),
		DeliveryConfirmationTemplateID: getEnv("DELIVERY_CONFIRMATION_TEMPLATE_ID", ""),
		DelayNotificationTemplateID:    getEnv("DELAY_NOTIFICATION_TEMPLATE_ID", ""),
	}

	// Validate required configuration
	if cfg.DatabaseURL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}

	if cfg.MetaPhoneNumberID == "" || cfg.MetaAccessToken == "" {
		return nil, errors.New("META_PHONE_NUMBER_ID and META_ACCESS_TOKEN are required")
	}

	return cfg, nil
}

// Helper functions to read environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}