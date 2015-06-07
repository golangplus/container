package heap

import (
	"sort"
)

// A heap for strings. The pointer to the zero value of Float64s is a heap with default less
// func which compares string values on the natual order.
// Use NewFloat64s to customize less func and initial capacity.
type Float64s struct {
	less func(i, j int) bool
	list []float64
}

// Len returns the number of elements in the current heap.
func (h *Float64s) Len() int {
	return len(h.list)
}

// Push inserts an element to the heap.
func (h *Float64s) Push(x float64) {
	h.list = append(h.list, x)

	if h.less == nil {
		PushLast(sort.Float64Slice(h.list))
	} else {
		PushLastF(len(h.list), h.less, sort.Float64Slice(h.list).Swap)
	}
}

// Pop removes the top element from the heap and returns it.
func (h *Float64s) Pop() float64 {
	if h.less == nil {
		PopToLast(sort.Float64Slice(h.list))
	} else {
		PopToLastF(len(h.list), h.less, sort.Float64Slice(h.list).Swap)
	}

	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// NewInts returns a *Float64s with customized less func and initial capacity.
func NewFloat64s(less func(x, y float64) bool, cap int) *Float64s {
	h := &Float64s{}

	h.less = func(i, j int) bool {
		return less(h.list[i], h.list[j])
	}
	if cap > 0 {
		h.list = make([]float64, 0, cap)
	}

	return h
}