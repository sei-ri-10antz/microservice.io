package pigeon

type Event struct {
	version int
	Data    interface{}
}

type EventHandler interface {
	Handle(Event) error
}

type EventProcessor func(Event) error
