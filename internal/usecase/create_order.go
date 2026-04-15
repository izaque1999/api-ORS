package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/izaque1999/api-ORS/internal/domain"
	"github.com/izaque1999/api-ORS/internal/repository"
)

type CreateOrderUseCase struct {
	Repo repository.OrderRepository
}

func (uc *CreateOrderUseCase) Execute(customerName string, items []domain.Item) (*domain.Order, error) {
	order := &domain.Order{
		ID:           uuid.NewString(),
		CustomerName: customerName,
		Items:        items,
		Status:       domain.StatusCreated,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	order.CalculateTotal()

	err := uc.Repo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil

}
