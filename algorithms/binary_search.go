package algorithms

// BinarySearch searches for a given element  from a collection of elements.
// Data structure: Array (sorted)
// Iterative approach: More efficient than recursive approach because of the overhead of recursive calls.
// It returns the index of the element if found, -1 otherwise.
// Time complexity: O(log n)
// Space complexity: O(1)

func BinarySearch(elements []int, targetElement int) int {
	left := 0
	right := len(elements) - 1

	for left <= right {
		mid := left + (right-left)/2 // To avoid overflow
		if elements[mid] == targetElement {
			return mid
		} else if elements[mid] > targetElement {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// BinarySearchRecursive searches for a given element in an array recursively.
// Data structure: Array (sorted)
// Recursive approach: Less efficient than iterative approach because of the overhead of recursive calls.
// It returns the index of the element if found, -1 otherwise
// Time complexity: O(log n)

func BinarySearchRecursive(elements []int, targetElement int, left int, right int) int {

	if left > right || len(elements) == 0 {
		return -1
	}

	mid := left + (right-left)/2 // To avoid overflow
	if elements[mid] == targetElement {
		return mid
	} else if elements[mid] > targetElement {
		return BinarySearchRecursive(elements, targetElement, left, mid-1)
	} else {
		return BinarySearchRecursive(elements, targetElement, mid+1, right)
	}
}
