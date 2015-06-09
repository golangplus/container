package heap

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/golangplus/testing/assert"
)

type intHeap []int

func (h *intHeap) Pop() int {
	PopToLast(sort.IntSlice(*h))
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *intHeap) Push(x int) {
	*h = append(*h, x)
	PushLast(sort.IntSlice(*h))
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

type Data struct {
	Value    string
	Priority int
}

type DataHeap []Data

func (h *DataHeap) Pop() Data {
	PopToLastF(len(*h), func(i, j int) bool {
		return (*h)[i].Priority < (*h)[j].Priority
	}, func(i, j int) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	})
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return res
}

func (h *DataHeap) Push(x Data) {
	*h = append(*h, x)
	PushLastF(len(*h), func(i, j int) bool {
		return (*h)[i].Priority < (*h)[j].Priority
	}, func(i, j int) {
		(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	})
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
			t.Errorf("%d should be larger than %d", cur, last)
		}
		last = cur
	}
}
