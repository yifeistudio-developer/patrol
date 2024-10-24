// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
)

// APIPort is an autogenerated mock type for the APIPort type
type APIPort struct {
	mock.Mock
}

// PlaceOrder provides a mock function with given fields: order
func (_m *APIPort) PlaceOrder(order domain.Order) (domain.Order, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for PlaceOrder")
	}

	var r0 domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Order) (domain.Order, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(domain.Order) domain.Order); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	if rf, ok := ret.Get(1).(func(domain.Order) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAPIPort creates a new instance of APIPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *APIPort {
	mock := &APIPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
