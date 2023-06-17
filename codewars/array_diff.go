package codewars

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
