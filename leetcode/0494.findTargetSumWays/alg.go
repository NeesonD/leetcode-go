package _494_findTargetSumWays

import "fmt"

var dict map[string]int

func findTargetSumWays(nums []int, target int) int {
	dict = make(map[string]int)
	return dp(nums, target, 0)
}

func dp(nums []int, target int, step int) int {

	if len(nums) == step {
		if target == 0 {
			return 1
		}
		return 0
	}
	b, ok := dict[fmt.Sprintf("%d-%d", target, step)]
	if ok {
		return b
	}
	hit1 := dp(nums, target+nums[step], step+1)
	hit2 := dp(nums, target-nums[step], step+1)
	dict[fmt.Sprintf("%d-%d", target, step)] = hit1 + hit2
	return hit1 + hit2
}
