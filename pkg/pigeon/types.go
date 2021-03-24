package pigeon

import (
	"reflect"
	"sync"

	"github.com/sei-ri/microservice.io/pkg/pigeon/internal"
)

// Types is event type memroy store
type Types interface {
	Put(i interface{})
	Get(s string) reflect.Type
}

// Types is event type memroy store
type types struct {
	mu    *sync.Mutex
	items map[string]reflect.Type
}

func newTypes() Types {
	return &types{
		mu:    &sync.Mutex{},
		items: map[string]reflect.Type{}}
}

func (a *types) Put(i interface{}) {
	a.mu.Lock()
	defer a.mu.Unlock()

	t := internal.ParseType(i)
	a.items[t.String()] = t
}

func (a *types) Get(s string) reflect.Type {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.items[s]
}
