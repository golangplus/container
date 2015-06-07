// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package heap is an alternative to the standard "heap" package.
It implements a very similar function to the builtin
heap(priority queue) package except the elements are not necessarily
interface{}, but can be any type.

The trick is to use the last element as the in/out place. Push/Pop/Remove are
replaced with PushLast/PopToLast/RemoveToLast, respectively. An heap with int
value can be easily implemented as follow:

    type IntHeap []int
    func (h *IntHeap) Pop() int {
        heap.PopToLast(sort.IntSlice(*h))
        res := (*h)[len(*h) - 1]
        *h = (*h)[:len(*h) - 1]

        return res
    }

    func (h *IntHeap) Push(x int) {
        *h = append(*h, x)
        heap.PushLast(sort.IntSlice(*h))
    }

Use of the IntHeap:

    hp := IntHeap{3, 1, 5}
	heap.Init(sort.Interface(h))
	hp.Push(4)
	...
	value := hp.Pop()
*/
package heap

import "sort"

// Init heapifies a non-empty array defined by the sort.Interface.
// The complexity is O(N), where N = h.Len().
//
func Init(h sort.Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		heapDown(h, i, n)
	}
}

// Push pushes the last element of the heap, which was not considered as part
// of the heap, onto the heap.
// The complexity is O(log(N)), where N = h.Len().
//
// NOTE You need to append the element to be pushed as the last element before
// calling to this method.
func PushLast(h sort.Interface) {
	heapUp(h, h.Len()-1)
}

// Pop removes the minimum element (according to Less) from the heap
// and place it as the last element of the heap.
// The complexity is O(log(N)), where N = h.Len().
//
// Same as Remove(h, 0).
//
// NOTE You need to remove the last element after calling to this method.
func PopToLast(h sort.Interface) {
	n1 := h.Len() - 1
	h.Swap(0, n1)
	heapDown(h, 0, n1)
}

// Fix re-establishes the heap ordering after the value of the element at the index has changed.
// Changing the value of the element at the index and then calling Fix is equivalent to,
// but less expensive than, calling RemoveToLast(h, index) followed by a PushLast.
// The complexity is O(log(N)), where N = h.Len().
func Fix(h sort.Interface, index int) {
	heapDown(h, index, h.Len())
	heapUp(h, index)
}

// Remove removes the element at index i from the heap and place it at the last
// element of the heap.
//
// The complexity is O(log(n)) where n = h.Len().
//
// NOTE You need to remove the last element after calling to this method.
func RemoveToLast(h sort.Interface, i int) {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		heapDown(h, i, n)
		heapUp(h, i)
	}
}

func heapUp(h sort.Interface, i int) {
	for i > 0 {
		p := (i - 1) / 2 // p is the parent of i
		if !h.Less(i, p) {
			// h[p] <= h[i], already in order
			break
		}
		h.Swap(i, p)
		i = p // move to upper level
	}
}

func heapDown(h sort.Interface, i, n int) {
	for {
		l := 2*i + 1 // left child
		if l >= n || l < 0 {
			break
		}
		c := l // c is initialized with l
		// Set c to r if h[r] < h[l]
		if r := l + 1; r < n && r > 0 && h.Less(r, l) {
			c = r
		}
		if !h.Less(c, i) {
			// h[i] <= h[c], already in order
			break
		}
		h.Swap(i, c)
		i = c // move to lower level
	}
}
