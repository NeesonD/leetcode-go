package _026_removeDuplicates

func removeDuplicates(nums []int) int {
	var slow int
	for i, num := range nums {
		if i > 0 && nums[i] != nums[i-1] {
			slow++
			nums[slow] = num
		}
	}
	return slow + 1
}
