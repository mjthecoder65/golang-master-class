package dsa

import "fmt"

// Heap is a data structure that is a complete binary tree.
// Heap is a binary tree that satisfies the heap property.

// Heap interface

type MinHeap struct {
	array []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Insert(value int) {
	h.array = append(h.array, value)
	index := len(h.array) - 1

	// Heapify Up
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] <= h.array[index] {
			break
		}
		h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
		index = parentIndex
	}
}

func (h *MinHeap) RemoveMin() (int, error) {
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	min := h.array[0]
	lastIndex := len(h.array) - 1

	// Swap the first and last elements
	h.array[0], h.array[lastIndex] = h.array[lastIndex], h.array[0]

	// Remove the last element
	h.array = h.array[:lastIndex]

	// Heapify Down
	index := 0
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < len(h.array) && h.array[leftChild] < h.array[smallest] {
			smallest = leftChild
		}
		if rightChild < len(h.array) && h.array[rightChild] < h.array[smallest] {
			smallest = rightChild
		}

		if smallest == index {
			break
		}

		h.array[index], h.array[smallest] = h.array[smallest], h.array[index]
		index = smallest
	}

	return min, nil
}

func (h *MinHeap) GetMin() (int, error) {
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.array[0], nil
}
