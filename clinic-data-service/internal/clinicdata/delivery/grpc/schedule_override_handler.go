package grpc

import (
	"context"

	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	audit "clinic-data-service/internal/clinicdata/delivery/grpc/utils"
	"clinic-data-service/internal/clinicdata/domain"
)

// CreateScheduleOverride handles creation
func (h *GRPCHandler) CreateScheduleOverride(ctx context.Context, req *clinicdatapb.CreateScheduleOverrideRequest) (*clinicdatapb.ScheduleOverride, error) {
	createdBy, createdName, createdEmail, createdRole := audit.ExtractAudit(ctx)

	override := &domain.ScheduleOverride{
		DoctorID:     req.DoctorId,
		RoomID:       req.RoomId,
		DayOfWeek:    int(req.DayOfWeek),
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Status:       req.Status,
		CreatedBy:    stringOrNil(createdBy),
		CreatedName:  createdName,
		CreatedEmail: createdEmail,
		CreatedRole:  createdRole,
		UpdatedBy:    stringOrNil(createdBy),
		UpdatedName:  createdName,
		UpdatedEmail: createdEmail,
		UpdatedRole:  createdRole,
	}

	created, err := h.scheduleOverrideService.Create(override)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}
	return mapOverrideToPB(created), nil
}

// GetOverrideByDoctorID fetches all override by doctor ID
func (h *GRPCHandler) GetOverrideByDoctorID(ctx context.Context, req *clinicdatapb.GetOverrideByDoctorIDRequest) (*clinicdatapb.ListScheduleOverrideResponse, error) {
	list, err := h.scheduleOverrideService.GetByDoctorID(req.DoctorId)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	var pbOverrides []*clinicdatapb.ScheduleOverride
	for _, o := range list {
		copy := o
		pbOverrides = append(pbOverrides, mapOverrideToPB(&copy))
	}

	return &clinicdatapb.ListScheduleOverrideResponse{Overrides: pbOverrides}, nil
}

// UpdateScheduleOverride updates an override
func (h *GRPCHandler) UpdateScheduleOverride(ctx context.Context, req *clinicdatapb.UpdateScheduleOverrideRequest) (*clinicdatapb.UpdateScheduleOverrideResponse, error) {
	updatedBy, updatedName, updatedEmail, updatedRole := audit.ExtractAudit(ctx)

	override := &domain.ScheduleOverride{
		ID:           req.Id,
		RoomID:       req.RoomId,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Status:       req.Status,
		UpdatedBy:    stringOrNil(updatedBy),
		UpdatedName:  updatedName,
		UpdatedEmail: updatedEmail,
		UpdatedRole:  updatedRole,
	}

	updated, err := h.scheduleOverrideService.Update(override)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return &clinicdatapb.UpdateScheduleOverrideResponse{
		Id:        updated.ID,
		DoctorId:  updated.DoctorID,
		DayOfWeek: int32(updated.DayOfWeek),
		Message:   "schedule override updated successfully",
	}, nil
}

// DeleteScheduleOverride deletes an override by ID
func (h *GRPCHandler) DeleteScheduleOverride(ctx context.Context, req *clinicdatapb.DeleteScheduleOverrideRequest) (*clinicdatapb.DeleteScheduleOverrideResponse, error) {
	err := h.scheduleOverrideService.Delete(req.Id)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}
	return &clinicdatapb.DeleteScheduleOverrideResponse{
		Id:      req.Id,
		Message: "schedule override deleted successfully",
	}, nil
}

// --- Helper ---

func mapOverrideToPB(o *domain.ScheduleOverride) *clinicdatapb.ScheduleOverride {
	return &clinicdatapb.ScheduleOverride{
		Id:        o.ID,
		DoctorId:  o.DoctorID,
		RoomId:    o.RoomID,
		DayOfWeek: int32(o.DayOfWeek),
		StartTime: o.StartTime,
		EndTime:   o.EndTime,
		Status:    o.Status,
	}
}
