package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"fmt"
	"log"
	"strings"
)

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
	repo      domain.DoctorRepository
	publisher domain.DoctorEventPublisher
}

// NewDoctorService is the constructor
func NewDoctorService(r domain.DoctorRepository, publisher domain.DoctorEventPublisher) DoctorService {
	return &doctorService{repo: r, publisher: publisher}
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

	// Buatkan command event ke AuthService (via event publisher)
	// Sekarang RabbitMQ logic DIABSTRAKSI agar mudah di-mock/unit-test
	userID, err := s.publisher.PublishDoctorCreated(doctor)
	if err != nil {
		return nil, fmt.Errorf("failed to publish doctor created event: %v", err)
	}
	doctor.UserID = &userID

	doctorNew, err := s.repo.Create(doctor)
	if err != nil {
		return nil, fmt.Errorf("failed to create doctor: %v", err)
	}

	// OPTIONAL: Kirim notifikasi/email lewat event lain kalau dibutuhkan,
	// bisa juga ditambahkan ke publisher interface lain

	return doctorNew, nil
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
