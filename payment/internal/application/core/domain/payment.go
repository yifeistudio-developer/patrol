package domain

import "time"

type Payment struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	OrderId    int64     `json:"order_id"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewPayment(userId int64, orderId int64, totalPrice float64) Payment {
	return Payment{
		CreatedAt:  time.Now(),
		Status:     "PENDING",
		UserId:     userId,
		TotalPrice: totalPrice,
		OrderId:    orderId,
	}
}
