package pigeon

import (
	"context"
	"log"

	"github.com/sei-ri/microservice.io/pkg/pigeon/internal"
)

type Client struct {
	commands map[string]CommandHandler
	events   map[string][]EventProcessor
	broker   Broker
	storage  Storage
	types    Types
	develop  bool
}

type Option func(c *Client) error

func WithBroker(v Broker) Option {
	return func(c *Client) error {
		c.broker = v
		return nil
	}
}

func WithStorage(v Storage) Option {
	return func(c *Client) error {
		c.storage = v
		return nil
	}
}

func WithDevelop(v bool) Option {
	return func(c *Client) error {
		c.develop = v
		return nil
	}
}

func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	c := &Client{
		commands: map[string]CommandHandler{},
		events:   map[string][]EventProcessor{},
		broker:   newBroker(true),
		storage:  newStorage(),
		types:    newTypes(),
		develop:  true,
	}

	for i := range opts {
		if err := opts[i](c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) AddCommands(handler CommandHandler, args ...Command) error {
	for i := range args {
		k := internal.ParseType(args[i]).String()

		if _, ok := c.commands[k]; ok {
			return ErrCommandDuplicated
		}

		c.commands[k] = handler

		if c.develop {
			log.Println("[PIGEON] command:", k)
		}
	}
	return nil
}

func (c *Client) AddEvents(typ interface{}, processor EventProcessor) error {
	c.types.Put(typ)
	k := internal.ParseType(typ)
	if err := c.broker.Subscribe(k, processor); err != nil {
		return err
	}
	return nil
}

func (c *Client) Dispatch(ctx context.Context, command Command) error {
	k := internal.ParseType(command).String()
	if c.develop {
		log.Println("[PIGEON] dispatch command:", k)
	}
	handler, ok := c.commands[k]
	if !ok {
		return ErrNoSuchCommand
	}
	return handler.Handle(ctx, command)
}

func (c *Client) Close() error {
	if c.broker != nil {
		return c.broker.Close()
	}
	if c.storage != nil {
		return c.storage.Close()
	}
	return nil
}

func (c *Client) Types() Types {
	return c.types
}

func (c *Client) Broker() Broker {
	return c.broker
}

func (c *Client) Storage() Storage {
	return c.storage
}
