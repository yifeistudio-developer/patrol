package payment

import (
	"context"
	"fmt"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"github.com/yifeistudio-developer/wharf/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	interceptor := grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(500*time.Millisecond))))
	opts = append(opts, interceptor)
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a Adapter) Charge(order *domain.Order) error {
	ctx := context.Background()
	result, err := a.payment.Create(ctx, &payment.CreatePaymentRequest{
		OrderId: 1,
	})
	if err != nil {
		// handle error
		fmt.Printf("handle charge error: %v\n", err)
		return err
	}
	fmt.Println(result)
	return nil
}
