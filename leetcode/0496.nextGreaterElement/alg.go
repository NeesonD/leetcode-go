package _496_nextGreaterElement

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	dict := make(map[int]int)
	for i := len(nums2) - 1; i > -1; i-- {
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			dict[nums2[i]] = -1
		} else {
			dict[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for i, v := range nums1 {
		nums1[i] = dict[v]
	}
	return nums1
}
