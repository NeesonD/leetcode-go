package _057_findContinuousSequence

func F() {
	findContinuousSequence(9)
}

var result [][]int

func findContinuousSequence(target int) [][]int {
	result = make([][]int, 0)
	windows := make([]int, 0)
	for i := 1; i < target; i++ {
		windows = append(windows, i)
		for check(windows, target) {
			windows = windows[1:]
		}
	}
	return result
}

func check(windows []int, target int) bool {
	sum := 0
	for _, v := range windows {
		sum += v
	}
	if sum < target {
		return false
	} else if sum == target {
		result = append(result, windows[:])
	}
	return true
}

func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		}
	}
	return nil
}
