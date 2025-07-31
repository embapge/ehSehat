package domain

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"ehSehat/libs/utils"

	"gorm.io/gorm"
)

type QueueRepository interface {
	FindByID(ctx context.Context, id uint) (*QueueModel, error)
	FindTodayByDoctor(ctx context.Context, doctorID uint) ([]*QueueModel, error)
	Create(ctx context.Context, queue *QueueModel) error
	Update(ctx context.Context, queue *QueueModel) error
	GetNextQueueNumber(ctx context.Context, doctorID string) (int64, time.Time, error)
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
	if err := r.db.WithContext(ctx).Create(q).Error; err != nil {
		return err
	}
	return nil
}

func (r *queueRepository) Update(ctx context.Context, q *QueueModel) error {
	if err := r.db.WithContext(ctx).Save(q).Error; err != nil {
		return err
	}
	return nil
}

func (r *queueRepository) GetNextQueueNumber(ctx context.Context, doctorID string) (int64, time.Time, error) {
	start, end := utils.TodayStartEnd()

	type result struct {
		MaxQueueNumber sql.NullInt64
		StartFrom      sql.NullTime
	}
	var res result
	err := r.db.WithContext(ctx).
		Model(&QueueModel{}).
		Select("queue_number as max_queue_number, start_from").
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("queue_number DESC").
		Limit(1).
		Row().
		Scan(&res.MaxQueueNumber, &res.StartFrom)

	// jika error dikarenakan tidak ada data, set nilai default
	if err == sql.ErrNoRows {
		return 0, time.Time{}, nil
	}

	fmt.Println("Lewat sini")

	return res.MaxQueueNumber.Int64, res.StartFrom.Time, nil
}
