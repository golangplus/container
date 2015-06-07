package heap

// A heap for ints. The pointer to the zero value of Ints is a heap with default less
// func which compare the int value upon its natual order.
// Use NewInts to customize less func and initial capacity.
type Ints struct {
	less func(x, y int) bool
	list []int
}

// Len returns the number of elements in the current heap.
func (h *Ints) Len() int {
	return len(h.list)
}

// Less compares the i-th and j-th elemnts in the heap.
func (h *Ints) Less(i, j int) bool {
	if h.less == nil {
		return h.list[i] < h.list[j]
	}

	return h.less(h.list[i], h.list[j])
}

// Swap exchanges values of the i-th and j-th elements.
func (h *Ints) Swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

// Push inserts an element to the heap.
func (h *Ints) Push(x int) {
	h.list = append(h.list, x)
	PushLast(h)
}

// Pop removes the top element from the heap and returns it.
func (h *Ints) Pop() int {
	PopToLast(h)
	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// NewInts returns a *Ints with customized less func and initial capacity.
// NOTE unlike Ints.Less, the parameters of less are the integer values to be compared
// not the indexes.
func NewInts(less func(x, y int) bool, cap int) *Ints {
	return &Ints{
		less: less,
		list: make([]int, 0, cap),
	}
}
