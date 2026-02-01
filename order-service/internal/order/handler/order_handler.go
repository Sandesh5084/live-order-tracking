package handler

import (
	"encoding/json"
	"net/http"
	"mylotapp/internal/order/service"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	domain "mylotapp/internal/order/model"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req domain.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	order, err := h.service.CreateOrder(r.Context(), req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := domain.OrderResponse{
		ID:         order.ID,
		UserID: order.UserID,
		Status:     string(order.Status),
		CreatedAt:  order.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	order, err := h.service.GetOrder(r.Context(), uuid.MustParse(id))
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	resp := domain.OrderResponse{
		ID:         order.ID,
		UserID:     order.UserID,
		Status:     string(order.Status),
		CreatedAt:  order.CreatedAt.Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *OrderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	statusDelivered := "delivered"

	err := h.service.UpdateStatus(r.Context(), uuid.MustParse(id), statusDelivered)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

