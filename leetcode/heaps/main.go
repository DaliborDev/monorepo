package main

import "fmt"

// MaxHeap struct has a slice that holds the data in an array-like format
type MaxHeap struct {
	data []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.maxHeapifyUp(len(h.data) - 1)
}

// Extract returns the largest element and removes it from the heap
func (h *MaxHeap) Extract() int {
	if len(h.data) <= 0 {
		panic("Cant extract from empty heap")
	}

	heapMax := h.data[0]
	last := len(h.data) - 1

	h.data[0] = h.data[last]
	h.data = h.data[:last]

	h.maxHeapifyDown(0)

	return heapMax
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.data[getParent(index)] < h.data[index] {
		h.swap(getParent(index), index)
		index = getParent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	last := len(h.data) - 1
	left, right := getLeft(index), getRight(index)
	childToCompare := 0

	for left <= last {
		if left == last {
			childToCompare = left
		} else if h.data[left] > h.data[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if h.data[index] < h.data[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			left, right = getLeft(index), getRight(index)
		} else {
			return
		}
	}
}

// get parent index
func getParent(i int) int {
	return (i - 1) / 2
}

// get left child index
func getLeft(i int) int {
	return i*2 + 1
}

// get right child index
func getRight(i int) int {
	return i*2 + 2
}

// swap values within data slice
func (h *MaxHeap) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}

func main() {
	m := &MaxHeap{}
	m2 := &MaxHeap{}

	buildHeap := []int{10, 20, 30}
	buildHeap2 := []int{10, 21, 2, 43, 12, 21, 55, 9, 0, 32}

	for _, v := range buildHeap {
		m.Insert(v)
	}

	for _, v := range buildHeap2 {
		m2.Insert(v)
		fmt.Println(m2)
	}

	for i := 0; i < len(m2.data); i++ {
		m2.Extract()
		fmt.Println(m2)
	}

}
