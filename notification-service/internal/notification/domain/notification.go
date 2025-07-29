package domain

import (
	"encoding/json"
	"time"
)

// Here notification columns:
// CREATE TABLE notifications (
//     id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
//     channel VARCHAR(30) NOT NULL,
//     recipient VARCHAR(255) NOT NULL,
//     template_name VARCHAR(100) NOT NULL,
//     subject TEXT,
//     body TEXT NOT NULL,
//     source_service VARCHAR(50) NOT NULL,
//     reference_id CHAR(36),
//     context JSON,
//     status VARCHAR(20) NOT NULL DEFAULT 'pending',
//     error_message TEXT,
//     retry_count INT DEFAULT 0,
//     sent_at DATETIME(3),
//     created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
//     updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
// );

type Notification struct {
	ID            string          `json:"id"`
	Channel       string          `json:"channel"`        // e.g., "email", "sms", "push"
	Recipient     string          `json:"recipient"`      // e.g., email address, phone number, or device token
	TemplateName  string          `json:"template_name"`  // Name of the template to use for the notification
	Subject       string          `json:"subject"`        // Subject of the notification (optional, e.g., for emails)
	Body          string          `json:"body"`           // Body of the notification (can be HTML or plain text)
	SourceService string          `json:"source_service"` // Service that generated the notification (e.g., "consultation", "payment")
	ReferenceID   string          `json:"reference_id"`   // ID of the related entity (e.g., consultation ID, payment ID)
	Context       json.RawMessage `json:"context"`        // Additional context for the notification, can be any JSON structure
	Status        string          `json:"status"`         // e.g., "pending", "sent", "failed"
	ErrorMessage  string          `json:"error_message"`  // Error message if the notification failed to send
	RetryCount    int             `json:"retry_count"`    // Number of retry attempts for sending the notification
	SentAt        time.Time       `json:"sent_at"`        // Timestamp when the notification was sent (if applicable)
	CreatedAt     time.Time       `json:"created_at"`     // Timestamp when the notification was created
	UpdatedAt     time.Time       `json:"updated_at"`     // Timestamp when the notification was last updated
}
