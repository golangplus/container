package heap

import (
	"sort"
)

// A heap for strings. The pointer to the zero value of Strings is a heap with default less
// func which compares string values on the natual order.
// Use NewStrings to customize less func and initial capacity.
type Strings struct {
	less func(i, j int) bool
	list []string
}

// Len returns the number of elements in the current heap.
func (h *Strings) Len() int {
	return len(h.list)
}

// Push inserts an element to the heap.
func (h *Strings) Push(x string) {
	h.list = append(h.list, x)

	if h.less == nil {
		PushLast(sort.StringSlice(h.list))
	} else {
		PushLastF(len(h.list), h.less, sort.StringSlice(h.list).Swap)
	}
}

// Pop removes the top element from the heap and returns it.
func (h *Strings) Pop() string {
	if h.less == nil {
		PopToLast(sort.StringSlice(h.list))
	} else {
		PopToLastF(len(h.list), h.less, sort.StringSlice(h.list).Swap)
	}

	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// NewInts returns a *Strings with customized less func and initial capacity.
func NewStrings(less func(x, y string) bool, cap int) *Strings {
	h := &Strings{}

	h.less = func(i, j int) bool {
		return less(h.list[i], h.list[j])
	}
	if cap > 0 {
		h.list = make([]string, 0, cap)
	}

	return h
}
