package memory

import (
	"errors"
	"sync"
)

type MemoryRepository[T any] struct {
	items []T
	mu    sync.RWMutex
	getID func(T) string
}

func NewMemoryRepository[T any](getID func(T) string, initialItems ...T) *MemoryRepository[T] {
	return &MemoryRepository[T]{
		items: initialItems,
		getID: getID,
	}
}

func (r *MemoryRepository[T]) GetAll() ([]T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.items, nil
}

func (r *MemoryRepository[T]) GetByID(id string) (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, item := range r.items {
		if r.getID(item) == id {
			return item, nil
		}
	}

	var zero T
	return zero, errors.New("not found")
}

func (r *MemoryRepository[T]) Create(entity T) (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items = append(r.items, entity)
	return entity, nil
}

func (r *MemoryRepository[T]) Update(entity T) (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	entityID := r.getID(entity)
	for i, item := range r.items {
		if r.getID(item) == entityID {
			r.items[i] = entity
			return entity, nil
		}
	}

	var zero T
	return zero, errors.New("not found")
}

func (r *MemoryRepository[T]) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, item := range r.items {
		if r.getID(item) == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}

	return errors.New("not found")
}
