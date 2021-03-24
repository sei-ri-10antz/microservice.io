package account

import (
	"context"
	"fmt"

	"github.com/sei-ri/microservice.io/api/v1/events"
	"github.com/sei-ri/microservice.io/api/v1/services"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
)

const (
	ErrAccountNotFound    = Error("account not found")
	ErrEmailAlreadyExists = Error("email already exists")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type AccountAggregate struct {
	pigeon.AggregateMixin

	Email    string
	Password string
}

func (a *AccountAggregate) ApplyChange(event pigeon.Event) error {
	switch e := event.Data.(type) {
	case *events.AccountCreated:
		a.ID = e.Id
		a.Email = e.Email
		a.Password = e.Password
	case *events.AccountPasswordChanged:
		a.ID = e.Id
		a.Password = e.Password
	default:
		fmt.Errorf("not support event: %v", e)
	}
	return nil
}

func (a *AccountAggregate) Handle(ctx context.Context, command pigeon.Command) error {
	var event pigeon.Event
	switch c := command.(type) {
	case *services.CreateAccountRequest:
		event.Data = &events.AccountCreated{
			Id:       c.Id.Value,
			Email:    c.Email,
			Password: c.Password,
		}
	case *services.ChangePasswordRequest:
		event.Data = &events.AccountPasswordChanged{
			Id:       c.Id,
			Password: c.Password,
		}
	default:
		return fmt.Errorf("not support command: %v", c)
	}
	return a.Apply(a, event, true)
}
