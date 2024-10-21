package ports

import "github.com/yifeistudio-developer/patrol/payment/internal/application/core/domain"

type APIPort interface {
	Charge(order domain.Payment) (domain.Payment, error)
}
