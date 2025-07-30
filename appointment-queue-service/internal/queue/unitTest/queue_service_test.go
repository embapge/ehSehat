package unitTest

import (
	"appointment-queue-service/internal/queue/domain"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQueueRepo implements domain.QueueRepository
type MockQueueRepo struct {
	mock.Mock
}

func (m *MockQueueRepo) FindByID(ctx context.Context, id uint) (*domain.QueueModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.QueueModel), args.Error(1)
}

func (m *MockQueueRepo) FindTodayByDoctor(ctx context.Context, doctorID uint) ([]*domain.QueueModel, error) {
	args := m.Called(ctx, doctorID)
	return args.Get(0).([]*domain.QueueModel), args.Error(1)
}

func (m *MockQueueRepo) Create(ctx context.Context, q *domain.QueueModel) error {
	args := m.Called(ctx, q)
	return args.Error(0)
}

func (m *MockQueueRepo) Update(ctx context.Context, q *domain.QueueModel) error {
	args := m.Called(ctx, q)
	return args.Error(0)
}

func (m *MockQueueRepo) GetNextQueueNumber(ctx context.Context, doctorID uint) (int, error) {
	args := m.Called(ctx, doctorID)
	return args.Int(0), args.Error(1)
}

func TestGenerateNextQueue_Success(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	doctorID := uint(1)
	userID := uint(2)
	userName := "John"
	userRole := "user"
	doctorName := "Dr. Smith"
	doctorSpecialization := "Umum"
	queueType := "appointment"
	appointmentID := uint(99)
	patientID := uint(3)
	patientName := "Anak John"

	mockRepo.On("GetNextQueueNumber", mock.Anything, doctorID).Return(5, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.QueueModel")).Return(nil)

	q, err := service.GenerateNextQueue(
		context.Background(),
		doctorID,
		userID,
		userName,
		userRole,
		&appointmentID,
		&patientID,
		&patientName,
		doctorName,
		doctorSpecialization,
		queueType,
	)

	assert.NoError(t, err)
	assert.Equal(t, 5, q.QueueNumber)
	assert.Equal(t, "active", q.Status)
	mockRepo.AssertExpectations(t)
}

func TestGenerateNextQueue_MissingRequiredData(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	_, err := service.GenerateNextQueue(
		context.Background(),
		0, 0, "", "", nil, nil, nil, "", "", "",
	)

	assert.Error(t, err)
}

func TestGenerateNextQueue_QueueNumberError(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	mockRepo.On("GetNextQueueNumber", mock.Anything, uint(1)).Return(0, errors.New("db error"))

	_, err := service.GenerateNextQueue(
		context.Background(),
		1, 1, "User", "user", nil, nil, nil, "Dr", "Spesialis", "appointment",
	)

	assert.Error(t, err)
}

func TestCreateQueue_Success(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	now := time.Now()
	patientName := "Budi"
	queue := &domain.QueueModel{
		UserID:               1,
		UserName:             "John",
		UserRole:             "member",
		DoctorID:             2,
		DoctorName:           "Dr. Smith",
		DoctorSpecialization: "Umum",
		Type:                 "online",
		PatientName:          &patientName,
		StartFrom:            now,
	}

	mockRepo.On("GetNextQueueNumber", mock.Anything, queue.DoctorID).Return(10, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.QueueModel")).Return(nil)

	err := service.CreateQueue(context.Background(), queue)
	assert.NoError(t, err)
	assert.Equal(t, 10, queue.QueueNumber)
	assert.Equal(t, "active", queue.Status)
	mockRepo.AssertExpectations(t)
}

func TestCreateQueue_MissingData(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	queue := &domain.QueueModel{} // kosong

	err := service.CreateQueue(context.Background(), queue)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "data user, dokter, dan tipe antrian wajib diisi")
}

func TestCreateQueue_MissingPatient(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	queue := &domain.QueueModel{
		UserID:               1,
		UserName:             "John",
		UserRole:             "member",
		DoctorID:             2,
		DoctorName:           "Dr. Smith",
		DoctorSpecialization: "Umum",
		Type:                 "online",
		// PatientName dan PatientID nil
	}

	err := service.CreateQueue(context.Background(), queue)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pasien harus diisi")
}

func TestGetQueueByID_Success(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	expected := &domain.QueueModel{ID: 1, QueueNumber: 5}
	mockRepo.On("FindByID", mock.Anything, uint(1)).Return(expected, nil)

	result, err := service.GetQueueByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTodayQueuesByDoctor(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	expected := []*domain.QueueModel{
		{ID: 1, DoctorID: 2},
		{ID: 2, DoctorID: 2},
	}

	mockRepo.On("FindTodayByDoctor", mock.Anything, uint(2)).Return(expected, nil)

	result, err := service.GetTodayQueuesByDoctor(context.Background(), 2)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestUpdateQueue_Success(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	queue := &domain.QueueModel{ID: 1, Status: "done"}

	mockRepo.On("Update", mock.Anything, queue).Return(nil)

	err := service.UpdateQueue(context.Background(), queue)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateQueue_InvalidID(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	queue := &domain.QueueModel{ID: 0, Status: "fail"}

	err := service.UpdateQueue(context.Background(), queue)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "id queue tidak boleh kosong")
}
