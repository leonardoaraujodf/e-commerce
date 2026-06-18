package ports

import "github.com/leonardoaraujodf/e-commerce/order/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
