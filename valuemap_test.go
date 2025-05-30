package valuemap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := New[string, int]()
	m1.Set("a", 1)
	m1.Set("b", 2)

	m2 := m1.Clone()
	m2.Set("a", 99)

	fmt.Println(m1.Get("a"))  // 1
	fmt.Println(m2.Get("a"))  // 99
	fmt.Println(m1.Equal(m2)) // false

	m1.Merge(m2)
	fmt.Println(m1.Get("a")) // 99
}
