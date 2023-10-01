package algorithms

// Quick sort is widely used comparison-based algorithm
// know for its efficiency and speed.
// It falls under the category of divide-and-conquer sorting algorithms
// The basic idea of quick sort is to partition an array into two subarrays
// one containing elements less than the chosen pivot, and the other
// container the elements greater than the pivot. Then, recursively apply the quick sort algorithm
// to the subarray.

func Partition(elements []int, low int, high int) int {
	pivot := elements[high]
	i := low - 1

	for j := low; j < high; j++ {
		if elements[j] < pivot {
			i++
			elements[i], elements[j] = elements[j], elements[i]
		}
	}

	elements[i+1], elements[high] = elements[high], elements[i+1]

	return i + 1
}

// low = 0, high = len(elements) - 1

func QuickSort(elements []int, low int, high int) {
	if low < high {
		pivotIndex := Partition(elements, low, high)
		QuickSort(elements, low, pivotIndex-1)
		QuickSort(elements, pivotIndex+1, high)
	}
}
