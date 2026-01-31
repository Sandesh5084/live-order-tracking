package service

import "context"
import domain "mylotapp/internal/order/model"
import "github.com/google/uuid"

type OrderService interface {
	CreateOrder(ctx context.Context, customerID uuid.UUID) (*domain.Order, error)
	GetOrder(ctx context.Context, id uuid.UUID) (*domain.Order, error)
	MarkOrderDelivered(ctx context.Context, id uuid.UUID, status string) error
}
