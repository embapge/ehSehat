package app

import (
	"context"
	"database/sql"
	consultationDomain "ehSehat/consultation-service/internal/consultation/domain"
	"ehSehat/consultation-service/internal/payment/domain"
	"ehSehat/libs/utils/grpcmetadata"
	"fmt"
)

type paymentApp struct {
	paymentRepo     domain.PaymentRepository
	consultationApp consultationDomain.ConsultationService
	pg              domain.PaymentGateway
}

func NewPaymentApp(repo domain.PaymentRepository, consultationApp consultationDomain.ConsultationService, pg domain.PaymentGateway) *paymentApp {
	return &paymentApp{
		paymentRepo:     repo,
		consultationApp: consultationApp,
		pg:              pg,
	}
}

func (app *paymentApp) CreatePayment(ctx context.Context, paymentReq *domain.CreatePaymentRequest) (*domain.Payment, error) {
	if app.consultationApp == nil {
		return nil, fmt.Errorf("consultationApp is nil")
	}

	consultation, _ := app.consultationApp.FindByIDConsultation(ctx, paymentReq.ConsultationID)
	if consultation.ID == "" {
		return nil, sql.ErrNoRows
	}

	result, err := app.pg.Create(consultation.ID, 600000.0)
	if err != nil {
		return nil, err
	}

	paymentLogs := domain.PaymentLog{
		Response: result,
	}

	md, _ := grpcmetadata.GetMetadataFromContext(ctx)

	userSnapshot := map[string]string{
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

	payment := &domain.Payment{
		ConsultationID:   consultation.ID,
		ConsultationDate: &consultation.Date,
		PatientID:        consultation.Patient.ID,
		PatientName:      &consultation.Patient.Name,
		DoctorID:         consultation.Doctor.ID,
		DoctorName:       &consultation.Doctor.Name,
		Amount:           paymentReq.Amount,
		Method:           "payment_link",
		Gateway:          app.pg.GetGatewayName(),
		PaymentLogs:      []domain.PaymentLog{paymentLogs},
		CreatedBy:        userSnapshot["ID"],
		CreatedName:      toPtr(userSnapshot["Name"]),
		CreatedEmail:     toPtr(userSnapshot["Email"]),
		CreatedRole:      toPtr(userSnapshot["Role"]),
		UpdatedBy:        userSnapshot["ID"],
		UpdatedName:      toPtr(userSnapshot["Name"]),
		UpdatedEmail:     toPtr(userSnapshot["Email"]),
		UpdatedRole:      toPtr(userSnapshot["Role"]),
	}

	err = app.paymentRepo.Create(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (app *paymentApp) UpdatePayment(ctx context.Context, id string, paymentReq *domain.UpdatePaymentRequest) error {
	payment, err := app.paymentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if payment.ID == "" {
		return sql.ErrNoRows
	}

	err = app.paymentRepo.Update(ctx, id, paymentReq)
	if err != nil {
		return err
	}

	return nil
}

func (app *paymentApp) FindByIDPayment(ctx context.Context, id string) (*domain.Payment, error) {
	payment, err := app.paymentRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if payment.ID == "" {
		return nil, sql.ErrNoRows
	}

	return payment, nil
}

func toPtr(s string) *string { return &s }
