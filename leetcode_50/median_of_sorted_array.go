package leetcode_50

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	midIndex := len(nums1) / 2 // This performs integer division in Go.

	if len(nums1)%2 == 1 {
		return float64(nums1[midIndex]) // Go is very strict with types.
	}

	middleSum := nums1[midIndex-1] + nums1[midIndex]

	return float64(middleSum) / float64(2)
}
