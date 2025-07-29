package grpc

import (
	"clinic-data-service/internal/clinicdata/app"
	"clinic-data-service/internal/clinicdata/delivery/grpc/clinicdatapb"
)

// GRPCHandler struct untuk gRPC server
type GRPCHandler struct {
	clinicdatapb.UnimplementedClinicDataServiceServer
	patientService        app.PatientService
	doctorService         app.DoctorService
	specializationService app.SpecializationService
}

// Constructor NewGRPCHandler menerima semua dependency service
func NewGRPCHandler(
	patientSvc app.PatientService,
	doctorSvc app.DoctorService,
	specSvc app.SpecializationService,
) *GRPCHandler {
	return &GRPCHandler{
		patientService:        patientSvc,
		doctorService:         doctorSvc,
		specializationService: specSvc,
	}
}
