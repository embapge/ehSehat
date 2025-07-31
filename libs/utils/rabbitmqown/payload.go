package rabbitmqown

import (
	"encoding/json"
	"time"
)

type NotificationPayload struct {
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
