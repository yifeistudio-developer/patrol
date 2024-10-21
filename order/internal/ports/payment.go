package ports

import "github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
