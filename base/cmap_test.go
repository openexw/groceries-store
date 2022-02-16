package base

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewUmap(t *testing.T) {
	m := NewUmap()
	m.Store("user", 1)
	m.Store("age", 2)
	m.Store("name", "jack")

	assert.Equal(t, m.Load("user"), 1)
	assert.Equal(t, m.Load("age"), 2)

	hi := m.LoadOrStore("hi", "a")
	assert.Equal(t, hi, "a")

	m.LoadAndDelete("hi")
}

func TestNewCMap(t *testing.T) {
	m := NewCMap()
	m.Store("user", 1)
	m.Store("age", 2)
	m.Store("name", "jack")

	assert.Equal(t, m.Load("user"), 1)
	assert.Equal(t, m.Load("age"), 2)

	hi := m.LoadOrStore("hi", "a")
	assert.Equal(t, hi, "a")

	m.LoadAndDelete("hi")
}

func BenchmarkCmap_Store(b *testing.B) {
	m := NewCMap()
	for i := 0; i < b.N; i++ {
		m.Store(strconv.Itoa(i), i)
	}
}

func BenchmarkUmap_Store(b *testing.B) {
	m := NewUmap()
	for i := 0; i < b.N; i++ {
		m.Store(strconv.Itoa(i), i)
	}
}
//BenchmarkCmap_Store-8   	 3818686	       344.9 ns/op
//BenchmarkUmap_Store-8   	 3842978	       355.8 ns/op