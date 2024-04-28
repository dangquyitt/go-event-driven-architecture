package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dangquyitt/go-event-driven-architecture/internal/application/dto"
	"github.com/dangquyitt/go-event-driven-architecture/internal/application/usecase"
)

type OrderController struct {
	createOrderUserCase *usecase.CreateOrderUseCase
}

func NewOrderController(createOrderUserCase *usecase.CreateOrderUseCase) *OrderController {
	return &OrderController{
		createOrderUserCase: createOrderUserCase,
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