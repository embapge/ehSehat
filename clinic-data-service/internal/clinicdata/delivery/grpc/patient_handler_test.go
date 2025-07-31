package grpc

import (
	"context"
	"testing"
	"time"

	"clinic-data-service/internal/clinicdata/app"
	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	"clinic-data-service/internal/clinicdata/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --------- MOCK SERVICE ---------
type MockPatientService struct {
	mock.Mock
}

func (m *MockPatientService) Create(p *domain.Patient) (*domain.Patient, error) {
	args := m.Called(p)
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientService) GetByID(id string) (*domain.Patient, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientService) GetAll() ([]domain.Patient, error) {
	args := m.Called()
	return args.Get(0).([]domain.Patient), args.Error(1)
}
func (m *MockPatientService) Update(p *domain.Patient) (*domain.Patient, error) {
	args := m.Called(p)
	return args.Get(0).(*domain.Patient), args.Error(1)
}
func (m *MockPatientService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// --------- HELPER ---------

func stringPtrPatient(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// --------- TEST CASES ---------
func TestCreatePatient_Handler_WithAuditMetadata(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	ctx := withAuditMD(context.Background(), "user123", "Jane Doe", "jane@example.com", "admin")
	req := &clinicdatapb.CreatePatientRequest{
		Name:        "Jane",
		Email:       "jane@example.com",
		BirthDate:   "2000-01-01",
		Gender:      "F",
		PhoneNumber: "0812345678",
		Address:     "Jakarta",
	}

	expected := &domain.Patient{
		Name:         "Jane",
		Email:        "jane@example.com",
		BirthDate:    "2000-01-01",
		Gender:       "F",
		PhoneNumber:  stringPtrPatient("0812345678"),
		Address:      stringPtrPatient("Jakarta"),
		CreatedBy:    stringPtrPatient("user123"),
		CreatedName:  "Jane Doe",
		CreatedEmail: "jane@example.com",
		CreatedRole:  "admin",
		UpdatedBy:    stringPtrPatient("user123"),
		UpdatedName:  "Jane Doe",
		UpdatedEmail: "jane@example.com",
		UpdatedRole:  "admin",
	}
	created := *expected
	created.ID = "patient-1"
	created.CreatedAt = time.Now()
	created.UpdatedAt = time.Now()

	mockSvc.On("Create", mock.MatchedBy(func(p *domain.Patient) bool {
		return p.Name == expected.Name &&
			p.Email == expected.Email &&
			p.BirthDate == expected.BirthDate &&
			p.Gender == expected.Gender &&
			p.PhoneNumber != nil && *p.PhoneNumber == "0812345678" &&
			p.Address != nil && *p.Address == "Jakarta" &&
			p.CreatedBy != nil && *p.CreatedBy == "user123" &&
			p.CreatedName == "Jane Doe" &&
			p.CreatedEmail == "jane@example.com" &&
			p.CreatedRole == "admin" &&
			p.UpdatedBy != nil && *p.UpdatedBy == "user123"
	})).Return(&created, nil)

	resp, err := handler.CreatePatient(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Jane", resp.Name)
	assert.Equal(t, "jane@example.com", resp.Email)
	assert.Equal(t, "2000-01-01", resp.BirthDate)
	assert.Equal(t, "F", resp.Gender)
	assert.Equal(t, "0812345678", resp.PhoneNumber)
	assert.Equal(t, "Jakarta", resp.Address)
}

func TestCreatePatient_Handler_ErrorFromService(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	ctx := withAuditMD(context.Background(), "user123", "Jane Doe", "jane@example.com", "admin")
	req := &clinicdatapb.CreatePatientRequest{
		Name:      "Jane",
		Email:     "jane@example.com",
		BirthDate: "2000-01-01",
		Gender:    "F",
	}

	mockSvc.On("Create", mock.AnythingOfType("*domain.Patient")).
		Return((*domain.Patient)(nil), app.ErrEmailAlreadyExists)

	resp, err := handler.CreatePatient(ctx, req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestGetPatientByID_Handler_Success(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	expected := &domain.Patient{
		ID:    "p-1",
		Name:  "Jane",
		Email: "jane@example.com",
	}
	mockSvc.On("GetByID", "p-1").Return(expected, nil)

	resp, err := handler.GetPatientByID(context.Background(), &clinicdatapb.GetPatientByIDRequest{Id: "p-1"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Jane", resp.Name)
	assert.Equal(t, "p-1", resp.Id)
}

func TestGetPatientByID_Handler_NotFound(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	mockSvc.On("GetByID", "not-found").Return((*domain.Patient)(nil), app.ErrNotFound)
	resp, err := handler.GetPatientByID(context.Background(), &clinicdatapb.GetPatientByIDRequest{Id: "not-found"})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestGetAllPatients_Handler(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	mockSvc.On("GetAll").Return([]domain.Patient{
		{ID: "1", Name: "Jane"},
		{ID: "2", Name: "Budi"},
	}, nil)

	resp, err := handler.GetAllPatients(context.Background(), &clinicdatapb.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Patients, 2)
	assert.Equal(t, "Jane", resp.Patients[0].Name)
}

func TestUpdatePatient_Handler(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	ctx := withAuditMD(context.Background(), "user123", "Jane Doe", "jane@example.com", "admin")
	req := &clinicdatapb.UpdatePatientRequest{
		Id:          "p-1",
		Name:        "Update",
		Email:       "update@mail.com",
		BirthDate:   "1990-01-01",
		Gender:      "F",
		PhoneNumber: "0811",
		Address:     "Yogyakarta",
	}

	updated := &domain.Patient{ID: "p-1", Name: "Update", Email: "update@mail.com"}
	mockSvc.On("Update", mock.AnythingOfType("*domain.Patient")).Return(updated, nil)

	resp, err := handler.UpdatePatient(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Update", resp.Name)
	assert.Equal(t, "p-1", resp.Id)
}

func TestDeletePatient_Handler(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	patient := &domain.Patient{ID: "p-1", Name: "Jane", Email: "jane@example.com"}
	mockSvc.On("GetByID", "p-1").Return(patient, nil)
	mockSvc.On("Delete", "p-1").Return(nil)

	req := &clinicdatapb.DeletePatientRequest{Id: "p-1"}
	resp, err := handler.DeletePatient(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Jane", resp.Name)
	assert.Equal(t, "p-1", resp.Id)
}

func TestDeletePatient_Handler_NotFound(t *testing.T) {
	mockSvc := new(MockPatientService)
	handler := &GRPCHandler{patientService: mockSvc}

	mockSvc.On("GetByID", "not-found").Return((*domain.Patient)(nil), app.ErrNotFound)
	req := &clinicdatapb.DeletePatientRequest{Id: "not-found"}
	resp, err := handler.DeletePatient(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}
