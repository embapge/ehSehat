package grpc

import (
	"context"
	consultationPb "ehSehat/consultation-service/internal/consultation/delivery/grpc/pb"
	"ehSehat/consultation-service/internal/consultation/domain"
	"ehSehat/libs/utils"
	"ehSehat/libs/utils/grpcmetadata"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type consultationHandler struct {
	consultationPb.UnimplementedConsultationServiceServer
	app domain.ConsultationService
	ch  *amqp.Channel
}

func NewConsultationHandler(app domain.ConsultationService, ch *amqp.Channel) *consultationHandler {
	return &consultationHandler{
		app: app,
		ch:  ch,
	}
}

func (h *consultationHandler) CreateConsultation(ctx context.Context, req *consultationPb.ConsultationRequest) (*consultationPb.ConsultationResponse, error) {
	if req == nil {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Request cannot be nil"))
	}
	if req.Patient == nil || req.Patient.Id == "" || req.Patient.Name == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Patient fields are required"))
	}
	if req.Room == nil || req.Room.Id == "" || req.Room.Name == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Room fields are required"))
	}
	if req.Doctor == nil || req.Doctor.Id == "" || req.Doctor.Name == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Doctor fields are required"))
	}
	if req.Symptoms == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Symptoms are required"))
	}
	if req.Diagnosis == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Diagnosis is required"))
	}
	if req.Date == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Date is required"))
	}
	if len(req.Prescription) == 0 {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Prescription is required"))
	}

	md, _ := grpcmetadata.GetMetadataFromContext(ctx)

	userSnapshot := map[string]interface{}{
		"ID":    "",
		"Name":  "",
		"Email": "",
		"Role":  "",
	}

	if v, ok := md["ts-user-id"]; ok && len(v) > 0 {
		userSnapshot["ID"] = v[0]
	}
	if v, ok := md["ts-user-name"]; ok && len(v) > 0 {
		userSnapshot["Name"] = v[0]
	}
	if v, ok := md["ts-user-email"]; ok && len(v) > 0 {
		userSnapshot["Email"] = v[0]
	}
	if v, ok := md["ts-user-role"]; ok && len(v) > 0 {
		userSnapshot["Role"] = v[0]
	}

	createdBySnapshot := domain.CreatedBySnapshot{
		ID:    userSnapshot["ID"].(string),
		Name:  userSnapshot["Name"].(string),
		Email: userSnapshot["Email"].(string),
		Role:  userSnapshot["Role"].(string),
	}

	updatedBySnapshot := domain.UpdatedBySnapshot{
		ID:    userSnapshot["ID"].(string),
		Name:  userSnapshot["Name"].(string),
		Email: userSnapshot["Email"].(string),
		Role:  userSnapshot["Role"].(string),
	}

	dateFormatted, _ := time.Parse("2006-01-02", req.Date)

	prescriptionFormatted := make([]domain.Prescription, len(req.Prescription))
	for i, p := range req.Prescription {
		prescriptionFormatted[i] = domain.Prescription{
			MedicineName: p.Name,
			Dose:         p.Dose,
		}
	}

	ageInt32 := int32(0)
	if req.Patient.Age != "" {
		if ageInt, err := strconv.Atoi(req.Patient.Age); err == nil {
			ageInt32 = int32(ageInt)
		}
	}

	patientSnapshot := domain.PatientSnapshot{
		ID:   req.Patient.Id,
		Name: req.Patient.Name,
		Age:  ageInt32,
	}

	doctorSnapshot := domain.DoctorSnapshot{
		ID:             req.Doctor.Id,
		Name:           req.Doctor.Name,
		Specialization: req.Doctor.Specialization,
	}

	roomSnapshot := domain.RoomSnapshot{
		ID:   req.Room.Id,
		Name: req.Room.Name,
	}

	consultation := &domain.Consultation{
		QueueID:       req.QueueId,
		AppointmentID: req.AppointmentId,
		CreatedBy:     createdBySnapshot,
		UpdatedBy:     updatedBySnapshot,
		Patient:       patientSnapshot,
		Doctor:        doctorSnapshot,
		Room:          roomSnapshot,
		Symptoms:      req.Symptoms,
		Prescription:  prescriptionFormatted,
		Diagnosis:     req.Diagnosis,
		Date:          dateFormatted,
	}

	err := h.app.CreateConsultation(ctx, consultation)
	if err != nil {
		return nil, err
	}

	var resQueID *wrapperspb.StringValue
	if consultation.QueueID != "" {
		resQueID = wrapperspb.String(consultation.QueueID)
	} else {
		resQueID = nil
	}

	var resAppID *wrapperspb.StringValue
	if consultation.AppointmentID != "" {
		resAppID = wrapperspb.String(consultation.AppointmentID)
	} else {
		resAppID = nil
	}

	return &consultationPb.ConsultationResponse{
		Id:            consultation.ID,
		QueueId:       resQueID,
		AppointmentId: resAppID,
		Patient: &consultationPb.PatientSnapshot{
			Id:   patientSnapshot.ID,
			Name: patientSnapshot.Name,
			Age:  strconv.Itoa(int(patientSnapshot.Age)),
		},
		Doctor: &consultationPb.DoctorSnapshot{
			Id:             doctorSnapshot.ID,
			Name:           doctorSnapshot.Name,
			Specialization: doctorSnapshot.Specialization,
		},
		Room: &consultationPb.RoomSnapshot{
			Id:   roomSnapshot.ID,
			Name: roomSnapshot.Name,
		},
		Symptoms:     consultation.Symptoms,
		Prescription: req.Prescription,
		Diagnosis:    consultation.Diagnosis,
		Date:         consultation.Date.Format("2006-01-02"),
		Status:       consultation.Status,
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

	userSnapshot := &consultationPb.UserSnapshot{
		Id:    consultation.User.ID,
		Name:  consultation.User.Name,
		Email: consultation.User.Email,
		Role:  consultation.User.Role,
	}

	patientSnap := &consultationPb.PatientSnapshot{
		Id:   consultation.Patient.ID,
		Name: consultation.Patient.Name,
		Age:  strconv.Itoa(int(consultation.Patient.Age)),
	}

	doctorSnap := &consultationPb.DoctorSnapshot{
		Id:             consultation.Doctor.ID,
		Name:           consultation.Doctor.Name,
		Specialization: consultation.Doctor.Specialization,
	}

	roomSnap := &consultationPb.RoomSnapshot{
		Id:   consultation.Room.ID,
		Name: consultation.Room.Name,
	}

	return &consultationPb.ConsultationResponse{
		Id:           consultation.ID,
		User:         userSnapshot,
		Patient:      patientSnap,
		Doctor:       doctorSnap,
		Room:         roomSnap,
		Symptoms:     consultation.Symptoms,
		Prescription: prescriptionFormatted,
		Diagnosis:    consultation.Diagnosis,
		Date:         consultation.Date.Format("2006-01-02"),
	}, nil
}

func (h *consultationHandler) UpdateConsultation(ctx context.Context, req *consultationPb.ConsultationRequest) (*consultationPb.ConsultationResponse, error) {
	if req.Id == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Consultation ID is required"))
	}

	consultation, err := h.app.FindByIDConsultation(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	md, _ := grpcmetadata.GetMetadataFromContext(ctx)

	userSnapshot := map[string]interface{}{
		"ID":    "",
		"Name":  "",
		"Email": "",
		"Role":  "",
	}

	if v, ok := md["ts-user-id"]; ok && len(v) > 0 {
		userSnapshot["ID"] = v[0]
	}
	if v, ok := md["ts-user-name"]; ok && len(v) > 0 {
		userSnapshot["Name"] = v[0]
	}
	if v, ok := md["ts-user-email"]; ok && len(v) > 0 {
		userSnapshot["Email"] = v[0]
	}
	if v, ok := md["ts-user-role"]; ok && len(v) > 0 {
		userSnapshot["Role"] = v[0]
	}

	updatedBySnapshot := domain.UpdatedBySnapshot{
		ID:    userSnapshot["ID"].(string),
		Name:  userSnapshot["Name"].(string),
		Email: userSnapshot["Email"].(string),
		Role:  userSnapshot["Role"].(string),
	}

	consultation.Symptoms = req.Symptoms
	consultation.Diagnosis = req.Diagnosis
	consultation.UpdatedBy = updatedBySnapshot
	consultation.UpdatedAt = time.Now()

	err = h.app.UpdateConsultation(ctx, consultation)
	if err != nil {
		return nil, err
	}

	var resQueID *wrapperspb.StringValue
	if consultation.QueueID != "" {
		resQueID = wrapperspb.String(consultation.QueueID)
	}

	var resAppID *wrapperspb.StringValue
	if consultation.AppointmentID != "" {
		resAppID = wrapperspb.String(consultation.AppointmentID)
	}

	return &consultationPb.ConsultationResponse{
		Id:            consultation.ID,
		QueueId:       resQueID,
		AppointmentId: resAppID,
	}, nil
}
