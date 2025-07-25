package app

import (
	"appointment-queue-service/internal/queue/domain"
	"context"
)

type QueueApp interface {
	FindByIDQueue(ctx context.Context, id uint) (*domain.QueueModel, error)
	FindTodayByDoctor(ctx context.Context, doctorID uint) ([]*domain.QueueModel, error)
	CreateQueue(ctx context.Context, queue *domain.QueueModel) error
	UpdateQueue(ctx context.Context, queue *domain.QueueModel) error
	GenerateNextQueue(
		ctx context.Context,
		doctorID uint,
		userID uint,
		userName, userRole string,
		appointmentID *uint,
		patientID *uint,
		patientName *string,
		doctorName, doctorSpecialization, queueType string,
	) (*domain.QueueModel, error)
}

type queueApp struct {
	service domain.QueueService
}

func NewQueueApp(service domain.QueueService) QueueApp {
	return &queueApp{service: service}
}

func (app *queueApp) FindByIDQueue(ctx context.Context, id uint) (*domain.QueueModel, error) {
	return app.service.GetQueueByID(ctx, id)
}

func (app *queueApp) FindTodayByDoctor(ctx context.Context, doctorID uint) ([]*domain.QueueModel, error) {
	return app.service.GetTodayQueuesByDoctor(ctx, doctorID)
}

func (app *queueApp) CreateQueue(ctx context.Context, queue *domain.QueueModel) error {
	return app.service.CreateQueue(ctx, queue)
}

func (app *queueApp) UpdateQueue(ctx context.Context, queue *domain.QueueModel) error {
	return app.service.UpdateQueue(ctx, queue)
}

func (app *queueApp) GenerateNextQueue(
	ctx context.Context,
	doctorID uint,
	userID uint,
	userName, userRole string,
	appointmentID *uint,
	patientID *uint,
	patientName *string,
	doctorName, doctorSpecialization, queueType string,
) (*domain.QueueModel, error) {
	return app.service.GenerateNextQueue(
		ctx,
		doctorID,
		userID,
		userName,
		userRole,
		appointmentID,
		patientID,
		patientName,
		doctorName,
		doctorSpecialization,
		queueType,
	)
}
