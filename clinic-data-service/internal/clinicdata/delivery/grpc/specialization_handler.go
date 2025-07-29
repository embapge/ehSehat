package grpc

import (
	"context"

	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	audit "clinic-data-service/internal/clinicdata/delivery/grpc/utils"
	"clinic-data-service/internal/clinicdata/domain"
)

// CreateSpecialization handles creation
func (h *GRPCHandler) CreateSpecialization(ctx context.Context, req *clinicdatapb.CreateSpecializationRequest) (*clinicdatapb.Specialization, error) {
	createdBy, createdName, createdEmail, createdRole := audit.ExtractAudit(ctx)

	s := &domain.Specialization{
		Name:         req.Name,
		CreatedBy:    stringOrNil(createdBy),
		CreatedName:  createdName,
		CreatedEmail: createdEmail,
		CreatedRole:  createdRole,
		UpdatedBy:    stringOrNil(createdBy),
		UpdatedName:  createdName,
		UpdatedEmail: createdEmail,
		UpdatedRole:  createdRole,
	}

	created, err := h.specializationService.Create(s)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return mapSpecToPB(created), nil
}

// GetSpecializationByID
func (h *GRPCHandler) GetSpecializationByID(ctx context.Context, req *clinicdatapb.GetSpecializationByIDRequest) (*clinicdatapb.Specialization, error) {
	s, err := h.specializationService.GetByID(req.Id)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}
	return mapSpecToPB(s), nil
}

// GetAllSpecializations
func (h *GRPCHandler) GetAllSpecializations(ctx context.Context, _ *clinicdatapb.Empty) (*clinicdatapb.ListSpecializationsResponse, error) {
	specs, err := h.specializationService.GetAll()
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	var pbSpecs []*clinicdatapb.Specialization
	for _, s := range specs {
		sCopy := s
		pbSpecs = append(pbSpecs, mapSpecToPB(&sCopy))
	}

	return &clinicdatapb.ListSpecializationsResponse{
		Specializations: pbSpecs,
	}, nil
}

// UpdateSpecialization
func (h *GRPCHandler) UpdateSpecialization(ctx context.Context, req *clinicdatapb.UpdateSpecializationRequest) (*clinicdatapb.UpdateSpecializationResponse, error) {
	updatedBy, updatedName, updatedEmail, updatedRole := audit.ExtractAudit(ctx)

	s := &domain.Specialization{
		ID:           req.Id,
		Name:         req.Name,
		UpdatedBy:    stringOrNil(updatedBy),
		UpdatedName:  updatedName,
		UpdatedEmail: updatedEmail,
		UpdatedRole:  updatedRole,
	}

	updated, err := h.specializationService.Update(s)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.UpdateSpecializationResponse{
		Id:      updated.ID,
		Name:    updated.Name,
		Message: "specialization updated successfully",
	}, nil
}

// DeleteSpecialization
func (h *GRPCHandler) DeleteSpecialization(ctx context.Context, req *clinicdatapb.DeleteSpecializationRequest) (*clinicdatapb.DeleteSpecializationResponse, error) {
	s, err := h.specializationService.GetByID(req.Id)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	if err := h.specializationService.Delete(req.Id); err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.DeleteSpecializationResponse{
		Id:      s.ID,
		Name:    s.Name,
		Message: "specialization deleted successfully",
	}, nil
}

// --- Helper ---

func mapSpecToPB(s *domain.Specialization) *clinicdatapb.Specialization {
	return &clinicdatapb.Specialization{
		Id:   s.ID,
		Name: s.Name,
	}
}
