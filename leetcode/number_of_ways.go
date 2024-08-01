package leetcode

// URL: https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/
// Title : Number of Ways to Stay in the Same Place After Some Steps

const MOD = 1000000007

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numWays(steps int, arrLen int) int {

	arrLen = findMin(arrLen, steps/2+1)

	prev := make([]int, arrLen)
	prev[0] = 1

	for step := 1; step <= steps; step++ {
		dp := make([]int, arrLen)

		for idx := arrLen - 1; idx >= 0; idx-- {
			res := prev[idx]

			if idx > 0 {
				res = (res + prev[idx-1]) % MOD
			}
			if idx < arrLen-1 {
				res = (res + prev[idx+1]) % MOD
			}

			dp[idx] = res
		}

		prev = dp
	}

	return prev[0]
}
