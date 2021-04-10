package _322_coinChange

import "math"

var dict map[int]int

func coinChange(coins []int, amount int) int {
	// 递归
	//dict = make(map[int]int)
	//return dp(coins, amount, dict)

	// 动态规划
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := range dp {
		for _, coin := range coins {
			if i < coin {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func dp(coins []int, amount int, dict map[int]int) int {

	result := math.MaxInt32
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	for _, coin := range coins {
		var subProblem int
		sub, ok := dict[amount-coin]
		if ok {
			subProblem = sub
		} else {
			subProblem = dp(coins, amount-coin, dict)
			dict[amount-coin] = subProblem
		}

		if subProblem == -1 {
			continue
		}
		result = int(math.Min(float64(result), float64(1+subProblem)))
	}
	if result == math.MaxInt32 {
		return -1
	}

	return result
}
