package pigeon

import (
	"context"
	"sync"
)

type Filter struct {
	ID      *string
	Version *int
}

type Data struct {
	ID      string
	Version int
	Type    string
	Dump    []byte
	Elem    interface{}
}

type Storage interface {
	Put(ctx context.Context, data ...*Data) error
	All(ctx context.Context, filter Filter) ([]Data, error)
	Get(ctx context.Context, filter Filter) (*Data, error)
	Close() error
}

type storage struct {
	mu    *sync.RWMutex
	all   []Data
	items map[string]Data
}

func newStorage() Storage {
	return &storage{
		mu:  &sync.RWMutex{},
		all: []Data{},
	}
}

func (s *storage) Put(ctx context.Context, data ...*Data) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range data {
		s.all = append(s.all, *data[i])
	}

	return nil
}

func (s *storage) All(ctx context.Context, filter Filter) ([]Data, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.all, nil
}

func (s *storage) Get(ctx context.Context, filter Filter) (*Data, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var row Data

	for i := range s.all {
		if v := filter.ID; v != nil {
			if s.all[i].ID == *v {
				row = s.all[i]
				break
			}
		}
		if v := filter.Version; v != nil {
			if s.all[i].Version == *v {
				row = s.all[i]
				break
			}
		}
	}

	return &row, nil
}

func (s *storage) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.all = make([]Data, 0)
	s.items = make(map[string]Data)

	return nil
}
