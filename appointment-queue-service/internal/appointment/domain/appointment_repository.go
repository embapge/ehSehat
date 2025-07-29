package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	FindByID(ctx context.Context, id uint) (*AppointmentModel, error)
	FindByUserID(ctx context.Context, userID uint) ([]*AppointmentModel, error)
	Create(ctx context.Context, appointment *AppointmentModel) error
	Update(ctx context.Context, appointment *AppointmentModel) error
	MarkAsPaid(ctx context.Context, appointmentID uint) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) FindByID(ctx context.Context, id uint) (*AppointmentModel, error) {
	var appointment AppointmentModel
	if err := r.db.WithContext(ctx).First(&appointment, id).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *appointmentRepository) FindByUserID(ctx context.Context, userID uint) ([]*AppointmentModel, error) {
	var appointments []*AppointmentModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *appointmentRepository) Create(ctx context.Context, a *AppointmentModel) error {
	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now
	if err := r.db.WithContext(ctx).Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (r *appointmentRepository) Update(ctx context.Context, a *AppointmentModel) error {
	a.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Save(a).Error; err != nil {
		return err
	}
	return nil
}

func (r *appointmentRepository) MarkAsPaid(ctx context.Context, appointmentID uint) error {
	return r.db.WithContext(ctx).
		Model(&AppointmentModel{}).
		Where("id = ?", appointmentID).
		Updates(map[string]interface{}{
			"is_paid":    true,
			"status":     "paid",
			"updated_at": time.Now(),
		}).Error
}
