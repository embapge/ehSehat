package domain

type NotificationRepository interface {
	// SaveNotification saves a new notification to the repository.
	Create(notification *Notification) error
	Update(notification *Notification) error

	// // GetNotification retrieves a notification by its ID.
	// GetNotification(id string) (*Notification, error)

	// // UpdateNotification updates an existing notification.

	// // DeleteNotification removes a notification from the repository.
	// DeleteNotification(id string) error
	// // ListNotifications retrieves all notifications, optionally filtered by channel.
	// ListNotifications(channel string) ([]*Notification, error)
}
