package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/sei-ri/microservice.io/pkg/pigeon"
	"github.com/sei-ri/microservice.io/pkg/pigeon/internal"
)

const (
	DefaultClusterID = "test-cluster"
	DefaultClientID  = "test-client"
)

type Broker struct {
	conn      stan.Conn
	opts      []stan.Option
	clusterID string
	clientID  string
}

type Option func(b *Broker) error

func WithClusterID(v string) Option {
	return func(b *Broker) error {
		b.clusterID = v
		return nil
	}
}

func WithClientID(v string) Option {
	return func(b *Broker) error {
		b.clientID = v
		return nil
	}
}

func WithNatsURL(v string) Option {
	return func(b *Broker) error {
		b.opts = append(b.opts, stan.NatsURL(v))
		return nil
	}
}

func WithNatsConn(nc *nats.Conn) Option {
	return func(b *Broker) error {
		b.opts = append(b.opts, stan.NatsConn(nc))
		return nil
	}
}

func New(ctx context.Context, opts ...Option) (*Broker, error) {
	b := &Broker{
		clusterID: DefaultClusterID,
		clientID:  DefaultClientID,
		opts: []stan.Option{
			stan.NatsURL(nats.DefaultURL),
			stan.SetConnectionLostHandler(func(_ stan.Conn, err error) {
				log.Fatalf("Connection lost, reason: %v", err)
			}),
		},
	}

	for i := range opts {
		if err := opts[i](b); err != nil {
			return nil, err
		}
	}

	return b, b.open(ctx)
}

func (b *Broker) open(ctx context.Context) error {
	sc, err := stan.Connect(b.clusterID, b.clientID, b.opts...)
	if err != nil {
		return err
	}
	b.conn = sc
	log.Printf("[PIGEON] nats-streaming connected to %s clusterID: [%s] clientID: [%s]\n",
		sc.NatsConn().ConnectedUrl(),
		b.clusterID,
		b.clientID,
	)
	return nil
}

func (b *Broker) Publish(event pigeon.Event) error {
	// TODO: need publisher options (for dynamic setting options)
	raw, err := json.Marshal(event.Data)
	if err != nil {
		return err
	}
	subj := internal.ParseType(event.Data).String()

	var guid string

	guid, err = b.conn.PublishAsync(subj, raw, func(ackedGuid string, err error) {
		if err != nil {
			log.Fatalf("Error in server ack for guid %s: %v\n", ackedGuid, err)
		}
		if ackedGuid != guid {
			log.Fatalf("Expected a matching guid in ack callback, got %s vs %s\n", ackedGuid, guid)
		}
	})
	log.Printf("[PIGEON] broker guid: %v", guid)
	return err
}

func (b *Broker) Subscribe(typ reflect.Type, processor pigeon.EventProcessor) error {
	// TODO: need subscriber options (for dynamic setting options)
	_, err := b.conn.QueueSubscribe(typ.String(), fmt.Sprintf("%s-group", b.clientID), func(msg *stan.Msg) {
		elem := reflect.New(typ).Interface()
		if err := json.Unmarshal(msg.Data, elem); err != nil {
			log.Printf("[PIGEON] json unmarshal: %v", err)
		}
		if err := processor(pigeon.Event{Data: elem}); err != nil {
			log.Printf("[PIGEON] failed to event processor: %v", err)
		}
	},
		stan.DurableName(fmt.Sprintf("%s-durable", b.clientID)),
		stan.StartWithLastReceived(),
	)
	return err
}

func (b *Broker) Close() error {
	if b.conn != nil {
		return b.conn.Close()
	}
	return nil
}
