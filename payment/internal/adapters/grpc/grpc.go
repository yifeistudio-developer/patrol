package grpc

import (
	"context"
	"fmt"
	"github.com/yifeistudio-developer/patrol/payment/internal/application/core/domain"
	"github.com/yifeistudio-developer/wharf/golang/payment"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	fmt.Println("handel create payment request.")
	result, err := a.api.Charge(domain.Payment{})
	if err != nil {
		return nil, err
	}
	return &payment.CreatePaymentResponse{
		BillId: result.Id,
	}, nil
}
