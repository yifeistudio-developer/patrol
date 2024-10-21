package api

import (
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"github.com/yifeistudio-developer/patrol/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	err = a.payment.Charge(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return domain.Order{}, nil
}
