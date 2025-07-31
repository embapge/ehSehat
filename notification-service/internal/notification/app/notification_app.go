package app

import (
	"ehSehat/libs/utils/rabbitmqown"
	"ehSehat/notification-service/internal/notification/domain"
	"fmt"
	"log"
	"net/smtp"
)

type notificationApp struct {
	repo domain.NotificationRepository
}

func NewNotificationApp(repo domain.NotificationRepository) *notificationApp {
	return &notificationApp{
		repo: repo,
	}
}

func (app *notificationApp) CreateNotification(notification *rabbitmqown.NotificationPayload) error {
	err := app.repo.Create(notification)
	if err != nil {
		log.Printf("Failed to create notification: %v", err)
	}

	from := "mohammadbarata.mb@gmail.com"
	password := "gehb nspg eamf okzt"
	to := []string{notification.Recipient}
	msg := []byte(fmt.Sprintf("To: %v\r\n"+
		"Subject: %v\r\n"+
		"\r\n"+
		"%v\r\n", notification.Recipient, notification.Subject, notification.Body))

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, to, msg)

	if err != nil {
		app.UpdateStatusNotification(notification.ID, false, err.Error())
		log.Fatal(err, "failed update status notification")
	}

	app.UpdateStatusNotification(notification.ID, true, "")
	return nil
}

func (app *notificationApp) UpdateStatusNotification(notificationID string, sent bool, errorMessage string) error {
	return app.repo.UpdateStatus(notificationID, sent, errorMessage)
}
