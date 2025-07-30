package grpc

import (
	"context"

	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	audit "clinic-data-service/internal/clinicdata/delivery/grpc/utils"
	"clinic-data-service/internal/clinicdata/domain"
)

// CreateScheduleFixed handles fixed schedule creation
func (h *GRPCHandler) CreateScheduleFixed(ctx context.Context, req *clinicdatapb.CreateScheduleFixedRequest) (*clinicdatapb.ScheduleFixed, error) {
	createdBy, createdName, createdEmail, createdRole := audit.ExtractAudit(ctx)

	schedule := &domain.ScheduleFixed{
		DoctorID:  req.DoctorId,
		RoomID:    req.RoomId,
		DayOfWeek: int(req.DayOfWeek),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,

		CreatedBy:    stringOrNil(createdBy),
		CreatedName:  createdName,
		CreatedEmail: createdEmail,
		CreatedRole:  createdRole,
		UpdatedBy:    stringOrNil(createdBy),
		UpdatedName:  createdName,
		UpdatedEmail: createdEmail,
		UpdatedRole:  createdRole,
	}

	created, err := h.scheduleFixedService.Create(schedule)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return mapScheduleToPB(created), nil
}

// GetFixedSchedulesByDoctorID fetches fixed schedules for a doctor
func (h *GRPCHandler) GetFixedSchedulesByDoctorID(ctx context.Context, req *clinicdatapb.GetFixedSchedulesByDoctorIDRequest) (*clinicdatapb.ListScheduleFixedResponse, error) {
	schedules, err := h.scheduleFixedService.GetByDoctorID(req.DoctorId)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	var pbSchedules []*clinicdatapb.ScheduleFixed
	for _, s := range schedules {
		sCopy := s
		pbSchedules = append(pbSchedules, mapScheduleToPB(&sCopy))
	}

	return &clinicdatapb.ListScheduleFixedResponse{Schedules: pbSchedules}, nil
}

// UpdateScheduleFixed updates a fixed schedule
func (h *GRPCHandler) UpdateScheduleFixed(ctx context.Context, req *clinicdatapb.UpdateScheduleFixedRequest) (*clinicdatapb.ScheduleFixed, error) {
	updatedBy, updatedName, updatedEmail, updatedRole := audit.ExtractAudit(ctx)

	schedule := &domain.ScheduleFixed{
		ID:        req.Id,
		RoomID:    req.RoomId,
		DayOfWeek: int(req.DayOfWeek),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,

		UpdatedBy:    stringOrNil(updatedBy),
		UpdatedName:  updatedName,
		UpdatedEmail: updatedEmail,
		UpdatedRole:  updatedRole,
	}

	updated, err := h.scheduleFixedService.Update(schedule)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return mapScheduleToPB(updated), nil
}

// --- Helpers ---

func mapScheduleToPB(s *domain.ScheduleFixed) *clinicdatapb.ScheduleFixed {
	return &clinicdatapb.ScheduleFixed{
		Id:        s.ID,
		DoctorId:  s.DoctorID,
		RoomId:    s.RoomID,
		DayOfWeek: int32(s.DayOfWeek),
		StartTime: s.StartTime,
		EndTime:   s.EndTime,
		Status:    s.Status,
	}
}
