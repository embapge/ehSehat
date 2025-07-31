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

type UserBodyRabbit struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

// DoctorService interface defines use-case logic for doctor
type DoctorService interface {
	Create(doctor *domain.Doctor) (*domain.Doctor, error)
	GetByID(id string) (*domain.Doctor, error)
	GetAll() ([]domain.Doctor, error)
	Update(doctor *domain.Doctor) (*domain.Doctor, error)
	Delete(id string) error
}

// doctorService implements DoctorService
type doctorService struct {
	repo domain.DoctorRepository
	ch   *amqp.Channel
}

// NewDoctorService is the constructor
func NewDoctorService(r domain.DoctorRepository, ch *amqp.Channel) DoctorService {
	return &doctorService{repo: r, ch: ch}
}

// Create inserts a new doctor after validation
func (s *doctorService) Create(doctor *domain.Doctor) (*domain.Doctor, error) {
	// Validasi wajib
	if strings.TrimSpace(doctor.Name) == "" ||
		strings.TrimSpace(doctor.Email) == "" ||
		strings.TrimSpace(doctor.SpecializationID) == "" ||
		doctor.Age <= 0 ||
		doctor.ConsultationFee <= 0 ||
		doctor.YearsOfExperience < 0 ||
		strings.TrimSpace(doctor.LicenseNumber) == "" {
		return nil, ErrMissingFields
	}

	existing, err := s.repo.GetByEmail(doctor.Email)
	if err != nil {
		log.Printf("ERROR checking existing email: %v", err)
		return nil, ErrInternal
	}
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}
	// Buatkan command reply rabbitMQ untuk membuat data user terlebih dahulu dengan menyertakan domain.Doctor dan mengembalikan userId
	if s.ch == nil || s.ch.IsClosed() {
		return nil, fmt.Errorf("rabbitmq channel is not open")
	}
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
		ID:    doctor.ID,
		Name:  doctor.Name,
		Email: doctor.Email,
		Role:  "doctor",
	}
	body, _ := json.Marshal(userBody)
	err = s.ch.Publish(
		"",              // default exchange
		"DoctorCreated", // queue name
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

	timeout := time.After(15 * time.Second)
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

			doctor.UserID = &resp.ID
			doctorNew, err := s.repo.Create(doctor)
			if err != nil {
				return nil, fmt.Errorf("failed to create doctor: %v", err)
			}

			doctorContext, _ := json.Marshal(doctorNew)
			payload := rabbitmqown.NotificationPayload{
				Channel:       "email",
				Recipient:     "baratagusti.bg@gmail.com", // Assuming recipient is the patient ID
				TemplateName:  "doctorCreated",            // Example template name
				Subject:       "Doctor Created!",
				Body:          fmt.Sprintf("Your doctor profile has been created. Please login into %v to create a consultation with email: %v, password: temansehat", os.Getenv("APP_URL"), doctorNew.Email),
				SourceService: "clinicDataService",
				Context:       doctorContext, // Additional context can be added here if needed
				Status:        "pending",     // Initial status
				ErrorMessage:  "",            // No error message initially
				RetryCount:    0,             // Initial retry count
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
			return doctorNew, nil
		case <-timeout:
			return nil, fmt.Errorf("timeout waiting for auth-service")
		}
	}
}

// GetByID returns a doctor by ID
func (s *doctorService) GetByID(id string) (*domain.Doctor, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrMissingID
	}
	return s.repo.GetByID(id)
}

// GetAll returns all doctors
func (s *doctorService) GetAll() ([]domain.Doctor, error) {
	return s.repo.GetAll()
}

// Update modifies an existing doctor
func (s *doctorService) Update(d *domain.Doctor) (*domain.Doctor, error) {
	if strings.TrimSpace(d.ID) == "" {
		return nil, ErrMissingID
	}
	return s.repo.Update(d)
}

// Delete removes a doctor by ID
func (s *doctorService) Delete(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrMissingID
	}
	return s.repo.Delete(id)
}
