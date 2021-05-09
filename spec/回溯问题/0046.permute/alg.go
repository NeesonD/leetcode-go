package _046_permute

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	choose := make([]int, 0)
	help(nums, choose)
	return result
}

func help(nums []int, choose []int) {
	if len(nums) == len(choose) {
		tmp := make([]int, len(choose))
		copy(tmp, choose)
		result = append(result, tmp)
		return
	}
	for _, num := range nums {
		var skip bool
		for _, i := range choose {
			if num == i {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		choose = append(choose, num)
		help(nums, choose)
		choose = choose[:len(choose)-1]
	}
}
