package domain

import "ehSehat/libs/utils/rabbitmqown"

// Here notification service interface:
type NotificationService interface {
	CreateNotification(notification *rabbitmqown.NotificationPayload) error
	UpdateStatusNotification(notificationID string, sent bool, errorMessage string) error
	// Delete(id string) error
	// GetByID(id string) (*rabbitmqown.NotificationPayload, error)
	// List(filter NotificationFilter) ([]*rabbitmqown.NotificationPayload, error)
}
