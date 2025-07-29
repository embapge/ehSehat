package main

import (
	"ehSehat/notification-service/config"
	"ehSehat/notification-service/internal/notification/domain"
	"encoding/json"
	"log"
	"os"

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
	conn, ch, err := config.InitRabbitMQ()
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

	msgs, err := config.ConsumeQueue(ch, os.Getenv("RABBITMQ_QUEUE"))
	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ messages: %v", err)
	}

	go func() {
		for msg := range msgs {
			var payload domain.Notification
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Printf("Failed to unmarshal RabbitMQ message: %v", err)
				continue
			} else {
				// listener
				log.Printf("Received message: %s", msg.Body)
			}
		}
	}()
	log.Println("Notification service is running...")
}
