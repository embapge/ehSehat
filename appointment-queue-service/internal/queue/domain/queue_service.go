package domain

import (
	"context"
	"errors"
	"time"
)

type QueueService interface {
	GetQueueByID(ctx context.Context, id uint) (*QueueModel, error)
	GetTodayQueuesByDoctor(ctx context.Context, doctorID uint) ([]*QueueModel, error)
	CreateQueue(ctx context.Context, q *QueueModel) error
	UpdateQueue(ctx context.Context, q *QueueModel) error
	GenerateNextQueue(ctx context.Context, doctorID uint, userID uint, userName, userRole string, appointmentID *uint, patientID *uint, patientName *string, doctorName, doctorSpecialization, queueType string) (*QueueModel, error)
}

type queueService struct {
	repo QueueRepository
}

func NewQueueService(r QueueRepository) QueueService {
	return &queueService{repo: r}
}

func (s *queueService) GetQueueByID(ctx context.Context, id uint) (*QueueModel, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *queueService) GetTodayQueuesByDoctor(ctx context.Context, doctorID uint) ([]*QueueModel, error) {
	return s.repo.FindTodayByDoctor(ctx, doctorID)
}

func (s *queueService) CreateQueue(ctx context.Context, q *QueueModel) error {
	if q.DoctorID == 0 || q.UserID == 0 || q.UserName == "" || q.UserRole == "" || q.DoctorName == "" || q.DoctorSpecialization == "" || q.Type == "" {
		return errors.New("data user, dokter, dan tipe antrian wajib diisi")
	}
	if q.PatientName == nil && q.PatientID == nil {
		return errors.New("pasien harus diisi")
	}
	q.Status = "active"
	q.CreatedAt = time.Now()

	// Ambil nomor antrian berikutnya
	nextNumber, err := s.repo.GetNextQueueNumber(ctx, q.DoctorID)
	if err != nil {
		return err
	}
	q.QueueNumber = nextNumber

	return s.repo.Create(ctx, q)
}

func (s *queueService) UpdateQueue(ctx context.Context, q *QueueModel) error {
	if q.ID == 0 {
		return errors.New("id queue tidak boleh kosong")
	}
	return s.repo.Update(ctx, q)
}

func (s *queueService) GenerateNextQueue(
	ctx context.Context,
	doctorID uint,
	userID uint,
	userName, userRole string,
	appointmentID *uint,
	patientID *uint,
	patientName *string,
	doctorName, doctorSpecialization, queueType string,
) (*QueueModel, error) {
	if doctorID == 0 || userID == 0 || userName == "" || userRole == "" || doctorName == "" || doctorSpecialization == "" || queueType == "" {
		return nil, errors.New("data wajib tidak lengkap")
	}

	if patientID == nil && patientName == nil {
		return nil, errors.New("pasien harus diisi")
	}

	nextNumber, err := s.repo.GetNextQueueNumber(ctx, doctorID)
	if err != nil {
		return nil, err
	}

	q := &QueueModel{
		UserID:               userID,
		UserName:             userName,
		UserRole:             userRole,
		PatientID:            patientID,
		PatientName:          patientName,
		DoctorID:             doctorID,
		DoctorName:           doctorName,
		DoctorSpecialization: doctorSpecialization,
		AppointmentID:        appointmentID,
		Type:                 queueType,
		QueueNumber:          nextNumber,
		Status:               "active",
		StartFrom:            time.Now().Add(time.Minute * 5), // dummy estimasi
		CreatedAt:            time.Now(),
	}

	if err := s.repo.Create(ctx, q); err != nil {
		return nil, err
	}
	return q, nil
}
