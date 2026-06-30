package ports

import (
	"context"

	"github.com/leonardoaraujodf/e-commerce/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
