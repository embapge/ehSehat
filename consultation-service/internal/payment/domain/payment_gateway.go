package domain

type PaymentGateway interface {
	Create(external_id string, amount float64) (interface{}, error)
	GetGatewayName() string
	// Update(external_id string, amount float64) (interface{}, error)
	// Expire(external_id string) (interface{}, error)
	// FindByID(external_id string) (interface{}, error)
}
