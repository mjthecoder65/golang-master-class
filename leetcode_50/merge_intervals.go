package leetcode_50

import "sort"

// URL: https://leetcode.com/problems/merge-intervals/

func Merge(intervals [][]int) [][]int {
	// Sort the intervals.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := [][]int{}

	// Merge Intervals;
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] <= end {
			end = Max(end, interval[1])
		} else {
			mergedIntervals = append(mergedIntervals, []int{start, end})
			start = interval[0]
			end = interval[1]
		}
	}

	mergedIntervals = append(mergedIntervals, []int{start, end})

	return mergedIntervals
}
