package leetcode

import "math"

// URL: https://leetcode.com/problems/max-dot-product-of-two-subsequences/
// Title : Max Dot Product of Two Subsequences
// Description : Given two arrays nums1 and nums2.
//               Return the maximum dot product between non-empty subsequences of nums1 and nums2 with the same length.

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	dp := make([]int, len(nums1)+1)

	for i := range dp {
		dp[i] = math.MinInt32
	}

	for i := range nums2 {
		topLeft, left := 0, 0
		for j := range nums1 {
			topLeft, left = left, dp[j+1]

			if topLeft == math.MinInt32 {
				topLeft = 0
			}

			currentProduct := nums1[j] * nums2[i]
			maxProduct := max(currentProduct, currentProduct+topLeft)
			dp[j+1] = max(dp[j+1], max(dp[j], maxProduct))
		}
	}

	return dp[len(nums1)]
}
