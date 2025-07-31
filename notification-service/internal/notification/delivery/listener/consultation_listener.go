package listener

import (
	"ehSehat/libs/utils/rabbitmqown"
	"ehSehat/notification-service/internal/notification/domain"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type consultationListener struct {
	app domain.NotificationService
	ch  *amqp.Channel
}

func NewConsultationListener(app domain.NotificationService, ch *amqp.Channel) *consultationListener {
	return &consultationListener{app: app, ch: ch}
}

func (l *consultationListener) Start() {
	queueName := "TemanSehatNotification"
	_, err := rabbitmqown.DeclareQueue(l.ch, queueName)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}
	msgs, err := rabbitmqown.ConsumeQueue(l.ch, queueName)
	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ messages: %v", err)
	}
	go func() {
		for msg := range msgs {
			var payload rabbitmqown.NotificationPayload
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Printf("Failed to unmarshal RabbitMQ message: %v", err)
				continue
			}
			err := l.app.CreateNotification(&payload)
			if err != nil {
				log.Printf("Failed to send notification email: %v", err)
				continue
			}
			log.Printf("Received message")
		}
	}()
}
