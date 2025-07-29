package app

import (
	"appointment-queue-service/internal/appointment/domain"
	"context"
)

type AppointmentApp interface {
	FindByIDAppointment(ctx context.Context, id uint) (*domain.AppointmentModel, error)
	FindByUserID(ctx context.Context, userID uint) ([]*domain.AppointmentModel, error)
	CreateAppointment(ctx context.Context, appointment *domain.AppointmentModel) (*domain.AppointmentModel, error)
	UpdateAppointment(ctx context.Context, id uint, appointment *domain.AppointmentModel) (*domain.AppointmentModel, error)
	MarkAsPaid(ctx context.Context, appointmentID uint) error
}

type appointmentApp struct {
	service domain.AppointmentService
}

func NewAppointmentApp(s domain.AppointmentService) AppointmentApp {
	return &appointmentApp{service: s}
}

func (app *appointmentApp) FindByIDAppointment(ctx context.Context, id uint) (*domain.AppointmentModel, error) {
	return app.service.GetByID(ctx, id)
}

func (app *appointmentApp) FindByUserID(ctx context.Context, userID uint) ([]*domain.AppointmentModel, error) {
	return app.service.GetAllByUser(ctx, userID)
}

func (app *appointmentApp) CreateAppointment(ctx context.Context, appointment *domain.AppointmentModel) (*domain.AppointmentModel, error) {
	return app.service.Create(ctx, appointment)
}

func (app *appointmentApp) UpdateAppointment(ctx context.Context, id uint, appointment *domain.AppointmentModel) (*domain.AppointmentModel, error) {
	return app.service.Update(ctx, id, appointment)
}

func (app *appointmentApp) MarkAsPaid(ctx context.Context, appointmentID uint) error {
	return app.service.MarkAsPaid(ctx, appointmentID)
}
