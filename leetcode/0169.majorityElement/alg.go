package _169_majorityElement

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}
	var result int
	var tmpCount int
	for num, count := range dict {
		if count > tmpCount {
			result = num
			tmpCount = count
		}
	}
	return result
}
