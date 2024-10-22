package api

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// happy path test.
func Test_Should_Place_Order(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(nil)
	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(domain.Order{})
	assert.Nil(t, err)
}

// db failed path
func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(errors.New("connection error"))
	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(domain.Order{})
	assert.EqualError(t, err, "connection error")
}

func Test_Should_Return_Error_When_Payment_Fail(t *testing.T) {
	payment := new(mockedPayment)
	db := new(mockedDb)
	payment.On("Charge", mock.Anything).Return(errors.New("insufficient funds"))
	db.On("Save", mock.Anything).Return(nil)
	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(domain.Order{})
	st, _ := status.FromError(err)
	assert.Equal(t, st.Message(), "order action failed.")
	assert.Equal(t, st.Details()[0].(*errdetails.BadRequest).FieldViolations[0].Description, "insufficient funds")
	assert.Equal(t, st.Code(), codes.InvalidArgument)
}
