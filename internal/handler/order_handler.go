package handler

import (
	"encoding/json"
	"net/http"

	"github.com/izaque1999/api-ORS/internal/domain"
	"github.com/izaque1999/api-ORS/internal/usecase"
)

type OrderHandler struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CustomerName string        `json:"customer_name"`
		Items        []domain.Item `json:"items"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.CreateOrderUseCase.Execute(req.CustomerName, req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
