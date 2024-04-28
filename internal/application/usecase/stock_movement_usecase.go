package usecase

import (
	"context"
	"log"

	"github.com/dangquyitt/go-event-driven-architecture/internal/domain/event"
)

type StockMovementUseCase struct {
}

func NewStockMovementUseCase() *StockMovementUseCase {
	return &StockMovementUseCase{}
}

func (h *StockMovementUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	log.Println("--- StockMovementUseCase ---")
	for _, item := range payload.Items {
		log.Printf("Removing %d items of product %s from stock\n", item.Quantity, item.ProductName)
	}
	return nil
}
