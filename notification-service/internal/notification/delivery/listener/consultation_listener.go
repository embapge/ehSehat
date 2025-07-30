package listener

import (
	"ehSehat/libs/utils/rabbitmqown"
	"ehSehat/notification-service/internal/notification/domain"
)

type consultationListener struct {
	app domain.NotificationService
}

func NewConsultationListener(app domain.NotificationService) *consultationListener {
	return &consultationListener{
		app: app,
	}
}

func (l *consultationListener) Handle(payload *rabbitmqown.RabbitPayload) error {
	switch payload.TemplateName {
	case "paymentCreated":
		l.app.CreateNotification(payload)
	}

	return nil
}
