package product

import (
	"context"
	"fmt"

	"github.com/sei-ri/microservice.io/api/v1/events"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
)

var _ pigeon.Aggregate = &ProductAggregate{}

const (
	ErrProductNotFound      = Error("product not found")
	ErrProductQtyBalanceOut = Error("product qty balance out")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type ProductAggregate struct {
	pigeon.AggregateMixin

	Name     string
	ImageURL string
	Price    int64
	Qty      int64
}

func (a *ProductAggregate) ApplyChange(event pigeon.Event) error {
	switch e := event.Data.(type) {
	case *events.ProductCreated:
		a.ID = e.Id
		a.Name = e.Name
		a.ImageURL = e.ImageUrl
		a.Price = e.Price
		a.Qty = e.Qty
	case *events.ProductQtyDeducted:
		a.ID = e.Id
		a.Qty = e.Qty
	default:
		return fmt.Errorf("not support event: %v", e)
	}
	return nil
}

func (a *ProductAggregate) Handle(ctx context.Context, command pigeon.Command) error {
	var event pigeon.Event
	switch c := command.(type) {
	case *services.CreateProductRequest:
		event.Data = &events.ProductCreated{
			Id:       c.Id.Value,
			Name:     c.Name,
			ImageUrl: c.ImageUrl,
			Price:    c.Price,
			Qty:      c.Qty,
		}
	case *services.DeductProductQtyRequest:
		event.Data = &events.ProductQtyDeducted{
			Id:  c.Id,
			Qty: c.Qty,
		}
	default:
		return fmt.Errorf("not support command: %v", c)
	}
	return a.Apply(a, event, true)
}
