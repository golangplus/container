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

	res := []string{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []string{"Abby", "Big Bird", "Count", "Elmo"})
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
	res := []string{h.Pop(), h.Pop(), h.Pop(), h.Pop()}
	assert.StringEqual(t, "res", res, []string{"Count", "Big Bird", "Elmo", "Abby"})
}
