package grpc

import (
	"appointment-queue-service/internal/appointment/app"
	"appointment-queue-service/internal/appointment/delivery/grpc/pb"
	"appointment-queue-service/internal/appointment/domain"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppointmentHandler struct {
	pb.UnimplementedAppointmentServiceServer
	App app.AppointmentApp
}

func NewAppointmentHandler(app app.AppointmentApp) pb.AppointmentServiceServer {
	return &AppointmentHandler{
		App: app,
	}
}

func (h *AppointmentHandler) GetAppointmentByID(ctx context.Context, req *pb.AppointmentIDRequest) (*pb.AppointmentResponse, error) {
	appt, err := h.App.FindByIDAppointment(ctx, uint(req.Id))
	if err != nil {
		// return nil, status.Errorf(codes.NotFound, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	return &pb.AppointmentResponse{
		Appointment: toProto(*appt),
	}, nil
}

func (h *AppointmentHandler) GetAppointmentsByUserID(ctx context.Context, req *pb.UserAppointmentsRequest) (*pb.AppointmentListResponse, error) {
	appointments, err := h.App.FindByUserID(ctx, uint(req.UserId))
	if err != nil {
		// return nil, status.Errorf(codes.Internal, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	var protoAppointments []*pb.Appointment
	for _, a := range appointments {
		protoAppointments = append(protoAppointments, toProto(*a))
	}

	return &pb.AppointmentListResponse{
		Appointments: protoAppointments,
	}, nil
}

func (h *AppointmentHandler) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.AppointmentResponse, error) {
	appt := &domain.AppointmentModel{
		UserID:   uint(req.UserId),
		DoctorID: uint(req.DoctorId),
		// ScheduleID: uint(req.ScheduleId),
		// Status:     req.Status,
		// Notes:      req.Notes,
	}

	createdAppt, err := h.App.CreateAppointment(ctx, appt)
	if err != nil {
		// return nil, status.Errorf(codes.InvalidArgument, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	return &pb.AppointmentResponse{
		Appointment: toProto(*createdAppt),
	}, nil
}

func (h *AppointmentHandler) UpdateAppointment(ctx context.Context, req *pb.UpdateAppointmentRequest) (*pb.AppointmentResponse, error) {
	appt := &domain.AppointmentModel{
		ID:       uint(req.Id),
		UserID:   uint(req.UserId),
		DoctorID: uint(req.DoctorId),
		// ScheduleID: uint(req.ScheduleId),
		Status: req.Status,
		// Notes:      req.Notes,
	}

	updatedAppt, err := h.App.UpdateAppointment(ctx, appt.ID, appt)
	if err != nil {
		// return nil, status.Errorf(codes.InvalidArgument, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	return &pb.AppointmentResponse{
		Appointment: toProto(*updatedAppt),
	}, nil
}

func (h *AppointmentHandler) MarkAppointmentAsPaid(ctx context.Context, req *pb.MarkAsPaidRequest) (*pb.EmptyResponse, error) {
	err := h.App.MarkAsPaid(ctx, uint(req.Id))
	if err != nil {
		// return nil, status.Errorf(codes.Internal, err.Error())
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	return &pb.EmptyResponse{}, nil
}

// helper function to convert domain.AppointmentModel ke pb.Appointment
func toProto(a domain.AppointmentModel) *pb.Appointment {
	return &pb.Appointment{
		Id:       uint32(a.ID),
		UserId:   uint32(a.UserID),
		DoctorId: uint32(a.DoctorID),
		// ScheduleId: uint32(a.ScheduleID),
		Status: a.Status,
		// Notes:      a.Notes,
		CreatedAt: a.CreatedAt.Format(time.RFC3339),
		UpdatedAt: a.UpdatedAt.Format(time.RFC3339),
	}
}
