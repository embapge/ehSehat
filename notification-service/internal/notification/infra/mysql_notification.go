package infra

import (
	"database/sql"
	"ehSehat/libs/utils/rabbitmqown"
	"fmt"
	"time"
)

type notificationMysql struct {
	db *sql.DB
}

func NewMySQLNotification(db *sql.DB) *notificationMysql {
	return &notificationMysql{db: db}
}

func (n *notificationMysql) Create(notification *rabbitmqown.NotificationPayload) error {
	sentAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := n.db.Exec(`
		INSERT INTO notifications (channel, recipient, template_name, subject, body, source_service, reference_id, context, status, error_message, retry_count, sent_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, notification.Channel, notification.Recipient, notification.TemplateName, notification.Subject, notification.Body, notification.SourceService, notification.ReferenceID, notification.Context, notification.Status, notification.ErrorMessage, notification.RetryCount, sentAt)

	if err != nil {
		return err
	}

	id, err := n.db.Query(`
		SELECT id FROM notifications
		WHERE reference_id = ? AND recipient = ? AND channel = ?
		ORDER BY sent_at DESC LIMIT 1
	`, notification.ReferenceID, notification.Recipient, notification.Channel)
	if err != nil {
		return err
	}
	defer id.Close()

	var uuid string
	if id.Next() {
		if err := id.Scan(&uuid); err != nil {
			return err
		}
		notification.ID = uuid
		return nil
	}

	return fmt.Errorf("failed to create notification, no rows affected")
}

// Update notifications when sent then retry count still same, when failed update status into failed and retrycount++
func (n *notificationMysql) UpdateStatus(notificationID string, sent bool, errorMessage string) error {
	var err error
	var result sql.Result
	var status string

	if sent {
		status = "sent"
	} else {
		status = "failed"
	}

	if sent {
		result, err = n.db.Exec(`
			UPDATE notifications
			SET status = ?, error_message = '', sent_at = NOW(), updated_at = NOW()
			WHERE id = ?
		`, status, notificationID)
	} else {
		result, err = n.db.Exec(`
			UPDATE notifications
			SET status = ?, error_message = ?, retry_count = retry_count + 1, updated_at = NOW()
			WHERE id = ?
		`, status, errorMessage, notificationID)
	}

	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no notification updated")
	}

	return nil
}
