package ports

import "github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"

type DBPort interface {
	Get(id int64) (domain.Order, error)
	Save(order *domain.Order) error
}
