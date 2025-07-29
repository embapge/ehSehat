package listener

import "ehSehat/notification-service/internal/notification/domain"

type consultationListener struct {
	// Add fields as necessary for the listener
	payload *domain.Notification
	app     domain.NotificationService
}

func NewConsultationListener(payload *domain.Notification, app domain.NotificationService) *consultationListener {
	return &consultationListener{
		payload: payload,
		app:     app,
	}
}

func (l *consultationListener) Handle() error {

	return nil
}
