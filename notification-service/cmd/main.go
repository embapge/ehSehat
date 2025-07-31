package main

import (
	"ehSehat/libs/utils/rabbitmqown"
	"ehSehat/notification-service/config"
	"ehSehat/notification-service/internal/notification/app"
	"ehSehat/notification-service/internal/notification/delivery/listener"
	"ehSehat/notification-service/internal/notification/infra"
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

	consultationListener := listener.NewConsultationListener(consultationApp, ch)
	consultationListener.Start()

	log.Println("Notification service is running...")
	select {}

}
