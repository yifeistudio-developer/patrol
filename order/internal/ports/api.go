package ports

import "github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
