package persistence

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// mockStorer simulates data for acceptance test and unit test
type mockStorer[K comparable, V any] struct {
	cache map[K]V
}

// NewMockStore returns a new instance of ports.Storer interface
func NewMockStorer[K comparable, V any]() ports.Storer[K, V] {
	return &mockStorer[K, V]{
		cache: make(map[K]V),
	}
}

// Save saves a resource in a map
func (m *mockStorer[K, V]) Save(_ context.Context, k K, v V) (err error) {
	_, ok := m.cache[k]

	if ok {
		err = fmt.Errorf("resource with id %v already existing", k)
		return
	}

	m.cache[k] = v

	return
}

// Remove removes the resources by an identifier of a map
func (m *mockStorer[K, V]) Remove(_ context.Context, k K) (err error) {
	_, ok := m.cache[k]

	if !ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("resource with id %v not found", k))
		return
	}

	delete(m.cache, k)

	return
}

// Search searchs a resource by id from a map
func (m *mockStorer[K, V]) Search(_ context.Context, k K) (v V, err error) {
	v, ok := m.cache[k]

	if !ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("resource with id %v not found", k))
		return
	}

	return
}
