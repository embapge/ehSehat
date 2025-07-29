package notification

import "ehSehat/notification-service/internal/notification/domain"

// type NotificationService interface {
// 	CreateNotification(notification *Notification) error
// 	UpdateNotification(notification *Notification) error
// 	// Delete(id string) error
// 	// GetByID(id string) (*Notification, error)
// 	// List(filter NotificationFilter) ([]*Notification, error)
// }

type notificationApp struct {
	repo domain.NotificationService
}

func NewNotificationApp(repo domain.NotificationService) *notificationApp {
	return &notificationApp{
		repo: repo,
	}
}

func (app *notificationApp) CreateNotification(notification *domain.Notification) error {
	return app.repo.CreateNotification(notification)
}

func (app *notificationApp) UpdateNotification(notification *domain.Notification) error {
	return app.repo.UpdateNotification(notification)
}
