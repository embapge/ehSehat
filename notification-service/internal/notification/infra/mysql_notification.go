package infra

import (
	"database/sql"
	"ehSehat/libs/utils/rabbitmqown"
	"fmt"
)

type notificationMysql struct {
	db *sql.DB
}

func NewMySQLNotification(db *sql.DB) *notificationMysql {
	return &notificationMysql{db: db}
}

// *domain.Notification
// ID            string          `json:"id"`
// 	Channel       string          `json:"channel"`        // e.g., "email", "sms", "push"
// 	Recipient     string          `json:"recipient"`      // e.g., email address, phone number, or device token
// 	TemplateName  string          `json:"template_name"`  // Name of the template to use for the notification
// 	Subject       string          `json:"subject"`        // Subject of the notification (optional, e.g., for emails)
// 	Body          string          `json:"body"`           // Body of the notification (can be HTML or plain text)
// 	SourceService string          `json:"source_service"` // Service that generated the notification (e.g., "consultation", "payment")
// 	ReferenceID   string          `json:"reference_id"`   // ID of the related entity (e.g., consultation ID, payment ID)
// 	Context       json.RawMessage `json:"context"`        // Additional context for the notification, can be any JSON structure
// 	Status        string          `json:"status"`         // e.g., "pending", "sent", "failed"
// 	ErrorMessage  string          `json:"error_message"`  // Error message if the notification failed to send
// 	RetryCount    int             `json:"retry_count"`    // Number of retry attempts for sending the notification
// 	SentAt        time.Time       `json:"sent_at"`        // Timestamp when the notification was sent (if applicable)
// 	CreatedAt     time.Time       `json:"created_at"`     // Timestamp when the notification was created
// 	UpdatedAt     time.Time       `json:"updated_at"`     // Timestamp when the notification was last updated

func (n *notificationMysql) Create(notification *rabbitmqown.RabbitPayload) error {
	// Implement the logic to save a new notification to the MySQL database
	// This is a placeholder implementation
	result, err := n.db.Exec(`
		INSERT INTO notifications (channel, recipient, template_name, subject, body, source_service, reference_id, context, status, error_message, retry_count, sent_at, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, notification.Channel, notification.Recipient, notification.TemplateName, notification.Subject, notification.Body, notification.SourceService, notification.ReferenceID, notification.Context, notification.Status, notification.ErrorMessage, notification.RetryCount, notification.SentAt, notification.CreatedAt, notification.UpdatedAt)

	if err != nil {
		return err
	}

	num, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if num <= 0 {
		return fmt.Errorf("failed to create notification, no rows affected")
	}

	notification.ID = fmt.Sprintf("%d", num)
	return nil
}

// Update notifications when sent then retry count still same, when failed update status into failed and retrycount++
func (n *notificationMysql) UpdateStatus(notificationID string, sent bool, errorMessage string) error {
	var (
		status string
		query  string
		args   []interface{}
	)
	if sent {
		status = "sent"
		query = `
			UPDATE notifications
			SET status = ?, error_message = ?, sent_at = NOW(), updated_at = NOW()
			WHERE id = ?
		`
		args = []interface{}{status, "", notificationID}
	} else {
		status = "failed"
		query = `
			UPDATE notifications
			SET status = ?, error_message = ?, retry_count = retry_count + 1, updated_at = NOW()
			WHERE id = ?
		`
		args = []interface{}{status, errorMessage, notificationID}
	}

	result, err := n.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no notification updated")
	}
	return nil
}
