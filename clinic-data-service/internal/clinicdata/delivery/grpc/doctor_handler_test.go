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
type MockDoctorService struct {
	mock.Mock
}

func (m *MockDoctorService) Create(d *domain.Doctor) (*domain.Doctor, error) {
	args := m.Called(d)
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorService) GetByID(id string) (*domain.Doctor, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorService) GetAll() ([]domain.Doctor, error) {
	args := m.Called()
	return args.Get(0).([]domain.Doctor), args.Error(1)
}
func (m *MockDoctorService) Update(d *domain.Doctor) (*domain.Doctor, error) {
	args := m.Called(d)
	return args.Get(0).(*domain.Doctor), args.Error(1)
}
func (m *MockDoctorService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// --------- HELPER: Context Metadata ---------

func stringPtrDoctor(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// --------- TEST CASES ---------

func TestCreateDoctor_Handler_WithAuditMetadata(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	ctx := withAuditMD(context.Background(), "user321", "dr. Budi", "budi@klinik.com", "admin")
	req := &clinicdatapb.CreateDoctorRequest{
		Name:              "dr. Budi",
		Email:             "budi@klinik.com",
		SpecializationId:  "sp-123",
		Age:               40,
		ConsultationFee:   120000,
		YearsOfExperience: 10,
		LicenseNumber:     "LIS1234567",
		PhoneNumber:       "0811000001",
	}

	expected := &domain.Doctor{
		Name:              "dr. Budi",
		Email:             "budi@klinik.com",
		SpecializationID:  "sp-123",
		Age:               40,
		ConsultationFee:   120000,
		YearsOfExperience: 10,
		LicenseNumber:     "LIS1234567",
		PhoneNumber:       stringPtrDoctor("0811000001"),
		IsActive:          false,
		CreatedBy:         stringPtrDoctor("user321"),
		CreatedName:       "dr. Budi",
		CreatedEmail:      "budi@klinik.com",
		CreatedRole:       "admin",
		UpdatedBy:         stringPtrDoctor("user321"),
		UpdatedName:       "dr. Budi",
		UpdatedEmail:      "budi@klinik.com",
		UpdatedRole:       "admin",
	}
	created := *expected
	created.ID = "dr-1"
	created.CreatedAt = time.Now()
	created.UpdatedAt = time.Now()

	mockSvc.On("Create", mock.MatchedBy(func(d *domain.Doctor) bool {
		return d.Name == expected.Name &&
			d.Email == expected.Email &&
			d.SpecializationID == expected.SpecializationID &&
			d.Age == expected.Age &&
			d.ConsultationFee == expected.ConsultationFee &&
			d.YearsOfExperience == expected.YearsOfExperience &&
			d.LicenseNumber == expected.LicenseNumber &&
			d.PhoneNumber != nil && *d.PhoneNumber == "0811000001" &&
			d.CreatedBy != nil && *d.CreatedBy == "user321" &&
			d.CreatedName == "dr. Budi" &&
			d.CreatedEmail == "budi@klinik.com" &&
			d.CreatedRole == "admin"
	})).Return(&created, nil)

	resp, err := handler.CreateDoctor(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "dr. Budi", resp.Name)
	assert.Equal(t, "budi@klinik.com", resp.Email)
	assert.Equal(t, "sp-123", resp.SpecializationId)
	assert.Equal(t, int32(40), resp.Age)
	assert.Equal(t, 120000.0, resp.ConsultationFee)
	assert.Equal(t, int32(10), resp.YearsOfExperience)
	assert.Equal(t, "LIS1234567", resp.LicenseNumber)
	assert.Equal(t, "0811000001", resp.PhoneNumber)
}

func TestCreateDoctor_Handler_ErrorFromService(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	ctx := withAuditMD(context.Background(), "user321", "dr. Budi", "budi@klinik.com", "admin")
	req := &clinicdatapb.CreateDoctorRequest{
		Name:              "dr. Budi",
		Email:             "budi@klinik.com",
		SpecializationId:  "sp-123",
		Age:               40,
		ConsultationFee:   120000,
		YearsOfExperience: 10,
		LicenseNumber:     "LIS1234567",
	}
	mockSvc.On("Create", mock.AnythingOfType("*domain.Doctor")).Return((*domain.Doctor)(nil), app.ErrEmailAlreadyExists)

	resp, err := handler.CreateDoctor(ctx, req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestGetDoctorByID_Handler_Success(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	expected := &domain.Doctor{
		ID:    "d-1",
		Name:  "dr. Budi",
		Email: "budi@klinik.com",
	}
	mockSvc.On("GetByID", "d-1").Return(expected, nil)

	resp, err := handler.GetDoctorByID(context.Background(), &clinicdatapb.GetDoctorByIDRequest{Id: "d-1"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "dr. Budi", resp.Name)
	assert.Equal(t, "d-1", resp.Id)
}

func TestGetDoctorByID_Handler_NotFound(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	mockSvc.On("GetByID", "notfound").Return((*domain.Doctor)(nil), app.ErrNotFound)
	resp, err := handler.GetDoctorByID(context.Background(), &clinicdatapb.GetDoctorByIDRequest{Id: "notfound"})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestGetAllDoctors_Handler(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	mockSvc.On("GetAll").Return([]domain.Doctor{
		{ID: "1", Name: "dr. A"},
		{ID: "2", Name: "dr. B"},
	}, nil)

	resp, err := handler.GetAllDoctors(context.Background(), &clinicdatapb.Empty{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Doctors, 2)
	assert.Equal(t, "dr. A", resp.Doctors[0].Name)
}

func TestUpdateDoctor_Handler(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	ctx := withAuditMD(context.Background(), "user321", "dr. Budi", "budi@klinik.com", "admin")
	req := &clinicdatapb.UpdateDoctorRequest{
		Id:                "d-1",
		Name:              "Update",
		Email:             "update@klinik.com",
		SpecializationId:  "sp-123",
		Age:               42,
		ConsultationFee:   130000,
		YearsOfExperience: 15,
		LicenseNumber:     "LIS7654321",
		PhoneNumber:       "08998877",
		IsActive:          true,
	}

	updated := &domain.Doctor{ID: "d-1", Name: "Update", Email: "update@klinik.com"}
	mockSvc.On("Update", mock.AnythingOfType("*domain.Doctor")).Return(updated, nil)

	resp, err := handler.UpdateDoctor(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Update", resp.Name)
	assert.Equal(t, "d-1", resp.Id)
}

func TestDeleteDoctor_Handler(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	doctor := &domain.Doctor{ID: "d-1", Name: "dr. Budi", Email: "budi@klinik.com"}
	mockSvc.On("GetByID", "d-1").Return(doctor, nil)
	mockSvc.On("Delete", "d-1").Return(nil)

	req := &clinicdatapb.DeleteDoctorRequest{Id: "d-1"}
	resp, err := handler.DeleteDoctor(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "dr. Budi", resp.Name)
	assert.Equal(t, "d-1", resp.Id)
}

func TestDeleteDoctor_Handler_NotFound(t *testing.T) {
	mockSvc := new(MockDoctorService)
	handler := &GRPCHandler{doctorService: mockSvc}

	mockSvc.On("GetByID", "notfound").Return((*domain.Doctor)(nil), app.ErrNotFound)
	req := &clinicdatapb.DeleteDoctorRequest{Id: "notfound"}
	resp, err := handler.DeleteDoctor(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}
