package domain

import "context"

type ConsultationService interface {
	FindByIDConsultation(ctx context.Context, id string) (*Consultation, error)
	CreateConsultation(ctx context.Context, consultation *Consultation) error
	UpdateConsultation(ctx context.Context, consultation *Consultation) error
}
