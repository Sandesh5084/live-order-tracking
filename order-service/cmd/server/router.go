package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"mylotapp/internal/order/handler"
)

func setupRouter(
	orderHandler *handler.OrderHandler,
) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", orderHandler.CreateOrder)
		r.Get("/{id}", orderHandler.GetOrder)
		r.Put("/{id}/deliver", orderHandler.MarkDelivered)
	})

	return r
}
