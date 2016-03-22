package heap

import (
	"sort"
)

// A heap for strings. The pointer to the zero value of Strings is a heap with default less
// func which compares string values on the natual order.
// Use NewStrings to customize less func and initial capacity.
type Strings struct {
	less func(i, j int) bool
	list sort.StringSlice
}

// Len returns the number of elements in the current heap.
func (h *Strings) Len() int {
	return len(h.list)
}

// Push inserts an element to the heap.
func (h *Strings) Push(x string) {
	h.list = append(h.list, x)

	if h.less == nil {
		PushLastF(len(h.list), h.list.Less, h.list.Swap)
	} else {
		PushLastF(len(h.list), h.less, h.list.Swap)
	}
}

// Peek returns the top most element. It panics if the heap is empty.
func (h *Strings) Peek() string {
	return h.list[0]
}

// Pop removes the top element from the heap and returns it.
func (h *Strings) Pop() string {
	if h.less == nil {
		PopToLastF(len(h.list), h.list.Less, h.list.Swap)
	} else {
		PopToLastF(len(h.list), h.less, h.list.Swap)
	}

	res := h.list[len(h.list)-1]
	h.list[len(h.list)-1] = "" // remove the reference in h.list
	h.list = h.list[:len(h.list)-1]

	return res
}

// PopAll pops and returns all elements of the heap in reverse order.
func (h *Strings) PopAll() []string {
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

// NewStrings returns a *Strings with a customized less func and the initial capacity.
func NewStrings(less func(x, y string) bool, cap int) *Strings {
	h := &Strings{}

	if less != nil {
		h.less = func(i, j int) bool {
			return less(h.list[i], h.list[j])
		}
	}
	if cap > 0 {
		h.list = make([]string, 0, cap)
	}
	return h
}
