package heap

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestIntHeap_DefLess(t *testing.T) {
	var h IntHeap

	assert.Equal(t, "len", h.Len(), 0)

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(3)

	assert.Equal(t, "len", h.Len(), 4)

	res := []int{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []int{1, 2, 3, 5})
}

func TestIntHeap_CustomLess(t *testing.T) {
	data := [...]int{5, 2, 1, 3}

	h := NewIntHeap(func(i, j int) bool {
		return data[i] < data[j]
	}, 5)

	assert.Equal(t, "len", h.Len(), 0)

	h.Push(0)
	h.Push(1)
	h.Push(2)
	h.Push(3)

	assert.Equal(t, "len", h.Len(), 4)
	res := []int{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []int{2, 1, 3, 0})
}
