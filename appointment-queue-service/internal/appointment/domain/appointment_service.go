package domain

import (
	"context"
	"errors"
)

type AppointmentService interface {
	FindAll(ctx context.Context) ([]*AppointmentModel, error)
	FindByUserID(ctx context.Context, userID uint) ([]*AppointmentModel, error)
	FindByID(ctx context.Context, id uint) (*AppointmentModel, error)
	CreateAppointment(ctx context.Context, appointment *AppointmentModel) (*AppointmentModel, error)
	Update(ctx context.Context, id uint, appointment *AppointmentModel) (*AppointmentModel, error)
	MarkAsPaid(ctx context.Context, id uint) error
}

type appointmentService struct {
	repo AppointmentRepository
}

func NewAppointmentService(r AppointmentRepository) AppointmentService {
	return &appointmentService{repo: r}
}

func (s *appointmentService) FindAll(ctx context.Context) ([]*AppointmentModel, error) {
	return s.repo.FindAll(ctx)
}

func (s *appointmentService) FindByUserID(ctx context.Context, userID uint) ([]*AppointmentModel, error) {
	if userID == 0 {
		return nil, errors.New("user_id tidak valid")
	}
	return s.repo.FindByUserID(ctx, userID)
}

func (s *appointmentService) FindByID(ctx context.Context, id uint) (*AppointmentModel, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.FindByID(ctx, id)
}

func (s *appointmentService) CreateAppointment(ctx context.Context, a *AppointmentModel) (*AppointmentModel, error) {
	if a.UserID == 0 || a.DoctorID == 0 || a.AppointmentAt.IsZero() {
		return nil, errors.New("user_id, doctor_id, dan appointment_at wajib diisi")
	}
	if a.UserFullName == "" || a.DoctorName == "" || a.DoctorSpecialization == "" {
		return nil, errors.New("nama user, nama dokter, dan spesialisasi dokter wajib diisi")
	}

	a.Status = "unpaid"

	err := s.repo.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *appointmentService) Update(ctx context.Context, id uint, a *AppointmentModel) (*AppointmentModel, error) {
	if id == 0 {
		return nil, errors.New("id appointment tidak valid")
	}
	a.ID = id
	err := s.repo.Update(ctx, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *appointmentService) MarkAsPaid(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("id appointment tidak valid")
	}
	return s.repo.MarkAsPaid(ctx, id)
}
