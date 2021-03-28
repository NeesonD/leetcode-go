package main

import "math"

func main() {
	findTargetSumWays([]int{1, 5, 5, 11},12)
}

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	n := len(nums)
	C := sum / 2
	dp := make([]bool, C+1)
	dp[0] = true
	if nums[0] <= C {
		dp[nums[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := C; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[C]
}

func findMaxForm(strs []string, m int, n int) int {
	num := len(strs)
	dp := make([][][]int, len(strs)+1)
	for f, _ := range dp {
		dp[f] = make([][]int, m+1)
		for i := 0; i <= m; i++ {
			dp[f][i] = make([]int, n+1)
			for j := 0; j <= n; j++ {
				dp[f][i][j] = 0
			}
		}
	}

	for i := 1; i <= num; i++ {
		a, b := getNum(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if a <= j && b <= k {
					dp[i][j][k] = int(math.Max(float64(dp[i][j][k]), float64(1+dp[i-1][j-a][k-b])))
				}
			}
		}
	}
	return dp[num][m][n]
}

func getNum(s string) (int, int) {
	sum1 := 0
	sum2 := 0
	for _, i := range s {
		if i == '0' {
			sum1++
		}
		if i == '1' {
			sum2++
		}
	}
	return sum1, sum2
}


func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < S || (sum+S)%2 != 0 {
		return 0
	}

	C := (sum + S) / 2

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, C+1)
	}
	for i := 0; i <= len(nums); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(nums); i++ {
		for j := 0; j <= C; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][C]
}
