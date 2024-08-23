package leetcode_50

import "math"

// URL: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func MaxProfit(prices []int) int {
	currentMaxProfit := 0
	minimumPrice := math.MaxInt

	for _, price := range prices {
		minimumPrice = Min(minimumPrice, price)
		currentMaxProfit = Max(currentMaxProfit, price-minimumPrice)
	}

	return currentMaxProfit
}
