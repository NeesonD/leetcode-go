package _014_cuttingRope

import "math"

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i < n+1; i++ {
		for j := 2; j < i; j++ {
			dp[i] = int(math.Max(float64(dp[i]), math.Max(float64(j*(i-j)), float64(j*dp[i-j]))))
		}
	}
	return dp[n]
}
