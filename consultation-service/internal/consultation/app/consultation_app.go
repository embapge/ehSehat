package app

import (
	"context"
	"ehSehat/consultation-service/internal/consultation/domain"
	"time"
)

type consultationApp struct {
	repo domain.ConsultationRepository
}

func NewConsultationApp(repo domain.ConsultationRepository) *consultationApp {
	return &consultationApp{repo: repo}
}

// FindByIDConsultation(ctx context.Context, id string) (*Consultation, error)
// 	CreateConsultation(ctx context.Context, consultation *Consultation) error
// 	UpdateConsultation(ctx context.Context, consultation *Consultation) error

func (app *consultationApp) FindByIDConsultation(ctx context.Context, id string) (*domain.Consultation, error) {
	return app.repo.FindByID(ctx, id)
}

func (app *consultationApp) CreateConsultation(ctx context.Context, consultation *domain.Consultation) error {
	consultation.Status = "unpaid"
	now := time.Now()
	consultation.CreatedAt = now
	consultation.UpdatedAt = now
	return app.repo.Create(ctx, consultation)
}

func (app *consultationApp) UpdateConsultation(ctx context.Context, consultation *domain.Consultation) error {
	return app.repo.Update(ctx, consultation)
}
