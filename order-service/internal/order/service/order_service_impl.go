package service
import (
	"context"
	domain "mylotapp/internal/order/model"
	"mylotapp/internal/order/repository"
	"mylotapp/internal/event"
	"github.com/google/uuid"
	"time"
)

type orderService struct {
	repo repository.OrderRepository
	publisher event.Publisher
}

func NewOrderService(repo repository.OrderRepository, publisher event.Publisher) OrderService {
	return &orderService{
		repo: repo,
		publisher: publisher,
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

func (s *orderService) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
    err := s.repo.UpdateStatus(ctx, id, status)
    if err != nil {
        return err
    }

    event := event.OrderStatusUpdatedEvent{
        EventType: "order_status_updated",
        OrderID:   id.String(),
        Status:    status,
        Timestamp: time.Now(),
    }

    err = s.publisher.Publish(event)
    if err != nil {
        return err
    }

    return nil
}

