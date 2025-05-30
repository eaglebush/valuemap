package valuemap

import (
	"maps"
	"sync"
)

type ValueMap[K comparable, V any] struct {
	data map[K]V
}

// New returns a new pointer to a thread-safe ValueMap.
func New[K comparable, V any]() *ValueMap[K, V] {
	return &ValueMap[K, V]{data: make(map[K]V)}
}

// FromMap returns a new ValueMap initialized with a copy of an existing map.
func FromMap[K comparable, V any](m map[K]V) *ValueMap[K, V] {
	cp := make(map[K]V, len(m))
	maps.Copy(cp, m)
	return &ValueMap[K, V]{data: cp}
}

// Set assigns a value to a key.
//
// mu is an external mutex to lock the internal map during value assigning
func (m *ValueMap[K, V]) Set(mu *sync.RWMutex, key K, value V) {
	mu.Lock()
	defer mu.Unlock()
	m.data[key] = value
}

// Get retrieves a value and a boolean indicating if the key exists.
//
// mu is an external mutex to lock the internal map during value retrieval
func (m *ValueMap[K, V]) Get(mu *sync.RWMutex, key K) (V, bool) {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

// Delete removes a key from the map.
//
// mu is an external mutex to lock the internal map during key deletion
func (m *ValueMap[K, V]) Delete(mu *sync.RWMutex, key K) {
	mu.Lock()
	defer mu.Unlock()
	delete(m.data, key)
}

// Clone returns a deep copy of the ValueMap.
//
// mu is an external mutex to lock the internal map during cloning
func (m *ValueMap[K, V]) Clone(mu *sync.RWMutex) *ValueMap[K, V] {
	mu.Lock()
	defer mu.Unlock()
	cp := make(map[K]V, len(m.data))
	maps.Copy(cp, m.data)
	return &ValueMap[K, V]{data: cp}
}

// Merge adds or overwrites keys from another ValueMap into this one.
//
// mu is an external mutex to lock the internal map during value merging
func (m *ValueMap[K, V]) Merge(mu *sync.RWMutex, other *ValueMap[K, V]) {
	mu.Lock()
	defer mu.Unlock()
	maps.Copy(m.data, other.data)
}

// Keys returns a slice of all keys.
//
// mu is an external mutex to lock the internal map during key retrieval
func (m *ValueMap[K, V]) Keys(mu *sync.RWMutex) []K {
	mu.RLock()
	defer mu.RUnlock()
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all values.
//
// mu is an external mutex to lock the internal map during value retrieval
func (m *ValueMap[K, V]) Values(mu *sync.RWMutex) []V {
	mu.RLock()
	defer mu.RUnlock()
	values := make([]V, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// Len returns the number of key-value pairs.
//
// mu is an external mutex to lock the internal map during map content counting
func (m *ValueMap[K, V]) Len(mu *sync.RWMutex) int {
	mu.RLock()
	defer mu.RUnlock()
	return len(m.data)
}

// Clear removes all entries from the map.
//
// mu is an external mutex to lock the internal map during map clearing
func (m *ValueMap[K, V]) Clear(mu *sync.RWMutex) {
	mu.Lock()
	defer mu.Unlock()
	m.data = make(map[K]V)
}

// Raw returns a read-only copy of the internal map.
//
// mu is an external mutex to lock the internal map during raw value retrieval
func (m *ValueMap[K, V]) Raw(mu *sync.RWMutex) map[K]V {
	mu.RLock()
	defer mu.RUnlock()
	cp := make(map[K]V, len(m.data))
	for k, v := range m.data {
		cp[k] = v
	}
	return cp
}
