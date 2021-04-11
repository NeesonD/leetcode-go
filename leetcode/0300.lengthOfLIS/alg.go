package _300_lengthOfLIS

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(1+dp[j])))
			}
		}
	}
	var result int
	for _, num := range dp {
		result = int(math.Max(float64(result), float64(num)))
	}
	return result
}
