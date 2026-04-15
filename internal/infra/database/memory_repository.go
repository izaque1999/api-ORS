package database

import (
	"errors"
	"sync"

	"github.com/izaque1999/api-ORS/internal/domain"
)

type MemoryRepository struct {
	data map[string]*domain.Order
	mu   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]*domain.Order),
	}
}

func (r *MemoryRepository) FindByID(id string) (*domain.Order, error) {
	r.mu.RLock()

	order, ok := r.data[id]
	if !ok {
		return nil, errors.New("pedido não encontrado")
	}

	return order, nil
}

func (r *MemoryRepository) FindAll() ([]*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var orders []*domain.Order
	for _, o := range r.data {
		orders = append(orders, o)
	}

	return orders, nil
}
