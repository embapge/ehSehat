package app

import (
	"testing"

	"clinic-data-service/internal/clinicdata/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --------- MOCK REPOSITORY ---------
type MockDoctorRepo struct {
	mock.Mock
}

func (m *MockDoctorRepo) Create(d *domain.Doctor) (*domain.Doctor, error) {
	args := m.Called(d)
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorRepo) GetByID(id string) (*domain.Doctor, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorRepo) GetByEmail(email string) (*domain.Doctor, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorRepo) GetAll() ([]domain.Doctor, error) {
	args := m.Called()
	return args.Get(0).([]domain.Doctor), args.Error(1)
}
func (m *MockDoctorRepo) Update(d *domain.Doctor) (*domain.Doctor, error) {
	args := m.Called(d)
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// --------- MOCK PUBLISHER (RabbitMQ/Event) ---------
type MockDoctorPublisher struct {
	mock.Mock
}

// Implementasi interface DoctorEventPublisher
func (m *MockDoctorPublisher) PublishDoctorCreated(doctor *domain.Doctor) (string, error) {
	args := m.Called(doctor)
	return args.String(0), args.Error(1)
}

// --------- UNIT TESTS ---------

func TestCreateDoctor_Success(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	doctor := &domain.Doctor{
		Name:              "dr. Budi",
		Email:             "budi@klinik.com",
		SpecializationID:  "sp-123",
		Age:               40,
		ConsultationFee:   100000,
		YearsOfExperience: 10,
		LicenseNumber:     "1234567",
	}

	// Setup: email tidak ada
	mockRepo.On("GetByEmail", "budi@klinik.com").Return(nil, nil)
	mockPublisher.On("PublishDoctorCreated", doctor).Return("user-123", nil) // <- userID yang didapat dari event
	mockRepo.On("Create", doctor).Return(doctor, nil)

	result, err := service.Create(doctor)
	assert.NoError(t, err)
	assert.Equal(t, doctor, result)
	mockRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestCreateDoctor_MissingField(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	doctor := &domain.Doctor{
		Name:              "", // missing
		Email:             "budi@klinik.com",
		SpecializationID:  "sp-123",
		Age:               40,
		ConsultationFee:   100000,
		YearsOfExperience: 10,
		LicenseNumber:     "1234567",
	}
	result, err := service.Create(doctor)
	assert.ErrorIs(t, err, ErrMissingFields)
	assert.Nil(t, result)
}

func TestCreateDoctor_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	doctor := &domain.Doctor{
		Name:              "dr. Budi",
		Email:             "budi@klinik.com",
		SpecializationID:  "sp-123",
		Age:               40,
		ConsultationFee:   100000,
		YearsOfExperience: 10,
		LicenseNumber:     "1234567",
	}
	mockRepo.On("GetByEmail", "budi@klinik.com").Return(doctor, nil)

	result, err := service.Create(doctor)
	assert.ErrorIs(t, err, ErrEmailAlreadyExists)
	assert.Nil(t, result)
}

func TestDoctorGetByID_Success(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	doctor := &domain.Doctor{ID: "id-1", Name: "Budi"}
	mockRepo.On("GetByID", "id-1").Return(doctor, nil)

	result, err := service.GetByID("id-1")
	assert.NoError(t, err)
	assert.Equal(t, doctor, result)
}

func TestDoctorGetByID_MissingID(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	result, err := service.GetByID("")
	assert.ErrorIs(t, err, ErrMissingID)
	assert.Nil(t, result)
}

func TestUpdateDoctor_Success(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	doctor := &domain.Doctor{ID: "id-1", Name: "Update"}
	mockRepo.On("Update", doctor).Return(doctor, nil)

	result, err := service.Update(doctor)
	assert.NoError(t, err)
	assert.Equal(t, doctor, result)
}

func TestUpdateDoctor_MissingID(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	doctor := &domain.Doctor{ID: "", Name: "NoID"}
	result, err := service.Update(doctor)
	assert.ErrorIs(t, err, ErrMissingID)
	assert.Nil(t, result)
}

func TestDeleteDoctor_Success(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	mockRepo.On("Delete", "321").Return(nil)
	err := service.Delete("321")
	assert.NoError(t, err)
}

func TestDeleteDoctor_MissingID(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	err := service.Delete("")
	assert.ErrorIs(t, err, ErrMissingID)
}

func TestGetAllDoctors_Success(t *testing.T) {
	mockRepo := new(MockDoctorRepo)
	mockPublisher := new(MockDoctorPublisher)
	service := NewDoctorService(mockRepo, mockPublisher)

	list := []domain.Doctor{{Name: "A"}, {Name: "B"}}
	mockRepo.On("GetAll").Return(list, nil)

	result, err := service.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}
