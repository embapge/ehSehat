package domain

import (
	"context"
	"database/sql"
	"time"

	"ehSehat/libs/utils"

	"gorm.io/gorm"
)

type QueueRepository interface {
	FindByID(ctx context.Context, id uint) (*QueueModel, error)
	FindTodayByDoctor(ctx context.Context, doctorID uint) ([]*QueueModel, error)
	Create(ctx context.Context, queue *QueueModel) error
	Update(ctx context.Context, queue *QueueModel) error
	GetNextQueueNumber(ctx context.Context, doctorID uint) (int, error)
}
type queueRepository struct {
	db *gorm.DB
}

func NewQueueRepository(db *gorm.DB) QueueRepository {
	return &queueRepository{db}
}

func (r *queueRepository) FindByID(ctx context.Context, id uint) (*QueueModel, error) {
	var q QueueModel
	if err := r.db.WithContext(ctx).First(&q, id).Error; err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *queueRepository) FindTodayByDoctor(ctx context.Context, doctorID uint) ([]*QueueModel, error) {
	start, end := utils.TodayStartEnd()

	var queues []*QueueModel
	if err := r.db.WithContext(ctx).
		Where("doctor_id = ? AND created_at BETWEEN ? AND ?", doctorID, start, end).
		Order("queue_number ASC").
		Find(&queues).Error; err != nil {
		return nil, err
	}
	return queues, nil
}

func (r *queueRepository) Create(ctx context.Context, q *QueueModel) error {
	now := time.Now()
	q.CreatedAt = now
	if err := r.db.WithContext(ctx).Create(q).Error; err != nil {
		return err
	}
	return nil
}

func (r *queueRepository) Update(ctx context.Context, q *QueueModel) error {
	q.UpdatedAt = time.Now()
	if err := r.db.WithContext(ctx).Save(q).Error; err != nil {
		return err
	}
	return nil
}

func (r *queueRepository) GetNextQueueNumber(ctx context.Context, doctorID uint) (int, error) {
	start, end := utils.TodayStartEnd()

	var max sql.NullInt64
	err := r.db.WithContext(ctx).
		Model(&QueueModel{}).
		Select("MAX(queue_number)").
		Where("doctor_id = ? AND created_at BETWEEN ? AND ?", doctorID, start, end).
		Scan(&max).Error
	if err != nil {
		return 0, err
	}

	if !max.Valid {
		return 1, nil
	}

	return int(max.Int64) + 1, nil
}
