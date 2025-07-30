package domain

import "ehSehat/libs/utils/rabbitmqown"

// Here notification service interface:
type NotificationService interface {
	CreateNotification(notification *rabbitmqown.RabbitPayload) error
	UpdateStatusNotification(notificationID string, sent bool, errorMessage string) error
	// Delete(id string) error
	// GetByID(id string) (*rabbitmqown.RabbitPayload, error)
	// List(filter NotificationFilter) ([]*rabbitmqown.RabbitPayload, error)
}
