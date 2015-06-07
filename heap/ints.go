package heap

// A heap for ints. The pointer to the zero value of IntHeap is a heap with default less
// func which compare the int value upon its natual order.
// Use NewIntHeap to customize less func and initial capacity.
type IntHeap struct {
	less func(x, y int) bool
	ints []int
}

// Len returns the number of elements in the current heap.
func (h *IntHeap) Len() int {
	return len(h.ints)
}

// Less compares the i-th and j-th elemnts in the heap.
func (h *IntHeap) Less(i, j int) bool {
	if h.less == nil {
		return h.ints[i] < h.ints[j]
	}

	return h.less(h.ints[i], h.ints[j])
}

// Swap exchanges values of the i-th and j-th elements.
func (h *IntHeap) Swap(i, j int) {
	h.ints[i], h.ints[j] = h.ints[j], h.ints[i]
}

// Push inserts an element to the heap.
func (h *IntHeap) Push(x int) {
	h.ints = append(h.ints, x)
	PushLast(h)
}

// Pop removes the top element from the heap and returns it.
func (h *IntHeap) Pop() int {
	PopToLast(h)
	res := h.ints[len(h.ints)-1]
	h.ints = h.ints[:len(h.ints)-1]

	return res
}

// NewIntHeap returns a *IntHeap with customized less func and initial capacity.
// NOTE unlike IntHeap.Less, the parameters of less are the integer values to be compared
// not the indexes.
func NewIntHeap(less func(x, y int) bool, cap int) *IntHeap {
	return &IntHeap{
		less: less,
		ints: make([]int, 0, cap),
	}
}
