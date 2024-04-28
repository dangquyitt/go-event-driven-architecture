package main

import (
	"log"
	"net/http"

	"github.com/dangquyitt/go-event-driven-architecture/internal/application/controller"
	"github.com/dangquyitt/go-event-driven-architecture/internal/application/usecase"
	"github.com/dangquyitt/go-event-driven-architecture/internal/infra/queue"
)

func main() {
	// initialize dependencies and implementations
	queue := queue.NewMemoryQueueAdapter()
	createOrderUseCase := usecase.NewCreateOrderUseCase(queue)
	orderController := controller.NewOrderController(createOrderUseCase)

	// ping
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// register routes
	http.HandleFunc("/create-order", orderController.CreateOrder)

	// start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
