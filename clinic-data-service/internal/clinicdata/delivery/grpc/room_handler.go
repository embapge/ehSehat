package grpc

import (
	"context"

	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
	audit "clinic-data-service/internal/clinicdata/delivery/grpc/utils"
	"clinic-data-service/internal/clinicdata/domain"
)

// CreateRoom handles room creation
func (h *GRPCHandler) CreateRoom(ctx context.Context, req *clinicdatapb.CreateRoomRequest) (*clinicdatapb.Room, error) {
	createdBy, createdName, createdEmail, createdRole := audit.ExtractAudit(ctx)

	room := &domain.Room{
		Name:     req.Name,
		IsActive: true,

		CreatedBy:    stringOrNil(createdBy),
		CreatedName:  createdName,
		CreatedEmail: createdEmail,
		CreatedRole:  createdRole,
		UpdatedBy:    stringOrNil(createdBy),
		UpdatedName:  createdName,
		UpdatedEmail: createdEmail,
		UpdatedRole:  createdRole,
	}

	created, err := h.roomService.Create(room)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	return mapRoomToPB(created), nil
}

// GetRoomByID returns room by ID
func (h *GRPCHandler) GetRoomByID(ctx context.Context, req *clinicdatapb.GetRoomByIDRequest) (*clinicdatapb.Room, error) {
	room, err := h.roomService.GetByID(req.Id)
	if err != nil {
		return nil, mapErrorToStatus(err)
	}
	return mapRoomToPB(room), nil
}

// GetAllRooms returns all rooms
func (h *GRPCHandler) GetAllRooms(ctx context.Context, _ *clinicdatapb.Empty) (*clinicdatapb.ListRoomsResponse, error) {
	rooms, err := h.roomService.GetAll()
	if err != nil {
		return nil, mapErrorToStatus(err)
	}

	var pbRooms []*clinicdatapb.Room
	for _, r := range rooms {
		rCopy := r
		pbRooms = append(pbRooms, mapRoomToPB(&rCopy))
	}

	return &clinicdatapb.ListRoomsResponse{Rooms: pbRooms}, nil
}

// --- Helper ---

func mapRoomToPB(r *domain.Room) *clinicdatapb.Room {
	return &clinicdatapb.Room{
		Id:       r.ID,
		Name:     r.Name,
		IsActive: r.IsActive,
	}
}
