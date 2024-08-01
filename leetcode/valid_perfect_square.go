package leetcode

// URL: https://leetcode.com/problems/valid-perfect-square/

// Time complexity: O(log n)

func isPerfectSquare(num int) bool {
	if num < 2 {
		return num == 1
	}

	left, right := 2, num/2

	for left <= right {
		mid := left + (right-left)/2
		guess := mid * mid

		if guess == num {
			return true
		} else if guess > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

// Brute force solution
// Time complexity: O(n)
func IsPerfectSquare(num int) bool {
	if num < 2 {
		return num == 1
	}

	start, end := 2, num/2

	for start <= end {
		guess := start * start
		if guess == num {
			return true
		}
		start++
	}

	return false
}
