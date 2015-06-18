package heap

import (
	"container/heap"
	"math/rand"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestInts_DefLess(t *testing.T) {
	var h Ints

	assert.Equal(t, "len", h.Len(), 0)

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(3)

	assert.Equal(t, "len", h.Len(), 4)

	res := []int{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []int{1, 2, 3, 5})
}

func TestInts_CustomLess(t *testing.T) {
	data := [...]int{5, 2, 1, 3}

	h := NewInts(func(i, j int) bool {
		return data[i] < data[j]
	}, 5)

	assert.Equal(t, "len", h.Len(), 0)

	h.Push(0)
	h.Push(1)
	h.Push(2)
	h.Push(3)

	assert.Equal(t, "len", h.Len(), 4)
	res := []int{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []int{2, 1, 3, 0})
}

type builtinIntHeap []int

func (h builtinIntHeap) Len() int           { return len(h) }
func (h builtinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h builtinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *builtinIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *builtinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const M = 10000

func BenchmarkBuiltinIntHeap(b *testing.B) {
	var data [M]int
	for i := range data {
		data[i] = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var h builtinIntHeap
		for _, vl := range data {
			heap.Push(&h, vl)
		}
		for len(h) > 0 {
			heap.Pop(&h)
		}
	}
}

func BenchmarkPlusIntHeap(b *testing.B) {
	var data [M]int
	for i := range data {
		data[i] = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var h Ints
		for _, vl := range data {
			h.Push(vl)
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}

func BenchmarkPlusIntHeap_Less(b *testing.B) {
	var data [M]int
	for i := range data {
		data[i] = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h := NewInts(func(x, y int) bool {
			return x < y
		}, 0)
		for _, vl := range data {
			h.Push(vl)
		}
		for h.Len() > 0 {
			h.Pop()
		}
	}
}
