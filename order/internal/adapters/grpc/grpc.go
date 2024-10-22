package grpc

import (
	"context"
	"fmt"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"github.com/yifeistudio-developer/wharf/golang/order"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	placeOrder, err := a.api.PlaceOrder(domain.Order{
		CustomerId: request.CustomerId,
		Status:     "PENDING",
		OrderItems: []domain.OrderItem{
			{
				ProductCode: "aaa",
				UnitPrice:   12.01,
				Quantity:    1,
			},
		},
	})
	if err != nil {
		fmt.Printf("handle place order error: %v", err)
		return nil, err
	}
	fmt.Println(placeOrder)
	return &order.CreateOrderResponse{
		OrderId: -1,
	}, nil
}
