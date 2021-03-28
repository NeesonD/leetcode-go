package _027_removeElement

func removeElement(nums []int, val int) int {
	var slow int
	for _, num := range nums {
		if num != val {
			nums[slow] = num
			slow++
		}
	}
	return slow
}
