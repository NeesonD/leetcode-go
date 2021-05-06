package _198_rob

func rob(nums []int) int {
	m := len(nums)
	n := 0
	for _, num := range nums {
		n += num
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = true
	}

	dp[1][nums[0]] = true

	for i := 2; i <= m; i++ {
		for j := nums[i-1]; j <= n; j++ {
			dp[i][j] = dp[i-1][j] || dp[i-2][j-nums[i-1]]
		}
	}
	for i := n; i >= 0; i-- {
		if dp[m][i] {
			return i
		}
	}
	return -1
}
