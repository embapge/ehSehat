package app

import (
	"clinic-data-service/internal/clinicdata/domain"
	"log"
	"strings"
)

// SpecializationService interface untuk business logic specialization
type SpecializationService interface {
	Create(s *domain.Specialization) (*domain.Specialization, error)
	GetByID(id string) (*domain.Specialization, error)
	GetAll() ([]domain.Specialization, error)
	Update(s *domain.Specialization) (*domain.Specialization, error)
	Delete(id string) error
}

// Implementasi struct service
type specializationService struct {
	repo domain.SpecializationRepository
}

// Constructor
func NewSpecializationService(r domain.SpecializationRepository) SpecializationService {
	return &specializationService{repo: r}
}

// Create specialization dengan validasi nama wajib
func (s *specializationService) Create(sp *domain.Specialization) (*domain.Specialization, error) {
	log.Printf("DEBUG create specialization: Name='%s'", sp.Name)

	if strings.TrimSpace(sp.Name) == "" {
		return nil, ErrMissingFields
	}

	created, err := s.repo.Create(sp)
	if err != nil {
		log.Printf("ERROR create specialization: %v", err)
		return nil, ErrInternal
	}

	return created, nil
}

// GetByID specialization by ID
func (s *specializationService) GetByID(id string) (*domain.Specialization, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrMissingID
	}

	result, err := s.repo.GetByID(id)
	if err != nil {
		log.Printf("ERROR get specialization by ID: %v", err)
		return nil, ErrInternal
	}
	if result == nil {
		return nil, ErrNotFound
	}
	return result, nil
}

// GetAll specializations
func (s *specializationService) GetAll() ([]domain.Specialization, error) {
	list, err := s.repo.GetAll()
	if err != nil {
		log.Printf("ERROR get all specializations: %v", err)
		return nil, ErrInternal
	}
	return list, nil
}

// Update specialization
func (s *specializationService) Update(sp *domain.Specialization) (*domain.Specialization, error) {
	if strings.TrimSpace(sp.ID) == "" || strings.TrimSpace(sp.Name) == "" {
		return nil, ErrMissingFields
	}

	updated, err := s.repo.Update(sp)
	if err != nil {
		log.Printf("ERROR update specialization: %v", err)
		return nil, ErrInternal
	}
	return updated, nil
}

// Delete specialization by ID
func (s *specializationService) Delete(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrMissingID
	}
	err := s.repo.Delete(id)
	if err != nil {
		log.Printf("ERROR delete specialization: %v", err)
		return ErrInternal
	}
	return nil
}
