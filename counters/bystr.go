// Package counters provide types supporting flexible counters within a map.
// Since counters are mergable, all types are designed not to be thread safe.
package counters

// ByString is a map of counters with a string as the key.
type ByString map[string]int

// Add increases the value of a specific key and returns the updated value.
func (bs *ByString) Add(key string, inc int) int {
	if *bs == nil {
		*bs = make(map[string]int)
	}
	v := (*bs)[key] + inc
	(*bs)[key] = v
	return v
}

// MergeWith adds values of another ByString counters into this one.
func (bs *ByString) MergeWith(that ByString) {
	if len(that) == 0 {
		return
	}
	if *bs == nil {
		*bs = make(map[string]int)
	}
	for k, v := range that {
		(*bs)[k] = (*bs)[k] + v
	}
}
