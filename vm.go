package valuemap

import (
	"reflect"
	"sync"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type ValueMap[K comparable, V any] struct {
	_    noCopy
	mu   sync.RWMutex
	data map[K]V
}

// New returns a new pointer to a thread-safe ValueMap.
func New[K comparable, V any]() *ValueMap[K, V] {
	return &ValueMap[K, V]{data: make(map[K]V)}
}

// FromMap returns a new ValueMap initialized with a copy of an existing map.
func FromMap[K comparable, V any](m map[K]V) *ValueMap[K, V] {
	cp := make(map[K]V, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return &ValueMap[K, V]{data: cp}
}

// Set assigns a value to a key.
func (m *ValueMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// Get retrieves a value and a boolean indicating if the key exists.
func (m *ValueMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

// Delete removes a key from the map.
func (m *ValueMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

// Clone returns a deep copy of the ValueMap.
func (m *ValueMap[K, V]) Clone() *ValueMap[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cp := make(map[K]V, len(m.data))
	for k, v := range m.data {
		cp[k] = v
	}
	return &ValueMap[K, V]{data: cp}
}

// Merge adds or overwrites keys from another ValueMap into this one.
func (m *ValueMap[K, V]) Merge(other *ValueMap[K, V]) {
	other.mu.RLock()
	defer other.mu.RUnlock()

	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range other.data {
		m.data[k] = v
	}
}

// Equal performs a deep equality check.
func (m *ValueMap[K, V]) Equal(other *ValueMap[K, V]) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	other.mu.RLock()
	defer other.mu.RUnlock()

	if len(m.data) != len(other.data) {
		return false
	}
	for k, v := range m.data {
		ov, ok := other.data[k]
		if !ok || !reflect.DeepEqual(v, ov) {
			return false
		}
	}
	return true
}

// Keys returns a slice of all keys.
func (m *ValueMap[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all values.
func (m *ValueMap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	values := make([]V, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// Len returns the number of key-value pairs.
func (m *ValueMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Clear removes all entries from the map.
func (m *ValueMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[K]V)
}

// Raw returns a read-only copy of the internal map.
func (m *ValueMap[K, V]) Raw() map[K]V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cp := make(map[K]V, len(m.data))
	for k, v := range m.data {
		cp[k] = v
	}
	return cp
}
