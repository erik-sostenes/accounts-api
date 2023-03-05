package persistence

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// mockStore simulates data for acceptance test and unit test
type mockStore[K comparable, V any] struct {
	cache map[K]V
}

// NewMockStore returns a new instance of ports.Store interface
func NewMockStore[K comparable, V any]() ports.Store[K, V] {
	return &mockStore[K, V]{
		cache: make(map[K]V),
	}
}

// Save saves a resource in a map
// if the resource already exist, returns a StatusBadRequest type error
func (m *mockStore[K, V]) Save(_ context.Context, k K, v V) (err error) {
	_, ok := m.cache[k]

	if ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("resource with id %v already existing", k))
		return
	}

	m.cache[k] = v

	return
}

// Remove removes the resources by an identifier of a map
// if the resource is not found, returns a Not Found type error
func (m *mockStore[K, V]) Remove(_ context.Context, k K) (err error) {
	_, ok := m.cache[k]

	if !ok {
		err = wrongs.StatusNotFound(fmt.Sprintf("resource with id %v not found", k))
		return
	}

	delete(m.cache, k)

	return
}

// Search searches a resource by id from a map
// if the resource is not found, returns a Not Found type error
func (m *mockStore[K, V]) Search(_ context.Context, k K) (v V, err error) {
	v, ok := m.cache[k]

	if !ok {
		err = wrongs.StatusNotFound(fmt.Sprintf("resource with id %v not found", k))
		return
	}

	return
}
