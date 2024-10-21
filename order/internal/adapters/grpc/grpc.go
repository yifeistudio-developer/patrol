package grpc

import (
	"context"
	"github.com/yifeistudio-developer/wharf/golang/order"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

	return &order.CreateOrderResponse{
		OrderId: -1,
	}, nil
}
