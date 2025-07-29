package app

import (
	"context"
	"ehSehat/consultation-service/internal/consultation/domain"
	"ehSehat/libs/utils"
	"fmt"
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
	fmt.Println("id consultation:", id)
	return app.repo.FindByID(ctx, id)
}

func (app *consultationApp) CreateConsultation(ctx context.Context, consultation *domain.Consultation) error {
	if consultation.Patient.ID == "" || consultation.Room.ID == "" || consultation.Doctor.ID == "" {
		return utils.NewBadRequestError("All fields are required")
	}

	consultation.Status = "unpaid"
	return app.repo.Create(ctx, consultation)
}

func (app *consultationApp) UpdateConsultation(ctx context.Context, consultation *domain.Consultation) error {
	return app.repo.Update(ctx, consultation)
}
