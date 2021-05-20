package _011_minArray

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := left + (right-left)/2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
