package infra

import (
	"database/sql"
	"ehSehat/notification-service/internal/notification/domain"
)

type notificationMysql struct {
	db *sql.DB
}

func NewMySQLNotification(db *sql.DB) *notificationMysql {
	return &notificationMysql{db: db}
}

func (n *notificationMysql) Create(notification *domain.Notification) error {
	// Implement the logic to save a new notification to the MySQL database
	// This is a placeholder implementation
	return nil
}

func (n *notificationMysql) Update(notification *domain.Notification) error {
	// Implement the logic to update an existing notification in the MySQL database
	// This is a placeholder implementation
	return nil
}
