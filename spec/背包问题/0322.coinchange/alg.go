package _322_coinchange

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
