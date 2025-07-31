package grpc

import (
	"context"

	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	audit "clinic-data-service/internal/clinicdata/delivery/grpc/utils"
	"clinic-data-service/internal/clinicdata/domain"
)

// CreateDoctor handles gRPC request to create a doctor
func (h *GRPCHandler) CreateDoctor(ctx context.Context, req *clinicdatapb.CreateDoctorRequest) (*clinicdatapb.Doctor, error) {
	createdBy, createdName, createdEmail, createdRole := audit.ExtractAudit(ctx)

	doctor := &domain.Doctor{
		Name:              req.GetName(),
		Email:             req.GetEmail(),
		SpecializationID:  req.GetSpecializationId(),
		Age:               int(req.GetAge()),
		ConsultationFee:   req.GetConsultationFee(),
		YearsOfExperience: int(req.GetYearsOfExperience()),
		LicenseNumber:     req.GetLicenseNumber(),
		PhoneNumber:       stringPtr(req.GetPhoneNumber()),
		IsActive:          false, // default inactive

		CreatedBy:    stringOrNil(createdBy),
		CreatedName:  createdName,
		CreatedEmail: createdEmail,
		CreatedRole:  createdRole,
		UpdatedBy:    stringOrNil(createdBy),
		UpdatedName:  createdName,
		UpdatedEmail: createdEmail,
		UpdatedRole:  createdRole,
	}

	created, err := h.doctorService.Create(doctor)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return mapDoctorToPB(created), nil
}

// GetDoctorByID returns a doctor by ID
func (h *GRPCHandler) GetDoctorByID(ctx context.Context, req *clinicdatapb.GetDoctorByIDRequest) (*clinicdatapb.Doctor, error) {
	doctor, err := h.doctorService.GetByID(req.GetId())
	if err != nil {
		return nil, mapErrorToStatus(err)
	}
	return mapDoctorToPB(doctor), nil
}

// GetAllDoctors returns all doctors
func (h *GRPCHandler) GetAllDoctors(ctx context.Context, _ *clinicdatapb.Empty) (*clinicdatapb.ListDoctorsResponse, error) {
	doctors, err := h.doctorService.GetAll()
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	var pbDoctors []*clinicdatapb.Doctor
	for _, d := range doctors {
		dCopy := d
		pbDoctors = append(pbDoctors, mapDoctorToPB(&dCopy))
	}
	return &clinicdatapb.ListDoctorsResponse{Doctors: pbDoctors}, nil
}

// UpdateDoctor updates a doctor's data
func (h *GRPCHandler) UpdateDoctor(ctx context.Context, req *clinicdatapb.UpdateDoctorRequest) (*clinicdatapb.UpdateDoctorResponse, error) {
	updatedBy, updatedName, updatedEmail, updatedRole := audit.ExtractAudit(ctx)

	doctor := &domain.Doctor{
		ID:                req.GetId(),
		Name:              req.GetName(),
		Email:             req.GetEmail(),
		SpecializationID:  req.GetSpecializationId(),
		Age:               int(req.GetAge()),
		ConsultationFee:   req.GetConsultationFee(),
		YearsOfExperience: int(req.GetYearsOfExperience()),
		LicenseNumber:     req.GetLicenseNumber(),
		PhoneNumber:       stringPtr(req.GetPhoneNumber()),
		IsActive:          req.GetIsActive(),

		UpdatedBy:    stringOrNil(updatedBy),
		UpdatedName:  updatedName,
		UpdatedEmail: updatedEmail,
		UpdatedRole:  updatedRole,
	}

	updated, err := h.doctorService.Update(doctor)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.UpdateDoctorResponse{
		Id:      updated.ID,
		UserId:  updated.UserIDOrEmpty(),
		Name:    updated.Name,
		Email:   updated.Email,
		Message: "doctor updated successfully",
	}, nil
}

// DeleteDoctor deletes a doctor
func (h *GRPCHandler) DeleteDoctor(ctx context.Context, req *clinicdatapb.DeleteDoctorRequest) (*clinicdatapb.DeleteDoctorResponse, error) {
	doctor, err := h.doctorService.GetByID(req.GetId())
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	if err := h.doctorService.Delete(req.GetId()); err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.DeleteDoctorResponse{
		Id:      doctor.ID,
		UserId:  doctor.UserIDOrEmpty(),
		Name:    doctor.Name,
		Email:   doctor.Email,
		Message: "doctor deleted successfully",
	}, nil
}

// --- Helper ---
func mapDoctorToPB(d *domain.Doctor) *clinicdatapb.Doctor {
	return &clinicdatapb.Doctor{
		Id:                d.ID,
		UserId:            d.UserIDOrEmpty(),
		Name:              d.Name,
		Email:             d.Email,
		SpecializationId:  d.SpecializationID,
		Age:               int32(d.Age),
		ConsultationFee:   d.ConsultationFee,
		YearsOfExperience: int32(d.YearsOfExperience),
		LicenseNumber:     d.LicenseNumber,
		PhoneNumber:       d.PhoneNumberOrEmpty(),
		IsActive:          d.IsActive,
	}
}
