package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"testing"
)

type mockedPayment struct {
	mock.Mock
}

func (p *mockedPayment) Charge(order *domain.Order) error {
	args := p.Called(order)
	return args.Error(0)
}

type mockedDb struct {
	mock.Mock
}

func (db *mockedDb) Save(order *domain.Order) error {
	return db.Called(order).Error(0)
}

func (db *mockedDb) Get(id int64) (domain.Order, error) {
	args := db.Called(id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func Test_Should_Place_Order(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(nil)
	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(domain.Order{})
	assert.Nil(t, err)
}
