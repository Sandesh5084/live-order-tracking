package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	domain "mylotapp/internal/order/model"
	"mylotapp/internal/order/repository"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) repository.OrderRepository {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(ctx context.Context, o *domain.Order) error {
	query := `
		INSERT INTO orders (id, customer_id, status, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		o.ID,
		o.UserID,
		o.Status,
		o.CreatedAt,
	)

	return err
}

func (r *OrderRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
	query := `
		SELECT id, customer_id, status, created_at
		FROM orders
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var o domain.Order
	err := row.Scan(
		&o.ID,
		&o.UserID,
		&o.Status,
		&o.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r *OrderRepo) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	query := `
		UPDATE orders
		SET status = $1
		WHERE id = $2
	`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}
