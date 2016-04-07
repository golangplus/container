package counters

import (
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestByString(t *testing.T) {
	var c1 ByString
	assert.Equal(t, "c1", c1, ByString{})

	c1.Add("one", 1)
	c1.Add("two", 2)
	assert.Equal(t, "c1", c1, ByString{"one": 1, "two": 2})

	c1.MergeWith(nil)

	c1.MergeWith(ByString{"two": 2, "three": 3})
	assert.Equal(t, "c1", c1, ByString{"one": 1, "two": 4, "three": 3})

	c1 = nil
	c1.MergeWith(ByString{"two": 2, "three": 3})
	assert.Equal(t, "c1", c1, ByString{"two": 2, "three": 3})
}
