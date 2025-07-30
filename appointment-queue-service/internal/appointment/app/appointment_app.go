package app

import (
	"appointment-queue-service/internal/appointment/domain"
	"context"
)

type AppointmentApp interface {
	FindAll(ctx context.Context) ([]*domain.AppointmentModel, error)
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

func (app *appointmentApp) FindAll(ctx context.Context) ([]*domain.AppointmentModel, error) {
	return app.service.FindAll(ctx)
}

func (app *appointmentApp) FindByIDAppointment(ctx context.Context, id uint) (*domain.AppointmentModel, error) {
	return app.service.FindByID(ctx, id)
}

func (app *appointmentApp) FindByUserID(ctx context.Context, userID uint) ([]*domain.AppointmentModel, error) {
	return app.service.FindByUserID(ctx, userID)
}

func (app *appointmentApp) CreateAppointment(ctx context.Context, appointment *domain.AppointmentModel) (*domain.AppointmentModel, error) {
	return app.service.CreateAppointment(ctx, appointment)
}

func (app *appointmentApp) UpdateAppointment(ctx context.Context, id uint, appointment *domain.AppointmentModel) (*domain.AppointmentModel, error) {
	return app.service.Update(ctx, id, appointment)
}

func (app *appointmentApp) MarkAsPaid(ctx context.Context, appointmentID uint) error {
	return app.service.MarkAsPaid(ctx, appointmentID)
}
