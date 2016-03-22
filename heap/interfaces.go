package heap

// A heap for interface{}. Use NewInterfaces to create an instance.
type Interfaces interface {
	// Len returns the number of elements in the current heap.
	Len() int
	// Push inserts an element to the heap.
	Push(x interface{})
	// Pop removes the top element from the heap and returns it.
	Pop() interface{}
	// PopAll pops and returns all elements of the heap in reverse order.
	PopAll() []interface{}
	// Peek returns the top most element. It panics if the heap is empty.
	Peek() interface{}
	// TopNPush inserts an element to the heap if the heap does not reach its
	// capacity. Otherwise, if the top element is less then the new element
	// the top element is removed and the new one is inserted.
	// This method is used to generate a top N largest elements where N is
	// the capacity of the heap.
	TopNPush(x interface{})
	// TopNPopAll is similar to PopAll but keep the capacity of the heap
	// unchanged.
	TopNPopAll() []interface{}
}

type interfaces struct {
	less func(i, j int) bool
	list []interface{}
}

// Interfaces.Len
func (h *interfaces) Len() int { return len(h.list) }

// Interfaces.Push
func (h *interfaces) Push(x interface{}) {
	h.list = append(h.list, x)

	PushLastF(len(h.list), h.less, func(i, j int) { h.list[i], h.list[j] = h.list[j], h.list[i] })
}

// Interfaces.TopNPush
func (h *interfaces) TopNPush(x interface{}) {
	if len(h.list) < cap(h.list)-1 {
		h.Push(x)
		return
	}
	h.list = append(h.list, x)
	if h.less(0, cap(h.list)-1) {
		h.list[0], h.list[cap(h.list)-1] = x, nil
		FixF(cap(h.list)-1, h.less, func(i, j int) { h.list[i], h.list[j] = h.list[j], h.list[i] }, 0)
	}
	h.list = h.list[:cap(h.list)-1]
}

// Interfaces.Pop
func (h *interfaces) Pop() interface{} {
	PopToLastF(len(h.list), h.less, func(i, j int) { h.list[i], h.list[j] = h.list[j], h.list[i] })

	res := h.list[len(h.list)-1]
	h.list[len(h.list)-1] = nil // remove the reference in h.list
	h.list = h.list[:len(h.list)-1]

	return res
}

// Interfaces.PopAll
func (h *interfaces) PopAll() []interface{} {
	for n := h.Len(); n > 1; n-- {
		PopToLastF(n, h.less, func(i, j int) { h.list[i], h.list[j] = h.list[j], h.list[i] })
	}
	res := h.list
	h.list = nil
	return res
}

// Interfaces.TopNPopAll
func (h *interfaces) TopNPopAll() []interface{} {
	for n := h.Len(); n > 1; n-- {
		PopToLastF(n, h.less, func(i, j int) { h.list[i], h.list[j] = h.list[j], h.list[i] })
	}
	res := append([]interface{}(nil), h.list...)
	for i := range h.list {
		h.list[i] = nil // remove the reference in h.list
	}
	h.list = h.list[:0]
	return res
}

// Interfaces.Peek
func (h *interfaces) Peek() interface{} {
	return h.list[0]
}

// NewInterfaces returns an instance of Interfaces with a customized less func and the initial capacity.
func NewInterfaces(less func(x, y interface{}) bool, cap int) Interfaces {
	h := &interfaces{}

	h.less = func(i, j int) bool {
		return less(h.list[i], h.list[j])
	}
	if cap > 0 {
		h.list = make([]interface{}, 0, cap+1)
	}

	return h
}
