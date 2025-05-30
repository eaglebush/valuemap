package valuemap

import "reflect"

type ValueMap[K comparable, V any] struct {
	data map[K]V
}

// New returns a new ValueMap
func New[K comparable, V any]() ValueMap[K, V] {
	return ValueMap[K, V]{data: make(map[K]V)}
}

// FromMap creates a ValueMap from an existing map (clones it)
func FromMap[K comparable, V any](m map[K]V) ValueMap[K, V] {
	newMap := make(map[K]V, len(m))
	for k, v := range m {
		newMap[k] = v
	}
	return ValueMap[K, V]{data: newMap}
}

// Set sets a key-value pair
func (m *ValueMap[K, V]) Set(key K, value V) {
	m.data[key] = value
}

// Get retrieves a value and ok flag
func (m ValueMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

// Delete removes a key
func (m *ValueMap[K, V]) Delete(key K) {
	delete(m.data, key)
}

// Clone returns a deep copy
func (m ValueMap[K, V]) Clone() ValueMap[K, V] {
	return FromMap(m.data)
}

// Merge combines another ValueMap into the current one (overwrites values)
func (m *ValueMap[K, V]) Merge(other ValueMap[K, V]) {
	for k, v := range other.data {
		m.data[k] = v
	}
}

func (m ValueMap[K, V]) Equal(other ValueMap[K, V]) bool {
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

// Keys returns a slice of keys
func (m ValueMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of values
func (m ValueMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// Len returns the number of entries
func (m ValueMap[K, V]) Len() int {
	return len(m.data)
}

// Clear removes all entries
func (m *ValueMap[K, V]) Clear() {
	m.data = make(map[K]V)
}

// Raw returns the internal map for read-only use
func (m ValueMap[K, V]) Raw() map[K]V {
	return m.data
}
