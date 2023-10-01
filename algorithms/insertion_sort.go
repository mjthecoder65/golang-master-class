package algorithms

// Insertion Sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.
// It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.
// Insertion sort is in-place algorithm, meaning that it doesn't require additional memory.
// Then best case, when the input array is already sorted, the time complexity of insertion sort is O(n)

/*
	ANALYSIS

Average Complexity: O(n ^ 2)
Worse Complexity: O(n ^ 2)
Best Complexity: O(n)
Class : Comparison sort
Data structure: Array
Space Complexity: O(1)
*/
func InsertionSort(elements []int) []int {

	for i := 1; i < len(elements); i++ {
		currentElement := elements[i]
		j := i - 1
		// Move elements that are greater than currentElement
		// to one position ahead of their current position.
		for j >= 0 && elements[j] > currentElement {
			elements[j+1] = elements[j]
			j--
		}
		elements[j+1] = currentElement
	}

	return elements
}

// Binary Insertion Sort is an efficient variation of the insertion sort algorithm
// that uses a binary search to find the correct position to insert element into a sorted portion or slice.
func BinaryInsertionSort(elements []int) []int {

	for i := 1; i < len(elements); i++ {
		// Store element to be inserted into the sorted portion.
		key := elements[i]

		// Find the position where the key should be inserted using binary search
		left, right := 0, i-1
		for left <= right {
			mid := left + (right-left)/2 // To avoid division overflow

			if elements[mid] > key {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		// Shift the elements to the right to make space for the key.
		for j := i; j < left; j-- {
			elements[j+1] = elements[i]
		}

		// Insert the key at the correct position
		elements[left] = key
	}

	return elements
}
