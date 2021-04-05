package _503_nextGreaterElements

func T() {

}

func nextGreaterElements(nums []int) []int {
	var stack []int
	result := make([]int, 2*len(nums))
	for i := 2*len(nums) - 1; i > -1; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i%len(nums)] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%len(nums)])
	}
	return result[:len(nums)]
}
