package domain

import (
	"errors"
	"time"
)

type Status string

const (
	StatusCreated   Status = "CREATED"
	StatusPaid      Status = "PAID"
	StatusShipped   Status = "SHIPPED"
	StatusDelivered Status = "DELIVERED"
	StatusCanceled  Status = "CANCELED"
)

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	ID           string
	CustomerName string
	Items        []Item
	TotalAmount  float64
	Status       Status
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (o *Order) CalculateTotal() {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.TotalAmount = total

}

func (o *Order) Pay() error {
	if o.Status != StatusCreated {
		return errors.New("pedido não pode ser pago")
	}
	o.Status = StatusPaid
	return nil
}
