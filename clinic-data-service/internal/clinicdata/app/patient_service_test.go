package app

import (
	"testing"

	"clinic-data-service/internal/clinicdata/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --------- MOCK REPOSITORY ---------
type MockPatientRepo struct {
	mock.Mock
}

func (m *MockPatientRepo) Create(p *domain.Patient) (*domain.Patient, error) {
	args := m.Called(p)
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientRepo) GetByID(id string) (*domain.Patient, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientRepo) GetByEmail(email string) (*domain.Patient, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientRepo) GetAll() ([]domain.Patient, error) {
	args := m.Called()
	return args.Get(0).([]domain.Patient), args.Error(1)
}
func (m *MockPatientRepo) Update(p *domain.Patient) (*domain.Patient, error) {
	args := m.Called(p)
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// --------- MOCK PUBLISHER ---------
type MockPatientEventPublisher struct {
	mock.Mock
}

func (m *MockPatientEventPublisher) PublishPatientCreated(p *domain.Patient) (string, error) {
	args := m.Called(p)
	return args.String(0), args.Error(1)
}

// --------- UNIT TESTS ---------
func TestCreatePatient_Success(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{
		Name:      "John",
		Email:     "john@example.com",
		BirthDate: "2000-01-01",
		Gender:    "M",
	}

	// Setup: email tidak ada di DB
	mockRepo.On("GetByEmail", "john@example.com").Return(nil, nil)
	// Setup: publisher akan mengembalikan userID
	mockPublisher.On("PublishPatientCreated", patient).Return("user-id-123", nil).Once()
	// Setup: insert OK
	mockRepo.On("Create", mock.MatchedBy(func(p *domain.Patient) bool {
		return p.UserID != nil && *p.UserID == "user-id-123"
	})).Return(patient, nil)

	result, err := service.Create(patient)
	assert.NoError(t, err)
	assert.Equal(t, patient, result)

	mockRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestCreatePatient_MissingField(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{
		Name:      "",
		Email:     "john@example.com",
		BirthDate: "2000-01-01",
		Gender:    "M",
	}

	result, err := service.Create(patient)
	assert.ErrorIs(t, err, ErrMissingFields)
	assert.Nil(t, result)
}

func TestCreatePatient_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{
		Name:      "Jane",
		Email:     "jane@example.com",
		BirthDate: "2000-01-01",
		Gender:    "F",
	}

	mockRepo.On("GetByEmail", "jane@example.com").Return(patient, nil)

	result, err := service.Create(patient)
	assert.ErrorIs(t, err, ErrEmailAlreadyExists)
	assert.Nil(t, result)
}

func TestCreatePatient_PublisherError(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{
		Name:      "John",
		Email:     "john@example.com",
		BirthDate: "2000-01-01",
		Gender:    "M",
	}

	mockRepo.On("GetByEmail", "john@example.com").Return(nil, nil)
	mockPublisher.On("PublishPatientCreated", patient).Return("", assert.AnError)

	result, err := service.Create(patient)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestPatientGetByID_Success(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{ID: "123", Name: "Budi"}
	mockRepo.On("GetByID", "123").Return(patient, nil)

	result, err := service.GetByID("123")
	assert.NoError(t, err)
	assert.Equal(t, patient, result)
}

func TestPatientGetByID_MissingID(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	result, err := service.GetByID("")
	assert.ErrorIs(t, err, ErrMissingID)
	assert.Nil(t, result)
}

func TestUpdatePatient_Success(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{ID: "123", Name: "Update"}
	mockRepo.On("Update", patient).Return(patient, nil)

	result, err := service.Update(patient)
	assert.NoError(t, err)
	assert.Equal(t, patient, result)
}

func TestUpdatePatient_MissingID(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	patient := &domain.Patient{ID: "", Name: "NoID"}
	result, err := service.Update(patient)
	assert.ErrorIs(t, err, ErrMissingID)
	assert.Nil(t, result)
}

func TestDeletePatient_Success(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	mockRepo.On("Delete", "321").Return(nil)
	err := service.Delete("321")
	assert.NoError(t, err)
}

func TestDeletePatient_MissingID(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	err := service.Delete("")
	assert.ErrorIs(t, err, ErrMissingID)
}

func TestGetAllPatients_Success(t *testing.T) {
	mockRepo := new(MockPatientRepo)
	mockPublisher := new(MockPatientEventPublisher)
	service := NewPatientService(mockRepo, mockPublisher)

	list := []domain.Patient{{Name: "A"}, {Name: "B"}}
	mockRepo.On("GetAll").Return(list, nil)

	result, err := service.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}
