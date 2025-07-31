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

type PaymentService interface {
	CreatePayment(ctx context.Context, payment *CreatePaymentRequest) (*Payment, error)
	UpdatePayment(ctx context.Context, id string, payment *UpdatePaymentRequest) error
	FindByIDPayment(ctx context.Context, id string) (*Payment, error)
	HandlePaymentWebhook(ctx context.Context, webhook *PaymentWebhook) error
}
