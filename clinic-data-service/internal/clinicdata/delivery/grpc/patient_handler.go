package grpc

import (
	"context"

	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	audit "clinic-data-service/internal/clinicdata/delivery/grpc/utils"
	"clinic-data-service/internal/clinicdata/domain"
)

// CreatePatient handles patient creation
func (h *GRPCHandler) CreatePatient(ctx context.Context, req *clinicdatapb.CreatePatientRequest) (*clinicdatapb.Patient, error) {
	createdBy, createdName, createdEmail, createdRole := audit.ExtractAudit(ctx)

	patient := &domain.Patient{
		Name:        req.Name,
		Email:       req.Email,
		BirthDate:   req.BirthDate,
		Gender:      req.Gender,
		PhoneNumber: stringPtr(req.PhoneNumber),
		Address:     stringPtr(req.Address),

		CreatedBy:    stringOrNil(createdBy),
		CreatedName:  createdName,
		CreatedEmail: createdEmail,
		CreatedRole:  createdRole,
		UpdatedBy:    stringOrNil(createdBy),
		UpdatedName:  createdName,
		UpdatedEmail: createdEmail,
		UpdatedRole:  createdRole,
	}

	created, err := h.patientService.Create(patient)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return mapDomainToPB(created), nil
}

// GetPatientByID fetches patient by ID
func (h *GRPCHandler) GetPatientByID(ctx context.Context, req *clinicdatapb.GetPatientByIDRequest) (*clinicdatapb.Patient, error) {
	patient, err := h.patientService.GetByID(req.Id)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}
	return mapDomainToPB(patient), nil
}

// GetAllPatients returns all patients
func (h *GRPCHandler) GetAllPatients(ctx context.Context, _ *clinicdatapb.Empty) (*clinicdatapb.ListPatientsResponse, error) {
	patients, err := h.patientService.GetAll()
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	var pbPatients []*clinicdatapb.Patient
	for _, p := range patients {
		pCopy := p
		pbPatients = append(pbPatients, mapDomainToPB(&pCopy))
	}

	return &clinicdatapb.ListPatientsResponse{Patients: pbPatients}, nil
}

// UpdatePatient updates existing patient
func (h *GRPCHandler) UpdatePatient(ctx context.Context, req *clinicdatapb.UpdatePatientRequest) (*clinicdatapb.UpdatePatientResponse, error) {
	updatedBy, updatedName, updatedEmail, updatedRole := audit.ExtractAudit(ctx)

	p := &domain.Patient{
		ID:          req.Id,
		Name:        req.Name,
		Email:       req.Email,
		BirthDate:   req.BirthDate,
		Gender:      req.Gender,
		PhoneNumber: stringPtr(req.PhoneNumber),
		Address:     stringPtr(req.Address),

		UpdatedBy:    stringOrNil(updatedBy),
		UpdatedName:  updatedName,
		UpdatedEmail: updatedEmail,
		UpdatedRole:  updatedRole,
	}

	updated, err := h.patientService.Update(p)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.UpdatePatientResponse{
		Id:      updated.ID,
		UserId:  updated.UserIDOrEmpty(),
		Name:    updated.Name,
		Email:   updated.Email,
		Message: "patient updated successfully",
	}, nil
}

// DeletePatient deletes a patient
func (h *GRPCHandler) DeletePatient(ctx context.Context, req *clinicdatapb.DeletePatientRequest) (*clinicdatapb.DeletePatientResponse, error) {
	patient, err := h.patientService.GetByID(req.Id)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	if err := h.patientService.Delete(req.Id); err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.DeletePatientResponse{
		Id:      patient.ID,
		UserId:  patient.UserIDOrEmpty(),
		Name:    patient.Name,
		Email:   patient.Email,
		Message: "patient deleted successfully",
	}, nil
}

// --- Helpers ---

func mapDomainToPB(p *domain.Patient) *clinicdatapb.Patient {
	return &clinicdatapb.Patient{
		Id:          p.ID,
		UserId:      p.UserIDOrEmpty(),
		Name:        p.Name,
		Email:       p.Email,
		BirthDate:   p.BirthDate,
		Gender:      p.Gender,
		PhoneNumber: p.PhoneNumberOrEmpty(),
		Address:     p.AddressOrEmpty(),
	}
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func stringOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
