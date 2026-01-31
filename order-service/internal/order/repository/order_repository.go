package repository

import (
	"context"
	domain "mylotapp/internal/order/model"
	"github.com/google/uuid"
)


type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Order, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}
