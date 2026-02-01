package event

import "time"

type OrderStatusUpdatedEvent struct {
	EventType string    `json:"type"`
	OrderID   string    `json:"order_id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
