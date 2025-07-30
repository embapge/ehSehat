package main

import (
	"ehSehat/libs/utils/rabbitmqown"
	"ehSehat/notification-service/config"
	"ehSehat/notification-service/internal/notification/app"
	"ehSehat/notification-service/internal/notification/delivery/listener"
	"ehSehat/notification-service/internal/notification/infra"
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := config.MySQLInit()
	if db == nil {
		log.Fatal("Failed to initialize MySQL connection")
	}

	consultationInfra := infra.NewMySQLNotification(db)
	consultationApp := app.NewNotificationApp(consultationInfra)
	consultationHandler := listener.NewConsultationListener(consultationApp)

	conn, ch, err := rabbitmqown.InitRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
		if err := ch.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ channel: %v", err)
		}
	}()

	queueName := "consultation_notifications"
	_, err = rabbitmqown.DeclareQueue(ch, queueName)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	msgs, err := rabbitmqown.ConsumeQueue(ch, queueName)
	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ messages: %v", err)
	}

	go func() {
		for msg := range msgs {
			var payload rabbitmqown.RabbitPayload
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Printf("Failed to unmarshal RabbitMQ message: %v", err)
				continue
			} else {
				consultationHandler.Handle(&payload)
				log.Printf("Received message: %s", msg.Body)
			}
		}
	}()

	log.Println("Notification service is running...")
	select {}

}
