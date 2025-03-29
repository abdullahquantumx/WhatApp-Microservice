// cmd/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/your-org/whatsapp-microservice/config"
	"github.com/your-org/whatsapp-microservice/internal/handler"
	"github.com/your-org/whatsapp-microservice/internal/queue"
	"github.com/your-org/whatsapp-microservice/internal/repository"
	"github.com/your-org/whatsapp-microservice/internal/service"
	"github.com/your-org/whatsapp-microservice/pkg/twilio"
	"github.com/your-org/whatsapp-microservice/pkg/utils"
	pb "github.com/your-org/whatsapp-microservice/proto"
)

func main() {
	// Initialize logger
	logger := utils.NewLogger()
	logger.Info("Starting WhatsApp Microservice")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", "error", err)
	}

	// Connect to database
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", "error", err)
	}
	defer db.Close()

	// Initialize repository
	messageRepo := repository.NewMessageRepository(db, logger)

	// Initialize WhatsApp client
	whatsappClient := twilio.NewClient(cfg.TwilioAccountSID, cfg.TwilioAuthToken, cfg.TwilioFromNumber, logger)

	// Initialize message queue
	messageProducer, err := queue.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic, logger)
	if err != nil {
		logger.Fatal("Failed to initialize Kafka producer", "error", err)
	}
	defer messageProducer.Close()

	// Initialize consumer
	messageConsumer, err := queue.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID, logger)
	if err != nil {
		logger.Fatal("Failed to initialize Kafka consumer", "error", err)
	}

	// Initialize services
	messageService := service.NewMessageService(messageRepo, whatsappClient, messageProducer, logger)
	webhookService := service.NewWebhookService(messageRepo, messageProducer, logger)

	// Start consumer
	go func() {
		logger.Info("Starting message consumer")
		messageConsumer.Consume(context.Background(), messageService.ProcessQueueMessage)
	}()

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
		if err != nil {
			logger.Fatal("Failed to listen for gRPC", "error", err)
		}

		grpcServer := grpc.NewServer()
		grpcHandler := handler.NewGrpcMessageHandler(messageService, logger)
		pb.RegisterWhatsAppServiceServer(grpcServer, grpcHandler)

		// Register reflection service on gRPC server (for debugging)
		if cfg.Environment != "production" {
			reflection.Register(grpcServer)
		}

		logger.Info("Starting gRPC server", "port", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve gRPC", "error", err)
		}
	}()

	// Initialize HTTP server for webhooks
	router := gin.Default()

	// Register middleware
	router.Use(gin.Recovery())
	router.Use(utils.RequestLogger(logger))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// Webhook handler
	webhookHandler := handler.NewWebhookHandler(webhookService, logger)
	router.POST("/webhook", webhookHandler.HandleWebhook)

	// Start HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start HTTP server", "error", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("HTTP server forced to shutdown", "error", err)
	}

	// Close consumer
	if err := messageConsumer.Close(); err != nil {
		logger.Error("Failed to close consumer", "error", err)
	}

	logger.Info("Server exited gracefully")
}