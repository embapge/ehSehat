package unitTest

import (
	"appointment-queue-service/internal/appointment/domain"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mock Repository ---
type MockAppointmentRepo struct {
	mock.Mock
}

func (m *MockAppointmentRepo) FindAll(ctx context.Context) ([]*domain.AppointmentModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.AppointmentModel), args.Error(1)
}

func (m *MockAppointmentRepo) FindByUserID(ctx context.Context, userID uint) ([]*domain.AppointmentModel, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]*domain.AppointmentModel), args.Error(1)
}

func (m *MockAppointmentRepo) FindByID(ctx context.Context, id uint) (*domain.AppointmentModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.AppointmentModel), args.Error(1)
}

func (m *MockAppointmentRepo) Create(ctx context.Context, a *domain.AppointmentModel) error {
	args := m.Called(ctx, a)
	return args.Error(0)
}

func (m *MockAppointmentRepo) Update(ctx context.Context, a *domain.AppointmentModel) error {
	args := m.Called(ctx, a)
	return args.Error(0)
}

func (m *MockAppointmentRepo) MarkAsPaid(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// --- Unit Tests ---

func TestFindAll(t *testing.T) {
	mockRepo := new(MockAppointmentRepo)
	service := domain.NewAppointmentService(mockRepo)

	expected := []*domain.AppointmentModel{{ID: 1}}

	mockRepo.On("FindAll", mock.Anything).Return(expected, nil)

	result, err := service.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestFindByUserID_Valid(t *testing.T) {
	mockRepo := new(MockAppointmentRepo)
	service := domain.NewAppointmentService(mockRepo)

	mockRepo.On("FindByUserID", mock.Anything, uint(10)).Return([]*domain.AppointmentModel{{ID: 1, UserID: 10}}, nil)

	result, err := service.FindByUserID(context.Background(), 10)

	assert.NoError(t, err)
	assert.Equal(t, uint(10), result[0].UserID)
}

func TestFindByUserID_Invalid(t *testing.T) {
	service := domain.NewAppointmentService(new(MockAppointmentRepo))

	_, err := service.FindByUserID(context.Background(), 0)

	assert.Error(t, err)
	assert.Equal(t, "user_id tidak valid", err.Error())
}

func TestFindByID_Valid(t *testing.T) {
	mockRepo := new(MockAppointmentRepo)
	service := domain.NewAppointmentService(mockRepo)

	mockRepo.On("FindByID", mock.Anything, uint(1)).Return(&domain.AppointmentModel{ID: 1}, nil)

	result, err := service.FindByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
}

func TestFindByID_Invalid(t *testing.T) {
	service := domain.NewAppointmentService(new(MockAppointmentRepo))

	_, err := service.FindByID(context.Background(), 0)

	assert.Error(t, err)
	assert.Equal(t, "id tidak valid", err.Error())
}

func TestCreateAppointment_Success(t *testing.T) {
	mockRepo := new(MockAppointmentRepo)
	service := domain.NewAppointmentService(mockRepo)

	appointment := &domain.AppointmentModel{
		UserID:               1,
		UserFullName:         "John",
		DoctorID:             2,
		DoctorName:           "Dr. Smith",
		DoctorSpecialization: "Cardiology",
		AppointmentAt:        time.Now().Add(time.Hour),
	}

	mockRepo.On("Create", mock.Anything, appointment).Return(nil)

	result, err := service.CreateAppointment(context.Background(), appointment)

	assert.NoError(t, err)
	assert.Equal(t, "unpaid", result.Status)
	mockRepo.AssertExpectations(t)
}

func TestCreateAppointment_Invalid(t *testing.T) {
	service := domain.NewAppointmentService(new(MockAppointmentRepo))

	appointment := &domain.AppointmentModel{}

	_, err := service.CreateAppointment(context.Background(), appointment)

	assert.Error(t, err)
}

func TestUpdate_Success(t *testing.T) {
	mockRepo := new(MockAppointmentRepo)
	service := domain.NewAppointmentService(mockRepo)

	appointment := &domain.AppointmentModel{
		UserID:               1,
		UserFullName:         "John",
		DoctorID:             2,
		DoctorName:           "Dr. Smith",
		DoctorSpecialization: "Cardiology",
		AppointmentAt:        time.Now(),
	}

	mockRepo.On("Update", mock.Anything, mock.Anything).Return(nil)

	result, err := service.Update(context.Background(), 1, appointment)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
}

func TestUpdate_InvalidID(t *testing.T) {
	service := domain.NewAppointmentService(new(MockAppointmentRepo))

	_, err := service.Update(context.Background(), 0, &domain.AppointmentModel{})
	assert.Error(t, err)
	assert.Equal(t, "id appointment tidak valid", err.Error())
}

func TestMarkAsPaid_Success(t *testing.T) {
	mockRepo := new(MockAppointmentRepo)
	service := domain.NewAppointmentService(mockRepo)

	mockRepo.On("MarkAsPaid", mock.Anything, uint(1)).Return(nil)

	err := service.MarkAsPaid(context.Background(), 1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMarkAsPaid_InvalidID(t *testing.T) {
	service := domain.NewAppointmentService(new(MockAppointmentRepo))

	err := service.MarkAsPaid(context.Background(), 0)

	assert.Error(t, err)
	assert.Equal(t, "id appointment tidak valid", err.Error())
}
