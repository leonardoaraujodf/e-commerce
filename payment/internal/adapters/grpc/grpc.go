package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/leonardoaraujodf/e-commerce.proto/golang/payment"
	"github.com/leonardoaraujodf/e-commerce/payment/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	log.Println("New create request received!")
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v", err)).Err()
	}

	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}
