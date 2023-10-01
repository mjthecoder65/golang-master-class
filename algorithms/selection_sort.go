package algorithms

// Selection Sort is an in-place comparision sorting algorithm.
// It find a mininum element in a array and place at the begining.

/* ANALYSIS
Worse-case performance (When array is reverse sorted)
	O(n^2) comparison
	O(n) swaps
Best-case performance (When array is sorted)
	O(n^2) comparison
	O(1) swaps

Average perforamance
	O(n^2) comparison
	O(n) swaps
*/

// Find index of mininum element in the slice.
func findIndexOfMinElement(startIndex int, elements []int) int {
	indexOfMinElement := startIndex

	for j := startIndex + 1; j < len(elements); j++ {
		if elements[j] < elements[indexOfMinElement] {
			indexOfMinElement = j
		}
	}
	return indexOfMinElement
}

// Sort a slice using selection sort algorithm.
func SelectionSort(elements []int) []int {
	for i := 0; i < len(elements); i++ {
		indexOfMinElement := findIndexOfMinElement(i, elements)                             // O(n^2) comparision and n swaps
		elements[i], elements[indexOfMinElement] = elements[indexOfMinElement], elements[i] // o(n) swaps
	}
	return elements
}
