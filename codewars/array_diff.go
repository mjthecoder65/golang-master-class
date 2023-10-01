package codewars

// Question: https://www.codewars.com/kata/523f5d21c841566fde000009/train/go

func ArrayDiff(a []int, b []int) []int {

	bucket := make(map[int]bool)
	results := []int{}
	// Adding all elements of slice b to bucket.
	// To make a lookup faster.
	for _, element := range b {
		bucket[element] = true
	}

	// Looping through all elements in a.
	// If the element is not in bucket, add it to results.
	for _, element := range a {

		if _, found := bucket[element]; !found {
			results = append(results, element)
		}
	}
	return results
}
