package domain

import "context"

type ConsultationRepository interface {
	FindByID(ctx context.Context, id string) (*Consultation, error)
	Create(ctx context.Context, consultation *Consultation) error
	Update(ctx context.Context, consultation *Consultation) error
	// FindByUserID(userID string) ([]*Consultation, error)
	// FindByPatientID(patientID string) ([]*Consultation, error)
	// FindByDoctorID(doctorID string) ([]*Consultation, error)
	// FindByRoomID(roomID string) ([]*Consultation, error)
}
