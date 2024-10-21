package api

import (
	"github.com/yifeistudio-developer/patrol/payment/internal/application/core/domain"
	"github.com/yifeistudio-developer/patrol/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(payment domain.Payment) (domain.Payment, error) {

	return domain.Payment{}, nil
}
