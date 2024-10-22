package payment

import (
	"context"
	"fmt"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"github.com/yifeistudio-developer/wharf/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	//defer func(conn *grpc.ClientConn) {
	//	err := conn.Close()
	//	if err != nil {
	//		// handle close error.
	//	}
	//}(conn)
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
