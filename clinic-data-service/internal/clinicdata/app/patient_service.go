package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"ehSehat/libs/utils/rabbitmqown"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Interface yang didefinisikan di domain service
type PatientService interface {
	Create(patient *domain.Patient) (*domain.Patient, error)
	GetByID(id string) (*domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(patient *domain.Patient) (*domain.Patient, error)
	Delete(id string) error
}

// Implementasi struct service
type patientService struct {
	repo domain.PatientRepository
	ch   *amqp.Channel
}

// Constructor untuk service
func NewPatientService(r domain.PatientRepository, ch *amqp.Channel) PatientService {
	return &patientService{repo: r, ch: ch}
}

// Create inserts a new patient after validating required fields and email uniqueness
func (s *patientService) Create(p *domain.Patient) (*domain.Patient, error) {
	log.Printf("DEBUG create validation: Name='%s' Email='%s' BirthDate='%s' Gender='%s'",
		p.Name, p.Email, p.BirthDate, p.Gender,
	)

	// Validasi field wajib
	if strings.TrimSpace(p.Name) == "" ||
		strings.TrimSpace(p.Email) == "" ||
		strings.TrimSpace(p.BirthDate) == "" ||
		strings.TrimSpace(p.Gender) == "" {
		return nil, ErrMissingFields
	}

	// Validasi email unik
	existing, err := s.repo.GetByEmail(p.Email)
	if err != nil {
		log.Printf("ERROR checking existing email: %v", err)
		return nil, ErrInternal
	}
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	// Buatkan command reply rabbitMQ untuk membuat data user terlebih dahulu dengan menyertakan domain.Doctor dan mengembalikan userId
	replyQueue, err := s.ch.QueueDeclare(
		"",                            // empty name, RabbitMQ akan generate nama unik
		false, true, true, false, nil, // exclusive, auto-delete
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare reply queue: %v", err)
	}

	msgs, err := s.ch.Consume(
		replyQueue.Name, "", true, false, false, false, nil, // consumer tag kosong, non-exclusive
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume reply queue: %v", err)
	}

	userBody := rabbitmqown.AuthRabbitBody{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
		Role:  "patient",
	}
	body, _ := json.Marshal(userBody)
	err = s.ch.Publish(
		"",               // default exchange
		"PatientCreated", // queue name
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			ReplyTo:     replyQueue.Name,
			Body:        body,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to publish message: %v", err)
	}

	timeout := time.After(5 * time.Second)
	var resp rabbitmqown.AuthRabbitBody
	for {
		select {
		case d := <-msgs:
			if err := json.Unmarshal(d.Body, &resp); err != nil {
				return nil, fmt.Errorf("failed to parse response: %v", err)
			}

			if resp.ID == "" {
				return nil, fmt.Errorf("auth-service error: empty user_id")
			}

			p.UserID = &resp.ID
			patientNew, err := s.repo.Create(p)
			if err != nil {
				log.Printf("ERROR creating patient: %v", err)
				return nil, fmt.Errorf("failed to create patient: %v", err)
			}

			patientContext, _ := json.Marshal(patientNew)
			payload := rabbitmqown.NotificationPayload{
				Channel: "email",
				// Recipient:     consultation.Patient.ID, // Assuming recipient is the patient ID
				Recipient:     "baratagusti.bg@gmail.com", // Assuming recipient is the patient ID
				TemplateName:  "patientCreated",           // Example template name
				Subject:       "Patient Created!",
				Body:          fmt.Sprintf("Your patient profile has been created. Please login into %v to consultate, queue and make appointment with email: %v, password: temansehat", os.Getenv("APP_URL"), patientNew.Email),
				SourceService: "clinicDataService",
				Context:       patientContext, // Additional context can be added here if needed
				Status:        "pending",      // Initial status
				ErrorMessage:  "",             // No error message initially
				RetryCount:    0,              // Initial retry count
			}

			payloadBytes, _ := json.Marshal(payload)

			err = s.ch.Publish(
				"",                        // default exchange
				os.Getenv("RABBIT_QUEUE"), // routing key (queue name)
				false,                     // mandatory
				false,                     // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        payloadBytes,
				},
			)

			if err != nil {
				return nil, err
			}
			return patientNew, nil
		case <-timeout:
			return nil, fmt.Errorf("timeout waiting for auth-service")
		}
	}

}

// GetByID returns a patient by ID
func (s *patientService) GetByID(id string) (*domain.Patient, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrMissingID
	}
	return s.repo.GetByID(id)
}

// GetAll returns all patients
func (s *patientService) GetAll() ([]domain.Patient, error) {
	return s.repo.GetAll()
}

// Update modifies an existing patient
func (s *patientService) Update(p *domain.Patient) (*domain.Patient, error) {
	if strings.TrimSpace(p.ID) == "" {
		return nil, ErrMissingID
	}
	return s.repo.Update(p)
}

// Delete removes a patient by ID
func (s *patientService) Delete(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrMissingID
	}
	return s.repo.Delete(id)
}
