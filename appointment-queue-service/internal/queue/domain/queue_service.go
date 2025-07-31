package domain

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type QueueService interface {
	GetQueueByID(ctx context.Context, id uint) (*QueueModel, error)
	GetTodayQueuesByDoctor(ctx context.Context, doctorID uint) ([]*QueueModel, error)
	CreateQueue(ctx context.Context, q *QueueModel) error
	UpdateQueue(ctx context.Context, q *QueueModel) error
	GenerateNextQueue(ctx context.Context, doctorID string, userID string, userName, userRole, userEmail string, appointmentID *uint, patientID *string, patientName *string, doctorName, doctorSpecialization, queueType string) (*QueueModel, error)
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
	if q.Type == "" {
		return errors.New("data user, dokter, dan tipe antrian wajib diisi")
	}
	// if q.PatientName == nil && q.PatientID == nil {
	// 	return errors.New("pasien harus diisi")
	// }
	q.Status = "active"
	// if q.StartFrom.IsZero() {
	// 	q.StartFrom = time.Now().Add(5 * time.Minute)
	// }
	// Ambil nomor antrian berikutnya
	queueNumber, startFrom, _ := s.repo.GetNextQueueNumber(ctx, q.DoctorID)

	fmt.Println("Cek nilai startFromStr:", startFrom)

	if queueNumber == 0 {
		queueNumber = 1 // Default to 1 if no previous queue number exists
		// set startFrom to 10 minutes from now
		startFrom = time.Now().Add(10 * time.Minute)
	} else if queueNumber > 0 {
		queueNumber++ // Increment the queue number
		if startFrom.Before(time.Now()) {
			startFrom = time.Now().Add(10 * time.Minute)
		} else {
			startFrom = startFrom.Add(40 * time.Minute)
		}
	}

	q.QueueNumber = int(queueNumber)
	q.StartFrom = startFrom // dummy estimasi

	fmt.Println("Cek nilai startFromStr2:", startFrom, queueNumber)

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
	doctorID string,
	userID string,
	userName, userRole, userEmail string,
	appointmentID *uint,
	patientID *string,
	patientName *string,
	doctorName, doctorSpecialization, queueType string,
) (*QueueModel, error) {
	if doctorID == "" || userID == "" || userName == "" || userRole == "" || userEmail == "" || doctorName == "" || doctorSpecialization == "" || queueType == "" {
		return nil, errors.New("data wajib tidak lengkap")
	}

	if patientID == nil && patientName == nil {
		return nil, errors.New("pasien harus diisi")
	}

	// Ambil nomor antrian berikutnya dari repository
	queueNumber, startFrom, err := s.repo.GetNextQueueNumber(ctx, doctorID)
	if err != nil {
		return &QueueModel{}, err
	}

	fmt.Println("Cek nilai startFromStr:", startFrom)

	if queueNumber == 0 {
		queueNumber = 1 // Default to 1 if no previous queue number exists
		// set startFrom to 10 minutes from now
		startFrom = time.Now().Add(10 * time.Minute)
	} else if queueNumber > 0 {
		queueNumber++ // Increment the queue number
		// tambah estimasi masuk ruangan 40 menit
		startFrom = startFrom.Add(40 * time.Minute)
	}

	q := &QueueModel{
		UserID:               userID,
		UserName:             userName,
		UserRole:             userRole,
		UserEmail:            userEmail,
		PatientID:            patientID,
		PatientName:          patientName,
		DoctorID:             doctorID,
		DoctorName:           doctorName,
		DoctorSpecialization: doctorSpecialization,
		AppointmentID:        appointmentID,
		Type:                 queueType,
		QueueNumber:          int(queueNumber),
		Status:               "active",
		StartFrom:            startFrom, // dummy estimasi
	}

	if err := s.repo.Create(ctx, q); err != nil {
		return nil, err
	}
	return q, nil
}
