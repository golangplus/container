package heap

import (
	"sort"
)

// A heap for strings. The pointer to the zero value of Float64s is a heap with default less
// func which compares string values on the natual order.
// Use NewFloat64s to customize less func and initial capacity.
type Float64s struct {
	less func(i, j int) bool
	list sort.Float64Slice
}

// Len returns the number of elements in the current heap.
func (h *Float64s) Len() int {
	return len(h.list)
}

// Push inserts an element to the heap.
func (h *Float64s) Push(x float64) {
	h.list = append(h.list, x)

	if h.less == nil {
		PushLastF(len(h.list), h.list.Less, h.list.Swap)
	} else {
		PushLastF(len(h.list), h.less, h.list.Swap)
	}
}

// Peek returns the top most element. It panics if the heap is empty.
func (h *Float64s) Peek() float64 {
	return h.list[0]
}

// Pop removes the top element from the heap and returns it.
func (h *Float64s) Pop() float64 {
	if h.less == nil {
		PopToLastF(len(h.list), h.list.Less, h.list.Swap)
	} else {
		PopToLastF(len(h.list), h.less, h.list.Swap)
	}

	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// PopAll pops and returns all elements of the heap in reverse order.
func (h *Float64s) PopAll() []float64 {
	for n := h.Len(); n > 1; n-- {
		if h.less == nil {
			PopToLastF(n, h.list.Less, h.list.Swap)
		} else {
			PopToLastF(n, h.less, h.list.Swap)
		}
	}
	res := h.list
	h.list = nil
	return res
}

// NewFloat64s returns a *Float64s with customized less func and initial capacity.
func NewFloat64s(less func(x, y float64) bool, cap int) *Float64s {
	h := &Float64s{}

	if less != nil {
		h.less = func(i, j int) bool {
			return less(h.list[i], h.list[j])
		}
	}
	if cap > 0 {
		h.list = make([]float64, 0, cap)
	}

	return h
}
