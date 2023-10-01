package algorithms

// LinearSearch searches for a given element in an array.
// Data structure: Array (sorted or unsorted)
// It returns the index of the element if found, -1 otherwise.
// Time complexity: O(n)
// Space complexity: O(1)

func LinearSearch(elements []int, targetElement int) int {
	for index, element := range elements {
		if element == targetElement {
			return index
		}
	}

	return -1
}

// LinearSearchRecursive searches for a given element in an array recursively.
// It returns the index of the element if found, -1 otherwise.
// Time complexity: O(n)
// Auxiliary space complexity: O(1)

func LinearSearchRecursive(elements []int, targetElement int, startIndex int) int {
	if startIndex >= len(elements) {
		return -1
	}

	if elements[startIndex] == targetElement {
		return int(startIndex)
	}

	return LinearSearchRecursive(elements, targetElement, startIndex+1)
}
