package domain

import "context"

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

// type PaymentService interface {
// 	CreatePayment(ctx context.Context, payment *CreatePaymentRequest) (*Payment, error)
// 	UpdatePayment(ctx context.Context, id string, payment *UpdatePaymentRequest) error
// 	FindByIDPayment(ctx context.Context, id string) (*Payment, error)
// 	HandlePaymentWebhook(ctx context.Context, webhook *PaymentWebhook) error
// }

// func (app *paymentApp) HandlePaymentWebhook(ctx context.Context, webhook *domain.PaymentWebhook) error {
// 	if webhook == nil || webhook.ID == "" || webhook.ExternalID == "" || webhook.PaymentID == "" || webhook.EventType == "" {
// 		return utils.NewBadRequestError("Invalid webhook request")
// 	}

// 	err := app.paymentRepo.HandlePaymentWebhook(ctx, webhook)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, id string, payment *UpdatePaymentRequest) error
	FindByID(ctx context.Context, id string) (*Payment, error)
	HandlePaymentWebhook(ctx context.Context, webhook *PaymentWebhook) error
	// tambahkan  cari payment webhook berdasarkan external_id
	FindWebhookByExternalID(ctx context.Context, externalID string) (*PaymentWebhook, error)
}
