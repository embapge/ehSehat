package domain

import "time"

// Doctor represents the doctor entity
type Doctor struct {
	ID                string
	UserID            *string
	Name              string
	Email             string
	SpecializationID  string
	Age               int
	ConsultationFee   float64
	YearsOfExperience int
	LicenseNumber     string
	PhoneNumber       *string
	IsActive          bool

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

// UserIDOrEmpty returns user_id or empty string if nil
func (d Doctor) UserIDOrEmpty() string {
	if d.UserID != nil {
		return *d.UserID
	}
	return ""
}

// PhoneNumberOrEmpty returns phone_number or empty string if nil
func (d Doctor) PhoneNumberOrEmpty() string {
	if d.PhoneNumber != nil {
		return *d.PhoneNumber
	}
	return ""
}

// DoctorRepository defines repository behavior for doctor entity
type DoctorRepository interface {
	Create(doctor *Doctor) (*Doctor, error)
	GetByID(id string) (*Doctor, error)
	GetByEmail(email string) (*Doctor, error)
	GetAll() ([]Doctor, error)
	Update(doctor *Doctor) (*Doctor, error)
	Delete(id string) error
}
