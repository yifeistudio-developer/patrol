package db

import (
	"fmt"
	"github.com/yifeistudio-developer/patrol/order/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func (a *Adapter) Get(id int64) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	order := domain.Order{
		Id:         int64(orderEntity.ID),
		CustomerId: orderEntity.CustomerId,
		Status:     orderEntity.Status,
		CreatedAt:  orderEntity.CreatedAt,
	}
	return order, res.Error
}

func (a *Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	orderModel := Order{
		CustomerId: order.CustomerId,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.Id = int64(orderModel.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", openErr)
	}
	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate order: %w", err)
	}
	return &Adapter{db}, nil
}

type Order struct {
	gorm.Model
	CustomerId int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderId     int64
}
