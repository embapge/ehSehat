package grpc

import (
	"clinic-data-service/internal/clinicdata/app"
	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
)

type GRPCHandler struct {
	clinicdatapb.UnimplementedClinicDataServiceServer
	patientService          app.PatientService
	doctorService           app.DoctorService
	specializationService   app.SpecializationService
	roomService             app.RoomService
	scheduleFixedService    app.ScheduleFixedService
	scheduleOverrideService app.ScheduleOverrideService
}

// NewGRPCHandler constructs GRPCHandler with all services
func NewGRPCHandler(
	patientSvc app.PatientService,
	doctorSvc app.DoctorService,
	specSvc app.SpecializationService,
	roomSvc app.RoomService,
	scheduleFixedSvc app.ScheduleFixedService,
	scheduleOverrideSvc app.ScheduleOverrideService,
) *GRPCHandler {
	return &GRPCHandler{
		patientService:          patientSvc,
		doctorService:           doctorSvc,
		specializationService:   specSvc,
		roomService:             roomSvc,
		scheduleFixedService:    scheduleFixedSvc,
		scheduleOverrideService: scheduleOverrideSvc,
	}
}
