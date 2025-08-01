package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"log"
	"strings"
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
	repo      domain.PatientRepository
	publisher domain.PatientEventPublisher // pakai interface baru di domain
}

// Constructor untuk service
func NewPatientService(r domain.PatientRepository, publisher domain.PatientEventPublisher) PatientService {
	return &patientService{repo: r, publisher: publisher}
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

	// Kirim event ke AuthService (request-reply), dapatkan userID balikan
	userID, err := s.publisher.PublishPatientCreated(p)
	if err != nil {
		return nil, err
	}
	p.UserID = &userID

	// Lanjutkan insert ke DB setelah dapat userID
	patientNew, err := s.repo.Create(p)
	if err != nil {
		log.Printf("ERROR creating patient: %v", err)
		return nil, err
	}

	// -- Jika butuh publish event notifikasi lain, tambahkan logic di sini (optional) --

	return patientNew, nil
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
