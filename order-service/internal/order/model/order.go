package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusPlaced         = "PLACED"
	StatusConfirmed      = "CONFIRMED"
	StatusPreparing      = "PREPARING"
	StatusOutForDelivery = "OUT_FOR_DELIVERY"
	StatusDelivered      = "DELIVERED"
)

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrder(userID uuid.UUID) *Order {
	now := time.Now()

	return &Order{
		ID:        uuid.New(),
		UserID:    userID,
		Status:    StatusPlaced,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func IsValidStatus(status string) bool {
	switch status {
	case StatusPlaced,
		StatusConfirmed,
		StatusPreparing,
		StatusOutForDelivery,
		StatusDelivered:
		return true
	default:
		return false
	}
}

type CreateOrderRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type OrderResponse struct {
	ID         uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}
