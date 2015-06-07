package heap

type Strings struct {
	less func(x, y string) bool
	list []string
}

// Len returns the number of elements in the current heap.
func (h *Strings) Len() int {
	return len(h.list)
}

// Less compares the i-th and j-th elemnts in the heap.
func (h *Strings) Less(i, j int) bool {
	if h.less == nil {
		return h.list[i] < h.list[j]
	}

	return h.less(h.list[i], h.list[j])
}

// Swap exchanges values of the i-th and j-th elements.
func (h *Strings) Swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

// Push inserts an element to the heap.
func (h *Strings) Push(x string) {
	h.list = append(h.list, x)
	PushLast(h)
}

// Pop removes the top element from the heap and returns it.
func (h *Strings) Pop() string {
	PopToLast(h)
	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// NewInts returns a *Strings with customized less func and initial capacity.
func NewStrings(less func(x, y string) bool, cap int) *Strings {
	return &Strings{
		less: less,
		list: make([]string, 0, cap),
	}
}
