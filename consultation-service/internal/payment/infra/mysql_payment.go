package infra

import (
	"context"
	"database/sql"
	"ehSehat/consultation-service/internal/payment/domain"
	"encoding/json"

	"github.com/google/uuid"
)

type mysqlPayment struct {
	db *sql.DB
}

func NewMySQLPayment(db *sql.DB) *mysqlPayment {
	return &mysqlPayment{db: db}
}

func (m *mysqlPayment) Create(ctx context.Context, payment *domain.Payment) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, `
		INSERT INTO payments 
		(id, consultation_id, consultation_date, patient_id, patient_name, doctor_id, doctor_name, amount, method, gateway, status, created_by, created_name, created_email, created_role, updated_by, updated_name, updated_email, updated_role) 
		VALUES 
		(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, payment.ID, payment.ConsultationID, payment.ConsultationDate, payment.PatientID, payment.PatientName, payment.DoctorID, payment.DoctorName, payment.Amount, payment.Method, payment.Gateway, "pending", payment.CreatedBy, payment.CreatedName, payment.CreatedEmail, payment.CreatedRole, payment.UpdatedBy, payment.UpdatedName, payment.UpdatedEmail, payment.UpdatedRole)
	if err != nil {
		tx.Rollback()
		return err
	}

	num, _ := result.RowsAffected()
	if num == 0 {
		tx.Rollback()
		return sql.ErrNoRows
	}

	// Insert payment logs
	for _, log := range payment.PaymentLogs {
		uuidLog, err := uuid.NewUUID()
		if err != nil {
			tx.Rollback()
			return err
		}
		log.ID = uuidLog.String()
		jsonResponse, err := json.Marshal(log.Response)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.ExecContext(ctx, `
					   INSERT INTO payment_logs (id, payment_id, response)
					   VALUES (?, ?, ?)`, log.ID, payment.ID, string(jsonResponse))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (m *mysqlPayment) Update(ctx context.Context, id string, payment *domain.UpdatePaymentRequest) error {
	_, err := m.db.ExecContext(ctx, "UPDATE payments SET status = ?, updated_by = ?, updated_name = ?, updated_email = ?, updated_role = ? WHERE id = ?", payment.Status, payment.UpdatedBy, payment.UpdatedName, payment.UpdatedEmail, payment.UpdatedRole, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlPayment) FindByID(ctx context.Context, id string) (*domain.Payment, error) {
	row := m.db.QueryRow("SELECT id, consultation_id, consultation_date, patient_id, patient_name, doctor_id, doctor_name, amount, method, gateway, status, created_by, created_name, created_email, created_role, updated_by, updated_name, updated_email, updated_role FROM payments WHERE id = ?", id)

	payment := &domain.Payment{}
	err := row.Scan(&payment.ID, &payment.ConsultationID, &payment.ConsultationDate, &payment.PatientID, &payment.PatientName, &payment.DoctorID, &payment.DoctorName, &payment.Amount, &payment.Method, &payment.Gateway, &payment.Status, &payment.CreatedBy, &payment.CreatedName, &payment.CreatedEmail, &payment.CreatedRole, &payment.UpdatedBy, &payment.UpdatedName, &payment.UpdatedEmail, &payment.UpdatedRole)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return payment, nil
}

func (m *mysqlPayment) HandlePaymentWebhook(ctx context.Context, webhook *domain.PaymentWebhook) error {
	payloadJSON, err := json.Marshal(webhook.Payload)
	if err != nil {
		return err
	}

	_, err = m.db.ExecContext(ctx, `
		INSERT INTO payment_webhooks (id, external_id, payment_id, event_type, payload)
		VALUES (?, ?, ?, ?, ?)`, webhook.ID, webhook.ExternalID, webhook.PaymentID, webhook.EventType, string(payloadJSON))

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlPayment) FindWebhookByExternalID(ctx context.Context, externalID string) (*domain.PaymentWebhook, error) {
	row := m.db.QueryRow("SELECT id, external_id, payment_id, event_type, payload FROM payment_webhooks WHERE external_id = ?", externalID)

	webhook := &domain.PaymentWebhook{}
	var payload string
	err := row.Scan(&webhook.ID, &webhook.ExternalID, &webhook.PaymentID, &webhook.EventType, &payload)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	webhook.Payload = payload
	return webhook, nil
}
