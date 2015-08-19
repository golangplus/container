// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package heap is an alternative to the standard "heap" package.
It implements a very similar function to the builtin
heap(priority queue) package except the elements are not necessarily
interface{}, but can be of any type.

The trick is to use the last element as the in/out place. Push/Pop/Remove are
replaced with PushLast/PopToLast/RemoveToLast, respectively.

A heap with int value can be easily implemented as follow:

    type IntHeap []int
    func (h *IntHeap) Pop() int {
      heap.PopToLast((*sort.IntSlice)(h))
      res := (*h)[len(*h) - 1]
      *h = (*h)[:len(*h) - 1]

      return res
    }

    func (h *IntHeap) Push(x int) {
      *h = append(*h, x)
      heap.PushLast((*sort.IntSlice)(h))
    }

Use of the IntHeap:

    hp := IntHeap{3, 1, 5}
    heap.Init(sort.Interface(h))
    hp.Push(4)
    ...
    value := hp.Pop()

PushLastF/PopToLastF/RemoveToLastF takes funcs other than sort.Interface as the argument.
E.g., a heap with type T value can be implemented as follow:

    type THeap []T
    func (h *THeap) Pop() T {
      heap.PopToLastF(len(*h), func(i, j int) bool {
        ti, tj := (*h)[i], (*h)h[j]
        // return whether ti < tj
      }, func(i, j int) {
        (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
      })

      res := (*h)[len(*h) - 1]
      *h = (*h)[:len(*h) - 1]

      return res
    }

    func (h *THeap) Push(x T) {
      *h = append(*h, x)
      heap.PushLastF(len(*h), func(i, j int) bool {
        ti, tj := (*h)[i], (*h)h[j]
        // return whether ti < tj
      }, func(i, j int) {
        (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
      })
    }
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
		heapDown(n, h.Less, h.Swap, i)
	}
}

// Similar to Init but with interface provided by funcs.
func InitF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {
	// heapify
	for i := Len/2 - 1; i >= 0; i-- {
		heapDown(Len, Less, Swap, i)
	}
}

// Push pushes the last element of the heap, which was not considered as part
// of the heap, onto the heap.
// The complexity is O(log(N)), where N = h.Len().
//
// NOTE You need to append the element to be pushed as the last element before
// calling to this method.
func PushLast(h sort.Interface) {
	heapUp(h.Less, h.Swap, h.Len()-1)
}

// Similar to PushLast but with interface provided by funcs.
func PushLastF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {
	heapUp(Less, Swap, Len-1)
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
	heapDown(n1, h.Less, h.Swap, 0)
}

// Similar to PopToLast but with interface provided by funcs.
func PopToLastF(Len int, Less func(i, j int) bool, Swap func(i, j int)) {
	Swap(0, Len-1)
	heapDown(Len-1, Less, Swap, 0)
}

// Fix re-establishes the heap ordering after the value of the element at the index has changed.
// Changing the value of the element at the index and then calling Fix is equivalent to,
// but less expensive than, calling RemoveToLast(h, index) followed by a PushLast.
// The complexity is O(log(N)), where N = h.Len().
func Fix(h sort.Interface, index int) {
	heapDown(h.Len(), h.Less, h.Swap, index)
	heapUp(h.Less, h.Swap, index)
}

// Similar to Fix but with interface provided by funcs.
func FixF(Len int, Less func(i, j int) bool, Swap func(i, j int), index int) {
	heapDown(Len, Less, Swap, index)
	heapUp(Less, Swap, index)
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
		heapDown(n, h.Less, h.Swap, i)
		heapUp(h.Less, h.Swap, i)
	}
}

// Similar to RemoveToLast but with interface provided by funcs.
func RemoveToLastF(Len int, Less func(i, j int) bool, Swap func(i, j int), i int) {
	n := Len - 1
	if n != i {
		Swap(i, n)
		heapDown(n, Less, Swap, i)
		heapUp(Less, Swap, i)
	}
}

func heapUp(less func(i, j int) bool, swap func(i, j int), i int) {
	for i > 0 {
		p := (i - 1) / 2 // p is the parent of i
		if !less(i, p) {
			// h[p] <= h[i], already in order
			break
		}
		swap(i, p)
		i = p // move to upper level
	}
}

func heapDown(n int, less func(i, j int) bool, swap func(i, j int), i int) {
	for {
		l := 2*i + 1 // left child
		if l >= n || l < 0 {
			break
		}
		c := l // c is initialized with l
		// Set c to r if h[r] < h[l]
		if r := l + 1; r < n && r > 0 && less(r, l) {
			c = r
		}
		if !less(c, i) {
			// h[i] <= h[c], already in order
			break
		}
		swap(i, c)
		i = c // move to lower level
	}
}
