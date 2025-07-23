package app

import (
	"context"
	"ehSehat/consultation-service/internal/consultation/domain"
	"ehSehat/libs/utils"
)

type consultationApp struct {
	repo domain.ConsultationRepository
}

func NewConsultationApp(repo domain.ConsultationRepository) *consultationApp {
	return &consultationApp{repo: repo}
}

func (app *consultationApp) FindByIDConsultation(ctx context.Context, id string) (*domain.Consultation, error) {
	return app.repo.FindByID(ctx, id)
}

func (app *consultationApp) CreateConsultation(ctx context.Context, consultation *domain.Consultation) error {
	if consultation.UserID == "" || consultation.PatientID == "" || consultation.DoctorID == "" || consultation.RoomID == "" || consultation.PatientName == "" || consultation.DoctorName == "" || consultation.RoomName == "" {
		return utils.NewBadRequestError("All fields are required")
	}

	return app.repo.Create(ctx, consultation)
}

func (app *consultationApp) UpdateConsultation(ctx context.Context, consultation *domain.Consultation) error {
	return app.repo.Update(ctx, consultation)
}
