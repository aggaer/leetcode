package dynamicprogramming

import "utils"

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	maxProfile := 0
	for i := 1; i < len(prices); i++ {
		min = utils.Min(min, prices[i])
		maxProfile = utils.Max(maxProfile, prices[i]-min)
	}
	return maxProfile
}
