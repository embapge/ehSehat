package app

import (
	"clinic-data-service/internal/clinicdata/domain"
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
	repo domain.DoctorRepository
}

// NewDoctorService is the constructor
func NewDoctorService(r domain.DoctorRepository) DoctorService {
	return &doctorService{repo: r}
}

// Create inserts a new doctor after validation
func (s *doctorService) Create(d *domain.Doctor) (*domain.Doctor, error) {

	// Validasi wajib
	if strings.TrimSpace(d.Name) == "" ||
		strings.TrimSpace(d.Email) == "" ||
		strings.TrimSpace(d.SpecializationID) == "" ||
		d.Age <= 0 ||
		d.ConsultationFee <= 0 ||
		d.YearsOfExperience < 0 ||
		strings.TrimSpace(d.LicenseNumber) == "" {
		return nil, ErrMissingFields
	}

	// Validasi email unik
	existing, err := s.repo.GetByEmail(d.Email)
	if err != nil {
		log.Printf("ERROR checking existing email: %v", err)
		return nil, ErrInternal
	}
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	return s.repo.Create(d)
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
