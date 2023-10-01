package utils

// Generate random slice of integers

func RandomIntSlice(min int, max int, size int) []int {
	if min > max {
		panic("min must be less than max")
	}

	var slice []int
	for i := 0; i < size; i++ {
		slice = append(slice, RandomInt(min, max))
	}

	return slice
}
