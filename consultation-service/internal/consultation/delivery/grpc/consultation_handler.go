package grpc

import (
	"context"
	"ehSehat/consultation-service/internal/consultation/domain"
	consultationPb "ehSehat/proto/consultation"
	"time"
)

type consultationHandler struct {
	consultationPb.UnimplementedConsultationServiceServer
	app domain.ConsultationService
}

func NewConsultationHandler(app domain.ConsultationService) *consultationHandler {
	return &consultationHandler{app: app}
}

// My proto file
// message Prescription {
//   string name = 1; // name of the medicine
//   string dose = 2; // dosage instructions
// }

// message ConsultationRequest {
//   string user_id = 1; // unique identifier for the consultation
//   string patient_id = 2; // from master-service
//   string patient_name = 3; // snapshot, optional
//   string doctor_id = 4; // from master-service
//   string doctor_name = 5; // snapshot, optional
//   string room_id = 6; // from master-service
//   string room_name = 7; // snapshot, optional
//   string symptoms = 9;
//   repeated Prescription prescription = 8; // list of prescriptions
//   string diagnosis = 10;
//   string date = 11; // date of the consultation
// }

// message ConsultationResponse {
//     string id = 1; // unique identifier for the consultation
//     string user_id = 2; // unique identifier for the consultation
//     string patient_id = 3; // from master-service
//     string patient_name = 4; // snapshot, optional
//     string doctor_id = 5; // from master-service
//     string doctor_name = 6; // snapshot, optional
//     string room_id = 7; // from master-service
//     string room_name = 8; // snapshot, optional
//     string symptoms = 9;
//     repeated Prescription prescription = 10; // list of prescriptions
//     string diagnosis = 11;
//     string date = 12;
//     google.protobuf.Timestamp created_at = 13;
//     google.protobuf.Timestamp updated_at = 14;
// }

// message ConsultationIDRequest {
//   string id = 1; // unique identifier for the consultation
// }

func (h *consultationHandler) CreateConsultation(ctx context.Context, req *consultationPb.ConsultationRequest) (*consultationPb.ConsultationResponse, error) {
	dateFormatted, _ := time.Parse("2006-01-02", req.Date)

	prescriptionFormatted := make([]domain.Prescription, len(req.Prescription))
	for i, p := range req.Prescription {
		prescriptionFormatted[i] = domain.Prescription{
			MedicineName: p.Name,
			Dose:         p.Dose,
		}
	}

	consultation := &domain.Consultation{
		UserID:       req.UserId,
		PatientID:    req.PatientId,
		PatientName:  req.PatientName,
		DoctorID:     req.DoctorId,
		DoctorName:   req.DoctorName,
		RoomID:       req.RoomId,
		RoomName:     req.RoomName,
		Symptoms:     req.Symptoms,
		Prescription: prescriptionFormatted,
		Diagnosis:    req.Diagnosis,
		Date:         dateFormatted,
	}

	err := h.app.CreateConsultation(ctx, consultation)
	if err != nil {
		return nil, err
	}

	return &consultationPb.ConsultationResponse{
		Id:           consultation.ID,
		UserId:       consultation.UserID,
		PatientId:    consultation.PatientID,
		PatientName:  consultation.PatientName,
		DoctorId:     consultation.DoctorID,
		DoctorName:   consultation.DoctorName,
		RoomId:       consultation.RoomID,
		RoomName:     consultation.RoomName,
		Symptoms:     consultation.Symptoms,
		Prescription: req.Prescription,
		Diagnosis:    consultation.Diagnosis,
		Date:         consultation.Date.Format("2006-01-02"),
	}, nil
}

func (h *consultationHandler) FindByIDConsultation(ctx context.Context, req *consultationPb.ConsultationIDRequest) (*consultationPb.ConsultationResponse, error) {
	consultation, err := h.app.FindByIDConsultation(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	prescriptionFormatted := make([]*consultationPb.Prescription, len(consultation.Prescription))
	for i, p := range consultation.Prescription {
		prescriptionFormatted[i] = &consultationPb.Prescription{
			Name: p.MedicineName,
			Dose: p.Dose,
		}
	}

	return &consultationPb.ConsultationResponse{
		Id:           consultation.ID,
		UserId:       consultation.UserID,
		PatientId:    consultation.PatientID,
		PatientName:  consultation.PatientName,
		DoctorId:     consultation.DoctorID,
		DoctorName:   consultation.DoctorName,
		RoomId:       consultation.RoomID,
		RoomName:     consultation.RoomName,
		Symptoms:     consultation.Symptoms,
		Prescription: prescriptionFormatted,
		Diagnosis:    consultation.Diagnosis,
		Date:         consultation.Date.Format("2006-01-02"),
	}, nil
}
