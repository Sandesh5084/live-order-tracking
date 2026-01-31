package service
import (
	"context"
	domain "mylotapp/internal/order/model"
	"mylotapp/internal/order/repository"
	"github.com/google/uuid"
)

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, customerID uuid.UUID) (*domain.Order, error) {
	order := domain.NewOrder(customerID)
	err := s.repo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}
func (s *orderService) GetOrder(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
    return s.repo.FindByID(ctx, id)
}

func (s *orderService) MarkOrderDelivered(ctx context.Context, id uuid.UUID, status string) error {
    return s.repo.UpdateStatus(ctx, id, status)
}

