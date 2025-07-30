package domain

import "ehSehat/libs/utils/rabbitmqown"

type NotificationRepository interface {
	// SaveNotification saves a new notification to the repository.
	Create(notification *rabbitmqown.RabbitPayload) error
	UpdateStatus(notificationID string, sent bool, errorMessage string) error

	// // GetNotification retrieves a notification by its ID.
	// GetNotification(id string) (*rabbitmqown.RabbitPayload, error)

	// // UpdateNotification updates an existing notification.

	// // DeleteNotification removes a notification from the repository.
	// DeleteNotification(id string) error
	// // ListNotifications retrieves all notifications, optionally filtered by channel.
	// ListNotifications(channel string) ([]*rabbitmqown.RabbitPayload, error)
}
