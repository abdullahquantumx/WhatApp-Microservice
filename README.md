# WhatsApp Messaging Microservice for Logistics Management

This microservice handles WhatsApp messaging for logistics operations, providing real-time updates to customers about their orders through both gRPC and HTTP interfaces.

## Features

- Send template-based WhatsApp messages
- Process incoming messages from customers
- gRPC API for high-performance integration
- REST webhook endpoint for WhatsApp callbacks
- Queue-based architecture for scalability
- Support for common logistics notifications:
  - Order confirmation
  - Shipment dispatched
  - Delivery ETAs
  - Delivery confirmation
  - Delay notifications
- Database persistence for message history
- Rate limiting to comply with WhatsApp policies

## Architecture

This microservice follows a clean, layered architecture:

- **Handler Layer**: gRPC endpoints and HTTP webhook handlers
- **Service Layer**: Business logic for message processing
- **Repository Layer**: Database operations for persistence
- **Queue Layer**: Asynchronous message processing with Kafka
- **WhatsApp Client**: Integration with Twilio's WhatsApp Business API

## Requirements

- Go 1.18+
- PostgreSQL 14+
- Kafka
- Protocol Buffers compiler (protoc)
- Twilio Account with WhatsApp Business API access

## Quick Start

### Using Docker Compose

1. Clone the repository
   ```
   git clone https://github.com/your-org/whatsapp-microservice.git
   cd whatsapp-microservice
   ```

2. Configure environment variables
   ```
   cp config/env.sample config/.env
   # Edit .env with your Twilio credentials
   ```

3. Start the services
   ```
   docker-compose up -d
   ```

4. The service will be available at:
   - HTTP webhook: http://localhost:8080/webhook
   - gRPC service: localhost:9090

### Manual Setup

1. Clone the repository
   ```
   git clone https://github.com/your-org/whatsapp-microservice.git
   cd whatsapp-microservice
   ```

2. Generate gRPC code
   ```
   protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      proto/whatsapp.proto
   ```

3. Configure environment variables
   ```
   cp config/env.sample config/.env
   # Edit .env with your configuration
   ```

4. Set up the database
   ```
   psql -U postgres -f db/init_db.sql
   ```

5. Run the application
   ```
   go run cmd/main.go
   ```

## API Endpoints

### gRPC API

The service exposes the following gRPC methods:

1. SendTemplateMessage - Send a template-based WhatsApp message
2. GetMessage - Retrieve a message by ID
3. ListMessages - List messages with filtering options

Example using grpcurl:

```bash
# Send a template message
grpcurl -d '{"phone_number": "+1234567890", "template_id": "order_confirmation", "parameters": {"order_id": "ORD-12345"}, "order_id": "ORD-12345", "customer_id": "CUST-6789"}' -plaintext localhost:9090 whatsapp.WhatsAppService/SendTemplateMessage

# Get a message by ID
grpcurl -d '{"message_id": 1}' -plaintext localhost:9090 whatsapp.WhatsAppService/GetMessage

# List messages
grpcurl -d '{"order_id": "ORD-12345", "limit": 10, "offset": 0}' -plaintext localhost:9090 whatsapp.WhatsAppService/ListMessages
```

### HTTP Webhook

```
POST /webhook
```

Used by Twilio to send delivery status updates and incoming messages.

## Development

### Project Structure

```
whatsapp-microservice/
│── cmd/
│   ├── main.go                 # Entry point of the application
│
│── config/
│   ├── config.go               # Loads environment variables & configurations
│   ├── env.sample              # Example .env file
│
│── proto/
│   ├── whatsapp.proto          # gRPC service definitions
│
│── internal/
│   ├── handler/                # Handles requests
│   │   ├── message_handler.go  # gRPC message handlers
│   │   ├── webhook_handler.go  # Webhook for WhatsApp API
│   │
│   ├── service/                # Business logic layer
│   │   ├── message_service.go  # Core logic for sending messages
│   │   ├── webhook_service.go  # Handles webhook events
│   │
│   ├── repository/             # Database interactions
│   │   ├── message_repository.go # Stores message logs
│   │
│   ├── queue/                  # Queue (Kafka)
│       ├── producer.go         # Publishes messages to the queue
│       ├── consumer.go         # Listens for messages
│
│── pkg/
│   ├── twilio/                 # API wrapper for Twilio
│   │   ├── client.go           # Handles WhatsApp API requests
│   │
│   ├── utils/                  # Utility functions
│       ├── logger.go           # Logs messages to console/file
│       ├── http_client.go      # Generic HTTP client wrapper
│       ├── helpers.go          # Helper functions
│
│── db/
│   ├── migrations/             # SQL migration scripts
│   ├── init_db.sql             # Initial database schema
│
│── test/
│   ├── message_service_test.go # Unit tests for message service
│   ├── queue_test.go           # Tests for Kafka consumer
│
│── Dockerfile                  # Dockerfile for containerization
│── docker-compose.yml          # Docker Compose for local testing
│── go.mod                      # Go module dependencies
│── go.sum                      # Dependency checksum
│── README.md                   # Project documentation
```

### Generating gRPC Code

To regenerate gRPC code after modifying the proto file:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/whatsapp.proto
```

### Running Tests

```
go test ./test/...
```

### Building the Docker Image

```
docker build -t whatsapp-service:latest .
```

## Security Considerations

- All WhatsApp API credentials are stored as environment variables
- Webhook signatures are validated to prevent spoofing
- Rate limiting is implemented to prevent abuse
- For production gRPC service, configure TLS

## Creating a gRPC Client

Here's an example of how to create a gRPC client in Go:

```go
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	
	pb "github.com/your-org/whatsapp-microservice/proto"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	// Create a client
	client := pb.NewWhatsAppServiceClient(conn)
	
	// Send a message
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	resp, err := client.SendTemplateMessage(ctx, &pb.SendTemplateMessageRequest{
		PhoneNumber: "+1234567890",
		TemplateId:  "order_confirmation",
		Parameters: map[string]string{
			"order_id": "ORD-12345",
		},
		OrderId:    "ORD-12345",
		CustomerId: "CUST-6789",
	})
	
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	
	log.Printf("Message sent with ID: %d", resp.MessageId)
}
```

## Deployment

### Kubernetes

1. Build and push the Docker image
   ```
   docker build -t your-registry.io/whatsapp-service:latest .
   docker push your-registry.io/whatsapp-service:latest
   ```

2. Apply Kubernetes manifests
   ```
   kubectl apply -f k8s/
   ```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests
5. Submit a pull request

## License

MIT