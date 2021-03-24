package pigeon

import (
	"log"
	"reflect"

	"github.com/sei-ri/microservice.io/pkg/pigeon/internal"
)

type Broker interface {
	Publish(event Event) error
	Subscribe(reflect.Type, EventProcessor) error
	Close() error
}

type broker struct {
	events       map[string][]EventProcessor
	debugEnabled bool
}

func newBroker(debug bool) Broker {
	return &broker{
		events:       map[string][]EventProcessor{},
		debugEnabled: debug,
	}
}

func (b *broker) Publish(event Event) error {
	k := internal.ParseType(event.Data).String()

	if b.debugEnabled {
		log.Println("[PIGEON] publish:", k)
	}

	processors, ok := b.events[k]
	if !ok {
		return ErrNoSuchEvent
	}
	for i := range processors {
		go func(processor EventProcessor) {
			processor(event)
		}(processors[i])
	}
	return nil
}

func (b *broker) Subscribe(typ reflect.Type, processor EventProcessor) error {
	k := typ.String()

	if b.debugEnabled {
		log.Println("[PIGEON] subscribe:", k)
	}

	processors, ok := b.events[k]
	if !ok {
		processors = make([]EventProcessor, 0)
	}

	for i := range processors {
		if reflect.DeepEqual(processors[i], processor) {
			return ErrEventDuplicated
		}
	}
	processors = append(processors, processor)
	b.events[k] = processors
	return nil
}

func (b *broker) Close() error {
	b.events = make(map[string][]EventProcessor)
	return nil
}
