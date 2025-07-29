package domain

import "context"

type PaymentService interface {
	CreatePayment(ctx context.Context, payment *CreatePaymentRequest) (*Payment, error)
	UpdatePayment(ctx context.Context, id string, payment *UpdatePaymentRequest) error
	FindByIDPayment(ctx context.Context, id string) (*Payment, error)
}
