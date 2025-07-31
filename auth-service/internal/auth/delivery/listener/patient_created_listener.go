package listener

import (
	"context"
	"encoding/json"
	"log"

	"ehSehat/auth-service/internal/auth/app"
	"ehSehat/libs/utils/rabbitmqown"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PatientCreatedListener listens for PatientCreated queue and processes user creation

type PatientCreatedListener struct {
	App *app.AuthApp
	Ch  *amqp.Channel
}

func NewPatientCreatedListener(app *app.AuthApp, ch *amqp.Channel) *PatientCreatedListener {
	return &PatientCreatedListener{App: app, Ch: ch}
}

func (l *PatientCreatedListener) Start() {
	queueName := "PatientCreated"
	_, err := rabbitmqown.DeclareQueue(l.Ch, queueName)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	msgs, err := l.Ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ messages: %v", err)
	}

	go func() {
		for msg := range msgs {
			var body rabbitmqown.AuthRabbitBody
			if err := json.Unmarshal(msg.Body, &body); err != nil {
				log.Printf("Failed to unmarshal RabbitMQ message: %v", err)
				continue
			}

			defaultPassword := "temansehat" // Ganti sesuai kebutuhan
			user, err := l.App.Register(
				context.Background(),
				body.Name,
				body.Email,
				defaultPassword,
				body.Role,
			)
			if err != nil {
				log.Println("Failed to create user")
				continue
			}

			userJSON, err := json.Marshal(user)
			if err != nil {
				log.Printf("Failed to marshal user: %v", err)
				continue
			}

			if msg.ReplyTo != "" {
				err = l.Ch.Publish(
					"",          // default exchange
					msg.ReplyTo, // reply queue
					false, false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        userJSON,
					},
				)
				if err != nil {
					log.Printf("failed to reply: %v", err)
					continue
				}
			}

			log.Printf("Processed PatientCreated: %s", msg.Body)
		}
	}()
}
