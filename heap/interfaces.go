package heap

// A heap for interface{}. Use NewInterfaces to create an instance.
type Interfaces interface {
	// Len returns the number of elements in the current heap.
	Len() int
	// Push inserts an element to the heap.
	Push(x interface{})
	// Pop removes the top element from the heap and returns it.
	Pop() interface{}
}

type interfaces struct {
	less func(i, j int) bool
	list []interface{}
}

// Interfaces.Len
func (h *interfaces) Len() int {
	return len(h.list)
}

// Interfaces.Push
func (h *interfaces) Push(x interface{}) {
	h.list = append(h.list, x)

	PushLastF(len(h.list), h.less, func(i, j int) {
		h.list[i], h.list[j] = h.list[j], h.list[i]
	})
}

// Interfaces.Pop
func (h *interfaces) Pop() interface{} {
	PopToLastF(len(h.list), h.less, func(i, j int) {
		h.list[i], h.list[j] = h.list[j], h.list[i]
	})

	res := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	return res
}

// NewInterfaces returns an instance of Interfaces with customized less func and initial capacity.
func NewInterfaces(less func(x, y interface{}) bool, cap int) Interfaces {
	h := &interfaces{}

	h.less = func(i, j int) bool {
		return less(h.list[i], h.list[j])
	}
	if cap > 0 {
		h.list = make([]interface{}, 0, cap)
	}

	return h
}
