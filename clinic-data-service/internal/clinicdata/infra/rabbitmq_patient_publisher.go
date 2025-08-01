package infra

import (
	"encoding/json"
	"fmt"
	"time"

	"clinic-data-service/internal/clinicdata/domain"
	"ehSehat/libs/utils/rabbitmqown"
	amqp "github.com/rabbitmq/amqp091-go"
)

// rabbitPatientPublisher adalah implementasi production dari PatientEventPublisher
type rabbitPatientPublisher struct {
	ch *amqp.Channel
}

// NewRabbitPatientPublisher constructor
func NewRabbitPatientPublisher(ch *amqp.Channel) *rabbitPatientPublisher {
	return &rabbitPatientPublisher{ch: ch}
}

// PublishPatientCreated mengirim event PatientCreated ke AuthService via RabbitMQ dan menunggu userID balikan.
func (p *rabbitPatientPublisher) PublishPatientCreated(patient *domain.Patient) (string, error) {
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

	userBody := rabbitmqown.AuthRabbitBody{
		ID:    patient.ID,
		Name:  patient.Name,
		Email: patient.Email,
		Role:  "patient",
	}
	body, _ := json.Marshal(userBody)
	err = p.ch.Publish(
		"", "PatientCreated", false, false,
		amqp.Publishing{
			ContentType: "application/json",
			ReplyTo:     replyQueue.Name,
			Body:        body,
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %v", err)
	}

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
