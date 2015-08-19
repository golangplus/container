package heap

import (
	"container/heap"
	"math/rand"
	"sort"
	"testing"

	"github.com/golangplus/testing/assert"
)

type intHeap []int

func (h *intHeap) Pop() int {
	PopToLast((*sort.IntSlice)(h))
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *intHeap) Push(x int) {
	*h = append(*h, x)
	PushLast((*sort.IntSlice)(h))
}

func TestIntHeap(t *testing.T) {
	var h intHeap

	for i := 0; i < 1000; i++ {
		h.Push(rand.Int())
	}

	assert.Equal(t, "len(h)", len(h), 1000)

	peek := h[0]
	last := h.Pop()
	assert.Equal(t, "h.Peek()", peek, last)
	for i := 1; i < 1000; i++ {
		cur := h.Pop()
		if cur < last {
			t.Errorf("%d should be larger than %d", cur, last)
		}
		last = cur
	}
}

func BenchmarkPlusIntHeapInter(b *testing.B) {
	var data [M]int
	for i := range data {
		data[i] = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var h intHeap
		for _, vl := range data {
			h.Push(vl)
		}
		for len(h) > 0 {
			h.Pop()
		}
	}
}

type Data struct {
	Value    string
	Priority int
}

type DataHeap []Data

func (h DataHeap) Len() int           { return len(h) }
func (h DataHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h DataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DataHeap) Pop() Data {
	/*
		PopToLastF(len(*h), func(i, j int) bool {
			return (*h)[i].Priority < (*h)[j].Priority
		}, func(i, j int) {
			(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		})
	*/
	PopToLast(h)

	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *DataHeap) Push(x Data) {
	*h = append(*h, x)
	PushLast(h)
	/*
		PushLastF(len(*h), func(i, j int) bool {
			return (*h)[i].Priority < (*h)[j].Priority
		}, func(i, j int) {
			(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
		})*/
}

type DataHeapF []Data

func (h *DataHeapF) Pop() Data {
	PopToLastF(len(*h), func(i, j int) bool {
		return (*h)[i].Priority < (*h)[j].Priority
	}, func(i, j int) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	})

	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *DataHeapF) Push(x Data) {
	*h = append(*h, x)
	PushLastF(len(*h), func(i, j int) bool {
		return (*h)[i].Priority < (*h)[j].Priority
	}, func(i, j int) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	})
}

type builtinDataHeap []Data

func (h builtinDataHeap) Len() int           { return len(h) }
func (h builtinDataHeap) Less(i, j int) bool { return h[i].Priority < h[j].Priority }
func (h builtinDataHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *builtinDataHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Data))
}

func (h *builtinDataHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestDataHeap(t *testing.T) {
	var h DataHeap

	for i := 0; i < 1000; i++ {
		h.Push(Data{"A", rand.Int()})
	}

	assert.Equal(t, "len(h)", len(h), 1000)

	peek := h[0]
	last := h.Pop()
	assert.Equal(t, "h.Peek()", peek, last)
	for i := 1; i < 1000; i++ {
		cur := h.Pop()
		if cur.Priority < last.Priority {
			t.Errorf("%v should be larger than %v", cur, last)
		}
		last = cur
	}
}

func BenchmarkDataHeap_Plus(b *testing.B) {
	var data [M]Data
	for i := range data {
		data[i].Priority = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var h DataHeap
		for _, vl := range data {
			h.Push(vl)
		}
		for len(h) > 0 {
			h.Pop()
		}
	}
}

func BenchmarkDataHeap_F(b *testing.B) {
	var data [M]Data
	for i := range data {
		data[i].Priority = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var h DataHeapF
		for _, vl := range data {
			h.Push(vl)
		}
		for len(h) > 0 {
			h.Pop()
		}
	}
}

func BenchmarkDataHeap_Pkg(b *testing.B) {
	var data [M]Data
	for i := range data {
		data[i].Priority = rand.Int()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var h builtinDataHeap
		for _, vl := range data {
			heap.Push(&h, vl)
		}
		for len(h) > 0 {
			heap.Pop(&h)
		}
	}
}
