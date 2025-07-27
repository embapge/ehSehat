package domain

import "context"

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, id string, payment *UpdatePaymentRequest) error
	FindByID(ctx context.Context, id string) (*Payment, error)
}
