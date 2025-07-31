package domain

import "time"

type Payment struct {
	ID               string       `json:"id"`
	ConsultationID   string       `json:"consultation_id"`
	ConsultationDate *time.Time   `json:"consultation_date"`
	PatientID        string       `json:"patient_id"`
	PatientName      *string      `json:"patient_name"`
	DoctorID         string       `json:"doctor_id"`
	DoctorName       *string      `json:"doctor_name"`
	Amount           float64      `json:"amount"`
	Method           string       `json:"method"`
	Gateway          string       `json:"gateway"`
	Identifier       string       `json:"identifier"`
	Status           string       `json:"status"`
	CreatedBy        string       `json:"created_by"`
	CreatedName      *string      `json:"created_name"`
	CreatedEmail     *string      `json:"created_email"`
	CreatedRole      *string      `json:"created_role"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedBy        string       `json:"updated_by"`
	UpdatedName      *string      `json:"updated_name"`
	UpdatedEmail     *string      `json:"updated_email"`
	UpdatedRole      *string      `json:"updated_role"`
	UpdatedAt        time.Time    `json:"updated_at"`
	PaymentLogs      []PaymentLog `json:"payment_logs,omitempty"`
}

type PaymentLog struct {
	ID        string      `json:"id"`
	PaymentID string      `json:"payment_id"`
	Response  interface{} `json:"response"`
}

type CreatePaymentRequest struct {
	ConsultationID string  `json:"consultation_id"`
	Amount         float64 `json:"amount"`
	Method         string  `json:"method"`
}

type UpdatePaymentRequest struct {
	Amount       float64   `json:"amount"`
	Method       string    `json:"method"`
	Status       string    `json:"status"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedName  *string   `json:"updated_name"`
	UpdatedEmail *string   `json:"updated_email"`
	UpdatedRole  *string   `json:"updated_role"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// webhook := &domain.PaymentWebhook{
// 		ID:         req.Id,
// 		ExternalID: req.ExternalId,
// 		PaymentID:  req.PaymentId,
// 		EventType:  req.EventType,
// 		Payload:    req.Payload.GetValue(),
// 	}

// 	err := h.app.HandlePaymentWebhook(ctx, webhook)
// 	if err != nil {
// 		return nil, utils.GRPCErrorToHTTPError(err)
// 	}
type PaymentWebhook struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	PaymentID  string `json:"payment_id"`
	EventType  string `json:"event_type"`
	Payload    string `json:"payload"`
}
