package _021_exchange

func exchange(nums []int) []int {
	slow := 0

	for i, num := range nums {
		if num%2 != 0 {
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
			continue
		}
	}
	return nums
}
