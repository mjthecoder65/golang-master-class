package algorithms

// Merge sort is a popular divid-and-conquer sorting algorithm
// that efficiently sorts a list or array of elements.
// It works by recursively diving the input into smaller sublist,
// sort those sublists and merge them back together in a sorted order.

// Merging Two subslices using two-finger Algorithms.

func MergeSubLists(leftSublist []int, rightSublist []int) []int {
	result := []int{}
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(leftSublist) && rightIndex < len(rightSublist) {
		if leftSublist[leftIndex] <= rightSublist[rightIndex] {
			result = append(result, leftSublist[leftIndex])
			leftIndex++
		} else {
			result = append(result, rightSublist[rightIndex])
			rightIndex++
		}
	}

	if leftIndex < len(leftSublist) {
		for leftIndex < len(leftSublist) {
			result = append(result, leftSublist[leftIndex])
			leftIndex++
		}
	}

	if rightIndex < len(rightSublist) {
		for rightIndex < len(rightSublist) {
			result = append(result, rightSublist[rightIndex])
			rightIndex++
		}
	}

	return result
}

// Implementing MergeSort.
func MergeSort(elements []int) []int {
	if len(elements) <= 1 {
		return elements
	}

	mid := len(elements) / 2
	leftSublist := elements[:mid]
	rightSublist := elements[mid:]

	leftSublist = MergeSort(leftSublist)
	rightSublist = MergeSort(rightSublist)

	return MergeSubLists(leftSublist, rightSublist)
}
