package domain

// Here notification service interface:
type NotificationService interface {
	CreateNotification(notification *Notification) error
	UpdateNotification(notification *Notification) error
	// Delete(id string) error
	// GetByID(id string) (*Notification, error)
	// List(filter NotificationFilter) ([]*Notification, error)
}
