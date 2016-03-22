package heap

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestStrings_DefLess(t *testing.T) {
	var h Strings

	assert.Equal(t, "len", h.Len(), 0)

	h.Push("Elmo")
	h.Push("Big Bird")
	h.Push("Abby")
	h.Push("Count")

	assert.Equal(t, "len", h.Len(), 4)
	assert.Equal(t, "peek", h.Peek(), "Abby")

	res := []string{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.Equal(t, "res", res, []string{"Abby", "Big Bird", "Count", "Elmo"})

	h.Push("Elmo")
	h.Push("Big Bird")
	h.Push("Abby")
	h.Push("Count")
	assert.Equal(t, "PopAll", h.PopAll(), []string{"Elmo", "Count", "Big Bird", "Abby"})
}

func TestStrings_UnRef(t *testing.T) {
	var h Strings
	h.Push("Hello")
	h.Pop()
	assert.Equal(t, "h.list[0]", h.list[:1][0], "")
}

func TestStrings_CustomLess(t *testing.T) {
	data := map[string]int{
		"Abby":     5,
		"Big Bird": 2,
		"Count":    1,
		"Elmo":     3,
	}

	h := NewStrings(func(x, y string) bool {
		return data[x] < data[y]
	}, 5)

	assert.Equal(t, "len", h.Len(), 0)

	h.Push("Abby")
	h.Push("Big Bird")
	h.Push("Count")
	h.Push("Elmo")

	assert.Equal(t, "len", h.Len(), 4)
	assert.Equal(t, "peek", h.Peek(), "Count")
	res := []string{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []string{"Count", "Big Bird", "Elmo", "Abby"})

	h.Push("Abby")
	h.Push("Big Bird")
	h.Push("Count")
	h.Push("Elmo")
	assert.Equal(t, "PopAll", h.PopAll(), []string{"Abby", "Elmo", "Big Bird", "Count"})
}
