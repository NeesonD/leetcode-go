package _053_maxSubArray

import "math"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))
	}
	result := nums[0]
	for _, num := range dp {
		result = int(math.Max(float64(result), float64(num)))
	}
	return result
}
