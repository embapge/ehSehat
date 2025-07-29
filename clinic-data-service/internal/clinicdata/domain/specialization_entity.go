package domain

import "time"

// Specialization represents the specialization entity
type Specialization struct {
	ID   string
	Name string

	CreatedBy    *string
	CreatedName  string
	CreatedEmail string
	CreatedRole  string
	CreatedAt    time.Time

	UpdatedBy    *string
	UpdatedName  string
	UpdatedEmail string
	UpdatedRole  string
	UpdatedAt    time.Time
}

// SpecializationRepository defines the contract for specialization DB operations
type SpecializationRepository interface {
	Create(s *Specialization) (*Specialization, error)
	GetByID(id string) (*Specialization, error)
	GetAll() ([]Specialization, error)
	Update(s *Specialization) (*Specialization, error)
	Delete(id string) error
}
