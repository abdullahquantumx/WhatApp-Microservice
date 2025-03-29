-- db/init_db.sql
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(50) NOT NULL,
    template_id VARCHAR(50) NOT NULL,
    parameters TEXT NOT NULL,
    order_id VARCHAR(50),
    customer_id VARCHAR(50),
    status VARCHAR(20) NOT NULL,
    error_message TEXT,
    external_id VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_messages_phone_number ON messages(phone_number);
CREATE INDEX IF NOT EXISTS idx_messages_order_id ON messages(order_id);
CREATE INDEX IF NOT EXISTS idx_messages_customer_id ON messages(customer_id);
CREATE INDEX IF NOT EXISTS idx_messages_external_id ON messages(external_id);
CREATE INDEX IF NOT EXISTS idx_messages_status ON messages(status);
CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at);

-- Create a table for message conversations
CREATE TABLE IF NOT EXISTS conversations (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(50) NOT NULL,
    customer_id VARCHAR(50),
    last_message_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_conversations_phone_number ON conversations(phone_number);
CREATE INDEX IF NOT EXISTS idx_conversations_customer_id ON conversations(customer_id);
CREATE INDEX IF NOT EXISTS idx_conversations_status ON conversations(status);

-- db/migrations/001_initial_schema.up.sql
-- Initial schema is already defined in init_db.sql

-- db/migrations/001_initial_schema.down.sql
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS conversations;

-- db/migrations/002_add_templates.up.sql
-- Table to store message templates
CREATE TABLE IF NOT EXISTS templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    content TEXT NOT NULL,
    parameters JSONB NOT NULL DEFAULT '[]',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Insert default templates
INSERT INTO templates (name, description, content, parameters) VALUES
('order_confirmation', 'Sent when an order is confirmed', 'Your order {{order_id}} has been confirmed. Thank you for your purchase!', '[{"name": "order_id", "type": "string", "required": true}]'),
('shipment_dispatched', 'Sent when an order is dispatched', 'Your order {{order_id}} has been dispatched! Track your package with tracking ID {{tracking_id}}. Estimated delivery: {{estimated_delivery}}', '[{"name": "order_id", "type": "string", "required": true}, {"name": "tracking_id", "type": "string", "required": true}, {"name": "estimated_delivery", "type": "string", "required": true}]'),
('delivery_eta', 'Sent to provide delivery ETA', 'Your order {{order_id}} is scheduled for delivery on {{estimated_delivery}}', '[{"name": "order_id", "type": "string", "required": true}, {"name": "estimated_delivery", "type": "string", "required": true}]'),
('delivery_confirmation', 'Sent when an order is delivered', 'Your order {{order_id}} has been delivered. Thank you for choosing our service!', '[{"name": "order_id", "type": "string", "required": true}]'),
('delay_notification', 'Sent when a delivery is delayed', 'We''re sorry, but your order {{order_id}} has been delayed due to {{reason}}. New estimated delivery: {{new_estimated_delivery}}', '[{"name": "order_id", "type": "string", "required": true}, {"name": "reason", "type": "string", "required": true}, {"name": "new_estimated_delivery", "type": "string", "required": true}]');

-- db/migrations/002_add_templates.down.sql
DROP TABLE IF EXISTS templates;

-- db/migrations/003_add_rate_limiting.up.sql
-- Table to track rate limiting
CREATE TABLE IF NOT EXISTS rate_limits (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(50) NOT NULL UNIQUE,
    message_count INTEGER NOT NULL DEFAULT 0,
    last_reset_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_rate_limits_phone_number ON rate_limits(phone_number);

-- db/migrations/003_add_rate_limiting.down.sql
DROP TABLE IF EXISTS rate_limits;