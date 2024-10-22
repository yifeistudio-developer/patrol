package ports

import "github.com/yifeistudio-developer/patrol/payment/internal/application/core/domain"

type APIPort interface {
	Charge(payment domain.Payment) (domain.Payment, error)
}
