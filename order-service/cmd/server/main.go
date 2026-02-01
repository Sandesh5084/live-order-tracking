package main

import (
	"log"
	"net/http"
    "mylotapp/internal/infra/postgres"
    "mylotapp/internal/order/handler"
    "mylotapp/internal/order/service"
    "mylotapp/internal/event"
    "mylotapp/internal/config"
    
)

func main() {
	cfg := config.LoadConfig()
	db := postgres.MustConnectDB(cfg)

    publisher := event.NewNoopPublisher()

	orderRepo := postgres.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo, publisher)
	orderHandler := handler.NewOrderHandler(orderService)

	router := setupRouter(orderHandler)

	log.Println("server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
