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

func (m *MockQueueRepo) GetNextQueueNumber(ctx context.Context, doctorID string) (int64, time.Time, error) {
	args := m.Called(ctx, doctorID)
	return args.Get(0).(int64), args.Get(1).(time.Time), args.Error(2)
}

func TestGenerateNextQueue_Success(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	doctorID := "1"
	userID := "2"
	userName := "John"
	userRole := "user"
	userEmail := "john@example.com"
	doctorName := "Dr. Smith"
	doctorSpecialization := "Umum"
	queueType := "appointment"
	appointmentID := uint(99)
	patientID := "3"
	patientName := "Anak John"

	expectedTime := time.Now()

	mockRepo.On("GetNextQueueNumber", mock.Anything, doctorID).Return(int64(5), expectedTime, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.QueueModel")).Return(nil)

	q, err := service.GenerateNextQueue(
		context.Background(),
		doctorID,
		userID,
		userName,
		userRole,
		userEmail,
		&appointmentID,
		&patientID,
		&patientName,
		doctorName,
		doctorSpecialization,
		queueType,
	)

	assert.NoError(t, err)
	assert.Equal(t, 6, q.QueueNumber)
	assert.Equal(t, "active", q.Status)
	mockRepo.AssertExpectations(t)
}

func TestGenerateNextQueue_MissingRequiredData(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	_, err := service.GenerateNextQueue(
		context.Background(),
		"",  // doctorID string
		"",  // userID string
		"",  // userName string
		"",  // userRole string
		"",  // userEmail string
		nil, // appointmentID *uint
		nil, // patientID *string
		nil, // patientName *string
		"",  // doctorName string
		"",  // doctorSpecialization string
		"",  // queueType string
	)

	assert.Error(t, err)
}

func TestGenerateNextQueue_QueueNumberError(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	doctorID := "1"
	patientID := "10"
	patientName := "Anak"

	mockRepo.On("GetNextQueueNumber", mock.Anything, doctorID).
		Return(int64(0), time.Time{}, errors.New("db error"))

	_, err := service.GenerateNextQueue(
		context.Background(),
		doctorID,         // doctorID
		"1",              // userID
		"User",           // userName
		"user",           // userRole
		"user@email.com", // userEmail
		nil,
		&patientID,
		&patientName,
		"Dr", "Spesialis", "appointment",
	)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
}

func TestCreateQueue_Success(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	now := time.Now()
	patientName := "Budi"
	queue := &domain.QueueModel{
		UserID:               "1",
		UserName:             "John",
		UserRole:             "member",
		DoctorID:             "2",
		DoctorName:           "Dr. Smith",
		DoctorSpecialization: "Umum",
		Type:                 "online",
		PatientName:          &patientName,
		StartFrom:            now,
	}

	mockRepo.On("GetNextQueueNumber", mock.Anything, queue.DoctorID).Return(int64(10), now, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.QueueModel")).Return(nil)

	err := service.CreateQueue(context.Background(), queue)
	assert.NoError(t, err)
	assert.Equal(t, 11, queue.QueueNumber)
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

func TestCreateQueue_MissingPatient_NotRequired(t *testing.T) {
	mockRepo := new(MockQueueRepo)
	service := domain.NewQueueService(mockRepo)

	queue := &domain.QueueModel{
		UserID:               "1",
		UserName:             "John",
		UserRole:             "member",
		DoctorID:             "2",
		DoctorName:           "Dr. Smith",
		DoctorSpecialization: "Umum",
		Type:                 "online",
		// PatientName dan PatientID nil (sekarang valid)
	}

	mockRepo.On("GetNextQueueNumber", mock.Anything, queue.DoctorID).Return(int64(3), time.Now(), nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.QueueModel")).Return(nil)

	err := service.CreateQueue(context.Background(), queue)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
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
		{ID: 1, DoctorID: "2"},
		{ID: 2, DoctorID: "2"},
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
