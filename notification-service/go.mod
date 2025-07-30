module ehSehat/notification-service

go 1.24.4

require (
	ehSehat v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.9.3
	github.com/joho/godotenv v1.5.1
)

replace ehSehat => ../

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
)
