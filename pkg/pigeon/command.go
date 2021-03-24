package pigeon

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"reflect"

	"github.com/sei-ri/microservice.io/pkg/pigeon/internal"
)

type Command interface{}

type CommandHandler interface {
	Handle(context.Context, Command) error
}

type simpleCommandHandler struct {
	aggregate reflect.Type
	types     Types
	storage   Storage
	broker    Broker
}

func NewSimpleCommandHandler(aggregate Aggregate, client *Client) CommandHandler {
	return &simpleCommandHandler{
		aggregate: reflect.TypeOf(aggregate).Elem(),
		types:     client.types,
		storage:   client.storage,
		broker:    client.broker,
	}
}

func (h *simpleCommandHandler) Handle(ctx context.Context, cmd Command) error {
	aggregate := reflect.New(h.aggregate).Interface().(Aggregate)

	if err := aggregate.Handle(ctx, cmd); err != nil {
		return err
	}

	aggregateID := aggregate.AggregateID()

	var version int

	// Gets last version by aggregateID
	if v, err := h.storage.Get(ctx, Filter{
		ID: &aggregateID,
	}); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	} else if v != nil && v.Version > 0 {
		version = v.Version
	}

	data := make([]*Data, len(aggregate.Uncommitted()))
	for i, event := range aggregate.Uncommitted() {
		version++
		event.version = version

		if v, err := json.Marshal(event.Data); err != nil {
			return err
		} else {
			data[i] = &Data{
				ID:      aggregateID,
				Version: event.version,
				Type:    internal.ParseType(event.Data).String(),
				Dump:    v,
			}
		}

		if err := h.broker.Publish(event); err != nil {
			return err
		}
	}

	if err := h.storage.Put(ctx, data...); err != nil {
		return err
	}

	return nil
}
