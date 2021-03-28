package _283_moveZeroes

func moveZeroes(nums []int) {
	var slow int
	for _, num := range nums {
		if num != 0 {
			nums[slow] = num
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
