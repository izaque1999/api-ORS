package repository

import "github.com/izaque1999/api-ORS/internal/domain"

type OrderRepository interface {
	Save(order *domain.Order) error
	FindByID(id string) (*domain.Order, error)
	FindAll() ([]*domain.Order, error)
}
