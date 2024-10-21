package db

import (
	"fmt"
	"github.com/yifeistudio-developer/patrol/payment/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func (a *Adapter) Get(id int64) (domain.Payment, error) {
	return domain.Payment{}, nil
}

func (a *Adapter) Save(payment *domain.Payment) error {
	return nil
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", openErr)
	}
	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate order: %w", err)
	}
	return &Adapter{db}, nil
}

type Payment struct {
	gorm.Model
	Status     string
	UserId     int64
	TotalPrice float64
	OrderId    int64
}
