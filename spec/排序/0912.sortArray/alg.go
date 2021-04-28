package _912_sortArray

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	index := partition(nums, l, r)
	quickSort(nums, l, index-1)
	quickSort(nums, index+1, r)

}

func partition(nums []int, l, r int) int {
	p := nums[l]
	left := l
	right := r
	index := l

	for right >= left {
		for right >= left {
			if nums[right] < p {
				nums[left] = nums[right]
				index = right
				left++
				break
			}
			right--
		}

		for right >= left {
			if nums[left] > p {
				nums[right] = nums[left]
				index = left
				right--
				break
			}
			left++
		}
	}
	nums[index] = p
	return index
}
