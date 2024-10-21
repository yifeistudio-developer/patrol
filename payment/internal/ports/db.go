package ports

import "github.com/yifeistudio-developer/patrol/payment/internal/application/core/domain"

type DBPort interface {
	Get(id int64) (domain.Payment, error)
	Save(order *domain.Payment) error
}
