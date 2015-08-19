package heap

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestFloat64s_DefLess(t *testing.T) {
	var h Float64s

	assert.Equal(t, "len", h.Len(), 0)

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(3)

	assert.Equal(t, "len", h.Len(), 4)
	assert.Equal(t, "peek", h.Peek(), 1.)

	res := []float64{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.Equal(t, "res", res, []float64{1, 2, 3, 5})

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(3)
	assert.Equal(t, "PopAll", h.PopAll(), []float64{5, 3, 2, 1})
}

func TestFloat64s_CustomLess(t *testing.T) {
	h := NewFloat64s(func(x, y float64) bool {
		return x > y
	}, 5)

	assert.Equal(t, "len", h.Len(), 0)

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(3)

	assert.Equal(t, "len", h.Len(), 4)
	assert.Equal(t, "peek", h.Peek(), 5.)
	res := []float64{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []float64{5, 3, 2, 1})

	h.Push(5)
	h.Push(2)
	h.Push(1)
	h.Push(3)
	assert.Equal(t, "PopAll", h.PopAll(), []float64{1, 2, 3, 5})
}
