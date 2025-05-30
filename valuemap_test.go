package valuemap

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	mu := sync.RWMutex{}
	m1 := New[string, int]()
	m1.Set(&mu, "a", 1)
	m1.Set(&mu, "b", 2)
	fmt.Println(m1.Get(&mu, "a")) // 1
	fmt.Println(m1.Get(&mu, "b")) // 2

	m2 := m1.Clone(&mu)
	m2.Set(&mu, "a", 99)

	fmt.Println(m1.Get(&mu, "a")) // 1
	fmt.Println(m2.Get(&mu, "a")) // 99

	m1.Merge(&mu, m2)
	fmt.Println(m1.Get(&mu, "a")) // 99
}
