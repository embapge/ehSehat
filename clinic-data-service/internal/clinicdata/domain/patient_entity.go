package domain

import "time"

// Patient represents the patient entity in the domain layer
type Patient struct {
<<<<<<< HEAD
	ID          string
	UserID      *string
	Name        string
	Email       string
	BirthDate   string
	Gender      string
	PhoneNumber *string
	Address     *string

=======
	ID           string
	UserID       *string
	Name         string
	Email        string
	BirthDate    string
	Gender       string
	PhoneNumber  *string
	Address      *string
>>>>>>> main
	CreatedBy    *string
	CreatedName  string
	CreatedEmail string
	CreatedRole  string
	CreatedAt    time.Time
<<<<<<< HEAD

=======
>>>>>>> main
	UpdatedBy    *string
	UpdatedName  string
	UpdatedEmail string
	UpdatedRole  string
	UpdatedAt    time.Time
}

// UserIDOrEmpty returns user_id or empty string if nil
func (p Patient) UserIDOrEmpty() string {
	if p.UserID != nil {
		return *p.UserID
	}
	return ""
}

// PhoneNumberOrEmpty returns phone_number or empty string if nil
func (p Patient) PhoneNumberOrEmpty() string {
	if p.PhoneNumber != nil {
		return *p.PhoneNumber
	}
	return ""
}

// AddressOrEmpty returns address or empty string if nil
func (p Patient) AddressOrEmpty() string {
	if p.Address != nil {
		return *p.Address
	}
	return ""
}

// PatientRepository defines repository behavior for patient entity
type PatientRepository interface {
	Create(patient *Patient) (*Patient, error)
	GetByID(id string) (*Patient, error)
	GetByEmail(email string) (*Patient, error)
	GetAll() ([]Patient, error)
	Update(patient *Patient) (*Patient, error)
	Delete(id string) error
}
