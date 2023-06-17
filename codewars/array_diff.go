package codewars

// Question: https://www.codewars.com/kata/523f5d21c841566fde000009/train/go

func ArrayDiff(a []int, b []int) []int {

	bucket := make(map[int]int)
	results := make([]int, 0)

	// Adding all elements in bucket.
	for _, key := range b {
		bucket[key] += 1
	}

	// Looping through all elements in a.
	for _, key := range a {
		_, ok := bucket[key]
		if !ok {
			results = append(results, key)
		}
	}

	return results
}
