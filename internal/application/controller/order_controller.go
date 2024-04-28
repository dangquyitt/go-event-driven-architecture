package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dangquyitt/go-event-driven-architecture/internal/application/dto"
	"github.com/dangquyitt/go-event-driven-architecture/internal/application/usecase"
	"github.com/dangquyitt/go-event-driven-architecture/internal/domain/event"
)

type OrderController struct {
	createOrderUserCase        *usecase.CreateOrderUseCase
	processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase
	stockMovementUseCase       *usecase.StockMovementUseCase
	sendOrderEmailUseCase      *usecase.SendOrderEmailUseCase
}

func NewOrderController(createOrderUserCase *usecase.CreateOrderUseCase,
	processOrderPaymentUseCase *usecase.ProcessOrderPaymentUseCase,
	stockMovementUseCase *usecase.StockMovementUseCase,
	sendOrderEmailUseCase *usecase.SendOrderEmailUseCase) *OrderController {
	return &OrderController{
		createOrderUserCase:        createOrderUserCase,
		processOrderPaymentUseCase: processOrderPaymentUseCase,
		stockMovementUseCase:       stockMovementUseCase,
		sendOrderEmailUseCase:      sendOrderEmailUseCase,
	}
}

func (u *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}
	var requestData dto.CreateOrderDTO
	json.NewDecoder(r.Body).Decode(&requestData)
	err := u.createOrderUserCase.Execute(r.Context(), requestData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *OrderController) ProcessOrderPayment(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := u.processOrderPaymentUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *OrderController) StockMovement(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := u.stockMovementUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (u *OrderController) SendOrderEmail(w http.ResponseWriter, r *http.Request) {
	var body event.OrderCreatedEvent
	json.NewDecoder(r.Body).Decode(&body)
	err := u.sendOrderEmailUseCase.Execute(r.Context(), &body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
