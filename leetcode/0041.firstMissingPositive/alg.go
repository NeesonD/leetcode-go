package _041_firstMissingPositive

func firstMissingPositive(nums []int) int {
	ints := make([]int, len(nums)+1)
	for _, num := range nums {
		if num <= len(nums) && num > 0 {
			ints[num] = 1
		}
	}
	for i, i2 := range ints {
		if i > 0 && i2 == 0 {
			return i
		}
	}
	return len(nums) + 1
}
