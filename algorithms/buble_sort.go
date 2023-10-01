package algorithms

// Bubble sort is a simple sorting algorithm that repeatedly steps through the list,
// compares adjacent elements and swaps them if they are in the wrong order.
// The pass through the list is repeated until the list is sorted.
// Time complexity: O(n^2)
// Space complexity: O(1)

func BubbleSort(elements []int) []int {
	for i := 0; i < len(elements); i++ {
		for j := 0; j < len(elements)-i-1; j++ {
			if elements[j] > elements[j+1] {
				elements[j], elements[j+1] = elements[j+1], elements[j]
			}
		}
	}
	return elements
}
