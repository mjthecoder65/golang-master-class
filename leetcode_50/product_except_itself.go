package leetcode_50

// Time complexity: O(N)
// Space complexity: O(N)
func ProductExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))
	result := make([]int, len(nums))

	// Initialize product with one.
	for i := 0; i < len(nums); i++ {
		leftProduct[i] = 1
		rightProduct[i] = 1
	}

	// Fill the left product array
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	// Fill the right product
	for i := len(nums) - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	// Construct the result array.
	for i := 0; i < len(nums); i++ {
		result[i] = leftProduct[i] * rightProduct[i]
	}

	return result
}

// Time complexity: O(N)
// Space Complexity: O(1)
func ProductExceptSelfSpaceOptimized(nums []int) []int {
	n := len(nums)

	// Create the result slice and initialize it with 1
	result := make([]int, n)

	// Step 1: Fill the result array with the product of elements to the left of each index
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Step 2: Update the result array with the product of elements to the right of each index
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}
