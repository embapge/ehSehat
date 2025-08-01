package infra

import (
	"encoding/json"
	"fmt"
	"time"

	"clinic-data-service/internal/clinicdata/domain"
	"ehSehat/libs/utils/rabbitmqown"

	amqp "github.com/rabbitmq/amqp091-go"
)

// rabbitDoctorPublisher adalah implementasi production dari DoctorEventPublisher
type rabbitDoctorPublisher struct {
	ch *amqp.Channel
}

// NewRabbitDoctorPublisher constructor
func NewRabbitDoctorPublisher(ch *amqp.Channel) *rabbitDoctorPublisher {
	return &rabbitDoctorPublisher{ch: ch}
}

// PublishDoctorCreated mengirim event DoctorCreated ke AuthService via RabbitMQ dan menunggu userID balikan.
func (p *rabbitDoctorPublisher) PublishDoctorCreated(doctor *domain.Doctor) (string, error) {
	if p.ch == nil || p.ch.IsClosed() {
		return "", fmt.Errorf("rabbitmq channel is not open")
	}

	// Buat reply queue unik (auto delete, exclusive)
	replyQueue, err := p.ch.QueueDeclare(
		"", false, true, true, false, nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to declare reply queue: %v", err)
	}

	msgs, err := p.ch.Consume(
		replyQueue.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to consume reply queue: %v", err)
	}

	// Siapkan payload user yang dikirim ke AuthService
	userBody := rabbitmqown.AuthRabbitBody{
		ID:    doctor.ID,
		Name:  doctor.Name,
		Email: doctor.Email,
		Role:  "doctor",
	}
	body, _ := json.Marshal(userBody)
	err = p.ch.Publish(
		"", "DoctorCreated", false, false,
		amqp.Publishing{
			ContentType: "application/json",
			ReplyTo:     replyQueue.Name,
			Body:        body,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %v", err)
	}

	// Tunggu reply dari AuthService (timeout 5 detik)
	timeout := time.After(5 * time.Second)
	var resp rabbitmqown.AuthRabbitBody
	for {
		select {
		case d := <-msgs:
			if err := json.Unmarshal(d.Body, &resp); err != nil {
				return "", fmt.Errorf("failed to parse response: %v", err)
			}
			if resp.ID == "" {
				return "", fmt.Errorf("auth-service error: empty user_id")
			}
			return resp.ID, nil
		case <-timeout:
			return "", fmt.Errorf("timeout waiting for auth-service")
		}
	}
}
