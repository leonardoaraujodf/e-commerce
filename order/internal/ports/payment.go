package ports

import "github.com/leonardoaraujodf/e-commerce/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
