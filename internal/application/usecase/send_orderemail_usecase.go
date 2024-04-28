package usecase

import (
	"context"
	"log"

	"github.com/dangquyitt/go-event-driven-architecture/internal/domain/event"
)

type SendOrderEmailUseCase struct {
}

func NewSendOrderEmailUseCase() *SendOrderEmailUseCase {
	return &SendOrderEmailUseCase{}
}

func (h *SendOrderEmailUseCase) Execute(ctx context.Context, payload *event.OrderCreatedEvent) error {
	log.Println("--- SendOrderEmailUseCase ---")
	log.Printf("--- MAIL Order Created: R$ %f \n", payload.TotalPrice)
	return nil
}
